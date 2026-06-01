---
layout: "ibm"
page_title: "IBM : ibm_byoc_engine"
description: |-
  Get information about byoc_engine
subcategory: "Broker Api"
---

# ibm_byoc_engine

Provides a read-only data source to retrieve information about a BYOC engine. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_byoc_engine" "byoc_engine" {
  subscription_id = "1e524f4a-d11b-4b16-8a3b-e5e703899a03"
  dataplane_id    = "1e524f4a-d11b-4b16-8a3b-e5e703899a03"
  engine_id       = "2e524f4a-d11b-4b16-8a3b-e5e703899a04"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `subscription_id` - (Required, String) Subscription ID.
* `dataplane_id` - (Required, String) Data Plane ID.
* `engine_id` - (Required, String) Engine ID.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the byoc_engine.
* `availability_zone` - (String) Availability zone for the engine.
* `compute_units` - (Integer) Number of compute units.
* `created_at` - (String) Creation timestamp.
* `endpoint_type` - (String) Endpoint type for the engine.
  * Constraints: Allowable values are: `private`, `public`.
* `engine_name` - (String) Name of the engine.
* `engine_short_id` - (String) Short ID of the engine.
* `engine_status` - (String) Current status of the engine.
* `engine_status_message` - (String) Status message for the engine.
* `engine_type` - (String) Type of the engine.
  * Constraints: Allowable values are: `db2`.
* `engine_ui_endpoint` - (List) Engine UI endpoint.
  
  Nested schema for `engine_ui_endpoint`:
  * `public` - (String) Public URL for the engine UI console.

* `instance_type` - (String) Instance type for the engine.
* `jdbc_endpoint` - (List) JDBC endpoint.
  
  Nested schema for `jdbc_endpoint`:
  * `private` - (String) Private JDBC connection string.

* `link_type` - (String) Type of engine link relationship (e.g., 'DR' for disaster recovery, 'STANDALONE' for standalone engines).
* `linked_dataplane_id` - (String) ID of the dataplane where the linked engine resides.
* `linked_engine_id` - (String) ID of the linked engine.
* `linked_engine_status` - (String) Status of the linked engine.
* `oracle_compatibility` - (Boolean) Whether Oracle compatibility is enabled.
* `plan` - (String) Plan for the engine (e.g., for DB2 Warehouse).
* `private_link_service_enabled` - (Boolean) Whether private link service is enabled.
* `profile_name` - (String) Profile name (for Netezza engines).
* `public_enabled` - (Boolean) Whether public access is enabled.
* `replicas` - (Integer) Number of replicas.
* `service_principals` - (List) Service principals (for Netezza on AWS).
* `storage_units` - (Integer) Number of storage units.
* `subscription_ids` - (List) Subscription IDs (for Netezza on Azure).
* `tags` - (List) Tags associated with the engine.
* `version` - (String) Version of the engine.