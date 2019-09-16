/*
 * Nodeum API
 *
 * Nodeum API  # Filter parameters When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)  
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nodeum

// Must have at most or less than %{count} occurrences (currently have %{value})
type OccurrenceLessThanOrEqualTo struct {
	Error_ string `json:"error"`
	// Expected maximum number of occurrences
	Count int32 `json:"count,omitempty"`
	// Current number of occurrences
	Value int32 `json:"value,omitempty"`
}
