package main

import (
	"fmt"

	"github.com/Gussy97/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	cfg.SetUser("Angus")

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Printf("%v", cfg)
}
