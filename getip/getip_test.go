package getip_test

import (
	"testing"

	"github.com/hsmtkk/squidicap/getip"
	"github.com/stretchr/testify/assert"
)

func TestGetIP(t *testing.T) {
	want := "93.184.216.34"
	got, err := getip.GetIP("www.example.com")
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
