package Firstock

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var userId, password, totp, vendorCode, apiKey = "", "", "", "", ""

func TestLogin(t *testing.T) {
	loginRequest := LoginRequest{
		UserId:     userId,
		Password:   password,
		TOTP:       totp,
		VendorCode: vendorCode,
		APIKey:     apiKey,
	}
	login, err := Login(loginRequest)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(login), &result); err != nil {
		t.Fatalf("Failed to unmarshal login response: %v", err)
	}
	// ✅ Top-level keys check
	expectedTopKeys := map[string]bool{
		"status":  true,
		"message": true,
		"data":    true,
	}

	for key := range result {
		assert.True(t, expectedTopKeys[key], "unexpected top-level key: %s", key)
	}
	for key := range expectedTopKeys {
		_, exists := result[key]
		assert.True(t, exists, "missing top-level key: %s", key)
	}

	// ✅ Check status value
	status, ok := result["status"].(string)
	assert.True(t, ok, `"status" should be a string`)
	assert.Equal(t, "success", status, `"status" must be "success"`)

	// ✅ Check data keys
	data, ok := result["data"].(map[string]interface{})
	assert.True(t, ok, `"data" should be a map`)

	expectedDataKeys := map[string]bool{
		"actid":      true,
		"userName":   true,
		"susertoken": true,
		"email":      true,
	}

	for key := range data {
		assert.True(t, expectedDataKeys[key], "unexpected key in data: %s", key)
	}
	for key := range expectedDataKeys {
		_, exists := data[key]
		assert.True(t, exists, "missing key in data: %s", key)
	}
	// You can add more assertions here to check the data fields if needed
}

func TestUserDetails(t *testing.T) {
	userDetails, err := UserDetails(userId)
	if err != nil {
		t.Fatalf("failed to get user details: %v", err)
		return
	}
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(userDetails), &result); err != nil {
		t.Fatalf("Failed to unmarshal login response: %v", err)
	}
	// ✅ Check top-level keys
	expectedTopKeys := map[string]bool{
		"status": true,
		"data":   true,
	}
	for key := range result {
		assert.True(t, expectedTopKeys[key], "unexpected top-level key: %s", key)
	}
	for key := range expectedTopKeys {
		_, exists := result[key]
		assert.True(t, exists, "missing top-level key: %s", key)
	}

	// ✅ Check status
	status, ok := result["status"].(string)
	assert.True(t, ok, `"status" should be a string`)
	assert.Equal(t, "success", status, `"status" must be 'success'`)

	// ✅ Check data keys
	data, ok := result["data"].(map[string]interface{})
	assert.True(t, ok, `"data" should be a map`)

	expectedDataKeys := map[string]bool{
		"actid":       true,
		"email":       true,
		"exchange":    true,
		"orarr":       true,
		"requestTime": true,
		"uprev":       true,
		"userName":    true,
	}

	for key := range data {
		assert.True(t, expectedDataKeys[key], "unexpected key in data: %s", key)
	}
	for key := range expectedDataKeys {
		_, exists := data[key]
		assert.True(t, exists, "missing key in data: %s", key)
	}

	// You can add more assertions here to check the data fields if needed
}

func TestPlaceOrder(t *testing.T) {
	exchange, retention, product, priceType, tradingSymbol, transactionType, price, triggerPrice, quantity, remarks := "BSE", "DAY", "C", "LMT", "SAWACA", "B", "0.49", "", "1", "Test Order"
	placeOrderRequest := PlaceOrderRequest{
		UserId:          userId,
		Exchange:        exchange,
		Retention:       retention,
		Product:         product,
		PriceType:       priceType,
		TradingSymbol:   tradingSymbol,
		TransactionType: transactionType,
		Price:           price,
		TriggerPrice:    triggerPrice,
		Quantity:        quantity,
		Remarks:         remarks,
	}
	placeOrder, err := PlaceOrder(placeOrderRequest)
	if err != nil {
		t.Error(err)
		return
	}

	// Parse JSON into a map
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(placeOrder), &result); err != nil {
		t.Fatalf("Failed to unmarshal login response: %v", err)
	}
	// Check top-level keys
	expectedKeys := []string{"status", "message", "data"}
	for _, key := range expectedKeys {
		_, exists := result[key]
		assert.True(t, exists, "Key '%s' should be present", key)
	}

	// Check for no extra keys
	assert.Equal(t, len(expectedKeys), len(result), "No extra keys should be present")

	// Check 'status' is "success"
	assert.Equal(t, "success", result["status"], "Status should be 'success'")

	// Check 'data' object contains required keys
	dataMap, ok := result["data"].(map[string]interface{})
	assert.True(t, ok, "Data should be a JSON object")

	expectedDataKeys := []string{"orderNumber", "requestTime"}
	for _, key := range expectedDataKeys {
		_, exists := dataMap[key]
		assert.True(t, exists, "Data key '%s' should be present", key)
	}

	// Optional: Check no extra keys in "data"
	assert.Equal(t, len(expectedDataKeys), len(dataMap), "No extra keys should be present in data")
}

func TestLimit(t *testing.T) {

	rmsLimitDetails, err := RMSLmit(userId)
	if err != nil {
		t.Error(err)
		return
	}

	// Parse JSON into a map
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(rmsLimitDetails), &result); err != nil {
		t.Fatalf("Failed to unmarshal login response: %v", err)
	}

	// ✅ Top-level keys
	expectedTopKeys := map[string]bool{
		"status":  true,
		"message": true,
		"data":    true,
	}
	for key := range result {
		assert.True(t, expectedTopKeys[key], "unexpected top-level key: %s", key)
	}
	for key := range expectedTopKeys {
		_, exists := result[key]
		assert.True(t, exists, "missing top-level key: %s", key)
	}

	// ✅ status must be "success"
	status, ok := result["status"].(string)
	assert.True(t, ok, `"status" must be a string`)
	assert.Equal(t, "success", status, `"status" must be 'success'`)

	// ✅ message must be string
	_, ok = result["message"].(string)
	assert.True(t, ok, `"message" must be a string`)

	// ✅ data must be a map
	data, ok := result["data"].(map[string]interface{})
	assert.True(t, ok, `"data" must be a map/object`)

	// ✅ data keys and types
	expectedDataKeys := map[string]bool{
		"brkcollamt":  true,
		"cash":        true,
		"collateral":  true,
		"expo":        true,
		"marginused":  true,
		"payin":       true,
		"peak_mar":    true,
		"premium":     true,
		"requestTime": true,
		"span":        true,
	}
	for key, val := range data {
		assert.True(t, expectedDataKeys[key], "unexpected key in data: %s", key)
		_, ok := val.(string)
		assert.True(t, ok, "value of %s should be string", key)
	}
	for key := range expectedDataKeys {
		_, exists := data[key]
		assert.True(t, exists, "missing key in data: %s", key)
	}
}

