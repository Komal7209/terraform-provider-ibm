// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package byoc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmByocDataplaneDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmByocDataplaneDataSourceConfigBasic(),
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

func testAccCheckIbmByocDataplaneDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_byoc_dataplane" "byoc_dataplane" {
			subscription_id = "test-subscription-id"
			dataplane_id = "test-dataplane-id"
		}
	`)
}
