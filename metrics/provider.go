package metrics

import (
	"os"
	"time"

	"github.com/swagchat/rtm-api/services"
	"github.com/swagchat/rtm-api/utils"
)

type Metrics struct {
	Hostname          string                    `json:"hostname"`
	AllCount          int                       `json:"allCount"`
	UserCount         int                       `json:"userCount"`
	RoomCount         int                       `json:"roomCount"`
	EachUserCount     map[string]int            `json:"eachUserCount,omitempty"`
	EachRoomCount     map[string]int            `json:"eachRoomCount,omitempty"`
	EachRoomUserCount map[string]map[string]int `json:"eachRoomUserCount,omitempty"`
	Timestamp         string                    `json:"timestamp"`
}

type Provider interface {
	Run()
}

func MetricsProvider() Provider {
	c := utils.Config()

	var provider Provider
	switch c.Metrics.Provider {
	case "":
		provider = &NotuseProvider{}
	case "stdout":
		provider = &StdoutProvider{}
	case "elasticsearch":
		provider = &ElasticsearchProvider{}
	}
	return provider
}

func makeMetrics(nowTime time.Time) *Metrics {
	c := utils.Config()
	m := &Metrics{}

	hostname, _ := os.Hostname()
	m.Hostname = hostname

	srv := services.GetServer()
	users := srv.Connection.Users()
	rooms := srv.Connection.Rooms()
	m.AllCount = srv.Connection.ConnectionCount()
	m.UserCount = len(users)
	m.RoomCount = len(rooms)

	m.Timestamp = nowTime.Format(time.RFC3339)

	if c.Metrics.Verbose {
		// TODO
	}

	return m
}

func exec(fn func(), interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fn()
		}
	}
}