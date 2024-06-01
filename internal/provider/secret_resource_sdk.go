// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SecretResourceModel) ToSharedSecretsV1SecretsPutRequest() *shared.SecretsV1SecretsPutRequest {
	description := r.Description.ValueString()
	name := r.Name.ValueString()
	secret := r.Secret.ValueString()
	out := shared.SecretsV1SecretsPutRequest{
		Description: description,
		Name:        name,
		Secret:      secret,
	}
	return &out
}

func (r *SecretResourceModel) RefreshFromSharedSecretsV1SecretsPutResponse(resp *shared.SecretsV1SecretsPutResponse) {
	if resp != nil {
		r.RequestID = types.StringPointerValue(resp.RequestID)
	}
}

func (r *SecretResourceModel) RefreshFromSharedSecretsV1Secret(resp *shared.SecretsV1Secret) {
	if resp != nil {
		r.Description = types.StringValue(resp.Description)
		r.ID = types.StringValue(resp.ID)
		if resp.Metadata.CreatedAt != nil {
			r.Metadata.CreatedAt = types.StringValue(resp.Metadata.CreatedAt.Format(time.RFC3339Nano))
		} else {
			r.Metadata.CreatedAt = types.StringNull()
		}
		r.Metadata.CreatedBy = types.StringPointerValue(resp.Metadata.CreatedBy)
		r.Metadata.CreatedThrough = types.StringPointerValue(resp.Metadata.CreatedThrough)
		if resp.Metadata.LastModifiedAt != nil {
			r.Metadata.LastModifiedAt = types.StringValue(resp.Metadata.LastModifiedAt.Format(time.RFC3339Nano))
		} else {
			r.Metadata.LastModifiedAt = types.StringNull()
		}
		r.Metadata.LastModifiedBy = types.StringPointerValue(resp.Metadata.LastModifiedBy)
		r.Metadata.LastModifiedThrough = types.StringPointerValue(resp.Metadata.LastModifiedThrough)
		r.Name = types.StringValue(resp.Name)
	}
}
