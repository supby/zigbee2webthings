// Code generated by gin-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/mikkeloscar/gin-swagger/api"
	"github.com/mikkeloscar/gin-swagger/tracing"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	strfmt "github.com/go-openapi/strfmt"
)

// GetThingsDeviceIDActionsActionNameActionIDEndpoint executes the core logic of the related
// route endpoint.
func GetThingsDeviceIDActionsActionNameActionIDEndpoint(handler func(ctx *gin.Context, params *GetThingsDeviceIDActionsActionNameActionIDParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span := opentracing.SpanFromContext(tracing.Context(ctx))

		// attach tags to opentracing span
		if span != nil {
			ext.HTTPMethod.Set(span, ctx.Request.Method)
			ext.HTTPUrl.Set(span, ctx.Request.URL.String())
		}

		// generate params from request
		params := NewGetThingsDeviceIDActionsActionNameActionIDParams()
		err := params.readRequest(ctx)
		if err != nil {
			errObj := err.(*errors.CompositeError)
			problem := api.Problem{
				Title:  "Unprocessable Entity.",
				Status: int(errObj.Code()),
				Detail: errObj.Error(),
			}

			// attach tags to opentracing span
			if span != nil {
				ext.HTTPStatusCode.Set(span, uint16(problem.Status))
			}

			ctx.Writer.Header().Set("Content-Type", "application/problem+json")
			ctx.JSON(problem.Status, problem)
			return
		}

		resp := handler(ctx, params)

		// attach tags to opentracing span
		if span != nil {
			ext.HTTPStatusCode.Set(span, uint16(resp.Code))
		}

		switch resp.Code {
		case http.StatusNoContent:
			ctx.AbortWithStatus(resp.Code)
		default:
			ctx.JSON(resp.Code, resp.Body)
		}
	}
}

// NewGetThingsDeviceIDActionsActionNameActionIDParams creates a new GetThingsDeviceIDActionsActionNameActionIDParams object
// with the default values initialized.
func NewGetThingsDeviceIDActionsActionNameActionIDParams() *GetThingsDeviceIDActionsActionNameActionIDParams {
	var ()
	return &GetThingsDeviceIDActionsActionNameActionIDParams{}
}

// GetThingsDeviceIDActionsActionNameActionIDParams contains all the bound params for the get things device ID actions action name action ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetThingsDeviceIDActionsActionNameActionID
type GetThingsDeviceIDActionsActionNameActionIDParams struct {

	/*Id of created action
	  Required: true
	  In: path
	*/
	ActionID string
	/*Name of action
	  Required: true
	  In: path
	*/
	ActionName string
	/*Id of device
	  Required: true
	  In: path
	*/
	DeviceID string
}

// readRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetThingsDeviceIDActionsActionNameActionIDParams) readRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	rActionID := []string{ctx.Param("actionId")}
	if err := o.bindActionID(rActionID, true, formats); err != nil {
		res = append(res, err)
	}

	rActionName := []string{ctx.Param("actionName")}
	if err := o.bindActionName(rActionName, true, formats); err != nil {
		res = append(res, err)
	}

	rDeviceID := []string{ctx.Param("deviceId")}
	if err := o.bindDeviceID(rDeviceID, true, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetThingsDeviceIDActionsActionNameActionIDParams) bindActionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ActionID = raw

	return nil
}

func (o *GetThingsDeviceIDActionsActionNameActionIDParams) bindActionName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ActionName = raw

	return nil
}

func (o *GetThingsDeviceIDActionsActionNameActionIDParams) bindDeviceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.DeviceID = raw

	return nil
}

// vim: ft=go
