package main

import (
	"context"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"log"
	"os"
	"time"
)

var html string

//var pages []string
var fPath = "./reg.html"

const targetURL = "https://www.ggcx.com/main/globalRegistrar"

func main() {
	fmt.Println("start chromedp ...")
	_, err := Spider()
	if err != nil {
		errors.Wrap(err, "internal err")
	}
}

func Spider() (context.Context, error) {
	// 禁用chrome headless
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(targetURL),
		chromedp.WaitVisible(`//*[@id="main"]/div/div[3]/div[5]/table`),
		chromedp.Sleep(time.Second),
		chromedp.OuterHTML(`//*[@id="main"]/div/div[3]/div[5]/table`, &html),
		chromedp.ActionFunc(func(ctx context.Context) error {
			for i := 1; i <= 3; i++ {
				LoadDoc(html)
				chromedp.Sleep(5 * time.Second).Do(ctx)
				chromedp.Click(`//*[@id="main"]/div/div[3]/div[5]/div[3]/div/div[2]`).Do(ctx)
				fmt.Println(`page+1`)

			}
			return nil
		}),
	)
	return ctx, err
}

func LoadDoc(html string) error {
	fmt.Println("loadDoc", len(html))
	var file *os.File
	defer func() {
		file.Close()
		os.Remove(fPath)
	}()
	if _, err := os.Stat(fPath); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			file, _ = os.Create(fPath)
		}
	}
	n, err := file.WriteString(html)
	if err != nil {
		fmt.Println("write err:", err)
		os.Exit(-1)
	}
	fmt.Println("write n:",n)
	doc, _ := htmlquery.LoadDoc(fPath)

	trs := htmlquery.Find(doc, ".//tr")
	for _, tr := range trs {
		tds := htmlquery.Find(tr, ".//td")
		for _, td := range tds {
			fmt.Println(htmlquery.InnerText(td))
		}
		fmt.Println("=======")
	}
	return nil
}
