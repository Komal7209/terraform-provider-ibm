// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.106.0-09823488-20250707-071701
 */

package brokerapi

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/cloud-go-sdk/brokerapiv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

func DataSourceIbmByocEngines() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmByocEnginesRead,

		Schema: map[string]*schema.Schema{
			"subscription_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subscription ID.",
			},
			"dataplane_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Data Plane ID.",
			},
			"engines": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of engines in the dataplane.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"engine_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Engine ID.",
						},
						"dataplane_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Data Plane ID.",
						},
						"availability_zone": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Availability zone for the engine.",
						},
						"storage_units": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of storage units.",
						},
						"compute_units": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of compute units.",
						},
						"engine_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the engine.",
						},
						"engine_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the engine.",
						},
						"endpoint_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Endpoint type for the engine.",
						},
						"instance_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance type for the engine.",
						},
						"replicas": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of replicas.",
						},
						"public_enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether public access is enabled.",
						},
						"private_link_service_enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether private link service is enabled.",
						},
						"oracle_compatibility": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether Oracle compatibility is enabled.",
						},
						"plan": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Plan for the engine.",
						},
						"profile_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Profile name for the engine.",
						},
						"service_principals": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Service principals.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"subscription_ids": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Subscription IDs.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"tags": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Tags associated with the engine.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"engine_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current status of the engine.",
						},
						"engine_status_message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status message for the engine.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation timestamp.",
						},
						"engine_short_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Short ID of the engine.",
						},
						"engine_ui_endpoint": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Engine UI endpoint.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"public": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Public URL for the engine UI console.",
									},
								},
							},
						},
						"jdbc_endpoint": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "JDBC endpoint.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"private": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Private JDBC connection string.",
									},
								},
							},
						},
						"linked_engine_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the linked engine.",
						},
						"link_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of engine link relationship.",
						},
						"linked_dataplane_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the dataplane where the linked engine resides.",
						},
						"linked_engine_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the linked engine.",
						},
						"version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Version of the engine.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmByocEnginesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_byoc_engines", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getEnginesOptions := brokerApiClient.NewGetEnginesOptions(
		d.Get("subscription_id").(string),
		d.Get("dataplane_id").(string),
	)

	getEnginesResponseIntf, response, err := brokerApiClient.GetEnginesWithContext(context, getEnginesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetEnginesWithContext failed: %s", err.Error()), "(Data) ibm_byoc_engines", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", d.Get("subscription_id").(string), d.Get("dataplane_id").(string)))

	engines := []map[string]interface{}{}
	for _, engineIntf := range getEnginesResponseIntf {
		engineMap := dataSourceIbmByocEnginesGetEnginesResponseToMap(engineIntf)
		engines = append(engines, engineMap)
	}

	if err = d.Set("engines", engines); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting engines: %s", err), "(Data) ibm_byoc_engines", "read", "set-engines").GetDiag()
	}

	return nil
}

