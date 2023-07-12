# Ostium_Test_Mono

This is a solidity project that implements a betting smart contract system. The project consists of two main contracts, Bet and Escrow, and a suite of unit tests.

## Description
The project facilitates bets between two users over the outcome of a price feed. Bets are opened with a stake, an expiry, and a closing time. Once the bet is placed, another user can join the bet, and the price at the joining time is considered the opening price. When the closing time is reached, the bet is settled based on the price at that moment compared to the opening price.

## Contracts
### Bet
Bet is the main contract where all betting operations happen. It has three main functionalities:

`openBet()`: A user can open a bet by specifying whether it's a long or short position, the amount they wish to bet, the expiry time of the bet and the closing time of the bet.

`joinBet()`: Another user can join an open bet. The price at the time of joining is recorded as the opening price.

`closeBet(`): Any participant can close the bet once the closing time has passed. The bet is settled based on the closing price and the winnings are transferred to the winner.

### Escrow
`Escrow` is a secondary contract that handles the transfer and holding of funds while the bet is ongoing.

## Testing
The project includes a comprehensive suite of unit tests, with a focus on edge cases and error conditions. The tests include expected behavior as well as testing for failure cases, such as attempting to join a bet after it has expired. We also employ fuzz testing to ensure robustness of our system under a wide range of inputs.

## Development Setup : Solidity
To set up the project for development, follow these steps:

Clone the repository to your local machine.
```
cd contract core
forge build
```
if there is dependecy issue run `forge install`

For Testing : 
```
forge test
```
Install the necessary dependencies using for install.

Dependencies
Solidity v0.8.17: The project's smart contracts are written in Solidity.
Foundry
Chainlink contract 0.8
Openzepplin v 4.9.2

## Development Setup: Go - Service and Bet Completion Bot 
Go to : [Go Service](./go-service/README.md)

--
## Limitations of this betting system. 
The following points discuss the limitation of the current betting system and potential ideas and changes to the architecture that would make this betting platform more interesting, and liquid, and offer more accurate pricing.

### Limited Liquidity:
In the current design, bets can only be made between two parties and both parties need to bet the same amount. This reduces liquidity and flexibility. A pool-based approach, similar to prediction markets, could be a solution to this. Users could buy shares of long or short outcomes from a liquidity pool. The odds could be determined by an **automated market maker** (AMM) based on the ratio of long to short shares in the pool.

### Limited Asset Support: 
The system currently supports betting only in USDC and only on the price of BTC. Increasing support for other assets (tokens & commodities) would make the platform more appealing to a wider audience.

### Limited Price Feeds: 
The system relies on a single Chainlink oracle for price feeds. The inclusion of more oracles or different types of price feeds can offer more accurate pricing and reduce the risk of manipulation or errors from a single oracle.

### Fixed Closing Times: 
Currently, the bet closing time is set when the bet is made. Implementing an option for users to close their bets early, possibly at a fee, could make the platform more flexible

### Lack of Leveraged Betting: 
Allowing leveraged betting could make the platform more appealing to certain traders. This would allow users to multiply their potential winnings (and losses) by betting more than they have in their accounts.

**In order to transform this simple betting system with more sophisticated futures, there are several key features and improvements to consider:**

### Leverage:
In traditional futures trading, investors can use leverage, which means they only need to put down a fraction of the total trade value (margin) to open a position. This could be implemented in our system, allowing users to make larger bets with less capital. However, it also introduces significant risk, so a liquidation mechanism would need to be implemented.

### Liquidation Mechanism:
If a user is using leverage and their position moves against them, there needs to be a mechanism to automatically close their position before their losses exceed their margin (the amount they put up to open the leveraged position).

### Futures Contracts: 
Instead of just betting on the price direction, we can implement actual futures contracts where users agree to buy or sell a certain amount of Bitcoin at a future date. The price of these contracts would move with the underlying asset, allowing users to close their positions before the expiry date.

### Price Bands: 
To prevent market manipulation and abrupt price movements, implementing price bands that limit the maximum price movement in a given time can be helpful. If the price hits these limits, trading could be paused to allow the market to calm down.

### Cross-Margin and Isolated Margin Modes: 
In cross-margin mode, all assets in the user's account are used to cover losses. In isolated margin mode, only the margin allocated to a specific position is at risk. Offering both modes provides flexibility to traders.

### Mark Price: 
In futures trading, the **mark price** is used to calculate unrealized PNL and liquidation prices. It can prevent market manipulation by ensuring these calculations are based on the average price, not just the last traded price.
