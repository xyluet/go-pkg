package run_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/xyluet/pkg/run"
)

func TestNoActors(t *testing.T) {
	var g run.Group
	res := make(chan error)
	go func() { res <- g.Run() }()
	select {
	case err := <-res:
		require.Nil(t, err)
	case <-time.After(100 * time.Millisecond):
		require.Fail(t, "timeout")
	}
}

func TestReturnError(t *testing.T) {
	myErr := errors.New("!")
	var g run.Group
	g.Add(func() error { return myErr }, func(err error) {})
	res := g.Run()
	require.Equal(t, myErr, res)
}

func TestReturnFirstError(t *testing.T) {
	myErr := errors.New("!")
	var g run.Group
	g.Add(func() error { return myErr }, func(err error) {})
	cancel := make(chan error)
	g.Add(func() error { <-cancel; return nil }, func(err error) { close(cancel) })
	res := g.Run()
	require.Equal(t, myErr, res)
}
