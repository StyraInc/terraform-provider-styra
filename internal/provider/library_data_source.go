// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/StyraInc/terraform-provider-styra/internal/provider/types"
	"github.com/StyraInc/terraform-provider-styra/internal/sdk"
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &LibraryDataSource{}
var _ datasource.DataSourceWithConfigure = &LibraryDataSource{}

func NewLibraryDataSource() datasource.DataSource {
	return &LibraryDataSource{}
}

// LibraryDataSource is the data source implementation.
type LibraryDataSource struct {
	client *sdk.StyraDas
}

// LibraryDataSourceModel describes the data model.
type LibraryDataSourceModel struct {
	DependantBundles types.String                             `tfsdk:"dependant_bundles"`
	ID               types.String                             `tfsdk:"id"`
	Result           tfTypes.LibrariesV1LibraryEntityExpanded `tfsdk:"result"`
}

// Metadata returns the data source type name.
func (r *LibraryDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_library"
}

// Schema defines the schema for the data source.
func (r *LibraryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Library DataSource",

		Attributes: map[string]schema.Attribute{
			"dependant_bundles": schema.StringAttribute{
				Optional:    true,
				Description: `level of report for bundles depending on the library. One of (none, active, all). "active" is the default`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `id`,
			},
			"result": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"datasources": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"category": schema.StringAttribute{
									Computed:    true,
									Description: `datasource category`,
								},
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `datasource ID`,
								},
								"optional": schema.BoolAttribute{
									Computed:    true,
									Description: `optional datasources can be deleted without being recreated automatically`,
								},
								"status": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"code": schema.StringAttribute{
											Computed: true,
										},
										"message": schema.StringAttribute{
											Computed: true,
										},
										"timestamp": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
					},
					"description": schema.StringAttribute{
						Computed: true,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"metadata": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"created_at": schema.StringAttribute{
								Computed: true,
							},
							"created_by": schema.StringAttribute{
								Computed: true,
							},
							"created_through": schema.StringAttribute{
								Computed: true,
							},
							"last_modified_at": schema.StringAttribute{
								Computed: true,
							},
							"last_modified_by": schema.StringAttribute{
								Computed: true,
							},
							"last_modified_through": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"policies": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"created": schema.StringAttribute{
									Computed:    true,
									Description: `policy on when to (re)generate the policy`,
								},
								"enforcement": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"enforced": schema.BoolAttribute{
											Computed:    true,
											Description: `true if the policy is enforced`,
										},
										"type": schema.StringAttribute{
											Computed:    true,
											Description: `enforcement type e.g. opa, test, mask`,
										},
									},
								},
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `policy ID (path)`,
								},
								"modules": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Computed:    true,
												Description: `module name`,
											},
											"placeholder": schema.BoolAttribute{
												Computed:    true,
												Description: `module is a placeholder`,
											},
											"read_only": schema.BoolAttribute{
												Computed:    true,
												Description: `true if module is read-only`,
											},
											"rules": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"allow": schema.Int64Attribute{
														Computed:    true,
														Description: `number of allow rules`,
													},
													"deny": schema.Int64Attribute{
														Computed:    true,
														Description: `number of deny rules`,
													},
													"enforce": schema.Int64Attribute{
														Computed:    true,
														Description: `number of enforce rules`,
													},
													"ignore": schema.Int64Attribute{
														Computed:    true,
														Description: `number of ignore rules`,
													},
													"monitor": schema.Int64Attribute{
														Computed:    true,
														Description: `number of monitor rules`,
													},
													"notify": schema.Int64Attribute{
														Computed:    true,
														Description: `number of notify rules`,
													},
													"other": schema.Int64Attribute{
														Computed:    true,
														Description: `number of unclassified rules`,
													},
													"test": schema.Int64Attribute{
														Computed:    true,
														Description: `number of test rules`,
													},
													"total": schema.Int64Attribute{
														Computed:    true,
														Description: `total number of rules`,
													},
												},
											},
										},
									},
									Description: `rego modules policy consists of`,
								},
								"rules": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"allow": schema.Int64Attribute{
											Computed:    true,
											Description: `number of allow rules`,
										},
										"deny": schema.Int64Attribute{
											Computed:    true,
											Description: `number of deny rules`,
										},
										"enforce": schema.Int64Attribute{
											Computed:    true,
											Description: `number of enforce rules`,
										},
										"ignore": schema.Int64Attribute{
											Computed:    true,
											Description: `number of ignore rules`,
										},
										"monitor": schema.Int64Attribute{
											Computed:    true,
											Description: `number of monitor rules`,
										},
										"notify": schema.Int64Attribute{
											Computed:    true,
											Description: `number of notify rules`,
										},
										"other": schema.Int64Attribute{
											Computed:    true,
											Description: `number of unclassified rules`,
										},
										"test": schema.Int64Attribute{
											Computed:    true,
											Description: `number of test rules`,
										},
										"total": schema.Int64Attribute{
											Computed:    true,
											Description: `total number of rules`,
										},
									},
								},
								"type": schema.StringAttribute{
									Computed:    true,
									Description: `policy type e.g. validating/rules`,
								},
							},
						},
					},
					"read_only": schema.BoolAttribute{
						Computed: true,
					},
					"source_control": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"library_origin": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"commit": schema.StringAttribute{
										Computed:    true,
										Description: `Commit SHA. Only one of reference or commit can be set at any time`,
									},
									"credentials": schema.StringAttribute{
										Computed:    true,
										Description: `Credentials are looked under the key <name>/<creds>`,
									},
									"path": schema.StringAttribute{
										Computed:    true,
										Description: `Path to limit the import to`,
									},
									"reference": schema.StringAttribute{
										Computed:    true,
										Description: `Remote reference. Only one of reference or commit can be set at any time`,
									},
									"ssh_credentials": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"passphrase": schema.StringAttribute{
												Computed:    true,
												Description: `Passphrase is looked under the key passphrase/<pass>`,
											},
											"private_key": schema.StringAttribute{
												Computed:    true,
												Description: `PrivateKey is looked under the key private-key/<key>`,
											},
										},
									},
									"url": schema.StringAttribute{
										Computed:    true,
										Description: `Repository URL`,
									},
								},
							},
							"use_workspace_settings": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
					"used_by": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bundles": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"bundle_id": schema.StringAttribute{
												Computed: true,
											},
											"version": schema.Int64Attribute{
												Computed: true,
											},
										},
									},
								},
								"system_id": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *LibraryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.StyraDas)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.StyraDas, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *LibraryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *LibraryDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var id string
	id = data.ID.ValueString()

	dependantBundles := new(string)
	if !data.DependantBundles.IsUnknown() && !data.DependantBundles.IsNull() {
		*dependantBundles = data.DependantBundles.ValueString()
	} else {
		dependantBundles = nil
	}
	request := operations.LibrariesGetRequest{
		ID:               id,
		DependantBundles: dependantBundles,
	}
	res, err := r.client.Libraries.LibrariesGet(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.LibrariesV1LibraryResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedLibrariesV1LibraryResponse(res.LibrariesV1LibraryResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
