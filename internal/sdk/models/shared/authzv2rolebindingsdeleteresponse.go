// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type AuthzV2RoleBindingsDeleteResponse struct {
	RequestID *string `json:"request_id,omitempty"`
}

func (o *AuthzV2RoleBindingsDeleteResponse) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}