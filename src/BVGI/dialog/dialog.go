package dialog

import (
	"github.com/michlabs/fbbot"
	//"expvar"
)
var insujs *fbbot.Dialog

func New() *fbbot.Dialog {
	d := fbbot.NewDialog()

	//var selectbot SelectBot

	var welcomeBot WelcomeBot

	var selectXCGBot SelectXCGBot

	//var getSoHDBot GetSoHDBot
	var selectTaisanBot SelectTaisanBot
	var selectTauthuyBot SelectTauthuyBot
	var selectHangBot SelectHangBot
	var get_TS_MST_Bot Get_TS_MST_Bot
    var confirmTS_MST_Bot ConfirmTS_MST_Bot

	var selectTypeOfRiskbot SelecTypeOfRiskBot
	//var getInfoSoGQQLbot GetInfoBySoGQQLBot
	var displayPolicyInfobot DisplayPolicyInfoBot
	//duydp
	//used for GetCarInfo service returning
	var getBienSoXeBot	GetBienSoXeBot
	var getconfirmBienSoXeBot ConfirmBienSoXeBot

	var getChassisNoMotorBot GetChassisNoMotorBot
	var confirmChassisNoBot  ConfirmChassisNoBot

	var getCMTNDBot		GetCMTNDBot
	var getConfirmCMTNDBot	ConfirmCMTNDBot
	var getGCNBHBot	GetGCNBHBot
	var getConfirmGCNBHBot ConfirmGCNBHBot
	var getCMNDCommonBot		GetCMNDCommonBot
	var getConfirmCMNDCommonBot	ConfirmCMNDCommonBot
	//used for GetPolicyFireInfo service returning
	var getPolicyURNFireBot GetPolicyURNFireBot
	var getConfirmPolicyURNFireBot ConfirmPolicyURNFireBot
	var getPolicyHolderIDFireBot GetPolicyHolderIDFireBot
	var confirmPolicyHolderIDFireBot ConfirmPolicyHolderIDFireBot
	var getPolicyHolderTaxCodeFireBot GetPolicyHolderTaxCodeFireBot
	var confirmPolicyHolderTaxCodeFireBot ConfirmPolicyHolderTaxCodeFireBot
	var getPolicyHolderNameFireBot GetPolicyHolderNameFireBot
	var confirmPolicyHolderNameFireBot ConfirmPolicyHolderNameFireBot
	//used BAGDV service returning
	var getSelectBAGDVBot SelectBAGDVBot
	var getContractNoBAGDVBot GetContractNoBAGDVBot
	var getConfirmContractNoBAGDVBot ConfirmContractNoBAGDVBot
	var getPolicyHolderName_BAGDV_Bot	GetPolicyHolderName_BAGDV_Bot
	var getConfirmPolicyHolderName_BAGDV_Bot	ConfirmPolicyHolderName_BAGDV_Bot

	var getLocationFireBot GetLocationFireBot

	//var getLocationFireBot1 GetLocationFireBot1

	var confirmLocationFireBot ConfirmLocationFireBot
	var confirmWebViewBot	ConfirmWebViewBot
	//used for GetPolicyVesselInfo returning
	var getPolicyURNVesselBot GetPolicyURNVesselBot
	var confirmPolicyURNVesselBot ConfirmPolicyURNVesselBot
	var getVesselNameBot GetVesselNameBot
	var confirmVesselNameBot ConfirmVesselNameBot
	var getRegNumberVesselBot GetRegNumberVesselBot
	var confirmRegNumberVesselBot ConfirmRegNumberVesselBot
	//used for GetPolicyCargoInfo returning
	var getPolicyURNCargoBot GetPolicyURNCargoBot
	var confirmPolicyURNCargoBot ConfirmPolicyURNCargoBot
	var getCargoNameBot GetCargoNameBot
	var confirmCargoNameBot ConfirmCargoNameBot
	var getRegNumberCargoBot GetRegNumberCargoBot
	var confirmRegNumberCargoBot ConfirmRegNumberCargoBot
	var getCargoCustomerNameBot	GetCargoCustomerNameBot
	var confirmCargoCustomerNameBot ConfirmCargoCustomerNameBot

	//duydp

	var getInfoBot GetInfoBot

	var developBot DevelopBot

	var confirmSyntaxBot ConfirmSyntaxBot

	var welcome Welcome
	var faq FAQ
	var silence Silence
	var err Error
	var noanswer NoAnswer
	var staffRegister StaffRegister
	var goodbye Goodbye


	d.SetBeginStep(welcomeBot)
	d.SetEndStep(goodbye)

	d.AddTransition(SelectBotEvent, welcomeBot)
	d.AddTransition(SelectXCGBotEvent, selectXCGBot)
	d.AddTransition(NhapTaiSanEvent,selectTaisanBot)
	d.AddTransition(NhapTauThuyEvent, selectTauthuyBot)
	d.AddTransition(NhapHangEvent,selectHangBot)
    d.AddTransition(Taisan_MST_Event,get_TS_MST_Bot)
    d.AddTransition(ConfirmTS_MST_Event,confirmTS_MST_Bot)
	d.AddTransition(InputBAGDVEvent,getSelectBAGDVBot)

	//d.AddTransition(NhapSoHDEvent, getSoHDBot)
	d.AddTransition(TracuuSyntaxEvent, getInfoBot)

	d.AddTransition(ConfirmSyntaxEvent, confirmSyntaxBot)

	d.AddTransition(SelecTypeOfRiskBotEvent, selectTypeOfRiskbot)
	//duydp added
	//used for GetCarInfo service returning
	d.AddTransition(NhapBienSoXeEvent,getBienSoXeBot)
	d.AddTransition(ConfirmBienSoXeEvent,getconfirmBienSoXeBot)

	d.AddTransition(InputChassisNoEvent,getChassisNoMotorBot)
	d.AddTransition(ConfirmChassisNoEvent,confirmChassisNoBot)

	d.AddTransition(NhapCMTNDEvent,getCMTNDBot)
	d.AddTransition(ConfirmCMTNDEvent,getConfirmCMTNDBot)
	d.AddTransition(NhapGCNBHEvent,getGCNBHBot)
	d.AddTransition(ConfirmGCNBHEvent,getConfirmGCNBHBot)
	d.AddTransition(NhapCMTNDEventCommon,getCMNDCommonBot)
	d.AddTransition(ConfirmCMTNDEventCommon,getConfirmCMNDCommonBot)
	//used for GetPolicyFireInfo service returning
	d.AddTransition(InputPolicyUrnFireEvent,getPolicyURNFireBot)
	d.AddTransition(ConfirmInputPolicyUrnFireEvent,getConfirmPolicyURNFireBot)
	d.AddTransition(InputPolicyHolderIDFireEvent,getPolicyHolderIDFireBot)
	d.AddTransition(ConfirmInputPolicyHolderIDFireEvent,confirmPolicyHolderIDFireBot)
	d.AddTransition(InputPolicyHolderTaxCodeFireEvent,getPolicyHolderTaxCodeFireBot)
	d.AddTransition(ConfirmInputPolicyHolderTaxCodeEvent,confirmPolicyHolderTaxCodeFireBot)
	//new added 12/4/2018: Location parameter
	d.AddTransition(InputLocationFireEvent,getLocationFireBot)

	//d.AddTransition(InputLocationFireEvent1,getLocationFireBot1)

	d.AddTransition(ConfirmWebViewBotEvent,confirmWebViewBot)
	d.AddTransition(ConfirmLocationFireEvent,confirmLocationFireBot)
	//new added 13/4/2018: Policy Holder Name parameter
	d.AddTransition(InputPolicyHolderNameFireEvent,getPolicyHolderNameFireBot)
	d.AddTransition(ConfirmPolicyHolderNamenFireEvent,confirmPolicyHolderNameFireBot)

	//used for GetPolicyVesselInfo returning
	d.AddTransition(InputPolicyUrnVesselEvent,getPolicyURNVesselBot)
	d.AddTransition(ConfirmInputPolicyUrnVesselEvent,confirmPolicyURNVesselBot)
	d.AddTransition(InputVesselNameEvent,getVesselNameBot)
	d.AddTransition(ConfirmInputVesselNameEvent,confirmVesselNameBot)
	d.AddTransition(InputregNumberVesselEvent,getRegNumberVesselBot)
	d.AddTransition(ConfirmInputregNumberVesselEvent,confirmRegNumberVesselBot)
	//used for GetPolicyCargoInfo returning
	d.AddTransition(InputPolicyUrnCargoEvent,getPolicyURNCargoBot)
	d.AddTransition(ConfirmInputPolicyUrnCargoEvent,confirmPolicyURNCargoBot)
	d.AddTransition(InputCargoNameEvent,getCargoNameBot)
	d.AddTransition(ConfirmInputCargoNameEvent,confirmCargoNameBot)
	d.AddTransition(InputregNumberCargoEvent,getRegNumberCargoBot)
	d.AddTransition(ConfirmInputregNumberCargoEvent,confirmRegNumberCargoBot)
	//add 28/8/2018 by duydp
	d.AddTransition(InputCargoCustomerNameEvent,getCargoCustomerNameBot)
	d.AddTransition(ConfirmInputCargoCustomerNameEvent,confirmCargoCustomerNameBot)
	//used for BAGDV returning 9/5/2018
	d.AddTransition(InputContractNo_BAGDV_Event,getContractNoBAGDVBot)
	d.AddTransition(ConfirmInputContractNo_BAGDV_Event,getConfirmContractNoBAGDVBot)
	d.AddTransition(InputPolicyHolderName_BAGDV_Event,getPolicyHolderName_BAGDV_Bot)
	d.AddTransition(ConfirmPolicyHolderNamen_BAGDV_Event,getConfirmPolicyHolderName_BAGDV_Bot)
	//duydp added
	d.AddTransition(DisplayPolicyInfoEvent, displayPolicyInfobot)

	d.AddTransition(DevelopEvent, developBot)

	d.AddTransition(GoWelcomeEvent, welcome)
	d.AddTransition(GoFAQEvent, faq)
	d.AddTransition(ErrorEvent, err)
	d.AddTransition(NoAnswerEvent, noanswer)
	d.AddTransition(GoSlienceEvent, silence)
	d.AddTransition(GoodbyeEvent, goodbye)
	d.AddTransition(StaffRegisterEvent, staffRegister)


	d.PreHandleMessageHook = PreHandlerMessageHook

	insujs = d
	return d
}