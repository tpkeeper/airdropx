pragma solidity 0.6.12;

// SPDX-License-Identifier: GPL-3.0-only

import "@openzeppelin/contracts/presets/ERC20PresetMinterPauser.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract TestToken is ERC20PresetMinterPauser {
    constructor() public ERC20PresetMinterPauser("dropTest", "TEST") {}
}
