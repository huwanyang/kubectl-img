package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var KubernetesConfigFlags *genericclioptions.ConfigFlags
var version = "v1.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-img",
	Short: "显示 k8s 资源的 image 信息命令插件",
	Long: `基于 Cobra 脚手架创建的 kubectl-img 命令插件，可以显示指定的 k8s 资源类型（deployments|daemonsets|statefulsets|jobs|cronjobs）
的 image 信息，并支持多种输出方式（json|yaml|xml|table）。`,
	Version: version,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 添加自定义的 Flags，也可以放在子命令的 init 方法里
	imageCmd.Flags().BoolP("deployments", "d", false, "Show resource deployment images")
	imageCmd.Flags().BoolP("daemonsets", "a", false, "Show resource daemonsets images")
	imageCmd.Flags().BoolP("statefulsets", "f", false, "Show resource statefulsets images")
	imageCmd.Flags().BoolP("jobs", "j", false, "Show resource jobs images")
	imageCmd.Flags().BoolP("cronjobs", "c", false, "Show resource cronjobs images")
	imageCmd.Flags().StringP("output", "o", "table", "Output format. One of: json|yaml|xml|table ")
	// 添加 kubernetes 默认的 ConfigFlags，以支持更多的命令扩展
	KubernetesConfigFlags = genericclioptions.NewConfigFlags(true)
	KubernetesConfigFlags.AddFlags(rootCmd.PersistentFlags())
}
