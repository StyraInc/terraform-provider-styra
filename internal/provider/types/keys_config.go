// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type KeysConfig struct {
	Algorithm  types.String `tfsdk:"algorithm"`
	Key        types.String `tfsdk:"key"`
	PrivateKey types.String `tfsdk:"private_key"`
	Scope      types.String `tfsdk:"scope"`
}
