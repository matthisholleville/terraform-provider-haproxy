package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceMaps(t *testing.T) {
	//var mapEntrieId string
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { preCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMapEntrie("test", "/metrics", "50", "test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("haproxy_maps.test", "map", "test"),
				),
			},
			{
				ResourceName:            "haproxy_maps.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_sync"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					mapEntrieKey := s.RootModule().Resources["haproxy_maps.test"].Primary.Attributes["id"]
					return fmt.Sprintf("map/%s/entrie/%s", "test", mapEntrieKey), nil
				},
			},
			{
				Config: testAccMapEntrie("test", "super_key", "50", "test-2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("haproxy_maps.test-2", "map", "test"),
				),
			},
			{
				ResourceName:            "haproxy_maps.test-2",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_sync"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					mapEntrieKey := s.RootModule().Resources["haproxy_maps.test-2"].Primary.Attributes["id"]
					return fmt.Sprintf("map/%s/entrie/%s", "test", mapEntrieKey), nil
				},
			},
			{
				Config: testAccMapEntrie("test", "/test1/test2", "50", "test-3"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("haproxy_maps.test-3", "map", "test"),
				),
			},
			{
				ResourceName:            "haproxy_maps.test-3",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_sync"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					mapEntrieKey := s.RootModule().Resources["haproxy_maps.test-3"].Primary.Attributes["id"]
					return fmt.Sprintf("map/%s/entrie/%s", "test", mapEntrieKey), nil
				},
			},
		},
	})
}

func TestCreateMapEntrie_with_invalid_key(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { preCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccMapEntrie("test", "test with bad key", "default_value", "test-bad1"),
				ExpectError: regexp.MustCompile("Cannot insert test with bad key. Space is not allowed."),
			},
			{
				Config:      testAccMapEntrie("test", "test-with-bad: key", "default_value", "test-bad2"),
				ExpectError: regexp.MustCompile("Cannot insert test-with-bad: key. Space is not allowed."),
			},
		},
	})
}

func testAccMapEntrie(mapName string, key string, value string, resourceName string) string {
	return fmt.Sprintf(`
resource "haproxy_maps" "%[4]s" {
	map   = "%[1]s"
	key   = "%[2]s"
	value = "%[3]s"
	}	
`, mapName, key, value, resourceName)
}
