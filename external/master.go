package external

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/ArpinFidel/go-utils/rest"
	"github.com/ArpinFidel/octafx-scraper/model"
	"github.com/PuerkitoBio/goquery"
)

type MasterHandler interface {
}

func (api *master) Refresh(master model.Master) (model.Master, error) {
	req := rest.Request{
		URL:    fmt.Sprintf("%v/%v", api.url, master.Code),
		Method: http.MethodGet,
	}
	res, code := req.Send()
	if code != 200 {
		return master, fmt.Errorf("status %v{%v}", code, string(res))
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(res))
	if err != nil {
		return master, err
	}

	doc.Find(".account-details__info-item").Each(func(i int, s *goquery.Selection) {
		key, _ := goquery.OuterHtml(s.Find(".account-details__info-key"))
		val, _ := goquery.OuterHtml(s.Find(".account-details__info-value"))
		print(key, val)
	})

	return master, nil
}

type master struct{ url string }

var masterExt *master

func NewMaster(url string) master {
	if masterExt == nil {
		masterExt = &master{url: url}
	}
	return *masterExt
}
