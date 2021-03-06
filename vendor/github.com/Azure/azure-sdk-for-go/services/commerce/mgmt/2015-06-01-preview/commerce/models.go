package commerce

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"net/http"
)

// AggregationGranularity enumerates the values for aggregation granularity.
type AggregationGranularity string

const (
	// Daily ...
	Daily AggregationGranularity = "Daily"
	// Hourly ...
	Hourly AggregationGranularity = "Hourly"
)

// Name enumerates the values for name.
type Name string

const (
	// NameMonetaryCommitment ...
	NameMonetaryCommitment Name = "Monetary Commitment"
	// NameMonetaryCredit ...
	NameMonetaryCredit Name = "Monetary Credit"
	// NameOfferTermInfo ...
	NameOfferTermInfo Name = "OfferTermInfo"
	// NameRecurringCharge ...
	NameRecurringCharge Name = "Recurring Charge"
)

// ErrorResponse describes the format of Error response.
type ErrorResponse struct {
	// Code - Error code
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// InfoField key-value pairs of instance details in the legacy format.
type InfoField struct {
	// Project - Identifies the name of the instance provisioned by the user.
	Project *string `json:"project,omitempty"`
}

// MeterInfo detailed information about the meter.
type MeterInfo struct {
	// MeterID - The unique identifier of the resource.
	MeterID *uuid.UUID `json:"MeterId,omitempty"`
	// MeterName - The name of the meter, within the given meter category
	MeterName *string `json:"MeterName,omitempty"`
	// MeterCategory - The category of the meter, e.g., 'Cloud services', 'Networking', etc..
	MeterCategory *string `json:"MeterCategory,omitempty"`
	// MeterSubCategory - The subcategory of the meter, e.g., 'A6 Cloud services', 'ExpressRoute (IXP)', etc..
	MeterSubCategory *string `json:"MeterSubCategory,omitempty"`
	// Unit - The unit in which the meter consumption is charged, e.g., 'Hours', 'GB', etc.
	Unit *string `json:"Unit,omitempty"`
	// MeterTags - Provides additional meter data. 'Third Party' indicates a meter with no discount. Blanks indicate First Party.
	MeterTags *[]string `json:"MeterTags,omitempty"`
	// MeterRegion - The region in which the Azure service is available.
	MeterRegion *string `json:"MeterRegion,omitempty"`
	// MeterRates - The list of key/value pairs for the meter rates, in the format 'key':'value' where key = the meter quantity, and value = the corresponding price
	MeterRates *map[string]*float64 `json:"MeterRates,omitempty"`
	// EffectiveDate - Indicates the date from which the meter rate is effective.
	EffectiveDate *date.Time `json:"EffectiveDate,omitempty"`
	// IncludedQuantity - The resource quantity that is included in the offer at no cost. Consumption beyond this quantity will be charged.
	IncludedQuantity *float64 `json:"IncludedQuantity,omitempty"`
}

// MonetaryCommitment indicates that a monetary commitment is required for this offer
type MonetaryCommitment struct {
	// EffectiveDate - Indicates the date from which the offer term is effective.
	EffectiveDate *date.Time `json:"EffectiveDate,omitempty"`
	// Name - Possible values include: 'NameOfferTermInfo', 'NameMonetaryCredit', 'NameMonetaryCommitment', 'NameRecurringCharge'
	Name Name `json:"Name,omitempty"`
	// TieredDiscount - The list of key/value pairs for the tiered meter rates, in the format 'key':'value' where key = price, and value = the corresponding discount percentage. This field is used only by offer terms of type 'Monetary Commitment'.
	TieredDiscount *map[string]*decimal.Decimal `json:"TieredDiscount,omitempty"`
	// ExcludedMeterIds - An array of meter ids that are excluded from the given offer terms.
	ExcludedMeterIds *[]uuid.UUID `json:"ExcludedMeterIds,omitempty"`
}

// MarshalJSON is the custom marshaler for MonetaryCommitment.
func (mc MonetaryCommitment) MarshalJSON() ([]byte, error) {
	mc.Name = NameMonetaryCommitment
	type Alias MonetaryCommitment
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(mc),
	})
}

