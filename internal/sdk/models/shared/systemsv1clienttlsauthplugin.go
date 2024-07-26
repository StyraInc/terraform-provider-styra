// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type SystemsV1ClientTLSAuthPlugin struct {
	// the path to the client certificate to authenticate with
	Cert string `json:"cert"`
	// the path to the private key of the client certificate
	PrivateKey string `json:"private_key"`
	// the passphrase to use for the private key
	PrivateKeyPassphrase *string `json:"private_key_passphrase,omitempty"`
}

func (o *SystemsV1ClientTLSAuthPlugin) GetCert() string {
	if o == nil {
		return ""
	}
	return o.Cert
}

func (o *SystemsV1ClientTLSAuthPlugin) GetPrivateKey() string {
	if o == nil {
		return ""
	}
	return o.PrivateKey
}

func (o *SystemsV1ClientTLSAuthPlugin) GetPrivateKeyPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.PrivateKeyPassphrase
}
