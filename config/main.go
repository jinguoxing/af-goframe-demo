package main


import (
    "flag"
    "fmt"
    "log"

    "af-go-frame/core/config"
    "af-go-frame/core/config/file"
)

var flagconf string

func init() {
    flag.StringVar(&flagconf, "conf", "config.yaml", "config path, eg: -conf config.yaml")
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

    var vv struct{
        Consul struct{
            Address string `json:"address"`
            Scheme  string `json:"scheme"`
        }`json:"consul"`
    }

    if err := c.Scan(&vv);err !=nil{
        panic(err)
    }

    log.Printf("config: %+v", vv)


    // Get a value associated with the key
    name, err := c.Value("service.name").String()
    if err != nil {
        panic(err)
    }
    log.Printf("service: %s", name)

    // watch key
    if err := c.Watch("service.name", func(key string, value config.Value) {
        log.Printf("config changed: %s = %v\n", key, value)
    }); err != nil {
        panic(err)
    }

    <-make(chan struct{})
}