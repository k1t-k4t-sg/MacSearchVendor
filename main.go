package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
    "os"
)


func main() {

	jsonFile, err := os.Open("mac_base.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
	mac_1c := "50ff.2009.e759"
	if !gjson.ValidBytes(byteValue) {
		fmt.Println("Invalid mac_base.json")
	}else{
		result := SearchVendor(&byteValue, MacParse(mac_1c))
	    fmt.Println("	\n", mac_1c, result)
		fmt.Println("")
		getMacaddress(mac_1c)
	}
}

func SearchVendor(json *[]byte, unique_mac string) string {
	value := gjson.GetBytes(*json, `root.row.#(Assignment=="`+unique_mac+`").Name`)
	if !value.Exists() {
		return "Vendor not found"
	} else {
		return strings.ToUpper(value.String())
	}
}

func MacParse(mac string) string {
	return strings.NewReplacer(":", "", ".", "").Replace(strings.ToUpper(mac))[:6]
}