package parts

type Part struct {
	Left      []string  `json:"left"`
	Symmetric [] string `json:"symmetric"`
	Right     []string  `json:"right"`
}

func (p *Part) PLeft() []string {
	return append([]string{" "}, append(p.Left, p.Symmetric...)...)
}

func (p *Part) PRight() []string {
	return append([]string{" "}, append(p.Right, p.Symmetric...)...)
}