func TestBasketMargin(t *testing.T) {
	basketMarginRequest := BasketMarginRequest{
		UserId:          userId,
		Exchange:        "NSE",
		TransactionType: "B",           // B = Buy, S = Sell
		Product:         "C",           // C = Delivery, I = Intraday, M = Margin Intraday (MIS)
		TradingSymbol:   "RELIANCE-EQ", // Ensure it's the correct symbol
		Quantity:        "1",           // As string
		PriceType:       "MKT",         // Example: "LMT" for Limit, "MKT" for Market
		Price:           "0",           // As string
		BasketListParams: []BasketListParam{
			{
				Exchange:        "NSE",
				TransactionType: "B",
				Product:         "C",
				TradingSymbol:   "IDEA-EQ",
				Quantity:        "1",
				PriceType:       "MKT",
				Price:           "0",
			},
		},
	}

	basketMargin, err := BasketMargin(basketMarginRequest)
	if err != nil {
		t.Error(err)
		return
	}

	var result map[string]interface{}
	errJson := json.Unmarshal([]byte(basketMargin), &result)
	assert.NoError(t, errJson, "should unmarshal JSON successfully")

	// ✅ Top-level keys
	expectedTopKeys := map[string]bool{
		"status":  true,
		"message": true,
		"data":    true,
	}
	for key := range result {
		assert.True(t, expectedTopKeys[key], "unexpected top-level key: %s", key)
	}
	for key := range expectedTopKeys {
		_, exists := result[key]
		assert.True(t, exists, "missing top-level key: %s", key)
	}

	// ✅ Validate status value
	status, ok := result["status"].(string)
	assert.True(t, ok, `"status" should be a string`)
	assert.Equal(t, "success", status, `"status" must be 'success'`)

	// ✅ Validate message type
	_, ok = result["message"].(string)
	assert.True(t, ok, `"message" should be a string`)

	// ✅ Validate data keys
	data, ok := result["data"].(map[string]interface{})
	assert.True(t, ok, `"data" should be a map`)

	expectedDataKeys := map[string]bool{
		"BasketMargin":     true,
		"MarginOnNewOrder": true,
		"PreviousMargin":   true,
		"Remarks":          true,
		"TradedMargin":     true,
		"requestTime":      true,
	}

	for key, val := range data {
		assert.True(t, expectedDataKeys[key], "unexpected key in data: %s", key)

		switch key {
		case "BasketMargin":
			arr, ok := val.([]interface{})
			assert.True(t, ok, `"BasketMargin" should be an array`)
			for i, item := range arr {
				_, ok := item.(string)
				assert.True(t, ok, "BasketMargin[%d] should be a string", i)
			}
		default:
			_, ok := val.(string)
			assert.True(t, ok, "value of %s should be a string", key)
		}
	}

	for key := range expectedDataKeys {
		_, exists := data[key]
		assert.True(t, exists, "missing key in data: %s", key)
	}
}

func TestHoldings(t *testing.T) {
	holdingsDetails, err := Holdings(userId)
	if err != nil {
		assert.Error(t, err)
		return
	}

	var result map[string]interface{}
	errJson := json.Unmarshal([]byte(holdingsDetails), &result)
	assert.NoError(t, errJson, "should unmarshal JSON successfully")

	if errJson != nil {
		assert.Error(t, err)
	}
	// ✅ Top-level keys
	expectedTopKeys := map[string]bool{
		"status":  true,
		"message": true,
		"data":    true,
	}
	for key := range result {
		assert.True(t, expectedTopKeys[key], "unexpected top-level key: %s", key)
	}
	for key := range expectedTopKeys {
		_, exists := result[key]
		assert.True(t, exists, "missing top-level key: %s", key)
	}

	// ✅ Validate status value
	status, ok := result["status"].(string)
	assert.True(t, ok, `"status" should be a string`)
	assert.Equal(t, "success", status, `"status" must be 'success'`)

	// ✅ Validate message type
	_, ok = result["message"].(string)
	assert.True(t, ok, `"message" should be a string`)

	// ✅ Validate data is array of objects
	dataArr, ok := result["data"].([]interface{})
	assert.True(t, ok, `"data" should be an array`)

	// ✅ Validate each item in data
	expectedItemKeys := map[string]bool{
		"exchange":      true,
		"tradingSymbol": true,
	}
	for i, item := range dataArr {
		entry, ok := item.(map[string]interface{})
		assert.True(t, ok, "each data entry should be a JSON object")

		for key := range entry {
			assert.True(t, expectedItemKeys[key], "unexpected key in data[%d]: %s", i, key)
		}
		for key := range expectedItemKeys {
			val, exists := entry[key]
			assert.True(t, exists, "missing key in data[%d]: %s", i, key)
			_, ok := val.(string)
			assert.True(t, ok, "data[%d].%s should be a string", i, key)
		}
	}
}

//--------------------------Market Connect APIs------------------------------

func TestGetQuote(t *testing.T) {
	getQuoteReq := GetInfoRequest{
		UserId:        userId,
		Exchange:      "NSE",
		TradingSymbol: "RELIANCE-EQ",
	}

	getQuoteDetails, err := GetQuote(getQuoteReq)
	if err != nil {
		t.Error(err)
		return
	}

	var response map[string]interface{}
	errJson := json.Unmarshal([]byte(getQuoteDetails), &response)
	assert.NoError(t, errJson, "should unmarshal JSON successfully")

	// Top-level key check
	expectedTopKeys := map[string]bool{
		"status":  true,
		"message": true,
		"data":    true,
	}
	for key := range response {
		assert.True(t, expectedTopKeys[key], "unexpected top-level key: %s", key)
	}
	for key := range expectedTopKeys {
		_, ok := response[key]
		assert.True(t, ok, "missing top-level key: %s", key)
	}

	// Status check
	status, ok := response["status"].(string)
	assert.True(t, ok)
	assert.Equal(t, "success", status)

	// Message check
	_, ok = response["message"].(string)
	assert.True(t, ok)

	// Data validation
	data, ok := response["data"].(map[string]interface{})
	assert.True(t, ok)

	// Check for tradingSymbol existence and value
	tradingSymbolVal, hasTradingSymbol := data["tradingSymbol"].(string)
	assert.True(t, hasTradingSymbol, "missing or invalid tradingSymbol")

	// Define index symbols that should skip ISIN check
	indexSymbols := map[string]bool{
		"BANKNIFTY":  true,
		"FINNIFTY":   true,
		"BANKEX":     true,
		"SENSEX50":   true,
		"NIFTY":      true,
		"MIDCPNIFTY": true,
		"SENSEX":     true,
		"NIFTYNXT50": true,
	}

	// Define expected data keys
	expectedDataKeys := map[string]bool{
		"VWAP": true, "bestBuyOrder1": true, "bestBuyOrder2": true, "bestBuyOrder3": true,
		"bestBuyOrder4": true, "bestBuyOrder5": true,
		"bestBuyPrice1": true, "bestBuyPrice2": true, "bestBuyPrice3": true,
		"bestBuyPrice4": true, "bestBuyPrice5": true,
		"bestBuyQuantity1": true, "bestBuyQuantity2": true, "bestBuyQuantity3": true,
		"bestBuyQuantity4": true, "bestBuyQuantity5": true,
		"bestSellOrder1": true, "bestSellOrder2": true, "bestSellOrder3": true,
		"bestSellOrder4": true, "bestSellOrder5": true,
		"bestSellPrice1": true, "bestSellPrice2": true, "bestSellPrice3": true,
		"bestSellPrice4": true, "bestSellPrice5": true,
		"bestSellQuantity1": true, "bestSellQuantity2": true, "bestSellQuantity3": true,
		"bestSellQuantity4": true, "bestSellQuantity5": true,
		"companyName": true, "dayClosePrice": true, "dayHighPrice": true, "dayLowPrice": true,
		"dayOpenPrice": true, "exchange": true, "lastTradedPrice": true,
		"lotSize": true, "multipler": true, "openInterest": true,
		"priceFactor": true, "pricePrecision": true, "requestTime": true,
		"segment": true, "symbolName": true, "tickSize": true,
		"token": true, "totalBuyQuantity": true, "totalSellQuantity": true,
		"tradingSymbol": true,
	}

	// Only check for ISIN if not an index
	if !indexSymbols[tradingSymbolVal] {
		expectedDataKeys["isin"] = true
		_, ok := data["isin"].(string)
		assert.True(t, ok, "isin must be present and string for: %s", tradingSymbolVal)
	} else {
		// Ensure isin is not present
		_, hasISIN := data["isin"]
		assert.False(t, hasISIN, "isin should NOT be present for index symbol: %s", tradingSymbolVal)
	}

	// Validate that no extra keys are present in data
	for key := range data {
		assert.True(t, expectedDataKeys[key], "unexpected key in data: %s", key)
		_, ok := data[key].(string)
		assert.True(t, ok, "value of key %s is not a string", key)
	}

	// Validate that all expected keys exist
	for key := range expectedDataKeys {
		_, ok := data[key]
		assert.True(t, ok, "missing expected key in data: %s", key)
	}
}

func TestGetQuotesLtp(t *testing.T) {
	getQuoteLtpReq := GetInfoRequest{
		UserId:        userId,
		Exchange:      "NSE",
		TradingSymbol: "NIFTY",
	}

	getQuoteDetails, err := GetQuoteLtp(getQuoteLtpReq)
	if err != nil {
		t.Error(err)
		return
	}
	// Parse JSON into map
	var response map[string]interface{}
	errJson := json.Unmarshal([]byte(getQuoteDetails), &response)
	assert.NoError(t, errJson, "response should unmarshal correctly")

	// 1. Check top-level keys
	expectedTopKeys := []string{"status", "message", "data"}

	for _, key := range expectedTopKeys {
		_, exists := response[key]
		assert.True(t, exists, "top-level key '%s' should be present", key)
	}

	// 2. Check for unexpected top-level keys
	for key := range response {
		assert.Contains(t, expectedTopKeys, key, "unexpected top-level key '%s' found", key)
	}

	// 3. Check "status" is "success"
	assert.Equal(t, "success", response["status"], "status should be 'success'")

	// 4. Check all expected keys in data
	expectedDataKeys := []string{"companyName", "exchange", "lastTradedPrice", "requestTime", "token"}

	dataMap, ok := response["data"].(map[string]interface{})
	assert.True(t, ok, "data should be a map")

	for _, key := range expectedDataKeys {
		_, exists := dataMap[key]
		assert.True(t, exists, "expected data key '%s' should be present", key)
	}

	// 5. Check for unexpected keys in data
	for key := range dataMap {
		assert.Contains(t, expectedDataKeys, key, "unexpected data key '%s' found", key)
	}
}

