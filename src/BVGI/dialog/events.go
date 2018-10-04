package dialog

import (
	"github.com/michlabs/fbbot"
)

const SelectBotEvent fbbot.Event    = "select bot"
const SelectXCGBotEvent fbbot.Event    = "select XCG bot"
const NhapSoHDEvent fbbot.Event     = "nhap SoHD"
const NhapTaiSanEvent fbbot.Event = "select tai san bot"


const InputBAGDVEvent fbbot.Event = "select BAGDV event"

const NhapTauThuyEvent fbbot.Event = "select tau thuy bot"
const NhapHangEvent fbbot.Event = "select hang bot"

const Taisan_MST_Event fbbot.Event = "nhap tai san theo MST bot"
const ConfirmTS_MST_Event fbbot.Event = "xac nhan sau khi nhap TS MST"
const TracuuSyntaxEvent fbbot.Event     = "Tracuu BSX"
const ConfirmSyntaxEvent  fbbot.Event     = "Confirm Syntax"
const ConfirmSoHDEvent fbbot.Event  = "xac nhan SoHD"
const SelecTypeOfRiskBotEvent fbbot.Event = "Select loai rui ro"
const GetInfoSoGQQLEvent fbbot.Event = "lay thong tin ho so tu so GQQL"
//duydp
//used for GetCarInfo service returning
const ConfirmBienSoXeEvent fbbot.Event="xac nhan bien so xe"
const NhapBienSoXeEvent		fbbot.Event="Nhập biển số xe"

const ConfirmChassisNoEvent fbbot.Event="xac nhan so khung"
const InputChassisNoEvent		fbbot.Event="Nhập so khung xe"

const NhapCMTNDEvent	fbbot.Event="Nhap CMND"
const ConfirmCMTNDEvent	fbbot.Event="xac nhan so CMND"
const NhapGCNBHEvent   fbbot.Event="Giay Chung Nhan Bao Hiem"
const ConfirmGCNBHEvent fbbot.Event="Xac nhan Giay Chung Nhan Bao Hiem"
const NhapCMTNDEventCommon	fbbot.Event="Nhap CMND common"
const ConfirmCMTNDEventCommon	fbbot.Event="Xac nhan so CMND common"
//used for GetPolicyFireInfo service returning
const InputPolicyUrnFireEvent fbbot.Event="Input policyURN Fire"
const ConfirmInputPolicyUrnFireEvent fbbot.Event="Confirm Input policyURN Fire"
const InputPolicyHolderIDFireEvent fbbot.Event="Input policy holder ID Fire"
const ConfirmInputPolicyHolderIDFireEvent fbbot.Event="Confirm Input policy holder ID Fire"
const InputPolicyHolderTaxCodeFireEvent fbbot.Event="Input policy holder TaxCode Fire"
const ConfirmInputPolicyHolderTaxCodeEvent fbbot.Event="Confirm Input policy holder TaxCode Fire"

const InputLocationFireEvent fbbot.Event="Input location Fire"

const InputLocationFireEvent1 fbbot.Event="Input location Fire1"
const InputLocationFireEvent_L2 fbbot.Event="Input location Fire L2"

const ConfirmLocationFireEvent fbbot.Event="Confirm location Fire"
const ConfirmWebViewBotEvent	fbbot.Event="Confirm Webview"

const InputPolicyHolderNameFireEvent fbbot.Event="Input Policy Holder Name Fire"
const ConfirmPolicyHolderNamenFireEvent fbbot.Event="Confirm Policy Holder Name Fire"

