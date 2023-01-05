package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExampleDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testEnabled,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.context.default", "enabled", "true"),
					// resource.TestCheckResourceAttr("data.context.this", "enabled", "false"),
					// resource.TestCheckResourceAttr("data.context.that", "enabled", "true"),
					// resource.TestCheckResourceAttr("data.context.foo", "enabled", "false"),
					// resource.TestCheckResourceAttr("data.context.bar", "enabled", "false"),
				),
			},
		},
	})
}

const testEnabled = `
data "context" "default" {
  enabled = true
}
`

// data "context" "this" {
//   enabled = false
// }

// data "context" "that" {
//   enabled = true
// }

// data "context" "foo" {
// 	context = data.context.this
// 	enabled = true
// }

// data "context" "bar" {
// 	context = data.context.that
//   enabled = false
// }
// `
