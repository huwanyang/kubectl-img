package cmd

import (
	"context"
	"fmt"
	"github.com/huwanyang/kubectl-img/pkg/kube"
	"github.com/huwanyang/kubectl-img/pkg/mtable"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	appv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beat1 "k8s.io/api/batch/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "显示 k8s 资源 image 信息",
	Long: `image 命令可以显示指定的 k8s 资源类型镜像，例如: deployments|daemonsets|statefulsets|jobs|cronjobs. 
同时可以指定输出格式，例如: json|yaml|xml|table`,
	Run: image,
}

func init() {
	rootCmd.AddCommand(imageCmd)
}

func image(cmd *cobra.Command, args []string) {
	fmt.Println("image command func called")
	clientSet := kube.ClientSet(KubernetesConfigFlags)
	ns, _ := rootCmd.Flags().GetString("namespace")
	var rList []interface{}
	if flag, _ := cmd.Flags().GetBool("deployments"); flag {
		deployments, err := clientSet.AppsV1().Deployments(ns).List(context.Background(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("list deployments error: %s", err.Error())
		}
		rList = append(rList, deployments)
	}

	if flag, _ := cmd.Flags().GetBool("daemonsets"); flag {
		daemonsets, err := clientSet.AppsV1().DaemonSets(ns).List(context.Background(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("list daemonsets err: %s", err.Error())
		}
		rList = append(rList, daemonsets)
	}

	if flag, _ := cmd.Flags().GetBool("statefulsets"); flag {
		statefulsets, err := clientSet.AppsV1().StatefulSets(ns).List(context.Background(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("list statefulsets err: %s", err.Error())
		}
		rList = append(rList, statefulsets)
	}

	if flag, _ := cmd.Flags().GetBool("jobs"); flag {
		jobs, err := clientSet.BatchV1().Jobs(ns).List(context.Background(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("list jobs err: %s", err.Error())
		}
		rList = append(rList, jobs)
	}

	if flag, _ := cmd.Flags().GetBool("cronjobs"); flag {
		cronjobs, err := clientSet.BatchV1beta1().CronJobs(ns).List(context.Background(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("list cronjobs err: %s", err.Error())
		}
		rList = append(rList, cronjobs)
	}
	fmt.Printf("length rList: %d\n", len(rList))

	resourceMap := make([]map[string]string, 0)
	for i := 0; i < len(rList); i++ {
		switch t := rList[i].(type) {
		case *appv1.DeploymentList:
			for k := 0; k < len(t.Items); k++ {
				containers := t.Items[k].Spec.Template.Spec.Containers
				for j := 0; j < len(containers); j++ {
					deployMap := make(map[string]string)
					deployMap["NAMESPACE"] = ns
					deployMap["TYPE"] = "deployment"
					deployMap["RESOURCE_NAME"] = t.Items[k].GetName()
					deployMap["CONTAINER_NAME"] = containers[j].Name
					deployMap["IMAGE"] = containers[j].Image
					resourceMap = append(resourceMap, deployMap)
				}
			}
		case *appv1.StatefulSetList:
			for k := 0; k < len(t.Items); k++ {
				containers := t.Items[k].Spec.Template.Spec.Containers
				for j := 0; j < len(containers); j++ {
					statefulSetMap := make(map[string]string)
					statefulSetMap["NAMESPACE"] = ns
					statefulSetMap["TYPE"] = "statefulset"
					statefulSetMap["RESOURCE_NAME"] = t.Items[k].GetName()
					statefulSetMap["CONTAINER_NAME"] = containers[j].Name
					statefulSetMap["IMAGE"] = containers[j].Image
					resourceMap = append(resourceMap, statefulSetMap)
				}
			}
		case *appv1.DaemonSetList:
			for k := 0; k < len(t.Items); k++ {
				containers := t.Items[k].Spec.Template.Spec.Containers
				for j := 0; j < len(containers); j++ {
					daemonSetMap := make(map[string]string)
					daemonSetMap["NAMESPACE"] = ns
					daemonSetMap["TYPE"] = "daemonset"
					daemonSetMap["RESOURCE_NAME"] = t.Items[k].GetName()
					daemonSetMap["CONTAINER_NAME"] = containers[j].Name
					daemonSetMap["IMAGE"] = containers[j].Image
					resourceMap = append(resourceMap, daemonSetMap)
				}
			}
		case *batchv1.JobList:
			for k := 0; k < len(t.Items); k++ {
				containers := t.Items[k].Spec.Template.Spec.Containers
				for j := 0; j < len(containers); j++ {
					jobMap := make(map[string]string)
					jobMap["NAMESPACE"] = ns
					jobMap["TYPE"] = "job"
					jobMap["RESOURCE_NAME"] = t.Items[k].GetName()
					jobMap["CONTAINER_NAME"] = containers[j].Name
					jobMap["IMAGE"] = containers[j].Image
					resourceMap = append(resourceMap, jobMap)
				}
			}
		case *batchv1beat1.CronJobList:
			for k := 0; k < len(t.Items); k++ {
				containers := t.Items[k].Spec.JobTemplate.Spec.Template.Spec.Containers
				for j := 0; j < len(containers); j++ {
					cronjobMap := make(map[string]string)
					cronjobMap["NAMESPACE"] = ns
					cronjobMap["TYPE"] = "cronjob"
					cronjobMap["RESOURCE_NAME"] = t.Items[k].GetName()
					cronjobMap["CONTAINER_NAME"] = containers[j].Name
					cronjobMap["IMAGE"] = containers[j].Image
					resourceMap = append(resourceMap, cronjobMap)
				}
			}
		}
	}
	output, _ := cmd.Flags().GetString("output")
	if len(resourceMap) == 0 {
		fmt.Printf(`
Usage:
  kubectl-img image [flags]

Flags:
  -c, --cronjobs        Show resource cronjobs images
  -a, --daemonsets      Show resource daemonsets images
  -d, --deployments     Show resource deployment images
  -j, --jobs            Show resource jobs images
  -f, --statefulsets    Show resource statefulsets images

  -o, --output string   Output format. One of: json|yaml
  -h, --help            help for image command
`)
	}
	//for i, m := range resourceMap {
	//	for k, v := range m {
	//		fmt.Printf("index: %d, key: %s, value: %s\n", i, k, v)
	//	}
	//}

	table := mtable.GenTable(resourceMap)
	if len(output) != 0 {
		switch output {
		case "json":
			jsonStr, err := table.JSON(2)
			if err != nil {
				fmt.Printf("table to json fail: %s", err.Error())
			}
			fmt.Println(jsonStr)
		case "yaml":
			bytes, err := yaml.Marshal(resourceMap)
			if err != nil {
				fmt.Printf("table to yaml fail: %s", err.Error())
			}
			fmt.Printf(string(bytes))
		case "xml":
			fmt.Println(table.XML(2))
		case "table":
			fmt.Printf(table.String())
		default:
			fmt.Printf("Unsupported output format '%s'. Only support: json|yaml|xml|table\n", output)
		}
	} else {
		fmt.Printf(table.String())
	}
}
