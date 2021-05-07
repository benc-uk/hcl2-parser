# HCLv2 parser for JS

This is a HCL v2 parser for JavaScript, there are several other JS parsers for HCL but none that support the newer HCL v2 syntax

# Usage

Install

```bash
npm install hcl2-parser
```

The module exports the following functions:

```ts
function parseToString(input: string): string
function parseToObject(input: string): any
```

Simple example of usage

```js
const hcl = require("hcl2-parser")

const hclString = `
# Create a resource group
variable "azureRegion" {
  type = string
  default = "uksouth"
}
resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = var.azureRegion
}
`

// Parse into a JSON string
stringResult = hcl.parseToString(hclString)
console.log(stringResult)

// Parse into an object, the actual result is in array index 0 for reasons I don't understand
objectResult = hcl.parseToObject(hclString)
console.log(objectResult[0].resource.azurerm_resource_group)
```
