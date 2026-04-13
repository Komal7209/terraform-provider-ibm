// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package byoc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/cloud-go-sdk/brokerapiv1"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMByocDataplane() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMByocDataplaneRead,

		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The subscription ID",
			},
			"dataplane_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dataplane ID",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the dataplane",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the dataplane",
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The region where the dataplane is deployed",
			},
			"cloud_provider": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The cloud provider (AWS, Azure, etc.)",
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

func dataSourceIBMByocDataplaneRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] GetDataplaneWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetDataplaneWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", subscriptionID, dataplaneID))

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
