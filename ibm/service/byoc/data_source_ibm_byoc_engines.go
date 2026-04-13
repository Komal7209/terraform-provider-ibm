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

func DataSourceIBMByocEngines() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMByocEnginesRead,

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
			"engines": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of engines",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"engine_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The engine ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the engine",
						},
						"engine_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of engine",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the engine",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The engine version",
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

func dataSourceIBMByocEnginesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	subscriptionID := d.Get("subscription_id").(string)
	dataplaneID := d.Get("dataplane_id").(string)

	subscriptionUUID := strfmt.UUID(subscriptionID)
	dataplaneUUID := strfmt.UUID(dataplaneID)

	getEnginesOptions := &brokerapiv1.GetEnginesOptions{
		SubscriptionID: &subscriptionUUID,
		DataplaneID:    &dataplaneUUID,
	}

	engineList, response, err := brokerApiClient.GetEnginesWithContext(context, getEnginesOptions)
	if err != nil {
		log.Printf("[DEBUG] GetEnginesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetEnginesWithContext failed %s\n%s", err, response))
	}

	var engines []interface{}
	for _, engineIntf := range engineList {
		// Convert interface to concrete type
		engine, ok := engineIntf.(*brokerapiv1.GetEnginesResponse)
		if !ok {
			continue
		}

		engineMap := map[string]interface{}{}
		if engine.EngineID != nil {
			engineMap["engine_id"] = engine.EngineID.String()
		}
		if engine.EngineName != nil {
			engineMap["name"] = *engine.EngineName
		}
		if engine.EngineType != nil {
			engineMap["engine_type"] = *engine.EngineType
		}
		if engine.EngineStatus != nil {
			engineMap["status"] = *engine.EngineStatus
		}
		if engine.CreatedAt != nil {
			engineMap["created_at"] = flex.DateTimeToString(engine.CreatedAt)
		}
		engines = append(engines, engineMap)
	}

	d.SetId(dataSourceIBMByocEnginesID(d))
	d.Set("engines", engines)

	return nil
}

func dataSourceIBMByocEnginesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
