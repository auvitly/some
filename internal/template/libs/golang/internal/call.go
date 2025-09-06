package internal

import "os"

type CallLocation struct {
	File    string
	Package string
	Line    string
	PWD     string
}

func GetCallLocation() CallLocation {
	return CallLocation{
		File:    os.Getenv("GOFILE"),
		Package: os.Getenv("GOPACKAGE"),
		Line:    os.Getenv("GOLINE"),
		PWD:     os.Getenv("PWD"),
	}
}