func TestGetMultiQuotes(t *testing.T) {
	getMultiQuotesReq := GetMultiQuotesRequest{
		UserId: userId, // replace with actual value
		Data: []MultiQuoteData{
			{
				Exchange:      "NSE",
				TradingSymbol: "Nifty 50", // Ensure this matches the broker’s expected format
			},
			{
				Exchange:      "NFO",
				TradingSymbol: "NIFTY03APR25C23500",
			},
		},
	}

	getMultiQuotes, err := GetMultiQuotes(getMultiQuotesReq)
	if err != nil {
		t.Error(err)
		return
	}
	var response map[string]interface{}
	errJson := json.Unmarshal([]byte(getMultiQuotes), &response)
	assert.NoError(t, errJson)

	// Check top-level keys
	expectedTopKeys := []string{"status", "message", "data"}
	for _, key := range expectedTopKeys {
		_, exists := response[key]
		assert.True(t, exists, "Top-level key '%s' missing", key)
	}
	for key := range response {
		assert.Contains(t, expectedTopKeys, key, "Unexpected top-level key '%s'", key)
	}

	// Check status value
	assert.Equal(t, "success", response["status"], "Expected status to be 'success'")

	// Check data is array
	dataArray, ok := response["data"].([]interface{})
	assert.True(t, ok, "data should be an array")

	// Loop through each item
	for _, item := range dataArray {
		entry, ok := item.(map[string]interface{})
		assert.True(t, ok, "Each item in data must be a map")

		// Determine if it's an error entry
		if status, exists := entry["status"]; exists && status == "error" {
			expectedErrorKeys := []string{
				"companyName", "error", "exchange", "identifier", "requestTime", "status", "tradingSymbol",
			}
			for _, key := range expectedErrorKeys {
				_, exists := entry[key]
				assert.True(t, exists, "Missing key '%s' in error item", key)
			}
			for key := range entry {
				assert.Contains(t, expectedErrorKeys, key, "Unexpected key '%s' in error item", key)
			}
		} else {
			// It's a valid quote item
			expectedQuoteKeys := []string{
				"VWAP", "bestBuyOrder1", "bestBuyOrder2", "bestBuyOrder3", "bestBuyOrder4", "bestBuyOrder5",
				"bestBuyPrice1", "bestBuyPrice2", "bestBuyPrice3", "bestBuyPrice4", "bestBuyPrice5",
				"bestBuyQuantity1", "bestBuyQuantity2", "bestBuyQuantity3", "bestBuyQuantity4", "bestBuyQuantity5",
				"bestSellOrder1", "bestSellOrder2", "bestSellOrder3", "bestSellOrder4", "bestSellOrder5",
				"bestSellPrice1", "bestSellPrice2", "bestSellPrice3", "bestSellPrice4", "bestSellPrice5",
				"bestSellQuantity1", "bestSellQuantity2", "bestSellQuantity3", "bestSellQuantity4", "bestSellQuantity5",
				"companyName", "dayClosePrice", "dayHighPrice", "dayLowPrice", "dayOpenPrice",
				"exchange", "identifier", "instrumentName", "lastTradedPrice", "lotSize", "multipler",
				"openInterest", "priceFactor", "pricePrecision", "requestTime", "segment", "symbolName",
				"tickSize", "token", "totalBuyQuantity", "totalSellQuantity", "tradingSymbol",
			}
			for _, key := range expectedQuoteKeys {
				_, exists := entry[key]
				assert.True(t, exists, "Missing key '%s' in quote item", key)
			}
			for key := range entry {
				assert.Contains(t, expectedQuoteKeys, key, "Unexpected key '%s' in quote item", key)
			}
		}
	}

}

func TestGetMultiQuotesLtp(t *testing.T) {
	getMultiQuotesLtpReq := GetMultiQuotesRequest{
		UserId: userId, // replace with actual value
		Data: []MultiQuoteData{
			{
				Exchange:      "NSE",
				TradingSymbol: "Nifty 50", // Ensure this matches the broker’s expected format
			},
			{
				Exchange:      "NFO",
				TradingSymbol: "NIFTY03APR25C23500",
			},
		},
	}

	getMultiQuotesLtp, err := GetMultiQuotesLtp(getMultiQuotesLtpReq)
	if err != nil {
		t.Error(err)
		return
	}
	var response map[string]interface{}
	errJson := json.Unmarshal([]byte(getMultiQuotesLtp), &response)
	assert.NoError(t, errJson, "JSON should unmarshal successfully")

	// 1. Check top-level keys
	expectedTopKeys := []string{"status", "message", "data"}
	for _, key := range expectedTopKeys {
		_, exists := response[key]
		assert.True(t, exists, "Top-level key '%s' should be present", key)
	}
	for key := range response {
		assert.Contains(t, expectedTopKeys, key, "Unexpected top-level key '%s'", key)
	}

	// 2. Check top-level status
	assert.Equal(t, "success", response["status"], "status should be 'success'")

	// 3. Validate data is array
	dataArray, ok := response["data"].([]interface{})
	assert.True(t, ok, "data should be an array")

	// 4. Validate each item in data
	for _, item := range dataArray {
		dataItem, ok := item.(map[string]interface{})
		assert.True(t, ok, "each item in data should be a map")

		// If status is "error" inside item
		if itemStatus, exists := dataItem["status"]; exists && itemStatus == "error" {
			expectedKeys := []string{
				"companyName", "exchange", "identifier", "lastTradedPrice",
				"requestTime", "status", "tradingSymbol", "error",
			}
			for _, key := range expectedKeys {
				_, exists := dataItem[key]
				assert.True(t, exists, "expected key '%s' should be present in error item", key)
			}
			for key := range dataItem {
				assert.Contains(t, expectedKeys, key, "unexpected key '%s' in error item", key)
			}
		} else {
			// Otherwise it's a successful quote
			expectedKeys := []string{
				"companyName", "exchange", "identifier", "lastTradedPrice",
				"requestTime", "token", "tradingSymbol",
			}
			for _, key := range expectedKeys {
				_, exists := dataItem[key]
				assert.True(t, exists, "expected key '%s' should be present in success item", key)
			}
			for key := range dataItem {
				assert.Contains(t, expectedKeys, key, "unexpected key '%s' in success item", key)
			}
		}
	}
}

