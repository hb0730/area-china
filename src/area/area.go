package area

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

/**
省
*/
// <td><a href='11.html'>北京市<br/></a></td>
const pReg string = "<td><a href='(.*?).html'>(.*?)<br/></a></td>"

/**
市
*/
const casReg string = "<tr class='.*?'><td><a href=.*?>(.*?)</a></td><td><a href=.*?>(.*?)</a></td></tr>"

// <tr class='countytr'><td>130101000000</td><td>市辖区</td></tr>
const vReg string = "<tr class='.*?'><td>(.*?)</td><td>.*?</td><td>(.*?)</td></tr>"

//Start
func Start() {
	province := getProvince()
	for _, p := range province {
		getCity(&p)
		for _, c := range p.Areas {
			getCounty(&c)
			for _, v := range c.Areas {
				fmt.Printf("%s %s %s \n", p.Name, c.Name, v.Name)
			}
		}
	}
	// 获取市
}

/**
省
*/
func getProvince() []Area {
	host := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm"
	url := "/2019/index.html"
	areas := fetch(host, url, pReg, 2)
	return areas
}

/**
市
*/
func getCity(area *Area) []Area {
	host := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm"
	url := "/2019/" + area.Code + ".html"
	areas := fetch(host, url, casReg, 4)
	area.Areas = areas
	return areas
}

/**

 */
func getCounty(area *Area) *Area {
	host := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm"
	cCode := area.Code[0:2]
	url := "/2019/" + cCode + "/" + area.Code + ".html"
	areas := fetch(host, url, casReg, 6)
	area.Areas = areas
	return area
}

func fetch(host string, route string, reg string, codeLen int) []Area {
	client := &http.Client{}
	request, err := http.NewRequest("GET", host+route, nil)
	if err != nil {
		fmt.Println("fatal error ", err.Error())
		os.Exit(0)
	}
	request.Header.Add("Accept-Language", "")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36")
	request.Header.Add("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	response, err := client.Do(request)
	if err != nil || response == nil {
		fmt.Print(err.Error())
	}
	defer response.Body.Close()
	byte2, _ := ioutil.ReadAll(response.Body)
	env := mahonia.NewDecoder("GBK")
	out := env.ConvertString(string(byte2))
	compile := regexp.MustCompile(reg)
	allString := compile.FindAllStringSubmatch(out, -1)
	areas := make([]Area, len(allString))
	for i, match := range allString {
		areas[i] = Area{match[1][0:codeLen], match[2], nil}
	}
	return areas
}

type Area struct {
	Code  string
	Name  string
	Areas []Area
}
