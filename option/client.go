package main

import (
	"fmt"
	"time"
)

type Client struct {
	url      string
	userName string
	passWord string
	opts     options
}

func NewClient(url string, userName string, passWord string, opts ...Option) *Client {
	//设置一个默认的Options
	op := options{
		charset:     "utf-8",
		readTimeout: time.Duration(20) * time.Second,
	}
	//调用动态传入的参数进行设置值
	for _, o := range opts {
		o(&op)
	}
	return &Client{
		url:      url,
		userName: userName,
		passWord: passWord,
		opts:     op,
	}
}

// options 是一个结构体，用于存储某些配置或选项。
type options struct {
	readTimeout time.Duration
	charset     string
}

// Option 是一个函数类型，它接受一个指向 options 的指针。
type Option func(*options)

func WithReadTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.readTimeout = timeout
	}
}

func WithCharset(charset string) Option {
	return func(o *options) {
		o.charset = charset
	}
}
func main() {
	client := NewClient("http://127.0.0.1:8080", "admin", "123456", WithCharset("www"), WithReadTimeout(time.Second))
	fmt.Println(client)

}
