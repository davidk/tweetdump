package main

import (
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/davecgh/go-spew/spew"
	"net/url"
	"os"
)

const (
	VERSION = "tweetdump version: 0.01\n"
)

var (
	tweetID        int64
	consumerKey    string
	consumerSecret string
	version        bool
)

func init() {
	flag.Int64Var(&tweetID, "tweetid", 0, "Unique ID of a tweet to dump")
	flag.StringVar(&consumerKey, "key", "", "Twitter API consumer key")
	flag.StringVar(&consumerSecret, "secret", "", "Twitter API consumer secret")
	flag.BoolVar(&version, "version", false, "Prints version information")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, VERSION)
		flag.PrintDefaults()
	}

	flag.Parse()

	if version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if tweetID == 0 {
		fmt.Fprintf(os.Stderr, "Error: tweetid must not be zero (usually values above 20 are valid)\n")
		flag.Usage()
		os.Exit(1)
	}

	if consumerKey == "" {
		consumerKey = os.Getenv("TWDUMP_CONSUMER_KEY")
	}

	if consumerSecret == "" {
		consumerSecret = os.Getenv("TWDUMP_CONSUMER_SECRET")
	}

	if consumerKey == "" || consumerSecret == "" {
		fmt.Fprintf(os.Stderr, "Error: key and secret need to be defined.\n")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(VERSION)
}

func dumpTweet(id int64, api *anaconda.TwitterApi) error {
	v := url.Values{}
	v.Set("tweet_mode", "extended")
	fmt.Printf("---- TWEET ID: %v ----\n", tweetID)
	tweet, err := api.GetTweet(id, v)

	if err != nil {
		return fmt.Errorf("Could not get tweet: %s\n", err)
	}

	spew.Dump(tweet)
	fmt.Printf("---- TWEET ID: %v ----\n", tweetID)
	return nil
}

func main() {

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	api := anaconda.NewTwitterApi("", "")

	err := dumpTweet(tweetID, api)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
