// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package byoc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

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
	// Set environment variable to skip deletion during this test
	// Engines take too long to provision and cannot be deleted while IN_PROGRESS
	t.Setenv("BYOC_SKIP_DELETE", "true")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		// Use a custom CheckDestroy that always succeeds since we're skipping deletion
		CheckDestroy: func(s *terraform.State) error {
			t.Log("Skipping destroy check - BYOC_SKIP_DELETE is set")
			t.Log("Manual cleanup required for test engines after provisioning completes")
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmByocEnginesDataSourceConfig(),
				// Allow non-empty plan since engine is still provisioning and computed fields are being populated
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "subscription_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engines.byoc_engines_data", "dataplane_id"),
					// Just verify that the data source returns some engines
					// The actual engine fields may vary depending on engine type and state
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["data.ibm_byoc_engines.byoc_engines_data"]
						if !ok {
							return fmt.Errorf("Not found: data.ibm_byoc_engines.byoc_engines_data")
						}

						enginesCount := rs.Primary.Attributes["engines.#"]
						if enginesCount == "" || enginesCount == "0" {
							return fmt.Errorf("No engines found in dataplane")
						}

						return nil
					},
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
			private_link_service_enabled = true
			public_enabled  = false
			replicas        = 1
			instance_type   = "Standard_D4s_v5"
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
			instance_type   = "Standard_D4s_v5"
			replicas        = 1
			private_link_service_enabled = true
			public_enabled  = false
			oracle_compatibility = false
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
			instance_type   = "Standard_D4s_v5"
			replicas        = 1
			private_link_service_enabled = true
			public_enabled  = false
			oracle_compatibility = false
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
