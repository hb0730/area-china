package area

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/qiniu/iconv"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
)

// Area  struct
type Area struct {
	// Code area code number
	Code string `json:"code"`
	// Name area name
	Name string `json:"name"`
	// Areas area children
	Areas []Area `json:"children"`
}

const (
	// 省级
	// <td><a href='11.html'>北京市<br/></a></td>
	pReg string = `<td><a href="(.*?).html">(.*?)<br></a></td>`
	// 地级，县级，乡级
	casReg string = `<tr class=".*?"><td><a href=.*?>(.*?)</a></td><td><a href=.*?>(.*?)</a></td></tr>`
	//村级
	vReg string = `<tr class=".*?"><td>(.*?)</td><td>.*?</td><td>(.*?)</td></tr>`
	host        = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm"

	//城乡规划默认编码长度
	defaultLength = "00000000000000000"
	minLength     = 2
	maxLength     = 17
)

// Spider fetch area  config
type Spider struct {
	year        string
	_codeLength int
	m           *minify.M
}

func NewSpider(year string, codeLength int) (Spider, error) {
	s := Spider{}
	if len(year) == 0 {
		return s, errors.New("year missing")
	}
	s.year = year
	//压缩
	s.m = minify.New()
	s.m.Add("text/html", &html.Minifier{KeepEndTags: true, KeepDocumentTags: true, KeepComments: true, KeepConditionalComments: true, KeepQuotes: true, KeepDefaultAttrVals: true})
	// 编码长度
	if codeLength < 2 {
		s._codeLength = minLength
	} else if codeLength > maxLength {
		s._codeLength = maxLength
	} else {
		s._codeLength = codeLength
	}
	return s, nil
}

//Start 开始
func (s Spider) Start() ([]Area, error) {
	province, err := s.GetProvince("")
	if err != nil {
		return nil, err
	}
	for i1, p := range province {
		city, err := s.GetCity(&p)
		if err != nil {
			return nil, err
		}
		province[i1] = p
		for i2, c := range city {
			county, err := s.GetCounty(&c)
			if err != nil {
				return nil, err
			}
			city[i2] = c
			for _, v := range county {
				fmt.Printf("%s %s %s \n", p.Name, c.Name, v.Name)
			}
		}
	}
	return province, nil
}

// GetProvince 获取省级地区,编码规则是1~2位
// 通过name只抓取当前数据
func (s Spider) GetProvince(name string) ([]Area, error) {
	url := fmt.Sprintf("/%s/%s", s.year, "index.html")
	areas, err := s.fetch(host, url, pReg)
	areaFilter := make([]Area, 0)
	if len(areas) != 0 && len(name) != 0 {
		for _, area := range areas {
			if area.Name == name {
				areaFilter = append(areaFilter, area)
				return areaFilter, nil
			}
		}
	}
	return areas, err
}

// GetCity 获取市级地区 编码规则是3~4位
func (s Spider) GetCity(area *Area) ([]Area, error) {
	if len(area.Code) == 0 {
		return nil, fmt.Errorf("code missing")
	}
	pCode := area.Code[0:2]
	//url := "/2019/" + cCode + ".html"
	url := fmt.Sprintf("/%s/%s.html", s.year, pCode)
	areas, err := s.fetch(host, url, casReg)
	area.Areas = areas
	return areas, err
}

// GetCounty 获取县级地区 编码规则是5~6位
func (s Spider) GetCounty(area *Area) ([]Area, error) {
	if len(area.Code) == 0 {
		return nil, fmt.Errorf("code missing")
	}
	pCode := area.Code[0:2]
	cCode := area.Code[0:4]
	//url := "/2019/" + cCode + "/" + aCode + ".html"
	url := fmt.Sprintf("/%s/%s/%s.html", s.year, pCode, cCode)
	areas, err := s.fetch(host, url, casReg)
	area.Areas = areas
	return areas, err
}

// GetStreet 抓取街道
func (s Spider) GetStreet(area *Area) ([]Area, error) {
	if len(area.Code) == 0 {
		return nil, fmt.Errorf("code missing")
	}
	pCode := area.Code[:2]
	cCodeSuffix := area.Code[2:4]
	//url:="/2019/11/01/110101.html"
	url := fmt.Sprintf("/%s/%s/%s/%s.html", s.year, pCode, cCodeSuffix, area.Code)
	areas, err := s.fetch(host, url, casReg)
	area.Areas = areas
	return areas, err
}

