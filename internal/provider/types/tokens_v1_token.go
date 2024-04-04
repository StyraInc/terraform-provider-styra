// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type TokensV1Token struct {
	AllowPathPatterns []types.String   `tfsdk:"allow_path_patterns"`
	Description       types.String     `tfsdk:"description"`
	Expires           types.String     `tfsdk:"expires"`
	ID                types.String     `tfsdk:"id"`
	Metadata          MetaV2ObjectMeta `tfsdk:"metadata"`
	Token             types.String     `tfsdk:"token"`
	TTL               types.String     `tfsdk:"ttl"`
	Uses              types.Int64      `tfsdk:"uses"`
}
