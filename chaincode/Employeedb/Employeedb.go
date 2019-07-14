
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	
        "github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)


type SmartContract struct {

}

type Emp_db struct {
	EmpID         int `json:"empID"` 
        Name          string `json:"name"`   
	Phone         int `json:"phone"`
	Email         string `json:"email"`
	Address       string `json:"address"`
        Designation   string `json:"designation"`
}

// Init initializes chaincode
// ===========================

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
        
	return shim.Success(nil)
}


// Invocation through Client Application
// =====================================

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	
        function, args := APIstub.GetFunctionAndParameters()
        fmt.Println("invoke is running " + function)
       
        // Ledger Initiation 
        //==================
        
        if function == "initLedger"{
           return s.initLedger(APIstub)
         }
        
        // Add new Employee to the db
        // ==========================
        
        if function == "newEmployee"{
           return s.newEmployee(APIstub,args)
        }

        // Get all Employees Info
        // ======================
        
        if function == "queryAllEmployees"{
            return s.queryAllEmployees(APIstub)
        } 
        
        // Get Employee Info by ID
        // =======================
         
        if function == "queryByID"{
            return s.queryByID(APIstub,args)
        }

        // Delete EMPLOYEE data from db
        // ============================
        
        if function == "DeleteEmployee"{
            return s.DeleteEmployee(APIstub,args)
        }
        
        // Update Employee Phone Number
        // ============================
        
         if function == "UpdateEmployeePhone"{
             return s.UpdateEmployeePhone(APIstub,args)
        }
        
        // History of Employee
        // ===================
        
        if function == "HistoryOfEmployees"{
             return s.HistoryOfEmployees(APIstub,args)
        }

        // Get Info by range
        // =================
        
        if function == "GetInfobyRange"{
             return s.GetInfobyRange(APIstub,args)
        }

        fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// Putting initial DATA int db
// ===========================

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
Emp := []Emp_db{ 
       
		Emp_db{EmpID:1, Name: "Brad"   ,  Phone: 9823553813 , Email  : "Brad@gmail.com"   ,Address:"38 Wall Street"      , Designation:"Manager"},
		Emp_db{EmpID:2, Name: "Jin Soo",  Phone: 9922327498 , Email  : "JinSoo@gmail.com" ,Address:"42 Mountain View"     , Designation:"Officer"},
		Emp_db{EmpID:3, Name: "Max"    ,  Phone: 9212327498, Email  : "Max@gmail.com"    ,Address:"254 St.joseph Church" ,Designation:"Asst. Officer"},
		Emp_db{EmpID:4, Name: "Adriana",  Phone: 9926727498, Email : "Adriana@gmail.com" ,Address:"23 Maria Hill"        , Designation:"Clerk"},
		Emp_db{EmpID:5, Name: "Michel" ,  Phone: 9678327498, Email : "Michel@gmail.com"  ,Address:"78 Churchgate"        , Designation:"Peon"},
       }		
       i := 0
       
        for i < len(Emp) {
		fmt.Println("i is ", i)
		EmpAsBytes, _ := json.Marshal(Emp[i])
		APIstub.PutState(strconv.Itoa(Emp[i].EmpID), EmpAsBytes)            //------PutState feeds into the database and we are walso assosciating each employee with id
		fmt.Println("Added", Emp[i])
		i = i + 1
	}

	return shim.Success(nil)
}

// ADD new EMPLOYEE in db
// ======================

func (s *SmartContract) newEmployee(APIstub shim.ChaincodeStubInterface,args[] string) sc.Response {
              
        if len(args)!=6 {
        return shim.Error("Sorry! Less arguments.Expecting 6.")
        } 

        mobile_no,_ := strconv.Atoi(args[2])
        empid,_ := strconv.Atoi(args[0])
        var employee = Emp_db{ EmpID: empid, Name: args[1], Phone: mobile_no,  Email: args[3], Address: args[4], Designation: args[5]}

	employeeAsBytes, _ := json.Marshal(employee)
	APIstub.PutState(args[0], employeeAsBytes)

	return shim.Success(nil)

}

// QUERY all employees
// ===================

func (s *SmartContract) queryAllEmployees(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "1"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllEmployees:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// QUERY employee by ID
// ====================

func (s *SmartContract) queryByID(APIstub shim.ChaincodeStubInterface,args[] string) sc.Response {

        if len(args)!=1 {
                 return shim.Error("Wrong Number of Arguments.Expecting 1")
        }
        employeeAsBytes, _ :=APIstub.GetState(args[0])
        return  shim.Success(employeeAsBytes)       
}

// DELETE employee form db
// =======================

func (s *SmartContract) DeleteEmployee(APIstub shim.ChaincodeStubInterface,args[] string) sc.Response {
        if len(args)!=1 {
                 return shim.Error("Wrong Number of Arguments.Expecting 1")
        }
        
        
        APIstub.DelState(args[0]) 
	
        return shim.Success(nil)
}

// Update Employee Phone
// =====================

func (s *SmartContract) UpdateEmployeePhone(APIstub shim.ChaincodeStubInterface, args[] string) sc.Response {
        if len(args)!=2 {
                 return shim.Error("Wrong Number of Arguments.Expecting 2")
        }
        EmployeeAsBytes,_:=APIstub.GetState(args[0])
        employee := Emp_db{}

	json.Unmarshal(EmployeeAsBytes, &employee)
	employee.Phone,_ = strconv.Atoi(args[1])

	EmployeeAsBytes, _ = json.Marshal(employee)
	APIstub.PutState(args[0],EmployeeAsBytes)

	return shim.Success(nil)
}


// History of Employee
// ===================

func (t *SmartContract) HistoryOfEmployees(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	

	fmt.Printf("- start getHistoryForEmployee: %s\n", args[0])

	resultsIterator, err := stub.GetHistoryForKey(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForEmployee returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

//------Get Info by Range--------

func (s *SmartContract) GetInfobyRange(APIstub shim.ChaincodeStubInterface,args[] string) sc.Response {

	startKey := args[0]
	endKey := args[1]

	resultsIterator,err := APIstub.GetStateByRange(startKey,endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
        
        var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		

		buffer.WriteString(", \"Record\":")
		
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- GetInfobyRange:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}



// ------------MAIN---------------------

func main() {

        err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

