// FILEPATH: /Users/pp/dev/go/go-expert/desafio-client-server-api/server/internal/awesome_api_test.go

package internal

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAwesomeApi_Do(t *testing.T) {

	api := NewAwesomeApi()

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	codeIn := "USD-BRL"

	quote, err := api.Do(ctx, codeIn)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.Equal(t, "USD", quote.Code)
	assert.Equal(t, "BRL", quote.CodeIn)
	assert.Equal(t, "Dólar Americano/Real Brasileiro", quote.Name)

	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel2()

	_, err2 := api.Do(ctx2, codeIn)
	assert.ErrorIs(t, err2, context.DeadlineExceeded)

}
