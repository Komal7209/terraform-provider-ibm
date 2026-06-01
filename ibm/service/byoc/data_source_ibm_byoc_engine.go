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

func DataSourceIbmByocEngine() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmByocEngineRead,

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
			"engine_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Engine ID.",
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
	}
}

func dataSourceIbmByocEngineRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_byoc_engine", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getEngineByIdOptions := brokerApiClient.NewGetEngineByIdOptions(
		d.Get("subscription_id").(string),
		d.Get("dataplane_id").(string),
		d.Get("engine_id").(string),
	)

	getEngineByIdResponseIntf, response, err := brokerApiClient.GetEngineByIDWithContext(context, getEngineByIdOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetEngineByIDWithContext failed: %s", err.Error()), "(Data) ibm_byoc_engine", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getEngineByIdResponse := getEngineByIdResponseIntf.(*brokerapiv1.GetEngineByIdResponse)

	d.SetId(fmt.Sprintf("%s/%s/%s", d.Get("subscription_id").(string), d.Get("dataplane_id").(string), d.Get("engine_id").(string)))

	if err = d.Set("dataplane_id", getEngineByIdResponse.DataplaneID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dataplane_id: %s", err), "(Data) ibm_byoc_engine", "read", "set-dataplane_id").GetDiag()
	}

	if !core.IsNil(getEngineByIdResponse.AvailabilityZone) {
		if err = d.Set("availability_zone", getEngineByIdResponse.AvailabilityZone); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting availability_zone: %s", err), "(Data) ibm_byoc_engine", "read", "set-availability_zone").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.StorageUnits) {
		if err = d.Set("storage_units", flex.IntValue(getEngineByIdResponse.StorageUnits)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting storage_units: %s", err), "(Data) ibm_byoc_engine", "read", "set-storage_units").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.ComputeUnits) {
		if err = d.Set("compute_units", flex.IntValue(getEngineByIdResponse.ComputeUnits)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting compute_units: %s", err), "(Data) ibm_byoc_engine", "read", "set-compute_units").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.EngineName) {
		if err = d.Set("engine_name", getEngineByIdResponse.EngineName); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting engine_name: %s", err), "(Data) ibm_byoc_engine", "read", "set-engine_name").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.EngineType) {
		if err = d.Set("engine_type", getEngineByIdResponse.EngineType); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting engine_type: %s", err), "(Data) ibm_byoc_engine", "read", "set-engine_type").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.EndpointType) {
		if err = d.Set("endpoint_type", getEngineByIdResponse.EndpointType); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting endpoint_type: %s", err), "(Data) ibm_byoc_engine", "read", "set-endpoint_type").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.InstanceType) {
		if err = d.Set("instance_type", getEngineByIdResponse.InstanceType); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting instance_type: %s", err), "(Data) ibm_byoc_engine", "read", "set-instance_type").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.Replicas) {
		if err = d.Set("replicas", flex.IntValue(getEngineByIdResponse.Replicas)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting replicas: %s", err), "(Data) ibm_byoc_engine", "read", "set-replicas").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.PublicEnabled) {
		if err = d.Set("public_enabled", getEngineByIdResponse.PublicEnabled); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting public_enabled: %s", err), "(Data) ibm_byoc_engine", "read", "set-public_enabled").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.PrivateLinkServiceEnabled) {
		if err = d.Set("private_link_service_enabled", getEngineByIdResponse.PrivateLinkServiceEnabled); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting private_link_service_enabled: %s", err), "(Data) ibm_byoc_engine", "read", "set-private_link_service_enabled").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.OracleCompatibility) {
		if err = d.Set("oracle_compatibility", getEngineByIdResponse.OracleCompatibility); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting oracle_compatibility: %s", err), "(Data) ibm_byoc_engine", "read", "set-oracle_compatibility").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.Plan) {
		if err = d.Set("plan", getEngineByIdResponse.Plan); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting plan: %s", err), "(Data) ibm_byoc_engine", "read", "set-plan").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.ProfileName) {
		if err = d.Set("profile_name", getEngineByIdResponse.ProfileName); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting profile_name: %s", err), "(Data) ibm_byoc_engine", "read", "set-profile_name").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.ServicePrincipals) {
		if err = d.Set("service_principals", getEngineByIdResponse.ServicePrincipals); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting service_principals: %s", err), "(Data) ibm_byoc_engine", "read", "set-service_principals").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.SubscriptionIds) {
		if err = d.Set("subscription_ids", getEngineByIdResponse.SubscriptionIds); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting subscription_ids: %s", err), "(Data) ibm_byoc_engine", "read", "set-subscription_ids").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.Tags) {
		if err = d.Set("tags", getEngineByIdResponse.Tags); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_byoc_engine", "read", "set-tags").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.EngineStatus) {
		if err = d.Set("engine_status", getEngineByIdResponse.EngineStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting engine_status: %s", err), "(Data) ibm_byoc_engine", "read", "set-engine_status").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.EngineStatusMessage) {
		if err = d.Set("engine_status_message", getEngineByIdResponse.EngineStatusMessage); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting engine_status_message: %s", err), "(Data) ibm_byoc_engine", "read", "set-engine_status_message").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(getEngineByIdResponse.CreatedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_byoc_engine", "read", "set-created_at").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.EngineShortID) {
		if err = d.Set("engine_short_id", getEngineByIdResponse.EngineShortID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting engine_short_id: %s", err), "(Data) ibm_byoc_engine", "read", "set-engine_short_id").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.EngineUiEndpoint) {
		engineUiEndpointMap, err := ResourceIbmByocEngineEngineUiEndpointToMap(getEngineByIdResponse.EngineUiEndpoint)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_byoc_engine", "read", "engine_ui_endpoint-to-map").GetDiag()
		}
		if err = d.Set("engine_ui_endpoint", []map[string]interface{}{engineUiEndpointMap}); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting engine_ui_endpoint: %s", err), "(Data) ibm_byoc_engine", "read", "set-engine_ui_endpoint").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.JdbcEndpoint) {
		jdbcEndpointMap, err := ResourceIbmByocEngineJdbcEndpointToMap(getEngineByIdResponse.JdbcEndpoint)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_byoc_engine", "read", "jdbc_endpoint-to-map").GetDiag()
		}
		if err = d.Set("jdbc_endpoint", []map[string]interface{}{jdbcEndpointMap}); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting jdbc_endpoint: %s", err), "(Data) ibm_byoc_engine", "read", "set-jdbc_endpoint").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.LinkedEngineID) {
		if err = d.Set("linked_engine_id", getEngineByIdResponse.LinkedEngineID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting linked_engine_id: %s", err), "(Data) ibm_byoc_engine", "read", "set-linked_engine_id").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.LinkType) {
		if err = d.Set("link_type", getEngineByIdResponse.LinkType); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting link_type: %s", err), "(Data) ibm_byoc_engine", "read", "set-link_type").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.LinkedDataplaneID) {
		if err = d.Set("linked_dataplane_id", getEngineByIdResponse.LinkedDataplaneID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting linked_dataplane_id: %s", err), "(Data) ibm_byoc_engine", "read", "set-linked_dataplane_id").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.LinkedEngineStatus) {
		if err = d.Set("linked_engine_status", getEngineByIdResponse.LinkedEngineStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting linked_engine_status: %s", err), "(Data) ibm_byoc_engine", "read", "set-linked_engine_status").GetDiag()
		}
	}

	if !core.IsNil(getEngineByIdResponse.Version) {
		if err = d.Set("version", getEngineByIdResponse.Version); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting version: %s", err), "(Data) ibm_byoc_engine", "read", "set-version").GetDiag()
		}
	}

	return nil
}
