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

// CRMCustomObjectEndpointRequest struct for CRMCustomObjectEndpointRequest
type CRMCustomObjectEndpointRequest struct {
	Model CustomObjectRequest `json:"model"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewCRMCustomObjectEndpointRequest instantiates a new CRMCustomObjectEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCRMCustomObjectEndpointRequest(model CustomObjectRequest) *CRMCustomObjectEndpointRequest {
	this := CRMCustomObjectEndpointRequest{}
	this.Model = model
	return &this
}

// NewCRMCustomObjectEndpointRequestWithDefaults instantiates a new CRMCustomObjectEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCRMCustomObjectEndpointRequestWithDefaults() *CRMCustomObjectEndpointRequest {
	this := CRMCustomObjectEndpointRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *CRMCustomObjectEndpointRequest) GetModel() CustomObjectRequest {
	if o == nil {
		var ret CustomObjectRequest
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *CRMCustomObjectEndpointRequest) GetModelOk() (*CustomObjectRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *CRMCustomObjectEndpointRequest) SetModel(v CustomObjectRequest) {
	o.Model = v
}

func (o CRMCustomObjectEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

func (v *CRMCustomObjectEndpointRequest) UnmarshalJSON(src []byte) error {
    type CRMCustomObjectEndpointRequestUnmarshalTarget CRMCustomObjectEndpointRequest

	var intermediateResult CRMCustomObjectEndpointRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = CRMCustomObjectEndpointRequest(intermediateResult)
	return nil
}
type NullableCRMCustomObjectEndpointRequest struct {
	value *CRMCustomObjectEndpointRequest
	isSet bool
}

func (v NullableCRMCustomObjectEndpointRequest) Get() *CRMCustomObjectEndpointRequest {
	return v.value
}

func (v *NullableCRMCustomObjectEndpointRequest) Set(val *CRMCustomObjectEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCRMCustomObjectEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCRMCustomObjectEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCRMCustomObjectEndpointRequest(val *CRMCustomObjectEndpointRequest) *NullableCRMCustomObjectEndpointRequest {
	return &NullableCRMCustomObjectEndpointRequest{value: val, isSet: true}
}

func (v NullableCRMCustomObjectEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCRMCustomObjectEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}

