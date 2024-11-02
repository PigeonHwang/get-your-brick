package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SearchResult struct {
	Total_count int `json:"total_count"`
	IdColor     int `json:"idColor"`
	Rpp         int `json:"rpp"`
	Pi          int `json:"pi"`
	List        []struct {
		IdInv                  int    `json:"idInv"`
		StrDesc                string `json:"strDesc"`
		CodeNew                string `json:"codeNew"`
		CodeComplete           string `json:"codeComplete"`
		StrInvImgUrl           string `json:"strInvImgUrl"`
		IdInvImg               int    `json:"idInvImg"`
		TypeInvImg             string `json:"typeInvImg"`
		N4Qty                  int    `json:"n4Qty"`
		IdColorDefault         int    `json:"idColorDefault"`
		TypeImgDefault         string `json:"typeImgDefault"`
		HasExtendedDescription int    `json:"hasExtendedDescription"`
		InstantCheckout        bool   `json:"instantCheckout"`
		MDisplaySalePrice      string `json:"mDisplaySalePrice"`
		MInvSalePrice          string `json:"mInvSalePrice"`
		NSalePct               int    `json:"nSalePct"`
		NTier1Qty              int    `json:"nTier1Qty"`
		NTier2Qty              int    `json:"nTier2Qty"`
		NTier3Qty              int    `json:"nTier3Qty"`
		NTier1DisplayPrice     string `json:"nTier1DisplayPrice"`
		NTier2DisplayPrice     string `json:"nTier2DisplayPrice"`
		NTier3DisplayPrice     string `json:"nTier3DisplayPrice"`
		NTier1InvPrice         string `json:"nTier1InvPrice"`
		NTier2InvPrice         string `json:"nTier2InvPrice"`
		NTier3InvPrice         string `json:"nTier3InvPrice"`
		IdColor                int    `json:"idColor"`
		StrCategory            string `json:"strCategory"`
		StrStorename           string `json:"strStorename"`
		IdCurrencyStore        int    `json:"idCurrencyStore"`
		MMinBuy                string `json:"mMinBuy"`
		StrSellerUsername      string `json:"strSellerUsername"`
		N4SellerFeedbackScore  int    `json:"n4SellerFeedbackScore"`
		StrSellerCountryName   string `json:"strSellerCountryName"`
		StrSellerCountryCode   string `json:"strSellerCountryCode"`
		StrColor               string `json:"strColor"`
	} `json:"list"`
	ReturnCode    int    `json:"returnCode"`
	ReturnMessage string `json:"returnMessage"`
	ErrorTicket   int    `json:"errorTicket"`
	ProcssingTime int    `json:"procssingTime"`
	StrRefNo      string `json:"strRefNo"`
}

func main() {
	sm := map[string][]int{}
	il := []string{
		"itemid=77556&color=155&ss=KR&cond=N&minqty=10&iconly=0",
		"itemid=77556&color=11&ss=KR&cond=N&minqty=10&iconly=0",
		"itemid=3710&color=9&ss=KR&cond=N&minqty=10&iconly=0",
		"itemid=683&color=9&ss=KR&cond=N&iconly=0",
	}
	link := "https://www.bricklink.com/ajax/clone/catalogifs.ajax?"

	for i := 0; i < len(il); i++ {
		resp, err := http.Get(link + il[i])
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}

		// json decode
		var r SearchResult
		err = json.Unmarshal([]byte(string(data)), &r)
		if err != nil {
			panic(err)
		}

		for j := 0; j < len(r.List); j++ {
			_, e := sm[r.List[j].StrStorename]
			if !e {
				sm[r.List[j].StrStorename] = []int{0, 0}
			}
			sm[r.List[j].StrStorename] = []int{sm[r.List[j].StrStorename][0] + 1, sm[r.List[j].StrStorename][1] + convertPrice(r.List[j].MDisplaySalePrice)}
		}
	}

	for k, v := range sm {
		fmt.Println("store name: ", k, "price: ", v[1], "count: ", v[0])
	}
}

func convertPrice(input string) int {
	rs := ""
	for i := 0; i < len(input); i++ {
		if input[i] >= 48 && input[i] <= 57 {
			rs = rs + string(input[i])
		}
	}

	r, _ := strconv.Atoi(rs)
	return r
}
