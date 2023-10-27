package config

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type InfoDatabaseSQL struct {
	Read struct {
		Hostname   string
		Name       string
		Username   string
		Password   string
		Port       string
		Parameter  string
		DriverConn string
	}
	Write struct {
		Hostname   string
		Name       string
		Username   string
		Password   string
		Port       string
		Parameter  string
		DriverConn string
	}
}

func (infoDb *InfoDatabaseSQL) getDriverConn(nameMap string) (err error) {
	viper.SetConfigFile("config.json")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = mapstructure.Decode(viper.GetStringMap(nameMap), infoDb)
	if err != nil {
		return err
	}

	infoDb.Read.DriverConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		infoDb.Read.Username, infoDb.Read.Password, infoDb.Read.Hostname, infoDb.Read.Port, infoDb.Read.Name)
	infoDb.Write.DriverConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		infoDb.Write.Username, infoDb.Write.Password, infoDb.Write.Hostname, infoDb.Write.Port, infoDb.Write.Name)

	println(infoDb.Read.DriverConn)
	println(infoDb.Write.DriverConn)

	return nil
}
