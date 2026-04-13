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
	"github.com/IBM/cloud-go-sdk/brokerapiv1"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMByocDataplanes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMByocDataplanesRead,

		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by subscription ID",
			},
			"dataplanes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of dataplanes",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataplane_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dataplane ID",
						},
						"subscription_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subscription ID",
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
							Description: "The cloud provider",
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
				},
			},
		},
	}
}

func dataSourceIBMByocDataplanesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	var dataplanes []interface{}

	if subscriptionID, ok := d.GetOk("subscription_id"); ok {
		// Get dataplanes for specific subscription
		subscriptionUUID := strfmt.UUID(subscriptionID.(string))
		getDataplanesOptions := &brokerapiv1.GetDataplanesBySubscriptionOptions{
			SubscriptionID: &subscriptionUUID,
		}

		dataplaneList, response, err := brokerApiClient.GetDataplanesBySubscriptionWithContext(context, getDataplanesOptions)
		if err != nil {
			log.Printf("[DEBUG] GetDataplanesBySubscriptionWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("GetDataplanesBySubscriptionWithContext failed %s\n%s", err, response))
		}

		for _, dataplane := range dataplaneList {
			dataplaneMap := map[string]interface{}{}
			if dataplane.DataplaneID != nil {
				dataplaneMap["dataplane_id"] = dataplane.DataplaneID.String()
			}
			if dataplane.McspSubscriptionID != nil {
				dataplaneMap["subscription_id"] = *dataplane.McspSubscriptionID
			}
			if dataplane.ServiceInstanceName != nil {
				dataplaneMap["name"] = *dataplane.ServiceInstanceName
			}
			if dataplane.ResourceProvisioningStatus != nil {
				dataplaneMap["status"] = *dataplane.ResourceProvisioningStatus
			}
			if dataplane.HyperscalerRegion != nil {
				dataplaneMap["region"] = *dataplane.HyperscalerRegion
			}
			if dataplane.Hyperscaler != nil {
				dataplaneMap["cloud_provider"] = *dataplane.Hyperscaler
			}
			if dataplane.CreatedAt != nil {
				dataplaneMap["created_at"] = flex.DateTimeToString(dataplane.CreatedAt)
			}
			if dataplane.LastUpdatedAt != nil {
				dataplaneMap["updated_at"] = flex.DateTimeToString(dataplane.LastUpdatedAt)
			}
			dataplanes = append(dataplanes, dataplaneMap)
		}
	} else {
		// Get all dataplanes
		getAllDataplanesOptions := &brokerapiv1.GetAllDataplanesOptions{}

		dataplaneList, response, err := brokerApiClient.GetAllDataplanesWithContext(context, getAllDataplanesOptions)
		if err != nil {
			log.Printf("[DEBUG] GetAllDataplanesWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("GetAllDataplanesWithContext failed %s\n%s", err, response))
		}

		for _, dataplane := range dataplaneList {
			dataplaneMap := map[string]interface{}{}
			if dataplane.DataplaneID != nil {
				dataplaneMap["dataplane_id"] = *dataplane.DataplaneID
			}
			if dataplane.McspSubscriptionID != nil {
				dataplaneMap["subscription_id"] = *dataplane.McspSubscriptionID
			}
			if dataplane.ServiceInstanceName != nil {
				dataplaneMap["name"] = *dataplane.ServiceInstanceName
			}
			if dataplane.ResourceProvisioningStatus != nil {
				dataplaneMap["status"] = *dataplane.ResourceProvisioningStatus
			}
			if dataplane.HyperscalerRegion != nil {
				dataplaneMap["region"] = *dataplane.HyperscalerRegion
			}
			if dataplane.Hyperscaler != nil {
				dataplaneMap["cloud_provider"] = *dataplane.Hyperscaler
			}
			if dataplane.CreatedAt != nil {
				dataplaneMap["created_at"] = *dataplane.CreatedAt
			}
			if dataplane.LastUpdatedAt != nil {
				dataplaneMap["updated_at"] = *dataplane.LastUpdatedAt
			}
			dataplanes = append(dataplanes, dataplaneMap)
		}
	}

	d.SetId(dataSourceIBMByocDataplanesID(d))
	d.Set("dataplanes", dataplanes)

	return nil
}

func dataSourceIBMByocDataplanesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
