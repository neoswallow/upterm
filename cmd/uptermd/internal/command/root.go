package command

import (
	"github.com/jingweno/upterm/server"
	"github.com/jingweno/upterm/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	flagHost        string
    flagWSHost       string
	flagHostKeys    []string
	flagNetwork     string
	flagNetworkOpts []string
	flagMetricAddr  string
)

func Root(logger log.FieldLogger) *cobra.Command {
	rootCmd := &rootCmd{logger}
	cmd := &cobra.Command{
		Use:   "uptermd",
		Short: "Upterm daemon",
		RunE:  rootCmd.Run,
	}

	cmd.PersistentFlags().StringVarP(&flagHost, "host", "", utils.DefaultLocalhost("2222"), "host (required)")
	cmd.PersistentFlags().StringVarP(&flagWSHost, "ws-host", "", "", "websocket host")
	cmd.PersistentFlags().StringSliceVarP(&flagHostKeys, "host-key", "", nil, "host private key")

	cmd.PersistentFlags().StringVarP(&flagNetwork, "network", "", "mem", "network provider")
	cmd.PersistentFlags().StringSliceVarP(&flagNetworkOpts, "network-opt", "", nil, "network provider option")

	cmd.PersistentFlags().StringVarP(&flagMetricAddr, "metric-addr", "", utils.DefaultLocalhost("9090"), "metric server address (required)")

	return cmd
}

type rootCmd struct {
	logger log.FieldLogger
}

func (cmd *rootCmd) Run(c *cobra.Command, args []string) error {
	opt := server.Opt{
		Addr:         flagHost,
		WSAddr:       flagWSHost,
		KeyFiles:     flagHostKeys,
		Network:      flagNetwork,
		NetworkOpt:   flagNetworkOpts,
		MetricAddr:   flagMetricAddr,
	}

	logger := cmd.logger.WithFields(log.Fields{
		"host":         flagHost,
		"metric-addr":  flagMetricAddr,
		"network":      flagNetwork,
		"network-opts": flagNetworkOpts,
	})
	logger.Info("starting server")
	defer logger.Info("shutting down sterver")

	return server.Start(opt)
}
