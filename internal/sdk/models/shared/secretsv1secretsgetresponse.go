// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SecretsV1SecretsGetResponse struct {
	RequestID *string          `json:"request_id,omitempty"`
	Result    *SecretsV1Secret `json:"result,omitempty"`
}

func (o *SecretsV1SecretsGetResponse) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}

func (o *SecretsV1SecretsGetResponse) GetResult() *SecretsV1Secret {
	if o == nil {
		return nil
	}
	return o.Result
}
