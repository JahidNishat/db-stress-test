package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestLoadBalancerE2E(t *testing.T) {
	for i := 0; i < 100; i++ {
		resp, err := http.Get(fmt.Sprintf("http://localhost:8000?user_id=%v", i))
		if err != nil {
			t.Fatal("error from lb: ", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Error("didn't success: ", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error("resp body error: ", err)
		}
		fmt.Println(string(body))
		resp.Body.Close()
	}
}
