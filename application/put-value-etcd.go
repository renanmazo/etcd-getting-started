package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)
import clientv3 "go.etcd.io/etcd/client/v3"

func main() {
	cli, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"0.0.0.0:2379"},
		DialTimeout: 5 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)

	t1 := time.Now()

	var wg sync.WaitGroup

	for i := 1; i < 5000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			createRegisters(&ctx, cli, 100)
		}()
	}

	wg.Wait()

	t2 := time.Now()

	fmt.Printf("Elapsed time: " + t2.Sub(t1).String())
	//readRegisters(&ctx, cli, 1000000)

	cancel()

	defer cli.Close()
}

func createRegisters(ctx *context.Context, client *clientv3.Client, records int) {
	for i := 1; i < records; i++ {
		_, err := client.Put(*ctx, "sample_key_"+strconv.Itoa(records), "sample_value_"+strconv.Itoa(records))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func readRegisters(ctx *context.Context, client *clientv3.Client, records int) {
	for i := 1; i < records; i++ {
		_, err := client.Get(*ctx, "sample_key_"+strconv.Itoa(records))

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}
}
