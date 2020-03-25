package qs

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/kejunmao/oji/parts"
	"reflect"
)

const (
	infoColor = "\033[1;34m%s\033[0m"
	boldColor = "\033[1m%s\033[0m"
)

var (
	Qs = []*survey.Question{
		{
			Name: "armLeft",
			Prompt: &survey.Select{
				Message: "选择左胳膊:",
				Options: parts.Arms.Left(),
			},
		},
		{
			Name: "armRight",
			Prompt: &survey.Select{
				Message: "选择右胳膊:",
				Options: parts.Arms.Right(),
			},
		},
		{
			Name: "bodyLeft",
			Prompt: &survey.Select{
				Message: "选择左身体:",
				Options: parts.Bodies.Left(),
			},
		},
		{
			Name: "bodyRight",
			Prompt: &survey.Select{
				Message: "选择右身体:",
				Options: parts.Bodies.Right(),
			},
		},
		{
			Name: "cheekLeft",
			Prompt: &survey.Select{
				Message: "选择左腮红:",
				Options: parts.Cheeks.Right(),
			},
		},
		{
			Name: "cheekRight",
			Prompt: &survey.Select{
				Message: "选择右腮红:",
				Options: parts.Cheeks.Right(),
			},
		},
		{
			Name: "eyeLeft",
			Prompt: &survey.Select{
				Message: "选择左眼睛:",
				Options: parts.Eyes.Right(),
			},
		},
		{
			Name: "eyeRight",
			Prompt: &survey.Select{
				Message: "选择右眼睛:",
				Options: parts.Eyes.Right(),
			},
		},
		{
			Name: "noseMouth",
			Prompt: &survey.Select{
				Message: "选择鼻子嘴巴:",
				Options: parts.NosesMouths.Right(),
			},
		},
	}
	Op = []*survey.Question{
		{
			Name: "copy",
			Prompt: &survey.Confirm{
				Message: "复制到剪贴板:",
				Default: true,
				Help:    "",
			},
		},
	}
	Answers = struct {
		ArmLeft    string
		BodyLeft   string
		CheekLeft  string
		EyeLeft    string
		NoseMouth  string
		EyeRight   string
		CheekRight string
		BodyRight  string
		ArmRight   string
	}{}
	Options = struct {
		Copy bool
	}{}
)

func PrintOji() string {
	fmt.Printf(infoColor, "! ")
	fmt.Printf(boldColor, "你的颜文字造好啦: ")
	v := reflect.ValueOf(Answers)
	var emoji string
	count := v.NumField()
	for i := 0; i < count; i++ {
		f := v.Field(i)
		emoji += f.String()
	}
	fmt.Println(emoji)
	return emoji
}
