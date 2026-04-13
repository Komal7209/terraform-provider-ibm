// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package byoc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/cloud-go-sdk/brokerapiv1"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMByocEngineDb2AWS(t *testing.T) {
	var engineID string
	rnd := fmt.Sprintf("tf-byoc-engine-db2-aws-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_byoc_engine." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMByocEngineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMByocEngineDb2AWS(testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMByocEngineExists(name, &engineID),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "engine_type", "db2"),
					resource.TestCheckResourceAttr(name, "version", "11.5"),
					resource.TestCheckResourceAttrSet(name, "subscription_id"),
					resource.TestCheckResourceAttrSet(name, "dataplane_id"),
					resource.TestCheckResourceAttrSet(name, "engine_id"),
					resource.TestCheckResourceAttrSet(name, "status"),
				),
			},
			{
				Config: testAccCheckIBMByocEngineDb2AWSUpdate(testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMByocEngineExists(name, &engineID),
					resource.TestCheckResourceAttr(name, "name", testName+"-updated"),
				),
			},
		},
	})
}

func TestAccIBMByocEngineDb2Azure(t *testing.T) {
	var engineID string
	rnd := fmt.Sprintf("tf-byoc-engine-db2-azure-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_byoc_engine." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMByocEngineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMByocEngineDb2Azure(testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMByocEngineExists(name, &engineID),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "engine_type", "db2"),
					resource.TestCheckResourceAttr(name, "version", "11.5"),
					resource.TestCheckResourceAttrSet(name, "subscription_id"),
					resource.TestCheckResourceAttrSet(name, "dataplane_id"),
					resource.TestCheckResourceAttrSet(name, "engine_id"),
					resource.TestCheckResourceAttrSet(name, "status"),
				),
			},
		},
	})
}

func TestAccIBMByocEngineDb2wh(t *testing.T) {
	var engineID string
	rnd := fmt.Sprintf("tf-byoc-engine-db2wh-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_byoc_engine." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMByocEngineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMByocEngineDb2wh(testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMByocEngineExists(name, &engineID),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "engine_type", "db2wh"),
					resource.TestCheckResourceAttr(name, "version", "11.5"),
				),
			},
		},
	})
}

func TestAccIBMByocEngineNetezza(t *testing.T) {
	var engineID string
	rnd := fmt.Sprintf("tf-byoc-engine-netezza-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_byoc_engine." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMByocEngineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMByocEngineNetezza(testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMByocEngineExists(name, &engineID),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "engine_type", "netezza"),
					resource.TestCheckResourceAttr(name, "version", "11.2"),
				),
			},
		},
	})
}

func testAccCheckIBMByocEngineDestroy(s *terraform.State) error {
	brokerApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrokerApiV1()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_byoc_engine" {
			continue
		}

		parts := strings.Split(rs.Primary.ID, "/")
		if len(parts) != 3 {
			return fmt.Errorf("Invalid ID format: %s", rs.Primary.ID)
		}

		subscriptionID := parts[0]
		dataplaneID := parts[1]
		engineID := parts[2]

		subscriptionUUID := strfmt.UUID(subscriptionID)
		dataplaneUUID := strfmt.UUID(dataplaneID)
		engineUUID := strfmt.UUID(engineID)

		getEngineOptions := &brokerapiv1.GetEngineByIdOptions{
			SubscriptionID: &subscriptionUUID,
			DataplaneID:    &dataplaneUUID,
			EngineID:       &engineUUID,
		}

		_, response, err := brokerApiClient.GetEngineByID(getEngineOptions)

		if err == nil {
			return fmt.Errorf("BYOC Engine still exists: %s", rs.Primary.ID)
		} else {
			if response != nil && response.StatusCode != 404 {
				return fmt.Errorf("[ERROR] Error checking if BYOC Engine (%s) has been destroyed: %s", rs.Primary.ID, err)
			}
		}
	}

	return nil
}

func testAccCheckIBMByocEngineExists(n string, engineID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		brokerApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrokerApiV1()
		if err != nil {
			return err
		}

		parts := strings.Split(rs.Primary.ID, "/")
		if len(parts) != 3 {
			return fmt.Errorf("Invalid ID format: %s", rs.Primary.ID)
		}

		subscriptionID := parts[0]
		dataplaneID := parts[1]
		engID := parts[2]

		subscriptionUUID := strfmt.UUID(subscriptionID)
		dataplaneUUID := strfmt.UUID(dataplaneID)
		engUUID := strfmt.UUID(engID)

		getEngineOptions := &brokerapiv1.GetEngineByIdOptions{
			SubscriptionID: &subscriptionUUID,
			DataplaneID:    &dataplaneUUID,
			EngineID:       &engUUID,
		}

		engine, _, err := brokerApiClient.GetEngineByID(getEngineOptions)
		if err != nil {
			return err
		}

		// Type assert to get the concrete type
		if engineResp, ok := engine.(*brokerapiv1.GetEngineByIdResponse); ok {
			*engineID = string(*engineResp.EngineID)
		} else {
			return fmt.Errorf("unexpected engine response type")
		}
		return nil
	}
}

func testAccCheckIBMByocEngineDb2AWS(name string) string {
	return fmt.Sprintf(`
		resource "ibm_byoc_engine" "%s" {
			subscription_id = "test-subscription-id"
			dataplane_id = "test-dataplane-id-aws"
			name = "%s"
			engine_type = "db2"
			version = "11.5"
			size = "small"
		}
	`, name, name)
}

func testAccCheckIBMByocEngineDb2AWSUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_byoc_engine" "%s" {
			subscription_id = "test-subscription-id"
			dataplane_id = "test-dataplane-id-aws"
			name = "%s-updated"
			engine_type = "db2"
			version = "11.5"
			size = "small"
		}
	`, name, name)
}

func testAccCheckIBMByocEngineDb2Azure(name string) string {
	return fmt.Sprintf(`
		resource "ibm_byoc_engine" "%s" {
			subscription_id = "test-subscription-id"
			dataplane_id = "test-dataplane-id-azure"
			name = "%s"
			engine_type = "db2"
			version = "11.5"
			size = "small"
		}
	`, name, name)
}

func testAccCheckIBMByocEngineDb2wh(name string) string {
	return fmt.Sprintf(`
		resource "ibm_byoc_engine" "%s" {
			subscription_id = "test-subscription-id"
			dataplane_id = "test-dataplane-id"
			name = "%s"
			engine_type = "db2wh"
			version = "11.5"
			size = "medium"
		}
	`, name, name)
}

func testAccCheckIBMByocEngineNetezza(name string) string {
	return fmt.Sprintf(`
		resource "ibm_byoc_engine" "%s" {
			subscription_id = "test-subscription-id"
			dataplane_id = "test-dataplane-id"
			name = "%s"
			engine_type = "netezza"
			version = "11.2"
			size = "large"
		}
	`, name, name)
}
