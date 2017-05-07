# tweetdump

### Installation

###### Requires a [golang distribution](https://golang.org/doc/install).

    $ go get github.com/davidk/tweetdump

### Updating

    $ go get -u github.com/davidk/tweetdump

### About

tweetdump unpacks [Anaconda's](https://github.com/ChimeraCoder/anaconda) Tweet structures using the [spew pretty printer](https://github.com/davecgh/go-spew/). Use this tool to debug data from Twitter, or to see all the data that can be had when using Anaconda.

### Usage

    # -- The ConsumerKey and SecretToken options need to be defined
    # -- in order to talk to the Twitter API
    # -tweetid is the number at the back of the status/ url
    # IE: https://twitter.com/jack/status/20 (-tweetid would be 20)

    $ bin/tweetdump -help
    tweetdump version: 0.01
      -key string
        	Twitter API consumer key
      -secret string
        	Twitter API consumer secret
      -tweetid int
        	Unique ID of a tweet to dump
      -version
        	Prints version information

### Environmental Variables

tweetdump supports the following environmental variables to set the consumer and secret keys. CLI invocations should override these:

    TWDUMP_CONSUMER_KEY
    TWDUMP_CONSUMER_SECRET

To get a secret token and consumer key, run through the gamut of creating a new
application [here](https://apps.twitter.com/).
