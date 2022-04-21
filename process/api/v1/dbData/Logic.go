package dbData

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"

	// "strconv"
	"strings"
	"time"

	"dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/config"
	dao "dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/dao"
	"github.com/Luzifer/go-openssl"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/parjom/goutil"
	"go.uber.org/zap"
)

type RegistrationFailure struct {
	User_account        string `json:"사용자 계정(필수)  gorm:"column:사용자 계정(필수)"`
	User_name           string `json:"사용자 이름(필수)" gorm:"column:사용자 이름(필수)"`
	Mobile_number       string `json:"사용자 휴대전화번호(필수)" gorm:"column:사용자 휴대전화번호(필수)"`
	Email               string `json:"이메일 주소" gorm:"이메일 주소"`
	Birth_day           string `json:"생년월일(필수)" gorm:"생년월일(필수)"`
	Department          string `json:"부서" gorm:"부서"`
	Position            string `json:"직급" gorm:"직급"`
	Office_phone_number string `json:"사무실 전화 번호" gorm:"사무실 전화 번호"`
	Group_name          string `json:"그룹명(필수)" gorm:"그룹명(필수)"`
	Password            string `json:"비밀번호" gorm:"비밀번호"`
	FailMessage         string `json:"FailMessage" gorm:"FailMessage"`
}

type CustomFloatUnmarshal struct {
	Value   string
	Percent float64
}

// func (c *CustomFloatUnmarshal) UnmarshalXLS(value string) error {
// 	c.Value = value
// 	f, err := toFloat(value)
// 	if err != nil {
// 		return err
// 	}
// 	c.Percent = f
// 	return nil
// }

const simplefile = "./test/Test1.xlsx"
const sheet = "Sheet 2"

type test1 []*RegistrationFailure

func (tests test1) String() string {
	s := "["
	for i, test := range tests {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%v", test)
	}
	return s + "]"
}

type BaseData struct {
	Id     string      `json:"id"`
	Pw     string      `json:"pw"`
	Detail *DetailData `json:"flag,omitempty"`
}

type DetailData struct {
	Flag bool `json:"value"`
}

