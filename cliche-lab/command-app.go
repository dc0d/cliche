package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func cmdApp(*cli.Context) error {
	fmt.Printf("%+v\n", &conf)
	return nil
}
