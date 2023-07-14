package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Block struct {
	Id string `json:"id"`
}

func Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Pong!\n")
}

func Daily(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	url := fmt.Sprintf("%s/boards/%s/blocks/%s/duplicate?asTemplate=false", os.Getenv("BASE_URL"), os.Getenv("BOARD_ID"), os.Getenv("TEMPLATE_BLOCK_ID"))
	headers := map[string]string{
		"Content-Type":     "application/json",
		"Accept":           "application/json",
		"X-Requested-With": "XMLHttpRequest",
		"Authorization":    "Bearer " + os.Getenv("API_TOKEN"),
	}
	data, err := SendRequest("POST", url, nil, headers)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()

	blocks := make([]Block, 0)
	err = json.NewDecoder(data.Body).Decode(&blocks)
	if err != nil {
		log.Fatal(err)
	}

	blockId := blocks[0].Id

	now, err := TimeIn(time.Now(), "Asia/Seoul")
	if err != nil {
		log.Fatal(err)
	}

	title := now.Format("Monday, 02/01/2006")

	url = fmt.Sprintf("%s/boards/%s/blocks/%s", os.Getenv("BASE_URL"), os.Getenv("BOARD_ID"), blockId)
	payload := map[string]interface{}{
		"title": title,
		"updatedFields": map[string]interface{}{
			"properties": map[string]interface{}{
				"a39x5cybshwrbjpc3juaakcyj6e": fmt.Sprintf(`{"from": %d}`, now.UnixMilli()),
				"ae9ar615xoknd8hw8py7mbyr7zo": "a1wj1kupmcnx3qbyqsdkyhkbzgr",
				"ao44fz8nf6z6tuj1x31t9yyehcc": []string{
					"ppmdhd5y138zbpqb3ocwy3r7rc",
					"cqmg9gu3ptyadk31s8wrdxhtsa",
					"rw4cxgm1qpgmzd5ts78dgts38e",
					"3wid7xjx73y63ct3wh3k61ot6w",
					"11ei69wt47fz9kg4rypwx4drse",
					"1jtucy8b7pfu7btahd3k141bgc",
					"51q6synw3jfduxeehd5bofbb7o",
					"pxeg7pzkupgg3pohu58n473zqa",
					"yo5gpm4kotyqfp3t7eggzusfyr",
					"boqnot1rr7gb5nuusfephm9jzh",
				},
			},
		},
	}
	if err != nil {
		log.Fatal(err)
	}

	data, err = SendRequest("PATCH", url, payload, headers)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()

	fmt.Fprint(w, "Success!\n")
}

func Reminder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	reminderType := p.ByName("type")
	payload := map[string]string{}

	switch reminderType {
	case "report":
		payload = map[string]string{
			"text": fmt.Sprintf(
				`@channel Don't forget to fill out the [Daily Report](%s) before the [1 PM KST meeting](%s).`,
				os.Getenv("BOARD_URL"),
				os.Getenv("ZOOM"),
			),
		}
	case "zoom":
		payload = map[string]string{
			"text": fmt.Sprintf("@channel Please join the meeting.\n%s", os.Getenv("ZOOM")),
		}
	default:
		fmt.Fprint(w, "Unknown reminder type.\n")

	}

	data, err := SendRequest("POST", os.Getenv("WEBHOOK"), payload, headers)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()

	fmt.Fprint(w, "Success!\n")
}
