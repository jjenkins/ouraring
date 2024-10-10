/*
Oura API Documentation

# Overview   The Oura API allows Oura users and partner applications to improve their user experience with Oura data.  This document describes the Oura API Version 2 (V2), which supports access to the latest Oura Ring data types.  For access to other data types—which have not yet migrated to V2—refer to the [Oura API Version 1 (V1)](https://cloud.ouraring.com/docs/) documentation.  # Data Access  Individual Oura users can access their own data through the API by using a [Personal Access Token](https://cloud.ouraring.com/personal-access-tokens).  If you want to retrieve data for multiple users, a registered [API Application](https://cloud.ouraring.com/oauth/applications) is required.  API Applications are limited to **10** users before requiring approval from Oura. There is no limit once an application is approved.  Additionally, Oura users **must provide consent** to share each data type an API Application has access to.  All data access requests through the Oura API require [Authentication](https://cloud.ouraring.com/docs/authentication).  Additionally, we recommend that Oura users keep their mobile app updated to support API access for the latest data types.  # Authentication  The Oura API provides two methods for Authentication: (1) the OAuth2 protocol and (2) Personal Access Tokens. For more information on the OAuth2 flow, see our [Authentication instructions](https://cloud.ouraring.com/docs/authentication).  Access tokens must be included in the request header as follows: ```http GET /v2/usercollection/personal_info HTTP/1.1 Host: api.ouraring.com Authorization: Bearer <token> ```  # Oura HTTP Response Codes  | Response Code                        | Description | | ------------------------------------ | - | | 200 OK                               | Successful Response         | | 400 Query Parameter Validation Error | The request contains query parameters that are invalid or incorrectly formatted. | | 429 Request Rate Limit Exceeded        | The API is rate limited to 5000 requests in a 5 minute period. You will receive a 429 error code if you exceed this limit. [Contact us](mailto:api-support@ouraring.com) if you expect your usage to exceed this limit.| 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oura

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the TagModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagModel{}

// TagModel A TagModel maps to an ASSANote. An ASSANote in ExtAPIV2 is called a Tag A TagModel will be populated by data from an ASSANote The fields in the TagModel map to fields in an ASSANote
type TagModel struct {
	Id string `json:"id"`
	// Day that the note belongs to.
	Day string `json:"day"`
	Text NullableString `json:"text"`
	// Timestamp of the note.
	Timestamp time.Time `json:"timestamp"`
	// Selected tags for the tag.
	Tags []string `json:"tags"`
}

type _TagModel TagModel

// NewTagModel instantiates a new TagModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagModel(id string, day string, text NullableString, timestamp time.Time, tags []string) *TagModel {
	this := TagModel{}
	this.Id = id
	this.Day = day
	this.Text = text
	this.Timestamp = timestamp
	this.Tags = tags
	return &this
}

// NewTagModelWithDefaults instantiates a new TagModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagModelWithDefaults() *TagModel {
	this := TagModel{}
	return &this
}

// GetId returns the Id field value
func (o *TagModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TagModel) SetId(v string) {
	o.Id = v
}

// GetDay returns the Day field value
func (o *TagModel) GetDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetDayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *TagModel) SetDay(v string) {
	o.Day = v
}

// GetText returns the Text field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TagModel) GetText() string {
	if o == nil || o.Text.Get() == nil {
		var ret string
		return ret
	}

	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagModel) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// SetText sets field value
func (o *TagModel) SetText(v string) {
	o.Text.Set(&v)
}

// GetTimestamp returns the Timestamp field value
func (o *TagModel) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *TagModel) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetTags returns the Tags field value
func (o *TagModel) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *TagModel) SetTags(v []string) {
	o.Tags = v
}

func (o TagModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["day"] = o.Day
	toSerialize["text"] = o.Text.Get()
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

func (o *TagModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"day",
		"text",
		"timestamp",
		"tags",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTagModel := _TagModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagModel)

	if err != nil {
		return err
	}

	*o = TagModel(varTagModel)

	return err
}

type NullableTagModel struct {
	value *TagModel
	isSet bool
}

func (v NullableTagModel) Get() *TagModel {
	return v.value
}

func (v *NullableTagModel) Set(val *TagModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTagModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTagModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagModel(val *TagModel) *NullableTagModel {
	return &NullableTagModel{value: val, isSet: true}
}

func (v NullableTagModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


