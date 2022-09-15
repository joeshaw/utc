package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var tzmap = map[string]string{
	"":    "Local",
	"UTC": "UTC",
	"EST": "America/New_York",
	"EDT": "America/New_York",
	"CST": "America/Chicago",
	"CDT": "America/Chicago",
	"MST": "America/Denver",
	"MDT": "America/Denver",
	"PST": "America/Los_Angeles",
	"PDT": "America/Los_Angeles",
}

var formats = []string{
	"15:04",
	"3:04PM",

	"15:04:05",
	"3:04:05PM",
}

func main() {
	var local bool
	flag.BoolVar(&local, "l", false, "output local time")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-l] <time> [timezone]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	when := strings.ToUpper(args[0])

	var tz string
	if len(args) > 1 {
		tz = args[1]
	}

	tzname, ok := tzmap[strings.ToUpper(tz)]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown timezone: %s\n", tz)
		os.Exit(1)
	}

	loc, err := time.LoadLocation(tzname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	now := time.Now()

	for _, f := range formats {
		t, err := time.Parse(f, when)
		if err == nil {
			t = time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)

			if local {
				t = t.Local()
			} else {
				t = t.UTC()
			}

			if t.Second() == 0 {
				fmt.Println(t.Format("15:04 MST"))
			} else {
				fmt.Println(t.Format("15:04:05 MST"))
			}
			return
		}
	}

	fmt.Fprintln(os.Stderr, "couldn't parse time")
	os.Exit(1)
}
