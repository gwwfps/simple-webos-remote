package main

import (
	"github.com/gwwfps/simple-webos-remote/config"
	"github.com/gwwfps/simple-webos-remote/gui"
	"github.com/gwwfps/simple-webos-remote/tvmanager"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	tvm := tvmanager.New(cfg)
	go tvm.Start()
	defer tvm.Close()

	g := gui.New(tvm, cfg)
	g.Run()
}
