package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	"log"
	"math/rand"
	"strings"
	"time"
)

const (
	// Preserve - Main
	AccountID  = "accountId"
	ContactsID = "contactsId"
	TripID     = "tripId"
	SeatType   = "seatType"
	//LoginToken = "loginToken"
	Date            = "date"
	From            = "from"
	To              = "to"
	Assurance       = "assurance"
	FoodType        = "foodType"
	StationName     = "stationName"
	StoreName       = "storeName"
	FoodName        = "foodName"
	FoodPrice       = "foodPrice"
	HandleDate      = "handleDate"
	ConsigneeName   = "consigneeName"
	ConsigneePhone  = "consigneePhone"
	ConsigneeWeight = "consigneeWeight"
	IsWithin        = "isWithin"

	// Assurance
	OrderId            = "orderId"
	AssuranceTypeIndex = "typeIndex"
	AssuranceTypeName  = "typeName"
	AssuranceTypePrice = "typePrice"

	// VerifyCode
	BooleanVerifyCode = "booleanVerifyCode"

	// User
	UserID       = "userId"
	UserName     = "userName"
	Password     = "password"
	Gender       = "gender"
	DocumentType = "documentType"
	DocumentNum  = "documentNum"
	Email        = "email"

	// Contacts
	//Id = "id" - ContactsID
	//AccountId = "accountId"
	Name = "name"
	//DocumentType   = "documentType"
	DocumentNumber = "documentNumber"
	PhoneNumber    = "phoneNumber"

	// Consign
	ID      = "id"
	OrderID = "orderId"
	//AccountID = "accountId"
	//HandleDate = "handleDate"
	TargetDate = "targetDate"
	//From = "from"
	//To = "to"
	Consignee = "consignee"
	Phone     = "phone"
	Weight    = "weight"
	//IsWithin = "isWithin"
	Price = "price"

	// FoodBehavior

	// Trip(Travel)
	TripId        = "tripId"
	TrainTypeName = "trainTypeName"
	//StartStation         = "startStation"
	TerminalStation = "terminalStation"
	StartTime       = "startTime"
	EndTime         = "endTime"
	//EconomyClass         = "economyClass"
	//ConfortClass         = "confortClass"
	PriceForEconomyClass = "priceForEconomyClass"
	PriceForConfortClass = "priceForConfortClass"

	// Train
	//Id           = "id" //Train-ID needed or not?
	//Name         = "name" //Train-Name needed or not?
	TrainTypName = "trainTypeName"
	ConfortClass = "confortClass"
	AverageSpeed = "averageSpeed"
	EconomyClass = "economyClass"

	// Route
	RouteID      = "routeId"
	Stations     = "stations"
	Distances    = "distances"
	StartStation = "startStation"
	EndStation   = "endStation"

	// Config
	ConfigName  = "name"
	Value       = "value"
	Description = "description"

	// Order
	//AccountId              = "accountId"
	BoughtDate             = "boughtDate"
	CoachNumber            = "coachNumber"
	ContactsDocumentNumber = "contactsDocumentNumber"
	//ContactsName           = "contactsName"
	DifferenceMoney = "differenceMoney"
	//DocumentType           = "documentType"
	//From                   = "from"
	//Id                     = "id"
	//Price                  = "price"
	SeatClass  = "seatClass"
	SeatNumber = "seatNumber"
	Status     = "status"
	//To                     = "to"
	TrainNumber = "trainNumber"
	TravelDate  = "travelDate"
	TravelTime  = "travelTime"

	// Price
	//Id                  = "id"
	TrainType           = "trainType"
	RouteId             = "routeId"
	BasicPriceRate      = "basicPriceRate"
	FirstClassPriceRate = "firstClassPriceRate"

	// Basic - Query
	//Status = "status"
	Percent = "percent"
	//TrainType = "trainType"
	//TrainType struct {
	//Id = "id"
	//Name = "name"
	//EconomyClass = "economyClass"
	//ConfortClass = "confortClass"
	//AverageSpeed = "averageSpeed"
	//} `json:"trainType"`
	Route = "route"
	//Route struct {
	//RouteID = "id"
	//Stations = "stations"
	//Distances = "distances"
	//StartStation = "startStation"
	//EndStation = "endStation"
	//} `json:"route"`
	Prices = "prices"
	//Prices struct {
	//ConfortClass = "confortClass"
	//EconomyClass = "economyClass"
	//} `json:"prices"`

	// Security
	SecurityID          = "id"
	SecurityName        = "name"
	SecurityValue       = "value"
	SecurityDescription = "description"

	// Station
	StationId    = "id"
	StationNames = "name"
	StayTime     = "stayTime"

	// Basic
	DepartureTime = "departureTime"

	// Seat
	SeatNo = "seatNo"
	//StartStation = "startStation"
	DestStation = "destStation"
)

var PreserveBehaviorChain *Chain

