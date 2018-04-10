package main

import (
	"fmt"
	"time"
	"math/rand"
	"encoding/json"
	"strconv"
	"bytes"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	
	//"github.com/hyperledger/fabric/common/util"
)


type Enrollment struct{
								// Attributes of Enrollment
	Participant_id string `json:"Participant_id"`
	Name string `json:"Name"`
	Role string `json:"Role"`
	Email string `json:"Email"`
	Enrollment_timestamp string `json:"Enrollment_timestamp"`
	First_line_of_address string `json:"First_line_of_address"`
	Second_line_of_address string `json:"Second_line_of_address"`
	City string `json:"City"`
	Province string `json:"Province"`
	Postcode string `json:"Postcode"`
	Country string `json:"Country"`

}

type Register_G_P struct{								// Attributes of a Register_G_P
	G_P_registration_id string `json:"G_P_registration_id"`
	Participant_id string `json:"Participant_id"`
	Type_of_plant string `json:"Type_of_plant"`
	Daily_injection_capacity int `json:"Daily_injection_capacity"`
	Installation_month string `json:"Installation_month"`
	Registration_timestamp string `json:"Registration_timestamp"`
	Remaining_token int `json:"Remaining_token"`
	Start_date string `json:"Start_date"`
	End_date string `json:"End_date"`
	S_0000_0030 int `json:"S_0000_0030"`
	S_0030_0100 int `json:"S_0030_0100"`
	S_0100_0130 int `json:"S_0100_0130"`
	S_0130_0200 int `json:"S_0130_0200"`
	S_0200_0230 int `json:"S_0200_0230"`
	S_0230_0300 int `json:"S_0230_0300"`
	S_0300_0330 int `json:"S_0300_0330"`
	S_0330_0400 int `json:"S_0330_0400"`
	S_0400_0430 int `json:"S_0400_0430"`
	S_0430_0500 int `json:"S_0430_0500"`
	S_0500_0530 int `json:"S_0500_0530"`
	S_0530_0600 int `json:"S_0530_0600"`
	S_0600_0630 int `json:"S_0600_0630"`
	S_0630_0700 int `json:"S_0630_0700"`
	S_0700_0730 int `json:"S_0700_0730"`
	S_0730_0800 int `json:"S_0730_0800"`
	S_0800_0830 int `json:"S_0800_0830"`
	S_0830_0900 int `json:"S_0830_0900"`
	S_0900_0930 int `json:"S_0900_0930"`
	S_0930_1000 int `json:"S_0930_1000"`
	S_1000_1030 int `json:"S_1000_1030"`
	S_1030_1100 int `json:"S_1030_1100"`
	S_1100_1130 int `json:"S_1100_1130"`
	S_1130_1200 int `json:"S_1130_1200"`
	S_1200_1230 int `json:"S_1200_1230"`
	S_1230_1300 int `json:"S_1230_1300"`
	S_1300_1330 int `json:"S_1300_1330"`
	S_1330_1400 int `json:"S_1330_1400"`
	S_1400_1430 int `json:"S_1400_1430"`
	S_1430_1500 int `json:"S_1430_1500"`
	S_1500_1530 int `json:"S_1500_1530"`
	S_1530_1600 int `json:"S_1530_1600"`
	S_1600_1630 int `json:"S_1600_1630"`
	S_1630_1700 int `json:"S_1630_1700"`
	S_1700_1730 int `json:"S_1700_1730"`
	S_1730_1800 int `json:"S_1730_1800"`
	S_1800_1830 int `json:"S_1800_1830"`
	S_1830_1900 int `json:"S_1830_1900"`
	S_1900_1930 int `json:"S_1900_1930"`
	S_1930_2000 int `json:"S_1930_2000"`
	S_2000_2030 int `json:"S_2000_2030"`
	S_2030_2100 int `json:"S_2030_2100"`
	S_2100_2130 int `json:"S_2100_2130"`
	S_2130_2200 int `json:"S_2130_2200"`
	S_2200_2230 int `json:"S_2200_2230"`
	S_2230_2300 int `json:"S_2230_2300"`
	S_2300_2330 int `json:"S_2300_2330"`
	S_2330_0000 int `json:"S_2330_0000"`
	Wallet int `json:"Wallet"`
}


type Register_R_C struct{
								// Attributes of a Register_R_C
	R_C_registration_id string `json:"R_C_registration_id"`
	Participant_id string `json:"Participant_id"`
	Type_of_plant string `json:"Type_of_plant"`
	Installation_month string `json:"Installation_month"`
	Registration_timestamp string `json:"Registration_timestamp"`
	Required_energy_token int `json:"Required_energy_token"`
	Remaining_token int `json:"Remaining_token"`
	Start_date string `json:"Start_date"`
	End_date string `json:"End_date"`
	Wallet int `json:"Wallet"`
	Regulator string `json:"Regulator"`
}


type RegisterCustomer struct{
								// Attributes of a RegisterCustomer
	Cregistration_id string `json:"Cregistration_id"`
	ConsumerParticipantID string `json:"ConsumerParticipantID"`
	PlantType string `json:"PlantType"`
	InstallationMonth string `json:"InstallationMonth"`
	Shift1Total int `json:"Shift1Total"`
	Shift2Total int `json:"Shift2Total"`
	Shift3Total int `json:"Shift3Total"`
	Shift4Total int `json:"Shift4Total"`
	Shift1 int `json:"Shift1"`
	Shift2 int `json:"Shift2"`
	Shift3 int `json:"Shift3"`
	Shift4 int `json:"Shift4"`
	Prosumer_participant_id string `json:"Prosumer_participant_id"`
	Prosumer_registration_id string `json:"Prosumer_registration_id"`
	Retailer_participant_id string `json:"Retailer_participant_id"`
	Retailer_registration_id string `json:"Retailer_registration_id"`
	Start_date string `json:"Start_date"`
	End_date string `json:"End_date"`
	Token_remuneration_type string `json:"Token_remuneration_type"`
	Value string `json:"Value"`
	InjectionCapacity int `json:"InjectionCapacity"`
	S_0000_0030Selected bool `json:"S_0000_0030Selected"`
	S_0030_0100Selected bool `json:"S_0030_0100Selected"`
	S_0100_0130Selected bool `json:"S_0100_0130Selected"`
	S_0130_0200Selected bool `json:"S_0130_0200Selected"`
	S_0200_0230Selected bool `json:"S_0200_0230Selected"`
	S_0230_0300Selected bool `json:"S_0230_0300Selected"`
	S_0300_0330Selected bool `json:"S_0300_0330Selected"`
	S_0330_0400Selected bool `json:"S_0330_0400Selected"`
	S_0400_0430Selected bool `json:"S_0400_0430Selected"`
	S_0430_0500Selected bool `json:"S_0430_0500Selected"`
	S_0500_0530Selected bool `json:"S_0500_0530Selected"`
	S_0530_0600Selected bool `json:"S_0530_0600Selected"`
	S_0600_0630Selected bool `json:"S_0600_0630Selected"`
	S_0630_0700Selected bool `json:"S_0630_0700Selected"`
	S_0700_0730Selected bool `json:"S_0700_0730Selected"`
	S_0730_0800Selected bool `json:"S_0730_0800Selected"`
	S_0800_0830Selected bool `json:"S_0800_0830Selected"`
	S_0830_0900Selected bool `json:"S_0830_0900Selected"`
	S_0900_0930Selected bool `json:"S_0900_0930Selected"`
	S_0930_1000Selected bool `json:"S_0930_1000Selected"`
	S_1000_1030Selected bool `json:"S_1000_1030Selected"`
	S_1030_1100Selected bool `json:"S_1030_1100Selected"`
	S_1100_1130Selected bool `json:"S_1100_1130Selected"`
	S_1130_1200Selected bool `json:"S_1130_1200Selected"`
	S_1200_1230Selected bool `json:"S_1200_1230Selected"`
	S_1230_1300Selected bool `json:"S_1230_1300Selected"`
	S_1300_1330Selected bool `json:"S_1300_1330Selected"`
	S_1330_1400Selected bool `json:"S_1330_1400Selected"`
	S_1400_1430Selected bool `json:"S_1400_1430Selected"`
	S_1430_1500Selected bool `json:"S_1430_1500Selected"`
	S_1500_1530Selected bool `json:"S_1500_1530Selected"`
	S_1530_1600Selected bool `json:"S_1530_1600Selected"`
	S_1600_1630Selected bool `json:"S_1600_1630Selected"`
	S_1630_1700Selected bool `json:"S_1630_1700Selected"`
	S_1700_1730Selected bool `json:"S_1700_1730Selected"`
	S_1730_1800Selected bool `json:"S_1730_1800Selected"`
	S_1800_1830Selected bool `json:"S_1800_1830Selected"`
	S_1830_1900Selected bool `json:"S_1830_1900Selected"`
	S_1900_1930Selected bool `json:"S_1900_1930Selected"`
	S_1930_2000Selected bool `json:"S_1930_2000Selected"`
	S_2000_2030Selected bool `json:"S_2000_2030Selected"`
	S_2030_2100Selected bool `json:"S_2030_2100Selected"`
	S_2100_2130Selected bool `json:"S_2100_2130Selected"`
	S_2130_2200Selected bool `json:"S_2130_2200Selected"`
	S_2200_2230Selected bool `json:"S_2200_2230Selected"`
	S_2230_2300Selected bool `json:"S_2230_2300Selected"`
	S_2300_2330Selected bool `json:"S_2300_2330Selected"`
	S_2330_0000Selected bool `json:"S_2330_0000Selected"`
	Wallet int `json:"Wallet"` 
}


type RegisterRegulator struct{
								// Attributes of a Register_R_C
	Rregistration_id string `json:"Rregistration_id"`
	Participant_id string `json:"Participant_id"`
	Type_of_plant string `json:"Type_of_plant"`
	Installation_month string `json:"Installation_month"`
	Registration_timestamp string `json:"Registration_timestamp"`
	Required_energy_token int `json:"Required_energy_token"`
	Remaining_token int `json:"Remaining_token"`
	Start_date string `json:"Start_date"`
	End_date string `json:"End_date"`
	Wallet int `json:"Wallet"`	
}


type Token struct{
								// Attributes of a Token
	Transaction_id string `json:"Transaction_id"`
	Sender_Registration_id string `json:"Sender_Registration_id"`
	Reciever_Registration_id string `json:"Reciever_Registration_id"`
	Status string `json:"Status"`	
	Prosumer_Registration_id string `json:"Prosumer_Registration_id"`
	Prosumer_Participant_id string `json:"Prosumer_Participant_id"`
	Retailer_Registration_id string `json:"Retailer_Registration_id"`
	Retailer_Participant_id string `json:"Retailer_Participant_id"`
	Customer_Registration_id string `json:"Customer_Registration_id"`
	Street string `json:"Street"`
	City string `json:"City"`
	Postcode string `json:"Postcode"`
	Country string `json:"Country"`
	Time_slot string `json:"Time_slot"`
	Agreement_id string `json:"Agreement_id"`
	Kw int `json:"Kw"`
	No_of_Token int `json:"No_of_Token"`
	Token_remuneration_type string `json:"Token_remuneration_type"`
	Token_remuneration_value string `json:"Token_remuneration_value"`
	Generate_Time_stamp int `json:"Generate_Time_stamp"`
	Wallet_Time_stamp int `json:"Wallet_Time_stamp"`
	Regulator string `json:"Regulator"`
}




type Sign_up struct{
								// Attributes of Sign up Agreement
	Agreement_id string `json:"Agreement_id"`
	Prosumer_participant_id string `json:"Prosumer_participant_id"`
	Prosumer_registration_id string `json:"Prosumer_registration_id"`
	Retailer_participant_id string `json:"Retailer_participant_id"`
	Retailer_registration_id string `json:"Retailer_registration_id"`
	Signed_up_energy_token int `json:"Signed_up_energy_token"`
	Time_period int `json:"Time_period"`
	Total_signed_up int `json:"Total_signed_up"`
	Start_date string `json:"Start_date"`
	End_date string `json:"End_date"`
	Token_remuneration_type string `json:"Token_remuneration_type"`
	Value string `json:"Value"`
	S_0000_0030 int `json:"S_0000_0030"`
	S_0030_0100 int `json:"S_0030_0100"`
	S_0100_0130 int `json:"S_0100_0130"`
	S_0130_0200 int `json:"S_0130_0200"`
	S_0200_0230 int `json:"S_0200_0230"`
	S_0230_0300 int `json:"S_0230_0300"`
	S_0300_0330 int `json:"S_0300_0330"`
	S_0330_0400 int `json:"S_0330_0400"`
	S_0400_0430 int `json:"S_0400_0430"`
	S_0430_0500 int `json:"S_0430_0500"`
	S_0500_0530 int `json:"S_0500_0530"`
	S_0530_0600 int `json:"S_0530_0600"`
	S_0600_0630 int `json:"S_0600_0630"`
	S_0630_0700 int `json:"S_0630_0700"`
	S_0700_0730 int `json:"S_0700_0730"`
	S_0730_0800 int `json:"S_0730_0800"`
	S_0800_0830 int `json:"S_0800_0830"`
	S_0830_0900 int `json:"S_0830_0900"`
	S_0900_0930 int `json:"S_0900_0930"`
	S_0930_1000 int `json:"S_0930_1000"`
	S_1000_1030 int `json:"S_1000_1030"`
	S_1030_1100 int `json:"S_1030_1100"`
	S_1100_1130 int `json:"S_1100_1130"`
	S_1130_1200 int `json:"S_1130_1200"`
	S_1200_1230 int `json:"S_1200_1230"`
	S_1230_1300 int `json:"S_1230_1300"`
	S_1300_1330 int `json:"S_1300_1330"`
	S_1330_1400 int `json:"S_1330_1400"`
	S_1400_1430 int `json:"S_1400_1430"`
	S_1430_1500 int `json:"S_1430_1500"`
	S_1500_1530 int `json:"S_1500_1530"`
	S_1530_1600 int `json:"S_1530_1600"`
	S_1600_1630 int `json:"S_1600_1630"`
	S_1630_1700 int `json:"S_1630_1700"`
	S_1700_1730 int `json:"S_1700_1730"`
	S_1730_1800 int `json:"S_1730_1800"`
	S_1800_1830 int `json:"S_1800_1830"`
	S_1830_1900 int `json:"S_1830_1900"`
	S_1900_1930 int `json:"S_1900_1930"`
	S_1930_2000 int `json:"S_1930_2000"`
	S_2000_2030 int `json:"S_2000_2030"`
	S_2030_2100 int `json:"S_2030_2100"`
	S_2100_2130 int `json:"S_2100_2130"`
	S_2130_2200 int `json:"S_2130_2200"`
	S_2200_2230 int `json:"S_2200_2230"`
	S_2230_2300 int `json:"S_2230_2300"`
	S_2300_2330 int `json:"S_2300_2330"`
	S_2330_0000 int `json:"S_2330_0000"`
	Customer string `json:"Customer"` 
}


type energyToken struct {
}

var Enroll_Formindexstr = "_Enrollforindex"

var Register_G_P_Formindexstr = "_RegisterGPforindex"

var Register_R_C_Formindexstr = "_RegisterRCforindex"

var CustomerFormindexstr = "_Customerindex"

var SignUp_Formindexstr = "_SignUpforindex"

var Token_indexstr = "_Tokenindex"

var Regulator_Formindexstr = "_Regulatorindex"


func main() {
	if err := shim.Start(new(energyToken)); err != nil {
		fmt.Printf("Main: Error starting Energy_Token chaincode: %s", err)
	}
}

// ===================================================================================
// Init
// ===================================================================================
func (t *energyToken) Init(stub shim.ChaincodeStubInterface) pb.Response {
 if _, args := stub.GetFunctionAndParameters(); len(args) > 0 {
	return shim.Error("Init: Incorrect number of arguments; none expected.")
	}
	var enroll_empty []string
	enrollAsBytes, _ := json.Marshal(enroll_empty)								//marshal an emtpy array of strings to clear the index
	enroll_err := stub.PutState(Enroll_Formindexstr, enrollAsBytes)
	if enroll_err != nil {
		return shim.Error("Error while initialising")
	}

	var regGP_empty []string
	regGPAsBytes, _ := json.Marshal(regGP_empty)								//marshal an emtpy array of strings to clear the index
	regGP_err := stub.PutState(Register_G_P_Formindexstr, regGPAsBytes)
	if regGP_err != nil {
		return shim.Error("Error while initialising")
	}

	var regRC_empty []string
	regRCAsBytes, _ := json.Marshal(regRC_empty)								//marshal an emtpy array of strings to clear the index
	regRC_err := stub.PutState(Register_R_C_Formindexstr, regRCAsBytes)
	if regRC_err != nil {
		return shim.Error("Error while initialising")
	}

	var SignUp_empty []string
	signupAsBytes, _ := json.Marshal(SignUp_empty)								//marshal an emtpy array of strings to clear the index
	signup_err := stub.PutState(SignUp_Formindexstr, signupAsBytes)
	if signup_err != nil {
		return shim.Error("Error while initialising")
	}


    var Token_empty []string
	tokenAsBytes, _ := json.Marshal(Token_empty)								//marshal an emtpy array of strings to clear the index
	token_err := stub.PutState(Token_indexstr, tokenAsBytes)
	if token_err != nil {
		return shim.Error("Error while initialising")
	}


    var Regulator_empty []string
	regulatorAsBytes, _ := json.Marshal(Regulator_empty)								//marshal an emtpy array of strings to clear the index
	regulator_err := stub.PutState(Regulator_Formindexstr, regulatorAsBytes)
	if regulator_err != nil {
		return shim.Error("Error while initialising")
	}


	var Customer_empty []string
	customerAsBytes, _ := json.Marshal(Customer_empty)								//marshal an emtpy array of strings to clear the index
	customer_err := stub.PutState(CustomerFormindexstr, customerAsBytes)
	if customer_err != nil {
		return shim.Error("Error while initialising")
	}


	return shim.Success(nil)
}

