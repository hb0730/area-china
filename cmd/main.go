package main

import (
	"flag"
	"github.com/hb0730/area-china/area"
)

var (
	year     string
	size     int
	filename string
)

func init() {
	flag.StringVar(&year, "year", "2020", "year")
	flag.IntVar(&size, "size", 6, "code min length")
	flag.StringVar(&filename, "filename", "../dist/area.json", "file path and file name")
}
func main() {
	flag.Parse()
	s, err := area.NewSpider(year, size)
	if err != nil {
		panic(err)
	}
	areas, err := s.Start()
	if err != nil {
		panic(err)
	}
	//写
	err = area.WriteJson(filename, areas)
	if err != nil {
		panic(err)
	}
}
