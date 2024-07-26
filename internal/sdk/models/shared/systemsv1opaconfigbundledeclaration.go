// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type SystemsV1OpaConfigBundleDeclaration struct {
	// persist activated bundles to disk
	Persist *bool                   `json:"persist,omitempty"`
	Polling *SystemsV1PollingConfig `json:"polling,omitempty"`
	// resource path to use to download bundle from configured service
	Resource *string `json:"resource,omitempty"`
	// name of service to use to contact remote server
	Service string                       `json:"service"`
	Signing *SystemsV1VerificationConfig `json:"signing,omitempty"`
	// size limit for individual files contained in the bundle
	SizeLimitBytes *int64 `json:"size_limit_bytes,omitempty"`
}

func (o *SystemsV1OpaConfigBundleDeclaration) GetPersist() *bool {
	if o == nil {
		return nil
	}
	return o.Persist
}

func (o *SystemsV1OpaConfigBundleDeclaration) GetPolling() *SystemsV1PollingConfig {
	if o == nil {
		return nil
	}
	return o.Polling
}

func (o *SystemsV1OpaConfigBundleDeclaration) GetResource() *string {
	if o == nil {
		return nil
	}
	return o.Resource
}

func (o *SystemsV1OpaConfigBundleDeclaration) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *SystemsV1OpaConfigBundleDeclaration) GetSigning() *SystemsV1VerificationConfig {
	if o == nil {
		return nil
	}
	return o.Signing
}

func (o *SystemsV1OpaConfigBundleDeclaration) GetSizeLimitBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.SizeLimitBytes
}
