package testutil

import (
	"testing"

	"github.com/relnod/fsa"
	"github.com/stretchr/testify/assert"
)

func TestCreateFiles(t *testing.T) {
	fs := fsa.NewTempFs(fsa.NewOsFs())
	defer fs.Cleanup()

	assert.NoError(t, CreateFiles(fs, `
	/a/
	/a/b
	/b/c/d:ln
	/b/c/e#bla
	`))

	assert.True(t, FileExists(fs, "/a/"))
	assert.True(t, FileExists(fs, "/a/b"))
	assert.True(t, FileExists(fs, "/b/c/d"))
	assert.True(t, IsSymlink(fs, "/b/c/d"), "Should be a symlink")
	assert.True(t, FileExists(fs, "/b/c/e"))
}

func TestCheckFikes(t *testing.T) {
	fs := fsa.NewTempFs(fsa.NewOsFs())
	defer fs.Cleanup()

	assert.NoError(t, CreateFiles(fs, `
	/a/
	/a/b
	/b/c/d:ln
	/b/c/e#bla
	`))
	assert.NoError(t, CheckFiles(fs, `
	/a/
	/a/b
	/b/c/d:ln
	/b/c/e
	`))
}