func init() {
	// ------------------------------------- init -------------------------------------------
	// ------------------------------------- init -------------------------------------------
	// ------------------------------------- init -------------------------------------------
	// Main Chain
	PreserveBehaviorChain = NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("PreserveBehaviors(Chain) Statrs. Starts time: %v", time.Now().String())
		return nil, nil
	}, "Dummy"))
	// AssuranceBehaviorChain -
	AssuranceBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("CreateAssuranceChain. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyAssurance"))
	// UserBehaviorsChain
	UserBehaviorsChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("UserBehaviorsChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyUserBehaviors"))
	// VerifyCodeBehaviorChain
	VerifyCodeBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("VerifyCodeBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyVerifyCodeBehavior"))
	// AuthBehaviorChain
	AuthBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("AuthBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyAuthBehavior"))
	// UserBehaviorChain
	UserBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("UserBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyUserBehavior"))
	// ContactsBehaviorChain
	ContactsBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("ContactsBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyContactsBehavior"))
	// ConsignBehaviorsChain
	ConsignBehaviorsChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("ConsignBehaviorsChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyConsignBehaviors"))
	//ConsignPriceBehaviorChain
	ConsignPriceBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("ConsignPriceBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyConsignPriceBehavior"))
	// FoodBehaviorChain
	FoodBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("FoodBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyFoodBehavior"))
	// TravelBehaviorChain
	TravelBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("TravelBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTravelBehavior"))
	TravelBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("TravelBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTravelBehavior"))
	// StationFoodBehaviorChain
	StationFoodBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("StationFoodBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationFoodBehavior"))
	//TrainFoodBehaviorChain
	TrainFoodBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("TrainFoodBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTrainFoodBehavior"))
	// TrainBehaviorChain
	TrainBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("TrainBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTrainBehavior"))
	TrainBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("TrainBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTrainBehavior"))
	//RouteBehaviorChain
	RouteBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("RouteBehaviorChain starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyRouteBehavior"))
	RouteBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("RouteBehaviorChain starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyRouteBehavior"))
	// BasicBehaviorChain
	BasicBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("BasicBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyBasicBehavior"))
	// SeatBehaviorChain
	SeatBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("SeatBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummySeatBehavior"))
	// StationBehaviorChain
	StationBehaviorChain0 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	StationBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	StationBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	StationBehaviorChain3 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	//PriceBehaviorChain
	PriceBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("PriceBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyPriceBehavior"))
	//ConfigBehaviorChain
	ConfigBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("ConfigBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyConfigBehavior"))
	//OrderBehaviorChain
	OrderBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("OrderBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderBehavior"))
	OrderBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("OrderBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderBehavior"))
	OrderBehaviorChain3 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("OrderBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderBehavior"))
	//OrderOtherBehaviorChain
	OrderOtherBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("OrderOtherBehaviorChain Starts. Strat time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderOtherBehavior"))
	OrderOtherBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("OrderOtherBehaviorChain Starts. Strat time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderOtherBehavior"))
	// SecurityBehaviorChain
	SecurityBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("SecurityBehaviorChain Satrts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummySecurityBehavior"))

	// ------------------------------------- NewFuncNode -------------------------------------------
	// ------------------------------------- NewFuncNode -------------------------------------------
	// ------------------------------------- NewFuncNode -------------------------------------------
	//AssuranceBehaviorChain - Assurance
	QueryAssuranceNode := NewFuncNode(QueryAssurance, "QueryAssurance")
	// CreateAssuranceNode := NewFuncNode(CreateAssurance, "CreateAssurance")

	//UserBehaviorsChain
	// AuthBehaviorChain - LoginAdmin/LoginBasic
	//LoginAdminNode := NewFuncNode(LoginAdmin, "LoginAdmin")
	LoginBasicNode := NewFuncNode(LoginBasic, "LoginBasic")
	//	VerifyCodeBehaviorChain
	VerifyCodeNode := NewFuncNode(VerifyCode, "VerifyCode")
	// UserBehaviorChain
	QueryUserNode := NewFuncNode(QueryUser, "QueryUser")

	//ContactsBehaviorChain - Contacts
	QueryContactsNode := NewFuncNode(QueryContacts, "QueryContacts")
	CreateContactsNode := NewFuncNode(CreateContacts, "CreateContacts")

	//ConsignBehaviorsChain
	//QueryConsignNode := NewFuncNode(QueryConsign, "QueryConsign")
	CreateConsignNode := NewFuncNode(CreateConsign, "CreateConsign")
	// ConsignPriceBehaviorChain
	//QueryConsignPriceNode := NewFuncNode(QueryConsignPric, "QueryConsignPrice")
	//CreateConsignPriceNode := NewFuncNode(CreateConsignPrice, "CreateConsignPrice")

	//FoodBehaviorChain
	QueryFoodNode := NewFuncNode(QueryFood, "QueryFood")
	// StationFoodBehaviorChain
	QueryStationFoodNode := NewFuncNode(QueryStationFood, "QueryStationFood")
	// TrainFoodBehaviorChain
	QueryTrainFoodNode := NewFuncNode(QueryTrainFood, "QueryTrainFood")
	// TravelBehaviorChain
	QueryTravelNode := NewFuncNode(QueryTrip, "QueryTrip")
	//CreateTravelNode := NewFuncNode(CreateTrip, "CreateTrip")

	//TravelBehaviorChain
	// TrainBehaviorChain
	QueryTrainNode := NewFuncNode(QueryTrain, "QueryTrain")
	// RouteBehaviorChain
	QueryRouteNode := NewFuncNode(QueryRoute, "QueryRoute")
	// BasicBehaviorChain
	QueryBasicNode := NewFuncNode(QueryBasic, "QueryBasic")
	// SeatBehaviorChain
	CreateSeatNode := NewFuncNode(CreateSeat, "QuerySeat")

	//BasicBehaviorChain
	// StationBehaviorChain
	QueryStationNode := NewFuncNode(QueryStation, "QueryStation")
	// PriceBehaviorChain
	QueryPriceNode := NewFuncNode(QueryPrice, "QueryPrice")
	// RouteBehaviorChain
	// TrainBehaviorChain

	//SeatBehaviorChain
	// ConfigBehaviorChain
	QueryConfigNode := NewFuncNode(QueryConfig, "QueryConfig")
	// OrderBehaviorChain
	QueryOrderNode := NewFuncNode(QueryOrder, "QueryOrder")
	// OrderOtherBehaviorChain
	QueryOrderOtherNode := NewFuncNode(QueryOrderOther, "QueryOrderOther")

	//OrderBehaviorChain
	// StationBehaviorChain

	//OrderOtherBehaviorChain
	// StationBehaviorChain

	//SecurityBehaviorChain
	// OrderBehaviorChain
	// OrderOtherBehaviorChain
	QuerySecurityNode := NewFuncNode(QuerySecurity, "QuerySecurity")

	// ******* Preserve ********
	PreserveNode := NewFuncNode(Preserve, "Preserve")

	// ------------------------------------- NewChain -------------------------------------------
	// ------------------------------------- NewChain -------------------------------------------
	// ------------------------------------- NewChain -------------------------------------------
	// ------------------------------------- NewChain -------------------------------------------
	// ------------------------------------- NewChain -------------------------------------------
	// AssuranceBehaviorChain - Assurance
	QueryAssuranceChain := NewChain(QueryAssuranceNode)
	//CreateAssuranceChain := NewChain(CreateAssuranceNode)

	// UserBehaviorsChain
	// AuthBehaviorChain - LoginAdmin/LoginBasic
	LoginBasicChain := NewChain(LoginBasicNode)
	// VerifyCodeBehaviorChain
	VerifyCodeChain := NewChain(VerifyCodeNode)
	// UserBehaviorChain
	QueryUserChain := NewChain(QueryUserNode)

	// ContactsBehaviorChain - Contacts
	QueryContactsChain := NewChain(QueryContactsNode)
	CreateContactsChain := NewChain(CreateContactsNode)

	// ConsignBehaviorsChain
	//QueryConsignChain := NewChain(QueryConsignNode)
	CreateConsignChain := NewChain(CreateConsignNode)
	// ConsignPriceBehaviorChain
	//QueryConsignPriceChain := NewChain(QueryConsignPriceNode)
	//CreateConsignPriceChain := NewChain(CreateConsignPriceNode)

	// FoodBehaviorChain
	QueryFoodChain := NewChain(QueryFoodNode)
	// StationFoodBehaviorChain
	QueryStationFoodChain := NewChain(QueryStationFoodNode)
	// TrainFoodBehaviorChain
	QueryTrainFoodChain := NewChain(QueryTrainFoodNode)
	// TravelBehaviorChain
	QueryTravelChain1 := NewChain(QueryTravelNode)
	QueryTravelChain2 := NewChain(QueryTravelNode)
	//CreateTravelChain := NewChain(CreateTravelNode)

	// TrainBehaviorChain
	QueryTrainChain1 := NewChain(QueryTrainNode)
	QueryTrainChain2 := NewChain(QueryTrainNode)
	// RouteBehaviorChain
	QueryRouteChain1 := NewChain(QueryRouteNode)
	QueryRouteChain2 := NewChain(QueryRouteNode)
	// BasicBehaviorChain
	QueryBasicChain := NewChain(QueryRouteNode)
	QueryBasicChain.AddNode(QueryStationNode)
	QueryBasicChain.AddNode(QueryTrainNode)
	QueryBasicChain.AddNode(QueryBasicNode)
	QueryBasicChain.AddNode(QueryPriceNode)
	// SeatBehaviorChain
	CreateSeatChain := NewChain(QueryConfigNode)
	CreateSeatChain.AddNode(QueryOrderNode)
	CreateSeatChain.AddNode(QueryOrderOtherNode)
	CreateSeatChain.AddNode(CreateSeatNode)

	// StationBehaviorChain
	QueryStationChain0 := NewChain(QueryStationNode)
	QueryStationChain1 := NewChain(QueryStationNode)
	QueryStationChain2 := NewChain(QueryStationNode)
	QueryStationChain3 := NewChain(QueryStationNode)
	// PriceBehaviorChain
	QueryPriceChain := NewChain(QueryPriceNode)

	// ConfigBehaviorChain
	QueryConfigChain := NewChain(QueryConfigNode)
	// OrderBehaviorChain
	QueryOrderChain1 := NewChain(QueryOrderNode)
	QueryOrderChain2 := NewChain(QueryOrderNode)
	QueryOrderChain3 := NewChain(QueryOrderNode)
	// OrderOtherBehaviorChain
	QueryOrderOtherChain1 := NewChain(QueryOrderOtherNode)
	QueryOrderOtherChain2 := NewChain(QueryOrderOtherNode)

	// SecurityBehaviorChain
	QuerySecurityChain := NewChain(QuerySecurityNode)

	// The Last Chain - Preserve Behavior Chain
	PreserveChain := NewChain(PreserveNode)

	// ------------------------------------- AddNextChain -------------------------------------------
	// ------------------------------------- AddNextChain -------------------------------------------
	// ------------------------------------- AddNextChain -------------------------------------------
	// 逆序 - 从处理逆序的第一层开始

	// -------------(AddNextChain)AssignEachChainWithItsCorrespondingBehaviorChain-------------------
	// -------------(AddNextChain)AssignEachChainWithItsCorrespondingBehaviorChain-------------------
	// -------------(AddNextChain)AssignEachChainWithItsCorrespondingBehaviorChain-------------------
	//AssuranceBehaviorChain
	AssuranceBehaviorChain.AddNextChain(QueryAssuranceChain, 1)
	//VerifyCodeBehaviorChain
	VerifyCodeBehaviorChain.AddNextChain(VerifyCodeChain, 1)
	//AuthBehaviorChain
	AuthBehaviorChain.AddNextChain(LoginBasicChain, 1)
	//UserBehaviorChain
	UserBehaviorChain.AddNextChain(QueryUserChain, 1)
	//UserBehaviorsChain
	UserBehaviorsChain.AddNextChain(VerifyCodeBehaviorChain, 1)
	VerifyCodeChain.AddNextChain(AuthBehaviorChain, 1)
	LoginBasicChain.AddNextChain(UserBehaviorChain, 1)
	//ContactsBehaviorChain
	ContactsBehaviorChain.AddNextChain(QueryContactsChain, 0.7)
	ContactsBehaviorChain.AddNextChain(CreateContactsChain, 0.3)
	//ConsignPriceBehaviorChain
	//ConsignBehaviorsChain
	ConsignBehaviorsChain.AddNextChain(ConsignPriceBehaviorChain, 1)
	ConsignPriceBehaviorChain.AddNextChain(CreateConsignChain, 1)
	//StationBehaviorChain
	StationBehaviorChain0.AddNextChain(QueryStationChain0, 1)
	StationBehaviorChain1.AddNextChain(QueryStationChain1, 1)
	StationBehaviorChain2.AddNextChain(QueryStationChain2, 1)
	StationBehaviorChain3.AddNextChain(QueryStationChain3, 1)
	//OrderOtherBehaviorChain
	OrderOtherBehaviorChain1.AddNextChain(QueryOrderOtherChain1, 1)
	OrderOtherBehaviorChain2.AddNextChain(QueryOrderOtherChain2, 1)
	//OrderBehaviorChain
	OrderBehaviorChain1.AddNextChain(QueryOrderChain1, 1)
	OrderBehaviorChain2.AddNextChain(QueryOrderChain2, 1)
	OrderBehaviorChain3.AddNextChain(QueryOrderChain3, 1)
	//SecurityBehaviorChain
	SecurityBehaviorChain.AddNextChain(QuerySecurityChain, 1)
	//ConfigBehaviorChain
	ConfigBehaviorChain.AddNextChain(QueryConfigChain, 1)
	//SeatBehaviorChain
	SeatBehaviorChain.AddNextChain(OrderBehaviorChain2, 1)
	QueryOrderChain2.AddNextChain(OrderOtherBehaviorChain2, 1)
	QueryOrderOtherChain2.AddNextChain(CreateSeatChain, 1)
	//PriceBehaviorChain
	PriceBehaviorChain.AddNextChain(QueryPriceChain, 1)
	//TrainBehaviorChain
	TrainBehaviorChain1.AddNextChain(QueryTrainChain1, 1)
	TrainBehaviorChain2.AddNextChain(QueryTrainChain2, 1)
	//RouteBehaviorChain
	RouteBehaviorChain1.AddNextChain(QueryRouteChain1, 1)
	RouteBehaviorChain2.AddNextChain(QueryRouteChain2, 1)
	//BasicBehaviorChain
	BasicBehaviorChain.AddNextChain(QueryBasicChain, 1)
	//TravelBehaviorChain
	TravelBehaviorChain1.AddNextChain(QueryTravelChain1, 1)
	TravelBehaviorChain2.AddNextChain(QueryTravelChain2, 1)
	//StationFoodBehaviorChain
	StationFoodBehaviorChain.AddNextChain(QueryStationFoodChain, 1)
	//TrainFoodBehaviorChain
	TrainFoodBehaviorChain.AddNextChain(QueryTrainFoodChain, 1)
	//FoodBehaviorChain
	//FoodBehaviorChain.AddNextChain(TravelBehaviorChain2, 1)
	FoodBehaviorChain.AddNextChain(StationFoodBehaviorChain, 1)
	QueryStationFoodChain.AddNextChain(TrainFoodBehaviorChain, 1)
	QueryTrainFoodChain.AddNextChain(QueryFoodChain, 1)

	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& Main Chain &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& Main Chain &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& Main Chain &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& Main Chain &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& Main Chain &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	//PreserveChain
	/*	AssuranceBehaviorChain | 1. BasicBehaviorChain | ConsignBehaviorsChain |
		ContactsBehaviorChain | FoodBehaviorChain | TravelBehaviorChain | SeatBehaviorChain |
		SecurityBehaviorChain | OrderBehaviorChain1 | StationBehaviorChain0 |
		0. UserBehaviorsChain */

	PreserveBehaviorChain.AddNextChain(UserBehaviorsChain, 1)
	// UserBehaviorsChain
	QueryUserChain.AddNextChain(BasicBehaviorChain, 1)
	// BasicBehaviorChain
	QueryBasicChain.AddNextChain(SeatBehaviorChain, 1)
	// SeatBehaviorChain
	CreateSeatChain.AddNextChain(ContactsBehaviorChain, 1)
	// TravelBehaviorChain
	QueryContactsChain.AddNextChain(TravelBehaviorChain1, 1)
	CreateContactsChain.AddNextChain(TravelBehaviorChain1, 1)
	// ContactsBehaviorChain
	QueryTravelChain1.AddNextChain(AssuranceBehaviorChain, 1)
	// AssuranceBehaviorChain
	QueryAssuranceChain.AddNextChain(FoodBehaviorChain, 1)
	// FoodBehaviorChain
	QueryFoodChain.AddNextChain(ConsignBehaviorsChain, 1)
	// ConsignBehaviorsChain
	CreateConsignChain.AddNextChain(PreserveChain, 1)

	// ------------------------------------- VisualizeChain -------------------------------------------
	// ------------------------------------- VisualizeChain -------------------------------------------
	fmt.Println(PreserveChain.VisualizeChain(0))
	fmt.Println()
}

