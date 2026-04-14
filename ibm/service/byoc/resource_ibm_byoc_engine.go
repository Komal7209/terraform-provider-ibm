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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	// "github.com/IBM-Cloud/terraform-provider-ibm/brokerapiv1"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/cloud-go-sdk/brokerapiv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
)

func ResourceIbmByocEngine() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmByocEngineCreate,
		ReadContext:   resourceIbmByocEngineRead,
		UpdateContext: resourceIbmByocEngineUpdate,
		DeleteContext: resourceIbmByocEngineDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"subscription_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Subscription ID.",
			},
			"dataplane_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Data Plane ID.",
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_units": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"compute_units": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"engine_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"engine_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_byoc_engine", "engine_type"),
			},
			"endpoint_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_byoc_engine", "endpoint_type"),
			},
			"instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_link_service_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"public_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"oracle_compatibility": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"replicas": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"plan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_principals": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"subscription_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"engine_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_status_message": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_short_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIbmByocEngineValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "engine_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "db2",
		},
		validate.ValidateSchema{
			Identifier:                 "endpoint_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "private, public, public-private",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_byoc_engine", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmByocEngineCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bodyModelMap := map[string]interface{}{}
	createEngineOptions := &brokerapiv1.CreateEngineOptions{}

	if _, ok := d.GetOk("engine_type"); ok {
		bodyModelMap["engine_type"] = d.Get("engine_type")
	}
	if _, ok := d.GetOk("engine_name"); ok {
		bodyModelMap["engine_name"] = d.Get("engine_name")
	}
	if _, ok := d.GetOk("admin_username"); ok {
		bodyModelMap["admin_username"] = d.Get("admin_username")
	}
	if _, ok := d.GetOk("admin_password"); ok {
		bodyModelMap["admin_password"] = d.Get("admin_password")
	}
	if _, ok := d.GetOk("admin_email"); ok {
		bodyModelMap["admin_email"] = d.Get("admin_email")
	}
	if _, ok := d.GetOk("availability_zone"); ok {
		bodyModelMap["availability_zone"] = d.Get("availability_zone")
	}
	if _, ok := d.GetOk("storage_units"); ok {
		bodyModelMap["storage_units"] = d.Get("storage_units")
	}
	if _, ok := d.GetOk("compute_units"); ok {
		bodyModelMap["compute_units"] = d.Get("compute_units")
	}
	if _, ok := d.GetOk("private_link_service_enabled"); ok {
		bodyModelMap["private_link_service_enabled"] = d.Get("private_link_service_enabled")
	}
	if _, ok := d.GetOk("public_enabled"); ok {
		bodyModelMap["public_enabled"] = d.Get("public_enabled")
	}
	if _, ok := d.GetOk("oracle_compatibility"); ok {
		bodyModelMap["oracle_compatibility"] = d.Get("oracle_compatibility")
	}
	if _, ok := d.GetOk("replicas"); ok {
		bodyModelMap["replicas"] = d.Get("replicas")
	}
	if _, ok := d.GetOk("instance_type"); ok {
		bodyModelMap["instance_type"] = d.Get("instance_type")
	}
	if _, ok := d.GetOk("plan"); ok {
		bodyModelMap["plan"] = d.Get("plan")
	}
	if _, ok := d.GetOk("endpoint_type"); ok {
		bodyModelMap["endpoint_type"] = d.Get("endpoint_type")
	}
	if _, ok := d.GetOk("profile_name"); ok {
		bodyModelMap["profile_name"] = d.Get("profile_name")
	}
	if _, ok := d.GetOk("service_principals"); ok {
		bodyModelMap["service_principals"] = d.Get("service_principals")
	}
	if _, ok := d.GetOk("subscription_ids"); ok {
		bodyModelMap["subscription_ids"] = d.Get("subscription_ids")
	}
	createEngineOptions.SetSubscriptionID(core.UUIDPtr(strfmt.UUID(d.Get("subscription_id").(string))))
	createEngineOptions.SetDataplaneID(core.UUIDPtr(strfmt.UUID(d.Get("dataplane_id").(string))))
	convertedModel, err := ResourceIbmByocEngineMapToCreateEngineBaseRequest(bodyModelMap)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "create", "parse-request-body").GetDiag()
	}
	createEngineOptions.CreateEngineBaseRequest = convertedModel

	createEngineResponse, _, err := brokerApiClient.CreateEngineWithContext(context, createEngineOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateEngineWithContext failed: %s", err.Error()), "ibm_byoc_engine", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *createEngineOptions.SubscriptionID, *createEngineOptions.DataplaneID, *createEngineResponse.EngineID))

	return resourceIbmByocEngineRead(context, d, meta)
}

func resourceIbmByocEngineRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getEngineByIdOptions := &brokerapiv1.GetEngineByIdOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "sep-id-parts").GetDiag()
	}

	subscriptionUUID := strfmt.UUID(parts[0])
	dataplaneUUID := strfmt.UUID(parts[1])
	engineUUID := strfmt.UUID(parts[2])

	getEngineByIdOptions.SetSubscriptionID(&subscriptionUUID)
	getEngineByIdOptions.SetDataplaneID(&dataplaneUUID)
	getEngineByIdOptions.SetEngineID(&engineUUID)

	getEngineByIdResponseIntf, response, err := brokerApiClient.GetEngineByIDWithContext(context, getEngineByIdOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetEngineByIDWithContext failed: %s", err.Error()), "ibm_byoc_engine", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getEngineByIdResponse := getEngineByIdResponseIntf.(*brokerapiv1.GetEngineByIdResponse)
	if err = d.Set("dataplane_id", getEngineByIdResponse.DataplaneID); err != nil {
		err = fmt.Errorf("Error setting dataplane_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-dataplane_id").GetDiag()
	}
	if !core.IsNil(getEngineByIdResponse.AvailabilityZone) {
		if err = d.Set("availability_zone", getEngineByIdResponse.AvailabilityZone); err != nil {
			err = fmt.Errorf("Error setting availability_zone: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-availability_zone").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.StorageUnits) {
		if err = d.Set("storage_units", flex.IntValue(getEngineByIdResponse.StorageUnits)); err != nil {
			err = fmt.Errorf("Error setting storage_units: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-storage_units").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.ComputeUnits) {
		if err = d.Set("compute_units", flex.IntValue(getEngineByIdResponse.ComputeUnits)); err != nil {
			err = fmt.Errorf("Error setting compute_units: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-compute_units").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.EngineName) {
		if err = d.Set("engine_name", getEngineByIdResponse.EngineName); err != nil {
			err = fmt.Errorf("Error setting engine_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-engine_name").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.EngineType) {
		if err = d.Set("engine_type", getEngineByIdResponse.EngineType); err != nil {
			err = fmt.Errorf("Error setting engine_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-engine_type").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.EndpointType) {
		if err = d.Set("endpoint_type", getEngineByIdResponse.EndpointType); err != nil {
			err = fmt.Errorf("Error setting endpoint_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-endpoint_type").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.InstanceType) {
		if err = d.Set("instance_type", getEngineByIdResponse.InstanceType); err != nil {
			err = fmt.Errorf("Error setting instance_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-instance_type").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.PrivateLinkServiceEnabled) {
		if err = d.Set("private_link_service_enabled", getEngineByIdResponse.PrivateLinkServiceEnabled); err != nil {
			err = fmt.Errorf("Error setting private_link_service_enabled: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-private_link_service_enabled").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.PublicEnabled) {
		if err = d.Set("public_enabled", getEngineByIdResponse.PublicEnabled); err != nil {
			err = fmt.Errorf("Error setting public_enabled: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-public_enabled").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.OracleCompatibility) {
		if err = d.Set("oracle_compatibility", getEngineByIdResponse.OracleCompatibility); err != nil {
			err = fmt.Errorf("Error setting oracle_compatibility: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-oracle_compatibility").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.Replicas) {
		if err = d.Set("replicas", flex.IntValue(getEngineByIdResponse.Replicas)); err != nil {
			err = fmt.Errorf("Error setting replicas: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-replicas").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.Plan) {
		if err = d.Set("plan", getEngineByIdResponse.Plan); err != nil {
			err = fmt.Errorf("Error setting plan: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-plan").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.ProfileName) {
		if err = d.Set("profile_name", getEngineByIdResponse.ProfileName); err != nil {
			err = fmt.Errorf("Error setting profile_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-profile_name").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.ServicePrincipals) {
		if err = d.Set("service_principals", getEngineByIdResponse.ServicePrincipals); err != nil {
			err = fmt.Errorf("Error setting service_principals: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-service_principals").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.SubscriptionIds) {
		if err = d.Set("subscription_ids", getEngineByIdResponse.SubscriptionIds); err != nil {
			err = fmt.Errorf("Error setting subscription_ids: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-subscription_ids").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.Tags) {
		if err = d.Set("tags", getEngineByIdResponse.Tags); err != nil {
			err = fmt.Errorf("Error setting tags: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-tags").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.EngineStatus) {
		if err = d.Set("engine_status", getEngineByIdResponse.EngineStatus); err != nil {
			err = fmt.Errorf("Error setting engine_status: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-engine_status").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.EngineStatusMessage) {
		if err = d.Set("engine_status_message", getEngineByIdResponse.EngineStatusMessage); err != nil {
			err = fmt.Errorf("Error setting engine_status_message: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-engine_status_message").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(getEngineByIdResponse.CreatedAt)); err != nil {
			err = fmt.Errorf("Error setting created_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-created_at").GetDiag()
		}
	}
	if !core.IsNil(getEngineByIdResponse.EngineShortID) {
		if err = d.Set("engine_short_id", getEngineByIdResponse.EngineShortID); err != nil {
			err = fmt.Errorf("Error setting engine_short_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "read", "set-engine_short_id").GetDiag()
		}
	}

	return nil
}

func resourceIbmByocEngineUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	performEngineOperationOptions := &brokerapiv1.PerformEngineOperationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "update", "sep-id-parts").GetDiag()
	}

	subscriptionUUIDUpd := strfmt.UUID(parts[0])
	dataplaneUUIDUpd := strfmt.UUID(parts[1])
	engineUUIDUpd := strfmt.UUID(parts[2])

	performEngineOperationOptions.SetSubscriptionID(&subscriptionUUIDUpd)
	performEngineOperationOptions.SetDataplaneID(&dataplaneUUIDUpd)
	performEngineOperationOptions.SetEngineID(&engineUUIDUpd)

	hasChange := false

	if d.HasChange("subscription_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "subscription_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_byoc_engine", "update", "subscription_id-forces-new").GetDiag()
	}
	if d.HasChange("dataplane_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "dataplane_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_byoc_engine", "update", "dataplane_id-forces-new").GetDiag()
	}
	if d.HasChange("operation_type") || d.HasChange("engine_name") || d.HasChange("linked_dataplane_id") {
		performEngineOperationOptions.SetOperationType(d.Get("operation_type").(string))
		performEngineOperationOptions.SetEngineName(d.Get("engine_name").(string))
		performEngineOperationOptions.SetLinkedDataplaneID(core.UUIDPtr(strfmt.UUID(d.Get("linked_dataplane_id").(string))))
		hasChange = true
	}
	if d.HasChange("tags") {
		tags := make(map[string]string)
		for k, v := range d.Get("tags").(map[string]interface{}) {
			tags[k] = v.(string)
		}
		performEngineOperationOptions.SetTags(tags)
		hasChange = true
	}
	if d.HasChange("backup_id") {
		performEngineOperationOptions.SetBackupID(d.Get("backup_id").(string))
		hasChange = true
	}
	if d.HasChange("backup_timestamp") {
		fmtDateTimeBackupTimestamp, err := core.ParseDateTime(d.Get("backup_timestamp").(string))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "update", "parse-backup_timestamp").GetDiag()
		}
		performEngineOperationOptions.SetBackupTimestamp(&fmtDateTimeBackupTimestamp)
		hasChange = true
	}
	if d.HasChange("instance_name") {
		performEngineOperationOptions.SetInstanceName(d.Get("instance_name").(string))
		hasChange = true
	}
	if d.HasChange("admin_username") {
		performEngineOperationOptions.SetAdminUsername(d.Get("admin_username").(string))
		hasChange = true
	}
	if d.HasChange("admin_password") {
		performEngineOperationOptions.SetAdminPassword(d.Get("admin_password").(string))
		hasChange = true
	}
	if d.HasChange("admin_email") {
		performEngineOperationOptions.SetAdminEmail(d.Get("admin_email").(string))
		hasChange = true
	}

	if hasChange {
		_, _, err = brokerApiClient.PerformEngineOperationWithContext(context, performEngineOperationOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PerformEngineOperationWithContext failed: %s", err.Error()), "ibm_byoc_engine", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmByocEngineRead(context, d, meta)
}

func resourceIbmByocEngineDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brokerApiClient, err := meta.(conns.ClientSession).BrokerApiV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteEngineOptions := &brokerapiv1.DeleteEngineOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_byoc_engine", "delete", "sep-id-parts").GetDiag()
	}

	subscriptionUUIDDel := strfmt.UUID(parts[0])
	dataplaneUUIDDel := strfmt.UUID(parts[1])
	engineUUIDDel := strfmt.UUID(parts[2])

	deleteEngineOptions.SetSubscriptionID(&subscriptionUUIDDel)
	deleteEngineOptions.SetDataplaneID(&dataplaneUUIDDel)
	deleteEngineOptions.SetEngineID(&engineUUIDDel)

	_, _, err = brokerApiClient.DeleteEngineWithContext(context, deleteEngineOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteEngineWithContext failed: %s", err.Error()), "ibm_byoc_engine", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmByocEngineMapToCreateEngineBaseRequest(modelMap map[string]interface{}) (brokerapiv1.CreateEngineBaseRequestIntf, error) {
	discValue, ok := modelMap["engine_type"]
	if ok {
		if discValue == "db2" {
			return ResourceIbmByocEngineMapToCreateEngineBaseRequestCreateDb2EngineRequest(modelMap)
		} else if discValue == "db2wh" {
			return ResourceIbmByocEngineMapToCreateEngineBaseRequestCreateDb2WhEngineRequest(modelMap)
		} else if discValue == "netezza-aws" {
			return ResourceIbmByocEngineMapToCreateEngineBaseRequestCreateNetezzaEngineRequest(modelMap)
		} else if discValue == "netezza-azure" {
			return ResourceIbmByocEngineMapToCreateEngineBaseRequestCreateNetezzaEngineRequest(modelMap)
		} else {
			return nil, fmt.Errorf("unexpected value for discriminator property 'engine_type' found in map: '%s'", discValue)
		}
	} else {
		return nil, fmt.Errorf("discriminator property 'engine_type' not found in map")
	}
}

func ResourceIbmByocEngineMapToCreateEngineBaseRequestCreateDb2WhEngineRequest(modelMap map[string]interface{}) (*brokerapiv1.CreateEngineBaseRequestCreateDb2WhEngineRequest, error) {
	model := &brokerapiv1.CreateEngineBaseRequestCreateDb2WhEngineRequest{}
	if modelMap["engine_type"] != nil && modelMap["engine_type"].(string) != "" {
		model.EngineType = core.StringPtr(modelMap["engine_type"].(string))
	}
	if modelMap["engine_name"] != nil && modelMap["engine_name"].(string) != "" {
		model.EngineName = core.StringPtr(modelMap["engine_name"].(string))
	}
	if modelMap["admin_username"] != nil && modelMap["admin_username"].(string) != "" {
		model.AdminUsername = core.StringPtr(modelMap["admin_username"].(string))
	}
	if modelMap["admin_password"] != nil && modelMap["admin_password"].(string) != "" {
		model.AdminPassword = core.StringPtr(modelMap["admin_password"].(string))
	}
	if modelMap["admin_email"] != nil && modelMap["admin_email"].(string) != "" {
		model.AdminEmail = core.StringPtr(modelMap["admin_email"].(string))
	}
	if modelMap["availability_zone"] != nil && modelMap["availability_zone"].(string) != "" {
		model.AvailabilityZone = core.StringPtr(modelMap["availability_zone"].(string))
	}
	if modelMap["storage_units"] != nil {
		model.StorageUnits = core.Int64Ptr(int64(modelMap["storage_units"].(int)))
	}
	if modelMap["compute_units"] != nil {
		model.ComputeUnits = core.Int64Ptr(int64(modelMap["compute_units"].(int)))
	}
	if modelMap["plan"] != nil && modelMap["plan"].(string) != "" {
		model.Plan = core.StringPtr(modelMap["plan"].(string))
	}
	return model, nil
}

func ResourceIbmByocEngineMapToCreateEngineBaseRequestCreateDb2EngineRequest(modelMap map[string]interface{}) (*brokerapiv1.CreateEngineBaseRequestCreateDb2EngineRequest, error) {
	model := &brokerapiv1.CreateEngineBaseRequestCreateDb2EngineRequest{}
	if modelMap["engine_type"] != nil && modelMap["engine_type"].(string) != "" {
		model.EngineType = core.StringPtr(modelMap["engine_type"].(string))
	}
	if modelMap["engine_name"] != nil && modelMap["engine_name"].(string) != "" {
		model.EngineName = core.StringPtr(modelMap["engine_name"].(string))
	}
	if modelMap["admin_username"] != nil && modelMap["admin_username"].(string) != "" {
		model.AdminUsername = core.StringPtr(modelMap["admin_username"].(string))
	}
	if modelMap["admin_password"] != nil && modelMap["admin_password"].(string) != "" {
		model.AdminPassword = core.StringPtr(modelMap["admin_password"].(string))
	}
	if modelMap["admin_email"] != nil && modelMap["admin_email"].(string) != "" {
		model.AdminEmail = core.StringPtr(modelMap["admin_email"].(string))
	}
	if modelMap["availability_zone"] != nil && modelMap["availability_zone"].(string) != "" {
		model.AvailabilityZone = core.StringPtr(modelMap["availability_zone"].(string))
	}
	if modelMap["storage_units"] != nil {
		model.StorageUnits = core.Int64Ptr(int64(modelMap["storage_units"].(int)))
	}
	if modelMap["compute_units"] != nil {
		model.ComputeUnits = core.Int64Ptr(int64(modelMap["compute_units"].(int)))
	}
	if modelMap["private_link_service_enabled"] != nil {
		model.PrivateLinkServiceEnabled = core.BoolPtr(modelMap["private_link_service_enabled"].(bool))
	}
	if modelMap["public_enabled"] != nil {
		model.PublicEnabled = core.BoolPtr(modelMap["public_enabled"].(bool))
	}
	if modelMap["oracle_compatibility"] != nil {
		model.OracleCompatibility = core.BoolPtr(modelMap["oracle_compatibility"].(bool))
	}
	if modelMap["replicas"] != nil {
		model.Replicas = core.Int64Ptr(int64(modelMap["replicas"].(int)))
	}
	if modelMap["instance_type"] != nil && modelMap["instance_type"].(string) != "" {
		model.InstanceType = core.StringPtr(modelMap["instance_type"].(string))
	}
	return model, nil
}

func ResourceIbmByocEngineMapToCreateEngineBaseRequestCreateNetezzaEngineRequest(modelMap map[string]interface{}) (*brokerapiv1.CreateEngineBaseRequestCreateNetezzaEngineRequest, error) {
	model := &brokerapiv1.CreateEngineBaseRequestCreateNetezzaEngineRequest{}
	if modelMap["engine_type"] != nil && modelMap["engine_type"].(string) != "" {
		model.EngineType = core.StringPtr(modelMap["engine_type"].(string))
	}
	if modelMap["engine_name"] != nil && modelMap["engine_name"].(string) != "" {
		model.EngineName = core.StringPtr(modelMap["engine_name"].(string))
	}
	if modelMap["admin_username"] != nil && modelMap["admin_username"].(string) != "" {
		model.AdminUsername = core.StringPtr(modelMap["admin_username"].(string))
	}
	if modelMap["admin_password"] != nil && modelMap["admin_password"].(string) != "" {
		model.AdminPassword = core.StringPtr(modelMap["admin_password"].(string))
	}
	if modelMap["admin_email"] != nil && modelMap["admin_email"].(string) != "" {
		model.AdminEmail = core.StringPtr(modelMap["admin_email"].(string))
	}
	if modelMap["availability_zone"] != nil && modelMap["availability_zone"].(string) != "" {
		model.AvailabilityZone = core.StringPtr(modelMap["availability_zone"].(string))
	}
	if modelMap["storage_units"] != nil {
		model.StorageUnits = core.Int64Ptr(int64(modelMap["storage_units"].(int)))
	}
	if modelMap["compute_units"] != nil {
		model.ComputeUnits = core.Int64Ptr(int64(modelMap["compute_units"].(int)))
	}
	if modelMap["endpoint_type"] != nil && modelMap["endpoint_type"].(string) != "" {
		model.EndpointType = core.StringPtr(modelMap["endpoint_type"].(string))
	}
	if modelMap["profile_name"] != nil && modelMap["profile_name"].(string) != "" {
		model.ProfileName = core.StringPtr(modelMap["profile_name"].(string))
	}
	if modelMap["service_principals"] != nil {
		servicePrincipals := []string{}
		for _, servicePrincipalsItem := range modelMap["service_principals"].([]interface{}) {
			servicePrincipals = append(servicePrincipals, servicePrincipalsItem.(string))
		}
		model.ServicePrincipals = servicePrincipals
	}
	if modelMap["subscription_ids"] != nil {
		subscriptionIds := []strfmt.UUID{}
		for _, subscriptionIdsItem := range modelMap["subscription_ids"].([]interface{}) {
			subscriptionIds = append(subscriptionIds, strfmt.UUID(subscriptionIdsItem.(string)))
		}
		model.SubscriptionIds = subscriptionIds
	}
	return model, nil
}
