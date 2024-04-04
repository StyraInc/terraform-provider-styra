// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GitV1GitRepoConfig struct {
	// Commit SHA. Only one of reference or commit can be set at any time
	Commit string `json:"commit"`
	// Credentials are looked under the key <name>/<creds>
	Credentials string `json:"credentials"`
	// Path to limit the import to
	Path string `json:"path"`
	// Remote reference. Only one of reference or commit can be set at any time
	Reference      string               `json:"reference"`
	SSHCredentials *GitV1SSHCredentials `json:"ssh_credentials,omitempty"`
	// Repository URL
	URL string `json:"url"`
}

func (o *GitV1GitRepoConfig) GetCommit() string {
	if o == nil {
		return ""
	}
	return o.Commit
}

func (o *GitV1GitRepoConfig) GetCredentials() string {
	if o == nil {
		return ""
	}
	return o.Credentials
}

func (o *GitV1GitRepoConfig) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *GitV1GitRepoConfig) GetReference() string {
	if o == nil {
		return ""
	}
	return o.Reference
}

func (o *GitV1GitRepoConfig) GetSSHCredentials() *GitV1SSHCredentials {
	if o == nil {
		return nil
	}
	return o.SSHCredentials
}

func (o *GitV1GitRepoConfig) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
