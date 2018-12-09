package anypoint

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"anypoint": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	assertAccEnv(t, "ANYPOINT_USERNAME")
	assertAccEnv(t, "ANYPOINT_PASSWORD")
}

func assertAccEnv(t *testing.T, name string) {
	if v := os.Getenv(name); v == "" {
		t.Fatal(fmt.Sprintf("%s must be set for acceptance tests", name))
	}
}
