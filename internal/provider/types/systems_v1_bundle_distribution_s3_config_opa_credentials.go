// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type SystemsV1BundleDistributionS3ConfigOpaCredentials struct {
	EnvironmentCredentials *TypeParameters                                                          `tfsdk:"environment_credentials"`
	MetadataCredentials    *SystemsV1BundleDistributionS3ConfigOpaCredentialsMetadataCredentials    `tfsdk:"metadata_credentials"`
	WebIdentityCredentials *SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials `tfsdk:"web_identity_credentials"`
}
