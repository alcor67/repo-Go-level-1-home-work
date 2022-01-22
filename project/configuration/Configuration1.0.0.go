package configuration

import (
	"errors"
	//"github.com/namsral/flag"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type myConfig struct {
	MyPort      string //8080
	MyUrl       string //http://jaeger:16686
	Some_app_id string //test
}

//создаем глобальные переменные для парсинга флагов

var (
	//первый параметр флаг, второй параметр значение по дефолту, третий параметр описание
	//todo библиотека flag всегда возвращает не nil а указатель на дефолтные значения
	PortFlag = flag.String("port", "", "port adress")
	UrlFlag  = flag.String("url", "", "url adress")
	IdFlag   = flag.String("id", "", "app id flag (test, debug, work)")
)

//todo все валидируются
func Load() (*myConfig, error) {

	conf := &myConfig{}

	//todo устанавливаем переменные окружения
	/*
		os.Setenv("MY_PORT", "1010")
		os.Setenv("MY_URL", "http://sentry:9000")
		os.Setenv("MY_ID", "test")
	*/
	//todo считываем переменные окружения из файла configuration.env
	//получаем путь до рабочего каталога
	path, errG := os.Getwd()
	if errG != nil {
		fmt.Printf("Не получен путь до рабочего каталога:  %s\n", errG.Error())
	}

	pathEnv := path + "\\configuration\\configuration.env"

	//считываем переменные окружения из файла configuration.env
	errL := godotenv.Load(pathEnv)

	if errL != nil {
		fmt.Printf("Ошибка открытия файла .env: %s\n", errL.Error())
	}

	//вызываем обработку флагов
	//если флаг не указан, то под этим флагом будет первоначальное значение по умолчанию - пустая строка
	//todo библиотека flag всегда возвращает не nil а указатель на дефолтные значения
	flag.Parse()

	//инициализация конфигурации значениями флагов
	conf.MyPort = *PortFlag
	conf.MyUrl = *UrlFlag
	conf.Some_app_id = *IdFlag

	//читаем и проверяем переменные окружения
	if varMyPort := os.Getenv("MY_PORT"); varMyPort != "" && conf.MyPort == "" {
		conf.MyPort = varMyPort
	}
	if varMyUrl := os.Getenv("MY_URL"); varMyUrl != "" && conf.MyUrl == "" {
		conf.MyUrl = varMyUrl
	}
	if varMyId := os.Getenv("MY_ID"); varMyId != "" && conf.Some_app_id == "" {
		conf.Some_app_id = varMyId
	}
	//устанавливаем значения флагов по умолчанию на пустые флаги: port=1, url=//http, id=work
	if conf.MyPort == "" {
		conf.MyPort = "1"
	}
	if conf.MyUrl == "" {
		conf.MyUrl = "//http"
	}
	if conf.Some_app_id == "" {
		conf.Some_app_id = "work"
	}

	fmt.Printf("Загружаемые значения: \n")
	fmt.Printf("       port: %+v \n", conf.MyPort)
	fmt.Printf("        url: %+v \n", conf.MyUrl)
	fmt.Printf("         id: %+v \n", conf.Some_app_id)

	//валидация значений флагов
	err := conf.validation()

	//обработка ошибки валидации значений флагов
	if err != nil {
		fmt.Printf("конфигурация не загружена: \n")
		return conf, err
	}
	/*
		fmt.Println("===================================")
		fmt.Printf("       valid port: %+v \n", conf.MyPort)
		fmt.Printf("        valid url: %+v \n", conf.MyUrl)
		fmt.Printf("valid Some_app_id: %+v \n", conf.Some_app_id)
	*/
	return conf, err
}

func (c *myConfig) validation() error {
	//формирование сообщения об ошибке
	errMsg := ""
	//валидация данных port
	_, errPort := strconv.ParseUint(c.MyPort, 10, 64)
	//обработка ошибки
	if errPort != nil {
		errMsg += "\n неверный формат -port\n"
	}

	//валидация данных url
	_, errUrl := url.ParseRequestURI(c.MyUrl)
	//обработка ошибки
	if errUrl != nil {
		errMsg += "\n неверный формат -url\n"
	}
	//валидация данных id
	//селекция значений флага IdFlag
	var errId bool
	switch c.Some_app_id {
	case "test":
		errId = false
	case "debug":
		errId = false
	case "work":
		errId = false
	default:
		errId = true
	}
	//обработка ошибки
	if errId {
		errMsg += "\n неверный формат -id, допустимые ключи (test, debug, work)\n"
	}
	if errPort != nil || errUrl != nil || errId {
		return errors.New(errMsg)
	}

	return nil
}
