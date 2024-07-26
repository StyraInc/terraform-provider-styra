// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type AuthzV2RoleBindingsPostResponse struct {
	RequestID   *string                  `json:"request_id,omitempty"`
	Rolebinding AuthzV2RoleBindingConfig `json:"rolebinding"`
}

func (o *AuthzV2RoleBindingsPostResponse) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}

func (o *AuthzV2RoleBindingsPostResponse) GetRolebinding() AuthzV2RoleBindingConfig {
	if o == nil {
		return AuthzV2RoleBindingConfig{}
	}
	return o.Rolebinding
}
