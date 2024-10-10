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

// checks if the DailyReadinessModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DailyReadinessModel{}

// DailyReadinessModel struct for DailyReadinessModel
type DailyReadinessModel struct {
	Id string `json:"id"`
	// Contributors of the daily readiness score.
	Contributors ReadinessContributors `json:"contributors"`
	// Day that the daily readiness belongs to.
	Day string `json:"day"`
	Score NullableInt32 `json:"score"`
	TemperatureDeviation NullableFloat32 `json:"temperature_deviation"`
	TemperatureTrendDeviation NullableFloat32 `json:"temperature_trend_deviation"`
	// Timestamp of the daily readiness.
	Timestamp string `json:"timestamp"`
}

type _DailyReadinessModel DailyReadinessModel

// NewDailyReadinessModel instantiates a new DailyReadinessModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyReadinessModel(id string, contributors ReadinessContributors, day string, score NullableInt32, temperatureDeviation NullableFloat32, temperatureTrendDeviation NullableFloat32, timestamp string) *DailyReadinessModel {
	this := DailyReadinessModel{}
	this.Id = id
	this.Contributors = contributors
	this.Day = day
	this.Score = score
	this.TemperatureDeviation = temperatureDeviation
	this.TemperatureTrendDeviation = temperatureTrendDeviation
	this.Timestamp = timestamp
	return &this
}

// NewDailyReadinessModelWithDefaults instantiates a new DailyReadinessModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyReadinessModelWithDefaults() *DailyReadinessModel {
	this := DailyReadinessModel{}
	return &this
}

// GetId returns the Id field value
func (o *DailyReadinessModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DailyReadinessModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DailyReadinessModel) SetId(v string) {
	o.Id = v
}

// GetContributors returns the Contributors field value
func (o *DailyReadinessModel) GetContributors() ReadinessContributors {
	if o == nil {
		var ret ReadinessContributors
		return ret
	}

	return o.Contributors
}

// GetContributorsOk returns a tuple with the Contributors field value
// and a boolean to check if the value has been set.
func (o *DailyReadinessModel) GetContributorsOk() (*ReadinessContributors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contributors, true
}

// SetContributors sets field value
func (o *DailyReadinessModel) SetContributors(v ReadinessContributors) {
	o.Contributors = v
}

// GetDay returns the Day field value
func (o *DailyReadinessModel) GetDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *DailyReadinessModel) GetDayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *DailyReadinessModel) SetDay(v string) {
	o.Day = v
}

// GetScore returns the Score field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DailyReadinessModel) GetScore() int32 {
	if o == nil || o.Score.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Score.Get()
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DailyReadinessModel) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Score.Get(), o.Score.IsSet()
}

// SetScore sets field value
func (o *DailyReadinessModel) SetScore(v int32) {
	o.Score.Set(&v)
}

// GetTemperatureDeviation returns the TemperatureDeviation field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *DailyReadinessModel) GetTemperatureDeviation() float32 {
	if o == nil || o.TemperatureDeviation.Get() == nil {
		var ret float32
		return ret
	}

	return *o.TemperatureDeviation.Get()
}

// GetTemperatureDeviationOk returns a tuple with the TemperatureDeviation field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DailyReadinessModel) GetTemperatureDeviationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemperatureDeviation.Get(), o.TemperatureDeviation.IsSet()
}

// SetTemperatureDeviation sets field value
func (o *DailyReadinessModel) SetTemperatureDeviation(v float32) {
	o.TemperatureDeviation.Set(&v)
}

// GetTemperatureTrendDeviation returns the TemperatureTrendDeviation field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *DailyReadinessModel) GetTemperatureTrendDeviation() float32 {
	if o == nil || o.TemperatureTrendDeviation.Get() == nil {
		var ret float32
		return ret
	}

	return *o.TemperatureTrendDeviation.Get()
}

// GetTemperatureTrendDeviationOk returns a tuple with the TemperatureTrendDeviation field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DailyReadinessModel) GetTemperatureTrendDeviationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemperatureTrendDeviation.Get(), o.TemperatureTrendDeviation.IsSet()
}

// SetTemperatureTrendDeviation sets field value
func (o *DailyReadinessModel) SetTemperatureTrendDeviation(v float32) {
	o.TemperatureTrendDeviation.Set(&v)
}

// GetTimestamp returns the Timestamp field value
func (o *DailyReadinessModel) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *DailyReadinessModel) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *DailyReadinessModel) SetTimestamp(v string) {
	o.Timestamp = v
}

func (o DailyReadinessModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailyReadinessModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["contributors"] = o.Contributors
	toSerialize["day"] = o.Day
	toSerialize["score"] = o.Score.Get()
	toSerialize["temperature_deviation"] = o.TemperatureDeviation.Get()
	toSerialize["temperature_trend_deviation"] = o.TemperatureTrendDeviation.Get()
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

func (o *DailyReadinessModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"contributors",
		"day",
		"score",
		"temperature_deviation",
		"temperature_trend_deviation",
		"timestamp",
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

	varDailyReadinessModel := _DailyReadinessModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDailyReadinessModel)

	if err != nil {
		return err
	}

	*o = DailyReadinessModel(varDailyReadinessModel)

	return err
}

type NullableDailyReadinessModel struct {
	value *DailyReadinessModel
	isSet bool
}

func (v NullableDailyReadinessModel) Get() *DailyReadinessModel {
	return v.value
}

func (v *NullableDailyReadinessModel) Set(val *DailyReadinessModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyReadinessModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyReadinessModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyReadinessModel(val *DailyReadinessModel) *NullableDailyReadinessModel {
	return &NullableDailyReadinessModel{value: val, isSet: true}
}

func (v NullableDailyReadinessModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyReadinessModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


