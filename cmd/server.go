package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "这是 Demo 命令 server 简单描述",
	Long: `这里是 server 命令的完整描述，支持可跨多行.
			
server 命令是基于 cobra 脚手架生成的，可快速开发我们想要的命令行工具，可实现输出内容。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server 命令被调用，这里开始执行业务逻辑！")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
