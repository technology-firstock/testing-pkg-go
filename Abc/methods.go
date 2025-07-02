// Copyright (c) [2025] [abc]
// SPDX-License-Identifier: MIT
package Abc

import (
	"encoding/json"
)

type firstock struct{}

var thefirstock = &apifunctions{}

// Call Login function to login to Firstock
// It takes a LoginRequest struct as input and returns a JSON response string and an error if any.
func (fs *firstock) Login(reqBody LoginRequest) (loginResponse string) {
	var login map[string]interface{}
	var status string = status_failed
	var result string

	var loginRequest LoginRequest = LoginRequest{
		UserId:     reqBody.UserId,
		Password:   encodePassword(reqBody.Password),
		TOTP:       reqBody.TOTP,
		VendorCode: reqBody.VendorCode,
		APIKey:     reqBody.APIKey,
	}
	login, err := thefirstock.LoginFunction(
		loginRequest,
	)
	if err != nil {
		loginResponse = internalServerErrorResponse()
		return
	}
	if login == nil {
		loginResponse = internalServerErrorResponse()
		return
	}

	s, ok := login[status_val].(string)
	if ok {
		status = s
	}
	loginStr, err := json.Marshal(login)
	if err != nil {
		loginResponse = internalServerErrorResponse()
		return
	}
	result = string(loginStr)
	if status != status_success {
		loginResponse = failureResponseStructure(result)
		return
	}
	// Extract SUserToken from login response
	dataMap, ok := login[data].(map[string]interface{})
	if !ok {
		loginResponse = internalServerErrorResponse() // "login[\"data\"] is not a map[string]interface{}"
		return
	}

	sUserToken, ok := dataMap[susertoken].(string)
	if !ok {
		loginResponse = internalServerErrorResponse() // "failed to extract SUserToken from login response"
		return
	}

	// Write the following to a config.json file. Create the file if it does not exist.
	err = saveJKeyToConfig(LogoutRequest{
		UserId: reqBody.UserId,
		JKey:   sUserToken,
	})

	loginResponse = successResponseStructure(result)

	return
}

// Call Logout function to logout from Firstock
// It takes a userId as input and returns a JSON response string and an error if any.
func (fs *firstock) Logout(userId string) (logoutResponse string) {
	var logout LogoutRequest
	var status string = status_failed
	var result string
	logout.UserId = userId
	logout.JKey = ""

	// Read jKey for userId from config.json
	jkey, errRead := readJKeyFromConfig(userId)
	if errRead != nil {
		logoutResponse = pleaseLoginToFirstock()
		return
	}

	logout.JKey = jkey
	logoutInfo, errLogout := thefirstock.LogoutFunction(logout)
	if errLogout != nil {
		logoutResponse = internalServerErrorResponse()
		return
	}
	if logoutInfo == nil {
		logoutResponse = internalServerErrorResponse()
		return
	}
	s, ok := logoutInfo[status_val].(string)
	if ok {
		status = s
	}

	logoutStr, errMarshal := json.Marshal(logoutInfo)

	if errMarshal != nil {
		logoutResponse = internalServerErrorResponse()
		return
	}

	result = string(logoutStr)
	if status != status_success {
		logoutResponse = failureResponseStructure(result)
		return
	}
	// Remove userId from config.json
	removeJKeyFromConfig(logout.UserId)
	logoutResponse = successResponseStructure(result)
	return
}

// Call UserDetails function to fetch user details from Firstock
// It takes a userId as input and returns a JSON response string and an error if any.
func (fs *firstock) UserDetails(userId string) (userDetailsResponse string) {
	var userDetailsRequest UserDetailsRequest
	var status string = status_failed
	var result string
	// Read jKey for userId from config.json
	jkey, errRead := readJKeyFromConfig(userId)
	if errRead != nil {
		userDetailsResponse = pleaseLoginToFirstock()
		return
	}
	userDetailsRequest.JKey = jkey
	userDetailsRequest.UserId = userId

	userDetails, errRes := thefirstock.UserDetailsFunction(userDetailsRequest)
	if errRes != nil {
		userDetailsResponse = internalServerErrorResponse()
		return
	}
	if userDetails == nil {
		userDetailsResponse = internalServerErrorResponse()
		return
	}
	s, ok := userDetails[status_val].(string)
	if ok {
		status = s
	}
	userDetailsStr, err := json.Marshal(userDetails)
	if err != nil {
		userDetailsResponse = internalServerErrorResponse()
		return
	}

	result = string(userDetailsStr)
	if status != status_success {
		userDetailsResponse = failureResponseStructure(result)
		return
	}

	userDetailsResponse = successResponseStructure(result)
	return
}

