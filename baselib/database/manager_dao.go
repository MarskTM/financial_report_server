package database

import "github.com/MarskTM/financial_report_server/baselib/database/postgre_sql"

type ManagerDBDao struct {
	PostConfig []postgre_sql.PostConfig
}