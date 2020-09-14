package handlers

import (
	"github.com/basurohit77/goproduct-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	// handle the request for add products
	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		p.l.Println("Put Method", r.URL.Path)
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("Invalid URI, more than one id")
			http.Error(rw, "Invalid URI Path", http.StatusBadRequest)
			return
		}
		if len(g[0]) !=2 {
			p.l.Println("Invalid URI, more than one id")
			http.Error(rw, "Invalid URI Path", http.StatusBadRequest)
			return
		}
		idstring := g[0][1]
		id, err := strconv.Atoi(idstring)
		if err != nil {
			http.Error(rw, "can't convert string to integer", http.StatusBadRequest)
			return
		}
		p.l.Println("product id: ", id)
		p.putProducts(id, rw, r)
		return
	}
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	gp := data.GetProducts()

	// serialize the list to JSON
	err := gp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// addProducts returns the products from the data store
func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw,"Unable to Marshel", http.StatusBadRequest)
		return
	}
	if prod.Name == "" {
		http.Error(rw,"Product has name as nil", http.StatusBadRequest)
		return
	}
	p.l.Printf("Product before add: %#v", prod)
	data.AddProducts(prod)
}

// putProducts returns the products from the data store
func (p *Products) putProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw,"Unable to Marshel", http.StatusBadRequest)
	}
	p.l.Printf("Product before update: %#v", prod)
	data.UpdateProducts(id, prod)
}
