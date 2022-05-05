// Code generated by goa v3.7.2, DO NOT EDIT.
//
// inventory gRPC server encoders and decoders
//
// Command:
// $ goa gen inventory-system/api/svc/design

package server

import (
	"context"
	inventorypb "inventory-system/gen/grpc/inventory/pb"
	inventory "inventory-system/gen/inventory"
	inventoryviews "inventory-system/gen/inventory/views"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeCreateResponse encodes responses from the "inventory" service "create"
// endpoint.
func EncodeCreateResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "create", "string", v)
	}
	resp := NewProtoCreateResponse(result)
	return resp, nil
}

// DecodeCreateRequest decodes requests sent to "inventory" service "create"
// endpoint.
func DecodeCreateRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *inventorypb.CreateRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventorypb.CreateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventory", "create", "*inventorypb.CreateRequest", v)
		}
	}
	var payload *inventory.CreateProductPayload
	{
		payload = NewCreatePayload(message)
	}
	return payload, nil
}

// EncodeUpdateResponse encodes responses from the "inventory" service "update"
// endpoint.
func EncodeUpdateResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "update", "string", v)
	}
	resp := NewProtoUpdateResponse(result)
	return resp, nil
}

// DecodeUpdateRequest decodes requests sent to "inventory" service "update"
// endpoint.
func DecodeUpdateRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *inventorypb.UpdateRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventorypb.UpdateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventory", "update", "*inventorypb.UpdateRequest", v)
		}
	}
	var payload *inventory.UpdateProductPayload
	{
		payload = NewUpdatePayload(message)
	}
	return payload, nil
}

// EncodeFindResponse encodes responses from the "inventory" service "find"
// endpoint.
func EncodeFindResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(*inventoryviews.Findproductresult)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "find", "*inventoryviews.Findproductresult", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoFindResponse(result)
	return resp, nil
}

// DecodeFindRequest decodes requests sent to "inventory" service "find"
// endpoint.
func DecodeFindRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *inventorypb.FindRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventorypb.FindRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventory", "find", "*inventorypb.FindRequest", v)
		}
	}
	var payload *inventory.FindPayload
	{
		payload = NewFindPayload(message)
	}
	return payload, nil
}

// EncodeDeleteResponse encodes responses from the "inventory" service "delete"
// endpoint.
func EncodeDeleteResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventory", "delete", "string", v)
	}
	resp := NewProtoDeleteResponse(result)
	return resp, nil
}

// DecodeDeleteRequest decodes requests sent to "inventory" service "delete"
// endpoint.
func DecodeDeleteRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *inventorypb.DeleteRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventorypb.DeleteRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventory", "delete", "*inventorypb.DeleteRequest", v)
		}
	}
	var payload *inventory.DeletePayload
	{
		payload = NewDeletePayload(message)
	}
	return payload, nil
}