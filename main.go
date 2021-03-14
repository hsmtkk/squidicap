package main

import (
	"log"
	"os"

	"github.com/hsmtkk/squidicap/embip"
	"github.com/hsmtkk/squidicap/getip"
)

func main() {
	tmplPath := getEnvVar("TEMPLATE_PATH")
	confPath := getEnvVar("CONF_PATH")
	icapHostName := getEnvVar("ICAP_HOST_NAME")
	icapHostIP, err := getip.GetIP(icapHostName)
	if err != nil {
		log.Fatal(err)
	}
	if err := embip.EmbedIPFile(tmplPath, icapHostIP, confPath); err != nil {
		log.Fatal(err)
	}
}

func getEnvVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("environment variable is not defined; %s", key)
	}
	return val
}
