// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/shared"
	"net/http"
)

type CreateRoleBindingResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	AuthzV2RoleBindingsPostResponse *shared.AuthzV2RoleBindingsPostResponse
	// Not Found
	MetaV1ErrorResponse *shared.MetaV1ErrorResponse
}

func (o *CreateRoleBindingResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateRoleBindingResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateRoleBindingResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateRoleBindingResponse) GetAuthzV2RoleBindingsPostResponse() *shared.AuthzV2RoleBindingsPostResponse {
	if o == nil {
		return nil
	}
	return o.AuthzV2RoleBindingsPostResponse
}

func (o *CreateRoleBindingResponse) GetMetaV1ErrorResponse() *shared.MetaV1ErrorResponse {
	if o == nil {
		return nil
	}
	return o.MetaV1ErrorResponse
}
