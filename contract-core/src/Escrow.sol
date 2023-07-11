// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

import "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-contracts/contracts/access/Ownable.sol";

contract Escrow is Ownable{
    IERC20 USDC;
    mapping(address => uint256) private _deposits;

    constructor(address _USDC) {
        USDC = IERC20(_USDC);
    }

    function deposit(address depositor, uint64 amount) external onlyOwner{
        USDC.transferFrom(depositor, address(this), amount);
        _deposits[depositor] += amount;
    }

    function release(address recipient, uint128 amount) external onlyOwner{
        require(_deposits[recipient] >= amount, "Insufficient deposit");
        _deposits[recipient] -= amount;
        USDC.transfer(recipient, amount);
    }
}