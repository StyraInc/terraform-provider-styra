// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SystemsV1AwsProfileCredentialService struct {
	AwsRegion types.String `tfsdk:"aws_region"`
	Path      types.String `tfsdk:"path"`
	Profile   types.String `tfsdk:"profile"`
}
