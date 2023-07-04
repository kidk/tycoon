package helpers

import (
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

//
// How I can improve this
//
// - Save all the timers to a yaml or JSON
// - Disable completely

type TimingHelper struct {
	timings  map[string]int64
	previous *Stack
	depth    int

	Disabled bool
}

func NewTimingHelper() *TimingHelper {
	return &TimingHelper{
		timings:  make(map[string]int64),
		previous: NewStack(),
		depth:    0,

		Disabled: false,
	}
}

func (th *TimingHelper) Log(message string) {
	if th.Disabled {
		return
	}
	logrus.Printf("%s%s\n", th.tabs(), message)
}

func (th *TimingHelper) Start(key string) {
	if th.Disabled {
		return
	}
	th.previous.Push(key)
	th.timings[key] = time.Now().UnixMicro()
	th.depth += 1
}

func (th *TimingHelper) Stop(key string) {
	if th.Disabled {
		return
	}
	if key != th.previous.Peek() {
		logrus.Printf("Can't stop, unknown key: %s\n", key)
		os.Exit(1)
	}
	th.previous.Pop()
	timing := float64(time.Now().UnixMicro()-th.timings[key]) / 1000.0
	logrus.Printf("%s%s: %.3f\n", th.tabs(), key, timing)
	th.depth -= 1
}

func (th *TimingHelper) tabs() string {
	return strings.Repeat("\t", th.depth-1)
}
