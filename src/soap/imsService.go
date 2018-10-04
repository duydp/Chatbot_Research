package soap

import (
	"fmt"
	"os"
	//"strconv"
)

func CheckFBIDExists(FbID string) (*SResult, error){
	var ParamXML = `<FbID>` + FbID + `</FbID>`
	return Post("CheckFBIDExists", ParamXML)
}

func GetInfoBySoHD(SoHD, FBID string) (*SResult, error){
	var ParamXML = `<SoHD>` + SoHD + `</SoHD>
		<FBID>` + FBID + `</FBID>`
	return Post("GetInfoFromSoHD", ParamXML)
}

func GetInfoBySoGQQL(SoGQQL, FBID string) (*SResult, error) {
	var ParamXML = `<SoGQQL>` + SoGQQL + `</SoGQQL> <FBID>` + FBID + `</FBID>`
	return Post("GetInfoBySoGQQL", ParamXML)
}
//author: duydp wrote on 12/3/2018
func GetInfoByBIEN_SOXE(Bien_SoXe string)(*SResult,error)  {
    fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<biensoxe>` + Bien_SoXe + `</biensoxe>
		<sokhung_may></sokhung_may>
		<chuhopdong></chuhopdong>
		<congty></congty>
		<phong></phong>
		<phone></phone>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("SearchCommonInfo",ParamXML)
}

//author: duydp wrote on 20/3/2018 CMND common
func GetInfoPolicyHolderID(CMTND string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<policyHolderID>`+CMTND+`</policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetPolicyInfo",ParamXML)
}


//author: duydp wrote on 21/3/2018 BKS
func GetregNumberCarInfo(regNumber string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<regNumber>`+regNumber+`</regNumber>
		<chassisNo></chassisNo>
		<certNumber></certNumber>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetCarInfo",ParamXML)
}
//author: duydp wrote on 21/3/2018 GCNBH
func GetcertNumberCarInfo(certNumber string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<regNumber></regNumber>
		<chassisNo></chassisNo>
		<certNumber>`+certNumber+`</certNumber>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetCarInfo",ParamXML)
}

//author: duydp wrote on 21/3/2018 CMND detail
func GetpolicyHolderIDCarInfo(policyHolderID string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<regNumber></regNumber>
		<chassisNo></chassisNo>
		<certNumber></certNumber>
		<policyHolderID>`+policyHolderID+`</policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetCarInfo",ParamXML)
}

//author: duydp wrote on 23/4/2018 SOKHUNG detail
func GetchassisNoCarInfo(chassisNo string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<regNumber></regNumber>
		<chassisNo>`+chassisNo+`</chassisNo>
		<certNumber></certNumber>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetCarInfo",ParamXML)
}

//begin***************Group : Fire**************************
//author: duydp wrote on 26/3/2018
// input: So don bao hiem (policyUrn)
// output: GetPolicyFireInfo
func GetPolicyURNFireInfo(policyUrn string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<regNumber></regNumber>
		<certNumber></certNumber>
		<policyUrn>`+policyUrn+`</policyUrn>
		<policyHolderName></policyHolderName>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<location></location>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetPolicyFireInfo",ParamXML)
}
//author: duydp wrote on 26/3/2018
// input: MST (policyHolderTaxCode)
// output: GetPolicyFireInfo
func GetPolicyHolderTaxCodeFireInfo(policyHolderTaxCode string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<regNumber></regNumber>
		<certNumber></certNumber>
		<policyUrn></policyUrn>
		<policyHolderName></policyHolderName>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode>`+policyHolderTaxCode+`</policyHolderTaxCode>
		<location></location>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetPolicyFireInfo",ParamXML)
}
//author: duydp wrote on 26/3/2018
// input: CMT (policyHolderID)
// output: GetPolicyFireInfo
func GetpolicyHolderIDFireInfo(policyHolderID string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<regNumber></regNumber>
		<certNumber></certNumber>
		<policyUrn></policyUrn>
		<policyHolderName></policyHolderName>
		<policyHolderID>`+policyHolderID+`</policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<location></location>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetPolicyFireInfo",ParamXML)
}
//author: duydp wrote on 12/4/2018
// input: vLocation (location)
// output: GetPolicyFireInfo
func GetLocationFireInfo(vLocation string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<regNumber></regNumber>
		<certNumber></certNumber>
		<policyUrn></policyUrn>
		<policyHolderName></policyHolderName>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<location>`+vLocation+`</location>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetPolicyFireInfo",ParamXML)
}
//author: duydp wrote on 13/4/2018
// input: vLocation (location)
// output: GetPolicyFireInfo
func GetPolicyHolderNameFireInfo(vPolicyHolderName string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<regNumber></regNumber>
		<certNumber></certNumber>
		<policyUrn></policyUrn>
		<policyHolderName>`+vPolicyHolderName+`</policyHolderName>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<location></location>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetPolicyFireInfo",ParamXML)
}
//End***************Group : Fire**************************

