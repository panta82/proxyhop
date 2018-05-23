package proxy

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

type Proxy struct {
	Target string
	Port string
}

func (p Proxy) SendRequest() error {
	resp, err := http.Get(p.Target)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(body)
	return nil
}