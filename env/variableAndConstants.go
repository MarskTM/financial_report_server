package env 

// ----------------------------------------------------------------
var (
	AccessTokenTime  int64 = 24
	RefreshTokenTime int64 = 72
	DefaultPassword        = "phenikaa@123"
)



var MapModelType = map[string]interface{}{
}

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
