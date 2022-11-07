package main

import (
	"fmt"
	"log"
	"os"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2022-01-01", "2022-11-06", quote.Daily, true)
	fmt.Println(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
	mva := talib.Ema(spy.Close, 9)
	fmt.Println(mva)

	file, err := os.Create("spy.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(spy.CSV())
}
