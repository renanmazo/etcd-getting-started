package main

import (
	"context"
	"fmt"
	"time"
)
import clientv3 "go.etcd.io/etcd/client/v3"

func main() {
	cli, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"172.16.1.2:2379"},
		DialTimeout: 5 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5000)
	_, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	defer cli.Close()
}
