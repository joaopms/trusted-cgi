package platform_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/reddec/trusted-cgi/application/lambda"
	"github.com/reddec/trusted-cgi/application/platform"
	"github.com/reddec/trusted-cgi/types"

	"github.com/stretchr/testify/assert"
)

type mockValidator struct {
	forbidden map[string]error
}

func (mv *mockValidator) Inspect(lambda string, request *types.Request) error {
	return mv.forbidden[lambda]
}

func TestPlatform_AddWithOldAliases(t *testing.T) {
	temp, err := ioutil.TempFile("", "")
	if !assert.NoError(t, err) {
		return
	}
	defer os.RemoveAll(temp.Name())
	defer temp.Close()

	workdir, err := ioutil.TempDir("", "")
	if !assert.NoError(t, err) {
		return
	}
	defer os.RemoveAll(workdir)

	dummy, err := lambda.DummyPublic(workdir, "/usr/bin/cat", "-")
	if !assert.NoError(t, err) {
		return
	}

	_, err = temp.WriteString(`{"links":{"xxx":"123", "yyy": "456", "zzz": "123"}}`)
	if !assert.NoError(t, err) {
		return
	}

	plato, err := platform.New(temp.Name(), &mockValidator{})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, plato.Config().Links, 3)

	err = plato.Add("123", dummy)
	if !assert.NoError(t, err) {
		return
	}

	byUID, err := plato.FindByUID("123")
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, byUID.Lambda, dummy)

	assert.Len(t, byUID.Aliases, 2)
	assert.Contains(t, byUID.Aliases, "xxx")
	assert.Contains(t, byUID.Aliases, "zzz")

	{
		byLink, err := plato.FindByLink("xxx")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, byLink, byUID)
	}
	{
		byLink, err := plato.FindByLink("zzz")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, byLink, byUID)
	}
}
