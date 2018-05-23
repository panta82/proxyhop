package proxy

import (
	"net/http"
	"io/ioutil"
	"fmt"
	. "proxyhop/tools"
	"io"
)

type Proxy struct {
	Target string
	Port string
	Verbosity int
}

func (p Proxy) Start() error {
	baseUrl := p.Target
	if p.Target[len(p.Target) - 1] == '/' {
		baseUrl = baseUrl[:len(p.Target) - 1]
	}

	fmt.Printf("Proxying %s ---> %s\n\n", EmText("http://localhost:" + p.Port), EmText(baseUrl))

	requestHandler := func (w http.ResponseWriter, r *http.Request) {
		destUrl := baseUrl + r.URL.Path

		onError := func(err error, statusCode int) {
			if p.Verbosity > 0 {
				PrintError(fmt.Sprintf("%s", destUrl), &err)
			}

			w.WriteHeader(statusCode)
			w.Write([]byte(err.Error()))
		}

		req, err := http.NewRequest(r.Method, destUrl, r.Body)
		if err != nil {
			onError(err, 500)
			return
		}

		for k, vList := range r.Header {
			for _, v := range vList {
				req.Header.Add(k, v)
			}
		}

		// Here is where proxying happens
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			onError(err, resp.StatusCode)
			return
		}

		w.WriteHeader(resp.StatusCode)
		for k, vList := range resp.Header {
			for _, v := range vList {
				w.Header().Add(k, v)
			}
		}

		io.Copy(w, resp.Body)

		if p.Verbosity > 0 {
			fmt.Printf(fmt.Sprintf("%s %s %s %s\n", r.Method, destUrl, MutedText(" --> "), resp.Status))
		}
	}

	return http.ListenAndServe(fmt.Sprintf(":%s", p.Port), http.HandlerFunc(requestHandler))
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