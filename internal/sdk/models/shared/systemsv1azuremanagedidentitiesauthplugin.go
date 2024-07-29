// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SystemsV1AzureManagedIdentitiesAuthPlugin struct {
	// API version to use
	APIVersion string `json:"api_version"`
	// optional client ID of the managed identity you would like the token for, required if your VM has multiple user-assigned managed identities
	ClientID string `json:"client_id"`
	// request endpoint
	Endpoint string `json:"endpoint"`
	// optional Azure Resource ID of the managed identity you would like the token for, required, if your VM has multiple user-assigned managed identities
	MiResID string `json:"mi_res_id"`
	// optional object ID of the managed identity you would like the token for, required if your VM has multiple user-assigned managed identities
	ObjectID string `json:"object_id"`
	// app ID URI of the target resource
	Resource string `json:"resource"`
}

func (o *SystemsV1AzureManagedIdentitiesAuthPlugin) GetAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.APIVersion
}

func (o *SystemsV1AzureManagedIdentitiesAuthPlugin) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SystemsV1AzureManagedIdentitiesAuthPlugin) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *SystemsV1AzureManagedIdentitiesAuthPlugin) GetMiResID() string {
	if o == nil {
		return ""
	}
	return o.MiResID
}

func (o *SystemsV1AzureManagedIdentitiesAuthPlugin) GetObjectID() string {
	if o == nil {
		return ""
	}
	return o.ObjectID
}

func (o *SystemsV1AzureManagedIdentitiesAuthPlugin) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}