// ===================================================================================
// Invoke
// ===================================================================================
func (t *energyToken) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
   fmt.Printf("Function Called = : %s", function)
	switch function {
		case "enrol": return t.enrol(stub, args)
		case "fetchData_By_EnrollID": return t.fetchData_By_EnrollID(stub, args)
		case "getAll_EnrollData": return t.getAll_EnrollData(stub, args)

		case "register_G_P": return t.register_G_P(stub, args);
		case "fetchData_By_Reg_G_P_ID": return t.fetchData_By_Reg_G_P_ID(stub, args);
		case "getAll_Reg_G_P_Data": return t.getAll_Reg_G_P_Data(stub, args);
		case "getData_by_Prosumers_remaining_token": return t.getData_by_Prosumers_remaining_token(stub, args);
        case "register_R_C": return t.register_R_C(stub, args);
        case "registerRegulator": return t.registerRegulator(stub, args);
		case "getAllRegulatorData": return t.getAllRegulatorData(stub, args);
		case "fetchData_By_Reg_R_C_ID": return t.fetchData_By_Reg_R_C_ID(stub, args);
		case "getAll_Reg_R_C_Data": return t.getAll_Reg_R_C_Data(stub, args);   
        case "registerCustomer": return t.registerCustomer(stub, args);
        case "getAllCustomerData": return t.getAllCustomerData(stub, args);
		case "sign_up": return t.sign_up(stub, args);
		case "fetchData_By_Sign_Up": return t.fetchData_By_Sign_Up(stub, args);
		case "getAll_signUP_Data": return t.getAll_signUP_Data(stub, args);

		case "generate_Token": return t.generate_Token(stub, args);
		case "getAll_TokenData": return t.getAll_TokenData(stub, args);
        case "fetchToken_By_TransactionID": return t.fetchToken_By_TransactionID(stub, args); 
        case "transferTokenToRetailer": return t.transferTokenToRetailer(stub, args);
        case "transferTokenToCustomer": return t.transferTokenToCustomer(stub, args);
        case "getHistoryDetails": return t.getHistoryDetails(stub, args);
        case "getTokensByAgreementID": return t.getTokensByAgreementID(stub, args);
        case "getTokensByProsumerID": return t.getTokensByProsumerID(stub, args); 
        case "getTokensByDateRange": return t.getTokensByDateRange(stub, args);
        case "getTokensByRetailerID": return t.getTokensByRetailerID(stub, args);
        case "updateProsumer": return t.updateProsumer(stub, args);
        case "updateRetailer": return t.updateRetailer(stub, args);
        case "updateAgreement": return t.updateAgreement(stub, args);
		default: return shim.Error("Invalid method!")
	}
}


//--------------------------------------------------------------------------------------------------------------------------------------
//------------------------------------------------------Enrollment starts---------------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------


// ===================================================================================
// enrol - Enrollment of Generators/Prosumers/Retailer/Consumer
// ===================================================================================
func (t *energyToken) enrol(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10.")
	}
	Name := args[0]
	Role := args[1]
	Email := args[2]
	Enrollment_timestamp := args[3]
	First_line_of_address := args[4]
	Second_line_of_address := args[5]
	City := args[6]
	Province := args[7]
	Postcode := args[8]
	Country := args[9]

	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000000000)

	_name := Name[0:3]
	_city := City[0:3]
	_role := Role[0:3]

	_randomNum := strconv.Itoa(randomNum)
	num := _randomNum[0:3]

    Participant_id := _name + _city + _role + num

    input := &Enrollment{Participant_id, Name, Role, Email, Enrollment_timestamp, First_line_of_address, Second_line_of_address, City, Province, Postcode, Country}

    inpuJSONasBytes, err := json.Marshal(input)
	if err != nil {
		return shim.Error(err.Error())
	}

	if messageAsBytes, err := stub.GetState(Participant_id); err != nil || messageAsBytes != nil {
		return shim.Error("This ID already exists.")
	}

	if err := stub.PutState(Participant_id, inpuJSONasBytes); err != nil {
		return shim.Error(err.Error())
		} else {
		fmt.Print("Block added")
	}


	EnrolIndexAsBytes, err := stub.GetState(Enroll_Formindexstr)
	if err != nil {
		return shim.Error("Failed to get Agreement index")
	}
	var Enrol_index []string
	fmt.Print("EnrolIndexAsBytes: ")
	fmt.Println(EnrolIndexAsBytes)

	json.Unmarshal(EnrolIndexAsBytes, &Enrol_index)
	fmt.Print("Enrol_index after unmarshal..before append: ")
	fmt.Println(Enrol_index)
	//append
	Enrol_index = append(Enrol_index, Participant_id)
	fmt.Println("! Enrollment index after appending Participant_id: ", Enrol_index)
	jsonAsBytes, _ := json.Marshal(Enrol_index)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(Enroll_Formindexstr, jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
	//return shim.Success([]byte(Participant_id))
    return shim.Success(inpuJSONasBytes)
}

// ===================================================================================
// Fetch_Data - fetchData_By_EnrollID
// ===================================================================================
func (t *energyToken) fetchData_By_EnrollID(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	Participant_id := args[0]
	if value, err := stub.GetState(Participant_id); err != nil || value == nil {
	return shim.Error("Read: invalid ID supplied.")
	} else {
		tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
		return shim.Success(value)
	}
}

// ===================================================================================
// getAllData - getAll_EnrollData
// ===================================================================================
func (t *energyToken) getAll_EnrollData(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("No arguements required")
	}

	var jsonResp, errResp string
	var err error
	var enrolRequestIndex []string

	fmt.Println("- start getAll_EnrollData")
	enrolRequestAsBytes, err := stub.GetState(Enroll_Formindexstr)
	if err != nil {
		return shim.Error("Failed to get Demand Request index")
	}
	fmt.Print("enrolRequestAsBytes : ")
	fmt.Println(enrolRequestAsBytes)
	json.Unmarshal(enrolRequestAsBytes, &enrolRequestIndex)								//un stringify it aka JSON.parse()
	fmt.Print("enrolRequestIndex : ")
	fmt.Println(enrolRequestIndex)
	fmt.Println("len(enrolRequestIndex) : ")
	fmt.Println(len(enrolRequestIndex))
	jsonResp = "{"
	for i,val := range enrolRequestIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Agreement request")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return shim.Error(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(enrolRequestIndex)-1 {
			jsonResp = jsonResp + ","
		}
	}
	jsonResp = jsonResp + "}"
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Println("end getAll_EnrollData")
	tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
	return shim.Success([]byte(jsonResp))
}

//**************************************************************************************************************************************
//--------------------------------------------------------------------------------------------------------------------------------------
//------------------------------------------------------Enrollment ends-----------------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------
//**************************************************************************************************************************************


//--------------------------------------------------------------------------------------------------------------------------------------
//------------------------------------------Registration for Generator and Prosumer starts----------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------

// ===================================================================================
// Registration of Generator and Prosumer
// ===================================================================================
func (t *energyToken) register_G_P(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 55 {
		return shim.Error("Incorrect number of arguments. Expecting 55.")
	}

	Participant_id := args[0]
	Type_of_plant := args[1]
	_Daily_injection_capacity := args[2]
	Installation_month := args[3]
	Registration_timestamp := args[4]
	Start_date := args[5]
	End_date := args[6]
	_S_0000_0030 := args[7]
	_S_0030_0100 := args[8]
	_S_0100_0130 := args[9]
	_S_0130_0200 := args[10]
	_S_0200_0230 := args[11]
	_S_0230_0300 := args[12]
	_S_0300_0330 := args[13]
	_S_0330_0400 := args[14]
	_S_0400_0430 := args[15]
	_S_0430_0500 := args[16]
	_S_0500_0530 := args[17]
	_S_0530_0600 := args[18]
	_S_0600_0630 := args[19]
	_S_0630_0700 := args[20]
	_S_0700_0730 := args[21]
	_S_0730_0800 := args[22]
	_S_0800_0830 := args[23]
	_S_0830_0900 := args[24]
	_S_0900_0930 := args[25]
	_S_0930_1000 := args[26]
	_S_1000_1030 := args[27]
	_S_1030_1100 := args[28]
	_S_1100_1130 := args[29]
	_S_1130_1200 := args[30]
	_S_1200_1230 := args[31]
	_S_1230_1300 := args[32]
	_S_1300_1330 := args[33]
	_S_1330_1400 := args[34]
	_S_1400_1430 := args[35]
	_S_1430_1500 := args[36]
	_S_1500_1530 := args[37]
	_S_1530_1600 := args[38]
	_S_1600_1630 := args[39]
	_S_1630_1700 := args[40]
	_S_1700_1730 := args[41]
	_S_1730_1800 := args[42]
	_S_1800_1830 := args[43]
	_S_1830_1900 := args[44]
	_S_1900_1930 := args[45]
	_S_1930_2000 := args[46]
	_S_2000_2030 := args[47]
	_S_2030_2100 := args[48]
	_S_2100_2130 := args[49]
	_S_2130_2200 := args[50]
	_S_2200_2230 := args[51]
	_S_2230_2300 := args[52]
	_S_2300_2330 := args[53]
	_S_2330_0000 := args[54]


	current_time := time.Now().Local()
	G_P_registration_id := Participant_id + current_time.Format("20060102150405")


	Daily_injection_capacity, _err := strconv.Atoi(_Daily_injection_capacity)
    if _err != nil {
		return shim.Error(_err.Error())
	}

	Remaining_token := Daily_injection_capacity

	S_0000_0030, _err := strconv.Atoi(_S_0000_0030)   //Slot 1
    if _err != nil {
		return shim.Error(_err.Error())
	}

    S_0030_0100, _err := strconv.Atoi(_S_0030_0100)   // Slot 2
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0100_0130, _err := strconv.Atoi(_S_0100_0130)   // Slot 3
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0130_0200, _err := strconv.Atoi(_S_0130_0200)   // Slot 4
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0200_0230, _err := strconv.Atoi(_S_0200_0230)   // Slot 5
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0230_0300, _err := strconv.Atoi(_S_0230_0300)   // Slot 6
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0300_0330, _err := strconv.Atoi(_S_0300_0330)   // Slot 7
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0330_0400, _err := strconv.Atoi(_S_0330_0400)   // Slot 8
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0400_0430, _err := strconv.Atoi(_S_0400_0430)   // Slot 9
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0430_0500, _err := strconv.Atoi(_S_0430_0500)   // Slot 10
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0500_0530, _err := strconv.Atoi(_S_0500_0530)   // Slot 11
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0530_0600, _err := strconv.Atoi(_S_0530_0600)   // Slot 12
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0600_0630, _err := strconv.Atoi(_S_0600_0630)   // Slot 13
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0630_0700, _err := strconv.Atoi(_S_0630_0700)   // Slot 14
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0700_0730, _err := strconv.Atoi(_S_0700_0730)   // Slot 15
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0730_0800, _err := strconv.Atoi(_S_0730_0800)   // Slot 16
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0800_0830, _err := strconv.Atoi(_S_0800_0830)   // Slot 17
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0830_0900, _err := strconv.Atoi(_S_0830_0900)   // Slot 18
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0900_0930, _err := strconv.Atoi(_S_0900_0930)   // Slot 19
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0930_1000, _err := strconv.Atoi(_S_0930_1000)   // Slot 20
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1000_1030, _err := strconv.Atoi(_S_1000_1030)   // Slot 21
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1030_1100, _err := strconv.Atoi(_S_1030_1100)   // Slot 22
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1100_1130, _err := strconv.Atoi(_S_1100_1130)   // Slot 23
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1130_1200, _err := strconv.Atoi(_S_1130_1200)   // Slot 24
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1200_1230, _err := strconv.Atoi(_S_1200_1230)   // Slot 25
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1230_1300, _err := strconv.Atoi(_S_1230_1300)   // Slot 26
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1300_1330, _err := strconv.Atoi(_S_1300_1330)   // Slot 27
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1330_1400, _err := strconv.Atoi(_S_1330_1400)   // Slot 28
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1400_1430, _err := strconv.Atoi(_S_1400_1430)   // Slot 29
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1430_1500, _err := strconv.Atoi(_S_1430_1500)   // Slot 30
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1500_1530, _err := strconv.Atoi(_S_1500_1530)   // Slot 31
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1530_1600, _err := strconv.Atoi(_S_1530_1600)   // Slot 32
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1600_1630, _err := strconv.Atoi(_S_1600_1630)   // Slot 33
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1630_1700, _err := strconv.Atoi(_S_1630_1700)   // Slot 34
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1700_1730, _err := strconv.Atoi(_S_1700_1730)   // Slot 35
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1730_1800, _err := strconv.Atoi(_S_1730_1800)   // Slot 36
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1800_1830, _err := strconv.Atoi(_S_1800_1830)   // Slot 37
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1830_1900, _err := strconv.Atoi(_S_1830_1900)   // Slot 38
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1900_1930, _err := strconv.Atoi(_S_1900_1930)   // Slot 39
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1930_2000, _err := strconv.Atoi(_S_1930_2000)   // Slot 40
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2000_2030, _err := strconv.Atoi(_S_2000_2030)   // Slot 41
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2030_2100, _err := strconv.Atoi(_S_2030_2100)   // Slot 42
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2100_2130, _err:= strconv.Atoi(_S_2100_2130)   // Slot 43
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2130_2200, _err:= strconv.Atoi(_S_2130_2200)   // Slot 44
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2200_2230, _err:= strconv.Atoi(_S_2200_2230)   // Slot 45
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2230_2300, _err:= strconv.Atoi(_S_2230_2300)   // Slot 46
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2300_2330, _err:= strconv.Atoi(_S_2300_2330)   // Slot 47
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2330_0000, _err:= strconv.Atoi(_S_2330_0000)   // Slot 48
    if _err != nil {
		return shim.Error(_err.Error())
	}


	input := &Register_G_P{G_P_registration_id, Participant_id, Type_of_plant, Daily_injection_capacity, Installation_month, Registration_timestamp, Remaining_token, Start_date, End_date, S_0000_0030, S_0030_0100, S_0100_0130, S_0130_0200, S_0200_0230, S_0230_0300, S_0300_0330, S_0330_0400, S_0400_0430,S_0430_0500, S_0500_0530, S_0530_0600, S_0600_0630, S_0630_0700, S_0700_0730, S_0730_0800, S_0800_0830, S_0830_0900, S_0900_0930, S_0930_1000, S_1000_1030, S_1030_1100, S_1100_1130, S_1130_1200, S_1200_1230, S_1230_1300, S_1300_1330, S_1330_1400, S_1400_1430, S_1430_1500, S_1500_1530, S_1530_1600, S_1600_1630,S_1630_1700, S_1700_1730, S_1730_1800, S_1800_1830, S_1830_1900, S_1900_1930, S_1930_2000, S_2000_2030, S_2030_2100, S_2100_2130, S_2130_2200, S_2200_2230, S_2230_2300, S_2300_2330, S_2330_0000,0}

	inpuJSONasBytes, err := json.Marshal(input)
	if err != nil {
		return shim.Error(err.Error())
	}


	if err := stub.PutState(G_P_registration_id, inpuJSONasBytes); err != nil {
		return shim.Error(err.Error())
		} else {
		fmt.Print("Block added")
	}


	RegisterGPIndexAsBytes, err := stub.GetState(Register_G_P_Formindexstr)
	if err != nil {
		return shim.Error("Failed to get Agreement index")
	}
	var RegisterGP_index []string
	fmt.Print("RegisterGPIndexAsBytes: ")
	fmt.Println(RegisterGPIndexAsBytes)

	json.Unmarshal(RegisterGPIndexAsBytes, &RegisterGP_index)
	fmt.Print("RegisterGP_index after unmarshal..before append: ")
	fmt.Println(RegisterGP_index)
	//append
	RegisterGP_index = append(RegisterGP_index, G_P_registration_id)
	fmt.Println("! Enrollment index after appending G_P_registration_id: ", RegisterGP_index)
	jsonAsBytes, _ := json.Marshal(RegisterGP_index)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(Register_G_P_Formindexstr, jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
	return shim.Success(inpuJSONasBytes)
}


// ===================================================================================
// Fetch_Data - fetch Data by registration ID of Generator or Prosumer
// ===================================================================================
func (t *energyToken) fetchData_By_Reg_G_P_ID(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	G_P_registration_id := args[0]
	if value, err := stub.GetState(G_P_registration_id); err != nil || value == nil {
	return shim.Error("Read: invalid ID supplied.")
	} else {
		tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
		return shim.Success(value)
	}
}


// ===================================================================================
// getAllData - Get all Registrations of Generators and Prosumers
// ===================================================================================
func (t *energyToken) getAll_Reg_G_P_Data(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("No arguements required")
	}

	var jsonResp, errResp string
	var err error
	var RegGPIndex []string

	fmt.Println("- start getAll_Reg_G_P_Data")
	Reg_G_P_AsBytes, err := stub.GetState(Register_G_P_Formindexstr)
	if err != nil {
		return shim.Error("Failed to get Demand Request index")
	}
	fmt.Print("Reg_G_P_AsBytes : ")
	fmt.Println(Reg_G_P_AsBytes)
	json.Unmarshal(Reg_G_P_AsBytes, &RegGPIndex)								//un stringify it aka JSON.parse()
	fmt.Print("RegGPIndex : ")
	fmt.Println(RegGPIndex)
	fmt.Println("len(RegGPIndex) : ")
	fmt.Println(len(RegGPIndex))
	jsonResp = "{"
	for i,val := range RegGPIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Reg request")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return shim.Error(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(RegGPIndex)-1 {
			jsonResp = jsonResp + ","
		}
	}
	jsonResp = jsonResp + "}"
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Println("end getAll_Reg_G_P_Data")
	tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
	return shim.Success([]byte(jsonResp))
}


