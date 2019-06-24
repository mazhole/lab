package main

import (
    "flag"
    "io/ioutil"
    "log"
    "net/http"
)

type Handler struct {
    http.Handler
}

func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    path := r.URL.Path

    c := http.Client{}
    resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)

    if err != nil {
        log.Println(err)
    }

    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    rw.WriteHeader(http.StatusOK)
    rw.Write(body)
}

func main() {
    var port string
    flag.StringVar(&port, "port", "", "Usage")
    flag.Parse()

    // Example 1
    if port == "8081" {
        http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
            path := r.URL.Path

            c := http.Client{}
            resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)

            if err != nil {
                log.Println(err)
            }

            defer resp.Body.Close()

            body, _ := ioutil.ReadAll(resp.Body)

            rw.WriteHeader(http.StatusOK)
            rw.Write(body)
        })
        http.ListenAndServe(":8081", nil)
        return
    }

    handler := Handler{}

    // Example 2
    if port == "8082" {
        http.HandleFunc("/", handler.ServeHTTP)
        http.ListenAndServe(":8082", nil)
        return
    }

    // Example 3
    if port == "8083" {
        http.Handle("/", handler)
        http.ListenAndServe(":8083", nil)
        return
    }

    // Example 4
    http.ListenAndServe(":8084", handler)
}