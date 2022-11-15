// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetStatusDeviceHandlerFunc turns a function with the right signature into a get status device handler
type GetStatusDeviceHandlerFunc func(GetStatusDeviceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStatusDeviceHandlerFunc) Handle(params GetStatusDeviceParams) middleware.Responder {
	return fn(params)
}

// GetStatusDeviceHandler interface for that can handle valid get status device params
type GetStatusDeviceHandler interface {
	Handle(GetStatusDeviceParams) middleware.Responder
}

// NewGetStatusDevice creates a new http.Handler for the get status device operation
func NewGetStatusDevice(ctx *middleware.Context, handler GetStatusDeviceHandler) *GetStatusDevice {
	return &GetStatusDevice{Context: ctx, Handler: handler}
}

/* GetStatusDevice swagger:route GET /status/device getStatusDevice

Get a basic info in the device

Get a basic info (linux command "uptime, hostnamectl") in the device or system.

*/
type GetStatusDevice struct {
	Context *middleware.Context
	Handler GetStatusDeviceHandler
}

func (o *GetStatusDevice) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetStatusDeviceParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
