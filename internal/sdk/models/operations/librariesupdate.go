// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/shared"
	"net/http"
)

type LibrariesUpdateRequest struct {
	// id
	ID                              string                                 `pathParam:"style=simple,explode=false,name=id"`
	LibrariesV1CreateLibraryRequest shared.LibrariesV1CreateLibraryRequest `request:"mediaType=application/json"`
}

func (o *LibrariesUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *LibrariesUpdateRequest) GetLibrariesV1CreateLibraryRequest() shared.LibrariesV1CreateLibraryRequest {
	if o == nil {
		return shared.LibrariesV1CreateLibraryRequest{}
	}
	return o.LibrariesV1CreateLibraryRequest
}

type LibrariesUpdateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LibrariesV1LibraryResponse *shared.LibrariesV1LibraryResponse
}

func (o *LibrariesUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LibrariesUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LibrariesUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *LibrariesUpdateResponse) GetLibrariesV1LibraryResponse() *shared.LibrariesV1LibraryResponse {
	if o == nil {
		return nil
	}
	return o.LibrariesV1LibraryResponse
}
