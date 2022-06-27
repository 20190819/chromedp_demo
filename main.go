package main

import (
	"context"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"io/fs"
	"log"
	"os"
	"time"
)

var htmlContent string

const targetURL = "https://www.ggcx.com/main/globalRegistrar"
const maxPage = 5

func main() {
	fmt.Println("start chromedp ...")
	_, err := Spider(maxPage)
	if err != nil {
		errors.Wrap(err, "internal err")
	}
}

func Spider(page int) (context.Context, error) {
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
		chromedp.ActionFunc(func(ctx context.Context) error {
			for i := 0; i <= page; i++ {
				if i > 0 {
					chromedp.OuterHTML(`//*[@id="main"]/div/div[3]/div[5]/table`, &htmlContent).Do(ctx)
					LoadDoc(htmlContent, i)
				}
				chromedp.Sleep(3 * time.Second).Do(ctx)
				chromedp.Click(`//*[@id="main"]/div/div[3]/div[5]/div[3]/div/div[2]`).Do(ctx)

			}
			return nil
		}),
	)
	return ctx, err
}

func LoadDoc(html string, number int) error {
	fmt.Println("loadDoc", len(html))
	var file *os.File
	var filePath = fmt.Sprintf("./reg_%d.html", number)
	file, err := os.OpenFile(filePath, os.O_CREATE, fs.ModePerm)
	if err != nil {
		fmt.Println(errors.Wrap(err, "打开文件报错"))
	}
	defer func() {
		file.Close()
		os.Remove(filePath)
	}()
	_, err = file.WriteString(html)
	if err != nil {
		fmt.Println(errors.Wrap(err, "写入文件报错"))
		os.Exit(-1)
	}
	doc, _ := htmlquery.LoadDoc(filePath)
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
