package parts


type Parts struct {
	left  []string
	symmetric [] string
	right []string
}

func (p *Parts) Left() []string {
	return append([]string{" "}, append(p.left, p.symmetric...)...)
}

func (p *Parts) Right() []string {
	return append([]string{" "}, append(p.right, p.symmetric...)...)
}

