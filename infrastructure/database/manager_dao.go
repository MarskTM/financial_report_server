package database

import (
	"fmt"

	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/system"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ManagerDBDao struct {
	postgre_sql *gorm.DB
}

func NewManagerDBDao() *ManagerDBDao {
	return &ManagerDBDao{}
}

func (m *ManagerDBDao) ConnectDB(env env.DBConfig, dbType string) error {
	var err error

	switch dbType {
	case system.PostgresDB:
		dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Ho_Chi_Minh", env.Host, env.Username, env.Password, env.Database, env.Port)
		m.postgre_sql, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	case system.MysqlDB:
		// Implement MySQL connection logic here
	}

	if err != nil {
		panic("failed to connect database")
	}

	// Apply database migration
	if env.IsMigratable {
		m.postgre_sql.AutoMigrate()
	}
	return nil
}
