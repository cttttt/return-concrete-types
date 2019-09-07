package main

import "returnconcrete/proxy"
import "returnconcrete/consumer"

func main() {
	proxy := proxy.New()
	consumer := consumer.New(proxy)
	consumer.Do()
}
