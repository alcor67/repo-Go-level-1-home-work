package main

import (
	"fmt"
	"os"

	//"project/configuration"
	"github.com/alcor67/repo-Go-level-1-home-work/configuration"
)

func main() {

	conf, errConf := configuration.Load()

	if errConf != nil {
		fmt.Printf("произошла ошибка: %s \n", errConf)
		os.Exit(1)
	}
	//fmt.Printf("%+v \n", conf)
	//fmt.Printf("%+T \n", conf)

	fmt.Println("===================================")
	fmt.Printf("Валидированные значения: \n")
	fmt.Printf(" valid port: %+v \n", conf.MyPort)
	fmt.Printf("  valid url: %+v \n", conf.MyUrl)
	fmt.Printf("   valid id: %+v \n", conf.Some_app_id)

}
