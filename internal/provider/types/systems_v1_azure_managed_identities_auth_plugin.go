// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SystemsV1AzureManagedIdentitiesAuthPlugin struct {
	APIVersion types.String `tfsdk:"api_version"`
	ClientID   types.String `tfsdk:"client_id"`
	Endpoint   types.String `tfsdk:"endpoint"`
	MiResID    types.String `tfsdk:"mi_res_id"`
	ObjectID   types.String `tfsdk:"object_id"`
	Resource   types.String `tfsdk:"resource"`
}
