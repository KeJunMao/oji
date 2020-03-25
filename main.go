package main

import (
	"github.com/kejunmao/oji/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	c := cli.NewApp()
	c.CustomAppHelpTemplate = `名称:
   {{.Name}} - {{.Usage}}
用法:
   {{.HelpName}}
选项:
   --help, h  帮助
`
	c.Name = "oji"
	c.HelpName = "oji"
	c.Usage = "(◕‿◕) 颜文字生成器"
	c.Action = cmd.Oji
	err := c.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
