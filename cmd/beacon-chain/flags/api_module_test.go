package flags

import (
	"testing"

	"github.com/sila-chain/Sila-Consensus-Core/v7/testing/assert"
)

func TestEnableHTTPSilaAPI(t *testing.T) {
	assert.Equal(t, true, EnableHTTPSilaAPI("sila"))
	assert.Equal(t, true, EnableHTTPSilaAPI("sila,foo"))
	assert.Equal(t, true, EnableHTTPSilaAPI("foo,sila"))
	assert.Equal(t, true, EnableHTTPSilaAPI("sila,sila"))
	assert.Equal(t, true, EnableHTTPSilaAPI("PrYsM"))
	assert.Equal(t, false, EnableHTTPSilaAPI("foo"))
	assert.Equal(t, false, EnableHTTPSilaAPI(""))
}

func TestEnableHTTPEthAPI(t *testing.T) {
	assert.Equal(t, true, EnableHTTPEthAPI("eth"))
	assert.Equal(t, true, EnableHTTPEthAPI("eth,foo"))
	assert.Equal(t, true, EnableHTTPEthAPI("foo,eth"))
	assert.Equal(t, true, EnableHTTPEthAPI("eth,eth"))
	assert.Equal(t, true, EnableHTTPEthAPI("EtH"))
	assert.Equal(t, false, EnableHTTPEthAPI("foo"))
	assert.Equal(t, false, EnableHTTPEthAPI(""))
}

func TestEnableApi(t *testing.T) {
	assert.Equal(t, true, enableAPI("foo", "foo"))
	assert.Equal(t, true, enableAPI("foo,bar", "foo"))
	assert.Equal(t, true, enableAPI("bar,foo", "foo"))
	assert.Equal(t, true, enableAPI("foo,foo", "foo"))
	assert.Equal(t, true, enableAPI("FoO", "foo"))
	assert.Equal(t, false, enableAPI("bar", "foo"))
	assert.Equal(t, false, enableAPI("", "foo"))
}
