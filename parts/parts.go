package parts

//type Parts map[string]*Part

type Parts struct {
	Arms        Part `json:"arms"`
	Bodies      Part `json:"bodies"`
	Cheeks      Part `json:"cheeks"`
	Eyes        Part `json:"eyes"`
	NosesMouths Part `json:"noses_mouths"`
}

var DefaultParts = Parts{
	Arms,
	Bodies,
	Cheeks,
	Eyes,
	NosesMouths,
}
