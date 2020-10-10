package main

import (
	"github.com/gwwfps/simple-webos-remote/config"
	"github.com/gwwfps/simple-webos-remote/tvmanager"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalln(err.Error())
	}

	tvm := tvmanager.New(cfg)
	tvm.TryConnect()
	defer tvm.Close()
	if tvm.ConnectionErr != nil {
		log.Fatalln(tvm.ConnectionErr.Error())
	}

	err = tvm.SwitchInput(os.Args[1])
	if err != nil {
		log.Fatalln(err.Error())
	}
}
