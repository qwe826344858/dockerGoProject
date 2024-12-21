package main

type EnvConfig struct {
	RedisConf RedisConfig `json:"redis_conf"`
	MysqlConf MysqlConfig `json:"mysql_conf"`
}

type RedisConfig struct {
	Host string `json:"host"`
}

type MysqlConfig struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func getEnvConfig() *EnvConfig {
	return &EnvConfig{
		RedisConf: RedisConfig{
			Host: "10.0.12.8",
		},
		MysqlConf: MysqlConfig{
			Host: "10.0.12.8",
			Username: "root",
			Password: "zoneslee",
		},
	}
}
