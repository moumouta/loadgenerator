package service

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"strconv"
	"testing"
)

func TestSvcImpl_ReqGetAllOrders(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, err := cli.ReqGetAllOrders()
	if err != nil {
		t.Errorf("err reponse: %v", err)
	}
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqAddOrder(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqAddOrder(&Order{
		AccountId:              uuid.NewString(),
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
		DifferenceMoney:        RandomDecimalStringBetween(1, 10),
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     uuid.NewString(),
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             GenerateSeatNumber(),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            "G111",
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}

func TestSvcImpl_ReqUpdateOrder(t *testing.T) {
	cli, _ := GetAdminClient()
	UpdateResp, err := cli.ReqUpdateOrder(&Order{
		AccountId:              "test1",
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
		DifferenceMoney:        RandomDecimalStringBetween(1, 10),
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     "790bcfd5-82d2-4717-aa9f-e00bef992268",
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             GenerateSeatNumber(),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            "G111",
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(UpdateResp.Msg)
}

func TestSvcImpl_ReqDeleteOrder(t *testing.T) {
	cli, _ := GetAdminClient()
	DeleteResp, err := cli.ReqDeleteOrder("790bcfd5-82d2-4717-aa9f-e00bef992268", "G111")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(DeleteResp.Msg)
}

func TestSvcImpl_End2End(t *testing.T) {
	cli, _ := GetAdminClient()
	newOrder := Order{
		AccountId:              uuid.New().String(),
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
		DifferenceMoney:        RandomDecimalStringBetween(1, 10),
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     uuid.New().String(),
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             GenerateSeatNumber(),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            GenerateTrainNumber(),
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	}
	AddResp, err := cli.ReqAddOrder(&newOrder)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(AddResp.Msg)
	oldOrder := AddResp.Data
	UpdateResp, err := cli.ReqUpdateOrder(&oldOrder)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(UpdateResp.Msg)
	DeleteResp, err := cli.ReqDeleteOrder(oldOrder.Id, oldOrder.TrainNumber)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(DeleteResp.Data.TrainNumber)
	fmt.Println(DeleteResp.Msg)
}
