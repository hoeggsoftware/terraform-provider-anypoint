package anypoint

import (
	"context"
	"fmt"
	"github.com/rhoegg/go-anypoint/anypointplatform"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAnypointBusinessGroup_Basic(t *testing.T) {
	var bg anypointplatform.BusinessGroup
	rInt := acctest.RandInt()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAnypointBusinessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckAnypointBusinessGroupConfig_basic, rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAnypointBusinessGroupExists("anypoint_business_group.testgroup", &bg),
					resource.TestCheckResourceAttr(
						"anypoint_business_group.testgroup", "name", fmt.Sprintf("TestGroup%d", rInt)),
				),
			},
		},
	})
}

func testAccCheckAnypointBusinessGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*anypointplatform.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "anypoint_business_group" {
			continue
		}

		// Try to find the key
		_, _, err := client.BusinessGroup.Get(context.Background(), rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Business Group still exists")
		}
	}

	return nil
}

func testAccCheckAnypointBusinessGroupExists(n string, bg *anypointplatform.BusinessGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		client := testAccProvider.Meta().(*anypointplatform.Client)

		foundBg, _, err := client.BusinessGroup.Get(context.Background(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if foundBg.ID != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		*bg = *foundBg

		return nil
	}
}

const testAccCheckAnypointBusinessGroupConfig_basic = `
resource "anypoint_business_group" "testgroup" {
    name = "TestGroup%d"
}`
