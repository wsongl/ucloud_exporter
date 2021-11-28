package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"github.com/wsongl/ucloud_exporter/config"
	_ "github.com/ucloud/ucloud-sdk-go"
)

var (
	// 配置文件
	configFlag string
	// port
	portFlag int
)

func main()  {
	flag.IntVar(&portFlag, "p", 8080, "the port of service")
	flag.StringVar(&configFlag, "c", "default.yaml", "the configfile of service")
	flag.Parse()

	viper.SetConfigFile(configFlag)
	viper.SetConfigType("yaml")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Fatal error config file: configFile not found")
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	viper.Unmarshal(config.ConfigVar)



	//prometheus.MustRegister()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":" + strconv.Itoa(portFlag), nil)
}
