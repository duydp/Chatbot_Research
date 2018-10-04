package main
import (
	"soap"
	"fmt"
	"encoding/xml"
	"BVGI/config"
	"strings"
	"strconv"

)

func mainbk()  {


	var policyURN_Vessel =strings.ToUpper("889210")

	result, _ :=  soap.GetPolicyURNCargoInfo(policyURN_Vessel)
	// Read xml
	byteValue := []byte(result.GetPolicyCargoInfoResult)

	var rowSets config.RowSet
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &rowSets)

	//if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication") {
	//	bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
	//	return ConfirmInputPolicyUrnCargoEvent
	//}

	var msgText  string
	var vesselNameVal = ""
	var vesselNameStr = ""
	var vesselNameStr1row = ""
	var vesselNameNo = 0

	var startCvr=0
	var vTradeName=""
	var vNgayNopPhi=""
	var vTinhTrangThuPhi=""
	var vPhoneTrade=""
	//var vSoVuBT=""
	//var vTyLeBT=""
	var vINCEPTION_DATE=""
	var vEXPIRY_DATE=""
	var vPOLICYURN=""
	var vPolicyHolderName=""
	var vBU=""
	var vGCN=""
	var vPhamVi=""
	var vCangDi =""
	var vCangDen =""


	for j:=0;j<len(rowSets.Rows); j++ {


		if (vesselNameVal != rowSets.Rows[j].VESSELORCONVEYANCE && rowSets.Rows[j].VESSELORCONVEYANCE != "") {
			vesselNameVal =  "Tàu " + rowSets.Rows[j].VESSELORCONVEYANCE
		}

		if (vPolicyHolderName != rowSets.Rows[j].POLICYHOLDER_NAME && rowSets.Rows[j].POLICYHOLDER_NAME != "") {
			vPolicyHolderName=  ", khách hàng " + rowSets.Rows[j].POLICYHOLDER_NAME
		}
		if (vBU != rowSets.Rows[j].BU_NAME && rowSets.Rows[j].BU_NAME != "") {
			vBU=  `, BH tại ` + rowSets.Rows[j].BU_NAME
		}


		if (vINCEPTION_DATE != rowSets.Rows[j].INCEPTION_DATE && rowSets.Rows[j].INCEPTION_DATE != "") {
			vINCEPTION_DATE=  `, từ ` + rowSets.Rows[j].INCEPTION_DATE
		}
		if (vEXPIRY_DATE != rowSets.Rows[j].INCEPTION_DATE && rowSets.Rows[j].EXPIRY_DATE != "") {
			vEXPIRY_DATE=  ` đến ` + rowSets.Rows[j].EXPIRY_DATE
		}

		if (vPOLICYURN != rowSets.Rows[j].POLICY_URN && rowSets.Rows[j].POLICY_URN != "") {
			vPOLICYURN=  ". Số đơn " + rowSets.Rows[j].POLICY_URN
		}

		if (vGCN != rowSets.Rows[j].CERTIFICATENUMBER && rowSets.Rows[j].CERTIFICATENUMBER !=""){
			vGCN = `, GCN ` + rowSets.Rows[j].CERTIFICATENUMBER
		}

		if (vPhamVi!= rowSets.Rows[j].SUMINSURED_AMT && rowSets.Rows[j].SUMINSURED_AMT!= "" && !strings.HasPrefix(strings.Trim(rowSets.Rows[j].SUMINSURED_AMT,""),"0 ")) {
			vPhamVi=`, tham gia BH ` + rowSets.Rows[j].COV_CLASS_NAME + " (" + rowSets.Rows[j].SUMINSURED_AMT + ")"
		}


		if (vCangDi != rowSets.Rows[j].PLACEOFDEPARTURES && rowSets.Rows[j].PLACEOFDEPARTURES != "") {
			vCangDi=  `, đi từ ` + rowSets.Rows[j].PLACEOFDEPARTURES
		}

		if (vCangDen != rowSets.Rows[j].FINALDESTINATIONS && rowSets.Rows[j].FINALDESTINATIONS != "") {
			vCangDen=  ` đến ` + rowSets.Rows[j].FINALDESTINATIONS
		}

		if (vNgayNopPhi!= rowSets.Rows[j].NGAY_NOP_PHI && rowSets.Rows[j].NGAY_NOP_PHI!= "") {
			vNgayNopPhi=`, nộp phí ngày`+rowSets.Rows[j].NGAY_NOP_PHI
		}

		if (vTinhTrangThuPhi!=rowSets.Rows[j].TINH_TRANG_THU_PHI && rowSets.Rows[j].TINH_TRANG_THU_PHI != "") {
			if rowSets.Rows[j].TINH_TRANG_THU_PHI != "Đã nộp phí" {
				vTinhTrangThuPhi = ` ` + strings.ToLower(rowSets.Rows[j].TINH_TRANG_THU_PHI)
			}
		}

		if (vTradeName!= rowSets.Rows[j].TRADE_NAME && rowSets.Rows[j].TRADE_NAME!= "") {
			vTradeName=`. LH CBKT ` +rowSets.Rows[j].TRADE_NAME
		}

		if (vPhoneTrade!= strings.Replace(rowSets.Rows[j].TRADE_PHONE, "/", "", 1) && strings.Replace(rowSets.Rows[j].TRADE_PHONE, "/", "", 1)!= "") {
			vPhoneTrade="("+strings.Replace(rowSets.Rows[j].TRADE_PHONE, "/", "", 1)+")"
		}

		//if (vSoVuBT!= rowSets.Rows[j].SO_VU_BT && rowSets.Rows[j].SO_VU_BT!= ""&& rowSets.Rows[j].SO_VU_BT!= "0") {
		//	vSoVuBT=`. Số vụ tổn thất: ` + rowSets.Rows[j].SO_VU_BT + " vụ,"
		//}
		//
		//if (vTyLeBT!= rowSets.Rows[j].TY_LE_BT && rowSets.Rows[j].TY_LE_BT!= "" && rowSets.Rows[j].TY_LE_BT!= "0%") {
		//	vTyLeBT=` tỷ lệ BT: ` + rowSets.Rows[j].TY_LE_BT
		//}

		if (vesselNameVal != rowSets.Rows[j].VESSELORCONVEYANCE && rowSets.Rows[j].VESSELORCONVEYANCE != "") {
			if vesselNameVal!="" {
				msgText += vesselNameStr
			}
			vesselNameNo +=1
			startCvr = 1
			vesselNameStr = "\r\n" + strconv.Itoa(vesselNameNo)+". " + `Tàu ` + rowSets.Rows[j].VESSELORCONVEYANCE

			if vesselNameNo==1 {
				vesselNameStr1row =`Tàu ` + rowSets.Rows[j].VESSELORCONVEYANCE
				//vesselNameStr1row += `, GCN ` + rowSets.Rows[j].CERTIFICATENUMBER
				//vesselNameStr1row += `, tham gia BH ` + rowSets.Rows[j].COV_CLASS_NAME + " (" + rowSets.Rows[j].SUMINSURED_AMT + ")"
				//vesselNameStr1row += `, đi từ ` + rowSets.Rows[j].PLACEOFDEPARTURES + ` đến ` + rowSets.Rows[j].FINALDESTINATIONS
				//vesselNameStr1row += `, nộp phí ngày`+rowSets.Rows[j].NGAY_NOP_PHI
				//vesselNameStr1row += ` ` + strings.ToLower(rowSets.Rows[j].TINH_TRANG_THU_PHI)
			}

			//vesselNameVal= `Tàu ` +rowSets.Rows[j].VESSELORCONVEYANCE

		}

		if startCvr==0 {
			vesselNameStr +=", "
			vesselNameStr1row += ", "
		}

		vesselNameStr += `, GCN ` + rowSets.Rows[j].CERTIFICATENUMBER

		//vesselNameStr += `, tham gia BH ` + rowSets.Rows[j].COV_CLASS_NAME + " (" + rowSets.Rows[j].SUMINSURED_AMT + ")"
		if (rowSets.Rows[j].SUMINSURED_AMT!= "" && !strings.HasPrefix(strings.Trim(rowSets.Rows[j].SUMINSURED_AMT,""),"0 ")) {
			vesselNameStr += ", tham gia BH " + rowSets.Rows[j].COV_CLASS_NAME + " (" + rowSets.Rows[j].SUMINSURED_AMT + ")"

			vesselNameStr1row	+= rowSets.Rows[j].COV_CLASS_NAME + " (" + rowSets.Rows[j].SUMINSURED_AMT + ")"
			startCvr=0
		}

		vesselNameStr += `, đi từ ` + rowSets.Rows[j].PLACEOFDEPARTURES + ` đến ` + rowSets.Rows[j].FINALDESTINATIONS
		vesselNameStr += `, nộp phí ngày`+rowSets.Rows[j].NGAY_NOP_PHI
		vesselNameStr += ` ` + strings.ToLower(rowSets.Rows[j].TINH_TRANG_THU_PHI)


	}



	if vesselNameNo==1 {
		msgText = "Thông tin hồ sơ:"+"\n"+vesselNameStr1row+vPolicyHolderName +vBU+vINCEPTION_DATE+vEXPIRY_DATE+vPOLICYURN+vGCN+vPhamVi+vCangDi+vCangDen+vNgayNopPhi +vTinhTrangThuPhi+vTradeName+vPhoneTrade
	} else {
		msgText = "Thông tin hồ sơ:"+"\n"+vesselNameVal + vPolicyHolderName +vBU+vINCEPTION_DATE+vEXPIRY_DATE+vPOLICYURN+vGCN+vPhamVi+vCangDi+vCangDen+vNgayNopPhi +vTinhTrangThuPhi+vTradeName+vPhoneTrade+ msgText + vesselNameStr
	}


	//bot.SendText(msg.Sender, msgText+".")
	fmt.Print(msgText)
}