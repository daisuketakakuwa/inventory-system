// Code generated by goa v3.7.2, DO NOT EDIT.
//
// inventory gRPC server
//
// Command:
// $ goa gen inventory-system/api/svc/design

package server

import (
	"context"
	inventorypb "inventory-system/gen/grpc/inventory/pb"
	inventory "inventory-system/gen/inventory"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the inventorypb.InventoryServer interface.
type Server struct {
	CreateH goagrpc.UnaryHandler
	UpdateH goagrpc.UnaryHandler
	FindH   goagrpc.UnaryHandler
	DeleteH goagrpc.UnaryHandler
	inventorypb.UnimplementedInventoryServer
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the inventory service endpoints.
func New(e *inventory.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		CreateH: NewCreateHandler(e.Create, uh),
		UpdateH: NewUpdateHandler(e.Update, uh),
		FindH:   NewFindHandler(e.Find, uh),
		DeleteH: NewDeleteHandler(e.Delete, uh),
	}
}

// NewCreateHandler creates a gRPC handler which serves the "inventory" service
// "create" endpoint.
func NewCreateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeCreateRequest, EncodeCreateResponse)
	}
	return h
}

// Create implements the "Create" method in inventorypb.InventoryServer
// interface.
func (s *Server) Create(ctx context.Context, message *inventorypb.CreateRequest) (*inventorypb.CreateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "create")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.CreateH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.CreateResponse), nil
}

// NewUpdateHandler creates a gRPC handler which serves the "inventory" service
// "update" endpoint.
func NewUpdateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeUpdateRequest, EncodeUpdateResponse)
	}
	return h
}

// Update implements the "Update" method in inventorypb.InventoryServer
// interface.
func (s *Server) Update(ctx context.Context, message *inventorypb.UpdateRequest) (*inventorypb.UpdateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "update")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.UpdateH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.UpdateResponse), nil
}

// NewFindHandler creates a gRPC handler which serves the "inventory" service
// "find" endpoint.
func NewFindHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeFindRequest, EncodeFindResponse)
	}
	return h
}

// Find implements the "Find" method in inventorypb.InventoryServer interface.
func (s *Server) Find(ctx context.Context, message *inventorypb.FindRequest) (*inventorypb.FindResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "find")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.FindH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.FindResponse), nil
}

// NewDeleteHandler creates a gRPC handler which serves the "inventory" service
// "delete" endpoint.
func NewDeleteHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDeleteRequest, EncodeDeleteResponse)
	}
	return h
}

// Delete implements the "Delete" method in inventorypb.InventoryServer
// interface.
func (s *Server) Delete(ctx context.Context, message *inventorypb.DeleteRequest) (*inventorypb.DeleteResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "delete")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.DeleteH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.DeleteResponse), nil
}