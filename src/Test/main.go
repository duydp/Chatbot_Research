package main
import (
	"soap"
	"fmt"
	"encoding/xml"
	"BVGI/config"
	"strings"


)

func main()  {


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
	var msgText, msgAll  string
	msgAll = ""
	//msgText = `Thông tin hồ sơ:`

	for i := 0; i < len(rowSets.Rows); i++ {
		if len(rowSets.Rows)%1==0{
			check_SUMINSURED_AMT:=rowSets.Rows[i].SUMINSURED_AMT
			check_PhoneTrade := strings.Replace(rowSets.Rows[i].TRADE_PHONE, "/", "", 1)
			check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
			check_NgayDongPhi:=rowSets.Rows[i].NGAY_NOP_PHI


			if rowSets.Rows[i].VESSELORCONVEYANCE != "" {
				msgText+=  "Tàu " + rowSets.Rows[i].VESSELORCONVEYANCE
			}

			if rowSets.Rows[i].POLICYHOLDER_NAME != "" {
				msgText+=  ", khách hàng " + rowSets.Rows[i].POLICYHOLDER_NAME
			}
			if rowSets.Rows[i].BU_NAME != "" {
				msgText+=  `, BH tại ` + rowSets.Rows[i].BU_NAME
			}


			if rowSets.Rows[i].INCEPTION_DATE != "" {
				msgText+=  `, từ ` + rowSets.Rows[i].INCEPTION_DATE
			}

			if rowSets.Rows[i].EXPIRY_DATE != "" {
				msgText+=  ` đến ` + rowSets.Rows[i].EXPIRY_DATE
			}

			if  rowSets.Rows[i].POLICY_URN != "" {
				msgText+=  ". Số đơn " + rowSets.Rows[i].POLICY_URN
			}

			if  rowSets.Rows[i].CERTIFICATENUMBER !=""{
				msgText+= `, GCN ` + rowSets.Rows[i].CERTIFICATENUMBER
			}

			if (check_SUMINSURED_AMT!= "" && !strings.HasPrefix(strings.Trim(rowSets.Rows[i].SUMINSURED_AMT,""),"0 ")) {
				msgText+=`, tham gia BH ` + rowSets.Rows[i].COV_CLASS_NAME + " (" + rowSets.Rows[i].SUMINSURED_AMT + ")"
			}


			if  rowSets.Rows[i].PLACEOFDEPARTURES != "" {
				msgText+=  `, đi từ ` + rowSets.Rows[i].PLACEOFDEPARTURES
			}

			if rowSets.Rows[i].FINALDESTINATIONS != "" {
				msgText+=  ` đến ` + rowSets.Rows[i].FINALDESTINATIONS
			}

			if (check_NopPhi!="0" && check_NgayDongPhi!= "") {
				msgText+=`, nộp phí ngày`+rowSets.Rows[i].NGAY_NOP_PHI
			}

			if rowSets.Rows[i].TINH_TRANG_THU_PHI != "" {
				if rowSets.Rows[i].TINH_TRANG_THU_PHI != "Đã nộp phí" {
					msgText+= ` ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)
				}
			}

			if (rowSets.Rows[i].TRADE_NAME!= "") {
				msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
			}

			if check_PhoneTrade!=""{
				msgText+="("+strings.Replace(rowSets.Rows[i].TRADE_PHONE, "/", "", 1)+")"
			}




			msgText += "."
			msgAll = msgAll + msgText
			//bot.SendText(msg.Sender, msgAll)
			fmt.Print(msgText)
		}
		msgText=""
		msgAll=""


	}


	//bot.SendText(msg.Sender, msgText+".")

}