// Code generated by go-swagger; DO NOT EDIT.

package appregistry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/operator-framework/go-appr/appregistry/appr"
	"github.com/operator-framework/go-appr/appregistry/blobs"
	"github.com/operator-framework/go-appr/appregistry/channel"
	"github.com/operator-framework/go-appr/appregistry/info"
	"github.com/operator-framework/go-appr/appregistry/package_appr"
)

// Default appregistry HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:5000"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new appregistry HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Appregistry {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new appregistry HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Appregistry {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new appregistry client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Appregistry {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Appregistry)
	cli.Transport = transport

	cli.Appr = appr.New(transport, formats)

	cli.Blobs = blobs.New(transport, formats)

	cli.Channel = channel.New(transport, formats)

	cli.Info = info.New(transport, formats)

	cli.PackageAppr = package_appr.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Appregistry is a client for appregistry
type Appregistry struct {
	Appr *appr.Client

	Blobs *blobs.Client

	Channel *channel.Client

	Info *info.Client

	PackageAppr *package_appr.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Appregistry) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Appr.SetTransport(transport)

	c.Blobs.SetTransport(transport)

	c.Channel.SetTransport(transport)

	c.Info.SetTransport(transport)

	c.PackageAppr.SetTransport(transport)

}
