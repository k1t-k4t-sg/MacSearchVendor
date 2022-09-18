package mac

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	pb "MacSearchVendor/pkg/api"

	"github.com/tidwall/gjson"
)

type GrpcServer struct {
	pb.UnimplementedSearchVendorServer 
}

func (g *GrpcServer) GetSearchVendor(ctx context.Context, reg *pb.Mac) (*pb.Vendor, error) {
	jsonFile, err := os.Open("static/mac_base.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	mac_1c := reg.GetQuery()
	if !gjson.ValidBytes(byteValue) {
		fmt.Println("Invalid mac_base.json")
		return &pb.Vendor{Query: mac_1c, Vendor: "nil"}, nil
	} else {
		result := SearchVendor(&byteValue, MacParse(mac_1c))
		return &pb.Vendor{
			Query: MacParse(mac_1c),
			Vendor: result}, nil
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
