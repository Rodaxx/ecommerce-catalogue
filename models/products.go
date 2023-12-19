package models

type ProductDatabase struct {
	Id              uint    `json:"id"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	Size            string  `json:"size"`
	Type            string  `json:"type"`
	Color           string  `json:"color"`
	Gender          string  `json:"gender"`
	ImageOne        string  `json:"imageOne"`
	ImageTwo        string  `json:"imageTwo"`
	BranchName      string  `json:"branchName"`
	BranchDirection string  `json:"branchDirection"`
	Quantity        uint    `json:"quantity"`
}
type GenericStructure struct {
	BranchName      string    `json:"branchName"`
	BranchDirection string    `json:"branchDirection"`
	Products        []Product `json:"products"`
}
type Product struct {
	Title          string           `json:"title"`
	Description    string           `json:"description"`
	Type           string           `json:"type"`
	Gender         string           `json:"gender"`
	Price          float64          `json:"price"`
	ColorImageSize []ColorImageSize `json:"colorImageSize"`
}
type ColorImageSize struct {
	Color    string      `json:"color"`
	ImageOne string      `json:"imageOne"`
	ImageTwo string      `json:"imageTwo"`
	Size     ProductSize `json:"size"`
}
type ProductSize struct {
	XS  uint `json:"XS"`
	S   uint `json:"S"`
	M   uint `json:"M"`
	L   uint `json:"L"`
	XL  uint `json:"XL"`
	XLL uint `json:"XLL"`
}