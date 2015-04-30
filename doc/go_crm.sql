/*
 Navicat MySQL Data Transfer

 Source Server         : localhost_root
 Source Server Version : 50617
 Source Host           : localhost
 Source Database       : go_crm

 Target Server Version : 50617
 File Encoding         : utf-8

 Date: 04/28/2015 10:09:34 AM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `CRM_Contact`
-- ----------------------------
DROP TABLE IF EXISTS `CRM_Contact`;
CREATE TABLE `CRM_Contact` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `C_name` varchar(250) DEFAULT NULL,
  `C_sex` varchar(250) DEFAULT NULL,
  `C_department` varchar(250) DEFAULT NULL,
  `C_position` varchar(250) DEFAULT NULL,
  `C_birthday` varchar(250) DEFAULT NULL,
  `C_tel` varchar(250) DEFAULT NULL,
  `C_fax` varchar(250) DEFAULT NULL,
  `C_email` varchar(250) DEFAULT NULL,
  `C_mob` varchar(250) DEFAULT NULL,
  `C_QQ` varchar(250) DEFAULT NULL,
  `C_add` varchar(250) DEFAULT NULL,
  `C_hobby` varchar(250) DEFAULT NULL,
  `C_remarks` varchar(250) DEFAULT NULL,
  `C_customerid` int(11) DEFAULT NULL,
  `C_customername` varchar(250) DEFAULT NULL,
  `C_createId` int(11) DEFAULT NULL,
  `C_createDate` datetime DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `CRM_Customer`
-- ----------------------------
DROP TABLE IF EXISTS `CRM_Customer`;
CREATE TABLE `CRM_Customer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Serialnumber` varchar(250) DEFAULT NULL,
  `Customer` varchar(250) DEFAULT NULL,
  `address` varchar(250) DEFAULT NULL,
  `tel` varchar(250) DEFAULT NULL,
  `fax` varchar(250) DEFAULT NULL,
  `site` varchar(250) DEFAULT NULL,
  `industry_id` int(11) DEFAULT NULL,
  `industry` varchar(250) DEFAULT NULL,
  `Provinces_id` int(11) DEFAULT NULL,
  `Provinces` varchar(250) DEFAULT NULL,
  `City_id` int(11) DEFAULT NULL,
  `City` varchar(250) DEFAULT NULL,
  `CustomerType_id` int(11) DEFAULT NULL,
  `CustomerType` varchar(250) DEFAULT NULL,
  `CustomerLevel_id` int(11) DEFAULT NULL,
  `CustomerLevel` varchar(250) DEFAULT NULL,
  `CustomerSource_id` int(11) DEFAULT NULL,
  `CustomerSource` varchar(250) DEFAULT NULL,
  `DesCripe` varchar(4000) DEFAULT NULL,
  `Remarks` varchar(4000) DEFAULT NULL,
  `Department_id` int(11) DEFAULT NULL,
  `Department` varchar(250) DEFAULT NULL,
  `Employee_id` int(11) DEFAULT NULL,
  `Employee` varchar(250) DEFAULT NULL,
  `privatecustomer` varchar(50) DEFAULT NULL,
  `lastfollow` datetime DEFAULT NULL,
  `Create_id` int(11) DEFAULT NULL,
  `Create_name` varchar(250) DEFAULT NULL,
  `Create_date` datetime DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `CRM_Follow`
-- ----------------------------
DROP TABLE IF EXISTS `CRM_Follow`;
CREATE TABLE `CRM_Follow` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Customer_id` int(11) DEFAULT NULL,
  `Customer_name` varchar(250) DEFAULT NULL,
  `Follow` varchar(250) DEFAULT NULL,
  `Follow_date` datetime DEFAULT NULL,
  `Follow_Type_id` int(11) DEFAULT NULL,
  `Follow_Type` varchar(250) DEFAULT NULL,
  `department_id` int(11) DEFAULT NULL,
  `department_name` varchar(250) DEFAULT NULL,
  `employee_id` int(11) DEFAULT NULL,
  `employee_name` varchar(250) DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `CRM_contract`
-- ----------------------------
DROP TABLE IF EXISTS `CRM_contract`;
CREATE TABLE `CRM_contract` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Contract_name` varchar(250) DEFAULT NULL,
  `Serialnumber` varchar(250) DEFAULT NULL,
  `Customer_id` int(11) DEFAULT NULL,
  `Customer_name` varchar(250) DEFAULT NULL,
  `C_depid` int(11) DEFAULT NULL,
  `C_depname` varchar(250) DEFAULT NULL,
  `C_empid` int(11) DEFAULT NULL,
  `C_empname` varchar(250) DEFAULT NULL,
  `Contract_amount` float DEFAULT NULL,
  `Pay_cycle` int(11) DEFAULT NULL,
  `Start_date` varchar(250) DEFAULT NULL,
  `End_date` varchar(250) DEFAULT NULL,
  `Sign_date` varchar(250) DEFAULT NULL,
  `Customer_Contractor` varchar(250) DEFAULT NULL,
  `Our_Contractor_depid` int(11) DEFAULT NULL,
  `Our_Contractor_depname` varchar(250) DEFAULT NULL,
  `Our_Contractor_id` int(11) DEFAULT NULL,
  `Our_Contractor_name` varchar(250) DEFAULT NULL,
  `Creater_id` int(11) DEFAULT NULL,
  `Creater_name` varchar(250) DEFAULT NULL,
  `Create_time` datetime DEFAULT NULL,
  `Main_Content` varchar(250) DEFAULT NULL,
  `Remarks` varchar(250) DEFAULT NULL,
  `File_serialnumber` varchar(250) DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `CRM_contract_attachment`
-- ----------------------------
DROP TABLE IF EXISTS `CRM_contract_attachment`;
CREATE TABLE `CRM_contract_attachment` (
  `contract_id` int(11) DEFAULT NULL,
  `page_id` varchar(250) DEFAULT NULL,
  `file_id` varchar(250) DEFAULT NULL,
  `file_name` varchar(250) DEFAULT NULL,
  `real_name` varchar(250) DEFAULT NULL,
  `file_size` int(11) DEFAULT NULL,
  `create_id` int(11) DEFAULT NULL,
  `create_name` varchar(250) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `CRM_invoice`
-- ----------------------------
DROP TABLE IF EXISTS `CRM_invoice`;
CREATE TABLE `CRM_invoice` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Customer_id` int(11) DEFAULT NULL,
  `Customer_name` varchar(250) DEFAULT NULL,
  `invoice_num` varchar(250) DEFAULT NULL,
  `invoice_type_id` int(11) DEFAULT NULL,
  `invoice_type` varchar(250) DEFAULT NULL,
  `invoice_amount` float DEFAULT NULL,
  `invoice_content` varchar(250) DEFAULT NULL,
  `invoice_date` datetime DEFAULT NULL,
  `C_depid` int(11) DEFAULT NULL,
  `C_depname` varchar(250) DEFAULT NULL,
  `C_empid` int(11) DEFAULT NULL,
  `C_empname` varchar(250) DEFAULT NULL,
  `create_id` int(11) DEFAULT NULL,
  `create_name` varchar(250) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  `order_id` int(11) DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `CRM_order`
-- ----------------------------
DROP TABLE IF EXISTS `CRM_order`;
CREATE TABLE `CRM_order` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Serialnumber` varchar(250) DEFAULT NULL,
  `Customer_id` int(11) DEFAULT NULL,
  `Customer_name` varchar(250) DEFAULT NULL,
  `Order_date` datetime DEFAULT NULL,
  `pay_type_id` int(11) DEFAULT NULL,
  `pay_type` varchar(250) DEFAULT NULL,
  `Order_details` varchar(250) DEFAULT NULL,
  `Order_status_id` int(11) DEFAULT NULL,
  `Order_status` varchar(250) DEFAULT NULL,
  `Order_amount` float DEFAULT NULL,
  `create_id` int(11) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  `C_dep_id` int(11) DEFAULT NULL,
  `C_dep_name` varchar(250) DEFAULT NULL,
  `C_emp_id` int(11) DEFAULT NULL,
  `C_emp_name` varchar(250) DEFAULT NULL,
  `F_dep_id` int(11) DEFAULT NULL,
  `F_dep_name` varchar(250) DEFAULT NULL,
  `F_emp_id` int(11) DEFAULT NULL,
  `F_emp_name` varchar(250) DEFAULT NULL,
  `receive_money` float DEFAULT NULL,
  `arrears_money` float DEFAULT NULL,
  `invoice_money` float DEFAULT NULL,
  `arrears_invoice` float DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `CRM_order_details`
-- ----------------------------
DROP TABLE IF EXISTS `CRM_order_details`;
CREATE TABLE `CRM_order_details` (
  `order_id` int(11) DEFAULT NULL,
  `product_id` int(11) DEFAULT NULL,
  `product_name` varchar(250) DEFAULT NULL,
  `price` float DEFAULT NULL,
  `quantity` int(11) DEFAULT NULL,
  `unit` varchar(250) DEFAULT NULL,
  `amount` float DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `CRM_product`
-- ----------------------------
DROP TABLE IF EXISTS `CRM_product`;
CREATE TABLE `CRM_product` (
  `product_id` int(11) NOT NULL AUTO_INCREMENT,
  `product_name` varchar(250) DEFAULT NULL,
  `category_id` int(11) DEFAULT NULL,
  `category_name` varchar(250) DEFAULT NULL,
  `specifications` varchar(250) DEFAULT NULL,
  `status` varchar(250) DEFAULT NULL,
  `unit` varchar(250) DEFAULT NULL,
  `remarks` varchar(250) DEFAULT NULL,
  `price` float DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `CRM_product_category`
-- ----------------------------
DROP TABLE IF EXISTS `CRM_product_category`;
CREATE TABLE `CRM_product_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_category` varchar(250) DEFAULT NULL,
  `parentid` int(11) DEFAULT NULL,
  `product_icon` varchar(250) DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_id` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `CRM_receive`
-- ----------------------------
DROP TABLE IF EXISTS `CRM_receive`;
CREATE TABLE `CRM_receive` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Customer_id` int(11) DEFAULT NULL,
  `Customer_name` varchar(250) DEFAULT NULL,
  `Receive_num` varchar(250) DEFAULT NULL,
  `Pay_type_id` int(11) DEFAULT NULL,
  `Pay_type` varchar(250) DEFAULT NULL,
  `Receive_amount` float DEFAULT NULL,
  `Receive_date` datetime DEFAULT NULL,
  `C_depid` int(11) DEFAULT NULL,
  `C_depname` varchar(250) DEFAULT NULL,
  `C_empid` int(11) DEFAULT NULL,
  `C_empname` varchar(250) DEFAULT NULL,
  `create_id` int(11) DEFAULT NULL,
  `create_name` varchar(250) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  `companyid` int(11) DEFAULT NULL,
  `order_id` int(11) DEFAULT NULL,
  `remarks` varchar(250) DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  `receive_direction_id` int(11) DEFAULT NULL,
  `receive_direction_name` varchar(250) DEFAULT NULL,
  `receive_type_id` int(11) DEFAULT NULL,
  `receive_type_name` varchar(250) DEFAULT NULL,
  `receive_real` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Param_City`
-- ----------------------------
DROP TABLE IF EXISTS `Param_City`;
CREATE TABLE `Param_City` (
  `id` int(11) NOT NULL,
  `parentid` int(11) DEFAULT NULL,
  `City` varchar(250) DEFAULT NULL,
  `City_order` int(11) DEFAULT NULL,
  `Create_id` int(11) DEFAULT NULL,
  `Create_date` datetime DEFAULT NULL,
  `Update_id` int(11) DEFAULT NULL,
  `Update_date` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Param_SysParam`
-- ----------------------------
DROP TABLE IF EXISTS `Param_SysParam`;
CREATE TABLE `Param_SysParam` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parentid` int(11) DEFAULT NULL,
  `params_name` varchar(250) DEFAULT NULL,
  `params_order` int(11) DEFAULT NULL,
  `Create_id` int(11) DEFAULT NULL,
  `Create_date` datetime DEFAULT NULL,
  `Update_id` int(11) DEFAULT NULL,
  `Update_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Param_SysParam_Type`
-- ----------------------------
DROP TABLE IF EXISTS `Param_SysParam_Type`;
CREATE TABLE `Param_SysParam_Type` (
  `id` int(11) NOT NULL,
  `params_name` varchar(250) DEFAULT NULL,
  `params_order` int(11) DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Personal_Calendar`
-- ----------------------------
DROP TABLE IF EXISTS `Personal_Calendar`;
CREATE TABLE `Personal_Calendar` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `emp_id` int(11) DEFAULT NULL,
  `emp_name` varchar(250) DEFAULT NULL,
  `companyid` int(11) DEFAULT NULL,
  `Subject` varchar(250) DEFAULT NULL,
  `Location` varchar(250) DEFAULT NULL,
  `MasterId` int(11) DEFAULT NULL,
  `Description` varchar(250) DEFAULT NULL,
  `CalendarType` tinyint(4) DEFAULT NULL,
  `StartTime` datetime DEFAULT NULL,
  `EndTime` datetime DEFAULT NULL,
  `IsAllDayEvent` bit(1) DEFAULT NULL,
  `HasAttachment` bit(1) DEFAULT NULL,
  `Category` varchar(250) DEFAULT NULL,
  `InstanceType` tinyint(4) DEFAULT NULL,
  `Attendees` varchar(250) DEFAULT NULL,
  `AttendeeNames` varchar(250) DEFAULT NULL,
  `OtherAttendee` varchar(250) DEFAULT NULL,
  `UPAccount` varchar(250) DEFAULT NULL,
  `UPName` varchar(250) DEFAULT NULL,
  `UPTime` datetime DEFAULT NULL,
  `RecurringRule` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Personal_notes`
-- ----------------------------
DROP TABLE IF EXISTS `Personal_notes`;
CREATE TABLE `Personal_notes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `emp_id` int(11) DEFAULT NULL,
  `emp_name` varchar(250) DEFAULT NULL,
  `note_content` varchar(250) DEFAULT NULL,
  `note_color` varchar(250) DEFAULT NULL,
  `xyz` varchar(250) DEFAULT NULL,
  `note_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Sys_App`
-- ----------------------------
DROP TABLE IF EXISTS `Sys_App`;
CREATE TABLE `Sys_App` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `App_name` varchar(100) DEFAULT NULL,
  `App_order` int(11) DEFAULT NULL,
  `App_url` varchar(250) DEFAULT NULL,
  `App_handler` varchar(250) DEFAULT NULL,
  `App_type` varchar(50) DEFAULT NULL,
  `App_icon` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Sys_Button`
-- ----------------------------
DROP TABLE IF EXISTS `Sys_Button`;
CREATE TABLE `Sys_Button` (
  `Btn_id` int(11) NOT NULL AUTO_INCREMENT,
  `Btn_name` varchar(255) DEFAULT NULL,
  `Btn_icon` varchar(50) DEFAULT NULL,
  `Btn_handler` varchar(255) DEFAULT NULL,
  `Menu_id` int(11) DEFAULT NULL,
  `Menu_name` varchar(255) DEFAULT NULL,
  `Btn_order` int(11) DEFAULT NULL,
  PRIMARY KEY (`Btn_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Sys_Menu`
-- ----------------------------
DROP TABLE IF EXISTS `Sys_Menu`;
CREATE TABLE `Sys_Menu` (
  `Menu_id` int(11) NOT NULL AUTO_INCREMENT,
  `Menu_name` varchar(255) DEFAULT NULL,
  `parentid` int(11) DEFAULT NULL,
  `parentname` varchar(255) DEFAULT NULL,
  `App_id` int(11) DEFAULT NULL,
  `Menu_url` varchar(255) DEFAULT NULL,
  `Menu_icon` varchar(50) DEFAULT NULL,
  `Menu_handler` varchar(50) DEFAULT NULL,
  `Menu_order` int(11) DEFAULT NULL,
  `Menu_type` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`Menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Sys_authority`
-- ----------------------------
DROP TABLE IF EXISTS `Sys_authority`;
CREATE TABLE `Sys_authority` (
  `Role_id` int(11) NOT NULL,
  `App_ids` varchar(250) DEFAULT NULL,
  `Menu_ids` varchar(250) DEFAULT NULL,
  `Button_ids` varchar(250) DEFAULT NULL,
  `Create_id` int(11) DEFAULT NULL,
  `Create_date` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Sys_data_authority`
-- ----------------------------
DROP TABLE IF EXISTS `Sys_data_authority`;
CREATE TABLE `Sys_data_authority` (
  `Role_id` int(11) DEFAULT NULL,
  `option_id` int(11) DEFAULT NULL,
  `Sys_option` varchar(250) DEFAULT NULL,
  `Sys_view` int(11) DEFAULT NULL,
  `Sys_add` int(11) DEFAULT NULL,
  `Sys_edit` int(11) DEFAULT NULL,
  `Sys_del` int(11) DEFAULT NULL,
  `Create_id` int(11) DEFAULT NULL,
  `Create_date` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Sys_log`
-- ----------------------------
DROP TABLE IF EXISTS `Sys_log`;
CREATE TABLE `Sys_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `EventType` varchar(250) DEFAULT NULL,
  `EventID` varchar(50) DEFAULT NULL,
  `EventTitle` varchar(250) DEFAULT NULL,
  `Original_txt` varchar(4000) DEFAULT NULL,
  `Current_txt` varchar(4000) DEFAULT NULL,
  `UserID` int(11) DEFAULT NULL,
  `UserName` varchar(50) DEFAULT NULL,
  `IPStreet` varchar(50) DEFAULT NULL,
  `EventDate` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Sys_log_Err`
-- ----------------------------
DROP TABLE IF EXISTS `Sys_log_Err`;
CREATE TABLE `Sys_log_Err` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Err_typeid` int(11) DEFAULT NULL,
  `Err_type` varchar(250) DEFAULT NULL,
  `Err_time` datetime DEFAULT NULL,
  `Err_url` varchar(500) DEFAULT NULL,
  `Err_message` varchar(250) DEFAULT NULL,
  `Err_source` varchar(500) DEFAULT NULL,
  `Err_trace` varchar(250) DEFAULT NULL,
  `Err_emp_id` int(11) DEFAULT NULL,
  `Err_emp_name` varchar(250) DEFAULT NULL,
  `Err_ip` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Sys_online`
-- ----------------------------
DROP TABLE IF EXISTS `Sys_online`;
CREATE TABLE `Sys_online` (
  `UserID` int(11) DEFAULT NULL,
  `UserName` varchar(50) DEFAULT NULL,
  `LastLogTime` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `Sys_role`
-- ----------------------------
DROP TABLE IF EXISTS `Sys_role`;
CREATE TABLE `Sys_role` (
  `RoleID` int(11) NOT NULL AUTO_INCREMENT,
  `RoleName` varchar(255) DEFAULT NULL,
  `RoleDscript` varchar(255) DEFAULT NULL,
  `RoleSort` int(11) DEFAULT NULL,
  `CreateID` int(11) DEFAULT NULL,
  `CreateDate` datetime DEFAULT NULL,
  `UpdateID` int(11) DEFAULT NULL,
  `UpdateDate` datetime DEFAULT NULL,
  PRIMARY KEY (`RoleID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `hr_department`
-- ----------------------------
DROP TABLE IF EXISTS `hr_department`;
CREATE TABLE `hr_department` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `d_name` varchar(50) DEFAULT NULL,
  `parentid` int(11) DEFAULT NULL,
  `parentname` varchar(50) DEFAULT NULL,
  `d_type` varchar(50) DEFAULT NULL,
  `d_icon` varchar(50) DEFAULT NULL,
  `d_fuzeren` varchar(50) DEFAULT NULL,
  `d_tel` varchar(50) DEFAULT NULL,
  `d_fax` varchar(50) DEFAULT NULL,
  `d_add` varchar(255) DEFAULT NULL,
  `d_email` varchar(50) DEFAULT NULL,
  `d_miaoshu` varchar(255) DEFAULT NULL,
  `d_order` int(11) DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `hr_employee`
-- ----------------------------
DROP TABLE IF EXISTS `hr_employee`;
CREATE TABLE `hr_employee` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` varchar(50) DEFAULT NULL,
  `pwd` varchar(50) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `idcard` varchar(50) DEFAULT NULL,
  `birthday` varchar(50) DEFAULT NULL,
  `d_id` int(11) DEFAULT NULL,
  `dname` varchar(50) DEFAULT NULL,
  `postid` int(11) DEFAULT NULL,
  `post` varchar(250) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `sex` varchar(50) DEFAULT NULL,
  `tel` varchar(50) DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL,
  `zhiwuid` int(11) DEFAULT NULL,
  `zhiwu` varchar(50) DEFAULT NULL,
  `sort` int(11) DEFAULT NULL,
  `EntryDate` varchar(50) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `remarks` varchar(255) DEFAULT NULL,
  `education` varchar(50) DEFAULT NULL,
  `level` varchar(50) DEFAULT NULL,
  `professional` varchar(50) DEFAULT NULL,
  `schools` varchar(50) DEFAULT NULL,
  `title` varchar(50) DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  `portal` varchar(250) DEFAULT NULL,
  `theme` varchar(250) DEFAULT NULL,
  `canlogin` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `hr_position`
-- ----------------------------
DROP TABLE IF EXISTS `hr_position`;
CREATE TABLE `hr_position` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `position_name` varchar(250) DEFAULT NULL,
  `position_order` int(11) DEFAULT NULL,
  `position_level` varchar(50) DEFAULT NULL,
  `create_id` int(11) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `hr_post`
-- ----------------------------
DROP TABLE IF EXISTS `hr_post`;
CREATE TABLE `hr_post` (
  `post_id` int(11) NOT NULL AUTO_INCREMENT,
  `post_name` varchar(255) DEFAULT NULL,
  `position_id` int(11) DEFAULT NULL,
  `position_name` varchar(255) DEFAULT NULL,
  `position_order` int(11) DEFAULT NULL,
  `dep_id` int(11) DEFAULT NULL,
  `depname` varchar(255) DEFAULT NULL,
  `emp_id` int(11) DEFAULT NULL,
  `emp_name` varchar(255) DEFAULT NULL,
  `default_post` int(11) DEFAULT NULL,
  `note` varchar(250) DEFAULT NULL,
  `post_descript` varchar(250) DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `public_news`
-- ----------------------------
DROP TABLE IF EXISTS `public_news`;
CREATE TABLE `public_news` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `news_title` varchar(250) DEFAULT NULL,
  `news_content` varchar(250) DEFAULT NULL,
  `create_id` int(11) DEFAULT NULL,
  `create_name` varchar(250) DEFAULT NULL,
  `dep_id` int(11) DEFAULT NULL,
  `dep_name` varchar(250) DEFAULT NULL,
  `news_time` datetime DEFAULT NULL,
  `isDelete` int(11) DEFAULT NULL,
  `Delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `public_notice`
-- ----------------------------
DROP TABLE IF EXISTS `public_notice`;
CREATE TABLE `public_notice` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `notice_title` varchar(250) DEFAULT NULL,
  `notice_content` varchar(250) DEFAULT NULL,
  `create_id` int(11) DEFAULT NULL,
  `create_name` varchar(250) DEFAULT NULL,
  `dep_id` int(11) DEFAULT NULL,
  `dep_name` varchar(250) DEFAULT NULL,
  `notice_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `sys_info`
-- ----------------------------
DROP TABLE IF EXISTS `sys_info`;
CREATE TABLE `sys_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sys_key` varchar(50) DEFAULT NULL,
  `sys_value` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `sys_role_emp`
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_emp`;
CREATE TABLE `sys_role_emp` (
  `RoleID` int(11) NOT NULL,
  `empID` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `tool_batch`
-- ----------------------------
DROP TABLE IF EXISTS `tool_batch`;
CREATE TABLE `tool_batch` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `batch_type` varchar(50) DEFAULT NULL,
  `o_dep_id` int(11) DEFAULT NULL,
  `o_dep` varchar(250) DEFAULT NULL,
  `o_emp_id` int(11) DEFAULT NULL,
  `o_emp` varchar(250) DEFAULT NULL,
  `c_dep_id` int(11) DEFAULT NULL,
  `c_dep` varchar(250) DEFAULT NULL,
  `c_emp_id` int(11) DEFAULT NULL,
  `c_emp` varchar(250) DEFAULT NULL,
  `b_count` int(11) DEFAULT NULL,
  `create_id` int(11) DEFAULT NULL,
  `create_name` varchar(250) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
