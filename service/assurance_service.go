package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type AssuranceService interface {
	GetAllAssurances() ([]AssuranceInfo, error)
	GetAllAssuranceTypes() ([]AssuranceType, error)
	DeleteAssuranceByID(assuranceID string) (*AssuranceResponse, error)
	DeleteAssuranceByOrderID(orderID string) (*AssuranceResponse, error)
	ModifyAssurance(assuranceID, orderID string, typeIndex int) (*AssuranceResponse, error)
	CreateNewAssurance(typeIndex int, orderID string) (*AssuranceResponse, error)
	GetAssuranceByID(assuranceID string) (*AssuranceInfo, error)
	FindAssuranceByOrderID(orderID string) ([]AssuranceInfo, error)
}

type AssuranceInfo struct {
	AssuranceID string `json:"assuranceId"`
	OrderID     string `json:"orderId"`
}

type AssuranceType struct {
	TypeID   int    `json:"typeId"`
	TypeName string `json:"typeName"`
}

type AssuranceResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		AssuranceID string `json:"assuranceId"`
	} `json:"data"`
}

type Modify_Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type GetAllResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id        string  `json:"id"`
		OrderId   string  `json:"orderId"`
		TypeIndex int     `json:"typeIndex"`
		TypeName  string  `json:"typeName"`
		TypePrice float64 `json:"typePrice"`
	} `json:"data"`
}

type AssuranceDeleteResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type GetAssuranceByIDeInfo struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id      string `json:"id"`
		OrderId string `json:"orderId"`
		Type    string `json:"type"`
	} `json:"data"`
}

func (s *SvcImpl) GetAllAssurances() (*GetAllResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/assuranceservice/assurances", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GetAllResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) GetAllAssuranceTypes() ([]AssuranceType, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/assuranceservice/assurances/types", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var types []AssuranceType
	err = json.NewDecoder(resp.Body).Decode(&types)
	if err != nil {
		return nil, err
	}

	return types, nil
}

func (s *SvcImpl) DeleteAssuranceByID(assuranceID string) (*AssuranceDeleteResponse, error) {
	url := fmt.Sprintf("%s/api/v1/assuranceservice/assurances/assuranceid/%s", s.BaseUrl, assuranceID)
	resp, err := s.cli.SendRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AssuranceDeleteResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) DeleteAssuranceByOrderID(orderID string) (*AssuranceResponse, error) {
	url := fmt.Sprintf("%s/api/v1/assuranceservice/assurances/orderid/%s", s.BaseUrl, orderID)
	resp, err := s.cli.SendRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AssuranceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) ModifyAssurance(assuranceID string, orderID string, typeIndex int) (*Modify_Response, error) {
	url := fmt.Sprintf("%s/api/v1/assuranceservice/assurances/%s/%s/%d", s.BaseUrl, assuranceID, orderID, typeIndex)
	resp, err := s.cli.SendRequest("PATCH", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Modify_Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) CreateNewAssurance(typeIndex int, orderID string) (*AssuranceResponse, error) {
	url := fmt.Sprintf("%s/api/v1/assuranceservice/assurances/%d/%s", s.BaseUrl, typeIndex, orderID)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AssuranceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) GetAssuranceByID(assuranceID string) (*GetAssuranceByIDeInfo, error) {
	url := fmt.Sprintf("%s/api/v1/assuranceservice/assurances/assuranceid/%s", s.BaseUrl, assuranceID)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var assurance GetAssuranceByIDeInfo
	err = json.Unmarshal(body, &assurance)
	if err != nil {
		return nil, err
	}

	return &assurance, nil
}

func (s *SvcImpl) FindAssuranceByOrderID(orderId string) (*GetAssuranceByIDeInfo, error) {
	url := fmt.Sprintf("%s/api/v1/assuranceservice/assurances/orderid/%s", s.BaseUrl, orderId)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var assurance GetAssuranceByIDeInfo
	err = json.Unmarshal(body, &assurance)
	if err != nil {
		return nil, err
	}

	return &assurance, nil
}
