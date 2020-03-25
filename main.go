package main

import (
	"github.com/kejunmao/oji/app"
	"github.com/kejunmao/oji/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	c := cli.NewApp()
	c.Name = app.Name()
	c.HelpName = app.Name()
	c.Usage = app.Usage()
	c.Action = cmd.Oji
	c.Flags = []cli.Flag{
		cmd.ConfigFlag,
	}
	c.Commands = []*cli.Command{
		cmd.InitConfig,
	}
	err := c.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
