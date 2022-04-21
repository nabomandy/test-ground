package dbData

import (
	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /dms-demo/api/v1/getDemoList	01. 데이터 조회
 * @apiName getDemoList
 * @apiGroup DB
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
func GetDataList(c *gin.Context) {

	result, _, _ := getDataList(c)
	c.JSON(result.Code, result)

}
