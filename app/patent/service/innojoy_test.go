package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newMockInnojoyClient() *InnojoyClient {
	return &InnojoyClient{
		email:    "zx@wei-ping.com",
		password: "405b0ceb6fe44c79a61e48374bb35fbc",
		hc:       newHttpClient(),
	}
}

func TestInnojoyAutoLogin(t *testing.T) {
	ic := newMockInnojoyClient()
	err := ic.autoLogin()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(ic.token))
}

func TestInnojoySimpleSearch(t *testing.T) {
	ic := newMockInnojoyClient()
	res, err := ic.SimpleSearch("请求处理方法", "wgzl,syxx,fmzl")
	assert.NoError(t, err)
	for _, r := range res {
		fmt.Println(r)
	}
}