// ===================================================================================
// getData_by_Prosumers_remaining_token - fetch data by Prosumers remaining token
// ===================================================================================
func (t *energyToken) getData_by_Prosumers_remaining_token(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var jsonResp, errResp string
  var err error
  fmt.Println("start getData_by_Prosumers_remaining_token")

  if len(args) != 0 {
    return shim.Error("No arguments required.")
  }

 	temp := Register_G_P{}


   ProsumerAsBytes, err := stub.GetState(Register_G_P_Formindexstr)                  //get the Participant_id from chaincode state
  if err != nil {
    jsonResp = "{\"Error\":\"Failed to get state "  + "\"}"
    return shim.Error(jsonResp)
  }

   var Prosumer_index []string
  json.Unmarshal(ProsumerAsBytes, &Prosumer_index)
  fmt.Print("Prosumer_index : ")
	fmt.Println(Prosumer_index)
  jsonResp = "{"
  for i,val := range Prosumer_index{

    fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for getData_by_remaining_token")
    valueAsBytes, err := stub.GetState(val)
    if err != nil {
      errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
      return shim.Error(errResp)
    }

    json.Unmarshal(valueAsBytes, &temp)


	fmt.Print("temp : ")
	fmt.Println(temp)

    if temp.Remaining_token > 0{
      fmt.Println("Prosumer found")
      jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
      if i < len(Prosumer_index)-1 {
        jsonResp = jsonResp + ","
      }
    }
  }
  //fmt.Print("valAsbytes : ")
  //fmt.Println(valAsbytes)
  jsonResp = jsonResp + "}"
  //fmt.Println("jsonResp : " + jsonResp)
  //fmt.Print("jsonResp in bytes : ")
  fmt.Println(jsonResp)
  fmt.Println("end getby getData_by_Prosumers_remaining_token")
  tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
  return shim.Success([]byte(jsonResp))
}

//**************************************************************************************************************************************
//--------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------Registration of Generator and Prosumer ends-----------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------
//**************************************************************************************************************************************


//--------------------------------------------------------------------------------------------------------------------------------------
//------------------------------------------Registration for Retailer and Consumer starts----------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------

// ===================================================================================
// Registration of Retailer and Consumer
// ===================================================================================
func (t *energyToken) register_R_C(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8.")
	}

	Participant_id := args[0]
	Type_of_plant := args[1]
	Installation_month := args[2]
	Registration_timestamp := args[3]
	_Required_energy_token := args[4]
	Start_date := args[5]
	End_date := args[6]
    Regulator := args[7]
	current_time := time.Now().Local()
	R_C_registration_id := Participant_id + current_time.Format("20060102150405")

	Required_energy_token, _err := strconv.Atoi(_Required_energy_token)
    if _err != nil {
		return shim.Error(_err.Error())
	}

	//Remaining_token := Required_energy_token 

	input := &Register_R_C{R_C_registration_id, Participant_id, Type_of_plant, Installation_month, Registration_timestamp, Required_energy_token, Required_energy_token, Start_date, End_date,0,Regulator}

	inputJSONasBytes, err := json.Marshal(input)
	if err != nil {
		return shim.Error(err.Error())
	}

	if err := stub.PutState(R_C_registration_id, inputJSONasBytes); err != nil {
		return shim.Error(err.Error())
		} else {
		fmt.Print("Block added")
	}


	RegisterRCIndexAsBytes, err := stub.GetState(Register_R_C_Formindexstr)
	if err != nil {
		return shim.Error("Failed to get Agreement index")
	}
	var RegisterRC_index []string
	fmt.Print("RegisterRCIndexAsBytes: ")
	fmt.Println(RegisterRCIndexAsBytes)

	json.Unmarshal(RegisterRCIndexAsBytes, &RegisterRC_index)
	fmt.Print("RegisterRC_index after unmarshal..before append: ")
	fmt.Println(RegisterRC_index)
	//append
	RegisterRC_index = append(RegisterRC_index, R_C_registration_id)
	fmt.Println("! Enrollment index after appending R_C_registration_id: ", RegisterRC_index)
	jsonAsBytes, _ := json.Marshal(RegisterRC_index)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(Register_R_C_Formindexstr, jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
	return shim.Success(inputJSONasBytes)
}





// ===================================================================================
// Registration of Regulator
// ===================================================================================
func (t *energyToken) registerRegulator(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7.")
	}

	Participant_id := args[0]
	Type_of_plant := args[1]
	Installation_month := args[2]
	Registration_timestamp := args[3]
	_Required_energy_token := args[4]
	Start_date := args[5]
	End_date := args[6]

	current_time := time.Now().Local()
	R_C_registration_id := Participant_id + current_time.Format("20060102150405")

	Required_energy_token, _err := strconv.Atoi(_Required_energy_token)
    if _err != nil {
		return shim.Error(_err.Error())
	}

	//Remaining_token := Required_energy_token 

	input := &RegisterRegulator{R_C_registration_id, Participant_id, Type_of_plant, Installation_month, Registration_timestamp, Required_energy_token, 0, Start_date, End_date,0}

	inputJSONasBytes, err := json.Marshal(input)
	if err != nil {
		return shim.Error(err.Error())
	}

	if err := stub.PutState(R_C_registration_id, inputJSONasBytes); err != nil {
		return shim.Error(err.Error())
		} else {
		fmt.Print("Block added")
	}


	RegulatorIndexAsBytes, err := stub.GetState(Regulator_Formindexstr)
	if err != nil {
		return shim.Error("Failed to get Agreement index")
	}
	var Regulator_index []string
	fmt.Print("RegulatorIndexAsBytes: ")
	fmt.Println(RegulatorIndexAsBytes)

	json.Unmarshal(RegulatorIndexAsBytes, &Regulator_index)
	fmt.Print("Regulator_index after unmarshal..before append: ")
	fmt.Println(Regulator_index)
	//append
	Regulator_index = append(Regulator_index, R_C_registration_id)
	fmt.Println("! Enrollment index after appending R_C_registration_id: ", Regulator_index)
	jsonAsBytes, _ := json.Marshal(Regulator_index)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(Regulator_Formindexstr, jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
	return shim.Success(inputJSONasBytes)
}


// ===================================================================================
// Fetch_Regulator Data - fetch Data by registration ID of Retailer or Consumer
// ===================================================================================
func (t *energyToken) getAllRegulatorData(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("No arguements required")
	}

	var jsonResp, errResp string
	var err error
	var RegRCIndex []string

	fmt.Println("- start getAllRegulatorData")
	Reg_R_C_AsBytes, err := stub.GetState(Regulator_Formindexstr)
	if err != nil {
		return shim.Error("Failed to get Demand Request index")
	}
	fmt.Print("Reg_R_C_AsBytes : ")
	fmt.Println(Reg_R_C_AsBytes)
	json.Unmarshal(Reg_R_C_AsBytes, &RegRCIndex)								//un stringify it aka JSON.parse()
	fmt.Print("RegRCIndex : ")
	fmt.Println(RegRCIndex)
	fmt.Println("len(RegRCIndex) : ")
	fmt.Println(len(RegRCIndex))
	jsonResp = "{"
	for i,val := range RegRCIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Reg request")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return shim.Error(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(RegRCIndex)-1 {
			jsonResp = jsonResp + ","
		}
	}
	jsonResp = jsonResp + "}"
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Println("end getAll_Reg_R_C_Data")
	tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
	return shim.Success([]byte(jsonResp))
}
// ===================================================================================
// Fetch_Data - fetch Data by registration ID of Retailer or Consumer
// ===================================================================================
func (t *energyToken) fetchData_By_Reg_R_C_ID(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	R_C_registration_id := args[0]
	if value, err := stub.GetState(R_C_registration_id); err != nil || value == nil {
	return shim.Error("Read: invalid ID supplied.")
	} else {
		tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
		return shim.Success(value)
	}
}


// ===================================================================================
// getAllData - Get all Registrations of Retailer 
// ===================================================================================
func (t *energyToken) getAll_Reg_R_C_Data(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("No arguements required")
	}

	var jsonResp, errResp string
	var err error
	var RegRCIndex []string

	fmt.Println("- start getAll_Reg_R_C_Data")
	Reg_R_C_AsBytes, err := stub.GetState(Register_R_C_Formindexstr)
	if err != nil {
		return shim.Error("Failed to get Demand Request index")
	}
	fmt.Print("Reg_R_C_AsBytes : ")
	fmt.Println(Reg_R_C_AsBytes)
	json.Unmarshal(Reg_R_C_AsBytes, &RegRCIndex)								//un stringify it aka JSON.parse()
	fmt.Print("RegRCIndex : ")
	fmt.Println(RegRCIndex)
	fmt.Println("len(RegRCIndex) : ")
	fmt.Println(len(RegRCIndex))
	jsonResp = "{"
	for i,val := range RegRCIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Reg request")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return shim.Error(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(RegRCIndex)-1 {
			jsonResp = jsonResp + ","
		}
	}
	jsonResp = jsonResp + "}"
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Println("end getAll_Reg_R_C_Data")
	tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
	return shim.Success([]byte(jsonResp))
}




func (t *energyToken) registerCustomer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
       fmt.Println("Before registerCustomer")
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}
     
     fmt.Println("INSIDE registerCustomer")

 
    CustomerData := args[0]


		current_time := time.Now().Local()

   //CustomerData := args[0]
  // RegisterCustomer 
   var customer RegisterCustomer


	json.Unmarshal([]byte(CustomerData), &customer)

	//Cregistration_id := ConsumerParticipantID + current_time.Format("20060102150405")

   

	   //Cregistration_id := customer.Participant_id + current_time.Format("20060102150405");
	   customer.Cregistration_id = customer.ConsumerParticipantID + current_time.Format("20060102150405")
        
	

	

	inputJSONasBytes, err := json.Marshal(customer)
	if err != nil {
		return shim.Error(err.Error())
	}

	if err := stub.PutState(customer.Cregistration_id, inputJSONasBytes); err != nil {
		return shim.Error(err.Error())
		} else {
		fmt.Print("Block added")
	}


	CustomerIndexAsBytes, err := stub.GetState(CustomerFormindexstr)
	if err != nil {
		return shim.Error("Failed to get Agreement index")
	}
	var Customer_index []string
	fmt.Print("CustomerIndexAsBytes: ")
	fmt.Println(CustomerIndexAsBytes)

	json.Unmarshal(CustomerIndexAsBytes, &Customer_index)
	fmt.Print("Customer_index after unmarshal..before append: ")
	fmt.Println(Customer_index)
	//append
	Customer_index = append(Customer_index, customer.Cregistration_id)
	fmt.Println("! Enrollment index after appending Cregistration_id: ", Customer_index)
	jsonAsBytes, _ := json.Marshal(Customer_index)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(CustomerFormindexstr, jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
	//return shim.Success(inputJSONasBytes)*/
	return shim.Success(inputJSONasBytes)
}








// ===================================================================================
// getAllCustomers - Get all Consumer
// ===================================================================================
func (t *energyToken) getAllCustomerData(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("No arguements required")
	}

	var jsonResp, errResp string
	var err error
	var CustomerIndex []string

	fmt.Println("- start getAllCustomerData")
	customerAsBytes, err := stub.GetState(CustomerFormindexstr)
	if err != nil {
		return shim.Error("Failed to get Demand Request index")
	}
	fmt.Print("customerAsBytes : ")
	fmt.Println(customerAsBytes)
	json.Unmarshal(customerAsBytes, &CustomerIndex)								//un stringify it aka JSON.parse()
	fmt.Print("CustomerIndex : ")
	fmt.Println(CustomerIndex)
	fmt.Println("len(CustomerIndex) : ")
	fmt.Println(len(CustomerIndex))
	jsonResp = "{"
	for i,val := range CustomerIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Reg request")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return shim.Error(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(CustomerIndex)-1 {
			jsonResp = jsonResp + ","
		}
	}
	jsonResp = jsonResp + "}"
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Println("end getAllCustomerData")
	tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
	return shim.Success([]byte(jsonResp))
}









//**************************************************************************************************************************************
//--------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------Registration of Retailer and Consumer ends-----------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------
//**************************************************************************************************************************************


//--------------------------------------------------------------------------------------------------------------------------------------
//------------------------------------------------------Sign up Agreement starts--------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------


// ===================================================================================
// sign_up - Sign_up agreement of Retailer and Prosumer
// ===================================================================================

