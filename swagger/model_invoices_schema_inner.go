/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InvoicesSchemaInner struct {
	AmountFormatted string `json:"amountFormatted"`
	Paid bool `json:"paid"`
	Status string `json:"status"`
	DueDate string `json:"dueDate,omitempty"`
	InvoiceURL string `json:"invoiceURL"`
	InvoicePDF string `json:"invoicePDF"`
}
