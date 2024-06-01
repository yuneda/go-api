package main

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBname     string
	JWTSecret  string
}

var ENVS = initConfig()

func initConfig() Confg {
	return Config{
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}