func (t *energyToken) sign_up(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 59 {
		return shim.Error("Incorrect number of arguments. Expecting 59.")
	}

	Prosumer_participant_id := args[0]
	Prosumer_registration_id := args[1]
	Retailer_participant_id := args[2]
	Retailer_registration_id := args[3]
	_Signed_up_energy_token := args[4]
	_Time_period := args[5]
	Start_date := args[6]
	End_date := args[7]
	Token_remuneration_type := args[8]
	Value := args[9]
	_S_0000_0030 := args[10]
	_S_0030_0100 := args[11]
	_S_0100_0130 := args[12]
	_S_0130_0200 := args[13]
	_S_0200_0230 := args[14]
	_S_0230_0300 := args[15]
	_S_0300_0330 := args[16]
	_S_0330_0400 := args[17]
	_S_0400_0430 := args[18]
	_S_0430_0500 := args[19]
	_S_0500_0530 := args[20]
	_S_0530_0600 := args[21]
	_S_0600_0630 := args[22]
	_S_0630_0700 := args[23]
	_S_0700_0730 := args[24]
	_S_0730_0800 := args[25]
	_S_0800_0830 := args[26]
	_S_0830_0900 := args[27]
	_S_0900_0930 := args[28]
	_S_0930_1000 := args[29]
	_S_1000_1030 := args[30]
	_S_1030_1100 := args[31]
	_S_1100_1130 := args[32]
	_S_1130_1200 := args[33]
	_S_1200_1230 := args[34]
	_S_1230_1300 := args[35]
	_S_1300_1330 := args[36]
	_S_1330_1400 := args[37]
	_S_1400_1430 := args[38]
	_S_1430_1500 := args[39]
	_S_1500_1530 := args[40]
	_S_1530_1600 := args[41]
	_S_1600_1630 := args[42]
	_S_1630_1700 := args[43]
	_S_1700_1730 := args[44]
	_S_1730_1800 := args[45]
	_S_1800_1830 := args[46]
	_S_1830_1900 := args[47]
	_S_1900_1930 := args[48]
	_S_1930_2000 := args[49]
	_S_2000_2030 := args[50]
	_S_2030_2100 := args[51]
	_S_2100_2130 := args[52]
	_S_2130_2200 := args[53]
	_S_2200_2230 := args[54]
	_S_2230_2300 := args[55]
	_S_2300_2330 := args[56]
	_S_2330_0000 := args[57]
	Customer     := args[58]

	current_time := time.Now().Local()
	Agreement_id := "AG" + current_time.Format("20060102150405")

	Signed_up_energy_token,err := strconv.Atoi(_Signed_up_energy_token)
    if err != nil {
		return shim.Error(err.Error())
	}

	Time_period,err := strconv.Atoi(_Time_period)
    if err != nil {
		return shim.Error(err.Error())
	}

	Total_signed_up := Signed_up_energy_token * Time_period

	//Total_signed_up := strconv.Itoa(_Total_signed_up)
	S_0000_0030, _err := strconv.Atoi(_S_0000_0030)   //Slot 1
    if _err != nil {
		return shim.Error(_err.Error())
	}

    S_0030_0100, _err := strconv.Atoi(_S_0030_0100)   // Slot 2
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0100_0130, _err := strconv.Atoi(_S_0100_0130)   // Slot 3
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0130_0200, _err := strconv.Atoi(_S_0130_0200)   // Slot 4
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0200_0230, _err := strconv.Atoi(_S_0200_0230)   // Slot 5
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0230_0300, _err := strconv.Atoi(_S_0230_0300)   // Slot 6
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0300_0330, _err := strconv.Atoi(_S_0300_0330)   // Slot 7
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0330_0400, _err := strconv.Atoi(_S_0330_0400)   // Slot 8
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0400_0430, _err := strconv.Atoi(_S_0400_0430)   // Slot 9
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0430_0500, _err := strconv.Atoi(_S_0430_0500)   // Slot 10
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0500_0530, _err := strconv.Atoi(_S_0500_0530)   // Slot 11
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0530_0600, _err := strconv.Atoi(_S_0530_0600)   // Slot 12
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0600_0630, _err := strconv.Atoi(_S_0600_0630)   // Slot 13
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0630_0700, _err := strconv.Atoi(_S_0630_0700)   // Slot 14
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0700_0730, _err := strconv.Atoi(_S_0700_0730)   // Slot 15
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0730_0800, _err := strconv.Atoi(_S_0730_0800)   // Slot 16
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0800_0830, _err := strconv.Atoi(_S_0800_0830)   // Slot 17
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0830_0900, _err := strconv.Atoi(_S_0830_0900)   // Slot 18
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0900_0930, _err := strconv.Atoi(_S_0900_0930)   // Slot 19
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0930_1000, _err := strconv.Atoi(_S_0930_1000)   // Slot 20
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1000_1030, _err := strconv.Atoi(_S_1000_1030)   // Slot 21
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1030_1100, _err := strconv.Atoi(_S_1030_1100)   // Slot 22
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1100_1130, _err := strconv.Atoi(_S_1100_1130)   // Slot 23
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1130_1200, _err := strconv.Atoi(_S_1130_1200)   // Slot 24
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1200_1230, _err := strconv.Atoi(_S_1200_1230)   // Slot 25
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1230_1300, _err := strconv.Atoi(_S_1230_1300)   // Slot 26
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1300_1330, _err := strconv.Atoi(_S_1300_1330)   // Slot 27
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1330_1400, _err := strconv.Atoi(_S_1330_1400)   // Slot 28
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1400_1430, _err := strconv.Atoi(_S_1400_1430)   // Slot 29
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1430_1500, _err := strconv.Atoi(_S_1430_1500)   // Slot 30
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1500_1530, _err := strconv.Atoi(_S_1500_1530)   // Slot 31
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1530_1600, _err := strconv.Atoi(_S_1530_1600)   // Slot 32
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1600_1630, _err := strconv.Atoi(_S_1600_1630)   // Slot 33
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1630_1700, _err := strconv.Atoi(_S_1630_1700)   // Slot 34
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1700_1730, _err := strconv.Atoi(_S_1700_1730)   // Slot 35
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1730_1800, _err := strconv.Atoi(_S_1730_1800)   // Slot 36
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1800_1830, _err := strconv.Atoi(_S_1800_1830)   // Slot 37
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1830_1900, _err := strconv.Atoi(_S_1830_1900)   // Slot 38
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1900_1930, _err := strconv.Atoi(_S_1900_1930)   // Slot 39
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1930_2000, _err := strconv.Atoi(_S_1930_2000)   // Slot 40
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2000_2030, _err := strconv.Atoi(_S_2000_2030)   // Slot 41
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2030_2100, _err := strconv.Atoi(_S_2030_2100)   // Slot 42
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2100_2130, _err:= strconv.Atoi(_S_2100_2130)   // Slot 43
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2130_2200, _err:= strconv.Atoi(_S_2130_2200)   // Slot 44
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2200_2230, _err:= strconv.Atoi(_S_2200_2230)   // Slot 45
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2230_2300, _err:= strconv.Atoi(_S_2230_2300)   // Slot 46
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2300_2330, _err:= strconv.Atoi(_S_2300_2330)   // Slot 47
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2330_0000, _err:= strconv.Atoi(_S_2330_0000)   // Slot 48
    if _err != nil {
		return shim.Error(_err.Error())
	}

	input := &Sign_up{Agreement_id, Prosumer_participant_id, Prosumer_registration_id, Retailer_participant_id, Retailer_registration_id, Signed_up_energy_token, Time_period, Total_signed_up, Start_date, End_date, Token_remuneration_type, Value, S_0000_0030, S_0030_0100, S_0100_0130, S_0130_0200, S_0200_0230, S_0230_0300, S_0300_0330, S_0330_0400, S_0400_0430, S_0430_0500, S_0500_0530, S_0530_0600, S_0600_0630, S_0630_0700, S_0700_0730, S_0730_0800, S_0800_0830, S_0830_0900, S_0900_0930, S_0930_1000, S_1000_1030, S_1030_1100, S_1100_1130, S_1130_1200, S_1200_1230, S_1230_1300, S_1300_1330, S_1330_1400, S_1400_1430, S_1430_1500, S_1500_1530, S_1530_1600, S_1600_1630, S_1630_1700, S_1700_1730, S_1730_1800, S_1800_1830, S_1830_1900, S_1900_1930, S_1930_2000, S_2000_2030, S_2030_2100, S_2100_2130, S_2130_2200, S_2200_2230, S_2230_2300, S_2300_2330, S_2330_0000,Customer}

	inputJSONasBytes, err := json.Marshal(input)
	if err != nil {
		return shim.Error(err.Error())
	}


	if err := stub.PutState(Agreement_id, inputJSONasBytes); err != nil {
		return shim.Error(err.Error())
		} else {
		fmt.Print("Block added")
	}


	SignUpIndexAsBytes, err := stub.GetState(SignUp_Formindexstr)
	if err != nil {
		return shim.Error("Failed to get Agreement index")
	}
	var SignUp_index []string
	fmt.Print("SignUpIndexAsBytes: ")
	fmt.Println(SignUpIndexAsBytes)

	json.Unmarshal(SignUpIndexAsBytes, &SignUp_index)
	fmt.Print("SignUp_index after unmarshal..before append: ")
	fmt.Println(SignUp_index)
	//append
	SignUp_index = append(SignUp_index, Agreement_id)
	fmt.Println("! Enrollment index after appending Agreement_id: ", SignUp_index)
	jsonAsBytes, _ := json.Marshal(SignUp_index)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(SignUp_Formindexstr, jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////
	///////////////////////updating remaining tokens of Prosumers////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////

	valAsbytes, err := stub.GetState(Prosumer_registration_id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + Prosumer_registration_id + "\"}"
		return shim.Error(jsonResp)
	}

	res := Register_G_P{}
	json.Unmarshal(valAsbytes, &res)
    fmt.Printf("res: ")
    fmt.Println(res)

    fmt.Printf("valAsbytes: ")
    fmt.Println(valAsbytes)

    /*_Total_signed_up,_err := strconv.Atoi(Total_signed_up)
    if _err != nil {
		return shim.Error(_err.Error())
	}
 */
	fmt.Printf("Before res.Remaining_token: ")
    fmt.Println(res.Remaining_token)

    res.Remaining_token = res.Remaining_token - Signed_up_energy_token

    fmt.Printf("After res.Remaining_token: ")
    fmt.Println(res.Remaining_token)


    res.S_0000_0030 = res.S_0000_0030 - S_0000_0030

	res.S_0030_0100 = res.S_0030_0100 - S_0030_0100

	res.S_0100_0130 = res.S_0100_0130 - S_0100_0130

	res.S_0130_0200 = res.S_0130_0200 - S_0130_0200

	res.S_0200_0230 = res.S_0200_0230 - S_0200_0230

	res.S_0230_0300 = res.S_0230_0300 - S_0230_0300

	res.S_0300_0330 = res.S_0300_0330 - S_0300_0330

	res.S_0330_0400 = res.S_0330_0400 - S_0330_0400

	res.S_0400_0430 = res.S_0400_0430 - S_0400_0430

	res.S_0430_0500 = res.S_0430_0500 - S_0430_0500

	res.S_0500_0530 = res.S_0500_0530 - S_0500_0530

	res.S_0530_0600 = res.S_0530_0600 - S_0530_0600

	res.S_0600_0630 = res.S_0600_0630 - S_0600_0630

	res.S_0630_0700 = res.S_0630_0700 - S_0630_0700

	res.S_0700_0730 = res.S_0700_0730 - S_0700_0730

	res.S_0730_0800 = res.S_0730_0800 - S_0730_0800

	res.S_0800_0830 = res.S_0800_0830 - S_0800_0830

	res.S_0830_0900 = res.S_0830_0900 - S_0830_0900

	res.S_0900_0930 = res.S_0900_0930 - S_0900_0930

	res.S_0930_1000 = res.S_0930_1000 - S_0930_1000

	res.S_1000_1030 = res.S_1000_1030 - S_1000_1030

	res.S_1030_1100 = res.S_1030_1100 - S_1030_1100

	res.S_1100_1130 = res.S_1100_1130 - S_1100_1130

	res.S_1130_1200 = res.S_1130_1200 - S_1130_1200

	res.S_1200_1230 = res.S_1200_1230 - S_1200_1230

	res.S_1230_1300 = res.S_1230_1300 - S_1230_1300

	res.S_1300_1330 = res.S_1300_1330 - S_1300_1330

	res.S_1330_1400 = res.S_1330_1400 - S_1330_1400

	res.S_1400_1430 = res.S_1400_1430 - S_1400_1430

	res.S_1430_1500 = res.S_1430_1500 - S_1430_1500

	res.S_1500_1530 = res.S_1500_1530 - S_1500_1530

	res.S_1530_1600 = res.S_1530_1600 - S_1530_1600

	res.S_1600_1630 = res.S_1600_1630 - S_1600_1630

	res.S_1630_1700 = res.S_1630_1700 - S_1630_1700

	res.S_1700_1730 = res.S_1700_1730 - S_1700_1730

	res.S_1730_1800 = res.S_1730_1800 - S_1730_1800

	res.S_1800_1830 = res.S_1800_1830 - S_1800_1830

	res.S_1830_1900 = res.S_1830_1900 - S_1830_1900

	res.S_1900_1930 = res.S_1900_1930 - S_1900_1930

	res.S_1930_2000 = res.S_1930_2000 - S_1930_2000

	res.S_2000_2030 = res.S_2000_2030 - S_2000_2030

	res.S_2030_2100 = res.S_2030_2100 - S_2030_2100

	res.S_2100_2130 = res.S_2100_2130 - S_2100_2130

	res.S_2130_2200 = res.S_2130_2200 - S_2130_2200

	res.S_2200_2230 = res.S_2200_2230 - S_2200_2230

	res.S_2230_2300 = res.S_2230_2300 - S_2230_2300

	res.S_2300_2330 = res.S_2300_2330 - S_2300_2330

	res.S_2330_0000 = res.S_2330_0000 - S_2330_0000


 // ==== marshal to JSON ====
		registerRequest := &Register_G_P{res.G_P_registration_id, res.Participant_id, res.Type_of_plant, res.Daily_injection_capacity, res.Installation_month, res.Registration_timestamp, res.Remaining_token, res.Start_date, res.End_date, res.S_0000_0030, res.S_0030_0100, res.S_0100_0130, res.S_0130_0200, res.S_0200_0230, res.S_0230_0300, res.S_0300_0330, res.S_0330_0400, res.S_0400_0430, res.S_0430_0500, res.S_0500_0530, res.S_0530_0600, res.S_0600_0630, res.S_0630_0700, res.S_0700_0730, res.S_0730_0800, res.S_0800_0830, res.S_0830_0900, res.S_0900_0930, res.S_0930_1000, res.S_1000_1030, res.S_1030_1100, res.S_1100_1130, res.S_1130_1200, res.S_1200_1230, res.S_1230_1300, res.S_1300_1330, res.S_1330_1400, res.S_1400_1430, res.S_1430_1500, res.S_1500_1530, res.S_1530_1600, res.S_1600_1630, res.S_1630_1700, res.S_1700_1730, res.S_1730_1800, res.S_1800_1830, res.S_1830_1900, res.S_1900_1930, res.S_1930_2000, res.S_2000_2030, res.S_2030_2100, res.S_2100_2130, res.S_2130_2200, res.S_2200_2230, res.S_2230_2300, res.S_2300_2330, res.S_2330_0000,0}

		RegisterGPJSONasBytes, err := json.Marshal(registerRequest)
		if err != nil {
			return shim.Error(err.Error())
		}

	    err = stub.PutState(res.G_P_registration_id, RegisterGPJSONasBytes)                                 //store Demand request with id as key
	    if err != nil {
	        return shim.Error(err.Error())
	    }else{

	    	//Updating remaining token 
               	retailerValAsbytes, err := stub.GetState(Retailer_registration_id)
				if err != nil {
					jsonResp := "{\"Error\":\"Failed to get state for " + Retailer_registration_id + "\"}"
					return shim.Error(jsonResp)
				}else{
					resRetailer := Register_R_C{}
					json.Unmarshal(retailerValAsbytes, &resRetailer)
           	   
                    resRetailer.Remaining_token = resRetailer.Remaining_token - Signed_up_energy_token

					retailerInput := &Register_R_C{resRetailer.R_C_registration_id, resRetailer.Participant_id, resRetailer.Type_of_plant, resRetailer.Installation_month, resRetailer.Registration_timestamp, resRetailer.Required_energy_token, resRetailer.Remaining_token, resRetailer.Start_date, resRetailer.End_date,0,resRetailer.Regulator}

                           RegisterGCJSONasBytes, err := json.Marshal(retailerInput)
							if err != nil {
								return shim.Error(err.Error())
							}

						    err = stub.PutState(resRetailer.R_C_registration_id, RegisterGCJSONasBytes)                                 //store Demand request with id as key
						    if err != nil {
						        return shim.Error(err.Error())
						    }else{  

						    }  
				}

	           }
			


	//////////////////////////////////////////////////////////////////////////////////////////////////////
	///////////////////////Remaining token of Prosumer updated successfully//////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////
	
	return shim.Success(inputJSONasBytes)
}



// ===================================================================================
// fetchData_by_Sign_up - fetch Data by Sign up
// ===================================================================================
func (t *energyToken) fetchData_By_Sign_Up(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	Agreement_id := args[0]
	if value, err := stub.GetState(Agreement_id); err != nil || value == nil {
	return shim.Error("Read: invalid ID supplied.")
	} else {
		tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
		return shim.Success(value)
	}
}


// ===================================================================================
// getAllData - Get all Registrations of Retailer or Consumer
// ===================================================================================
func (t *energyToken) getAll_signUP_Data(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("No arguements required")
	}

	var jsonResp, errResp string
	var err error
	var signUPIndex []string

	fmt.Println("- start getAll_signUP_Data")
	SignUP_AsBytes, err := stub.GetState(SignUp_Formindexstr)
	if err != nil {
		return shim.Error("Failed to get Request index")
	}
	fmt.Print("SignUP_AsBytes : ")
	fmt.Println(SignUP_AsBytes)
	json.Unmarshal(SignUP_AsBytes, &signUPIndex)								//un stringify it aka JSON.parse()
	fmt.Print("signUPIndex : ")
	fmt.Println(signUPIndex)
	fmt.Println("len(signUPIndex) : ")
	fmt.Println(len(signUPIndex))
	jsonResp = "{"
	for i,val := range signUPIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Reg request")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return shim.Error(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(signUPIndex)-1 {
			jsonResp = jsonResp + ","
		}
	}
	jsonResp = jsonResp + "}"
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Println("end getAll_signUP_Data")
	tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
	return shim.Success([]byte(jsonResp))
}


