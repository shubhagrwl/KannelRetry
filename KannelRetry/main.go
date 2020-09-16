package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/KannelRetry/controllers"
	"github.com/KannelRetry/dataservices"
	"github.com/KannelRetry/prometheus"
	"github.com/KannelRetry/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var serviceName = "<<=============MSG Retry Mechanism=============>>"

func main() {
	fmt.Printf("Starting %v\n", serviceName)
	initializeLogrus()
	initializeViper()
	initializeDatabase()
	prometheus.RegisterPrometheusMetrics()
	startServer("9090")
}

func initializeDatabase() {
	controllers.DataService = &dataservices.MysqlClient{}
	controllers.DataService.Connect()
}

func initializeViper() {
	viper.SetConfigName("appConfig")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't open config file.")
	}
}

func initializeLogrus() {
	Formatter := new(log.JSONFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	log.SetFormatter(Formatter)

	var filename = "/tmp/kannelretry/log/logfile_rolling.log"
	dir, errWD := os.Getwd()
	if errWD != nil {
		log.Error("Error occurred while getting present working directory, logging to stderror" + " - " + errWD.Error())
	}
	log.Info(dir)
	f, err := os.OpenFile(dir+"/"+filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Error("Error occurred while opening log file, logging to stderror")
	} else {
		multiWriter := io.MultiWriter(os.Stdout, f)
		log.SetOutput(multiWriter)
	}
}

func startServer(port string) {
	go http.ListenAndServe(":9092", router.NewHealthcheckRouter())

	r := router.NewRouter()
	http.Handle("/", promhttp.InstrumentHandlerCounter(prometheus.AllMetrics.HttpRequestsTotal, r))
	http.Handle("/metrics", promhttp.Handler())
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"X-Api-Key", "X-Client-Key", "Content-Type"},
	})
	handler := c.Handler(r)
	log.Info("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Error("An error occurred starting HTTP listener at port " + port + " error: " + err.Error())
	}
}
