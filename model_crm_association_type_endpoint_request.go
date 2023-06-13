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

// CRMAssociationTypeEndpointRequest struct for CRMAssociationTypeEndpointRequest
type CRMAssociationTypeEndpointRequest struct {
	Model AssociationTypeRequestRequest `json:"model"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewCRMAssociationTypeEndpointRequest instantiates a new CRMAssociationTypeEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCRMAssociationTypeEndpointRequest(model AssociationTypeRequestRequest) *CRMAssociationTypeEndpointRequest {
	this := CRMAssociationTypeEndpointRequest{}
	this.Model = model
	return &this
}

// NewCRMAssociationTypeEndpointRequestWithDefaults instantiates a new CRMAssociationTypeEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCRMAssociationTypeEndpointRequestWithDefaults() *CRMAssociationTypeEndpointRequest {
	this := CRMAssociationTypeEndpointRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *CRMAssociationTypeEndpointRequest) GetModel() AssociationTypeRequestRequest {
	if o == nil {
		var ret AssociationTypeRequestRequest
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *CRMAssociationTypeEndpointRequest) GetModelOk() (*AssociationTypeRequestRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *CRMAssociationTypeEndpointRequest) SetModel(v AssociationTypeRequestRequest) {
	o.Model = v
}

func (o CRMAssociationTypeEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

func (v *CRMAssociationTypeEndpointRequest) UnmarshalJSON(src []byte) error {
    type CRMAssociationTypeEndpointRequestUnmarshalTarget CRMAssociationTypeEndpointRequest

	var intermediateResult CRMAssociationTypeEndpointRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = CRMAssociationTypeEndpointRequest(intermediateResult)
	return nil
}
type NullableCRMAssociationTypeEndpointRequest struct {
	value *CRMAssociationTypeEndpointRequest
	isSet bool
}

func (v NullableCRMAssociationTypeEndpointRequest) Get() *CRMAssociationTypeEndpointRequest {
	return v.value
}

func (v *NullableCRMAssociationTypeEndpointRequest) Set(val *CRMAssociationTypeEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCRMAssociationTypeEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCRMAssociationTypeEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCRMAssociationTypeEndpointRequest(val *CRMAssociationTypeEndpointRequest) *NullableCRMAssociationTypeEndpointRequest {
	return &NullableCRMAssociationTypeEndpointRequest{value: val, isSet: true}
}

func (v NullableCRMAssociationTypeEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCRMAssociationTypeEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}

