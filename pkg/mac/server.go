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

var ByteValue []byte

func init(){
	jsonFile, err := os.Open("static/mac_base.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	ByteValue, _ = ioutil.ReadAll(jsonFile)
	if !gjson.ValidBytes(ByteValue) {
		fmt.Println("Invalid mac_base.json")
	} else {
		fmt.Println("Succes mac_base.json")
	}
}

type GrpcServer struct {
	pb.UnimplementedSearchVendorServer 
}

func (g *GrpcServer) GetSearchVendor(ctx context.Context, reg *pb.Mac) (*pb.Vendor, error) {
	mac_1c := reg.GetQuery()
	return &pb.Vendor{
		Query: MacParse(mac_1c),
		Vendor: SearchVendor(&ByteValue, MacParse(mac_1c))}, nil
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
