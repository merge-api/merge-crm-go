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

// CommonModelScopesBodyRequest struct for CommonModelScopesBodyRequest
type CommonModelScopesBodyRequest struct {
	ModelId string `json:"model_id"`
	EnabledActions []EnabledActionsEnum `json:"enabled_actions"`
	DisabledFields []string `json:"disabled_fields"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewCommonModelScopesBodyRequest instantiates a new CommonModelScopesBodyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonModelScopesBodyRequest(modelId string, enabledActions []EnabledActionsEnum, disabledFields []string) *CommonModelScopesBodyRequest {
	this := CommonModelScopesBodyRequest{}
	this.ModelId = modelId
	this.EnabledActions = enabledActions
	this.DisabledFields = disabledFields
	return &this
}

// NewCommonModelScopesBodyRequestWithDefaults instantiates a new CommonModelScopesBodyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonModelScopesBodyRequestWithDefaults() *CommonModelScopesBodyRequest {
	this := CommonModelScopesBodyRequest{}
	return &this
}

// GetModelId returns the ModelId field value
func (o *CommonModelScopesBodyRequest) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *CommonModelScopesBodyRequest) GetModelIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *CommonModelScopesBodyRequest) SetModelId(v string) {
	o.ModelId = v
}

// GetEnabledActions returns the EnabledActions field value
func (o *CommonModelScopesBodyRequest) GetEnabledActions() []EnabledActionsEnum {
	if o == nil {
		var ret []EnabledActionsEnum
		return ret
	}

	return o.EnabledActions
}

// GetEnabledActionsOk returns a tuple with the EnabledActions field value
// and a boolean to check if the value has been set.
func (o *CommonModelScopesBodyRequest) GetEnabledActionsOk() (*[]EnabledActionsEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EnabledActions, true
}

// SetEnabledActions sets field value
func (o *CommonModelScopesBodyRequest) SetEnabledActions(v []EnabledActionsEnum) {
	o.EnabledActions = v
}

// GetDisabledFields returns the DisabledFields field value
func (o *CommonModelScopesBodyRequest) GetDisabledFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DisabledFields
}

// GetDisabledFieldsOk returns a tuple with the DisabledFields field value
// and a boolean to check if the value has been set.
func (o *CommonModelScopesBodyRequest) GetDisabledFieldsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisabledFields, true
}

// SetDisabledFields sets field value
func (o *CommonModelScopesBodyRequest) SetDisabledFields(v []string) {
	o.DisabledFields = v
}

func (o CommonModelScopesBodyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model_id"] = o.ModelId
	}
	if true {
		toSerialize["enabled_actions"] = o.EnabledActions
	}
	if true {
		toSerialize["disabled_fields"] = o.DisabledFields
	}
	return json.Marshal(toSerialize)
}

func (v *CommonModelScopesBodyRequest) UnmarshalJSON(src []byte) error {
    type CommonModelScopesBodyRequestUnmarshalTarget CommonModelScopesBodyRequest

	var intermediateResult CommonModelScopesBodyRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = CommonModelScopesBodyRequest(intermediateResult)
	return nil
}
type NullableCommonModelScopesBodyRequest struct {
	value *CommonModelScopesBodyRequest
	isSet bool
}

func (v NullableCommonModelScopesBodyRequest) Get() *CommonModelScopesBodyRequest {
	return v.value
}

func (v *NullableCommonModelScopesBodyRequest) Set(val *CommonModelScopesBodyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonModelScopesBodyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonModelScopesBodyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonModelScopesBodyRequest(val *CommonModelScopesBodyRequest) *NullableCommonModelScopesBodyRequest {
	return &NullableCommonModelScopesBodyRequest{value: val, isSet: true}
}

func (v NullableCommonModelScopesBodyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonModelScopesBodyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


