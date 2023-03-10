package integration

import (
	"fmt"
	"net/http"
	"time"
)

func awaitForServer(port int) {
	for {
		resp, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
		if err != nil {
			continue
		}
		if resp.StatusCode == http.StatusNotFound {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}
