//author: duydp
//wrote at: 21/3/2018
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

type Get_TS_MST_Bot struct {
	fbbot.BaseStep
}
func (s Get_TS_MST_Bot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, util.Personalize(T("getMST_title"), &msg.Sender))
	return StayEvent
}

func (s Get_TS_MST_Bot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation") == "" {
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "true")
	}
	var strFileName,strParametername,strIDSender string
	strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
	strFileName="Fire.log"
	strParametername=" |"+string(util.Personalize(T("getMST_title"), &msg.Sender))+":"
	soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,msg.Text)
	var strMST =strings.ToUpper(msg.Text)
	//var strMST ="`"+msg.Text+"`"
	//strMST=strings.Replace(strconv.Quote(strings.ToUpper(strMST)), `"`, " ", 2)
	//
	//pr := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	//gOUTPUT, _, _ := transform.String(pr ,strMST)
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
	result, _ :=  soap.GetPolicyHolderTaxCodeFireInfo(strMST)
	if (result==nil) {
		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		return ConfirmTS_MST_Event
	}
	// Read xmlGetPolicyHolderTaxCodeFireInfo
	byteValue := []byte(result.GetPolicyFireInfoResult)

	var rowSets config.RowSet
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &rowSets)
	if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
		//+ rowSets.Rows[0].DATA
		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		return ConfirmTS_MST_Event
	}
	var msgText, msgAll  string

	msgAll = ""
	msgText = `Thông tin hồ sơ: ` + "\n"
	for i := 0; i < len(rowSets.Rows); i++ {
		if len(rowSets.Rows)%1==0{
			check_NopPhi:=rowSets.Rows[i].SUMINSURED_AMT

			check_NgayNopPhi:=strings.Replace(rowSets.Rows[i].NGAY_NOP_PHI,"/", "",1)


			msgText+=`Số đơn ` + rowSets.Rows[i].POLICY_URN

			msgText+=`, khách hàng ` + rowSets.Rows[i].TRADE_NAME

			msgText+=`, tham gia ` + rowSets.Rows[i].PRODUCT_NAME

			if check_NopPhi!="0"{
				msgText+=` ` + "("+rowSets.Rows[i].SUMINSURED_AMT+")"}

			msgText+=`, địa điểm : `+rowSets.Rows[i].LOCATION

			msgText+=`, từ ` + rowSets.Rows[i].INCEPTION_DATE

			msgText+=` đến ` + rowSets.Rows[i].EXPIRY_DATE

			if check_NgayNopPhi!=""{
				msgText+=` ngày ` + rowSets.Rows[i].NGAY_NOP_PHI
			}

			msgText+=`, tại ` + rowSets.Rows[i].BU_NAME

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


	return ConfirmTS_MST_Event
}

