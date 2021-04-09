package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

type Product struct {
	SerialNo        string `json:"SerialNo"`
	BatchId         string `json:"batchId"`
	ModelNo         string `json:"modelNo"`
	ModelName       string `json:"modelName"`
	Date            string `json:"date"`
	CftId           string `json:"cftId"`
	Price           string `json:"price"`
	Spec            string `json:"spec"`
	Status          string `json:"status"`
	Ownership       string `json:"ownership"`
	DealerId        string `json:"dealerId"`
	DealerName      string `json:"dealerName"`
	Insurance       string `json:"insurance"`
	InsuranceId     string `json:"insuranceId"`
	InsuranceExpiry string `json:"insuranceExpiry"`
	ServiceId       string `json:"serviceId"`
	ServiceHistory  string `json:"serviceHistory"`
	SerCentre       string `json:"serCentre"`
	DateOfService   string `json:"dateOfService"`
	ServiceDetails  string `json:"serviceDetails"`
}

type Service struct {
	ServiceId      string `json:"serviceId"`
	SerCentre      string `json:"serCentre"`
	DateOfService  string `json:"dateOfService"`
	ServiceDetails string `json:"serviceDetails"`
}

type AllWatches struct {
	AllWatches []string
}

type AllWatchesDetails struct {
	Watches []Product
}

//type TimeTracker struct{
//	DispachedTime	string
//	ReachedTime	string
//	TPTemprature	string
//}

func (s *SmartContract) AddProduct(ctx contractapi.TransactionContextInterface, productData string) (string, error) {

	if len(productData) == 0 {
		return "", fmt.Errorf("Please pass the correct productData")
	}

	var product Product
	err := json.Unmarshal([]byte(productData), &product)
	if err != nil {
		return "", fmt.Errorf("Failed while unmarshling product records %s", err.Error())
	}

	productAsBytes, err := json.Marshal(product)
	if err != nil {
		return "", fmt.Errorf("Failed while marshling product records %s", err.Error())
	}

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(product.SerialNo, productAsBytes)

	//update all batches to the array whose key is "AllWatches"

	AllWatchesAsBytes, err := ctx.GetStub().GetState("AllWatches")

	var allWatches AllWatches

	err = json.Unmarshal(AllWatchesAsBytes, &allWatches)
	allWatches.AllWatches = append(allWatches.AllWatches, product.SerialNo)

	allwatchesAsBytes, _ := json.Marshal(allWatches)
	err = ctx.GetStub().PutState("AllWatches", allwatchesAsBytes)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	//-----------------------------------------------------------
	return "", nil

}

//-----------------------------------------------------------

// TransferAsset updates the owner field of asset with given id in world state.
func (s *SmartContract) TransferMtoD(ctx contractapi.TransactionContextInterface, args []string) error {
	// must be an invoke
	var err error
	var product Product
	bAsBytes, err := ctx.GetStub().GetState(args[0])

	err = json.Unmarshal(bAsBytes, &product)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	product.DealerId = args[1]
	product.DealerName = args[2]
	product.Date = args[3]
	product.Status = "Dealer"
	product.Ownership = args[2]

	//Commit updates batch to ledger
	btAsBytes, _ := json.Marshal(product)
	err = ctx.GetStub().PutState(product.SerialNo, btAsBytes)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}
func (s *SmartContract) transferDtoC(ctx contractapi.TransactionContextInterface, args []string) error {
	// must be an invoke
	var err error
	var product Product
	bAsBytes, err := ctx.GetStub().GetState(args[0])

	err = json.Unmarshal(bAsBytes, &product)

	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}

	product.Date = args[1]
	product.Insurance = args[2]
	product.InsuranceId = args[5]
	product.Status = "Customer"
	product.Ownership = args[3]
	product.InsuranceExpiry = args[4] //expiry date field updated

	//Commit updates batch to ledger
	btAsBytes, _ := json.Marshal(product)
	err = ctx.GetStub().PutState(product.SerialNo, btAsBytes)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (s *SmartContract) serviceUpdate(ctx contractapi.TransactionContextInterface, args []string) error {
	// must be an invoke
	var err error
	var service Service
	var product Product
	productAsBytes, err := ctx.GetStub().GetState(args[0])

	err = json.Unmarshal(productAsBytes, &product)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	product.ServiceId = args[1]
	product.ServiceHistory = args[2]
	product.SerCentre = args[3]
	product.DateOfService = args[4]
	product.ServiceDetails = args[5]

	// must be an invoke

	//serviceAsBytes, err := stub.GetState(args[1])

	//err = json.Unmarshal(serviceAsBytes, &service)

	//if err != nil {
	//	return shim.Error(err.Error())
	//}

	service.ServiceId = args[1]
	service.SerCentre = args[3]
	service.DateOfService = args[4]
	service.ServiceDetails = args[5]

	//Commit updates batch to ledger
	stAsBytes, _ := json.Marshal(service)
	err = ctx.GetStub().PutState(product.SerialNo+service.ServiceId, stAsBytes)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	//Commit updates batch to ledger
	ptAsBytes, _ := json.Marshal(product)
	err = ctx.GetStub().PutState(product.SerialNo, ptAsBytes) //ServiceId or SerialNo
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

// ============================================================================================================================
// Get All Batches Details for Transporter
// ============================================================================================================================
func (s *SmartContract) GetAllWatches(ctx contractapi.TransactionContextInterface, args []string) error {

	//get the AllBatches index
	var owner string
	owner = args[0]
	fmt.Printf("Value of Owner: %s", owner)
	allBAsBytes, _ := ctx.GetStub().GetState("AllWatches")

	var res AllWatches
	json.Unmarshal(allBAsBytes, &res)

	var rab AllWatchesDetails

	for i := range res.AllWatches {

		sbAsBytes, _ := ctx.GetStub().GetState(res.AllWatches[i])

		var sb Product
		json.Unmarshal(sbAsBytes, &sb)

		if sb.Status == owner {
			fmt.Printf("Value of Owner-1: %s", sb.Status)
			rab.Watches = append(rab.Watches, sb)
		}
	}
	rabAsBytes, _ := json.Marshal(rab)
	return ctx.GetStub().PutState(args[0], rabAsBytes)
}

//End of changing the Batch ID

/// Query callback representing the query of a chaincode
func (s *SmartContract) Query(ctx contractapi.TransactionContextInterface, args []string) error {

	var serl string // Entities
	var err error

	if len(args) != 1 {
		return fmt.Errorf("Incorrect number of arguments. Expecting name of the person to query")
	}

	serl = args[0]

	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetState(serl)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + serl + "\"}"
		return fmt.Errorf(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + serl + "\"}"
		return fmt.Errorf(jsonResp)
	}

	//	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	//	logger.Infof("Query Response:%s\n", jsonResp)
	return ctx.GetStub().PutState(args[0], Avalbytes)

}

//Query to get the history of the BAtchID

func (s *SmartContract) GetProductHistory(ctx contractapi.TransactionContextInterface, args []string) error {

	fmt.Printf("In getproducthistory Function")

	if len(args) < 1 {
		fmt.Printf("In getproducthistory Error Function")
		return fmt.Errorf("Incorrect number of arguments. Expecting 1")
	}

	SerialNo := args[0]

	fmt.Printf("- start getHistoryForWatch: %s\n", SerialNo)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(SerialNo)

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer

	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForMarble returning:\n%s\n", buffer.String())

	return ctx.GetStub().PutState(args[0], buffer.Bytes())
}

// ============================================================================================================================

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create luxury watch chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting luxury watch chaincode: %s", err.Error())
	}
}