// ************************************* NewFuncNode_Function *******************************************
// ************************************* NewFuncNode_Function *******************************************
// ************************************* NewFuncNode_Function *******************************************

// AssuranceBehaviorChain
func QueryAssurance(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	Assurances, err := cli.GetAllAssurances()
	if err != nil {
		log.Println("GetAllAssurances failed: %v", err)
		return nil, err
	}
	if Assurances.Status != 1 {
		log.Println("Assurances status is not 1: %d", Assurances.Status)
		return nil, nil
	}

	randomIndex := rand.Intn(len(Assurances.Data))
	ctx.Set(OrderId, Assurances.Data[randomIndex].OrderId)
	ctx.Set(AssuranceTypeIndex, Assurances.Data[randomIndex].TypeIndex)
	ctx.Set(AssuranceTypeName, Assurances.Data[randomIndex].TypeName)
	ctx.Set(AssuranceTypePrice, Assurances.Data[randomIndex].TypePrice)

	return nil, nil
}

func CreateAssurance(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	//Create a new assurance
	TheOrderID := ctx.Get(OrderId).(string)
	addAssuranceResp, err := cli.CreateNewAssurance(1, TheOrderID) // typeIndex 1 -> TRAFFIC_ACCIDENT
	if err != nil {
		log.Println("CreateNewAssurance failed: %v", err)
		return nil, err
	}
	if addAssuranceResp.Msg == "Already exists" {
		log.Println("Order ID found, skip")
		return nil, err
	}
	if addAssuranceResp.Data.OrderId != TheOrderID {
		log.Println("Request failed, addAssuranceResp.Data.OrderId:%s, expected: %s", addAssuranceResp.Data.OrderId, TheOrderID)
		return nil, err
	}
	if addAssuranceResp.Data.Type != "TRAFFIC_ACCIDENT" {
		log.Println("Request failed, addAssuranceResp.Data.Type are expected to be 'TRAFFIC_ACCIDENT' but actually: %v", addAssuranceResp.Data.Type)
		return nil, err
	}

	ctx.Set(OrderId, addAssuranceResp.Data.OrderId)
	//ctx.Set(TypeIndex, addAssuranceResp.Data.)
	//ctx.Set(TypeName, Assurances.Data[randomIndex].TypeName)
	//ctx.Set(TypePrice, Assurances.Data[randomIndex].TypePrice)

	return nil, nil
}

//UserBehaviorsChain
// LoginBasicChain
//func LoginBasic(ctx *Context) (*NodeResult, error) {
//	cli, ok := ctx.Get(Client).(*service.SvcImpl)
//	if !ok {
//		return nil, fmt.Errorf("service client not found in context")
//	}
//	// login
//	loginResult, err := cli.ReqUserLogin(&service.UserLoginInfoReq{
//		Password:         "111111",
//		UserName:         "fdse_microservice",
//		VerificationCode: "123",
//	})
//	if err != nil {
//		return nil, err
//	}
//	ctx.Set(LoginToken, loginResult.Data.Token)
//	return nil, nil
//}

// VerifyCodeBehaviorChain
func VerifyCode(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	verifyCode := generateVerifyCode()
	verifyCodeResp, err := cli.VerifyCode(verifyCode)
	if err != nil {
		log.Println("Request failed, err %s", err)
		return nil, err
	}
	if !verifyCodeResp {
		log.Println("Verification failed")
		return nil, err
	}
	//log.Println("Verification code verified. The result is %v and verifyCode: %v", verifyCodeResp, verifyCode)

	ctx.Set(BooleanVerifyCode, verifyCodeResp)

	return nil, nil
}

func QueryUser(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	allUsersResp, err := cli.GetAllUsers()
	if err != nil {
		log.Println("Request failed, err1 %s", err)
		return nil, err
	}
	if allUsersResp.Status != 1 {
		log.Println("Expected status 200, got %d", allUsersResp.Status)
		return nil, err
	}

	randomIndex := rand.Intn(len(allUsersResp.Data))
	ctx.Set(UserID, allUsersResp.Data[randomIndex].UserID)
	ctx.Set(UserName, allUsersResp.Data[randomIndex].UserName)
	ctx.Set(Password, allUsersResp.Data[randomIndex].Password)
	ctx.Set(Gender, allUsersResp.Data[randomIndex].Gender)
	ctx.Set(DocumentNum, allUsersResp.Data[randomIndex].DocumentNum)
	ctx.Set(DocumentType, allUsersResp.Data[randomIndex].DocumentType)
	ctx.Set(Email, allUsersResp.Data[randomIndex].Email)

	return nil, nil
}

// ContactsBehaviorChain
func QueryContacts(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	var contactsSvc service.ContactsService = cli
	TheAccountId := ctx.Get(UserId).(string)
	GetAllContacts, err := contactsSvc.GetContactByAccountId(TheAccountId)
	if err != nil {
		log.Println("[Mock AccountID]GetAllContacts fail. The error occurs: %v", err)
		return nil, err
	}
	if GetAllContacts.Status != 1 {
		log.Println("[Mock AccountID]GetAllContacts.Status != 1")
		return nil, err
	}

	randomIndex := rand.Intn(len(GetAllContacts.Data))
	ctx.Set(AccountID, GetAllContacts.Data[randomIndex].AccountId)
	ctx.Set(ContactsID, GetAllContacts.Data[randomIndex].Id)
	ctx.Set(Name, GetAllContacts.Data[randomIndex].Name)
	ctx.Set(DocumentType, GetAllContacts.Data[randomIndex].DocumentType)
	ctx.Set(DocumentNumber, GetAllContacts.Data[randomIndex].DocumentNumber)
	ctx.Set(PhoneNumber, GetAllContacts.Data[randomIndex].PhoneNumber)

	return nil, nil
}

