# airdropx

airdop tool for erc20 token.

fast, stable and secure. 

Has been used many times in practice.

## use step

1 use main account deploy `AirDroX.sol`

2 send erc20 token to main account

3 config `conf.toml` refer to `conf.exmaple.toml`

4 config `xxx.xlsx` refer to `example.xlsx`

5 `./build/dropperd -C conf.toml`

6 just wait

## notice

deploy airdropx contract need gas: 406,894

transfer erc20 token to 100 addresses with aridropx in 1 tx need gas: 2,855,985

transfer erc20 token to 1 address directly need gas 29,935;