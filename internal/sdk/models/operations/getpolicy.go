// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/shared"
	"net/http"
)

type GetPolicyRequest struct {
	// policy name
	Policy string `pathParam:"style=simple,explode=false,name=policy"`
	// include dependencies
	Dependencies *bool `queryParam:"style=form,explode=true,name=dependencies"`
}

func (o *GetPolicyRequest) GetPolicy() string {
	if o == nil {
		return ""
	}
	return o.Policy
}

func (o *GetPolicyRequest) GetDependencies() *bool {
	if o == nil {
		return nil
	}
	return o.Dependencies
}

type GetPolicyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	PoliciesV1PolicyGetResponse *shared.PoliciesV1PolicyGetResponse
	// Not Found
	MetaV1ErrorResponse *shared.MetaV1ErrorResponse
}

func (o *GetPolicyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPolicyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPolicyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPolicyResponse) GetPoliciesV1PolicyGetResponse() *shared.PoliciesV1PolicyGetResponse {
	if o == nil {
		return nil
	}
	return o.PoliciesV1PolicyGetResponse
}

func (o *GetPolicyResponse) GetMetaV1ErrorResponse() *shared.MetaV1ErrorResponse {
	if o == nil {
		return nil
	}
	return o.MetaV1ErrorResponse
}
