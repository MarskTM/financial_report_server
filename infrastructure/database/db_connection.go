// File này sẽ tiền hành quản lý các Database được dùng trong một server.
// Hiện tại cấu chúc hệ thống chỉ sử dụng chung một server db nền chưa cần quản láy ManagerDBDAO dưới dạng một mảng các đb connect.

package database

import (
	"fmt"

	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/database/dao"
	"github.com/MarskTM/financial_report_server/infrastructure/database/do"
	"github.com/golang/glog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ManagerDAO struct {
	UserDAO dao.UserDAO
	Postgre *gorm.DB
}

func (m *ManagerDAO) ConnectDB(config env.DBConfig, dbType string) error {
	var err error

	switch dbType {
	case env.PostgresType:
		dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Ho_Chi_Minh", config.Host, config.Username, config.Password, config.Database, config.Port)

		// glog.V(1).Infof("(+) gateway::DB - Error: %s", dns)
		m.Postgre, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err != nil {
			glog.V(1).Infof("(-) ConnectDB - Error: %s", err.Error())
			return err
		}

		m.UserDAO = *dao.NewUserDAO(m.Postgre)

		// Apply database migration
		if config.IsMigratable {
			glog.V(1).Infoln("(+) Migrating database...")
			m.Postgre.AutoMigrate(
				&do.Role{},
				&do.User{},
				&do.Profile{},
				&do.UserRole{},

				&do.Document{},
				&do.Company{},
				&do.SubCompany{},
				&do.CompanyManagement{},
				&do.CompanyTiding{},
				&do.Shareholder{},

				// &do.FinancialReport{},
				// &do.BalanceSheet{},
				// &do.IncomeStatement{},
                // &do.CashFlowStatement{},
			)
			glog.V(1).Infoln("(+) Done migrating database")
		}

	case env.MysqlType:
		// Implement MySQL connection logic here
	}

	if err != nil {
		panic("failed to connect database")
	}
	return nil
}