func (fs *firstock) PlaceOrder(req PlaceOrderRequest) (placeOrderResponse string) {
	var status string = status_failed
	var result string
	jkey, errRead := readJKeyFromConfig(req.UserId)
	if errRead != nil {
		placeOrderResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := PlaceOrderRequestBody{
		UserId:          req.UserId,
		JKey:            jkey,
		Exchange:        req.Exchange,
		Retention:       req.Retention,
		Product:         req.Product,
		PriceType:       req.PriceType,
		TradingSymbol:   req.TradingSymbol,
		TransactionType: req.TransactionType,
		Price:           req.Price,
		TriggerPrice:    req.TriggerPrice,
		Quantity:        req.Quantity,
		Remarks:         req.Remarks,
	}

	placeOrderDetails, errPlaceOrder := thefirstock.PlaceOrderFunction(reqBody)
	if errPlaceOrder != nil {
		placeOrderResponse = internalServerErrorResponse()
		return
	}
	if placeOrderDetails == nil {
		placeOrderResponse = internalServerErrorResponse()
		return
	}

	s, ok := placeOrderDetails[status_val].(string)
	if ok {
		status = s
	}

	placeOrderDetailsVal, err := json.Marshal(placeOrderDetails)
	if err != nil {
		placeOrderResponse = internalServerErrorResponse()
		return
	}

	result = string(placeOrderDetailsVal)
	if status != status_success {
		placeOrderResponse = failureResponseStructure(result)
		return
	}
	placeOrderResponse = successResponseStructure(result)

	return
}

func (fs *firstock) OrderMargin(req OrderMarginRequest) (orderMarginResponse string) {
	var status string = status_failed
	var result string

	jkey, errRead := readJKeyFromConfig(req.UserId)
	if errRead != nil {
		orderMarginResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := OrderMarginRequestBody{
		UserId:          req.UserId,
		JKey:            jkey,
		Exchange:        req.Exchange,
		TransactionType: req.TransactionType,
		Product:         req.Product,
		TradingSymbol:   req.TradingSymbol,
		Quantity:        req.Quantity,
		PriceType:       req.PriceType,
		Price:           req.Price,
	}

	orderMarginDetails, errOrder := thefirstock.OrderMarginFunction(reqBody)
	if errOrder != nil {
		orderMarginResponse = internalServerErrorResponse()
		return
	}

	if orderMarginDetails == nil {
		orderMarginResponse = internalServerErrorResponse()
		return
	}

	s, ok := orderMarginDetails[status_val].(string)
	if ok {
		status = s
	}

	orderMarginDetailsVal, err := json.Marshal(orderMarginDetails)
	if err != nil {
		orderMarginResponse = internalServerErrorResponse()
		return
	}
	result = string(orderMarginDetailsVal)
	if status != status_success {
		orderMarginResponse = failureResponseStructure(result)
		return
	}

	orderMarginResponse = successResponseStructure(result)
	return
}

func (fs *firstock) SingleOrderHistory(req OrderRequest) (singleOrderHistoryResponse string) {
	var status string = status_failed
	var result string
	jkey, errRead := readJKeyFromConfig(req.UserId)
	if errRead != nil {
		singleOrderHistoryResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := OrderRequestBody{
		UserId:      req.UserId,
		JKey:        jkey,
		OrderNumber: req.OrderNumber,
	}

	singleOrderHistoryDetails, errOrder := thefirstock.SingleOrderHistoryFunction(reqBody)
	if errOrder != nil {
		singleOrderHistoryResponse = internalServerErrorResponse()
		return
	}
	if singleOrderHistoryDetails == nil {
		singleOrderHistoryResponse = internalServerErrorResponse()
		return
	}

	s, ok := singleOrderHistoryDetails[status_val].(string)
	if ok {
		status = s
	}

	singleOrderHistoryDetailsVal, err := json.Marshal(singleOrderHistoryDetails)
	if err != nil {
		singleOrderHistoryResponse = internalServerErrorResponse()
		return
	}
	result = string(singleOrderHistoryDetailsVal)

	if status != status_success {
		singleOrderHistoryResponse = failureResponseStructure(result)
		return
	}

	singleOrderHistoryResponse = successResponseStructure(result)

	return
}

func (fs *firstock) CancelOrder(req OrderRequest) (cancelOrderResponse string) {
	var status string = status_failed
	var result string
	jkey, errRead := readJKeyFromConfig(req.UserId)

	if errRead != nil {
		cancelOrderResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := OrderRequestBody{
		UserId:      req.UserId,
		JKey:        jkey,
		OrderNumber: req.OrderNumber,
	}

	cancelOrderDetails, errOrder := thefirstock.CancelOrderFunction(reqBody)
	if errOrder != nil {
		cancelOrderResponse = internalServerErrorResponse()
		return
	}
	if cancelOrderDetails == nil {
		cancelOrderResponse = internalServerErrorResponse()
		return
	}

	s, ok := cancelOrderDetails[status_val].(string)
	if ok {
		status = s
	}

	cancelOrderDetailsVal, err := json.Marshal(cancelOrderDetails)
	if err != nil {
		cancelOrderResponse = internalServerErrorResponse()
		return
	}
	result = string(cancelOrderDetailsVal)

	if status != status_success {
		cancelOrderResponse = failureResponseStructure(result)
		return
	}

	cancelOrderResponse = successResponseStructure(result)

	return
}

func (fs *firstock) ModifyOrder(req ModifyOrderRequest) (modifyOrderResponse string) {
	var status string = status_failed
	var result string
	jkey, errRead := readJKeyFromConfig(req.UserId)
	if errRead != nil {
		modifyOrderResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := ModifyOrderRequestBody{
		UserId:         req.UserId,
		JKey:           jkey,
		OrderNumber:    req.OrderNumber,
		PriceType:      req.PriceType,
		TradingSymbol:  req.TradingSymbol,
		Price:          req.Price,
		TriggerPrice:   req.TriggerPrice,
		Quantity:       req.Quantity,
		Product:        req.Product,
		Retention:      req.Retention,
		Mkt_protection: req.Mkt_protection,
	}

	modifyOrderDetails, errOrder := thefirstock.ModifyOrderFunction(reqBody)
	if errOrder != nil {
		modifyOrderResponse = internalServerErrorResponse()
		return
	}

	if modifyOrderDetails == nil {
		modifyOrderResponse = internalServerErrorResponse()
		return
	}

	s, ok := modifyOrderDetails[status_val].(string)
	if ok {
		status = s
	}

	modifyOrderDetailsVal, err := json.Marshal(modifyOrderDetails)
	if err != nil {
		modifyOrderResponse = internalServerErrorResponse()
		return
	}
	result = string(modifyOrderDetailsVal)

	if status != status_success {
		modifyOrderResponse = failureResponseStructure(result)
		return
	}

	modifyOrderResponse = successResponseStructure(result)

	return
}

func (fs *firstock) TradeBook(userId string) (tradeBookResponse string) {
	var status string = status_failed
	var result string

	jkey, errRead := readJKeyFromConfig(userId)
	if errRead != nil {
		tradeBookResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	tradeBookDetails, errTradeBook := thefirstock.TradeBookFunction(reqBody)
	if errTradeBook != nil {
		tradeBookResponse = internalServerErrorResponse()
		return
	}

	if tradeBookDetails == nil {
		tradeBookResponse = internalServerErrorResponse()
		return
	}

	s, ok := tradeBookDetails[status_val].(string)
	if ok {
		status = s
	}

	tradeBookDetailsVal, err := json.Marshal(tradeBookDetails)
	if err != nil {
		tradeBookResponse = internalServerErrorResponse()
		return
	}
	result = string(tradeBookDetailsVal)

	if status != status_success {
		tradeBookResponse = failureResponseStructure(result)
		return
	}

	tradeBookResponse = successResponseStructure(result)

	return
}

func (fs *firstock) RMSLmit(userId string) (rmsLmitResponse string) {
	var status string = status_failed
	var result string
	jkey, errRead := readJKeyFromConfig(userId)
	if errRead != nil {
		rmsLmitResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	rmsLimitDetails, errRmsLimit := thefirstock.RmsLimitFunction(reqBody)
	if errRmsLimit != nil {
		rmsLmitResponse = internalServerErrorResponse()
		return
	}

	if rmsLimitDetails == nil {
		rmsLmitResponse = internalServerErrorResponse()
		return
	}

	s, ok := rmsLimitDetails[status_val].(string)
	if ok {
		status = s
	}

	rmsLimitDetailsVal, err := json.Marshal(rmsLimitDetails)
	if err != nil {
		rmsLmitResponse = internalServerErrorResponse()
		return
	}
	result = string(rmsLimitDetailsVal)

	if status != status_success {
		rmsLmitResponse = failureResponseStructure(result)
		return
	}

	rmsLmitResponse = successResponseStructure(result)

	return
}

func (fs *firstock) PositionBook(userId string) (positionBookResponse string) {
	var status string = status_failed
	var result string
	jkey, errRead := readJKeyFromConfig(userId)
	if errRead != nil {
		positionBookResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	positionBookDetails, errPositionBook := thefirstock.PositionBookFunction(reqBody)
	if errPositionBook != nil {
		positionBookResponse = internalServerErrorResponse()
		return
	}

	if positionBookDetails == nil {
		positionBookResponse = internalServerErrorResponse()
		return
	}

	s, ok := positionBookDetails[status_val].(string)
	if ok {
		status = s
	}

	positionBookDetailsVal, err := json.Marshal(positionBookDetails)
	if err != nil {
		positionBookResponse = internalServerErrorResponse()
		return
	}
	result = string(positionBookDetailsVal)

	if status != status_success {
		positionBookResponse = failureResponseStructure(result)
		return
	}

	positionBookResponse = successResponseStructure(result)
	return
}

func (fs *firstock) Holdings(userId string) (holdingsResponse string) {
	var status string = status_failed
	var result string
	jkey, errRead := readJKeyFromConfig(userId)
	if errRead != nil {
		holdingsResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	holdingsDetails, errHoldings := thefirstock.HoldingsFunction(reqBody)
	if errHoldings != nil {
		holdingsResponse = internalServerErrorResponse()
		return
	}

	if holdingsDetails == nil {
		holdingsResponse = internalServerErrorResponse()
		return
	}

	s, ok := holdingsDetails[status_val].(string)
	if ok {
		status = s
	}

	holdingsDetailsVal, err := json.Marshal(holdingsDetails)
	if err != nil {
		holdingsResponse = internalServerErrorResponse()
		return
	}
	result = string(holdingsDetailsVal)

	if status != status_success {
		holdingsResponse = failureResponseStructure(result)
		return
	}

	holdingsResponse = successResponseStructure(result)
	return
}

func (fs *firstock) OrderBook(userId string) (orderBookResponse string) {
	var status string = status_failed
	var result string
	jkey, errRead := readJKeyFromConfig(userId)
	if errRead != nil {
		orderBookResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	orderBookDetails, errOrderBook := thefirstock.OrderBookFunction(reqBody)
	if errOrderBook != nil {
		orderBookResponse = internalServerErrorResponse()
		return
	}

	if orderBookDetails == nil {
		orderBookResponse = internalServerErrorResponse()
		return
	}

	s, ok := orderBookDetails[status_val].(string)
	if ok {
		status = s
	}

	orderBookDetailsVal, err := json.Marshal(orderBookDetails)
	if err != nil {
		orderBookResponse = internalServerErrorResponse()
		return
	}
	result = string(orderBookDetailsVal)

	if status != status_success {
		orderBookResponse = failureResponseStructure(result)
		return
	}

	orderBookResponse = successResponseStructure(result)
	return
}

func (fs *firstock) GetExpiry(getExpiryRequest GetInfoRequest) (getExpiryResponse string) {
	var status string = status_failed
	var result string
	jkey, errRead := readJKeyFromConfig(getExpiryRequest.UserId)
	if errRead != nil {
		getExpiryResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := GetInfoRequestBody{
		UserId:        getExpiryRequest.UserId,
		JKey:          jkey,
		Exchange:      getExpiryRequest.Exchange,
		TradingSymbol: getExpiryRequest.TradingSymbol,
	}

	getExpiryDetails, errGetExpiry := thefirstock.GetExpiryFunction(reqBody)
	if errGetExpiry != nil {
		getExpiryResponse = internalServerErrorResponse()
		return
	}

	if getExpiryDetails == nil {
		getExpiryResponse = internalServerErrorResponse()
		return
	}

	s, ok := getExpiryDetails[status_val].(string)
	if ok {
		status = s
	}

	getExpiryDetailsVal, err := json.Marshal(getExpiryDetails)
	if err != nil {
		getExpiryResponse = internalServerErrorResponse()
		return
	}
	result = string(getExpiryDetailsVal)

	if status != status_success {
		getExpiryResponse = failureResponseStructure(result)
		return
	}

	getExpiryResponse = successResponseStructure(result)
	return
}

func (fs *firstock) BrokerageCalculator(brokerageCalculatorRequest BrokerageCalculatorRequest) (brokerageCalculatorResponse string) {
	var status string = status_failed
	var result string
	jkey, errRead := readJKeyFromConfig(brokerageCalculatorRequest.UserId)
	if errRead != nil {
		brokerageCalculatorResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := BrokerageCalculatorRequestBody{
		UserId:          brokerageCalculatorRequest.UserId,
		JKey:            jkey,
		Exchange:        brokerageCalculatorRequest.Exchange,
		TradingSymbol:   brokerageCalculatorRequest.TradingSymbol,
		TransactionType: brokerageCalculatorRequest.TransactionType,
		Product:         brokerageCalculatorRequest.Product,
		Quantity:        brokerageCalculatorRequest.Quantity,
		Price:           brokerageCalculatorRequest.Price,
		StrikePrice:     brokerageCalculatorRequest.StrikePrice,
		InstName:        brokerageCalculatorRequest.InstName,
		LotSize:         brokerageCalculatorRequest.LotSize,
	}

	brockerageCalculatorDetails, errbrockerageCalculator := thefirstock.BrokerageCalculatorFunction(reqBody)
	if errbrockerageCalculator != nil {
		brokerageCalculatorResponse = internalServerErrorResponse()
		return
	}

	if brockerageCalculatorDetails == nil {
		brokerageCalculatorResponse = internalServerErrorResponse()
		return
	}

	s, ok := brockerageCalculatorDetails[status_val].(string)
	if ok {
		status = s
	}

	brockerageCalculatorDetailsVal, err := json.Marshal(brockerageCalculatorDetails)
	if err != nil {
		brokerageCalculatorResponse = internalServerErrorResponse()
		return
	}
	result = string(brockerageCalculatorDetailsVal)

	if status != status_success {
		brokerageCalculatorResponse = failureResponseStructure(result)
		return
	}

	brokerageCalculatorResponse = successResponseStructure(result)
	return
}

func (fs *firstock) BasketMargin(basketMarginRequest BasketMarginRequest) (basketMarginResponse string) {
	var status string = status_failed
	var result string
	jkey, errRead := readJKeyFromConfig(basketMarginRequest.UserId)
	if errRead != nil {
		basketMarginResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := BasketMarginRequestBody{
		UserId:           basketMarginRequest.UserId,
		JKey:             jkey,
		Exchange:         basketMarginRequest.Exchange,
		TradingSymbol:    basketMarginRequest.TradingSymbol,
		TransactionType:  basketMarginRequest.TransactionType,
		Product:          basketMarginRequest.Product,
		Quantity:         basketMarginRequest.Quantity,
		Price:            basketMarginRequest.Price,
		PriceType:        basketMarginRequest.PriceType,
		BasketListParams: basketMarginRequest.BasketListParams,
	}

	basketMarginDetails, errbasketMargin := thefirstock.BasketMarginFunction(reqBody)
	if errbasketMargin != nil {
		basketMarginResponse = internalServerErrorResponse()
		return
	}

	if basketMarginDetails == nil {
		basketMarginResponse = internalServerErrorResponse()
		return
	}

	s, ok := basketMarginDetails[status_val].(string)
	if ok {
		status = s
	}

	brockerageCalculatorDetailsVal, err := json.Marshal(basketMarginDetails)
	if err != nil {
		basketMarginResponse = internalServerErrorResponse()
		return
	}
	result = string(brockerageCalculatorDetailsVal)

	if status != status_success {
		basketMarginResponse = failureResponseStructure(result)
		return
	}

	basketMarginResponse = successResponseStructure(result)
	return
}

func (fs *firstock) GetSecurityInfo(getSecurityInfoRequest GetInfoRequest) (getSecurityInfoResponse string) {
	var status string = status_failed
	var result string
	jkey, errRead := readJKeyFromConfig(getSecurityInfoRequest.UserId)
	if errRead != nil {
		getSecurityInfoResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := GetInfoRequestBody{
		UserId:        getSecurityInfoRequest.UserId,
		JKey:          jkey,
		Exchange:      getSecurityInfoRequest.Exchange,
		TradingSymbol: getSecurityInfoRequest.TradingSymbol,
	}

	getSecurityInfoDetails, errGetSecurityInfo := thefirstock.GetSecurityInfoFunction(reqBody)
	if errGetSecurityInfo != nil {
		getSecurityInfoResponse = internalServerErrorResponse()
		return
	}

	if getSecurityInfoDetails == nil {
		getSecurityInfoResponse = internalServerErrorResponse()
		return
	}

	s, ok := getSecurityInfoDetails[status_val].(string)
	if ok {
		status = s
	}

	getSecurityInfoDetailsVal, err := json.Marshal(getSecurityInfoDetails)
	if err != nil {
		getSecurityInfoResponse = internalServerErrorResponse()
		return
	}
	result = string(getSecurityInfoDetailsVal)

	if status != status_success {
		getSecurityInfoResponse = failureResponseStructure(result)
		return
	}

	getSecurityInfoResponse = successResponseStructure(result)
	return
}

func (fs *firstock) ProductConversion(productConversionRequest ProductConversionRequest) (productConversionResponse string) {
	var status string = status_failed
	var result string

	jkey, errRead := readJKeyFromConfig(productConversionRequest.UserId)

	if errRead != nil {
		productConversionResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := ProductConversionRequestBody{
		UserId:          productConversionRequest.UserId,
		JKey:            jkey,
		TradingSymbol:   productConversionRequest.TradingSymbol,
		Exchange:        productConversionRequest.Exchange,
		PreviousProduct: productConversionRequest.PreviousProduct,
		Product:         productConversionRequest.Product,
		Quantity:        productConversionRequest.Quantity,
	}

	productConversionDetails, errproductConversion := thefirstock.ProductConversionFunction(reqBody)
	if errproductConversion != nil {
		productConversionResponse = internalServerErrorResponse()
		return
	}
	if productConversionDetails == nil {
		productConversionResponse = internalServerErrorResponse()
		return
	}

	s, ok := productConversionDetails[status_val].(string)
	if ok {
		status = s
	}

	productConversionDetailsVal, err := json.Marshal(productConversionDetails)
	if err != nil {
		productConversionResponse = internalServerErrorResponse()
		return
	}
	result = string(productConversionDetailsVal)

	if status != status_success {
		productConversionResponse = failureResponseStructure(result)
		return
	}

	productConversionResponse = successResponseStructure(result)
	return
}

// Market Connect
func (fs *firstock) GetQuote(getQuoteRequest GetInfoRequest) (getQuoteResponse string) {
	var status string = status_failed
	var result string

	jkey, errRead := readJKeyFromConfig(getQuoteRequest.UserId)
	if errRead != nil {
		getQuoteResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := GetInfoRequestBody{
		UserId:        getQuoteRequest.UserId,
		JKey:          jkey,
		Exchange:      getQuoteRequest.Exchange,
		TradingSymbol: getQuoteRequest.TradingSymbol,
	}

	getQuoteDetails, errGetQuote := thefirstock.GetQuoteFunction(reqBody)
	if errGetQuote != nil {
		getQuoteResponse = internalServerErrorResponse()
		return
	}
	if getQuoteDetails == nil {
		getQuoteResponse = internalServerErrorResponse()
		return
	}

	s, ok := getQuoteDetails[status_val].(string)
	if ok {
		status = s
	}

	getQuoteDetailsVal, err := json.Marshal(getQuoteDetails)
	if err != nil {
		getQuoteResponse = internalServerErrorResponse()
		return
	}
	result = string(getQuoteDetailsVal)

	if status != status_success {
		getQuoteResponse = failureResponseStructure(result)
		return
	}

	getQuoteResponse = successResponseStructure(result)
	return

}

func (fs *firstock) GetQuoteLtp(getQuoteLtpRequest GetInfoRequest) (getQuoteLtpResponse string) {
	var status string = status_failed
	var result string

	jkey, errRead := readJKeyFromConfig(getQuoteLtpRequest.UserId)
	if errRead != nil {
		getQuoteLtpResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := GetInfoRequestBody{
		UserId:        getQuoteLtpRequest.UserId,
		JKey:          jkey,
		Exchange:      getQuoteLtpRequest.Exchange,
		TradingSymbol: getQuoteLtpRequest.TradingSymbol,
	}

	getQuoteLtpDetails, errGetQuoteLtp := thefirstock.GetQuoteLtpFunction(reqBody)
	if errGetQuoteLtp != nil {
		getQuoteLtpResponse = internalServerErrorResponse()
		return
	}

	if getQuoteLtpDetails == nil {
		getQuoteLtpResponse = internalServerErrorResponse()
		return
	}

	s, ok := getQuoteLtpDetails[status_val].(string)
	if ok {
		status = s
	}

	getQuoteLtpDetailsVal, err := json.Marshal(getQuoteLtpDetails)
	if err != nil {
		getQuoteLtpResponse = internalServerErrorResponse()
		return
	}
	result = string(getQuoteLtpDetailsVal)

	if status != status_success {
		getQuoteLtpResponse = failureResponseStructure(result)
		return
	}

	getQuoteLtpResponse = successResponseStructure(result)
	return
}

func (fs *firstock) GetMultiQuotes(getMultiQuotesRequest GetMultiQuotesRequest) (getMultiQuotesResponse string) {
	var status string = status_failed
	var result string

	jkey, errRead := readJKeyFromConfig(getMultiQuotesRequest.UserId)
	if errRead != nil {
		getMultiQuotesResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := GetMultiQuotesRequestBody{
		UserId: getMultiQuotesRequest.UserId,
		JKey:   jkey,
		Data:   getMultiQuotesRequest.Data,
	}

	getMultiQuotesDetails, errGetMultiQuotes := thefirstock.GetMultiQuotesFunction(reqBody)
	if errGetMultiQuotes != nil {
		getMultiQuotesResponse = internalServerErrorResponse()
		return
	}

	if getMultiQuotesDetails == nil {
		getMultiQuotesResponse = internalServerErrorResponse()
		return
	}

	s, ok := getMultiQuotesDetails[status_val].(string)
	if ok {
		status = s
	}

	getMultiQuotesDetailsVal, err := json.Marshal(getMultiQuotesDetails)
	if err != nil {
		getMultiQuotesResponse = internalServerErrorResponse()
		return
	}
	result = string(getMultiQuotesDetailsVal)

	if status != status_success {
		getMultiQuotesResponse = failureResponseStructure(result)
		return
	}

	getMultiQuotesResponse = successResponseStructure(result)
	return

}

func (fs *firstock) GetMultiQuotesLtp(getMultiQuotesRequest GetMultiQuotesRequest) (getMultiQuotesLtpResponse string) {
	var status string = status_failed
	var result string

	jkey, errRead := readJKeyFromConfig(getMultiQuotesRequest.UserId)
	if errRead != nil {
		getMultiQuotesLtpResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := GetMultiQuotesRequestBody{
		UserId: getMultiQuotesRequest.UserId,
		JKey:   jkey,
		Data:   getMultiQuotesRequest.Data,
	}

	getMultiQuotesLtpDetails, errGetMultiQuotesLtp := thefirstock.GetMultiQuotesLtpFunction(reqBody)
	if errGetMultiQuotesLtp != nil {
		getMultiQuotesLtpResponse = internalServerErrorResponse()
		return
	}

	if getMultiQuotesLtpDetails == nil {
		getMultiQuotesLtpResponse = internalServerErrorResponse()
		return
	}

	s, ok := getMultiQuotesLtpDetails[status_val].(string)
	if ok {
		status = s
	}

	getMultiQuotesLtpDetailsVal, err := json.Marshal(getMultiQuotesLtpDetails)
	if err != nil {
		getMultiQuotesLtpResponse = internalServerErrorResponse()
		return
	}
	result = string(getMultiQuotesLtpDetailsVal)

	if status != status_success {
		getMultiQuotesLtpResponse = failureResponseStructure(result)
		return
	}

	getMultiQuotesLtpResponse = successResponseStructure(result)
	return

}

func (fs *firstock) IndexList(userId string) (indexListResponse string) {
	var status string = status_failed
	var result string

	jkey, errRead := readJKeyFromConfig(userId)
	if errRead != nil {
		indexListResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	indexListDetails, errIndexList := thefirstock.IndexListFunction(reqBody)
	if errIndexList != nil {
		indexListResponse = internalServerErrorResponse()
		return
	}

	if indexListDetails == nil {
		indexListResponse = internalServerErrorResponse()
		return
	}

	s, ok := indexListDetails[status_val].(string)
	if ok {
		status = s
	}

	indexListDetailsVal, err := json.Marshal(indexListDetails)
	if err != nil {
		indexListResponse = internalServerErrorResponse()
		return
	}
	result = string(indexListDetailsVal)

	if status != status_success {
		indexListResponse = failureResponseStructure(result)
		return
	}

	indexListResponse = successResponseStructure(result)
	return

}

func (fs *firstock) SearchScrips(searchScripsRequest SearchScripsRequest) (searchScripsResponse string) {
	var status string = status_failed
	var result string

	jkey, errRead := readJKeyFromConfig(searchScripsRequest.UserId)
	if errRead != nil {
		searchScripsResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := SearchScripsBody{
		UserId: searchScripsRequest.UserId,
		JKey:   jkey,
		SText:  searchScripsRequest.SText,
	}

	searchScripsDetails, errsearchScrips := thefirstock.SearchScripsFunction(reqBody)
	if errsearchScrips != nil {
		searchScripsResponse = internalServerErrorResponse()
		return
	}

	if searchScripsDetails == nil {
		searchScripsResponse = internalServerErrorResponse()
		return
	}

	s, ok := searchScripsDetails[status_val].(string)
	if ok {
		status = s
	}

	searchScripsDetailsVal, err := json.Marshal(searchScripsDetails)
	if err != nil {
		searchScripsResponse = internalServerErrorResponse()
		return
	}
	result = string(searchScripsDetailsVal)

	if status != status_success {
		searchScripsResponse = failureResponseStructure(result)
		return
	}

	searchScripsResponse = successResponseStructure(result)
	return
}

func (fs *firstock) OptionChain(optionChainRequest OptionChainRequest) (optionChainResponse string) {
	var status string = status_failed
	var result string

	jkey, errRead := readJKeyFromConfig(optionChainRequest.UserId)
	if errRead != nil {
		optionChainResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := OptionChainRequestBody{
		UserId:      optionChainRequest.UserId,
		JKey:        jkey,
		Exchange:    optionChainRequest.Exchange,
		Symbol:      optionChainRequest.Symbol,
		Expiry:      optionChainRequest.Expiry,
		Count:       optionChainRequest.Count,
		StrikePrice: optionChainRequest.StrikePrice,
	}

	optionChainDetails, erroptionChain := thefirstock.OptionChainFunction(reqBody)
	if erroptionChain != nil {
		optionChainResponse = internalServerErrorResponse()
		return
	}
	if optionChainDetails == nil {
		optionChainResponse = internalServerErrorResponse()
		return
	}

	s, ok := optionChainDetails[status_val].(string)
	if ok {
		status = s
	}

	optionChainDetailsVal, err := json.Marshal(optionChainDetails)
	if err != nil {
		optionChainResponse = internalServerErrorResponse()
		return
	}
	result = string(optionChainDetailsVal)

	if status != status_success {
		optionChainResponse = failureResponseStructure(result)
		return
	}

	optionChainResponse = successResponseStructure(result)
	return

}

func (fs *firstock) TimePriceSeriesRegularInterval(req TimePriceSeriesIntervalRequest) (timePriceSeriesRegularIntervalResponse string) {
	var status string = status_failed
	var result string

	jkey, errRead := readJKeyFromConfig(req.UserId)
	if errRead != nil {
		timePriceSeriesRegularIntervalResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := TimePriceSeriesIntervalRequestBody{
		UserId:        req.UserId,
		JKey:          jkey,
		Exchange:      req.Exchange,
		Interval:      req.Interval,
		TradingSymbol: req.TradingSymbol,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
	}

	timePriceSeriesDetails, errtimePriceSeries := thefirstock.TimePriceSeriesRegularIntervalFunction(reqBody)
	if errtimePriceSeries != nil {
		timePriceSeriesRegularIntervalResponse = internalServerErrorResponse()
		return
	}

	if timePriceSeriesDetails == nil {
		timePriceSeriesRegularIntervalResponse = internalServerErrorResponse()
		return
	}

	s, ok := timePriceSeriesDetails[status_val].(string)
	if ok {
		status = s
	}

	timePriceSeriesDetailsVal, err := json.Marshal(timePriceSeriesDetails)
	if err != nil {
		timePriceSeriesRegularIntervalResponse = internalServerErrorResponse()
		return
	}
	result = string(timePriceSeriesDetailsVal)

	if status != status_success {
		timePriceSeriesRegularIntervalResponse = failureResponseStructure(result)
		return
	}

	timePriceSeriesRegularIntervalResponse = successResponseStructure(result)
	return
}

func (fs *firstock) TimePriceSeriesDayInterval(req TimePriceSeriesIntervalRequest) (timePriceSeriesDayIntervalResponse string) {
	var status string = status_failed
	var result string

	jkey, errRead := readJKeyFromConfig(req.UserId)
	if errRead != nil {
		timePriceSeriesDayIntervalResponse = pleaseLoginToFirstock()
		return
	}

	reqBody := TimePriceSeriesIntervalRequestBody{
		UserId:        req.UserId,
		JKey:          jkey,
		Exchange:      req.Exchange,
		Interval:      req.Interval,
		TradingSymbol: req.TradingSymbol,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
	}

	timePriceSeriesDetails, errtimePriceSeries := thefirstock.TimePriceSeriesDayIntervalFunction(reqBody)
	if errtimePriceSeries != nil {
		timePriceSeriesDayIntervalResponse = internalServerErrorResponse()
		return
	}
	if timePriceSeriesDetails == nil {
		timePriceSeriesDayIntervalResponse = internalServerErrorResponse()
		return
	}

	s, ok := timePriceSeriesDetails[status_val].(string)
	if ok {
		status = s
	}

	timePriceSeriesDetailsVal, err := json.Marshal(timePriceSeriesDetails)
	if err != nil {
		timePriceSeriesDayIntervalResponse = internalServerErrorResponse()
		return
	}
	result = string(timePriceSeriesDetailsVal)

	if status != status_success {
		timePriceSeriesDayIntervalResponse = failureResponseStructure(result)
		return
	}

	timePriceSeriesDayIntervalResponse = successResponseStructure(result)
	return

}

type FirstockAPI interface {
	Login(reqBody LoginRequest) (jsonResponse string)
	Logout(userId string) (jsonResponse string)
	UserDetails(userId string) (userDetailsResponse string)
	PlaceOrder(req PlaceOrderRequest) (jsonResponse string)
	OrderMargin(req OrderMarginRequest) (jsonResponse string)
	SingleOrderHistory(req OrderRequest) (jsonResponse string)
	CancelOrder(req OrderRequest) (jsonResponse string)
	ModifyOrder(req ModifyOrderRequest) (jsonResponse string)
	TradeBook(userId string) (jsonResponse string)
	RMSLmit(userId string) (jsonResponse string)
	PositionBook(userId string) (jsonResponse string)
	Holdings(userId string) (jsonResponse string)
	OrderBook(userId string) (jsonResponse string)
	GetExpiry(getExpiryRequest GetInfoRequest) (jsonResponse string)
	BrokerageCalculator(brokerageCalculatorRequest BrokerageCalculatorRequest) (jsonResponse string)
	BasketMargin(basketMarginRequest BasketMarginRequest) (jsonResponse string)
	GetSecurityInfo(getSecurityInfoRequest GetInfoRequest) (jsonResponse string)
	ProductConversion(productConversionRequest ProductConversionRequest) (jsonResponse string)
	GetQuote(getQuoteRequest GetInfoRequest) (jsonResponse string)
	GetQuoteLtp(getQuoteLtpRequest GetInfoRequest) (jsonResponse string)
	GetMultiQuotes(getMultiQuotesRequest GetMultiQuotesRequest) (jsonResponse string)
	GetMultiQuotesLtp(getMultiQuotesRequest GetMultiQuotesRequest) (jsonResponse string)
	IndexList(userId string) (jsonResponse string)
	SearchScrips(searchScripsRequest SearchScripsRequest) (jsonResponse string)
	OptionChain(optionChainRequest OptionChainRequest) (jsonResponse string)
	TimePriceSeriesRegularInterval(req TimePriceSeriesIntervalRequest) (jsonResponse string)
	TimePriceSeriesDayInterval(req TimePriceSeriesIntervalRequest) (jsonResponse string)
}

// internal instance, not exported
var firstockAPI FirstockAPI = &firstock{}

func Login(reqBody LoginRequest) (jsonResponse string) {
	return firstockAPI.Login(reqBody)
}

func Logout(userId string) (jsonResponse string) {
	return firstockAPI.Logout(userId)
}

func UserDetails(userId string) (userDetailsResponse string) {
	return firstockAPI.UserDetails(userId)
}

func PlaceOrder(req PlaceOrderRequest) (jsonResponse string) {
	return firstockAPI.PlaceOrder(req)
}

func OrderMargin(req OrderMarginRequest) (jsonResponse string) {
	return firstockAPI.OrderMargin(req)
}

func SingleOrderHistory(req OrderRequest) (jsonResponse string) {
	return firstockAPI.SingleOrderHistory(req)
}

func CancelOrder(req OrderRequest) (jsonResponse string) {
	return firstockAPI.CancelOrder(req)
}

func ModifyOrder(req ModifyOrderRequest) (jsonResponse string) {
	return firstockAPI.ModifyOrder(req)
}

func TradeBook(userId string) (jsonResponse string) {
	return firstockAPI.TradeBook(userId)
}

func RMSLmit(userId string) (jsonResponse string) {
	return firstockAPI.RMSLmit(userId)
}

func PositionBook(userId string) (jsonResponse string) {
	return firstockAPI.PositionBook(userId)
}

func Holdings(userId string) (jsonResponse string) {
	return firstockAPI.Holdings(userId)
}

func OrderBook(userId string) (jsonResponse string) {
	return firstockAPI.OrderBook(userId)
}

func GetExpiry(getExpiryRequest GetInfoRequest) (jsonResponse string) {
	return firstockAPI.GetExpiry(getExpiryRequest)
}

func BrokerageCalculator(brokerageCalculatorRequest BrokerageCalculatorRequest) (jsonResponse string) {
	return firstockAPI.BrokerageCalculator(brokerageCalculatorRequest)
}

func BasketMargin(basketMarginRequest BasketMarginRequest) (jsonResponse string) {
	return firstockAPI.BasketMargin(basketMarginRequest)
}

func GetSecurityInfo(getSecurityInfoRequest GetInfoRequest) (jsonResponse string) {
	return firstockAPI.GetSecurityInfo(getSecurityInfoRequest)
}

func ProductConversion(productConversionRequest ProductConversionRequest) (jsonResponse string) {
	return firstockAPI.ProductConversion(productConversionRequest)
}

func GetQuote(getQuoteRequest GetInfoRequest) (jsonResponse string) {
	return firstockAPI.GetQuote(getQuoteRequest)
}

func GetQuoteLtp(getQuoteLtpRequest GetInfoRequest) (jsonResponse string) {
	return firstockAPI.GetQuoteLtp(getQuoteLtpRequest)
}

func GetMultiQuotes(getMultiQuotesRequest GetMultiQuotesRequest) (jsonResponse string) {
	return firstockAPI.GetMultiQuotes(getMultiQuotesRequest)
}

func GetMultiQuotesLtp(getMultiQuotesRequest GetMultiQuotesRequest) (jsonResponse string) {
	return firstockAPI.GetMultiQuotesLtp(getMultiQuotesRequest)
}

func IndexList(userId string) (jsonResponse string) {
	return firstockAPI.IndexList(userId)
}

func SearchScrips(searchScripsRequest SearchScripsRequest) (jsonResponse string) {
	return firstockAPI.SearchScrips(searchScripsRequest)
}

func OptionChain(optionChainRequest OptionChainRequest) (jsonResponse string) {
	return firstockAPI.OptionChain(optionChainRequest)
}

func TimePriceSeriesRegularInterval(req TimePriceSeriesIntervalRequest) (jsonResponse string) {
	return firstockAPI.TimePriceSeriesRegularInterval(req)
}

func TimePriceSeriesDayInterval(req TimePriceSeriesIntervalRequest) (jsonResponse string) {
	return firstockAPI.TimePriceSeriesDayInterval(req)
}