//used for GetPolicyVesselInfo service returning
const InputPolicyUrnVesselEvent fbbot.Event="Input policyURN Vessel"
const ConfirmInputPolicyUrnVesselEvent fbbot.Event="Confirm Input policyURN Vessel"
const InputVesselNameEvent fbbot.Event="Input  Vessel name"
const ConfirmInputVesselNameEvent fbbot.Event="Confirm Input Vessel name"
const InputregNumberVesselEvent fbbot.Event="Input  reg Number Vessel"
const ConfirmInputregNumberVesselEvent fbbot.Event="Confirm Input reg Number Vessel"
//used for GetPolicyCargoInfo service returning
const InputPolicyUrnCargoEvent fbbot.Event="Input policyURN Cargo"
const ConfirmInputPolicyUrnCargoEvent fbbot.Event="Confirm Input policyURN Cargo"
const InputCargoNameEvent fbbot.Event="Input  Cargo name"


const ConfirmInputCargoNameEvent fbbot.Event="Confirm Input Cargo name"
const InputregNumberCargoEvent fbbot.Event="Input  reg Number Cargo"
const ConfirmInputregNumberCargoEvent fbbot.Event="Confirm Input reg Number Cargo"

const InputCargoCustomerNameEvent fbbot.Event="Input Cargo customer name"
const ConfirmInputCargoCustomerNameEvent fbbot.Event="Confirm Input Cargo Customer Name"

//used for Get BAGDV services
const InputContractNo_BAGDV_Event fbbot.Event="Input Contract No BAGDV"
const ConfirmInputContractNo_BAGDV_Event fbbot.Event="Confirm Input Contract No BAGDV"

const InputPolicyHolderName_BAGDV_Event fbbot.Event="Input Policy Holder Name BAGDV"
const ConfirmPolicyHolderNamen_BAGDV_Event fbbot.Event="Confirm Policy Holder Name BAGDV"

//duydp
const ImageDocDialogEvent fbbot.Event  	= "upload tai lieu image"
const ConfirmSoGQQLEvent fbbot.Event = "xac nhan So GQQL"
const CreateRequestEvent fbbot.Event = "create request"
const DisplayPolicyInfoEvent fbbot.Event = "Hien thi thong tin ho so"

const StayEvent fbbot.Event         = ""

const DevelopEvent fbbot.Event      = "Dang xay dung"
const GoWelcomeEvent fbbot.Event    = "go welcome"
const GoFAQEvent fbbot.Event        = "go to FAQ"
const ErrorEvent fbbot.Event        = "error"
const GoodbyeEvent fbbot.Event      = "goodbye"
const GoSlienceEvent fbbot.Event    = "go slience"
const NoAnswerEvent fbbot.Event		= "No answer"
const StaffRegisterEvent fbbot.Event= "Staff register"

//booking add
const CheckBookedEvent fbbot.Event 	= "check booked"
const FillSportEvent fbbot.Event  	= "fillsport"
const FillAddressEvent fbbot.Event  	= "filladdress"
const FillDatetimeEvent fbbot.Event  	= "filldatetime"

const NewBookEvent fbbot.Event  	  = "Booking-common-request"
const NewBookWithdetailEvent fbbot.Event  = "Booking-detailed-request"
const BookL2Event fbbot.Event  		= "booklayer2"
const BookL3Event fbbot.Event  		= "booklayer3"
const BookL4Event fbbot.Event  		= "booklayer4"
const FaQStadiumEvent fbbot.Event  	= "faq stadium"
const FaQAthleteEvent fbbot.Event  	= "faq athlete"
const FaQSportEvent fbbot.Event  	= "faq sport"
const GuideEvent fbbot.Event  		= "guide"
//intent
const CheckBookedIntent  		= "checkBooked"
const BookIntent 			= "Booking-common-request"
const BookWithDetailIntent 		= "Booking-detailed-request"
const Address_intent 			= "address"
const Datetime_intent 			= "datetime"
const BookFillIntent 			= "book_fill"
const FaQStadiumIntent 			= "faq_stadium"
const FaQAthleteIntent 			= "faq_athlete"
const FaQSportIntent 			= "faq_sport"
const GuideIntent 			= "guide"
const OfftopicIntent			= "off_topic"