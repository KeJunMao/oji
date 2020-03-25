package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/atotto/clipboard"
	"github.com/kejunmao/oji/qs"
	"github.com/urfave/cli/v2"
)

func Oji(c *cli.Context) error {
	err := survey.Ask(qs.Qs, &qs.Answers)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	emoji := qs.PrintOji()
	err = survey.Ask(qs.Op,&qs.Options)
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
