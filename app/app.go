package app

import (
	"github.com/kejunmao/oji/parts"
)

var (
	name = "oji"
	usage = "(◕‿◕) 颜文字生成器"
	p parts.Parts
)

func Name() string {
	return name
}

func Usage() string {
	return usage
}

func Parts() parts.Parts {
	return p
}

func SetParts(pa parts.Parts)  {
	p = pa
}