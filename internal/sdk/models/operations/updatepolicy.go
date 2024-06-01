// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/shared"
	"net/http"
)

type UpdatePolicyRequest struct {
	// policy name
	Policy string `pathParam:"style=simple,explode=false,name=policy"`
	// etag
	IfNoneMatch                  *string                             `header:"style=simple,explode=false,name=If-None-Match"`
	PoliciesV1PoliciesPutRequest shared.PoliciesV1PoliciesPutRequest `request:"mediaType=application/json"`
}

func (o *UpdatePolicyRequest) GetPolicy() string {
	if o == nil {
		return ""
	}
	return o.Policy
}

func (o *UpdatePolicyRequest) GetIfNoneMatch() *string {
	if o == nil {
		return nil
	}
	return o.IfNoneMatch
}

func (o *UpdatePolicyRequest) GetPoliciesV1PoliciesPutRequest() shared.PoliciesV1PoliciesPutRequest {
	if o == nil {
		return shared.PoliciesV1PoliciesPutRequest{}
	}
	return o.PoliciesV1PoliciesPutRequest
}

type UpdatePolicyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	PoliciesV1PolicyPutResponse *shared.PoliciesV1PolicyPutResponse
}

func (o *UpdatePolicyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePolicyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePolicyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdatePolicyResponse) GetPoliciesV1PolicyPutResponse() *shared.PoliciesV1PolicyPutResponse {
	if o == nil {
		return nil
	}
	return o.PoliciesV1PolicyPutResponse
}
