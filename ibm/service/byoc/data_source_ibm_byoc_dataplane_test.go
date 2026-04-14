// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package byoc_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmByocDataplaneDataSourceBasic(t *testing.T) {
	subscriptionID := os.Getenv("BYOC_SUBSCRIPTION_ID")
	dataplaneID := os.Getenv("BYOC_DATAPLANE_ID")

	if subscriptionID == "" {
		subscriptionID = "f038d818-2358-42f0-a84a-c32656597586" // Default for demo
		t.Logf("[WARN] BYOC_SUBSCRIPTION_ID not set, using default: %s", subscriptionID)
	}
	if dataplaneID == "" {
		dataplaneID = "ffc234dd-5af0-408c-bcd6-add46b47c86f" // Default for demo
		t.Logf("[WARN] BYOC_DATAPLANE_ID not set, using default: %s", dataplaneID)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmByocDataplaneDataSourceConfigBasic(subscriptionID, dataplaneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_byoc_dataplane.byoc_dataplane", "subscription_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_dataplane.byoc_dataplane", "dataplane_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_dataplane.byoc_dataplane", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_dataplane.byoc_dataplane", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_dataplane.byoc_dataplane", "region"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_dataplane.byoc_dataplane", "cloud_provider"),
				),
			},
		},
	})
}

func testAccCheckIbmByocDataplaneDataSourceConfigBasic(subscriptionID, dataplaneID string) string {
	return fmt.Sprintf(`
		data "ibm_byoc_dataplane" "byoc_dataplane" {
			subscription_id = "%s"
			dataplane_id = "%s"
		}
	`, subscriptionID, dataplaneID)
}
