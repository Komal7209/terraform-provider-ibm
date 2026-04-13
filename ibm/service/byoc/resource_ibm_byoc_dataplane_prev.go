// // Copyright IBM Corp. 2024 All Rights Reserved.
// // Licensed under the Mozilla Public License v2.0

package byoc

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
// 	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
// 	"github.com/IBM/cloud-go-sdk/brokerapiv1"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// )

// func ResourceIBMByocDataplane() *schema.Resource {
// 	return &schema.Resource{
// 		CreateContext: resourceIBMByocDataplaneCreate,
// 		ReadContext:   resourceIBMByocDataplaneRead,
// 		UpdateContext: resourceIBMByocDataplaneUpdate,
// 		DeleteContext: resourceIBMByocDataplaneDelete,
// 		Importer:      &schema.ResourceImporter{},

// 		Timeouts: &schema.ResourceTimeout{
// 			Create: schema.DefaultTimeout(30 * time.Minute),
// 			Update: schema.DefaultTimeout(30 * time.Minute),
// 			Delete: schema.DefaultTimeout(30 * time.Minute),
// 		},

// 		Schema: map[string]*schema.Schema{
// 			"subscription_id": {
// 				Type:        schema.TypeString,
// 				Required:    true,
// 				ForceNew:    true,
// 				Description: "The subscription ID",
// 			},
// 			"dataplane_id": {
// 				Type:        schema.TypeString,
// 				Computed:    true,
// 				Description: "The dataplane ID",
// 			},
// 			"name": {
// 				Type:        schema.TypeString,
// 				Required:    true,
// 				Description: "The name of the dataplane",
// 			},
// 			"region": {
// 				Type:        schema.TypeString,
// 				Required:    true,
// 				ForceNew:    true,
// 				Description: "The region where the dataplane will be deployed",
// 			},
// 			"cloud_provider": {
// 				Type:        schema.TypeString,
// 				Required:    true,
// 				ForceNew:    true,
// 				Description: "The cloud provider (AWS, Azure, etc.)",
// 			},
// 			"status": {
// 				Type:        schema.TypeString,
// 				Computed:    true,
// 				Description: "The status of the dataplane",
// 			},
// 			"created_at": {
// 				Type:        schema.TypeString,
// 				Computed:    true,
// 				Description: "The creation timestamp",
// 			},
// 			"updated_at": {
// 				Type:        schema.TypeString,
// 				Computed:    true,
// 				Description: "The last update timestamp",
// 			},
// 		},
// 	}
// }

// func resourceIBMByocDataplaneCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	subscriptionID := d.Get("subscription_id").(string)
// 	dataplaneID := d.Get("dataplane_id").(string)

// 	updateDataplaneOptions := &brokerapiv1.UpdateDataplaneOptions{
// 		SubscriptionID: &subscriptionID,
// 		DataplaneID:    &dataplaneID,
// 	}

// 	// Build the request body with supported fields
// 	requestBody := &brokerapiv1.PutDataplaneRequestBody{}

// 	// Set supported fields from the schema
// 	if vpcCidr, ok := d.GetOk("vpc_cidr"); ok {
// 		requestBody.VpcCidr = flex.PtrToString(vpcCidr.(string))
// 	}
// 	if vpcNetworkType, ok := d.GetOk("vpc_network_type"); ok {
// 		requestBody.VpcNetworkType = flex.PtrToString(vpcNetworkType.(string))
// 	}

// 	updateDataplaneOptions.PutDataplaneRequestBody = requestBody

// 	_, response, err := brokerApiClient.UpdateDataplaneWithContext(context, updateDataplaneOptions)
// 	if err != nil {
// 		log.Printf("[DEBUG] UpdateDataplaneWithContext failed %s\n%s", err, response)
// 		return diag.FromErr(fmt.Errorf("UpdateDataplaneWithContext failed %s\n%s", err, response))
// 	}

// 	d.SetId(fmt.Sprintf("%s/%s", subscriptionID, dataplaneID))

// 	return resourceIBMByocDataplaneRead(context, d, meta)
// }

// func resourceIBMByocDataplaneRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	subscriptionID := d.Get("subscription_id").(string)
// 	dataplaneID := d.Get("dataplane_id").(string)

// 	getDataplaneOptions := &brokerapiv1.GetDataplaneByIdOptions{
// 		SubscriptionID: &subscriptionID,
// 		DataplaneID:    &dataplaneID,
// 	}

// 	dataplane, response, err := brokerApiClient.GetDataplaneByIDWithContext(context, getDataplaneOptions)
// 	if err != nil {
// 		if response != nil && response.StatusCode == 404 {
// 			d.SetId("")
// 			return nil
// 		}
// 		log.Printf("[DEBUG] GetDataplaneWithContext failed %s\n%s", err, response)
// 		return diag.FromErr(fmt.Errorf("GetDataplaneWithContext failed %s\n%s", err, response))
// 	}

// 	if dataplane.ServiceInstanceName != nil {
// 		d.Set("name", *dataplane.ServiceInstanceName)
// 	}
// 	if dataplane.ResourceProvisioningStatus != nil {
// 		d.Set("status", *dataplane.ResourceProvisioningStatus)
// 	}
// 	if dataplane.HyperscalerRegion != nil {
// 		d.Set("region", *dataplane.HyperscalerRegion)
// 	}
// 	if dataplane.Hyperscaler != nil {
// 		d.Set("cloud_provider", *dataplane.Hyperscaler)
// 	}
// 	if dataplane.CreatedAt != nil {
// 		d.Set("created_at", flex.DateTimeToString(dataplane.CreatedAt))
// 	}
// 	if dataplane.LastUpdatedAt != nil {
// 		d.Set("updated_at", flex.DateTimeToString(dataplane.LastUpdatedAt))
// 	}

// 	return nil
// }

// func resourceIBMByocDataplaneUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	// Most dataplane fields are read-only
// 	// Only re-read the resource to get the latest state
// 	return resourceIBMByocDataplaneRead(context, d, meta)
// }

// func resourceIBMByocDataplaneDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	subscriptionID := d.Get("subscription_id").(string)
// 	dataplaneID := d.Get("dataplane_id").(string)

// 	deleteDataplaneOptions := &brokerapiv1.DeleteDataplaneOptions{
// 		SubscriptionID: &subscriptionID,
// 		DataplaneID:    &dataplaneID,
// 	}

// 	_, response, err := brokerApiClient.DeleteDataplaneWithContext(context, deleteDataplaneOptions)
// 	if err != nil {
// 		log.Printf("[DEBUG] DeleteDataplaneWithContext failed %s\n%s", err, response)
// 		return diag.FromErr(fmt.Errorf("DeleteDataplaneWithContext failed %s\n%s", err, response))
// 	}

// 	d.SetId("")
// 	return nil
// }
