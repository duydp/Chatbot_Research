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
type GetVesselNameBot struct {
	fbbot.BaseStep
}
func (s GetVesselNameBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, util.Personalize(T("getVesselName_title"), &msg.Sender))
	return StayEvent
}

func (s GetVesselNameBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation") == "" {
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "true")
	}

	var strFileName,strParametername,strIDSender string
	strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
	strFileName="Vessel.log"
	strParametername=" |"+string(util.Personalize(T("getVesselName_title"), &msg.Sender))+":"
	soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,msg.Text)


	var Vessel_name =strings.ToUpper(msg.Text)

	//var Vessel_name ="`"+msg.Text+"`"
	//Vessel_name=strings.Replace(strconv.Quote(strings.ToUpper(Vessel_name)), `"`, " ", 2)
	//
	//pr := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	//gOUTPUT, _, _ := transform.String(pr ,Vessel_name)
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
	result, _ :=  soap.GetvesselNameInfo(Vessel_name)
	if (result==nil) {
		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		return ConfirmInputVesselNameEvent
	}
	// Read xml
	byteValue := []byte(result.GetPolicyVesselInfoResult)

	var rowSets config.RowSet
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &rowSets)
	if (rowSets.Rows[0].DATA == "NO DATA FOUND" || rowSets.Rows[0].DATA == "False Authentication" || rowSets.Rows[0].ERROR != "") {
		bot.SendText(msg.Sender, util.Personalize(T("NoData_Found") + "" , &msg.Sender))
		return ConfirmInputVesselNameEvent
	}
	var msgText, msgAll  string
	msgAll = ""
	if len(rowSets.Rows) > 5 {
		iTotal:=len(rowSets.Rows)
		var strCount=strconv.Itoa(iTotal)+" kết quả"
		bot.SendText(msg.Sender, util.Personalize(T("getwebview")+strCount, &msg.Sender))
		bot.STMemory.For(msg.Sender.ID).Set("skey", msg.Text)
		bot.STMemory.For(msg.Sender.ID).Set("svtype", "tau_ten")

		return ConfirmWebViewBotEvent
	} else {
		msgText = "Thông tin hồ sơ: "
		for i := 0; i < len(rowSets.Rows); i++ {
			if len(rowSets.Rows)%1==0{
				check_SUMINSURED_AMT:=rowSets.Rows[i].SUMINSURED_AMT
//dsds
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



	return ConfirmInputVesselNameEvent
}