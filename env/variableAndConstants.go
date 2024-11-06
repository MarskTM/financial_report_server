package env

import "time"

// ----------------------------------------------------------------
var (
	ExtendHour        int64 = 720
	ExtendRefreshHour int64 = 1440

	CacheExpiresAt = 3 * time.Hour //

	DefaultPassword = "phenikaa@123"
)

var (
	// DB Type
	PostgresType = "postgres"
	MysqlType    = "mysql"
)

var MapModelType = map[string]interface{}{}

var MapAssociation = map[string]map[string]interface{}{ // Alown preload association 2 level model
	"users": {
		"UserRoles":      "",
		"UserRoles.Role": "",
	},
	"roles":    {},
	"userRole": {},
	"profiles": {
		"User":                "",
		"User.UserRoles.Role": "",
		"Recruitment":         "",
		"InternShip":          "",
	},
	"recruitments": {
		"Profile":           "",
		"InternJob":         "",
		"InternJob.Company": "",
	},
	"internshipEvaluates": {},
	"internShips": {
		"Profile":            "",
		"Company":            "",
		"InternShipEvaluate": "",
	},
}
