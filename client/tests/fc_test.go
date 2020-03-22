package tests

import (
	"testing"

	"github.com/JointFaaS/Client-go/client"
)

func TestFcCreate(t *testing.T) {
	_, err := client.NewClient(client.Config{
		Host: "http://127.0.0.1:8080",
	})
    if err != nil {
       t.Errorf("NewClient Error: %s", err.Error())
    }
}