// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/shared"
	"net/http"
)

type DeleteRoleBindingV2Request struct {
	// rolebinding ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteRoleBindingV2Request) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteRoleBindingV2Response struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	AuthzV2RoleBindingsDeleteResponse *shared.AuthzV2RoleBindingsDeleteResponse
	// Not Found
	MetaV1ErrorResponse *shared.MetaV1ErrorResponse
}

func (o *DeleteRoleBindingV2Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteRoleBindingV2Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteRoleBindingV2Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteRoleBindingV2Response) GetAuthzV2RoleBindingsDeleteResponse() *shared.AuthzV2RoleBindingsDeleteResponse {
	if o == nil {
		return nil
	}
	return o.AuthzV2RoleBindingsDeleteResponse
}

func (o *DeleteRoleBindingV2Response) GetMetaV1ErrorResponse() *shared.MetaV1ErrorResponse {
	if o == nil {
		return nil
	}
	return o.MetaV1ErrorResponse
}