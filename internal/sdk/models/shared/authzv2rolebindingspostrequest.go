// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type AuthzV2RoleBindingsPostRequest struct {
	// if present, implies updating existing rolebinding in its entirety, otherwise create new
	ID             *string               `json:"id,omitempty"`
	ResourceFilter AuthzV2ResourceFilter `json:"resource_filter"`
	// role ID e.g., SystemOwner
	RoleID string `json:"role_id"`
	// list of subjects
	Subjects []AuthzV2Subject `json:"subjects"`
}

func (o *AuthzV2RoleBindingsPostRequest) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AuthzV2RoleBindingsPostRequest) GetResourceFilter() AuthzV2ResourceFilter {
	if o == nil {
		return AuthzV2ResourceFilter{}
	}
	return o.ResourceFilter
}

func (o *AuthzV2RoleBindingsPostRequest) GetRoleID() string {
	if o == nil {
		return ""
	}
	return o.RoleID
}

func (o *AuthzV2RoleBindingsPostRequest) GetSubjects() []AuthzV2Subject {
	if o == nil {
		return []AuthzV2Subject{}
	}
	return o.Subjects
}