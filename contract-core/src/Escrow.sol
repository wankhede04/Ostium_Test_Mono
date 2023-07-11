// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-contracts/contracts/access/Ownable.sol";

contract Escrow is Ownable{
    IERC20 USDC;

    constructor(address _USDC) {
        USDC = IERC20(_USDC);
    }

    function deposit(address depositor, uint64 amount) external onlyOwner{
        USDC.transferFrom(depositor, address(this), amount);
    }

    function release(address recipient, uint128 amount) external onlyOwner{
        require(USDC.balanceOf(address(this)) >= amount, "Insufficient deposit");
        USDC.transfer(recipient, amount);
    }
}