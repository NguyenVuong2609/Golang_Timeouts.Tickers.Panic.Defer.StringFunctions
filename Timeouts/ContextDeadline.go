package Timeouts

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func callExternalService(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	fmt.Println("response hello")
}
func execCall() error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(req.Context(), 2*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	_, err = client.Do(req)
	return err
}
func ContextDeadline() {
	go func() {
		http.HandleFunc("/", callExternalService)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(1 * time.Second) // wait for server to run
	// call server
	err := execCall()
	if errors.Is(err, context.DeadlineExceeded) {
		log.Println("ContextDeadlineExceeded: true")
	}
	if os.IsTimeout(err) {
		log.Println("IsTimeoutError: true")
	}
	if err != nil {
		log.Fatal(err)
	}
}
