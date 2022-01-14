package main

import (
	"airdropx/config"
	"airdropx/contract_binds/airdropx"
	"airdropx/contract_binds/erc20"
	"airdropx/log"
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const oneGwei = 1e9

type TransInfo struct {
	Address common.Address
	Amount  *big.Int
}

func _main() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	logrus.Infof("config: %+v", cfg)
	log.InitLogFile(cfg.LogFilePath)

	f, err := excelize.OpenFile(cfg.ExcelFilePath)
	if err != nil {
		return err
	}
	client, err := ethclient.Dial(cfg.EthApi)
	if err != nil {
		return err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return err
	}
	callOpts := &bind.CallOpts{
		BlockNumber: nil,
		Context:     context.Background(),
	}
	privKeyBts, err := hexutil.Decode(cfg.Seed)
	if err != nil {
		return err
	}

	_, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBts)
	from := crypto.PubkeyToAddress(*pubKey.ToECDSA())
	gasPriceLimit := new(big.Int).Mul(big.NewInt(cfg.GasLimit), big.NewInt(oneGwei))
	increaseGas := big.NewInt(cfg.IncreaseGas * oneGwei)

	transacOpts := bind.TransactOpts{
		From: from,
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return signTx(tx, privKeyBts, chainId)
		},
		Context: context.Background(),
	}

	var willUseGasPrice *big.Int
	for {
		willUseGasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			return err
		}
		logrus.Info("suggest gas price: ", new(big.Int).Div(willUseGasPrice, big.NewInt(oneGwei)))
		willUseGasPrice = new(big.Int).Add(willUseGasPrice, increaseGas)
		if willUseGasPrice.Cmp(gasPriceLimit) > 0 {
			time.Sleep(time.Second * 5)
			continue
		} else {
			break
		}
	}
	transacOpts.GasPrice = willUseGasPrice
	logrus.Info("final use gas price: ", new(big.Int).Div(willUseGasPrice, big.NewInt(oneGwei)))
	airdropAddress, tx, contractAirdrop, err := contract_airdropx.DeployAirDropX(&transacOpts, client)
	if err != nil {
		return fmt.Errorf("deploy err: %s", err.Error())
	}
	logrus.Info("deploy contract tx hash: ", tx.Hash().String())
	for {
		_, pending, err := client.TransactionByHash(context.Background(), tx.Hash())
		if err == nil && !pending {
			break
		} else {
			if err != nil {
				logrus.Info(" tx status: ", err)
			} else {
				logrus.Info("tx status: is pending...")
			}
			time.Sleep(time.Second * 8)
			continue
		}
	}
	logrus.Info("deploy contract tx ok: ", tx.Hash().String())

	contractErc20, err := contract_erc20.NewErc20(common.HexToAddress(cfg.TokenContractAddress), client)
	if err != nil {
		return err
	}

	symbol, err := contractErc20.Symbol(callOpts)
	if err != nil {
		return err
	}
	logrus.Infof("token symbol: %s", symbol)
	decimals, err := contractErc20.Decimals(callOpts)
	if err != nil {
		return err
	}
	decimalsDeci := decimal.New(1, int32(decimals))
	logrus.Infof("token decimals: %d", decimals)

	rows, err := f.GetRows(cfg.SheetName)
	if err != nil {
		panic(err)
	}
	rowsLen := len(rows)
	logrus.Infof("address number: %d", rowsLen)
	transInfoChan := make(chan []TransInfo, rowsLen/int(cfg.TransNumberLimit)+1)

	totalAmount := big.NewInt(0)
	willTrans := make([]TransInfo, 0)
	for i, row := range rows {
		if uint64(i) < cfg.StartIndex {
			continue
		}
		if !common.IsHexAddress(row[0]) {
			return err
		}
		addr := common.HexToAddress(row[0])
		amountDec, err := decimal.NewFromString(row[1])
		if err != nil {
			return err
		}
		amountBig := amountDec.Mul(decimalsDeci).BigInt()
		newTransInfo := TransInfo{
			Address: addr,
			Amount:  amountBig,
		}
		willTrans = append(willTrans, newTransInfo)
		totalAmount = new(big.Int).Add(totalAmount, amountBig)
		if uint64(len(willTrans)) >= cfg.TransNumberLimit || i == rowsLen-1 {
			transInfoChan <- willTrans
			willTrans = make([]TransInfo, 0)
		}
	}
	logrus.Info("transInfoChan len: ", len(transInfoChan))

	allowance, err := contractErc20.Allowance(callOpts, from, airdropAddress)
	if err != nil {
		return err
	}

	// check approve needed
	if totalAmount.Cmp(allowance) > 0 {
		var willUseGasPrice *big.Int
		var err error
		for {
			willUseGasPrice, err = client.SuggestGasPrice(context.Background())
			if err != nil {
				return err
			}
			willUseGasPrice = new(big.Int).Add(willUseGasPrice, increaseGas)
			if willUseGasPrice.Cmp(gasPriceLimit) > 0 {
				time.Sleep(time.Second * 5)
				continue
			} else {
				break
			}
		}
		transacOpts.GasPrice = willUseGasPrice
		tx, err := contractErc20.Approve(&transacOpts, airdropAddress, totalAmount)
		if err != nil {
			return err
		}
		logrus.Info("approve tx hash: ", tx.Hash().String())
		for {
			_, pending, err := client.TransactionByHash(context.Background(), tx.Hash())
			if err == nil && !pending {
				break
			} else {
				if err != nil {
					logrus.Info("approve tx status: ", err)
				} else {
					logrus.Info("approve tx status: is pending...")
				}
				time.Sleep(time.Second * 8)
				continue
			}
		}
		logrus.Info("approve tx ok: ", tx.Hash().String())
	}

