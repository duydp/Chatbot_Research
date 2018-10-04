//author: duydp
//wrote at: 28/8/2018
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
type GetCargoCustomerNameBot struct {
	fbbot.BaseStep
}
func (s GetCargoCustomerNameBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, util.Personalize(T("getPolicyHolderName_title"), &msg.Sender))
	return StayEvent
}

func (s GetCargoCustomerNameBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation") == "" {
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "true")
	}
	var strFileName,strParametername,strIDSender string
	strIDSender=string(msg.Sender.ID+":"+msg.Sender.Name()+" |")
	strFileName="Cargo.log"
	strParametername=" |"+string(util.Personalize(T("getPolicyHolderName_title"), &msg.Sender))+":"
	soap.AppendStringToFileServer(strIDSender,strFileName,strParametername,msg.Text)

	var customer_name =strings.ToUpper(msg.Text)

	result, _ :=  soap.GetCargoCustomerNameInfo(customer_name)
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
		bot.STMemory.For(msg.Sender.ID).Set("skey", msg.Text)
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



	return ConfirmInputCargoCustomerNameEvent
}