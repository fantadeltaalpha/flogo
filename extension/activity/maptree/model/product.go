package model

type Products struct{
	Products []*Product `json:"products"`
}

type Product struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	PCO []string `json:"-"`
	Attribute map[string]string `json:"attributes,omitempty"`
	ProductComprisedOf []Product `json:"pco,omitempty"`
}

type Attribute struct{
	Value string `json:"value"`
	Language string `json:"language"`
}