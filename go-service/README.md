# bet-ty

### Watcher Service
Watcher service maintains a DB with bet info, along with their status, by watching events emmited by bet contract, it watchers over the following events:

`BetOpened(uint256,address,uint64,bool,uint32,uint32)`: Event emmited when a bet is opened in contract.

`BetJoined(uint256,address,uint256)`: Event emitted when a bet is joined in contract.

`BetClosed(uint256,address,uint256)`: Event emitted when a bet is closed.

### Liquidation Bot
Bot that keeps a watch over all joined bets, maintains a priority queue for them, and works as a event closing bot, and settles the bet.

## Development Setup : Go
To set up the project for development, follow these steps:

Clone the repository to your local machine.

Create a config file from template in go-service, with required rpc and credentials
```
cd go-service
go mod tidy
go run .
```
Install the necessary dependencies using for install.

Dependencies
Go 1.20
Postgresql 14
