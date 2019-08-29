package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

// Define the DocRecord Structure, which holds the signature of the document
// signed by issuer, and the time when this record is created
type DocRecord struct {
	TrxUid string `json:"uid"`
	Hash   string `json:"hash"`
	Time   string `json:"time"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryDocRecord" {
		return s.queryDocRecord(APIstub, args)
	} else if function == "createDocRecord" {
		result, err := s.createDocRecord(APIstub, args)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Return the result as success payload
		return shim.Success([]byte(result))
	} else if function == "hashDocRecord" {
		return s.hashDocRecord(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryDocRecord(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	docrecordAsBytes, _ := APIstub.GetState(args[0])

	if docrecordAsBytes == nil {
		return shim.Error("Document not found: " + args[0])
	}

	return shim.Success(docrecordAsBytes)
}

func (s *SmartContract) createDocRecord(APIstub shim.ChaincodeStubInterface, args []string) (string, error) {

	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect number of arguments. Expecting 1")
	}

	hashBytes := sha256.Sum256([]byte(args[0]))
	hashStr := string(hashBytes[:])

	if len(hashStr) == 0 {
		return "", fmt.Errorf("hashStr was not calculated: %s", args[0])
	}
	trxId := APIstub.GetTxID()

	if len(trxId) == 0 {
		return "", fmt.Errorf("trxId was not calculated: %s", args[0])
	}

	var docrecord = DocRecord{TrxUid: trxId, Hash: hashStr, Time: time.Now().Format(time.RFC3339)}
	docrecordAsBytes, _ := json.Marshal(docrecord)
	APIstub.PutState(trxId, docrecordAsBytes)

	return string(trxId), nil
}

func (s *SmartContract) hashDocRecord(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	hashBytes := sha256.Sum256([]byte(args[0]))
	return shim.Success(hashBytes[:])
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}