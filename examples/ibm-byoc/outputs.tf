// ============================================================================
// Outputs - Display important information about the engines
// ============================================================================

// Uncomment these outputs when you create an engine resource
// output "byoc_engine_id" {
//   description = "The composite ID of the created BYOC engine (subscription_id/dataplane_id/engine_id)"
//   value       = ibm_byoc_engine.byoc_engine_instance.id
// }

// output "byoc_engine_id_only" {
//   description = "The engine ID extracted from the composite ID"
//   value       = split("/", ibm_byoc_engine.byoc_engine_instance.id)[2]
// }

// output "byoc_engine_status" {
//   description = "The status of the created BYOC engine"
//   value       = ibm_byoc_engine.byoc_engine_instance.engine_status
// }

// output "byoc_engine_name" {
//   description = "The name of the created BYOC engine"
//   value       = ibm_byoc_engine.byoc_engine_instance.engine_name
// }

// ============================================================================
// Data Source Outputs - Full JSON Response
// ============================================================================

output "byoc_engine_full_response" {
  description = "JSON response from the BYOC engine data source"
  value       = data.ibm_byoc_engine.byoc_engine_data
}

# required for list engines
# output "all_engines_count" {
#   description = "Total number of engines in the dataplane"
#   value       = length(data.ibm_byoc_engines.byoc_engines_list.engines)
# }

# output "all_engines" {
#   description = "Complete list of all engines with all details"
#   value       = data.ibm_byoc_engines.byoc_engines_list.engines
# }
