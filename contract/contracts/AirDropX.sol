pragma solidity 0.6.12;

// SPDX-License-Identifier: GPL-3.0-only

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract AirDropX is Ownable {
    using SafeERC20 for IERC20;

    function HelpTransfer(
        address tokenAddress,
        address from,
        address[] memory toList,
        uint256[] memory amountList
    ) external onlyOwner {
        require(toList.length == amountList.length, "param err");
        for (uint256 i = 0; i < toList.length; i++) {
            require(
                IERC20(tokenAddress).transferFrom(
                    from,
                    toList[i],
                    amountList[i]
                ),
                "trans failed"
            );
        }
    }
}
