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

// checks if the MultiDocumentResponseRingConfigurationModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiDocumentResponseRingConfigurationModel{}

// MultiDocumentResponseRingConfigurationModel struct for MultiDocumentResponseRingConfigurationModel
type MultiDocumentResponseRingConfigurationModel struct {
	Data []RingConfigurationModel `json:"data"`
	NextToken NullableString `json:"next_token"`
}

type _MultiDocumentResponseRingConfigurationModel MultiDocumentResponseRingConfigurationModel

// NewMultiDocumentResponseRingConfigurationModel instantiates a new MultiDocumentResponseRingConfigurationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiDocumentResponseRingConfigurationModel(data []RingConfigurationModel, nextToken NullableString) *MultiDocumentResponseRingConfigurationModel {
	this := MultiDocumentResponseRingConfigurationModel{}
	this.Data = data
	this.NextToken = nextToken
	return &this
}

// NewMultiDocumentResponseRingConfigurationModelWithDefaults instantiates a new MultiDocumentResponseRingConfigurationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiDocumentResponseRingConfigurationModelWithDefaults() *MultiDocumentResponseRingConfigurationModel {
	this := MultiDocumentResponseRingConfigurationModel{}
	return &this
}

// GetData returns the Data field value
func (o *MultiDocumentResponseRingConfigurationModel) GetData() []RingConfigurationModel {
	if o == nil {
		var ret []RingConfigurationModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MultiDocumentResponseRingConfigurationModel) GetDataOk() ([]RingConfigurationModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *MultiDocumentResponseRingConfigurationModel) SetData(v []RingConfigurationModel) {
	o.Data = v
}

// GetNextToken returns the NextToken field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MultiDocumentResponseRingConfigurationModel) GetNextToken() string {
	if o == nil || o.NextToken.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextToken.Get()
}

// GetNextTokenOk returns a tuple with the NextToken field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MultiDocumentResponseRingConfigurationModel) GetNextTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextToken.Get(), o.NextToken.IsSet()
}

// SetNextToken sets field value
func (o *MultiDocumentResponseRingConfigurationModel) SetNextToken(v string) {
	o.NextToken.Set(&v)
}

func (o MultiDocumentResponseRingConfigurationModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiDocumentResponseRingConfigurationModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["next_token"] = o.NextToken.Get()
	return toSerialize, nil
}

func (o *MultiDocumentResponseRingConfigurationModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"next_token",
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

	varMultiDocumentResponseRingConfigurationModel := _MultiDocumentResponseRingConfigurationModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMultiDocumentResponseRingConfigurationModel)

	if err != nil {
		return err
	}

	*o = MultiDocumentResponseRingConfigurationModel(varMultiDocumentResponseRingConfigurationModel)

	return err
}

type NullableMultiDocumentResponseRingConfigurationModel struct {
	value *MultiDocumentResponseRingConfigurationModel
	isSet bool
}

func (v NullableMultiDocumentResponseRingConfigurationModel) Get() *MultiDocumentResponseRingConfigurationModel {
	return v.value
}

func (v *NullableMultiDocumentResponseRingConfigurationModel) Set(val *MultiDocumentResponseRingConfigurationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiDocumentResponseRingConfigurationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiDocumentResponseRingConfigurationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiDocumentResponseRingConfigurationModel(val *MultiDocumentResponseRingConfigurationModel) *NullableMultiDocumentResponseRingConfigurationModel {
	return &NullableMultiDocumentResponseRingConfigurationModel{value: val, isSet: true}
}

func (v NullableMultiDocumentResponseRingConfigurationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiDocumentResponseRingConfigurationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


