package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

//json to struct
func (p *Product) FromJSON(r io.Reader) error  {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

//      ADD -----------------
func getNextID() int {
	prod := productList[len(productList) - 1]
	return prod.ID + 1
}
// AddProducts add a list of products
func AddProducts(p *Product)  {
	p.ID= getNextID()
	productList =append(productList, p)
}

//         Update -------------------
var ErrProductNotFound = fmt.Errorf("Product not found")
func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func UpdateProducts(id int, p *Product) error  {
	fp, pos, err := findProduct(id)
	//println(fp, pos, err)
	if fp == nil {
		return err
	}
	if err != nil {
		return err
	}
	p.ID=id
	productList[pos]=p

	return nil
}