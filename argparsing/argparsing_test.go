package argparsing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseArgs(t *testing.T) {
	a := ParseArgs()
	assert.Equal(t, *(a.DbIp), "127.0.0.1", "Should equal default value")
	assert.Equal(t, *(a.DbPort), 5432, "Should equal default value")
	assert.Equal(t, *(a.DbUsername), "postgres", "Should equal default value")
	assert.Equal(t, *(a.DbPassword), "Password1", "Should equal default value")
	assert.Equal(t, *(a.DbName), "postgres", "Should equal default value")
	assert.Equal(t, *(a.WebListenIp), "0.0.0.0", "Should equal default value")
	assert.Equal(t, *(a.WebListenPort), 443, "Should equal default value")
	assert.Equal(t, *(a.WebKeyPath), "./key", "Should equal default value")
	assert.Equal(t, *(a.WebCertPath), "./cert", "Should equal default value")
	assert.Equal(t, *(a.RunTls), true, "Should equal default value")
	assert.Equal(t, *(a.WebRootDir), "./web/", "Should equal default value")
}

