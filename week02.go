package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)


func main()  {
	ctx, cancelFunc := context.WithCancel(context.Background())

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		err := httpService(ctx)
		 if err != nil {
		 	cancelFunc()
		 	return err
		 }
		 return nil
	})

	if err := g.Wait();err != nil {
		log.Fatalf("service err %v",err)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	o := <-quit
	log.Fatal("os exit",o.String())
}

func httpService(ctx context.Context) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("week 02")
	})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		return errors.Wrapf(err,"http service err")
	}
	return nil
}

