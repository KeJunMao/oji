package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/atotto/clipboard"
	"github.com/kejunmao/oji/app"
	"github.com/kejunmao/oji/parts"
	"github.com/kejunmao/oji/qs"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
)

var (
	configPath string
	initConfigPath string
	ConfigFlag = &cli.StringFlag{
		Name:        "config",
		Aliases:     []string{"c"},
		Usage:       "使用指定的配置文件 `FILE`",
		Destination: &configPath,
	}
	InitConfig = &cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "输出默认配置文件 `FILE`",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "指定配置文件路径 `FILE`",
				Destination: &initConfigPath,
				Required:    true,
			},
		},
		Action: initConfig,
	}
)

func Oji(c *cli.Context) error {
	if configPath != "" {
		configData, err := ioutil.ReadFile(configPath)
		if err != nil {
			return err
		}
		var customParts parts.Parts
		err = json.Unmarshal(configData, &customParts)
		if err != nil {
			return err
		}
		app.SetParts(customParts)
	} else {
		app.SetParts(parts.DefaultParts)
	}

	err := survey.Ask(qs.Qs(app.Parts()), &qs.Answers)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	emoji := qs.PrintOji()
	err = survey.Ask(qs.Op, &qs.Options)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if qs.Options.Copy {
		if err := clipboard.WriteAll(emoji); err != nil {
			return err
		}
	}
	return nil
}

func initConfig(c *cli.Context) error {
	byteBuf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(byteBuf)
	encoder.SetIndent("", "    ")
	err := encoder.Encode(parts.DefaultParts)
	if err != nil {
		return err
	}
	f, err := os.Create(initConfigPath)
	if err != nil {
		return err
	}
	_, err = f.Write(byteBuf.Bytes())
	if err != nil {
		return err
	}
	return nil
}
