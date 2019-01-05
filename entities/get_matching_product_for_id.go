package entity

import "encoding/xml"

type MwsMatAmount__ns2 struct {
	XMLName xml.Name `xml:"Amount,omitempty" json:"Amount,omitempty"`
	Number float32 `xml:",chardata" json:",omitempty"`
}

type MwsMatAspectRatio__ns2 struct {
	XMLName xml.Name `xml:"AspectRatio,omitempty" json:"AspectRatio,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatBinding__ns2 struct {
	XMLName xml.Name `xml:"Binding,omitempty" json:"Binding,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatBlurayRegion__ns2 struct {
	XMLName xml.Name `xml:"BlurayRegion,omitempty" json:"BlurayRegion,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatBrand__ns2 struct {
	XMLName xml.Name `xml:"Brand,omitempty" json:"Brand,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatCreator__ns2 struct {
	XMLName xml.Name `xml:"Creator,omitempty" json:"Creator,omitempty"`
	AttrRole string`xml:"Role,attr"  json:",omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatCurrencyCode__ns2 struct {
	XMLName xml.Name `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatFormat__ns2 struct {
	XMLName xml.Name `xml:"Format,omitempty" json:"Format,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatHeight__ns2 struct {
	XMLName xml.Name `xml:"Height,omitempty" json:"Height,omitempty"`
	AttrUnits string`xml:"Units,attr"  json:",omitempty"`
	Number float32 `xml:",chardata" json:",omitempty"`
}

type MwsMatIsAdultProduct__ns2 struct {
	XMLName xml.Name `xml:"IsAdultProduct,omitempty" json:"IsAdultProduct,omitempty"`
	Flag bool `xml:",chardata" json:",omitempty"`
}

type MwsMatItemAttributes__ns2 struct {
	XMLName xml.Name `xml:"ItemAttributes,omitempty" json:"ItemAttributes,omitempty"`
	AttrHttp_colon__slash__slash_www_dot_w3_dot_org_slash_XML_slash_1998_slash_namespacelang string`xml:"http://www.w3.org/XML/1998/namespace lang,attr"  json:",omitempty"`
	MwsMatAspectRatio__ns2 *MwsMatAspectRatio__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd AspectRatio,omitempty" json:"AspectRatio,omitempty"`
	MwsMatBinding__ns2 *MwsMatBinding__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Binding,omitempty" json:"Binding,omitempty"`
	MwsMatBlurayRegion__ns2 *MwsMatBlurayRegion__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd BlurayRegion,omitempty" json:"BlurayRegion,omitempty"`
	MwsMatBrand__ns2 *MwsMatBrand__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Brand,omitempty" json:"Brand,omitempty"`
	MwsMatCreator__ns2 []*MwsMatCreator__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Creator,omitempty" json:"Creator,omitempty"`
	MwsMatFormat__ns2 []*MwsMatFormat__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Format,omitempty" json:"Format,omitempty"`
	MwsMatIsAdultProduct__ns2 *MwsMatIsAdultProduct__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd IsAdultProduct,omitempty" json:"IsAdultProduct,omitempty"`
	MwsMatItemDimensions__ns2 *MwsMatItemDimensions__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd ItemDimensions,omitempty" json:"ItemDimensions,omitempty"`
	MwsMatLabel__ns2 *MwsMatLabel__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Label,omitempty" json:"Label,omitempty"`
	MwsMatLanguages__ns2 *MwsMatLanguages__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Languages,omitempty" json:"Languages,omitempty"`
	MwsMatListPrice__ns2 *MwsMatListPrice__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd ListPrice,omitempty" json:"ListPrice,omitempty"`
	MwsMatManufacturer__ns2 *MwsMatManufacturer__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Manufacturer,omitempty" json:"Manufacturer,omitempty"`
	MwsMatNumberOfDiscs__ns2 *MwsMatNumberOfDiscs__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd NumberOfDiscs,omitempty" json:"NumberOfDiscs,omitempty"`
	MwsMatNumberOfItems__ns2 *MwsMatNumberOfItems__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd NumberOfItems,omitempty" json:"NumberOfItems,omitempty"`
	MwsMatPackageDimensions__ns2 *MwsMatPackageDimensions__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd PackageDimensions,omitempty" json:"PackageDimensions,omitempty"`
	MwsMatProductGroup__ns2 *MwsMatProductGroup__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd ProductGroup,omitempty" json:"ProductGroup,omitempty"`
	MwsMatProductTypeName__ns2 *MwsMatProductTypeName__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd ProductTypeName,omitempty" json:"ProductTypeName,omitempty"`
	MwsMatPublisher__ns2 *MwsMatPublisher__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Publisher,omitempty" json:"Publisher,omitempty"`
	MwsMatRegionCode__ns2 *MwsMatRegionCode__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd RegionCode,omitempty" json:"RegionCode,omitempty"`
	MwsMatReleaseDate__ns2 *MwsMatReleaseDate__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd ReleaseDate,omitempty" json:"ReleaseDate,omitempty"`
	MwsMatRunningTime__ns2 *MwsMatRunningTime__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd RunningTime,omitempty" json:"RunningTime,omitempty"`
	MwsMatSeikodoProductCode__ns2 *MwsMatSeikodoProductCode__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd SeikodoProductCode,omitempty" json:"SeikodoProductCode,omitempty"`
	MwsMatSmallImage__ns2 *MwsMatSmallImage__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd SmallImage,omitempty" json:"SmallImage,omitempty"`
	MwsMatStudio__ns2 *MwsMatStudio__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Studio,omitempty" json:"Studio,omitempty"`
	MwsMatTitle__ns2 *MwsMatTitle__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Title,omitempty" json:"Title,omitempty"`
}

type MwsMatItemDimensions__ns2 struct {
	XMLName xml.Name `xml:"ItemDimensions,omitempty" json:"ItemDimensions,omitempty"`
	MwsMatHeight__ns2 *MwsMatHeight__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Height,omitempty" json:"Height,omitempty"`
	MwsMatLength__ns2 *MwsMatLength__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Length,omitempty" json:"Length,omitempty"`
	MwsMatWeight__ns2 *MwsMatWeight__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Weight,omitempty" json:"Weight,omitempty"`
	MwsMatWidth__ns2 *MwsMatWidth__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Width,omitempty" json:"Width,omitempty"`
}

type MwsMatLabel__ns2 struct {
	XMLName xml.Name `xml:"Label,omitempty" json:"Label,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatLanguage__ns2 struct {
	XMLName xml.Name `xml:"Language,omitempty" json:"Language,omitempty"`
	MwsMatName__ns2 *MwsMatName__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Name,omitempty" json:"Name,omitempty"`
	MwsMatType__ns2 *MwsMatType__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Type,omitempty" json:"Type,omitempty"`
}

type MwsMatLanguages__ns2 struct {
	XMLName xml.Name `xml:"Languages,omitempty" json:"Languages,omitempty"`
	MwsMatLanguage__ns2 []*MwsMatLanguage__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Language,omitempty" json:"Language,omitempty"`
}

type MwsMatLength__ns2 struct {
	XMLName xml.Name `xml:"Length,omitempty" json:"Length,omitempty"`
	AttrUnits string`xml:"Units,attr"  json:",omitempty"`
	Number float32 `xml:",chardata" json:",omitempty"`
}

type MwsMatListPrice__ns2 struct {
	XMLName xml.Name `xml:"ListPrice,omitempty" json:"ListPrice,omitempty"`
	MwsMatAmount__ns2 *MwsMatAmount__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Amount,omitempty" json:"Amount,omitempty"`
	MwsMatCurrencyCode__ns2 *MwsMatCurrencyCode__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd CurrencyCode,omitempty" json:"CurrencyCode,omitempty"`
}

type MwsMatManufacturer__ns2 struct {
	XMLName xml.Name `xml:"Manufacturer,omitempty" json:"Manufacturer,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatName__ns2 struct {
	XMLName xml.Name `xml:"Name,omitempty" json:"Name,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatNumberOfDiscs__ns2 struct {
	XMLName xml.Name `xml:"NumberOfDiscs,omitempty" json:"NumberOfDiscs,omitempty"`
	Number int8 `xml:",chardata" json:",omitempty"`
}

type MwsMatNumberOfItems__ns2 struct {
	XMLName xml.Name `xml:"NumberOfItems,omitempty" json:"NumberOfItems,omitempty"`
	Number int8 `xml:",chardata" json:",omitempty"`
}

type MwsMatPackageDimensions__ns2 struct {
	XMLName xml.Name `xml:"PackageDimensions,omitempty" json:"PackageDimensions,omitempty"`
	MwsMatHeight__ns2 *MwsMatHeight__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Height,omitempty" json:"Height,omitempty"`
	MwsMatLength__ns2 *MwsMatLength__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Length,omitempty" json:"Length,omitempty"`
	MwsMatWeight__ns2 *MwsMatWeight__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Weight,omitempty" json:"Weight,omitempty"`
	MwsMatWidth__ns2 *MwsMatWidth__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Width,omitempty" json:"Width,omitempty"`
}

type MwsMatProductGroup__ns2 struct {
	XMLName xml.Name `xml:"ProductGroup,omitempty" json:"ProductGroup,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatProductTypeName__ns2 struct {
	XMLName xml.Name `xml:"ProductTypeName,omitempty" json:"ProductTypeName,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatPublisher__ns2 struct {
	XMLName xml.Name `xml:"Publisher,omitempty" json:"Publisher,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatRegionCode__ns2 struct {
	XMLName xml.Name `xml:"RegionCode,omitempty" json:"RegionCode,omitempty"`
	Number int8 `xml:",chardata" json:",omitempty"`
}

type MwsMatReleaseDate__ns2 struct {
	XMLName xml.Name `xml:"ReleaseDate,omitempty" json:"ReleaseDate,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatRunningTime__ns2 struct {
	XMLName xml.Name `xml:"RunningTime,omitempty" json:"RunningTime,omitempty"`
	AttrUnits string`xml:"Units,attr"  json:",omitempty"`
	Number int16 `xml:",chardata" json:",omitempty"`
}

type MwsMatSeikodoProductCode__ns2 struct {
	XMLName xml.Name `xml:"SeikodoProductCode,omitempty" json:"SeikodoProductCode,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatSmallImage__ns2 struct {
	XMLName xml.Name `xml:"SmallImage,omitempty" json:"SmallImage,omitempty"`
	MwsMatHeight__ns2 *MwsMatHeight__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Height,omitempty" json:"Height,omitempty"`
	MwsMatURL__ns2 *MwsMatURL__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd URL,omitempty" json:"URL,omitempty"`
	MwsMatWidth__ns2 *MwsMatWidth__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd Width,omitempty" json:"Width,omitempty"`
}

type MwsMatStudio__ns2 struct {
	XMLName xml.Name `xml:"Studio,omitempty" json:"Studio,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatTitle__ns2 struct {
	XMLName xml.Name `xml:"Title,omitempty" json:"Title,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatType__ns2 struct {
	XMLName xml.Name `xml:"Type,omitempty" json:"Type,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatURL__ns2 struct {
	XMLName xml.Name `xml:"URL,omitempty" json:"URL,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatWeight__ns2 struct {
	XMLName xml.Name `xml:"Weight,omitempty" json:"Weight,omitempty"`
	AttrUnits string`xml:"Units,attr"  json:",omitempty"`
	Number float32 `xml:",chardata" json:",omitempty"`
}

type MwsMatWidth__ns2 struct {
	XMLName xml.Name `xml:"Width,omitempty" json:"Width,omitempty"`
	AttrUnits string`xml:"Units,attr"  json:",omitempty"`
	Number float32 `xml:",chardata" json:",omitempty"`
}

type MwsMatASIN struct {
	XMLName xml.Name `xml:"ASIN,omitempty" json:"ASIN,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatAttributeSets struct {
	XMLName xml.Name `xml:"AttributeSets,omitempty" json:"AttributeSets,omitempty"`
	MwsMatItemAttributes__ns2 *MwsMatItemAttributes__ns2 `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01/default.xsd ItemAttributes,omitempty" json:"ItemAttributes,omitempty"`
}

type MwsMatGetMatchingProductForIdResponse struct {
	XMLName xml.Name `xml:"GetMatchingProductForIdResponse,omitempty" json:"GetMatchingProductForIdResponse,omitempty"`
	Attrxmlns string`xml:"xmlns,attr"  json:",omitempty"`
	MwsMatGetMatchingProductForIdResult []*MwsMatGetMatchingProductForIdResult `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 GetMatchingProductForIdResult,omitempty" json:"GetMatchingProductForIdResult,omitempty"`
	MwsMatResponseMetadata *MwsMatResponseMetadata `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ResponseMetadata,omitempty" json:"ResponseMetadata,omitempty"`
}

type MwsMatGetMatchingProductForIdResult struct {
	XMLName xml.Name `xml:"GetMatchingProductForIdResult,omitempty" json:"GetMatchingProductForIdResult,omitempty"`
	AttrId string`xml:"Id,attr"  json:",omitempty"`
	AttrIdType string`xml:"IdType,attr"  json:",omitempty"`
	Attrstatus string`xml:"status,attr"  json:",omitempty"`
	MwsMatProducts *MwsMatProducts `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Products,omitempty" json:"Products,omitempty"`
}

type MwsMatIdentifiers struct {
	XMLName xml.Name `xml:"Identifiers,omitempty" json:"Identifiers,omitempty"`
	MwsMatMarketplaceASIN *MwsMatMarketplaceASIN `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 MarketplaceASIN,omitempty" json:"MarketplaceASIN,omitempty"`
}

type MwsMatMarketplaceASIN struct {
	XMLName xml.Name `xml:"MarketplaceASIN,omitempty" json:"MarketplaceASIN,omitempty"`
	MwsMatASIN *MwsMatASIN `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ASIN,omitempty" json:"ASIN,omitempty"`
	MwsMatMarketplaceId *MwsMatMarketplaceId `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 MarketplaceId,omitempty" json:"MarketplaceId,omitempty"`
}

type MwsMatMarketplaceId struct {
	XMLName xml.Name `xml:"MarketplaceId,omitempty" json:"MarketplaceId,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatProduct struct {
	XMLName xml.Name `xml:"Product,omitempty" json:"Product,omitempty"`
	MwsMatAttributeSets *MwsMatAttributeSets `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 AttributeSets,omitempty" json:"AttributeSets,omitempty"`
	MwsMatIdentifiers *MwsMatIdentifiers `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Identifiers,omitempty" json:"Identifiers,omitempty"`
	MwsMatRelationships *MwsMatRelationships `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Relationships,omitempty" json:"Relationships,omitempty"`
	MwsMatSalesRankings *MwsMatSalesRankings `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 SalesRankings,omitempty" json:"SalesRankings,omitempty"`
}

type MwsMatProductCategoryId struct {
	XMLName xml.Name `xml:"ProductCategoryId,omitempty" json:"ProductCategoryId,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatProducts struct {
	XMLName xml.Name `xml:"Products,omitempty" json:"Products,omitempty"`
	AttrXmlnsns2 string`xml:"xmlns ns2,attr"  json:",omitempty"`
	MwsMatProduct *MwsMatProduct `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Product,omitempty" json:"Product,omitempty"`
}

type MwsMatRank struct {
	XMLName xml.Name `xml:"Rank,omitempty" json:"Rank,omitempty"`
	Number int32 `xml:",chardata" json:",omitempty"`
}

type MwsMatRelationships struct {
	XMLName xml.Name `xml:"Relationships,omitempty" json:"Relationships,omitempty"`
}

type MwsMatRequestId struct {
	XMLName xml.Name `xml:"RequestId,omitempty" json:"RequestId,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type MwsMatResponseMetadata struct {
	XMLName xml.Name `xml:"ResponseMetadata,omitempty" json:"ResponseMetadata,omitempty"`
	MwsMatRequestId *MwsMatRequestId `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 RequestId,omitempty" json:"RequestId,omitempty"`
}

type MwsMatSalesRank struct {
	XMLName xml.Name `xml:"SalesRank,omitempty" json:"SalesRank,omitempty"`
	MwsMatProductCategoryId *MwsMatProductCategoryId `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 ProductCategoryId,omitempty" json:"ProductCategoryId,omitempty"`
	MwsMatRank *MwsMatRank `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 Rank,omitempty" json:"Rank,omitempty"`
}

type MwsMatSalesRankings struct {
	XMLName xml.Name `xml:"SalesRankings,omitempty" json:"SalesRankings,omitempty"`
	MwsMatSalesRank []*MwsMatSalesRank `xml:"http://mws.amazonservices.com/schema/Products/2011-10-01 SalesRank,omitempty" json:"SalesRank,omitempty"`
}

