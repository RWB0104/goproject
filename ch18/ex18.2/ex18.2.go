package main

import "../fedex/fedex.go"

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {

}
