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

	"strconv"
)
type GetLocationFireBot struct {
	fbbot.BaseStep
	soap.MessageItemAddContentDialog
}

func (s GetLocationFireBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, util.Personalize(T("getLocationFire_title"), &msg.Sender))

	return StayEvent
}
func (s GetLocationFireBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation") == "" {
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "true")
	}

	var strFileName, strParametername, strIDSender string
	strIDSender = string(msg.Sender.ID + ":" + msg.Sender.Name() + " |")
	strFileName = "Fire.log"
	strParametername = " |" + string(util.Personalize(T("getLocationFire_title"), &msg.Sender)) + ":"
	soap.AppendStringToFileServer(strIDSender, strFileName, strParametername, msg.Text)

	//var Location_Fire ="`"+msg.Text+"`"
	var Location_Fire= strings.ToUpper(msg.Text)
	//Location_Fire=strings.Replace(strconv.Quote(strings.ToUpper(Location_Fire)), `"`, " ", 2)
	//fmt.Println(Location_Fire)
	//pr := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	//gOUTPUT, _, _ := transform.String(pr ,strings.Replace(strings.ToTitle(Location_Fire),"Đ","D",-1))
	//
	//reg, _ := regexp.Compile("[^a-zA-Z0-9-%]+")
	//
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
	result, _ := soap.GetLocationFireInfo(Location_Fire)
	if (result == nil) {
		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found")+"", &msg.Sender))
		return ConfirmLocationFireEvent
	}
	byteValue := []byte(result.GetPolicyFireInfoResult)
	var rowSets config.RowSet
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &rowSets)

	if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found")+"", &msg.Sender))
		return ConfirmLocationFireEvent
	}
	var msgText, msgAll string

	if len(rowSets.Rows) > 5 {
		iTotal:=len(rowSets.Rows)
		var strCount=strconv.Itoa(iTotal)+" kết quả"
		bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
		bot.STMemory.For(msg.Sender.ID).Set("skey", msg.Text)

		bot.STMemory.For(msg.Sender.ID).Set("svtype", "ts_location")

		return ConfirmWebViewBotEvent
	} else {

		msgText = "Thông tin hồ sơ: "
		for j := 0; j < len(rowSets.Rows); j++ {
			if len(rowSets.Rows)%1 == 0 {
				msgText += "\n"
				if (strings.Contains(msgText, rowSets.Rows[j].LOCATION) && rowSets.Rows[j].LOCATION == "") {
				} else {
					msgText += rowSets.Rows[j].LOCATION
				}

				if (strings.Contains(msgText, rowSets.Rows[j].POLICYHOLDER_NAME) && rowSets.Rows[j].POLICYHOLDER_NAME == "") {
				} else {
					msgText += ", khách hàng " + rowSets.Rows[j].POLICYHOLDER_NAME
				}

				if (strings.Contains(msgText, rowSets.Rows[j].BU_NAME) && rowSets.Rows[j].BU_NAME == "") {
				} else {
					msgText += `, BH tại ` + rowSets.Rows[j].BU_NAME
				}

				if (strings.Contains(msgText, rowSets.Rows[j].INCEPTION_DATE) && rowSets.Rows[j].INCEPTION_DATE == "") {
				} else {
					msgText += `, từ ngày ` + rowSets.Rows[j].INCEPTION_DATE
				}

				if (strings.Contains(msgText, rowSets.Rows[j].EXPIRY_DATE) && rowSets.Rows[j].EXPIRY_DATE == "") {
				} else {
					msgText += ` đến ngày ` + rowSets.Rows[j].EXPIRY_DATE
				}

				if (strings.Contains(msgText, rowSets.Rows[j].POLICY_URN) && rowSets.Rows[j].POLICY_URN == "") {
				} else {
					msgText += ". Số đơn " + rowSets.Rows[j].POLICY_URN
				}

				if (strings.Contains(msgText, rowSets.Rows[j].COV_CLASS_NAME) && rowSets.Rows[j].COV_CLASS_NAME == "") {
				} else {
					msgText += ", tham gia BH " + rowSets.Rows[j].COV_CLASS_NAME
				}
				if (strings.HasPrefix(strings.Trim(rowSets.Rows[j].SUMINSURED_AMT, ""), "0") && rowSets.Rows[j].SUMINSURED_AMT == "") {
				} else {
					msgText += " (" + rowSets.Rows[j].SUMINSURED_AMT + ")"
				}

				if (rowSets.Rows[j].TINH_TRANG_THU_PHI != "" && rowSets.Rows[j].TINH_TRANG_THU_PHI == "Đã nộp phí") {
				} else {
					msgText += ", " + strings.ToLower(rowSets.Rows[j].TINH_TRANG_THU_PHI)
				}
				if (strings.Contains(msgText, rowSets.Rows[j].NGAY_NOP_PHI) && rowSets.Rows[j].NGAY_NOP_PHI == "") {
				} else {
					msgText += `, nộp phí ngày ` + rowSets.Rows[j].NGAY_NOP_PHI
				}

				if (strings.Contains(msgText, rowSets.Rows[j].SO_VU_BT) || rowSets.Rows[j].SO_VU_BT == "" || rowSets.Rows[j].SO_VU_BT == "0") {
				} else {

					msgText += `. Số vụ tổn thất: ` + rowSets.Rows[j].SO_VU_BT + " vụ,"
				}

				if (strings.Contains(msgText, rowSets.Rows[j].TY_LE_BT) || rowSets.Rows[j].TY_LE_BT == "" || rowSets.Rows[j].TY_LE_BT == "0%") {
				} else {
					msgText += ` tỷ lệ BT: ` + rowSets.Rows[j].TY_LE_BT
				}

				if (strings.Contains(msgText, rowSets.Rows[j].TRADE_NAME) && rowSets.Rows[j].TRADE_NAME != "") {
				} else {
					msgText += `. LH CBKT ` + rowSets.Rows[j].TRADE_NAME
				}

				if (rowSets.Rows[j].TRADE_PHONE != "" && len(rowSets.Rows[j].TRADE_PHONE) > 1) {
					var v_string_p string = rowSets.Rows[j].TRADE_PHONE
					var rst string
					v_reg, _ := regexp.Compile("[^Z0-9]+")
					prc_str := v_reg.ReplaceAllString(v_string_p, " ")
					rst = strings.TrimSpace(prc_str)
					rst = strings.Replace(rst, ` `, `/`, -1)
					msgText += " (" + rst + ")"
				}

				msgText += "."
				msgAll = msgAll + msgText

				bot.SendText(msg.Sender, msgAll)
			}
			msgText = ""
			msgAll = ""
		}
	}
		return ConfirmLocationFireEvent

}
