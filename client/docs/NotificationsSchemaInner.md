# NotificationsSchemaInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The id of this notification | 
**Message** | **string** | The actual notification message | 
**Link** | **string** | The link to change request or feature toggle the notification refers to | 
**NotificationType** | **string** | The type of the notification used e.g. for the graphical hints | 
**CreatedBy** | [**NotificationsSchemaInnerCreatedBy**](NotificationsSchemaInnerCreatedBy.md) |  | 
**CreatedAt** | **time.Time** | The date and time when the notification was created | 
**ReadAt** | **NullableTime** | The date and time when the notification was read or marked as read, otherwise &#x60;null&#x60; | 

## Methods

### NewNotificationsSchemaInner

`func NewNotificationsSchemaInner(id float32, message string, link string, notificationType string, createdBy NotificationsSchemaInnerCreatedBy, createdAt time.Time, readAt NullableTime, ) *NotificationsSchemaInner`

NewNotificationsSchemaInner instantiates a new NotificationsSchemaInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSchemaInnerWithDefaults

`func NewNotificationsSchemaInnerWithDefaults() *NotificationsSchemaInner`

NewNotificationsSchemaInnerWithDefaults instantiates a new NotificationsSchemaInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationsSchemaInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsSchemaInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsSchemaInner) SetId(v float32)`

SetId sets Id field to given value.


### GetMessage

`func (o *NotificationsSchemaInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationsSchemaInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationsSchemaInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetLink

`func (o *NotificationsSchemaInner) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *NotificationsSchemaInner) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *NotificationsSchemaInner) SetLink(v string)`

SetLink sets Link field to given value.


### GetNotificationType

`func (o *NotificationsSchemaInner) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotificationsSchemaInner) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotificationsSchemaInner) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetCreatedBy

`func (o *NotificationsSchemaInner) GetCreatedBy() NotificationsSchemaInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *NotificationsSchemaInner) GetCreatedByOk() (*NotificationsSchemaInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *NotificationsSchemaInner) SetCreatedBy(v NotificationsSchemaInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *NotificationsSchemaInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationsSchemaInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationsSchemaInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetReadAt

`func (o *NotificationsSchemaInner) GetReadAt() time.Time`

GetReadAt returns the ReadAt field if non-nil, zero value otherwise.

### GetReadAtOk

`func (o *NotificationsSchemaInner) GetReadAtOk() (*time.Time, bool)`

GetReadAtOk returns a tuple with the ReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAt

`func (o *NotificationsSchemaInner) SetReadAt(v time.Time)`

SetReadAt sets ReadAt field to given value.


### SetReadAtNil

`func (o *NotificationsSchemaInner) SetReadAtNil(b bool)`

 SetReadAtNil sets the value for ReadAt to be an explicit nil

### UnsetReadAt
`func (o *NotificationsSchemaInner) UnsetReadAt()`

UnsetReadAt ensures that no value is present for ReadAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


