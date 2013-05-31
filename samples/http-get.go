package main

import (
    "net/http"
    "io/ioutil"
    "log"
    "os"
)

func main() {
    resp, err := http.Get("http://stfuawsc.com")
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    os.Stdout.Write(body)
}

