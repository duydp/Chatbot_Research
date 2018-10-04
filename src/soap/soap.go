package soap

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/clbanning/mxj"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"github.com/michlabs/fbbot"
	"encoding/json"
	"os"
	"io"
	"unicode/utf8"
	"strings"

	"log"
	"time"
)

type SoapResult struct {
	Status  string
	Message string
	Result  string
}
type SResult struct {
	SearchCommonInfoResult  string
	GetPolicyInfoResult string
	GetCarInfoResult string
	GetPolicyFireInfoResult string
	GetPolicyVesselInfoResult string
	GetPolicyCargoInfoResult string
	GetContractInfoResult string
}
type LOCATION struct {
	LOC_PROVINCE	string//tinh,tp
	LOC_WARD	string//phuong,xa
	LOC_DISTRICT string//quan,huyen
	LOC_FULL 	string//addr,number of street
	LOC_STREET	string//street
	BU_NAME	string
	TRADE_NAME	string
	TRADE_PHONE	string
	PRODUCT_NAME	string
	COV_CLASS_NAME	string
	POLICY_URN	string
	POLICYHOLDER_NAME	string
	INCEPTION_DATE	string
	EXPIRY_DATE	string
	SUMINSURED_AMT	string
	PREMIUM_AMT	string
	PREMIUM_AMT_NO_VAT	string
	PREMIUM_PAYMENT_AMT	string
	NGAY_NOP_PHI	string
	TINH_TRANG_THU_PHI	string
	SO_VU_BT	string
	DA_TRA_BT	string
	TY_LE_BT	string
}


type OptList struct{
	Title string
	Buttons []fbbot.Button
}

type ConfirmAddMessageItems struct {
	fbbot.BaseStep
	ListItems []OptList
}
type MessageItemAddContentDialog struct {
	ConfirmAddMessageItems
}
func CompareChars(word string) {
	s := []byte(word)
	for utf8.RuneCount(s) > 1 {
		r, size := utf8.DecodeRune(s)
		s = s[size:]
		nextR, size := utf8.DecodeRune(s)
		fmt.Print(r == nextR, ",")
	}
}

func RemoveUnicodeFont(strUnicode string) string{
	var strTCVN3Vowels,strVNWithoutMarkVowels string
	strTCVN3Vowels = "áàảãạăắằẳẵặâấầẩẫậéèẻẽẹêếềểễệíìỉĩịóòỏõọôốồổỗộơớờởỡợúùủũụưứừửữựýỳỷỹỵđÁÀẢÃẠĂẮẰẲẴẶÂẤẦẨẪẬÉÈẺẼẸÊẾỀỂỄỆÍÌỈĨỊÓÒỎÕỌÔỐỒỔỖỘƠỚỜỞỠỢÚÙỦŨỤƯỨỪỬỮỰÝỲỶỸỴĐ";
	fmt.Print(len(strTCVN3Vowels))
	strVNWithoutMarkVowels = "aaaaaaaaaaaaaaaaaeeeeeeeeeeeiiiiiooooooooooooooooouuuuuuuuuuuyyyyydAAAAAAAAAAAAAAAAAEEEEEEEEEEEIIIIIOOOOOOOOOOOOOOOOOUUUUUUUUUUUYYYYYD";
	fmt.Print(len(strVNWithoutMarkVowels))
	var strTemp  string
	//re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	//re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	//strTemp = re_leadclose_whtsp.ReplaceAllString(strUnicode, "")
	//strTemp = re_inside_whtsp.ReplaceAllString(strTemp, " ")
	strTemp=strUnicode
	if len(strTemp)==0{
		return " "
	}
	for j:=0;j<len(strTCVN3Vowels);j++{
		var OldStr =strTCVN3Vowels[j]

		var NewStr=strVNWithoutMarkVowels[j]

		strTemp=strings.Replace(strTemp,string(OldStr),string(NewStr),-1)
	}

	return strTemp

}
func WriteDataOnCacheJSon(dataJson []byte, bot *fbbot.Bot, msg *fbbot.Message) (error) {
	stMemoryKey := msg.Sender.ID
	stMemoryValue := bot.STMemory.For(msg.Sender.ID).Get(stMemoryKey)
	if stMemoryValue != "" {
		var memoryJson []byte
		err := json.Unmarshal([]byte(stMemoryValue),&memoryJson)
		if err != nil {
			bot.Logger.Errorf(" unmarshal json value fail!")
			fmt.Print("cant unmarshal json value!")
		}
		for i := 0; i < len(memoryJson); i++ {
			dataJson = append(dataJson, memoryJson[i])
		}
	}
	jsonvar, _ := json.Marshal(dataJson)
	bot.STMemory.For(msg.Sender.ID).Set(msg.Sender.ID, string(jsonvar))
	return nil
}

