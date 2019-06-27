package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    _ "net/http/pprof"
    "strconv"
    "io/ioutil"
)

type Todo struct {
    Name string `json:"name"`
    Done bool   `json:"done"`
}

func main() {
    todos := []Todo{
        {"Learn Go", false},
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fileContents, err := ioutil.ReadFile("index.html")
        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusNotFound)
            return
        }
        w.Write(fileContents)
    })

    http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("=> Request:", r.URL.Path)
        defer r.Body.Close()

        switch r.Method {
        case http.MethodGet:
            data, _ := json.Marshal(todos)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            w.Write(data)
        case http.MethodPost:
            decoder := json.NewDecoder(r.Body)
            todo := Todo{}
            err := decoder.Decode(&todo)
            if err != nil {
                log.Println(err)
                w.WriteHeader(http.StatusBadRequest)
                return
            }
            todos = append(todos, todo)
        case http.MethodPut:
            id := r.URL.Path[len("/todos/"):]
            index, _ := strconv.ParseInt(id, 10, 0)
            todos[index].Done = true
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
        }
    })

    http.ListenAndServe(":8081", nil)
}
