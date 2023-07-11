package watchers

import (
	"context"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/m/v2/config"
	eventhandlers "github.com/m/v2/event-handlers"
)

func InitializeWatcher(wg sync.WaitGroup) {
	config := config.ReadWatcherConfig()

	client, err := ethclient.Dial(config.RPC)
	contractAddress := common.HexToAddress(config.Contract)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(3870791),
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	WatchEvents(sub, logs)

	wg.Done()
}

func WatchEvents(sub ethereum.Subscription, logs chan types.Log) {

	BetOpenedSig := "BetOpened(uint256,address,uint64,bool,uint32,uint32)"
	BetOpenedSigHash := crypto.Keccak256Hash([]byte(BetOpenedSig))

	BetJoinedSig := "BetJoined(uint256,address,uint256)"
	BetJoinedSigHash := crypto.Keccak256Hash([]byte(BetJoinedSig))

	BetClosedSig := "BetClosed(uint256,address,uint256)"
	BetClosedSigHash := crypto.Keccak256Hash([]byte(BetClosedSig))

	for {
		println(BetOpenedSigHash.Hex())
		println(BetJoinedSigHash.Hex())
		println(BetClosedSigHash.Hex())
		select {
		case vLog := <-logs:
			println("Inside")
			switch vLog.Topics[0].Hex() {
			case BetOpenedSigHash.Hex():
				eventhandlers.BetInitializedHandler(vLog)
			case BetJoinedSigHash.Hex():
				eventhandlers.BetJoinHandler(vLog)
			case BetClosedSigHash.Hex():
				eventhandlers.BetClosedHandler(vLog)
			}
		}
	}
}
