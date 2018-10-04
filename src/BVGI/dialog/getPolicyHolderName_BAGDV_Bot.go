//author: duydp
//wrote at: 09/5/2018
package dialog
import (
	"util"
	"github.com/michlabs/fbbot"
	"BVGI/config"
	"soap"
	"encoding/xml"
	//"fmt"
	"strings"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/runes"
	"unicode"
	"regexp"
	"strconv"
)
type GetPolicyHolderName_BAGDV_Bot struct {
	fbbot.BaseStep
}
func (s GetPolicyHolderName_BAGDV_Bot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, util.Personalize(T("getPolicyHolderName_title"), &msg.Sender))
	return StayEvent
}
func (s GetPolicyHolderName_BAGDV_Bot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation") == "" {
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "true")
	}
	//var PolicyHolderName_BAGDV =strings.ToUpper(msg.Text)
	var PolicyHolderName_BAGDV ="`"+msg.Text+"`"
	PolicyHolderName_BAGDV=strings.Replace(strconv.Quote(strings.ToUpper(PolicyHolderName_BAGDV)), `"`, " ", 2)

	pr := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	gOUTPUT, _, _ := transform.String(pr ,PolicyHolderName_BAGDV)
	reg, _ := regexp.Compile("[^a-zA-Z0-9-%]+")
	//if err != nil {
	//	log.Fatal(err)
	//}
	prc_Str:= reg.ReplaceAllString(gOUTPUT, " ")

	var v_processedString string
	v_processedString=""
	var b_pre bool=strings.HasPrefix(strings.TrimSpace(prc_Str),`%`)
	var b_sur bool=strings.HasSuffix(strings.TrimSpace(prc_Str),`%`)

	if (b_pre==true && b_sur==false){
		prc_Str:=strings.TrimSpace(prc_Str)
		prc_Str=prc_Str[1:]
		prc_Str=strings.Replace(prc_Str,"%","",-1)
		v_processedString=`%`+strings.TrimSpace(prc_Str)
	}
	if (b_sur==true && b_pre==false){
		prc_Str:=strings.TrimSpace(prc_Str)
		prc_Str=prc_Str[0:len(prc_Str)-1]
		prc_Str=strings.Replace(prc_Str,"%","",-1)
		v_processedString=strings.TrimSpace(prc_Str)+`%`
	}
	if (b_pre==true && b_sur==true){
		prc_Str:=strings.TrimSpace(prc_Str)
		prc_Str=prc_Str[1:]
		prc_Str=prc_Str[0:len(prc_Str)-1]
		prc_Str=strings.Replace(prc_Str,"%","",-1)
		v_processedString=`%`+strings.TrimSpace(prc_Str)+`%`
	}

	if (b_pre==false && b_sur==false){
		prc_Str:=strings.TrimSpace(prc_Str)
		v_processedString=prc_Str
	}
	result, _ :=  soap.Get_BAGDV_CustomerName_Info(v_processedString)



	// Read xml
	byteValue := []byte(result.GetContractInfoResult)

	var rowSets config.RowSet
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &rowSets)
	if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication") {
		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		return ConfirmPolicyHolderNamen_BAGDV_Event
	}
	var msgText, msgAll  string
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

	return ConfirmPolicyHolderNamen_BAGDV_Event
}

