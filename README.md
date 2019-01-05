# Amazon Marketplace Web Services (MWS) API [![Build Status](https://travis-ci.org/ezkl/go-amazon-mws-api.png?branch=master)](https://travis-ci.org/ezkl/go-amazon-mws-api)

This Amazon MWS API client is based heavily on @DDRBoxman's [go-amazon-product-api](https://github.com/DDRBoxman/go-amazon-product-api).

## Documentation

You can view the auto-generated documentation at [http://godoc.org/github.com/ezkl/go-amazon-mws-api](http://godoc.org/github.com/ezkl/go-amazon-mws-api).

## Usage

```go
package main

import (
       "fmt"
       "github.com/ezkl/go-amazon-mws-api"
)

func main() {
       var api amazonmws.AmazonMWSAPI

       api.AccessKey = ""
       api.SecretKey = ""
       api.Host = "mws.amazonservices.com"
       api.MarketplaceId = "ATVPDKIKX0DER"
       api.SellerId = ""

       asins := []string{"0195019199"}

       result,err := api.GetLowestOfferListingsForASIN(asins)

       if (err != nil) {
           fmt.Println(err)
       }

       fmt.Println(result)
}
```

## Create entity struct

1. save amazon mws api xml  
  
2. create struct using chidley  
https://github.com/gnewton/chidley  
  
```
$ go get -u github.com/gnewton/chidley
```
  
3. create struct from xml
```
$ sh xml_to_struct.sh get_matching_product_for_id Mat
```
