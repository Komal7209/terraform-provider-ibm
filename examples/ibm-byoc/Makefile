# ============================================================================
# IBM BYOC Terraform Provider - Example Makefile
# ============================================================================
# This Makefile provides convenient commands for building the provider,
# initializing Terraform, and managing infrastructure.
#
# Usage: make [target]
# Run 'make help' to see all available commands
# ============================================================================

# Variables
PROVIDER_VERSION := 1.0.14
OS_ARCH := $(shell go env GOOS)_$(shell go env GOARCH)
PLUGIN_DIR := $(HOME)/.terraform.d/plugins/registry.terraform.io/ibm-cloud/ibm/$(PROVIDER_VERSION)/$(OS_ARCH)
ROOT_DIR := ../..
TERRAFORM := terraform
GO := go

# Colors for output
COLOR_RESET := \033[0m
COLOR_BOLD := \033[1m
COLOR_GREEN := \033[32m
COLOR_YELLOW := \033[33m
COLOR_BLUE := \033[34m
COLOR_CYAN := \033[36m

# Default target
.DEFAULT_GOAL := help

# ============================================================================
# Help Target - Display all available commands
# ============================================================================

.PHONY: help
help: ## Show this help message
	@echo "$(COLOR_BOLD)$(COLOR_CYAN)═══════════════════════════════════════════════════════════════════════$(COLOR_RESET)"
	@echo "$(COLOR_BOLD)$(COLOR_CYAN)  IBM BYOC Terraform Provider - Available Commands$(COLOR_RESET)"
	@echo "$(COLOR_BOLD)$(COLOR_CYAN)═══════════════════════════════════════════════════════════════════════$(COLOR_RESET)"
	@echo ""
	@echo "$(COLOR_BOLD)$(COLOR_GREEN)🔨 Build Commands:$(COLOR_RESET)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; /build|install|clean/ {printf "  $(COLOR_YELLOW)%-20s$(COLOR_RESET) %s\n", $$1, $$2}'
	@echo ""
	@echo "$(COLOR_BOLD)$(COLOR_GREEN)🚀 Terraform Commands:$(COLOR_RESET)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; /init-locally|plan|apply|destroy|validate|fmt|output/ {printf "  $(COLOR_YELLOW)%-20s$(COLOR_RESET) %s\n", $$1, $$2}'
	@echo ""
	@echo "$(COLOR_BOLD)$(COLOR_GREEN)🧪 Test Commands:$(COLOR_RESET)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; /test/ {printf "  $(COLOR_YELLOW)%-20s$(COLOR_RESET) %s\n", $$1, $$2}'
	@echo ""
	@echo "$(COLOR_BOLD)$(COLOR_GREEN)🔄 Workflow Commands:$(COLOR_RESET)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; /all|refresh|reset/ {printf "  $(COLOR_YELLOW)%-20s$(COLOR_RESET) %s\n", $$1, $$2}'
	@echo ""
	@echo "$(COLOR_BOLD)$(COLOR_BLUE)Examples:$(COLOR_RESET)"
	@echo "  $(COLOR_CYAN)make all$(COLOR_RESET)              # Build provider and deploy infrastructure"
	@echo "  $(COLOR_CYAN)make build-and-init$(COLOR_RESET)   # Build provider and initialize Terraform"
	@echo "  $(COLOR_CYAN)make plan$(COLOR_RESET)             # Preview infrastructure changes"
	@echo "  $(COLOR_CYAN)make apply$(COLOR_RESET)            # Apply infrastructure changes"
	@echo "  $(COLOR_CYAN)make destroy$(COLOR_RESET)          # Destroy all infrastructure"
	@echo ""
	@echo "$(COLOR_BOLD)$(COLOR_BLUE)Prerequisites:$(COLOR_RESET)"
	@echo "  • Go $(shell go version | awk '{print $$3}')"
	@echo "  • Terraform $(shell terraform version -json 2>/dev/null | grep -o '"terraform_version":"[^"]*' | cut -d'"' -f4 || echo 'not installed')"
	@echo "  • IBM Cloud credentials configured"
	@echo ""

# ============================================================================
# Build Commands
# ============================================================================

.PHONY: build
build: ## Build the IBM Terraform provider from source
	@echo "$(COLOR_BOLD)$(COLOR_GREEN)🔨 Building IBM Terraform Provider...$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Changing to root directory: $(ROOT_DIR)$(COLOR_RESET)"
	@cd $(ROOT_DIR) && $(MAKE) build
	@echo "$(COLOR_GREEN)✓ Provider built successfully$(COLOR_RESET)"

.PHONY: build-local
build-local: ## Build and install provider locally for testing
	@echo "$(COLOR_BOLD)$(COLOR_GREEN)🔨 Building and installing provider locally...$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Changing to root directory: $(ROOT_DIR)$(COLOR_RESET)"
	@cd $(ROOT_DIR) && $(MAKE) build-local
	@echo "$(COLOR_GREEN)✓ Provider installed to: $(PLUGIN_DIR)$(COLOR_RESET)"

.PHONY: install
install: build-local ## Alias for build-local (build and install provider)
	@echo "$(COLOR_GREEN)✓ Provider ready for use$(COLOR_RESET)"

.PHONY: clean
clean: ## Remove Terraform state and cache files
	@echo "$(COLOR_BOLD)$(COLOR_YELLOW)🧹 Cleaning Terraform files...$(COLOR_RESET)"
	@rm -rf .terraform .terraform.lock.hcl terraform.tfstate terraform.tfstate.backup
	@echo "$(COLOR_GREEN)✓ Cleanup complete$(COLOR_RESET)"

# ============================================================================
# Terraform Commands
# ============================================================================

.PHONY: init-locally 
init-locally: ## Initialize Terraform (download providers and modules)
	@echo "$(COLOR_BOLD)$(COLOR_BLUE)🚀 Initializing Terraform...$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Running in current directory: examples/ibm-byoc$(COLOR_RESET)"
	@$(TERRAFORM) init -upgrade -plugin-dir=$(HOME)/.terraform.d/plugins
	@echo "$(COLOR_GREEN)✓ Terraform initialized$(COLOR_RESET)"

.PHONY: validate
validate: ## Validate Terraform configuration files
	@echo "$(COLOR_BOLD)$(COLOR_BLUE)✓ Validating Terraform configuration...$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Running in current directory: examples/ibm-byoc$(COLOR_RESET)"
	@$(TERRAFORM) validate
	@echo "$(COLOR_GREEN)✓ Configuration is valid$(COLOR_RESET)"

# .PHONY: fmt
# fmt: ## Format Terraform configuration files
# 	@echo "$(COLOR_BOLD)$(COLOR_BLUE)📝 Formatting Terraform files...$(COLOR_RESET)"
# 	@echo "$(COLOR_BLUE)→ Running in current directory: examples/ibm-byoc$(COLOR_RESET)"
# 	@$(TERRAFORM) fmt -recursive
# 	@echo "$(COLOR_GREEN)✓ Files formatted$(COLOR_RESET)"

.PHONY: plan
plan: ## Show Terraform execution plan (preview changes)
	@echo "$(COLOR_BOLD)$(COLOR_BLUE)📋 Generating Terraform plan...$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Running in current directory: examples/ibm-byoc$(COLOR_RESET)"
	@$(TERRAFORM) plan
	@echo "$(COLOR_GREEN)✓ Plan generated$(COLOR_RESET)"

.PHONY: apply
apply: ## Apply Terraform changes (create/update infrastructure)
	@echo "$(COLOR_BOLD)$(COLOR_GREEN)🚀 Applying Terraform changes...$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Running in current directory: examples/ibm-byoc$(COLOR_RESET)"
	@$(TERRAFORM) apply
	@echo "$(COLOR_GREEN)✓ Changes applied successfully$(COLOR_RESET)"

.PHONY: apply-auto
apply-auto: ## Apply Terraform changes without confirmation prompt
	@echo "$(COLOR_BOLD)$(COLOR_GREEN)🚀 Auto-applying Terraform changes...$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Running in current directory: examples/ibm-byoc$(COLOR_RESET)"
	@$(TERRAFORM) apply -auto-approve
	@echo "$(COLOR_GREEN)✓ Changes applied successfully$(COLOR_RESET)"

.PHONY: destroy
destroy: ## Destroy all Terraform-managed infrastructure
	@echo "$(COLOR_BOLD)$(COLOR_YELLOW)⚠️  Destroying infrastructure...$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Running in current directory: examples/ibm-byoc$(COLOR_RESET)"
	@$(TERRAFORM) destroy
	@echo "$(COLOR_GREEN)✓ Infrastructure destroyed$(COLOR_RESET)"

.PHONY: destroy-auto
destroy-auto: ## Destroy infrastructure without confirmation prompt
	@echo "$(COLOR_BOLD)$(COLOR_YELLOW)⚠️  Auto-destroying infrastructure...$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Running in current directory: examples/ibm-byoc$(COLOR_RESET)"
	@$(TERRAFORM) destroy -auto-approve
	@echo "$(COLOR_GREEN)✓ Infrastructure destroyed$(COLOR_RESET)"

.PHONY: output
output: ## Show Terraform output values
	@echo "$(COLOR_BOLD)$(COLOR_BLUE)📤 Terraform Outputs:$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Running in current directory: examples/ibm-byoc$(COLOR_RESET)"
	@$(TERRAFORM) output

.PHONY: show
show: ## Show current Terraform state
	@echo "$(COLOR_BOLD)$(COLOR_BLUE)📊 Current Terraform State:$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Running in current directory: examples/ibm-byoc$(COLOR_RESET)"
	@$(TERRAFORM) show

.PHONY: refresh
refresh: ## Refresh Terraform state to match real infrastructure
	@echo "$(COLOR_BOLD)$(COLOR_BLUE)🔄 Refreshing Terraform state...$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Running in current directory: examples/ibm-byoc$(COLOR_RESET)"
	@$(TERRAFORM) refresh
	@echo "$(COLOR_GREEN)✓ State refreshed$(COLOR_RESET)"

# ============================================================================
# Test Commands
# ============================================================================

.PHONY: test-provider
test-provider: ## Run provider unit tests
	@echo "$(COLOR_BOLD)$(COLOR_CYAN)🧪 Running provider tests...$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Changing to root directory: $(ROOT_DIR)$(COLOR_RESET)"
	@cd $(ROOT_DIR) && $(GO) test ./ibm/service/byoc/... -v
	@echo "$(COLOR_GREEN)✓ Tests completed$(COLOR_RESET)"

.PHONY: test-integration
test-integration: ## Run integration tests for BYOC service
	@echo "$(COLOR_BOLD)$(COLOR_CYAN)🧪 Running integration tests...$(COLOR_RESET)"
	@echo "$(COLOR_BLUE)→ Changing to root directory: $(ROOT_DIR)$(COLOR_RESET)"
	@cd $(ROOT_DIR) && TF_ACC=1 $(GO) test ./ibm/service/byoc/... -v -timeout 120m
	@echo "$(COLOR_GREEN)✓ Integration tests completed$(COLOR_RESET)"

# ============================================================================
# Workflow Commands (Combined Operations)
# ============================================================================

.PHONY: build-and-init
build-and-init: build-local clean init ## Build provider, clean, and initialize Terraform
	@echo "$(COLOR_GREEN)✓ Ready to run terraform plan$(COLOR_RESET)"

.PHONY: all
all: build-and-init validate plan apply ## Complete workflow: build, init, validate, plan, and apply
	@echo "$(COLOR_BOLD)$(COLOR_GREEN)✓ Complete deployment finished!$(COLOR_RESET)"

.PHONY: reset
reset: destroy clean ## Destroy infrastructure and clean all Terraform files
	@echo "$(COLOR_GREEN)✓ Environment reset complete$(COLOR_RESET)"

.PHONY: redeploy
redeploy: destroy-auto apply-auto ## Destroy and recreate infrastructure (auto-approve)
	@echo "$(COLOR_GREEN)✓ Redeployment complete$(COLOR_RESET)"

# ============================================================================
# Development Commands
# ============================================================================

.PHONY: dev-setup
dev-setup: build-local init ## Quick setup for development (build and init)
	@echo "$(COLOR_GREEN)✓ Development environment ready$(COLOR_RESET)"

.PHONY: quick-test
quick-test: validate plan ## Quick validation and plan check
	@echo "$(COLOR_GREEN)✓ Quick test complete$(COLOR_RESET)"

# ============================================================================
# Information Commands
# ============================================================================

.PHONY: version
version: ## Show provider and Terraform versions
	@echo "$(COLOR_BOLD)$(COLOR_BLUE)Version Information:$(COLOR_RESET)"
	@echo "  Provider Version: $(PROVIDER_VERSION)"
	@echo "  OS/Arch: $(OS_ARCH)"
	@echo "  Plugin Directory: $(PLUGIN_DIR)"
	@echo "  Go Version: $(shell go version)"
	@echo "  Terraform Version: $(shell terraform version | head -n1)"

.PHONY: status
status: ## Show current infrastructure status
	@echo "$(COLOR_BOLD)$(COLOR_BLUE)📊 Infrastructure Status:$(COLOR_RESET)"
	@if [ -f terraform.tfstate ]; then \
		echo "$(COLOR_GREEN)✓ State file exists$(COLOR_RESET)"; \
		$(TERRAFORM) show -json | jq -r '.values.root_module.resources[] | select(.mode=="managed") | "  • \(.type).\(.name) [\(.values.engine_status // "N/A")]"' 2>/dev/null || echo "  State file present but no resources found"; \
	else \
		echo "$(COLOR_YELLOW)⚠ No state file found - infrastructure not deployed$(COLOR_RESET)"; \
	fi

# ============================================================================
# Utility Targets
# ============================================================================

.PHONY: check-prereqs
check-prereqs: ## Check if all prerequisites are installed
	@echo "$(COLOR_BOLD)$(COLOR_BLUE)Checking prerequisites...$(COLOR_RESET)"
	@command -v go >/dev/null 2>&1 || { echo "$(COLOR_YELLOW)⚠ Go is not installed$(COLOR_RESET)"; exit 1; }
	@echo "$(COLOR_GREEN)✓ Go is installed$(COLOR_RESET)"
	@command -v terraform >/dev/null 2>&1 || { echo "$(COLOR_YELLOW)⚠ Terraform is not installed$(COLOR_RESET)"; exit 1; }
	@echo "$(COLOR_GREEN)✓ Terraform is installed$(COLOR_RESET)"
	@command -v jq >/dev/null 2>&1 || echo "$(COLOR_YELLOW)⚠ jq is not installed (optional, for better status output)$(COLOR_RESET)"
	@echo "$(COLOR_GREEN)✓ All required prerequisites are installed$(COLOR_RESET)"

# Prevent make from treating files as targets
.PHONY: all build build-local install clean init validate fmt plan apply apply-auto \
        destroy destroy-auto output show refresh test-provider test-integration \
        build-and-init reset redeploy dev-setup quick-test version status check-prereqs help

