package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	confType string
	confFile string
	confETCD string
)

var vers bool

var RootCmd = &cobra.Command{
	Use: "cmdbdemo-api",
	Short: "cmdbdemo-api 管理系统",
	Long: "cmdbdemo-api......",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("no flags find")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env/etcd]")
	RootCmd.PersistentFlags().StringVarP(&confFile, "config-file", "f", "etc/demo.toml", "the service config from file")
	RootCmd.PersistentFlags().StringVarP(&confETCD, "config-etcd", "e", "127.0.0.1:2379", "the service config from etcd")
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the demo version")
}
