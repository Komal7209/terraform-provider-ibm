# IBM BYOC (Bring Your Own Cloud) Terraform Example

This example demonstrates how to use the IBM BYOC Terraform provider to manage dataplanes and database engines across multiple cloud providers (AWS, Azure, GCP).

## Overview

The IBM BYOC service allows you to deploy IBM database engines (Db2, Db2 Warehouse, Netezza) in your own cloud infrastructure while maintaining IBM Cloud management capabilities.

## Supported Configurations

### Cloud Providers
- **AWS** - Amazon Web Services
- **Azure** - Microsoft Azure
- **GCP** - Google Cloud Platform (coming soon)

### Database Engines
- **Db2** - IBM Db2 Database
- **Db2 Warehouse** - IBM Db2 Warehouse
- **Netezza** - IBM Netezza Performance Server

## Prerequisites

1. **IBM Cloud Account** with BYOC service access
2. **Cloud Provider Account** (AWS or Azure)
3. **Terraform** >= 1.0
4. **Provider Credentials**:
   - For Azure: Service Principal with appropriate permissions
   - For AWS: IAM Role with cross-account access

## Quick Start - Azure with Db2

### Step 1: Setup

```bash
cd examples/ibm-byoc
cp variables.tfvars.example terraform.tfvars
```

### Step 2: Configure Variables

Edit `terraform.tfvars` with your Azure and IBM Cloud details:

```hcl
# IBM Cloud
subscription_id = "your-ibm-subscription-uuid"
dataplane_id    = "your-dataplane-uuid"
dataplane_name  = "my-azure-dataplane"
engine_id       = "your-engine-uuid"
engine_name     = "my-db2-engine"

# Azure Configuration
azure_region          = "eastus"
azure_subscription_id = "12345678-1234-1234-1234-123456789012"
azure_tenant_id       = "87654321-4321-4321-4321-210987654321"
azure_client_id       = "abcdef12-3456-7890-abcd-ef1234567890"
azure_client_secret   = "your-azure-client-secret"

# Engine Configuration
availability_zone = "eastus-1"
storage_units     = 100
compute_units     = 4
instance_type     = "Standard_D4s_v3"
```

### Step 3: Initialize and Apply

```bash
terraform init
terraform plan
terraform apply
```

## CRUD Operations

### CREATE
Deploy a new dataplane and engine:
```bash
terraform apply
```

The provider will:
1. **Update hyperscaler account** credentials
2. **Check prerequisites** (quotas, permissions)
3. **Verify credentials** (Azure only - explicit verification before creation)
4. **Create dataplane** resources in your cloud
5. **Deploy engine** on the dataplane

### READ
View current state:
```bash
# Show all resources
terraform show

# Show specific resources
terraform state show ibm_byoc_dataplane.byoc_dataplane
terraform state show ibm_byoc_engine.db2_engine

# Use data sources
terraform console
> data.ibm_byoc_dataplane.dataplane_info
> data.ibm_byoc_engines.all_engines
```

### UPDATE
Modify configuration and apply changes:
```bash
# Edit terraform.tfvars (e.g., change compute_units from 4 to 8)
terraform apply
```

### DELETE
Remove all resources:
```bash
terraform destroy
```

## Hyperscaler-Specific Flows

### AWS Flow
1. Update hyperscaler account (AWS account ID and role ARN)
2. Check prerequisites (KMS keys, quotas)
3. Create dataplane (verification happens internally)

### Azure Flow
1. Update hyperscaler account (subscription ID, tenant ID)
2. Check prerequisites (credentials, vCPU quotas)
3. **Verify credentials explicitly** (required before creation)
4. Create dataplane

## Testing Different Configurations

### Test AWS with Db2

1. In `main.tf`, comment out Azure dataplane and uncomment AWS dataplane
2. Update `terraform.tfvars` with AWS credentials:
```hcl
aws_region     = "us-east-1"
aws_account_id = "123456789012"
aws_role_arn   = "arn:aws:iam::123456789012:role/byoc-role"
```
3. Run `terraform apply`

### Test Db2 Warehouse

1. In `main.tf`, uncomment the `db2wh_engine` resource
2. Add to `terraform.tfvars`:
```hcl
db2wh_engine_id = "your-db2wh-engine-uuid"
```
3. Run `terraform apply`

### Test Netezza

1. In `main.tf`, uncomment the `netezza_engine` resource
2. Add to `terraform.tfvars`:
```hcl
netezza_engine_id = "your-netezza-engine-uuid"
```
3. Run `terraform apply`

## File Structure

```
examples/ibm-byoc/
├── main.tf                      # Main configuration with all resources
├── variables.tf                 # Variable definitions
├── variables.tfvars.example     # Example values (copy to terraform.tfvars)
├── provider.tf                  # Provider configuration
├── versions.tf                  # Terraform and provider version constraints
├── outputs.tf                   # Output definitions
└── README.md                    # This file
```

## Important Notes

### UUID Requirements
- `subscription_id`, `dataplane_id`, and `engine_id` must be valid UUIDs
- Generate UUIDs using: `uuidgen` (macOS/Linux) or online UUID generators

### Authentication
- The BYOC API requires a **bearer token** for authentication (different from IBM Cloud IAM token)
- Set via environment variable: `export TF_VAR_bearer_token="your-bearer-token"`
- The bearer token is passed to the provider's `iam_token` parameter
- **Note**: Despite the parameter name, BYOC uses bearer token authentication, not standard IAM tokens

### Network Configuration
- **AWS**: Requires VPC CIDR and subnet configurations
- **Azure**: Network configuration handled automatically
- **Private endpoints**: Recommended for production deployments

### Timeouts
- Dataplane creation: 30 minutes (configurable)
- Engine creation: 60 minutes (configurable)
- Adjust timeouts based on your cloud provider and region

## Troubleshooting

### Common Issues

1. **"Invalid UUID format"**
   - Ensure all IDs are valid UUIDs
   - Check for extra spaces or invalid characters

2. **"Prerequisites check failed"**
   - Verify cloud provider quotas
   - Check IAM permissions
   - For Azure: Ensure service principal has Contributor role

3. **"Verification failed" (Azure)**
   - Verify tenant ID and subscription ID are correct
   - Check client ID and client secret
   - Ensure service principal is not expired

4. **"Timeout waiting for dataplane"**
   - Increase timeout in `timeouts` block
   - Check cloud provider console for resource creation status
   - Review provider logs for detailed error messages

### Debug Mode

Enable detailed logging:
```bash
export TF_LOG=DEBUG
terraform apply
```

## API Flow Details

### UpdateDataplane API Actions

The provider uses the `UpdateDataplane` (PUT) API with different `action` parameters:

1. **action=updateHyperscalerAccount**: Sets cloud credentials
2. **action=prereqs**: Validates prerequisites
3. **action=verify**: Verifies credentials (Azure only, before creation)
4. **No action**: Creates/updates actual dataplane resources

### Why No Separate Create API?

The BYOC API follows RESTful PUT semantics where PUT is idempotent and handles both create and update operations. This design:
- Simplifies the API surface
- Provides fine-grained control via action parameters
- Supports hyperscaler-specific flows naturally

## Additional Resources

- [IBM BYOC Documentation](https://cloud.ibm.com/docs/byoc)
- [Terraform IBM Provider](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs)
- [BYOC API Reference](https://cloud.ibm.com/apidocs/byoc)

## Support

For issues or questions:
1. Check the troubleshooting section above
2. Review Terraform and provider logs
3. Contact IBM Cloud Support
4. Open an issue in the provider repository

## License

This example is provided under the Mozilla Public License v2.0.