// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.106.0-09823488-20250707-071701
 */

package byoc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/cloud-go-sdk/brokerapiv1"
	"github.com/go-openapi/strfmt"
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

	subscriptionID := strfmt.UUID(d.Get("subscription_id").(string))
	dataplaneID := strfmt.UUID(d.Get("dataplane_id").(string))

	getEnginesOptions := brokerApiClient.NewGetEnginesOptions(
		&subscriptionID,
		&dataplaneID,
	)

	getEnginesResponseIntf, response, err := brokerApiClient.GetEnginesWithContext(context, getEnginesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetEnginesWithContext failed: %s", err.Error()), "(Data) ibm_byoc_engines", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Log raw API response for debugging
	if response != nil {
		log.Printf("[DEBUG] HTTP Status Code: %d", response.StatusCode)
		if response.Result != nil {
			log.Printf("[DEBUG] Raw API Response Result: %+v", response.Result)
		}
		// Try to log the raw HTTP response body if available
		if response.RawResult != nil {
			log.Printf("[DEBUG] Raw HTTP Response Body available")
		}
	}
	log.Printf("[DEBUG] Number of engines returned by SDK: %d", len(getEnginesResponseIntf))
	if len(getEnginesResponseIntf) > 0 {
		log.Printf("[DEBUG] First engine type: %T", getEnginesResponseIntf[0])
		log.Printf("[DEBUG] First engine value: %+v", getEnginesResponseIntf[0])
	}

	d.SetId(fmt.Sprintf("%s/%s", d.Get("subscription_id").(string), d.Get("dataplane_id").(string)))

	engines := []map[string]interface{}{}
	for i, engineIntf := range getEnginesResponseIntf {
		log.Printf("[DEBUG] Processing engine %d, type: %T, value: %+v", i, engineIntf, engineIntf)
		engineMap := dataSourceIbmByocEnginesGetEnginesResponseToMap(engineIntf)
		log.Printf("[DEBUG] Engine %d map: %+v", i, engineMap)
		engines = append(engines, engineMap)
	}

	log.Printf("[DEBUG] Total engines processed: %d", len(engines))
	if len(engines) > 0 {
		log.Printf("[DEBUG] First engine map keys: %v", engines[0])
	}

	if err = d.Set("engines", engines); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting engines: %s", err), "(Data) ibm_byoc_engines", "read", "set-engines").GetDiag()
	}

	return nil
}

