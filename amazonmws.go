// Package amazonmws provides methods for interacting with the Amazon Marketplace Services API.
package amazonmws

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strconv"
)

// FeeEstimateRequest is struct
type FeeEstimateRequest struct {
	IDValue             string
	PriceToEstimateFees float64
	Currency            string
	MarketplaceID       string
	IDType              string
	Identifier          string
	IsAmazonFulfilled   bool
}

func (f *FeeEstimateRequest) requestString(index int, key string) string {
	var buffer bytes.Buffer
	buffer.WriteString("FeesEstimateRequestList.FeesEstimateRequest.")
	buffer.WriteString(strconv.Itoa(index))
	buffer.WriteString(".")
	buffer.WriteString(key)
	return buffer.String()
}

func (f *FeeEstimateRequest) setDefaults(mid string) {
	if f.Currency == "" {
		f.Currency = "USD"
	}

	if f.MarketplaceID == "" {
		f.MarketplaceID = mid
	}

	if f.IDType == "" {
		f.IDType = "ASIN"
	}

	if f.Identifier == "" {
		f.Identifier = f.IDValue
	}

	f.IsAmazonFulfilled = true
}

func (f *FeeEstimateRequest) toQuery(index int, marketplaceID string) map[string]string {
	output := make(map[string]string)

	f.setDefaults(marketplaceID)
	output[f.requestString(index+1, "IdValue")] = f.IDValue
	output[f.requestString(index+1, "PriceToEstimateFees.ListingPrice.Amount")] = strconv.FormatFloat(f.PriceToEstimateFees, 'f', 2, 32)
	output[f.requestString(index+1, "PriceToEstimateFees.ListingPrice.CurrencyCode")] = f.Currency
	output[f.requestString(index+1, "PriceToEstimateFees.Shipping.Amount")] = "0"
	output[f.requestString(index+1, "PriceToEstimateFees.Shipping.CurrencyCode")] = f.Currency
	output[f.requestString(index+1, "PriceToEstimateFees.Points.PointsNumber")] = "0"
	output[f.requestString(index+1, "PriceToEstimateFees.Points.PointsMonetaryValue.Amount")] = "0"
	output[f.requestString(index+1, "PriceToEstimateFees.Points.PointsMonetaryValue.CurrencyCode")] = f.Currency
	output[f.requestString(index+1, "MarketplaceId")] = f.MarketplaceID
	output[f.requestString(index+1, "IdType")] = f.IDType
	output[f.requestString(index+1, "Identifier")] = f.Identifier

	var isFba string
	if f.IsAmazonFulfilled {
		isFba = "true"
	} else {
		isFba = "false"
	}

	output[f.requestString(index+1, "IsAmazonFulfilled")] = isFba

	return output
}

// ListMatchingProducts - returns a list of products and their attributes, based on a search query.
func (api AmazonMWSAPI) ListMatchingProducts(query, queryContextID string) (string, error) {
	params := make(map[string]string)

	params["MarketplaceId"] = string(api.MarketplaceId)
	params["Query"] = query

	if queryContextID != "" {
		params["QueryContextId"] = queryContextID
	}

	return api.genSignAndFetch("ListMatchingProducts", "/Products/2011-10-01", params)
}

/*
GetLowestOfferListingsForASIN takes a list of ASINs and returns the result.
*/
func (api AmazonMWSAPI) GetLowestOfferListingsForASIN(items []string, condition *string) (*MwsLolGetLowestOfferListingsForASINResponse, error) {
	params := make(map[string]string)

	for k, v := range items {
		key := fmt.Sprintf("ASINList.ASIN.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = string(api.MarketplaceId)

	if condition != nil {
		params["ItemCondition"] = *condition
	}

	xmlStr, err := api.genSignAndFetch("GetLowestOfferListingsForASIN", "/Products/2011-10-01", params)
	var resp MwsLolGetLowestOfferListingsForASINResponse
	xml.Unmarshal([]byte(xmlStr), &resp)
	if resp.Attrxmlns == "" {
		return nil, fmt.Errorf("error xml: %s", xmlStr)
	}
	return &resp, err
}

/*
GetCompetitivePricingForASIN takes a list of ASINs and returns the result.
*/
func (api AmazonMWSAPI) GetCompetitivePricingForASIN(items []string, condition *string) (*MwsComGetCompetitivePricingForASINResponse, error) {
	params := make(map[string]string)

	for k, v := range items {
		key := fmt.Sprintf("ASINList.ASIN.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = string(api.MarketplaceId)

	if condition != nil {
		params["ItemCondition"] = *condition
	}

	xmlStr, err := api.genSignAndFetch("GetCompetitivePricingForASIN", "/Products/2011-10-01", params)
	if err != nil {
		return nil, err
	}
	var resp MwsComGetCompetitivePricingForASINResponse
	xml.Unmarshal([]byte(xmlStr), &resp)
	if resp.Attrxmlns == "" {
		return nil, fmt.Errorf("error xml: %s", xmlStr)
	}
	return &resp, err
}

// GetMatchingProductForID is mws api
func (api AmazonMWSAPI) GetMatchingProductForID(idType string, idList []string) (*MwsMatGetMatchingProductForIdResponse, error) {
	params := make(map[string]string)

	for k, v := range idList {
		key := fmt.Sprintf("IdList.Id.%d", (k + 1))
		params[key] = string(v)
	}

	params["IdType"] = idType
	params["MarketplaceId"] = string(api.MarketplaceId)

	xmlStr, err := api.genSignAndFetch("GetMatchingProductForId", "/Products/2011-10-01", params)
	if err != nil {
		return nil, err
	}
	var resp MwsMatGetMatchingProductForIdResponse
	xml.Unmarshal([]byte(xmlStr), &resp)
	if resp.Attrxmlns == "" {
		return nil, fmt.Errorf("error xml: %s", xmlStr)
	}
	return &resp, nil
}

// GetMyFeesEstimate is mws api
func (api AmazonMWSAPI) GetMyFeesEstimate(items []FeeEstimateRequest) (string, error) {
	params := make(map[string]string)

	for index, item := range items {
		queryItems := item.toQuery(index, api.MarketplaceId)

		for key, value := range queryItems {
			params[key] = value
		}
	}

	return api.genSignAndFetchViaPost("GetMyFeesEstimate", "/Products/2011-10-01", params)
}

// GetReportRequestStatus is mws api
func (api AmazonMWSAPI) GetReportRequestStatus(reportID string) (string, error) {
	params := make(map[string]string)

	params["ReportRequestIdList.Id.1"] = reportID

	return api.genSignAndFetch("GetReportRequestList", "/Reports/2009-01-01", params)
}
