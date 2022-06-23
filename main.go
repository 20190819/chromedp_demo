package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/davecgh/go-spew/spew"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	connDB()
	err := cdpHandler()
	if err != nil {
		fmt.Println("cdp run err", err)
	}
	fmt.Println("main end")
}

func connDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		"root", "888888", "crawl")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		os.Exit(-1)
	}
	return db
}

func createData() {
	// todo
}

var trs []*cdp.Node

func cdpHandler() error {
	fmt.Println("cdpHandler")
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))
	alloc, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(
		alloc,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// 开启一个Chrome
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(`https://www.ggcx.com/main/globalRegistrar`),
		chromedp.WaitVisible(`div[class="tabb"]`),
		chromedp.Nodes(`.//tr`, &trs),
	)
	for _,tr:=range trs{
		spew.Dump(tr)
	}
	if err != nil {
		return err
	}
	return nil
}
