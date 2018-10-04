package dialog

import (

	"util"
	"github.com/michlabs/fbbot"
	"soap"
	"encoding/xml"
	"BVGI/config"
	"strings"
	"strconv"
	"regexp"
)

type GetInfoBot struct {
	fbbot.BaseStep
}

func (s GetInfoBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	if (bot.STMemory.For(msg.Sender.ID).Get("Redirect_From_Hook")!="true") {
		bot.SendText(msg.Sender, util.Personalize(T("getSyntax_title"), &msg.Sender))
	} else {
		bot.STMemory.For(msg.Sender.ID).Set("Redirect_From_Hook","false")
	}
	return StayEvent

}

func (s GetInfoBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	strMsg := strings.ToUpper(msg.Text)
	//var vMsg ="`"+msg.Text+"`"
	//vMsg=strings.Replace(strconv.Quote(strings.ToUpper(vMsg)), `"`, " ", 2)
	//var strMsg string
	//pr := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	//gOUTPUT, _, _ := transform.String(pr ,strings.Replace(vMsg,"Đ","D",-1))
	//reg, _ := regexp.Compile("[^a-zA-Z0-9-%]+")
	////if err != nil {
	////	log.Fatal(err)
	////}
	//v_processedString := reg.ReplaceAllString(gOUTPUT, " ")
	//strMsg=strings.TrimSpace(v_processedString)


	if (!strings.HasPrefix(strings.ToUpper(strMsg), "CAR ") && !strings.HasPrefix(strings.ToUpper(strMsg), "GCN ") && !strings.HasPrefix(strings.ToUpper(strMsg), "SK ") && !strings.HasPrefix(strings.ToUpper(strMsg), "CMT ") && !strings.HasPrefix(strings.ToUpper(strMsg), "MST ") && !strings.HasPrefix(strings.ToUpper(strMsg), "SD ") && !strings.HasPrefix(strings.ToUpper(strMsg), "TAU ") && !strings.HasPrefix(strings.ToUpper(strMsg), "HH ") && !strings.HasPrefix(strings.ToUpper(strMsg), "IMO ") && !strings.HasPrefix(strings.ToUpper(strMsg), "DK ") && !strings.HasPrefix(strings.ToUpper(strMsg), "DD ")&& !strings.HasPrefix(strings.ToUpper(strMsg), "TEN ") && !strings.HasPrefix(strings.ToUpper(strMsg), "KH ") && !strings.HasPrefix(strings.ToUpper(strMsg), "HD ")) {
		bot.STMemory.For(msg.Sender.ID).Set("SyntaxInvalid", "true")
		return ConfirmSyntaxEvent
	}else{
		bot.STMemory.For(msg.Sender.ID).Set("SyntaxInvalid", "false")
	}

	var strKey = bot.STMemory.For(msg.Sender.ID).Get("KEY_BSX")

	//stringKey="`"+stringKey+"`"
	//stringKey=strings.Replace(strconv.Quote(strings.ToUpper(stringKey)), `"`, " ", 2)
	//rp := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	//gKEY, _, _ := transform.String(rp ,strings.Replace(stringKey,"Đ","D",-1))
	//
	//registrator, _ := regexp.Compile("[^a-zA-Z0-9-%]+")
	////if err_key != nil {
	////	log.Fatal(err_key)
	////}
	//prc_Str:= registrator.ReplaceAllString(gKEY, " ")
	//
	//var v_processedKey string
	//v_processedKey=""
	//var b_pre bool=strings.HasPrefix(strings.TrimSpace(prc_Str),`%`)
	//var b_sur bool=strings.HasSuffix(strings.TrimSpace(prc_Str),`%`)
	//
	//if (b_pre==true && b_sur==false){
	//	prc_Str:=strings.TrimSpace(prc_Str)
	//	prc_Str=prc_Str[1:]
	//	prc_Str=strings.Replace(prc_Str,"%","",-1)
	//	v_processedKey=`%`+strings.TrimSpace(prc_Str)
	//}
	//if (b_sur==true && b_pre==false){
	//	prc_Str:=strings.TrimSpace(prc_Str)
	//	prc_Str=prc_Str[0:len(prc_Str)-1]
	//	prc_Str=strings.Replace(prc_Str,"%","",-1)
	//	v_processedKey=strings.TrimSpace(prc_Str)+`%`
	//}
	//if (b_pre==true && b_sur==true){
	//	prc_Str:=strings.TrimSpace(prc_Str)
	//	prc_Str=prc_Str[1:]
	//	prc_Str=prc_Str[0:len(prc_Str)-1]
	//	prc_Str=strings.Replace(prc_Str,"%","",-1)
	//	v_processedKey=`%`+strings.TrimSpace(prc_Str)+`%`
	//}
	//
	//if (b_pre==false && b_sur==false){
	//	prc_Str:=strings.TrimSpace(prc_Str)
	//	v_processedKey=prc_Str
	//}
	//
	//var strKey=v_processedKey
	//var strKey = bot.STMemory.For(msg.Sender.ID).Get("KEY_BSX")
	var msgText, msgAll  string
	msgAll = ""


	if (bot.STMemory.For(msg.Sender.ID).Get("Action")=="xecogioi") || (bot.STMemory.For(msg.Sender.ID).Get("Action")=="biensoxe") || (bot.STMemory.For(msg.Sender.ID).Get("Action")=="soGCNBH") || (bot.STMemory.For(msg.Sender.ID).Get("Action")=="soK") || (bot.STMemory.For(msg.Sender.ID).Get("Action")=="soCMT"){

		if (strings.HasPrefix(strings.ToUpper(strMsg), "GCN ")) {

			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Car.log"
			strParametername=" |GCN: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)

			result, _ :=  soap.GetcertNumberCarInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}

			// Read xml
			byteValue := []byte(result.GetCarInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				//+ rowSets.Rows[0].DATA
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}

			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "sogcnbh")

				return ConfirmWebViewBotEvent
			} else {
				msgText="Thông tin hồ sơ: "
				for i := 0; i < len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{

						check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
						check_CertNum:=rowSets.Rows[i].CERT_NUMBER


						check_TNDSBB:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].TNDS_BAT_BUOC,""),"0tr")
						check_TNDSTN:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].TNDS_TU_NGUYEN,""),"0tr")
						check_sotienVCX:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].VAT_CHAT_XE,""),"0tr")
						check_LaiPhu:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].LAI_PHU,""),"0tr")
						check_HangHoa:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].HANG_HOA,""),"0tr")

						check_SoVuBT:=rowSets.Rows[i].SO_VU_BT
						check_TyLeBT:=rowSets.Rows[i].TY_LE_BT

						if(rowSets.Rows[i].REG_NUMBER!=""){
							msgText += "\n" + `BKS: ` + rowSets.Rows[i].REG_NUMBER
						}

						msgText+=`, chủ xe ` + rowSets.Rows[i].POLICYHOLDER_NAME

						if (rowSets.Rows[i].POLICYHOLDER_PHONE!="" && len(rowSets.Rows[i].POLICYHOLDER_PHONE)>1){
							var v_string string=rowSets.Rows[i].POLICYHOLDER_PHONE
							var rst_hP string
							v_reg1, _ := regexp.Compile("[^Z0-9]+")
							prc_str1:= v_reg1.ReplaceAllString(v_string, " ")
							rst_hP=strings.TrimSpace(prc_str1)
							rst_hP=strings.Replace(rst_hP,` `,`/`,-1)
							msgText+=" ("+rst_hP+")"
						}
						msgText+=`, BH tại ` + rowSets.Rows[i].BU_NAME +`-`+ rowSets.Rows[i].DEPT_NAME
						msgText+=`, từ ` + rowSets.Rows[i].INCEPTION_DATE

						msgText+=` đến ` + rowSets.Rows[i].EXPIRY_DATE

						msgText+=`. Số đơn ` +rowSets.Rows[i].POLICY_URN
						if check_CertNum!=""{
							if check_CertNum!="0"{
								msgText+=`, GCN ` + rowSets.Rows[i].CERT_NUMBER
							}
						}

						if check_TNDSBB !=true{
							msgText+=`, TNDSBB ` + "(" +rowSets.Rows[i].TNDS_BAT_BUOC+ ")"
						}
						if check_TNDSTN !=true{
							msgText+=`, TNDSTN ` + "(" +rowSets.Rows[i].TNDS_TU_NGUYEN+ ")"
						}
						//check_sotienVCX:=rowSets.Rows[i].VAT_CHAT_XE
						if check_sotienVCX!=true{
							msgText+=`, VCXE ` + "("+rowSets.Rows[i].VAT_CHAT_XE+")"}

						if check_LaiPhu!=true{
							msgText+=`, NTX ` + "("+rowSets.Rows[i].LAI_PHU+")"
						}

						if check_HangHoa!=true{
							msgText+=`, TNDS HH ` + "("+rowSets.Rows[i].HANG_HOA+")"
						}

						if check_NopPhi!="0"{
							msgText+=`, nộp phí ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
						}

						if rowSets.Rows[i].TINH_TRANG_THU_PHI!="" {
							if rowSets.Rows[i].TINH_TRANG_THU_PHI!="Đã nộp phí"{
								msgText+=`, ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)
							}
						}


						if check_SoVuBT!="0" {
							msgText += `. Số vụ tổn thất: ` + rowSets.Rows[i].SO_VU_BT + " vụ,"
						}
						if (check_TyLeBT!="0%") && (check_TyLeBT!="%") {
							msgText += ` tỷ lệ BT: ` + rowSets.Rows[i].TY_LE_BT
						}

						if (rowSets.Rows[i].TRADE_NAME!= "") {
							msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
						}

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}
						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
					}
					msgText=""
					msgAll=""
				}
			}

		}

		if (strings.HasPrefix(strings.ToUpper(strMsg), "SK ")) {
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Car.log"
			strParametername=" |SK: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)

			result, _ :=  soap.GetchassisNoCarInfo(strKey)

			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}

			// Read xml
			byteValue := []byte(result.GetCarInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				//+ rowSets.Rows[0].DATA
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}

			msgAll = ""

			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)
				bot.STMemory.For(msg.Sender.ID).Set("svtype", "sok")
				return ConfirmWebViewBotEvent
			} else {
				msgText="Thông tin hồ sơ: "
				for i := 0; i < len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{

						check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
						//check_CertNum:=rowSets.Rows[i].CERT_NUMBER


						check_TNDSBB:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].TNDS_BAT_BUOC,""),"0tr")
						check_TNDSTN:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].TNDS_TU_NGUYEN,""),"0tr")
						check_sotienVCX:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].VAT_CHAT_XE,""),"0tr")
						check_LaiPhu:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].LAI_PHU,""),"0tr")
						check_HangHoa:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].HANG_HOA,""),"0tr")

						check_SoVuBT:=rowSets.Rows[i].SO_VU_BT
						check_TyLeBT:=rowSets.Rows[i].TY_LE_BT

						if(rowSets.Rows[i].CHASSIS_NO!=""){
							msgText += "\n" + `BKS/SK: ` + rowSets.Rows[i].REG_NUMBER + "/" + rowSets.Rows[i].CHASSIS_NO
						}

						msgText+=`, chủ xe ` + rowSets.Rows[i].POLICYHOLDER_NAME

						if (rowSets.Rows[i].POLICYHOLDER_PHONE!="" && len(rowSets.Rows[i].POLICYHOLDER_PHONE)>1){
							var v_string string=rowSets.Rows[i].POLICYHOLDER_PHONE
							var rst_hP string
							v_reg1, _ := regexp.Compile("[^Z0-9]+")
							prc_str1:= v_reg1.ReplaceAllString(v_string, " ")
							rst_hP=strings.TrimSpace(prc_str1)
							rst_hP=strings.Replace(rst_hP,` `,`/`,-1)
							msgText+=" ("+rst_hP+")"
						}
						msgText+=`, BH tại ` + rowSets.Rows[i].BU_NAME +`-`+ rowSets.Rows[i].DEPT_NAME
						msgText+=`, từ ` + rowSets.Rows[i].INCEPTION_DATE

						msgText+=` đến ` + rowSets.Rows[i].EXPIRY_DATE

						msgText+=`. Số đơn ` +rowSets.Rows[i].POLICY_URN
						//if check_CertNum!=""{
						//	if check_CertNum!="0"{
						//		msgText+=`, GCN ` + rowSets.Rows[i].CERT_NUMBER
						//	}
						//}

						if check_TNDSBB !=true{
							msgText+=`, TNDSBB ` + "(" +rowSets.Rows[i].TNDS_BAT_BUOC+ ")"
						}
						if check_TNDSTN !=true{
							msgText+=`, TNDSTN ` + "(" +rowSets.Rows[i].TNDS_TU_NGUYEN+ ")"
						}
						//check_sotienVCX:=rowSets.Rows[i].VAT_CHAT_XE
						if check_sotienVCX!=true{
							msgText+=`, VCXE ` + "("+rowSets.Rows[i].VAT_CHAT_XE+")"}

						if check_LaiPhu!=true{
							msgText+=`, NTX ` + "("+rowSets.Rows[i].LAI_PHU+")"
						}

						if check_HangHoa!=true{
							msgText+=`, TNDS HH ` + "("+rowSets.Rows[i].HANG_HOA+")"
						}

						if check_NopPhi!="0"{
							msgText+=`, nộp phí ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
						}

						if rowSets.Rows[i].TINH_TRANG_THU_PHI!="" {
							if rowSets.Rows[i].TINH_TRANG_THU_PHI!="Đã nộp phí"{
								msgText+=`, ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)
							}
						}


						if check_SoVuBT!="0" {
							msgText += `. Số vụ tổn thất: ` + rowSets.Rows[i].SO_VU_BT + " vụ,"
						}
						if (check_TyLeBT!="0%") && (check_TyLeBT!="%") {
							msgText += ` tỷ lệ BT: ` + rowSets.Rows[i].TY_LE_BT
						}

						if (rowSets.Rows[i].TRADE_NAME!= "") {
							msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
						}

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}
						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
					}
					msgText=""
					msgAll=""
				}
			}

			//bot.STMemory.For(msg.Sender.ID).Set("Action","")
		}

	}else if (bot.STMemory.For(msg.Sender.ID).Get("Action")=="taisan") || (bot.STMemory.For(msg.Sender.ID).Get("Action")=="ts_sodon") || (bot.STMemory.For(msg.Sender.ID).Get("Action")=="ts_mst") || (bot.STMemory.For(msg.Sender.ID).Get("Action")=="ts_cmt") || (bot.STMemory.For(msg.Sender.ID).Get("Action")=="ts_location")|| (bot.STMemory.For(msg.Sender.ID).Get("Action")=="ts_tkh"){

		if (strings.HasPrefix(strings.ToUpper(strMsg), "SD ")) {
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Fire.log"
			strParametername=" |SD: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
			result, _ :=  soap.GetPolicyURNFireInfo(strKey)
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

			var locationVal = ""
			var locationStr = ""
			var locationStr1row = ""
			var locationNo = 0
			var startCvr=0
			var vTradeName=""
			var vNgayNopPhi=""
			var vTinhTrangThuPhi=""
			var vPhoneTrade=""
			var vSoVuBT=""
			var vTyLeBT=""
			var vINCEPTION_DATE=""
			var vEXPIRY_DATE=""
			var vPOLICYURN=""
			var vPolicyHolderName=""
			var vBU=""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "ts_sodon")

				return ConfirmWebViewBotEvent
			} else {
				msgText="Thông tin hồ sơ: "
				for j:=0;j<len(rowSets.Rows); j++ {

					if (vPolicyHolderName != rowSets.Rows[j].POLICYHOLDER_NAME && rowSets.Rows[j].POLICYHOLDER_NAME != "") {
						vPolicyHolderName=  ", khách hàng " + rowSets.Rows[j].POLICYHOLDER_NAME
					}
					if (vBU != rowSets.Rows[j].BU_NAME && rowSets.Rows[j].BU_NAME != "") {
						vBU=  `, BH tại ` + rowSets.Rows[j].BU_NAME
					}
					if (vTradeName!= rowSets.Rows[j].TRADE_NAME && rowSets.Rows[j].TRADE_NAME!= "") {
						vTradeName=`. LH CBKT ` +rowSets.Rows[j].TRADE_NAME
					}

					if (vINCEPTION_DATE != rowSets.Rows[j].INCEPTION_DATE && rowSets.Rows[j].INCEPTION_DATE != "") {
						vINCEPTION_DATE=  `, từ ngày ` + rowSets.Rows[j].INCEPTION_DATE
					}
					if (vEXPIRY_DATE != rowSets.Rows[j].INCEPTION_DATE && rowSets.Rows[j].EXPIRY_DATE != "") {
						vEXPIRY_DATE=  ` đến ngày ` + rowSets.Rows[j].EXPIRY_DATE
					}
					if (vPOLICYURN != rowSets.Rows[j].POLICY_URN && rowSets.Rows[j].POLICY_URN != "") {
						vPOLICYURN=  "Số đơn " + rowSets.Rows[j].POLICY_URN
					}
					if (vNgayNopPhi!= rowSets.Rows[j].NGAY_NOP_PHI && rowSets.Rows[j].NGAY_NOP_PHI!= "") {
						vNgayNopPhi=`, nộp phí ngày `+rowSets.Rows[j].NGAY_NOP_PHI
					}
					if (rowSets.Rows[j].TINH_TRANG_THU_PHI != "" && rowSets.Rows[j].TINH_TRANG_THU_PHI == "Đã nộp phí" ) {
					}else {
						vTinhTrangThuPhi = ", "+strings.ToLower(rowSets.Rows[j].TINH_TRANG_THU_PHI)
					}


					if (rowSets.Rows[j].TRADE_PHONE !="" && len(rowSets.Rows[j].TRADE_PHONE)>1){
						var v_string_p string=rowSets.Rows[j].TRADE_PHONE
						var rst string
						v_reg, _ := regexp.Compile("[^Z0-9]+")
						prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
						rst=strings.TrimSpace(prc_str)
						rst=strings.Replace(rst,` `,`/`,-1)
						vPhoneTrade=" ("+rst+")"
					}

					if (vSoVuBT!= rowSets.Rows[j].SO_VU_BT && rowSets.Rows[j].SO_VU_BT!= ""&& rowSets.Rows[j].SO_VU_BT!= "0") {
						vSoVuBT=`. Số vụ tổn thất: ` + rowSets.Rows[j].SO_VU_BT + " vụ,"
					}

					if (vTyLeBT!= rowSets.Rows[j].TY_LE_BT && rowSets.Rows[j].TY_LE_BT!= "" && rowSets.Rows[j].TY_LE_BT!= "0%") {
						vTyLeBT=` tỷ lệ BT: ` + rowSets.Rows[j].TY_LE_BT
					}

					if (locationVal != rowSets.Rows[j].LOCATION && rowSets.Rows[j].LOCATION != "") {
						if locationVal!="" {
							msgText += locationStr
						}
						locationNo +=1
						startCvr = 1
						locationStr = "\r\n" + strconv.Itoa(locationNo)+"." + rowSets.Rows[j].LOCATION
						if locationNo==1 {
							locationStr1row ="Số đơn " + rowSets.Rows[j].POLICY_URN+", " +rowSets.Rows[j].LOCATION + ", khách hàng " + rowSets.Rows[j].POLICYHOLDER_NAME + `, BH tại ` + rowSets.Rows[j].BU_NAME
							locationStr1row += `, từ ngày ` + rowSets.Rows[j].INCEPTION_DATE + ` đến ngày ` + rowSets.Rows[j].EXPIRY_DATE
							locationStr1row +=  ", tham gia BH "
						}

						locationVal=rowSets.Rows[j].LOCATION

					}

					if startCvr==0 {
						locationStr +=", "
						locationStr1row += ", "
					}
					locationStr += " tham gia BH " + rowSets.Rows[j].COV_CLASS_NAME
					locationStr1row	+= rowSets.Rows[j].COV_CLASS_NAME
					if (rowSets.Rows[j].SUMINSURED_AMT!= "" && !strings.HasPrefix(strings.Trim(rowSets.Rows[j].SUMINSURED_AMT,""),"0 ")) {
						locationStr += " (" + rowSets.Rows[j].SUMINSURED_AMT + ")"
						locationStr1row	+= " (" + rowSets.Rows[j].SUMINSURED_AMT + ")"

					}
					startCvr=0

				}

				if locationNo==1 {
					msgText = "Thông tin hồ sơ:"+"\n"+locationStr1row+vTinhTrangThuPhi+vNgayNopPhi+vSoVuBT+vTyLeBT+vTradeName+vPhoneTrade
				} else {
					msgText = "Thông tin hồ sơ:"+"\n"+vPOLICYURN+vPolicyHolderName +vBU+vINCEPTION_DATE+vEXPIRY_DATE+vTinhTrangThuPhi+vNgayNopPhi+vSoVuBT+vTyLeBT+vTradeName+vPhoneTrade+" .Địa điểm bảo hiểm:"+ msgText + locationStr
				}

				bot.SendText(msg.Sender, msgText+".")
			}
		}

		//if (strings.HasPrefix(strings.ToUpper(strMsg), "MST ")) {
		//	var strFileName,strParametername,strIDSender string
		//	strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
		//	strFileName="Fire.log"
		//	strParametername=" |MST: "
		//	soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
		//
		//	result, _ :=  soap.GetPolicyHolderTaxCodeFireInfo(strKey)
		//
		//	if (result==nil) {
		//		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		//		return ConfirmTS_MST_Event
		//	}
		//	// Read xml
		//	byteValue := []byte(result.GetPolicyFireInfoResult)
		//
		//	var rowSets config.RowSet
		//	// we unmarshal our byteArray which contains our
		//	// xmlFiles content into 'users' which we defined above
		//	xml.Unmarshal(byteValue, &rowSets)
		//	if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
		//		//+ rowSets.Rows[0].DATA
		//		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		//		return ConfirmTS_MST_Event
		//	}
		//	//var msgText, msgAll  string
		//	msgAll = ""
		//
		//	msgText = `Thông tin hồ sơ:` + "\n"
		//	for i := 0; i < len(rowSets.Rows); i++ {
		//		msgText="Thông tin hồ sơ:"
		//		if len(rowSets.Rows)%1==0{
		//			check_NopPhi:=rowSets.Rows[i].SUMINSURED_AMT
		//
		//			check_NgayNopPhi:=strings.Replace(rowSets.Rows[i].NGAY_NOP_PHI,"/", "",1)
		//
		//
		//			msgText+=`Số đơn ` + rowSets.Rows[i].POLICY_URN
		//
		//			msgText+=`, khách hàng ` + rowSets.Rows[i].TRADE_NAME
		//
		//			msgText+=`, tham gia ` + rowSets.Rows[i].PRODUCT_NAME
		//
		//			if check_NopPhi!="0"{
		//				msgText+=` ` + "("+rowSets.Rows[i].SUMINSURED_AMT+")"}
		//
		//			msgText+=`, địa điểm : `+rowSets.Rows[i].LOCATION
		//
		//			msgText+=`, từ ` + rowSets.Rows[i].INCEPTION_DATE
		//
		//			msgText+=` đến ` + rowSets.Rows[i].EXPIRY_DATE
		//
		//			if check_NgayNopPhi!=""{
		//				msgText+=` ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
		//			}
		//
		//			msgText+=`, tại ` + rowSets.Rows[i].BU_NAME
		//
		//			if (rowSets.Rows[i].TRADE_NAME!= "") {
		//				msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
		//			}
		//
		//			if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
		//				var v_string_p string=rowSets.Rows[i].TRADE_PHONE
		//				var rst string
		//				v_reg, _ := regexp.Compile("[^Z0-9]+")
		//				prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
		//				rst=strings.TrimSpace(prc_str)
		//				rst=strings.Replace(rst,` `,`/`,-1)
		//				msgText+=" ("+rst+")"
		//			}
		//
		//
		//			msgText += "."
		//			msgAll = msgAll + msgText
		//			bot.SendText(msg.Sender, msgAll)
		//		}
		//		msgText=""
		//		msgAll=""
		//	}
		//}

		//if (strings.HasPrefix(strings.ToUpper(strMsg), "CMT ")) {
		//	var strFileName,strParametername,strIDSender string
		//	strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
		//	strFileName="Fire.log"
		//	strParametername=" |CMT: "
		//	soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
		//	result, _ :=  soap.GetpolicyHolderIDFireInfo(strKey)
		//	if (result==nil) {
		//		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		//		return ConfirmInputPolicyHolderIDFireEvent
		//	}
		//	// Read xml
		//	byteValue := []byte(result.GetPolicyFireInfoResult)
		//
		//	var rowSets config.RowSet
		//	// we unmarshal our byteArray which contains our
		//	// xmlFiles content into 'users' which we defined above
		//	xml.Unmarshal(byteValue, &rowSets)
		//	if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
		//		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		//		return ConfirmInputPolicyHolderIDFireEvent
		//	}
		//	//var msgText, msgAll  string
		//	msgAll = ""
		//	msgText = `Thông tin hồ sơ:` + "\n"
		//
		//	for i := 0; i < len(rowSets.Rows); i++ {
		//		if len(rowSets.Rows)%1==0{
		//			check_SUMINSURED_AMT:=rowSets.Rows[i].SUMINSURED_AMT
		//
		//			check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
		//			check_NgayDongPhi:=rowSets.Rows[i].NGAY_NOP_PHI
		//
		//			msgText+="\n" +"Số đơn "+rowSets.Rows[i].POLICY_URN+", khách hàng "+ rowSets.Rows[i].POLICYHOLDER_NAME
		//
		//			msgText +=`, tham gia `
		//			if check_SUMINSURED_AMT !="0"{
		//				msgText+= rowSets.Rows[i].PRODUCT_NAME
		//				msgText+=" (" +check_SUMINSURED_AMT+ ")"
		//			}
		//
		//
		//			msgText+=`, địa điểm : `+rowSets.Rows[i].LOCATION
		//
		//			msgText+=`, thời hạn BH từ ` + rowSets.Rows[i].INCEPTION_DATE
		//
		//			msgText+=` đến ` + rowSets.Rows[i].EXPIRY_DATE
		//
		//			msgText+=`, ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)
		//
		//			if check_NopPhi!="0"{
		//				if check_NgayDongPhi!=""{msgText+=` ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
		//				}
		//			}
		//			msgText+=`, tại ` + rowSets.Rows[i].BU_NAME
		//			if (rowSets.Rows[i].TRADE_NAME!= "") {
		//				msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
		//			}
		//
		//			if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
		//				var v_string_p string=rowSets.Rows[i].TRADE_PHONE
		//				var rst string
		//				v_reg, _ := regexp.Compile("[^Z0-9]+")
		//				prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
		//				rst=strings.TrimSpace(prc_str)
		//				rst=strings.Replace(rst,` `,`/`,-1)
		//				msgText+=" ("+rst+")"
		//			}
		//			msgText += "."
		//			msgAll = msgAll + msgText
		//			bot.SendText(msg.Sender, msgAll)
		//		}
		//		msgText=""
		//		msgAll=""
		//
		//	}
		//
		//}

		if (strings.HasPrefix(strings.ToUpper(strMsg), "DD ")) {
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Fire.log"
			strParametername=" |DD: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
			result, _ :=  soap.GetLocationFireInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmLocationFireEvent
			}
			// Read xml
			byteValue := []byte(result.GetPolicyFireInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmLocationFireEvent
			}


			msgText=""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "ts_location")

				return ConfirmWebViewBotEvent
			} else {
		//process records less than 5
				msgText="Thông tin hồ sơ: "
				for j:=0;j<len(rowSets.Rows); j++ {
					if len(rowSets.Rows)%1==0{
						msgText += "\n"
						if (strings.Contains(msgText, rowSets.Rows[j].LOCATION) && rowSets.Rows[j].LOCATION == ""){
						}else{
							msgText+=rowSets.Rows[j].LOCATION
						}

						if (strings.Contains(msgText, rowSets.Rows[j].POLICYHOLDER_NAME) && rowSets.Rows[j].POLICYHOLDER_NAME == ""){
						}else{
							msgText+=", khách hàng " + rowSets.Rows[j].POLICYHOLDER_NAME
						}

						if (strings.Contains(msgText, rowSets.Rows[j].BU_NAME) && rowSets.Rows[j].BU_NAME == ""){
						}else{
							msgText+=`, BH tại ` + rowSets.Rows[j].BU_NAME
						}

						if (strings.Contains(msgText, rowSets.Rows[j].INCEPTION_DATE) && rowSets.Rows[j].INCEPTION_DATE == ""){
						}else{
							msgText+=`, từ ngày ` + rowSets.Rows[j].INCEPTION_DATE
						}

						if (strings.Contains(msgText, rowSets.Rows[j].EXPIRY_DATE) && rowSets.Rows[j].EXPIRY_DATE == ""){
						}else{
							msgText+= ` đến ngày ` + rowSets.Rows[j].EXPIRY_DATE
						}

						if (strings.Contains(msgText, rowSets.Rows[j].POLICY_URN) && rowSets.Rows[j].POLICY_URN == ""){
						}else{
							msgText+= ". Số đơn " + rowSets.Rows[j].POLICY_URN
						}

						if (strings.Contains(msgText, rowSets.Rows[j].COV_CLASS_NAME) && rowSets.Rows[j].COV_CLASS_NAME == ""){
						}else{
							msgText+= ", tham gia BH "+rowSets.Rows[j].COV_CLASS_NAME
						}
						if (strings.HasPrefix(strings.Trim(rowSets.Rows[j].SUMINSURED_AMT,""),"0") && rowSets.Rows[j].SUMINSURED_AMT == ""){
						}else{
							msgText+= " (" +rowSets.Rows[j].SUMINSURED_AMT + ")"
						}

						if (rowSets.Rows[j].TINH_TRANG_THU_PHI != "" && rowSets.Rows[j].TINH_TRANG_THU_PHI == "Đã nộp phí") {
						}else{
							msgText+= ", "+strings.ToLower(rowSets.Rows[j].TINH_TRANG_THU_PHI)
						}
						if (strings.Contains(msgText, rowSets.Rows[j].NGAY_NOP_PHI) && rowSets.Rows[j].NGAY_NOP_PHI == ""){
						}else{
							msgText+= `, nộp phí ngày `+rowSets.Rows[j].NGAY_NOP_PHI
						}

						if (strings.Contains(msgText, rowSets.Rows[j].SO_VU_BT) || rowSets.Rows[j].SO_VU_BT== "" || rowSets.Rows[j].SO_VU_BT== "0"){
						}else{

							msgText+= `. Số vụ tổn thất: ` + rowSets.Rows[j].SO_VU_BT + " vụ,"
						}

						if (strings.Contains(msgText, rowSets.Rows[j].TY_LE_BT) || rowSets.Rows[j].TY_LE_BT== "" || rowSets.Rows[j].TY_LE_BT== "0%"){
						}else{
							msgText+= ` tỷ lệ BT: ` + rowSets.Rows[j].TY_LE_BT
						}

						if (strings.Contains(msgText, rowSets.Rows[j].TRADE_NAME) && rowSets.Rows[j].TRADE_NAME != ""){
						}else{
							msgText+= `. LH CBKT ` +rowSets.Rows[j].TRADE_NAME
						}

						if (rowSets.Rows[j].TRADE_PHONE !="" && len(rowSets.Rows[j].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[j].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}
						msgText += "."
						msgAll = msgAll + msgText

						bot.SendText(msg.Sender, msgAll)
					}
					msgText=""
					msgAll=""
				}
			}

		}
		if (strings.HasPrefix(strings.ToUpper(strMsg), "KH ")) {
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Fire.log"
			strParametername=" |KH: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
			result, _ :=  soap.GetPolicyHolderNameFireInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmPolicyHolderNamenFireEvent
			}
			// Read xml
			byteValue := []byte(result.GetPolicyFireInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmPolicyHolderNamenFireEvent
			}
			msgAll = ""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "ts_tkh")

				return ConfirmWebViewBotEvent
			} else {
				msgText="Thông tin hồ sơ: "
				for j:=0;j<len(rowSets.Rows); j++ {
					if len(rowSets.Rows)%1==0{
						msgText += "\n"
						if (strings.Contains(msgText, rowSets.Rows[j].POLICYHOLDER_NAME) || rowSets.Rows[j].POLICYHOLDER_NAME == ""){
						}else{
							msgText+="Khách hàng " + rowSets.Rows[j].POLICYHOLDER_NAME
						}
						if (strings.Contains(msgText, rowSets.Rows[j].LOCATION) || rowSets.Rows[j].LOCATION == ""){
						}else{
							msgText+=", "+rowSets.Rows[j].LOCATION
						}

						if (strings.Contains(msgText, rowSets.Rows[j].BU_NAME) || rowSets.Rows[j].BU_NAME == ""){
						}else{
							msgText+=`, BH tại ` + rowSets.Rows[j].BU_NAME
						}

						if (strings.Contains(msgText, rowSets.Rows[j].INCEPTION_DATE) || rowSets.Rows[j].INCEPTION_DATE == ""){
						}else{
							msgText+=`, từ ngày ` + rowSets.Rows[j].INCEPTION_DATE
						}

						if (strings.Contains(msgText, rowSets.Rows[j].EXPIRY_DATE) || rowSets.Rows[j].EXPIRY_DATE == ""){
						}else{
							msgText+= ` đến ngày ` + rowSets.Rows[j].EXPIRY_DATE
						}

						if (strings.Contains(msgText, rowSets.Rows[j].POLICY_URN) || rowSets.Rows[j].POLICY_URN == ""){
						}else{
							msgText+= ". Số đơn " + rowSets.Rows[j].POLICY_URN
						}

						if (strings.Contains(msgText, rowSets.Rows[j].COV_CLASS_NAME) || rowSets.Rows[j].COV_CLASS_NAME == ""){
						}else{
							msgText+= ", tham gia BH "+rowSets.Rows[j].COV_CLASS_NAME
						}
						if (strings.HasPrefix(strings.Trim(rowSets.Rows[j].SUMINSURED_AMT,""),"0") || rowSets.Rows[j].SUMINSURED_AMT == ""){
						}else{
							msgText+= " (" +rowSets.Rows[j].SUMINSURED_AMT + ")"
						}

						if (rowSets.Rows[j].TINH_TRANG_THU_PHI != "" || rowSets.Rows[j].TINH_TRANG_THU_PHI == "Đã nộp phí") {
						}else{
							msgText+= ", "+strings.ToLower(rowSets.Rows[j].TINH_TRANG_THU_PHI)
						}
						if (strings.Contains(msgText, rowSets.Rows[j].NGAY_NOP_PHI) || rowSets.Rows[j].NGAY_NOP_PHI == ""){
						}else{
							msgText+= `, nộp phí ngày `+rowSets.Rows[j].NGAY_NOP_PHI
						}

						if (strings.Contains(msgText, rowSets.Rows[j].SO_VU_BT) || rowSets.Rows[j].SO_VU_BT== "" || rowSets.Rows[j].SO_VU_BT== "0"){
						}else{
							msgText+= `. Số vụ tổn thất: ` + rowSets.Rows[j].SO_VU_BT + " vụ,"
						}

						if (strings.Contains(msgText, rowSets.Rows[j].TY_LE_BT) || rowSets.Rows[j].TY_LE_BT== "" || rowSets.Rows[j].TY_LE_BT== "0%"){
						}else{
							msgText+= ` tỷ lệ BT: ` + rowSets.Rows[j].TY_LE_BT
						}

						if (strings.Contains(msgText, rowSets.Rows[j].TRADE_NAME) || rowSets.Rows[j].TRADE_NAME == ""){
						}else{
							msgText+= `. LH CBKT ` +rowSets.Rows[j].TRADE_NAME
						}

						if (rowSets.Rows[j].TRADE_PHONE !="" && len(rowSets.Rows[j].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[j].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}
						msgText += "."
						msgAll = msgAll + msgText

						bot.SendText(msg.Sender, msgAll)
					}
					msgText=""
					msgAll=""
				}
			}


		}
		//bot.STMemory.For(msg.Sender.ID).Set("Action","")
	} else if (bot.STMemory.For(msg.Sender.ID).Get("Action")=="hanghoa")|| (bot.STMemory.For(msg.Sender.ID).Get("Action")=="cargo_sodon") || (bot.STMemory.For(msg.Sender.ID).Get("Action")=="cargo_ten")|| (bot.STMemory.For(msg.Sender.ID).Get("Action")=="cargo_kh"){
		if (strings.HasPrefix(strings.ToUpper(strMsg), "SD ")){

			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Cargo.log"
			strParametername=" |SD: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)

			result, _ :=  soap.GetPolicyURNCargoInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmInputPolicyUrnCargoEvent
			}
			// Read xml
			byteValue := []byte(result.GetPolicyCargoInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmInputPolicyUrnCargoEvent
			}

			msgAll = ""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "cargo_sodon")

				return ConfirmWebViewBotEvent
			} else {
				msgText="Thông tin hồ sơ: "
				for i := 0; i < len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{
						check_SUMINSURED_AMT:=rowSets.Rows[i].SUMINSURED_AMT

						check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
						check_NgayDongPhi:=rowSets.Rows[i].NGAY_NOP_PHI

						if  rowSets.Rows[i].POLICY_URN != "" {
							msgText+=  "Số đơn " + rowSets.Rows[i].POLICY_URN
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

						if rowSets.Rows[i].VESSELORCONVEYANCE != "" {
							msgText+=  ". Tàu " + rowSets.Rows[i].VESSELORCONVEYANCE
						}

						if  rowSets.Rows[i].CERTIFICATENUMBER !=""{
							msgText+= `, GCN ` + rowSets.Rows[i].CERTIFICATENUMBER
						}
						msgText+=`, tham gia BH ` + rowSets.Rows[i].PRODUCT_NAME
						if (check_SUMINSURED_AMT!= "" && !strings.HasPrefix(strings.Trim(rowSets.Rows[i].SUMINSURED_AMT,""),"0 ")) {
							msgText+= " (" + rowSets.Rows[i].SUMINSURED_AMT + ")"
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

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}


						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
						//fmt.Print(msgText)
					}
					msgText=""
					msgAll=""


				}
			}


		}
		//case customer Name on Cargo added 28/8/2018 by duydp
		if (strings.HasPrefix(strings.ToUpper(strMsg), "TEN ")){

			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Cargo.log"
			strParametername=" |"+string(util.Personalize(T("getPolicyHolderName_title"), &msg.Sender))+":"

			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)

			result, _ :=  soap.GetCargoCustomerNameInfo(strKey)

			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmInputCargoCustomerNameEvent
			}

			// Read xml
			byteValue := []byte(result.GetPolicyCargoInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmInputCargoCustomerNameEvent
			}

			var msgText, msgAll  string
			msgAll = ""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)
				bot.STMemory.For(msg.Sender.ID).Set("svtype", "cargo_kh")

				return ConfirmWebViewBotEvent
			} else {
				msgText = "Thông tin hồ sơ: "
				for i := 0; i < len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{
						check_SUMINSURED_AMT:=rowSets.Rows[i].SUMINSURED_AMT

						check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
						check_NgayDongPhi:=rowSets.Rows[i].NGAY_NOP_PHI
						if rowSets.Rows[i].POLICYHOLDER_NAME != "" {
							msgText+=  "Khách hàng " + rowSets.Rows[i].POLICYHOLDER_NAME
						}

						if rowSets.Rows[i].VESSELORCONVEYANCE != "" {
							msgText+=  ", Tàu " + rowSets.Rows[i].VESSELORCONVEYANCE
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

						if  rowSets.Rows[i].PRODUCT_NAME !=""{
							msgText+= `, tham gia ` + rowSets.Rows[i].PRODUCT_NAME
						}
						if (check_SUMINSURED_AMT!= "" && !strings.HasPrefix(strings.Trim(rowSets.Rows[i].SUMINSURED_AMT,""),"0 ")) {
							msgText+= " (" + rowSets.Rows[i].SUMINSURED_AMT + ")"
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

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}


						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
						//fmt.Print(msgText)
					}
					msgText=""
					msgAll=""


				}
			}


		}
		//bot.STMemory.For(msg.Sender.ID).Set("Action","")
	} else if (bot.STMemory.For(msg.Sender.ID).Get("Action")=="bagdv")|| (bot.STMemory.For(msg.Sender.ID).Get("Action")=="bagdv_hdbh") || (bot.STMemory.For(msg.Sender.ID).Get("Action")=="bagdv_kh"){
		if (strings.HasPrefix(strings.ToUpper(strMsg), "HD ")){


			result, _ :=  soap.Get_BAGDV_ConstractNo_Info(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmInputContractNo_BAGDV_Event
			}
			// Read xml
			byteValue := []byte(result.GetContractInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmInputContractNo_BAGDV_Event
			}

			var InsuredNameVal = ""
			var InsuredNameStr = ""
			var InsuredNameStr1row = ""
			var InsuredNameNo = 0
			var vINCEPTION_DATE=""
			var vEXPIRY_DATE=""
			var vPOLICYURN=""
			var vBU=""
			var vSumInsured=""
			var vList=""
			vList="Danh sách người được bảo hiểm: "
			for j:=0;j<len(rowSets.Rows); j++ {
				if (vPOLICYURN != rowSets.Rows[j].CONTRACT_NUMBER && rowSets.Rows[j].CONTRACT_NUMBER != "") {
					vPOLICYURN=  "HĐBH: " + rowSets.Rows[j].CONTRACT_NUMBER
				}
				if (vBU != rowSets.Rows[j].BU_NAME && rowSets.Rows[j].BU_NAME != "") {
					vBU=  `, BH tại ` + rowSets.Rows[j].BU_NAME
				}
				if (vINCEPTION_DATE != rowSets.Rows[j].INCEPTION_DATE && rowSets.Rows[j].INCEPTION_DATE != "") {
					vINCEPTION_DATE=  `, từ ngày ` + rowSets.Rows[j].INCEPTION_DATE
				}
				if (vEXPIRY_DATE != rowSets.Rows[j].EXPIRED_DATE && rowSets.Rows[j].EXPIRED_DATE != "") {
					vEXPIRY_DATE=  ` đến ngày ` + rowSets.Rows[j].EXPIRED_DATE
				}
				if (rowSets.Rows[j].SUM_INSURED!= "" && !strings.HasPrefix(strings.Trim(rowSets.Rows[j].SUM_INSURED,""),"0 ")) {
					vSumInsured= ", STBH theo HĐ:" + rowSets.Rows[j].SUM_INSURED
				}

				if (InsuredNameVal != rowSets.Rows[j].INSURED_NAME && rowSets.Rows[j].INSURED_NAME != "") {
					if InsuredNameVal!="" {
						msgText += InsuredNameStr
					}

					InsuredNameNo +=1
					InsuredNameStr = "\r\n" + strconv.Itoa(InsuredNameNo)+"." + rowSets.Rows[j].INSURED_NAME+vSumInsured
					if InsuredNameNo==1 {
						InsuredNameStr1row =rowSets.Rows[j].INSURED_NAME+vSumInsured
					}

					InsuredNameVal=rowSets.Rows[j].INSURED_NAME

				}

			}

			if InsuredNameNo==1 {
				msgText = "Thông tin hồ sơ: "+"\n"+vPOLICYURN+vBU+vINCEPTION_DATE+vEXPIRY_DATE+"\n"+vList+InsuredNameStr1row
			} else {
				msgText = "Thông tin hồ sơ: "+"\n"+vPOLICYURN+vBU+vINCEPTION_DATE+vEXPIRY_DATE+"\n"+vList+ msgText + InsuredNameStr
			}

			bot.SendText(msg.Sender, msgText)

		}

		if (strings.HasPrefix(strings.ToUpper(strMsg), "BAGDV ") ||strings.HasPrefix(strings.ToUpper(strMsg), "BA ")||strings.HasPrefix(strings.ToUpper(strMsg), "GDV ")) {


			result, _ :=  soap.Get_BAGDV_CustomerName_Info(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmPolicyHolderNamen_BAGDV_Event
			}
			// Read xml
			byteValue := []byte(result.GetContractInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmPolicyHolderNamen_BAGDV_Event
			}
			msgAll = ""
			msgText="Thông tin hồ sơ:"
			for j:=0;j<len(rowSets.Rows); j++ {
				if len(rowSets.Rows)%1==0{
					msgText += "\n"
					if (strings.Contains(msgText, rowSets.Rows[j].INSURED_NAME) || rowSets.Rows[j].INSURED_NAME == ""){
					}else{
						msgText+="NĐBH " + rowSets.Rows[j].INSURED_NAME
					}

					if (strings.Contains(msgText, rowSets.Rows[j].CONTRACT_NUMBER) || rowSets.Rows[j].CONTRACT_NUMBER == ""){
					}else{
						msgText+= ", HĐBH " + rowSets.Rows[j].CONTRACT_NUMBER
					}

					if (strings.Contains(msgText, rowSets.Rows[j].BU_NAME) || rowSets.Rows[j].BU_NAME == ""){
					}else{
						msgText+=`, BH tại ` + rowSets.Rows[j].BU_NAME
					}

					if (strings.Contains(msgText, rowSets.Rows[j].INCEPTION_DATE) || rowSets.Rows[j].INCEPTION_DATE == ""){
					}else{
						msgText+=`, từ ngày ` + rowSets.Rows[j].INCEPTION_DATE
					}

					if (strings.Contains(msgText, rowSets.Rows[j].EXPIRED_DATE) || rowSets.Rows[j].EXPIRED_DATE == ""){
					}else{
						msgText+= ` đến ngày ` + rowSets.Rows[j].EXPIRED_DATE
					}

					msgText+="\n"+"Chi tiết bồi thường: "


					if (strings.HasPrefix(strings.Trim(rowSets.Rows[j].SUM_INSURED,""),"0") || rowSets.Rows[j].SUM_INSURED == ""){
					}else{
						msgText+= "\r\n"+"  - Số tiền BH theo HĐ: " +rowSets.Rows[j].SUM_INSURED
					}

					if (strings.HasPrefix(strings.Trim(rowSets.Rows[j].PAYMENT_MONEY,""),"0") || rowSets.Rows[j].PAYMENT_MONEY == ""){
					}else{
						msgText+= "\r\n"+"  - Số tiền đã được chi trả: " +rowSets.Rows[j].PAYMENT_MONEY
					}

					if (strings.HasPrefix(strings.Trim(rowSets.Rows[j].QUYEN_LOI_CON_LAI,""),"0") || rowSets.Rows[j].QUYEN_LOI_CON_LAI == ""){
					}else{
						msgText+= "\r\n"+"  - Quyền lợi còn lại: " +rowSets.Rows[j].QUYEN_LOI_CON_LAI
					}

					msgAll = msgAll + msgText

					bot.SendText(msg.Sender, msgAll)
				}
				msgText=""
				msgAll=""
			}

		}
		//bot.STMemory.For(msg.Sender.ID).Set("Action","")
	}else if (bot.STMemory.For(msg.Sender.ID).Get("Action")=="tauthuy")|| (bot.STMemory.For(msg.Sender.ID).Get("Action")=="tau_sodon") || (bot.STMemory.For(msg.Sender.ID).Get("Action")=="tau_ten"){
		if (strings.HasPrefix(strings.ToUpper(strMsg), "SD ")){
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Vessel.log"
			strParametername=" |SD: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
			result, _ :=  soap.GetPolicyURNVesselInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmInputPolicyUrnVesselEvent
			}
			// Read xml
			byteValue := []byte(result.GetPolicyVesselInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmInputPolicyUrnVesselEvent
			}

			msgAll = ""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "tau_sodon")

				return ConfirmWebViewBotEvent
			} else {
				msgText="Thông tin hồ sơ: "
				for i := 0; i < len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{
						check_SUMINSURED_AMT:=rowSets.Rows[i].SUMINSURED_AMT

						check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
						//check_NgayDongPhi:=rowSets.Rows[i].NGAY_NOP_PHI

						check_SoVuBT:=rowSets.Rows[i].SO_VU_BT
						check_TyLeBT:=rowSets.Rows[i].TY_LE_BT

						if rowSets.Rows[i].NAME_OF_VESSEL!=""{
							msgText+="\n" +`Tàu `+rowSets.Rows[i].NAME_OF_VESSEL
						}

						if rowSets.Rows[i].REGISTRATIONNO_IMO!=""{
							msgText+=`, số đăng ký/ IMO ` + rowSets.Rows[i].REGISTRATIONNO_IMO
						}

						msgText +=`, KH `+ rowSets.Rows[i].POLICYHOLDER_NAME
						msgText+=`, BH tại ` + rowSets.Rows[i].BU_NAME

						msgText+=`, từ ` + rowSets.Rows[i].INCEPTION_DATE

						msgText+=` đến ` + rowSets.Rows[i].EXPIRY_DATE

						msgText+=`. Số đơn ` + rowSets.Rows[i].POLICY_URN

						msgText +=`, tham gia BH `
						msgText+= rowSets.Rows[i].COV_CLASS_NAME
						if check_SUMINSURED_AMT !="0"{
							msgText+=" (" +check_SUMINSURED_AMT+ ")"
						}

						if check_NopPhi!="0"{
							msgText+=`, nộp phí ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
						}

						if rowSets.Rows[i].TINH_TRANG_THU_PHI!="" {
							if rowSets.Rows[i].TINH_TRANG_THU_PHI!="Đã nộp phí"{
								msgText+=`, ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)
							}
						}


						if (check_SoVuBT!="0") && (check_SoVuBT!="") {
							msgText += `. Số vụ tổn thất: ` + rowSets.Rows[i].SO_VU_BT + " vụ,"
						}
						if (check_TyLeBT!="0%") && (check_TyLeBT!="%") && (check_TyLeBT!=""){
							msgText += ` tỷ lệ BT: ` + rowSets.Rows[i].TY_LE_BT
						}

						if (rowSets.Rows[i].TRADE_NAME!= "") {
							msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
						}

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}
						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
					}
					msgText=""
					msgAll=""
				}
			}



		}


		//bot.STMemory.For(msg.Sender.ID).Set("Action","")
		//else if (bot.STMemory.For(msg.Sender.ID).Get("Action")=="")
	}else{

		//fmt.Print("Cuphapchung"+strMsg)
		//fmt.Print("giatri Action = " + bot.STMemory.For(msg.Sender.ID).Get("Action"))

		if (strings.HasPrefix(strings.ToUpper(strMsg), "CAR ")){
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Car.log"
			strParametername=" |CAR: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)

			result, _ :=  soap.GetregNumberCarInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}
			// Read xml
			byteValue := []byte(result.GetCarInfoResult)
			var rowSets config.RowSet
			// we unmarshal our byteArray which contains
			xml.Unmarshal(byteValue, &rowSets)

			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}

			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "biensoxe")

				return ConfirmWebViewBotEvent
			} else {
				msgText="Thông tin hồ sơ: "
				for i := 0; i < len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{

						check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
						//check_CertNum:=rowSets.Rows[i].CERT_NUMBER


						check_TNDSBB:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].TNDS_BAT_BUOC,""),"0tr")
						check_TNDSTN:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].TNDS_TU_NGUYEN,""),"0tr")
						check_sotienVCX:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].VAT_CHAT_XE,""),"0tr")
						check_LaiPhu:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].LAI_PHU,""),"0tr")
						check_HangHoa:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].HANG_HOA,""),"0tr")

						check_SoVuBT:=rowSets.Rows[i].SO_VU_BT
						check_TyLeBT:=rowSets.Rows[i].TY_LE_BT

						if(rowSets.Rows[i].REG_NUMBER!=""){
							msgText += "\n" + `BKS: ` + rowSets.Rows[i].REG_NUMBER
						}

						msgText+=`, chủ xe ` + rowSets.Rows[i].POLICYHOLDER_NAME

						if (rowSets.Rows[i].POLICYHOLDER_PHONE!="" && len(rowSets.Rows[i].POLICYHOLDER_PHONE)>1){
							var v_string string=rowSets.Rows[i].POLICYHOLDER_PHONE
							var rst_hP string
							v_reg1, _ := regexp.Compile("[^Z0-9]+")
							prc_str1:= v_reg1.ReplaceAllString(v_string, " ")
							rst_hP=strings.TrimSpace(prc_str1)
							rst_hP=strings.Replace(rst_hP,` `,`/`,-1)
							msgText+=" ("+rst_hP+")"
						}
						msgText+=`, BH tại ` + rowSets.Rows[i].BU_NAME +`-`+ rowSets.Rows[i].DEPT_NAME
						msgText+=`, từ ` + rowSets.Rows[i].INCEPTION_DATE

						msgText+=` đến ` + rowSets.Rows[i].EXPIRY_DATE

						msgText+=`. Số đơn ` +rowSets.Rows[i].POLICY_URN
						//if check_CertNum!=""{
						//	if check_CertNum!="0"{
						//		msgText+=`, GCN ` + rowSets.Rows[i].CERT_NUMBER
						//	}
						//}

						if check_TNDSBB !=true{
							msgText+=`, TNDSBB ` + "(" +rowSets.Rows[i].TNDS_BAT_BUOC+ ")"
						}
						if check_TNDSTN !=true{
							msgText+=`, TNDSTN ` + "(" +rowSets.Rows[i].TNDS_TU_NGUYEN+ ")"
						}
						//check_sotienVCX:=rowSets.Rows[i].VAT_CHAT_XE
						if check_sotienVCX!=true{
							msgText+=`, VCXE ` + "("+rowSets.Rows[i].VAT_CHAT_XE+")"}

						if check_LaiPhu!=true{
							msgText+=`, NTX ` + "("+rowSets.Rows[i].LAI_PHU+")"
						}

						if check_HangHoa!=true{
							msgText+=`, TNDS HH ` + "("+rowSets.Rows[i].HANG_HOA+")"
						}

						if check_NopPhi!="0"{
							msgText+=`, nộp phí ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
						}

						if rowSets.Rows[i].TINH_TRANG_THU_PHI!="" {
							if rowSets.Rows[i].TINH_TRANG_THU_PHI!="Đã nộp phí"{
								msgText+=`, ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)
							}
						}


						if check_SoVuBT!="0" {
							msgText += `. Số vụ tổn thất: ` + rowSets.Rows[i].SO_VU_BT + " vụ,"
						}
						if (check_TyLeBT!="0%") && (check_TyLeBT!="%") {
							msgText += ` tỷ lệ BT: ` + rowSets.Rows[i].TY_LE_BT
						}

						if (rowSets.Rows[i].TRADE_NAME!= "") {
							msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
						}

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}
						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
					}
					msgText=""
					msgAll=""
				}
			}



			//bot.STMemory.For(msg.Sender.ID).Set("Action","")

		}

		if (strings.HasPrefix(strings.ToUpper(strMsg), "GCN ")) {

			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Car.log"
			strParametername=" |GCN: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)

			result, _ :=  soap.GetcertNumberCarInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}

			// Read xml
			byteValue := []byte(result.GetCarInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				//+ rowSets.Rows[0].DATA
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "sogcnbh")

				return ConfirmWebViewBotEvent
			} else {
				msgText="Thông tin hồ sơ: "
				for i := 0; i < len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{

						check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
						check_CertNum:=rowSets.Rows[i].CERT_NUMBER


						check_TNDSBB:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].TNDS_BAT_BUOC,""),"0tr")
						check_TNDSTN:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].TNDS_TU_NGUYEN,""),"0tr")
						check_sotienVCX:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].VAT_CHAT_XE,""),"0tr")
						check_LaiPhu:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].LAI_PHU,""),"0tr")
						check_HangHoa:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].HANG_HOA,""),"0tr")

						check_SoVuBT:=rowSets.Rows[i].SO_VU_BT
						check_TyLeBT:=rowSets.Rows[i].TY_LE_BT

						if(rowSets.Rows[i].REG_NUMBER!=""){
							msgText += "\n" + `BKS: ` + rowSets.Rows[i].REG_NUMBER
						}

						msgText+=`, chủ xe ` + rowSets.Rows[i].POLICYHOLDER_NAME

						if (rowSets.Rows[i].POLICYHOLDER_PHONE!="" && len(rowSets.Rows[i].POLICYHOLDER_PHONE)>1){
							var v_string string=rowSets.Rows[i].POLICYHOLDER_PHONE
							var rst_hP string
							v_reg1, _ := regexp.Compile("[^Z0-9]+")
							prc_str1:= v_reg1.ReplaceAllString(v_string, " ")
							rst_hP=strings.TrimSpace(prc_str1)
							rst_hP=strings.Replace(rst_hP,` `,`/`,-1)
							msgText+=" ("+rst_hP+")"
						}
						msgText+=`, BH tại ` + rowSets.Rows[i].BU_NAME +` `+ rowSets.Rows[i].DEPT_NAME
						msgText+=`, từ ` + rowSets.Rows[i].INCEPTION_DATE

						msgText+=` đến ` + rowSets.Rows[i].EXPIRY_DATE

						msgText+=`. Số đơn ` +rowSets.Rows[i].POLICY_URN
						if check_CertNum!=""{
							if check_CertNum!="0"{
								msgText+=`, GCN ` + rowSets.Rows[i].CERT_NUMBER
							}
						}

						if check_TNDSBB !=true{
							msgText+=`, TNDSBB ` + "(" +rowSets.Rows[i].TNDS_BAT_BUOC+ ")"
						}
						if check_TNDSTN !=true{
							msgText+=`, TNDSTN ` + "(" +rowSets.Rows[i].TNDS_TU_NGUYEN+ ")"
						}
						//check_sotienVCX:=rowSets.Rows[i].VAT_CHAT_XE
						if check_sotienVCX!=true{
							msgText+=`, VCXE ` + "("+rowSets.Rows[i].VAT_CHAT_XE+")"}

						if check_LaiPhu!=true{
							msgText+=`, NTX ` + "("+rowSets.Rows[i].LAI_PHU+")"
						}

						if check_HangHoa!=true{
							msgText+=`, TNDS HH ` + "("+rowSets.Rows[i].HANG_HOA+")"
						}

						if check_NopPhi!="0"{
							msgText+=`, nộp phí ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
						}

						if rowSets.Rows[i].TINH_TRANG_THU_PHI!="" {
							if rowSets.Rows[i].TINH_TRANG_THU_PHI!="Đã nộp phí"{
								msgText+=`, ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)
							}
						}


						if check_SoVuBT!="0" {
							msgText += `. Số vụ tổn thất: ` + rowSets.Rows[i].SO_VU_BT + " vụ,"
						}
						if (check_TyLeBT!="0%") && (check_TyLeBT!="%") {
							msgText += ` tỷ lệ BT: ` + rowSets.Rows[i].TY_LE_BT
						}

						if (rowSets.Rows[i].TRADE_NAME!= "") {
							msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
						}

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}
						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
					}
					msgText=""
					msgAll=""
				}
			}


		}
		//Syntax :Chassis No (Motor)
		//Wrote:duydp at 23/4/2014
		if (strings.HasPrefix(strings.ToUpper(strMsg), "SK ")) {
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Car.log"
			strParametername=" |SK: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
			result, _ :=  soap.GetchassisNoCarInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}

			// Read xml
			byteValue := []byte(result.GetCarInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				//+ rowSets.Rows[0].DATA
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}

			msgAll = ""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "sok")

				return ConfirmWebViewBotEvent
			} else {
				msgText="Thông tin hồ sơ: "
				for i := 0; i < len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{

						check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
						//check_CertNum:=rowSets.Rows[i].CERT_NUMBER


						check_TNDSBB:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].TNDS_BAT_BUOC,""),"0tr")
						check_TNDSTN:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].TNDS_TU_NGUYEN,""),"0tr")
						check_sotienVCX:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].VAT_CHAT_XE,""),"0tr")
						check_LaiPhu:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].LAI_PHU,""),"0tr")
						check_HangHoa:=strings.HasPrefix(strings.Trim(rowSets.Rows[i].HANG_HOA,""),"0tr")

						check_SoVuBT:=rowSets.Rows[i].SO_VU_BT
						check_TyLeBT:=rowSets.Rows[i].TY_LE_BT

						if(rowSets.Rows[i].CHASSIS_NO!=""){
							msgText += "\n" + `BKS/SK: ` + rowSets.Rows[i].REG_NUMBER + "/" + rowSets.Rows[i].CHASSIS_NO
						}

						msgText+=`, chủ xe ` + rowSets.Rows[i].POLICYHOLDER_NAME

						if (rowSets.Rows[i].POLICYHOLDER_PHONE!="" && len(rowSets.Rows[i].POLICYHOLDER_PHONE)>1){
							var v_string string=rowSets.Rows[i].POLICYHOLDER_PHONE
							var rst_hP string
							v_reg1, _ := regexp.Compile("[^Z0-9]+")
							prc_str1:= v_reg1.ReplaceAllString(v_string, " ")
							rst_hP=strings.TrimSpace(prc_str1)
							rst_hP=strings.Replace(rst_hP,` `,`/`,-1)
							msgText+=" ("+rst_hP+")"
						}
						msgText+=`, BH tại ` + rowSets.Rows[i].BU_NAME +`-`+ rowSets.Rows[i].DEPT_NAME
						msgText+=`, từ ` + rowSets.Rows[i].INCEPTION_DATE

						msgText+=` đến ` + rowSets.Rows[i].EXPIRY_DATE

						msgText+=`. Số đơn ` +rowSets.Rows[i].POLICY_URN
						//if check_CertNum!=""{
						//	if check_CertNum!="0"{
						//		msgText+=`, GCN ` + rowSets.Rows[i].CERT_NUMBER
						//	}
						//}

						if check_TNDSBB !=true{
							msgText+=`, TNDSBB ` + "(" +rowSets.Rows[i].TNDS_BAT_BUOC+ ")"
						}
						if check_TNDSTN !=true{
							msgText+=`, TNDSTN ` + "(" +rowSets.Rows[i].TNDS_TU_NGUYEN+ ")"
						}
						//check_sotienVCX:=rowSets.Rows[i].VAT_CHAT_XE
						if check_sotienVCX!=true{
							msgText+=`, VCXE ` + "("+rowSets.Rows[i].VAT_CHAT_XE+")"}

						if check_LaiPhu!=true{
							msgText+=`, NTX ` + "("+rowSets.Rows[i].LAI_PHU+")"
						}

						if check_HangHoa!=true{
							msgText+=`, TNDS HH ` + "("+rowSets.Rows[i].HANG_HOA+")"
						}

						if check_NopPhi!="0"{
							msgText+=`, nộp phí ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
						}

						if rowSets.Rows[i].TINH_TRANG_THU_PHI!="" {
							if rowSets.Rows[i].TINH_TRANG_THU_PHI!="Đã nộp phí"{
								msgText+=`, ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)
							}
						}


						if check_SoVuBT!="0" {
							msgText += `. Số vụ tổn thất: ` + rowSets.Rows[i].SO_VU_BT + " vụ,"
						}
						if (check_TyLeBT!="0%") && (check_TyLeBT!="%") {
							msgText += ` tỷ lệ BT: ` + rowSets.Rows[i].TY_LE_BT
						}

						if (rowSets.Rows[i].TRADE_NAME!= "") {
							msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
						}

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}
						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
					}
					msgText=""
					msgAll=""
				}
			}


			//bot.STMemory.For(msg.Sender.ID).Set("Action","")
		}
		//
		if (strings.HasPrefix(strings.ToUpper(strMsg), "IMO ")){
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Vessel.log"
			strParametername=" |IMO: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
			result, _ :=  soap.GetregNumberVesselInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}
			// Read xml
			byteValue := []byte(result.GetPolicyVesselInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
				//ConfirmInputregNumberVesselEvent
			}

			msgAll = ""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "tau_imo")

				return ConfirmWebViewBotEvent
			} else {
				msgText="Thông tin hồ sơ: "
				for i := 0; i < len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{
						check_SUMINSURED_AMT:=rowSets.Rows[i].SUMINSURED_AMT

						check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
						//check_NgayDongPhi:=rowSets.Rows[i].NGAY_NOP_PHI

						check_SoVuBT:=rowSets.Rows[i].SO_VU_BT
						check_TyLeBT:=rowSets.Rows[i].TY_LE_BT

						if rowSets.Rows[i].NAME_OF_VESSEL!=""{
							msgText+="\n" +`Tàu `+rowSets.Rows[i].NAME_OF_VESSEL
						}

						if rowSets.Rows[i].REGISTRATIONNO_IMO!=""{
							msgText+=`, số đăng ký/ IMO ` + rowSets.Rows[i].REGISTRATIONNO_IMO
						}

						msgText +=`, KH `+ rowSets.Rows[i].POLICYHOLDER_NAME
						msgText+=`, BH tại ` + rowSets.Rows[i].BU_NAME

						msgText+=`, từ ` + rowSets.Rows[i].INCEPTION_DATE

						msgText+=` đến ` + rowSets.Rows[i].EXPIRY_DATE

						msgText+=`. Số đơn ` + rowSets.Rows[i].POLICY_URN

						msgText +=`, tham gia BH `
						msgText+= rowSets.Rows[i].COV_CLASS_NAME
						if check_SUMINSURED_AMT !="0"{
							msgText+=" (" +check_SUMINSURED_AMT+ ")"
						}

						if check_NopPhi!="0"{
							msgText+=`, nộp phí ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
						}

						if rowSets.Rows[i].TINH_TRANG_THU_PHI!="" {
							if rowSets.Rows[i].TINH_TRANG_THU_PHI!="Đã nộp phí"{
								msgText+=`, ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)
							}
						}


						if (check_SoVuBT!="0") && (check_SoVuBT!="") {
							msgText += `. Số vụ tổn thất: ` + rowSets.Rows[i].SO_VU_BT + " vụ,"
						}
						if (check_TyLeBT!="0%") && (check_TyLeBT!="%") && (check_TyLeBT!=""){
							msgText += ` tỷ lệ BT: ` + rowSets.Rows[i].TY_LE_BT
						}

						if (rowSets.Rows[i].TRADE_NAME!= "") {
							msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
						}

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}
						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
					}
					msgText=""
					msgAll=""
				}
			}



		}
		if (strings.HasPrefix(strings.ToUpper(strMsg), "HH ")){
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Cargo.log"
			strParametername=" |HH: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)

			result, _ :=  soap.GetCargoNameInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}
			// Read xml
			byteValue := []byte(result.GetPolicyCargoInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
				//ConfirmInputCargoNameEvent
			}

			msgAll = ""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "cargo_ten")

				return ConfirmWebViewBotEvent
			} else {
				msgText="Thông tin hồ sơ: "
				for i := 0; i < len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{
						check_SUMINSURED_AMT:=rowSets.Rows[i].SUMINSURED_AMT

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
						msgText+=`, tham gia ` + rowSets.Rows[i].PRODUCT_NAME
						if (check_SUMINSURED_AMT!= "" && !strings.HasPrefix(strings.Trim(rowSets.Rows[i].SUMINSURED_AMT,""),"0 ")) {
							msgText+= " (" + rowSets.Rows[i].SUMINSURED_AMT + ")"
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

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}


						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
						//fmt.Print(msgText)
					}
					msgText=""
					msgAll=""


				}
			}



		}
		if (strings.HasPrefix(strings.ToUpper(strMsg), "DK ")){
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Vessel.log"
			strParametername=" |DK: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
			result, _ :=  soap.GetregNumberVesselInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}
			// Read xml
			byteValue := []byte(result.GetPolicyVesselInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
				//ConfirmInputregNumberVesselEvent
			}

			msgAll = ""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "tau_imo")

				return ConfirmWebViewBotEvent
			} else {
				msgText="Thông tin hồ sơ: "
				for i := 0; i < len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{
						check_SUMINSURED_AMT:=rowSets.Rows[i].SUMINSURED_AMT

						check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
						//check_NgayDongPhi:=rowSets.Rows[i].NGAY_NOP_PHI

						check_SoVuBT:=rowSets.Rows[i].SO_VU_BT
						check_TyLeBT:=rowSets.Rows[i].TY_LE_BT

						if rowSets.Rows[i].NAME_OF_VESSEL!=""{
							msgText+="\n" +`Tàu `+rowSets.Rows[i].NAME_OF_VESSEL
						}

						if rowSets.Rows[i].REGISTRATIONNO_IMO!=""{
							msgText+=`, số đăng ký/ IMO ` + rowSets.Rows[i].REGISTRATIONNO_IMO
						}

						msgText +=`, KH `+ rowSets.Rows[i].POLICYHOLDER_NAME
						msgText+=`, BH tại ` + rowSets.Rows[i].BU_NAME

						msgText+=`, từ ` + rowSets.Rows[i].INCEPTION_DATE

						msgText+=` đến ` + rowSets.Rows[i].EXPIRY_DATE

						msgText+=`. Số đơn ` + rowSets.Rows[i].POLICY_URN

						msgText +=`, tham gia BH `
						msgText+= rowSets.Rows[i].COV_CLASS_NAME
						if check_SUMINSURED_AMT !="0"{
							msgText+=" (" +check_SUMINSURED_AMT+ ")"
						}

						if check_NopPhi!="0"{
							msgText+=`, nộp phí ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
						}

						if rowSets.Rows[i].TINH_TRANG_THU_PHI!="" {
							if rowSets.Rows[i].TINH_TRANG_THU_PHI!="Đã nộp phí"{
								msgText+=`, ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)
							}
						}


						if (check_SoVuBT!="0") && (check_SoVuBT!="") {
							msgText += `. Số vụ tổn thất: ` + rowSets.Rows[i].SO_VU_BT + " vụ,"
						}
						if (check_TyLeBT!="0%") && (check_TyLeBT!="%") && (check_TyLeBT!=""){
							msgText += ` tỷ lệ BT: ` + rowSets.Rows[i].TY_LE_BT
						}

						if (rowSets.Rows[i].TRADE_NAME!= "") {
							msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
						}

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}

						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
					}
					msgText=""
					msgAll=""
				}
			}



		}
		//Cu phap DD location
		if (strings.HasPrefix(strings.ToUpper(strMsg), "DD ")) {
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Fire.log"
			strParametername=" |DD: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
			result, _ :=  soap.GetLocationFireInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}
			// Read xml
			byteValue := []byte(result.GetPolicyFireInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
				//ConfirmLocationFireEvent
			}
			msgText=""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "ts_location")

				return ConfirmWebViewBotEvent
			} else {

				//process records less than 5
				msgText="Thông tin hồ sơ: "
				for j:=0;j<len(rowSets.Rows); j++ {
					if len(rowSets.Rows)%1==0{
						msgText += "\n"
						if (strings.Contains(msgText, rowSets.Rows[j].LOCATION) && rowSets.Rows[j].LOCATION == ""){
						}else{
							msgText+=rowSets.Rows[j].LOCATION
						}

						if (strings.Contains(msgText, rowSets.Rows[j].POLICYHOLDER_NAME) && rowSets.Rows[j].POLICYHOLDER_NAME == ""){
						}else{
							msgText+=", khách hàng " + rowSets.Rows[j].POLICYHOLDER_NAME
						}

						if (strings.Contains(msgText, rowSets.Rows[j].BU_NAME) && rowSets.Rows[j].BU_NAME == ""){
						}else{
							msgText+=`, BH tại ` + rowSets.Rows[j].BU_NAME
						}

						if (strings.Contains(msgText, rowSets.Rows[j].INCEPTION_DATE) && rowSets.Rows[j].INCEPTION_DATE == ""){
						}else{
							msgText+=`, từ ngày ` + rowSets.Rows[j].INCEPTION_DATE
						}

						if (strings.Contains(msgText, rowSets.Rows[j].EXPIRY_DATE) && rowSets.Rows[j].EXPIRY_DATE == ""){
						}else{
							msgText+= ` đến ngày ` + rowSets.Rows[j].EXPIRY_DATE
						}

						if (strings.Contains(msgText, rowSets.Rows[j].POLICY_URN) && rowSets.Rows[j].POLICY_URN == ""){
						}else{
							msgText+= ". Số đơn " + rowSets.Rows[j].POLICY_URN
						}

						if (strings.Contains(msgText, rowSets.Rows[j].COV_CLASS_NAME) && rowSets.Rows[j].COV_CLASS_NAME == ""){
						}else{
							msgText+= ", tham gia BH "+rowSets.Rows[j].COV_CLASS_NAME
						}
						if (strings.HasPrefix(strings.Trim(rowSets.Rows[j].SUMINSURED_AMT,""),"0") && rowSets.Rows[j].SUMINSURED_AMT == ""){
						}else{
							msgText+= " (" +rowSets.Rows[j].SUMINSURED_AMT + ")"
						}

						if (rowSets.Rows[j].TINH_TRANG_THU_PHI != "" && rowSets.Rows[j].TINH_TRANG_THU_PHI == "Đã nộp phí") {
						}else{
							msgText+= ", "+strings.ToLower(rowSets.Rows[j].TINH_TRANG_THU_PHI)
						}
						if (strings.Contains(msgText, rowSets.Rows[j].NGAY_NOP_PHI) && rowSets.Rows[j].NGAY_NOP_PHI == ""){
						}else{
							msgText+= `, nộp phí ngày `+rowSets.Rows[j].NGAY_NOP_PHI
						}

						if (strings.Contains(msgText, rowSets.Rows[j].SO_VU_BT) || rowSets.Rows[j].SO_VU_BT== "" || rowSets.Rows[j].SO_VU_BT== "0"){
						}else{

							msgText+= `. Số vụ tổn thất: ` + rowSets.Rows[j].SO_VU_BT + " vụ,"
						}

						if (strings.Contains(msgText, rowSets.Rows[j].TY_LE_BT) || rowSets.Rows[j].TY_LE_BT== "" || rowSets.Rows[j].TY_LE_BT== "0%"){
						}else{
							msgText+= ` tỷ lệ BT: ` + rowSets.Rows[j].TY_LE_BT
						}

						if (strings.Contains(msgText, rowSets.Rows[j].TRADE_NAME) && rowSets.Rows[j].TRADE_NAME != ""){
						}else{
							msgText+= `. LH CBKT ` +rowSets.Rows[j].TRADE_NAME
						}

						if (rowSets.Rows[j].TRADE_PHONE !="" && len(rowSets.Rows[j].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[j].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}
						msgText += "."
						msgAll = msgAll + msgText

						bot.SendText(msg.Sender, msgAll)
					}
					msgText=""
					msgAll=""
				}
			}

		}
		//add customer for Fire here
		//duydp added at 20180419
		if (strings.HasPrefix(strings.ToUpper(strMsg), "KH ")) {
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Fire.log"
			strParametername=" |KH: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
			result, _ :=  soap.GetPolicyHolderNameFireInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmPolicyHolderNamenFireEvent
			}
			// Read xml
			byteValue := []byte(result.GetPolicyFireInfoResult)

			var rowSetFire config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSetFire)
			if (rowSetFire.Rows[0].DATA == "NO DATA FOUND" || rowSetFire.Rows[0].DATA == "False Authentication" || rowSetFire.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmPolicyHolderNamenFireEvent
			}
			if len(rowSetFire.Rows) > 5 {
				iTotal:=len(rowSetFire.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "ts_tkh")

				return ConfirmWebViewBotEvent
			} else {
				msgText="Thông tin hồ sơ: "

				for j:=0;j<len(rowSetFire.Rows); j++ {
					if len(rowSetFire.Rows)%1==0{
						msgText += "\n"
						if (strings.Contains(msgText, rowSetFire.Rows[j].POLICYHOLDER_NAME) || rowSetFire.Rows[j].POLICYHOLDER_NAME == ""){
						}else{
							msgText+="Khách hàng " + rowSetFire.Rows[j].POLICYHOLDER_NAME
						}
						if (strings.Contains(msgText, rowSetFire.Rows[j].LOCATION) || rowSetFire.Rows[j].LOCATION == ""){
						}else{
							msgText+=", "+rowSetFire.Rows[j].LOCATION
						}

						if (strings.Contains(msgText, rowSetFire.Rows[j].BU_NAME) || rowSetFire.Rows[j].BU_NAME == ""){
						}else{
							msgText+=`, BH tại ` + rowSetFire.Rows[j].BU_NAME
						}

						if (strings.Contains(msgText, rowSetFire.Rows[j].INCEPTION_DATE) || rowSetFire.Rows[j].INCEPTION_DATE == ""){
						}else{
							msgText+=`, từ ngày ` + rowSetFire.Rows[j].INCEPTION_DATE
						}

						if (strings.Contains(msgText, rowSetFire.Rows[j].EXPIRY_DATE) || rowSetFire.Rows[j].EXPIRY_DATE == ""){
						}else{
							msgText+= ` đến ngày ` + rowSetFire.Rows[j].EXPIRY_DATE
						}

						if (strings.Contains(msgText, rowSetFire.Rows[j].POLICY_URN) || rowSetFire.Rows[j].POLICY_URN == ""){
						}else{
							msgText+= ". Số đơn " + rowSetFire.Rows[j].POLICY_URN
						}

						if (strings.Contains(msgText, rowSetFire.Rows[j].COV_CLASS_NAME) || rowSetFire.Rows[j].COV_CLASS_NAME == ""){
						}else{
							msgText+= ", tham gia BH "+rowSetFire.Rows[j].COV_CLASS_NAME
						}
						if (strings.HasPrefix(strings.Trim(rowSetFire.Rows[j].SUMINSURED_AMT,""),"0") || rowSetFire.Rows[j].SUMINSURED_AMT == ""){
						}else{
							msgText+= " (" +rowSetFire.Rows[j].SUMINSURED_AMT + ")"
						}

						if (rowSetFire.Rows[j].TINH_TRANG_THU_PHI != "" || rowSetFire.Rows[j].TINH_TRANG_THU_PHI == "Đã nộp phí") {
						}else{
							msgText+= ", "+strings.ToLower(rowSetFire.Rows[j].TINH_TRANG_THU_PHI)
						}
						if (strings.Contains(msgText, rowSetFire.Rows[j].NGAY_NOP_PHI) || rowSetFire.Rows[j].NGAY_NOP_PHI == ""){
						}else{
							msgText+= `, nộp phí ngày `+rowSetFire.Rows[j].NGAY_NOP_PHI
						}

						if (strings.Contains(msgText, rowSetFire.Rows[j].SO_VU_BT) || rowSetFire.Rows[j].SO_VU_BT== "" || rowSetFire.Rows[j].SO_VU_BT== "0"){
						}else{
							msgText+= `. Số vụ tổn thất: ` + rowSetFire.Rows[j].SO_VU_BT + " vụ,"
						}

						if (strings.Contains(msgText, rowSetFire.Rows[j].TY_LE_BT) || rowSetFire.Rows[j].TY_LE_BT== "" || rowSetFire.Rows[j].TY_LE_BT== "0%"){
						}else{
							msgText+= ` tỷ lệ BT: ` + rowSetFire.Rows[j].TY_LE_BT
						}

						if (strings.Contains(msgText, rowSetFire.Rows[j].TRADE_NAME) || rowSetFire.Rows[j].TRADE_NAME == ""){
						}else{
							msgText+= `. LH CBKT ` +rowSetFire.Rows[j].TRADE_NAME
						}

						if (rowSetFire.Rows[j].TRADE_PHONE !="" && len(rowSetFire.Rows[j].TRADE_PHONE)>1){
							var v_string_p string=rowSetFire.Rows[j].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}
						msgText += "."
						msgAll = msgAll + msgText

						bot.SendText(msg.Sender, msgAll)
					}
					msgText=""
					msgAll=""
				}
			}


		}
		//case customer Name on Cargo added 28/8/2018 by duydp
		if (strings.HasPrefix(strings.ToUpper(strMsg), "TEN ")){

			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Cargo.log"
			strParametername=" |"+string(util.Personalize(T("getPolicyHolderName_title"), &msg.Sender))+":"
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
			result, _ :=  soap.GetCargoCustomerNameInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmInputCargoCustomerNameEvent
			}

			// Read xml
			byteValue := []byte(result.GetPolicyCargoInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmInputCargoCustomerNameEvent
			}

			var msgText, msgAll  string
			msgAll = ""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)
				bot.STMemory.For(msg.Sender.ID).Set("svtype", "cargo_kh")

				return ConfirmWebViewBotEvent
			} else {
				msgText = "Thông tin hồ sơ: "
				for i := 0; i < len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{
						check_SUMINSURED_AMT:=rowSets.Rows[i].SUMINSURED_AMT

						check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
						check_NgayDongPhi:=rowSets.Rows[i].NGAY_NOP_PHI
						if rowSets.Rows[i].POLICYHOLDER_NAME != "" {
							msgText+=  "Khách hàng " + rowSets.Rows[i].POLICYHOLDER_NAME
						}

						if rowSets.Rows[i].VESSELORCONVEYANCE != "" {
							msgText+=  ", Tàu " + rowSets.Rows[i].VESSELORCONVEYANCE
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

						if  rowSets.Rows[i].PRODUCT_NAME !=""{
							msgText+= `, tham gia ` + rowSets.Rows[i].PRODUCT_NAME
						}
						if (check_SUMINSURED_AMT!= "" && !strings.HasPrefix(strings.Trim(rowSets.Rows[i].SUMINSURED_AMT,""),"0 ")) {
							msgText+= " (" + rowSets.Rows[i].SUMINSURED_AMT + ")"
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

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}


						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
						//fmt.Print(msgText)
					}
					msgText=""
					msgAll=""


				}
			}


		}

		// Cu phap TAU
		if (strings.HasPrefix(strings.ToUpper(strMsg), "TAU ")){
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Vessel.log"
			strParametername=" |TAU: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
			result, _ :=  soap.GetvesselNameInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}
			// Read xml
			byteValue := []byte(result.GetPolicyVesselInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
				//ConfirmInputVesselNameEvent
			}

			msgAll = ""

			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "tau_ten")

				return ConfirmWebViewBotEvent
			} else {
				msgText = `Thông tin hồ sơ: `
				for i := 0; i < len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{
						check_SUMINSURED_AMT:=rowSets.Rows[i].SUMINSURED_AMT

						check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
						//check_NgayDongPhi:=rowSets.Rows[i].NGAY_NOP_PHI
						check_SoVuBT:=rowSets.Rows[i].SO_VU_BT
						check_TyLeBT:=rowSets.Rows[i].TY_LE_BT

						if rowSets.Rows[i].NAME_OF_VESSEL!=""{
							msgText+="\n" +`Tàu `+rowSets.Rows[i].NAME_OF_VESSEL
						}

						if rowSets.Rows[i].REGISTRATIONNO_IMO!=""{
							msgText+=`, số đăng ký/ IMO ` + rowSets.Rows[i].REGISTRATIONNO_IMO
						}

						msgText +=`, KH `+ rowSets.Rows[i].POLICYHOLDER_NAME
						msgText+=`, BH tại ` + rowSets.Rows[i].BU_NAME

						msgText+=`, từ ` + rowSets.Rows[i].INCEPTION_DATE

						msgText+=` đến ` + rowSets.Rows[i].EXPIRY_DATE

						msgText+=`. Số đơn ` + rowSets.Rows[i].POLICY_URN

						msgText +=`, tham gia BH `
						msgText+= rowSets.Rows[i].COV_CLASS_NAME
						if check_SUMINSURED_AMT !="0"{
							msgText+=" (" +check_SUMINSURED_AMT+ ")"
						}

						if check_NopPhi!="0"{
							msgText+=`, nộp phí ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
						}

						if rowSets.Rows[i].TINH_TRANG_THU_PHI!="" {
							if rowSets.Rows[i].TINH_TRANG_THU_PHI!="Đã nộp phí"{
								msgText+=`, ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)
							}
						}


						if (check_SoVuBT!="0") && (check_SoVuBT!="") {
							msgText += `. Số vụ tổn thất: ` + rowSets.Rows[i].SO_VU_BT + " vụ,"
						}
						if (check_TyLeBT!="0%") && (check_TyLeBT!="%") && (check_TyLeBT!=""){
							msgText += ` tỷ lệ BT: ` + rowSets.Rows[i].TY_LE_BT
						}

						if (rowSets.Rows[i].TRADE_NAME!= "") {
							msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
						}

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}

						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
					}
					msgText=""
					msgAll=""

				}
			}


		}
		//syntax HDBH-BAGDV
		//Ten khach hang: ba,gdv,bagdv: 10/5/2018
		if (strings.HasPrefix(strings.ToUpper(strMsg), "BAGDV ")||strings.HasPrefix(strings.ToUpper(strMsg), "BA ")||strings.HasPrefix(strings.ToUpper(strMsg), "GDV ")){

			result, _ :=  soap.Get_BAGDV_CustomerName_Info(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmPolicyHolderNamen_BAGDV_Event
			}
			// Read xml
			byteValue := []byte(result.GetContractInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmPolicyHolderNamen_BAGDV_Event
			}
			msgAll = ""
			msgText="Thông tin hồ sơ: "
			for j:=0;j<len(rowSets.Rows); j++ {
				if len(rowSets.Rows)%1==0{
					msgText += "\n"
					if (strings.Contains(msgText, rowSets.Rows[j].INSURED_NAME) || rowSets.Rows[j].INSURED_NAME == ""){
					}else{
						msgText+="NĐBH " + rowSets.Rows[j].INSURED_NAME
					}

					if (strings.Contains(msgText, rowSets.Rows[j].CONTRACT_NUMBER) || rowSets.Rows[j].CONTRACT_NUMBER == ""){
					}else{
						msgText+= ", HĐBH " + rowSets.Rows[j].CONTRACT_NUMBER
					}

					if (strings.Contains(msgText, rowSets.Rows[j].BU_NAME) || rowSets.Rows[j].BU_NAME == ""){
					}else{
						msgText+=`, BH tại ` + rowSets.Rows[j].BU_NAME
					}

					if (strings.Contains(msgText, rowSets.Rows[j].INCEPTION_DATE) || rowSets.Rows[j].INCEPTION_DATE == ""){
					}else{
						msgText+=`, từ ngày ` + rowSets.Rows[j].INCEPTION_DATE
					}

					if (strings.Contains(msgText, rowSets.Rows[j].EXPIRED_DATE) || rowSets.Rows[j].EXPIRED_DATE == ""){
					}else{
						msgText+= ` đến ngày ` + rowSets.Rows[j].EXPIRED_DATE
					}

					msgText+="\n"+"Chi tiết bồi thường: "


					if (strings.HasPrefix(strings.Trim(rowSets.Rows[j].SUM_INSURED,""),"0") || rowSets.Rows[j].SUM_INSURED == ""){
					}else{
						msgText+= "\r\n"+"  - Số tiền BH theo HĐ: " +rowSets.Rows[j].SUM_INSURED
					}

					if (strings.HasPrefix(strings.Trim(rowSets.Rows[j].PAYMENT_MONEY,""),"0") || rowSets.Rows[j].PAYMENT_MONEY == ""){
					}else{
						msgText+= "\r\n"+"  - Số tiền đã được chi trả: " +rowSets.Rows[j].PAYMENT_MONEY
					}

					if (strings.HasPrefix(strings.Trim(rowSets.Rows[j].QUYEN_LOI_CON_LAI,""),"0") || rowSets.Rows[j].QUYEN_LOI_CON_LAI == ""){
					}else{
						msgText+= "\r\n"+"  - Quyền lợi còn lại: " +rowSets.Rows[j].QUYEN_LOI_CON_LAI
					}

					msgAll = msgAll + msgText

					bot.SendText(msg.Sender, msgAll)
				}
				msgText=""
				msgAll=""
			}
		}
		// So HopDong BH: 9/5/2018 duydp
		if (strings.HasPrefix(strings.ToUpper(strMsg), "HD ")){


			result, _ :=  soap.Get_BAGDV_ConstractNo_Info(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmInputContractNo_BAGDV_Event
			}
			// Read xml
			byteValue := []byte(result.GetContractInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmInputContractNo_BAGDV_Event
			}

			var InsuredNameVal = ""
			var InsuredNameStr = ""
			var InsuredNameStr1row = ""
			var InsuredNameNo = 0
			var vINCEPTION_DATE=""
			var vEXPIRY_DATE=""
			var vPOLICYURN=""
			var vBU=""
			var vSumInsured=""
			var vList=""
			vList="Danh sách người được bảo hiểm: "
			for j:=0;j<len(rowSets.Rows); j++ {
				if (vPOLICYURN != rowSets.Rows[j].CONTRACT_NUMBER && rowSets.Rows[j].CONTRACT_NUMBER != "") {
					vPOLICYURN=  "HĐBH: " + rowSets.Rows[j].CONTRACT_NUMBER
				}
				if (vBU != rowSets.Rows[j].BU_NAME && rowSets.Rows[j].BU_NAME != "") {
					vBU=  `, BH tại ` + rowSets.Rows[j].BU_NAME
				}
				if (vINCEPTION_DATE != rowSets.Rows[j].INCEPTION_DATE && rowSets.Rows[j].INCEPTION_DATE != "") {
					vINCEPTION_DATE=  `, từ ngày ` + rowSets.Rows[j].INCEPTION_DATE
				}
				if (vEXPIRY_DATE != rowSets.Rows[j].EXPIRED_DATE && rowSets.Rows[j].EXPIRED_DATE != "") {
					vEXPIRY_DATE=  ` đến ngày ` + rowSets.Rows[j].EXPIRED_DATE
				}
				if (rowSets.Rows[j].SUM_INSURED!= "" && !strings.HasPrefix(strings.Trim(rowSets.Rows[j].SUM_INSURED,""),"0 ")) {
					vSumInsured= ", STBH theo HĐ:" + rowSets.Rows[j].SUM_INSURED
				}

				if (InsuredNameVal != rowSets.Rows[j].INSURED_NAME && rowSets.Rows[j].INSURED_NAME != "") {
					if InsuredNameVal!="" {
						msgText += InsuredNameStr
					}

					InsuredNameNo +=1
					InsuredNameStr = "\r\n" + strconv.Itoa(InsuredNameNo)+"." + rowSets.Rows[j].INSURED_NAME+vSumInsured
					if InsuredNameNo==1 {
						InsuredNameStr1row =rowSets.Rows[j].INSURED_NAME+vSumInsured
					}

					InsuredNameVal=rowSets.Rows[j].INSURED_NAME

				}

			}

			if InsuredNameNo==1 {
				msgText = "Thông tin hồ sơ: "+"\n"+vPOLICYURN+vBU+vINCEPTION_DATE+vEXPIRY_DATE+"\n"+vList+InsuredNameStr1row
			} else {
				msgText = "Thông tin hồ sơ: "+"\n"+vPOLICYURN+vBU+vINCEPTION_DATE+vEXPIRY_DATE+"\n"+vList+ msgText + InsuredNameStr
			}

			bot.SendText(msg.Sender, msgText)
		}
		//common CMT & MST
		//case 1: CMT
		if (strings.HasPrefix(strings.ToUpper(strMsg), "CMT ")) {
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="PolicyInfo.log"
			strParametername=" |CMT: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
			result, _ := soap.GetCommonPolicyInfo(1, strKey ,"")
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}

			// Read xml
			byteValue := []byte(result.GetPolicyInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found")+"", &msg.Sender))
				return ConfirmSyntaxEvent
			}

			msgAll= ""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "CMT")

				return ConfirmWebViewBotEvent
			} else {
				msgText= `Thông tin hồ sơ: `
				for i := 0; i <len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{


						//check_NopPhi := rowSets.Rows[i].PREMIUM_AMT
						check_NgayDongPhi := rowSets.Rows[i].NGAY_NOP_PHI
						msgText += "\n" + "Số đơn " + rowSets.Rows[i].POLICY_URN + ", khách hàng " + rowSets.Rows[i].POLICYHOLDER_NAME


						msgText += `, tham gia BH `

						msgText += rowSets.Rows[i].PRODUCT_NAME
						msgText += `, nhóm sản phẩm ` + rowSets.Rows[i].BUSINESS_LINE
						msgText += `, thời hạn BH từ ` + rowSets.Rows[i].INCEPTION_DATE
						msgText += ` đến ` + rowSets.Rows[i].EXPIRY_DATE
						msgText += `, ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)

						if check_NgayDongPhi != "" {
							msgText += ` ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
						}
						msgText += `, tại ` + rowSets.Rows[i].BU_NAME
						if (rowSets.Rows[i].TRADE_NAME!= "") {
							msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
						}

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}
						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
					}
					msgAll=""
					msgText=""
				}
			}


		}
		//
		//case 2: MST
		if (strings.HasPrefix(strings.ToUpper(strMsg), "MST ")) {
			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="PolicyInfo.log"
			strParametername=" |MST: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
			result, _ := soap.GetCommonPolicyInfo(2,"" , strKey)

			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}

			// Read xml
			byteValue := []byte(result.GetPolicyInfoResult)

			var rowSets config.RowSet
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'users' which we defined above
			xml.Unmarshal(byteValue, &rowSets)
			if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found")+"", &msg.Sender))
				return ConfirmSyntaxEvent
			}

			msgAll= ""
			if len(rowSets.Rows) > 5 {
				iTotal:=len(rowSets.Rows)
				var strCount=strconv.Itoa(iTotal)+" kết quả"
				bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
				bot.STMemory.For(msg.Sender.ID).Set("skey", strKey)

				bot.STMemory.For(msg.Sender.ID).Set("svtype", "MST")

				return ConfirmWebViewBotEvent
			} else {
				msgText= `Thông tin hồ sơ: `
				for i := 0; i <len(rowSets.Rows); i++ {
					if len(rowSets.Rows)%1==0{


						//check_NopPhi := rowSets.Rows[i].PREMIUM_AMT
						check_NgayDongPhi := rowSets.Rows[i].NGAY_NOP_PHI
						msgText += "\n" + "Số đơn " + rowSets.Rows[i].POLICY_URN + ", khách hàng " + rowSets.Rows[i].POLICYHOLDER_NAME


						msgText += `, tham gia `

						msgText += rowSets.Rows[i].PRODUCT_NAME
						msgText += `, nhóm sản phẩm ` + rowSets.Rows[i].BUSINESS_LINE
						msgText += `, thời hạn BH từ ` + rowSets.Rows[i].INCEPTION_DATE
						msgText += ` đến ` + rowSets.Rows[i].EXPIRY_DATE
						msgText += `, ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)

						if check_NgayDongPhi != "" {
							msgText += ` ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
						}
						msgText += `, tại ` + rowSets.Rows[i].BU_NAME
						if (rowSets.Rows[i].TRADE_NAME!= "") {
							msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
						}

						if (rowSets.Rows[i].TRADE_PHONE !="" && len(rowSets.Rows[i].TRADE_PHONE)>1){
							var v_string_p string=rowSets.Rows[i].TRADE_PHONE
							var rst string
							v_reg, _ := regexp.Compile("[^Z0-9]+")
							prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
							rst=strings.TrimSpace(prc_str)
							rst=strings.Replace(rst,` `,`/`,-1)
							msgText+=" ("+rst+")"
						}
						msgText += "."
						msgAll = msgAll + msgText
						bot.SendText(msg.Sender, msgAll)
					}
					msgAll=""
					msgText=""
				}
			}


		}
		// CASE cu phap SD chung
		if (strings.HasPrefix(strings.ToUpper(strMsg), "SD ")) {

			var strFileName,strParametername,strIDSender string
			strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
			strFileName="Fire.log"
			strParametername=" |SD: "
			soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)

			var rowTS config.RowSet
			var rowHH config.RowSet
			var rowTau config.RowSet
			// case 1 SD Tai san
			result, _ :=  soap.GetPolicyURNFireInfo(strKey)
			if (result==nil) {
				bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
				return ConfirmSyntaxEvent
			}
			byteValue := []byte(result.GetPolicyFireInfoResult)
			xml.Unmarshal(byteValue, &rowTS)

			if (rowTS.Rows[0].DATA == "NO DATA FOUND" || rowTS.Rows[0].DATA == "False Authentication" || rowTS.Rows[0].ERROR != "") {
				var strFileName,strParametername,strIDSender string
				strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
				strFileName="Cargo.log"
				strParametername=" |SD: "
				soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)

				// Case 2 SD Hang hoa
				result, _ :=  soap.GetPolicyURNCargoInfo(strKey)
				if (result==nil) {
					bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
					return ConfirmSyntaxEvent
				}
				byteValue := []byte(result.GetPolicyCargoInfoResult)
				xml.Unmarshal(byteValue, &rowHH)

				if (rowHH.Rows[0].DATA == "NO DATA FOUND" || rowHH.Rows[0].DATA == "False Authentication" || rowHH.Rows[0].ERROR != "") {
					var strFileName,strParametername,strIDSender string
					strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
					strFileName="Vessel.log"
					strParametername=" |SD: "
					soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,strKey)
					// Case 3 SD Tau thuy
					result, _ :=  soap.GetPolicyURNVesselInfo(strKey)
					if (result==nil) {
						bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
						return ConfirmSyntaxEvent
					}

					//var rowSet1 config.RowSet
					byteValue := []byte(result.GetPolicyVesselInfoResult)
					xml.Unmarshal(byteValue, &rowTau)
                    //fmt.Print("KQ:" + rowTau.Rows[0].DATA)

					if (rowTau.Rows[0].DATA == "NO DATA FOUND" || rowTau.Rows[0].DATA == "False Authentication" || rowTau.Rows[0].ERROR != "") {
						//fmt.Print("In")
						bot.SendText(msg.Sender, util.Personalize(T("NoData_Found")+"", &msg.Sender))
						return ConfirmSyntaxEvent
					}else{
						//fmt.Print("Ketthuc")
						msgAll = ""
						//msgText = `Thông tin hồ sơ:`

						for i := 0; i < len(rowTau.Rows); i++ {
							if len(rowTau.Rows)%1==0{
								check_SUMINSURED_AMT:=rowTau.Rows[i].SUMINSURED_AMT

								check_NopPhi:=rowTau.Rows[i].PREMIUM_PAYMENT_AMT
								//check_NgayDongPhi:=rowTau.Rows[i].NGAY_NOP_PHI

								check_SoVuBT:=rowTau.Rows[i].SO_VU_BT
								check_TyLeBT:=rowTau.Rows[i].TY_LE_BT

								if rowTau.Rows[i].NAME_OF_VESSEL!=""{
									msgText+="\n" +`Tàu `+rowTau.Rows[i].NAME_OF_VESSEL
								}

								if rowTau.Rows[i].REGISTRATIONNO_IMO!=""{
									msgText+=`, số đăng ký/ IMO ` + rowTau.Rows[i].REGISTRATIONNO_IMO
								}

								msgText +=`, KH `+ rowTau.Rows[i].POLICYHOLDER_NAME
								msgText+=`, BH tại ` + rowTau.Rows[i].BU_NAME

								msgText+=`, từ ` + rowTau.Rows[i].INCEPTION_DATE

								msgText+=` đến ` + rowTau.Rows[i].EXPIRY_DATE

								msgText+=`. Số đơn ` + rowTau.Rows[i].POLICY_URN

								msgText +=`, tham gia BH `
								msgText+= rowTau.Rows[i].PRODUCT_NAME
								if check_SUMINSURED_AMT !="0"{

									msgText+=" (" +check_SUMINSURED_AMT+ ")"
								}

								if check_NopPhi!="0"{
									msgText+=`, nộp phí ngày ` + rowTau.Rows[i].NGAY_NOP_PHI
								}

								if rowTau.Rows[i].TINH_TRANG_THU_PHI!="" {
									if rowTau.Rows[i].TINH_TRANG_THU_PHI!="Đã nộp phí"{
										msgText+=`, ` + strings.ToLower(rowTau.Rows[i].TINH_TRANG_THU_PHI)
									}
								}


								if (check_SoVuBT!="0") && (check_SoVuBT!="") {
									msgText += `. Số vụ tổn thất: ` + rowTau.Rows[i].SO_VU_BT + " vụ,"
								}
								if (check_TyLeBT!="0%") && (check_TyLeBT!="%") && (check_TyLeBT!=""){
									msgText += ` tỷ lệ BT: ` + rowTau.Rows[i].TY_LE_BT
								}

								if (rowTau.Rows[i].TRADE_NAME!= "") {
									msgText+=`. LH CBKT ` +rowTau.Rows[i].TRADE_NAME
								}

								if (rowTau.Rows[i].TRADE_PHONE !="" && len(rowTau.Rows[i].TRADE_PHONE)>1){
									var v_string_p string=rowTau.Rows[i].TRADE_PHONE
									var rst string
									v_reg, _ := regexp.Compile("[^Z0-9]+")
									prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
									rst=strings.TrimSpace(prc_str)
									rst=strings.Replace(rst,` `,`/`,-1)
									msgText+=" ("+rst+")"
								}
								msgText += "."
								msgAll = msgAll + msgText
								bot.SendText(msg.Sender, msgAll)
							}
							msgText=""
							msgAll=""
						}

					}

				}else{

					msgAll = ""
					//msgText = `Thông tin hồ sơ:`

					for i := 0; i < len(rowHH.Rows); i++ {
						if len(rowHH.Rows)%1==0{
							check_SUMINSURED_AMT:=rowHH.Rows[i].SUMINSURED_AMT

							check_NopPhi:=rowHH.Rows[i].PREMIUM_PAYMENT_AMT
							check_NgayDongPhi:=rowHH.Rows[i].NGAY_NOP_PHI

							if  rowHH.Rows[i].POLICY_URN != "" {
								msgText+=  "Số đơn " + rowHH.Rows[i].POLICY_URN
							}



							if rowHH.Rows[i].POLICYHOLDER_NAME != "" {
								msgText+=  ", khách hàng " + rowHH.Rows[i].POLICYHOLDER_NAME
							}
							if rowHH.Rows[i].BU_NAME != "" {
								msgText+=  `, BH tại ` + rowHH.Rows[i].BU_NAME
							}


							if rowHH.Rows[i].INCEPTION_DATE != "" {
								msgText+=  `, từ ` + rowHH.Rows[i].INCEPTION_DATE
							}

							if rowHH.Rows[i].EXPIRY_DATE != "" {
								msgText+=  ` đến ` + rowHH.Rows[i].EXPIRY_DATE
							}

							if rowHH.Rows[i].VESSELORCONVEYANCE != "" {
								msgText+=  ". Tàu " + rowHH.Rows[i].VESSELORCONVEYANCE
							}

							if  rowHH.Rows[i].CERTIFICATENUMBER !=""{
								msgText+= `, GCN ` + rowHH.Rows[i].CERTIFICATENUMBER
							}
							if  rowHH.Rows[i].PRODUCT_NAME !=""{
								msgText+=`, tham gia ` + rowHH.Rows[i].PRODUCT_NAME
							}

							if (check_SUMINSURED_AMT!= "" && !strings.HasPrefix(strings.Trim(rowHH.Rows[i].SUMINSURED_AMT,""),"0 ")) {
								msgText+= " (" + rowHH.Rows[i].SUMINSURED_AMT + ")"
							}


							if  rowHH.Rows[i].PLACEOFDEPARTURES != "" {
								msgText+=  `, đi từ ` + rowHH.Rows[i].PLACEOFDEPARTURES
							}

							if rowHH.Rows[i].FINALDESTINATIONS != "" {
								msgText+=  ` đến ` + rowHH.Rows[i].FINALDESTINATIONS
							}

							if (check_NopPhi!="0" && check_NgayDongPhi!= "") {
								msgText+=`, nộp phí ngày`+rowHH.Rows[i].NGAY_NOP_PHI
							}

							if rowHH.Rows[i].TINH_TRANG_THU_PHI != "" {
								if rowHH.Rows[i].TINH_TRANG_THU_PHI != "Đã nộp phí" {
									msgText+= ` ` + strings.ToLower(rowHH.Rows[i].TINH_TRANG_THU_PHI)
								}
							}

							if (rowHH.Rows[i].TRADE_NAME!= "") {
								msgText+=`. LH CBKT ` +rowHH.Rows[i].TRADE_NAME
							}

							if (rowHH.Rows[i].TRADE_PHONE !="" && len(rowHH.Rows[i].TRADE_PHONE)>1){
								var v_string_p string=rowHH.Rows[i].TRADE_PHONE
								var rst string
								v_reg, _ := regexp.Compile("[^Z0-9]+")
								prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
								rst=strings.TrimSpace(prc_str)
								rst=strings.Replace(rst,` `,`/`,-1)
								msgText+=" ("+rst+")"
							}

							msgText += "."
							msgAll = msgAll + msgText
							bot.SendText(msg.Sender, msgAll)
							//fmt.Print(msgText)
						}
						msgText=""
						msgAll=""


					}


				}

			}else{

				var locationVal = ""
				var locationStr = ""
				var locationStr1row = ""
				var locationNo = 0
				var startCvr=0
				var vTradeName=""
				var vNgayNopPhi=""
				var vTinhTrangThuPhi=""
				var vPhoneTrade=""
				var vSoVuBT=""
				var vTyLeBT=""
				var vINCEPTION_DATE=""
				var vEXPIRY_DATE=""
				var vPOLICYURN=""
				var vPolicyHolderName=""
				var vBU=""
				for j:=0;j<len(rowTS.Rows); j++ {

					if (vPolicyHolderName != rowTS.Rows[j].POLICYHOLDER_NAME && rowTS.Rows[j].POLICYHOLDER_NAME != "") {
						vPolicyHolderName=  ", khách hàng " + rowTS.Rows[j].POLICYHOLDER_NAME
					}
					if (vBU != rowTS.Rows[j].BU_NAME && rowTS.Rows[j].BU_NAME != "") {
						vBU=  `, BH tại ` + rowTS.Rows[j].BU_NAME
					}
					if (vTradeName!= rowTS.Rows[j].TRADE_NAME && rowTS.Rows[j].TRADE_NAME!= "") {
						vTradeName=`. LH CBKT ` +rowTS.Rows[j].TRADE_NAME
					}

					if (vINCEPTION_DATE != rowTS.Rows[j].INCEPTION_DATE && rowTS.Rows[j].INCEPTION_DATE != "") {
						vINCEPTION_DATE=  `, từ ngày ` + rowTS.Rows[j].INCEPTION_DATE
					}
					if (vEXPIRY_DATE != rowTS.Rows[j].INCEPTION_DATE && rowTS.Rows[j].EXPIRY_DATE != "") {
						vEXPIRY_DATE=  ` đến ngày ` + rowTS.Rows[j].EXPIRY_DATE
					}
					if (vPOLICYURN != rowTS.Rows[j].POLICY_URN && rowTS.Rows[j].POLICY_URN != "") {
						vPOLICYURN=  "Số đơn " + rowTS.Rows[j].POLICY_URN
					}
					if (vNgayNopPhi!= rowTS.Rows[j].NGAY_NOP_PHI && rowTS.Rows[j].NGAY_NOP_PHI!= "") {
						vNgayNopPhi=`, nộp phí ngày `+rowTS.Rows[j].NGAY_NOP_PHI
					}
					if (rowTS.Rows[j].TINH_TRANG_THU_PHI != "" && rowTS.Rows[j].TINH_TRANG_THU_PHI == "Đã nộp phí" ) {
					}else {
						vTinhTrangThuPhi = ", "+strings.ToLower(rowTS.Rows[j].TINH_TRANG_THU_PHI)
					}

					if (rowTS.Rows[j].TRADE_PHONE !="" && len(rowTS.Rows[j].TRADE_PHONE)>1){
						var v_string_p string=rowTS.Rows[j].TRADE_PHONE
						var rst string
						v_reg, _ := regexp.Compile("[^Z0-9]+")
						prc_str:= v_reg.ReplaceAllString(v_string_p, " ")
						rst=strings.TrimSpace(prc_str)
						rst=strings.Replace(rst,` `,`/`,-1)
						vPhoneTrade=" ("+rst+")"
					}



					if (vSoVuBT!= rowTS.Rows[j].SO_VU_BT && rowTS.Rows[j].SO_VU_BT!= ""&& rowTS.Rows[j].SO_VU_BT!= "0") {
						vSoVuBT=`. Số vụ tổn thất: ` + rowTS.Rows[j].SO_VU_BT + " vụ,"
					}

					if (vTyLeBT!= rowTS.Rows[j].TY_LE_BT && rowTS.Rows[j].TY_LE_BT!= "" && rowTS.Rows[j].TY_LE_BT!= "0%") {
						vTyLeBT=` tỷ lệ BT: ` + rowTS.Rows[j].TY_LE_BT
					}

					if (locationVal != rowTS.Rows[j].LOCATION && rowTS.Rows[j].LOCATION != "") {
						if locationVal!="" {
							msgText += locationStr
						}
						locationNo +=1
						startCvr = 1
						locationStr = "\r\n" + strconv.Itoa(locationNo)+"." + rowTS.Rows[j].LOCATION
						if locationNo==1 {
							locationStr1row ="Số đơn " + rowTS.Rows[j].POLICY_URN+", " +rowTS.Rows[j].LOCATION + ", khách hàng " + rowTS.Rows[j].POLICYHOLDER_NAME + `, BH tại ` + rowTS.Rows[j].BU_NAME
							locationStr1row += `, từ ngày ` + rowTS.Rows[j].INCEPTION_DATE + ` đến ngày ` + rowTS.Rows[j].EXPIRY_DATE
							locationStr1row +=  ", tham gia BH "
						}

						locationVal=rowTS.Rows[j].LOCATION

					}

					if startCvr==0 {
						locationStr +=", "
						locationStr1row += ", "
					}
					locationStr += " tham gia BH " + rowTS.Rows[j].COV_CLASS_NAME
					locationStr1row	+= rowTS.Rows[j].COV_CLASS_NAME
					if (rowTS.Rows[j].SUMINSURED_AMT!= "" && !strings.HasPrefix(strings.Trim(rowTS.Rows[j].SUMINSURED_AMT,""),"0 ")) {
						locationStr +=  " (" + rowTS.Rows[j].SUMINSURED_AMT + ")"
						locationStr1row	+=  " (" + rowTS.Rows[j].SUMINSURED_AMT + ")"

					}
					startCvr=0

				}

				if locationNo==1 {
					msgText = "Thông tin hồ sơ: "+"\n"+locationStr1row+vTinhTrangThuPhi+vNgayNopPhi+vSoVuBT+vTyLeBT+vTradeName+vPhoneTrade
				} else {
					msgText = "Thông tin hồ sơ: "+"\n"+vPOLICYURN+vPolicyHolderName +vBU+vINCEPTION_DATE+vEXPIRY_DATE+vTinhTrangThuPhi+vNgayNopPhi+vSoVuBT+vTyLeBT+vTradeName+vPhoneTrade+" .Địa điểm bảo hiểm:"+ msgText + locationStr
				}

				bot.SendText(msg.Sender, msgText+".")

			}

		}

	}
	bot.STMemory.For(msg.Sender.ID).Set("Action","")
	return ConfirmSyntaxEvent

}
