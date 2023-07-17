/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A strategy constraint. For more information, refer to [the strategy constraint reference documentation](https://docs.getunleash.io/reference/strategy-constraints)
type ConstraintSchema struct {
	// The name of the context field that this constraint should apply to.
	ContextName string `json:"contextName"`
	// The operator to use when evaluating this constraint. For more information about the various operators, refer to [the strategy constraint operator documentation](https://docs.getunleash.io/reference/strategy-constraints#strategy-constraint-operators).
	Operator string `json:"operator"`
	// Whether the operator should be case sensitive or not. Defaults to `false` (being case sensitive).
	CaseInsensitive bool `json:"caseInsensitive,omitempty"`
	// Whether the result should be negated or not. If `true`, will turn a `true` result into a `false` result and vice versa.
	Inverted bool `json:"inverted,omitempty"`
	// The context values that should be used for constraint evaluation. Use this property instead of `value` for properties that accept multiple values.
	Values []string `json:"values,omitempty"`
	// The context value that should be used for constraint evaluation. Use this property instead of `values` for properties that only accept single values.
	Value string `json:"value,omitempty"`
}
