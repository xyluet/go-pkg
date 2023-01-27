package contextext_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	contextext "github.com/xyluet/pkg/context"
)

func TestDetach(t *testing.T) {
	key := struct{ name string }{name: "key"}
	ctx := context.WithValue(context.Background(), key, 12)
	ctx, cancel := context.WithTimeout(ctx, time.Nanosecond)
	cancel()

	ctx = contextext.Detach(ctx)
	select {
	case <-ctx.Done():
		require.Fail(t, "context should be detached from parents cancellation")
	default:
	}

	value := ctx.Value(key)
	require.Equal(t, 12, value.(int))
}
