package cli

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"

	"github.com/jack139/go-infer/http"
	"github.com/jack139/go-infer/server"
)

var (
	// http 服务
	HttpCmd = &cobra.Command{
		Use:   "http",
		Short: "start http service",
		RunE: func(cmd *cobra.Command, args []string) error {
			// 启动 http 服务
			http.RunServer()

			return nil
		},
	}

	// Dispatcher server
	ServerCmd = &cobra.Command{
		Use:   "server <queue No.>",
		Short: "start dispatcher service",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("queue number needed")
			}

			_, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("queue number should be a integer")
			}

			// 启动 分发服务
			server.RunServer(args[0])

			return nil
		},
	}
)
