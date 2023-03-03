package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperits/subscription-sets/cve/application/dto"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://services.nvd.nist.gov/rest/json/cves/2.0/?resultsPerPage=100&startIndex=0")
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	var cveList dto.CveDTO
	err = json.Unmarshal(content, &cveList)
	if err != nil {
		fmt.Println("marshal error ", err.Error())
	}
	for _, v := range cveList.Vulnerabilities {
		fmt.Println(v.Info.ID)
	}

}
