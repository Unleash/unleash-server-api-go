# InvoicesSchemaInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountFormatted** | **string** |  | 
**Paid** | **bool** |  | 
**Status** | **string** |  | 
**DueDate** | Pointer to **string** |  | [optional] 
**InvoiceURL** | **string** |  | 
**InvoicePDF** | **string** |  | 

## Methods

### NewInvoicesSchemaInner

`func NewInvoicesSchemaInner(amountFormatted string, paid bool, status string, invoiceURL string, invoicePDF string, ) *InvoicesSchemaInner`

NewInvoicesSchemaInner instantiates a new InvoicesSchemaInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicesSchemaInnerWithDefaults

`func NewInvoicesSchemaInnerWithDefaults() *InvoicesSchemaInner`

NewInvoicesSchemaInnerWithDefaults instantiates a new InvoicesSchemaInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountFormatted

`func (o *InvoicesSchemaInner) GetAmountFormatted() string`

GetAmountFormatted returns the AmountFormatted field if non-nil, zero value otherwise.

### GetAmountFormattedOk

`func (o *InvoicesSchemaInner) GetAmountFormattedOk() (*string, bool)`

GetAmountFormattedOk returns a tuple with the AmountFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFormatted

`func (o *InvoicesSchemaInner) SetAmountFormatted(v string)`

SetAmountFormatted sets AmountFormatted field to given value.


### GetPaid

`func (o *InvoicesSchemaInner) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *InvoicesSchemaInner) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *InvoicesSchemaInner) SetPaid(v bool)`

SetPaid sets Paid field to given value.


### GetStatus

`func (o *InvoicesSchemaInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoicesSchemaInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoicesSchemaInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDueDate

`func (o *InvoicesSchemaInner) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoicesSchemaInner) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoicesSchemaInner) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoicesSchemaInner) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetInvoiceURL

`func (o *InvoicesSchemaInner) GetInvoiceURL() string`

GetInvoiceURL returns the InvoiceURL field if non-nil, zero value otherwise.

### GetInvoiceURLOk

`func (o *InvoicesSchemaInner) GetInvoiceURLOk() (*string, bool)`

GetInvoiceURLOk returns a tuple with the InvoiceURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceURL

`func (o *InvoicesSchemaInner) SetInvoiceURL(v string)`

SetInvoiceURL sets InvoiceURL field to given value.


### GetInvoicePDF

`func (o *InvoicesSchemaInner) GetInvoicePDF() string`

GetInvoicePDF returns the InvoicePDF field if non-nil, zero value otherwise.

### GetInvoicePDFOk

`func (o *InvoicesSchemaInner) GetInvoicePDFOk() (*string, bool)`

GetInvoicePDFOk returns a tuple with the InvoicePDF field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePDF

`func (o *InvoicesSchemaInner) SetInvoicePDF(v string)`

SetInvoicePDF sets InvoicePDF field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


