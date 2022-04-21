package dbData

var businessPeriod = `SELECT "User"."uid", "User"."user_account", "User"."user_name", "User"."mobile_number", "User"."email", "User"."birth_day", "User"."area", "User"."department", "User"."position", "User"."office_phone_number", "User"."account_state", "User"."outer_account", "User"."outer_auth_info", "User"."group_code", "User"."reg_date", COALESCE("User"."upd_date", "User"."reg_date") AS "upd_date", "User"."is_temp_pwd", 
COALESCE(DATE_PART('day', now()- "User"."pwd_update_date")>=0, true) AS "is_expire", 
"User"."pwd_update_date", "User"."confirm_id", "User"."is_sms_agree", "Group"."group_code" AS "Group.group_code", "Group"."group_name" AS "Group.group_name", "Group"."enabled" AS "Group.enabled", "Group"."area" AS "Group.area", "Group"."reg_date" AS "Group.reg_date", "Group"."upd_date" AS "Group.upd_date", "Group"."group_type" AS "Group.group_type", "Group"."event_auto_check" AS "Group.event_auto_check", "Group"."layer_auto_check" AS "Group.layer_auto_check", "Code"."code" AS "Code.code", "Code"."p_code" AS "Code.p_code", "Code"."code_name" AS "Code.code_name", "Code"."memo" AS "Code.memo", "Code"."is_used" AS "Code.is_used", "Code"."reg_date" AS "Code.reg_date", "Code"."upd_date" AS "Code.upd_date", "Code"."is_modifiable" AS "Code.is_modifiable" 
FROM "user_info" AS "User" 
LEFT OUTER JOIN "group_info" AS "Group" 
ON "User"."group_code" = "Group"."group_code" 
LEFT OUTER JOIN "code_info" AS "Code" 
ON "User"."account_state" = "Code"."code" 
WHERE "User"."uid" = 'b7e6cf8a-465a-4ca9-b5b2-39fb7902ce04';`

// var test = ` INSERT INTO "user_info"
// (uid, user_account, user_name, mobile_number, birth_day, area, group_code,
// `

var test = `select * from `

// var FindGlobalProperty = `
// SELECT "prop_group", "prop_key", "prop_value", "prop_name", "etc", "upd_date" FROM "global_prop" AS "GlobalProp" WHERE "GlobalProp"."prop_group" IN ('SMS_INFO', 'MISSING_INFO', 'LIVESTOCK_INFO');
// `

// var FindValue = `
// SELECT "prop_group", "prop_key", "prop_value", "prop_name", "etc", "upd_date"
// FROM "global_prop" AS "GlobalProp"
// --WHERE "GlobalProp"."prop_group" IN ('SMS_INFO', 'MISSING_INFO', 'LIVESTOCK_INFO');
// WHERE "GlobalProp"."prop_group" IN ('SMS_INFO')
// AND "GlobalProp"."prop_key" = 'SMS_INFO_SMS_AUTH_URL';
// `

// var FindValue = `
// SELECT "prop_group", "prop_key", "prop_value", "prop_name", "etc", "upd_date"
// FROM "global_prop" AS "GlobalProp"
// WHERE "GlobalProp"."prop_group" IN ('SMS_INFO', 'MISSING_INFO', 'LIVESTOCK_INFO');
// --WHERE "GlobalProp"."prop_group" IN ('SMS_INFO')
// --AND "GlobalProp"."prop_key" = 'SMS_INFO_SMS_AUTH_URL';
// `

var FindGlobalProperty = `
SELECT "prop_key", "prop_value" 
FROM "global_prop" AS "GlobalProp" 
WHERE "GlobalProp"."prop_group" IN ('SMS_INFO', 'MISSING_INFO','LIVESTOCK_INFO');
`
