package liquidationbot

import (
	"container/heap"
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
	// We want Pop to give us the earliest time, so we use Before here.
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

func InitializeLiquidationBot(wg sync.WaitGroup) {
	pq, err := db.PostgresDBGlobalInstance.GetBetsSortedByPriority([]models.BetStatus{models.BetStatusPlaced})
	if err != nil {
		fmt.Println("Cannot get bets priority queue: %s", err)
		os.Exit(1)
	}
	fmt.Println(len(pq))
	client, err := ethclient.Dial("<>")
	betContract, err := bet.NewBet(common.HexToAddress("<>"), client)

	if err != nil {
		fmt.Println("Cannot get contract")
		os.Exit(1)
	}
	priorityQueueShort = make(PriorityQueueShort, len(pq))
	fmt.Println(len(pq))
	for index, value := range pq {
		priorityQueueShort[index] = &BetCover{
			Bet:   value,
			index: index,
		}
	}
	fmt.Println(len(priorityQueueShort))
	heap.Init(&priorityQueueShort)

	Run(betContract, client)

	wg.Done()
}

func Run(betContract *bet.Bet, client *ethclient.Client) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		fmt.Println(len(priorityQueueShort))
		select {
		case <-ticker.C:
			item := priorityQueueShort.Peek()
			if item == nil {
				continue
			}
			if item.Bet.LiquidationTime.Before(time.Now()) {
				fmt.Println("Liquidation time has passed for BetID:", item.Bet.BetID)
				_, err := betContract.CloseBet(&bind.TransactOpts{
					From: common.HexToAddress("<>"),
					Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
						chainId, _ := client.NetworkID(context.Background())
						prvk, _ := crypto.HexToECDSA("<>")
						return types.SignTx(tx, types.LatestSignerForChainID(chainId), prvk)
					},
				}, big.NewInt(int64(priorityQueueShort.Peek().Bet.BetID)))
				fmt.Println(err)
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
