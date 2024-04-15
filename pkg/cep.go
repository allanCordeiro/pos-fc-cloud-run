package pkg

import (
	"fmt"
	"regexp"
	"strings"
)

type Cep struct {
	value string
}

func NewCep(cep string) *Cep {
	c := &Cep{value: cep}
	c.format()
	return c
}

func (c *Cep) IsCepCodeValid() bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(c.value)
}

func (c *Cep) GetCode() string {
	return c.value
}

func (c *Cep) format() {
	c.value = strings.Replace(c.value, "-", "", 1)
	c.value = fmt.Sprintf("%08s", c.value)
}
