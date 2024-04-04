// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SystemsV1ExternalBundleConfig struct {
	// externally configured bundles, use name of bundle as key
	Bundles map[string]SystemsV1OpaConfigBundleDeclaration `json:"bundles,omitempty"`
	// externally configured services
	Services []SystemsV1OpaConfigServiceDeclaration `json:"services"`
}

func (o *SystemsV1ExternalBundleConfig) GetBundles() map[string]SystemsV1OpaConfigBundleDeclaration {
	if o == nil {
		return nil
	}
	return o.Bundles
}

func (o *SystemsV1ExternalBundleConfig) GetServices() []SystemsV1OpaConfigServiceDeclaration {
	if o == nil {
		return []SystemsV1OpaConfigServiceDeclaration{}
	}
	return o.Services
}
