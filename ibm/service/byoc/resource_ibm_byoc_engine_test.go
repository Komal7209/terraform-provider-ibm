// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package brokerapi_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/brokerapi"
	"github.com/IBM/cloud-go-sdk/brokerapiv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmByocEngineBasic(t *testing.T) {
	var conf brokerapiv1.GetEngineByIdResponse

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmByocEngineDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmByocEngineConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmByocEngineExists("ibm_byoc_engine.byoc_engine_instance", conf),
				),
			},
		},
	})
}

func TestAccIbmByocEngineAllArgs(t *testing.T) {
	var conf brokerapiv1.GetEngineByIdResponse
	availabilityZone := fmt.Sprintf("tf_availability_zone_%d", acctest.RandIntRange(10, 100))
	storageUnits := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	computeUnits := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	engineName := fmt.Sprintf("test-engine-%d", time.Now().Unix())
	engineType := "db2"
	endpointType := "public"
	instanceType := fmt.Sprintf("tf_instance_type_%d", acctest.RandIntRange(10, 100))
	replicas := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	publicEnabled := "true"
	privateLinkServiceEnabled := "true"
	oracleCompatibility := "true"
	plan := fmt.Sprintf("tf_plan_%d", acctest.RandIntRange(10, 100))
	profileName := fmt.Sprintf("tf_profile_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmByocEngineDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmByocEngineConfig(availabilityZone, storageUnits, computeUnits, engineName, engineType, endpointType, instanceType, replicas, publicEnabled, privateLinkServiceEnabled, oracleCompatibility, plan, profileName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmByocEngineExists("ibm_byoc_engine.byoc_engine_instance", conf),
					resource.TestCheckResourceAttr("ibm_byoc_engine.byoc_engine_instance", "availability_zone", availabilityZone),
					resource.TestCheckResourceAttr("ibm_byoc_engine.byoc_engine_instance", "storage_units", storageUnits),
					resource.TestCheckResourceAttr("ibm_byoc_engine.byoc_engine_instance", "compute_units", computeUnits),
					resource.TestCheckResourceAttr("ibm_byoc_engine.byoc_engine_instance", "engine_name", engineName),
					resource.TestCheckResourceAttr("ibm_byoc_engine.byoc_engine_instance", "engine_type", engineType),
					resource.TestCheckResourceAttr("ibm_byoc_engine.byoc_engine_instance", "endpoint_type", endpointType),
					resource.TestCheckResourceAttr("ibm_byoc_engine.byoc_engine_instance", "instance_type", instanceType),
					resource.TestCheckResourceAttr("ibm_byoc_engine.byoc_engine_instance", "replicas", replicas),
					resource.TestCheckResourceAttr("ibm_byoc_engine.byoc_engine_instance", "public_enabled", publicEnabled),
					resource.TestCheckResourceAttr("ibm_byoc_engine.byoc_engine_instance", "private_link_service_enabled", privateLinkServiceEnabled),
					resource.TestCheckResourceAttr("ibm_byoc_engine.byoc_engine_instance", "oracle_compatibility", oracleCompatibility),
					resource.TestCheckResourceAttr("ibm_byoc_engine.byoc_engine_instance", "plan", plan),
					resource.TestCheckResourceAttr("ibm_byoc_engine.byoc_engine_instance", "profile_name", profileName),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_byoc_engine.byoc_engine_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmByocEngineConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_byoc_engine" "byoc_engine_instance" {
			subscription_id = "9aafe1f3-9f83-4e31-b99f-c12a119e364e"
			dataplane_id = "8ccfce03-cdeb-4b48-a45f-a2995a41e859"
		}
	`)
}

func testAccCheckIbmByocEngineConfig(availabilityZone string, storageUnits string, computeUnits string, engineName string, engineType string, endpointType string, instanceType string, replicas string, publicEnabled string, privateLinkServiceEnabled string, oracleCompatibility string, plan string, profileName string) string {
	return fmt.Sprintf(`

		resource "ibm_byoc_engine" "byoc_engine_instance" {
			subscription_id = "9aafe1f3-9f83-4e31-b99f-c12a119e364e"
			dataplane_id = "8ccfce03-cdeb-4b48-a45f-a2995a41e859"
			availability_zone = "%s"
			storage_units = %s
			compute_units = %s
			engine_name = "%s"
			engine_type = "%s"
			endpoint_type = "%s"
			instance_type = "%s"
			replicas = %s
			public_enabled = %s
			private_link_service_enabled = %s
			oracle_compatibility = %s
			plan = "%s"
			profile_name = "%s"
			service_principals = "FIXME"
			subscription_ids = "FIXME"
		}
	`, availabilityZone, storageUnits, computeUnits, engineName, engineType, endpointType, instanceType, replicas, publicEnabled, privateLinkServiceEnabled, oracleCompatibility, plan, profileName)
}

func testAccCheckIbmByocEngineExists(n string, obj brokerapiv1.GetEngineByIdResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		brokerApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrokerApiV1()
		if err != nil {
			return err
		}

		getEngineByIdOptions := &brokerapiv1.GetEngineByIdOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getEngineByIdOptions.SetSubscriptionID(parts[0])
		getEngineByIdOptions.SetDataplaneID(parts[1])
		getEngineByIdOptions.SetEngineID(parts[2])

		createEngineResponseIntf, _, err := brokerApiClient.GetEngineByID(getEngineByIdOptions)
		if err != nil {
			return err
		}

		createEngineResponse := createEngineResponseIntf.(*brokerapiv1.GetEngineByIdResponse)
		obj = *createEngineResponse
		return nil
	}
}

func testAccCheckIbmByocEngineDestroy(s *terraform.State) error {
	brokerApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrokerApiV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_byoc_engine" {
			continue
		}

		getEngineByIdOptions := &brokerapiv1.GetEngineByIdOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getEngineByIdOptions.SetSubscriptionID(parts[0])
		getEngineByIdOptions.SetDataplaneID(parts[1])
		getEngineByIdOptions.SetEngineID(parts[2])

		// Try to find the key
		_, response, err := brokerApiClient.GetEngineByID(getEngineByIdOptions)

		if err == nil {
			return fmt.Errorf("byoc_engine still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for byoc_engine (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmByocEngineEngineUiEndpointToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["public"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(brokerapiv1.EngineUiEndpoint)
	model.Public = core.StringPtr("testString")

	result, err := brokerapi.ResourceIbmByocEngineEngineUiEndpointToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmByocEngineJdbcEndpointToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["private"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(brokerapiv1.JdbcEndpoint)
	model.Private = core.StringPtr("testString")

	result, err := brokerapi.ResourceIbmByocEngineJdbcEndpointToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmByocEngineMapToCreateEngineBaseRequest(t *testing.T) {
	// Checking the result is disabled for this model, because it has a discriminator
	// and there are separate tests for each child model below.
	model := make(map[string]interface{})
	model["engine_type"] = "db2"
	model["engine_name"] = "db2-engine"
	model["admin_username"] = "admin"
	model["admin_password"] = "{SHA2}R/dfwhLaP217XwTB3IBjoqH3G1oxMA=="
	model["admin_email"] = "user@example.com"
	model["storage_units"] = int(50)
	model["compute_units"] = int(2)
	model["private_link_service_enabled"] = true
	model["public_enabled"] = false
	model["oracle_compatibility"] = false
	model["replicas"] = int(1)
	model["instance_type"] = "Standard_D4s_v5"
	model["availability_zone"] = "us-south-1"
	model["plan"] = "db2wh-small"
	model["endpoint_type"] = "private"
	model["profile_name"] = "netezza-profile"
	model["service_principals"] = []interface{}{"service-principal-1"}
	model["subscription_ids"] = []interface{}{"9aafe1f3-9f83-4e31-b99f-c12a119e364e"}

	_, err := brokerapi.ResourceIbmByocEngineMapToCreateEngineBaseRequest(model)
	assert.Nil(t, err)
}

// func TestResourceIbmByocEngineMapToCreateEngineBaseRequestCreateNetezzaEngineRequest(t *testing.T) {
// 	checkResult := func(result *brokerapiv1.CreateEngineBaseRequestCreateNetezzaEngineRequest) {
// 		model := new(brokerapiv1.CreateEngineBaseRequestCreateNetezzaEngineRequest)
// 		model.EngineType = core.StringPtr("netezza-azure")
// 		model.EngineName = core.StringPtr("netezza-engine")
// 		model.AdminUsername = core.StringPtr("admin")
// 		model.AdminPassword = core.StringPtr("{SHA2}R/dfwhLaP217XwTB3IBjoqH3G1oxMA==")
// 		model.AdminEmail = core.StringPtr("user@example.com")
// 		model.AvailabilityZone = core.StringPtr("us-south-1")
// 		model.StorageUnits = core.Int64Ptr(int64(50))
// 		model.ComputeUnits = core.Int64Ptr(int64(2))
// 		model.EndpointType = core.StringPtr("private")
// 		model.ProfileName = core.StringPtr("netezza-profile")
// 		model.ServicePrincipals = []string{"service-principal-1"}
// 		model.SubscriptionIds = []strfmt.UUID{"550e8400-e29b-41d4-a716-446655440000"}

// 		assert.Equal(t, result, model)
// 	}

// 	model := make(map[string]interface{})
// 	model["engine_type"] = "netezza-azure"
// 	model["engine_name"] = "netezza-engine"
// 	model["admin_username"] = "admin"
// 	model["admin_password"] = "{SHA2}R/dfwhLaP217XwTB3IBjoqH3G1oxMA=="
// 	model["admin_email"] = "user@example.com"
// 	model["availability_zone"] = "us-south-1"
// 	model["storage_units"] = int(50)
// 	model["compute_units"] = int(2)
// 	model["endpoint_type"] = "private"
// 	model["profile_name"] = "netezza-profile"
// 	model["service_principals"] = []interface{}{"service-principal-1"}
// 	model["subscription_ids"] = []interface{}{"9aafe1f3-9f83-4e31-b99f-c12a119e364e"}

// 	result, err := brokerapi.ResourceIbmByocEngineMapToCreateEngineBaseRequestCreateNetezzaEngineRequest(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestResourceIbmByocEngineMapToCreateEngineBaseRequestCreateDb2WhEngineRequest(t *testing.T) {
// 	checkResult := func(result *brokerapiv1.CreateEngineBaseRequestCreateDb2WhEngineRequest) {
// 		model := new(brokerapiv1.CreateEngineBaseRequestCreateDb2WhEngineRequest)
// 		model.EngineType = core.StringPtr("db2wh")
// 		model.EngineName = core.StringPtr("db2wh-engine")
// 		model.AdminUsername = core.StringPtr("admin")
// 		model.AdminPassword = core.StringPtr("{SHA2}R/dfwhLaP217XwTB3IBjoqH3G1oxMA==")
// 		model.AdminEmail = core.StringPtr("user@example.com")
// 		model.AvailabilityZone = core.StringPtr("us-south-1")
// 		model.StorageUnits = core.Int64Ptr(int64(100))
// 		model.ComputeUnits = core.Int64Ptr(int64(4))
// 		model.Plan = core.StringPtr("db2wh-small")

// 		assert.Equal(t, result, model)
// 	}

// 	model := make(map[string]interface{})
// 	model["engine_type"] = "db2wh"
// 	model["engine_name"] = "db2wh-engine"
// 	model["admin_username"] = "admin"
// 	model["admin_password"] = "{SHA2}R/dfwhLaP217XwTB3IBjoqH3G1oxMA=="
// 	model["admin_email"] = "user@example.com"
// 	model["availability_zone"] = "us-south-1"
// 	model["storage_units"] = int(100)
// 	model["compute_units"] = int(4)
// 	model["plan"] = "db2wh-small"

// 	result, err := brokerapi.ResourceIbmByocEngineMapToCreateEngineBaseRequestCreateDb2WhEngineRequest(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmByocEngineMapToCreateEngineBaseRequestCreateDb2EngineRequest(t *testing.T) {
	checkResult := func(result *brokerapiv1.CreateEngineBaseRequestCreateDb2EngineRequest) {
		model := new(brokerapiv1.CreateEngineBaseRequestCreateDb2EngineRequest)
		model.EngineType = core.StringPtr("db2")
		model.EngineName = core.StringPtr("db2-engine")
		model.AdminUsername = core.StringPtr("admin")
		model.AdminPassword = core.StringPtr("{SHA2}R/dfwhLaP217XwTB3IBjoqH3G1oxMA==")
		model.AdminEmail = core.StringPtr("user@example.com")
		model.StorageUnits = core.Int64Ptr(int64(50))
		model.ComputeUnits = core.Int64Ptr(int64(2))
		model.PrivateLinkServiceEnabled = core.BoolPtr(true)
		model.PublicEnabled = core.BoolPtr(false)
		model.OracleCompatibility = core.BoolPtr(false)
		model.Replicas = core.Int64Ptr(int64(1))
		model.InstanceType = core.StringPtr("Standard_D4s_v5")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["engine_type"] = "db2"
	model["engine_name"] = "db2-engine"
	model["admin_username"] = "admin"
	model["admin_password"] = "{SHA2}R/dfwhLaP217XwTB3IBjoqH3G1oxMA=="
	model["admin_email"] = "user@example.com"
	model["storage_units"] = int(50)
	model["compute_units"] = int(2)
	model["private_link_service_enabled"] = true
	model["public_enabled"] = false
	model["oracle_compatibility"] = false
	model["replicas"] = int(1)
	model["instance_type"] = "Standard_D4s_v5"

	result, err := brokerapi.ResourceIbmByocEngineMapToCreateEngineBaseRequestCreateDb2EngineRequest(model)
	assert.Nil(t, err)
	checkResult(result)
}
