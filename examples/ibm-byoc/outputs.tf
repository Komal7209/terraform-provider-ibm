// This output allows byoc_engine data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_byoc_engine" {
  value       = ibm_byoc_engine.byoc_engine_instance
  description = "byoc_engine resource instance"
}