func CreateContacts(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	CreateContactsInput := service.AdminContacts{
		Id:             faker.UUIDHyphenated(),
		AccountId:      ctx.Get(UserId).(string),
		Name:           faker.Name(),
		DocumentType:   rand.Intn(1),
		DocumentNumber: generateDocumentNumber(),
		PhoneNumber:    faker.PhoneNumber,
	}
	CreateContacts, err := cli.AddContact(&CreateContactsInput)
	if err != nil {
		log.Println("[Mock AccountID] CreateContacts error occurs: %v", err)
		return nil, err
	}
	if CreateContacts.Status != 1 {
		log.Println("[Mock AccountID] CreateContacts.Status != 1")
		return nil, err
	}

	ctx.Set(AccountID, CreateContacts.Data.AccountId)
	ctx.Set(ContactsID, CreateContacts.Data.Id)
	ctx.Set(Name, CreateContacts.Data.Name)
	ctx.Set(DocumentType, CreateContacts.Data.DocumentType)
	ctx.Set(DocumentNumber, CreateContacts.Data.DocumentNumber)
	ctx.Set(PhoneNumber, CreateContacts.Data.PhoneNumber)

	return nil, nil
}

// ConsignBehaviorsChain
func QueryConsign(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	//TheAccountId := ctx.Get(AccountID).(string)
	//// Query consign records by account ID
	//consignsByAccountId, err := cli.QueryByAccountId(TheAccountId)
	//if err != nil {
	//	log.Println("QueryByAccountId failed: %v", err)
	//}
	//if consignsByAccountId.Status != 1 {
	//	log.Println("consignsByAccountId failed")
	//
	//}
	///*found := false
	//for _, consign := range consignsByAccountId.Data {
	//	if consign.IsWithin == existedConsign.IsWithin &&
	//		consign.To == existedConsign.To &&
	//		consign.Weight == existedConsign.Weight &&
	//		consign.ID == existedConsign.ID &&
	//		consign.Phone == existedConsign.Phone &&
	//		consign.HandleDate == existedConsign.HandleDate &&
	//		consign.TargetDate == existedConsign.TargetDate &&
	//		consign.OrderID == existedConsign.OrderID &&
	//		consign.Consignee == existedConsign.Consignee &&
	//		consign.From == existedConsign.From &&
	//		consign.AccountID == existedConsign.AccountID {
	//		found = true
	//	}
	//}
	//if !found {
	//	log.Println("Can not find consign by accountId.")
	//}*/
	//log.Println("QueryByAccountId response: %+v", consignsByAccountId)

	// Query consign records by order ID
	TheOrderId := ctx.Get(OrderId).(string)
	consignsByOrderId, err := cli.QueryByOrderId(TheOrderId)
	if err != nil {
		log.Println("QueryByOrderId failed: %v", err)
		return nil, err
	}
	if consignsByOrderId.Status != 1 {
		log.Println("consignsByOrderId.Status = 1")
		return nil, err
	}
	/*isMatch1 := false
	if consignsByOrderId.Data.OrderId == existedConsign.OrderID &&
		consignsByOrderId.Data.Id == existedConsign.ID &&
		consignsByOrderId.Data.From == existedConsign.From &&
		consignsByOrderId.Data.To == existedConsign.To &&
		consignsByOrderId.Data.Phone == existedConsign.Phone &&
		consignsByOrderId.Data.Consignee == existedConsign.Consignee &&
		consignsByOrderId.Data.TargetDate == existedConsign.TargetDate &&
		consignsByOrderId.Data.HandleDate == existedConsign.HandleDate &&
		consignsByOrderId.Data.Weight == existedConsign.Weight &&
		consignsByOrderId.Data.AccountId == existedConsign.AccountID {
		isMatch1 = true
	}
	if !isMatch1 {
		log.Println("Can not find consign by orderId.")
	}*/
	//log.Println("QueryByOrderId response: %+v", consignsByOrderId)

	// Query consign records by consignee
	//TheConsignee := ctx.Get(Name).(string)
	//consignsByConsignee, err := cli.QueryByConsignee(TheConsignee)
	//if err != nil {
	//	log.Println("QueryByConsignee failed: %v", err)
	//}
	//if consignsByConsignee.Status != 1 {
	//	log.Println("consignsByConsignee failed.")
	//}
	//isMatch2 := false
	//for _, consign := range consignsByConsignee.Data {
	//	if consign.Id == existedConsign.ID &&
	//		consign.AccountId == existedConsign.AccountID &&
	//		consign.OrderId == existedConsign.OrderID &&
	//		consign.To == existedConsign.To &&
	//		consign.From == existedConsign.From &&
	//		consign.Weight == existedConsign.Weight &&
	//		consign.HandleDate == existedConsign.HandleDate &&
	//		consign.TargetDate == existedConsign.TargetDate &&
	//		consign.Phone == existedConsign.Phone &&
	//		consign.Consignee == existedConsign.Consignee {
	//		isMatch2 = true
	//	}
	//}
	//if !isMatch2 {
	//	log.Println("Can not find consign by consignee.")
	//}
	//log.Println("QueryByConsignee response: %+v", consignsByConsignee)

	ctx.Set(ID, consignsByOrderId.Data.Id)
	ctx.Set(OrderId, consignsByOrderId.Data.OrderId)
	ctx.Set(AccountID, consignsByOrderId.Data.AccountId)
	ctx.Set(HandleDate, consignsByOrderId.Data.HandleDate)
	ctx.Set(TargetDate, consignsByOrderId.Data.TargetDate)
	ctx.Set(From, consignsByOrderId.Data.From)
	ctx.Set(To, consignsByOrderId.Data.To)
	ctx.Set(Consignee, consignsByOrderId.Data.Consignee)
	ctx.Set(Phone, consignsByOrderId.Data.Phone)
	ctx.Set(Weight, consignsByOrderId.Data.Weight)
	ctx.Set(Price, consignsByOrderId.Data.Price)

	return nil, nil
}

func CreateConsign(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Mock data
	MockedId := faker.UUIDHyphenated()
	MockedAccountId := ctx.Get(AccountID).(string)
	MockedOrderId := ctx.Get(OrderId).(string)
	MockedHandleDate := ctx.Get(DepartureTime).(string)
	//MockedHandleDate := ""
	MockedTargetDate := ctx.Get(DepartureTime).(string)
	//MockedTargetDate := ""
	MockedFromPlace := ctx.Get(StartStation).(string)
	MockedToPlace := ctx.Get(EndStation).(string)
	MockedConsignee := faker.Name()
	MockedPhone := faker.Phonenumber()
	MockedWeight := GenerateWeight()

	// Insert a new consign record
	insertReq := service.Consign{
		ID:         MockedId,
		OrderID:    MockedOrderId,
		AccountID:  MockedAccountId,
		HandleDate: MockedHandleDate,
		TargetDate: MockedTargetDate,
		From:       MockedFromPlace,
		To:         MockedToPlace,
		Consignee:  MockedConsignee,
		Phone:      MockedPhone,
		Weight:     MockedWeight,
		IsWithin:   BooleanIsWithin(MockedWeight),
	}
	insertResp, err := cli.InsertConsignRecord(&insertReq)
	if err != nil {
		log.Println("InsertConsignRecord failed: %v", err)
		return nil, err
	}
	if insertResp.Msg == "Already exists" {
		return nil, fmt.Errorf("Consign already exists")
	}
	if insertResp.Status != 1 {
		log.Println("InsertConsignRecord failed: %v", insertResp.Status)
		return nil, err
	}
	isMatch := false
	if /*insertResp.Data.ID == insertReq.ID &&*/
	/*insertResp.Data.IsWithin == insertReq.IsWithin &&*/
	insertResp.Data.AccountID == insertReq.AccountID &&
		insertResp.Data.From == insertReq.From &&
		insertResp.Data.Consignee == insertReq.Consignee &&
		insertResp.Data.OrderID == insertReq.OrderID &&
		insertResp.Data.Phone == insertReq.Phone &&
		insertResp.Data.TargetDate == insertReq.TargetDate &&
		insertResp.Data.HandleDate == insertReq.HandleDate &&
		insertResp.Data.To == insertReq.To &&
		insertResp.Data.Weight == insertReq.Weight {
		isMatch = true
	}
	if !isMatch {
		log.Println("Creation not match. Expect: %v, but get: %v", insertReq, insertResp.Data)
		return nil, err
	}
	//log.Println("InsertConsignRecord response: %+v", insertResp)
	//existedConsign := insertResp.Data

	ctx.Set(ID, insertResp.Data.ID)
	//ctx.Set(OrderID, insertResp.Data.OrderID)
	//ctx.Set(AccountID, insertResp.Data.AccountID)
	ctx.Set(HandleDate, insertResp.Data.HandleDate)
	ctx.Set(TargetDate, insertResp.Data.TargetDate)
	ctx.Set(From, insertResp.Data.From)
	ctx.Set(To, insertResp.Data.To)
	ctx.Set(Consignee, insertResp.Data.Consignee)
	ctx.Set(Phone, insertResp.Data.Phone)
	ctx.Set(Weight, insertResp.Data.Weight)
	ctx.Set(IsWithin, insertResp.Data.IsWithin)

	return nil, nil
}

