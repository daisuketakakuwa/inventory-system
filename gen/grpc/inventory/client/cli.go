// Code generated by goa v3.7.2, DO NOT EDIT.
//
// inventory gRPC client CLI support package
//
// Command:
// $ goa gen inventory-system/api/svc/design

package client

import (
	"encoding/json"
	"fmt"
	inventorypb "inventory-system/gen/grpc/inventory/pb"
	inventory "inventory-system/gen/inventory"
)

// BuildCreatePayload builds the payload for the inventory create endpoint from
// CLI flags.
func BuildCreatePayload(inventoryCreateMessage string) (*inventory.CreateProductPayload, error) {
	var err error
	var message inventorypb.CreateRequest
	{
		if inventoryCreateMessage != "" {
			err = json.Unmarshal([]byte(inventoryCreateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"productDescription\": \"Ut molestias dolor et fuga.\",\n      \"productMinStock\": 283978419,\n      \"productName\": \"Dolorem saepe occaecati rerum sequi.\"\n   }'")
			}
		}
	}
	v := &inventory.CreateProductPayload{
		ProductName: message.ProductName,
	}
	if message.ProductDescription != "" {
		v.ProductDescription = &message.ProductDescription
	}
	if message.ProductMinStock != 0 {
		v.ProductMinStock = &message.ProductMinStock
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the inventory update endpoint from
// CLI flags.
func BuildUpdatePayload(inventoryUpdateMessage string) (*inventory.UpdateProductPayload, error) {
	var err error
	var message inventorypb.UpdateRequest
	{
		if inventoryUpdateMessage != "" {
			err = json.Unmarshal([]byte(inventoryUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"productDescription\": \"Placeat labore.\",\n      \"productId\": 6325241227651595690,\n      \"productMinStock\": 39024873,\n      \"productName\": \"Enim ut nemo non debitis.\"\n   }'")
			}
		}
	}
	v := &inventory.UpdateProductPayload{
		ProductID: int(message.ProductId),
	}
	if message.ProductName != "" {
		v.ProductName = &message.ProductName
	}
	if message.ProductDescription != "" {
		v.ProductDescription = &message.ProductDescription
	}
	if message.ProductMinStock != 0 {
		v.ProductMinStock = &message.ProductMinStock
	}

	return v, nil
}

// BuildFindPayload builds the payload for the inventory find endpoint from CLI
// flags.
func BuildFindPayload(inventoryFindMessage string) (*inventory.FindPayload, error) {
	var err error
	var message inventorypb.FindRequest
	{
		if inventoryFindMessage != "" {
			err = json.Unmarshal([]byte(inventoryFindMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"productId\": 6190521017900561320\n   }'")
			}
		}
	}
	v := &inventory.FindPayload{
		ProductID: int(message.ProductId),
	}

	return v, nil
}

// BuildDeletePayload builds the payload for the inventory delete endpoint from
// CLI flags.
func BuildDeletePayload(inventoryDeleteMessage string) (*inventory.DeletePayload, error) {
	var err error
	var message inventorypb.DeleteRequest
	{
		if inventoryDeleteMessage != "" {
			err = json.Unmarshal([]byte(inventoryDeleteMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"productId\": 848980590140449328\n   }'")
			}
		}
	}
	v := &inventory.DeletePayload{
		ProductID: int(message.ProductId),
	}

	return v, nil
}
