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

// OpportunityEndpointRequest struct for OpportunityEndpointRequest
type OpportunityEndpointRequest struct {
	Model OpportunityRequest `json:"model"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewOpportunityEndpointRequest instantiates a new OpportunityEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpportunityEndpointRequest(model OpportunityRequest) *OpportunityEndpointRequest {
	this := OpportunityEndpointRequest{}
	this.Model = model
	return &this
}

// NewOpportunityEndpointRequestWithDefaults instantiates a new OpportunityEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpportunityEndpointRequestWithDefaults() *OpportunityEndpointRequest {
	this := OpportunityEndpointRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *OpportunityEndpointRequest) GetModel() OpportunityRequest {
	if o == nil {
		var ret OpportunityRequest
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *OpportunityEndpointRequest) GetModelOk() (*OpportunityRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *OpportunityEndpointRequest) SetModel(v OpportunityRequest) {
	o.Model = v
}

func (o OpportunityEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

func (v *OpportunityEndpointRequest) UnmarshalJSON(src []byte) error {
    type OpportunityEndpointRequestUnmarshalTarget OpportunityEndpointRequest

	var intermediateResult OpportunityEndpointRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = OpportunityEndpointRequest(intermediateResult)
	return nil
}
type NullableOpportunityEndpointRequest struct {
	value *OpportunityEndpointRequest
	isSet bool
}

func (v NullableOpportunityEndpointRequest) Get() *OpportunityEndpointRequest {
	return v.value
}

func (v *NullableOpportunityEndpointRequest) Set(val *OpportunityEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOpportunityEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOpportunityEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpportunityEndpointRequest(val *OpportunityEndpointRequest) *NullableOpportunityEndpointRequest {
	return &NullableOpportunityEndpointRequest{value: val, isSet: true}
}

func (v NullableOpportunityEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpportunityEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


