# terraform {
#   required_version = ">= 1.0"
#   required_providers {
#     ibm = {
#       source  = "IBM-Cloud/ibm"
#       version = ">= 1.12.0"
#     }
#   }
# }

terraform { 
  required_providers { 
    ibm = { 
      source  = "IBM-Cloud/ibm" 
      version = "2.20.2" 
    } 
  } 
} 