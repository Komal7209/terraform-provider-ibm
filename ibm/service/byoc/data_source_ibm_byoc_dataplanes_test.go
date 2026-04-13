// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package byoc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmByocDataplanesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmByocDataplanesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_byoc_dataplanes.byoc_dataplanes", "dataplanes.#"),
				),
			},
		},
	})
}

func TestAccIbmByocDataplanesDataSourceWithSubscription(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmByocDataplanesDataSourceConfigWithSubscription(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_byoc_dataplanes.byoc_dataplanes", "subscription_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_dataplanes.byoc_dataplanes", "dataplanes.#"),
				),
			},
		},
	})
}

func testAccCheckIbmByocDataplanesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_byoc_dataplanes" "byoc_dataplanes" {
		}
	`)
}

func testAccCheckIbmByocDataplanesDataSourceConfigWithSubscription() string {
	return fmt.Sprintf(`
		data "ibm_byoc_dataplanes" "byoc_dataplanes" {
			subscription_id = "test-subscription-id"
		}
	`)
}
