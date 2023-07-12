package db

import (
	"fmt"
	"os"
	"time"

	"github.com/m/v2/config"
	"github.com/m/v2/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDataBase struct {
	DB *gorm.DB
}

var PostgresDBGlobalInstance *PostgresDataBase

func InitializeDB() {
	dBConfig := config.ReadDBConfig()
	dbURL := "host=" + dBConfig.Host + " user=" + dBConfig.User + " password=" + dBConfig.Password + " dbname=" + dBConfig.Name + " port=" + dBConfig.Port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error unable to open DB: %s", err)
		os.Exit(1)
	}

	if err := db.AutoMigrate(&models.Bet{}); err != nil {
		fmt.Printf("Failed to automigrate tables %s", err)
		os.Exit(1)
	}

	PostgresDBGlobalInstance = &PostgresDataBase{
		DB: db,
	}
}

func (db *PostgresDataBase) CreateBet(bet *models.Bet) error {
	bet.Status = models.BetStatusInit

	if err := PostgresDBGlobalInstance.DB.Create(&bet); err.Error != nil {
		return err.Error
	}

	return nil
}

func (db *PostgresDataBase) UpdateBet(bet *models.Bet) error {

	result := PostgresDBGlobalInstance.DB.
		Model(models.Bet{}).
		Where("bet_id = ?", bet.BetID).
		Updates(bet)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *PostgresDataBase) GetBetByID(id uint64) (*models.Bet, error) {
	bet := models.Bet{}

	result := PostgresDBGlobalInstance.DB.
		Model(models.Bet{}).
		Where("bet_id = ?", id).
		Limit(1).
		Take(&bet)
	if result.Error != nil {
		return &bet, result.Error
	}

	return &bet, nil
}

func (db *PostgresDataBase) GetBetsSortedByPriority(statuses []models.BetStatus) ([]*models.Bet, error) {
	var PriorityList []*models.Bet

	result := db.DB.
		Model(&models.Bet{}).
		Where(
			"status IN (?)",
			statuses,
		).
		Order(struct {
			LiquidationTime time.Time
		}{}).
		Find(&PriorityList)
	if result.Error != nil {
		return nil, fmt.Errorf("CreatePriorityList: unable to get list from db %w", result.Error)
	}

	return PriorityList, nil
}
