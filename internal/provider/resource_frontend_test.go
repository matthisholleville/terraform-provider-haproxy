package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceFrontend(t *testing.T) {
	var frontendId string
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { preCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFrontendConfig("tfacc-frontend1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("haproxy_frontend.test", "name", "tfacc-frontend1"),

					// extract frontendId for future use
					func(s *terraform.State) error {
						frontendId = s.RootModule().Resources["haproxy_frontend.test"].Primary.Attributes["id"]
						return nil
					},
				),
			},
			importStep("haproxy_frontend.test"),
			{
				ResourceName: "haproxy_frontend.test",
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return frontendId, nil
				},
				ImportState:       true,
				ImportStateVerify: true,
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
