package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceAcl(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { preCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAclConfig("tfacc-acl1", "tfacc-acl1", "hdr(host)", "test.com", "stats", "frontend"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("haproxy_acl.tfacc-acl1", "name", "tfacc-acl1"),
				),
			},
			{
				Config: testAccAclConfig("tfacc-acl2", "tfacc-acl1", "hdr(host)", "test2.com", "stats", "frontend"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("haproxy_acl.tfacc-acl2", "name", "tfacc-acl2"),
				),
			},
		},
	})
}

func testAccAclConfig(resource_name string, name string, criterion string, value string, parent_name string, parent_type string) string {
	return fmt.Sprintf(`
resource "haproxy_acl" "%[1]s" {
	name = "%[2]s"
	criterion = "%[3]s"
	value = "%[4]s"
	parent_name = "%[5]s"
	parent_type = "%[6]s"
}	
`, resource_name, name, criterion, value, parent_name, parent_type)
}
