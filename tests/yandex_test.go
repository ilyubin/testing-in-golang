package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testing-in-golang/fw"
)

func Test_yandex_search(t *testing.T) {
	url := fw.Url()
	assert.Equal(t, url, "https://ya.ru/")

	fw.Type(".input__input", "Agouti")
	fw.Click(".search2__button button")
	text := fw.GetText("li.serp-item:first-of-type .organic__url-text")
	assert.Contains(t, text, "Агути")
}
