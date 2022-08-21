module github.com/finiteloopme/xds-from-scratch

go 1.19

require (
	github.com/envoyproxy/go-control-plane v0.10.2-0.20220325020618-49ff273808a1
	github.com/finiteloopme/goutils v0.0.0-20220818081022-dcee2b6467e0
	github.com/golang/protobuf v1.5.2
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.1
)

require google.golang.org/genproto v0.0.0-20220222213610-43724f9ea8cf

require (
	github.com/census-instrumentation/opencensus-proto v0.2.1 // indirect
	github.com/cncf/xds/go v0.0.0-20211011173535-cb28da3451f1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.1.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220209214540-3681064d5158 // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace github.com/finiteloopme/xds-from-scratch => ../
