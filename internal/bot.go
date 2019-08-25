package internal

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/spf13/viper"
)

// Bot bot
type Bot struct {
	CrawlingURL string
	Selector    string
	Receiver    string
	Client      *http.Client
}

// Crawl Request a CrawlingURL
func (b *Bot) Crawl() (string, error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var d string
	err := chromedp.Run(ctx,
		chromedp.Navigate(b.CrawlingURL),
		chromedp.InnerHTML(b.Selector, &d),
	)

	if err != nil {
		return "", err
	}

	return d, nil
}

// SendToSlack Send result to clients
func (b *Bot) SendToSlack(m string) error {
	json := fmt.Sprintf(`{
		"icon_emoji": ":ghost:",
		"text": "%s\n%s",
	}`, b.CrawlingURL, m)

	req, err := http.NewRequest("POST", b.Receiver, bytes.NewBufferString(json))
	if err != nil {
		return err
	}

	resp, err := b.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// Run Start crawling bot
func Run() {
	b := &Bot{
		CrawlingURL: viper.Get("crawling_url").(string),
		Selector:    viper.Get("selector").(string),
		Receiver:    viper.Get("receiver").(string),
		Client:      &http.Client{},
	}
	d, err := b.Crawl()

	if err != nil {
		fmt.Printf("Cannot crawl an url: %+v\nError: %+v\n", b.CrawlingURL, err)
		return
	}

	err = b.SendToSlack(d)
	if err != nil {
		fmt.Printf("Cannot send an message\nError: %+v\n", err)
		return
	}

}
