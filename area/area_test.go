package area

import "testing"

func TestWriteJson(t *testing.T) {
	type args struct {
		filename string
		year     string
		area     []Area
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "empty data", args: args{year: "2022", filename: "../dist/area.json", area: []Area{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteJson(tt.args.filename, tt.args.area)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
