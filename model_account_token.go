/*
 * Merge CRM API
 *
 * The unified API for building rich integrations with multiple CRM platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_crm_client

import (
	"encoding/json"
)

// AccountToken struct for AccountToken
type AccountToken struct {
	AccountToken string `json:"account_token"`
	Integration AccountIntegration `json:"integration"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewAccountToken instantiates a new AccountToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountToken(accountToken string, integration AccountIntegration) *AccountToken {
	this := AccountToken{}
	this.AccountToken = accountToken
	this.Integration = integration
	return &this
}

// NewAccountTokenWithDefaults instantiates a new AccountToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountTokenWithDefaults() *AccountToken {
	this := AccountToken{}
	return &this
}

// GetAccountToken returns the AccountToken field value
func (o *AccountToken) GetAccountToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountToken
}

// GetAccountTokenOk returns a tuple with the AccountToken field value
// and a boolean to check if the value has been set.
func (o *AccountToken) GetAccountTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountToken, true
}

// SetAccountToken sets field value
func (o *AccountToken) SetAccountToken(v string) {
	o.AccountToken = v
}

// GetIntegration returns the Integration field value
func (o *AccountToken) GetIntegration() AccountIntegration {
	if o == nil {
		var ret AccountIntegration
		return ret
	}

	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *AccountToken) GetIntegrationOk() (*AccountIntegration, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value
func (o *AccountToken) SetIntegration(v AccountIntegration) {
	o.Integration = v
}

func (o AccountToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_token"] = o.AccountToken
	}
	if true {
		toSerialize["integration"] = o.Integration
	}
	return json.Marshal(toSerialize)
}

func (v *AccountToken) UnmarshalJSON(src []byte) error {
    type AccountTokenUnmarshalTarget AccountToken

	var intermediateResult AccountTokenUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = AccountToken(intermediateResult)
	return nil
}
type NullableAccountToken struct {
	value *AccountToken
	isSet bool
}

func (v NullableAccountToken) Get() *AccountToken {
	return v.value
}

func (v *NullableAccountToken) Set(val *AccountToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountToken(val *AccountToken) *NullableAccountToken {
	return &NullableAccountToken{value: val, isSet: true}
}

func (v NullableAccountToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


