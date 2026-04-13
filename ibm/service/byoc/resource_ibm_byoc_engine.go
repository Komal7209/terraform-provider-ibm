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
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMByocEngine() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMByocEngineCreate,
		ReadContext:   resourceIBMByocEngineRead,
		UpdateContext: resourceIBMByocEngineUpdate,
		DeleteContext: resourceIBMByocEngineDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subscription ID",
			},
			"dataplane_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The dataplane ID",
			},
			"engine_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The engine ID",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the engine",
			},
			"engine_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of engine (db2, db2wh, netezza)",
			},
			"version": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The engine version",
			},
			"size": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The size/tier of the engine",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the engine",
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

func resourceIBMByocEngineCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	subscriptionID := d.Get("subscription_id").(string)
	dataplaneID := d.Get("dataplane_id").(string)

	createEngineOptions := &brokerapiv1.CreateEngineOptions{
		SubscriptionID: &subscriptionID,
		DataplaneID:    &dataplaneID,
	}

	// Build the request body based on engine type
	engineType := d.Get("engine_type").(string)
	
	var requestBody interface{}
	
	switch engineType {
	case "db2":
		db2Request := &brokerapiv1.CreateEngineBaseRequestCreateDb2EngineRequest{}
		if name, ok := d.GetOk("name"); ok {
			db2Request.EngineName = flex.PtrToString(name.(string))
		}
		requestBody = db2Request
	case "db2wh":
		db2whRequest := &brokerapiv1.CreateEngineBaseRequestCreateDb2WhEngineRequest{}
		if name, ok := d.GetOk("name"); ok {
			db2whRequest.EngineName = flex.PtrToString(name.(string))
		}
		requestBody = db2whRequest
	case "netezza":
		netezzaRequest := &brokerapiv1.CreateEngineBaseRequestCreateNetezzaEngineRequest{}
		if name, ok := d.GetOk("name"); ok {
			netezzaRequest.EngineName = flex.PtrToString(name.(string))
		}
		requestBody = netezzaRequest
	default:
		return diag.FromErr(fmt.Errorf("unsupported engine type: %s", engineType))
	}

	createEngineOptions.CreateEngineBaseRequest = requestBody.(brokerapiv1.CreateEngineBaseRequestIntf)

	engine, response, err := brokerApiClient.CreateEngineWithContext(context, createEngineOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateEngineWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateEngineWithContext failed %s\n%s", err, response))
	}

	engineID := *engine.EngineID
	d.SetId(fmt.Sprintf("%s/%s/%s", subscriptionID, dataplaneID, engineID))
	d.Set("engine_id", engineID)

	return resourceIBMByocEngineRead(context, d, meta)
}

func resourceIBMByocEngineRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	subscriptionID := d.Get("subscription_id").(string)
	dataplaneID := d.Get("dataplane_id").(string)
	engineID := d.Get("engine_id").(string)

	getEngineOptions := &brokerapiv1.GetEngineByIdOptions{
		SubscriptionID: &subscriptionID,
		DataplaneID:    &dataplaneID,
		EngineID:       &engineID,
	}

	engineIntf, response, err := brokerApiClient.GetEngineByIDWithContext(context, getEngineOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetEngineByIDWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetEngineByIDWithContext failed %s\n%s", err, response))
	}

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

func resourceIBMByocEngineUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// The BYOC SDK does not support updating engines directly
	// Any changes to engine properties require recreation
	// This function just refreshes the state
	return resourceIBMByocEngineRead(context, d, meta)
}

func resourceIBMByocEngineDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	subscriptionID := d.Get("subscription_id").(string)
	dataplaneID := d.Get("dataplane_id").(string)
	engineID := d.Get("engine_id").(string)

	deleteEngineOptions := &brokerapiv1.DeleteEngineOptions{
		SubscriptionID: &subscriptionID,
		DataplaneID:    &dataplaneID,
		EngineID:       &engineID,
	}

	_, response, err := brokerApiClient.DeleteEngineWithContext(context, deleteEngineOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteEngineWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteEngineWithContext failed %s\n%s", err, response))
	}

	d.SetId("")
	return nil
}

