// Copyright (c) [2025] [abc]
// SPDX-License-Identifier: MIT
package Abc

import (
	"encoding/json"
	"errors"
	"fmt"
)

type firstock struct{}

var thefirstock = &apifunctions{}

// Call Login function to login to Firstock
// It takes a LoginRequest struct as input and returns a JSON response string and an error if any.
func (fs *firstock) Login(reqBody LoginRequest) (jsonResponse string, err error) {
	var login map[string]interface{}
	var status string

	var loginRequest LoginRequest = LoginRequest{
		UserId:     reqBody.UserId,
		Password:   EncodePassword(reqBody.Password),
		TOTP:       reqBody.TOTP,
		VendorCode: reqBody.VendorCode,
		APIKey:     reqBody.APIKey,
	}
	login, err = thefirstock.LoginFunction(
		loginRequest,
	)

	if login != nil {
		s, ok := login[status_val].(string)
		if ok {
			status = s
		}
	}
	if err != nil || login == nil || status != success_status {
		err = errors.New(login_failed)
		return
	} else {

		// Extract SUserToken from login response
		dataMap, ok := login[data].(map[string]interface{})
		if !ok {
			err = errors.New(login_failed)
			fmt.Println("login[\"data\"] is not a map[string]interface{}")
			return
		}

		sUserToken, ok := dataMap[susertoken].(string)
		if !ok {
			fmt.Printf("failed to extract SUserToken from login response: %v", err)
			err = errors.New(login_failed)
			return
		}

		// Write the following to a config.json file. Create the file if it does not exist.
		SaveJKeyToConfig(LogoutRequest{
			UserId: reqBody.UserId,
			JKey:   sUserToken,
		})

		jsonBytes, errRes := json.Marshal(login)
		if errRes != nil {
			fmt.Println(errRes)
			err = errors.New(login_failed)
			return
		}
		jsonResponse = string(jsonBytes)
	}
	return
}

// Call Logout function to logout from Firstock
// It takes a userId as input and returns a JSON response string and an error if any.
func (fs *firstock) Logout(userId string) (jsonResponse string, err error) {
	var logout LogoutRequest
	logout.UserId = userId
	logout.JKey = ""

	config_file_path := getPackageConfigPath()
	// Read jKey for userId from config.json
	jkey, errRead := ReadJKeyFromConfig(config_file_path, userId)
	if errRead != nil {
		fmt.Println("Error reading jkey from config:", errRead)
		err = errors.New(login_first_to_logout)
		return
	}

	logout.JKey = jkey
	logoutResponse, errLogout := thefirstock.LogoutFunction(logout)
	if errLogout != nil {
		fmt.Println("Logout failed:", errLogout)
		err = errors.New(failed_to_logout)
		return
	}
	if logoutResponse[status_val] == "success" {
		// Remove userId from config.json
		RemoveJKeyFromConfig(logout.UserId)
	}
	jsonBytes, errRes := json.Marshal(logoutResponse)
	if errRes != nil {
		err = errors.New(failed_to_logout)
		fmt.Println("Error marshalling logout response:", errRes)
		return
	}
	jsonResponse = string(jsonBytes)
	return jsonResponse, nil
}

// Call UserDetails function to fetch user details from Firstock
// It takes a userId as input and returns a JSON response string and an error if any.
func (fs *firstock) UserDetails(userId string) (userDetailsResponse string, err error) {
	var userDetailsRequest UserDetailsRequest

	config_file_path := getPackageConfigPath()
	// Read jKey for userId from config.json
	jkey, errRead := ReadJKeyFromConfig(config_file_path, userId)
	if errRead != nil {
		fmt.Println("Error reading jkey from config:", errRead)
		err = errors.New(login_first_to_fetch_user_details)
		return
	}

	userDetailsRequest.JKey = jkey
	userDetailsRequest.UserId = userId

	userDetails, errRes := thefirstock.UserDetailsFunction(userDetailsRequest)
	if errRes != nil {
		fmt.Println("Failed to fetch user details:", errRes)
		err = errors.New(failed_to_fetch_user_details)
		return
	}

	jsonBytes, errorVal := json.Marshal(userDetails)
	if errorVal != nil {
		fmt.Println("Error marshalling user details response:", errorVal)
		err = errors.New(failed_to_fetch_user_details)
		return
	}
	var jsonResponse = string(jsonBytes)
	return jsonResponse, nil
}

