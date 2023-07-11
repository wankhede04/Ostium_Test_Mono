// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
import "chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";
import "./Escrow.sol";

/**
* @title Bet Contract
* @notice This contract allows users to bet on the price movement of Bitcoin
*/
contract Bet {
    struct BetStruct {
        address userA; // 160 bits
        address userB; // 160 bits
        uint256 btcOpenPrice;  // 256 bits
        uint64 betAmount;      // 64 bits
        uint32 expirationTime; // 32 bits
        uint32 closingTime;    // 32 bits
        bool isLong;           // 8 bits
        bool active;           // 8 bits
    }

    Escrow public immutable escrow;
    AggregatorV3Interface internal priceFeed; // Chainlink's price feed contract instance
    BetStruct[] public bets; // Array of all bets

    event BetOpened(uint256 betIndex, address indexed userA, uint64 betAmount, bool isLong, uint32 expirationTime, uint32 closingTime);
    event BetJoined(uint256 betIndex, address indexed userB, uint256 openPrice);
    event BetClosed(uint256 betIndex, address indexed winner, uint256 closePrice);

    constructor(address _priceFeed, address _USDC) {
        escrow = new Escrow(_USDC);
        priceFeed = AggregatorV3Interface(_priceFeed);
    }

    function openBet(bool _isLong, uint64 _betAmount, uint32 _expirationTime, uint32 _closingTime) external {
        escrow.deposit(msg.sender, _betAmount);
        bets.push(BetStruct(msg.sender, address(0), 0, _betAmount, _expirationTime, _closingTime, _isLong, false));
        emit BetOpened(bets.length - 1, msg.sender, _betAmount, _isLong, _expirationTime, _closingTime);
    }

    function joinBet(uint256 betIndex) external {
        BetStruct storage bet = bets[betIndex];
        require(!bet.active, "Bet already active");
        require(bet.userA != msg.sender, "Can't join own bet");
        require(block.timestamp <= bet.expirationTime, "Bet expired");
        escrow.deposit(msg.sender, bet.betAmount);
        (, int256 price,,,) = priceFeed.latestRoundData();
        bet.btcOpenPrice = uint256(price);
        bet.userB = msg.sender;
        bet.active = true;
        emit BetJoined(betIndex, msg.sender, bet.btcOpenPrice);
    }

    function closeBet(uint256 betIndex) external {
        BetStruct storage bet = bets[betIndex];
        require(bet.active, "Bet not active");
        require(block.timestamp >= bet.closingTime, "Bet not yet closed");
        (, int256 price,,,) = priceFeed.latestRoundData();
        uint256 closePrice = uint256(price);
        address winner = closePrice > bet.btcOpenPrice ? (bet.isLong ? bet.userA : bet.userB) : (bet.isLong ? bet.userB : bet.userA);
        escrow.release(winner,bet.betAmount*2);
        bet.active = false;
        emit BetClosed(betIndex, winner, closePrice);
    }

    function getBetInfo(uint256 betIndex) external view returns (address userA, address userB, uint256 btcOpenPrice, uint64 betAmount, uint32 expirationTime, uint32 closingTime, bool isLong, bool active) {
        require(betIndex < bets.length, "Bet does not exist");
        BetStruct storage bet = bets[betIndex];
        return (bet.userA, bet.userB, bet.btcOpenPrice, bet.betAmount, bet.expirationTime, bet.closingTime, bet.isLong, bet.active);
    }
}
