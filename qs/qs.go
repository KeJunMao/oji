package qs

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/kejunmao/oji/app"
	gparts "github.com/kejunmao/oji/parts"
	"reflect"
)

const (
	infoColor = "\033[1;34m%s\033[0m"
	boldColor = "\033[1m%s\033[0m"
)

var (
	parts = app.Parts()
	qs    []*survey.Question
	Op    = []*survey.Question{
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

func Qs(p gparts.Parts) []*survey.Question {
	parts = p
	initQs()
	return qs
}

func initQs() {
	qs = []*survey.Question{
		{
			Name: "armLeft",
			Prompt: &survey.Select{
				Message: "选择左胳膊:",
				Options: parts.Arms.PLeft(),
			},
		},
		{
			Name: "armRight",
			Prompt: &survey.Select{
				Message: "选择右胳膊:",
				Options: parts.Arms.PRight(),
			},
		},
		{
			Name: "bodyLeft",
			Prompt: &survey.Select{
				Message: "选择左身体:",
				Options: parts.Bodies.PLeft(),
			},
		},
		{
			Name: "bodyRight",
			Prompt: &survey.Select{
				Message: "选择右身体:",
				Options: parts.Bodies.PRight(),
			},
		},
		{
			Name: "cheekLeft",
			Prompt: &survey.Select{
				Message: "选择左腮红:",
				Options: parts.Cheeks.PRight(),
			},
		},
		{
			Name: "cheekRight",
			Prompt: &survey.Select{
				Message: "选择右腮红:",
				Options: parts.Cheeks.PRight(),
			},
		},
		{
			Name: "eyeLeft",
			Prompt: &survey.Select{
				Message: "选择左眼睛:",
				Options: parts.Eyes.PRight(),
			},
		},
		{
			Name: "eyeRight",
			Prompt: &survey.Select{
				Message: "选择右眼睛:",
				Options: parts.Eyes.PRight(),
			},
		},
		{
			Name: "noseMouth",
			Prompt: &survey.Select{
				Message: "选择鼻子嘴巴:",
				Options: parts.NosesMouths.PRight(),
			},
		},
	}
}