func dataSourceIbmByocEnginesGetEnginesResponseToMap(model brokerapiv1.GetEnginesResponseIntf) map[string]interface{} {
	modelMap := make(map[string]interface{})

	// Handle discriminated union types
	if db2Engine, ok := model.(*brokerapiv1.GetEnginesResponseGetDb2EngineResponse); ok {
		if db2Engine.EngineID != nil {
			modelMap["engine_id"] = *db2Engine.EngineID
		}
		if db2Engine.DataplaneID != nil {
			modelMap["dataplane_id"] = *db2Engine.DataplaneID
		}
		if db2Engine.AvailabilityZone != nil {
			modelMap["availability_zone"] = *db2Engine.AvailabilityZone
		}
		if db2Engine.StorageUnits != nil {
			modelMap["storage_units"] = flex.IntValue(db2Engine.StorageUnits)
		}
		if db2Engine.ComputeUnits != nil {
			modelMap["compute_units"] = flex.IntValue(db2Engine.ComputeUnits)
		}
		if db2Engine.EngineName != nil {
			modelMap["engine_name"] = *db2Engine.EngineName
		}
		if db2Engine.EngineType != nil {
			modelMap["engine_type"] = *db2Engine.EngineType
		}
		if db2Engine.EndpointType != nil {
			modelMap["endpoint_type"] = *db2Engine.EndpointType
		}
		if db2Engine.InstanceType != nil {
			modelMap["instance_type"] = *db2Engine.InstanceType
		}
		if db2Engine.Replicas != nil {
			modelMap["replicas"] = flex.IntValue(db2Engine.Replicas)
		}
		if db2Engine.PublicEnabled != nil {
			modelMap["public_enabled"] = *db2Engine.PublicEnabled
		}
		if db2Engine.PrivateLinkServiceEnabled != nil {
			modelMap["private_link_service_enabled"] = *db2Engine.PrivateLinkServiceEnabled
		}
		if db2Engine.OracleCompatibility != nil {
			modelMap["oracle_compatibility"] = *db2Engine.OracleCompatibility
		}
		if db2Engine.Plan != nil {
			modelMap["plan"] = *db2Engine.Plan
		}
		if db2Engine.ProfileName != nil {
			modelMap["profile_name"] = *db2Engine.ProfileName
		}
		if db2Engine.ServicePrincipals != nil {
			modelMap["service_principals"] = db2Engine.ServicePrincipals
		}
		if db2Engine.SubscriptionIds != nil {
			modelMap["subscription_ids"] = db2Engine.SubscriptionIds
		}
		if db2Engine.Tags != nil {
			modelMap["tags"] = db2Engine.Tags
		}
		if db2Engine.EngineStatus != nil {
			modelMap["engine_status"] = *db2Engine.EngineStatus
		}
		if db2Engine.EngineStatusMessage != nil {
			modelMap["engine_status_message"] = *db2Engine.EngineStatusMessage
		}
		if db2Engine.CreatedAt != nil {
			modelMap["created_at"] = flex.DateTimeToString(db2Engine.CreatedAt)
		}
		if db2Engine.EngineShortID != nil {
			modelMap["engine_short_id"] = *db2Engine.EngineShortID
		}
		if db2Engine.EngineUiEndpoint != nil {
			engineUiEndpointMap, _ := ResourceIbmByocEngineEngineUiEndpointToMap(db2Engine.EngineUiEndpoint)
			modelMap["engine_ui_endpoint"] = []map[string]interface{}{engineUiEndpointMap}
		}
		if db2Engine.JdbcEndpoint != nil {
			jdbcEndpointMap, _ := ResourceIbmByocEngineJdbcEndpointToMap(db2Engine.JdbcEndpoint)
			modelMap["jdbc_endpoint"] = []map[string]interface{}{jdbcEndpointMap}
		}
		if db2Engine.LinkedEngineID != nil {
			modelMap["linked_engine_id"] = *db2Engine.LinkedEngineID
		}
		if db2Engine.LinkType != nil {
			modelMap["link_type"] = *db2Engine.LinkType
		}
		if db2Engine.LinkedDataplaneID != nil {
			modelMap["linked_dataplane_id"] = *db2Engine.LinkedDataplaneID
		}
		if db2Engine.LinkedEngineStatus != nil {
			modelMap["linked_engine_status"] = *db2Engine.LinkedEngineStatus
		}
		if db2Engine.Version != nil {
			modelMap["version"] = *db2Engine.Version
		}
	} else if db2whEngine, ok := model.(*brokerapiv1.GetEnginesResponseGetDb2WhEngineResponse); ok {
		// Handle DB2 Warehouse engine type
		if db2whEngine.EngineID != nil {
			modelMap["engine_id"] = *db2whEngine.EngineID
		}
		if db2whEngine.DataplaneID != nil {
			modelMap["dataplane_id"] = *db2whEngine.DataplaneID
		}
		if db2whEngine.AvailabilityZone != nil {
			modelMap["availability_zone"] = *db2whEngine.AvailabilityZone
		}
		if db2whEngine.StorageUnits != nil {
			modelMap["storage_units"] = flex.IntValue(db2whEngine.StorageUnits)
		}
		if db2whEngine.ComputeUnits != nil {
			modelMap["compute_units"] = flex.IntValue(db2whEngine.ComputeUnits)
		}
		if db2whEngine.EngineName != nil {
			modelMap["engine_name"] = *db2whEngine.EngineName
		}
		if db2whEngine.EngineType != nil {
			modelMap["engine_type"] = *db2whEngine.EngineType
		}
		if db2whEngine.Plan != nil {
			modelMap["plan"] = *db2whEngine.Plan
		}
		if db2whEngine.Tags != nil {
			modelMap["tags"] = db2whEngine.Tags
		}
		if db2whEngine.EngineStatus != nil {
			modelMap["engine_status"] = *db2whEngine.EngineStatus
		}
		if db2whEngine.EngineStatusMessage != nil {
			modelMap["engine_status_message"] = *db2whEngine.EngineStatusMessage
		}
		if db2whEngine.CreatedAt != nil {
			modelMap["created_at"] = flex.DateTimeToString(db2whEngine.CreatedAt)
		}
		if db2whEngine.EngineShortID != nil {
			modelMap["engine_short_id"] = *db2whEngine.EngineShortID
		}
		if db2whEngine.Version != nil {
			modelMap["version"] = *db2whEngine.Version
		}
	} else if netezzaEngine, ok := model.(*brokerapiv1.GetEnginesResponseGetNetezzaEngineResponse); ok {
		// Handle Netezza engine type
		if netezzaEngine.EngineID != nil {
			modelMap["engine_id"] = *netezzaEngine.EngineID
		}
		if netezzaEngine.DataplaneID != nil {
			modelMap["dataplane_id"] = *netezzaEngine.DataplaneID
		}
		if netezzaEngine.AvailabilityZone != nil {
			modelMap["availability_zone"] = *netezzaEngine.AvailabilityZone
		}
		if netezzaEngine.StorageUnits != nil {
			modelMap["storage_units"] = flex.IntValue(netezzaEngine.StorageUnits)
		}
		if netezzaEngine.ComputeUnits != nil {
			modelMap["compute_units"] = flex.IntValue(netezzaEngine.ComputeUnits)
		}
		if netezzaEngine.EngineName != nil {
			modelMap["engine_name"] = *netezzaEngine.EngineName
		}
		if netezzaEngine.EngineType != nil {
			modelMap["engine_type"] = *netezzaEngine.EngineType
		}
		if netezzaEngine.EndpointType != nil {
			modelMap["endpoint_type"] = *netezzaEngine.EndpointType
		}
		if netezzaEngine.ProfileName != nil {
			modelMap["profile_name"] = *netezzaEngine.ProfileName
		}
		if netezzaEngine.ServicePrincipals != nil {
			modelMap["service_principals"] = netezzaEngine.ServicePrincipals
		}
		if netezzaEngine.SubscriptionIds != nil {
			modelMap["subscription_ids"] = netezzaEngine.SubscriptionIds
		}
		if netezzaEngine.Tags != nil {
			modelMap["tags"] = netezzaEngine.Tags
		}
		if netezzaEngine.EngineStatus != nil {
			modelMap["engine_status"] = *netezzaEngine.EngineStatus
		}
		if netezzaEngine.EngineStatusMessage != nil {
			modelMap["engine_status_message"] = *netezzaEngine.EngineStatusMessage
		}
		if netezzaEngine.CreatedAt != nil {
			modelMap["created_at"] = flex.DateTimeToString(netezzaEngine.CreatedAt)
		}
		if netezzaEngine.EngineShortID != nil {
			modelMap["engine_short_id"] = *netezzaEngine.EngineShortID
		}
		if netezzaEngine.Version != nil {
			modelMap["version"] = *netezzaEngine.Version
		}
	}

	return modelMap
}
