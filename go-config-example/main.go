package main

import(
  config "github.com/kayac/go-config"
  "fmt"
)

type Config struct {
  APIToken string `yaml:"api_token"`
  Profile  string `yaml:"profile"`
}

func main() {
  var conf Config
  if err := config.LoadWithEnv(&conf, "config.yaml"); err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println("api_token:", conf.APIToken)
  fmt.Println("profile:", conf.Profile)
}
