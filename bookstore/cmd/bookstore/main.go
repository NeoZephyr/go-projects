package main

import (
	_ "bookstore/internal/store"
	"bookstore/server"
	"bookstore/store/factory"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s, err := factory.New("mem")

	if err != nil {
		panic(err)
	}

	srv := server.New(":8000", s)
	errChan, err := srv.ListenAndServe()

	if err != nil {
		log.Println("book server start failed:", err)
		return
	}

	log.Println("book server start ok")

	// 监视系统信号实现 http 服务实例的优雅退出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <- errChan:
		log.Println("book server run failed:", err)
		return
	case <- c:
		log.Println("book server is exiting...")
		ctx, callback := context.WithTimeout(context.Background(), time.Second)
		defer callback()
		err = srv.Shutdown(ctx)
	}

	if err != nil {
		log.Println("book server exit error:", err)
		return
	}

	log.Println("book server exit ok")
}
