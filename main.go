package main

import (
	"errors"
	"net/http"
	"time"
	
	"github.com/sivdead/apiserver/config"
	"github.com/sivdead/apiserver/pkg/constant"
	"github.com/sivdead/apiserver/model"
	
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/sivdead/apiserver/router"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	
	pflag.Parse()
	
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	
	// Create the Gin engine.
	g := gin.New()
	
	middlewares := []gin.HandlerFunc{}
	
	// Routes.
	router.Load(
		// Cores.
		g,
		
		// middlewares.
		middlewares...,
	)
	
	// init db
	model.DB.Init()
	defer model.DB.Close()
	
	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Debug("The router has been deployed successfully.")
	}()
	
	log.Debugf("Start to listening the incoming requests on http address: %s", viper.GetString(constant.ADDR))
	log.Debugf(http.ListenAndServe(viper.GetString(constant.ADDR), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt(constant.MAX_PING_COUNT); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString(constant.URL) + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		
		// Sleep for a second to continue the next ping.
		log.Debug("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}
