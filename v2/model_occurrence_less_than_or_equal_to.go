/*
 * Nodeum API
 *
 * # About  This document describes the Nodeum API version 2:  If you are looking for any information about the product itself, reach the product website https://www.nodeum.io. You can also contact us at this email address : info@nodeum.io  # Filter parameters When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)  
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nodeum
// OccurrenceLessThanOrEqualTo Must have at most or less than %{count} occurrences (currently have %{value})
type OccurrenceLessThanOrEqualTo struct {
	Error string `json:"error"`
	// Expected maximum number of occurrences
	Count int32 `json:"count,omitempty"`
	// Current number of occurrences
	Value int32 `json:"value,omitempty"`
}
