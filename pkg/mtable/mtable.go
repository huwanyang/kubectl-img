package mtable

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
)

var title = []string{
	"NAMESPACE", "TYPE", "RESOURCE_NAME", "CONTAINER_NAME", "IMAGE",
}

/*
 * 通过 gotable 组件将结果美化输出
 */
func GenTable(mapList []map[string]string) *table.Table {
	table, err := gotable.Create(title...)
	if err != nil {
		fmt.Printf("create table fail: %s", err.Error())
		return nil
	}
	table.AddRows(mapList)
	return table
}
