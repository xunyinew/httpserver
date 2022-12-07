package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func main() {
	HttpServerStart(8090)
	log.Fatalln(http.ListenAndServe("localhost:8090", nil))
}

func HttpServerStart(port int) {
	log.SetPrefix("info:")
	log.SetFlags(log.Ldate | log.Llongfile)

	http.HandleFunc("/", httpaccessFunc)
	http.HandleFunc("healthz", healthzFunc)
	addr := ":" + strconv.Itoa(port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func healthzFunc(w http.ResponseWriter, req *http.Request) {
	HealthzCode := "200"
	w.Write([]byte(HealthzCode))
}

func httpaccessFunc(w http.ResponseWriter, r *http.Request) {
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			log.Printf("%s=%s", k, v[0])
			w.Header().Set(k, v[0])
		}
		log.Printf("\n\n\n")
		r.ParseForm()
		if len(r.Form) > 0 {
			for k, v := range r.Form {
				log.Printf("%s=%s", k, v[0])
			}
		}
		log.Printf("\n\n\n")
		os.Setenv("VERSION", "JDK version 1.11.0")
		name := os.Getenv("VERSION")
		log.Printf("VERSION Env:", name)
		log.Printf("\n\n\n")

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			fmt.Println("err:", err)
		}
		if net.ParseIP(ip) != nil {
			fmt.Printf("ip==>%s\n", ip)
			log.Printf(ip)
		}
		fmt.Printf("http status code===>%s", http.StatusOK)
		log.Println(http.StatusOK)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server Access,Success!"))
	}
}
