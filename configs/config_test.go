package configs

import (
	"os"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	Config.Read("testing")
}
func TestParam_Read(t *testing.T) {
	type fields struct {
		Port           string
		SiteURL        string
		Environment    int8
		DatabaseServer string
		DatabaseName   string
		Username       string
		Password       string
		Handler        string
	}
	tests := []struct {
		name   string
		fields *fields
		flag   int
	}{
		// TODO: Add test cases.
		{"read config", &fields{}, 0},
		{"read config", &fields{}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Param{
				Port:           tt.fields.Port,
				SiteURL:        tt.fields.SiteURL,
				Environment:    tt.fields.Environment,
				DatabaseServer: tt.fields.DatabaseServer,
				DatabaseName:   tt.fields.DatabaseName,
				Username:       tt.fields.Username,
				Password:       tt.fields.Password,
				Handler:        tt.fields.Handler,
			}
			if tt.flag == 1 {
				os.Setenv("configpath", "/home/rahul/goworkspace/src/article-get-post_rahul/configs/")
			}
			c.Read("testing")
		})
	}
}
