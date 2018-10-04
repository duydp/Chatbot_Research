//author: duydp
//wrote at: 26/3/2018
package dialog
import (
	"util"
	"github.com/michlabs/fbbot"
	"BVGI/config"
	"soap"
	"encoding/xml"
	"strings"
	"strconv"

	"regexp"
)
type GetPolicyURNFireBot struct {
	fbbot.BaseStep
}
func (s GetPolicyURNFireBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, util.Personalize(T("getSoHopDong_title"), &msg.Sender))
	return StayEvent
}

func (s GetPolicyURNFireBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation") == "" {
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "true")
	}

	var strFileName,strParametername,strIDSender string
	strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
	strFileName="Fire.log"
	strParametername=" |"+string(util.Personalize(T("getSoHopDong_title"), &msg.Sender))+":"
	soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,msg.Text)

	var policyURN_Fire =strings.ToUpper(msg.Text)
	//var policyURN_Fire ="`"+msg.Text+"`"
	//policyURN_Fire=strings.Replace(strconv.Quote(strings.ToUpper(policyURN_Fire)), `"`, " ", 2)
	//
	//pr := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	//gOUTPUT, _, _ := transform.String(pr ,policyURN_Fire)
	//reg, _ := regexp.Compile("[^a-zA-Z0-9-%]+")
	////if err != nil {
	////	log.Fatal(err)
	////}
	//prc_Str:= reg.ReplaceAllString(gOUTPUT, " ")
	//
	//var v_processedString string
	//v_processedString=""
	//var b_pre bool=strings.HasPrefix(strings.TrimSpace(prc_Str),`%`)
	//var b_sur bool=strings.HasSuffix(strings.TrimSpace(prc_Str),`%`)
	//
	//if (b_pre==true && b_sur==false){
	//	prc_Str:=strings.TrimSpace(prc_Str)
	//	prc_Str=prc_Str[1:]
	//	prc_Str=strings.Replace(prc_Str,"%","",-1)
	//	v_processedString=`%`+strings.TrimSpace(prc_Str)
	//}
	//if (b_sur==true && b_pre==false){
	//	prc_Str:=strings.TrimSpace(prc_Str)
	//	prc_Str=prc_Str[0:len(prc_Str)-1]
	//	prc_Str=strings.Replace(prc_Str,"%","",-1)
	//	v_processedString=strings.TrimSpace(prc_Str)+`%`
	//}
	//if (b_pre==true && b_sur==true){
	//	prc_Str:=strings.TrimSpace(prc_Str)
	//	prc_Str=prc_Str[1:]
	//	prc_Str=prc_Str[0:len(prc_Str)-1]
	//	prc_Str=strings.Replace(prc_Str,"%","",-1)
	//	v_processedString=`%`+strings.TrimSpace(prc_Str)+`%`
	//}
	//
	//if (b_pre==false && b_sur==false){
	//	prc_Str:=strings.TrimSpace(prc_Str)
	//	v_processedString=prc_Str
	//}
	result, _ :=  soap.GetPolicyURNFireInfo(policyURN_Fire)


	// Read xml
	if (result==nil) {
		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		return ConfirmInputPolicyUrnFireEvent
	}

	byteValue := []byte(result.GetPolicyFireInfoResult)

	var rowSets config.RowSet
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &rowSets)


	if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		return ConfirmInputPolicyUrnFireEvent
	}

	if len(rowSets.Rows) > 5 {
		iTotal:=len(rowSets.Rows)
		var strCount=strconv.Itoa(iTotal)+" kết quả"
		bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
		bot.STMemory.For(msg.Sender.ID).Set("skey", msg.Text)

		bot.STMemory.For(msg.Sender.ID).Set("svtype", "ts_sodon")

		return ConfirmWebViewBotEvent
	} else {
		var msgText string
		var locationVal= ""
		var locationStr= ""
		var locationStr1row= ""
		var locationNo= 0
		var startCvr = 0
		var vTradeName = ""
		var vNgayNopPhi = ""
		var vTinhTrangThuPhi = ""
		var vPhoneTrade = ""
		var vSoVuBT = ""
		var vTyLeBT = ""
		var vINCEPTION_DATE = ""
		var vEXPIRY_DATE = ""
		var vPOLICYURN = ""
		var vPolicyHolderName = ""
		var vBU = ""
		for j := 0; j < len(rowSets.Rows); j++ {

			if (vPolicyHolderName != rowSets.Rows[j].POLICYHOLDER_NAME && rowSets.Rows[j].POLICYHOLDER_NAME != "") {
				vPolicyHolderName = ", khách hàng " + rowSets.Rows[j].POLICYHOLDER_NAME
			}
			if (vBU != rowSets.Rows[j].BU_NAME && rowSets.Rows[j].BU_NAME != "") {
				vBU = `, BH tại ` + rowSets.Rows[j].BU_NAME
			}
			if (vTradeName != rowSets.Rows[j].TRADE_NAME && rowSets.Rows[j].TRADE_NAME != "") {
				vTradeName = `. LH CBKT ` + rowSets.Rows[j].TRADE_NAME
			}

			if (vINCEPTION_DATE != rowSets.Rows[j].INCEPTION_DATE && rowSets.Rows[j].INCEPTION_DATE != "") {
				vINCEPTION_DATE = `, từ ngày ` + rowSets.Rows[j].INCEPTION_DATE
			}
			if (vEXPIRY_DATE != rowSets.Rows[j].INCEPTION_DATE && rowSets.Rows[j].EXPIRY_DATE != "") {
				vEXPIRY_DATE = ` đến ngày ` + rowSets.Rows[j].EXPIRY_DATE
			}
			if (vPOLICYURN != rowSets.Rows[j].POLICY_URN && rowSets.Rows[j].POLICY_URN != "") {
				vPOLICYURN = "Số đơn " + rowSets.Rows[j].POLICY_URN
			}
			if (vNgayNopPhi != rowSets.Rows[j].NGAY_NOP_PHI && rowSets.Rows[j].NGAY_NOP_PHI != "") {
				vNgayNopPhi = `, nộp phí ngày ` + rowSets.Rows[j].NGAY_NOP_PHI
			}

			if (rowSets.Rows[j].TINH_TRANG_THU_PHI == "" || rowSets.Rows[j].TINH_TRANG_THU_PHI == "Đã nộp phí" ) {
			} else {
				vTinhTrangThuPhi = ", " + strings.ToLower(rowSets.Rows[j].TINH_TRANG_THU_PHI)
			}

			if (rowSets.Rows[j].TRADE_PHONE != "" && len(rowSets.Rows[j].TRADE_PHONE) > 1) {
				var v_string_p string = rowSets.Rows[j].TRADE_PHONE
				var rst string
				v_reg, _ := regexp.Compile("[^Z0-9]+")
				prc_str := v_reg.ReplaceAllString(v_string_p, " ")
				rst = strings.TrimSpace(prc_str)
				rst = strings.Replace(rst, ` `, `/`, -1)
				vPhoneTrade = " (" + rst + ")"
			}

			if (vSoVuBT != rowSets.Rows[j].SO_VU_BT && rowSets.Rows[j].SO_VU_BT != "" && rowSets.Rows[j].SO_VU_BT != "0") {
				vSoVuBT = `. Số vụ tổn thất: ` + rowSets.Rows[j].SO_VU_BT + " vụ,"
			}

			if (vTyLeBT != rowSets.Rows[j].TY_LE_BT && rowSets.Rows[j].TY_LE_BT != "" && rowSets.Rows[j].TY_LE_BT != "0%") {
				vTyLeBT = ` tỷ lệ BT: ` + rowSets.Rows[j].TY_LE_BT
			}

			if (locationVal != rowSets.Rows[j].LOCATION && rowSets.Rows[j].LOCATION != "") {
				if locationVal != "" {
					msgText += locationStr
				}
				locationNo += 1
				startCvr = 1
				locationStr = "\r\n" + strconv.Itoa(locationNo) + "." + rowSets.Rows[j].LOCATION
				if locationNo == 1 {
					locationStr1row = "Số đơn " + rowSets.Rows[j].POLICY_URN + ", " + rowSets.Rows[j].LOCATION + ", khách hàng " + rowSets.Rows[j].POLICYHOLDER_NAME + `, BH tại ` + rowSets.Rows[j].BU_NAME
					locationStr1row += `, từ ngày ` + rowSets.Rows[j].INCEPTION_DATE + ` đến ngày ` + rowSets.Rows[j].EXPIRY_DATE
					locationStr1row += ", tham gia BH "
				}

				locationVal = rowSets.Rows[j].LOCATION

			}

			if startCvr == 0 {
				locationStr += ", "
				locationStr1row += ", "
			}
			locationStr += " tham gia BH " + rowSets.Rows[j].COV_CLASS_NAME
			locationStr1row += rowSets.Rows[j].COV_CLASS_NAME
			if (rowSets.Rows[j].SUMINSURED_AMT != "" && !strings.HasPrefix(strings.Trim(rowSets.Rows[j].SUMINSURED_AMT, ""), "0 ")) {
				locationStr += " (" + rowSets.Rows[j].SUMINSURED_AMT + ")"
				locationStr1row += " (" + rowSets.Rows[j].SUMINSURED_AMT + ")"
			}
			startCvr = 0

		}

		if locationNo == 1 {
			msgText = "Thông tin hồ sơ:" + "\n" + locationStr1row + vTinhTrangThuPhi + vNgayNopPhi + vSoVuBT + vTyLeBT + vTradeName + vPhoneTrade
		} else {
			msgText = "Thông tin hồ sơ:" + "\n" + vPOLICYURN + vPolicyHolderName + vBU + vINCEPTION_DATE + vEXPIRY_DATE + vTinhTrangThuPhi + vNgayNopPhi + vSoVuBT + vTyLeBT + vTradeName + vPhoneTrade + " .Địa điểm bảo hiểm:" + msgText + locationStr
		}

		bot.SendText(msg.Sender, msgText+".")
	}
	return ConfirmInputPolicyUrnFireEvent
}