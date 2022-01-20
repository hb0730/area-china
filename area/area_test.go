package area

import (
	"fmt"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestCode(t *testing.T) {
	prCode := "110100000000"
	for strings.HasSuffix(prCode, "0") {
		prCode = strings.TrimSuffix(prCode, "0")
	}
	prCode = "441900003000"
	for strings.HasSuffix(prCode, "0") {
		prCode = strings.TrimSuffix(prCode, "0")
	}
	fmt.Println(prCode)
}

func TestComplement(t *testing.T) {
	prCode := "11010000000"
	defaultLength := 6
	defaultCodeLength := "00000000000000000"
	if len(prCode) < defaultLength {
		prCode += defaultCodeLength[0:(defaultLength - len(prCode))]
		fmt.Printf(prCode)
	}
}

func TestStart(t *testing.T) {
	Start("2020", 6)
}

func TestWriteJson(t *testing.T) {
	_year = "2020"
	WriteJson(nil)
}

func Test_getStreet(t *testing.T) {
	_year = "2020"
	a := &Area{
		Code: "110101",
		Name: "东城区",
	}
	getStreet(a)
}
func Test_pReg(t *testing.T) {
	body := `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3c.org/TR/1999/REC-html401-19991224/loose.dtd">
<HTML>

<HEAD>
    <META content="text/html; charset=utf-8" http-equiv=Content-Type>
    <TITLE>2021年统计用区划代码</TITLE>
    <STYLE type=text/css>
        BODY {
            MARGIN: 0px
        }
        
        BODY {
            FONT-SIZE: 12px
        }
        
        TD {
            FONT-SIZE: 12px
        }
        
        TH {
            FONT-SIZE: 12px
        }
        
        .redBig {
            COLOR: #d00018;
            FONT-SIZE: 18px;
            FONT-WEIGHT: bold
        }
        
        .STYLE3 a {
            COLOR: #fff;
            text-decoration: none;
        }
        
        .STYLE5 {
            COLOR: #236fbe;
            FONT-WEIGHT: bold
        }
        
        .content {
            LINE-HEIGHT: 1.5;
            FONT-SIZE: 10.4pt
        }
        
        .tdPading {
            PADDING-LEFT: 30px
        }
        
        .blue {
            COLOR: #0000ff
        }
        
        .STYLE6 {
            COLOR: #ffffff
        }
        
        .a2 {
            LINE-HEIGHT: 1.5;
            COLOR: #2a6fbd;
            FONT-SIZE: 12px
        }
        
        a2:link {
            LINE-HEIGHT: 1.5;
            COLOR: #2a6fbd;
            FONT-SIZE: 12px
        }
        
        a2:hover {
            LINE-HEIGHT: 1.5;
            COLOR: #2a6fbd;
            FONT-SIZE: 12px;
            TEXT-DECORATION: underline
        }
        
        a2:visited {
            LINE-HEIGHT: 1.5;
            COLOR: #2a6fbd;
            FONT-SIZE: 12px
        }
    </STYLE>
    <SCRIPT language=javascript>
        function doZoom(size) {
            document.getElementById("zoom").style.fontSize = size + "px";
        }
    </SCRIPT>
    <META name=GENERATOR content="MSHTML 8.00.7600.16700">
</HEAD>

<BODY>
    <TABLE border=0 cellSpacing=0 cellPadding=0 width=778 align=center>
        <TBODY>
            <TR>
                <TD colSpan=2>
                    <IMG src="http://www.stats.gov.cn/images/banner.jpg" width=778 height=135>
                </TD>
            </TR>
        </TBODY>
    </TABLE>
    <MAP id=Map name=Map><AREA href="http://www.stats.gov.cn/english/" shape=rect coords=277,4,328,18><AREA href="http://www.stats.gov.cn:82/" shape=rect coords=181,4,236,18><AREA href="http://www.stats.gov.cn/" shape=rect coords=85,4,140,17></MAP>
    <TABLE border=0 cellSpacing=0 cellPadding=0 width=778 align=center>
        <TBODY>
            <TR>
                <TD vAlign=top>
                    <TABLE style="MARGIN-TOP: 15px; MARGIN-BOTTOM: 18px" border=0 cellSpacing=0 cellPadding=0 width="100%" align=center>
                        <TBODY>
                            <TR>
                                <TD style=" BACKGROUND-REPEAT: repeat-x; BACKGROUND-POSITION: 50% top" background=/images/topLine.gif align=right></TD>
                            </TR>
                            <TR>
                                <TD style="BACKGROUND-REPEAT: repeat-y; BACKGROUND-POSITION: right 50%" vAlign=top background=images/rightBorder.gif>
                                    <TABLE border=0 cellSpacing=0 cellPadding=0 width="100%">
                                        <TBODY>
                                            <TR>
                                                <TD width="1%" height="200" vAlign=top>
                                                    <table class="citytable">
                                                        <tr class="cityhead">
                                                            <td width=150>统计用区划代码</td>
                                                            <td>名称</td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4401.html">440100000000</a></td>
                                                            <td><a href="44/4401.html">广州市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4402.html">440200000000</a></td>
                                                            <td><a href="44/4402.html">韶关市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4403.html">440300000000</a></td>
                                                            <td><a href="44/4403.html">深圳市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4404.html">440400000000</a></td>
                                                            <td><a href="44/4404.html">珠海市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4405.html">440500000000</a></td>
                                                            <td><a href="44/4405.html">汕头市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4406.html">440600000000</a></td>
                                                            <td><a href="44/4406.html">佛山市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4407.html">440700000000</a></td>
                                                            <td><a href="44/4407.html">江门市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4408.html">440800000000</a></td>
                                                            <td><a href="44/4408.html">湛江市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4409.html">440900000000</a></td>
                                                            <td><a href="44/4409.html">茂名市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4412.html">441200000000</a></td>
                                                            <td><a href="44/4412.html">肇庆市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4413.html">441300000000</a></td>
                                                            <td><a href="44/4413.html">惠州市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4414.html">441400000000</a></td>
                                                            <td><a href="44/4414.html">梅州市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4415.html">441500000000</a></td>
                                                            <td><a href="44/4415.html">汕尾市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4416.html">441600000000</a></td>
                                                            <td><a href="44/4416.html">河源市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4417.html">441700000000</a></td>
                                                            <td><a href="44/4417.html">阳江市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4418.html">441800000000</a></td>
                                                            <td><a href="44/4418.html">清远市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4419.html">441900000000</a></td>
                                                            <td><a href="44/4419.html">东莞市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4420.html">442000000000</a></td>
                                                            <td><a href="44/4420.html">中山市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4451.html">445100000000</a></td>
                                                            <td><a href="44/4451.html">潮州市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4452.html">445200000000</a></td>
                                                            <td><a href="44/4452.html">揭阳市</a></td>
                                                        </tr>
                                                        <tr class="citytr">
                                                            <td><a href="44/4453.html">445300000000</a></td>
                                                            <td><a href="44/4453.html">云浮市</a></td>
                                                        </tr>
                                                    </table>
                                                </TD>
                                            </TR>
                                        </TBODY>
                                    </TABLE>
                                </TD>
                            </TR>
                            <TR>
                                <TD style="BACKGROUND-REPEAT: repeat-x; BACKGROUND-POSITION: 50% top" background=images/borderBottom.gif></TD>
                            </TR>
                        </TBODY>
                    </TABLE>
                </TD>
            </TR>
            <TR>
                <TD bgColor=#e2eefc height=2></TD>
            </TR>
            <TR>
                <TD class=STYLE3 height=60>
                    <DIV align=center style="background-color:#1E67A7; height:75px; color:#fff;"><br/>版权所有：国家统计局
                        <A class=STYLE3 href="http://www.miibeian.gov.cn/" target=_blank>京ICP备05034670号</A>
                        <BR>
                        <BR>地址：北京市西城区月坛南街57号（100826）
                        <BR>
                    </DIV>
                </TD>
            </TR>
        </TBODY>
    </TABLE>
</BODY>

</HTML>`
	m := minify.New()
	m.Add("text/html", &html.Minifier{KeepEndTags: true, KeepDocumentTags: true, KeepComments: true, KeepConditionalComments: true, KeepQuotes: true, KeepDefaultAttrVals: true})
	body, _ = m.String("text/html", body)
	//body, _ = minify.New().Add("text/html",&html.Minifier{})
	compile := regexp.MustCompile(`<tr class=".*?"><td><a href=.*?>(.*?)</a></td><td><a href=.*?>(.*?)</a></td></tr>`)
	allString := compile.FindAllStringSubmatch(body, -1)
	if len(allString) == 0 {
		t.Errorf("表达式错误")
	}
}

func TestConvert(t *testing.T) {
	type args struct {
		src        []byte
		srcCode    string
		targetCode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert(tt.args.src, tt.args.srcCode, tt.args.targetCode); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStart1(t *testing.T) {
	type args struct {
		year   string
		length int
	}
	tests := []struct {
		name string
		args args
		want []Area
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Start(tt.args.year, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteJson1(t *testing.T) {
	type args struct {
		area []Area
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteJson(tt.args.area)
		})
	}
}

func Test_fetch(t *testing.T) {
	type args struct {
		host  string
		route string
		reg   string
	}
	tests := []struct {
		name string
		args args
		want []Area
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fetch(tt.args.host, tt.args.route, tt.args.reg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBody(t *testing.T) {
	type args struct {
		host  string
		route string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBody(tt.args.host, tt.args.route); got != tt.want {
				t.Errorf("getBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCity(t *testing.T) {
	_year = "2021"
	area := Area{Code: "440000", Name: "广东省"}
	getCity(&area)
}

func Test_getCounty(t *testing.T) {
	_year = "2021"
	area := Area{Code: "659000", Name: "广东省"}
	getCounty(&area)
}

func Test_getProvince(t *testing.T) {
	_year = "2021"
	area := Area{Code: "440000", Name: "广东省"}
	getCounty(&area)
}

func Test_getStreet1(t *testing.T) {
	type args struct {
		area *Area
	}
	tests := []struct {
		name string
		args args
		want []Area
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStreet(tt.args.area); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStreet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toUtf8(t *testing.T) {
	type args struct {
		content     []byte
		contentType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toUtf8(tt.args.content, tt.args.contentType); got != tt.want {
				t.Errorf("toUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}
