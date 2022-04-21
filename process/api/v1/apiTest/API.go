package apiTest

import "github.com/gin-gonic/gin"

/**
 * @api {GET} /dms-demo/api/v1/getAPIData	01. 데이터 조회
 * @apiName getAPIData
 * @apiGroup API
 * @apiVersion  1.0.0
 * @apiDescription GET_API
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
 *				"test": [
 *					{
 *						"demo_id": "0001",
 *						"demo_name": "첫번째"
 *
 *					},
 *					{
 *
 *						"demo_id": "0002",
 *						"demo_name": "두번째"
 *					}
 *				],
 *				"test2": []
 *			},
 *			"totalCount": 2
 *		}
 * }
 *
 */
func GetAPIData(c *gin.Context) {
	result, _, _ := getAPIData(c)
	c.JSON(result.Code, result)
}

/**
 * @api {GET} /dms-demo/api/v1/getAPIData2	02. 데이터 조회
 * @apiName getAPIData2
 * @apiGroup API
 * @apiVersion  1.0.0
 * @apiDescription GET_API
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
 *			"results": "0001" ,
 *			"totalCount": 2
 *		}
 * }
 *
 */
func GetAPIData2(c *gin.Context) {
	result, _, _ := getAPIData2(c)
	c.JSON(result.Code, result)
}

/**
 * @api {GET} /dms-demo/api/v1/postAPIData	03. 데이터 조회
 * @apiName postAPIData
 * @apiGroup API
 * @apiVersion  1.0.0
 * @apiDescription POST_TEST
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
 *				"test": [
 *					{
 *						"demo_id": "0001",
 *						"demo_name": "첫번째"
 *
 *					},
 *					{
 *
 *						"demo_id": "0002",
 *						"demo_name": "두번째"
 *					}
 *				],
 *				"test2": []
 *			},
 *			"totalCount": 2
 *		}
 * }
 *
 */
func PostAPIData(c *gin.Context) {
	result, _, _ := postAPIData(c)
	c.JSON(result.Code, result)
}
