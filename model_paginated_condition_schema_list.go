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

// PaginatedConditionSchemaList struct for PaginatedConditionSchemaList
type PaginatedConditionSchemaList struct {
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results *[]ConditionSchema `json:"results,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewPaginatedConditionSchemaList instantiates a new PaginatedConditionSchemaList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedConditionSchemaList() *PaginatedConditionSchemaList {
	this := PaginatedConditionSchemaList{}
	return &this
}

// NewPaginatedConditionSchemaListWithDefaults instantiates a new PaginatedConditionSchemaList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedConditionSchemaListWithDefaults() *PaginatedConditionSchemaList {
	this := PaginatedConditionSchemaList{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedConditionSchemaList) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedConditionSchemaList) GetNextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedConditionSchemaList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedConditionSchemaList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedConditionSchemaList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedConditionSchemaList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedConditionSchemaList) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedConditionSchemaList) GetPreviousOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedConditionSchemaList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedConditionSchemaList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedConditionSchemaList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedConditionSchemaList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedConditionSchemaList) GetResults() []ConditionSchema {
	if o == nil || o.Results == nil {
		var ret []ConditionSchema
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedConditionSchemaList) GetResultsOk() (*[]ConditionSchema, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedConditionSchemaList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ConditionSchema and assigns it to the Results field.
func (o *PaginatedConditionSchemaList) SetResults(v []ConditionSchema) {
	o.Results = &v
}

func (o PaginatedConditionSchemaList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

func (v *PaginatedConditionSchemaList) UnmarshalJSON(src []byte) error {
    type PaginatedConditionSchemaListUnmarshalTarget PaginatedConditionSchemaList

	var intermediateResult PaginatedConditionSchemaListUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = PaginatedConditionSchemaList(intermediateResult)
	return nil
}
type NullablePaginatedConditionSchemaList struct {
	value *PaginatedConditionSchemaList
	isSet bool
}

func (v NullablePaginatedConditionSchemaList) Get() *PaginatedConditionSchemaList {
	return v.value
}

func (v *NullablePaginatedConditionSchemaList) Set(val *PaginatedConditionSchemaList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedConditionSchemaList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedConditionSchemaList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedConditionSchemaList(val *PaginatedConditionSchemaList) *NullablePaginatedConditionSchemaList {
	return &NullablePaginatedConditionSchemaList{value: val, isSet: true}
}

func (v NullablePaginatedConditionSchemaList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedConditionSchemaList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


