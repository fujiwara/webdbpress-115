package main

import (
  "flag"
  "fmt"
  "os"
  "strings"
)

func main() {
	var port int
	var host string
	flag.IntVar(&port, "port", 8080, "port number")
	flag.StringVar(&host, "host", "localhost", "hostname")

  // 設定されているflagについて環境変数があればそれでデフォルトを埋める
	flag.VisitAll(func(f *flag.Flag) {
		name := "MYAPP_" + strings.ToUpper(f.Name)
		if v, e := os.LookupEnv(name); e {
			f.Value.Set(v)
		}
	})
	flag.Parse()
	fmt.Printf("%s:%d\n", host, port)
}
