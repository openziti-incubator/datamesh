module github.com/openziti-incubator/datamesh

go 1.16

replace github.com/openziti-incubator/cf => ../cf

replace github.com/michaelquigley/pfxlog => ../../q/products/pfxlog

replace github.com/openziti/foundation => ../foundation

require (
	github.com/michaelquigley/pfxlog v0.5.0
	github.com/openziti-incubator/cf v0.0.1
	github.com/openziti/foundation v0.15.56
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
)