package entity

import "encoding/xml"

type MwsComASIN struct {
	XMLName xml.Name `xml:"ASIN,omitempty" json:"ASIN,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsComAmount struct {
	XMLName xml.Name `xml:"Amount,omitempty" json:"Amount,omitempty"`
	Number float32 `xml:",chardata" json:",omitempty"`
}

type MwsComCompetitivePrice struct {
	XMLName xml.Name `xml:"CompetitivePrice,omitempty" json:"CompetitivePrice,omitempty"`
	AttrbelongsToRequester string`xml:"belongsToRequester,attr"  json:",omitempty"`
	Attrcondition string`xml:"condition,attr"  json:",omitempty"`
	Attrsubcondition string`xml:"subcondition,attr"  json:",omitempty"`
	MwsComCompetitivePriceId *MwsComCompetitivePriceId `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 CompetitivePriceId,omitempty" json:"CompetitivePriceId,omitempty"`
	MwsComPrice *MwsComPrice `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Price,omitempty" json:"Price,omitempty"`
}

type MwsComCompetitivePriceId struct {
	XMLName xml.Name `xml:"CompetitivePriceId,omitempty" json:"CompetitivePriceId,omitempty"`
	Flag bool `xml:",chardata" json:",omitempty"`
}

type MwsComCompetitivePrices struct {
	XMLName xml.Name `xml:"CompetitivePrices,omitempty" json:"CompetitivePrices,omitempty"`
	MwsComCompetitivePrice *MwsComCompetitivePrice `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 CompetitivePrice,omitempty" json:"CompetitivePrice,omitempty"`
}

type MwsComCompetitivePricing struct {
	XMLName xml.Name `xml:"CompetitivePricing,omitempty" json:"CompetitivePricing,omitempty"`
	MwsComCompetitivePrices *MwsComCompetitivePrices `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 CompetitivePrices,omitempty" json:"CompetitivePrices,omitempty"`
	MwsComNumberOfOfferListings *MwsComNumberOfOfferListings `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 NumberOfOfferListings,omitempty" json:"NumberOfOfferListings,omitempty"`
}

type MwsComCurrencyCode struct {
	XMLName xml.Name `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsComGetCompetitivePricingForASINResponse struct {
	XMLName xml.Name `xml:"GetCompetitivePricingForASINResponse,omitempty" json:"GetCompetitivePricingForASINResponse,omitempty"`
	Attrxmlns string`xml:"xmlns,attr"  json:",omitempty"`
	MwsComGetCompetitivePricingForASINResult []*MwsComGetCompetitivePricingForASINResult `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 GetCompetitivePricingForASINResult,omitempty" json:"GetCompetitivePricingForASINResult,omitempty"`
	MwsComResponseMetadata *MwsComResponseMetadata `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ResponseMetadata,omitempty" json:"ResponseMetadata,omitempty"`
}

type MwsComGetCompetitivePricingForASINResult struct {
	XMLName xml.Name `xml:"GetCompetitivePricingForASINResult,omitempty" json:"GetCompetitivePricingForASINResult,omitempty"`
	AttrASIN string`xml:"ASIN,attr"  json:",omitempty"`
	Attrstatus string`xml:"status,attr"  json:",omitempty"`
	MwsComProduct *MwsComProduct `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Product,omitempty" json:"Product,omitempty"`
}

type MwsComIdentifiers struct {
	XMLName xml.Name `xml:"Identifiers,omitempty" json:"Identifiers,omitempty"`
	MwsComMarketplaceASIN *MwsComMarketplaceASIN `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 MarketplaceASIN,omitempty" json:"MarketplaceASIN,omitempty"`
}

type MwsComLandedPrice struct {
	XMLName xml.Name `xml:"LandedPrice,omitempty" json:"LandedPrice,omitempty"`
	MwsComAmount *MwsComAmount `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Amount,omitempty" json:"Amount,omitempty"`
	MwsComCurrencyCode *MwsComCurrencyCode `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 CurrencyCode,omitempty" json:"CurrencyCode,omitempty"`
}

type MwsComListingPrice struct {
	XMLName xml.Name `xml:"ListingPrice,omitempty" json:"ListingPrice,omitempty"`
	MwsComAmount *MwsComAmount `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Amount,omitempty" json:"Amount,omitempty"`
	MwsComCurrencyCode *MwsComCurrencyCode `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 CurrencyCode,omitempty" json:"CurrencyCode,omitempty"`
}

