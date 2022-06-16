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

// AccountDetails struct for AccountDetails
type AccountDetails struct {
	Id *string `json:"id,omitempty"`
	Integration *string `json:"integration,omitempty"`
	IntegrationSlug *string `json:"integration_slug,omitempty"`
	Category NullableCategoryEnum `json:"category,omitempty"`
	EndUserOriginId *string `json:"end_user_origin_id,omitempty"`
	EndUserOrganizationName *string `json:"end_user_organization_name,omitempty"`
	EndUserEmailAddress *string `json:"end_user_email_address,omitempty"`
	Status *string `json:"status,omitempty"`
	WebhookListenerUrl *string `json:"webhook_listener_url,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewAccountDetails instantiates a new AccountDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDetails() *AccountDetails {
	this := AccountDetails{}
	return &this
}

// NewAccountDetailsWithDefaults instantiates a new AccountDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDetailsWithDefaults() *AccountDetails {
	this := AccountDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountDetails) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountDetails) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountDetails) SetId(v string) {
	o.Id = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *AccountDetails) GetIntegration() string {
	if o == nil || o.Integration == nil {
		var ret string
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetIntegrationOk() (*string, bool) {
	if o == nil || o.Integration == nil {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *AccountDetails) HasIntegration() bool {
	if o != nil && o.Integration != nil {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given string and assigns it to the Integration field.
func (o *AccountDetails) SetIntegration(v string) {
	o.Integration = &v
}

// GetIntegrationSlug returns the IntegrationSlug field value if set, zero value otherwise.
func (o *AccountDetails) GetIntegrationSlug() string {
	if o == nil || o.IntegrationSlug == nil {
		var ret string
		return ret
	}
	return *o.IntegrationSlug
}

// GetIntegrationSlugOk returns a tuple with the IntegrationSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetIntegrationSlugOk() (*string, bool) {
	if o == nil || o.IntegrationSlug == nil {
		return nil, false
	}
	return o.IntegrationSlug, true
}

// HasIntegrationSlug returns a boolean if a field has been set.
func (o *AccountDetails) HasIntegrationSlug() bool {
	if o != nil && o.IntegrationSlug != nil {
		return true
	}

	return false
}

// SetIntegrationSlug gets a reference to the given string and assigns it to the IntegrationSlug field.
func (o *AccountDetails) SetIntegrationSlug(v string) {
	o.IntegrationSlug = &v
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountDetails) GetCategory() CategoryEnum {
	if o == nil || o.Category.Get() == nil {
		var ret CategoryEnum
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountDetails) GetCategoryOk() (*CategoryEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *AccountDetails) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableCategoryEnum and assigns it to the Category field.
func (o *AccountDetails) SetCategory(v CategoryEnum) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *AccountDetails) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *AccountDetails) UnsetCategory() {
	o.Category.Unset()
}

// GetEndUserOriginId returns the EndUserOriginId field value if set, zero value otherwise.
func (o *AccountDetails) GetEndUserOriginId() string {
	if o == nil || o.EndUserOriginId == nil {
		var ret string
		return ret
	}
	return *o.EndUserOriginId
}

// GetEndUserOriginIdOk returns a tuple with the EndUserOriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetEndUserOriginIdOk() (*string, bool) {
	if o == nil || o.EndUserOriginId == nil {
		return nil, false
	}
	return o.EndUserOriginId, true
}

// HasEndUserOriginId returns a boolean if a field has been set.
func (o *AccountDetails) HasEndUserOriginId() bool {
	if o != nil && o.EndUserOriginId != nil {
		return true
	}

	return false
}

// SetEndUserOriginId gets a reference to the given string and assigns it to the EndUserOriginId field.
func (o *AccountDetails) SetEndUserOriginId(v string) {
	o.EndUserOriginId = &v
}

// GetEndUserOrganizationName returns the EndUserOrganizationName field value if set, zero value otherwise.
func (o *AccountDetails) GetEndUserOrganizationName() string {
	if o == nil || o.EndUserOrganizationName == nil {
		var ret string
		return ret
	}
	return *o.EndUserOrganizationName
}

// GetEndUserOrganizationNameOk returns a tuple with the EndUserOrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetEndUserOrganizationNameOk() (*string, bool) {
	if o == nil || o.EndUserOrganizationName == nil {
		return nil, false
	}
	return o.EndUserOrganizationName, true
}

// HasEndUserOrganizationName returns a boolean if a field has been set.
func (o *AccountDetails) HasEndUserOrganizationName() bool {
	if o != nil && o.EndUserOrganizationName != nil {
		return true
	}

	return false
}

// SetEndUserOrganizationName gets a reference to the given string and assigns it to the EndUserOrganizationName field.
func (o *AccountDetails) SetEndUserOrganizationName(v string) {
	o.EndUserOrganizationName = &v
}

// GetEndUserEmailAddress returns the EndUserEmailAddress field value if set, zero value otherwise.
func (o *AccountDetails) GetEndUserEmailAddress() string {
	if o == nil || o.EndUserEmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EndUserEmailAddress
}

// GetEndUserEmailAddressOk returns a tuple with the EndUserEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetEndUserEmailAddressOk() (*string, bool) {
	if o == nil || o.EndUserEmailAddress == nil {
		return nil, false
	}
	return o.EndUserEmailAddress, true
}

// HasEndUserEmailAddress returns a boolean if a field has been set.
func (o *AccountDetails) HasEndUserEmailAddress() bool {
	if o != nil && o.EndUserEmailAddress != nil {
		return true
	}

	return false
}

// SetEndUserEmailAddress gets a reference to the given string and assigns it to the EndUserEmailAddress field.
func (o *AccountDetails) SetEndUserEmailAddress(v string) {
	o.EndUserEmailAddress = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountDetails) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountDetails) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AccountDetails) SetStatus(v string) {
	o.Status = &v
}

// GetWebhookListenerUrl returns the WebhookListenerUrl field value if set, zero value otherwise.
func (o *AccountDetails) GetWebhookListenerUrl() string {
	if o == nil || o.WebhookListenerUrl == nil {
		var ret string
		return ret
	}
	return *o.WebhookListenerUrl
}

// GetWebhookListenerUrlOk returns a tuple with the WebhookListenerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDetails) GetWebhookListenerUrlOk() (*string, bool) {
	if o == nil || o.WebhookListenerUrl == nil {
		return nil, false
	}
	return o.WebhookListenerUrl, true
}

// HasWebhookListenerUrl returns a boolean if a field has been set.
func (o *AccountDetails) HasWebhookListenerUrl() bool {
	if o != nil && o.WebhookListenerUrl != nil {
		return true
	}

	return false
}

// SetWebhookListenerUrl gets a reference to the given string and assigns it to the WebhookListenerUrl field.
func (o *AccountDetails) SetWebhookListenerUrl(v string) {
	o.WebhookListenerUrl = &v
}

func (o AccountDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Integration != nil {
		toSerialize["integration"] = o.Integration
	}
	if o.IntegrationSlug != nil {
		toSerialize["integration_slug"] = o.IntegrationSlug
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.EndUserOriginId != nil {
		toSerialize["end_user_origin_id"] = o.EndUserOriginId
	}
	if o.EndUserOrganizationName != nil {
		toSerialize["end_user_organization_name"] = o.EndUserOrganizationName
	}
	if o.EndUserEmailAddress != nil {
		toSerialize["end_user_email_address"] = o.EndUserEmailAddress
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.WebhookListenerUrl != nil {
		toSerialize["webhook_listener_url"] = o.WebhookListenerUrl
	}
	return json.Marshal(toSerialize)
}

func (v *AccountDetails) UnmarshalJSON(src []byte) error {
    type AccountDetailsUnmarshalTarget AccountDetails

	var intermediateResult AccountDetailsUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = AccountDetails(intermediateResult)
	return nil
}
type NullableAccountDetails struct {
	value *AccountDetails
	isSet bool
}

func (v NullableAccountDetails) Get() *AccountDetails {
	return v.value
}

func (v *NullableAccountDetails) Set(val *AccountDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDetails(val *AccountDetails) *NullableAccountDetails {
	return &NullableAccountDetails{value: val, isSet: true}
}

func (v NullableAccountDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


