# go-dcmi
### Overview
包装了一些常用的dcmi函数，可以方便地获取一些NPU信息

### Quick Start
```bash
go get "github.com/huyiqiu/go-dcmi/dcmi"
```

```go
package main

import (
	"fmt"
	"github.com/huyiqiu/go-dcmi/dcmi"
)

func main() {
	err := dcmi.DcmiInit()
	if err != nil {
		fmt.Println(err)
		return
	}
	cnt, ids, _ := dcmi.GetCardList()
	fmt.Printf("%d NPUs found\n", cnt)
	for _, id := range ids {
		deviceNum, err := dcmi.GetDeviceNumInCard(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := 0; i < deviceNum; i++ {
			device := dcmi.GetDeviceByCardIdAndDeviceId(id, i)
			t, err := device.GetDeviceType()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("NpuId: %d, chipId: %d device type: %s\n", id, i, t)
		}
	}
}
```

执行后你可能会得到如下信息
```bash
 go run main.go 
4 NPUs found
NpuId: 2, chipId: 0 device type: NPU
NpuId: 2, chipId: 1 device type: NPU
NpuId: 3, chipId: 0 device type: NPU
NpuId: 3, chipId: 1 device type: NPU
NpuId: 5, chipId: 0 device type: NPU
NpuId: 5, chipId: 1 device type: NPU
NpuId: 6, chipId: 0 device type: NPU
NpuId: 6, chipId: 1 device type: NPU
```