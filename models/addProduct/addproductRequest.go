package models

import "encoding/xml"

// AddProductEnvelope ...
type AddProductEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Xsi     string   `xml:"xsi,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Soap12  string   `xml:"soap12,attr"`
	Body    AddProductBody
}

type AddProductBody struct {
	XMLName    xml.Name `xml:"Body"`
	Text       string   `xml:",chardata"`
	AddProduct AddProduct
}

type AddProduct struct {
	XMLName       xml.Name `xml:"AddProduct"`
	Text          string   `xml:",chardata"`
	Xmlns         string   `xml:"xmlns,attr"`
	ProductName   string   `xml:"ProductName"`
	ListPrice     float64  `xml:"ListPrice"`
	ProductNumber int      `xml:"ProductNumber"`
}

func getAddProductRequest(price float64, productname string, productnumber int) (string, error) {
	payload := AddProductEnvelope{
		Body: AddProductBody{
			AddProduct: AddProduct{
				ListPrice:     price,
				ProductName:   "Agua Jane",
				ProductNumber: productnumber,
			},
		},
	}
	stringified, err := xml.MarshalIndent(payload, "", "  ")
	if err != nil {
		return "", err
	}
	return string(stringified), nil
}
