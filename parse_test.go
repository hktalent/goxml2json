package xml2json

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Product struct {
	ID      int     `json:"id"`
	Price   float64 `json:"price"`
	Deleted bool    `json:"deleted"`
}

type StringProduct struct {
	ID      string `json:"id"`
	Price   string `json:"price"`
	Deleted string `json:"deleted"`
}

type MixedProduct struct {
	ID      string  `json:"id"`
	Price   float64 `json:"price"`
	Deleted string  `json:"deleted"`
}

const (
	productString = `
	<?xml version="1.0" encoding="UTF-8"?>	
		<id>42</id>
		<price>13.32</price>
		<deleted>true</deleted>
		`
)

func TestAllJSTypeParsing(t *testing.T) {
	xml := strings.NewReader(productString)
	jsBuf, err := Convert(xml, WithTypeConverter(Bool, Int, Float))
	assert.NoError(t, err, "could not parse test xml")
	product := Product{}
	err = json.Unmarshal(jsBuf.Bytes(), &product)
	assert.NoError(t, err, "could not unmarshal test json")
	assert.Equal(t, 42, product.ID, "price should match")
	assert.Equal(t, 13.32, product.Price, "price should match")
	assert.Equal(t, true, product.Deleted, "price should match")
}

func TestStringParsing(t *testing.T) {
	xml := strings.NewReader(productString)
	jsBuf, err := Convert(xml)
	assert.NoError(t, err, "could not parse test xml")
	product := StringProduct{}
	err = json.Unmarshal(jsBuf.Bytes(), &product)
	assert.NoError(t, err, "could not unmarshal test json")
	assert.Equal(t, "42", product.ID, "price should match")
	assert.Equal(t, "13.32", product.Price, "price should match")
	assert.Equal(t, "true", product.Deleted, "price should match")
}

func TestMixedParsing(t *testing.T) {
	xml := strings.NewReader(productString)
	jsBuf, err := Convert(xml, WithTypeConverter(Float))
	assert.NoError(t, err, "could not parse test xml")
	product := MixedProduct{}
	err = json.Unmarshal(jsBuf.Bytes(), &product)
	assert.NoError(t, err, "could not unmarshal test json")
	assert.Equal(t, "42", product.ID, "price should match")
	assert.Equal(t, 13.32, product.Price, "price should match")
	assert.Equal(t, "true", product.Deleted, "price should match")
}
