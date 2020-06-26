package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pencil/route"
	"syscall"
	"time"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/14
 */
func main() {
	srv := &http.Server{
		Addr:    ":8000",
		Handler: route.GetInitRouter(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	//退出信号捕获
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	//创建超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	//关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 1 seconds.")
	}
	log.Println("Server exiting")
}
