# Login401Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the error instance | [optional] 
**Name** | Pointer to **string** | The name of the error kind | [optional] 
**Message** | Pointer to **string** | A description of what went wrong. | [optional] 

## Methods

### NewLogin401Response

`func NewLogin401Response() *Login401Response`

NewLogin401Response instantiates a new Login401Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogin401ResponseWithDefaults

`func NewLogin401ResponseWithDefaults() *Login401Response`

NewLogin401ResponseWithDefaults instantiates a new Login401Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Login401Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Login401Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Login401Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Login401Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Login401Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Login401Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Login401Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Login401Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMessage

`func (o *Login401Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Login401Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Login401Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Login401Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


