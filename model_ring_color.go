/*
Oura API Documentation

# Overview   The Oura API allows Oura users and partner applications to improve their user experience with Oura data.  This document describes the Oura API Version 2 (V2), which supports access to the latest Oura Ring data types.  For access to other data types—which have not yet migrated to V2—refer to the [Oura API Version 1 (V1)](https://cloud.ouraring.com/docs/) documentation.  # Data Access  Individual Oura users can access their own data through the API by using a [Personal Access Token](https://cloud.ouraring.com/personal-access-tokens).  If you want to retrieve data for multiple users, a registered [API Application](https://cloud.ouraring.com/oauth/applications) is required.  API Applications are limited to **10** users before requiring approval from Oura. There is no limit once an application is approved.  Additionally, Oura users **must provide consent** to share each data type an API Application has access to.  All data access requests through the Oura API require [Authentication](https://cloud.ouraring.com/docs/authentication).  Additionally, we recommend that Oura users keep their mobile app updated to support API access for the latest data types.  # Authentication  The Oura API provides two methods for Authentication: (1) the OAuth2 protocol and (2) Personal Access Tokens. For more information on the OAuth2 flow, see our [Authentication instructions](https://cloud.ouraring.com/docs/authentication).  Access tokens must be included in the request header as follows: ```http GET /v2/usercollection/personal_info HTTP/1.1 Host: api.ouraring.com Authorization: Bearer <token> ```  # Oura HTTP Response Codes  | Response Code                        | Description | | ------------------------------------ | - | | 200 OK                               | Successful Response         | | 400 Query Parameter Validation Error | The request contains query parameters that are invalid or incorrectly formatted. | | 429 Request Rate Limit Exceeded        | The API is rate limited to 5000 requests in a 5 minute period. You will receive a 429 error code if you exceed this limit. [Contact us](mailto:api-support@ouraring.com) if you expect your usage to exceed this limit.| 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oura

import (
	"encoding/json"
	"fmt"
)

// RingColor the model 'RingColor'
type RingColor string

// List of RingColor
const (
	BRUSHED_SILVER RingColor = "brushed_silver"
	GLOSSY_BLACK RingColor = "glossy_black"
	GLOSSY_GOLD RingColor = "glossy_gold"
	GLOSSY_WHITE RingColor = "glossy_white"
	GUCCI RingColor = "gucci"
	MATT_GOLD RingColor = "matt_gold"
	ROSE RingColor = "rose"
	SILVER RingColor = "silver"
	STEALTH_BLACK RingColor = "stealth_black"
	TITANIUM RingColor = "titanium"
	TITANIUM_AND_GOLD RingColor = "titanium_and_gold"
)

// All allowed values of RingColor enum
var AllowedRingColorEnumValues = []RingColor{
	"brushed_silver",
	"glossy_black",
	"glossy_gold",
	"glossy_white",
	"gucci",
	"matt_gold",
	"rose",
	"silver",
	"stealth_black",
	"titanium",
	"titanium_and_gold",
}

func (v *RingColor) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RingColor(value)
	for _, existing := range AllowedRingColorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RingColor", value)
}

// NewRingColorFromValue returns a pointer to a valid RingColor
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRingColorFromValue(v string) (*RingColor, error) {
	ev := RingColor(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RingColor: valid values are %v", v, AllowedRingColorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RingColor) IsValid() bool {
	for _, existing := range AllowedRingColorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RingColor value
func (v RingColor) Ptr() *RingColor {
	return &v
}

type NullableRingColor struct {
	value *RingColor
	isSet bool
}

func (v NullableRingColor) Get() *RingColor {
	return v.value
}

func (v *NullableRingColor) Set(val *RingColor) {
	v.value = val
	v.isSet = true
}

func (v NullableRingColor) IsSet() bool {
	return v.isSet
}

func (v *NullableRingColor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRingColor(val *RingColor) *NullableRingColor {
	return &NullableRingColor{value: val, isSet: true}
}

func (v NullableRingColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRingColor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

