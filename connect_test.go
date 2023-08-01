package test_ci

import (
	"context"
	"testing"
)

func TestOpenGauss(t *testing.T) {
	err := New(context.TODO(), Config{
		URL:      "postgres://root:Kubevela-123@127.0.0.1:5432/kubevela?sslmode=disable&client_encoding=UTF-8&connect_timeout=1",
		Database: "kubevela",
	})
	if err != nil {
		t.Fatal(err)
	}

}
