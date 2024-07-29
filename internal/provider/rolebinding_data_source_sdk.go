// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	tfTypes "github.com/StyraInc/terraform-provider-styra/internal/provider/types"
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *RoleBindingDataSourceModel) RefreshFromSharedAuthzV2RoleBindingConfig(resp *shared.AuthzV2RoleBindingConfig) {
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
	r.ResourceFilter.ID = types.StringValue(resp.ResourceFilter.ID)
	r.ResourceFilter.Kind = types.StringValue(resp.ResourceFilter.Kind)
	r.RoleID = types.StringValue(resp.RoleID)
	r.Subjects = []tfTypes.AuthzV2Subject{}
	if len(r.Subjects) > len(resp.Subjects) {
		r.Subjects = r.Subjects[:len(resp.Subjects)]
	}
	for subjectsCount, subjectsItem := range resp.Subjects {
		var subjects1 tfTypes.AuthzV2Subject
		if subjectsItem.ClaimConfig == nil {
			subjects1.ClaimConfig = nil
		} else {
			subjects1.ClaimConfig = &tfTypes.AuthzV2ClaimConfig{}
			subjects1.ClaimConfig.IdentityProvider = types.StringPointerValue(subjectsItem.ClaimConfig.IdentityProvider)
			subjects1.ClaimConfig.Key = types.StringValue(subjectsItem.ClaimConfig.Key)
			subjects1.ClaimConfig.Value = types.StringValue(subjectsItem.ClaimConfig.Value)
		}
		subjects1.ID = types.StringPointerValue(subjectsItem.ID)
		subjects1.Kind = types.StringValue(subjectsItem.Kind)
		if subjectsCount+1 > len(r.Subjects) {
			r.Subjects = append(r.Subjects, subjects1)
		} else {
			r.Subjects[subjectsCount].ClaimConfig = subjects1.ClaimConfig
			r.Subjects[subjectsCount].ID = subjects1.ID
			r.Subjects[subjectsCount].Kind = subjects1.Kind
		}
	}
}
