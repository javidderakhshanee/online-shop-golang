package config

import (
	//"github.com/spf13/viper"
	"fmt"
	"os"
)

type ConfigurationsWrapper struct {
	MySql MySqlConfigurations
}

type MySqlConfigurations struct {
	ConnectionString string
}

type Configuration struct {
	ConfigurationsWrapper
}

func NewConfiguration() Configuration {
	dbhost := os.Getenv("DB_HOST")
	dbuid := os.Getenv("DB_USER")
	dbpwd := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	dbcharset := os.Getenv("DB_CHARSET")
	dbloc := os.Getenv("DB_LOC")
	mysqlConfiguration := MySqlConfigurations{
		ConnectionString: fmt.Sprintf("%s:%s@@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
			dbuid, dbpwd, dbhost, dbport, dbname, dbcharset, dbloc),
	}

	/*viper.SetConfigName("config")
	viper.AddConfigPath(".")
	*/

	/*if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}**/

	configurationWrapper := ConfigurationsWrapper{
		MySql: mysqlConfiguration,
	}

	return Configuration{configurationWrapper}
}
