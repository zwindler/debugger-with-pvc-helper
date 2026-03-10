package main

import (
	"github.com/zwindler/kubectl-debug-pvc/cmd"
	// Register client-go authentication plugins (OIDC, GCP, Azure).
	// Required by krew best practices so the plugin can authenticate to
	// clusters that use in-tree credential providers.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// version is set at build time via -ldflags "-X main.version=<tag>".
var version = "dev"

func main() {
	cmd.Execute(version)
}
