/*
Package "pushbullet" provides interfaces for Pushbullet HTTP API.

Pushbullet is a web service,
which makes your devices work better together by allowing you to move things between them easily.

The official url: https://www.pushbullet.com/

Currently, this package supports only "pushes" except file.

See the API documentation for the details: https://docs.pushbullet.com/#http
*/
package pushbullet

import (
	"golang.org/x/net/websocket"
	"log"
	"strings"
)

/*
Realtime event stream WebSocket API.
This requires the access token.
The token is found in account settings.

Account settings: https://www.pushbullet.com/account
*/
func SubscribeStream(token string, newPushNotify chan bool) {
	ws, err := websocket.Dial(ENDPOINT_STREAM+token, "", ENDPOINT_STREAM_ORIGIN)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	data := make([]byte, 512)
	for {
		n, err := ws.Read(data)
		if err != nil {
			log.Fatal(err)
		}

		msg := string(data[:n])
		if strings.Count(msg, `"nop"`) > 0 { // ignore "nop".
			continue
		}

		newPushNotify <- true
		break
	}
}