func dataSourceIbmByocEnginesGetEnginesResponseToMap(model brokerapiv1.GetEnginesResponseIntf) map[string]interface{} {
	modelMap := make(map[string]interface{})

	log.Printf("[DEBUG] Converting engine model, type: %T, value: %+v", model, model)

	// Handle base GetEnginesResponse type (most common case)
	if baseEngine, ok := model.(*brokerapiv1.GetEnginesResponse); ok {
		log.Printf("[DEBUG] Processing base GetEnginesResponse, EngineID: %v, EngineName: %v", baseEngine.EngineID, baseEngine.EngineName)

		// Only set fields that are actually present in the API response
		// This avoids showing Netezza-specific fields for DB2 engines and vice versa

		if baseEngine.EngineID != nil {
			modelMap["engine_id"] = baseEngine.EngineID.String()
			log.Printf("[DEBUG] Set engine_id to: %s", baseEngine.EngineID.String())
		}
		if baseEngine.DataplaneID != nil {
			modelMap["dataplane_id"] = baseEngine.DataplaneID.String()
		}
		if baseEngine.StorageUnits != nil {
			modelMap["storage_units"] = flex.IntValue(baseEngine.StorageUnits)
		}
		if baseEngine.ComputeUnits != nil {
			modelMap["compute_units"] = flex.IntValue(baseEngine.ComputeUnits)
		}
		if baseEngine.EngineName != nil {
			modelMap["engine_name"] = *baseEngine.EngineName
			log.Printf("[DEBUG] Set engine_name to: %s", *baseEngine.EngineName)
		}
		if baseEngine.EngineType != nil {
			modelMap["engine_type"] = *baseEngine.EngineType
		}
		if baseEngine.EndpointType != nil {
			modelMap["endpoint_type"] = *baseEngine.EndpointType
		}
		if baseEngine.InstanceType != nil {
			modelMap["instance_type"] = *baseEngine.InstanceType
		}
		if baseEngine.Replicas != nil {
			modelMap["replicas"] = flex.IntValue(baseEngine.Replicas)
		}
		if baseEngine.PublicEnabled != nil {
			modelMap["public_enabled"] = *baseEngine.PublicEnabled
		}
		if baseEngine.PrivateLinkServiceEnabled != nil {
			modelMap["private_link_service_enabled"] = *baseEngine.PrivateLinkServiceEnabled
		}
		if baseEngine.OracleCompatibility != nil {
			modelMap["oracle_compatibility"] = *baseEngine.OracleCompatibility
		}
		if baseEngine.Tags != nil && len(baseEngine.Tags) > 0 {
			modelMap["tags"] = baseEngine.Tags
		}
		if baseEngine.EngineStatus != nil {
			modelMap["engine_status"] = *baseEngine.EngineStatus
		}
		if baseEngine.EngineStatusMessage != nil {
			modelMap["engine_status_message"] = *baseEngine.EngineStatusMessage
		}
		if baseEngine.CreatedAt != nil {
			modelMap["created_at"] = baseEngine.CreatedAt.String()
		}
		if baseEngine.EngineShortID != nil {
			modelMap["engine_short_id"] = *baseEngine.EngineShortID
		}
		if baseEngine.EngineUiEndpoint != nil {
			engineUiEndpointMap, err := ResourceIbmByocEngineEngineUiEndpointToMap(baseEngine.EngineUiEndpoint)
			if err != nil {
				log.Printf("[ERROR] Error converting EngineUiEndpoint: %s", err)
			} else {
				modelMap["engine_ui_endpoint"] = []map[string]interface{}{engineUiEndpointMap}
			}
		}
		if baseEngine.JdbcEndpoint != nil {
			jdbcEndpointMap, err := ResourceIbmByocEngineJdbcEndpointToMap(baseEngine.JdbcEndpoint)
			if err != nil {
				log.Printf("[ERROR] Error converting JdbcEndpoint: %s", err)
			} else {
				modelMap["jdbc_endpoint"] = []map[string]interface{}{jdbcEndpointMap}
			}
		}
		if baseEngine.LinkedEngineID != nil {
			modelMap["linked_engine_id"] = baseEngine.LinkedEngineID.String()
		}
		if baseEngine.LinkType != nil {
			modelMap["link_type"] = *baseEngine.LinkType
		}
		if baseEngine.LinkedDataplaneID != nil {
			modelMap["linked_dataplane_id"] = baseEngine.LinkedDataplaneID.String()
		}
		if baseEngine.LinkedEngineStatus != nil {
			modelMap["linked_engine_status"] = *baseEngine.LinkedEngineStatus
		}
		if baseEngine.Version != nil {
			modelMap["version"] = *baseEngine.Version
		}

		return modelMap
	}

	// Handle discriminated union types (fallback for specific engine types)
	if db2Engine, ok := model.(*brokerapiv1.GetEnginesResponseGetDb2EngineResponse); ok {
		log.Printf("[DEBUG] Processing DB2 engine, EngineID: %v, EngineName: %v", db2Engine.EngineID, db2Engine.EngineName)
		// Always set engine_id, even if empty, to ensure the field exists in the map
		if db2Engine.EngineID != nil {
			modelMap["engine_id"] = *db2Engine.EngineID
			log.Printf("[DEBUG] Set engine_id to: %s", *db2Engine.EngineID)
		} else {
			modelMap["engine_id"] = ""
			log.Printf("[WARN] DB2 EngineID was nil, set to empty string")
		}
		if db2Engine.DataplaneID != nil {
			modelMap["dataplane_id"] = *db2Engine.DataplaneID
		} else {
			modelMap["dataplane_id"] = ""
		}
		if db2Engine.StorageUnits != nil {
			modelMap["storage_units"] = flex.IntValue(db2Engine.StorageUnits)
		} else {
			modelMap["storage_units"] = 0
		}
		if db2Engine.ComputeUnits != nil {
			modelMap["compute_units"] = flex.IntValue(db2Engine.ComputeUnits)
		} else {
			modelMap["compute_units"] = 0
		}
		if db2Engine.EngineName != nil {
			modelMap["engine_name"] = *db2Engine.EngineName
			log.Printf("[DEBUG] Set engine_name to: %s", *db2Engine.EngineName)
		} else {
			modelMap["engine_name"] = ""
			log.Printf("[WARN] DB2 EngineName was nil, set to empty string")
		}
		if db2Engine.EngineType != nil {
			modelMap["engine_type"] = *db2Engine.EngineType
		} else {
			modelMap["engine_type"] = ""
		}
		if db2Engine.EndpointType != nil {
			modelMap["endpoint_type"] = *db2Engine.EndpointType
		} else {
			modelMap["endpoint_type"] = ""
		}
		if db2Engine.InstanceType != nil {
			modelMap["instance_type"] = *db2Engine.InstanceType
		} else {
			modelMap["instance_type"] = ""
		}
		if db2Engine.Replicas != nil {
			modelMap["replicas"] = flex.IntValue(db2Engine.Replicas)
		} else {
			modelMap["replicas"] = 0
		}
		if db2Engine.PublicEnabled != nil {
			modelMap["public_enabled"] = *db2Engine.PublicEnabled
		} else {
			modelMap["public_enabled"] = false
		}
		if db2Engine.PrivateLinkServiceEnabled != nil {
			modelMap["private_link_service_enabled"] = *db2Engine.PrivateLinkServiceEnabled
		} else {
			modelMap["private_link_service_enabled"] = false
		}
		if db2Engine.OracleCompatibility != nil {
			modelMap["oracle_compatibility"] = *db2Engine.OracleCompatibility
		} else {
			modelMap["oracle_compatibility"] = false
		}
		// Note: Plan, ProfileName, ServicePrincipals, and SubscriptionIds are not available for DB2 engine type
		if db2Engine.Tags != nil {
			modelMap["tags"] = db2Engine.Tags
		} else {
			modelMap["tags"] = []string{}
		}
		if db2Engine.EngineStatus != nil {
			modelMap["engine_status"] = *db2Engine.EngineStatus
		} else {
			modelMap["engine_status"] = ""
		}
		if db2Engine.EngineStatusMessage != nil {
			modelMap["engine_status_message"] = *db2Engine.EngineStatusMessage
		} else {
			modelMap["engine_status_message"] = ""
		}
		if db2Engine.CreatedAt != nil {
			modelMap["created_at"] = flex.DateTimeToString(db2Engine.CreatedAt)
		} else {
			modelMap["created_at"] = ""
		}
		if db2Engine.EngineShortID != nil {
			modelMap["engine_short_id"] = *db2Engine.EngineShortID
		} else {
			modelMap["engine_short_id"] = ""
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
	} else {
		// None of the known engine types matched
		log.Printf("[ERROR] Unknown engine type received: %T", model)
		log.Printf("[ERROR] Engine value: %+v", model)
		// Return empty map with at least the type information
		modelMap["engine_type"] = "unknown"
		modelMap["engine_status"] = "unknown"
	}

	log.Printf("[DEBUG] Final modelMap: %+v", modelMap)
	return modelMap
}
