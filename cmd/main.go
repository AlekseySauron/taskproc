package main

import (
	"context"
	"fmt"

	"github.com/AlekseySauron/taskproc/app"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("..")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Ошибка viper", err.Error())
		return
	}
	//fmt.Println(viper.GetString("sysname"))

	ctx, _ := context.WithCancel(context.Background())
	app := app.NewApplication(ctx)
	app.Run()
	defer app.Stop()

}
