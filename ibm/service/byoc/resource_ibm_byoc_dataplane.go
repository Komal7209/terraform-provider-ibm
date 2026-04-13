// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package byoc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/cloud-go-sdk/brokerapiv1"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMByocDataplane() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMByocDataplaneCreate,
		ReadContext:   resourceIBMByocDataplaneRead,
		UpdateContext: resourceIBMByocDataplaneUpdate,
		DeleteContext: resourceIBMByocDataplaneDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subscription ID",
			},
			"dataplane_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_byoc_dataplane", "dataplane_id"),
				Description:  "The dataplane ID (must be a valid UUID)",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the dataplane",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The region where the dataplane will be deployed",
			},
			"cloud_provider": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The cloud provider (AWS, Azure, etc.)",
			},
			"hyperscaler_account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The hyperscaler account ID (AWS account ID)",
			},
			"hyperscaler_subscription_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Azure subscription ID (required for Azure)",
			},
			"hyperscaler_tenant_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Azure tenant ID (required for Azure)",
			},
			"vpc_cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "VPC CIDR block (required for AWS + Netezza)",
			},
			"vpc_network_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "VPC network type: public or private (required for AWS + Netezza)",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v != "" && v != "public" && v != "private" {
						errs = append(errs, fmt.Errorf("%q must be either 'public' or 'private', got: %s", key, v))
					}
					return
				},
			},
			"private_subnet_cidrs": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Private subnet CIDR blocks (required for AWS + Netezza)",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"public_subnet_cidrs": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Public subnet CIDR blocks (required for AWS + Netezza)",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the dataplane",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation timestamp",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The last update timestamp",
			},
		},
	}
}

func ResourceIBMByocDataplaneValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "dataplane_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		})

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_byoc_dataplane", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMByocDataplaneCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	subscriptionID := d.Get("subscription_id").(string)
	dataplaneID := d.Get("dataplane_id").(string)
	cloudProvider := d.Get("cloud_provider").(string)

	log.Printf("[DEBUG] Creating dataplane with ID: %s for cloud provider: %s", dataplaneID, cloudProvider)

	// Hyperscaler-specific flow (AWS and Azure have different sequences)
	switch cloudProvider {
	case "AWS":
		// AWS Flow: updateAccount -> prereqs -> create (verify happens internally)

		// Step 1: Update AWS hyperscaler account
		if hyperscalerAccountID, ok := d.GetOk("hyperscaler_account_id"); ok {
			log.Printf("[DEBUG] AWS Step 1: Updating AWS hyperscaler account")
			if err := updateHyperscalerAccountAWS(context, brokerApiClient, subscriptionID, dataplaneID, hyperscalerAccountID.(string)); err != nil {
				return diag.FromErr(err)
			}
		}

		// Step 2: Check prerequisites
		log.Printf("[DEBUG] AWS Step 2: Checking prerequisites")
		if err := checkPrerequisites(context, brokerApiClient, subscriptionID, dataplaneID); err != nil {
			return diag.FromErr(err)
		}

		// Step 3: Create dataplane (verify happens internally for AWS)
		log.Printf("[DEBUG] AWS Step 3: Creating dataplane (verify happens internally)")
		if err := createOrUpdateDataplane(context, brokerApiClient, subscriptionID, dataplaneID, d, false); err != nil {
			return diag.FromErr(err)
		}

	case "Azure":
		// Azure Flow: updateAccount -> prereqs -> verify -> create

		// Step 1: Update Azure hyperscaler account (returns AppConsentUrl)
		if hyperscalerSubscriptionID, ok := d.GetOk("hyperscaler_subscription_id"); ok {
			if hyperscalerTenantID, ok := d.GetOk("hyperscaler_tenant_id"); ok {
				log.Printf("[DEBUG] Azure Step 1: Updating Azure hyperscaler account")
				if err := updateHyperscalerAccountAzure(context, brokerApiClient, subscriptionID, dataplaneID,
					hyperscalerSubscriptionID.(string), hyperscalerTenantID.(string)); err != nil {
					return diag.FromErr(err)
				}
				log.Printf("[INFO] Azure App Registration created. User must grant consent via Azure portal.")
			}
		}

		// Step 2: Check prerequisites
		log.Printf("[DEBUG] Azure Step 2: Checking prerequisites")
		if err := checkPrerequisites(context, brokerApiClient, subscriptionID, dataplaneID); err != nil {
			return diag.FromErr(err)
		}

		// Step 3: Verify dataplane (MUST be explicit for Azure before createDataplane)
		log.Printf("[DEBUG] Azure Step 3: Verifying dataplane")
		if err := verifyDataplane(context, brokerApiClient, subscriptionID, dataplaneID); err != nil {
			return diag.FromErr(err)
		}

		// Step 4: Create dataplane (skips verify since already done)
		log.Printf("[DEBUG] Azure Step 4: Creating dataplane (verification already completed)")
		if err := createOrUpdateDataplane(context, brokerApiClient, subscriptionID, dataplaneID, d, false); err != nil {
			return diag.FromErr(err)
		}
	}

	// Set the resource ID
	d.SetId(fmt.Sprintf("%s/%s", subscriptionID, dataplaneID))
	log.Printf("[DEBUG] Dataplane creation initiated with ID: %s", dataplaneID)

	// Wait for dataplane to be ready
	log.Printf("[DEBUG] Waiting for dataplane to be ready")
	subscriptionUUID := strfmt.UUID(subscriptionID)
	dataplaneUUID := strfmt.UUID(dataplaneID)

	stateConf := &resource.StateChangeConf{
		Pending: []string{"pending", "provisioning", "creating"},
		Target:  []string{"active", "ready"},
		Refresh: func() (interface{}, string, error) {
			dataplane, _, err := brokerApiClient.GetDataplaneByIDWithContext(context, &brokerapiv1.GetDataplaneByIdOptions{
				SubscriptionID: &subscriptionUUID,
				DataplaneID:    &dataplaneUUID,
			})
			if err != nil {
				return nil, "", err
			}
			status := "unknown"
			if dataplane.ResourceProvisioningStatus != nil {
				status = *dataplane.ResourceProvisioningStatus
			}
			log.Printf("[DEBUG] Dataplane status: %s", status)
			return dataplane, status, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(context)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error waiting for dataplane (%s) to be ready: %s", dataplaneID, err))
	}

	// Read the final dataplane state
	return resourceIBMByocDataplaneRead(context, d, meta)
}

