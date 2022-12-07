package main

import (
	"log"
	"net/http"
)

func listing1() error {
	var client *http.Client
	if tracing {
		// 因为err，所以这里需要用:=，但是这样就会导致这里的client其实是一个内部变量
		client, err := createClientWithTracing()
		if err != nil {
			return err
		}
		// 如果没有这句，其实这里的内部变量client是会被提示unused
		log.Println(client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			return err
		}
		log.Println(client)
	}

	// 没有被定义的client
	_ = client
	return nil
}

func listing2() error {
	var client *http.Client
	if tracing {
		c, err := createClientWithTracing()
		if err != nil {
			return err
		}
		// 这是一种方式，但是缺点是代码看起来很多余
		client = c
	} else {
		c, err := createDefaultClient()
		if err != nil {
			return err
		}
		client = c
	}

	_ = client
	return nil
}

func listing3() error {
	var client *http.Client
	var err error
	if tracing {
		// 因为err在外部定义了，所以直接=，这种方式的最大缺点就是这个err的范围太大了
		// 如果是顺序代码，我建议还是用err1，err2这类的方式（listing2）
		client, err = createClientWithTracing()
		if err != nil {
			return err
		}
	} else {
		client, err = createDefaultClient()
		if err != nil {
			return err
		}
	}

	_ = client
	return nil
}

func listing4() error {
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}
	// 因为这里的err只会走一个分支，所以不需要每个分支都err handling
	if err != nil {
		return err
	}

	_ = client
	return nil
}

var tracing bool

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}
