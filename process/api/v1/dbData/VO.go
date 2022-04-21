package dbData

import (
	"encoding/json"
	"time"

	"gopkg.in/guregu/null.v4"
)

var tableName string = `TEST `

type InputTest struct {
	Page  string `json:"page" gorm:"column:page" binding:"required,uuid"`
	Count string `json:"count" gorm:"column:count" binding:"required"`
}

type InputData struct {
	Page                string      `json:"page" gorm:"param_page"`
	Count               string      `json:"count" gorm:"param_count"`
	Account_state       string      `json:"account_state" gorm:"param_account_state"`
	Group_code          string      `json:"group_code" gorm:"param_group_code"`
	Search_keyword      string      `json:"search_keyword" gorm:"param_search_keyword"`
	Area                string      `json:"area" gorm:"param_area"`
	UID                 string      `json:"uid" gorm:"param_uid"`
	User_account        string      `json:"user_account" gorm:"param_user_account"`
	User_name           string      `json:"user_name" gorm:"param_user_name"`
	Mobile_number       string      `json:"mobile_number" gorm:"param_mobile_number"`
	Birth_day           string      `json:"birth_day" gorm:"param_birth_day"`
	Outer_auth_info     string      `json:"outer_auth_info" gorm:"param_outer_auth_info"`
	Email               string      `json:"email" gorm:"param_email"`
	Department          string      `json:"department" gorm:"param_department"`
	Office_phone_number string      `json:"office_phone_number" gorm:"param_office_phone_number"`
	Outer_account       string      `json:"outer_account" gorm:"param_outer_account"`
	Position            string      `json:"position" gorm:"param_position"`
	Password            string      `json:"password" gorm:"param_password"`
	Send_sms_password   interface{} `json:"send_sms_password" gorm:"end_sms_password"`
	Send_sms_password2  interface{} `json:"send_sms_password" gorm:"end_sms_password"`
	Send_sms_password3  interface{} `json:"send_sms_password" gorm:"end_sms_password"`
	Is_sms_agree        *bool       `json:"is_sms_agree" gorm:"is_sms_agree"`
	Is_sms_agree2       *bool       `json:"is_sms_agree2" gorm:"is_sms_agree"`
	Is_sms_agree3       null.Bool   `json:"is_sms_agree3" gorm:"is_sms_agree3"`
}

type User_Info struct {
	UID                 string           `json:"uid" gorm:"column:ew_uuid;type:uuid;primary_key;default:uuid_generate_v4()"`
	User_account        string           `json:"user_account" gorm:"param_user_account"`
	User_name           string           `json:"user_name" gorm:"param_user_name"`
	Mobile_number       string           `json:"mobile_number" gorm:"param_mobile_number"`
	Email               *string          `json:"email" gorm:"param_email"`
	Birth_day           string           `json:"birth_day" gorm:"param_birth_day"`
	Area                string           `json:"area" gorm:"param_area"`
	Department          *string          `json:"department" gorm:"param_department"`
	Position            *string          `json:"position" gorm:"param_position"`
	Office_phone_number *string          `json:"office_phone_number" gorm:"param_office_phone_number"`
	Account_state       string           `json:"account_state" gorm:"param_account_state"`
	Outer_account       bool             `json:"outer_account" gorm:"param_outer_account"`
	Outer_auth_info     *json.RawMessage `json:"outer_auth_info" gorm:"param_outer_auth_info"`
	Group_code          string           `json:"group_code" gorm:"param_group_code"`
	Reg_date            string           `json:"reg_date" gorm:"param_reg_date"`
	Upd_date            string           `json:"upd_date" gorm:"param_upd_date"`
	Is_temp_pwd         bool             `json:"is_temp_pwd" gorm:"param_is_temp_pwd"`
	Is_expire           string           `json:"is_expire" gorm:"param_is_expire"`
	Pwd_update_date     string           `json:"pwd_update_date" gorm:"param_pwd_update_date"`
	Confirm_id          *string          `json:"confirm_id" gorm:"param_confirm_id"`
	Is_sms_agree        bool             `json:"is_sms_agree" gorm:"param_is_sms_agree"`
}

type Group_Infos struct { // -> 클라이언트가 보낸 requestbody를 받는 구조체
	Group_code       string `json:"group_code" gorm:"param_group_code"`
	Group_name       string `json:"group_name" gorm:"param_group_name"`
	Enabled          string `json:"enabled" gorm:"param_enabled"`
	Area             string `json:"area" gorm:"param_area"`
	Reg_date         string `json:"reg_date" gorm:"param_reg_date"`
	Upd_date         string `json:"upd_date" gorm:"param_upd_date"`
	Group_type       string `json:"group_type" gorm:"param_group_type"`
	Event_auto_check string `json:"event_auto_check" gorm:"param_event_auto_check"`
	Layer_auto_check string `json:"layer_auto_check" gorm:"param_layer_auto_check"`
}

type Code_Infos struct { // -> 클라이언트가 보낸 requestbody를 받는 구조체
	Code          string `json:"code" gorm:"column:ew_uuid;type:uuid;primary_key;default:uuid_generate_v4()"`
	P_code        string `json:"p_code" gorm:"param_p_code"`
	Code_name     string `json:"code_name" gorm:"param_code_name"`
	Memo          string `json:"memo" gorm:"param_memo"`
	Is_used       bool   `json:"is_used" gorm:"param_is_used"`
	Reg_date      string `json:"reg_date" gorm:"param_reg_date"`
	Upd_date      string `json:"upd_date" gorm:"param_upd_date"`
	Is_modifiable string `json:"is_modifiable" gorm:"param_is_modifiable"`
}

type File_Infos struct { // -> 클라이언트가 보낸 requestbody를 받는 구조체
	File_id       *string `json:"file_id" gorm:"column:ew_uuid;type:uuid;primary_key;default:uuid_generate_v4()"`
	File_path     *string `json:"file_path" gorm:"param_file_path"`
	File_type     *string `json:"file_type" gorm:"param_file_type"`
	File_name     *string `json:"file_name" gorm:"param_file_name"`
	File_origname *string `json:"file_origname" gorm:"param_file_origname"`
	Reg_date      *string `json:"reg_date" gorm:"param_reg_date"`
	Ref_id        *string `json:"ref_id" gorm:"param_ref_id"`
	File_size     *string `json:"file_size" gorm:"param_file_size"`
	Ref_table     *string `json:"ref_table" gorm:"param_ref_table"`
	Ordering      *string `json:"ordering" gorm:"param_ordering"`
}

type Global_Property struct {
	Prop_group string    `json:"prop_group" gorm:"prop_group"`
	Prop_key   string    `json:"prop_key" gorm:"prop_key"`
	Prop_value string    `json:"prop_value" gorm:"prop_value"`
	Prop_name  string    `json:"prop_name" gorm:"prop_name"`
	Etc        string    `json:"etc" gorm:"etc"`
	Upd_date   time.Time `json:"upd_date" gorm:"upd_date"`
}

type Global_Prop struct {
	Prop_key   string      `json:"prop_key" gorm:"prop_key"`
	Prop_value interface{} `json:"prop_value" gorm:"prop_value"`
}
