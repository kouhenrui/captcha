package main

import "io"

type Storetor interface {
	// Set sets the digits for the captcha id.
	Set(id string, value string) error
	// Get returns stored digits for the captcha id. Clear indicates
	// whether the captcha must be deleted from the store.
	Get(id string, clear bool) string
	//Verify captcha's answer directly
	Verify(id, answer string, clear bool) bool
}
type Drivertor interface {
	//DrawCaptcha draws binary item
	DrawCaptcha(content string) (image Imagetor, err error)
	//GenerateIdQuestionAnswer creates rand id, content and answer
	GenerateIdQuestionAnswer() (id, q, a string)
}

// Item is captcha item interface
type Imagetor interface {
	//WriteTo writes to a writer
	WriteTo(w io.Writer) (n int64, err error)
	//EncodeB64string encodes as base64 string
	EncodeB64string() string
}
