//author: duydp
//wrote at: 28/3/2018
package dialog
import (
	"util"
	"github.com/michlabs/fbbot"
	"BVGI/config"
	"soap"
	"encoding/xml"
	"strings"
	"regexp"
)
type GetPolicyHolderTaxCodeFireBot struct {
	fbbot.BaseStep
}
func (s GetPolicyHolderTaxCodeFireBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, util.Personalize(T("getMST_title"), &msg.Sender))
	return StayEvent
}
func (s GetPolicyHolderTaxCodeFireBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation") == "" {
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "true")
	}

	var strFileName,strParametername,strIDSender string
	strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
	strFileName="Fire.log"
	strParametername=" |"+string(util.Personalize(T("getMST_title"), &msg.Sender))+":"
	soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,msg.Text)

	var policyHolderID_Fire =strings.ToUpper(msg.Text)
	//var policyHolderID_Fire ="`"+msg.Text+"`"
	//policyHolderID_Fire=strings.Replace(strconv.Quote(strings.ToUpper(policyHolderID_Fire)), `"`, " ", 2)
	//
	//pr := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	//gOUTPUT, _, _ := transform.String(pr ,policyHolderID_Fire)
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
	result, _ :=  soap.GetPolicyHolderTaxCodeFireInfo(policyHolderID_Fire)
	if (result==nil) {
		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		return ConfirmInputPolicyHolderTaxCodeEvent
	}


	// Read xml
	byteValue := []byte(result.GetPolicyFireInfoResult)

	var rowSets config.RowSet
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &rowSets)
	if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		return ConfirmInputPolicyHolderTaxCodeEvent
	}
	var msgText, msgAll  string
	msgAll = ""
	msgText = `Thông tin hồ sơ: `

	for i := 0; i < len(rowSets.Rows); i++ {
		if len(rowSets.Rows)%1==0{
			check_SUMINSURED_AMT:=rowSets.Rows[i].SUMINSURED_AMT

			check_NopPhi:=rowSets.Rows[i].PREMIUM_PAYMENT_AMT
			check_NgayDongPhi:=rowSets.Rows[i].NGAY_NOP_PHI

			msgText+="\n" +"Số đơn "+rowSets.Rows[i].POLICY_URN+", khách hàng "+ rowSets.Rows[i].POLICYHOLDER_NAME
			/*if check_PhoneTrade != "" {
				msgText += " (" + check_PhoneTrade + ")"
			}*/
			msgText +=`, tham gia `
			if check_SUMINSURED_AMT !="0"{
				msgText+= rowSets.Rows[i].PRODUCT_NAME
				msgText+=" (" +check_SUMINSURED_AMT+ ")"
			}


			msgText+=`, địa điểm BH: `+rowSets.Rows[i].LOCATION

			msgText+=`, thời hạn BH từ ` + rowSets.Rows[i].INCEPTION_DATE

			msgText+=` đến ` + rowSets.Rows[i].EXPIRY_DATE

			msgText+=`, ` + strings.ToLower(rowSets.Rows[i].TINH_TRANG_THU_PHI)

			if check_NopPhi!="0"{
				if check_NgayDongPhi!=""{msgText+=`, ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
				}
			}
			msgText+=`, tại ` + rowSets.Rows[i].BU_NAME
			if (rowSets.Rows[i].TRADE_NAME!= "") {
				msgText+=`. LH CBKT ` +rowSets.Rows[i].TRADE_NAME
			}

			if rowSets.Rows[i].TRADE_PHONE !=""{
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

	return ConfirmInputPolicyHolderTaxCodeEvent
}