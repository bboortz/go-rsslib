package rsslib

import ()

type RssItem struct {
	Id          int64  `json:"Id"`
	Uuid        string `json:"Uuid"`
	Channel     string `json:"Channel"`
	Title       string `json:"Title"`
	Link        string `json:"Link"`
	Description string `json:"Description"`
	Thumbnail   string `json:"Thumbnail"`
	PublishDate string `json:"PublishDate"`
	UpdateDate  string `json:"UpdateDate"`
}

type RssItems []RssItem

func (s RssItems) Len() int64 {
	return int64(len(s))
}
