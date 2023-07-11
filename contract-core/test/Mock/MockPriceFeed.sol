// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract MockPriceFeed is AggregatorV3Interface {
    uint8 public constant DECIMALS = 18;
    string public constant DESCRIPTION = "Mock Price Feed";
    uint256 public constant VERSION = 1;

    int256[] public prices;

    // Events
    event PriceStored(uint80 roundId, int256 price);

  function storePrice(int256 price) external returns(uint256 roundId) {
    roundId = prices.length;
    prices.push(price);
    emit PriceStored(uint80(roundId), price);
  }

  function decimals() external pure returns (uint8) {
    return DECIMALS;
  }

  function description() external pure returns (string memory) {
    return DESCRIPTION;
  }

  function version() external pure returns (uint256) {
    return VERSION;
  }

  function getRoundData(
    uint80 _roundId
  ) external view returns (uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound) {
    return(_roundId, prices[_roundId], 0, 0, 0);
  }

  function latestRoundData()
    external
    view
    returns (uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound) {
        uint80 length = uint80(prices.length);
        return(length, prices[length - 1], 0, 0, 0);
    }

}