func QueryConsignPric(ctx *Context) (*NodeResult, error) {
	_, ok := ctx.Get(Client).(*service.SvcImpl)
	//cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part

	return nil, nil
}

func CreateConsignPrice(ctx *Context) (*NodeResult, error) {
	_, ok := ctx.Get(Client).(*service.SvcImpl)
	//cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")

	}

	// TODO part

	return nil, nil
}

// FoodBehaviorChain
func QueryFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Query all
	allFoodOrders, err := cli.FindAllFoodOrder()
	if err != nil {
		log.Println("FindAllFoodOrder request failed, err %s", err)
		return nil, err
	}
	if len(allFoodOrders.Data) == 0 {
		log.Println("FindAllFoodOrder returned empty results")
		return nil, err
	}
	if allFoodOrders.Status != 1 {
		log.Println("FindAllFoodOrder failed: %v", allFoodOrders.Status)
		return nil, err
	}

	randomIndex := rand.Intn(len(allFoodOrders.Data))
	//ctx.Set(OrderId, allFoodOrders.Data[randomIndex].OrderId)
	ctx.Set(FoodType, allFoodOrders.Data[randomIndex].FoodType)
	//ctx.Set(StationName, allFoodOrders.Data[randomIndex].StationName)
	ctx.Set(StoreName, allFoodOrders.Data[randomIndex].StoreName)
	ctx.Set(FoodName, allFoodOrders.Data[randomIndex].FoodName)
	ctx.Set(Price, allFoodOrders.Data[randomIndex].Price)

	return nil, nil
}

func CreateFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Mock data
	MockedOrderID := ctx.Get(OrderId).(string)
	MockedID := faker.UUIDHyphenated()
	foodOrder := service.FoodOrder{
		ID:          MockedID,
		OrderID:     MockedOrderID,
		FoodType:    rand.Intn(1),
		FoodName:    generateRandomFood(),
		StationName: ctx.Get(StationName).(string),
		StoreName:   ctx.Get(StoreName).(string),
		Price:       ctx.Get(Price).(float64),
	}

	// Create Test
	newCreateResp, err := cli.CreateFoodOrder(&foodOrder)
	if err != nil {
		log.Println("NewCreateFoodOrder request failed, err %s", err)
		return nil, err
	}
	if newCreateResp.Status != 1 {
		log.Println("NEwCreateFoodOrder failed")
		return nil, err
	}

	ctx.Set(OrderId, newCreateResp.Data.OrderId)
	ctx.Set(FoodType, newCreateResp.Data.FoodType)
	ctx.Set(StationName, newCreateResp.Data.StationName)
	ctx.Set(StoreName, newCreateResp.Data.StoreName)
	ctx.Set(FoodName, newCreateResp.Data.FoodName)
	ctx.Set(Price, newCreateResp.Data.Price)

	return nil, nil
}

func QueryStationFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.GetAllStationFood()
	if err != nil {
		log.Println("Resp returns err: %v", err)
		return nil, err
	}
	if resp.Status != 1 {
		log.Println("GetAllStationFood status should be 1, but is %d", resp.Status)
		return nil, err
	}

	//Id           string  `json:"id"`
	//StationName  string  `json:"stationName"`
	//StoreName    string  `json:"storeName"`
	//Telephone    string  `json:"telephone"`
	//BusinessTime string  `json:"businessTime"`
	//DeliveryFee  float64 `json:"deliveryFee"`
	//FoodList     []struct {
	//	FoodName string  `json:"foodName"`
	//	Price    float64 `json:"price"`
	//} `json:"foodList"`

	randomIndex := rand.Intn(len(resp.Data))
	//ctx.Set(ID, resp.Data[randomIndex].Id)
	ctx.Set(StationName, resp.Data[randomIndex].StationName)
	ctx.Set(StoreName, resp.Data[randomIndex].StoreName)
	ctx.Set(Phone, resp.Data[randomIndex].Telephone)
	ctx.Set(Price, resp.Data[randomIndex].DeliveryFee)

	return nil, nil

}

func QueryTrainFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.GetAllTrainFood()
	if err != nil {
		log.Println("resp returns err: %v", err)
		return nil, err
	}
	if resp.Status != 1 {
		log.Println("GetAllTrainFood's status should be 1 but got %d", resp.Status)
		return nil, nil
	}

	//	Id       string `json:"id"`
	//	TripId   string `json:"tripId"`
	//	FoodList []struct {
	//		FoodName string  `json:"foodName"`
	//		Price    float64 `json:"price"`
	//	} `json:"foodList"`
	//} `json:"data"`

	randomIndex := rand.Intn(len(resp.Data))
	randomFoodlistIndex := rand.Intn(len(resp.Data[randomIndex].FoodList))
	ctx.Set(TripID, resp.Data[randomIndex].TripId)
	ctx.Set(FoodName, resp.Data[randomIndex].FoodList[randomFoodlistIndex].FoodName)
	ctx.Set(FoodPrice, resp.Data[randomIndex].FoodList[randomFoodlistIndex].Price)

	return nil, nil
}

// TravelBehaviorChain
func QueryTrip(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	/*QueryAllTripResp, err := cli.QueryAllTrip()
	if err != nil {
		log.Println("Request failed, err %s", err)
		return nil, err
	}
	if QueryAllTripResp.Status != 1 {
		log.Println("Request failed, status: %d", QueryAllTripResp.Status)
		return nil, err
	}*/
	// Test QueryInfo
	tripInfo := service.TripInfo{
		StartPlace:    ctx.Get(StartStation).(string),
		EndPlace:      ctx.Get(EndStation).(string),
		DepartureTime: ctx.Get(DepartureTime).(string), // make no sence
		//DepartureTime: "",                              //????????????????????????????????????
	}
	queryInfoResp, err := cli.QueryInfo(tripInfo)
	if err != nil {
		log.Println("QueryInfo request failed, err %s", err)
		return nil, err
	}
	if queryInfoResp.Status != 1 {
		log.Println("QueryInfo failed, status: %d", queryInfoResp.Status)
		return nil, err
	}

	/*	TripId = "tripId"
		TrainTypeName = "trainTypeName"
		StartStation = "startStation"
		TerminalStation = "terminalStation"
		StartTime = "startTime"
		EndTime = "endTime"
		EconomyClass = "economyClass"
		ConfortClass = "confortClass"
		PriceForEconomyClass = "priceForEconomyClass"
		PriceForConfortClass = "priceForConfortClass"*/

	randomIndex := rand.Intn(len(queryInfoResp.Data))
	ctx.Set(TripId, fmt.Sprintf("%s%s", queryInfoResp.Data[randomIndex].TripId.Type, queryInfoResp.Data[randomIndex].TripId.Number))
	//ctx.Set(TrainTypeName, queryInfoResp.Data[randomIndex].TrainTypeName)
	ctx.Set(StartStation, queryInfoResp.Data[randomIndex].StartStation)
	ctx.Set(TerminalStation, queryInfoResp.Data[randomIndex].TerminalStation)
	ctx.Set(StartTime, queryInfoResp.Data[randomIndex].StartTime)
	ctx.Set(EndTime, queryInfoResp.Data[randomIndex].EndTime)
	ctx.Set(EconomyClass, queryInfoResp.Data[randomIndex].EconomyClass)
	ctx.Set(ConfortClass, queryInfoResp.Data[randomIndex].ConfortClass)
	ctx.Set(PriceForEconomyClass, queryInfoResp.Data[randomIndex].PriceForEconomyClass)
	ctx.Set(PriceForConfortClass, queryInfoResp.Data[randomIndex].PriceForConfortClass)
	// ????????????????????????

	return nil, nil
}