out:
	for {
		select {
		case transInfo := <-transInfoChan:
			logrus.Info("will send transInfo: ", transInfo)
			logrus.Info("will send after 6 seconds...")
			time.Sleep(time.Second * 6)

			willUseAddrList := make([]common.Address, len(transInfo))
			willUseAmountList := make([]*big.Int, len(transInfo))
			for i, t := range transInfo {
				willUseAddrList[i] = t.Address
				willUseAmountList[i] = t.Amount
			}

			var willUseGasPrice *big.Int
			var err error
			for {
				willUseGasPrice, err = client.SuggestGasPrice(context.Background())
				if err != nil {
					return err
				}
				logrus.Info("suggest gas price: ", new(big.Int).Div(willUseGasPrice, big.NewInt(oneGwei)))
				willUseGasPrice = new(big.Int).Add(willUseGasPrice, increaseGas)
				if willUseGasPrice.Cmp(gasPriceLimit) > 0 {
					time.Sleep(time.Second * 5)
					continue
				} else {
					break
				}
			}
			transacOpts.GasPrice = willUseGasPrice
			logrus.Info("final use gas price: ", new(big.Int).Div(willUseGasPrice, big.NewInt(oneGwei)))

			tx, err := contractAirdrop.HelpTransfer(
				&transacOpts,
				common.HexToAddress(cfg.TokenContractAddress),
				from,
				willUseAddrList, willUseAmountList,
			)
			if err != nil {
				return err
			}
			logrus.Info("tx hash: ", tx.Hash().String())
			for {
				_, pending, err := client.TransactionByHash(context.Background(), tx.Hash())
				if err == nil && !pending {
					break
				} else {
					if err != nil {
						logrus.Info("tx status: ", err)
					} else {
						logrus.Info("tx status: is pending...")
					}
					time.Sleep(time.Second * 8)
					continue
				}
			}
			logrus.Info("tx ok: ", tx.Hash().String())

		default:
			break out
		}
	}

	logrus.Info("all send success")
	return nil
}

func main() {
	err := _main()
	if err != nil {
		panic(err)
	}
}

func signTx(rawTx *types.Transaction, privateKeyBts []byte, chainId *big.Int) (signedTx *types.Transaction, err error) {
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyBts)
	// Sign the transaction and verify the sender to avoid hardware fault surprises
	signer := types.NewEIP155Signer(chainId)
	signedTx, err = types.SignTx(rawTx, signer, privKey.ToECDSA())
	return
}
