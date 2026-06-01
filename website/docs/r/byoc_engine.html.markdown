---
layout: "ibm"
page_title: "IBM : ibm_byoc_engine"
description: |-
  Manages byoc_engine.
subcategory: "Broker Api"
---

# ibm_byoc_engine

Create, update, and delete byoc_engines with this resource.

## Example Usage

```hcl
resource "ibm_byoc_engine" "byoc_engine_instance" {
  availability_zone = "us-south-1"
  compute_units = 2
  dataplane_id = 1e524f4a-d11b-4b16-8a3b-e5e703899a03
  endpoint_type = "private"
  engine_name = "db2-engine"
  engine_type = "db2"
  instance_type = "Standard_D4s_v5"
  oracle_compatibility = false
  plan = "db2wh-small"
  private_link_service_enabled = true
  profile_name = "netezza-profile"
  public_enabled = false
  replicas = 1
  storage_units = 50
  subscription_id = 1e524f4a-d11b-4b16-8a3b-e5e703899a03
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `availability_zone` - (Optional, Forces new resource, String) 
* `compute_units` - (Optional, Forces new resource, Integer) 
* `dataplane_id` - (Required, Forces new resource, String) Data Plane ID.
* `endpoint_type` - (Optional, Forces new resource, String) 
  * Constraints: Allowable values are: `public`, `private`.
* `engine_name` - (Optional, Forces new resource, String) 
* `engine_type` - (Optional, Forces new resource, String) 
  * Constraints: Allowable values are: `db2`, `db2wh`, `netezza-aws`, `netezza-azure`.
* `instance_type` - (Optional, Forces new resource, String) 
* `oracle_compatibility` - (Optional, Forces new resource, Boolean) 
* `plan` - (Optional, Forces new resource, String) 
* `private_link_service_enabled` - (Optional, Forces new resource, Boolean) 
* `profile_name` - (Optional, Forces new resource, String) 
* `public_enabled` - (Optional, Forces new resource, Boolean) 
* `replicas` - (Optional, Forces new resource, Integer) 
* `service_principals` - (Optional, Forces new resource, List) 
* `storage_units` - (Optional, Forces new resource, Integer) 
* `subscription_id` - (Required, Forces new resource, String) Subscription ID.
* `subscription_ids` - (Optional, Forces new resource, List) 

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the byoc_engine.
* `created_at` - (String) 
* `engine_short_id` - (String) 
* `engine_status` - (String) 
* `engine_status_message` - (String) 
* `engine_ui_endpoint` - (List) 
Nested schema for **engine_ui_endpoint**:
	* `public` - (String) Public URL for the engine UI console.
* `jdbc_endpoint` - (List) 
Nested schema for **jdbc_endpoint**:
	* `private` - (String) Private JDBC connection string.
* `link_type` - (String) Type of engine link relationship (e.g., 'DR' for disaster recovery, 'STANDALONE' for standalone engines).
* `linked_dataplane_id` - (String) ID of the dataplane where the linked engine resides (NULL if linked engine status is DELETE_COMPLETE).
* `linked_engine_id` - (String) ID of the linked engine (NULL if linked engine status is DELETE_COMPLETE).
* `linked_engine_status` - (String) Status of the linked engine (NULL if linked engine status is DELETE_COMPLETE).
* `tags` - (List) 
* `version` - (String) 


## Import

You can import the `ibm_byoc_engine` resource by using `engine_id`.
The `engine_id` property can be formed from `subscription_id`, and `dataplane_id` in the following format:

<pre>
&lt;subscription_id&gt;/&lt;dataplane_id&gt;
</pre>
* `subscription_id`: A strfmt.UUID in the format `1e524f4a-d11b-4b16-8a3b-e5e703899a03`. Subscription ID.
* `dataplane_id`: A strfmt.UUID in the format `1e524f4a-d11b-4b16-8a3b-e5e703899a03`. Data Plane ID.

# Syntax
<pre>
$ terraform import ibm_byoc_engine.byoc_engine &lt;subscription_id&gt;/&lt;dataplane_id&gt;
</pre>