func TestIndexList(t *testing.T) {
	indexList, err := IndexList(userId)
	if err != nil {
		t.Error(err)
		return
	}
	// Parse the JSON
	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(indexList), &resp)
	assert.NoError(t, errJson)

	// Top-level keys check
	assert.Contains(t, resp, "status", "Missing key: status")
	assert.Contains(t, resp, "message", "Missing key: message")
	assert.Contains(t, resp, "data", "Missing key: data")

	// Check status is success
	assert.Equal(t, "success", resp["status"], "Expected status to be 'success'")

	// Ensure data is an array
	data, ok := resp["data"].([]interface{})
	assert.True(t, ok, "Expected 'data' to be an array")

	expectedKeys := map[string]bool{
		"exchange":      true,
		"token":         true,
		"tradingSymbol": true,
		"symbol":        true,
		"idxname":       true,
	}

	// Check each item
	for i, item := range data {
		index, ok := item.(map[string]interface{})
		assert.True(t, ok, "Item %d in data is not a JSON object", i)

		for key := range index {
			assert.Contains(t, expectedKeys, key, "Unexpected key '%s' in item %d", key, i)
		}

		// Ensure all expected keys are present
		for key := range expectedKeys {
			assert.Contains(t, index, key, "Missing expected key '%s' in item %d", key, i)
		}
	}
}

func TestSearchScrips(t *testing.T) {
	searchScripsRequest := SearchScripsRequest{
		UserId: userId,
		SText:  "RELIANCE",
	}
	searchScrips, err := SearchScrips(searchScripsRequest)
	if err != nil {
		t.Error(err)
		return
	}
	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(searchScrips), &resp)
	assert.NoError(t, errJson)

	assert.Equal(t, "success", resp["status"])
	assert.Equal(t, "Search symbols", resp["message"])

	data, ok := resp["data"].([]interface{})
	assert.True(t, ok, "data should be a list")

	for _, item := range data {
		entry, ok := item.(map[string]interface{})
		assert.True(t, ok)

		expectedKeys := []string{"token", "exchange", "companyName", "representationName", "instrumentName", "tradingSymbol"}
		for _, key := range expectedKeys {
			_, exists := entry[key]
			assert.True(t, exists, "missing key: %s", key)
		}

		// Detect unexpected keys
		for key := range entry {
			assert.Contains(t, expectedKeys, key, "unexpected key: %s", key)
		}
	}
}

func TestOptionChain(t *testing.T) {
	optionChainRequest := OptionChainRequest{
		UserId:      userId,
		Exchange:    "NFO",
		Symbol:      "NIFTY",
		Expiry:      "12JUN25", // Format must match broker format
		Count:       "5",       // Number of strikes above/below
		StrikePrice: "23150",   // ATM strike price
	}
	optionChain, err := OptionChain(optionChainRequest)
	if err != nil {
		t.Error(err)
		return
	}
	// Unmarshal into a generic map
	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(optionChain), &resp)
	assert.NoError(t, errJson, "Response JSON should be valid")

	// Basic field assertions
	assert.Equal(t, "success", resp["status"], "Status should be 'success'")
	assert.Equal(t, "Option chain data retrieved successfully", resp["message"], "Message should match")

	// Validate data is an array
	dataArr, ok := resp["data"].([]interface{})
	assert.True(t, ok, "'data' should be an array")

	// Allowed keys for each option data object
	allowedKeys := map[string]bool{
		"exchange":        true,
		"lastTradedPrice": true,
		"lotSize":         true,
		"optionType":      true,
		"parentToken":     true,
		"pricePrecision":  true,
		"strikePrice":     true,
		"tickSize":        true,
		"token":           true,
		"tradingSymbol":   true,
	}

	// Validate each item
	for i, item := range dataArr {
		itemMap, ok := item.(map[string]interface{})
		assert.True(t, ok, "Each item in 'data' must be an object (index %d)", i)

		for key := range itemMap {
			assert.Truef(t, allowedKeys[key], "Unexpected key '%s' in item at index %d", key, i)
		}
	}
}

func TestTimePriceSeriesRegularInterval(t *testing.T) {
	timePriceSeriesRegularIntervalRequest := TimePriceSeriesIntervalRequest{
		UserId:        userId,
		Exchange:      "NSE",
		TradingSymbol: "NIFTY",
		Interval:      "1mi", // 5 minutes interval
		StartTime:     "09:15:00 23-04-2025",
		EndTime:       "15:29:00 23-04-2025",
	}
	timePriceSeriesRegularInterval, err := TimePriceSeriesRegularInterval(timePriceSeriesRegularIntervalRequest)
	if err != nil {
		t.Error(err)
		return
	}
	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(timePriceSeriesRegularInterval), &resp)
	assert.NoError(t, errJson, "JSON should unmarshal correctly")

	// Check top-level keys
	assert.Equal(t, "success", resp["status"], "Status should be 'success'")
	_, msgOk := resp["message"]
	assert.True(t, msgOk, "Message key should be present")

	// Check data array
	data, ok := resp["data"].([]interface{})
	assert.True(t, ok, "Data should be an array")

	expectedKeys := map[string]bool{
		"time": true, "epochTime": true, "open": true,
		"high": true, "low": true, "close": true,
		"volume": true, "oi": true,
	}

	for i, item := range data {
		record, ok := item.(map[string]interface{})
		assert.True(t, ok, "Each data item should be a map")

		for key := range record {
			assert.True(t, expectedKeys[key], "Unexpected key '%s' in data[%d]", key, i)
		}

		for key := range expectedKeys {
			_, present := record[key]
			assert.True(t, present, "Missing expected key '%s' in data[%d]", key, i)
		}
	}
}

func TestTimePriceSeriesDayInterval(t *testing.T) {
	timePriceSeriesDayIntervalRequest := TimePriceSeriesIntervalRequest{
		UserId:        userId,
		Exchange:      "NSE",
		TradingSymbol: "NIFTY",
		Interval:      "1d", // 5 minutes interval
		StartTime:     "09:15:00 20-04-2025",
		EndTime:       "15:29:00 23-04-2025",
	}
	timePriceSeriesDayInterval, err := TimePriceSeriesDayInterval(timePriceSeriesDayIntervalRequest)
	if err != nil {
		t.Error(err)
		return
	}
	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(timePriceSeriesDayInterval), &resp)
	assert.NoError(t, errJson, "JSON should unmarshal correctly")

	// Top-level checks
	assert.Equal(t, "success", resp["status"], "Status should be 'success'")
	_, msgExists := resp["message"]
	assert.True(t, msgExists, "Message key should be present")

	// Data checks
	data, ok := resp["data"].([]interface{})
	assert.True(t, ok, "Data should be an array")

	expectedKeys := map[string]bool{
		"time": true, "epochTime": true, "open": true,
		"high": true, "low": true, "close": true,
		"volume": true, "oi": true,
	}

	for i, item := range data {
		entry, ok := item.(map[string]interface{})
		assert.True(t, ok, "Each data[%d] should be a map", i)

		// No unexpected keys
		for key := range entry {
			assert.True(t, expectedKeys[key], "Unexpected key '%s' in data[%d]", key, i)
		}

		// All expected keys should be present
		for key := range expectedKeys {
			_, present := entry[key]
			assert.True(t, present, "Missing key '%s' in data[%d]", key, i)
		}
	}
}

