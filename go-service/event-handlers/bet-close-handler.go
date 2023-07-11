package eventhandlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/m/v2/db"
	"github.com/m/v2/db/models"
	"github.com/m/v2/watchers/abi/bet"
)

func BetClosedHandler(vLog types.Log) {
	betContractAbi, err := abi.JSON(strings.NewReader(bet.BetABI))
	if err != nil {
		fmt.Println("Cannot get abi of contract: %s", err)
		os.Exit(1)
	}
	betClosedEvent := bet.BetBetClosed{}
	betContractAbi.UnpackIntoInterface(&betClosedEvent, "BetClosed", vLog.Data)

	initBet, err := db.PostgresDBGlobalInstance.GetBetByID(betClosedEvent.BetIndex.Uint64())

	if err != nil {
		fmt.Println("Cannot fetch bet from db: %s", err)
		os.Exit(1)
	}

	initBet.Status = models.BetStatusFulfilled

	err = db.PostgresDBGlobalInstance.UpdateBet(initBet)
	if err != nil {
		fmt.Println("Cannot insert bet: %s", err)
		os.Exit(1)
	}

}
