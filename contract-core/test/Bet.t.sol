// SPDX-License-Identifier: Unlicense
pragma solidity 0.8.17;

import "forge-std/Test.sol";

import {Bet, Escrow} from "src/Bet.sol";
import {MockPriceFeed} from "./Mock/MockPriceFeed.sol";
import {MockUSDC} from "./Mock/MockUSDC.sol";
import "forge-std/Vm.sol";

contract BetTest is Test {

    Bet public bet;
    Escrow public escrow;
    MockPriceFeed public priceFeed;
    MockUSDC public usdc;
    address public  MOCK_ADDR_1 = vm.addr(1);
    address public  MOCK_ADDR_2 = vm.addr(2);
    uint32 constant UNIT_TEST_EXPIRY_TIME = 3600;
    uint32 constant UNIT_TEST_CLOSING_TIME = 7776000;

    //--------EVENTS--------//
    event BetOpened(uint256 betIndex, address indexed userA, uint64 betAmount, bool isLong, uint32 expirationTime, uint32 closingTime);
    event BetJoined(uint256 betIndex, address indexed userB, uint256 openPrice);
    event BetClosed(uint256 betIndex, address indexed winner, uint256 closePrice);

    //--------SETUP--------//
    function setUp() public {
        priceFeed = new MockPriceFeed();
        usdc = new MockUSDC("Mock USDC", "USDC");
        bet = new Bet(address(priceFeed), address(usdc));
        escrow = bet.escrow();
        vm.prank(MOCK_ADDR_1);
        usdc.mint(1e24);
        vm.prank(MOCK_ADDR_2);
        usdc.mint(1e24);
    }

    function testOpenBet(bool long, uint64 betAmount, uint16 utExpiry, uint16 utClosing) public {
        vm.assume(betAmount>0);
        vm.assume(utClosing>=utExpiry);
        uint32 expiry = uint32(block.timestamp + utExpiry);
        uint32 closingTime = uint32(block.timestamp + utClosing);
        _openBet(long, betAmount, expiry, closingTime, MOCK_ADDR_1);
    }

    function testOpenBet_InvalidExpiryTime(bool long, uint64 betAmount, uint16 utExpiry, uint16 utClosing) public {
        vm.assume(betAmount>0);
        vm.assume(utClosing<utExpiry);
        uint32 expiry = uint32(block.timestamp + utExpiry);
        uint32 closingTime = uint32(block.timestamp + utClosing);
        _openBet(long, betAmount, expiry, closingTime, MOCK_ADDR_1);
    }

    function testJoinBet(bool long, uint64 betAmount, uint16 utExpiry, uint16 utClosing, uint16 joinTime) public {
        vm.assume(betAmount>0);
        vm.assume(utClosing>=utExpiry && joinTime<=utExpiry);
        uint32 expiry = uint32(block.timestamp + utExpiry);
        uint32 closingTime = uint32(block.timestamp + utClosing);
        uint256 betIndex = _openBet(long, betAmount, expiry, closingTime, MOCK_ADDR_1);

        // Increasing some time
        vm.warp(block.timestamp+joinTime);
        _joinBet(betIndex, betAmount, MOCK_ADDR_2);
    }

    function testJoinBetAfterExpiry(bool long, uint64 betAmount, uint16 utExpiry, uint16 utClosing, uint16 joinTime) public {
        vm.assume(betAmount>0);
        vm.assume(utClosing>=utExpiry && joinTime > utExpiry);
        uint32 expiry = uint32(block.timestamp + utExpiry);
        uint32 closingTime = uint32(block.timestamp + utClosing);
        uint256 betIndex = _openBet(long, betAmount, expiry, closingTime, MOCK_ADDR_1);

        // Increasing time to exceed expiry
        vm.warp(block.timestamp + joinTime);

        vm.startPrank(MOCK_ADDR_2);
        usdc.approve(address(escrow), betAmount);
        vm.expectRevert(); // bet expired
        bet.joinBet(betIndex);
        vm.stopPrank();
    }

    function testJoinBetAfterBetClosed() public {
        bool long = true;
        uint64 betAmount = 1e18;
        uint32 expiry = uint32(block.timestamp + UNIT_TEST_EXPIRY_TIME);
        uint32 closingTime = uint32(block.timestamp + UNIT_TEST_CLOSING_TIME);
        uint256 betIndex = _openBet(long, betAmount, expiry, closingTime, MOCK_ADDR_1);

        // Increasing some time
        vm.warp(block.timestamp + 1000);

        uint256 openingPrice = _joinBet(betIndex, betAmount, MOCK_ADDR_2);

        // Increasing time to exceed closingTime
        vm.warp(block.timestamp + UNIT_TEST_CLOSING_TIME);

        _closeBet(betIndex, openingPrice, betAmount, MOCK_ADDR_1, MOCK_ADDR_2, long);

        vm.startPrank(MOCK_ADDR_2);
        usdc.approve(address(escrow), betAmount);
        vm.expectRevert(); // bet expired
        bet.joinBet(betIndex);
        vm.stopPrank();
    }

    function testCloseBet(bool long, uint32 betAmount, uint16 utExpiry, uint16 utClosing, uint joiningTime) public {
        vm.assume(betAmount>0);
        vm.assume(utExpiry<utClosing && joiningTime<utExpiry);
        uint32 expiry = uint32(block.timestamp + utExpiry);
        uint32 closingTime = uint32(block.timestamp + utClosing);
        uint256 betIndex = _openBet(long, betAmount, expiry, closingTime, MOCK_ADDR_1);

        // Increasing some time
        vm.warp(block.timestamp + joiningTime);

        uint256 openingPrice = _joinBet(betIndex, betAmount, MOCK_ADDR_2);

        // Increasing time to exceed closingTime
        vm.warp(block.timestamp + utClosing);

        _closeBet(betIndex, openingPrice, betAmount, MOCK_ADDR_1, MOCK_ADDR_2, long);
    }

    function testCLoseBetAfterBetClosed() public {
        bool long = true;
        uint64 betAmount = 1e18;
        uint32 expiry = uint32(block.timestamp + UNIT_TEST_EXPIRY_TIME);
        uint32 closingTime = uint32(block.timestamp + UNIT_TEST_CLOSING_TIME);
        uint256 betIndex = _openBet(long, betAmount, expiry, closingTime, MOCK_ADDR_1);

        // Increasing some time
        vm.warp(block.timestamp + 1000);

        uint256 openingPrice = _joinBet(betIndex, betAmount, MOCK_ADDR_2);

        // Increasing time to exceed closingTime
        vm.warp(block.timestamp + UNIT_TEST_CLOSING_TIME);

        _closeBet(betIndex, openingPrice, betAmount, MOCK_ADDR_1, MOCK_ADDR_2, long);

        vm.startPrank(MOCK_ADDR_2);
        vm.expectRevert(); // bet not active
        bet.closeBet(betIndex);
        vm.stopPrank();
    }

     function _openBet(bool long, uint64 betAmount, uint32 expiry, uint32 closingTime, address maker) internal returns(uint256) {
         if(expiry>closingTime){
             vm.startPrank(maker);
             usdc.approve(address(escrow), betAmount);
             vm.expectRevert("Expiration time should be less than or equal to closing time");
             bet.openBet(long, betAmount, expiry, closingTime);
             vm.stopPrank();
             return 0;
         }
         uint256 length = bet.totalBets();
         uint256 userBalanceBefore = usdc.balanceOf(maker);
        uint256 expectedDecreaseInBalance = betAmount;
        vm.startPrank(maker);
        usdc.approve(address(escrow), betAmount);
        vm.expectEmit(true, true, true, true);
        emit BetOpened(length, maker, betAmount, long, expiry, closingTime);
        bet.openBet(long, betAmount, expiry, closingTime);
        vm.stopPrank();
        uint256 userBalanceAfter = usdc.balanceOf(maker);
        assertEq(
            userBalanceBefore,
            userBalanceAfter + expectedDecreaseInBalance,
            "USDC balance after closing bet is not correct"
        );
        return length;
     }

     function _joinBet(uint256 betIndex, uint64 betAmount, address taker) internal returns(uint256) {
        int256 price = 1e21;
        priceFeed.storePrice(price);
        uint256 userBalanceBefore = usdc.balanceOf(taker);
        uint256 expectedDecreaseInBalance = betAmount;
        vm.startPrank(taker);
        usdc.approve(address(escrow), betAmount);
        vm.expectEmit(true, true, true, true);
        emit BetJoined(betIndex, taker, uint256(price));
        bet.joinBet(betIndex);
        vm.stopPrank();
        uint256 userBalanceAfter = usdc.balanceOf(taker);
        assertEq(
            userBalanceBefore,
            userBalanceAfter + expectedDecreaseInBalance,
            "USDC balance after closing bet is not correct"
        );
        return uint256(price);
     }

     function _closeBet(uint256 betIndex, uint256 openingPrice, uint64 betAmount, address maker, address taker, bool isMakerLong) internal {
        uint256 price = 2e21;
        priceFeed.storePrice(int256(price));
        address winner;
        if(price > openingPrice) {
            winner = isMakerLong ? maker : taker;
        }
        else {
            winner = isMakerLong ? taker : maker;
        }
        uint256 winnerBalanceBefore = usdc.balanceOf(winner);
        uint256 expectedBalanceIncrease = betAmount * 2;
        vm.expectEmit(true, true, true, true);
        emit BetClosed(betIndex, winner, price);
        bet.closeBet(betIndex);
        uint256 winnerBalanceAfter = usdc.balanceOf(winner);
        assertEq(
            winnerBalanceAfter,
            winnerBalanceBefore + expectedBalanceIncrease,
            "USDC balance after closing bet is not correct"
        );
     }
}