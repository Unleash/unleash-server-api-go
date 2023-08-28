/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.4.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the NotificationsSchemaInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsSchemaInner{}

// NotificationsSchemaInner struct for NotificationsSchemaInner
type NotificationsSchemaInner struct {
	// The id of this notification
	Id float32 `json:"id"`
	// The actual notification message
	Message string `json:"message"`
	// The link to change request or feature toggle the notification refers to
	Link string `json:"link"`
	// The type of the notification used e.g. for the graphical hints
	NotificationType string                            `json:"notificationType"`
	CreatedBy        NotificationsSchemaInnerCreatedBy `json:"createdBy"`
	// The date and time when the notification was created
	CreatedAt time.Time `json:"createdAt"`
	// The date and time when the notification was read or marked as read, otherwise `null`
	ReadAt NullableTime `json:"readAt"`
}

// NewNotificationsSchemaInner instantiates a new NotificationsSchemaInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSchemaInner(id float32, message string, link string, notificationType string, createdBy NotificationsSchemaInnerCreatedBy, createdAt time.Time, readAt NullableTime) *NotificationsSchemaInner {
	this := NotificationsSchemaInner{}
	this.Id = id
	this.Message = message
	this.Link = link
	this.NotificationType = notificationType
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	this.ReadAt = readAt
	return &this
}

// NewNotificationsSchemaInnerWithDefaults instantiates a new NotificationsSchemaInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSchemaInnerWithDefaults() *NotificationsSchemaInner {
	this := NotificationsSchemaInner{}
	return &this
}

// GetId returns the Id field value
func (o *NotificationsSchemaInner) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NotificationsSchemaInner) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NotificationsSchemaInner) SetId(v float32) {
	o.Id = v
}

// GetMessage returns the Message field value
func (o *NotificationsSchemaInner) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *NotificationsSchemaInner) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *NotificationsSchemaInner) SetMessage(v string) {
	o.Message = v
}

// GetLink returns the Link field value
func (o *NotificationsSchemaInner) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *NotificationsSchemaInner) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *NotificationsSchemaInner) SetLink(v string) {
	o.Link = v
}

// GetNotificationType returns the NotificationType field value
func (o *NotificationsSchemaInner) GetNotificationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *NotificationsSchemaInner) GetNotificationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *NotificationsSchemaInner) SetNotificationType(v string) {
	o.NotificationType = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *NotificationsSchemaInner) GetCreatedBy() NotificationsSchemaInnerCreatedBy {
	if o == nil {
		var ret NotificationsSchemaInnerCreatedBy
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *NotificationsSchemaInner) GetCreatedByOk() (*NotificationsSchemaInnerCreatedBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *NotificationsSchemaInner) SetCreatedBy(v NotificationsSchemaInnerCreatedBy) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NotificationsSchemaInner) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NotificationsSchemaInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NotificationsSchemaInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetReadAt returns the ReadAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *NotificationsSchemaInner) GetReadAt() time.Time {
	if o == nil || o.ReadAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ReadAt.Get()
}

// GetReadAtOk returns a tuple with the ReadAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationsSchemaInner) GetReadAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReadAt.Get(), o.ReadAt.IsSet()
}

// SetReadAt sets field value
func (o *NotificationsSchemaInner) SetReadAt(v time.Time) {
	o.ReadAt.Set(&v)
}

func (o NotificationsSchemaInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsSchemaInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["message"] = o.Message
	toSerialize["link"] = o.Link
	toSerialize["notificationType"] = o.NotificationType
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["readAt"] = o.ReadAt.Get()
	return toSerialize, nil
}

type NullableNotificationsSchemaInner struct {
	value *NotificationsSchemaInner
	isSet bool
}

func (v NullableNotificationsSchemaInner) Get() *NotificationsSchemaInner {
	return v.value
}

func (v *NullableNotificationsSchemaInner) Set(val *NotificationsSchemaInner) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSchemaInner) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSchemaInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSchemaInner(val *NotificationsSchemaInner) *NullableNotificationsSchemaInner {
	return &NullableNotificationsSchemaInner{value: val, isSet: true}
}

func (v NullableNotificationsSchemaInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSchemaInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
