package main

import (
	"Area-China/area"
	"flag"
)

var (
	year string
	size int
)

func init() {
	flag.StringVar(&year, "year", "2020", "year")
	flag.IntVar(&size, "size", 6, "code min length")
}
func main() {
	flag.Parse()
	area.Start(year, size)
}