func CreateTrip(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Mock para
	MockedLoginId := ctx.Get(LoginToken).(string)
	//MockedTripId := GenerateTripId()
	MockedTripId := ctx.Get(TripId).(string)
	//MockedTrainTypeName := generateTrainTypeName(MockedTripId) /*"GaoTieSeven"*/
	MockedTrainTypeName := ctx.Get(TrainTypeName).(string)
	MockedRouteID := ctx.Get(RouteID).(string)
	MockedStartStationName := ctx.Get(From).(string)
	MockedStationsName := /*strings.Join(AllRoutesByQuery.Data[0].Stations, ",")*/ ctx.Get(StationName).(string)
	MockedTerminalStationName := ctx.Get(To).(string)
	MockedStartTime := ctx.Get(StartTime).(string)
	MockedEndTime := ctx.Get(EndTime).(string)

	// Mock input
	travelInfo := service.TravelInfo{
		LoginID:          MockedLoginId,
		TripID:           MockedTripId,
		TrainTypeName:    MockedTrainTypeName,
		RouteID:          MockedRouteID,
		StartStationName: MockedStartStationName,
		//StationsName:        fmt.Sprintf("%v,%v", MockedStartStationName, MockedTerminalStationName),
		StationsName:        MockedStationsName,
		TerminalStationName: MockedTerminalStationName,
		StartTime:           MockedStartTime,
		EndTime:             MockedEndTime,
	}

	// Create Test
	createResp, err := cli.CreateTrip(&travelInfo)
	if err != nil {
		log.Println("CreateTrip request failed, err %s", err)
		return nil, err
	}
	if createResp.Status != 1 {
		log.Println("CreateTrip failed: %s", createResp.Msg)
		return nil, err
	}
	if createResp.Msg == "Already exists" {
		log.Println("Already exists: %s", createResp.Msg)
		return nil, err
	}
	isMatch := false
	if /*createResp.Data.Id == travelInfo.LoginID &&*/
	createResp.Data.StationsName == toLowerCaseAndRemoveSpaces(travelInfo.StationsName) &&
		createResp.Data.StartStationName == toLowerCaseAndRemoveSpaces(travelInfo.StartStationName) &&
		createResp.Data.TerminalStationName == toLowerCaseAndRemoveSpaces(travelInfo.TerminalStationName) &&
		createResp.Data.StartTime == travelInfo.StartTime &&
		createResp.Data.EndTime == travelInfo.EndTime &&
		createResp.Data.TrainTypeName == travelInfo.TrainTypeName &&
		createResp.Data.RouteId == travelInfo.RouteID {
		isMatch = true
	}
	if !isMatch {
		log.Println("CreateTrip failed: %s. Except: %v, but get: %v", createResp.Msg, travelInfo, createResp.Data)
		return nil, err
	}

	/*	EndTime             string `json:"endTime"`
		Id                  string `json:"id"`
		RouteId             string `json:"routeId"`
		StartStationName    string `json:"startStationName"`
		StartTime           string `json:"startTime"`
		StationsName        string `json:"stationsName"`
		TerminalStationName string `json:"terminalStationName"`
		TrainTypeName       string `json:"trainTypeName"`
		TripId              TripId `json:"tripId"`*/

	//ctx.Set(TripId, createResp.Data.TripId)
	//ctx.Set(TrainTypeName, createResp.Data.TrainTypeName)
	//ctx.Set(StartStation, createResp.Data.StartStation)
	//ctx.Set(TerminalStation, createResp.Data.TerminalStation)
	//ctx.Set(StartTime, queryInfoResp.Data[randomIndex].StartTime)
	//ctx.Set(EndTime, queryInfoResp.Data[randomIndex].EndTime)
	//ctx.Set(EconomyClass, queryInfoResp.Data[randomIndex].EconomyClass)
	//ctx.Set(ConfortClass, queryInfoResp.Data[randomIndex].ConfortClass)
	//ctx.Set(PriceForEconomyClass, queryInfoResp.Data[randomIndex].PriceForEconomyClass)
	//ctx.Set(PriceForConfortClass, queryInfoResp.Data[randomIndex].PriceForConfortClass)
	return nil, nil
}

func QueryTrain(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Query all
	allTrainTypes, err := cli.Query()
	if err != nil {
		log.Println("Query all request failed, err %s", err)
		return nil, err
	}
	if allTrainTypes.Status != 1 {
		log.Println("allTrainTypes.Status != 1")
		return nil, err
	}
	if len(allTrainTypes.Data) == 0 {
		log.Println("Query all returned no results")
		return nil, err
	}
	/*found := false
	for _, trainTypeElement := range allTrainTypes.Data {
		if trainTypeElement.Id == createResp.Data.Id &&
			trainTypeElement.Name == existedtrainType.Name &&
			trainTypeElement.AverageSpeed == existedtrainType.AverageSpeed &&
			trainTypeElement.ConfortClass == existedtrainType.ConfortClass &&
			trainTypeElement.EconomyClass == existedtrainType.ConfortClass {
			found = true
		}
	}
	if !found {
		t.Errorf("Query all not get the corresponsing result, whcih means 'Creation Fails'")
	}*/

	/*	Id           string `json:"id"`
		Name         string `json:"name"`
		ConfortClass int    `json:"confortClass"`
		AverageSpeed int    `json:"averageSpeed"`
		EconomyClass int    `json:"economyClass"`*/
	randomIndex := rand.Intn(len(allTrainTypes.Data))
	ctx.Set(TrainTypName, allTrainTypes.Data[randomIndex].Name)
	ctx.Set(ConfortClass, allTrainTypes.Data[randomIndex].ConfortClass)
	ctx.Set(AverageSpeed, allTrainTypes.Data[randomIndex].AverageSpeed)
	ctx.Set(EconomyClass, allTrainTypes.Data[randomIndex].EconomyClass)

	return nil, nil
}

func QueryRoute(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	AllRoutesByQuery, err := cli.QueryAllRoutes()
	if err != nil {
		log.Println("Request failed, err2 %s", err)
		return nil, err
	}
	if AllRoutesByQuery.Status != 1 {
		log.Fatal("AllRoutes_By_Query.Status != 1")
		return nil, err
	}

	/*	Id           string   `json:"id"`
		Stations     []string `json:"stations"`
		Distances    []int    `json:"distances"`
		StartStation string   `json:"startStation"`
		EndStation   string   `json:"endStation"`*/

	randomIndex := rand.Intn(len(AllRoutesByQuery.Data))
	ctx.Set(RouteID, AllRoutesByQuery.Data[randomIndex].Id)
	ctx.Set(StartStation, AllRoutesByQuery.Data[randomIndex].StartStation)
	ctx.Set(EndStation, AllRoutesByQuery.Data[randomIndex].EndStation)
	//ctx.Set(StationName, getMiddleElements(strings.Join(AllRoutesByQuery.Data[randomIndex].Stations, ",")))
	ctx.Set(StationName, AllRoutesByQuery.Data[randomIndex].Stations)
	ctx.Set(Distances, AllRoutesByQuery.Data[randomIndex].Distances)

	return nil, nil
}

//func CreateRoute(ctx *Context) (*NodeResult, error) {
//	cli, ok := ctx.Get(Client).(*service.SvcImpl)

//	if !ok {
//		return nil, fmt.Errorf("service client not found in context")
//	}
//
//	// Create
//	MockedID := faker.UUIDHyphenated()
//	MockedStartStation := faker.GetRealAddress().City
//	MockedEndStation := faker.GetRealAddress().City
//	MockedStationList := fmt.Sprintf("%s,%s,%s", MockedStartStation, faker.GetRealAddress().City, MockedEndStation)
//	MockedDistanceList := fmt.Sprintf("%d,%d,%d", rand.Intn(30), rand.Intn(30), rand.Intn(30))
//	input := service.RouteInfo{
//		ID:           MockedID,
//		StartStation: MockedStartStation,
//		EndStation:   MockedEndStation,
//		StationList:  MockedStationList,
//		DistanceList: MockedDistanceList,
//	}
//	resp, err := cli.CreateAndModifyRoute(&input)
//	if err != nil {
//		log.Println("Request failed, err %s", err)
//		return nil, err
//	}
//	if resp.Msg == "Already exists" {
//		log.Println("Route already exists, skip")
//		return nil, err
//	}
//	if resp.Data.Id != input.ID {
//		log.Println("Route ID does not match, expect %s, got %s", input.ID, resp.Data.Id)
//		return nil, err
//	}
//	if resp.Data.StartStation != input.StartStation {
//		log.Println("StartStation does not match, expect %s, got %s", input.StartStation, resp.Data.StartStation)
//		return nil, err
//	}
//	if resp.Data.EndStation != input.EndStation {
//		log.Println("StartStation does not match, expect %s, got %s", input.StartStation, resp.Data.StartStation)
//		return nil, err
//	}
//	if StringSliceToString(resp.Data.Stations) != ConvertCommaSeparatedToBracketed(input.StationList) {
//		log.Println("StationList does not match, expect %s, got %s", ConvertCommaSeparatedToBracketed(input.StationList), StringSliceToString(resp.Data.Stations))
//		return nil, err
//	}
//	if IntSliceToString(resp.Data.Distances) != ConvertCommaSeparatedToBracketed(input.DistanceList) {
//		log.Println("DistanceList does not match, expect %s, got %s", ConvertCommaSeparatedToBracketed(input.DistanceList), IntSliceToString(resp.Data.Distances))
//		return nil, err
//	}
//
//	ctx.Set(From, resp.Data.StartStation)
//	ctx.Set(To, resp.Data.EndStation)
//	ctx.Set(StationName, getMiddleElements(strings.Join(resp.Data.Stations, ",")))
//	ctx.Set(RouteID, resp.Data.Id)
//
//	return nil, nil
//}

