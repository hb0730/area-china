package area

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hb0730/go-request"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	// 省级
	// <td><a href='11.html'>北京市<br/></a></td>
	pReg string = "<td><a href='(.*?).html'>(.*?)<br/></a></td>"
	// 地级，县级，乡级
	casReg string = "<tr class='.*?'><td><a href=.*?>(.*?)</a></td><td><a href=.*?>(.*?)</a></td></tr>"
	//村级
	vReg string = "<tr class='.*?'><td>(.*?)</td><td>.*?</td><td>(.*?)</td></tr>"
	host        = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm"

	//城乡规划默认编码长度
	defaultLength = "00000000000000000"
	minLength     = 2
	maxLength     = 17
)

var _year string
var _length int

//Start 开始
func Start(year string, length int) []Area {
	if length < 2 {
		_length = minLength
	} else if length > maxLength {
		length = maxLength
	} else {
		_length = length
	}
	_year = year
	province := getProvince()
	for i1, p := range province {
		city := getCity(&p)
		province[i1] = p
		for i2, c := range city {
			county := getCounty(&c)
			city[i2] = c
			for _, v := range county {
				fmt.Printf("%s %s %s \n", p.Name, c.Name, v.Name)
			}
		}
	}
	// 导出json
	WriteJson(province)
	return province
}

//getProvince 获取省级地区,编码规则是1~2位
func getProvince() []Area {
	// /2019/index.html
	url := fmt.Sprintf("/%s/%s", _year, "index.html")
	areas := fetch(host, url, pReg)
	return areas
}

//getCity 获取市级地区 编码规则是3~4位
func getCity(area *Area) []Area {
	pCode := area.Code[0:2]
	//url := "/2019/" + cCode + ".html"
	url := fmt.Sprintf("/%s/%s.html", _year, pCode)
	areas := fetch(host, url, casReg)
	area.Areas = areas
	return areas
}

//getCounty 获取县级地区 编码规则是5~6位
func getCounty(area *Area) []Area {
	pCode := area.Code[0:2]
	cCode := area.Code[0:4]
	//url := "/2019/" + cCode + "/" + aCode + ".html"
	url := fmt.Sprintf("/%s/%s/%s.html", _year, pCode, cCode)
	areas := fetch(host, url, casReg)
	area.Areas = areas
	return areas
}

//getStreet 抓取街道
func getStreet(area *Area) []Area {
	pCode := area.Code[:2]
	cCodeSuffix := area.Code[2:4]
	//url:="/2019/11/01/110101.html"
	url := fmt.Sprintf("/%s/%s/%s/%s.html", _year, pCode, cCodeSuffix, area.Code)
	areas := fetch(host, url, casReg)
	area.Areas = areas
	return nil
}

// 获取网页地区信息
// @params host
// @params route path
// @params reg 表达式
// @params codeLen 编码长度
func fetch(host string, route string, reg string) []Area {
	out := getBody(host, route)
	compile := regexp.MustCompile(reg)
	allString := compile.FindAllStringSubmatch(out, -1)
	areas := make([]Area, len(allString))
	for i, match := range allString {
		code := match[1]
		for strings.HasSuffix(code, "0") && len(code) > _length {
			code = strings.TrimSuffix(code, "0")
		}
		if len(code) < _length {
			code += defaultLength[0:(_length - len(code))]
		}

		areas[i] = Area{code, match[2], nil}
	}
	return areas
}

func getBody(host string, route string) string {
	for {
		req, err := request.CreateRequest("GET", host+route, "")
		if err != nil {
			fmt.Println("fatal error ", err.Error())
			os.Exit(0)
		}
		headers := map[string]string{
			"Accept-Language": "zh-CN,zh;q=0.9",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
			"Accept-Charset":  "GBK,utf-8;q=0.7,*;q=0.3",
			"Accept-Encoding": "gzip, deflate",
		}
		req.SetHeaders(headers)
		time.Sleep(time.Second * 2)
		err = req.Do()
		if err != nil {
			fmt.Println("fatal error")
			panic(err)
		}
		resp := req.GetResponse()
		// 熔断或者超时或者404等
		if resp.StatusCode != 200 && resp.StatusCode != 304 {
			fmt.Printf("[Error] %d 休眠 30 秒重试 \n", resp.StatusCode)
			time.Sleep(time.Duration(30) * time.Second)
		} else {

			body, _ := req.GetBody()
			utf8Body, _ := gbk2Utf8(body)
			return string(utf8Body)
		}

	}
}

//WriteJson 写入json file
func WriteJson(area []Area) {
	areaBytes, err := json.Marshal(area)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fileName := "dist/area%s-%d.json"
	currentTime := time.Now().UnixNano() / 1e6
	fileName = fmt.Sprintf(fileName, _year, currentTime)
	err = ioutil.WriteFile(fileName, areaBytes, 0666)
	if err != nil {
		fmt.Printf("create file error: %s", err.Error())
		return
	}
}

//gbk2Utf8 gbk转utf-8
func gbk2Utf8(body []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(body), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//Area 地区
type Area struct {
	Code  string `json:"code "`    //编码
	Name  string `json:"name"`     //名称
	Areas []Area `json:"children"` //下级行政
}
