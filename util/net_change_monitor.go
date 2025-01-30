package util

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

type Subscriber func(signal *dbus.Signal)

type NetworkMonitor struct {
	subscribers []Subscriber
	conn        *dbus.Conn
}

func NewNetworkMonitor() (*NetworkMonitor, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to system bus: %w", err)
	}

	err = conn.AddMatchSignal(
		dbus.WithMatchInterface("org.freedesktop.NetworkManager"),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to add match signal: %w", err)
	}

	return &NetworkMonitor{
		subscribers: []Subscriber{},
		conn:        conn,
	}, nil
}

func (nm *NetworkMonitor) Subscribe(subscriber Subscriber) {
	nm.subscribers = append(nm.subscribers, subscriber)
}

func (nm *NetworkMonitor) StartListening() {
	c := make(chan *dbus.Signal, 10)
	nm.conn.Signal(c)

	go func() {
		for signal := range c {
			for _, subscriber := range nm.subscribers {
				subscriber(signal)
			}
		}
	}()
}
