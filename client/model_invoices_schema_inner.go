/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the InvoicesSchemaInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoicesSchemaInner{}

// InvoicesSchemaInner struct for InvoicesSchemaInner
type InvoicesSchemaInner struct {
	AmountFormatted string  `json:"amountFormatted"`
	Paid            bool    `json:"paid"`
	Status          string  `json:"status"`
	DueDate         *string `json:"dueDate,omitempty"`
	InvoiceURL      string  `json:"invoiceURL"`
	InvoicePDF      string  `json:"invoicePDF"`
}

// NewInvoicesSchemaInner instantiates a new InvoicesSchemaInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoicesSchemaInner(amountFormatted string, paid bool, status string, invoiceURL string, invoicePDF string) *InvoicesSchemaInner {
	this := InvoicesSchemaInner{}
	this.AmountFormatted = amountFormatted
	this.Paid = paid
	this.Status = status
	this.InvoiceURL = invoiceURL
	this.InvoicePDF = invoicePDF
	return &this
}

// NewInvoicesSchemaInnerWithDefaults instantiates a new InvoicesSchemaInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicesSchemaInnerWithDefaults() *InvoicesSchemaInner {
	this := InvoicesSchemaInner{}
	return &this
}

// GetAmountFormatted returns the AmountFormatted field value
func (o *InvoicesSchemaInner) GetAmountFormatted() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmountFormatted
}

// GetAmountFormattedOk returns a tuple with the AmountFormatted field value
// and a boolean to check if the value has been set.
func (o *InvoicesSchemaInner) GetAmountFormattedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountFormatted, true
}

// SetAmountFormatted sets field value
func (o *InvoicesSchemaInner) SetAmountFormatted(v string) {
	o.AmountFormatted = v
}

// GetPaid returns the Paid field value
func (o *InvoicesSchemaInner) GetPaid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Paid
}

// GetPaidOk returns a tuple with the Paid field value
// and a boolean to check if the value has been set.
func (o *InvoicesSchemaInner) GetPaidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Paid, true
}

// SetPaid sets field value
func (o *InvoicesSchemaInner) SetPaid(v bool) {
	o.Paid = v
}

// GetStatus returns the Status field value
func (o *InvoicesSchemaInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InvoicesSchemaInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InvoicesSchemaInner) SetStatus(v string) {
	o.Status = v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *InvoicesSchemaInner) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesSchemaInner) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *InvoicesSchemaInner) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *InvoicesSchemaInner) SetDueDate(v string) {
	o.DueDate = &v
}

// GetInvoiceURL returns the InvoiceURL field value
func (o *InvoicesSchemaInner) GetInvoiceURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceURL
}

// GetInvoiceURLOk returns a tuple with the InvoiceURL field value
// and a boolean to check if the value has been set.
func (o *InvoicesSchemaInner) GetInvoiceURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceURL, true
}

// SetInvoiceURL sets field value
func (o *InvoicesSchemaInner) SetInvoiceURL(v string) {
	o.InvoiceURL = v
}

// GetInvoicePDF returns the InvoicePDF field value
func (o *InvoicesSchemaInner) GetInvoicePDF() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoicePDF
}

// GetInvoicePDFOk returns a tuple with the InvoicePDF field value
// and a boolean to check if the value has been set.
func (o *InvoicesSchemaInner) GetInvoicePDFOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoicePDF, true
}

// SetInvoicePDF sets field value
func (o *InvoicesSchemaInner) SetInvoicePDF(v string) {
	o.InvoicePDF = v
}

func (o InvoicesSchemaInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoicesSchemaInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amountFormatted"] = o.AmountFormatted
	toSerialize["paid"] = o.Paid
	toSerialize["status"] = o.Status
	if !IsNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
	toSerialize["invoiceURL"] = o.InvoiceURL
	toSerialize["invoicePDF"] = o.InvoicePDF
	return toSerialize, nil
}

type NullableInvoicesSchemaInner struct {
	value *InvoicesSchemaInner
	isSet bool
}

func (v NullableInvoicesSchemaInner) Get() *InvoicesSchemaInner {
	return v.value
}

func (v *NullableInvoicesSchemaInner) Set(val *InvoicesSchemaInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoicesSchemaInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoicesSchemaInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoicesSchemaInner(val *InvoicesSchemaInner) *NullableInvoicesSchemaInner {
	return &NullableInvoicesSchemaInner{value: val, isSet: true}
}

func (v NullableInvoicesSchemaInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoicesSchemaInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
