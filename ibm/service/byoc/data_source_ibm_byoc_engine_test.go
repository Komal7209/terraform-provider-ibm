// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package byoc_test

import (
	"fmt"
	"testing"
	"time"

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
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "subscription_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "dataplane_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "engine_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "engine_name"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "engine_type"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "engine_status"),
				),
			},
		},
	})
}

func TestAccIbmByocEngineDataSourceAllArgs(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             testAccCheckIbmByocEngineDataSourceConfig(),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "subscription_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "dataplane_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "engine_id"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "engine_name"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "engine_type"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "engine_status"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "engine_status_message"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "engine_short_id"),
					// Version may not be immediately available for newly created engines
					// resource.TestCheckResourceAttrSet("data.ibm_byoc_engine.byoc_engine_data", "version"),
				),
			},
		},
	})
}

func testAccCheckIbmByocEngineDataSourceConfigBasic() string {
	return `
		data "ibm_byoc_engine" "byoc_engine_data" {
			subscription_id = "9aafe1f3-9f83-4e31-b99f-c12a119e364e"
			dataplane_id    = "8ccfce03-cdeb-4b48-a45f-a2995a41e859"
			engine_id       = "e51ad19c-cf6e-499d-9952-d03ca2d31f1c"
		}
	`
}

func testAccCheckIbmByocEngineDataSourceConfig() string {
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
			instance_type   = "Standard_D4s_v5"
			replicas        = 1
			private_link_service_enabled = true
			public_enabled  = false
			oracle_compatibility = false
		}

		data "ibm_byoc_engine" "byoc_engine_data" {
			subscription_id = ibm_byoc_engine.byoc_engine_instance.subscription_id
			dataplane_id    = ibm_byoc_engine.byoc_engine_instance.dataplane_id
			engine_id       = element(split("/", ibm_byoc_engine.byoc_engine_instance.id), 2)
		}
	`, timestamp)
}
