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
	"fmt"
)

// AccountDetailsAndActionsStatusEnum * `COMPLETE` - COMPLETE * `INCOMPLETE` - INCOMPLETE * `RELINK_NEEDED` - RELINK_NEEDED
type AccountDetailsAndActionsStatusEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of AccountDetailsAndActionsStatusEnum
const (
    ACCOUNTDETAILSANDACTIONSSTATUSENUM_MERGE_NONSTANDARD_VALUE AccountDetailsAndActionsStatusEnum = "MERGE_NONSTANDARD_VALUE"
    
	ACCOUNTDETAILSANDACTIONSSTATUSENUM_COMPLETE AccountDetailsAndActionsStatusEnum = "COMPLETE"
	ACCOUNTDETAILSANDACTIONSSTATUSENUM_INCOMPLETE AccountDetailsAndActionsStatusEnum = "INCOMPLETE"
	ACCOUNTDETAILSANDACTIONSSTATUSENUM_RELINK_NEEDED AccountDetailsAndActionsStatusEnum = "RELINK_NEEDED"
)

var allowedAccountDetailsAndActionsStatusEnumEnumValues = []AccountDetailsAndActionsStatusEnum{
	"COMPLETE",
	"INCOMPLETE",
	"RELINK_NEEDED",
}

func (v *AccountDetailsAndActionsStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountDetailsAndActionsStatusEnum(value)
	for _, existing := range allowedAccountDetailsAndActionsStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = ACCOUNTDETAILSANDACTIONSSTATUSENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewAccountDetailsAndActionsStatusEnumFromValue returns a pointer to a valid AccountDetailsAndActionsStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountDetailsAndActionsStatusEnumFromValue(v string) (*AccountDetailsAndActionsStatusEnum, error) {
	ev := AccountDetailsAndActionsStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := ACCOUNTDETAILSANDACTIONSSTATUSENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountDetailsAndActionsStatusEnum) IsValid() bool {
	for _, existing := range allowedAccountDetailsAndActionsStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountDetailsAndActionsStatusEnum value
func (v AccountDetailsAndActionsStatusEnum) Ptr() *AccountDetailsAndActionsStatusEnum {
	return &v
}

type NullableAccountDetailsAndActionsStatusEnum struct {
	value *AccountDetailsAndActionsStatusEnum
	isSet bool
}

func (v NullableAccountDetailsAndActionsStatusEnum) Get() *AccountDetailsAndActionsStatusEnum {
	return v.value
}

func (v *NullableAccountDetailsAndActionsStatusEnum) Set(val *AccountDetailsAndActionsStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDetailsAndActionsStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDetailsAndActionsStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDetailsAndActionsStatusEnum(val *AccountDetailsAndActionsStatusEnum) *NullableAccountDetailsAndActionsStatusEnum {
	return &NullableAccountDetailsAndActionsStatusEnum{value: val, isSet: true}
}

func (v NullableAccountDetailsAndActionsStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDetailsAndActionsStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

