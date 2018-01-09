package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const baseUrl string = "https://www.lox24.eu/API/httpsms.php?konto=%s&password=%s&service=%s&text=%s&from=%s&to=%s"

type GrafanaMessage struct {
	Title    string
	Message  string
	ImageUrl string
	RuleUrl  string
}

func main() {
	user_id := os.Getenv("USER_ID")
	password := os.Getenv("PASSWORD")
	service := os.Getenv("SERVICE")
	from := os.Getenv("FROM")
	recipients := strings.Split(os.Getenv("RECIPIENTS"), ",")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var m GrafanaMessage
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			log.Println(err)
		}

		for _, recipient := range recipients {
			reqUrl := fmt.Sprintf(baseUrl, user_id, password, service, url.QueryEscape(m.Title), from, recipient)
			log.Println(reqUrl)

			resp, err := http.Get(reqUrl)
			if err != nil {
				log.Println(err)
			} else {
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				bodyString := string(body)
				log.Println(bodyString)
			}
		}
	})

	log.Fatal(http.ListenAndServe(":4444", nil))
}
