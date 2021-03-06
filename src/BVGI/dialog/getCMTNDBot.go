//author: duydp
//wrote at: 20/3/2018
package dialog
import (
	"util"
	"github.com/michlabs/fbbot"
	"BVGI/config"
	//"strings"
	"soap"
	"encoding/xml"

	"strings"
	"regexp"
)
type GetCMTNDBot struct {
	fbbot.BaseStep
}
func (s GetCMTNDBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, util.Personalize(T("getCMTND_title"), &msg.Sender))
	return StayEvent
}
func (s GetCMTNDBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation") == "" {
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "true")
	}
	var strFileName,strParametername,strIDSender string
	strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
	strFileName="Car.log"
	strParametername=" |"+string(util.Personalize(T("getCMTND_title"), &msg.Sender))+":"
	soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,msg.Text)

	var policyHolderID =strings.ToUpper(msg.Text)
	//var policyHolderID ="`"+msg.Text+"`"
	//policyHolderID=strings.Replace(strconv.Quote(strings.ToUpper(policyHolderID)), `"`, " ", 2)
	//
	//pr := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	//gOUTPUT, _, _ := transform.String(pr ,policyHolderID)
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
	result, _ :=  soap.GetpolicyHolderIDCarInfo(policyHolderID)
	if (result==nil) {
		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		return ConfirmCMTNDEvent
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
		return ConfirmCMTNDEvent
	}
	var msgText, msgAll  string
	msgAll = ""
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

			msgText+="\n" +`BKS/SK: ` + rowSets.Rows[i].REG_NUMBER

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


			msgText+=`, BH tại ` + rowSets.Rows[i].BU_NAME
			msgText+=`, từ ngày ` + rowSets.Rows[i].INCEPTION_DATE

			msgText+=` đến ngày ` + rowSets.Rows[i].EXPIRY_DATE

			msgText+=`, số đơn ` +rowSets.Rows[i].POLICY_URN
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
			if check_TyLeBT!="0%" {
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


	return ConfirmCMTNDEvent
}