func QueryBasic(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// ----------------------------------  Mock data  -----------------------------------------
	// ----------------------------------  Mock data  -----------------------------------------
	// ----------------------------------  Mock data  -----------------------------------------
	//MockedTripId := faker.UUIDHyphenated()
	MockedTripTripId := GenerateTripId()
	MockedTripTripIdType := MockedTripTripId[0]
	MockedTripTripIdNumber := MockedTripTripId[1:]

	// ********************************* Input_By_Mock ****************************************
	// ********************************* Input_By_Mock ****************************************
	// ********************************* Input_By_Mock ****************************************
	travelQuery := &service.Travel{
		Trip: service.Trip{
			Id: faker.UUIDHyphenated(), // randomly generated
			TripId: service.TripId{
				Type:   fmt.Sprintf("%c", MockedTripTripIdType),
				Number: MockedTripTripIdNumber,
			},
			TrainTypeName:       GenerateTrainTypeName(),
			RouteId:             ctx.Get(RouteID).(string),
			StartStationName:    ctx.Get(StartStation).(string),
			StationsName:        getMiddleElements(strings.Join(ctx.Get(StationName).([]string), ",")), // only ok when there is exactly three stations
			TerminalStationName: ctx.Get(EndStation).(string),
			StartTime:           "", // can be any
			EndTime:             "", // can be any
		},
		StartPlace: ctx.Get(StartStation).(string),
		EndPlace:   ctx.Get(EndStation).(string),
		//DepartureTime: extractDate(getRandomTime()), // 生成1小时到1天之后的时间
		DepartureTime: getRandomTime(), // 生成1小时到1天之后的时间
	}

	var basicSvc service.BasicService = cli
	travel, err := basicSvc.QueryForTravel(travelQuery)
	if err != nil {
		log.Println("Query travel request failed, err %s", err)
		return nil, err
	}
	if travel.Status != 1 {
		log.Println("travel.Status != 1")
		return nil, err
	}

	/*	//Status = "status"
		Percent = "percent"
		//TrainType struct {
		//Id = "id"
		//Name = "name"
		//EconomyClass = "economyClass"
		//ConfortClass = "confortClass"
		//AverageSpeed = "averageSpeed"
		//} `json:"trainType"`
		//Route struct {
		//RouteID = "id"
		//Stations = "stations"
		//Distances = "distances"
		//StartStation = "startStation"
		//EndStation = "endStation"
		//} `json:"route"`
		//Prices struct {
		//ConfortClass = "confortClass"
		//EconomyClass = "economyClass"
		//} `json:"prices"`*/
	ctx.Set(Status, travel.Data.Status)
	ctx.Set(Percent, travel.Data.Percent)
	ctx.Set(TrainType, travel.Data.TrainType)
	ctx.Set(Route, travel.Data.Route)
	ctx.Set(Prices, travel.Data.Prices)
	//
	ctx.Set(DepartureTime, travelQuery.DepartureTime)
	ctx.Set(TrainTypeName, travelQuery.Trip.TrainTypeName)

	return nil, nil
}

func CreateSeat(ctx *Context) (*NodeResult, error) {
	// cli, ok := ctx.Get(Client).(*service.SvcImpl)
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	totalNum := rand.Intn(10)
	seatCreateInfoReq := &service.SeatCreateInfoReq{
		TravelDate:  ctx.Get(TravelDate).(string),
		TrainNumber: ctx.Get(TrainNumber).(string),
		DestStation: ctx.Get(EndStation).(string),
		SeatType:    ctx.Get(SeatClass).(int),
		TotalNum:    totalNum,
		Stations:    ctx.Get(StationName).([]string),
	}

	resp, err := cli.ReqSeatCreate(seatCreateInfoReq)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	if resp.Status != 1 {
		log.Println("SeatCreateInfoReq.Status != 1")
		return nil, err
	}

	ctx.Set(SeatNo, resp.Data.SeatNo)
	ctx.Set(DestStation, resp.Data.DestStation)
	//ctx.Set(StartStation, resp.Data.StartStation)

	return nil, nil
}

func QueryStation(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	QueryAll, err7 := cli.QueryStations()
	if err7 != nil {
		log.Println("Request failed, err7 %s", err7)
		return nil, err7
	}
	if QueryAll.Status != 1 {
		log.Println("Request failed, QueryAll.Status: %d, expected: %d", QueryAll.Status, 1)
		return nil, err7
	}
	/*found := false
	for _, station := range QueryAll.Data {
		if station.Name == existedStation.Name {
			found = true
		}
	}
	if !found {
		t.Errorf("Request failed, station not found")
	}*/

	/*	StationId       = "id"
		StationNames     = "name"
		StayTime = "stayTime"*/
	randomIndex := rand.Intn(len(QueryAll.Data))
	ctx.Set(StationId, QueryAll.Data[randomIndex].Id)
	ctx.Set(StationNames, QueryAll.Data[randomIndex].Name)
	ctx.Set(StayTime, QueryAll.Data[randomIndex].StayTime)

	return nil, nil
}

func QueryPrice(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	/*	// Query all price configs
		allPriceConfigs, err1 := priceSvc.FindAllPriceConfig()
		if err1 != nil {
			t.Errorf("FindAllPriceConfig failed: %v", err1)
		}
		found := false
		for _, price := range allPriceConfigs.Data {
			if price.Id == existedPrice.Id {
				found = true
			}
		}
		if !found {
			t.Errorf("Request failed, station not found")
		}*/
	// Query price config by route ID and train type
	//println("RouteID: %v, TrainType: %v", ctx.Get(RouteID).(string), ctx.Get(TrainTypeName).(string))
	TheRouteId := ctx.Get(RouteID).(string)
	TheTrainType := ctx.Get(TrainTypeName).(string)
	priceByRouteAndTrain, err := cli.FindByRouteIdAndTrainType(TheRouteId, TheTrainType)
	if err != nil {
		log.Println("FindByRouteIdAndTrainType failed: %v", err)
		return nil, err
	}
	if priceByRouteAndTrain.Status != 1 {
		log.Println("There is not corresponding Ticket available.")
		return &(NodeResult{false}), err
	}
	/*	if priceByRouteAndTrain.Data.Id != ctx.Get(ID) {
		log.Println("priceByRouteAndTrain.Data.Id != existedPrice.Id")
		return nil, err
	}*/

	/*	Id                  string  `json:"id"`
		TrainType           string  `json:"trainType"`
		RouteId             string  `json:"routeId"`
		BasicPriceRate      float64 `json:"basicPriceRate"`
		FirstClassPriceRate float64 `json:"firstClassPriceRate"`*/
	//ctx.Set(TrainType, priceByRouteAndTrain.Data.TrainType)
	//ctx.Set(RouteID, priceByRouteAndTrain.Data.RouteId)
	ctx.Set(BasicPriceRate, priceByRouteAndTrain.Data.BasicPriceRate)
	ctx.Set(FirstClassPriceRate, priceByRouteAndTrain.Data.FirstClassPriceRate)

	return nil, nil
}

func QuerySecurity(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Get All Security Configs
	configs, err3 := cli.FindAllSecurityConfig()
	if err3 != nil {
		log.Println("FindAllSecurityConfig failed: %v", err3)
		return nil, err3
	}
	if configs.Status != 1 {
		log.Println("[Security Service]Status != 1")
		return nil, err3
	}
	/*found := false
	for _, security := range configs.Data {
		if security.ID == existedSecurity.ID &&
			security.Name == existedSecurity.Name &&
			security.Value == existedSecurity.Value &&
			security.Description == existedSecurity.Description {
			found = true
		}
	}
	if !found {
		log.Println("[Security Service]Cannot find existed security config")
		return nil, err3
	}*/

	/*	ID          string `json:"id"`
		Name        string `json:"name"`
		Value       string `json:"value"`
		Description string `json:"description"`*/
	randomIndex := rand.Intn(len(configs.Data))
	ctx.Set(SecurityID, configs.Data[randomIndex].ID)
	ctx.Set(SecurityName, configs.Data[randomIndex].Name)
	ctx.Set(SecurityValue, configs.Data[randomIndex].Value)
	ctx.Set(SecurityDescription, configs.Data[randomIndex].Description)

	return nil, nil
}

// ConfigBehaviorChain
func QueryConfig(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Query All Configs Test
	queryAllResp, err := cli.QueryAllConfigs()
	if err != nil {
		log.Println("QueryAllConfigs request failed, err %s", err)
		return nil, err
	}
	if queryAllResp.Status != 1 {
		log.Println("QueryAllConfigs status != 1")
		return nil, err
	}

	/*	Name        string `json:"name"`
		Value       string `json:"value"`
		Description string `json:"description"`*/
	randomIndex := rand.Intn(len(queryAllResp.Data))
	ctx.Set(ConfigName, queryAllResp.Data[randomIndex].Name)
	ctx.Set(Value, queryAllResp.Data[randomIndex].Value)
	ctx.Set(Description, queryAllResp.Data[randomIndex].Description)

	return nil, nil
}

func QueryOrder(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	Resp, err := cli.ReqFindAllOrder()
	if err != nil {
		log.Println("Request failed, err %s", err)
		return nil, err
	}
	if len(Resp.Data) == 0 {
		log.Println("no data found.")
		return nil, err
	}

	/*	//AccountId              = "accountId"
		BoughtDate             = "boughtDate"
		CoachNumber            = "coachNumber"
		ContactsDocumentNumber = "contactsDocumentNumber"
		ContactsName           = "contactsName"
		DifferenceMoney        = "differenceMoney"
		//DocumentType           = "documentType"
		//From                   = "from"
		//Id                     = "id"
		//Price                  = "price"
		SeatClass  = "seatClass"
		SeatNumber = "seatNumber"
		Status     = "status"
		//To                     = "to"
		TrainNumber = "trainNumber"
		TravelDate  = "travelDate"
		TravelTime  = "travelTime"*/
	randomIndex := rand.Intn(len(Resp.Data))
	//ctx.Set(AccountID, Resp.Data[randomIndex].AccountId)
	ctx.Set(BoughtDate, Resp.Data[randomIndex].BoughtDate)
	ctx.Set(CoachNumber, Resp.Data[randomIndex].CoachNumber)
	ctx.Set(ContactsDocumentNumber, Resp.Data[randomIndex].ContactsDocumentNumber)
	//ctx.Set(ContactsName, Resp.Data[randomIndex].ContactsName)
	ctx.Set(Name, Resp.Data[randomIndex].ContactsName)
	ctx.Set(DifferenceMoney, Resp.Data[randomIndex].DifferenceMoney)
	ctx.Set(SeatClass, Resp.Data[randomIndex].SeatClass)
	ctx.Set(SeatNumber, Resp.Data[randomIndex].SeatNumber)
	ctx.Set(Status, Resp.Data[randomIndex].Status)
	ctx.Set(TrainNumber, Resp.Data[randomIndex].TrainNumber)
	ctx.Set(TravelDate, Resp.Data[randomIndex].TravelDate)
	ctx.Set(TravelTime, Resp.Data[randomIndex].TravelTime)

	return nil, nil
}

