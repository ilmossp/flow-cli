package fdclient

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

type FlowDeamonInterface struct {
	NewDownloadConfirmed (string)
}

func GenerateSession() (*dbus.Conn, error) {
	session, err := dbus.SessionBus()

	if err != nil {
		return nil, err
	}

	return session, nil
}

func NewDownloadConfirmed(url string) {
	session, err := GenerateSession()

	if err != nil {
		panic(err)
	}

	obj := session.Object("com.github.essmehdi.Flowd", "/com/github/essmehdi/Flowd/Listener")
	var download string

	if callErr := obj.Call("com.github.essmehdi.Flowd.NewDownloadConfirmed", 0, url).Store(&download); callErr != nil {
		fmt.Println("this a method call error")
		panic(callErr)
	}

	fmt.Println(download)

}

func SubscribeToDownloads() {
	conn, err := GenerateSession()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	if err := conn.AddMatchSignal(dbus.WithMatchObjectPath("/com/github/essmehdi/Flowd/Listener"), dbus.WithMatchInterface("com.github.essmehdi.Flowd.Listener")); err != nil {
		panic(err)
	}

	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)
	for v := range c {
		fmt.Println(v)
	}
}
