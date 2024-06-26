package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type StationFoodService interface {
	GetAllStationFood() (*GetStationFoodResp, error)
	GetStationFoodByName(stationName string) (*GetStationFoodResp, error)
	GetStationFoodByNames(stationNames []string) (*GetStationFoodResp, error)
	GetStationFoodById(storeId string) (*GetStationFoodSingleResp, error)
}
type GetStationFoodResp struct {
	Status int           `json:"status"`
	Msg    string        `json:"msg"`
	Data   []StationFood `json:"data"`
}

type GetStationFoodSingleResp struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   StationFood `json:"data"`
}

type StationFood struct {
	Id           string  `json:"id"`
	StationName  string  `json:"stationName"`
	StoreName    string  `json:"storeName"`
	Telephone    string  `json:"telephone"`
	BusinessTime string  `json:"businessTime"`
	DeliveryFee  float64 `json:"deliveryFee"`
	FoodList     []struct {
		FoodName string  `json:"foodName"`
		Price    float64 `json:"price"`
	} `json:"foodList"`
}

func (s *SvcImpl) GetAllStationFood() (*GetStationFoodResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/stationfoodservice/stationfoodstores", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result GetStationFoodResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
func (s *SvcImpl) GetStationFoodByName(stationName string) (*GetStationFoodResp, error) {
	resp, err := s.cli.SendRequest("GET", fmt.Sprintf("%s/api/v1/stationfoodservice/stationfoodstores/%s", s.BaseUrl, stationName), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result GetStationFoodResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
func (s *SvcImpl) GetStationFoodByNames(stationNames []string) (*GetStationFoodResp, error) {
	resp, err := s.cli.SendRequest("POST", fmt.Sprintf("%s/api/v1/stationfoodservice/stationfoodstores", s.BaseUrl), stationNames)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result GetStationFoodResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) GetStationFoodById(storeId string) (*GetStationFoodSingleResp, error) {
	resp, err := s.cli.SendRequest("GET", fmt.Sprintf("%s/api/v1/stationfoodservice/stationfoodstores/bystoreid/%s", s.BaseUrl, storeId), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result GetStationFoodSingleResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
