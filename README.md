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

## Development Setup : Go - Service and Bet Completion Bot 
Go to : [Go Service](./go-service/README.md)
