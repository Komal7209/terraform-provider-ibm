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

func DataSourceIBMByocEngine() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMByocEngineRead,

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
			"engine_id": {
				Type:        schema.TypeString,
				Required:    true,
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
				Description: "The type of engine (db2, db2wh, netezza)",
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
	}
}

func dataSourceIBMByocEngineRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	subscriptionID := d.Get("subscription_id").(string)
	dataplaneID := d.Get("dataplane_id").(string)
	engineID := d.Get("engine_id").(string)

	subscriptionUUID := strfmt.UUID(subscriptionID)
	dataplaneUUID := strfmt.UUID(dataplaneID)
	engineUUID := strfmt.UUID(engineID)

	getEngineOptions := &brokerapiv1.GetEngineByIdOptions{
		SubscriptionID: &subscriptionUUID,
		DataplaneID:    &dataplaneUUID,
		EngineID:       &engineUUID,
	}

	engineIntf, response, err := brokerApiClient.GetEngineByIDWithContext(context, getEngineOptions)
	if err != nil {
		log.Printf("[DEBUG] GetEngineByIDWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetEngineByIDWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", subscriptionID, dataplaneID, engineID))

	// Convert interface to concrete type
	engine, ok := engineIntf.(*brokerapiv1.GetEngineByIdResponse)
	if !ok {
		return diag.FromErr(fmt.Errorf("Failed to convert engine response to expected type"))
	}

	if engine.EngineName != nil {
		d.Set("name", *engine.EngineName)
	}
	if engine.EngineType != nil {
		d.Set("engine_type", *engine.EngineType)
	}
	if engine.EngineStatus != nil {
		d.Set("status", *engine.EngineStatus)
	}
	if engine.CreatedAt != nil {
		d.Set("created_at", flex.DateTimeToString(engine.CreatedAt))
	}

	return nil
}
