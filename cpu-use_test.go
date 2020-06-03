package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shimtest"
)

func TestBaseFunctinality(t *testing.T){
	simpleChaincode := new(SimpleChaincode)
	mockStub := shimtest.NewMockStub("Test Feature AddCPU", simpleChaincode)
	testTable := []struct{
		name	string
		assetName	string
		functionName	string
		txID	string
	}{
		{"Adding test","test1","AddCPU","1"},
		{"Get test","test1","GetUsage","2"},
	}

	for _, testCase := range testTable{
		t.Run(testCase.name,func(t *testing.T){
		res := mockStub.MockInvoke(testCase.txID,[][]byte{[]byte(testCase.functionName),[]byte(testCase.assetName)})
		if res.Status != 200{
			t.Fatalf("There was an error status = %v : msg = %v",res.Status,res.Message)
		}
		fmt.Println(string(res.Payload))
		if res.Payload == nil{
			t.Fatalf("Its nil")
		}
		})
	}
}

func TestGetHistory(t *testing.T){
	simpleChaincode := new(SimpleChaincode)
	mockStub := shimtest.NewMockStub("Test Feature AddCPU", simpleChaincode)

	err, res := mockStub.GetHistoryForKey("test1")

	if err != nil{
		t.Errorf("The error is %v",err)
	}

	if res == nil{
		t.Errorf("Its empty")
	}
}