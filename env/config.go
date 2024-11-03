package env

// -----------------------------------------------------------------
// gateway config
type GatewayConfig struct {
	Addr           string `json:"addr" toml:"addr"`
	ReadTimeout    int    `json:"read_timeout" toml:"read_timeout"`
	WriteTimeout   int    `json:"write_timeout" toml:"write_timeout"`
	MaxHeaderBytes int    `json:"max_header_bytes" toml:"max_header_bytes"`

	DBConfig DBConfig    `json:"db_config"`
	Redis    RedisConfig `json:"redis"`
}

// biz server config
type BizServerConfig struct {
	Addr string `json:"addr" toml:"addr"`

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
	Host         string `json:"host" toml:"host"`
	Port         int    `json:"port" toml:"port"`
	Username     string `json:"username" toml:"username"`
	Password     string `json:"password" toml:"password"`
	Database     string `json:"database" toml:"database"`
	IsMigratable bool   `json:"is_migratable" toml:"is_migratable"`
}

// redis config
type RedisConfig struct {
	Addr     string `json:"addr" toml:"addr"`
	Password string `json:"password" toml:"password"`
	DB       int    `json:"db" toml:"db"`
}
