// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/shared"
	"net/http"
)

type GetSystemRequest struct {
	// system ID
	ID string `pathParam:"style=simple,explode=false,name=system"`
}

func (o *GetSystemRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetSystemResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	SystemsV1SystemsGetResponse *shared.SystemsV1SystemsGetResponse
	// Not Found
	MetaV1ErrorResponse *shared.MetaV1ErrorResponse
}

func (o *GetSystemResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSystemResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSystemResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSystemResponse) GetSystemsV1SystemsGetResponse() *shared.SystemsV1SystemsGetResponse {
	if o == nil {
		return nil
	}
	return o.SystemsV1SystemsGetResponse
}

func (o *GetSystemResponse) GetMetaV1ErrorResponse() *shared.MetaV1ErrorResponse {
	if o == nil {
		return nil
	}
	return o.MetaV1ErrorResponse
}
