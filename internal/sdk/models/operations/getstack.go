// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/shared"
	"net/http"
)

type GetStackRequest struct {
	// stack id
	ID string `pathParam:"style=simple,explode=false,name=stack"`
}

func (o *GetStackRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetStackResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	StacksV1StacksGetResponse *shared.StacksV1StacksGetResponse
}

func (o *GetStackResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetStackResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetStackResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetStackResponse) GetStacksV1StacksGetResponse() *shared.StacksV1StacksGetResponse {
	if o == nil {
		return nil
	}
	return o.StacksV1StacksGetResponse
}
