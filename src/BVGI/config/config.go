package config

import(
	"github.com/kelseyhightower/envconfig"
	"encoding/xml"
)

var Bot BotConfig
var UI UIConfig
var DB DBConfig
var SC	SearchCommonInfo
//var RowSets RowSet
//var Rows	Row
//author: duydp wrote on 12/3/2018 (note write: capital ex: Authen_key)
type SearchCommonInfo struct{
	Authen_key	string `default:"D1D606D57489A58D3E26DDF5FD2D3E00"`
	PageNumber	int		`default:"1"`
	PageSize	int		`default:"1"`
}
//author: duydp wrote on 12/3/2018

type BotConfig struct {
	Port            int    `default:"1203"`
	VerifyToken     string `required:"true"`
	PageAccessToken string `required:"true"`
	LogFile string `default:"bot.log"`
	LogLevel        string `default:"info"`
	Wit        map[string]string `required:"true"`
	FPTAI	map[string]string `required:"true"`
	LanguageFile        string `required:"true"`
	DebugFile	string `required:"true"`
	Hex string `required:"true"`
	ConversationTimeout float64 `required:"true"`
	SilenceDuration float64 `required:"true"`
}

type UIConfig struct {
	Port       int    `required:"true"`
	Username   string `required:"true"`
	Userpass   string `required:"true"`
	StaticDir  string `required:"true"`
	BotName	   string `required:"true"`
}

type DBConfig struct {
	Host     string `required:"true"`
	Port	 string `default:"3306"`
	Name     string `required:"true"`
	User     string `required:"true"`
	Pass string `required:"true"`
}

func LoadFromEnv() error {
	if err := envconfig.Process("bvgi", &Bot); err != nil {
		return err
	}

	if err := envconfig.Process("ui", &UI); err != nil {
		return err
	}

	if err := envconfig.Process("db", &DB); err != nil {
		return err
	}

	return nil
}

type RowSet struct {

	XMLName xml.Name `xml:"ROWSET"`
	Rows   []Row   `xml:"ROW"`

}