// 获取网页地区信息转换成结构信息
func (s Spider) fetch(host string, route, reg string) ([]Area, error) {
	content, err := s.getBody(host, route)
	if err != nil {
		return nil, err
	}
	//压缩
	content, _ = s.m.String("text/html", content)
	// 提取
	compile := regexp.MustCompile(reg)
	area := compile.FindAllStringSubmatch(content, -1)
	areas := make([]Area, len(area))
	for i, match := range area {
		code := match[1]
		for strings.HasSuffix(code, "0") && len(code) > s._codeLength {
			code = strings.TrimSuffix(code, "0")
		}
		if len(code) < s._codeLength {
			code += defaultLength[0:(s._codeLength - len(code))]
		}

		areas[i] = Area{code, match[2], nil}
	}
	return areas, nil
}

func (s Spider) getBody(host, route string) (string, error) {
	request := resty.New().R()
	for {
		// 不频繁请求
		time.Sleep(3 * time.Second)

		response, err := request.Get(host + route)
		if err != nil {
			return "", fmt.Errorf("request error:%v", err.Error())
		}
		// 熔断或者超时或者404等
		if response.StatusCode() != 200 && response.StatusCode() != 304 {
			fmt.Printf("[Error] %d 休眠 30 秒重试 \n", response.StatusCode())
			time.Sleep(30 * time.Second)
		} else {
			return toUtf8(response.Body(), response.Header().Get("Content-Type")), nil
		}
	}
}

// 内部编码判断和转换，会自动判断传入的字符串编码，并将它转换成utf-8
func toUtf8(content []byte, contentType string) string {
	var htmlEncode string
	contentBody := string(content)

	if strings.Contains(contentType, "gbk") || strings.Contains(contentType, "gb2312") || strings.Contains(contentType, "gb18030") || strings.Contains(contentType, "windows-1252") {
		htmlEncode = "gb18030"
	} else if strings.Contains(contentType, "big5") {
		htmlEncode = "big5"
	} else if strings.Contains(contentType, "utf-8") {
		htmlEncode = "utf-8"
	}
	if htmlEncode == "" {
		//先尝试读取charset
		reg := regexp.MustCompile(`(?is)<meta[^>]*charset\s*=["']?\s*([A-Za-z0-9\-]+)`)
		match := reg.FindStringSubmatch(contentBody)
		if len(match) > 1 {
			contentType = strings.ToLower(match[1])
			if strings.Contains(contentType, "gbk") || strings.Contains(contentType, "gb2312") || strings.Contains(contentType, "gb18030") || strings.Contains(contentType, "windows-1252") {
				htmlEncode = "gb18030"
			} else if strings.Contains(contentType, "big5") {
				htmlEncode = "big5"
			} else if strings.Contains(contentType, "utf-8") {
				htmlEncode = "utf-8"
			}
		}
		if htmlEncode == "" {
			reg = regexp.MustCompile(`(?is)<title[^>]*>(.*?)</title>`)
			match = reg.FindStringSubmatch(contentBody)
			if len(match) > 1 {
				aa := match[1]
				_, contentType, _ = charset.DetermineEncoding([]byte(aa), "")
				htmlEncode = strings.ToLower(htmlEncode)
				if strings.Contains(contentType, "gbk") || strings.Contains(contentType, "gb2312") || strings.Contains(contentType, "gb18030") || strings.Contains(contentType, "windows-1252") {
					htmlEncode = "gb18030"
				} else if strings.Contains(contentType, "big5") {
					htmlEncode = "big5"
				} else if strings.Contains(contentType, "utf-8") {
					htmlEncode = "utf-8"
				}
			}
		}
	}
	return Convert(content, htmlEncode, "utf-8")

}

// Convert 编码转换
func Convert(src []byte, srcCode, targetCode string) string {
	cd, err := iconv.Open(targetCode, srcCode)
	if err != nil {
		return ""
	}
	defer cd.Close()
	var outbuf [512]byte
	s1, _, err := cd.Conv(src, outbuf[:])
	if err != nil {
		return ""
	}
	return string(s1)
}

// WriteJson 写入json file
func WriteJson(filename string, area []Area) error {
	if len(filename) == 0 {
		panic("filename missing...")
	}
	areaBytes, err := json.Marshal(area)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, areaBytes, 0666)
	if err != nil {
		return fmt.Errorf("create file error: %s", err.Error())
	}
	return nil
}
