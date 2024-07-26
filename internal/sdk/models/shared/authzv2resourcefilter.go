// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type AuthzV2ResourceFilter struct {
	// resource ID e.g., system123
	ID string `json:"id"`
	// resource type e.g., system
	Kind string `json:"kind"`
}

func (o *AuthzV2ResourceFilter) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AuthzV2ResourceFilter) GetKind() string {
	if o == nil {
		return ""
	}
	return o.Kind
}