// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/shared"
	"net/http"
)

type DeleteSecretRequest struct {
	// secret ID
	SecretID string `pathParam:"style=simple,explode=false,name=secretId"`
}

func (o *DeleteSecretRequest) GetSecretID() string {
	if o == nil {
		return ""
	}
	return o.SecretID
}

type DeleteSecretResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	SecretsV1SecretsDeleteResponse *shared.SecretsV1SecretsDeleteResponse
}

func (o *DeleteSecretResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteSecretResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteSecretResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteSecretResponse) GetSecretsV1SecretsDeleteResponse() *shared.SecretsV1SecretsDeleteResponse {
	if o == nil {
		return nil
	}
	return o.SecretsV1SecretsDeleteResponse
}
