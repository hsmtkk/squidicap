package embip_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/hsmtkk/squidicap/embip"
	"github.com/stretchr/testify/assert"
)

func TestEmbedIP(t *testing.T) {
	tmplBytes, err := os.ReadFile("./squid.conf.tmpl")
	assert.Nil(t, err)
	want, err := os.ReadFile("./squid.conf")
	assert.Nil(t, err)
	var buf bytes.Buffer
	err = embip.EmbedIP(string(tmplBytes), "192.168.1.2", &buf)
	assert.Nil(t, err)
	got := buf.Bytes()
	assert.Equal(t, want, got)
}
