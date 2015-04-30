package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["CoCRM/controllers:PersonalNotesController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PersonalNotesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PersonalNotesController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PersonalNotesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PersonalNotesController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PersonalNotesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PersonalNotesController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PersonalNotesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PersonalNotesController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PersonalNotesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysButtonController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysButtonController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysButtonController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysButtonController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysButtonController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysButtonController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysButtonController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysButtonController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysButtonController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysButtonController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:ParamSysParamController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:ParamSysParamController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:ParamSysParamController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:ParamSysParamController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:ParamSysParamController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:ParamSysParamController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:ParamSysParamController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:ParamSysParamController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:ParamSysParamController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:ParamSysParamController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PersonalCalendarController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PersonalCalendarController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PersonalCalendarController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PersonalCalendarController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PersonalCalendarController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PersonalCalendarController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PersonalCalendarController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PersonalCalendarController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PersonalCalendarController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PersonalCalendarController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysLogErrController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysLogErrController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysLogErrController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysLogErrController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysLogErrController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysLogErrController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysLogErrController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysLogErrController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysLogErrController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysLogErrController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMReceiveController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMReceiveController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMReceiveController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMReceiveController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMReceiveController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMReceiveController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMReceiveController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMReceiveController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMReceiveController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMReceiveController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMFollowController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMFollowController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMFollowController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMFollowController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMFollowController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMFollowController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMFollowController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMFollowController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMFollowController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMFollowController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysRoleController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysRoleController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysRoleController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysRoleController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysRoleController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMProductCategoryController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMProductCategoryController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMProductCategoryController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMProductCategoryController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMProductCategoryController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMProductCategoryController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMProductCategoryController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMProductCategoryController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMProductCategoryController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMProductCategoryController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysMenuController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysMenuController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysMenuController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysMenuController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysMenuController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysMenuController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysMenuController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysMenuController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysMenuController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysMenuController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMContactController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMContactController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMContactController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMContactController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMContactController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMContactController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMContactController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMContactController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMContactController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMContactController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysLogController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysLogController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysLogController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysLogController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysLogController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysLogController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysLogController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysLogController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysLogController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysLogController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMContractController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMContractController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMContractController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMContractController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMContractController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMContractController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMContractController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMContractController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMContractController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMContractController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMInvoiceController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMInvoiceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMInvoiceController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMInvoiceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMInvoiceController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMInvoiceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMInvoiceController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMInvoiceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMInvoiceController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMInvoiceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMOrderController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMOrderController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMOrderController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMOrderController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMOrderController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMOrderController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMOrderController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMOrderController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMOrderController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMOrderController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMCustomerController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMCustomerController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMCustomerController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMCustomerController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMCustomerController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMCustomerController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMCustomerController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMCustomerController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMCustomerController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMCustomerController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrPositionController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrPositionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrPositionController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrPositionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrPositionController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrPositionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrPositionController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrPositionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrPositionController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrPositionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PublicNewsController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PublicNewsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PublicNewsController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PublicNewsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PublicNewsController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PublicNewsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PublicNewsController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PublicNewsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PublicNewsController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PublicNewsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:ToolBatchController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:ToolBatchController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:ToolBatchController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:ToolBatchController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:ToolBatchController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:ToolBatchController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:ToolBatchController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:ToolBatchController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:ToolBatchController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:ToolBatchController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysInfoController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysInfoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysInfoController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysInfoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysInfoController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysInfoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysInfoController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysInfoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysInfoController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysInfoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMProductController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMProductController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMProductController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMProductController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMProductController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMProductController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMProductController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMProductController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:CRMProductController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:CRMProductController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrEmployeeController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrEmployeeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrEmployeeController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrEmployeeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrEmployeeController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrEmployeeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrEmployeeController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrEmployeeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrEmployeeController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrEmployeeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrPostController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrPostController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrPostController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrPostController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrPostController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrPostController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrPostController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrPostController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrPostController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrPostController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysAppController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysAppController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysAppController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysAppController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysAppController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysAppController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysAppController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysAppController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:SysAppController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:SysAppController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PublicNoticeController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PublicNoticeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PublicNoticeController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PublicNoticeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PublicNoticeController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PublicNoticeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PublicNoticeController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PublicNoticeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:PublicNoticeController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:PublicNoticeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrDepartmentController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrDepartmentController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrDepartmentController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrDepartmentController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrDepartmentController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrDepartmentController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrDepartmentController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrDepartmentController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["CoCRM/controllers:HrDepartmentController"] = append(beego.GlobalControllerRouter["CoCRM/controllers:HrDepartmentController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