// ===================================================================================
// Token Generation API
// ===================================================================================
func (t *energyToken) generate_Token(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8.")
	}

	//Participant_id := args[0]
	Prosumer_Registration_id := args[0]
	Time_slot := args[1]
    //Kw := args[2]
    Kw, errs_Kw := strconv.Atoi(args[2])   // Kw
    if errs_Kw != nil {
		return shim.Error("Read: invalid Kw supplied.")
	}
    Generate_Time_stamp , errs_Generate_Time_stamp := strconv.Atoi(args[3])
    if errs_Generate_Time_stamp != nil {
		return shim.Error("Read: invalid Time_stamp  supplied.")
	}
    Street := args[4]
    City := args[5]
    Postcode := args[6]
    Country := args[7]

   if value, err := stub.GetState(Prosumer_Registration_id); err != nil || value == nil {
	return shim.Error("Read: invalid ID supplied.")
	} else {
		res := Register_G_P{}
	    json.Unmarshal(value, &res)
	    fmt.Print("res(Prosumer_Registration_id1): ")
		fmt.Println(res)
	    Sender_Registration_id := Prosumer_Registration_id
	    Reciever_Registration_id := Prosumer_Registration_id
		Status := "Created"
		Agreement_id := ""
		No_of_Token := Kw
   		Token_remuneration_type  := ""
   	 	Token_remuneration_value := ""
   	 	Prosumer_Participant_id  := ""
        Retailer_Registration_id := ""
        Retailer_Participant_id  := ""
        Customer_Registration_id := ""
        Regulator :=""
   		current_time := time.Now().Local()
		Wallet_Time_stamp := Generate_Time_stamp
		Transaction_id := Prosumer_Registration_id + current_time.Format("20060102150405")
		input := &Token{Transaction_id, Sender_Registration_id, Reciever_Registration_id, Status, Prosumer_Registration_id,Prosumer_Participant_id,Retailer_Registration_id,Retailer_Participant_id,Customer_Registration_id,Street,City,Postcode,Country,Time_slot, Agreement_id, Kw, No_of_Token,Token_remuneration_type,Token_remuneration_value,Generate_Time_stamp,Wallet_Time_stamp,Regulator}

		tokeninputJSONasBytes, err := json.Marshal(input)
		if err != nil {
			return shim.Error("Failed to get Data for Prosumer Registeration ID")
		}
		if err := stub.PutState(Transaction_id, tokeninputJSONasBytes); err != nil {
			return shim.Error(err.Error())
		} else {
			fmt.Print("Block added")
		}
		TokenIndexAsBytes, err := stub.GetState(Token_indexstr)
		if err != nil {
			return shim.Error("Failed to get Agreement index")
		}
		var Token_index []string
		
		json.Unmarshal(TokenIndexAsBytes, &Token_index)
		fmt.Print("Token_index after unmarshal..before append: ")
		fmt.Println(Token_index)
		//append
		Token_index = append(Token_index, Transaction_id)
		fmt.Println("! Token index after appending Transaction_id: ", Transaction_id)
		jsonAsBytes, _ := json.Marshal(Token_index)
		fmt.Println("Token_index AFTER Prosumer_Registration_id1: ")
		fmt.Println(Token_index)
		err = stub.PutState(Token_indexstr, jsonAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
		tosend := "Event send"
        err = stub.SetEvent("evtsender", []byte(tosend))
        if err != nil {
            return shim.Error(err.Error())
        }
     // Token Generation Done
     // Updating Wallet of Prosumer
      Wallet := res.Wallet + Kw

       fmt.Println("BEFORE UPDATING WALLET of Prosumer  "+res.G_P_registration_id+" VALUE = "+ strconv.Itoa(res.Wallet))
      registerRequest := &Register_G_P{res.G_P_registration_id, res.Participant_id, res.Type_of_plant, res.Daily_injection_capacity, res.Installation_month, res.Registration_timestamp, res.Remaining_token, res.Start_date, res.End_date, res.S_0000_0030, res.S_0030_0100, res.S_0100_0130, res.S_0130_0200, res.S_0200_0230, res.S_0230_0300, res.S_0300_0330, res.S_0330_0400, res.S_0400_0430, res.S_0430_0500, res.S_0500_0530, res.S_0530_0600, res.S_0600_0630, res.S_0630_0700, res.S_0700_0730, res.S_0730_0800, res.S_0800_0830, res.S_0830_0900, res.S_0900_0930, res.S_0930_1000, res.S_1000_1030, res.S_1030_1100, res.S_1100_1130, res.S_1130_1200, res.S_1200_1230, res.S_1230_1300, res.S_1300_1330, res.S_1330_1400, res.S_1400_1430, res.S_1430_1500, res.S_1500_1530, res.S_1530_1600, res.S_1600_1630, res.S_1630_1700, res.S_1700_1730, res.S_1730_1800, res.S_1800_1830, res.S_1830_1900, res.S_1900_1930, res.S_1930_2000, res.S_2000_2030, res.S_2030_2100, res.S_2100_2130, res.S_2130_2200, res.S_2200_2230, res.S_2230_2300, res.S_2300_2330, res.S_2330_0000,Wallet}
		RegisterGPJSONasBytes, err := json.Marshal(registerRequest)
		if err != nil {
			return shim.Error(err.Error())
		}
	    err = stub.PutState(res.G_P_registration_id, RegisterGPJSONasBytes)                                   //store Demand request with id as key
	    if err != nil {
	        return shim.Error(err.Error())
	    }else{
	    	 fmt.Println("UPDATING WALLET of Prosumer  "+res.G_P_registration_id+" SUCCESSS1")


             if retailerValue, errRetailer := stub.GetState(Prosumer_Registration_id); errRetailer != nil || retailerValue == nil {
	    	 fmt.Println("ERROR IN UPDATION")
			//return "false"
			} else {
          		 res2 := Register_G_P{}
	   		     json.Unmarshal(retailerValue, &res2)
	    		fmt.Print("Prosumer_Registration_id AFTER UPDATION: ")
				fmt.Println(res2)
			}


	    }
    
     //  found := true;
       fmt.Println("BEFORE: getAgreementData")
       fmt.Println("BEFORE: Prosumer_Registration_id "+Prosumer_Registration_id+" Time_slot "+Time_slot+" Transaction_id "+Transaction_id)
       /*transferStatus := getAgreementData(stub,Prosumer_Registration_id,Time_slot,Transaction_id);
       fmt.Println("transferStatus  for token : "+transferStatus)*/
       





      /*if found{
      	return shim.Success(tokeninputJSONasBytes)
      }*/

       //End
		return shim.Success(tokeninputJSONasBytes)

	}
}

// ===================================================================================
// getAgreementData - get Agreement Data
// ===================================================================================


/*func (t *energyToken) getAgreementData(stub shim.ChaincodeStubInterface,prosumer_id string,time_slot string,Transaction_id string)  pb.Response {*/

  func (t *energyToken) transferTokenToRetailer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
      prosumer_id := args[0]
      time_slot   := args[1]
      Transaction_id := args[2]
	 //Search for all the agreement ID which has the prosumer Registeration
      fmt.Print("INSIDE getAGREEMENTData ")
       //var jsonResp, errResp string
       var jsonResp  string
	   var err1 error
	   var signUPIndex []string
       res   := Sign_up{}
       found := false 
	   // 
       resToken    := Token{}
       resProsumer := Register_G_P{}
       resRetailer := Register_R_C{}
	   fmt.Println("- start getAll_signUP_Data")
	   SignUP_AsBytes, err1 := stub.GetState(SignUp_Formindexstr)
	   if err1 != nil {
		return shim.Error("false")
	   }
		fmt.Print("SignUP_AsBytes : ")
	
		json.Unmarshal(SignUP_AsBytes, &signUPIndex)								//un stringify it aka JSON.parse()
		fmt.Println(signUPIndex)
		fmt.Println("len(signUPIndex) : ")
		fmt.Println(len(signUPIndex))
		fmt.Println("Prosumer_id : "+prosumer_id+"time_slot "+time_slot+"Transaction_id "+Transaction_id)
        // Get Token Data
     
		jsonResp = "{"
		for i,val := range signUPIndex{

			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Reg request")
			valueAsBytes, err := stub.GetState(val)
			if err != nil {
			//	errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return shim.Error("false")
			}
			json.Unmarshal(valueAsBytes, &res)
			fmt.Println("res(SignUp_Formindexstr1): ")
		fmt.Println(res)
			   fmt.Println("res.Prosumer_registration_id "+res.Prosumer_registration_id+" "+" prosumer_id = "+prosumer_id)
			if(res.Prosumer_registration_id == prosumer_id ){
				
				fmt.Println("FOUND the Agreement with Agreement_id = "+res.Agreement_id)
				if("S_0000_0030" == time_slot){
					if(res.S_0000_0030>0){
						found = true 
						break
                      
					}
				}else if("S_0030_0100" == time_slot){
					if(res.S_0030_0100>0){
                      found = true 
						break
					}
				}else if("S_0100_0130" == time_slot){
					if(res.S_0100_0130>0){
                     found = true 
						break
					}
				}else if("S_0130_0200" == time_slot){
					if(res.S_0130_0200>0){
                      found = true 
						break
					}
				}else if("S_0200_0230" == time_slot){
					if(res.S_0200_0230>0){
                    found = true 
						break
					}
				}else if("S_0230_0300" == time_slot){
					if(res.S_0230_0300>0){
                      found = true 
						break
					}
				}else if("S_0300_0330" == time_slot){
					if(res.S_0300_0330>0){
                      found = true 
						break
					}
				}else if("S_0330_0400" == time_slot){
					if(res.S_0330_0400>0){
                    found = true 
						break
					}
				}else if("S_0400_0430" == time_slot){
					if(res.S_0400_0430>0){
                      found = true 
						break
					}
				}else if("S_0430_0500" == time_slot){
					if(res.S_0430_0500>0){
                     found = true 
						break
					}
				}else if("S_0500_0530" == time_slot){
					if(res.S_0500_0530>0){
                  found = true 
						break
					}
				}else if("S_0530_0600" == time_slot){
					if(res.S_0530_0600>0){
                      found = true 
						break
					}
				}else if("S_0600_0630" == time_slot){
					if(res.S_0600_0630>0){
                      found = true 
						break
					}
				}else if("S_0630_0700" == time_slot){
					if(res.S_0630_0700>0){
                      found = true 
						break
					}
				}else if("S_0700_0730" == time_slot){
					if(res.S_0700_0730>0){
                      found = true 
						break
					}
				}else if("S_0730_0800" == time_slot){
					if(res.S_0730_0800>0){
                      found = true 
						break
					}
				}else if("S_0800_0830" == time_slot){
					if(res.S_0800_0830>0){
                      found = true 
						break
					}
				}else if("S_0830_0900" == time_slot){
					if(res.S_0830_0900>0){
                      found = true 
						break
					}
				}else if("S_0900_0930" == time_slot){
					if(res.S_0900_0930>0){
                      found = true 
						break
					}
				}else if("S_0930_1000" == time_slot){
					if(res.S_0930_1000>0){
                      found = true 
						break
					}
				}else if("S_1000_1030" == time_slot){
					if(res.S_1000_1030>0){
                      found = true 
						break
					}
				}else if("S_1030_1100" == time_slot){
					if(res.S_1030_1100>0){
                      found = true 
						break
					}
				}else if("S_1100_1130" == time_slot){
					if(res.S_1100_1130>0){
                      found = true 
						break
					}
				}else if("S_1130_1200" == time_slot){
					if(res.S_1130_1200>0){
                      found = true 
						break
					}
				}else if("S_1200_1230" == time_slot){
					if(res.S_1200_1230>0){
                      found = true 
						break
					}
				}else if("S_1230_1300" == time_slot){
					if(res.S_1230_1300>0){
                      found = true 
						break
					}
				}else if("S_1300_1330" == time_slot){
					if(res.S_1300_1330>0){
                      found = true 
						break
					}
				}else if("S_1330_1400" == time_slot){
					if(res.S_1330_1400>0){
                      found = true 
						break
					}
				}else if("S_1400_1430" == time_slot){
					if(res.S_1400_1430>0){
                      found = true 
						break
					}
				}else if("S_1430_1500" == time_slot){
					if(res.S_1430_1500>0){
                      found = true 
						break
					}
				}else if("S_1500_1530" == time_slot){
					if(res.S_1500_1530>0){
                      found = true 
						break
					}
				}else if("S_1530_1600" == time_slot){
					if(res.S_1530_1600>0){
                      found = true 
						break
					}
				}else if("S_1600_1630" == time_slot){
					if(res.S_1600_1630>0){
                      found = true 
						break
					}
				}else if("S_1630_1700" == time_slot){
					if(res.S_1630_1700>0){
                     found = true 
						break
					}
				}else if("S_1700_1730" == time_slot){
					if(res.S_1700_1730>0){
                      found = true 
						break
					}
				}else if("S_1730_1800" == time_slot){
					if(res.S_1730_1800>0){
                      found = true 
						break
					}
				}else if("S_1800_1830" == time_slot){
					if(res.S_1800_1830>0){
                      found = true 
						break
					}
				}else if("S_1830_1900" == time_slot){
					if(res.S_1830_1900>0){
                      found = true 
						break
					}
				}else if("S_1900_1930" == time_slot){
					if(res.S_1900_1930>0){
                      found = true 
						break
					}
				}else if("S_1930_2000" == time_slot){
					if(res.S_1930_2000>0){
                      found = true 
						break
					}
				}else if("S_2000_2030" == time_slot){
					if(res.S_2000_2030>0){
                      found = true 
						break
					}
				}else if("S_2030_2100" == time_slot){
					if(res.S_2030_2100>0){
                      found = true 
						break
					}
				}else if("S_2100_2130" == time_slot){
					if(res.S_2100_2130>0){
                      found = true 
						break
					}
				}else if("S_2130_2200" == time_slot){
					if(res.S_2130_2200>0){
                      found = true 
						break
					}
				}else if("S_2200_2230" == time_slot){
					if(res.S_2200_2230>0){
                      found = true 
						break
					}
				}else if("S_2230_2300" == time_slot){
					if(res.S_2230_2300>0){
                      found = true 
						break
					}
				}else if("S_2300_2330" == time_slot){
					if(res.S_2300_2330>0){
                      found = true 
						break
					}
				}else if("S_2330_0000" == time_slot){
					if(res.S_2330_0000>0){
                      found = true 
						break
					}
				}
				/*fmt.Print("valueAsBytes : ")
			    fmt.Println(valueAsBytes)*/
			    jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
			    if i < len(signUPIndex)-1 {
				 jsonResp = jsonResp + ","
				}
			
			}
			jsonResp = jsonResp + "}"
			fmt.Println("jsonResp : " + jsonResp)
		
	 //return shim.Success([]byte(jsonResp))
      
	  }

     if(found){
    	fmt.Println("AGREEMENT FOUND 1")  
          fmt.Println("Transaction_id")
          fmt.Println(Transaction_id)
    	   if valueToken, errToken := stub.GetState(Transaction_id); errToken != nil || valueToken == nil {
        	fmt.Println("ERROR 1 : ")
		   return shim.Error( "Read: invalid ID supplied.")
		   } else {
		   	fmt.Println("valueToken2")
       		fmt.Println(valueToken)
            fmt.Println("GOT TOKEN Info : ")			 
	        json.Unmarshal(valueToken, &resToken)
			fmt.Println("resToken2")
       	    fmt.Println(resToken)
    	     

               ///RETAILER UPDATION
				if prosumerValue, errProsumer := stub.GetState(resToken.Prosumer_Registration_id); errProsumer != nil || prosumerValue == nil {
					 fmt.Println("ERROR 2")
					return shim.Error( "false" )
				}else{
					fmt.Println("GOT PROSUMER DATA")

					json.Unmarshal(prosumerValue, &resProsumer)
					fmt.Println("resProsumer1")
					fmt.Println(resProsumer)
					if retailerValue, errRetailer := stub.GetState(res.Retailer_registration_id); errRetailer != nil || retailerValue == nil {
						fmt.Println("ERROR 3")
						return shim.Error( "false" )
						} else {
						fmt.Println("GOT RETAILER DATA")
						json.Unmarshal(retailerValue, &resRetailer)
						fmt.Println("resRetailer1")
						fmt.Println(resRetailer)
						fmt.Println("GOT PROSUMER WALLET BEFORE Updation "+ strconv.Itoa(resProsumer.Wallet) +" PROSUMER "+resProsumer.G_P_registration_id)
						resProsumer.Wallet = resProsumer.Wallet - resToken.Kw
						registerRequest := &Register_G_P{resProsumer.G_P_registration_id, resProsumer.Participant_id, resProsumer.Type_of_plant, resProsumer.Daily_injection_capacity, resProsumer.Installation_month, resProsumer.Registration_timestamp, resProsumer.Remaining_token, resProsumer.Start_date, resProsumer.End_date, resProsumer.S_0000_0030, resProsumer.S_0030_0100, resProsumer.S_0100_0130, resProsumer.S_0130_0200, resProsumer.S_0200_0230, resProsumer.S_0230_0300, resProsumer.S_0300_0330, resProsumer.S_0330_0400, resProsumer.S_0400_0430, resProsumer.S_0430_0500, resProsumer.S_0500_0530, resProsumer.S_0530_0600, resProsumer.S_0600_0630, resProsumer.S_0630_0700, resProsumer.S_0700_0730, resProsumer.S_0730_0800, resProsumer.S_0800_0830, resProsumer.S_0830_0900, resProsumer.S_0900_0930, resProsumer.S_0930_1000, resProsumer.S_1000_1030, resProsumer.S_1030_1100, resProsumer.S_1100_1130, resProsumer.S_1130_1200, resProsumer.S_1200_1230, resProsumer.S_1230_1300, resProsumer.S_1300_1330, resProsumer.S_1330_1400, resProsumer.S_1400_1430, resProsumer.S_1430_1500, resProsumer.S_1500_1530, resProsumer.S_1530_1600, resProsumer.S_1600_1630, resProsumer.S_1630_1700, resProsumer.S_1700_1730, resProsumer.S_1730_1800, resProsumer.S_1800_1830, resProsumer.S_1830_1900, resProsumer.S_1900_1930, resProsumer.S_1930_2000, resProsumer.S_2000_2030, resProsumer.S_2030_2100, resProsumer.S_2100_2130, resProsumer.S_2130_2200, resProsumer.S_2200_2230, resProsumer.S_2230_2300, resProsumer.S_2300_2330, resProsumer.S_2330_0000,resProsumer.Wallet}
						fmt.Println("registerRequest1")
						fmt.Println(registerRequest)
						RegisterGPJSONasBytes, err := json.Marshal(registerRequest)
						if err != nil {
							return shim.Error( "false" )
						}
						err = stub.PutState(resProsumer.G_P_registration_id, RegisterGPJSONasBytes)                                   //store Demand request with id as key
						if err != nil {
							return shim.Error( "false" )
						}else{
							fmt.Println(" PROSUMER WALLET UPDATED  ")
						}
						fmt.Println("GOT RETAILER WALLET BEFORE Updation "+ strconv.Itoa(resRetailer.Wallet) +" RETAILER "+resRetailer.R_C_registration_id)
						resRetailer.Wallet = resRetailer.Wallet + resToken.Kw
						retailerInput := &Register_R_C{resRetailer.R_C_registration_id, resRetailer.Participant_id, resRetailer.Type_of_plant, resRetailer.Installation_month, resRetailer.Registration_timestamp, resRetailer.Required_energy_token, resRetailer.Remaining_token, resRetailer.Start_date, resRetailer.End_date,resRetailer.Wallet,resRetailer.Regulator}
						fmt.Println("retailerInput1")
						fmt.Println(retailerInput)
						RetailerRJSONasBytes, err1 := json.Marshal(retailerInput)
						if err1 != nil {
							return shim.Error( "false" )
						}
						err1 = stub.PutState(resRetailer.R_C_registration_id, RetailerRJSONasBytes)                                   // store Demand request with id as key
						if err1 != nil {
								return shim.Error( "false" )
						}else{
							fmt.Println(" RETAILER WALLET UPDATED 1 ")

							input := &Token{resToken.Transaction_id,resToken.Sender_Registration_id, res.Retailer_registration_id, "Transfered To Retailer", resToken.Prosumer_Registration_id,res.Prosumer_participant_id,res.Retailer_registration_id,res.Retailer_participant_id,res.Customer,resToken.Street,resToken.City,resToken.Postcode,resToken.Country,resToken.Time_slot, res.Agreement_id, resToken.Kw, resToken.No_of_Token,res.Token_remuneration_type,res.Value,resToken.Generate_Time_stamp,resToken.Wallet_Time_stamp,resRetailer.Regulator}
							fmt.Println("input2")
							fmt.Println(input)
							tokeninputJSONasBytes, err := json.Marshal(input)
							if err != nil {
									return shim.Error("false")
							}
							if err := stub.PutState(Transaction_id, tokeninputJSONasBytes); err != nil {
									return shim.Error("false")
							} else {
							  fmt.Println(" Token Updated ")
							  return shim.Success(tokeninputJSONasBytes)
							}
						}
	           	
             
							return shim.Success([]byte("Token Updation Failed"))
					}

					}

						return shim.Success([]byte("Token Transfered"))
				}
    

					return shim.Success([]byte("Token Transfered"))
					//return shim.Success( "false")
			
			}
			    return shim.Success([]byte("Agreement Not Found"))
		}
 
 
 // ===================================================================================
 // Transfer - transfer token 
 // ===================================================================================
    /*
    func  transferTokenToCustomer(stub shim.ChaincodeStubInterface,prosumer_id string , retailer_id string,customer_id string,value int,token_id string)  string {*/
