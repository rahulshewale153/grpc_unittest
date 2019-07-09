package connection

import (
	"context"
	"fmt"
	"grpc_unittest/configs"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
)

func init() {
	configs.Config.Read("testing")
}
func mockDBConnection() *gorm.DB {
	DB, err := gorm.Open("mysql", configs.Config.Username+":"+configs.Config.Password+"@/"+configs.Config.DatabaseName+"?charset=utf8&parseTime=True&loc=Local")
	fmt.Println("Connection Sucessfull!")
	if err != nil {
		configs.Ld.Logger(context.Background(), configs.ERROR, "Failed to connect database!", err)
	}
	return DB
}
func TestConnectionService_DBConnect(t *testing.T) {

	connection := NewDatabaseConnection()
	mockdb := mockDBConnection()
	tests := []struct {
		name string
		conn *ConnectionService
		want interface{}
		flag int
	}{
		{"Not connect DB", connection, mockdb, 1},
		{"connect DB", connection, mockdb, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("configpath", "./configs/")
			if tt.flag == 1 {
				//pass worong path
				os.Setenv("configpath", "./configs1/")
			}
			configs.Config.Read("testing")
			if got := connection.DBConnect(); got == nil {
				t.Errorf("ConnectionService.DBConnect() = %v, want %v", got, tt.want)
			}
		})
	}
}
