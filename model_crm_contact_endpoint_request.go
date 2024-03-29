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

// CRMContactEndpointRequest struct for CRMContactEndpointRequest
type CRMContactEndpointRequest struct {
	Model ContactRequest `json:"model"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewCRMContactEndpointRequest instantiates a new CRMContactEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCRMContactEndpointRequest(model ContactRequest) *CRMContactEndpointRequest {
	this := CRMContactEndpointRequest{}
	this.Model = model
	return &this
}

// NewCRMContactEndpointRequestWithDefaults instantiates a new CRMContactEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCRMContactEndpointRequestWithDefaults() *CRMContactEndpointRequest {
	this := CRMContactEndpointRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *CRMContactEndpointRequest) GetModel() ContactRequest {
	if o == nil {
		var ret ContactRequest
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *CRMContactEndpointRequest) GetModelOk() (*ContactRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *CRMContactEndpointRequest) SetModel(v ContactRequest) {
	o.Model = v
}

func (o CRMContactEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

func (v *CRMContactEndpointRequest) UnmarshalJSON(src []byte) error {
    type CRMContactEndpointRequestUnmarshalTarget CRMContactEndpointRequest

	var intermediateResult CRMContactEndpointRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = CRMContactEndpointRequest(intermediateResult)
	return nil
}
type NullableCRMContactEndpointRequest struct {
	value *CRMContactEndpointRequest
	isSet bool
}

func (v NullableCRMContactEndpointRequest) Get() *CRMContactEndpointRequest {
	return v.value
}

func (v *NullableCRMContactEndpointRequest) Set(val *CRMContactEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCRMContactEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCRMContactEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCRMContactEndpointRequest(val *CRMContactEndpointRequest) *NullableCRMContactEndpointRequest {
	return &NullableCRMContactEndpointRequest{value: val, isSet: true}
}

func (v NullableCRMContactEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCRMContactEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