// checkPrerequisites checks prerequisites using UpdateDataplane with action=prereqs
func checkPrerequisites(ctx context.Context, client *brokerapiv1.BrokerApiV1, subscriptionID, dataplaneID string) error {
	subscriptionUUID := strfmt.UUID(subscriptionID)
	dataplaneUUID := strfmt.UUID(dataplaneID)
	action := "prereqs"

	updateOptions := &brokerapiv1.UpdateDataplaneOptions{
		SubscriptionID:          &subscriptionUUID,
		DataplaneID:             &dataplaneUUID,
		Action:                  &action,
		PutDataplaneRequestBody: &brokerapiv1.PutDataplaneRequestBody{},
	}

	_, response, err := client.UpdateDataplaneWithContext(ctx, updateOptions)
	if err != nil {
		responseBody := ""
		statusCode := 0
		if response != nil {
			statusCode = response.StatusCode
			if rawResult := response.GetRawResult(); rawResult != nil {
				responseBody = string(rawResult)
			}
		}
		log.Printf("[DEBUG] CheckPrerequisites (action=prereqs) failed: %s\nStatus: %d\nResponse: %s", err, statusCode, responseBody)
		return fmt.Errorf("checkPrerequisites failed: %s (Status: %d, Response: %s)", err, statusCode, responseBody)
	}
	log.Printf("[DEBUG] Prerequisites check completed successfully")
	return nil
}

// verifyDataplane verifies dataplane using UpdateDataplane with action=verify
func verifyDataplane(ctx context.Context, client *brokerapiv1.BrokerApiV1, subscriptionID, dataplaneID string) error {
	subscriptionUUID := strfmt.UUID(subscriptionID)
	dataplaneUUID := strfmt.UUID(dataplaneID)
	action := "verify"

	updateOptions := &brokerapiv1.UpdateDataplaneOptions{
		SubscriptionID:          &subscriptionUUID,
		DataplaneID:             &dataplaneUUID,
		Action:                  &action,
		PutDataplaneRequestBody: &brokerapiv1.PutDataplaneRequestBody{},
	}

	_, response, err := client.UpdateDataplaneWithContext(ctx, updateOptions)
	if err != nil {
		log.Printf("[DEBUG] VerifyDataplane (action=verify) failed: %s\n%s", err, response)
		return fmt.Errorf("verifyDataplane failed: %s", err)
	}
	log.Printf("[DEBUG] Dataplane verification completed successfully")
	return nil
}

