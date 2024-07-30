package main

import (
	"github.com/mukezhz/geng/cmd"
	"github.com/mukezhz/geng/pkg"
)

func main() {

	logger := pkg.GetLogger()

	// read configurations and return config struct
	config, err := pkg.NewConfig()
	if err != nil {
    logger.Fatal("configuration generation issue", "err", err)
	}

	// cobra cmd
	cmd.ExecuteWith(config)
}