type Row struct {
	//SearchCommonInfoResult
	XMLName xml.Name `xml:"ROW"`
	ContactID string `xml:"CONTACT_ID"`
	ChuHopDong string `xml:"CHU_HOPDONG"`
	SoHopDong string `xml:"SO_HOP_DONG"`
	BienSoXe string `xml:"BIEN_SOXE"`
	TrangThaiHopDong string `xml:"TRANGTHAI_HOPDONG"`
	NgayHieuLuc string `xml:"NGAY_HIEU_LUC"`
	NgayHetHieuLuc string `xml:"NGAY_HET_HIEU_LUC"`
	TenSanPham string `xml:"TEN_SANPHAM"`
	MaDVKD string `xml:"MA_DVKD"`
	DONVIKINHDOANH string `xml:"DONVI_KINHDOANH"`
	MACNDVKD string `xml:"MA_CN_DVKD"`
	CHINHANHDONVIKINHDOANH string `xml:"CHINHANH_DONVI_KINHDOANH"`
	PHONE string `xml:"PHONE"`
	SOKHUNG string `xml:"SOKHUNG"`
	SOMAY string `xml:"SOMAY"`
	CAR_GROUP string `xml:"CAR_GROUP"`
	MAKE string `xml:"MAKE"`
	MODEL string `xml:"MODEL"`
	EC_NEWFOROLD string `xml:"EC_NEWFOROLD"`
	EC_LOSS_OF_USE_EXTENSION string `xml:"EC_LOSS_OF_USE_EXTENSION"`
	EC_REPAIRER_OPTION string `xml:"EC_REPAIRER_OPTION"`
	EC_OPTIONAL_DEDUCTIBLE string `xml:"EC_OPTIONAL_DEDUCTIBLE"`
	EC_ACC_OUT_OF_VIETNAM string `xml:"EC_ACC_OUT_OF_VIETNAM"`
	EC_INDEMNITY_LIMIT string `xml:"EC_INDEMNITY_LIMIT"`
	EC_PARTIAL_THEFT string `xml:"EC_PARTIAL_THEFT"`
	EC_ENGINE_DAMAGED_BY_FLOOD string `xml:"EC_ENGINE_DAMAGED_BY_FLOOD"`
	DATA string `xml:"DATA"`
	//GetPolicyInfo
	TRADE_NAME string `xml:"TRADE_NAME"`
	PRODUCT_NAME string `xml:"PRODUCT_NAME"`
	//GetCarInfoResult
	POLICY_URN string`xml:"POLICY_URN"`
	POLICYHOLDER_NAME string`xml:"POLICYHOLDER_NAME"`
	POLICYHOLDER_PHONE string`xml:"POLICYHOLDER_PHONE"`
	BU_NAME string`xml:"BU_NAME"`
	TRADE_PHONE string`xml:"TRADE_PHONE"`
	INCEPTION_DATE string`xml:"INCEPTION_DATE"`
	EXPIRY_DATE string`xml:"EXPIRY_DATE"`
	REG_NUMBER string`xml:"REG_NUMBER"`
	CERT_NUMBER string`xml:"CERT_NUMBER"`
	CHASSIS_NO string`xml:"CHASSIS_NO"`
	VAT_CHAT_XE string`xml:"VAT_CHAT_XE"`
	TNDS_BAT_BUOC string`xml:"TNDS_BAT_BUOC"`
	TNDS_TU_NGUYEN string`xml:"TNDS_TU_NGUYEN"`
	LAI_PHU string`xml:"LAI_PHU"`
	HANG_HOA string`xml:"HANG_HOA"`
	PREMIUM_AMT string`xml:"PREMIUM_AMT"`
	PREMIUM_PAYMENT_AMT string`xml:"PREMIUM_PAYMENT_AMT"`
	PREMIUM_AMT_NO_VAT string`xml:"PREMIUM_AMT_NO_VAT"`
	NGAY_NOP_PHI string`xml:"NGAY_NOP_PHI"`
	TINH_TRANG_THU_PHI string`xml:"TINH_TRANG_THU_PHI"`
	SO_VU_BT string`xml:"SO_VU_BT"`
	DA_TRA_BT string`xml:"DA_TRA_BT"`
	UOC_BT string`xml:"UOC_BT"`
	THU_HOI string`xml:"THU_HOI"`
	TY_LE_BT string`xml:"TY_LE_BT"`

	//GetPolicyFireInfo
	COV_CLASS_NAME string`xml:"COV_CLASS_NAME"`
	SUMINSURED_AMT string`xml:"SUMINSURED_AMT"`
	LOCATION string`xml:"LOCATION"`
	LOC_STREETNO	string`xml:"LOC_STREETNO"`
	LOC_STREETNAME	string`xml:"LOC_STREETNAME"`
	LOC_STREETTYPE	string`xml:"LOC_STREETTYPE"`
	LOC_SUBURB	string`xml:"LOC_SUBURB"`
	LOC_CITY	string`xml:"LOC_CITY"`
	LOC_STATE	string`xml:"LOC_STATE"`
	//GetPolicyVesselInfo
	NAME_OF_VESSEL string`xml:"NAME_OF_VESSEL"`
	REGISTRATIONNO_IMO string`xml:"REGISTRATIONNO_IMO"`
	CERTIFICATENUMBER string`xml:"CERTIFICATENUMBER"`
	//GetPolicyCargoInfo added 3/4/2018
	PLACEOFDEPARTURES string`xml:"PLACEOFDEPARTURES"`
	FINALDESTINATIONS string`xml:"FINALDESTINATIONS"`
	VESSELORCONVEYANCE string`xml:"VESSELORCONVEYANCE"`
	//GetPolicyInfo
	BUSINESS_LINE string`xml:"BUSINESS_LINE"`
	//Get BAGDV
	INSURED_NAME string`xml:"INSURED_NAME"`
	CONTRACT_NUMBER string`xml:"CONTRACT_NUMBER"`
	EXPIRED_DATE string`xml:"EXPIRED_DATE"`
	SUM_INSURED string`xml:"SUM_INSURED"`
	PREMIUM string`xml:"PREMIUM"`
	PAYMENT_MONEY string`xml:"PAYMENT_MONEY"`
	SO_TIEN_BT string`xml:"SO_TIEN_BT"`
	QUYEN_LOI_CON_LAI string`xml:"QUYEN_LOI_CON_LAI"`

	ERROR string `xml:"ERROR"`
	DEPT_NAME string `xml:"DEPT_NAME"`
}

