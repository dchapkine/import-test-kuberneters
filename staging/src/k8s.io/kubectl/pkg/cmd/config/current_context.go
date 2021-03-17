/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"fmt"

	"github.com/spf13/cobra"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

// CurrentContextOptions holds the command-line options for 'config current-context' sub command
type CurrentContextOptions struct {
	ConfigAccess clientcmd.ConfigAccess

	genericclioptions.IOStreams
}

var (
	currentContextLong = templates.LongDesc(i18n.T(`
		Displays the current-context`))

	currentContextExample = templates.Examples(`
		# Display the current-context
		kubectl config current-context`)
)

// NewCmdConfigCurrentContext returns a Command instance for 'config current-context' sub command
func NewCmdConfigCurrentContext(streams genericclioptions.IOStreams, configAccess clientcmd.ConfigAccess) *cobra.Command {
	options := &CurrentContextOptions{
		ConfigAccess: configAccess,
		IOStreams:    streams,
	}

	cmd := &cobra.Command{
		Use:     "current-context",
		Short:   i18n.T("Displays the current-context"),
		Long:    currentContextLong,
		Example: currentContextExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(options.RunCurrentContext())
		},
	}

	return cmd
}

// RunCurrentContext performs the execution of 'config current-context' sub command
func (o *CurrentContextOptions) RunCurrentContext() error {
	config, err := o.ConfigAccess.GetStartingConfig()
	if err != nil {
		return err
	}

	if config.CurrentContext == "" {
		err = fmt.Errorf("current-context is not set")
		return err
	}

	fmt.Fprintf(o.Out, "%s\n", config.CurrentContext)
	return nil
}