// AsMonetaryCredit is the BasicOfferTermInfo implementation for MonetaryCommitment.
func (mc MonetaryCommitment) AsMonetaryCredit() (*MonetaryCredit, bool) {
	return nil, false
}

// AsMonetaryCommitment is the BasicOfferTermInfo implementation for MonetaryCommitment.
func (mc MonetaryCommitment) AsMonetaryCommitment() (*MonetaryCommitment, bool) {
	return &mc, true
}

// AsRecurringCharge is the BasicOfferTermInfo implementation for MonetaryCommitment.
func (mc MonetaryCommitment) AsRecurringCharge() (*RecurringCharge, bool) {
	return nil, false
}

// AsOfferTermInfo is the BasicOfferTermInfo implementation for MonetaryCommitment.
func (mc MonetaryCommitment) AsOfferTermInfo() (*OfferTermInfo, bool) {
	return nil, false
}

// AsBasicOfferTermInfo is the BasicOfferTermInfo implementation for MonetaryCommitment.
func (mc MonetaryCommitment) AsBasicOfferTermInfo() (BasicOfferTermInfo, bool) {
	return &mc, true
}

// MonetaryCredit indicates that this is a monetary credit offer.
type MonetaryCredit struct {
	// EffectiveDate - Indicates the date from which the offer term is effective.
	EffectiveDate *date.Time `json:"EffectiveDate,omitempty"`
	// Name - Possible values include: 'NameOfferTermInfo', 'NameMonetaryCredit', 'NameMonetaryCommitment', 'NameRecurringCharge'
	Name Name `json:"Name,omitempty"`
	// Credit - The amount of credit provided under the terms of the given offer level.
	Credit *decimal.Decimal `json:"Credit,omitempty"`
	// ExcludedMeterIds - An array of meter ids that are excluded from the given offer terms.
	ExcludedMeterIds *[]uuid.UUID `json:"ExcludedMeterIds,omitempty"`
}

// MarshalJSON is the custom marshaler for MonetaryCredit.
func (mc MonetaryCredit) MarshalJSON() ([]byte, error) {
	mc.Name = NameMonetaryCredit
	type Alias MonetaryCredit
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(mc),
	})
}

// AsMonetaryCredit is the BasicOfferTermInfo implementation for MonetaryCredit.
func (mc MonetaryCredit) AsMonetaryCredit() (*MonetaryCredit, bool) {
	return &mc, true
}

// AsMonetaryCommitment is the BasicOfferTermInfo implementation for MonetaryCredit.
func (mc MonetaryCredit) AsMonetaryCommitment() (*MonetaryCommitment, bool) {
	return nil, false
}

// AsRecurringCharge is the BasicOfferTermInfo implementation for MonetaryCredit.
func (mc MonetaryCredit) AsRecurringCharge() (*RecurringCharge, bool) {
	return nil, false
}

// AsOfferTermInfo is the BasicOfferTermInfo implementation for MonetaryCredit.
func (mc MonetaryCredit) AsOfferTermInfo() (*OfferTermInfo, bool) {
	return nil, false
}

// AsBasicOfferTermInfo is the BasicOfferTermInfo implementation for MonetaryCredit.
func (mc MonetaryCredit) AsBasicOfferTermInfo() (BasicOfferTermInfo, bool) {
	return &mc, true
}

// BasicOfferTermInfo describes the offer term.
type BasicOfferTermInfo interface {
	AsMonetaryCredit() (*MonetaryCredit, bool)
	AsMonetaryCommitment() (*MonetaryCommitment, bool)
	AsRecurringCharge() (*RecurringCharge, bool)
	AsOfferTermInfo() (*OfferTermInfo, bool)
}

