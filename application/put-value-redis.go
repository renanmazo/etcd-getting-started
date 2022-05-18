package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"sync"
	"time"
)

var ctx = context.Background()

func main() {
	cli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	t1 := time.Now()

	var wg sync.WaitGroup

	for i := 0; i < 5000; i++ {
		wg.Add(1)

		i := i
		go func() {
			defer wg.Done()
			createRegisters(&ctx, cli, 100, i)
		}()
	}

	wg.Wait()

	t2 := time.Now()

	fmt.Printf("Elapsed time: " + t2.Sub(t1).String())
}

func createRegisters(ctx *context.Context, client *redis.Client, records int, offset int) {
	for i := 0; i < records; i++ {
		err := client.Set(
			*ctx,
			"sample_key_"+strconv.Itoa(offset)+"_"+strconv.Itoa(i),
			"sample_value_"+strconv.Itoa(i),
			time.Minute*15).Err()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
