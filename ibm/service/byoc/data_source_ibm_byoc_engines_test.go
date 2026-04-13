// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package byoc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmByocEnginesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmByocEnginesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines", "subscription_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines", "dataplane_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines", "engines.#"),
				),
			},
		},
	})
}

func testAccCheckIbmByocEnginesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_byoc_engines" "byoc_engines" {
			subscription_id = "test-subscription-id"
			dataplane_id = "test-dataplane-id"
		}
	`)
}
