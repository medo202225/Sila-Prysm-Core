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

func TestEnableHTTPSilaCompatAPI(t *testing.T) {
	assert.Equal(t, true, EnableHTTPSilaCompatAPI("eth"))
	assert.Equal(t, true, EnableHTTPSilaCompatAPI("eth,foo"))
	assert.Equal(t, true, EnableHTTPSilaCompatAPI("foo,eth"))
	assert.Equal(t, true, EnableHTTPSilaCompatAPI("eth,eth"))
	assert.Equal(t, true, EnableHTTPSilaCompatAPI("EtH"))
	assert.Equal(t, false, EnableHTTPSilaCompatAPI("foo"))
	assert.Equal(t, false, EnableHTTPSilaCompatAPI(""))
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
