package main

import (
	"fmt"
	"log"
)

func main() {
	var p Publisher
	p = newPublisher()
	p.notification(msg, "hello")

	r := newReader(ReaderId: "0001")
	r1 := newReader(ReaderId: "0002")
	r2 := newReader(ReaderId: "0003")
	p.addReader(r)
	p.addReader(r1)
	p.notification(msg, "you are subscribed")

	p.removeReader(r.ReaderId("0003"))
	p.notification(msg, "your subscription is expired")
}

type Publisher interface {
	addReader(reader Reader)
	removeReader(readerId string)
	notification(msg string)
}

type Reader interface {
	id() string
	feedback(msg string)
}

//Observable

type publisher struct {
	readers map[string]Reader
}

func newPublisher() publisher {
	return publisher{readers: make(map[string]Reader)}
}

func (p publisher) addReader(reader Reader) {
	p.readers[reader.id()] = reader
}

func (p publisher) removeReader(readerId string) {
	delete(p.readers, readerId)
}

func (p publisher) notification(msg string) {
	for Reader := range p.readers {
		Reader.feedback(msg)
	}
}

//Obsrever

type reader struct {
	ReaderId string
}

func newReader(readerId string) reader {
	return reader{ReaderId: readerId}
}

func (r reader) id() string {
	return r.ReaderId
}

func (r reader) feedback(msg string) {
	log.Printf(format: "your Id %r - received to: %r" , r.ReaderId, msg)
}