func (t *energyToken) transferTokenToCustomer(stub shim.ChaincodeStubInterface, args []string) pb.Response{
      Transaction_id := args[0]

	fmt.Println("INSIDE   TRANSFERTOKEN To Customer")	
	resRetailer  := Register_R_C{}
	resCustomer  := RegisterCustomer{}
	resToken     := Token{}	
	resRegulator := RegisterRegulator{}
    if tokenValue, tokenErr := stub.GetState(Transaction_id); tokenErr != nil || tokenValue == nil {
			return shim.Error( "false" )
	} else {
		json.Unmarshal(tokenValue, &resToken)			
		fmt.Println(tokenValue)
		fmt.Println("resToken:")
		fmt.Println(resToken)
		 if retailerValue, errRetailer := stub.GetState(resToken.Retailer_Registration_id); errRetailer != nil || retailerValue == nil {
	    		 fmt.Println("ERROR 3")
				 return shim.Error( "false" )
				} else {
					fmt.Println("GOT RETAILER DATA")
            	 	json.Unmarshal(retailerValue, &resRetailer)

                    if(resToken.Customer_Registration_id != "null"){          

						if customerValue, customerErr := stub.GetState(resToken.Customer_Registration_id); customerErr != nil || customerValue == nil {
						fmt.Println("ERROR 4")
						return shim.Error( "false" )
						} else {
							fmt.Println(" GOT CUSTOMER DATA   ")
							fmt.Println("customerValue1: ")
							fmt.Println(customerValue)
							json.Unmarshal(customerValue, &resCustomer)
                            fmt.Println("resCustomer1: ")
							fmt.Println(resCustomer)
							resCustomer.Wallet = resCustomer.Wallet + resToken.Kw
							/*customerInput := &RegisterCustomer{resCustomer.Cregistration_id,resCustomer.ConsumerParticipantID,resCustomer.PlantType, resCustomer.InstallationMonth, resCustomer.Shift1Total, resCustomer.Shift2Total, resCustomer.Shift3Total, resCustomer.Shift4Total, resCustomer.Shift1, resCustomer.Shift2,resCustomer.Shift3,resCustomer.Shift4,resCustomer.Prosumer_participant_id, resCustomer.Prosumer_registration_id, resCustomer.Retailer_participant_id, resCustomer.Retailer_registration_id, resCustomer.Signed_up_energy_token, resCustomer.Time_period, resCustomer.Total_signed_up, resCustomer.Start_date, resCustomer.End_date, resCustomer.Token_remuneration_type, resCustomer.Value, resCustomer.S_0000_0030Selected, resCustomer.S_0030_0100Selected, resCustomer.S_0100_0130Selected, resCustomer.S_0130_0200Selected, resCustomer.S_0200_0230Selected, resCustomer.S_0230_0300Selected, resCustomer.S_0300_0330Selected, resCustomer.S_0330_0400Selected, resCustomer.S_0400_0430Selected, resCustomer.S_0430_0500Selected, resCustomer.S_0500_0530Selected, resCustomer.S_0530_0600Selected, resCustomer.S_0600_0630Selected, resCustomer.S_0630_0700Selected, resCustomer.S_0700_0730Selected, resCustomer.S_0730_0800Selected, resCustomer.S_0800_0830Selected, resCustomer.S_0830_0900Selected, resCustomer.S_0900_0930Selected, resCustomer.S_0930_1000Selected, resCustomer.S_1000_1030Selected, resCustomer.S_1030_1100Selected, resCustomer.S_1100_1130Selected, resCustomer.S_1130_1200Selected, resCustomer.S_1200_1230Selected, resCustomer.S_1230_1300Selected, resCustomer.S_1300_1330Selected, resCustomer.S_1330_1400Selected, resCustomer.S_1400_1430Selected, resCustomer.S_1430_1500Selected, resCustomer.S_1500_1530Selected, resCustomer.S_1530_1600Selected, resCustomer.S_1600_1630Selected, resCustomer.S_1630_1700Selected, resCustomer.S_1700_1730Selected, resCustomer.S_1730_1800Selected, resCustomer.S_1800_1830Selected, resCustomer.S_1830_1900Selected, resCustomer.S_1900_1930Selected, resCustomer.S_1930_2000Selected, resCustomer.S_2000_2030Selected, resCustomer.S_2030_2100Selected, resCustomer.S_2100_2130Selected, resCustomer.S_2130_2200Selected, resCustomer.S_2200_2230Selected, resCustomer.S_2230_2300Selected, resCustomer.S_2300_2330, resCustomer.S_2330_0000,resCustomer.Wallet}*/
							fmt.Println("customerInput1: ")
							fmt.Println(resCustomer)
							//CustomerRJSONasBytes, err2 := json.Marshal(customerInput)
							CustomerRJSONasBytes, err2 := json.Marshal(resCustomer)
							if err2 != nil {
								return shim.Error( "false" )
							}
							err2 = stub.PutState(resCustomer.Cregistration_id, CustomerRJSONasBytes)                                   //store Demand request with id as key
							if err2 != nil {
								return shim.Error( "false" )
							}else{
								fmt.Println(" CUSTOMER WALLET UPDATED  ")
							}
				         //updating the wallet of retailer 
							resRetailer.Wallet = resRetailer.Wallet - resToken.Kw
							retailerInput := &Register_R_C{resRetailer.R_C_registration_id, resRetailer.Participant_id, resRetailer.Type_of_plant, resRetailer.Installation_month, resRetailer.Registration_timestamp, resRetailer.Required_energy_token, resRetailer.Remaining_token, resRetailer.Start_date, resRetailer.End_date,resRetailer.Wallet,resRetailer.Regulator}
							fmt.Println("retailerInput2: ")
							fmt.Println(retailerInput)
							RetailerRJSONasBytes, err3 := json.Marshal(retailerInput)
							if err3 != nil {
								return shim.Error( "false" )
							}
							err3 = stub.PutState(resRetailer.R_C_registration_id, RetailerRJSONasBytes)                                   //store Demand request with id as key
							if err3 != nil {
								return shim.Error( "false" )
							}else{
								fmt.Println(" RETAILER WALLET UPDATED  2")
							}
							fmt.Println("token_id:")
							fmt.Println(resToken.Transaction_id)
							//updating Token                      
							
				         
                            input := &Token{resToken.Transaction_id,resRetailer.R_C_registration_id, resCustomer.Cregistration_id, "Retired", resToken.Prosumer_Registration_id,resToken.Prosumer_Participant_id,resToken.Retailer_Registration_id,resToken.Retailer_Participant_id,resToken.Customer_Registration_id,resToken.Street,resToken.City,resToken.Postcode,resToken.Country, resToken.Time_slot, resToken.Agreement_id, resToken.Kw, resToken.No_of_Token,resToken.Token_remuneration_type,resToken.Token_remuneration_value,resToken.Generate_Time_stamp,resToken.Wallet_Time_stamp,resRetailer.Regulator}
                             fmt.Println("input:")
				            fmt.Println(input)
							   	tokeninputJSONasBytes, err := json.Marshal(input)
								if err != nil {
									return shim.Error( "false" )
								}
								if err := stub.PutState(resToken.Transaction_id, tokeninputJSONasBytes); err != nil {
								return shim.Error( "false" )
								} else {
								   return shim.Success(tokeninputJSONasBytes)
								}
								
						}
					}else{
					  fmt.Println(" NO CUSTOMER DATA   ")

                     fmt.Println(" Fetching Regulator DATA   ")
                     if regulatorValue, regulatorErr := stub.GetState(resToken.Regulator); regulatorErr != nil || regulatorValue == nil {
						fmt.Println("ERROR 4")
						return shim.Error( "false" )
						} else {
							fmt.Println(" GOT REGULATOR DATA   ")
							fmt.Println("regulatorValue: ")
							fmt.Println(regulatorValue)
							json.Unmarshal(regulatorValue, &resRegulator)
                            fmt.Println("resRegulator: ")
							fmt.Println(resRegulator)
							resRegulator.Wallet = resRegulator.Wallet + resToken.Kw
							regulatorInput := &RegisterRegulator{resRegulator.Rregistration_id, resRegulator.Participant_id, resRegulator.Type_of_plant, resRegulator.Installation_month, resRegulator.Registration_timestamp, resRegulator.Required_energy_token, resRegulator.Remaining_token, resRegulator.Start_date, resRegulator.End_date,resRegulator.Wallet}
							fmt.Println("regulatorInput: ")
							fmt.Println(regulatorInput)
							RegulatorRJSONasBytes, err2 := json.Marshal(regulatorInput)
							if err2 != nil {
								return shim.Error( "false" )
							}
							err2 = stub.PutState(resRegulator.Rregistration_id, RegulatorRJSONasBytes)                                   //store Demand request with id as key
							if err2 != nil {
								return shim.Error( "false" )
							}else{
								fmt.Println(" Regulator WALLET UPDATED  ")
							}
				         //updating the wallet of retailer 
							resRetailer.Wallet = resRetailer.Wallet - resToken.Kw
							retailerInput := &Register_R_C{resRetailer.R_C_registration_id, resRetailer.Participant_id, resRetailer.Type_of_plant, resRetailer.Installation_month, resRetailer.Registration_timestamp, resRetailer.Required_energy_token, resRetailer.Remaining_token, resRetailer.Start_date, resRetailer.End_date,resRetailer.Wallet,resRetailer.Regulator}
							fmt.Println("retailerInput2: ")
							fmt.Println(retailerInput)
							RetailerRJSONasBytes, err3 := json.Marshal(retailerInput)
							if err3 != nil {
								return shim.Error( "false" )
							}
							err3 = stub.PutState(resRetailer.R_C_registration_id, RetailerRJSONasBytes)                                   //store Demand request with id as key
							if err3 != nil {
								return shim.Error( "false" )
							}else{
								fmt.Println(" RETAILER WALLET UPDATED  3")
							}
							fmt.Println("token_id:")
							fmt.Println(resToken.Transaction_id)
							//updating Token 
					  
					    input := &Token{resToken.Transaction_id,resToken.Sender_Registration_id, resRegulator.Rregistration_id, "Retired", resToken.Prosumer_Registration_id,resToken.Prosumer_Participant_id,resToken.Retailer_Registration_id,resToken.Retailer_Participant_id,resToken.Customer_Registration_id,resToken.Street,resToken.City,resToken.Postcode,resToken.Country, resToken.Time_slot, resToken.Agreement_id, resToken.Kw, resToken.No_of_Token,resToken.Token_remuneration_type,resToken.Token_remuneration_value,resToken.Generate_Time_stamp,resToken.Wallet_Time_stamp,resRegulator.Rregistration_id}
                             fmt.Println("input:")
				            fmt.Println(input)
							   	tokeninputJSONasBytes, err := json.Marshal(input)
								if err != nil {
									return shim.Error( "false" )
								}
								if err := stub.PutState(resToken.Transaction_id, tokeninputJSONasBytes); err != nil {
								return shim.Error( "false" )
								} else {
								   return shim.Success(tokeninputJSONasBytes)
								}
					  
					  
					}
              }
			}
		   return shim.Success([]byte("Token Transfered"))
    }
       return shim.Success([]byte("Token Transfered"))
}



// ===================================================================================
// getAllToken - getAll_TokenData
// ===================================================================================
func (t *energyToken) getAll_TokenData(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("No arguements required")
	}

	var jsonResp, errResp string
	var err error
	var tokenRequestIndex []string

	fmt.Println("- start getAll_TokenData")
	tokenAsBytes, err := stub.GetState(Token_indexstr)
	if err != nil {
		return shim.Error("Failed to get Token Request index")
	}
	fmt.Print("tokenAsBytes : ")
	fmt.Println(tokenAsBytes)
	json.Unmarshal(tokenAsBytes, &tokenRequestIndex)								//un stringify it aka JSON.parse()
	fmt.Print("tokenRequestIndex : ")
	fmt.Println(tokenRequestIndex)
	fmt.Println("len(tokenRequestIndex) : ")
	fmt.Println(len(tokenRequestIndex))
	jsonResp = "{"
	for i,val := range tokenRequestIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Token request")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return shim.Error(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(tokenRequestIndex)-1 {
			jsonResp = jsonResp + ","
		}
	}
	jsonResp = jsonResp + "}"
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Println("end getAll_TokenData")
	tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
	return shim.Success([]byte(jsonResp))
}



// ===================================================================================
// fetchToken_By_TransactionID - fetchToken_By_TransactionID
// ===================================================================================
func (t *energyToken) fetchToken_By_TransactionID(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	Transaction_id := args[0]
	if value, err := stub.GetState(Transaction_id); err != nil || value == nil {
	return shim.Error("Read: invalid ID supplied.")
	} else {
		tosend := "Event send"
                        err = stub.SetEvent("evtsender", []byte(tosend))
                        if err != nil {
                            return shim.Error(err.Error())
                        }
		return shim.Success(value)
	}
}



// ===================================================================================
// GetHistroyDetails - Function to  Get Histroy Details of an ID
// ===================================================================================

func (t *energyToken) getHistoryDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    if len(args) < 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
    }

    trid := args[0]

    fmt.Printf("- start getTradeRequestHistory: %s\n", trid)

    resultsIterator, err := stub.GetHistoryForKey(trid)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

    // buffer is a JSON array containing historic values for the Trade
    var buffer bytes.Buffer
    buffer.WriteString("[")

    bArrayMemberAlreadyWritten := false
    for resultsIterator.HasNext() {
        response, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
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
        //as-is (as the Value itself a JSON Demand request)
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

    fmt.Printf("- getHistoryDetails returning:\n%s\n", buffer.String())

    return shim.Success(buffer.Bytes())
}

// ===================================================================================
// getTokensByAgreementID - Function to Get Tokens By AgreementID
// ===================================================================================

