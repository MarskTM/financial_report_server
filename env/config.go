package env


// -----------------------------------------------------------------
// gateway config
type GatewayConfig struct {
	Addr           string `json:"addr"`
	ReadTimeout    int    `json:"read_timeout"`
	WriteTimeout   int    `json:"write_timeout"`
	MaxHeaderBytes int    `json:"max_header_bytes"`

	DB    *DBConfig
	Redis *RedisConfig
}

// authen config
type AuthenConfig struct {
	URL string `json:"url"`

	DB    *DBConfig
	Redis *RedisConfig
}

//  document config
type DocumentConfig struct {
	URL string `json:"url"`

	DB    *DBConfig
	Redis *RedisConfig
}

// ----------------------------------------------------------------
// db config
type DBConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Database     string `json:"database"`
	IsMigratable bool   `json:"is_migratable"`
}

// redis config
type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}


