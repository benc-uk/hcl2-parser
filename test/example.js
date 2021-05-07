const hcl = require("../dist")

const input = `
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=2.46.0"
    }
  }
}

# Configure the Microsoft Azure Provider
provider "azurerm" {
  features {}
}

# Create a resource group
variable "azureRegion" {
  type = string
  default = "uksouth"
}
resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = var.azureRegion
}

# Create a virtual network within the resource group
resource "azurerm_virtual_network" "example" {
  name                = "example-network"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  address_space       = ["10.0.0.0/16"]
}`

result = hcl.parseToObject(input)

console.log(result)
console.log("=======================")
console.log(`azureRegion == ${result[0].variable.azureRegion[0].default}`)
console.log(result[0].resource.azurerm_resource_group)
