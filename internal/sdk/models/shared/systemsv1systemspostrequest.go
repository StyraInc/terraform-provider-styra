// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/internal/utils"
)

// SystemsV1SystemsPostRequestTypeParameters - system type parameter values (for template.* types)
type SystemsV1SystemsPostRequestTypeParameters struct {
}

type SystemsV1SystemsPostRequest struct {
	BundleDownload *SystemsV1BundleDownloadConfig `json:"bundle_download,omitempty"`
	BundleRegistry *SystemsV1BundleRegistryConfig `json:"bundle_registry,omitempty"`
	// only put data in the context bundle
	ContextBundleDataOnly *bool `json:"context_bundle_data_only,omitempty"`
	// list of path prefixes for policies/datasources that go into the second (context) bundle
	ContextBundleRoots []string `json:"context_bundle_roots,omitempty"`
	// location of key attributes and additional columns in the decisions grouped by policy entry point path
	DecisionMappings     map[string]SystemsV1RuleDecisionMappings `json:"decision_mappings,omitempty"`
	DeploymentParameters *SystemsV1SystemDeploymentParameters     `json:"deployment_parameters,omitempty"`
	// description for the system
	Description *string `json:"description,omitempty"`
	// error/warning configuration: one of "all", "errors", "none"
	ErrorSetting    *string                        `json:"error_setting,omitempty"`
	ExternalBundles *SystemsV1ExternalBundleConfig `json:"external_bundles,omitempty"`
	// optional parameter to map Styra DAS system ID to external IDs used by a customer. (mapping can be retrieved with TranslateExternalIds operation)
	ExternalID *string `json:"external_id,omitempty"`
	// when set, stacks that are not linked to this system will be filtered out of its bundles
	FilterStacks *bool `json:"filter_stacks,omitempty"`
	// optional parameter to specify the Kafka topic where the decision logs for this system should be published (ignored if Kafka is not configured for the workspace for decision export)
	KafkaTopic *string `json:"kafka_topic,omitempty"`
	// enable mock OPAs for this system
	MockOpaEnabled *bool `json:"mock_opa_enabled,omitempty"`
	// system name
	Name string `json:"name"`
	// prevents users from modifying policies using Styra UIs
	ReadOnly      *bool                     `default:"false" json:"read_only"`
	SourceControl *GitV1SourceControlConfig `json:"source_control,omitempty"`
	// system type e.g. kubernetes
	Type string `json:"type"`
	// system type parameter values (for template.* types)
	TypeParameters *SystemsV1SystemsPostRequestTypeParameters `json:"type_parameters,omitempty"`
}

func (s SystemsV1SystemsPostRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SystemsV1SystemsPostRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SystemsV1SystemsPostRequest) GetBundleDownload() *SystemsV1BundleDownloadConfig {
	if o == nil {
		return nil
	}
	return o.BundleDownload
}

func (o *SystemsV1SystemsPostRequest) GetBundleRegistry() *SystemsV1BundleRegistryConfig {
	if o == nil {
		return nil
	}
	return o.BundleRegistry
}

func (o *SystemsV1SystemsPostRequest) GetContextBundleDataOnly() *bool {
	if o == nil {
		return nil
	}
	return o.ContextBundleDataOnly
}

func (o *SystemsV1SystemsPostRequest) GetContextBundleRoots() []string {
	if o == nil {
		return nil
	}
	return o.ContextBundleRoots
}

func (o *SystemsV1SystemsPostRequest) GetDecisionMappings() map[string]SystemsV1RuleDecisionMappings {
	if o == nil {
		return nil
	}
	return o.DecisionMappings
}

func (o *SystemsV1SystemsPostRequest) GetDeploymentParameters() *SystemsV1SystemDeploymentParameters {
	if o == nil {
		return nil
	}
	return o.DeploymentParameters
}

func (o *SystemsV1SystemsPostRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SystemsV1SystemsPostRequest) GetErrorSetting() *string {
	if o == nil {
		return nil
	}
	return o.ErrorSetting
}

func (o *SystemsV1SystemsPostRequest) GetExternalBundles() *SystemsV1ExternalBundleConfig {
	if o == nil {
		return nil
	}
	return o.ExternalBundles
}

func (o *SystemsV1SystemsPostRequest) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *SystemsV1SystemsPostRequest) GetFilterStacks() *bool {
	if o == nil {
		return nil
	}
	return o.FilterStacks
}

func (o *SystemsV1SystemsPostRequest) GetKafkaTopic() *string {
	if o == nil {
		return nil
	}
	return o.KafkaTopic
}

func (o *SystemsV1SystemsPostRequest) GetMockOpaEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.MockOpaEnabled
}

func (o *SystemsV1SystemsPostRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SystemsV1SystemsPostRequest) GetReadOnly() *bool {
	if o == nil {
		return nil
	}
	return o.ReadOnly
}

func (o *SystemsV1SystemsPostRequest) GetSourceControl() *GitV1SourceControlConfig {
	if o == nil {
		return nil
	}
	return o.SourceControl
}

func (o *SystemsV1SystemsPostRequest) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *SystemsV1SystemsPostRequest) GetTypeParameters() *SystemsV1SystemsPostRequestTypeParameters {
	if o == nil {
		return nil
	}
	return o.TypeParameters
}
