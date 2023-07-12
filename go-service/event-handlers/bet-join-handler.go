package eventhandlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/m/v2/db"
	"github.com/m/v2/db/models"
	liquidationbot "github.com/m/v2/liquidation-bot"

	//liquidationbot "github.com/m/v2/liquidation-bot"
	"github.com/m/v2/watchers/abi/bet"
)

func BetJoinHandler(vLog types.Log) {
	betContractAbi, err := abi.JSON(strings.NewReader(bet.BetABI))
	if err != nil {
		fmt.Printf("Cannot get abi of contract: %s", err)
		os.Exit(1)
	}

	betJoinedEvent := bet.BetBetJoined{}
	betContractAbi.UnpackIntoInterface(&betJoinedEvent, "BetJoined", vLog.Data)

	initBet, err := db.PostgresDBGlobalInstance.GetBetByID(betJoinedEvent.BetIndex.Uint64())

	if err != nil {
		fmt.Printf("Cannot fetch bet from db: %s", err)
		os.Exit(1)
	}

	initBet.Status = models.BetStatusPlaced

	err = db.PostgresDBGlobalInstance.UpdateBet(initBet)
	if err != nil {
		fmt.Printf("Cannot insert bet: %s", err)
		os.Exit(1)
	}

	liquidationbot.AddToPriorityQueue(initBet)
}
