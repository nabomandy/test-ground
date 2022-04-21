package dao

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type Result struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	ResponseTime string      `json:"responseTime"`
	Response     interface{} `json:"response"`
}

type Results struct {
	Results    interface{} `json:"results"`
	TotalCount int         `json:"totalCount"`
}

type PageResults struct {
	Results    interface{} `json:"results"`
	Count      int         `json:"count"`
	TotalCount int         `json:"totalCount"`
}

var CronScheduler *cron.Cron

func JsonResp(resp *resty.Response) (Result, error) {
	var jsonBody Result = Result{}
	err := json.Unmarshal(resp.Body(), &jsonBody)
	return jsonBody, err
}

func InterfaceToStrArr(src interface{}) []string {
	tempInterface := src.([]interface{})
	strArr := make([]string, len(tempInterface))
	for i, v := range tempInterface {
		strArr[i] = v.(string)
	}
	return strArr
}

func UnixTimestampToTimeString(unixtime string) string {
	i, err := strconv.ParseInt(unixtime, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	tmStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	return tmStr
}

func TimestampTimeZoneToString(tm time.Time) string {
	var tmStr string = ""
	if tm.Year() != 1 {
		tmStr = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	}
	return tmStr
}

func Contains(strSlice []string, str string) bool {
	for _, v := range strSlice {
		if v == str {
			return true
		}
	}
	return false
}

// const ErrorCode {
//         /**
//          * @apiDefine Success
//          */
//         Success = 200,
//         /**
//          * @apiDefine PartialSuccess
//          */
//         PartialSuccess = 206,

//         /**
//          * @apiDefine InvalidRequestFormat
//          * @apiErrorExample {json} Error-Response:
//          *  HTTP/1.1 250 Failure
//          *  {
//          *      "code": 250,
//          *      "message": "invalid request format",
//          *      "responseTime": "2019-11-26 12:51:32"
//          *  }
//          */
//         InvalidRequestFormat = 250,

//         /**
//          * @apiDefine InvalidValue
//          * @apiErrorExample {json} Error-Response:
//          *  HTTP/1.1 260 Failure
//          *  {
//          *      "code": 260,
//          *      "message": "invalid value",
//          *      "responseTime": "2019-11-26 12:51:32"
//          *  }
//          *
//          */
//         InvalidValue = 260,

//         /**
//          * @apiDefine InvalidFileType
//          * @apiErrorExample {json} Error-Response:
//          *  HTTP/1.1 261 Failure
//          *  {
//          *      "code": 261,
//          *      "message": "invalid file type",
//          *      "responseTime": "2019-11-26 12:51:32"
//          *  }
//          */
//         InvalidFileType = 261,

//         /**
//          * @apiDefine BadRequest
//          */
//         BadRequest = 400,
//         Unauthorized = 401,
//         PermissionDenined = 402,
//         /**
//          * @apiDefine NotAllowed
//          */
//         NotAllowed = 403,
//         NotFound = 404,
//         RequestTimeout = 408,

//         /**
//          * @apiDefine InternalServerError
//          * @apiErrorExample {json} Error-Response:
//          *  HTTP/1.1 500 Failure
//          *  {
//          *      "code": 500,
//          *      "message": "Internal Server Error",
//          *      "responseTime": "2018-03-19 11:28:06"
//          *  }
//          */
//         InternalServerError = 500
//     }

func ResponseFunc(code int, msg string, response Results) Result {
	zap.S().Infof(`Code[%v] Massage : %v`, code, msg)
	result := Result{
		Code:         code,
		Message:      msg,
		ResponseTime: time.Now().Format(time.RFC3339),
		Response:     response,
	}
	return result
}

func ResponseFunc2(code int, msg string, response Result) Result {
	zap.S().Infof(`Code[%v] Massage : %v`, code, msg)
	result := Result{
		Code:         code,
		Message:      msg,
		ResponseTime: time.Now().Format(time.RFC3339),
		Response:     response,
	}
	return result
}
