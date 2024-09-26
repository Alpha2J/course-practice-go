package main

import "bytes"

type ReadWrite interface {
	Read(b bytes.Buffer) bool
	Write(b bytes.Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}

type MyFile struct {
}

func (mf *MyFile) Read(b bytes.Buffer) bool {
	return false
}

func (mf *MyFile) Write(b bytes.Buffer) bool {
	return false
}

func (mf *MyFile) Lock() {
}

func (mf *MyFile) Unlock() {
}

func (mf *MyFile) Close() {
}

// 1. 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样
func t1() {
	var myFile File = &MyFile{}
	myFile.Close()
}

func main() {
	t1()
}