func (t *energyToken) getTokensByAgreementID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
 var jsonResp, errResp string
    var err error
      if len(args) 	!= 3 {
        return shim.Error("Incorrect number of arguments. Expecting 3")
    }


    aggrementID := args[0]
    D1 := args[1]
    D2 := args[2]
    Date1, _err := strconv.Atoi(D1)   //Slot 1
    if _err != nil {
		return shim.Error(_err.Error())
	}
	Date2, _err1 := strconv.Atoi(D2)   //Slot 1
    if _err1 != nil {
		return shim.Error(_err1.Error())
	}
    var tokenRequestIndex []string

    fmt.Println("- start getTokensByAgreementID")
    tokenRequestAsBytes, err := stub.GetState(Token_indexstr)
    if err != nil {
        return shim.Error("Failed to get Token Request index")
    }
    fmt.Print("tradeRequestAsBytes : ")
    fmt.Println(tokenRequestAsBytes)
    json.Unmarshal(tokenRequestAsBytes, &tokenRequestIndex)                           //un stringify it aka JSON.parse()
    fmt.Print("tokenRequestIndex : ")
    fmt.Println(tokenRequestIndex)
    fmt.Println("len(tokenRequestIndex) : ")
    fmt.Println(len(tokenRequestIndex))
    res := Token{}
   
    
    jsonResp = "["
    for i,val := range tokenRequestIndex{
         fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tokens ")
        valueAsBytes, err := stub.GetState(val)
         json.Unmarshal(valueAsBytes, &res)
         if res.Agreement_id == aggrementID && res.Generate_Time_stamp > Date1 &&  res.Generate_Time_stamp < Date2  {
         jsonResp = jsonResp + string(valueAsBytes[:])   
          if i < len(tokenRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }    
            
    }      

        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
       
    }
    jsonResponse := jsonResp
    if last := len(jsonResponse) - 1; last >= 0 && jsonResponse[last] == ',' {
        jsonResponse = jsonResponse[:last]
    }
    fmt.Println("jsonResponse:", jsonResponse)

    jsonResponse = jsonResponse + "]"
    fmt.Println("jsonResponse : " + jsonResponse)
    fmt.Println("end getTokensByAgreementID")
    return shim.Success([]byte(jsonResponse))
}


// ===================================================================================
// getTokensByProsumerID - Function to Get Tokens By ProsumerID
// ===================================================================================

func (t *energyToken) getTokensByProsumerID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
 var jsonResp, errResp string
    var err error
     if len(args) != 3 {
        return shim.Error("Incorrect number of arguments. Expecting 3")
    }

    prosumerID := args[0]
    D1      := args[1]
    D2      := args[2]
    Date1, _err := strconv.Atoi(D1)   //Slot 1
    if _err != nil {
		return shim.Error(_err.Error())
	}
	Date2, _err1 := strconv.Atoi(D2)   //Slot 1
    if _err1 != nil {
		return shim.Error(_err1.Error())
	}
    var tokenRequestIndex []string

    fmt.Println("- start getTokensByProsumerID")
    tokenRequestAsBytes, err := stub.GetState(Token_indexstr)
    if err != nil {
        return shim.Error("Failed to get Token Request index")
    }
    fmt.Print("tradeRequestAsBytes : ")
    fmt.Println(tokenRequestAsBytes)
    json.Unmarshal(tokenRequestAsBytes, &tokenRequestIndex)                           //un stringify it aka JSON.parse()
    fmt.Print("tokenRequestIndex : ")
    fmt.Println(tokenRequestIndex)
    fmt.Println("len(tokenRequestIndex) : ")
    fmt.Println(len(tokenRequestIndex))
    res := Token{}
   
    
    jsonResp = "["
    for i,val := range tokenRequestIndex{
         fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tokens ")
        valueAsBytes, err := stub.GetState(val)
         json.Unmarshal(valueAsBytes, &res)
         if res.Prosumer_Registration_id == prosumerID && res.Generate_Time_stamp > Date1 &&  res.Generate_Time_stamp < Date2  {
         jsonResp = jsonResp + string(valueAsBytes[:])   
          if i < len(tokenRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }    
            
    }      

        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
       
    }

    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getTokensByProsumerID")
    return shim.Success([]byte(jsonResp))
}


// ===================================================================================
// getTokensByRetailerID - Function to Get Tokens By RetailerID
// ===================================================================================

func (t *energyToken) getTokensByRetailerID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
 var jsonResp, errResp string
    var err error
     if len(args) != 3 {
        return shim.Error("Incorrect number of arguments. Expecting 3")
    }


    retailerID := args[0]
    D1 := args[1]
    D2 := args[2]

    Date1, _err := strconv.Atoi(D1)   //Slot 1
    if _err != nil {
		return shim.Error(_err.Error())
	}
	Date2, _err1 := strconv.Atoi(D2)   //Slot 1
    if _err1 != nil {
		return shim.Error(_err1.Error())
	}
    var tokenRequestIndex []string

    fmt.Println("- start getTokensByRetailerID")
    tokenRequestAsBytes, err := stub.GetState(Token_indexstr)
    if err != nil {
        return shim.Error("Failed to get Token Request index")
    }
    fmt.Print("tradeRequestAsBytes : ")
    fmt.Println(tokenRequestAsBytes)
    json.Unmarshal(tokenRequestAsBytes, &tokenRequestIndex)                           //un stringify it aka JSON.parse()
    fmt.Print("tokenRequestIndex : ")
    fmt.Println(tokenRequestIndex)
    fmt.Println("len(tokenRequestIndex) : ")
    fmt.Println(len(tokenRequestIndex))
    res := Token{}
   
    
    jsonResp = "["
    for i,val := range tokenRequestIndex{
         fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tokens ")
        valueAsBytes, err := stub.GetState(val)
         json.Unmarshal(valueAsBytes, &res)
         if res.Retailer_Registration_id == retailerID && res.Generate_Time_stamp > Date1 &&  res.Generate_Time_stamp < Date2  {
         jsonResp = jsonResp + string(valueAsBytes[:])   
          if i < len(tokenRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }    
            
    }      

        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
       
    }
     
      if last := len(jsonResp) - 1; last >= 0 && jsonResp[last] == ',' {
        jsonResp = jsonResp[:last]
    }
    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getTokensByRetailerID")
    return shim.Success([]byte(jsonResp))
}



// ===================================================================================
// getTokensByDateRange - Function to get Tokens By DateRange
// ===================================================================================

func (t *energyToken) getTokensByDateRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {
 var jsonResp, errResp string
    var err error
    if len(args) != 2 {
        return shim.Error("Incorrect number of arguments. Expecting 2")
    }


    D1 := args[0]
    D2 := args[1]
    Date1, _err := strconv.Atoi(D1)   //Slot 1
    if _err != nil {
		return shim.Error(_err.Error())
	}
	Date2, _err1 := strconv.Atoi(D2)   //Slot 1
    if _err1 != nil {
		return shim.Error(_err1.Error())
	}
    var tokenRequestIndex []string

    fmt.Println("- start getTokensByDateRange")
    tokenRequestAsBytes, err := stub.GetState(Token_indexstr)
    if err != nil {
        return shim.Error("Failed to get Token Request index")
    }
    fmt.Print("tradeRequestAsBytes : ")
    fmt.Println(tokenRequestAsBytes)
    json.Unmarshal(tokenRequestAsBytes, &tokenRequestIndex)                           //un stringify it aka JSON.parse()
    fmt.Print("tokenRequestIndex : ")
    fmt.Println(tokenRequestIndex)
    fmt.Println("len(tokenRequestIndex) : ")
    fmt.Println(len(tokenRequestIndex))
    res := Token{}
   
    
    jsonResp = "["
    for i,val := range tokenRequestIndex{
         fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tokens ")
        valueAsBytes, err := stub.GetState(val)
         json.Unmarshal(valueAsBytes, &res)
         if res.Generate_Time_stamp > Date1 &&  res.Generate_Time_stamp < Date2 {
         jsonResp = jsonResp + string(valueAsBytes[:])   
          if i < len(tokenRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }    
            
     }      

        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
       
    }

    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getTokensByDateRange")
    return shim.Success([]byte(jsonResp))
}


// ===================================================================================
// updateProsumer - Function to update Prosumer
// ===================================================================================


func (t *energyToken) updateProsumer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 55 {
		return shim.Error("Incorrect number of arguments. Expecting 55.")
	}



	G_P_registration_id := args[0]
	Type_of_plant := args[1]
	_Daily_injection_capacity := args[2]
	Installation_month := args[3]
	Registration_timestamp := args[4]
	Start_date := args[5]
	End_date := args[6]
	_S_0000_0030 := args[7]
	_S_0030_0100 := args[8]
	_S_0100_0130 := args[9]
	_S_0130_0200 := args[10]
	_S_0200_0230 := args[11]
	_S_0230_0300 := args[12]
	_S_0300_0330 := args[13]
	_S_0330_0400 := args[14]
	_S_0400_0430 := args[15]
	_S_0430_0500 := args[16]
	_S_0500_0530 := args[17]
	_S_0530_0600 := args[18]
	_S_0600_0630 := args[19]
	_S_0630_0700 := args[20]
	_S_0700_0730 := args[21]
	_S_0730_0800 := args[22]
	_S_0800_0830 := args[23]
	_S_0830_0900 := args[24]
	_S_0900_0930 := args[25]
	_S_0930_1000 := args[26]
	_S_1000_1030 := args[27]
	_S_1030_1100 := args[28]
	_S_1100_1130 := args[29]
	_S_1130_1200 := args[30]
	_S_1200_1230 := args[31]
	_S_1230_1300 := args[32]
	_S_1300_1330 := args[33]
	_S_1330_1400 := args[34]
	_S_1400_1430 := args[35]
	_S_1430_1500 := args[36]
	_S_1500_1530 := args[37]
	_S_1530_1600 := args[38]
	_S_1600_1630 := args[39]
	_S_1630_1700 := args[40]
	_S_1700_1730 := args[41]
	_S_1730_1800 := args[42]
	_S_1800_1830 := args[43]
	_S_1830_1900 := args[44]
	_S_1900_1930 := args[45]
	_S_1930_2000 := args[46]
	_S_2000_2030 := args[47]
	_S_2030_2100 := args[48]
	_S_2100_2130 := args[49]
	_S_2130_2200 := args[50]
	_S_2200_2230 := args[51]
	_S_2230_2300 := args[52]
	_S_2300_2330 := args[53]
	_S_2330_0000 := args[54]
 //Getting Details of Prosumer

	if value, err := stub.GetState(G_P_registration_id); err != nil || value == nil {
	return shim.Error("Read: invalid ID supplied.")
	} else {
		res := Register_G_P{}
		json.Unmarshal(value, &res)
	    fmt.Print("res(Prosumer_Registration_id1): ")
		fmt.Println(res)

	Daily_injection_capacity, _err := strconv.Atoi(_Daily_injection_capacity)
    if _err != nil {
		return shim.Error(_err.Error())
	}

	Remaining_token := Daily_injection_capacity

	S_0000_0030, _err := strconv.Atoi(_S_0000_0030)   //Slot 1
    if _err != nil {
		return shim.Error(_err.Error())
	}

    S_0030_0100, _err := strconv.Atoi(_S_0030_0100)   // Slot 2
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0100_0130, _err := strconv.Atoi(_S_0100_0130)   // Slot 3
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0130_0200, _err := strconv.Atoi(_S_0130_0200)   // Slot 4
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0200_0230, _err := strconv.Atoi(_S_0200_0230)   // Slot 5
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0230_0300, _err := strconv.Atoi(_S_0230_0300)   // Slot 6
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0300_0330, _err := strconv.Atoi(_S_0300_0330)   // Slot 7
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0330_0400, _err := strconv.Atoi(_S_0330_0400)   // Slot 8
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0400_0430, _err := strconv.Atoi(_S_0400_0430)   // Slot 9
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0430_0500, _err := strconv.Atoi(_S_0430_0500)   // Slot 10
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0500_0530, _err := strconv.Atoi(_S_0500_0530)   // Slot 11
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0530_0600, _err := strconv.Atoi(_S_0530_0600)   // Slot 12
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0600_0630, _err := strconv.Atoi(_S_0600_0630)   // Slot 13
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0630_0700, _err := strconv.Atoi(_S_0630_0700)   // Slot 14
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0700_0730, _err := strconv.Atoi(_S_0700_0730)   // Slot 15
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0730_0800, _err := strconv.Atoi(_S_0730_0800)   // Slot 16
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0800_0830, _err := strconv.Atoi(_S_0800_0830)   // Slot 17
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0830_0900, _err := strconv.Atoi(_S_0830_0900)   // Slot 18
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0900_0930, _err := strconv.Atoi(_S_0900_0930)   // Slot 19
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0930_1000, _err := strconv.Atoi(_S_0930_1000)   // Slot 20
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1000_1030, _err := strconv.Atoi(_S_1000_1030)   // Slot 21
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1030_1100, _err := strconv.Atoi(_S_1030_1100)   // Slot 22
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1100_1130, _err := strconv.Atoi(_S_1100_1130)   // Slot 23
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1130_1200, _err := strconv.Atoi(_S_1130_1200)   // Slot 24
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1200_1230, _err := strconv.Atoi(_S_1200_1230)   // Slot 25
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1230_1300, _err := strconv.Atoi(_S_1230_1300)   // Slot 26
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1300_1330, _err := strconv.Atoi(_S_1300_1330)   // Slot 27
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1330_1400, _err := strconv.Atoi(_S_1330_1400)   // Slot 28
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1400_1430, _err := strconv.Atoi(_S_1400_1430)   // Slot 29
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1430_1500, _err := strconv.Atoi(_S_1430_1500)   // Slot 30
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1500_1530, _err := strconv.Atoi(_S_1500_1530)   // Slot 31
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1530_1600, _err := strconv.Atoi(_S_1530_1600)   // Slot 32
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1600_1630, _err := strconv.Atoi(_S_1600_1630)   // Slot 33
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1630_1700, _err := strconv.Atoi(_S_1630_1700)   // Slot 34
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1700_1730, _err := strconv.Atoi(_S_1700_1730)   // Slot 35
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1730_1800, _err := strconv.Atoi(_S_1730_1800)   // Slot 36
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1800_1830, _err := strconv.Atoi(_S_1800_1830)   // Slot 37
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1830_1900, _err := strconv.Atoi(_S_1830_1900)   // Slot 38
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1900_1930, _err := strconv.Atoi(_S_1900_1930)   // Slot 39
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1930_2000, _err := strconv.Atoi(_S_1930_2000)   // Slot 40
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2000_2030, _err := strconv.Atoi(_S_2000_2030)   // Slot 41
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2030_2100, _err := strconv.Atoi(_S_2030_2100)   // Slot 42
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2100_2130, _err:= strconv.Atoi(_S_2100_2130)   // Slot 43
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2130_2200, _err:= strconv.Atoi(_S_2130_2200)   // Slot 44
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2200_2230, _err:= strconv.Atoi(_S_2200_2230)   // Slot 45
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2230_2300, _err:= strconv.Atoi(_S_2230_2300)   // Slot 46
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2300_2330, _err:= strconv.Atoi(_S_2300_2330)   // Slot 47
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2330_0000, _err:= strconv.Atoi(_S_2330_0000)   // Slot 48
    if _err != nil {
		return shim.Error(_err.Error())
	}


	input := &Register_G_P{res.G_P_registration_id, res.Participant_id, Type_of_plant, Daily_injection_capacity, Installation_month, Registration_timestamp, Remaining_token, Start_date, End_date, S_0000_0030, S_0030_0100, S_0100_0130, S_0130_0200, S_0200_0230, S_0230_0300, S_0300_0330, S_0330_0400, S_0400_0430,S_0430_0500, S_0500_0530, S_0530_0600, S_0600_0630, S_0630_0700, S_0700_0730, S_0730_0800, S_0800_0830, S_0830_0900, S_0900_0930, S_0930_1000, S_1000_1030, S_1030_1100, S_1100_1130, S_1130_1200, S_1200_1230, S_1230_1300, S_1300_1330, S_1330_1400, S_1400_1430, S_1430_1500, S_1500_1530, S_1530_1600, S_1600_1630,S_1630_1700, S_1700_1730, S_1730_1800, S_1800_1830, S_1830_1900, S_1900_1930, S_1930_2000, S_2000_2030, S_2030_2100, S_2100_2130, S_2130_2200, S_2200_2230, S_2230_2300, S_2300_2330, S_2330_0000,res.Wallet}

	inpuJSONasBytes, err := json.Marshal(input)
	if err != nil {
		return shim.Error(err.Error())
	}


	if err := stub.PutState(res.G_P_registration_id, inpuJSONasBytes); err != nil {
		return shim.Error(err.Error())
		} else {
		fmt.Print("Block added")
	}


	return shim.Success(inpuJSONasBytes)
 }
}



// ===================================================================================
// updateRetailer - Function to update Retailer
// ===================================================================================

func (t *energyToken) updateRetailer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 7.")
	}

	R_C_registration_id := args[0]
	Type_of_plant := args[1]
	Installation_month := args[2]
	Registration_timestamp := args[3]
	_Required_energy_token := args[4]
	Start_date := args[5]
	End_date := args[6]
    Regulator := args[7]
	/*current_time := time.Now().Local()
	R_C_registration_id := Participant_id + current_time.Format("20060102150405")
 */
	Required_energy_token, _err := strconv.Atoi(_Required_energy_token)
    if _err != nil {
		return shim.Error(_err.Error())
	}

	//Remaining_token := Required_energy_token 
	if value, err := stub.GetState(R_C_registration_id); err != nil || value == nil {
	return shim.Error("Read: invalid ID supplied.")
	} else {
		res := Register_R_C{}
		json.Unmarshal(value, &res)
	    fmt.Print("res(Prosumer_Registration_id1): ")
		fmt.Println(res)


	input := &Register_R_C{res.R_C_registration_id, res.Participant_id, Type_of_plant, Installation_month, Registration_timestamp, Required_energy_token, res.Remaining_token, Start_date, End_date,res.Wallet,Regulator}

	inputJSONasBytes, err := json.Marshal(input)
	if err != nil {
		return shim.Error(err.Error())
	}

	if err := stub.PutState(R_C_registration_id, inputJSONasBytes); err != nil {
		return shim.Error(err.Error())
		} else {
		fmt.Print("Block added")
	}
	return shim.Success(inputJSONasBytes)
 }
}