// OfferTermInfo describes the offer term.
type OfferTermInfo struct {
	// EffectiveDate - Indicates the date from which the offer term is effective.
	EffectiveDate *date.Time `json:"EffectiveDate,omitempty"`
	// Name - Possible values include: 'NameOfferTermInfo', 'NameMonetaryCredit', 'NameMonetaryCommitment', 'NameRecurringCharge'
	Name Name `json:"Name,omitempty"`
}

func unmarshalBasicOfferTermInfo(body []byte) (BasicOfferTermInfo, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["Name"] {
	case string(NameMonetaryCredit):
		var mc MonetaryCredit
		err := json.Unmarshal(body, &mc)
		return mc, err
	case string(NameMonetaryCommitment):
		var mc MonetaryCommitment
		err := json.Unmarshal(body, &mc)
		return mc, err
	case string(NameRecurringCharge):
		var rc RecurringCharge
		err := json.Unmarshal(body, &rc)
		return rc, err
	default:
		var oti OfferTermInfo
		err := json.Unmarshal(body, &oti)
		return oti, err
	}
}
func unmarshalBasicOfferTermInfoArray(body []byte) ([]BasicOfferTermInfo, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	otiArray := make([]BasicOfferTermInfo, len(rawMessages))

	for index, rawMessage := range rawMessages {
		oti, err := unmarshalBasicOfferTermInfo(*rawMessage)
		if err != nil {
			return nil, err
		}
		otiArray[index] = oti
	}
	return otiArray, nil
}

// MarshalJSON is the custom marshaler for OfferTermInfo.
func (oti OfferTermInfo) MarshalJSON() ([]byte, error) {
	oti.Name = NameOfferTermInfo
	type Alias OfferTermInfo
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(oti),
	})
}

// AsMonetaryCredit is the BasicOfferTermInfo implementation for OfferTermInfo.
func (oti OfferTermInfo) AsMonetaryCredit() (*MonetaryCredit, bool) {
	return nil, false
}

// AsMonetaryCommitment is the BasicOfferTermInfo implementation for OfferTermInfo.
func (oti OfferTermInfo) AsMonetaryCommitment() (*MonetaryCommitment, bool) {
	return nil, false
}

// AsRecurringCharge is the BasicOfferTermInfo implementation for OfferTermInfo.
func (oti OfferTermInfo) AsRecurringCharge() (*RecurringCharge, bool) {
	return nil, false
}

// AsOfferTermInfo is the BasicOfferTermInfo implementation for OfferTermInfo.
func (oti OfferTermInfo) AsOfferTermInfo() (*OfferTermInfo, bool) {
	return &oti, true
}

// AsBasicOfferTermInfo is the BasicOfferTermInfo implementation for OfferTermInfo.
func (oti OfferTermInfo) AsBasicOfferTermInfo() (BasicOfferTermInfo, bool) {
	return &oti, true
}

// RateCardQueryParameters parameters that are used in the odata $filter query parameter for providing RateCard
// information.
type RateCardQueryParameters struct {
	// OfferDurableID - The Offer ID parameter consists of the 'MS-AZR-' prefix, plus the Offer ID number (e.g., MS-AZR-0026P). See https://azure.microsoft.com/en-us/support/legal/offer-details/ for more information on the list of available Offer IDs, country/region availability, and billing currency.
	OfferDurableID *string `json:"OfferDurableId,omitempty"`
	// Currency - The currency in which the rates need to be provided.
	Currency *string `json:"Currency,omitempty"`
	// Locale - The culture in which the resource metadata needs to be localized.
	Locale *string `json:"Locale,omitempty"`
	// RegionInfo - 2 letter ISO code where the offer was purchased.
	RegionInfo *string `json:"RegionInfo,omitempty"`
}