func getDataList(c *gin.Context) (dao.Result, int, []interface{}) {

	var result dao.Result
	var resultResponse dao.Results
	var totalcount int = 0
	//var responseData interface{}
	var resultBody []interface{}

	var dmsDBConn *gorm.DB
	// var totalCnt int = 0
	var err error = nil
	// var rows *sql.Rows = nil
	var inputData InputTest
	// var inputData InputData
	err = c.ShouldBind(&inputData) // 타입이 안맞거나, 값이 없으면 걸린다.
	if err != nil {
		result = dao.ResponseFunc(200, err.Error(), resultResponse)
		return result, totalcount, resultBody
	}

	// 키 값만 소문자만들기 시작
	lck := "CN=056박범섭001,ou=people,ou=경기도,o=Government of Korea,c=kr"
	// fmt.Println(lck)
	t := strings.Split(lck, ",")
	// fmt.Println(t)
	t[0] = strings.Replace(t[0], "cn=", "", -1)
	// test123 := strings.Split(t, "=")
	// fmt.Println(test123)

	t1oner := []string{}

	for i, _ := range t {
		t1 := strings.Split(t[i], "=") // 나누고
		// t1[0] = strings.ToLower(t1[0])
		t[i] = strings.ReplaceAll(t[i], t1[0], strings.ToLower(t1[0])) // 소문자로 바꿔주고
		fmt.Println(t[i])
		// 넣어줄 차례
		t1oner = append(t1oner, t[i])
	}

	fmt.Println(t1oner)

	t1Zeus := strings.Join(t1oner[:], ",")
	fmt.Println(t1Zeus)
	// 키 값만 소문자만들기 끝

	// if reflect.TypeOf(inputData.Outer_auth_info).Kind() == reflect.Array {
	// 	fmt.Println("되네")
	// }

	testquery := []string{"test", "test", "test"}
	testquery = append(testquery, "user_info")

	testquery = append(testquery, "findone")
	testquery = append(testquery, "findall")
	fmt.Println(testquery)
	fmt.Println(strings.Join(testquery[:], ","))

	// //json to string
	// outer_auth, _ := json.Marshal(inputData.Outer_auth_info) //
	// fmt.Println("doc:", string(outer_auth))
	// outer_auth_info := string(outer_auth)
	// fmt.Println("outer_auth_info:", outer_auth_info)
	// fmt.Println(reflect.TypeOf(outer_auth_info))

	var querytesting = test

	fmt.Println("(" + "test" + "," + "test")

	type BaseData struct {
		Id     string      `json:"id"`
		Pw     string      `json:"pw"`
		Detail *DetailData `json:"flag,omitempty"`
	}

	type DetailData struct {
		Flag bool `json:"value"`
	}

	data2 := "{\"id\": \"test\", \"pw\": \"pass\", \"flag\": { \"value\":true}}"
	// data := "{\"id\": \"test\", \"pw\": \"pass\", \"flag\": { \"value\":false}}"
	// data := "{\"id\": \"test\", \"pw\": \"pass\"}"

	var baseData BaseData
	json.Unmarshal([]byte(data2), &baseData)
	fmt.Println(baseData)
	if baseData.Detail != nil {
		fmt.Println(baseData.Detail.Flag)
	}
	////////////////

	// if *inputData.Is_sms_agree2 == true {
	// 	fmt.Println("starbucks")
	// }

	// fmt.Println(cast.ToBool(inputData.Send_sms_password))

	fmt.Println(inputData.Page)
	// json.rawmassege를 string으로 변환
	// doc, _ := json.Marshal(inputData.Outer_auth_info) // 맵을 JSON 문서로 변환
	// fmt.Println("doc:", string(doc))
	// fmt.Println(reflect.TypeOf(doc))
	// json.rawmassege를 string으로 변환 끝

	var dst2 interface{}
	// var hello string
	// hello = `{"id": "admin","url": "http://www.naver.com"}`
	// json.Unmarshal(inputData.Outer_auth_info, &dst2)
	fmt.Println(dst2)

	// test33333 := strconv.FormatBool(inputData.Is_sms_agree)
	// test12345 := strconv.FormatBool(inputData.Send_sms_password)

	// // fmt.Println("test12345", test12345)
	// if inputData.Outer_auth_info == nil {
	// 	fmt.Println("ok")
	// }

	// fmt.Println(inputData.Outer_auth_info)
	// fmt.Println(reflect.TypeOf(inputData.Outer_auth_info))
	fmt.Println(reflect.TypeOf(dst2))
	// fmt.Println(test33333)
	//대문자 소문자로 변환
	var str string = "ABCD,EF=Ghijklmn"
	fmt.Println(strings.ToLower(str))

	//encDN 목표값 : rppsEOKCKwUrmkDhLyKk/p8VftjuQQe/kA2AIliH7/N8lbPViuakM3oplW40v8vecG3n0UDyLR2rPmKeFmmjlrmYPGTJ/wEqutObxIYeQ9dWbTpDGnwRKVzEzj82kSU03yM=:&&

	// 암호화 버전
	euc_key := config.GetInstance().EUC_KEY // 시크릿키는 env파일에 설정해서, config로 받아오고
	signedData := "12ec4e47db2ad0e79de62cfc2bdc406a2c0913ed933f0b5979a4bd69f50a54a6"

	o := openssl.New()

	enc, err := o.EncryptString(euc_key, signedData)

	// enc, err := o.EncryptString(euc_key, signedData)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}

	var DecryptionResult string = string(enc)
	var ParseDecryption map[string]string
	json.Unmarshal([]byte(DecryptionResult), &ParseDecryption)
	group_code := ParseDecryption["group_code"]
	fmt.Println(group_code)

	// 암호화 버전 끝

	// //복호화 버전
	// euc_key := config.GetInstance().EUC_KEY // 시크릿키는 env파일에 설정해서, config로 받아오고

	// signedData := "U2FsdGVkX18Kf9lRaxi/vxabByTOQ1DiOAKVJSJD9+swO6p8OB8SX3oWqgsq9ewZEgTj0uydcliqJVAANrpwDJj/M+pII/7NSq+FXc5YuHKU4aJVtcBxLy+oqzk25v9ZPkCmz9B+pissnakxUbGvSglyyXeIsByNrw1OEj3/NAtUNc+ZgYgII5DCJxTW6n2eSnnw328b4OfaFgAg2IDJuRLachcNwbYCgTk3pcwjWsLhS3AWvToHVkyNemtYhAuS"
	// loginfo_string := signedData
	// o := openssl.New()
	// dec, err := o.DecryptBytes(euc_key, []byte(loginfo_string))
	// if err != nil {
	// 	result = dao.ResponseFunc(http.StatusBadRequest, "invalid Request Format", resultResponse) // 어떤 에러코드로 작성할지 JWJ선임님께 문의
	// }
	// // 파싱완료된 로그인포 변수 선언
	// var DecryptionResult string = string(dec)
	// var ParseDecryption map[string]string
	// json.Unmarshal([]byte(DecryptionResult), &ParseDecryption)
	// group_code := ParseDecryption["group_code"]
	// fmt.Println(group_code)
	// // 복호화 버전 끝

	strtmp := "cn=김혜원()008804520190315188002010,ou=SHB,ou=personal4IB,o=yessign,c=kr"
	slice := strings.Split(strtmp, ",")

	for _, str := range slice {
		fmt.Println(str)
	}

	//HmacSHA256
	secret := "7ad99acc28f54f9d9aa68d571e249edc"
	data := "cn=김혜원()008804520190315188002010,ou=SHB,ou=personal4IB,o=yessign,c=kr"
	// data := "12ec4e47db2ad0e79de62cfc2bdc406a2c0913ed933f0b5979a4bd69f50a54a6"

	fmt.Printf("Secret: %s Data: %s\n", secret, data)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	fmt.Println("Result: " + sha)

	////HmacSHA256 끝

	query := businessPeriod

	// test := inputData.Outer_auth_info

	fmt.Println(test)

	var src []interface{}
	// json.Unmarshal([]byte(inputData.Outer_auth_info), &src)
	fmt.Println(src) // 꺼냈고
	// fmt.Println(event_subtype[1]) // 꺼냈고

	res := goutil.JsonGetValue(src, "event_subtype")

	fmt.Println(res)

	fmt.Println(res)

	// type Codes struct {
	// 	one string `json:"test"`  //
	// 	two string `json:"test2"` //
	// }

	// type AuthTokenClaims struct {
	// 	TokenUUID          string   `json:"tid"`  // 토큰 UUID
	// 	UserID             string   `json:"id"`   // 유저 ID
	// 	Name               string   `json:"name"` // 유저 이름
	// 	Email              string   `json:"mail"` // 유저 메일
	// 	UserUUID           string   `json:"uid"`  // 유저 UUID
	// 	Role               []inter `json:"role"` // 유저 역할
	// 	Codes              []Codes  `json:"Code"` //
	// 	jwt.StandardClaims          // 표준 토큰 Claims
	// }

	// at := AuthTokenClaims{
	// 	TokenUUID: uuid.NewString(),
	// 	UserID:    "vompressor",
	// 	// uuid 참고 // https://www.vompressor.com/uuid-go/
	// 	UserUUID: uuid.NewString(),
	// 	Name:     "vompressor",
	// 	Role: {
	// 		"username": "jeeva@getrightcare.com",
	// 		"password": "admin",
	// 	},
	// 	Email: "vompressor@gmail.com",
	// 	// Codes: Codes{
	// 	// 	one: "1",
	// 	// 	two: "2",
	// 	// },
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: jwt.At(time.Now().Add(time.Minute * 15)), // 만료시간 15분
	// 	},
	// }

	// atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &at)
	// signedAuthToken, err := atoken.SignedString([]byte("SecretCode"))

	// fmt.Println("signedAuthToken: ", signedAuthToken)
	// type Outer_auth_info struct {
	// 	Event_subtype      string
	// 	Event_subtype_name string
	// }

	// var outer_auth_info Outer_auth_info
	// err := json.Unmarshal(inputData.Outer_auth_info, &outer_auth_info)
	// if err != nil {
	// 	log.Fatalf("JSON unmarshaling failed: %s", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("%+v\n", movie)

	// // point_info에서 point_list를 제외한 tmpl_code와 score로 구성된 custom json을 생성하여 출력
	// var jsonb []interface{}
	// pointInfo := make(map[string]interface{})
	// json.Unmarshal([]byte(inputData.Outer_auth_info), &jsonb)
	// for _, v := range jsonb {
	// 	obj, _ := json.Marshal(v)
	// 	var jsonbItem interface{}
	// 	json.Unmarshal([]byte(obj), &jsonbItem)
	// 	fmt.Println(jsonbItem)
	// 	fmt.Println(goutil.JsonGetValue(jsonbItem, "event_subtype"))

	// 	// pointInfo[cast.ToString(goutil.JsonGetValue(jsonbItem, "tmpl_code"))] = goutil.JsonGetValue(jsonbItem, "score")
	// }
	// pointRaw, _ := json.Marshal(pointInfo)
	// inputData.Outer_auth_info = pointRaw
	// res = append(res, row)

	// goutil.JsonGetValue()

	// fmt.Println("user" + string(inputData.Outer_auth_info))

	// fmt.Println(inputData.Outer_auth_info)

	// fmt.Println(reflect.TypeOf(inputData.Is_sms_agree))

	// fmt.Println(strconv.FormatBool(inputData.Is_sms_agree))

	// fmt.Println(strconv.FormatBool(inputData.Is_sms_agree) == "false")

	// fmt.Println(reflect.TypeOf(strconv.FormatBool(inputData.Is_sms_agree)))

	// is_sms_agree, sms_agree_err := strconv.ParseBool(inputData.Is_sms_agree.string())

	// fmt.Println("is_sms_agree : ", is_sms_agree)
	// fmt.Println("is_sms_agree : ", sms_agree_err)

	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "!@#$%^&*()"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		digits + specials
	length := 10
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	// str := string(buf) // E.g. "3i[g0|)z"

	fmt.Println("str : ", str)

	// db select
	dmsDBConn, err = dao.ConnectDB(config.GetInstance().VurixDmsDB)
	if err != nil {
		zap.S().Errorf("Fail to DB Connection : %v", err)
		return dao.ResponseFunc(500, err.Error(), resultResponse), totalcount, resultBody
	}
	// var users = []User_Info{{UID: "jinzhu1"}}
	// dmsDBConn.Create(&users)

	defer dmsDBConn.Close()
	var rows *sql.Rows = nil
	var queryTest = FindGlobalProperty
	var Test2 struct {
		Prop_key   string      `json:"prop_key" gorm:"prop_key"`
		Prop_value interface{} `json:"prop_value" gorm:"prop_value"`
	}

	// dmsDBConn.ScanRows(queryTest, &rows)

	rows, err = dmsDBConn.Raw(queryTest).Rows()
	if err != nil {
		result = dao.ResponseFunc(500, "Internal Server Error", resultResponse)

	} else {
		for rows.Next() { // 존재하는 데이터를 스캐닝
			dmsDBConn.ScanRows(rows, &Test2)
		}
	}

	fmt.Println("Test2 : ", Test2)

	tt := dmsDBConn.Exec(querytesting).RowsAffected
	fmt.Println("tt : ", tt)

	dmsDBConn.Raw(query).Scan(&result)

	resultResponse = dao.Results{
		Results: result,
	}

	result = dao.ResponseFunc(200, "success", resultResponse)
	return result, totalcount, resultBody
}
