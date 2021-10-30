package quay

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func Tags() ([]gjson.Result, error) {
	resp, err := http.Get("https://quay.io/api/v1/repository/jetstack/cert-manager-controller/tag/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	log.Info(bodyString)
	value := gjson.Get(bodyString, "tags.#.name")
	println(value.String())
	return value.Array(), nil
}