func TestBrokerageCalculator(t *testing.T) {
	brokerageCalculatorRequest := BrokerageCalculatorRequest{
		UserId:          userId,
		Exchange:        "NSE",
		TradingSymbol:   "SAWACA",
		TransactionType: "B",
		Product:         "C",
		Quantity:        "1",
		Price:           "0.50",
		StrikePrice:     "0.00",
		InstName:        "EQ",
		LotSize:         "1",
	}

	brokerageCalculator, err := BrokerageCalculator(brokerageCalculatorRequest)
	if err != nil {
		t.Error(err)
		return
	}

	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(brokerageCalculator), &resp)
	assert.NoError(t, errJson, "JSON should unmarshal correctly")

	// Top-level checks
	assert.Equal(t, "success", resp["status"], "Status should be 'success'")
	_, msgExists := resp["message"]
	assert.True(t, msgExists, "Message key should be present")

	// Data checks
	data, ok := resp["data"].(map[string]interface{})
	assert.True(t, ok, "Data should be an object")

	// Expected keys
	expectedKeys := map[string]bool{
		"brokerage":       true,
		"exchange_charge": true,
		"gst":             true,
		"remarks":         true,
		"sebi_charge":     true,
		"stamp_duty":      true,
	}

	// Raise error for unexpected keys
	for key := range data {
		if !expectedKeys[key] {
			t.Errorf("Unexpected key '%s' found in data", key)
		}
	}

	// Raise error if expected keys are missing
	for key := range expectedKeys {
		if _, present := data[key]; !present {
			t.Errorf("Expected key '%s' missing in data", key)
		}
	}
}

func TestGetExpiry(t *testing.T) {
	getExpiryReq := GetInfoRequest{
		UserId:        userId,
		Exchange:      "NSE",
		TradingSymbol: "NIFTY",
	}

	getExpiryDetails, err := GetExpiry(getExpiryReq)
	if err != nil {
		t.Error(err)
		return
	}

	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(getExpiryDetails), &resp)
	assert.NoError(t, errJson, "JSON should unmarshal correctly")

	// Top-level checks
	assert.Equal(t, "success", resp["status"], "Status should be 'success'")
	_, msgExists := resp["message"]
	assert.True(t, msgExists, "Message key should be present")

	// Data checks
	data, ok := resp["data"].(map[string]interface{})
	assert.True(t, ok, "'data' should be an object")

	// Expected keys in data
	expectedKeys := map[string]bool{
		"expiryDates": true,
	}

	// Check for missing expected keys
	for key := range expectedKeys {
		if _, exists := data[key]; !exists {
			t.Errorf("Expected key '%s' missing in 'data'", key)
		}
	}

	// Fail test if unexpected keys are present
	for key := range data {
		if !expectedKeys[key] {
			t.Errorf("Unexpected key found in 'data': %s", key)
		}
	}

	// Validate expiryDates array
	expiryDates, ok := data["expiryDates"].([]interface{})
	assert.True(t, ok, "'expiryDates' should be an array")
	assert.NotEmpty(t, expiryDates, "'expiryDates' should not be empty")

	// Ensure all expiryDates are strings
	for i, v := range expiryDates {
		_, isString := v.(string)
		assert.Truef(t, isString, "expiryDates[%d] should be a string", i)
	}
}

func TestOrderBook(t *testing.T) {
	orderBookDetails, err := OrderBook(userId)
	if err != nil {
		t.Error(err)
		return
	}
	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(orderBookDetails), &resp)
	assert.NoError(t, errJson, "JSON should unmarshal correctly")

	// Top-level checks
	assert.Equal(t, "success", resp["status"], "Status should be 'success'")
	_, messageExists := resp["message"]
	assert.True(t, messageExists, "'message' key should be present")

	// Data must be an array
	data, ok := resp["data"].([]interface{})
	assert.True(t, ok, "'data' should be an array")
	assert.NotEmpty(t, data, "'data' array should not be empty")

	// Required keys and types
	expectedKeys := map[string]bool{
		"averagePrice":    true,
		"exchange":        true,
		"fillShares":      true,
		"lotSize":         true,
		"orderNumber":     true,
		"orderTime":       true,
		"price":           true,
		"priceType":       true,
		"product":         true,
		"quantity":        true,
		"rejectReason":    true,
		"remarks":         true,
		"retention":       true,
		"status":          true,
		"tickSize":        true,
		"token":           true,
		"tradingSymbol":   true,
		"transactionType": true,
		"userId":          true,
	}

	// Check each order's structure
	for i, item := range data {
		order, ok := item.(map[string]interface{})
		assert.True(t, ok, "Each order should be an object")
		for key := range order {
			if !expectedKeys[key] {
				t.Errorf("Unexpected key found in order[%d]: %s", i, key)
			}
		}
		for key := range expectedKeys {
			_, exists := order[key]
			assert.Truef(t, exists, "Key '%s' should exist in order[%d]", key, i)
		}
	}
}
func TestModifyOrder(t *testing.T) {
	modify_order := ModifyOrderRequest{
		UserId:         userId,
		OrderNumber:    "25060900005629",
		PriceType:      "MKT",
		TradingSymbol:  "SHALPRO",
		Price:          "",
		TriggerPrice:   "",
		Quantity:       "2",
		Product:        "C",
		Retention:      "DAY",
		Mkt_protection: "0.5",
	}
	modifyOrder, err := ModifyOrder(modify_order)
	if err != nil {
		t.Error(err)
		return
	}

	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(modifyOrder), &resp)
	assert.NoError(t, errJson, "JSON should unmarshal correctly")

	// Allowed keys at top-level
	expectedTopLevelKeys := map[string]bool{
		"status":  true,
		"message": true,
		"data":    true,
	}

	// Check for unexpected keys at top-level
	for key := range resp {
		if !expectedTopLevelKeys[key] {
			t.Errorf("Unexpected key at top-level: '%s'", key)
		}
	}

	// Top-level validation
	assert.Equal(t, "success", resp["status"], "Status should be 'success'")
	assert.Equal(t, "Order modification details", resp["message"], "Message should match expected value")

	// Data object validation
	data, ok := resp["data"].(map[string]interface{})
	assert.True(t, ok, "'data' should be a JSON object")

	// Allowed keys inside data
	expectedDataKeys := map[string]bool{
		"orderNumber": true,
		"requestTime": true,
	}

	// Check for unexpected keys inside data
	for key := range data {
		if !expectedDataKeys[key] {
			t.Errorf("Unexpected key in 'data': '%s'", key)
		}
	}

	// Check presence of expected keys in data
	_, orderExists := data["orderNumber"]
	_, timeExists := data["requestTime"]
	assert.True(t, orderExists, "'orderNumber' should be present in 'data'")
	assert.True(t, timeExists, "'requestTime' should be present in 'data'")
}

