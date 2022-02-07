package area

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strings"
	"time"
)

// Area  struct
type (
	Area struct {
		// Code area code number
		Code string `json:"code"`
		// Alia 网页名称
		N2 string `json:"-"`
		// Name area name
		Name string `json:"name"`
		// Areas area children
		Areas []Area `json:"children"`
	}
	T struct {
		Children   []interface{} `json:"children"`
		Diji       string        `json:"diji"`
		QuHuaDaiMa string        `json:"quHuaDaiMa"`
		Quhao      string        `json:"quhao"`
		Shengji    string        `json:"shengji"`
		Xianji     string        `json:"xianji"`
	}
)

const (
	url     = "http://xzqh.mca.gov.cn/selectJson"
	jsonStr = "[{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"110000\",\"quhao\":\"NULL\",\"shengji\":\"北京市（京）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"120000\",\"quhao\":\"NULL\",\"shengji\":\"天津市（津）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"130000\",\"quhao\":\"\",\"shengji\":\"河北省（冀）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"140000\",\"quhao\":\"NULL\",\"shengji\":\"山西省（晋）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"150000\",\"quhao\":\"NULL\",\"shengji\":\"内蒙古自治区（内蒙古）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"210000\",\"quhao\":\"NULL\",\"shengji\":\"辽宁省（辽）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"220000\",\"quhao\":\"NULL\",\"shengji\":\"吉林省（吉）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"230000\",\"quhao\":\"NULL\",\"shengji\":\"黑龙江省（黑）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"310000\",\"quhao\":\"NULL\",\"shengji\":\"上海市（沪）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"320000\",\"quhao\":\"\",\"shengji\":\"江苏省（苏）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"330000\",\"quhao\":\"NULL\",\"shengji\":\"浙江省（浙）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"340000\",\"quhao\":\"NULL\",\"shengji\":\"安徽省（皖）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"350000\",\"quhao\":\"NULL\",\"shengji\":\"福建省（闽）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"360000\",\"quhao\":\"\",\"shengji\":\"江西省（赣）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"370000\",\"quhao\":\"NULL\",\"shengji\":\"山东省（鲁）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"410000\",\"quhao\":\"\",\"shengji\":\"河南省（豫）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"420000\",\"quhao\":\"NULL\",\"shengji\":\"湖北省（鄂）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"430000\",\"quhao\":\"NULL\",\"shengji\":\"湖南省（湘）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"440000\",\"quhao\":\"\",\"shengji\":\"广东省（粤）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"450000\",\"quhao\":\"\",\"shengji\":\"广西壮族自治区（桂）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"460000\",\"quhao\":\"\",\"shengji\":\"海南省（琼）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"500000\",\"quhao\":\"NULL\",\"shengji\":\"重庆市（渝）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"510000\",\"quhao\":\"NULL\",\"shengji\":\"四川省（川、蜀）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"520000\",\"quhao\":\"NULL\",\"shengji\":\"贵州省（黔、贵）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"530000\",\"quhao\":\"NULL\",\"shengji\":\"云南省（滇、云）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"540000\",\"quhao\":\"\",\"shengji\":\"西藏自治区（藏）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"610000\",\"quhao\":\"NULL\",\"shengji\":\"陕西省（陕、秦）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"620000\",\"quhao\":\"NULL\",\"shengji\":\"甘肃省（甘、陇）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"630000\",\"quhao\":\"\",\"shengji\":\"青海省（青）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"640000\",\"quhao\":\"NULL\",\"shengji\":\"宁夏回族自治区（宁）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"650000\",\"quhao\":\"\",\"shengji\":\"新疆维吾尔自治区（新）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"810000\",\"quhao\":\"00852\",\"shengji\":\"香港特别行政区（港）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"820000\",\"quhao\":\"00853\",\"shengji\":\"澳门特别行政区（澳）\",\"xianji\":\"\"},{\"children\":[],\"diji\":\"\",\"quHuaDaiMa\":\"710000\",\"quhao\":\"\",\"shengji\":\"台湾省（台）\",\"xianji\":\"\"}]"
)

var areas = make([]Area, 0)

func init() {
	area := make([]T, 0)
	err := json.Unmarshal([]byte(jsonStr), &area)
	if err != nil {
		panic(err)
	}
	for _, area := range area {
		a := Area{Code: area.QuHuaDaiMa, Name: area.Shengji[:strings.Index(area.Shengji, "（")], N2: area.Shengji}
		areas = append(areas, a)
	}
}
func GetArea() ([]Area, error) {
	for i, area := range areas {
		fmt.Printf("省: %s", area.Name)
		city, err := GetCity(area.N2)
		if err != nil {
			return nil, err
		}
		for j, c := range city {
			fmt.Printf(" 市: %s", c.Name)
			count, err := GetCounty(area.N2, c.N2)
			if err != nil {
				return nil, err
			}
			fmt.Printf(" 区:%v \n", count)
			city[j].Areas = count
		}
		areas[i].Areas = city
	}
	return areas, nil
}
func GetCity(province string) ([]Area, error) {
	body := map[string]string{"shengji": province}
	result, err := fetch(body)
	if err != nil {
		return nil, err
	}
	city := make([]Area, 0)
	for _, area := range result {
		a := Area{Code: area.QuHuaDaiMa, Name: area.Diji, N2: area.Diji}
		city = append(city, a)
	}
	return city, nil
}
func GetCounty(province, city string) ([]Area, error) {
	body := map[string]string{"shengji": province, "diji": city}
	result, err := fetch(body)
	if err != nil {
		return nil, err
	}
	county := make([]Area, 0)
	for _, area := range result {
		a := Area{Code: area.QuHuaDaiMa, Name: area.Xianji, N2: area.Xianji}
		county = append(county, a)
	}
	return county, nil
}

func fetch(data map[string]string) ([]T, error) {
	request := resty.New().SetContentLength(true).SetPreRequestHook(func(client *resty.Client, request *http.Request) error {
		if request.Method == http.MethodPost {
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
		}
		return nil
	}).R()
	var area []T
	for {
		response, err := request.SetFormData(data).Post(url)
		if err != nil {
			return nil, fmt.Errorf("request error:%v", err.Error())
		}
		if response.StatusCode() >= 200 && response.StatusCode() < 300 {
			err := json.Unmarshal(response.Body(), &area)
			if err != nil {
				return nil, fmt.Errorf("request error:%v", err.Error())
			}
			return area, nil
		} else {
			fmt.Printf("[Error] %d 休眠 30 秒重试 \n", response.StatusCode())
			time.Sleep(30 * time.Second)
		}
	}
}