// RecurringCharge indicates a recurring charge is present for this offer.
type RecurringCharge struct {
	// EffectiveDate - Indicates the date from which the offer term is effective.
	EffectiveDate *date.Time `json:"EffectiveDate,omitempty"`
	// Name - Possible values include: 'NameOfferTermInfo', 'NameMonetaryCredit', 'NameMonetaryCommitment', 'NameRecurringCharge'
	Name Name `json:"Name,omitempty"`
	// RecurringCharge - The amount of recurring charge as per the offer term.
	RecurringCharge *int32 `json:"RecurringCharge,omitempty"`
}

// MarshalJSON is the custom marshaler for RecurringCharge.
func (rc RecurringCharge) MarshalJSON() ([]byte, error) {
	rc.Name = NameRecurringCharge
	type Alias RecurringCharge
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(rc),
	})
}

// AsMonetaryCredit is the BasicOfferTermInfo implementation for RecurringCharge.
func (rc RecurringCharge) AsMonetaryCredit() (*MonetaryCredit, bool) {
	return nil, false
}

// AsMonetaryCommitment is the BasicOfferTermInfo implementation for RecurringCharge.
func (rc RecurringCharge) AsMonetaryCommitment() (*MonetaryCommitment, bool) {
	return nil, false
}

// AsRecurringCharge is the BasicOfferTermInfo implementation for RecurringCharge.
func (rc RecurringCharge) AsRecurringCharge() (*RecurringCharge, bool) {
	return &rc, true
}

// AsOfferTermInfo is the BasicOfferTermInfo implementation for RecurringCharge.
func (rc RecurringCharge) AsOfferTermInfo() (*OfferTermInfo, bool) {
	return nil, false
}

// AsBasicOfferTermInfo is the BasicOfferTermInfo implementation for RecurringCharge.
func (rc RecurringCharge) AsBasicOfferTermInfo() (BasicOfferTermInfo, bool) {
	return &rc, true
}

// ResourceRateCardInfo price and Metadata information for resources
type ResourceRateCardInfo struct {
	autorest.Response `json:"-"`
	// Currency - The currency in which the rates are provided.
	Currency *string `json:"Currency,omitempty"`
	// Locale - The culture in which the resource information is localized.
	Locale *string `json:"Locale,omitempty"`
	// IsTaxIncluded - All rates are pretax, so this will always be returned as 'false'.
	IsTaxIncluded *bool `json:"IsTaxIncluded,omitempty"`
	// OfferTerms - A list of offer terms.
	OfferTerms *[]BasicOfferTermInfo `json:"OfferTerms,omitempty"`
	// Meters - A list of meters.
	Meters *[]MeterInfo `json:"Meters,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for ResourceRateCardInfo struct.
func (rrci *ResourceRateCardInfo) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["Currency"]
	if v != nil {
		var currency string
		err = json.Unmarshal(*m["Currency"], &currency)
		if err != nil {
			return err
		}
		rrci.Currency = &currency
	}

	v = m["Locale"]
	if v != nil {
		var locale string
		err = json.Unmarshal(*m["Locale"], &locale)
		if err != nil {
			return err
		}
		rrci.Locale = &locale
	}

	v = m["IsTaxIncluded"]
	if v != nil {
		var isTaxIncluded bool
		err = json.Unmarshal(*m["IsTaxIncluded"], &isTaxIncluded)
		if err != nil {
			return err
		}
		rrci.IsTaxIncluded = &isTaxIncluded
	}

	v = m["OfferTerms"]
	if v != nil {
		offerTerms, err := unmarshalBasicOfferTermInfoArray(*m["OfferTerms"])
		if err != nil {
			return err
		}
		rrci.OfferTerms = &offerTerms
	}

	v = m["Meters"]
	if v != nil {
		var meters []MeterInfo
		err = json.Unmarshal(*m["Meters"], &meters)
		if err != nil {
			return err
		}
		rrci.Meters = &meters
	}

	return nil
}

// UsageAggregation describes the usageAggregation.
type UsageAggregation struct {
	// ID - Unique Id for the usage aggregate.
	ID *string `json:"id,omitempty"`
	// Name - Name of the usage aggregate.
	Name *string `json:"name,omitempty"`
	// Type - Type of the resource being returned.
	Type *string `json:"type,omitempty"`
	// UsageSample - Usage data.
	*UsageSample `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for UsageAggregation struct.
func (ua *UsageAggregation) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		ua.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		ua.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		ua.Type = &typeVar
	}

	v = m["properties"]
	if v != nil {
		var properties UsageSample
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		ua.UsageSample = &properties
	}

	return nil
}

