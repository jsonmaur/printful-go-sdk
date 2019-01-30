package printful

import (
	"fmt"
	"sync"
)

func GetProducts() ([]Product, error) {
	body, err := Request("store/products")
	if err != nil {
		return nil, err
	}

	var parsed ResponseProduct
	if err = unmarshal(body, &parsed); err != nil {
		return nil, err
	}

	return parsed.Result, nil
}

func GetVariants(product Product) (Variants, error) {
	body, err := Request(fmt.Sprintf("store/products/%v", product.ID))
	if err != nil {
		return Variants{}, err
	}

	var parsed ResponseVariants
	if err = unmarshal(body, &parsed); err != nil {
		return Variants{}, err
	}

	return parsed.Result, nil
}

func GetAllVariants(products []Product) ([]Variants, bool, []error) {
	length := cap(products)
	allVariants := make([]Variants, length)
	errors := make([]error, length)

	wg := sync.WaitGroup{}
	wg.Add(length)

	for i, v := range products {
		go func(index int, product Product) {
			defer wg.Done()

			variants, err := GetVariants(product)

			allVariants[index] = variants
			errors[index] = err
		}(i, v)
	}

	wg.Wait()

	ok := true
	for _, v := range errors {
		if v != nil {
			ok = false
			break
		}
	}

	return allVariants, ok, errors
}
