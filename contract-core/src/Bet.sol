// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";
import "./Escrow.sol";
import "openzeppelin-contracts/contracts/access/Ownable.sol";

/**
 * @title Bet Contract
 * @dev This contract allows users to bet on the price movement of Bitcoin.
 * Each bet consists of two parties: User A and User B.
 * The bet amount, expiration time, closing time and other bet parameters are captured in a struct.
 * The contract uses Chainlink Price Feeds for getting the BTC/USD price and ERC20 token (USDC) for betting.
 */
contract Bet is Ownable{
    struct BetStruct {
        address userA; // Address of User A - 160 bits
        address userB; // Address of User B - 160 bits
        uint256 btcOpenPrice;  // BTC price at the time of opening the bet - 356 bits
        uint64 betAmount;  // Amount of bet in USDC - 64 bits
        uint32 expirationTime;  // Expiration time of the bet offer - 32 bits
        uint32 closingTime;  // Time when the bet is supposed to close - 32 bits
        bool isLong;  // Direction of the bet. If true, User A thinks price will be higher at closing time - 8 bitd
        bool active;  // True if the bet is active, false if it is waiting for a second participant or already closed - 8 bits
    }

    // Instance of Chainlink's BTC/USD price feed contract
    AggregatorV3Interface internal priceFeed;

    // Instance of Escrow contract to handle USDC deposits and withdrawals
    Escrow public immutable escrow;

    // Array of all bets
    BetStruct[] public bets;

    uint256 public callerRewardPercentage = 1e16;  // 1% initially, represented in a fixed-point format where 1e18 is 100%

    event BetOpened(uint256 betIndex, address indexed userA, uint64 betAmount, bool isLong, uint32 expirationTime, uint32 closingTime);
    event BetJoined(uint256 betIndex, address indexed userB, uint256 openPrice);
    event BetClosed(uint256 betIndex, address indexed winner, uint256 closePrice);

    constructor(address _priceFeed, address _USDC) {
        priceFeed = AggregatorV3Interface(_priceFeed);
        escrow = new Escrow(_USDC);
    }

    function setCallerRewardPercentage(uint256 newPercentage) external onlyOwner {
        require(newPercentage <= 1e18, "Reward percentage cannot be more than 100%");
        callerRewardPercentage = newPercentage;
    }

    /**
     * @dev Opens a new bet
     * @param _isLong Direction of the bet
     * @param _betAmount Amount of bet in USDC
     * @param _expirationTime Expiration time of the bet offer
     * @param _closingTime Time when the bet is supposed to close
     */
    function openBet(bool _isLong, uint64 _betAmount, uint32 _expirationTime, uint32 _closingTime) external {
        require(_expirationTime <= _closingTime, "Expiration time should be less than or equal to closing time");
        require(_betAmount > 0, "Bet amount should be greater than zero");

        escrow.deposit(msg.sender, _betAmount);

        bets.push(BetStruct(msg.sender, address(0), 0, _betAmount, _expirationTime, _closingTime, _isLong, false));

        emit BetOpened(bets.length - 1, msg.sender, _betAmount, _isLong, _expirationTime, _closingTime);
    }

    /**
     * @dev Joins an existing bet
     * @param betIndex Index of the bet to join in the `bets` array
     */
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

    /**
     * @dev Closes a bet, sends the funds to the winner
     * @param betIndex Index of the bet to close in the `bets` array
     */
    function setCallerRewardPercentage(uint256 newPercentage) external onlyOwner {
        require(newPercentage <= 1e18, "Reward percentage cannot be more than 100%");
        callerRewardPercentage = newPercentage;
    }

    function closeBet(uint256 betIndex) external {
        BetStruct storage bet = bets[betIndex];

        require(bet.active, "Bet not active");
        require(block.timestamp >= bet.closingTime, "Bet not yet closed");

        (, int256 price,,,) = priceFeed.latestRoundData();
        uint256 closePrice = uint256(price);
        uint256 callerReward = bet.betAmount * 2 * callerRewardPercentage / 1e18;  // Calculate the caller's reward based on the percentage
        uint256 payout = bet.betAmount * 2 - callerReward;  // The remaining after subtracting the caller's reward

        if(closePrice == bet.btcOpenPrice) {
            // Draw case, return betAmount to each user
            escrow.release(bet.userA, bet.betAmount - callerReward/2); // deduct half of the caller reward from each user
            escrow.release(bet.userB, bet.betAmount - callerReward/2);
        } else {
            // Normal case, winner takes all (after subtracting the caller's reward)
            address winner = closePrice > bet.btcOpenPrice ? (bet.isLong ? bet.userA : bet.userB) : (bet.isLong ? bet.userB : bet.userA);
            escrow.release(winner, payout);
            emit BetClosed(betIndex, winner, closePrice);
        }

        // Reward the caller
        escrow.release(msg.sender, callerReward);

        bet.active = false;
    }

    /**
     * @dev Returns information about a specific bet
     * @param betIndex Index of the bet in the `bets` array
     */
    function getBetInfo(uint256 betIndex) external view returns (address userA, address userB, uint256 btcOpenPrice, uint64 betAmount, uint32 expirationTime, uint32 closingTime, bool isLong, bool active) {
        require(betIndex < bets.length, "Bet does not exist");

        BetStruct storage bet = bets[betIndex];

        return (bet.userA, bet.userB, bet.btcOpenPrice, bet.betAmount, bet.expirationTime, bet.closingTime, bet.isLong, bet.active);
    }

    /**
     * @dev Returns total number of bets
     */
    function totalBets() external view returns (uint256) {
        return bets.length;
    }
}
