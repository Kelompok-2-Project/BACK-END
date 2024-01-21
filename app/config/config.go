package config

import (
	"log"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	// "github.com/midtrans/midtrans-go"
	"github.com/spf13/viper"
)

var (
	JWT_SECRET string
)

// type MidtransConfig struct {
// 	ApiKey string
// 	Env    midtrans.EnvironmentType
// }

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("DBUSER"); found {
		app.DB_USERNAME = val
		isRead = false
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		app.DB_PASSWORD = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DB_HOSTNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DB_PORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DB_NAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("JWTSECRET"); found {
		JWT_SECRET = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}

		JWT_SECRET = viper.GetString("JWTSECRET")
		app.DB_USERNAME = viper.Get("DBUSER").(string)
		app.DB_PASSWORD = viper.Get("DBPASS").(string)
		app.DB_HOSTNAME = viper.Get("DBHOST").(string)
		app.DB_PORT, _ = strconv.Atoi(viper.Get("DBPORT").(string))
		app.DB_NAME = viper.Get("DBNAME").(string)
	}

	return &app
}

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	cldName := viper.GetString("CLDNAME")
	cldKey := viper.GetString("CLDKEY")
	cldSecret := viper.GetString("CLDSECRET")

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}

// func (cfg *MidtransConfig) LoadFromEnv(file ...string) error {
// 	cfg.ApiKey = viper.GetString("MIDKEY")
// 	midtransEnv, _ := strconv.Atoi(viper.Get("MIDSANDBOX").(string))

// 	if midtransEnv == 0 {
// 		cfg.Env = midtrans.Production
// 	} else {
// 		cfg.Env = midtrans.Sandbox
// 	}

// 	return nil
// }

