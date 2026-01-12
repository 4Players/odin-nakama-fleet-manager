# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** | The name of the city | 
**CityDisplay** | **string** | The display-friendly name of the city | 
**Continent** | **string** | The continent name | 
**Country** | **string** | The country name | 
**IsProtected** | **bool** | Indicates if the location is password protected | 

## Methods

### NewLocation

`func NewLocation(city string, cityDisplay string, continent string, country string, isProtected bool, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *Location) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Location) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Location) SetCity(v string)`

SetCity sets City field to given value.


### GetCityDisplay

`func (o *Location) GetCityDisplay() string`

GetCityDisplay returns the CityDisplay field if non-nil, zero value otherwise.

### GetCityDisplayOk

`func (o *Location) GetCityDisplayOk() (*string, bool)`

GetCityDisplayOk returns a tuple with the CityDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityDisplay

`func (o *Location) SetCityDisplay(v string)`

SetCityDisplay sets CityDisplay field to given value.


### GetContinent

`func (o *Location) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *Location) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *Location) SetContinent(v string)`

SetContinent sets Continent field to given value.


### GetCountry

`func (o *Location) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Location) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Location) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetIsProtected

`func (o *Location) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *Location) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *Location) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)