func CreateOrder(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	originOrder0 := service.Order{
		AccountId:              ctx.Get(AccountID).(string),
		BoughtDate:             faker.Date(),
		CoachNumber:            generateCoachNumber(),
		ContactsDocumentNumber: generateDocumentNumber(),
		//ContactsName:           ctx.Get(ContactsName).(string),
		ContactsName:    ctx.Get(Name).(string),
		DifferenceMoney: "",
		DocumentType:    0,
		//From:                   ctx.Get(From).(string),
		From: ctx.Get(From).(string), // First, create/query get the station;
		// then put them here -> If you want to create a new Order, you have to do the whole process.
		Id:         "nil",
		Price:      RandomDecimalStringBetween(1, 10),
		SeatClass:  GetTrainTicketClass(),
		SeatNumber: service.GenerateSeatNumber(),
		Status:     0,
		//To:                     ctx.Get(To).(string),
		To:          ctx.Get(To).(string),
		TrainNumber: ctx.Get(TrainNumber).(string),
		TravelDate:  getRandomTime(),
		TravelTime:  generateRandomTime(),
	}

	CreateNewOrderResp, err := cli.ReqCreateNewOrder(&originOrder0)
	if err != nil {
		log.Println("Request failed, err %s", err)
		return nil, err
	}
	if CreateNewOrderResp.Status != 1 {
		log.Println("Request failed, CreateNewOrder status != 1")
		return nil, err
	}

	//ctx.Set(AccountID, Resp.Data[randomIndex].AccountId)
	ctx.Set(BoughtDate, CreateNewOrderResp.Data.BoughtDate)
	ctx.Set(CoachNumber, CreateNewOrderResp.Data.CoachNumber)
	ctx.Set(ContactsDocumentNumber, CreateNewOrderResp.Data.ContactsDocumentNumber)
	//ctx.Set(ContactsName, Resp.Data[randomIndex].ContactsName)
	ctx.Set(Name, CreateNewOrderResp.Data.ContactsName)
	ctx.Set(DifferenceMoney, CreateNewOrderResp.Data.DifferenceMoney)
	ctx.Set(SeatClass, CreateNewOrderResp.Data.SeatClass)
	ctx.Set(SeatNumber, CreateNewOrderResp.Data.SeatNumber)
	ctx.Set(Status, CreateNewOrderResp.Data.Status)
	ctx.Set(TrainNumber, CreateNewOrderResp.Data.TrainNumber)
	ctx.Set(TravelDate, CreateNewOrderResp.Data.TravelDate)
	ctx.Set(TravelTime, CreateNewOrderResp.Data.TravelTime)

	return nil, nil
}

func QueryOrderOther(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	GetResp, err := cli.ReqFindAllOrderOther()

	if err != nil {
		log.Println("Request failed, err %s", err)
		return nil, err
	}
	if GetResp.Status != 1 {
		log.Println("Request failed, CreateNewOrder status != 1")
		return nil, err
	}

	randomIndex := rand.Intn(len(GetResp.Data))
	//ctx.Set(AccountID, Resp.Data[randomIndex].AccountId)
	ctx.Set(BoughtDate, GetResp.Data[randomIndex].BoughtDate)
	ctx.Set(CoachNumber, GetResp.Data[randomIndex].CoachNumber)
	ctx.Set(ContactsDocumentNumber, GetResp.Data[randomIndex].ContactsDocumentNumber)
	//ctx.Set(ContactsName, Resp.Data[randomIndex].ContactsName)
	ctx.Set(Name, GetResp.Data[randomIndex].ContactsName)
	ctx.Set(DifferenceMoney, GetResp.Data[randomIndex].DifferenceMoney)
	ctx.Set(SeatClass, GetResp.Data[randomIndex].SeatClass)
	ctx.Set(SeatNumber, GetResp.Data[randomIndex].SeatNumber)
	ctx.Set(Status, GetResp.Data[randomIndex].Status)
	ctx.Set(TrainNumber, GetResp.Data[randomIndex].TrainNumber)
	ctx.Set(TravelDate, GetResp.Data[randomIndex].TravelDate)
	ctx.Set(TravelTime, GetResp.Data[randomIndex].TravelTime)

	return nil, nil
}

func CreateOrderOther(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	AddResp, err := cli.ReqCreateNewOrderOther(&service.Order{
		AccountId:              ctx.Get(AccountID).(string),
		BoughtDate:             faker.Date(),
		CoachNumber:            generateCoachNumber(),
		ContactsDocumentNumber: generateDocumentNumber(),
		//ContactsName:           ctx.Get(ContactsName).(string),
		ContactsName:    ctx.Get(Name).(string),
		DifferenceMoney: "",
		DocumentType:    0,
		//From:                   ctx.Get(From).(string),
		From: ctx.Get(From).(string), // First, create/query get the station;
		// then put them here -> If you want to create a new Order, you have to do the whole process.
		Id:         "nil",
		Price:      RandomDecimalStringBetween(1, 10),
		SeatClass:  GetTrainTicketClass(),
		SeatNumber: service.GenerateSeatNumber(),
		Status:     0,
		//To:                     ctx.Get(To).(string),
		To:          ctx.Get(To).(string),
		TrainNumber: ctx.Get(TrainNumber).(string),
		TravelDate:  getRandomTime(),
		TravelTime:  generateRandomTime(),
	})

	if err != nil {
		log.Println("Request failed, err %s", err)
		return nil, err
	}
	if AddResp.Status != 1 {
		log.Println("Request failed, CreateNewOrder status != 1")
		return nil, err
	}

	//ctx.Set(AccountID, Resp.Data[randomIndex].AccountId)
	ctx.Set(BoughtDate, AddResp.Data.BoughtDate)
	ctx.Set(CoachNumber, AddResp.Data.CoachNumber)
	ctx.Set(ContactsDocumentNumber, AddResp.Data.ContactsDocumentNumber)
	//ctx.Set(ContactsName, Resp.Data[randomIndex].ContactsName)
	ctx.Set(Name, AddResp.Data.ContactsName)
	ctx.Set(DifferenceMoney, AddResp.Data.DifferenceMoney)
	ctx.Set(SeatClass, AddResp.Data.SeatClass)
	ctx.Set(SeatNumber, AddResp.Data.SeatNumber)
	ctx.Set(Status, AddResp.Data.Status)
	ctx.Set(TrainNumber, AddResp.Data.TrainNumber)
	ctx.Set(TravelDate, AddResp.Data.TravelDate)
	ctx.Set(TravelTime, AddResp.Data.TravelTime)

	return nil, nil
}

/////////////////////////////--- Preserve Behaviors ---///////////////////////////////
/////////////////////////////--- Preserve Behaviors ---///////////////////////////////
/////////////////////////////--- Preserve Behaviors ---///////////////////////////////

// Preserve Behaviors - The Last One
func Preserve(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	OrderTicketsInfo := service.OrderTicketsInfo{
		AccountID:       ctx.Get(AccountID).(string),  // Query:Create = 1 : 0
		ContactsID:      ctx.Get(ContactsID).(string), // Query:Create = 1 : 0
		TripID:          ctx.Get(TripId).(string),
		SeatType:        ctx.Get(SeatClass).(int),
		LoginToken:      ctx.Get(LoginToken).(string),
		Date:            extractDate(ctx.Get(DepartureTime).(string)),
		From:            ctx.Get(From).(string),
		To:              ctx.Get(To).(string),
		Assurance:       ctx.Get(AssuranceTypeIndex).(int),
		FoodType:        ctx.Get(FoodType).(int),
		StationName:     ctx.Get(StationName).(string),
		StoreName:       ctx.Get(StoreName).(string),
		FoodName:        ctx.Get(FoodName).(string),
		FoodPrice:       ctx.Get(FoodPrice).(float64),
		HandleDate:      extractDate(ctx.Get(HandleDate).(string)),
		ConsigneeName:   ctx.Get(Consignee).(string),
		ConsigneePhone:  ctx.Get(Phone).(string),
		ConsigneeWeight: ctx.Get(Weight).(float64),
		IsWithin:        ctx.Get(IsWithin).(bool),
	}
	PreserveResp, err := cli.Preserve(&OrderTicketsInfo)
	if err != nil {
		return nil, err
	}
	if PreserveResp.Status != 1 {
		return nil, fmt.Errorf("preserve order tickets fail. PreserveResp.Status != 1, get %v", PreserveResp.Status)
	}
	log.Println("The Status is: %v, and PreserveResp Data: %v\n", PreserveResp.Status, PreserveResp.Data)
	log.Println("PreserveBehaviors(Chain) Ends. End time: %v", time.Now().String())

	//return nil, nil
	return &(NodeResult{false}), nil // Chain End :D
}
