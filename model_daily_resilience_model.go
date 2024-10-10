/*
Oura API Documentation

# Overview   The Oura API allows Oura users and partner applications to improve their user experience with Oura data.  This document describes the Oura API Version 2 (V2), which supports access to the latest Oura Ring data types.  For access to other data types—which have not yet migrated to V2—refer to the [Oura API Version 1 (V1)](https://cloud.ouraring.com/docs/) documentation.  # Data Access  Individual Oura users can access their own data through the API by using a [Personal Access Token](https://cloud.ouraring.com/personal-access-tokens).  If you want to retrieve data for multiple users, a registered [API Application](https://cloud.ouraring.com/oauth/applications) is required.  API Applications are limited to **10** users before requiring approval from Oura. There is no limit once an application is approved.  Additionally, Oura users **must provide consent** to share each data type an API Application has access to.  All data access requests through the Oura API require [Authentication](https://cloud.ouraring.com/docs/authentication).  Additionally, we recommend that Oura users keep their mobile app updated to support API access for the latest data types.  # Authentication  The Oura API provides two methods for Authentication: (1) the OAuth2 protocol and (2) Personal Access Tokens. For more information on the OAuth2 flow, see our [Authentication instructions](https://cloud.ouraring.com/docs/authentication).  Access tokens must be included in the request header as follows: ```http GET /v2/usercollection/personal_info HTTP/1.1 Host: api.ouraring.com Authorization: Bearer <token> ```  # Oura HTTP Response Codes  | Response Code                        | Description | | ------------------------------------ | - | | 200 OK                               | Successful Response         | | 400 Query Parameter Validation Error | The request contains query parameters that are invalid or incorrectly formatted. | | 429 Request Rate Limit Exceeded        | The API is rate limited to 5000 requests in a 5 minute period. You will receive a 429 error code if you exceed this limit. [Contact us](mailto:api-support@ouraring.com) if you expect your usage to exceed this limit.| 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oura

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DailyResilienceModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DailyResilienceModel{}

// DailyResilienceModel struct for DailyResilienceModel
type DailyResilienceModel struct {
	Id string `json:"id"`
	// Day when the resilience record was recorded.
	Day string `json:"day"`
	// Contributors to the resilience score.
	Contributors ResilienceContributors `json:"contributors"`
	// Resilience level.
	Level LongTermResilienceLevel `json:"level"`
}

type _DailyResilienceModel DailyResilienceModel

// NewDailyResilienceModel instantiates a new DailyResilienceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyResilienceModel(id string, day string, contributors ResilienceContributors, level LongTermResilienceLevel) *DailyResilienceModel {
	this := DailyResilienceModel{}
	this.Id = id
	this.Day = day
	this.Contributors = contributors
	this.Level = level
	return &this
}

// NewDailyResilienceModelWithDefaults instantiates a new DailyResilienceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyResilienceModelWithDefaults() *DailyResilienceModel {
	this := DailyResilienceModel{}
	return &this
}

// GetId returns the Id field value
func (o *DailyResilienceModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DailyResilienceModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DailyResilienceModel) SetId(v string) {
	o.Id = v
}

// GetDay returns the Day field value
func (o *DailyResilienceModel) GetDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *DailyResilienceModel) GetDayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *DailyResilienceModel) SetDay(v string) {
	o.Day = v
}

// GetContributors returns the Contributors field value
func (o *DailyResilienceModel) GetContributors() ResilienceContributors {
	if o == nil {
		var ret ResilienceContributors
		return ret
	}

	return o.Contributors
}

// GetContributorsOk returns a tuple with the Contributors field value
// and a boolean to check if the value has been set.
func (o *DailyResilienceModel) GetContributorsOk() (*ResilienceContributors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contributors, true
}

// SetContributors sets field value
func (o *DailyResilienceModel) SetContributors(v ResilienceContributors) {
	o.Contributors = v
}

// GetLevel returns the Level field value
func (o *DailyResilienceModel) GetLevel() LongTermResilienceLevel {
	if o == nil {
		var ret LongTermResilienceLevel
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *DailyResilienceModel) GetLevelOk() (*LongTermResilienceLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *DailyResilienceModel) SetLevel(v LongTermResilienceLevel) {
	o.Level = v
}

func (o DailyResilienceModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailyResilienceModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["day"] = o.Day
	toSerialize["contributors"] = o.Contributors
	toSerialize["level"] = o.Level
	return toSerialize, nil
}

func (o *DailyResilienceModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"day",
		"contributors",
		"level",
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

	varDailyResilienceModel := _DailyResilienceModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDailyResilienceModel)

	if err != nil {
		return err
	}

	*o = DailyResilienceModel(varDailyResilienceModel)

	return err
}

type NullableDailyResilienceModel struct {
	value *DailyResilienceModel
	isSet bool
}

func (v NullableDailyResilienceModel) Get() *DailyResilienceModel {
	return v.value
}

func (v *NullableDailyResilienceModel) Set(val *DailyResilienceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyResilienceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyResilienceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyResilienceModel(val *DailyResilienceModel) *NullableDailyResilienceModel {
	return &NullableDailyResilienceModel{value: val, isSet: true}
}

func (v NullableDailyResilienceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyResilienceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


