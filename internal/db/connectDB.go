package db

import (
	"fmt"
	"just_simple_api/internal/config"
	"just_simple_api/internal/utils"
)

func ConnectString() string {
	connDB := config.Cfg.Database
	passwordDB := utils.LoadENV("DB_PASSWORD")
	conn := fmt.Sprintf("host=%s port=%d password=%s user=%s dbname=%s sslmode=%s",
		connDB.Host,
		connDB.Port,
		passwordDB,
		connDB.User,
		connDB.DBName,
		connDB.SSlModel)

	return conn
}
