package liquidationbot

import (
	"container/heap"
	"context"
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/m/v2/config"
	"github.com/m/v2/db"
	"github.com/m/v2/db/models"
	"github.com/m/v2/watchers/abi/bet"
)

type BetCover struct {
	Bet   *models.Bet
	index int
}

type PriorityQueueShort []*BetCover

var priorityQueueShort PriorityQueueShort

func (pq PriorityQueueShort) Len() int { return len(pq) }
func (pq PriorityQueueShort) Less(i, j int) bool {
	return pq[i].Bet.LiquidationTime.Before(pq[j].Bet.LiquidationTime)
}

func (pq PriorityQueueShort) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueueShort) Push(x any) {
	n := len(*pq)
	item := x.(*BetCover)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueueShort) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueueShort) Peek() *BetCover {
	if len(*pq) == 0 {
		return nil
	}
	return (*pq)[0]
}

func InitializeLiquidationBot(wg *sync.WaitGroup) {
	botConfig := config.ReadBotConfig()

	pq, err := db.PostgresDBGlobalInstance.GetBetsSortedByPriority([]models.BetStatus{models.BetStatusPlaced})
	if err != nil {
		fmt.Printf("Cannot get bets priority queue: %s", err)
		os.Exit(1)
	}

	client, err := ethclient.Dial(botConfig.RPC)
	if err != nil {
		fmt.Printf("Cannot connect to botConfig: %s", err)
		os.Exit(1)
	}

	betContract, err := bet.NewBet(common.HexToAddress(botConfig.Contract), client)
	if err != nil {
		fmt.Printf("Cannot get contract: %s", err)
		os.Exit(1)
	}

	priorityQueueShort = make(PriorityQueueShort, len(pq))
	for index, value := range pq {
		priorityQueueShort[index] = &BetCover{
			Bet:   value,
			index: index,
		}
	}
	heap.Init(&priorityQueueShort)

	Run(betContract, client, botConfig)

	wg.Done()
}

func Run(betContract *bet.Bet, client *ethclient.Client, botConfig config.BotConfig) {
	ticker := time.NewTicker(time.Second)

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Printf("Cannot get chain id: %s", err)
		os.Exit(1)
	}

	privateKey, _ := crypto.HexToECDSA(botConfig.PrivateKeyLiquidator)
	if err != nil {
		fmt.Printf("Cannot get private key: %s", err)
		os.Exit(1)
	}

	defer ticker.Stop()

	fmt.Println("Liquidation Bot is Running")
	for {
		select {
		case <-ticker.C:
			item := priorityQueueShort.Peek()
			if item == nil {
				continue
			}

			if item.Bet.LiquidationTime.Before(time.Now()) {
				fmt.Println("Liquidation time has passed for BetID:", item.Bet.BetID)
				_, err := betContract.CloseBet(&bind.TransactOpts{
					From: common.HexToAddress(botConfig.PublicKeyLiquidator),
					Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
						return types.SignTx(tx, types.LatestSignerForChainID(chainId), privateKey)
					},
				}, big.NewInt(int64(priorityQueueShort.Peek().Bet.BetID)))
				if err != nil {
					fmt.Printf("Cannot close bet: %s", err)
					os.Exit(1)
				}
				heap.Pop(&priorityQueueShort)
			}
		}
	}
}

func AddToPriorityQueue(bet *models.Bet) {
	newBet := &BetCover{
		Bet: bet,
	}
	priorityQueueShort.Push(newBet)
}
