package main

import (
    "fmt"
    "os"
    "sort"
    "text/tabwriter"
    "time"
)

type Event struct {
    Target         string
    Name           string
    MostRecently   time.Duration
    SecondRecently time.Duration
}

var events = []*Event{
    {"Menu", "Start Menu", length("3m38s"), length("3h2m40s")},
    {"Folder", "My Document", length("1m30s"), length("56m10s")},
    {"Button", "Shutdown Button", length("3m38s"), length("2h2m30s")},
    {"Icon", "Trash", length("4m27s"), length("1h20m11s")},
}

func length(s string) time.Duration {
    d, err := time.ParseDuration(s)
    if err != nil {
        panic(s)
    }
    return d
}

func printTracks(events []*Event) {
    const format = "%v\t%v\t%v\t%v\t\n"
    tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
    fmt.Fprintf(tw, format, "Target", "Name", "MostRecently", "SecondRecently")
    fmt.Fprintf(tw, format, "-----", "------", "-----", "----")
    for _, t := range events {
        fmt.Fprintf(tw, format, t.Target, t.Name, t.MostRecently, t.SecondRecently)
    }
    tw.Flush()
}

func main() {

    eventsCopy := make([]*Event, len(events))
    copy(eventsCopy, events)
    fmt.Println("\nCustom:")
    sort.Sort(customSort{events, func(x, y *Event) bool {
        if x.MostRecently != y.MostRecently {
            return x.MostRecently < y.MostRecently
        }
        if x.SecondRecently != y.SecondRecently {
            return x.SecondRecently < y.SecondRecently
        }
        return false
    }})
    printTracks(events)

    fmt.Println("\nStable:")
    sort.Stable(customSort{eventsCopy, func(x, y *Event) bool {
        if x.MostRecently != y.MostRecently {
            return x.MostRecently < y.MostRecently
        }
        // if x.SecondRecently != y.SecondRecently {
        //     return x.SecondRecently < y.SecondRecently
        // }
        return false
    }})
    printTracks(eventsCopy)
}

type customSort struct {
    t    []*Event
    less func(x, y *Event) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
