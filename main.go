package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hero5512/encrypt-tool/cmd/crypto"
	"github.com/urfave/cli"
)

func InitApp() error {
	app := cli.NewApp()
	app.Name = "encrypt-tool"
	app.Usage = "crypto for important data!"
	app.Version = "1.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Allen",
			Email: "lvshuaino@gmail.com",
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Printf("%s-%s", app.Name, app.Version)
		fmt.Printf("\n%s", app.Usage)
		return nil
	}
	app.Commands = crypto.CryptoCommands

	return app.Run(os.Args)
}

func main() {
	err := InitApp()
	if err != nil {
		log.Fatal(err)
	}
}