type MwsComMarketplaceASIN struct {
	XMLName xml.Name `xml:"MarketplaceASIN,omitempty" json:"MarketplaceASIN,omitempty"`
	MwsComASIN *MwsComASIN `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ASIN,omitempty" json:"ASIN,omitempty"`
	MwsComMarketplaceId *MwsComMarketplaceId `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 MarketplaceId,omitempty" json:"MarketplaceId,omitempty"`
}

type MwsComMarketplaceId struct {
	XMLName xml.Name `xml:"MarketplaceId,omitempty" json:"MarketplaceId,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsComNumberOfOfferListings struct {
	XMLName xml.Name `xml:"NumberOfOfferListings,omitempty" json:"NumberOfOfferListings,omitempty"`
	MwsComOfferListingCount []*MwsComOfferListingCount `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 OfferListingCount,omitempty" json:"OfferListingCount,omitempty"`
}

type MwsComOfferListingCount struct {
	XMLName xml.Name `xml:"OfferListingCount,omitempty" json:"OfferListingCount,omitempty"`
	Attrcondition string`xml:"condition,attr"  json:",omitempty"`
	Number int8 `xml:",chardata" json:",omitempty"`
}

type MwsComPrice struct {
	XMLName xml.Name `xml:"Price,omitempty" json:"Price,omitempty"`
	MwsComLandedPrice *MwsComLandedPrice `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 LandedPrice,omitempty" json:"LandedPrice,omitempty"`
	MwsComListingPrice *MwsComListingPrice `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ListingPrice,omitempty" json:"ListingPrice,omitempty"`
	MwsComShipping *MwsComShipping `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Shipping,omitempty" json:"Shipping,omitempty"`
}

type MwsComProduct struct {
	XMLName xml.Name `xml:"Product,omitempty" json:"Product,omitempty"`
	AttrXmlnsns2 string`xml:"xmlns ns2,attr"  json:",omitempty"`
	MwsComCompetitivePricing *MwsComCompetitivePricing `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 CompetitivePricing,omitempty" json:"CompetitivePricing,omitempty"`
	MwsComIdentifiers *MwsComIdentifiers `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Identifiers,omitempty" json:"Identifiers,omitempty"`
	MwsComSalesRankings *MwsComSalesRankings `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 SalesRankings,omitempty" json:"SalesRankings,omitempty"`
}

type MwsComProductCategoryId struct {
	XMLName xml.Name `xml:"ProductCategoryId,omitempty" json:"ProductCategoryId,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsComRank struct {
	XMLName xml.Name `xml:"Rank,omitempty" json:"Rank,omitempty"`
	Number int32 `xml:",chardata" json:",omitempty"`
}

type MwsComRequestId struct {
	XMLName xml.Name `xml:"RequestId,omitempty" json:"RequestId,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsComResponseMetadata struct {
	XMLName xml.Name `xml:"ResponseMetadata,omitempty" json:"ResponseMetadata,omitempty"`
	MwsComRequestId *MwsComRequestId `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 RequestId,omitempty" json:"RequestId,omitempty"`
}

type MwsComSalesRank struct {
	XMLName xml.Name `xml:"SalesRank,omitempty" json:"SalesRank,omitempty"`
	MwsComProductCategoryId *MwsComProductCategoryId `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ProductCategoryId,omitempty" json:"ProductCategoryId,omitempty"`
	MwsComRank *MwsComRank `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Rank,omitempty" json:"Rank,omitempty"`
}

type MwsComSalesRankings struct {
	XMLName xml.Name `xml:"SalesRankings,omitempty" json:"SalesRankings,omitempty"`
	MwsComSalesRank []*MwsComSalesRank `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 SalesRank,omitempty" json:"SalesRank,omitempty"`
}

type MwsComShipping struct {
	XMLName xml.Name `xml:"Shipping,omitempty" json:"Shipping,omitempty"`
	MwsComAmount *MwsComAmount `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Amount,omitempty" json:"Amount,omitempty"`
	MwsComCurrencyCode *MwsComCurrencyCode `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 CurrencyCode,omitempty" json:"CurrencyCode,omitempty"`
}

