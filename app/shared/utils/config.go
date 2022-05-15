package utils

type Config struct {
	MongodbUsername string `mapstructure:"MONGODB_USERNAME"`
	MongodbPw       string `mapstructure:"MONGODB_PW"`
	MongodbHost     string `mapstructure:"MONGODB_HOST"`
	MongodbPort     string `mapstructure:"MONGODB_PORT"`
	MongodbName     string `mapstructure:"MONGODB_NAME"`
}