// UsageAggregationListResult the Get UsageAggregates operation response.
type UsageAggregationListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets details for the requested aggregation.
	Value *[]UsageAggregation `json:"value,omitempty"`
	// NextLink - Gets or sets the link to the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// UsageAggregationListResultIterator provides access to a complete listing of UsageAggregation values.
type UsageAggregationListResultIterator struct {
	i    int
	page UsageAggregationListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *UsageAggregationListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter UsageAggregationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter UsageAggregationListResultIterator) Response() UsageAggregationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter UsageAggregationListResultIterator) Value() UsageAggregation {
	if !iter.page.NotDone() {
		return UsageAggregation{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (ualr UsageAggregationListResult) IsEmpty() bool {
	return ualr.Value == nil || len(*ualr.Value) == 0
}

// usageAggregationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ualr UsageAggregationListResult) usageAggregationListResultPreparer() (*http.Request, error) {
	if ualr.NextLink == nil || len(to.String(ualr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ualr.NextLink)))
}

// UsageAggregationListResultPage contains a page of UsageAggregation values.
type UsageAggregationListResultPage struct {
	fn   func(UsageAggregationListResult) (UsageAggregationListResult, error)
	ualr UsageAggregationListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *UsageAggregationListResultPage) Next() error {
	next, err := page.fn(page.ualr)
	if err != nil {
		return err
	}
	page.ualr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page UsageAggregationListResultPage) NotDone() bool {
	return !page.ualr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page UsageAggregationListResultPage) Response() UsageAggregationListResult {
	return page.ualr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page UsageAggregationListResultPage) Values() []UsageAggregation {
	if page.ualr.IsEmpty() {
		return nil
	}
	return *page.ualr.Value
}

// UsageSample describes a sample of the usageAggregation.
type UsageSample struct {
	// SubscriptionID - The subscription identifier for the Azure user.
	SubscriptionID *uuid.UUID `json:"subscriptionId,omitempty"`
	// MeterID - Unique ID for the resource that was consumed (aka ResourceID).
	MeterID *string `json:"meterId,omitempty"`
	// UsageStartTime - UTC start time for the usage bucket to which this usage aggregate belongs.
	UsageStartTime *date.Time `json:"usageStartTime,omitempty"`
	// UsageEndTime - UTC end time for the usage bucket to which this usage aggregate belongs.
	UsageEndTime *date.Time `json:"usageEndTime,omitempty"`
	// Quantity - The amount of the resource consumption that occurred in this time frame.
	Quantity *map[string]interface{} `json:"quantity,omitempty"`
	// Unit - The unit in which the usage for this resource is being counted, e.g. Hours, GB.
	Unit *string `json:"unit,omitempty"`
	// MeterName - Friendly name of the resource being consumed.
	MeterName *string `json:"meterName,omitempty"`
	// MeterCategory - Category of the consumed resource.
	MeterCategory *string `json:"meterCategory,omitempty"`
	// MeterSubCategory - Sub-category of the consumed resource.
	MeterSubCategory *string `json:"meterSubCategory,omitempty"`
	// MeterRegion - Region of the meterId used for billing purposes
	MeterRegion *string `json:"meterRegion,omitempty"`
	// InfoFields - Key-value pairs of instance details (legacy format).
	InfoFields *InfoField `json:"infoFields,omitempty"`
	// InstanceData - Key-value pairs of instance details represented as a string.
	InstanceData *string `json:"instanceData,omitempty"`
}
