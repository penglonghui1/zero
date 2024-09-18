package tools

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pengcainiao2/zero/rest/httprouter"
	"golang.org/x/crypto/bcrypt"

	"github.com/963204765/httpclient/httplib"
	json "github.com/json-iterator/go"
)

var (
	week = map[int]string{
		0: "周日",
		1: "周一",
		2: "周二",
		3: "周三",
		4: "周四",
		5: "周五",
		6: "周六",
	}
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

/**
 * 计算两个日期相差天数
 * @synopsis GetSubDay
 * @param string
 * @param ...string
 * @return
 */
func GetSubDay(current interface{}, befores ...string) int64 {
	before := "2020-05-06"
	if len(befores) > 0 {
		before = befores[0]
	}

	var (
		loc, _   = time.LoadLocation("Local")
		start, _ = time.ParseInLocation("2006-01-02", before, loc)
		end      time.Time
	)

	switch x := current.(type) {
	case int64:
		end, _ = time.ParseInLocation("2006-01-02", time.Unix(x, 0).Format("2006-01-02"), loc)
	case string:
		end, _ = time.ParseInLocation("2006-01-02", x, loc)
	}

	return (end.Unix() - start.Unix()) / 86400
}

/**
 * 计算两个日期相差月数
 * @synopsis GetSubMonth
 * @param string
 * @param ...string
 * @return
 */
func GetSubMonth(current interface{}, befores ...string) int64 {
	before := "2020-05-06"
	if len(befores) > 0 {
		before = befores[0]
	}

	var (
		loc, _   = time.LoadLocation("Local")
		start, _ = time.ParseInLocation("2006-01-02", before, loc)
		end      time.Time
	)

	switch x := current.(type) {
	case int64:
		end, _ = time.ParseInLocation("2006-01-02", time.Unix(x, 0).Format("2006-01-02"), loc)
	case string:
		end, _ = time.ParseInLocation("2006-01-02", x, loc)
	}

	y1 := end.Year()
	y2 := start.Year()
	m1 := int(end.Month())
	m2 := int(start.Month())
	d1 := 1
	d2 := 1

	yearInterval := y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}

	monthInterval %= 12
	return int64(yearInterval*12 + monthInterval)
}

/**
 * 判断字符串是否在数组中
 * @synopsis InArray
 * @param str string
 * @param arr []string
 * @return bool
 */
func InArray(str interface{}, arr interface{}) bool {
	switch arr := arr.(type) {
	case []string:
		if str, ok := str.(string); ok {
			for _, val := range arr {
				if val == str {
					return true
				}
			}
		}
	case []int8:
		if str, ok := str.(int8); ok {
			for _, val := range arr {
				if val == str {
					return true
				}
			}
		}
	case []int:
		if str, ok := str.(int); ok {
			for _, val := range arr {
				if val == str {
					return true
				}
			}
		}
	case []int16:
		if str, ok := str.(int16); ok {
			for _, val := range arr {
				if val == str {
					return true
				}
			}
		}
	case []int32:
		if str, ok := str.(int32); ok {
			for _, val := range arr {
				if val == str {
					return true
				}
			}
		}
	case []int64:
		if str, ok := str.(int64); ok {
			for _, val := range arr {
				if val == str {
					return true
				}
			}
		}
	}
	return false
}

/**
 * 删除数组元素
 * @synopsis RemoveArray
 * @param arr []string
 * @param val string
 * @return []string
 */
func RemoveArray(arr []string, val string) []string {
	var key = -1
	for k, v := range arr {
		if val == v {
			key = k
			break
		}
	}
	if key == -1 {
		return arr
	} else {
		var result = make([]string, 0)
		if len(arr[:key]) > 0 {
			result = append(result, arr[:key]...)
		}
		if len(arr[key+1:]) > 0 {
			result = append(result, arr[key+1:]...)
		}
		return result
	}
}

/**
 * 数组去重
 * @synopsis ArrayUnique
 * @param []string
 * @return
 */
func ArrayUnique(arr []string) []string {
	result := make([]string, 0)
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range arr {
		if e == "" {
			continue
		}
		if _, ok := tempMap[e]; !ok {
			tempMap[e] = 0
			result = append(result, e)
		}
	}
	return result
}

/**
 * 数组差集
 * @synopsis ArrayDiff
 * @param []string
 * @param []string
 * @return []string
 */
func ArrayDiff(a, b []string) []string {
	var diffArray []string
	temp := map[string]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}
	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}
	return diffArray
}

