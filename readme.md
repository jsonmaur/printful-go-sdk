# Printful Go SDK

The non-official Go SDK for the [Printful API](https://www.printful.com/docs/index). Currently in active development.

## Usage

```go
package main

import (
	"github.com/jsonmaur/printful-go-sdk/printful"
)

func main() {
	printful.Register("API_KEY")
	
	// get all products
	products, err := printful.GetProducts()
	if err != nil {
		log.Fatal(err)
	}

	// get all variants of a product
	variants, err := printful.GetVariants(products[0])
	if err != nil {
		log.Fatal(err)
	}

	// get all variants of all products
	allVariants, errs, ok := printful.GetAllVariants(products)
	if !ok {
		log.Fatal(errs)
	}
}
```