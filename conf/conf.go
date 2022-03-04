package conf

import (
	"Bilibili-project/model"
	"fmt"
	"github.com/spf13/viper"
)


//type sqlDb struct {
//	host string `yaml:"host"`
//	user string	`yaml:"user"`
//	password string `yaml:"password"`
//	port int `yaml:"port"`
//	database string `yaml:"database"`
//}

type sqlDb struct {
	Host string `mapstructure:"host"`
	User string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port int `mapstructure:"port"`
	Database string `mapstructure:"database"`
}

type RedisDb struct {
	addr string
	password string
	db int
}

type Config struct {
	Sql sqlDb
	Redis RedisDb
}

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil{
		panic(err)
	}
	var myConfig Config
	err = viper.Unmarshal(&myConfig)
	if err != nil{
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?" +
		"charset=utf8mb4&parseTime=True&loc=Local", myConfig.Sql.User,
		myConfig.Sql.Password, myConfig.Sql.Host, myConfig.Sql.Port, myConfig.Sql.Database)
	model.Database(dsn)
	addr := myConfig.Redis.addr
	password := myConfig.Redis.password
	db := myConfig.Redis.db
	_ = model.RedisDb(addr, password, db)
}

