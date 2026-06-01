# Examples for Broker Api

These examples illustrate how to use the resources and data sources associated with Broker Api.

The following resources are supported:
* ibm_byoc_engine

## Usage

To run this example, execute the following commands:

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
| subscription_id | Subscription ID. | `` | true |
| dataplane_id | Data Plane ID. | `` | true |
| availability_zone |  | `string` | false |
| storage_units |  | `number` | false |
| compute_units |  | `number` | false |
| engine_name |  | `string` | false |
| engine_type |  | `string` | false |
| endpoint_type |  | `string` | false |
| instance_type |  | `string` | false |
| replicas |  | `number` | false |
| public_enabled |  | `bool` | false |
| private_link_service_enabled |  | `bool` | false |
| oracle_compatibility |  | `bool` | false |
| plan |  | `string` | false |
| profile_name |  | `string` | false |
| service_principals |  | `list(string)` | false |
| subscription_ids |  | `list(string)` | false |

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

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |
