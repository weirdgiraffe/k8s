//
// main.go
// Copyright (C) 2017 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	server := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			buf := new(bytes.Buffer)
			fmt.Fprintf(buf, "%s %s\n", r.Method, r.URL.Path)
			for k, v := range r.Header {
				fmt.Fprintf(buf, "%s: %s\n", k, strings.Join(v, ","))
			}
			body, _ := ioutil.ReadAll(r.Body)
			fmt.Fprintf(buf, "\n%s\n", body)
			os.Stderr.Write(buf.Bytes())
			w.Write(buf.Bytes())
		}),
	}
	server.ListenAndServe()
}
