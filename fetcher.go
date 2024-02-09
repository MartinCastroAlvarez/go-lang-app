package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func fetchAPI() {
    url := "https://httpbin.org/get"
    response, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error reading body:", err)
        return
    }
    fmt.Println("Response Body:", string(body))
}
