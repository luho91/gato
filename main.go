package main

import (
	"github.com/luho91/gato/internal/config"
	"fmt"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		fmt.Println("Error reading config file :(", err)
		return
	}

	fmt.Println(cfg)

	err = cfg.SetUser("user")

	if err != nil {
		fmt.Println("Error writing config :(", err)
		return
	}

	cfg, err = config.Read()

	if err != nil {
		fmt.Println("Error reading config file :(", err)
		return
	}

	fmt.Println(cfg)
	
}
