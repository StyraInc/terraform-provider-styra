// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type SystemsV1SystemsDeleteResponse struct {
	RequestID *string `json:"request_id,omitempty"`
}

func (o *SystemsV1SystemsDeleteResponse) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}
