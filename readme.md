# HCL v2 parser for JS

This is a [HCL version 2](https://pkg.go.dev/github.com/hashicorp/hcl/v2#readme-changes-in-2-0) parser for JavaScript, there are several other JS parsers for HCL but none that support the newer HCL v2 syntax

It wraps the very helpful [tmccombs/hcl2json](https://github.com/tmccombs/hcl2json) and calls the convert package in order to parse HCL input strings to JSON strings. The Go code in `parser.go` is converted to JS using [GopherJS](https://github.com/gopherjs/gopherjs)

TypeScript definitions are included

## Reference

The module exports the following functions:

```ts
function parseToString(input: string): string
function parseToObject(input: string): any
```

## Usage

Install as normal with NPM

```bash
npm install hcl2-parser
```

Importing into your project

```ts
// Good old fashioned Node.js CommonJS require
const hcl = require("hcl2-parser")

// Import with ES6 or TypeScript
import * as hcl = from "hcl2-parser"
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
