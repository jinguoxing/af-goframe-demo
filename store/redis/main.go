package redis

import (
    "flag"
    "fmt"
    data "google.golang.org/genproto/googleapis/analytics/data/v1beta"
    "log"

    "af-go-frame/core/config"
    "af-go-frame/core/config/file"
)

var flagconf string

func init() {
    flag.StringVar(&flagconf, "conf", "config.yaml", "config path, eg: -conf config.yaml")
}

type Bootstrap struct {

    Data *Data
}

type Data struct {


}

func main() {
    flag.Parse()

    fmt.Println(flagconf)
    c := config.New(
        config.WithSource(
            file.NewSource(flagconf),
        ),
    )
    if err := c.Load(); err != nil {
        panic(err)
    }

    var bc conf.Bootstrap


    // Defines the config JSON Field
    var v struct {
        Service struct {
            Name    string `json:"name"`
            Version string `json:"version"`
        } `json:"service"`
    }

    // Unmarshal the config to struct
    if err := c.Scan(&v); err != nil {
        panic(err)
    }
    log.Printf("config: %+v", v)


    // Get a value associated with the key
    name, err := c.Value("service.name").String()
    if err != nil {
        panic(err)
    }
    log.Printf("service: %s", name)



    <-make(chan struct{})