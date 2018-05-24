package proxy

import (
	"net/http"
	"fmt"
	. "proxyhop/tools"
	"io"
	"github.com/rs/cors"
	"time"
)

type Proxy struct {
	Target string
	Port string
	Verbosity int
	CORSBusting bool
}

func (p Proxy) Start() error {
	baseUrl := p.Target
	if p.Target[len(p.Target) - 1] == '/' {
		baseUrl = baseUrl[:len(p.Target) - 1]
	}

	if p.Verbosity > 0 {
		fmt.Printf("Proxying %s ---> %s\n\n", EmText("http://localhost:" + p.Port), EmText(baseUrl))
	}

	requestHandler := func (w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		destUrl := baseUrl + r.URL.Path
		if r.URL.RawQuery != "" {
			destUrl += "?" + r.URL.RawQuery
		}
		if r.URL.Fragment != "" {
			// This will probably never happen, but can't hurt to add
			destUrl += "#" + r.URL.Fragment
		}

		onError := func(err error) {
			if p.Verbosity > 0 {
				PrintError(fmt.Sprintf("%s", destUrl), &err)
			}

			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}

		req, err := http.NewRequest(r.Method, destUrl, r.Body)
		if err != nil {
			onError(err)
			return
		}

		for k, vList := range r.Header {
			for _, v := range vList {
				if k == "Accept-Encoding" {
					// TODO: Support GZIP
					continue
				}
				req.Header.Add(k, v)
			}
		}

		// Here is where proxying happens
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			onError(err)
			return
		}

		w.WriteHeader(resp.StatusCode)
		for k, vList := range resp.Header {
			for _, v := range vList {
				w.Header().Add(k, v)
			}
		}

		defer resp.Body.Close()
		io.Copy(w, resp.Body)

		if p.Verbosity > 0 {
			now := time.Now()
			elapsed := float64(now.Sub(start)) / float64(time.Second)
			timeArrow := MutedText(fmt.Sprintf(">--(%.2f ms)-->", elapsed))
			fmt.Printf(fmt.Sprintf("[%s] %s %s %s %s\n", MutedText(now.Format("2006-01-02 15:04:05")),
				r.Method, destUrl, timeArrow, resp.Status))
		}
	}

	handlerFunc := http.HandlerFunc(requestHandler)

	if p.CORSBusting {
		tmp := cors.AllowAll().Handler(handlerFunc)
		handlerFunc = (tmp).(http.HandlerFunc)
	}

	return http.ListenAndServe(fmt.Sprintf(":%s", p.Port), handlerFunc)
}
