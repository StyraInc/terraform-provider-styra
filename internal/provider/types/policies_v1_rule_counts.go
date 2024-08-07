// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type PoliciesV1RuleCounts struct {
	Allow   types.Int64 `tfsdk:"allow"`
	Deny    types.Int64 `tfsdk:"deny"`
	Enforce types.Int64 `tfsdk:"enforce"`
	Ignore  types.Int64 `tfsdk:"ignore"`
	Monitor types.Int64 `tfsdk:"monitor"`
	Notify  types.Int64 `tfsdk:"notify"`
	Other   types.Int64 `tfsdk:"other"`
	Test    types.Int64 `tfsdk:"test"`
	Total   types.Int64 `tfsdk:"total"`
}
