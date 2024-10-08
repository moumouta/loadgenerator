package main

import (
	"github.com/Lincyaw/loadgenerator/behaviors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	// 定义命令行选项变量
	var debug bool

	// 创建一个新的cobra命令
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A load generator application",
		Run: func(cmd *cobra.Command, args []string) {
			if debug {
				logrus.SetLevel(logrus.DebugLevel)
				logrus.SetReportCaller(true)
			} else {
				logrus.SetLevel(logrus.InfoLevel)
			}

			logrus.SetFormatter(&logrus.TextFormatter{
				FullTimestamp: true,
			})

			lg := &behaviors.LoadGenerator{}
			composedChain := behaviors.NewChain(behaviors.NewFuncNode(func(ctx *behaviors.Context) (*behaviors.NodeResult, error) {
				return nil, nil
			}, "dummy"))
			composedChain.AddNextChain(behaviors.NormalPreserveChain, 10)
			composedChain.AddNextChain(behaviors.NormalOrderPayChain, 10)
			composedChain.AddNextChain(behaviors.OrderConsignChain, 10)
			composedChain.AddNextChain(behaviors.TicketCollectAndEnterStationChain, 10)

			composedChain.AddNextChain(behaviors.AdvancedSearchChain, 20)
			composedChain.AddNextChain(behaviors.ConsignListChain, 8)
			composedChain.AddNextChain(behaviors.OrderChangeChain, 3)
			composedChain.AddNextChain(behaviors.OrderCancelChain, 2)
			lg.Start(behaviors.WithThread(1), behaviors.WithSleep(0), behaviors.WithChain(composedChain))
		},
	}

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("Error executing command: %v", err)
		os.Exit(1)
	}
}
