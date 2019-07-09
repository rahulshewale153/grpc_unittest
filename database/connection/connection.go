package connection

import (
	"context"
	"fmt"
	"grpc_unittest/configs"

	"github.com/jinzhu/gorm"
)

type ConnectionService struct {
}

func NewDatabaseConnection() *ConnectionService {
	return &ConnectionService{}
}

//DbConncet Database Connection String
func (conn *ConnectionService) DBConnect() *gorm.DB {
	DB, err := gorm.Open("mysql", configs.Config.Username+":"+configs.Config.Password+"@/"+configs.Config.DatabaseName+"?charset=utf8&parseTime=True&loc=Local")
	fmt.Println("Connection Sucessfull!")
	if err != nil {
		configs.Ld.Logger(context.Background(), configs.ERROR, "Failed to connect database!", err)
	}
	return DB
}
