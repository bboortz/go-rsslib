package rsslib

import ()

type RssItem struct {
	Id          int64  `json:"id"`
	Channel     string `json:"Channel"`
	Title       string `json:"Title"`
	Link        string `json:"Link"`
	Description string `json:"Description"`
	Thumbnail   string `json:"Thumbnail"`
}

type RssItems []RssItem