//begin***************Group : Vessel**********************

//author: duydp wrote on 28/3/2018
// input: so don tau (policyUrn)
// output: GetPolicyVesselInfo
func GetPolicyURNVesselInfo(policyUrn string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<policyUrn>`+policyUrn+`</policyUrn>
		<policyHolderName></policyHolderName>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<vesselName></vesselName>
		<regNumber></regNumber>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetPolicyVesselInfo",ParamXML)
}
//author: duydp wrote on 28/3/2018
// input: ten tau (vesselName)
// output: GetPolicyVesselInfo
func GetvesselNameInfo(vesselName string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<policyUrn></policyUrn>
		<policyHolderName></policyHolderName>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<vesselName>`+vesselName+`</vesselName>
		<regNumber></regNumber>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetPolicyVesselInfo",ParamXML)
}

//author: duydp wrote on 28/3/2018
// input: so dang ky (vesselName)
// output: GetPolicyVesselInfo
func GetregNumberVesselInfo(regNumber string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<policyUrn></policyUrn>
		<policyHolderName></policyHolderName>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<vesselName></vesselName>
		<regNumber>`+regNumber+`</regNumber>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetPolicyVesselInfo",ParamXML)
}
//End***************Group : Vessel**********************
//Begin***************Group : Cargo**********************
//author: duydp wrote on 30/3/2018
// input: so don hang (policyUrn)
// output: GetPolicyCargoInfo
func GetPolicyURNCargoInfo(policyUrn string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<policyUrn>`+policyUrn+`</policyUrn>
		<policyHolderName></policyHolderName>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<vesselName></vesselName>
		<regNumber></regNumber>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetPolicyCargoInfo",ParamXML)
}

//author: duydp wrote on 30/3/2018
// input: ten tau Hang (Cargo Name)
// output: GetPolicyCargoInfo
func GetCargoNameInfo(vesselName string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))
	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<policyUrn></policyUrn>
		<policyHolderName></policyHolderName>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<vesselName>`+vesselName+`</vesselName>
		<regNumber></regNumber>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetPolicyCargoInfo",ParamXML)
}
//author: duydp wrote on 28/8/2018
// input: ten khach hang f or HangHoa (Cargo's Customer Name)
// output: GetPolicyCargoInfoResult

func GetCargoCustomerNameInfo(CustomerName string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))
	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<policyUrn></policyUrn>
		<policyHolderName>`+CustomerName+`</policyHolderName>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<vesselName></vesselName>
		<regNumber></regNumber>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`
	return Post("GetPolicyCargoInfo",ParamXML)
}

//author: duydp wrote on 30/3/2018
// input: so dang ky hang (Cargo Name)
// output: GetPolicyCargoInfo
func GetRegNumberCargoInfo(regNumber string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<policyUrn></policyUrn>
		<policyHolderName></policyHolderName>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<vesselName></vesselName>
		<regNumber>`+regNumber+`</regNumber>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetPolicyCargoInfo",ParamXML)
}
//End***************Group : Cargo**********************

//Begin common information group : GetPolicyInfo
//author: duydp wrote on 03/04/2018
// input: thong tin chung theo CMT,MST (GetPolicyInfo)
// output: GetPolicyInfo
func GetCommonPolicyInfo(opt int,policyHolID string,policyHolTax string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))
	if opt==1{
		//get policyHolID
		var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<policyHolderID>`+policyHolID+`</policyHolderID>
		<policyHolderTaxCode></policyHolderTaxCode>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`
		return Post("GetPolicyInfo",ParamXML)
		}else{
		//get policyHolTax
		var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<policyHolderID></policyHolderID>
		<policyHolderTaxCode>`+policyHolTax+`</policyHolderTaxCode>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`
		return Post("GetPolicyInfo",ParamXML)
	}

}

//End

//Begin***************Group : BAGDV**********************
//author: duydp wrote on 8/5/2018
// input: so hop dong BH (ConstractNo)
// output: GetConstractNo_BAGDV
func Get_BAGDV_ConstractNo_Info(v_ConstractNo string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<contractNumber>`+v_ConstractNo+`</contractNumber>
		<insuredName></insuredName>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetContractInfo",ParamXML)
}
//author: duydp wrote on 8/5/2018
// input: ten khach hang(CustomerName)
// output: GetCusomterName_BAGDV
func Get_BAGDV_CustomerName_Info(v_insuredName string)(*SResult,error)  {
	fmt.Print(os.Getenv("Authen_key"))

	var ParamXML = `<authen_key>`+os.Getenv("Authen_key")+`</authen_key>
		<contractNumber></contractNumber>
		<insuredName>`+v_insuredName+`</insuredName>
		<pageNumber>`+os.Getenv("PageNumber")+` </pageNumber>
		<pageSize>`+os.Getenv("PageSize")+`</pageSize>`

	return Post("GetContractInfo",ParamXML)
}
//End***************Group : BAGDV**********************