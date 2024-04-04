// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SystemsV1ServerTLSConfig struct {
	CaCert           types.String `tfsdk:"ca_cert"`
	SystemCaRequired types.Bool   `tfsdk:"system_ca_required"`
}
