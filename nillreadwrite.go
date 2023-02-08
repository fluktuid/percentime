package main

type nilReader struct{}
type nilWriter struct{}

func (nilReader) Read(p []byte) (n int, err error)  { return len(p), nil }
func (nilWriter) Write(p []byte) (n int, err error) { return len(p), nil }
