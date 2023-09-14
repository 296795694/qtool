package str

import (
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

const (
	letterBytes = "0123456789abcdefghijklmnopqrstuvwxyz"
	intBytes    = "0123456789"
)

// RandString 随机字符串 相同时间返回的值一样
func RandString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	const (
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// RandInt 随机数
func RandInt(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	const (
		intIdxBits = 6
		intIdxMask = 1<<intIdxBits - 1
		intIdxMax  = 63 / intIdxBits
	)
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), intIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), intIdxMax
		}
		if idx := int(cache & intIdxMask); idx < len(intBytes) {
			b[i] = intBytes[idx]
			i--
		}
		cache >>= intIdxBits
		remain--
	}
	return string(b)
}

// FirstToUpper 首字母转大写
func FirstToUpper(str string) string {
	if str == "" {
		return str
	}
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// GenerateRandomNumber 生成随机数
func GenerateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

// Substr 截取字符串
func Substr(str string, start int, length int) string {
	if length == 0 {
		return ""
	}
	rune_str := []rune(str)
	len_str := len(rune_str)
	if start < 0 {
		start = len_str + start
	}
	if start > len_str {
		start = len_str
	}
	end := start + length
	if end > len_str {
		end = len_str
	}
	if length < 0 {
		end = len_str + length
	}
	if start > end {
		start, end = end, start
	}
	return string(rune_str[start:end])
}

// StrToInt string 转 uint
func StrToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return i
}

//StrToUInt string 转int
func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}

// UIntToStr uint 转 str
func UIntToStr(i uint) string {
	return strconv.Itoa(int(i))
}

// IntToStr int 转 str
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

// Hump 字符串转驼峰
func Hump(str string) (hump string) {
	if str == "id" {
		hump = "ID"
		return
	}
	strArr := strings.Split(str, "_")
	for _, item := range strArr {
		hump += FirstToUpper(item)
	}
	return
}

// DeletePreAndSufSpace 删除左右两边的所有空格
func DeletePreAndSufSpace(str string) string {
	strList := []byte(str)
	spaceCount, count := 0, len(strList)
	for i := 0; i <= len(strList)-1; i++ {
		if strList[i] == 32 {
			spaceCount++
		} else {
			break
		}
	}
	strList = strList[spaceCount:]
	spaceCount, count = 0, len(strList)
	for i := count - 1; i >= 0; i-- {
		if strList[i] == 32 {
			spaceCount++
		} else {
			break
		}
	}
	return string(strList[:count-spaceCount])
}

// Float64ToStr StrToFloat64 str 转 float64
func Float64ToStr(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}

// StrToFloat64 str 转 float64
func StrToFloat64(str string) float64 {
	i, e := strconv.ParseFloat(str, 64)
	if e != nil {
		return 0.0
	}
	return i
}

// FormatPrice 格式化价格
func FormatPrice(str float64) float64 {
	price, _ := decimal.NewFromFloat(str).Round(2).Float64()
	return price
}

// FormatPriceAsString 格式化价格 返回字符串
func FormatPriceAsString(price float64, symbol bool, intFlag bool) string {
	str := decimal.NewFromFloat(price).Round(2).String()
	if symbol == true {
		if price > 0 {
			str = "￥" + str
		} else {
			str = "-￥" + Float64ToStr(math.Abs(price), 2)
		}
	}
	if price == 0 && !intFlag {
		str = ""
	}
	return str
}

// GetIncrementId 生成编号
func GetIncrementId() string {
	return time.Now().Format("20060102150405") + RandInt(6)
}

// TrimHtml 去掉html
func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

//GetDisplayNickName 隐藏中间字符串
func GetDisplayNickName(nickname string) string {
	r := []rune(nickname)
	if len(r) < 1 {
		return nickname
	} else if len(r) == 2 {
		return string(r[0]) + "*"
	} else if len(r) == 11 {
		rs := []rune(nickname)
		return string(rs[0:3]) + "****" + string(rs[7:11])
	} else if len(r) == 18 {
		rs := []rune(nickname)
		return string(rs[0:3]) + "***********" + string(rs[14:18])
	} else {
		return string(r[0]) + "*" + string(r[len(r)-1])
	}
}

func GenerateCardNum(count int) []string {
	var nums []string
	if count > 0 {
		cardsMap := make(map[string]int)
		for len(cardsMap) < count {
			cardNum := fmt.Sprintf("%s%s", time.Now().Format("20060102150405")[2:], RandInt(3))
			cardsMap[cardNum] = 1
		}

		for key, _ := range cardsMap {
			nums = append(nums, key)
		}
	}

	return nums
}

func GetBase64AndReplaceStr(str string) (base64str string) {
	if str != "" {
		base64str = base64.StdEncoding.EncodeToString([]byte(str))
		base64str = strings.ReplaceAll(base64str, "+", "-")
		base64str = strings.ReplaceAll(base64str, "/", "_")
	}
	return
}

func FormatTime(s int) string {
	if s < 10 {
		return "0" + IntToStr(s)
	}
	return IntToStr(s)

}
