// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AuthzV2RoleBindingConfig struct {
	// rolebinding ID
	ID             string                `json:"id"`
	Metadata       MetaV1ObjectMeta      `json:"metadata"`
	ResourceFilter AuthzV2ResourceFilter `json:"resource_filter"`
	// role ID e.g., SystemOwner
	RoleID string `json:"role_id"`
	// list of subjects
	Subjects []AuthzV2Subject `json:"subjects"`
}

func (o *AuthzV2RoleBindingConfig) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AuthzV2RoleBindingConfig) GetMetadata() MetaV1ObjectMeta {
	if o == nil {
		return MetaV1ObjectMeta{}
	}
	return o.Metadata
}

func (o *AuthzV2RoleBindingConfig) GetResourceFilter() AuthzV2ResourceFilter {
	if o == nil {
		return AuthzV2ResourceFilter{}
	}
	return o.ResourceFilter
}

func (o *AuthzV2RoleBindingConfig) GetRoleID() string {
	if o == nil {
		return ""
	}
	return o.RoleID
}

func (o *AuthzV2RoleBindingConfig) GetSubjects() []AuthzV2Subject {
	if o == nil {
		return []AuthzV2Subject{}
	}
	return o.Subjects
}
