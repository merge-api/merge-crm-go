# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Street1** | Pointer to **NullableString** | Line 1 of the address&#39;s street. | [optional] 
**Street2** | Pointer to **NullableString** | Line 2 of the address&#39;s street. | [optional] 
**City** | Pointer to **NullableString** | The address&#39;s city. | [optional] 
**State** | Pointer to **NullableString** | The address&#39;s state. | [optional] 
**PostalCode** | Pointer to **NullableString** | The address&#39;s postal code. | [optional] 
**Country** | Pointer to [**NullableCountryEnum**](CountryEnum.md) | The address&#39;s country. | [optional] 
**AddressType** | Pointer to [**NullableAddressTypeEnum**](AddressTypeEnum.md) | The address type. | [optional] 

## Methods

### NewAddress

`func NewAddress() *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreet1

`func (o *Address) GetStreet1() string`

GetStreet1 returns the Street1 field if non-nil, zero value otherwise.

### GetStreet1Ok

`func (o *Address) GetStreet1Ok() (*string, bool)`

GetStreet1Ok returns a tuple with the Street1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet1

`func (o *Address) SetStreet1(v string)`

SetStreet1 sets Street1 field to given value.

### HasStreet1

`func (o *Address) HasStreet1() bool`

HasStreet1 returns a boolean if a field has been set.

### SetStreet1Nil

`func (o *Address) SetStreet1Nil(b bool)`

 SetStreet1Nil sets the value for Street1 to be an explicit nil

### UnsetStreet1
`func (o *Address) UnsetStreet1()`

UnsetStreet1 ensures that no value is present for Street1, not even an explicit nil
### GetStreet2

`func (o *Address) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *Address) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *Address) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *Address) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### SetStreet2Nil

`func (o *Address) SetStreet2Nil(b bool)`

 SetStreet2Nil sets the value for Street2 to be an explicit nil

### UnsetStreet2
`func (o *Address) UnsetStreet2()`

UnsetStreet2 ensures that no value is present for Street2, not even an explicit nil
### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Address) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *Address) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *Address) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *Address) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Address) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Address) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Address) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *Address) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Address) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetPostalCode

`func (o *Address) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Address) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *Address) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *Address) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *Address) GetCountry() CountryEnum`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*CountryEnum, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v CountryEnum)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Address) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Address) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Address) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetAddressType

`func (o *Address) GetAddressType() AddressTypeEnum`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *Address) GetAddressTypeOk() (*AddressTypeEnum, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *Address) SetAddressType(v AddressTypeEnum)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *Address) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### SetAddressTypeNil

`func (o *Address) SetAddressTypeNil(b bool)`

 SetAddressTypeNil sets the value for AddressType to be an explicit nil

### UnsetAddressType
`func (o *Address) UnsetAddressType()`

UnsetAddressType ensures that no value is present for AddressType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


