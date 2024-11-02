package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.bricklink.com/ajax/clone/catalogifs.ajax?itemid=77556&color=155&ss=KR&cond=N&minqty=10&iconly=0")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))
}
