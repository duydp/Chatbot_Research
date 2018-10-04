//author: duydp
//wrote at: 8/5/2018
package dialog
import (
	"util"
	"github.com/michlabs/fbbot"
	"BVGI/config"
	"soap"
	"encoding/xml"
	"strings"
	"strconv"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/runes"
	"unicode"
	"regexp"
)
type GetContractNoBAGDVBot struct {
	fbbot.BaseStep
}
func (s GetContractNoBAGDVBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, util.Personalize(T("getSoHDBH_title"), &msg.Sender))
	return StayEvent
}

func (s GetContractNoBAGDVBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation") == "" {
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "true")
	}
	//var ContractNo =strings.ToUpper(msg.Text)
	var ContractNo ="`"+msg.Text+"`"
	ContractNo=strings.Replace(strconv.Quote(strings.ToUpper(ContractNo)), `"`, " ", 2)

	pr := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	gOUTPUT, _, _ := transform.String(pr ,ContractNo)
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
	result, _ :=  soap.Get_BAGDV_ConstractNo_Info(v_processedString)
	// Read xml
	byteValue := []byte(result.GetContractInfoResult)

	var rowSets config.RowSet
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &rowSets)
	if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication") {
		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		return ConfirmInputContractNo_BAGDV_Event
	}
	var msgText  string
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
		msgText = "Thông tin hồ sơ:"+"\n"+vPOLICYURN+vBU+vINCEPTION_DATE+vEXPIRY_DATE+"\n"+vList+InsuredNameStr1row
	} else {
		msgText = "Thông tin hồ sơ:"+"\n"+vPOLICYURN+vBU+vINCEPTION_DATE+vEXPIRY_DATE+"\n"+vList+ msgText + InsuredNameStr
	}

	bot.SendText(msg.Sender, msgText)

	return ConfirmInputContractNo_BAGDV_Event
}