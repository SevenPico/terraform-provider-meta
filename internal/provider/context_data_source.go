package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &ContextDataSource{}

type ContextDataSource struct {
	Enabled       types.Bool              `tfsdk:"enabled"`
	Attributes    []types.String          `tfsdk:"attributes"`
	AttributesMap map[string]types.String `tfsdk:"attributes_map"`
	//Context       *ContextDataSource      `tfsdk:"context"`
	//Descriptors   map[string]Descriptor   `tfsdk:"descriptors"`

	// Outputs
	Id types.String `tfsdk:"id"`
}

//type ContextDataModel struct {
//	Enabled       types.Bool              `tfsdk:"enabled"`
//	Attributes    []types.String          `tfsdk:"attributes"`
//	AttributesMap map[string]types.String `tfsdk:"attributes_map"`
//	Context       *ContextDataSource      `tfsdk:"context"`
//	//Descriptors   map[string]Descriptor   `tfsdk:"descriptors"`
//}

// type Descriptor struct {
// 	Delimiter  *string  `tfsdk:"delimiter"`
// 	Order      []string `tfsdk:"order"`
// 	Upper      *bool    `tfsdk:"upper"`
// 	Lower      *bool    `tfsdk:"lower"`
// 	Title      *bool    `tfsdk:"title"`
// 	Reverse    *bool    `tfsdk:"reverse"`
// 	Attributes *bool    `tfsdk:"attributes"`
// 	Limit      *int     `tfsdk:"limit"`
// }

func NewContextDataSource() datasource.DataSource {
	return &ContextDataSource{}
}

func (d *ContextDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "context"
}

func (d *ContextDataSource) getPartialSchema() tfsdk.Schema {
	// descriptorAttributes := map[string]tfsdk.Attribute{
	// 	"order": {
	// 		MarkdownDescription: "Ordered list of keys in `tags` for which the corresponding value should be included in the output.\nDefault: `[]`",
	// 		Type:                types.ListType{ElemType: types.StringType},
	// 		Optional:            true,
	// 	},
	// 	"delimiter": {
	// 		MarkdownDescription: "String to separate `tags` values specified in `order`.\nDefault: \"\"",
	// 		Type:                types.StringType,
	// 		Optional:            true,
	// 	},
	// 	"lower": {
	// 		MarkdownDescription: "Set `true` to force output to lower-case",
	// 		Type:                types.BoolType,
	// 		Optional:            true,
	// 	},
	// 	"upper": {
	// 		MarkdownDescription: "Set `true` to force output to upper-case",
	// 		Type:                types.BoolType,
	// 		Optional:            true,
	// 	},
	// 	"title": {
	// 		MarkdownDescription: "Set `true` to force output to title-case",
	// 		Type:                types.BoolType,
	// 		Optional:            true,
	// 	},
	// 	"reverse": {
	// 		MarkdownDescription: "Set `true` to reverse the order of `tags` values specified in `order`",
	// 		Type:                types.BoolType,
	// 		Optional:            true,
	// 	},
	// 	"attributes": {
	// 		MarkdownDescription: "Set `true` to include `attributes` in output",
	// 		Type:                types.BoolType,
	// 		Optional:            true,
	// 	},
	// 	"limit": {
	// 		MarkdownDescription: "Character limit of output. Tail characters are trimmed.",
	// 		Type:                types.NumberType,
	// 		Optional:            true,
	// 	},
	// }

	// id, prefix, dns_name, title
	attributes := map[string]tfsdk.Attribute{
		"enabled": {
			MarkdownDescription: "Set `true` if resources using this context should be created.",
			Type:                types.BoolType,
			Optional:            true,
		},
		"attributes": {
			MarkdownDescription: "TODO",
			Type:                types.ListType{ElemType: types.StringType},
			Optional:            true,
		},
		"attributes_map": {
			MarkdownDescription: "TODO",
			Type:                types.MapType{ElemType: types.StringType},
			Optional:            true,
		},
		// "descriptors": {
		// 	MarkdownDescription: "TODO",
		// 	Attributes:          tfsdk.MapNestedAttributes(descriptorAttributes),
		// 	Optional:            true,
		// },

		// Outputs
		"id": {
			MarkdownDescription: "TODO",
			Type:                types.StringType,
			Computed:            true,
		},
	}

	s := tfsdk.Schema{
		Attributes: attributes,
	}

	return s
}