func TestCancelOrder(t *testing.T) {
	order_number := "25060900005545"

	cancel_order := OrderRequest{
		UserId:      userId,
		OrderNumber: order_number,
	}
	cancelOrder, err := CancelOrder(cancel_order)
	if err != nil {
		t.Error(err)
		return
	}

	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(cancelOrder), &resp)
	assert.NoError(t, errJson, "JSON should unmarshal correctly")

	// Top-level checks
	assert.Equal(t, "success", resp["status"], "Status should be 'success'")
	assert.Equal(t, "Order cancellation details", resp["message"], "Message should match expected value")

	// Data object checks
	data, ok := resp["data"].(map[string]interface{})
	assert.True(t, ok, "'data' should be a JSON object")

	// Expected keys
	expectedKeys := map[string]bool{
		"orderNumber": true,
		"rejreason":   true,
		"requestTime": true,
	}

	// Check for missing expected keys
	for key := range expectedKeys {
		if _, exists := data[key]; !exists {
			t.Errorf("Expected key '%s' missing in 'data'", key)
		}
	}

	// Fail test if unexpected keys are found
	for key := range data {
		if !expectedKeys[key] {
			t.Errorf("Unexpected key found in 'data': %s", key)
		}
	}
}

func TestSingleOrderHistory(t *testing.T) {
	order_number := "25060900004354"
	single_order_history := OrderRequest{
		UserId:      userId,
		OrderNumber: order_number,
	}
	singleOrderHistory, err := SingleOrderHistory(single_order_history)
	if err != nil {
		t.Error(err)
		return
	}

	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(singleOrderHistory), &resp)
	assert.NoError(t, errJson, "JSON should unmarshal properly")

	// Check top-level fields
	assert.Equal(t, "success", resp["status"], "Status should be 'success'")
	assert.Equal(t, "Order history data received", resp["message"], "Message should match")

	// Validate 'data' as an array
	data, ok := resp["data"].([]interface{})
	assert.True(t, ok, "'data' should be a JSON array")
	assert.NotEmpty(t, data, "'data' should not be empty")

	// Validate each item in 'data' has expected fields
	expectedFields := []string{
		"averagePrice", "exchange", "exchangeOrderNum", "exchangeTime", "fillShares",
		"orderNumber", "orderTime", "price", "priceType", "product", "quantity",
		"rejectReason", "remarks", "reportType", "retention", "status", "tickSize",
		"token", "tradingSymbol", "transactionType", "userId",
	}

	for i, item := range data {
		order, ok := item.(map[string]interface{})
		assert.True(t, ok, "Each order item should be an object")
		for _, field := range expectedFields {
			_, exists := order[field]
			assert.True(t, exists, "Order #%d should contain key: %s", i+1, field)
		}
		for key := range order {
			found := false
			for _, field := range expectedFields {
				if key == field {
					found = true
					break
				}
			}
			if !found {
				t.Error("Unexpected key found in position #", i+1, ":", key)
			}
		}
	}
}

func TestTradeBook(t *testing.T) {
	tradeBook, err := TradeBook(userId)
	if err != nil {
		t.Error(err)
		return
	}
	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(tradeBook), &resp)
	assert.NoError(t, errJson, "JSON should unmarshal properly")

	// Check top-level fields
	assert.Equal(t, "success", resp["status"], "Status should be 'success'")
	assert.Equal(t, "Trade history data received", resp["message"], "Message should match")

	// Validate 'data' is a non-empty array
	data, ok := resp["data"].([]interface{})
	assert.True(t, ok, "'data' should be a JSON array")
	assert.NotEmpty(t, data, "'data' should not be empty")

	// Validate each item has expected fields
	expectedFields := []string{
		"exchange", "exchangeUpdateTime", "exchordid", "fillId", "fillPrice", "fillQuantity",
		"fillTime", "fillshares", "lotSize", "orderNumber", "orderTime", "priceFactor",
		"pricePrecision", "priceType", "product", "quantity", "retention", "tickSize",
		"token", "tradingSymbol", "transactionType", "userId",
	}

	for i, item := range data {
		trade, ok := item.(map[string]interface{})
		assert.True(t, ok, "Each trade item should be an object")
		for _, field := range expectedFields {
			_, exists := trade[field]
			assert.True(t, exists, "Trade #%d should contain key: %s", i+1, field)
		}

		for key := range trade {
			found := false
			for _, field := range expectedFields {
				if key == field {
					found = true
					break
				}
			}
			if !found {
				t.Error("Unexpected key found in position #", i+1, ":", key)
			}
		}
	}
}

