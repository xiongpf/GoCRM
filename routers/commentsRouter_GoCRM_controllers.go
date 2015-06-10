package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["GoCRM/controllers:PersonalCalendarController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PersonalCalendarController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PersonalCalendarController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PersonalCalendarController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PersonalCalendarController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PersonalCalendarController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PersonalCalendarController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PersonalCalendarController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PersonalCalendarController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PersonalCalendarController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:ToolBatchController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:ToolBatchController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:ToolBatchController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:ToolBatchController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:ToolBatchController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:ToolBatchController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:ToolBatchController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:ToolBatchController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:ToolBatchController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:ToolBatchController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PublicNoticeController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PublicNoticeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PublicNoticeController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PublicNoticeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PublicNoticeController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PublicNoticeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PublicNoticeController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PublicNoticeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PublicNoticeController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PublicNoticeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysButtonController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysButtonController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysButtonController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysButtonController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysButtonController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysButtonController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysButtonController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysButtonController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysButtonController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysButtonController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMOrderController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMOrderController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMOrderController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMOrderController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMOrderController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMOrderController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMOrderController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMOrderController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMOrderController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMOrderController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrPositionController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrPositionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrPositionController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrPositionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrPositionController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrPositionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrPositionController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrPositionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrPositionController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrPositionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysLogErrController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysLogErrController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysLogErrController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysLogErrController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysLogErrController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysLogErrController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysLogErrController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysLogErrController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysLogErrController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysLogErrController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:ParamSysParamController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:ParamSysParamController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:ParamSysParamController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:ParamSysParamController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:ParamSysParamController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:ParamSysParamController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:ParamSysParamController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:ParamSysParamController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:ParamSysParamController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:ParamSysParamController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMContactController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMContactController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMContactController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMContactController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMContactController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMContactController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMContactController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMContactController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMContactController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMContactController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMContractController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMContractController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMContractController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMContractController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMContractController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMContractController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMContractController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMContractController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMContractController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMContractController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrDepartmentController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrDepartmentController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrDepartmentController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrDepartmentController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrDepartmentController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrDepartmentController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrDepartmentController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrDepartmentController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrDepartmentController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrDepartmentController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMProductCategoryController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMProductCategoryController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMProductCategoryController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMProductCategoryController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMProductCategoryController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMProductCategoryController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMProductCategoryController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMProductCategoryController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMProductCategoryController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMProductCategoryController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysRoleController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysRoleController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysRoleController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysRoleController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysRoleController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysRoleController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysMenuController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysMenuController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysMenuController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysMenuController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysMenuController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysMenuController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysMenuController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysMenuController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysMenuController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysMenuController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMReceiveController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMReceiveController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMReceiveController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMReceiveController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMReceiveController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMReceiveController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMReceiveController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMReceiveController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMReceiveController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMReceiveController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrPostController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrPostController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrPostController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrPostController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrPostController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrPostController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrPostController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrPostController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrPostController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrPostController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMCustomerController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMCustomerController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMCustomerController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMCustomerController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMCustomerController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMCustomerController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMCustomerController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMCustomerController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMCustomerController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMCustomerController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PublicNewsController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PublicNewsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PublicNewsController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PublicNewsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PublicNewsController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PublicNewsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PublicNewsController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PublicNewsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PublicNewsController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PublicNewsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMProductController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMProductController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMProductController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMProductController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMProductController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMProductController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMProductController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMProductController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMProductController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMProductController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMInvoiceController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMInvoiceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMInvoiceController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMInvoiceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMInvoiceController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMInvoiceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMInvoiceController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMInvoiceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMInvoiceController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMInvoiceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysLogController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysLogController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysLogController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysLogController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysLogController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysLogController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysLogController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysLogController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysLogController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysLogController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysAppController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysAppController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysAppController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysAppController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysAppController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysAppController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysAppController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysAppController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysAppController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysAppController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMFollowController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMFollowController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMFollowController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMFollowController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMFollowController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMFollowController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMFollowController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMFollowController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:CRMFollowController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:CRMFollowController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrEmployeeController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrEmployeeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrEmployeeController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrEmployeeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrEmployeeController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrEmployeeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrEmployeeController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrEmployeeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:HrEmployeeController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:HrEmployeeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PersonalNotesController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PersonalNotesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PersonalNotesController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PersonalNotesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PersonalNotesController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PersonalNotesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PersonalNotesController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PersonalNotesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:PersonalNotesController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:PersonalNotesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysInfoController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysInfoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysInfoController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysInfoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysInfoController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysInfoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysInfoController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysInfoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["GoCRM/controllers:SysInfoController"] = append(beego.GlobalControllerRouter["GoCRM/controllers:SysInfoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
