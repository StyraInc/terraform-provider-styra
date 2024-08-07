// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SystemsV1ServerTLSConfig struct {
	// the path to the root CA certificate. If not provided, this defaults to TLS using the host’s root CA set
	CaCert *string `json:"ca_cert,omitempty"`
	// require system certificate appended with root CA certificate
	SystemCaRequired *bool `json:"system_ca_required,omitempty"`
}

func (o *SystemsV1ServerTLSConfig) GetCaCert() *string {
	if o == nil {
		return nil
	}
	return o.CaCert
}

func (o *SystemsV1ServerTLSConfig) GetSystemCaRequired() *bool {
	if o == nil {
		return nil
	}
	return o.SystemCaRequired
}
