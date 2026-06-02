package main

import (
	"context"
	"net/http"
	"io"
	"encoding/xml"
	"html"
)

type RSSFeed struct {
	Channel struct {
		Title		string		`xml:"title"`
		Link		string		`xml:"link"`
		Description	string		`xml:"description"`
		Item		[]RSSItem	`xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title		string	`xml:"title"`
	Link		string	`xml:"link"`
	Description	string	`xml:"description"`
	PubDate		string	`xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	f := RSSFeed{}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)

	if err != nil {
		return &f, err
	}

	req.Header.Set("User-Agent", "gator")

	c := &http.Client {}

	res, err := c.Do(req)

	if err != nil {
		return &f, err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)

	if err != nil {
		return &f, err
	}

	if err = xml.Unmarshal(b, &f); err != nil {
		return &f, err
	}

	f.Channel.Title = html.UnescapeString(f.Channel.Title)
	f.Channel.Description = html.UnescapeString(f.Channel.Description)

	for i, el := range f.Channel.Item {
		f.Channel.Item[i].Title = html.UnescapeString(el.Title)
		f.Channel.Item[i].Description = html.UnescapeString(el.Description)
	}
	
	return &f, nil
}
