package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"net/http"
	"strings"

	// "github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"helm.sh/helm/v3/pkg/cli"
	"k8s.io/client-go/tools/clientcmd"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/registry"
	"helm.sh/helm/v3/pkg/repo"
	// "helm.sh/helm/v3/internal/tlsutil"
	"oras.land/oras-go/pkg/auth"
	registryauth "oras.land/oras-go/pkg/registry/remote/auth"
	"github.com/containerd/containerd/remotes"
)

var settings = cli.New()

func newRegistryClient(certFile, keyFile, caFile string, insecureSkipTLSverify bool) (*registry.Client, error) {
	if certFile != "" && keyFile != "" || caFile != "" || insecureSkipTLSverify {
		registryClient, err := newRegistryClientWithTLS(certFile, keyFile, caFile, insecureSkipTLSverify)
		if err != nil {
			return nil, err
		}
		return registryClient, nil
	}
	registryClient, err := newDefaultRegistryClient()
	if err != nil {
		return nil, err
	}
	return registryClient, nil
}

func newDefaultRegistryClient() (*registry.Client, error) {
	// Create a new registry client
	registryClient, err := registry.NewClient(
		registry.ClientOptDebug(settings.Debug),
		registry.ClientOptEnableCache(true),
		registry.ClientOptWriter(os.Stderr),
		registry.ClientOptCredentialsFile(settings.RegistryConfig),
	)
	if err != nil {
		return nil, err
	}
	return registryClient, nil
}

func newRegistryClientWithTLS(certFile, keyFile, caFile string, insecureSkipTLSverify bool) (*registry.Client, error) {
	// Create a new registry client
	
	registryClient, err := NewRegistryClientWithTLS(os.Stderr, certFile, keyFile, caFile, insecureSkipTLSverify,
		settings.RegistryConfig, settings.Debug,
	)
	if err != nil {
		return nil, err
	}
	return registryClient, nil
}