package main

import (
	"log"
	
	"net/http"
	_ "net/http/pprof"
	"time"
	"github.com/SwanHtetAungPhyo/Scache/constants"
	"github.com/SwanHtetAungPhyo/Scache/server"
	"github.com/SwanHtetAungPhyo/Scache/utils"
)


func main(){
	go func ()  {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	cacheConfig, err := server.NewCofig(
		server.WithPort(":9000"),
		server.WithCapacity(200),
		server.WithExpiration(10 * time.Minute),
	)
	if err != nil {
		panic("Cache Server Configuration Failed")
	}
	_, err = server.NewScacheServer(cacheConfig)
	if err != nil{
		 utils.LogMessage(constants.ERROR, err.Error())
	}	
	select{}
}