func (fs *firstock) PlaceOrder(req PlaceOrderRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, req.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_place_order)
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

	orderDetails, errOrder := thefirstock.PlaceOrderFunction(reqBody)
	if errOrder != nil {
		fmt.Println("Error placing order:", errOrder)
		err = errors.New(error_placing_order)
		return
	}

	jsonPayload, errRes := json.Marshal(orderDetails)
	if errRes != nil {
		fmt.Println("Error marshalling order details response:", errRes)
		err = errors.New(error_placing_order)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) OrderMargin(req OrderMarginRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, req.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_order_margin)
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
		fmt.Println("failed to fetch order margin details: %w", errOrder)
		err = errors.New(error_fetching_order_margin)
		return
	}

	jsonPayload, errRes := json.Marshal(orderMarginDetails)
	if errRes != nil {
		fmt.Println("Error marshalling order margin details response:", errRes)
		err = errors.New(error_fetching_order_margin)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) SingleOrderHistory(req OrderRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, req.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_single_order_history)
		return
	}

	reqBody := OrderRequestBody{
		UserId:      req.UserId,
		JKey:        jkey,
		OrderNumber: req.OrderNumber,
	}

	orderMarginDetails, errOrder := thefirstock.SingleOrderHistoryFunction(reqBody)
	if errOrder != nil {
		fmt.Println("failed to fetch single order history details: %w", errOrder)
		err = errors.New(error_fetching_single_order_history)
		return
	}

	jsonPayload, errRes := json.Marshal(orderMarginDetails)
	if errRes != nil {
		fmt.Println("Error marshalling single order history details response:", errRes)
		err = errors.New(error_fetching_single_order_history)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) CancelOrder(req OrderRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, req.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_cancel_order)
		return
	}

	reqBody := OrderRequestBody{
		UserId:      req.UserId,
		JKey:        jkey,
		OrderNumber: req.OrderNumber,
	}

	cancelOrderDetails, errOrder := thefirstock.CancelOrderFunction(reqBody)
	if errOrder != nil {
		fmt.Println("failed to fetch cancel order details: %w", errOrder)
		err = errors.New(error_cancelling_order)
		return
	}

	jsonPayload, errRes := json.Marshal(cancelOrderDetails)
	if errRes != nil {
		fmt.Println("Error marshalling cancel order details response:", errRes)
		err = errors.New(error_cancelling_order)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) ModifyOrder(req ModifyOrderRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, req.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_modify_order)
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
		fmt.Println("failed to fetch modify order details: %w", errOrder)
		err = errors.New(error_modifying_order)
		return
	}

	jsonPayload, errRes := json.Marshal(modifyOrderDetails)
	if errRes != nil {
		fmt.Println("Error marshalling modify order details response:", errRes)
		err = errors.New(error_modifying_order)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) TradeBook(userId string) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, userId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_trade_book_details)
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	tradeBookDetails, errTradeBook := thefirstock.TradeBookFunction(reqBody)
	if errTradeBook != nil {
		fmt.Println("failed to fetch trade book details: %w", errTradeBook)
		err = errors.New(error_fetching_trade_book_details)
		return
	}

	jsonPayload, errRes := json.Marshal(tradeBookDetails)
	if errRes != nil {
		fmt.Println("Error marshalling trade book details response:", errRes)
		err = errors.New(error_fetching_trade_book_details)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) RMSLmit(userId string) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, userId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_rms_limit)
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	rmsLimitDetails, errRmsLimit := thefirstock.RmsLimitFunction(reqBody)
	if errRmsLimit != nil {
		fmt.Println("failed to fetch Rms Limit details: %w", errRmsLimit)
		err = errors.New(error_fetching_rms_limit)
		return
	}

	jsonPayload, errRes := json.Marshal(rmsLimitDetails)
	if errRes != nil {
		fmt.Println("Error marshalling rms limit response:", errRes)
		err = errors.New(error_fetching_rms_limit)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) PositionBook(userId string) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, userId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_position_book)
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	positionBookDetails, errPositionBook := thefirstock.PositionBookFunction(reqBody)
	if errPositionBook != nil {
		fmt.Println("failed to fetch position book details: %w", errPositionBook)
		err = errors.New(error_fetching_position_book)
		return
	}

	jsonPayload, errRes := json.Marshal(positionBookDetails)
	if errRes != nil {
		fmt.Println("Error marshalling position book response:", errRes)
		err = errors.New(error_fetching_position_book)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) Holdings(userId string) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, userId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_holdings)
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	holdings, errHoldings := thefirstock.HoldingsFunction(reqBody)
	if errHoldings != nil {
		fmt.Println("failed to fetch view holdings details: %w", errHoldings)
		err = errors.New(error_fetching_holdings)
		return
	}

	jsonPayload, errRes := json.Marshal(holdings)
	if errRes != nil {
		fmt.Println("Error marshalling view holdings response:", errRes)
		err = errors.New(error_fetching_holdings)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) OrderBook(userId string) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, userId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_order_book)
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	orderBook, errOrderBook := thefirstock.OrderBookFunction(reqBody)
	if errOrderBook != nil {
		fmt.Println("failed to fetch order book details: %w", errOrderBook)
		err = errors.New(error_fetching_order_book)
		return
	}

	jsonPayload, errRes := json.Marshal(orderBook)
	if errRes != nil {
		fmt.Println("Error marshalling order book response:", errRes)
		err = errors.New(error_fetching_order_book)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) GetExpiry(getExpiryRequest GetInfoRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, getExpiryRequest.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_expiry_details)
		return
	}

	reqBody := GetInfoRequestBody{
		UserId:        getExpiryRequest.UserId,
		JKey:          jkey,
		Exchange:      getExpiryRequest.Exchange,
		TradingSymbol: getExpiryRequest.TradingSymbol,
	}

	getExpiry, errGetExpiry := thefirstock.GetExpiryFunction(reqBody)
	if errGetExpiry != nil {
		fmt.Println("failed to fetch order book details: %w", errGetExpiry)
		err = errors.New(error_fetching_expiry_details)
		return
	}

	jsonPayload, errRes := json.Marshal(getExpiry)
	if errRes != nil {
		fmt.Println("Error marshalling order book response:", errRes)
		err = errors.New(error_fetching_expiry_details)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) BrokerageCalculator(brokerageCalculatorRequest BrokerageCalculatorRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, brokerageCalculatorRequest.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_brokerage_calculator_details)
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

	brockerageCalculator, errbrockerageCalculator := thefirstock.BrokerageCalculatorFunction(reqBody)
	if errbrockerageCalculator != nil {
		fmt.Println("failed to fetch brokerage calculator details: %w", errbrockerageCalculator)
		err = errors.New(error_fetching_brokerage_calculator_details)
		return
	}

	jsonPayload, errRes := json.Marshal(brockerageCalculator)
	if errRes != nil {
		fmt.Println("Error marshalling brokerage calculator response:", errRes)
		err = errors.New(error_fetching_brokerage_calculator_details)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) BasketMargin(basketMarginRequest BasketMarginRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, basketMarginRequest.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_basket_margin_details)
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

	basketMargin, errbasketMargin := thefirstock.BasketMarginFunction(reqBody)
	if errbasketMargin != nil {
		fmt.Println("failed to fetch basket margin details: %w", errbasketMargin)
		err = errors.New(error_fetching_basket_margin_details)
		return
	}

	jsonPayload, errRes := json.Marshal(basketMargin)
	if errRes != nil {
		fmt.Println("Error marshalling basket margin response:", errRes)
		err = errors.New(error_fetching_basket_margin_details)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) GetSecurityInfo(getSecurityInfoRequest GetInfoRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, getSecurityInfoRequest.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_get_security_info)
		return
	}

	reqBody := GetInfoRequestBody{
		UserId:        getSecurityInfoRequest.UserId,
		JKey:          jkey,
		Exchange:      getSecurityInfoRequest.Exchange,
		TradingSymbol: getSecurityInfoRequest.TradingSymbol,
	}

	getSecurityInfo, errGetSecurityInfo := thefirstock.GetSecurityInfoFunction(reqBody)
	if errGetSecurityInfo != nil {
		fmt.Println("failed to fetch security info details: %w", errGetSecurityInfo)
		err = errors.New(error_getting_security_info)
		return
	}

	jsonPayload, errRes := json.Marshal(getSecurityInfo)
	if errRes != nil {
		fmt.Println("Error marshalling security info response:", errRes)
		err = errors.New(error_getting_security_info)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) ProductConversion(productConversionRequest ProductConversionRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, productConversionRequest.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_product_conversion_details)
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

	productConversion, errproductConversion := thefirstock.ProductConversionFunction(reqBody)
	if errproductConversion != nil {
		fmt.Println("failed to fetch security info details: %w", errproductConversion)
		err = errors.New(error_fetching_product_conversion_details)
		return
	}

	jsonPayload, errRes := json.Marshal(productConversion)
	if errRes != nil {
		fmt.Println("Error marshalling security info response:", errRes)
		err = errors.New(error_fetching_product_conversion_details)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

// Market Connect
func (fs *firstock) GetQuote(getQuoteRequest GetInfoRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, getQuoteRequest.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_get_quote)
		return
	}

	reqBody := GetInfoRequestBody{
		UserId:        getQuoteRequest.UserId,
		JKey:          jkey,
		Exchange:      getQuoteRequest.Exchange,
		TradingSymbol: getQuoteRequest.TradingSymbol,
	}

	getQuote, errGetQuote := thefirstock.GetQuoteFunction(reqBody)
	if errGetQuote != nil {
		fmt.Println("failed to fetch security info details: %w", errGetQuote)
		err = errors.New(error_get_quote)
		return
	}

	jsonPayload, errRes := json.Marshal(getQuote)
	if errRes != nil {
		fmt.Println("Error marshalling security info response:", errRes)
		err = errors.New(error_get_quote)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) GetQuoteLtp(getQuoteLtpRequest GetInfoRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, getQuoteLtpRequest.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_get_quote_ltp)
		return
	}

	reqBody := GetInfoRequestBody{
		UserId:        getQuoteLtpRequest.UserId,
		JKey:          jkey,
		Exchange:      getQuoteLtpRequest.Exchange,
		TradingSymbol: getQuoteLtpRequest.TradingSymbol,
	}

	getQuoteLtp, errGetQuoteLtp := thefirstock.GetQuoteLtpFunction(reqBody)
	if errGetQuoteLtp != nil {
		fmt.Println("failed to fetch get quote ltp details: %w", errGetQuoteLtp)
		err = errors.New(error_get_quote_ltp)
		return
	}

	jsonPayload, errRes := json.Marshal(getQuoteLtp)
	if errRes != nil {
		fmt.Println("Error marshalling get quote ltp response:", errRes)
		err = errors.New(error_get_quote_ltp)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) GetMultiQuotes(getMultiQuotesRequest GetMultiQuotesRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, getMultiQuotesRequest.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_get_multi_quotes)
		return
	}

	reqBody := GetMultiQuotesRequestBody{
		UserId: getMultiQuotesRequest.UserId,
		JKey:   jkey,
		Data:   getMultiQuotesRequest.Data,
	}

	getMultiQuotes, errGetMultiQuotes := thefirstock.GetMultiQuotesFunction(reqBody)
	if errGetMultiQuotes != nil {
		fmt.Println("failed to fetch get multi quotes details: %w", errGetMultiQuotes)
		err = errors.New(error_get_multi_quotes)
		return
	}

	jsonPayload, errRes := json.Marshal(getMultiQuotes)
	if errRes != nil {
		fmt.Println("Error marshalling get quote ltp response:", errRes)
		err = errors.New(error_get_multi_quotes)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) GetMultiQuotesLtp(getMultiQuotesRequest GetMultiQuotesRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, getMultiQuotesRequest.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_get_multi_quotes_ltp)
		return
	}

	reqBody := GetMultiQuotesRequestBody{
		UserId: getMultiQuotesRequest.UserId,
		JKey:   jkey,
		Data:   getMultiQuotesRequest.Data,
	}

	getMultiQuotesLtp, errGetMultiQuotesLtp := thefirstock.GetMultiQuotesLtpFunction(reqBody)
	if errGetMultiQuotesLtp != nil {
		fmt.Println("failed to fetch get multi quotes details: %w", errGetMultiQuotesLtp)
		err = errors.New(error_get_multi_quotes_ltp)
		return
	}

	jsonPayload, errRes := json.Marshal(getMultiQuotesLtp)
	if errRes != nil {
		fmt.Println("Error marshalling get quote ltp response:", errRes)
		err = errors.New(error_get_multi_quotes_ltp)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) IndexList(userId string) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, userId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_index_list)
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	indexList, errIndexList := thefirstock.IndexListFunction(reqBody)
	if errIndexList != nil {
		fmt.Println("failed to fetch index list details: %w", errIndexList)
		err = errors.New(error_fetching_index_list)
		return
	}

	jsonPayload, errRes := json.Marshal(indexList)
	if errRes != nil {
		fmt.Println("Error marshalling index list response:", errRes)
		err = errors.New(error_fetching_index_list)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) SearchScrips(searchScripsRequest SearchScripsRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, searchScripsRequest.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_search_scrips)
		return
	}

	reqBody := SearchScripsBody{
		UserId: searchScripsRequest.UserId,
		JKey:   jkey,
		SText:  searchScripsRequest.SText,
	}

	searchScrips, errsearchScrips := thefirstock.SearchScripsFunction(reqBody)
	if errsearchScrips != nil {
		fmt.Println("failed to fetch index list details: %w", errsearchScrips)
		err = errors.New(error_fetching_search_scrips)
		return
	}

	jsonPayload, errRes := json.Marshal(searchScrips)
	if errRes != nil {
		fmt.Println("Error marshalling index list response:", errRes)
		err = errors.New(error_fetching_search_scrips)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) OptionChain(optionChainRequest OptionChainRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, optionChainRequest.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_option_chain)
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

	optionChain, erroptionChain := thefirstock.OptionChainFunction(reqBody)
	if erroptionChain != nil {
		fmt.Println("failed to fetch option chain details: %w", erroptionChain)
		err = errors.New(error_fetching_option_chain)
		return
	}

	jsonPayload, errRes := json.Marshal(optionChain)
	if errRes != nil {
		fmt.Println("Error marshalling option chain response:", errRes)
		err = errors.New(error_fetching_option_chain)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) TimePriceSeriesRegularInterval(req TimePriceSeriesIntervalRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, req.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_time_price_series_regular_interval)
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

	timePriceSeries, errtimePriceSeries := thefirstock.TimePriceSeriesRegularIntervalFunction(reqBody)
	if errtimePriceSeries != nil {
		fmt.Println("failed to fetch time price series details: %w", errtimePriceSeries)
		err = errors.New(error_fetching_time_price_series_regular_interval)
		return
	}

	jsonPayload, errRes := json.Marshal(timePriceSeries)
	if errRes != nil {
		fmt.Println("Error marshalling time price series response:", errRes)
		err = errors.New(error_fetching_time_price_series_regular_interval)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

func (fs *firstock) TimePriceSeriesDayInterval(req TimePriceSeriesIntervalRequest) (jsonResponse string, err error) {
	config_file_path := getPackageConfigPath()
	jkey, errRead := ReadJKeyFromConfig(config_file_path, req.UserId)
	if errRead != nil {
		fmt.Println("failed to read jkey from config: %w", errRead)
		err = errors.New(login_first_to_fetch_time_price_series_day_interval)
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

	timePriceSeries, errtimePriceSeries := thefirstock.TimePriceSeriesDayIntervalFunction(reqBody)
	if errtimePriceSeries != nil {
		fmt.Println("failed to fetch time price series details: %w", errtimePriceSeries)
		err = errors.New(error_fetching_time_price_series_day_interval)
		return
	}

	jsonPayload, errRes := json.Marshal(timePriceSeries)
	if errRes != nil {
		fmt.Println("Error marshalling time price series response:", errRes)
		err = errors.New(error_fetching_time_price_series_day_interval)
		return
	}

	jsonResponse = string(jsonPayload)

	return
}

type FirstockAPI interface {
	Login(reqBody LoginRequest) (jsonResponse string, err error)
	Logout(userId string) (jsonResponse string, err error)
	UserDetails(userId string) (userDetailsResponse string, err error)
	PlaceOrder(req PlaceOrderRequest) (jsonResponse string, err error)
	OrderMargin(req OrderMarginRequest) (jsonResponse string, err error)
	SingleOrderHistory(req OrderRequest) (jsonResponse string, err error)
	CancelOrder(req OrderRequest) (jsonResponse string, err error)
	ModifyOrder(req ModifyOrderRequest) (jsonResponse string, err error)
	TradeBook(userId string) (jsonResponse string, err error)
	RMSLmit(userId string) (jsonResponse string, err error)
	PositionBook(userId string) (jsonResponse string, err error)
	Holdings(userId string) (jsonResponse string, err error)
	OrderBook(userId string) (jsonResponse string, err error)
	GetExpiry(getExpiryRequest GetInfoRequest) (jsonResponse string, err error)
	BrokerageCalculator(brokerageCalculatorRequest BrokerageCalculatorRequest) (jsonResponse string, err error)
	BasketMargin(basketMarginRequest BasketMarginRequest) (jsonResponse string, err error)
	GetSecurityInfo(getSecurityInfoRequest GetInfoRequest) (jsonResponse string, err error)
	ProductConversion(productConversionRequest ProductConversionRequest) (jsonResponse string, err error)
	GetQuote(getQuoteRequest GetInfoRequest) (jsonResponse string, err error)
	GetQuoteLtp(getQuoteLtpRequest GetInfoRequest) (jsonResponse string, err error)
	GetMultiQuotes(getMultiQuotesRequest GetMultiQuotesRequest) (jsonResponse string, err error)
	GetMultiQuotesLtp(getMultiQuotesRequest GetMultiQuotesRequest) (jsonResponse string, err error)
	IndexList(userId string) (jsonResponse string, err error)
	SearchScrips(searchScripsRequest SearchScripsRequest) (jsonResponse string, err error)
	OptionChain(optionChainRequest OptionChainRequest) (jsonResponse string, err error)
	TimePriceSeriesRegularInterval(req TimePriceSeriesIntervalRequest) (jsonResponse string, err error)
	TimePriceSeriesDayInterval(req TimePriceSeriesIntervalRequest) (jsonResponse string, err error)
}

// internal instance, not exported
var firstockAPI FirstockAPI = &firstock{}

func Login(reqBody LoginRequest) (jsonResponse string, err error) {
	return firstockAPI.Login(reqBody)
}

func Logout(userId string) (jsonResponse string, err error) {
	return firstockAPI.Logout(userId)
}

func UserDetails(userId string) (userDetailsResponse string, err error) {
	return firstockAPI.UserDetails(userId)
}

func PlaceOrder(req PlaceOrderRequest) (jsonResponse string, err error) {
	return firstockAPI.PlaceOrder(req)
}

func OrderMargin(req OrderMarginRequest) (jsonResponse string, err error) {
	return firstockAPI.OrderMargin(req)
}

func SingleOrderHistory(req OrderRequest) (jsonResponse string, err error) {
	return firstockAPI.SingleOrderHistory(req)
}

func CancelOrder(req OrderRequest) (jsonResponse string, err error) {
	return firstockAPI.CancelOrder(req)
}

func ModifyOrder(req ModifyOrderRequest) (jsonResponse string, err error) {
	return firstockAPI.ModifyOrder(req)
}

func TradeBook(userId string) (jsonResponse string, err error) {
	return firstockAPI.TradeBook(userId)
}

func RMSLmit(userId string) (jsonResponse string, err error) {
	return firstockAPI.RMSLmit(userId)
}

func PositionBook(userId string) (jsonResponse string, err error) {
	return firstockAPI.PositionBook(userId)
}

func Holdings(userId string) (jsonResponse string, err error) {
	return firstockAPI.Holdings(userId)
}

func OrderBook(userId string) (jsonResponse string, err error) {
	return firstockAPI.OrderBook(userId)
}

func GetExpiry(getExpiryRequest GetInfoRequest) (jsonResponse string, err error) {
	return firstockAPI.GetExpiry(getExpiryRequest)
}

func BrokerageCalculator(brokerageCalculatorRequest BrokerageCalculatorRequest) (jsonResponse string, err error) {
	return firstockAPI.BrokerageCalculator(brokerageCalculatorRequest)
}

func BasketMargin(basketMarginRequest BasketMarginRequest) (jsonResponse string, err error) {
	return firstockAPI.BasketMargin(basketMarginRequest)
}

func GetSecurityInfo(getSecurityInfoRequest GetInfoRequest) (jsonResponse string, err error) {
	return firstockAPI.GetSecurityInfo(getSecurityInfoRequest)
}

func ProductConversion(productConversionRequest ProductConversionRequest) (jsonResponse string, err error) {
	return firstockAPI.ProductConversion(productConversionRequest)
}

func GetQuote(getQuoteRequest GetInfoRequest) (jsonResponse string, err error) {
	return firstockAPI.GetQuote(getQuoteRequest)
}

func GetQuoteLtp(getQuoteLtpRequest GetInfoRequest) (jsonResponse string, err error) {
	return firstockAPI.GetQuoteLtp(getQuoteLtpRequest)
}

func GetMultiQuotes(getMultiQuotesRequest GetMultiQuotesRequest) (jsonResponse string, err error) {
	return firstockAPI.GetMultiQuotes(getMultiQuotesRequest)
}

func GetMultiQuotesLtp(getMultiQuotesRequest GetMultiQuotesRequest) (jsonResponse string, err error) {
	return firstockAPI.GetMultiQuotesLtp(getMultiQuotesRequest)
}

func IndexList(userId string) (jsonResponse string, err error) {
	return firstockAPI.IndexList(userId)
}

func SearchScrips(searchScripsRequest SearchScripsRequest) (jsonResponse string, err error) {
	return firstockAPI.SearchScrips(searchScripsRequest)
}

func OptionChain(optionChainRequest OptionChainRequest) (jsonResponse string, err error) {
	return firstockAPI.OptionChain(optionChainRequest)
}

func TimePriceSeriesRegularInterval(req TimePriceSeriesIntervalRequest) (jsonResponse string, err error) {
	return firstockAPI.TimePriceSeriesRegularInterval(req)
}

func TimePriceSeriesDayInterval(req TimePriceSeriesIntervalRequest) (jsonResponse string, err error) {
	return firstockAPI.TimePriceSeriesDayInterval(req)
}
