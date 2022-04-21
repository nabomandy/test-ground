package config

import (
	"encoding/json"
	"log"
	"reflect"
	"strings"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	//DEBUG_MODE
	Debug_Mode string `json:"DEBUG_MODE" env:"DEBUG_MODE" envDefault:"true"`
	//
	AppPath string `json:"APP_PATH"      env:"APP_PATH" envDefault="./app"`
	// http Port
	Port int `json:"SERVICE_PORT" env:"SERVICE_PORT" envDefault:"3000"`
	// DB 접속 정보 포맷				[dbtype]::[host]::[port]::[dbname]::[user]::[pass]
	VurixDmsDB []string `json:"VURIX_DMS_DB" env:"VURIX_DMS_DB" envSeparator:"::" envDefault="postgres::172.16.36.179::5432::gis::geoserver::geoserver`

	// 지역코드
	//AreaCode string `json:"AREA_CODE" env:"AREA_CODE" envDefault:"11410"`

	// vurix-dms-service, dms-gis-service 호스트
	//VurixDmsHostUrl string `json:"VURIX_DMS_HOST" env:"VURIX_DMS_HOST" envDefault:"http://vurix-dms-service:8080"`
	//DmsGisHostUrl   string `json:"DMS_GIS_HOST" env:"DMS_GIS_HOST" envDefault:"http://dms-gis-service:8080"`
	// 이벤트 브로커 URL
	// CommonEventBrokerURL string `json:"COMMON_EVENT_BROKER_URL"     env:"COMMON_EVENT_BROKER_URL"`

	// Seeder File Path
	SeederFilePath string `json:"SEEDER_FILE_PATH" env:"SEEDER_FILE_PATH" envDefault:"/usr/app/seeds/sql/"`

	EUC_KEY string `json:"EUC_KEY"      env:"EUC_KEY" envDefault:"54a45f13-2a4d-40fe-91a4-f8386aa13629"`
}

var instance *Config

func GetInstance() *Config {
	if instance == nil {
		instance = new(Config)
	}
	return instance
}

func (c *Config) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := env.Parse(c); err != nil {
		log.Printf("%+v\n", err)
	}

	conf, _ := json.Marshal(c)
	log.Println("Config Info : ", string(conf))
}

func GetField(c *Config, field string) string {
	r := reflect.ValueOf(c)
	f := reflect.Indirect(r).FieldByName(strings.ToUpper(field))
	return f.String()
}
