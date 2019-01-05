package amazonmws

import "encoding/xml"

type MwsLolASIN struct {
	XMLName xml.Name `xml:"ASIN,omitempty" json:"ASIN,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsLolAllOfferListingsConsidered struct {
	XMLName xml.Name `xml:"AllOfferListingsConsidered,omitempty" json:"AllOfferListingsConsidered,omitempty"`
	Flag bool `xml:",chardata" json:",omitempty"`
}

type MwsLolAmount struct {
	XMLName xml.Name `xml:"Amount,omitempty" json:"Amount,omitempty"`
	Number float32 `xml:",chardata" json:",omitempty"`
}

type MwsLolCurrencyCode struct {
	XMLName xml.Name `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsLolFulfillmentChannel struct {
	XMLName xml.Name `xml:"FulfillmentChannel,omitempty" json:"FulfillmentChannel,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsLolGetLowestOfferListingsForASINResponse struct {
	XMLName xml.Name `xml:"GetLowestOfferListingsForASINResponse,omitempty" json:"GetLowestOfferListingsForASINResponse,omitempty"`
	Attrxmlns string`xml:"xmlns,attr"  json:",omitempty"`
	MwsLolGetLowestOfferListingsForASINResult []*MwsLolGetLowestOfferListingsForASINResult `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 GetLowestOfferListingsForASINResult,omitempty" json:"GetLowestOfferListingsForASINResult,omitempty"`
	MwsLolResponseMetadata *MwsLolResponseMetadata `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ResponseMetadata,omitempty" json:"ResponseMetadata,omitempty"`
}

type MwsLolGetLowestOfferListingsForASINResult struct {
	XMLName xml.Name `xml:"GetLowestOfferListingsForASINResult,omitempty" json:"GetLowestOfferListingsForASINResult,omitempty"`
	AttrASIN string`xml:"ASIN,attr"  json:",omitempty"`
	Attrstatus string`xml:"status,attr"  json:",omitempty"`
	MwsLolAllOfferListingsConsidered *MwsLolAllOfferListingsConsidered `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 AllOfferListingsConsidered,omitempty" json:"AllOfferListingsConsidered,omitempty"`
	MwsLolProduct *MwsLolProduct `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Product,omitempty" json:"Product,omitempty"`
}

type MwsLolIdentifiers struct {
	XMLName xml.Name `xml:"Identifiers,omitempty" json:"Identifiers,omitempty"`
	MwsLolMarketplaceASIN *MwsLolMarketplaceASIN `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 MarketplaceASIN,omitempty" json:"MarketplaceASIN,omitempty"`
}

type MwsLolItemCondition struct {
	XMLName xml.Name `xml:"ItemCondition,omitempty" json:"ItemCondition,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsLolItemSubcondition struct {
	XMLName xml.Name `xml:"ItemSubcondition,omitempty" json:"ItemSubcondition,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsLolLandedPrice struct {
	XMLName xml.Name `xml:"LandedPrice,omitempty" json:"LandedPrice,omitempty"`
	MwsLolAmount *MwsLolAmount `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Amount,omitempty" json:"Amount,omitempty"`
	MwsLolCurrencyCode *MwsLolCurrencyCode `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 CurrencyCode,omitempty" json:"CurrencyCode,omitempty"`
}

type MwsLolListingPrice struct {
	XMLName xml.Name `xml:"ListingPrice,omitempty" json:"ListingPrice,omitempty"`
	MwsLolAmount *MwsLolAmount `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Amount,omitempty" json:"Amount,omitempty"`
	MwsLolCurrencyCode *MwsLolCurrencyCode `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 CurrencyCode,omitempty" json:"CurrencyCode,omitempty"`
}

type MwsLolLowestOfferListing struct {
	XMLName xml.Name `xml:"LowestOfferListing,omitempty" json:"LowestOfferListing,omitempty"`
	MwsLolMultipleOffersAtLowestPrice *MwsLolMultipleOffersAtLowestPrice `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 MultipleOffersAtLowestPrice,omitempty" json:"MultipleOffersAtLowestPrice,omitempty"`
	MwsLolNumberOfOfferListingsConsidered *MwsLolNumberOfOfferListingsConsidered `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 NumberOfOfferListingsConsidered,omitempty" json:"NumberOfOfferListingsConsidered,omitempty"`
	MwsLolPrice *MwsLolPrice `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Price,omitempty" json:"Price,omitempty"`
	MwsLolQualifiers *MwsLolQualifiers `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Qualifiers,omitempty" json:"Qualifiers,omitempty"`
	MwsLolSellerFeedbackCount *MwsLolSellerFeedbackCount `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 SellerFeedbackCount,omitempty" json:"SellerFeedbackCount,omitempty"`
}

type MwsLolLowestOfferListings struct {
	XMLName xml.Name `xml:"LowestOfferListings,omitempty" json:"LowestOfferListings,omitempty"`
	MwsLolLowestOfferListing []*MwsLolLowestOfferListing `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 LowestOfferListing,omitempty" json:"LowestOfferListing,omitempty"`
}

type MwsLolMarketplaceASIN struct {
	XMLName xml.Name `xml:"MarketplaceASIN,omitempty" json:"MarketplaceASIN,omitempty"`
	MwsLolASIN *MwsLolASIN `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ASIN,omitempty" json:"ASIN,omitempty"`
	MwsLolMarketplaceId *MwsLolMarketplaceId `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 MarketplaceId,omitempty" json:"MarketplaceId,omitempty"`
}

type MwsLolMarketplaceId struct {
	XMLName xml.Name `xml:"MarketplaceId,omitempty" json:"MarketplaceId,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsLolMax struct {
	XMLName xml.Name `xml:"Max,omitempty" json:"Max,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsLolMultipleOffersAtLowestPrice struct {
	XMLName xml.Name `xml:"MultipleOffersAtLowestPrice,omitempty" json:"MultipleOffersAtLowestPrice,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsLolNumberOfOfferListingsConsidered struct {
	XMLName xml.Name `xml:"NumberOfOfferListingsConsidered,omitempty" json:"NumberOfOfferListingsConsidered,omitempty"`
	Number int8 `xml:",chardata" json:",omitempty"`
}

type MwsLolPoints struct {
	XMLName xml.Name `xml:"Points,omitempty" json:"Points,omitempty"`
	MwsLolPointsMonetaryValue *MwsLolPointsMonetaryValue `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 PointsMonetaryValue,omitempty" json:"PointsMonetaryValue,omitempty"`
	MwsLolPointsNumber *MwsLolPointsNumber `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 PointsNumber,omitempty" json:"PointsNumber,omitempty"`
}

type MwsLolPointsMonetaryValue struct {
	XMLName xml.Name `xml:"PointsMonetaryValue,omitempty" json:"PointsMonetaryValue,omitempty"`
	MwsLolAmount *MwsLolAmount `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Amount,omitempty" json:"Amount,omitempty"`
	MwsLolCurrencyCode *MwsLolCurrencyCode `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 CurrencyCode,omitempty" json:"CurrencyCode,omitempty"`
}

type MwsLolPointsNumber struct {
	XMLName xml.Name `xml:"PointsNumber,omitempty" json:"PointsNumber,omitempty"`
	Number int8 `xml:",chardata" json:",omitempty"`
}

type MwsLolPrice struct {
	XMLName xml.Name `xml:"Price,omitempty" json:"Price,omitempty"`
	MwsLolLandedPrice *MwsLolLandedPrice `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 LandedPrice,omitempty" json:"LandedPrice,omitempty"`
	MwsLolListingPrice *MwsLolListingPrice `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ListingPrice,omitempty" json:"ListingPrice,omitempty"`
	MwsLolPoints *MwsLolPoints `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Points,omitempty" json:"Points,omitempty"`
	MwsLolShipping *MwsLolShipping `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Shipping,omitempty" json:"Shipping,omitempty"`
}

type MwsLolProduct struct {
	XMLName xml.Name `xml:"Product,omitempty" json:"Product,omitempty"`
	AttrXmlnsns2 string`xml:"xmlns ns2,attr"  json:",omitempty"`
	MwsLolIdentifiers *MwsLolIdentifiers `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Identifiers,omitempty" json:"Identifiers,omitempty"`
	MwsLolLowestOfferListings *MwsLolLowestOfferListings `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 LowestOfferListings,omitempty" json:"LowestOfferListings,omitempty"`
}

type MwsLolQualifiers struct {
	XMLName xml.Name `xml:"Qualifiers,omitempty" json:"Qualifiers,omitempty"`
	MwsLolFulfillmentChannel *MwsLolFulfillmentChannel `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 FulfillmentChannel,omitempty" json:"FulfillmentChannel,omitempty"`
	MwsLolItemCondition *MwsLolItemCondition `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ItemCondition,omitempty" json:"ItemCondition,omitempty"`
	MwsLolItemSubcondition *MwsLolItemSubcondition `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ItemSubcondition,omitempty" json:"ItemSubcondition,omitempty"`
	MwsLolSellerPositiveFeedbackRating *MwsLolSellerPositiveFeedbackRating `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 SellerPositiveFeedbackRating,omitempty" json:"SellerPositiveFeedbackRating,omitempty"`
	MwsLolShippingTime *MwsLolShippingTime `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ShippingTime,omitempty" json:"ShippingTime,omitempty"`
	MwsLolShipsDomestically *MwsLolShipsDomestically `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ShipsDomestically,omitempty" json:"ShipsDomestically,omitempty"`
}

type MwsLolRequestId struct {
	XMLName xml.Name `xml:"RequestId,omitempty" json:"RequestId,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsLolResponseMetadata struct {
	XMLName xml.Name `xml:"ResponseMetadata,omitempty" json:"ResponseMetadata,omitempty"`
	MwsLolRequestId *MwsLolRequestId `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 RequestId,omitempty" json:"RequestId,omitempty"`
}

type MwsLolSellerFeedbackCount struct {
	XMLName xml.Name `xml:"SellerFeedbackCount,omitempty" json:"SellerFeedbackCount,omitempty"`
	Number int32 `xml:",chardata" json:",omitempty"`
}

type MwsLolSellerPositiveFeedbackRating struct {
	XMLName xml.Name `xml:"SellerPositiveFeedbackRating,omitempty" json:"SellerPositiveFeedbackRating,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsLolShipping struct {
	XMLName xml.Name `xml:"Shipping,omitempty" json:"Shipping,omitempty"`
	MwsLolAmount *MwsLolAmount `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Amount,omitempty" json:"Amount,omitempty"`
	MwsLolCurrencyCode *MwsLolCurrencyCode `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 CurrencyCode,omitempty" json:"CurrencyCode,omitempty"`
}

type MwsLolShippingTime struct {
	XMLName xml.Name `xml:"ShippingTime,omitempty" json:"ShippingTime,omitempty"`
	MwsLolMax *MwsLolMax `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Max,omitempty" json:"Max,omitempty"`
}

type MwsLolShipsDomestically struct {
	XMLName xml.Name `xml:"ShipsDomestically,omitempty" json:"ShipsDomestically,omitempty"`
	Flag bool `xml:",chardata" json:",omitempty"`
}

