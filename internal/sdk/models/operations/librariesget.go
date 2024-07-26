// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/shared"
	"net/http"
)

type LibrariesGetRequest struct {
	// id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// level of report for bundles depending on the library. One of (none, active, all). "active" is the default
	DependantBundles *string `queryParam:"style=form,explode=true,name=dependant_bundles"`
}

func (o *LibrariesGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *LibrariesGetRequest) GetDependantBundles() *string {
	if o == nil {
		return nil
	}
	return o.DependantBundles
}

type LibrariesGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	LibrariesV1LibraryResponse *shared.LibrariesV1LibraryResponse
}

func (o *LibrariesGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LibrariesGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LibrariesGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *LibrariesGetResponse) GetLibrariesV1LibraryResponse() *shared.LibrariesV1LibraryResponse {
	if o == nil {
		return nil
	}
	return o.LibrariesV1LibraryResponse
}
