// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type GitV1SSHCredentials struct {
	// Passphrase is looked under the key passphrase/<pass>
	Passphrase string `json:"passphrase"`
	// PrivateKey is looked under the key private-key/<key>
	PrivateKey string `json:"private_key"`
}

func (o *GitV1SSHCredentials) GetPassphrase() string {
	if o == nil {
		return ""
	}
	return o.Passphrase
}

func (o *GitV1SSHCredentials) GetPrivateKey() string {
	if o == nil {
		return ""
	}
	return o.PrivateKey
}
