package function

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/stretchr/testify/assert"
)

// TestHandle ensures that Handle executes without error and returns the
// HTTP 200 status code indicating no errors.
func TestHandle(t *testing.T) {
	var (
		w   = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "http://example.com/test", nil)
		res *http.Response
	)

	Handle(context.Background(), w, req)
	res = w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("unexpected response code: %v", res.StatusCode)
	}
}

func TestCurlRun(t *testing.T) {
	t.Run("Run cURL test", func(t *testing.T) {
		w := httptest.NewRecorder()
		execution := testkube.Execution{
			TestName:      "curl-test",
			TestNamespace: "testkube",
			TestType:      "curl/test",
			Name:          "curl-test-1",
			Number:        1,
			Content: &testkube.TestContent{
				Type_: "string",
				Data: `{
				"command": [
				  "curl",
				  "https://testkube.io/"
				],
				"expected_status": "200"
				  }`,
				// Uri: "https://github.com/kubeshop/testkube.git",
				// Repository: &testkube.Repository{
				// Branch: "main",
				// Path:   "test/curl/executor-tests/curl-smoke-test.json",
				// Uri:    "https://github.com/kubeshop/testkube.git",
				// },
			},
		}
		b, err := json.Marshal(execution)
		assert.NoError(t, err)
		bodyReader := bytes.NewReader(b)
		req := httptest.NewRequest("POST", "http://example.com/test", bodyReader)

		Handle(context.Background(), w, req)
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 200 {
			t.Fatalf("unexpected response code: %v", res.StatusCode)
		}

		decoder := json.NewDecoder(res.Body)
		var responseResult testkube.ExecutionResult
		err = decoder.Decode(&responseResult)
		if err != nil {
			t.Fatalf("couldn't decode response: %v", err)
		}

		// assert.Equal(t, execution, nil)
		assert.Equal(t, testkube.FAILED_ExecutionStatus, *responseResult.Status)
	})
}
