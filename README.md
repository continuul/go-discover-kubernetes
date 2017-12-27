# Go Discover Nodes for Cloud Providers [![Build Status](https://travis-ci.org/hashicorp/go-discover.svg?branch=master)](https://travis-ci.org/hashicorp/go-discover) [![GoDoc](https://godoc.org/github.com/hashicorp/go-discover?status.svg)](https://godoc.org/github.com/hashicorp/go-discover)

`go-discover` is a Go (golang) library and command line tool to discover
ip addresses of nodes in cloud environments based on meta information
like tags provided by the environment.

See [go-discover](https://github.com/hashicorp/go-discover) for full details.

This library is an extension to `go-discover` to add Kubernetes support,
specifically to discover addresses of load-balanced applications from its
related [Service](https://kubernetes.io/docs/concepts/services-networking/service/) object.

### Supported Providers

Additional cloud providers herein include:

 * Kubernetes [Config options](https://github.com/continuul/go-discover/blob/master/provider/kubernetes/kubernetes_discover.go#L15-L22)

### Config Example

```
# Kubernetes
provider=kubernetes namespace=default service=demo

```

## Command Line Tool Usage

Install the command line tool with:

```
go get -u github.com/continuul/go-discover/cmd/discover
```

The build it with:

```
make build
```

Then run it with:

```
$ kubectl create -f nginx.yaml
$ kubectl create -f discover.yaml
```

## Library Usage

Install the library with:

```
go get -u github.com/continuul/go-discover
```

You can then either support discovery for all available providers
or only for some of them.

```go
// support discovery for all supported providers
d := discover.Discover{}

// support discovery for AWS and GCE only
d := discover.Discover{
	Providers : map[string]discover.Provider{
		"aws": discover.Providers["aws"],
		"gce": discover.Providers["gce"],
		"kubernetes": &kubernetes.Provider{},
	}
}

// use ioutil.Discard for no log output
l := log.New(os.Stderr, "", log.LstdFlags)

cfg := "provider=kubernetes namespace=default ..."
addrs, err := d.Addrs(cfg, l)
```
