package pagedata

type PageData struct {
	Title    string     `xml:"Title"`
	Body     string     `xml:"Body"`
	Products []*Product `xml:"Products>Product"`
}

type Product struct {
	Title     string `xml:"Title"`
	ImageName string `xml:"ImageName"`
	Url       string `xml:"Url"`
}
