// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type LibrariesDeleteRequest struct {
	// id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *LibrariesDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type LibrariesDeleteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *LibrariesDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LibrariesDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LibrariesDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
