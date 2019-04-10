package main

import (
	"log"
	"os"
        "time"
        "strconv"

	"github.com/mg64ve/goanda"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key := os.Getenv("OANDA_API_KEY")
	accountID := os.Getenv("OANDA_ACCOUNT_ID")
	oanda := goanda.NewConnection(accountID, key, false)
	timeStart := time.Date(2017, 12, 1, 0, 0, 0, 0, time.UTC)
        timeEnd := time.Date(2017, 12, 1, 1, 0, 0, 0, time.UTC)
        ts := timeStart.Unix()
        te := timeEnd.Unix()
        granularity := "M1"
	history := oanda.GetCandlesFromTime("EUR_USD", strconv.FormatInt(ts,10), strconv.FormatInt(te,10), granularity)
	spew.Dump(history)
}