// updateHyperscalerAccountAWS updates AWS account credentials using action=updateHyperscalerAccount
func updateHyperscalerAccountAWS(ctx context.Context, client *brokerapiv1.BrokerApiV1, subscriptionID, dataplaneID, accountID string) error {
	subscriptionUUID := strfmt.UUID(subscriptionID)
	dataplaneUUID := strfmt.UUID(dataplaneID)
	action := "updateHyperscalerAccount"

	updateOptions := &brokerapiv1.UpdateDataplaneOptions{
		SubscriptionID: &subscriptionUUID,
		DataplaneID:    &dataplaneUUID,
		Action:         &action,
		PutDataplaneRequestBody: &brokerapiv1.PutDataplaneRequestBody{
			HyperscalerAccountID: &accountID,
		},
	}

	_, response, err := client.UpdateDataplaneWithContext(ctx, updateOptions)
	if err != nil {
		log.Printf("[DEBUG] UpdateHyperscalerAccount (AWS) failed: %s\n%s", err, response)
		return fmt.Errorf("UpdateHyperscalerAccount (AWS) failed: %s", err)
	}
	log.Printf("[DEBUG] AWS hyperscaler account updated successfully")
	return nil
}

// updateHyperscalerAccountAzure updates Azure account credentials using action=updateHyperscalerAccount
func updateHyperscalerAccountAzure(ctx context.Context, client *brokerapiv1.BrokerApiV1, subscriptionID, dataplaneID, azureSubscriptionID, tenantID string) error {
	subscriptionUUID := strfmt.UUID(subscriptionID)
	dataplaneUUID := strfmt.UUID(dataplaneID)
	action := "updateHyperscalerAccount"

	updateOptions := &brokerapiv1.UpdateDataplaneOptions{
		SubscriptionID: &subscriptionUUID,
		DataplaneID:    &dataplaneUUID,
		Action:         &action,
		PutDataplaneRequestBody: &brokerapiv1.PutDataplaneRequestBody{
			HyperscalerSubscriptionID: &azureSubscriptionID,
			HyperscalerTenantID:       &tenantID,
		},
	}

	_, response, err := client.UpdateDataplaneWithContext(ctx, updateOptions)
	if err != nil {
		log.Printf("[DEBUG] UpdateHyperscalerAccount (Azure) failed: %s\n%s", err, response)
		return fmt.Errorf("UpdateHyperscalerAccount (Azure) failed: %s", err)
	}
	log.Printf("[DEBUG] Azure hyperscaler account updated successfully")
	return nil
}

