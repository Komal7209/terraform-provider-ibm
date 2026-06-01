// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package brokerapi_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

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
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "subscription_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "dataplane_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "engines.#"),
				),
			},
		},
	})
}

func TestAccIbmByocEnginesDataSourceAllArgs(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmByocEnginesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "subscription_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "dataplane_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "engines.#"),
					resource.TestCheckResourceAttr("data.ibm_byoc_engines.byoc_engines_data", "engines.#", "1"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "engines.0.engine_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "engines.0.engine_name"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "engines.0.engine_type"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "engines.0.engine_status"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "engines.0.dataplane_id"),
				),
			},
		},
	})
}

func TestAccIbmByocEnginesDataSourceMultipleEngines(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmByocEnginesDataSourceConfigMultiple(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "engines.#"),
					// Check that we have at least 2 engines
					resource.TestCheckResourceAttr("data.ibm_byoc_engines.byoc_engines_data", "engines.#", "2"),
					// Check first engine
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "engines.0.engine_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "engines.0.engine_name"),
					// Check second engine
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "engines.1.engine_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "engines.1.engine_name"),
				),
			},
		},
	})
}

func testAccCheckIbmByocEnginesDataSourceConfigBasic() string {
	return `
		data "ibm_byoc_engines" "byoc_engines_data" {
			subscription_id = "9aafe1f3-9f83-4e31-b99f-c12a119e364e"
			dataplane_id    = "8ccfce03-cdeb-4b48-a45f-a2995a41e859"
		}
	`
}

func testAccCheckIbmByocEnginesDataSourceConfig() string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf(`
		resource "ibm_byoc_engine" "byoc_engine_instance" {
			subscription_id = "9aafe1f3-9f83-4e31-b99f-c12a119e364e"
			dataplane_id    = "8ccfce03-cdeb-4b48-a45f-a2995a41e859"
			engine_name     = "test-engine-%d"
			engine_type     = "db2"
			admin_username  = "admin"
			admin_password  = "{SHA2}R/dfwhLaP217XwTB3IBjoqH3G1oxMA=="
			admin_email     = "admin@example.com"
			storage_units   = 50
			compute_units   = 2
			endpoint_type   = "private"
		}

		data "ibm_byoc_engines" "byoc_engines_data" {
			subscription_id = ibm_byoc_engine.byoc_engine_instance.subscription_id
			dataplane_id    = ibm_byoc_engine.byoc_engine_instance.dataplane_id
			
			depends_on = [ibm_byoc_engine.byoc_engine_instance]
		}
	`, timestamp)
}

func testAccCheckIbmByocEnginesDataSourceConfigMultiple() string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf(`
		resource "ibm_byoc_engine" "byoc_engine_instance_1" {
			subscription_id = "9aafe1f3-9f83-4e31-b99f-c12a119e364e"
			dataplane_id    = "8ccfce03-cdeb-4b48-a45f-a2995a41e859"
			engine_name     = "test-engine-%d-1"
			engine_type     = "db2"
			admin_username  = "admin"
			admin_password  = "SecurePassword123!"
			admin_email     = "admin@example.com"
			storage_units   = 50
			compute_units   = 2
			endpoint_type   = "private"
		}

		resource "ibm_byoc_engine" "byoc_engine_instance_2" {
			subscription_id = "9aafe1f3-9f83-4e31-b99f-c12a119e364e"
			dataplane_id    = "8ccfce03-cdeb-4b48-a45f-a2995a41e859"
			engine_name     = "test-engine-%d-2"
			engine_type     = "db2"
			admin_username  = "admin"
			admin_password  = "SecurePassword123!"
			admin_email     = "admin@example.com"
			storage_units   = 50
			compute_units   = 2
			endpoint_type   = "private"
		}

		data "ibm_byoc_engines" "byoc_engines_data" {
			subscription_id = ibm_byoc_engine.byoc_engine_instance_1.subscription_id
			dataplane_id    = ibm_byoc_engine.byoc_engine_instance_1.dataplane_id
			
			depends_on = [
				ibm_byoc_engine.byoc_engine_instance_1,
				ibm_byoc_engine.byoc_engine_instance_2
			]
		}
	`, timestamp, timestamp)
}
