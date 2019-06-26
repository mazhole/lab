package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "html/template"
    "net/http"
    "os"
    "strconv"
)

type Handler struct {
    http.Handler
}

type Todo struct {
    Name string
    IsDone bool
}

func IsNotDone(todo Todo) bool {
    return !todo.IsDone
}

func main() {

    port, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("=> error:", err)
    }

    fmt.Println("=> port:", port)
    fmt.Println()

    if port == 8082 {
        Template1(port)
    }

    if port == 8083 {
        Template2(port)
    }
}

func Template1(port int) {
    t := template.New("main")
    t, _ = t.Parse(
        `<div style="display: inline-block; border: 1px solid #aaa; border-radius: 3px; padding: 30px; margin: 20px;">
            <pre>{{.}}</pre>
        </div>`)

    http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
        path := r.URL.Path

        c := http.Client{}
        resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)

        if err != nil {
            log.Println(err)
            rw.WriteHeader(http.StatusInternalServerError)
            rw.Write([]byte("error"))
            return
        }

        defer resp.Body.Close()

        body, _ := ioutil.ReadAll(resp.Body)

        //rw.WriteHeader(http.StatusOK)
        t.Execute(rw, string(body))
    })

    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func Template2(port int) {

    t, err := template.New("template.html").
        Funcs(template.FuncMap{"IsNotDone": IsNotDone}).
        ParseFiles("template.html")

    if err != nil {
        log.Fatal("=> Can not expand template", err)
        return
    }

    todos := []Todo{
        {"Learn Go", false},
        {"...", false},
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            param := r.FormValue("id")
            index, _ := strconv.ParseInt(param, 10, 0)
            todos[index].IsDone = true
        }

        err := t.Execute(w, todos)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}