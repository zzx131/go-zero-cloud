package breakerlearn

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/breaker"
	"github.com/zeromicro/go-zero/core/stat"
	"strconv"
	"testing"
)

func init() {
	stat.SetReporter(nil)
}

func TestCircuitBreaker_Allow(t *testing.T) {
	b := breaker.NewBreaker()
	assert.True(t, len(b.Name()) > 0)
	_, err := b.Allow()
	assert.Nil(t, err)
}

func TestLogReason(t *testing.T) {
	b := breaker.NewBreaker()
	assert.True(t, len(b.Name()) > 0)

	for i := 0; i < 1000; i++ {
		_ = b.Do(func() error {
			return errors.New(strconv.Itoa(i))
		})
	}
}

func BenchmarkGoogleBreaker(b *testing.B) {
	br := breaker.NewBreaker()
	for i := 0; i < b.N; i++ {
		_ = br.Do(func() error {
			return nil
		})

	}
	fmt.Println("ene")
}
