// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type LibrariesV1SystemUsingLibrary struct {
	Bundles  []LibrariesV1SystemBundle `tfsdk:"bundles"`
	SystemID types.String              `tfsdk:"system_id"`
}
