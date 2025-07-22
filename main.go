package main

import (
	fiberHandler "jastip-core/router/http"
	grpcHandler "jastip-core/router/tcp"
	"sync"

	"github.com/alfisar/jastip-import/config"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		grpcHandler.Start()
	}()

	go func() {
		defer wg.Done()
		config.Init()
		router := fiberHandler.NewRouter()

		router.Listen("0.0.0.0:8802")
	}()
	wg.Wait()

}
