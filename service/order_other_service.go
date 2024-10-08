package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type OrderOtherService interface {
	ReqFindAllOrderOther() (*OrderArrResp, error)
	ReqCreateNewOrderOther(input *Order) (*OrderResp, error)
	ReqSaveOrderInfoOther(input *Order) (*OrderResp, error)
	ReqAddCreateNewOrderOther(input *Order) (*OrderResp, error)
	ReqUpdateOrderOrderServiceOther(input *Order) (*OrderResp, error)
	ReqPayOrderOther(orderId string) (*OrderResp, error)
	ReqGetOrderPriceOther(orderId string) (*GetOrderPriceResp, error)
	ReqQueryOrdersOther(input *OrderInfo) (*OrderArrResp, error)
	ReqQueryOrderForRefreshOther(input *OrderInfo) (*OrderArrResp, error)
	ReqSecurityInfoCheckOther(checkDate string, accountId string) (*OrderSecurityResp, error)
	ReqModifyOrderOther(orderId string, status int) (*OrderResp, error)
	ReqGetTicketsListOther(input *Seat) (*TicketResp, error)
	ReqDeleteOrderOrderServiceOther(orderId string) (*OrderResp, error)
	ReqGetOrderByIdOther(orderId string) (*OrderResp, error)
	ReqCalculateSoldTicketOther(travelDate string, travelNumber string) (*OrderResp, error)
}

func (s *SvcImpl) ReqFindAllOrderOther() (*OrderArrResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderArrResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqCreateNewOrderOther(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderOtherService/orderOther", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqSaveOrderInfoOther(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/orderOtherService/orderOther", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqAddCreateNewOrderOther(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderOtherService/orderOther/admin", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqUpdateOrderOrderServiceOther(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/orderOtherService/orderOther/admin", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqPayOrderOther(orderId string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther/orderpay/"+orderId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetOrderPriceOther(orderId string) (*GetOrderPriceResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther/price/"+orderId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result GetOrderPriceResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqQueryOrdersOther(input *OrderInfo) (*OrderArrResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderOtherService/orderOther/query", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderArrResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqQueryOrderForRefreshOther(input *OrderInfo) (*OrderArrResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderOtherService/orderOther/refresh", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderArrResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqSecurityInfoCheckOther(checkDate string, accountId string) (*OrderSecurityResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther/security/"+checkDate+"/"+accountId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderSecurityResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqModifyOrderOther(orderId string, status int) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther/status/"+orderId+"/"+strconv.Itoa(status), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetTicketsListOther(input *Seat) (*TicketResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderOtherService/orderOther/tickets", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TicketResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqDeleteOrderOrderServiceOther(orderId string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+"/api/v1/orderOtherService/orderOther/"+orderId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetOrderByIdOther(orderId string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther/"+orderId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqCalculateSoldTicketOther(travelDate string, travelNumber string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther/"+travelDate+"/"+travelNumber, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
