package area

import (
	"fmt"
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
