// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// StacksV1StackConfigTypeParameters - stack type parameter values (for template.* types)
type StacksV1StackConfigTypeParameters struct {
}

type StacksV1StackConfig struct {
	Authz *SystemsV1AuthzConfig `json:"authz,omitempty"`
	// datasources created for the stack
	Datasources []SystemsV1DatasourceConfig `json:"datasources,omitempty"`
	Description string                      `json:"description"`
	// current stack deployment errors
	Errors map[string]SystemsV1AgentErrors `json:"errors,omitempty"`
	ID     string                          `json:"id"`
	// IDs of systems matching the stack
	MatchingSystems []string         `json:"matching_systems,omitempty"`
	Metadata        MetaV1ObjectMeta `json:"metadata"`
	// A history of any migrations performed on this stack
	MigrationHistory []SystemsV1MigrationRecord `json:"migration_history,omitempty"`
	// minimum running OPA version for any of the matching systems
	MinimumOpaVersion *string                      `json:"minimum_opa_version,omitempty"`
	Name              string                       `json:"name"`
	Policies          []SystemsV1PolicyConfig      `json:"policies"`
	ReadOnly          bool                         `json:"read_only"`
	SourceControl     *StacksV1SourceControlConfig `json:"source_control,omitempty"`
	Status            string                       `json:"status"`
	Type              string                       `json:"type"`
	// stack type parameter values (for template.* types)
	TypeParameters *StacksV1StackConfigTypeParameters `json:"type_parameters,omitempty"`
}

func (o *StacksV1StackConfig) GetAuthz() *SystemsV1AuthzConfig {
	if o == nil {
		return nil
	}
	return o.Authz
}

func (o *StacksV1StackConfig) GetDatasources() []SystemsV1DatasourceConfig {
	if o == nil {
		return nil
	}
	return o.Datasources
}

func (o *StacksV1StackConfig) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *StacksV1StackConfig) GetErrors() map[string]SystemsV1AgentErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *StacksV1StackConfig) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *StacksV1StackConfig) GetMatchingSystems() []string {
	if o == nil {
		return nil
	}
	return o.MatchingSystems
}

func (o *StacksV1StackConfig) GetMetadata() MetaV1ObjectMeta {
	if o == nil {
		return MetaV1ObjectMeta{}
	}
	return o.Metadata
}

func (o *StacksV1StackConfig) GetMigrationHistory() []SystemsV1MigrationRecord {
	if o == nil {
		return nil
	}
	return o.MigrationHistory
}

func (o *StacksV1StackConfig) GetMinimumOpaVersion() *string {
	if o == nil {
		return nil
	}
	return o.MinimumOpaVersion
}

func (o *StacksV1StackConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *StacksV1StackConfig) GetPolicies() []SystemsV1PolicyConfig {
	if o == nil {
		return []SystemsV1PolicyConfig{}
	}
	return o.Policies
}

func (o *StacksV1StackConfig) GetReadOnly() bool {
	if o == nil {
		return false
	}
	return o.ReadOnly
}

func (o *StacksV1StackConfig) GetSourceControl() *StacksV1SourceControlConfig {
	if o == nil {
		return nil
	}
	return o.SourceControl
}

func (o *StacksV1StackConfig) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *StacksV1StackConfig) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *StacksV1StackConfig) GetTypeParameters() *StacksV1StackConfigTypeParameters {
	if o == nil {
		return nil
	}
	return o.TypeParameters
}
