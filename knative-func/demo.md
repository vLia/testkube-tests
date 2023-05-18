# Testkube Tests running on Knative

This is a POC for running Testkube tests on Knative.

## Usage instructions

### Deploy this function

```bash
kn func run
```

### Deploy this function when already built

```bash
kn func run --build=false
```

### Apply the executor

```bash
kubectl apply -f executor.yaml
```

### Create a test

```bash
testkube create test --file curl-test.json --type knative --name curl-test-knative
```

### Run the test

```bash
testkube run test curl-test-knative
```

### Check test status

Check the result in the function logs first.

```bash
testkube get test curl-test-knative
```
