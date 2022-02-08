package provider

import (
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy"
)

var providerFactories = map[string]func() (*schema.Provider, error){
	"haproxy": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}

var testClient haproxy.Client

func TestMain(m *testing.M) {
	if os.Getenv("TF_ACC") == "" {
		// short circuit non-acceptance test runs
		os.Exit(m.Run())
	}

	serverAddr := os.Getenv("HAPROXY_SERVER")
	username := os.Getenv("HAPROXY_USERNAME")
	password := os.Getenv("HAPROXY_PASSWORD")
	insecure, _ := strconv.ParseBool(os.Getenv("HAPROXY_INSECURE"))

	testClient := haproxy.NewClient(username, password, serverAddr, insecure)

	err := testClient.TestApiCall()
	if err != nil {
		panic(err)
	}

	resource.TestMain(m)
}

func preCheck(t *testing.T) {
	variables := []string{
		"HAPROXY_SERVER",
		"HAPROXY_USERNAME",
		"HAPROXY_PASSWORD",
	}

	for _, variable := range variables {
		value := os.Getenv(variable)
		if value == "" {
			t.Fatalf("`%s` must be set for acceptance tests!", variable)
		}
	}
}

func importStep(name string, ignore ...string) resource.TestStep {
	step := resource.TestStep{
		ResourceName:      name,
		ImportState:       true,
		ImportStateVerify: true,
	}

	if len(ignore) > 0 {
		step.ImportStateVerifyIgnore = ignore
	}

	return step
}
