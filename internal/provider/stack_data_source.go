// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

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
var _ datasource.DataSource = &StackDataSource{}
var _ datasource.DataSourceWithConfigure = &StackDataSource{}

func NewStackDataSource() datasource.DataSource {
	return &StackDataSource{}
}

// StackDataSource is the data source implementation.
type StackDataSource struct {
	client *sdk.StyraDas
}

// StackDataSourceModel describes the data model.
type StackDataSourceModel struct {
	Authz             *tfTypes.SystemsV1AuthzConfig           `tfsdk:"authz"`
	Datasources       []tfTypes.SystemsV1DatasourceConfig     `tfsdk:"datasources"`
	Description       types.String                            `tfsdk:"description"`
	Errors            map[string]tfTypes.SystemsV1AgentErrors `tfsdk:"errors"`
	ID                types.String                            `tfsdk:"id"`
	MatchingSystems   []types.String                          `tfsdk:"matching_systems"`
	Metadata          tfTypes.MetaV2ObjectMeta                `tfsdk:"metadata"`
	MigrationHistory  []tfTypes.SystemsV1MigrationRecord      `tfsdk:"migration_history"`
	MinimumOpaVersion types.String                            `tfsdk:"minimum_opa_version"`
	Name              types.String                            `tfsdk:"name"`
	Policies          []tfTypes.SystemsV1PolicyConfig         `tfsdk:"policies"`
	ReadOnly          types.Bool                              `tfsdk:"read_only"`
	SourceControl     *tfTypes.StacksV1SourceControlConfig    `tfsdk:"source_control"`
	Status            types.String                            `tfsdk:"status"`
	Type              types.String                            `tfsdk:"type"`
	TypeParameters    *tfTypes.TypeParameters                 `tfsdk:"type_parameters"`
}

// Metadata returns the data source type name.
func (r *StackDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_stack"
}

// Schema defines the schema for the data source.
func (r *StackDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Stack DataSource",

		Attributes: map[string]schema.Attribute{
			"authz": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"role_bindings": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `role binding ID`,
								},
								"role_name": schema.StringAttribute{
									Computed:    true,
									Description: `role name`,
								},
							},
						},
						Description: `a list of role binding configs`,
					},
				},
			},
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
				Description: `datasources created for the stack`,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"errors": schema.MapNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"errors": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
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
							Description: `list of system errors`,
						},
						"waiting": schema.BoolAttribute{
							Computed:    true,
							Description: `true if the the system is waiting for error to be resolved`,
						},
					},
				},
				Description: `current stack deployment errors`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `stack id`,
			},
			"matching_systems": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `IDs of systems matching the stack`,
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
			"migration_history": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"from": schema.StringAttribute{
							Computed: true,
						},
						"initiated_by": schema.StringAttribute{
							Computed: true,
						},
						"initiating_user": schema.StringAttribute{
							Computed: true,
						},
						"migrated_at": schema.StringAttribute{
							Computed: true,
						},
						"recovered": schema.BoolAttribute{
							Computed: true,
						},
						"to": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				Description: `A history of any migrations performed on this stack`,
			},
			"minimum_opa_version": schema.StringAttribute{
				Computed:    true,
				Description: `minimum running OPA version for any of the matching systems`,
			},
			"name": schema.StringAttribute{
				Computed: true,
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
					"stack_origin": schema.SingleNestedAttribute{
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
			"status": schema.StringAttribute{
				Computed: true,
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
			"type_parameters": schema.SingleNestedAttribute{
				Computed:    true,
				Attributes:  map[string]schema.Attribute{},
				Description: `stack type parameter values (for template.* types)`,
			},
		},
	}
}

func (r *StackDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *StackDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *StackDataSourceModel
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

	id := data.ID.ValueString()
	request := operations.GetStackRequest{
		ID: id,
	}
	res, err := r.client.Stacks.GetStack(ctx, request)
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
	if !(res.StacksV1StacksGetResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedStacksV1StackConfig(&res.StacksV1StacksGetResponse.Result)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
