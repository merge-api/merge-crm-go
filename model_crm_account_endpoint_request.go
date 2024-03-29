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

// CRMAccountEndpointRequest struct for CRMAccountEndpointRequest
type CRMAccountEndpointRequest struct {
	Model AccountRequest `json:"model"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewCRMAccountEndpointRequest instantiates a new CRMAccountEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCRMAccountEndpointRequest(model AccountRequest) *CRMAccountEndpointRequest {
	this := CRMAccountEndpointRequest{}
	this.Model = model
	return &this
}

// NewCRMAccountEndpointRequestWithDefaults instantiates a new CRMAccountEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCRMAccountEndpointRequestWithDefaults() *CRMAccountEndpointRequest {
	this := CRMAccountEndpointRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *CRMAccountEndpointRequest) GetModel() AccountRequest {
	if o == nil {
		var ret AccountRequest
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *CRMAccountEndpointRequest) GetModelOk() (*AccountRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *CRMAccountEndpointRequest) SetModel(v AccountRequest) {
	o.Model = v
}

func (o CRMAccountEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

func (v *CRMAccountEndpointRequest) UnmarshalJSON(src []byte) error {
    type CRMAccountEndpointRequestUnmarshalTarget CRMAccountEndpointRequest

	var intermediateResult CRMAccountEndpointRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = CRMAccountEndpointRequest(intermediateResult)
	return nil
}
type NullableCRMAccountEndpointRequest struct {
	value *CRMAccountEndpointRequest
	isSet bool
}

func (v NullableCRMAccountEndpointRequest) Get() *CRMAccountEndpointRequest {
	return v.value
}

func (v *NullableCRMAccountEndpointRequest) Set(val *CRMAccountEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCRMAccountEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCRMAccountEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCRMAccountEndpointRequest(val *CRMAccountEndpointRequest) *NullableCRMAccountEndpointRequest {
	return &NullableCRMAccountEndpointRequest{value: val, isSet: true}
}

func (v NullableCRMAccountEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCRMAccountEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


