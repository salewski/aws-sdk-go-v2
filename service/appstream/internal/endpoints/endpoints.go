// Code generated by smithy-go-codegen DO NOT EDIT.

package endpoints

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	endpoints "github.com/aws/aws-sdk-go-v2/aws/endpoints/v2"
	"regexp"
)

// Options is the endpoint resolver configuration options
type Options struct {
	DisableHTTPS bool
}

// Resolver AppStream endpoint resolver
type Resolver struct {
	partitions endpoints.Partitions
}

// ResolveEndpoint resolves the service endpoint for the given region and options
func (r *Resolver) ResolveEndpoint(region string, options Options) (aws.Endpoint, error) {
	opt := endpoints.Options{
		DisableHTTPS: options.DisableHTTPS,
	}
	return r.partitions.ResolveEndpoint(region, opt)
}

// New returns a new Resolver
func New() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var defaultPartitions = endpoints.Partitions{
	{
		ID: "aws",
		Defaults: endpoints.Endpoint{
			Hostname:          "appstream2.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
			CredentialScope: endpoints.CredentialScope{
				Service: "appstream",
			},
		},
		RegionRegex:    regexp.MustCompile("^(us|eu|ap|sa|ca|me)\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"ap-northeast-1": endpoints.Endpoint{},
			"ap-northeast-2": endpoints.Endpoint{},
			"ap-southeast-1": endpoints.Endpoint{},
			"ap-southeast-2": endpoints.Endpoint{},
			"eu-central-1":   endpoints.Endpoint{},
			"eu-west-1":      endpoints.Endpoint{},
			"fips": endpoints.Endpoint{
				Hostname: "appstream2-fips.us-west-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-2",
				},
			},
			"us-east-1": endpoints.Endpoint{},
			"us-west-2": endpoints.Endpoint{},
		},
	},
	{
		ID: "aws-cn",
		Defaults: endpoints.Endpoint{
			Hostname:          "appstream2.{region}.amazonaws.com.cn",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^cn\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
	},
	{
		ID: "aws-iso",
		Defaults: endpoints.Endpoint{
			Hostname:          "appstream2.{region}.c2s.ic.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^us\\-iso\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
	},
	{
		ID: "aws-iso-b",
		Defaults: endpoints.Endpoint{
			Hostname:          "appstream2.{region}.sc2s.sgov.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^us\\-isob\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
	},
	{
		ID: "aws-us-gov",
		Defaults: endpoints.Endpoint{
			Hostname:          "appstream2.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
			CredentialScope: endpoints.CredentialScope{
				Service: "appstream",
			},
		},
		RegionRegex:    regexp.MustCompile("^us\\-gov\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"fips": endpoints.Endpoint{
				Hostname: "appstream2-fips.us-gov-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-west-1",
				},
			},
			"us-gov-west-1": endpoints.Endpoint{},
		},
	},
}