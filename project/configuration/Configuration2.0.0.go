package configuration

import (
	"encoding/json"
	"errors"

	//"github.com/namsral/flag"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type myConfig struct {
	MyPort      string //8080
	MyUrl       string //http://jaeger:16686
	Some_app_id string //test
}

//создаем глобальные переменные для парсинга флагов

var (
	//первый параметр флаг, второй параметр значение по дефолту, третий параметр описание
	//библиотека flag всегда возвращает не nil а указатель на дефолтные значения
	PortFlag = flag.String("port", "", "port adress")
	UrlFlag  = flag.String("url", "", "url adress")
	IdFlag   = flag.String("id", "", "app id flag (test, debug, work)")
)

type ConfigurationJson struct {
	MyID   string `json:"MyId"`
	MyPort string `json:"MyPort"`
	MyURL  string `json:"MyUrl"`
}

type ConfigurationYaml struct {
	MyPort string `yaml:"MyPort"`
	MyURL  string `yaml:"MyUrl"`
	MyID   string `yaml:"MyId"`
}

//функция загрузки флагов
func Load() (*myConfig, error) {

	conf := &myConfig{}

	//вызываем обработку флагов
	//если флаг не указан, то под этим флагом будет первоначальное значение по умолчанию - пустая строка
	// библиотека flag всегда возвращает не nil а указатель на дефолтные значения
	flag.Parse()

	//инициализация конфигурации значениями флагов
	conf.MyPort = *PortFlag
	conf.MyUrl = *UrlFlag
	conf.Some_app_id = *IdFlag

	//получаем путь до рабочего каталога
	path, errG := os.Getwd()
	if errG != nil {
		fmt.Printf("Не получен путь до рабочего каталога:  %s\n", errG.Error())
	}

	pathJson := path + "\\configuration\\configuration.json"
	//fmt.Println(pathJson)
	//fmt.Println(filepath.Ext(pathJson))

	pathYaml := path + "\\configuration\\configuration.yaml"
	//fmt.Println(pathYaml)
	//fmt.Println(filepath.Ext(pathYaml))

	//открываем файл configuration.json
	var errx int = 1
	var errFile *int = &errx
	var configJson ConfigurationJson
	fileJson, errJson := os.Open(pathJson)
	if errJson != nil {
		*errFile = 2
		fmt.Printf("Ошибка открытия файла .json: %s\n", errJson.Error())
	}

	//открываем файл configuration.yaml
	var configYaml ConfigurationYaml
	fileYaml, errYaml := os.ReadFile(pathYaml)
	if errYaml != nil {
		*errFile = 3
		fmt.Printf("Ошибка открытия файла .yaml: %s\n", errYaml.Error())
	}

	//fmt.Println("переменная error: ", errx)
	switch errx {
	case 1:

		//считывает конфигурацию из файла json и в процессе на ходу декодирует
		if err := json.NewDecoder(fileJson).Decode(&configJson); err != nil {
			fmt.Printf("Ошибка декодирования файла .json: %s\n", err.Error())
		}
		if configJson.MyPort != "" && conf.MyPort == "" {
			conf.MyPort = configJson.MyPort
		}

		if configJson.MyURL != "" && conf.MyUrl == "" {
			conf.MyUrl = configJson.MyURL
		}

		if configJson.MyID != "" && conf.Some_app_id == "" {
			conf.Some_app_id = configJson.MyID
		}

	case 2:
		//считывает конфигурацию из файла yaml
		if err := yaml.Unmarshal(fileYaml, &configYaml); err != nil {
			fmt.Printf("Ошибка декодирования файла .yaml: %s\n", err.Error())
		}
		if configYaml.MyPort != "" && conf.MyPort == "" {
			conf.MyPort = configYaml.MyPort
		}

		if configYaml.MyURL != "" && conf.MyUrl == "" {
			conf.MyUrl = configYaml.MyURL
		}

		if configYaml.MyID != "" && conf.Some_app_id == "" {
			conf.Some_app_id = configYaml.MyID
		}

	default:
		// устанавливаем переменные окружения
		/*
			os.Setenv("MY_PORT", "1010")
			os.Setenv("MY_URL", "http://sentry:9000")
			os.Setenv("MY_ID", "test")
		*/
		//считываем переменные окружения из файла configuration.env
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

		//todo =================

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
	}

	//валидация значений флагов
	err := conf.validation()

	//обработка ошибки валидации значений флагов
	if err != nil {
		fmt.Printf("конфигурация не загружена: \n")
		return conf, err
	}

	fmt.Printf("Загружаемые значения: \n")
	fmt.Printf("       port: %+v \n", conf.MyPort)
	fmt.Printf("        url: %+v \n", conf.MyUrl)
	fmt.Printf("         id: %+v \n", conf.Some_app_id)

	return conf, err

	//валидация значений флагов
	err = conf.validation()

	//обработка ошибки валидации значений флагов
	if err != nil {
		fmt.Printf("конфигурация не загружена: \n")
		return conf, err
	}

	return conf, err
}

func (c *myConfig) validation() error {
	//формирование сообщения об ошибке
	errMsg := ""
	//валидация данных port
	_, errPort := strconv.ParseUint(c.MyPort, 10, 64)
	//обработка ошибки
	if errPort != nil {
		errMsg = "\n неверный формат -port: " + string(c.MyPort) + "\n"
		//errMsg = errMsg1
	}

	//валидация данных url
	_, errUrl := url.ParseRequestURI(c.MyUrl)
	//обработка ошибки
	if errUrl != nil {
		errMsg += "\n неверный формат  -url: " + string(c.MyUrl) + "\n"
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
		errMsg += "\n неверный формат   -id: " + string(c.Some_app_id) + "\n допустимые ключи (test, debug, work)\n"
	}
	if errPort != nil || errUrl != nil || errId {
		return errors.New(errMsg)
	}

	return nil
}
