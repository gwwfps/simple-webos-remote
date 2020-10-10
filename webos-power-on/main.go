package main

import (
	"github.com/gwwfps/simple-webos-remote/config"
	"github.com/gwwfps/simple-webos-remote/tvmanager"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalln(err.Error())
	}

	tvm := tvmanager.New(cfg)
	err = tvm.PowerOn()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
