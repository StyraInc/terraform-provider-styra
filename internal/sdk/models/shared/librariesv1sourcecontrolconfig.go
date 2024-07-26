// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type LibrariesV1SourceControlConfig struct {
	LibraryOrigin        *GitV1GitRepoConfig `json:"library_origin,omitempty"`
	UseWorkspaceSettings bool                `json:"use_workspace_settings"`
}

func (o *LibrariesV1SourceControlConfig) GetLibraryOrigin() *GitV1GitRepoConfig {
	if o == nil {
		return nil
	}
	return o.LibraryOrigin
}

func (o *LibrariesV1SourceControlConfig) GetUseWorkspaceSettings() bool {
	if o == nil {
		return false
	}
	return o.UseWorkspaceSettings
}
