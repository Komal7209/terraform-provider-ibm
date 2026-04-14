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

func TestAccIBMByocDataplaneAWS(t *testing.T) {
	var dataplaneID string
	rnd := fmt.Sprintf("tf-byoc-dp-aws-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_byoc_dataplane." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMByocDataplaneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMByocDataplaneAWS(testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMByocDataplaneExists(name, &dataplaneID),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "region", "us-east-1"),
					resource.TestCheckResourceAttr(name, "cloud_provider", "AWS"),
					resource.TestCheckResourceAttrSet(name, "subscription_id"),
					resource.TestCheckResourceAttrSet(name, "dataplane_id"),
					resource.TestCheckResourceAttrSet(name, "status"),
				),
			},
			{
				Config: testAccCheckIBMByocDataplaneAWSUpdate(testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMByocDataplaneExists(name, &dataplaneID),
					resource.TestCheckResourceAttr(name, "name", testName+"-updated"),
				),
			},
		},
	})
}

func TestAccIBMByocDataplaneAzure(t *testing.T) {
	var dataplaneID string
	rnd := fmt.Sprintf("tf-byoc-dp-azure-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_byoc_dataplane." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMByocDataplaneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMByocDataplaneAzure(testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMByocDataplaneExists(name, &dataplaneID),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "region", "eastus"),
					resource.TestCheckResourceAttr(name, "cloud_provider", "Azure"),
					resource.TestCheckResourceAttrSet(name, "subscription_id"),
					resource.TestCheckResourceAttrSet(name, "dataplane_id"),
					resource.TestCheckResourceAttrSet(name, "status"),
				),
			},
			{
				Config: testAccCheckIBMByocDataplaneAzureUpdate(testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMByocDataplaneExists(name, &dataplaneID),
					resource.TestCheckResourceAttr(name, "name", testName+"-updated"),
				),
			},
		},
	})
}

func testAccCheckIBMByocDataplaneDestroy(s *terraform.State) error {
	brokerApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrokerApiV1()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_byoc_dataplane" {
			continue
		}

		parts := strings.Split(rs.Primary.ID, "/")
		if len(parts) != 2 {
			return fmt.Errorf("Invalid ID format: %s", rs.Primary.ID)
		}

		subscriptionID := parts[0]
		dataplaneID := parts[1]

		subscriptionUUID := strfmt.UUID(subscriptionID)
		dataplaneUUID := strfmt.UUID(dataplaneID)

		getDataplaneOptions := &brokerapiv1.GetDataplaneByIdOptions{
			SubscriptionID: &subscriptionUUID,
			DataplaneID:    &dataplaneUUID,
		}

		_, response, err := brokerApiClient.GetDataplaneByID(getDataplaneOptions)

		if err == nil {
			return fmt.Errorf("BYOC Dataplane still exists: %s", rs.Primary.ID)
		} else {
			if response != nil && response.StatusCode != 404 {
				return fmt.Errorf("[ERROR] Error checking if BYOC Dataplane (%s) has been destroyed: %s", rs.Primary.ID, err)
			}
		}
	}

	return nil
}

func testAccCheckIBMByocDataplaneExists(n string, dataplaneID *string) resource.TestCheckFunc {
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
		if len(parts) != 2 {
			return fmt.Errorf("Invalid ID format: %s", rs.Primary.ID)
		}

		subscriptionID := parts[0]
		dpID := parts[1]

		subscriptionUUID := strfmt.UUID(subscriptionID)
		dpUUID := strfmt.UUID(dpID)

		getDataplaneOptions := &brokerapiv1.GetDataplaneByIdOptions{
			SubscriptionID: &subscriptionUUID,
			DataplaneID:    &dpUUID,
		}

		dataplane, _, err := brokerApiClient.GetDataplaneByID(getDataplaneOptions)
		if err != nil {
			return err
		}

		*dataplaneID = *dataplane.ServiceInstanceID
		return nil
	}
}

func testAccCheckIBMByocDataplaneAWS(name string) string {
	return fmt.Sprintf(`
		resource "ibm_byoc_dataplane" "%s" {
			subscription_id = "f038d818-2358-42f0-a84a-c32656597586"
			dataplane_id = "b2c3d4e5-6789-01bc-defg-234567890abc"
			name = "%s"
			region = "us-east-1"
			cloud_provider = "AWS"
			hyperscaler_account_id = "123456789012"
		}
	`, name, name)
}

func testAccCheckIBMByocDataplaneAWSUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_byoc_dataplane" "%s" {
			subscription_id = "f038d818-2358-42f0-a84a-c32656597586"
			dataplane_id = "b2c3d4e5-6789-01bc-defg-234567890abc"
			name = "%s-updated"
			region = "us-east-1"
			cloud_provider = "AWS"
			hyperscaler_account_id = "123456789012"
		}
	`, name, name)
}

func testAccCheckIBMByocDataplaneAzure(name string) string {
	return fmt.Sprintf(`
		resource "ibm_byoc_dataplane" "%s" {
			subscription_id = "f038d818-2358-42f0-a84a-c32656597586"
			dataplane_id = "a1b2c3d4-5678-90ab-cdef-1234567890ab"
			name = "%s"
			region = "eastus"
			cloud_provider = "Azure"
		}
	`, name, name)
}

// hyperscaler_subscription_id = "f038d818-2358-42f0-a84a-c32656597586"
// hyperscaler_tenant_id = "ffc234dd-5af0-408c-bcd6-add46b47c86f"

func testAccCheckIBMByocDataplaneAzureUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_byoc_dataplane" "%s" {
			subscription_id = "f038d818-2358-42f0-a84a-c32656597586"
			dataplane_id = "a1b2c3d4-5678-90ab-cdef-1234567890ab"
			name = "%s-updated"
			region = "eastus"
			cloud_provider = "Azure"
		}
	`, name, name)
}

// hyperscaler_subscription_id = "f038d818-2358-42f0-a84a-c32656597586"
// hyperscaler_tenant_id = "ffc234dd-5af0-408c-bcd6-add46b47c86f"
