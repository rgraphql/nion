module github.com/rgraphql/nion

go 1.14

replace github.com/rgraphql/rgraphql => github.com/rgraphql/rgraphql v1.0.1-0.20181030233530-86a2aefc6fe6 // rewrite

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/golang/protobuf v1.4.2
	github.com/graphql-go/graphql v0.7.9
	github.com/hashicorp/golang-lru v0.5.4
	github.com/pkg/errors v0.9.1
	github.com/rgraphql/magellan v0.5.0 // indirect
	github.com/rgraphql/rgraphql v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.6.0
	github.com/urfave/cli v1.22.4
	golang.org/x/tools v0.0.0-20200619210111-0f592d2728bb
)
