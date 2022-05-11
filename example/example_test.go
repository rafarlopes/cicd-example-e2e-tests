package example

import (
	"net/http"
	"os"
	"testing"
)

func TestHappyPath(t *testing.T) {

	t.Run("should return http 200", func(t *testing.T) {
		resp, err := http.Get(getURL(t))
		if err != nil {
			t.Fatalf("Should get 200 but got error %+v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Should get 200 but got %d", resp.StatusCode)
		}
	})
}

func getURL(t *testing.T) string {
	t.Helper()
	if url := os.Getenv("APPLICATION_URL"); url != "" {
		return url
	}

	t.Fatal("Missing environment variable APPLICATION_URL")
	return ""
}
