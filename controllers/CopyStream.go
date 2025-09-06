package controllers

import "io"

func CopyStream(source io.Reader, view io.Writer) {
	go io.Copy(view, source)
}
