// Code generated by goa v3.7.2, DO NOT EDIT.
//
// inventory endpoints
//
// Command:
// $ goa gen inventory-system/api/svc/design

package inventory

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "inventory" service endpoints.
type Endpoints struct {
	Create goa.Endpoint
	Update goa.Endpoint
	Find   goa.Endpoint
	Delete goa.Endpoint
}

// NewEndpoints wraps the methods of the "inventory" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Create: NewCreateEndpoint(s),
		Update: NewUpdateEndpoint(s),
		Find:   NewFindEndpoint(s),
		Delete: NewDeleteEndpoint(s),
	}
}

// Use applies the given middleware to all the "inventory" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
	e.Update = m(e.Update)
	e.Find = m(e.Find)
	e.Delete = m(e.Delete)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "inventory".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateProductPayload)
		return s.Create(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "inventory".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateProductPayload)
		return s.Update(ctx, p)
	}
}

// NewFindEndpoint returns an endpoint function that calls the method "find" of
// service "inventory".
func NewFindEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*FindPayload)
		res, err := s.Find(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedFindproductresult(res, "default")
		return vres, nil
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "inventory".
func NewDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeletePayload)
		return s.Delete(ctx, p)
	}
}
