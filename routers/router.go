// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"CoCRM/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/crm",
		//主页
		beego.NSRouter("/index", &controllers.MainController{}),

		//联系人
		beego.NSNamespace("/contact",
			beego.NSRouter("/getall", &controllers.CRMContactController{}, "get:GetAll"),
			beego.NSRouter("/showAddpage", &controllers.CRMContactController{}, "get:ShowAddPage"),
			beego.NSRouter("/addAction", &controllers.CRMContactController{}, "post:Post"),
			beego.NSRouter("/delete", &controllers.CRMContactController{}, "post:Delete"),
		),

		beego.NSNamespace("/CRM_Customer",
			beego.NSInclude(
				&controllers.CRMCustomerController{},
			),
		),

		beego.NSNamespace("/CRM_Follow",
			beego.NSInclude(
				&controllers.CRMFollowController{},
			),
		),

		beego.NSNamespace("/CRM_contract",
			beego.NSInclude(
				&controllers.CRMContractController{},
			),
		),

		beego.NSNamespace("/CRM_invoice",
			beego.NSInclude(
				&controllers.CRMInvoiceController{},
			),
		),

		beego.NSNamespace("/CRM_order",
			beego.NSInclude(
				&controllers.CRMOrderController{},
			),
		),

		beego.NSNamespace("/CRM_product",
			beego.NSInclude(
				&controllers.CRMProductController{},
			),
		),

		beego.NSNamespace("/CRM_product_category",
			beego.NSInclude(
				&controllers.CRMProductCategoryController{},
			),
		),

		beego.NSNamespace("/CRM_receive",
			beego.NSInclude(
				&controllers.CRMReceiveController{},
			),
		),

		beego.NSNamespace("/Param_SysParam",
			beego.NSInclude(
				&controllers.ParamSysParamController{},
			),
		),

		beego.NSNamespace("/Personal_Calendar",
			beego.NSInclude(
				&controllers.PersonalCalendarController{},
			),
		),

		beego.NSNamespace("/Personal_notes",
			beego.NSInclude(
				&controllers.PersonalNotesController{},
			),
		),

		beego.NSNamespace("/Sys_App",
			beego.NSInclude(
				&controllers.SysAppController{},
			),
		),

		beego.NSNamespace("/Sys_Button",
			beego.NSInclude(
				&controllers.SysButtonController{},
			),
		),

		beego.NSNamespace("/Sys_Menu",
			beego.NSInclude(
				&controllers.SysMenuController{},
			),
		),

		beego.NSNamespace("/Sys_log",
			beego.NSInclude(
				&controllers.SysLogController{},
			),
		),

		beego.NSNamespace("/Sys_log_Err",
			beego.NSInclude(
				&controllers.SysLogErrController{},
			),
		),

		beego.NSNamespace("/Sys_role",
			beego.NSInclude(
				&controllers.SysRoleController{},
			),
		),

		beego.NSNamespace("/hr_department",
			beego.NSInclude(
				&controllers.HrDepartmentController{},
			),
		),

		beego.NSNamespace("/hr_employee",
			beego.NSInclude(
				&controllers.HrEmployeeController{},
			),
		),

		beego.NSNamespace("/hr_position",
			beego.NSInclude(
				&controllers.HrPositionController{},
			),
		),

		beego.NSNamespace("/hr_post",
			beego.NSInclude(
				&controllers.HrPostController{},
			),
		),

		beego.NSNamespace("/public_news",
			beego.NSInclude(
				&controllers.PublicNewsController{},
			),
		),

		beego.NSNamespace("/public_notice",
			beego.NSInclude(
				&controllers.PublicNoticeController{},
			),
		),

		beego.NSNamespace("/sys_info",
			beego.NSInclude(
				&controllers.SysInfoController{},
			),
		),

		beego.NSNamespace("/tool_batch",
			beego.NSInclude(
				&controllers.ToolBatchController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