func TestPositionBook(t *testing.T) {
	positionBookDetails, err := PositionBook(userId)
	if err != nil {
		t.Error(err)
		return
	}

	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(positionBookDetails), &resp)
	assert.NoError(t, errJson, "JSON should unmarshal properly")

	// Check top-level fields
	assert.Equal(t, "success", resp["status"], "Status should be 'success'")
	assert.Equal(t, "Positions retrieved successfully", resp["message"], "Message should match")

	// Validate 'data' is a non-empty array
	data, ok := resp["data"].([]interface{})
	assert.True(t, ok, "'data' should be a JSON array")
	assert.NotEmpty(t, data, "'data' should not be empty")

	// Expected keys in each position
	expectedFields := []string{
		"RealizedPNL", "dayBuyAmount", "dayBuyAveragePrice", "dayBuyQuantity",
		"daySellAmount", "daySellAveragePrice", "daySellQuantity", "exchange",
		"lastTradedPrice", "lotSize", "netAveragePrice", "netQuantity",
		"netUploadPrice", "product", "tickSize", "token", "totalMTM",
		"totalPNL", "tradingSymbol", "uploadPrice", "userId",
	}

	// Validate keys in each position object

	for i, item := range data {
		pos, ok := item.(map[string]interface{})
		assert.True(t, ok, "Each position item should be an object")
		for _, field := range expectedFields {
			_, exists := pos[field]
			assert.True(t, exists, "Position #%d should contain key: %s", i+1, field)
		}

		for key := range pos {
			found := false
			for _, field := range expectedFields {
				if key == field {
					found = true
					break
				}
			}
			if !found {
				t.Error("Unexpected key found in position #", i+1, ":", key)
			}
		}
	}

}

func TestOrderMargin(t *testing.T) {
	exchange, product, priceType, tradingSymbol, transactionType, price, quantity := "BSE", "C", "MKT", "SAWACA", "B", "0.50", "1"
	orderMarginRequest := OrderMarginRequest{
		UserId:          userId,
		Exchange:        exchange,
		TransactionType: transactionType,
		Product:         product,
		TradingSymbol:   tradingSymbol,
		Quantity:        quantity,
		PriceType:       priceType,
		Price:           price,
	}
	orderMargin, err := OrderMargin(orderMarginRequest)
	if err != nil {
		t.Error(err)
		return
	}
	// Parse into map
	var result map[string]interface{}
	errJson := json.Unmarshal([]byte(orderMargin), &result)
	assert.NoError(t, errJson, "JSON unmarshalling failed")

	// Check top-level status
	assert.Equal(t, "success", result["status"], "status should be 'success'")

	// Check data key exists and is a map
	data, ok := result["data"].(map[string]interface{})
	assert.True(t, ok, "'data' should be a JSON object")

	// Expected keys in the data
	expectedKeys := map[string]bool{
		"availableMargin":  true,
		"cash":             true,
		"marginOnNewOrder": true,
		"remarks":          true,
		"requestTime":      true,
	}

	// Check each expected key is present
	for key := range expectedKeys {
		_, exists := data[key]
		assert.True(t, exists, "Key '"+key+"' should be present in data")
	}

	// Fail if unexpected keys are found
	for key := range data {
		if !expectedKeys[key] {
			t.Errorf("Unexpected key found in data: %s", key)
		}
	}
}

func TestProductConversion(t *testing.T) {
	productConversionRequest := ProductConversionRequest{
		UserId:          userId,
		TradingSymbol:   "AVANCE",
		Exchange:        "BSE",
		PreviousProduct: "I", // B = Buy, S = Sell
		Product:         "C", // C = Delivery, I = Intraday, M = Margin Intraday (MIS)
		Quantity:        "1", // As string
	}

	productConversion, err := ProductConversion(productConversionRequest)
	if err != nil {
		t.Error(err)
		return
	}

	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(productConversion), &resp)
	assert.NoError(t, errJson, "JSON should unmarshal correctly")

	// Expected top-level keys
	expectedTopKeys := map[string]bool{
		"status":  true,
		"message": true,
		"data":    true,
	}

	// Check for unexpected top-level keys
	for key := range resp {
		if !expectedTopKeys[key] {
			t.Errorf("Unexpected top-level key: %s", key)
		}
	}

	// Check for missing top-level keys
	for key := range expectedTopKeys {
		if _, ok := resp[key]; !ok {
			t.Errorf("Missing expected top-level key: %s", key)
		}
	}

	assert.Equal(t, "success", resp["status"], "Status should be 'success'")
	assert.Equal(t, "Partial position conversion info fetched", resp["message"], "Message should match")

	// Validate 'data' is a map
	data, ok := resp["data"].(map[string]interface{})
	assert.True(t, ok, "'data' should be a JSON object")

	// Expected keys in data
	expectedDataKeys := map[string]bool{
		"Status":      true,
		"requestTime": true,
	}

	// Check for unexpected keys in data
	for key := range data {
		if !expectedDataKeys[key] {
			t.Errorf("Unexpected key in data: %s", key)
		}
	}

	// Check for missing expected keys in data
	for key := range expectedDataKeys {
		if _, ok := data[key]; !ok {
			t.Errorf("Missing expected key in data: %s", key)
		}
	}
}

func TestLogout(t *testing.T) {
	logout, err := Logout(userId)
	if err != nil {
		t.Error(err)
		return
	}

	var resp map[string]interface{}
	errJson := json.Unmarshal([]byte(logout), &resp)
	assert.NoError(t, errJson, "JSON should unmarshal correctly")

	// Define expected top-level keys
	expectedKeys := map[string]bool{
		"status":  true,
		"message": true,
	}

	// Check that all expected keys are present
	for key := range expectedKeys {
		if _, exists := resp[key]; !exists {
			t.Errorf("Missing expected key: '%s'", key)
		}
	}

	// Check for unexpected top-level keys
	for key := range resp {
		if !expectedKeys[key] {
			t.Errorf("Unexpected key at top-level: '%s'", key)
		}
	}

	// Assert expected values
	assert.Equal(t, "success", resp["status"], "Status should be 'success'")
}
