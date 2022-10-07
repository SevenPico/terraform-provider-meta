package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &ContextData{}

type ContextData struct {
	Enabled      *bool                 `tfsdk:"enabled"`
	Attributes   []string              `tfsdk:"attributes"`
	Tags         map[string]string     `tfsdk:"tags"`
	IdDescriptor *Descriptor           `tfsdk:"id_descriptor"`
	Descriptors  map[string]Descriptor `tfsdk:"descriptors"`

	// Computed
	Id      *string            `tfsdk:"id"`
	Outputs map[string]*string `tfsdk:"outputs"`
}

type Descriptor struct {
	Delimiter  *string  `tfsdk:"delimiter"`
	Order      []string `tfsdk:"order"`
	Upper      *bool    `tfsdk:"upper"`
	Lower      *bool    `tfsdk:"lower"`
	Title      *bool    `tfsdk:"title"`
	Reverse    *bool    `tfsdk:"reverse"`
	Attributes *bool    `tfsdk:"attributes"`
	Limit      *int     `tfsdk:"limit"`
	// Replace map[string]string
}

func NewContextDataSource() datasource.DataSource {
	return &ContextData{}
}

func (d *ContextData) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "context"
}

func (d *ContextData) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	descriptorAttribute := map[string]tfsdk.Attribute{
		"order": {
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"delimiter": {
			Type:     types.StringType,
			Optional: true,
		},
		"lower": {
			Type:     types.BoolType,
			Optional: true,
		},
		"upper": {
			Type:     types.BoolType,
			Optional: true,
		},
		"title": {
			Type:     types.BoolType,
			Optional: true,
		},
		"reverse": {
			Type:     types.BoolType,
			Optional: true,
		},
		"attributes": {
			Type:     types.BoolType,
			Optional: true,
		},
		"limit": {
			Type:     types.NumberType,
			Optional: true,
		},
	}

	attributes := map[string]tfsdk.Attribute{
		"enabled": {
			MarkdownDescription: "Set `true` if resources using this context should be created.",
			Type:                types.BoolType,
			Optional:            true,
		},
		"attributes": {
			MarkdownDescription: "List of strings.",
			Type:                types.ListType{ElemType: types.StringType},
			Optional:            true,
		},
		"tags": {
			MarkdownDescription: "Map of strings.",
			Type:                types.MapType{ElemType: types.StringType},
			Optional:            true,
		},

		"descriptors": {
			Attributes: tfsdk.MapNestedAttributes(descriptorAttribute),
			Optional:   true,
		},

		"id_descriptor": {
			Attributes: tfsdk.SingleNestedAttributes(descriptorAttribute),
			Optional:   true,
		},

		// Computed
		"id": {
			Type:     types.StringType,
			Computed: true,
		},
		"outputs": {
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
	}

	s := tfsdk.Schema{
		MarkdownDescription: "Context data source",
		Attributes:          attributes,
	}

	return s, nil
}

func (d *ContextData) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	resp.Diagnostics.AddWarning("This is a warning!", "Beware!")

	// Read Terraform configuration data into the model
	diag := req.Config.Get(ctx, d)
	resp.Diagnostics.Append(diag...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "read a data source")

	d.Id = d.BuildDescriptor(d.IdDescriptor)

	d.Outputs = map[string]*string{}

	for k, v := range d.Descriptors {
		d.Outputs[k] = d.BuildDescriptor(&v)
	}

	// Save data into Terraform state
	diag = resp.State.Set(ctx, d)
	resp.Diagnostics.Append(diag...)
}

func (d *ContextData) BuildDescriptor(descriptor *Descriptor) *string {
	idParts := []string{}

	for _, tag_key := range descriptor.Order {
		if tag_val, ok := d.Tags[tag_key]; ok {
			idParts = append(idParts, tag_val)
		} else {
			// TODO - warn
		}
	}

	if descriptor.Attributes != nil && *descriptor.Attributes {
		idParts = append(idParts, d.Attributes...)
	}

	if descriptor.Reverse != nil && *descriptor.Reverse {
		for i, j := 0, len(idParts)-1; i < j; i, j = i+1, j-1 {
			idParts[i], idParts[j] = idParts[j], idParts[i]
		}
	}

	id := strings.Join(idParts, *descriptor.Delimiter)

	if descriptor.Upper != nil && *descriptor.Upper {
		id = strings.ToUpper(id)
	}
	if descriptor.Lower != nil && *descriptor.Lower {
		id = strings.ToLower(id)
	}
	if descriptor.Title != nil && *descriptor.Title {
		id = strings.Title(id)
	}

	if descriptor.Limit != nil && len(id) > *descriptor.Limit {
		id = id[:*descriptor.Limit]
		id = strings.TrimRight(id, *descriptor.Delimiter)
	}

	return &id
}