// SendHTTP 发送http请求
// Deprecated: 使用 httprouter.PerformanceRequest 代替，新的方法中集成了sentry
func SendHTTP(ctx *httprouter.Context, method, url string) *httplib.BeegoHTTPRequest {
	req := httplib.NewBeegoRequest(url, method)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if ctx != nil {
		header := ctx.Data
		req.Header("Authorization", header.Authorization)
		req.Header("X-Auth-User", header.UserID)
		req.Header("X-Auth-Platform", header.Platform)
		req.Header("X-Auth-Version", header.ClientVersion)
		req.Header("X-Auth-ClientIP", header.ClientIP)
		req.Header("X-Request-ID", header.RequestID)
	}
	return req
}

/**
 * 生成随机字符串
 * @synopsis RandString
 * @param int
 * @return
 */
func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

/**
 * 转换日期
 * @synopsis ConvertDate
 * @param string
 * @return
 */
func ConvertDate(date string) string {
	var (
		t, _   = time.Parse("2006-01-02", date)
		year   = t.Year()
		toYeay = time.Now().Year()
		format string
	)

	if year == toYeay {
		format = "1月2日"
	} else {
		format = "2006年1月2日"
	}

	dateFormat := t.Format(format)
	if time.Now().Format("2006-01-02") == date {
		dateFormat = dateFormat + "（今日）"
	} else {
		dateFormat = dateFormat + " " + week[int(t.Weekday())]
	}
	return dateFormat
}

/**
 * 删除html标签
 * @synopsis RemoveHtmlTag
 * @param str string
 * @return string
 */
func RemoveHtmlTag(str string) string {
	re, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>") //nolint
	str = re.ReplaceAllString(str, "")
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>") //nolint
	str = re.ReplaceAllString(str, "\n")
	re, _ = regexp.Compile("\\s{2,}") //nolint
	str = strings.TrimSpace(strings.ReplaceAll(str, "\n", ""))
	return strings.ReplaceAll(str, "&nbsp;", "")
}

/**
 * 日期时间解析为时间戳
 * @synopsis Strtotime
 * @param date string
 * @param now ...string
 * @return int64
 */
func Strtotime(date string, now ...string) int64 {
	nowTime, _ := time.Parse("2006-01-02", date)
	if len(now) > 0 {
		arr := strings.Split(now[0], " ")
		if len(arr) == 2 {
			num, _ := strconv.Atoi(arr[0])
			switch arr[1] {
			case "day":
				nowTime = nowTime.AddDate(0, 0, num)
			case "month":
				nowTime = nowTime.AddDate(0, num, 0)
			case "year":
				nowTime = nowTime.AddDate(num, 0, 0)
			}
		}
	}
	return time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Now().Location()).Unix()
}

// RandStringRunes - generate random string using random int
func RandStringRunes(n int) string {
	b := make([]rune, n)
	l := len(letterRunes)
	for i := range b {
		b[i] = letterRunes[rand.Intn(l)]
	}
	return string(b)
}

// StructToMap 结构体转map
func StructToMap(s interface{}, number ...bool) (map[string]interface{}, error) {
	var result = make(map[string]interface{})
	if bytesArr, err := json.Marshal(s); err != nil {
		return nil, err
	} else {
		d := json.NewDecoder(bytes.NewReader(bytesArr))
		// 设置将float64转为一个number
		if len(number) > 0 && number[0] {
			d.UseNumber()
		}
		if err := d.Decode(&result); err != nil {
			return nil, err
		}
		return result, nil
	}
}

// MapToStruct map解析到结构体
func MapToStruct(m interface{}, s interface{}) error {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonByte, s)
}

// ExtractStructFields 提取结构体字段
func ExtractStructFields(s interface{}, dest interface{}, fields ...string) error {

	bytesArr, err := json.Marshal(s)
	if err != nil {
		return err
	}

	d := json.NewDecoder(bytes.NewReader(bytesArr))
	d.UseNumber()
	var data interface{}
	if err := d.Decode(&data); err != nil {
		return err
	}

	switch data.(type) {
	case map[string]interface{}:
		var result = make(map[string]interface{})
		for _, field := range fields {
			if val, ok := data.(map[string]interface{})[field]; ok {
				result[field] = val
			}
		}
		*dest.(*map[string]interface{}) = result
	case []interface{}:
		var result = make([]map[string]interface{}, 0)
		for _, item := range data.([]interface{}) {
			var maps = make(map[string]interface{})
			for _, field := range fields {
				if val, ok := item.(map[string]interface{})[field]; ok {
					maps[field] = val
				}
			}
			result = append(result, maps)
		}
		*dest.(*[]map[string]interface{}) = result
	default:
		return errors.New("错误类型")
	}

	return nil
}

// MD5 生成md5
func MD5(str string) string {
	has := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", has)
}

// PasswordHash 生成hash密码
func PasswordHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// PasswordVerify 验证密码
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