func ChooseList(msg *fbbot.Message,bot *fbbot.Bot,strMsg string,itemPage int, totalPage int, currentPage int,Lists []OptList,vALL string){
	var ListItem []fbbot.Bubble

	for _,List:=range Lists{
		ListItem=append(ListItem,fbbot.Bubble{
			Title:List.Title,
			Buttons:List.Buttons,
		})

	}
	//set List has only one elements
	if len(Lists)==1{
		ListItem=append(ListItem,fbbot.Bubble{
			Title: "",
			SubTitle: "",
		})
	}

	var v_MAXItems int
	v_MAXItems=itemPage*(currentPage+1)
	var Buttons []fbbot.Button
	if (v_MAXItems >= totalPage) {
		Buttons = append(Buttons, fbbot.Button{
			Type: "postback",
			Title: ".",
			Payload:vALL,
		})
	}
	botListItem:=new(fbbot.ListTemplateMessage)
	botListItem.Bubbles=ListItem
	botListItem.Buttons=Buttons
	botListItem.Text=strMsg
	bot.Send(msg.Sender,botListItem)

}
//duydp copied & modified
type BinaryNode struct{
	left *BinaryNode
	right *BinaryNode
	data string
}

type BinaryTree struct{
	root *BinaryNode
}

func (t *BinaryTree) insert(data string) *BinaryTree{
	if t.root==nil{
		t.root=&BinaryNode{data: data,left:nil,right:nil}
	}else{
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data string){
	if n==nil{
		return
	}else if data <=n.data{
		if n.left==nil{
			n.left=&BinaryNode{data: data,left:nil,right:nil}
		}else{
			n.left.insert(data)
		}
	}else{
		if n.right==nil{
			n.right=&BinaryNode{data: data,left:nil,right:nil}
		}else{
			n.right.insert(data)
		}
	}
}
func print(w io.Writer,node *BinaryNode,ns int,ch rune){
	if node==nil{
		return
	}
	for i:=0;i<ns;i++{
		fmt.Fprint(w," ")
	}
	fmt.Fprint(w,"%c:%v\n",ch,node.data)
	print(w,node.left,ns+2,'L')
	print(w,node.right,ns+2,'R')
}

//end
func generateRequestContent(funcName, ParamXML string) string {
	var getTemplate = `<?xml version="1.0" encoding="utf-8"?>
    <soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
      <soap:Body>
        <`+funcName+` xmlns="http://tempuri.org/">
        	`+ParamXML+`
        </`+funcName+`>
      </soap:Body>
    </soap:Envelope>`
	/*
	tmpl, err := template.New("getService").Parse(getTemplate)
	if err != nil {
		panic(err)
	}
	var doc bytes.Buffer
	err = tmpl.Execute(&doc, nil)
	if err != nil {
		panic(err)
	}
	return doc.String()
	*/
	return getTemplate
}

func Post(funcName, ParamXML string) (*SResult, error) {
	url :=  os.Getenv("BVGI_SERVICE_URL")

	timeout := time.Duration(10 * time.Minute)
	client := http.Client{
		Timeout: timeout,
	}
	sRequestContent := generateRequestContent(funcName, ParamXML)
	requestContent := []byte(sRequestContent)



	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestContent))


	//httpClient := httpclient.NewWithTimeout(500*time.Millisecond, 1*time.Second)
	//resp, err := httpClient.Get("http://google.com")
	//if err != nil {
	//	fmt.Println("Rats! Google is down.")
	//}

	//ctx := context.WithTimeout(context.Background(), 5 * time.Second)
	//req, _ := http.NewRequest("GET", "https://mota.cf", nil)
	//req = req.WithContext(ctx)
	//
	//resp, _ := http.DefaultClient.Do(req)



	if err != nil {
		fmt.Println("Loi 1")
		return nil, err
	}

	req.Header.Add("SOAPAction", "http://tempuri.org/"+funcName)
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Accept", "text/xml")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Loi 3, StatusCode=" + string(resp.StatusCode) + ", Status=" + resp.Status + "")
		return nil, errors.New("Error Respose " + resp.Status)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Loi 4")
		return nil, err
	}
	m, _ := mxj.NewMapXml(contents, true)
	return convertResults(&m, funcName)
	//return string(contents), nil
}

//filename: name of file product
//pre_message: name's module has written.
//out_message: msg.text
func AppendStringToFileServer(IDSender,filename,parametername,parametervalue string)error{

	stTypeOfOption := os.Getenv("BVGI_PATH")+"/logs/"+filename

	//case : File does not exist
	if _, err := os.Stat(stTypeOfOption); os.IsNotExist(err) {
//| log.Lshortfile
		out, _ := os.Create(stTypeOfOption)
		flag := log.LstdFlags | log.Lmicroseconds
		prefix := IDSender
		newLog := log.New(out, prefix, flag)
		newLog.Println(parametername+parametervalue)
		//case: File exists
	} else {
		//create your file with desired read/write permissions
		f, err := os.OpenFile(stTypeOfOption, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			log.Fatal(err)
		}
		//defer to close when you're done with it, not because you think it's idiomatic!
		defer f.Close()
		//set output of logs to f
		log.SetOutput(f)
		prefix := parametername+parametervalue
		log.Println(IDSender,prefix)


	}
	return nil
}
func convertResults(soapResponse *mxj.Map, funcName string) (*SResult, error) {
	// Read tag contains results
	objResult, err := soapResponse.ValueForPath("Envelope.Body."+funcName+"Response")

	var result SResult
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
		// add a DecodeHook here if you need complex Decoding of results -> DecodeHook: yourfunc,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(objResult); err != nil {
		return nil, err
	}
	return &result, nil
}

