package area

import (
	"testing"
)

func TestGetArea(t *testing.T) {
	areas, err := GetArea()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", areas)
}

func Test_fetch(t *testing.T) {
	type args struct {
		data map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "省份", args: args{map[string]string{"shengji": "新疆维吾尔自治区（新）"}}},
		{name: "地级市", args: args{map[string]string{"shengji": "新疆维吾尔自治区（新）", "diji": "乌鲁木齐市"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetch(tt.args.data)
			if err != nil {
				t.Errorf("%v", err)
			}
			t.Logf("%v", got)
		})
	}
}

func TestGetCounty(t *testing.T) {
	type args struct {
		province string
		city     string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"地级市", args{province: "北京市（京）", city: "北京市"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCounty(tt.args.province, tt.args.city)
			if err != nil {
				t.Error(err)
			}
			t.Log(got)
		})
	}
}

func TestGetCity(t *testing.T) {
	type args struct {
		province string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"区", args{province: "北京市（京）"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCity(tt.args.province)
			if err != nil {
				t.Error(err)
			}
			t.Log(got)
		})
	}
}
