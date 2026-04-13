// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package byoc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmByocEngineDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmByocEngineDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine", "subscription_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine", "dataplane_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine", "engine_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine", "engine_type"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine", "status"),
				),
			},
		},
	})
}

func testAccCheckIbmByocEngineDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_byoc_engine" "byoc_engine" {
			subscription_id = "test-subscription-id"
			dataplane_id = "test-dataplane-id"
			engine_id = "test-engine-id"
		}
	`)
}
