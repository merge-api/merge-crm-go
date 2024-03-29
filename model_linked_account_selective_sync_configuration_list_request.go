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

// LinkedAccountSelectiveSyncConfigurationListRequest struct for LinkedAccountSelectiveSyncConfigurationListRequest
type LinkedAccountSelectiveSyncConfigurationListRequest struct {
	// The selective syncs associated with a linked account.
	SyncConfigurations []LinkedAccountSelectiveSyncConfigurationRequest `json:"sync_configurations"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewLinkedAccountSelectiveSyncConfigurationListRequest instantiates a new LinkedAccountSelectiveSyncConfigurationListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedAccountSelectiveSyncConfigurationListRequest(syncConfigurations []LinkedAccountSelectiveSyncConfigurationRequest) *LinkedAccountSelectiveSyncConfigurationListRequest {
	this := LinkedAccountSelectiveSyncConfigurationListRequest{}
	this.SyncConfigurations = syncConfigurations
	return &this
}

// NewLinkedAccountSelectiveSyncConfigurationListRequestWithDefaults instantiates a new LinkedAccountSelectiveSyncConfigurationListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedAccountSelectiveSyncConfigurationListRequestWithDefaults() *LinkedAccountSelectiveSyncConfigurationListRequest {
	this := LinkedAccountSelectiveSyncConfigurationListRequest{}
	return &this
}

// GetSyncConfigurations returns the SyncConfigurations field value
func (o *LinkedAccountSelectiveSyncConfigurationListRequest) GetSyncConfigurations() []LinkedAccountSelectiveSyncConfigurationRequest {
	if o == nil {
		var ret []LinkedAccountSelectiveSyncConfigurationRequest
		return ret
	}

	return o.SyncConfigurations
}

// GetSyncConfigurationsOk returns a tuple with the SyncConfigurations field value
// and a boolean to check if the value has been set.
func (o *LinkedAccountSelectiveSyncConfigurationListRequest) GetSyncConfigurationsOk() (*[]LinkedAccountSelectiveSyncConfigurationRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SyncConfigurations, true
}

// SetSyncConfigurations sets field value
func (o *LinkedAccountSelectiveSyncConfigurationListRequest) SetSyncConfigurations(v []LinkedAccountSelectiveSyncConfigurationRequest) {
	o.SyncConfigurations = v
}

func (o LinkedAccountSelectiveSyncConfigurationListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sync_configurations"] = o.SyncConfigurations
	}
	return json.Marshal(toSerialize)
}

func (v *LinkedAccountSelectiveSyncConfigurationListRequest) UnmarshalJSON(src []byte) error {
    type LinkedAccountSelectiveSyncConfigurationListRequestUnmarshalTarget LinkedAccountSelectiveSyncConfigurationListRequest

	var intermediateResult LinkedAccountSelectiveSyncConfigurationListRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = LinkedAccountSelectiveSyncConfigurationListRequest(intermediateResult)
	return nil
}
type NullableLinkedAccountSelectiveSyncConfigurationListRequest struct {
	value *LinkedAccountSelectiveSyncConfigurationListRequest
	isSet bool
}

func (v NullableLinkedAccountSelectiveSyncConfigurationListRequest) Get() *LinkedAccountSelectiveSyncConfigurationListRequest {
	return v.value
}

func (v *NullableLinkedAccountSelectiveSyncConfigurationListRequest) Set(val *LinkedAccountSelectiveSyncConfigurationListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedAccountSelectiveSyncConfigurationListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedAccountSelectiveSyncConfigurationListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedAccountSelectiveSyncConfigurationListRequest(val *LinkedAccountSelectiveSyncConfigurationListRequest) *NullableLinkedAccountSelectiveSyncConfigurationListRequest {
	return &NullableLinkedAccountSelectiveSyncConfigurationListRequest{value: val, isSet: true}
}

func (v NullableLinkedAccountSelectiveSyncConfigurationListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedAccountSelectiveSyncConfigurationListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


