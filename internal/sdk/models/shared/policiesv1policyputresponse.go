// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PoliciesV1PolicyPutResponse struct {
	RequestID *string `json:"request_id,omitempty"`
}

func (o *PoliciesV1PolicyPutResponse) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}
