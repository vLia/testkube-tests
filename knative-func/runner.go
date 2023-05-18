package function

import (
	"context"
	"log"
	"os"

	curlRunnerPkg "github.com/kubeshop/testkube/contrib/executor/curl/pkg/runner"
	initRunnerPkg "github.com/kubeshop/testkube/contrib/executor/init/pkg/runner"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/envs"
	"github.com/kubeshop/testkube/pkg/executor/output"
	"github.com/kubeshop/testkube/pkg/ui"
)

// SetUpCurlRunner prepares a curl runner
func SetUpCurlRunner() *curlRunnerPkg.CurlRunner {
	ctx := context.Background()
	params := LoadVariables()
	r, err := curlRunnerPkg.NewCurlRunner(ctx, params)
	if err != nil {
		log.Fatalf("%s Could not run cURL tests: %s", ui.IconCross, err.Error())
	}

	return r
}

// RunCurl run the execution
func RunCurl(ctx context.Context, r *curlRunnerPkg.CurlRunner, execution testkube.Execution) testkube.ExecutionResult {
	// if r.GetType().IsMain() && execution.PreRunScript != "" {
	// 	output.PrintEvent("running script", execution.Id)

	// 	if err = runScript(e.PreRunScript); err != nil {
	// 		output.PrintError(os.Stderr, err)
	// 		os.Exit(1)
	// 	}
	// }

	output.PrintEvent("running test", execution.Id)
	result, err := r.Run(ctx, execution)
	if err != nil {
		output.PrintError(os.Stderr, err)
		os.Exit(1)
	}

	output.PrintResult(result)
	return result
}

// SetUpInitRunner prepares a curl runner
func SetUpInitRunner() *initRunnerPkg.InitRunner {
	params := LoadVariables()
	r := initRunnerPkg.NewRunner(params)
	return r
}

func LoadVariables() envs.Params {
	return envs.Params{
		DataDir: os.TempDir(),
	}
}
