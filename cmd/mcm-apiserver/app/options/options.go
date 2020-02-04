// licensed Materials - Property of IBM
// 5737-E67
// (C) Copyright IBM Corporation 2016, 2019 All Rights Reserved
// US Government Users Restricted Rights - Use, duplication or disclosure restricted by GSA ADP Schedule Contract with IBM Corp.

// Package options contains flags and options for initializing hcm-apiserver.
package options

import (
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/apiserver/pkg/storage/storagebackend"

	"github.com/spf13/pflag"

	klusterlet "github.com/open-cluster-management/multicloud-operators-foundation/pkg/klusterlet/client"
	mcmstorage "github.com/open-cluster-management/multicloud-operators-foundation/pkg/storage"
)

const (
	// DefaultEtcdPathPrefix is the default prefix that is prepended to all
	// resource paths in etcd.  It is intended to allow an operator to
	// differentiate the storage of different API servers from one another in
	// a single etcd.
	DefaultEtcdPathPrefix = "/registry"
)

// ServerRunOptions for the hcm-apiserver.
type ServerRunOptions struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions
	Etcd                    *genericoptions.EtcdOptions
	SecureServing           *genericoptions.SecureServingOptionsWithLoopback
	Audit                   *genericoptions.AuditOptions
	Admission               *genericoptions.AdmissionOptions
	Authentication          *genericoptions.DelegatingAuthenticationOptions
	Authorization           *genericoptions.DelegatingAuthorizationOptions
	APIEnablement           *genericoptions.APIEnablementOptions
	MCMStorage              *mcmstorage.Options
	KlusterletClientOptions *klusterlet.ClientOptions

	StandAlone         bool
	AuthorizationQPS   float32
	AuthorizationBurst int
}

// NewServerRunOptions creates a new ServerRunOptions object with default values.
func NewServerRunOptions() *ServerRunOptions {
	s := ServerRunOptions{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		Admission:               genericoptions.NewAdmissionOptions(),
		Etcd:                    genericoptions.NewEtcdOptions(storagebackend.NewDefaultConfig(DefaultEtcdPathPrefix, nil)),
		SecureServing:           genericoptions.NewSecureServingOptions().WithLoopback(),
		Audit:                   genericoptions.NewAuditOptions(),
		Authentication:          genericoptions.NewDelegatingAuthenticationOptions(),
		Authorization:           genericoptions.NewDelegatingAuthorizationOptions(),
		APIEnablement:           genericoptions.NewAPIEnablementOptions(),
		MCMStorage:              mcmstorage.NewStorageOptions(),
		KlusterletClientOptions: klusterlet.NewClientOptions(),
	}

	registerAllAdmissionPlugins(s.Admission.Plugins, &s.KlusterletClientOptions.CAFile)

	return &s
}

// AddFlags adds flags for ServerRunOptions fields to be specified via FlagSet.
func (options *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&options.StandAlone, "standalone", false, "Enable when apiserver is run in standalone mode")

	fs.Float32Var(&options.AuthorizationQPS, "authorization_qps", 500, "Maximum QPS for authorization requests")
	fs.IntVar(&options.AuthorizationBurst, "authorization_burst", 1000, "Maximum burst for authorization requests throttle")
	// Add the generic flags.
	options.GenericServerRunOptions.AddUniversalFlags(fs)
	options.Etcd.AddFlags(fs)
	options.SecureServing.AddFlags(fs)
	options.Audit.AddFlags(fs)
	options.Authentication.AddFlags(fs)
	options.Authorization.AddFlags(fs)
	options.APIEnablement.AddFlags(fs)
	options.Admission.AddFlags(fs)
	options.MCMStorage.AddFlags(fs)
	options.KlusterletClientOptions.AddFlags(fs)
}
