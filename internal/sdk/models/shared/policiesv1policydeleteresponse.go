// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PoliciesV1PolicyDeleteResponse struct {
	RequestID *string `json:"request_id,omitempty"`
}

func (o *PoliciesV1PolicyDeleteResponse) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}
