package internal

import (
	"net/http"
	"time"
)

func WaitForServer(t *TestContext, host string, seconds int) {
	if seconds == 0 {
		t.Fatal("server did not start")
		return
	}

	httpClient := &http.Client{}
	_, err := httpClient.Get("http://" + host + "/status")
	if err != nil {
		time.Sleep(1 * time.Second)
		WaitForServer(t, host, seconds-1)
	}
}