func (d *ContextDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	attributes := d.getPartialSchema().Attributes

	attributes["context"] = tfsdk.Attribute{
		MarkdownDescription: "TODO",
		Attributes:          tfsdk.SingleNestedAttributes(attributes),
		Optional:            true,
	}

	s := tfsdk.Schema{
		MarkdownDescription: "Context Data Source",
		Attributes:          attributes,
	}

	return s, nil
}

func (d *ContextDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &d)...)

	// s := fwschemadata.Data{
	// 	Description:    fwschemadata.DataDescriptionConfiguration,
	// 	Schema:         getPartialSchema(), //req.Config.Schema,
	// 	TerraformValue: req.Config.Raw,
	// }
	// resp.Diagnostics.Append(s.Get(ctx, &d)...)

	//return reflect.Into(ctx, getPartialSchema.Type(), req.Config.Raw, &d, reflect.Options{}, path.Empty())
	// Read Terraform configuration data into the model
	// parent := ContextDataSource{}
	// resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("context"), &parent)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// diag := resp.State.SetAttribute(ctx, path.Root("id"), "test123")
	// resp.Diagnostics.Append(diag...)

	// // merge enabled
	// var newEnabled types.Bool
	// var parentEnabled types.Bool

	// diag := req.Config.GetAttribute(ctx, path.Root("enabled"), &newEnabled)
	// resp.Diagnostics.Append(diag...)
	// if newEnabled.Null || newEnabled.Unknown {
	// 	newEnabled.Value = true
	// }

	// diag = req.Config.GetAttribute(ctx, path.Root("context").AtName("enabled"), &parentEnabled)
	// resp.Diagnostics.Append(diag...)
	// if parentEnabled.Null || parentEnabled.Unknown {
	// 	parentEnabled.Value = true
	// }

	// enabled := newEnabled.Value && parentEnabled.Value

	// diag = resp.State.SetAttribute(ctx, path.Root("enabled"), enabled)
	// resp.Diagnostics.Append(diag...)

	// // merge attributes
	// attributes := []string{}
	// var newAttributes types.List
	// var parentAttributes types.List

	// diag = req.Config.GetAttribute(ctx, path.Root("attributes"), &newAttributes)
	// resp.Diagnostics.Append(diag...)

	// diag = req.Config.GetAttribute(ctx, path.Root("context").AtName("attributes"), &parentAttributes)
	// resp.Diagnostics.Append(diag...)

	// if !(parentAttributes.Null || parentAttributes.Unknown) {
	// 	diag = parentAttributes.ElementsAs(ctx, &attributes, false)
	// 	resp.Diagnostics.Append(diag...)
	// }

	// if !(newAttributes.Null || newAttributes.Unknown) {
	// 	var appendAttributes []string
	// 	diag = newAttributes.ElementsAs(ctx, &appendAttributes, false)
	// 	resp.Diagnostics.Append(diag...)
	// 	attributes = append(attributes, appendAttributes...)
	// }

	// diag = resp.State.SetAttribute(ctx, path.Root("attributes"), attributes)

	// // merge tags
	// tags := map[string]string{}
	// var newTags types.Map
	// var parentTags types.Map

	// diag = req.Config.GetAttribute(ctx, path.Root("tags"), &newTags)
	// resp.Diagnostics.Append(diag...)

	// diag = req.Config.GetAttribute(ctx, path.Root("context").AtName("tags"), &parentTags)
	// resp.Diagnostics.Append(diag...)

	// if !(parentTags.Null || parentTags.Unknown) {
	// 	diag = parentTags.ElementsAs(ctx, &tags, false)
	// 	resp.Diagnostics.Append(diag...)
	// }

	// if !(newTags.Null || newTags.Unknown) {
	// 	var appendTags map[string]string
	// 	diag = newTags.ElementsAs(ctx, &appendTags, false)
	// 	resp.Diagnostics.Append(diag...)
	// 	for k, v := range appendTags {
	// 		tags[k] = v
	// 	}
	// }

	// diag = resp.State.SetAttribute(ctx, path.Root("tags"), tags)
	// resp.Diagnostics.Append(diag...)

	// // merge descriptors
	// descriptors := map[string]Descriptor{}

	// // default id descriptor
	// defaultDelimiter := "-"
	// defaultLimit := 64
	// defaultAttributes := true
	// defaultLower := true
	// descriptors["id"] = Descriptor{
	// 	Delimiter:  &defaultDelimiter,
	// 	Lower:      &defaultLower,
	// 	Attributes: &defaultAttributes,
	// 	Limit:      &defaultLimit,
	// }

	// var newDescriptors types.Map
	// var parentDescriptors types.Map

	// diag = req.Config.GetAttribute(ctx, path.Root("descriptors"), &newDescriptors)
	// resp.Diagnostics.Append(diag...)

	// diag = req.Config.GetAttribute(ctx, path.Root("context").AtName("descriptors"), &parentDescriptors)
	// resp.Diagnostics.Append(diag...)

	// if !(parentDescriptors.Null || parentDescriptors.Unknown) {
	// 	diag = parentDescriptors.ElementsAs(ctx, &descriptors, false)
	// 	resp.Diagnostics.Append(diag...)
	// }

	// if !(newDescriptors.Null || newDescriptors.Unknown) {
	// 	var appendDescriptors map[string]Descriptor
	// 	diag = newDescriptors.ElementsAs(ctx, &appendDescriptors, false)
	// 	resp.Diagnostics.Append(diag...)
	// 	for k, v := range appendDescriptors {
	// 		descriptors[k] = v
	// 	}
	// }

	// diag = resp.State.SetAttribute(ctx, path.Root("descriptors"), descriptors)
	// resp.Diagnostics.Append(diag...)

	// // outputs from descriptors
	// outputs := map[string]string{}
	// for name, descriptor := range descriptors {
	// 	idParts := []string{}

	// 	for _, tag_key := range descriptor.Order {
	// 		if tag_val, ok := tags[tag_key]; ok {
	// 			idParts = append(idParts, tag_val)
	// 		} else {
	// 			// TODO - warn
	// 		}
	// 	}

	// 	if descriptor.Attributes != nil && *descriptor.Attributes {
	// 		idParts = append(idParts, attributes...)
	// 	}

	// 	if descriptor.Reverse != nil && *descriptor.Reverse {
	// 		for i, j := 0, len(idParts)-1; i < j; i, j = i+1, j-1 {
	// 			idParts[i], idParts[j] = idParts[j], idParts[i]
	// 		}
	// 	}

	// 	id := strings.Join(idParts, *descriptor.Delimiter)

	// 	if descriptor.Upper != nil && *descriptor.Upper {
	// 		id = strings.ToUpper(id)
	// 	}
	// 	if descriptor.Lower != nil && *descriptor.Lower {
	// 		id = strings.ToLower(id)
	// 	}
	// 	if descriptor.Title != nil && *descriptor.Title {
	// 		id = strings.Title(id)
	// 	}

	// 	if descriptor.Limit != nil && len(id) > *descriptor.Limit {
	// 		id = id[:*descriptor.Limit]
	// 		id = strings.TrimRight(id, *descriptor.Delimiter)
	// 	}

	// 	outputs[name] = id
	// }

	// diag = resp.State.SetAttribute(ctx, path.Root("outputs"), outputs)
	// resp.Diagnostics.Append(diag...)

	// diag = resp.State.SetAttribute(ctx, path.Root("id"), outputs["id"])
	// resp.Diagnostics.Append(diag...)

	// // legacy output
	// diag = resp.State.SetAttribute(ctx, path.Root("legacy").AtName("enabled"), enabled)
	// resp.Diagnostics.Append(diag...)

	// diag = resp.State.SetAttribute(ctx, path.Root("legacy").AtName("attributes"), attributes)
	// resp.Diagnostics.Append(diag...)

	// legacyTags := map[string]string{}
	// for tag_key, tag_value := range tags {
	// 	legacyTags[strings.Title(tag_key)] = tag_value
	// }
	// diag = resp.State.SetAttribute(ctx, path.Root("legacy").AtName("tags"), legacyTags)
	// resp.Diagnostics.Append(diag...)

	// diag = resp.State.SetAttribute(ctx, path.Root("legacy").AtName("label_order"), descriptors["id"].Order)
	// resp.Diagnostics.Append(diag...)

	// for _, tag_key := range []string{"namespace", "tenant", "environment", "stage", "name", "domain_name", "dns_name_format", "delimiter", "regex_replace_chars"} {
	// 	if val, ok := tags[tag_key]; ok {
	// 		diag = resp.State.SetAttribute(ctx, path.Root("legacy").AtName(tag_key), val)
	// 		resp.Diagnostics.Append(diag...)
	// 	}
	// }
}