// createOrUpdateDataplane creates or updates the dataplane (no action parameter)
// This updates the DB and makes a gRPC call to WFC to create/update actual dataplane resources
func createOrUpdateDataplane(ctx context.Context, client *brokerapiv1.BrokerApiV1, subscriptionID, dataplaneID string, d *schema.ResourceData, isUpdate bool) error {
	subscriptionUUID := strfmt.UUID(subscriptionID)
	dataplaneUUID := strfmt.UUID(dataplaneID)

	// No action parameter = updates DB and makes gRPC call to WFC
	updateOptions := &brokerapiv1.UpdateDataplaneOptions{
		SubscriptionID: &subscriptionUUID,
		DataplaneID:    &dataplaneUUID,
	}

	// Build the request body with network configuration
	requestBody := &brokerapiv1.PutDataplaneRequestBody{}

	cloudProvider := d.Get("cloud_provider").(string)

	// AWS-specific configuration
	if cloudProvider == "AWS" {
		if vpcCidr, ok := d.GetOk("vpc_cidr"); ok {
			requestBody.VpcCidr = flex.PtrToString(vpcCidr.(string))
		}
		if vpcNetworkType, ok := d.GetOk("vpc_network_type"); ok {
			requestBody.VpcNetworkType = flex.PtrToString(vpcNetworkType.(string))
		}
		if privateSubnetCidrs, ok := d.GetOk("private_subnet_cidrs"); ok {
			cidrs := flex.ExpandStringList(privateSubnetCidrs.([]interface{}))
			requestBody.PrivateSubnetCidrs = cidrs
		}
		if publicSubnetCidrs, ok := d.GetOk("public_subnet_cidrs"); ok {
			cidrs := flex.ExpandStringList(publicSubnetCidrs.([]interface{}))
			requestBody.PublicSubnetCidrs = cidrs
		}
	} else if cloudProvider == "Azure" {
		// Azure-specific configuration
		// Azure does not require network configuration in the request body
		// Azure hyperscaler credentials (subscription_id and tenant_id) are handled
		// separately in Step 2 via updateHyperscalerAccountAzure function
		log.Printf("[DEBUG] Azure dataplane - no network configuration needed in request body")
	}
	// For other cloud providers (GCP, etc.), request body remains empty

	updateOptions.PutDataplaneRequestBody = requestBody

	operation := "Create"
	if isUpdate {
		operation = "Update"
	}

	_, response, err := client.UpdateDataplaneWithContext(ctx, updateOptions)
	if err != nil {
		log.Printf("[DEBUG] %s dataplane failed: %s\n%s", operation, err, response)
		return fmt.Errorf("%s dataplane failed (gRPC call to WFC): %s", operation, err)
	}
	log.Printf("[DEBUG] %s dataplane successful (gRPC call to WFC initiated)", operation)
	return nil
}

func resourceIBMByocDataplaneRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	subscriptionID := d.Get("subscription_id").(string)
	dataplaneID := d.Get("dataplane_id").(string)

	subscriptionUUID := strfmt.UUID(subscriptionID)
	dataplaneUUID := strfmt.UUID(dataplaneID)

	getDataplaneOptions := &brokerapiv1.GetDataplaneByIdOptions{
		SubscriptionID: &subscriptionUUID,
		DataplaneID:    &dataplaneUUID,
	}

	dataplane, response, err := brokerApiClient.GetDataplaneByIDWithContext(context, getDataplaneOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			log.Printf("[DEBUG] Dataplane %s not found, removing from state", dataplaneID)
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetDataplaneByIDWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetDataplaneByIDWithContext failed %s\n%s", err, response))
	}

	// Set all computed and optional fields from the response
	if dataplane.ServiceInstanceName != nil {
		d.Set("name", *dataplane.ServiceInstanceName)
	}
	if dataplane.ResourceProvisioningStatus != nil {
		d.Set("status", *dataplane.ResourceProvisioningStatus)
	}
	if dataplane.HyperscalerRegion != nil {
		d.Set("region", *dataplane.HyperscalerRegion)
	}
	if dataplane.Hyperscaler != nil {
		d.Set("cloud_provider", *dataplane.Hyperscaler)
	}
	if dataplane.CreatedAt != nil {
		d.Set("created_at", flex.DateTimeToString(dataplane.CreatedAt))
	}
	if dataplane.LastUpdatedAt != nil {
		d.Set("updated_at", flex.DateTimeToString(dataplane.LastUpdatedAt))
	}

	return nil
}

func resourceIBMByocDataplaneUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	subscriptionID := d.Get("subscription_id").(string)
	dataplaneID := d.Get("dataplane_id").(string)
	cloudProvider := d.Get("cloud_provider").(string)

	// Check if hyperscaler account credentials changed
	if d.HasChange("hyperscaler_account_id") && cloudProvider == "AWS" {
		if hyperscalerAccountID, ok := d.GetOk("hyperscaler_account_id"); ok {
			log.Printf("[DEBUG] Updating AWS hyperscaler account")
			if err := updateHyperscalerAccountAWS(context, brokerApiClient, subscriptionID, dataplaneID, hyperscalerAccountID.(string)); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if (d.HasChange("hyperscaler_subscription_id") || d.HasChange("hyperscaler_tenant_id")) && cloudProvider == "Azure" {
		if hyperscalerSubscriptionID, ok := d.GetOk("hyperscaler_subscription_id"); ok {
			if hyperscalerTenantID, ok := d.GetOk("hyperscaler_tenant_id"); ok {
				log.Printf("[DEBUG] Updating Azure hyperscaler account")
				if err := updateHyperscalerAccountAzure(context, brokerApiClient, subscriptionID, dataplaneID,
					hyperscalerSubscriptionID.(string), hyperscalerTenantID.(string)); err != nil {
					return diag.FromErr(err)
				}
			}
		}
	}

	// Check if any network configuration fields have changed
	if d.HasChanges("vpc_cidr", "vpc_network_type", "private_subnet_cidrs", "public_subnet_cidrs") {
		log.Printf("[DEBUG] Updating dataplane network configuration (will trigger gRPC call to WFC)")
		if err := createOrUpdateDataplane(context, brokerApiClient, subscriptionID, dataplaneID, d, true); err != nil {
			return diag.FromErr(err)
		}

		// Wait for update to complete
		log.Printf("[DEBUG] Waiting for dataplane update to complete")
		subscriptionUUID := strfmt.UUID(subscriptionID)
		dataplaneUUID := strfmt.UUID(dataplaneID)

		stateConf := &resource.StateChangeConf{
			Pending: []string{"updating", "provisioning"},
			Target:  []string{"active", "ready"},
			Refresh: func() (interface{}, string, error) {
				dataplane, _, err := brokerApiClient.GetDataplaneByIDWithContext(context, &brokerapiv1.GetDataplaneByIdOptions{
					SubscriptionID: &subscriptionUUID,
					DataplaneID:    &dataplaneUUID,
				})
				if err != nil {
					return nil, "", err
				}
				status := "unknown"
				if dataplane.ResourceProvisioningStatus != nil {
					status = *dataplane.ResourceProvisioningStatus
				}
				log.Printf("[DEBUG] Dataplane status: %s", status)
				return dataplane, status, nil
			},
			Timeout:    d.Timeout(schema.TimeoutUpdate),
			Delay:      10 * time.Second,
			MinTimeout: 5 * time.Second,
		}

		_, err = stateConf.WaitForStateContext(context)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error waiting for dataplane (%s) update to complete: %s", dataplaneID, err))
		}
	}

	// Read the updated dataplane state
	return resourceIBMByocDataplaneRead(context, d, meta)
}

func resourceIBMByocDataplaneDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	subscriptionID := d.Get("subscription_id").(string)
	dataplaneID := d.Get("dataplane_id").(string)

	subscriptionUUID := strfmt.UUID(subscriptionID)
	dataplaneUUID := strfmt.UUID(dataplaneID)

	deleteDataplaneOptions := &brokerapiv1.DeleteDataplaneOptions{
		SubscriptionID: &subscriptionUUID,
		DataplaneID:    &dataplaneUUID,
	}

	_, response, err := brokerApiClient.DeleteDataplaneWithContext(context, deleteDataplaneOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			// Dataplane already deleted
			log.Printf("[DEBUG] Dataplane %s already deleted", dataplaneID)
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] DeleteDataplaneWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteDataplaneWithContext failed %s\n%s", err, response))
	}

	// Wait for deletion to complete
	log.Printf("[DEBUG] Waiting for dataplane deletion to complete")

	subscriptionUUIDDel := strfmt.UUID(subscriptionID)
	dataplaneUUIDDel := strfmt.UUID(dataplaneID)

	stateConf := &resource.StateChangeConf{
		Pending: []string{"deleting", "active", "ready"},
		Target:  []string{"deleted"},
		Refresh: func() (interface{}, string, error) {
			dataplane, response, err := brokerApiClient.GetDataplaneByIDWithContext(context, &brokerapiv1.GetDataplaneByIdOptions{
				SubscriptionID: &subscriptionUUIDDel,
				DataplaneID:    &dataplaneUUIDDel,
			})
			if err != nil {
				if response != nil && response.StatusCode == 404 {
					return dataplane, "deleted", nil
				}
				return nil, "", err
			}
			status := "unknown"
			if dataplane.ResourceProvisioningStatus != nil {
				status = *dataplane.ResourceProvisioningStatus
			}
			log.Printf("[DEBUG] Dataplane status: %s", status)
			return dataplane, status, nil
		},
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(context)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error waiting for dataplane (%s) to be deleted: %s", dataplaneID, err))
	}

	log.Printf("[DEBUG] Deleted dataplane %s", dataplaneID)
	d.SetId("")
	return nil
}
