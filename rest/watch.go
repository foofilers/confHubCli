package rest

import (
	"golang.org/x/net/websocket"
	"strings"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"time"
)

type WatchCommand struct {
	GenericWsCommand
	Data struct {
		     Applications []string `json:"applications"`
	     }
}

type GenericWsCommand struct {
	Op int `json:"operation"`
}

type ApplicationChangedNotification struct {
	Application string `json:"application"`
}

func (cli *ConfHubClient) WatchApp(appNames []string) (chan ApplicationChangedNotification, error) {
	respChan := make(chan ApplicationChangedNotification)
	if err := cli.watchChanges(respChan, appNames); err != nil {
		return nil, err
	}
	return respChan, nil
}

func (cli *ConfHubClient) watchChanges(respChan chan ApplicationChangedNotification, appNames []string) error {
	wsUrl := strings.Replace(cli.Url, "https://", "ws://", 1)
	wsUrl = strings.Replace(wsUrl, "http://", "ws://", 1)
	wsUrl += "/ws"
	logrus.Debugf("Connecting to web socket:%s", wsUrl)
	wsConn, err := websocket.Dial(wsUrl, "", "http://localhost")
	if err != nil {
		logrus.Error("error connecting to websocket",err)
		//restart myself
		time.Sleep(time.Second)
		go cli.watchChanges(respChan,appNames)
		return  err
	}
	watchCmd := &WatchCommand{}
	watchCmd.Op = 0
	watchCmd.Data.Applications = appNames
	jsonCmd, err := json.Marshal(watchCmd)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if _, err := wsConn.Write(jsonCmd); err != nil {
		logrus.Error(err)
		return err
	}
	go func() {
		for {
			appChangedNotif := &ApplicationChangedNotification{}
			if err := json.NewDecoder(wsConn).Decode(appChangedNotif); err != nil {
				logrus.Errorf("error deconding websocket message, restart myself", err)
				//restart myself
				time.Sleep(time.Second)
				cli.watchChanges(respChan, appNames)
				return;
			}
			respChan <- *appChangedNotif
		}

	}()
	return nil
}