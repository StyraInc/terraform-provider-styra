// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SystemsV1AwsSigningAuthPlugin struct {
	EnvironmentCredentials types.String                                                             `tfsdk:"environment_credentials"`
	MetadataCredentials    *SystemsV1BundleDistributionS3ConfigOpaCredentialsMetadataCredentials    `tfsdk:"metadata_credentials"`
	ProfileCredentials     *SystemsV1AwsProfileCredentialService                                    `tfsdk:"profile_credentials"`
	Service                types.String                                                             `tfsdk:"service"`
	WebIdentityCredentials *SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials `tfsdk:"web_identity_credentials"`
}
