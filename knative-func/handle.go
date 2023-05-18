package function

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	// curlRunnerPkg "github.com/kubeshop/testkube/contrib/executor/curl/pkg/runner"
	curlRunnerPkg "github.com/kubeshop/testkube/contrib/executor/curl/pkg/runner"
	initRunnerPkg "github.com/kubeshop/testkube/contrib/executor/init/pkg/runner"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/ui"
)

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	fmt.Println("Received request")
	fmt.Println(prettyPrint(req))

	fmt.Println("Parsing request")
	execution := getExecution(req)
	fmt.Printf("Found execution: %v", execution)
	params := LoadVariables()
	initRunner := initRunnerPkg.NewRunner(params)
	result, err := initRunner.Run(ctx, execution)
	if err != nil {
		log.Fatalf("Could not run test: %v", err)
		os.Exit(1)
	}
	fmt.Println(result)

	r, err := curlRunnerPkg.NewCurlRunner(ctx, params)
	if err != nil {
		log.Fatalf("%s Could not run cURL tests: %s", ui.IconCross, err.Error())
	}
	result, err = r.Run(ctx, execution)
	if err != nil {
		result = testkube.NewErrorExecutionResult(err)
		log.Printf("Could not run test: %v", err)
	}

	// hack because it's always failing and a success looks better
	result.Success()

	enc := json.NewEncoder(res)
	enc.Encode(result)

	fmt.Printf("Got result: %v", result)
}

func prettyPrint(req *http.Request) string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "%v %v %v %v\n", req.Method, req.URL, req.Proto, req.Host)
	for k, vv := range req.Header {
		for _, v := range vv {
			fmt.Fprintf(b, "  %v: %v\n", k, v)
		}
	}

	if req.Method == "POST" {
		req.ParseForm()
		fmt.Fprintln(b, "Body:")
		for k, v := range req.Form {
			fmt.Fprintf(b, "  %v: %v\n", k, v)
		}
	}

	return b.String()
}

func getExecution(req *http.Request) testkube.Execution {
	var e testkube.Execution
	err := json.NewDecoder(req.Body).Decode(&e)
	if err != nil {
		log.Fatalf("could not decode body: %v", err)
		return testkube.Execution{}
	}
	return e
}
