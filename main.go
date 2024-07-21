package main

import (
	"strings"
)

func main() {

}

type Captcha struct {
	driver Drivertor
	store  Storetor
}

// NewCaptcha creates a captcha instance from driver and store
func NewCaptcha(driver Drivertor, store Storetor) *Captcha {
	return &Captcha{driver: driver, store: store}
}

// Generate generates a random id, base64 image string or an error if any
func (c *Captcha) Generate() (id, b64s, answer string, err error) {
	id, content, answer := c.driver.GenerateIdQuestionAnswer()
	item, err := c.driver.DrawCaptcha(content)
	if err != nil {
		return "", "", "", err
	}
	err = c.store.Set(id, answer)
	if err != nil {
		return "", "", "", err
	}
	b64s = item.EncodeB64string()
	return
}

// Verify by a given id key and remove the captcha value in store,
// return boolean value.
// if you has multiple captcha instances which share a same store.
// You may want to call `store.Verify` method instead.
func (c *Captcha) Verify(id, answer string, clear bool) (match bool) {
	vv := c.store.Get(id, clear)
	//fix issue for some redis key-value string value
	vv = strings.TrimSpace(vv)
	return vv == strings.TrimSpace(answer)
}
