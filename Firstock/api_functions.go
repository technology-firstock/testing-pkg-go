// Copyright (c) [2025] [abc]
// SPDX-License-Identifier: MIT
package Firstock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type apifunctions struct{}

func (fs *apifunctions) LoginFunction(
	reqBody LoginRequest) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(reqBody)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(login_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var loginResp map[string]interface{}
	if err := json.Unmarshal(body, &loginResp); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return loginResp, nil
}

func (fs *apifunctions) LogoutFunction(
	reqBody LogoutRequest) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(reqBody)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(logout_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var logoutResp map[string]interface{}
	if err := json.Unmarshal(body, &logoutResp); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return logoutResp, nil
}

func (fs *apifunctions) UserDetailsFunction(reqBody UserDetailsRequest) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(reqBody)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(user_details_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var userDetailsResponse map[string]interface{}
	if err := json.Unmarshal(body, &userDetailsResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return userDetailsResponse, nil

}

func (fs *apifunctions) PlaceOrderFunction(req PlaceOrderRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(place_order_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var placeOrderResponse map[string]interface{}
	if err := json.Unmarshal(body, &placeOrderResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return placeOrderResponse, nil
}

func (fs *apifunctions) OrderMarginFunction(req OrderMarginRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(order_margin_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var orderMarginResponse map[string]interface{}
	if err := json.Unmarshal(body, &orderMarginResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return orderMarginResponse, nil
}

func (fs *apifunctions) SingleOrderHistoryFunction(req OrderRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(single_order_history_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var singleOrderHistoryResponse map[string]interface{}
	if err := json.Unmarshal(body, &singleOrderHistoryResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return singleOrderHistoryResponse, nil
}

func (fs *apifunctions) CancelOrderFunction(req OrderRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(cancel_order_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var cancelOrderResponse map[string]interface{}
	if err := json.Unmarshal(body, &cancelOrderResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return cancelOrderResponse, nil
}

func (fs *apifunctions) ModifyOrderFunction(req ModifyOrderRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(modify_order_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var modifyOrderResponse map[string]interface{}
	if err := json.Unmarshal(body, &modifyOrderResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return modifyOrderResponse, nil
}

func (fs *apifunctions) TradeBookFunction(req BaseRequest) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(trade_book_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var tradeBookResponse map[string]interface{}
	if err := json.Unmarshal(body, &tradeBookResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return tradeBookResponse, nil
}

func (fs *apifunctions) RmsLimitFunction(req BaseRequest) (map[string]interface{}, error) {
	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(rms_limit_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var rmsLimitResponse map[string]interface{}
	if err := json.Unmarshal(body, &rmsLimitResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return rmsLimitResponse, nil
}

func (fs *apifunctions) PositionBookFunction(req BaseRequest) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(position_book_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var positionBookResponse map[string]interface{}
	if err := json.Unmarshal(body, &positionBookResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return positionBookResponse, nil
}

func (fs *apifunctions) HoldingsFunction(req BaseRequest) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(holdings_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var holdingsResponse map[string]interface{}
	if err := json.Unmarshal(body, &holdingsResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return holdingsResponse, nil
}

func (fs *apifunctions) OrderBookFunction(req BaseRequest) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(order_book_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var orderBookResponse map[string]interface{}
	if err := json.Unmarshal(body, &orderBookResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return orderBookResponse, nil
}

func (fs *apifunctions) GetExpiryFunction(req GetInfoRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(get_expiry_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var getExpiryResponse map[string]interface{}
	if err := json.Unmarshal(body, &getExpiryResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return getExpiryResponse, nil
}

func (fs *apifunctions) BrokerageCalculatorFunction(req BrokerageCalculatorRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(brokerage_calculator_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var brokerageCalculatorResponse map[string]interface{}
	if err := json.Unmarshal(body, &brokerageCalculatorResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return brokerageCalculatorResponse, nil
}

func (fs *apifunctions) BasketMarginFunction(req BasketMarginRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(basket_margin_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var basketMarginResponse map[string]interface{}
	if err := json.Unmarshal(body, &basketMarginResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return basketMarginResponse, nil
}

func (fs *apifunctions) GetSecurityInfoFunction(req GetInfoRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(get_security_info, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var getSecurityInfoResponse map[string]interface{}
	if err := json.Unmarshal(body, &getSecurityInfoResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return getSecurityInfoResponse, nil
}

func (fs *apifunctions) ProductConversionFunction(req ProductConversionRequestBody) (map[string]interface{}, error) {
	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(product_conversion_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var productConversionResponse map[string]interface{}
	if err := json.Unmarshal(body, &productConversionResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return productConversionResponse, nil
}

// ---------------------------------------Connect---------------------------------
func (fs *apifunctions) GetQuoteFunction(req GetInfoRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(get_quote_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var getQuoteResponse map[string]interface{}
	if err := json.Unmarshal(body, &getQuoteResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return getQuoteResponse, nil
}

func (fs *apifunctions) GetQuoteLtpFunction(req GetInfoRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(get_quote_ltp_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var getQuoteLtpResponse map[string]interface{}
	if err := json.Unmarshal(body, &getQuoteLtpResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return getQuoteLtpResponse, nil
}

func (fs *apifunctions) GetMultiQuotesFunction(req GetMultiQuotesRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(get_multi_quotes_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var getMultiQuotesResponse map[string]interface{}
	if err := json.Unmarshal(body, &getMultiQuotesResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return getMultiQuotesResponse, nil
}

func (fs *apifunctions) GetMultiQuotesLtpFunction(req GetMultiQuotesRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(get_multi_quotes_ltp_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var getMultiQuotesResponse map[string]interface{}
	if err := json.Unmarshal(body, &getMultiQuotesResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return getMultiQuotesResponse, nil
}

func (fs *apifunctions) IndexListFunction(req BaseRequest) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(index_list_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var indexListResponse map[string]interface{}
	if err := json.Unmarshal(body, &indexListResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return indexListResponse, nil
}

func (fs *apifunctions) SearchScripsFunction(req SearchScripsBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(search_scrips_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var searchScripsResponse map[string]interface{}
	if err := json.Unmarshal(body, &searchScripsResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return searchScripsResponse, nil
}

func (fs *apifunctions) OptionChainFunction(req OptionChainRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(option_chain_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var optionChainResponse map[string]interface{}
	if err := json.Unmarshal(body, &optionChainResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return optionChainResponse, nil
}

func (fs *apifunctions) TimePriceSeriesRegularIntervalFunction(req TimePriceSeriesIntervalRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(time_price_series_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var timePriceSeriesResponse map[string]interface{}
	if err := json.Unmarshal(body, &timePriceSeriesResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return timePriceSeriesResponse, nil
}

func (fs *apifunctions) TimePriceSeriesDayIntervalFunction(req TimePriceSeriesIntervalRequestBody) (map[string]interface{}, error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(time_price_series_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var timePriceSeriesResponse map[string]interface{}
	if err := json.Unmarshal(body, &timePriceSeriesResponse); err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return timePriceSeriesResponse, nil
}
