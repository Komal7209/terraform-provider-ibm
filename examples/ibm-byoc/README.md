# Examples for Byoc Broker Api

These examples illustrate how to use the resources and data sources associated with Broker Api.

The following resources are supported:
* ibm_byoc_engine

## Usage

**Important**: Before running this example, you must set the BYOC bearer token as an environment variable:

```bash
export BYOC_BEARER_TOKEN="your-bearer-token-here"
# or
export IBMCLOUD_BYOC_TOKEN="your-bearer-token-here"
```

Then execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Broker Api resources

### Resource: ibm_byoc_engine

```hcl
resource "ibm_byoc_engine" "byoc_engine_instance" {
  subscription_id = var.byoc_engine_subscription_id
  dataplane_id = var.byoc_engine_dataplane_id
  availability_zone = var.byoc_engine_availability_zone
  storage_units = var.byoc_engine_storage_units
  compute_units = var.byoc_engine_compute_units
  engine_name = var.byoc_engine_engine_name
  engine_type = var.byoc_engine_engine_type
  admin_username = var.byoc_engine_admin_username
  admin_password = var.byoc_engine_admin_password
  admin_email = var.byoc_engine_admin_email
  endpoint_type = var.byoc_engine_endpoint_type
  instance_type = var.byoc_engine_instance_type
  replicas = var.byoc_engine_replicas
  public_enabled = var.byoc_engine_public_enabled
  private_link_service_enabled = var.byoc_engine_private_link_service_enabled
  oracle_compatibility = var.byoc_engine_oracle_compatibility
  plan = var.byoc_engine_plan
  profile_name = var.byoc_engine_profile_name
  service_principals = var.byoc_engine_service_principals
  subscription_ids = var.byoc_engine_subscription_ids
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| subscription_id | Subscription ID | `string` | true |
| dataplane_id | Data Plane ID | `string` | true |
| engine_type | Type of the engine (db2, db2wh, netezza) | `string` | true |
| admin_username | Admin username for the engine | `string` | true |
| admin_password | Admin password for the engine (SHA-2 hashed) | `string` | true |
| admin_email | Admin email for the engine | `string` | true |
| instance_type | Azure instance type (e.g., Standard_D4s_v5) | `string` | true |
| replicas | Number of replicas | `number` | true |
| private_link_service_enabled | Enable private link service | `bool` | true |
| public_enabled | Enable public access | `bool` | true |
| availability_zone | Availability zone | `string` | false |
| storage_units | Number of storage units | `number` | false |
| compute_units | Number of compute units | `number` | false |
| engine_name | Name of the engine | `string` | false |
| endpoint_type | Endpoint type | `string` | false |
| oracle_compatibility | Enable Oracle compatibility | `bool` | false |
| plan | Plan name | `string` | false |
| profile_name | Profile name | `string` | false |
| service_principals | Service principals | `list(string)` | false |
| subscription_ids | Subscription IDs | `list(string)` | false |

#### Outputs

| Name | Description |
|------|-------------|
| tags |  |
| engine_status |  |
| engine_status_message |  |
| created_at |  |
| engine_short_id |  |
| engine_ui_endpoint |  |
| jdbc_endpoint |  |
| linked_engine_id | ID of the linked engine (NULL if linked engine status is DELETE_COMPLETE). |
| link_type | Type of engine link relationship (e.g., 'DR' for disaster recovery, 'STANDALONE' for standalone engines). |
| linked_dataplane_id | ID of the dataplane where the linked engine resides (NULL if linked engine status is DELETE_COMPLETE). |
| linked_engine_status | Status of the linked engine (NULL if linked engine status is DELETE_COMPLETE). |
| version |  |


## Assumptions

1. You have a valid IBM Cloud account with access to BYOC services
2. You have obtained a BYOC bearer token and set it as an environment variable
3. You have valid subscription_id and dataplane_id values

## Notes

1. The `admin_password` must be SHA-2 hashed
2. The `engine_type` field is required and determines the type of engine to create (db2, db2wh, or netezza)
3. Admin credentials (`admin_username`, `admin_password`, `admin_email`) are required for engine creation
4. Engines take time to provision and cannot be deleted while in IN_PROGRESS state
5. Use valid Azure instance types for the `instance_type` field (e.g., Standard_D4s_v5)

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |
