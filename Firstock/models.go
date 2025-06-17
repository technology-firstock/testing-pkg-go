// Copyright (c) [2025] [abc]
// SPDX-License-Identifier: MIT
package Firstock

// Models for Login
type LoginRequest struct {
	UserId     string `json:"userId"`
	Password   string `json:"password"`
	TOTP       string `json:"totp"`
	VendorCode string `json:"vendorCode"`
	APIKey     string `json:"apiKey"`
}

type LoginResponse struct {
	Status  string `json:"status"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type Data struct {
	Actid      string `json:"actid"`
	UserName   string `json:"userName"`
	SUserToken string `json:"susertoken"`
	Email      string `json:"email"`
}

// Model for User Details
type UserDetailsRequest struct {
	UserId string `json:"userId"`
	JKey   string `json:"jkey"`
}

// Model for Logout
type LogoutRequest struct {
	UserId string `json:"userId"`
	JKey   string `json:"jkey"`
}

// Models for Place Order
type PlaceOrderRequest struct {
	UserId          string `json:"userId"`
	Exchange        string `json:"exchange"`
	Retention       string `json:"retention"`
	Product         string `json:"product"`
	PriceType       string `json:"priceType"`
	TradingSymbol   string `json:"tradingSymbol"`
	TransactionType string `json:"transactionType"`
	Price           string `json:"price"`
	TriggerPrice    string `json:"triggerPrice"`
	Quantity        string `json:"quantity"`
	Remarks         string `json:"remarks"`
}

type PlaceOrderRequestBody struct {
	UserId          string `json:"userId"`
	JKey            string `json:"jkey"`
	Exchange        string `json:"exchange"`
	Retention       string `json:"retention"`
	Product         string `json:"product"`
	PriceType       string `json:"priceType"`
	TradingSymbol   string `json:"tradingSymbol"`
	TransactionType string `json:"transactionType"`
	Price           string `json:"price"`
	TriggerPrice    string `json:"triggerPrice"`
	Quantity        string `json:"quantity"`
	Remarks         string `json:"remarks"`
}

// Models for OrderMargin
type OrderMarginRequest struct {
	UserId          string `json:"userId"`
	Exchange        string `json:"exchange"`
	TransactionType string `json:"transactionType"`
	Product         string `json:"product"`
	TradingSymbol   string `json:"tradingSymbol"`
	Quantity        string `json:"quantity"`
	PriceType       string `json:"priceType"`
	Price           string `json:"price"`
}

type OrderMarginRequestBody struct {
	UserId          string `json:"userId"`
	JKey            string `json:"jkey"`
	Exchange        string `json:"exchange"`
	TransactionType string `json:"transactionType"`
	Product         string `json:"product"`
	TradingSymbol   string `json:"tradingSymbol"`
	Quantity        string `json:"quantity"`
	PriceType       string `json:"priceType"`
	Price           string `json:"price"`
}

// Models for Single Order History

type OrderRequest struct {
	UserId      string `json:"userId"`
	OrderNumber string `json:"orderNumber"`
}

type OrderRequestBody struct {
	UserId      string `json:"userId"`
	JKey        string `json:"jkey"`
	OrderNumber string `json:"orderNumber"`
}

// Model for Trade Book, RmsLimit
type BaseRequest struct {
	UserId string `json:"userId"`
	JKey   string `json:"jkey"`
}

// Models for Get Expiry

type GetInfoRequest struct {
	UserId        string `json:"userId"`
	Exchange      string `json:"exchange"`
	TradingSymbol string `json:"tradingSymbol"`
}

type GetInfoRequestBody struct {
	UserId        string `json:"userId"`
	JKey          string `json:"jkey"`
	Exchange      string `json:"exchange"`
	TradingSymbol string `json:"tradingSymbol"`
}

// Model for ModifyOrder
type ModifyOrderRequest struct {
	UserId         string `json:"userId"`
	OrderNumber    string `json:"orderNumber"`
	PriceType      string `json:"priceType"`
	TradingSymbol  string `json:"tradingSymbol"`
	Price          string `json:"price"`
	TriggerPrice   string `json:"triggerPrice"`
	Quantity       string `json:"quantity"`
	Product        string `json:"product"`
	Retention      string `json:"retention"`
	Mkt_protection string `json:"mkt_protection"`
}

type ModifyOrderRequestBody struct {
	UserId         string `json:"userId"`
	JKey           string `json:"jkey"`
	OrderNumber    string `json:"orderNumber"`
	PriceType      string `json:"priceType"`
	TradingSymbol  string `json:"tradingSymbol"`
	Price          string `json:"price"`
	TriggerPrice   string `json:"triggerPrice"`
	Quantity       string `json:"quantity"`
	Product        string `json:"product"`
	Retention      string `json:"retention"`
	Mkt_protection string `json:"mkt_protection"`
}

type BrokerageCalculatorRequest struct {
	UserId          string `json:"userId"`
	Exchange        string `json:"exchange"`
	TradingSymbol   string `json:"tradingSymbol"`
	TransactionType string `json:"transactionType"`
	Product         string `json:"Product"`
	Quantity        string `json:"quantity"`
	Price           string `json:"price"`
	StrikePrice     string `json:"strike_price"`
	InstName        string `json:"inst_name"`
	LotSize         string `json:"lot_size"`
}

type BrokerageCalculatorRequestBody struct {
	UserId          string `json:"userId"`
	JKey            string `json:"jkey"`
	Exchange        string `json:"exchange"`
	TradingSymbol   string `json:"tradingSymbol"`
	TransactionType string `json:"transactionType"`
	Product         string `json:"Product"`
	Quantity        string `json:"quantity"`
	Price           string `json:"price"`
	StrikePrice     string `json:"strike_price"`
	InstName        string `json:"inst_name"`
	LotSize         string `json:"lot_size"`
}

type BasketListParam struct {
	Exchange        string `json:"exchange"`
	TransactionType string `json:"transactionType"`
	Product         string `json:"product"`
	TradingSymbol   string `json:"tradingSymbol"`
	Quantity        string `json:"quantity"`
	PriceType       string `json:"priceType"`
	Price           string `json:"price"`
}

type BasketMarginRequest struct {
	UserId           string            `json:"userId"`
	Exchange         string            `json:"exchange"`
	TransactionType  string            `json:"transactionType"`
	Product          string            `json:"product"`
	TradingSymbol    string            `json:"tradingSymbol"`
	Quantity         string            `json:"quantity"`
	PriceType        string            `json:"priceType"`
	Price            string            `json:"price"`
	BasketListParams []BasketListParam `json:"BasketList_Params"`
}

type BasketMarginRequestBody struct {
	UserId           string            `json:"userId"`
	JKey             string            `json:"jKey"`
	Exchange         string            `json:"exchange"`
	TransactionType  string            `json:"transactionType"`
	Product          string            `json:"product"`
	TradingSymbol    string            `json:"tradingSymbol"`
	Quantity         string            `json:"quantity"`
	PriceType        string            `json:"priceType"`
	Price            string            `json:"price"`
	BasketListParams []BasketListParam `json:"BasketList_Params"`
}

type ProductConversionRequest struct {
	UserId          string `json:"userId"`
	TradingSymbol   string `json:"tradingSymbol"`
	Exchange        string `json:"exchange"`
	PreviousProduct string `json:"previousProduct"`
	Product         string `json:"product"`
	Quantity        string `json:"quantity"`
}

type ProductConversionRequestBody struct {
	UserId          string `json:"userId"`
	JKey            string `json:"jkey"`
	TradingSymbol   string `json:"tradingSymbol"`
	Exchange        string `json:"exchange"`
	PreviousProduct string `json:"previousProduct"`
	Product         string `json:"product"`
	Quantity        string `json:"quantity"`
}

type SearchScripsRequest struct {
	UserId string `json:"userId"`
	SText  string `json:"stext"`
}

type SearchScripsBody struct {
	UserId string `json:"userId"`
	JKey   string `json:"jkey"`
	SText  string `json:"stext"`
}

type OptionChainRequest struct {
	UserId      string `json:"userId"`
	Exchange    string `json:"exchange"`
	Symbol      string `json:"symbol"`
	Expiry      string `json:"expiry"`
	Count       string `json:"count"`
	StrikePrice string `json:"strikePrice"`
}

type OptionChainRequestBody struct {
	UserId      string `json:"userId"`
	JKey        string `json:"jkey"`
	Exchange    string `json:"exchange"`
	Symbol      string `json:"symbol"`
	Expiry      string `json:"expiry"`
	Count       string `json:"count"`
	StrikePrice string `json:"strikePrice"`
}

type MultiQuoteData struct {
	Exchange      string `json:"exchange"`
	TradingSymbol string `json:"tradingSymbol"`
}

type GetMultiQuotesRequest struct {
	UserId string           `json:"userId"`
	Data   []MultiQuoteData `json:"data"`
}

type GetMultiQuotesRequestBody struct {
	UserId string           `json:"userId"`
	JKey   string           `json:"jkey"`
	Data   []MultiQuoteData `json:"data"`
}

type TimePriceSeriesIntervalRequest struct {
	UserId        string `json:"userId"`
	Exchange      string `json:"exchange"`
	Interval      string `json:"interval"`
	TradingSymbol string `json:"tradingSymbol"`
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
}

type TimePriceSeriesIntervalRequestBody struct {
	UserId        string `json:"userId"`
	JKey          string `json:"jkey"`
	Exchange      string `json:"exchange"`
	Interval      string `json:"interval"`
	TradingSymbol string `json:"tradingSymbol"`
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
}
