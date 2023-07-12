package eventhandlers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/m/v2/db"
	"github.com/m/v2/db/models"
	"github.com/m/v2/watchers/abi/bet"
)

func BetInitializedHandler(vLog types.Log) {
	betContractAbi, err := abi.JSON(strings.NewReader(bet.BetABI))
	if err != nil {
		fmt.Printf("Cannot get abi of contract: %s", err)
		os.Exit(1)
	}

	betOpenedEvent := bet.BetBetOpened{}
	betContractAbi.UnpackIntoInterface(&betOpenedEvent, "BetOpened", vLog.Data)

	initBet := &models.Bet{
		BetID:           betOpenedEvent.BetIndex.Uint64(),
		Trader:          string(vLog.Topics[1].Hex()),
		Amount:          betOpenedEvent.BetAmount,
		IsShort:         !betOpenedEvent.IsLong,
		Deadline:        time.Unix(int64(betOpenedEvent.ExpirationTime), 0),
		LiquidationTime: time.Unix(int64(betOpenedEvent.ClosingTime), 0),
	}

	err = db.PostgresDBGlobalInstance.CreateBet(initBet)
	if err != nil {
		fmt.Printf("Cannot insert bet: %s", err)
		os.Exit(1)
	}

}
