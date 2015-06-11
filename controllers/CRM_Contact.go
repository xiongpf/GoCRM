package controllers

import (
	"GoCRM/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"fmt"
	"time"
)

// oprations for CRMContact
type CRMContactController struct {
	beego.Controller
}

func (c *CRMContactController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *CRMContactController) ShowAddPage() {
	c.TplNames = "contact/add.html"
}

// @Title Post
// @Description create CRMContact
// @Param	body		body 	models.CRMContact	true		"body for CRMContact content"
// @Success 200 {int} models.CRMContact.Id
// @Failure 403 body is empty
// @router / [post]
func (c *CRMContactController) Post() {
	ct := new(models.CRMContact)
	ct.CAdd = c.GetString("CAdd")
	ct.CName = c.GetString("CName")
	ct.CBirthday = c.GetString("CBirthday")
	ct.CCreateDate = time.Now()
	ct.CDepartment = c.GetString("CDepartment")
	ct.CEmail = c.GetString("CEmail")
	ct.CFax = c.GetString("CFax")
	ct.CHobby = c.GetString("CHobby")
	ct.CMob = c.GetString("CMob")
	ct.CQQ = c.GetString("CQQ")
	ct.CSex = c.GetString("CSex")
	ct.CTel = c.GetString("CTel")
    ct.CHobby = c.GetString("CHobby")
    ct.CRemarks = c.GetString("CRemarks")

	if id, err := models.AddCRMContact(ct); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
	//c.Redirect("/crm/contact/getall",200)
}

// @Title Get
// @Description get CRMContact by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.CRMContact
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CRMContactController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetCRMContactById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.TplNames = "contact/edit.html"
    c.Render()
}

// @Title Get All
// @Description get CRMContact
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.CRMContact
// @Failure 403
// @router / [get]
func (c *CRMContactController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("numPerPage"); err == nil {
		limit = v
		fmt.Println(limit)
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJson()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllCRMContact(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	//c.ServeJson()
	c.Data["totalCount"] = len(l)
	c.TplNames = "contact/index.html"
}

// @Title Update
// @Description update the CRMContact
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.CRMContact	true		"body for CRMContact content"
// @Success 200 {object} models.CRMContact
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CRMContactController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.CRMContact{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateCRMContactById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the CRMContact
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CRMContactController) Delete() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteCRMContact(id); err == nil {
		c.Data["json"] = map[string]string{"statusCode": "200", "message": "删除成功", "navTabId": "", "rel": "", "callbackType": "", "forwardUrl": "", "confirmMsg": ""}
	} else {
		c.Data["json"] = err.Error()
	}

	c.ServeJson()
}
