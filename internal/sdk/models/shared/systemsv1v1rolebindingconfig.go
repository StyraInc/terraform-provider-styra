// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SystemsV1V1RoleBindingConfig struct {
	// role binding ID
	ID string `json:"id"`
	// role name
	RoleName string `json:"role_name"`
}

func (o *SystemsV1V1RoleBindingConfig) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SystemsV1V1RoleBindingConfig) GetRoleName() string {
	if o == nil {
		return ""
	}
	return o.RoleName
}
