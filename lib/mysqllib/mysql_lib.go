/*
*

	@author: jixiaogang
	@since: 2024/1/5
	@desc: //TODO

*
*/
package mysqllib

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Client *gorm.DB
)

type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Config   string `json:"config"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ExampleNewClient() {
	m := Mysql{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetString("mysql.port"),
		Config:   viper.GetString("mysql.config"),
		Database: viper.GetString("mysql.database"),
		Username: viper.GetString("mysql.username"),
		Password: viper.GetString("mysql.password"),
	}

	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?" + m.Config
	Client, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
