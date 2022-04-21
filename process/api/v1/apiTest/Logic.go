package apiTest

import (
	"io"

	"dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/dao"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func getAPIData(c *gin.Context) (dao.Result, int, []interface{}) {

	var result dao.Result
	var resultResponse dao.Results
	var totalcount int = 0
	//var responseData interface{}
	var resultBody []interface{}
	var params map[string][]string = c.Request.URL.Query()

	if params["name"] == nil {
		result = dao.ResponseFunc(250, "invalid request format", resultResponse)
		return result, totalcount, resultBody

	}

	name := params["name"][0]
	resultBody = append(resultBody, map[string]interface{}{
		"name": name,
	})

	resultResponse = dao.Results{
		Results:    resultBody,
		TotalCount: totalcount,
	}
	result = dao.ResponseFunc(200, "sucess", resultResponse)

	return result, totalcount, resultBody
}

func getAPIData2(c *gin.Context) (dao.Result, int, []interface{}) {

	var result dao.Result
	var resultResponse dao.Results
	var totalcount int = 0
	//var responseData interface{}
	var resultBody []interface{}
	var params map[string][]string = c.Request.URL.Query()

	reader := c.Request.Body.(io.Reader)
	bodyString := dao.StreamToString(reader)
	zap.S().Infof("%v", bodyString)

	if params["name"] == nil {
		result = dao.ResponseFunc(250, "invalid request format", resultResponse)

	} else {

		name := params["name"][0]
		resultResponse = dao.Results{
			Results:    name,
			TotalCount: totalcount,
		}
		result = dao.ResponseFunc(200, "success", resultResponse)
	}

	return result, totalcount, resultBody
}

/**
 * @api {GET} /dms-demo/api/v1/getDemoList	01. 데이터 조회
 * @apiName getDemoList
 * @apiGroup API
 * @apiVersion  1.0.0
 * @apiDescription DB 데모 데이터 조회
 *
 * @apiSuccess {number} code 응답 코드(200:성공)
 * @apiSuccess {string} message 응답 메시지
 * @apiSuccess {string} responseTime 응답 완료 시간
 * @apiSuccess {object} response 응답 값
 * @apiSuccess {number} response.totalCount 조회 데이터 수
 * @apiSuccess {object} response.results 결과
 * @apiSuccess {object[]} response.results."test" 결과리스트
 * @apiSuccess {string} response.results."test".demo_id  고유아이디
 * @apiSuccess {string} response.results."test".demo_name  명칭
 *
 * @apiUse Success
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Success
 * {
 * 		"code": 200,
 *		"message": "success",
 *		"responseTime": "2021-03-12T13:56:43+09:00",
 *		"response": {
 *			"results": {
 *				"test": name
 *			},
 *			"totalCount": 0
 *		}
 * }
 *
 */
func postAPIData(c *gin.Context) (dao.Result, int, []interface{}) {

	var result dao.Result
	var resultResponse dao.Results
	var totalcount int = 0
	//var responseData interface{}
	var resultBody []interface{}
	var params map[string][]string = c.Request.URL.Query()

	reader := c.Request.Body.(io.Reader)
	bodyString := dao.StreamToString(reader)
	zap.S().Infof("%v", bodyString)

	if params["name"] == nil {
		result = dao.ResponseFunc(250, "invalid request format", resultResponse)

	} else {

		name := params["name"][0]
		resultResponse = dao.Results{
			Results:    name,
			TotalCount: totalcount,
		}
		result = dao.ResponseFunc(200, "success", resultResponse)
	}

	return result, totalcount, resultBody
}
