package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceFrontend(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { preCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFrontendConfig("tfacc-frontend1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("haproxy_frontend.test", "name", "tfacc-frontend1"),
				),
			},
		},
	})
}

func testAccFrontendConfig(name string) string {
	return fmt.Sprintf(`
resource "haproxy_frontend" "test" {
	name = "%[1]s"
}	
`, name)
}
