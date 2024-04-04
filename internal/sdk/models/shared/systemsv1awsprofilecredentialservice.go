// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SystemsV1AwsProfileCredentialService struct {
	// the AWS region to use for the AWS signing service credential method
	AwsRegion string `json:"aws_region"`
	// the path to the shared credentials file
	Path *string `json:"path,omitempty"`
	// AWS Profile to extract credentials from the credentials file
	Profile *string `json:"profile,omitempty"`
}

func (o *SystemsV1AwsProfileCredentialService) GetAwsRegion() string {
	if o == nil {
		return ""
	}
	return o.AwsRegion
}

func (o *SystemsV1AwsProfileCredentialService) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *SystemsV1AwsProfileCredentialService) GetProfile() *string {
	if o == nil {
		return nil
	}
	return o.Profile
}
