// Code generated by goa v3.7.2, DO NOT EDIT.
//
// inventory HTTP client CLI support package
//
// Command:
// $ goa gen inventory-system/api/svc/design

package client

import (
	"encoding/json"
	"fmt"
	inventory "inventory-system/gen/inventory"
	"strconv"
)

// BuildCreatePayload builds the payload for the inventory create endpoint from
// CLI flags.
func BuildCreatePayload(inventoryCreateBody string) (*inventory.CreateProductPayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(inventoryCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"productDescription\": \"Voluptatum maxime debitis soluta.\",\n      \"productMinStock\": 1041205383,\n      \"productName\": \"Occaecati consequuntur amet inventore.\"\n   }'")
		}
	}
	v := &inventory.CreateProductPayload{
		ProductName:        body.ProductName,
		ProductDescription: body.ProductDescription,
		ProductMinStock:    body.ProductMinStock,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the inventory update endpoint from
// CLI flags.
func BuildUpdatePayload(inventoryUpdateBody string, inventoryUpdateProductID string) (*inventory.UpdateProductPayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(inventoryUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"productDescription\": \"Itaque ut sunt.\",\n      \"productMinStock\": 136779013,\n      \"productName\": \"Sint rerum ab.\"\n   }'")
		}
	}
	var productID int
	{
		var v int64
		v, err = strconv.ParseInt(inventoryUpdateProductID, 10, strconv.IntSize)
		productID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for productID, must be INT")
		}
	}
	v := &inventory.UpdateProductPayload{
		ProductName:        body.ProductName,
		ProductDescription: body.ProductDescription,
		ProductMinStock:    body.ProductMinStock,
	}
	v.ProductID = productID

	return v, nil
}

// BuildFindPayload builds the payload for the inventory find endpoint from CLI
// flags.
func BuildFindPayload(inventoryFindProductID string) (*inventory.FindPayload, error) {
	var err error
	var productID int
	{
		var v int64
		v, err = strconv.ParseInt(inventoryFindProductID, 10, strconv.IntSize)
		productID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for productID, must be INT")
		}
	}
	v := &inventory.FindPayload{}
	v.ProductID = productID

	return v, nil
}

// BuildDeletePayload builds the payload for the inventory delete endpoint from
// CLI flags.
func BuildDeletePayload(inventoryDeleteProductID string) (*inventory.DeletePayload, error) {
	var err error
	var productID int
	{
		var v int64
		v, err = strconv.ParseInt(inventoryDeleteProductID, 10, strconv.IntSize)
		productID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for productID, must be INT")
		}
	}
	v := &inventory.DeletePayload{}
	v.ProductID = productID

	return v, nil
}