// ===================================================================================
// updateAgreement - Function to update Agreement
// ===================================================================================

func (t *energyToken) updateAgreement(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 60 {
		return shim.Error("Incorrect number of arguments. Expecting 60.")
	}

	Prosumer_participant_id := args[0]
	Prosumer_registration_id := args[1]
	Retailer_participant_id := args[2]
	Retailer_registration_id := args[3]
	_Signed_up_energy_token := args[4]
	_Time_period := args[5]
	Start_date := args[6]
	End_date := args[7]
	Token_remuneration_type := args[8]
	Value := args[9]
	_S_0000_0030 := args[10]
	_S_0030_0100 := args[11]
	_S_0100_0130 := args[12]
	_S_0130_0200 := args[13]
	_S_0200_0230 := args[14]
	_S_0230_0300 := args[15]
	_S_0300_0330 := args[16]
	_S_0330_0400 := args[17]
	_S_0400_0430 := args[18]
	_S_0430_0500 := args[19]
	_S_0500_0530 := args[20]
	_S_0530_0600 := args[21]
	_S_0600_0630 := args[22]
	_S_0630_0700 := args[23]
	_S_0700_0730 := args[24]
	_S_0730_0800 := args[25]
	_S_0800_0830 := args[26]
	_S_0830_0900 := args[27]
	_S_0900_0930 := args[28]
	_S_0930_1000 := args[29]
	_S_1000_1030 := args[30]
	_S_1030_1100 := args[31]
	_S_1100_1130 := args[32]
	_S_1130_1200 := args[33]
	_S_1200_1230 := args[34]
	_S_1230_1300 := args[35]
	_S_1300_1330 := args[36]
	_S_1330_1400 := args[37]
	_S_1400_1430 := args[38]
	_S_1430_1500 := args[39]
	_S_1500_1530 := args[40]
	_S_1530_1600 := args[41]
	_S_1600_1630 := args[42]
	_S_1630_1700 := args[43]
	_S_1700_1730 := args[44]
	_S_1730_1800 := args[45]
	_S_1800_1830 := args[46]
	_S_1830_1900 := args[47]
	_S_1900_1930 := args[48]
	_S_1930_2000 := args[49]
	_S_2000_2030 := args[50]
	_S_2030_2100 := args[51]
	_S_2100_2130 := args[52]
	_S_2130_2200 := args[53]
	_S_2200_2230 := args[54]
	_S_2230_2300 := args[55]
	_S_2300_2330 := args[56]
	_S_2330_0000 := args[57]
	Customer     := args[58]
	Agreement_id := args[59]
	

	Signed_up_energy_token,err := strconv.Atoi(_Signed_up_energy_token)
    if err != nil {
		return shim.Error(err.Error())
	}

	Time_period,err := strconv.Atoi(_Time_period)
    if err != nil {
		return shim.Error(err.Error())
	}

	Total_signed_up := Signed_up_energy_token * Time_period

	//Total_signed_up := strconv.Itoa(_Total_signed_up)
	S_0000_0030, _err := strconv.Atoi(_S_0000_0030)   //Slot 1
    if _err != nil {
		return shim.Error(_err.Error())
	}

    S_0030_0100, _err := strconv.Atoi(_S_0030_0100)   // Slot 2
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0100_0130, _err := strconv.Atoi(_S_0100_0130)   // Slot 3
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0130_0200, _err := strconv.Atoi(_S_0130_0200)   // Slot 4
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0200_0230, _err := strconv.Atoi(_S_0200_0230)   // Slot 5
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0230_0300, _err := strconv.Atoi(_S_0230_0300)   // Slot 6
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0300_0330, _err := strconv.Atoi(_S_0300_0330)   // Slot 7
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0330_0400, _err := strconv.Atoi(_S_0330_0400)   // Slot 8
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0400_0430, _err := strconv.Atoi(_S_0400_0430)   // Slot 9
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0430_0500, _err := strconv.Atoi(_S_0430_0500)   // Slot 10
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0500_0530, _err := strconv.Atoi(_S_0500_0530)   // Slot 11
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0530_0600, _err := strconv.Atoi(_S_0530_0600)   // Slot 12
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0600_0630, _err := strconv.Atoi(_S_0600_0630)   // Slot 13
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0630_0700, _err := strconv.Atoi(_S_0630_0700)   // Slot 14
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0700_0730, _err := strconv.Atoi(_S_0700_0730)   // Slot 15
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0730_0800, _err := strconv.Atoi(_S_0730_0800)   // Slot 16
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0800_0830, _err := strconv.Atoi(_S_0800_0830)   // Slot 17
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0830_0900, _err := strconv.Atoi(_S_0830_0900)   // Slot 18
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0900_0930, _err := strconv.Atoi(_S_0900_0930)   // Slot 19
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_0930_1000, _err := strconv.Atoi(_S_0930_1000)   // Slot 20
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1000_1030, _err := strconv.Atoi(_S_1000_1030)   // Slot 21
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1030_1100, _err := strconv.Atoi(_S_1030_1100)   // Slot 22
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1100_1130, _err := strconv.Atoi(_S_1100_1130)   // Slot 23
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1130_1200, _err := strconv.Atoi(_S_1130_1200)   // Slot 24
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1200_1230, _err := strconv.Atoi(_S_1200_1230)   // Slot 25
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1230_1300, _err := strconv.Atoi(_S_1230_1300)   // Slot 26
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1300_1330, _err := strconv.Atoi(_S_1300_1330)   // Slot 27
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1330_1400, _err := strconv.Atoi(_S_1330_1400)   // Slot 28
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1400_1430, _err := strconv.Atoi(_S_1400_1430)   // Slot 29
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1430_1500, _err := strconv.Atoi(_S_1430_1500)   // Slot 30
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1500_1530, _err := strconv.Atoi(_S_1500_1530)   // Slot 31
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1530_1600, _err := strconv.Atoi(_S_1530_1600)   // Slot 32
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1600_1630, _err := strconv.Atoi(_S_1600_1630)   // Slot 33
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1630_1700, _err := strconv.Atoi(_S_1630_1700)   // Slot 34
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1700_1730, _err := strconv.Atoi(_S_1700_1730)   // Slot 35
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1730_1800, _err := strconv.Atoi(_S_1730_1800)   // Slot 36
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1800_1830, _err := strconv.Atoi(_S_1800_1830)   // Slot 37
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1830_1900, _err := strconv.Atoi(_S_1830_1900)   // Slot 38
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1900_1930, _err := strconv.Atoi(_S_1900_1930)   // Slot 39
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_1930_2000, _err := strconv.Atoi(_S_1930_2000)   // Slot 40
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2000_2030, _err := strconv.Atoi(_S_2000_2030)   // Slot 41
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2030_2100, _err := strconv.Atoi(_S_2030_2100)   // Slot 42
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2100_2130, _err:= strconv.Atoi(_S_2100_2130)   // Slot 43
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2130_2200, _err:= strconv.Atoi(_S_2130_2200)   // Slot 44
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2200_2230, _err:= strconv.Atoi(_S_2200_2230)   // Slot 45
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2230_2300, _err:= strconv.Atoi(_S_2230_2300)   // Slot 46
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2300_2330, _err:= strconv.Atoi(_S_2300_2330)   // Slot 47
    if _err != nil {
		return shim.Error(_err.Error())
	}

	S_2330_0000, _err:= strconv.Atoi(_S_2330_0000)   // Slot 48
    if _err != nil {
		return shim.Error(_err.Error())
	}

  if value, err := stub.GetState(Agreement_id); err != nil || value == nil {
	return shim.Error("Read: invalid ID supplied.")
	} else {
		resAgreement := Sign_up{}
		json.Unmarshal(value, &resAgreement)
	    fmt.Print("resAgreement(Agreement_id): ")
		fmt.Println(resAgreement)





	input := &Sign_up{resAgreement.Agreement_id, Prosumer_participant_id, Prosumer_registration_id, Retailer_participant_id, Retailer_registration_id, Signed_up_energy_token, Time_period, Total_signed_up, Start_date, End_date, Token_remuneration_type, Value, S_0000_0030, S_0030_0100, S_0100_0130, S_0130_0200, S_0200_0230, S_0230_0300, S_0300_0330, S_0330_0400, S_0400_0430, S_0430_0500, S_0500_0530, S_0530_0600, S_0600_0630, S_0630_0700, S_0700_0730, S_0730_0800, S_0800_0830, S_0830_0900, S_0900_0930, S_0930_1000, S_1000_1030, S_1030_1100, S_1100_1130, S_1130_1200, S_1200_1230, S_1230_1300, S_1300_1330, S_1330_1400, S_1400_1430, S_1430_1500, S_1500_1530, S_1530_1600, S_1600_1630, S_1630_1700, S_1700_1730, S_1730_1800, S_1800_1830, S_1830_1900, S_1900_1930, S_1930_2000, S_2000_2030, S_2030_2100, S_2100_2130, S_2130_2200, S_2200_2230, S_2230_2300, S_2300_2330, S_2330_0000,Customer}

	inputJSONasBytes, err := json.Marshal(input)
	if err != nil {
		return shim.Error(err.Error())
	}


	if err := stub.PutState(Agreement_id, inputJSONasBytes); err != nil {
		return shim.Error(err.Error())
		} else {
		fmt.Print("Block added")
	}
	
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	///////////////////////updating remaining tokens of Prosumers////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////

	valAsbytes, err := stub.GetState(Prosumer_registration_id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + Prosumer_registration_id + "\"}"
		return shim.Error(jsonResp)
	}

	res := Register_G_P{}
	json.Unmarshal(valAsbytes, &res)
    fmt.Printf("res: ")
    fmt.Println(res)

    fmt.Printf("valAsbytes: ")
    fmt.Println(valAsbytes)

    /*_Total_signed_up,_err := strconv.Atoi(Total_signed_up)
    if _err != nil {
		return shim.Error(_err.Error())
	}
 */
	fmt.Printf("Before res.Remaining_token: ")
    fmt.Println(res.Remaining_token)

    res.Remaining_token = res.Remaining_token - Signed_up_energy_token

    fmt.Printf("After res.Remaining_token: ")
    fmt.Println(res.Remaining_token)


    res.S_0000_0030 = res.S_0000_0030 - S_0000_0030

	res.S_0030_0100 = res.S_0030_0100 - S_0030_0100

	res.S_0100_0130 = res.S_0100_0130 - S_0100_0130

	res.S_0130_0200 = res.S_0130_0200 - S_0130_0200

	res.S_0200_0230 = res.S_0200_0230 - S_0200_0230

	res.S_0230_0300 = res.S_0230_0300 - S_0230_0300

	res.S_0300_0330 = res.S_0300_0330 - S_0300_0330

	res.S_0330_0400 = res.S_0330_0400 - S_0330_0400

	res.S_0400_0430 = res.S_0400_0430 - S_0400_0430

	res.S_0430_0500 = res.S_0430_0500 - S_0430_0500

	res.S_0500_0530 = res.S_0500_0530 - S_0500_0530

	res.S_0530_0600 = res.S_0530_0600 - S_0530_0600

	res.S_0600_0630 = res.S_0600_0630 - S_0600_0630

	res.S_0630_0700 = res.S_0630_0700 - S_0630_0700

	res.S_0700_0730 = res.S_0700_0730 - S_0700_0730

	res.S_0730_0800 = res.S_0730_0800 - S_0730_0800

	res.S_0800_0830 = res.S_0800_0830 - S_0800_0830

	res.S_0830_0900 = res.S_0830_0900 - S_0830_0900

	res.S_0900_0930 = res.S_0900_0930 - S_0900_0930

	res.S_0930_1000 = res.S_0930_1000 - S_0930_1000

	res.S_1000_1030 = res.S_1000_1030 - S_1000_1030

	res.S_1030_1100 = res.S_1030_1100 - S_1030_1100

	res.S_1100_1130 = res.S_1100_1130 - S_1100_1130

	res.S_1130_1200 = res.S_1130_1200 - S_1130_1200

	res.S_1200_1230 = res.S_1200_1230 - S_1200_1230

	res.S_1230_1300 = res.S_1230_1300 - S_1230_1300

	res.S_1300_1330 = res.S_1300_1330 - S_1300_1330

	res.S_1330_1400 = res.S_1330_1400 - S_1330_1400

	res.S_1400_1430 = res.S_1400_1430 - S_1400_1430

	res.S_1430_1500 = res.S_1430_1500 - S_1430_1500

	res.S_1500_1530 = res.S_1500_1530 - S_1500_1530

	res.S_1530_1600 = res.S_1530_1600 - S_1530_1600

	res.S_1600_1630 = res.S_1600_1630 - S_1600_1630

	res.S_1630_1700 = res.S_1630_1700 - S_1630_1700

	res.S_1700_1730 = res.S_1700_1730 - S_1700_1730

	res.S_1730_1800 = res.S_1730_1800 - S_1730_1800

	res.S_1800_1830 = res.S_1800_1830 - S_1800_1830

	res.S_1830_1900 = res.S_1830_1900 - S_1830_1900

	res.S_1900_1930 = res.S_1900_1930 - S_1900_1930

	res.S_1930_2000 = res.S_1930_2000 - S_1930_2000

	res.S_2000_2030 = res.S_2000_2030 - S_2000_2030

	res.S_2030_2100 = res.S_2030_2100 - S_2030_2100

	res.S_2100_2130 = res.S_2100_2130 - S_2100_2130

	res.S_2130_2200 = res.S_2130_2200 - S_2130_2200

	res.S_2200_2230 = res.S_2200_2230 - S_2200_2230

	res.S_2230_2300 = res.S_2230_2300 - S_2230_2300

	res.S_2300_2330 = res.S_2300_2330 - S_2300_2330

	res.S_2330_0000 = res.S_2330_0000 - S_2330_0000


 // ==== marshal to JSON ====
		registerRequest := &Register_G_P{res.G_P_registration_id, res.Participant_id, res.Type_of_plant, res.Daily_injection_capacity, res.Installation_month, res.Registration_timestamp, res.Remaining_token, res.Start_date, res.End_date, res.S_0000_0030, res.S_0030_0100, res.S_0100_0130, res.S_0130_0200, res.S_0200_0230, res.S_0230_0300, res.S_0300_0330, res.S_0330_0400, res.S_0400_0430, res.S_0430_0500, res.S_0500_0530, res.S_0530_0600, res.S_0600_0630, res.S_0630_0700, res.S_0700_0730, res.S_0730_0800, res.S_0800_0830, res.S_0830_0900, res.S_0900_0930, res.S_0930_1000, res.S_1000_1030, res.S_1030_1100, res.S_1100_1130, res.S_1130_1200, res.S_1200_1230, res.S_1230_1300, res.S_1300_1330, res.S_1330_1400, res.S_1400_1430, res.S_1430_1500, res.S_1500_1530, res.S_1530_1600, res.S_1600_1630, res.S_1630_1700, res.S_1700_1730, res.S_1730_1800, res.S_1800_1830, res.S_1830_1900, res.S_1900_1930, res.S_1930_2000, res.S_2000_2030, res.S_2030_2100, res.S_2100_2130, res.S_2130_2200, res.S_2200_2230, res.S_2230_2300, res.S_2300_2330, res.S_2330_0000,res.Wallet}

		RegisterGPJSONasBytes, err := json.Marshal(registerRequest)
		if err != nil {
			return shim.Error(err.Error())
		}

	    err = stub.PutState(res.G_P_registration_id, RegisterGPJSONasBytes)                                 //store Demand request with id as key
	    if err != nil {
	        return shim.Error(err.Error())
	    }else{

	    	//Updating remaining token 
               	retailerValAsbytes, err := stub.GetState(Retailer_registration_id)
				if err != nil {
					jsonResp := "{\"Error\":\"Failed to get state for " + Retailer_registration_id + "\"}"
					return shim.Error(jsonResp)
				}else{
					resRetailer := Register_R_C{}
					json.Unmarshal(retailerValAsbytes, &resRetailer)           	   		
                    resRetailer.Remaining_token = resRetailer.Remaining_token - Signed_up_energy_token
					retailerInput := &Register_R_C{resRetailer.R_C_registration_id, resRetailer.Participant_id, resRetailer.Type_of_plant, resRetailer.Installation_month, resRetailer.Registration_timestamp, resRetailer.Required_energy_token, resRetailer.Remaining_token, resRetailer.Start_date, resRetailer.End_date,resRetailer.Wallet,resRetailer.Regulator}

                           RegisterGCJSONasBytes, err := json.Marshal(retailerInput)
							if err != nil {
								return shim.Error(err.Error())
							}

						    err = stub.PutState(resRetailer.R_C_registration_id, RegisterGCJSONasBytes)                                 //store Demand request with id as key
						    if err != nil {
						        return shim.Error(err.Error())
						    }else{  

						    }  
				}

	           }
			


	//////////////////////////////////////////////////////////////////////////////////////////////////////
	///////////////////////Remaining token of Prosumer updated successfully//////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////
	
	return shim.Success(inputJSONasBytes)
 }
}
