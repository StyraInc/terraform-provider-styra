// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SystemsV1BundleDistributionS3ConfigOpaCredentialsMetadataCredentials struct {
	AwsRegion string  `json:"aws_region"`
	IamRole   *string `json:"iam_role,omitempty"`
}

func (o *SystemsV1BundleDistributionS3ConfigOpaCredentialsMetadataCredentials) GetAwsRegion() string {
	if o == nil {
		return ""
	}
	return o.AwsRegion
}

func (o *SystemsV1BundleDistributionS3ConfigOpaCredentialsMetadataCredentials) GetIamRole() *string {
	if o == nil {
		return nil
	}
	return o.IamRole
}
