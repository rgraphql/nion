#!/bin/bash

go run \
	 github.com/rgraphql/nion/cmd/nion \
   analyze --schema ./schema.graphql \
   --go-pkg github.com/rgraphql/nion/example/simple \
   --go-query-type RootResolver \
   --go-output ./resolve/resolve_generated.go
