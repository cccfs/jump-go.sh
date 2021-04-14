package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	env "github.com/alibabacloud-go/darabonba-env/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/viper"
	"os"
)

func CreateClient (accessKeyId *string, accessKeySecret *string) (_result *ecs20140526.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId: accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	config.Endpoint = tea.String("ecs-cn-hangzhou.aliyuncs.com")
	_result = &ecs20140526.Client{}
	_result, _err = ecs20140526.NewClient(config)
	return _result, _err
}

type Info struct {
	Name string `json:"name"`
	SshUser interface{} `json:"ssh_user"`
	PublicIpAddress interface{} `json:"public_ip_address"`
	PrivateIpAddress interface{} `json:"private_ip_address"`
}

func RequestInstances () error {
	client, _err := CreateClient(env.GetEnv(tea.String("ALICLOUD_ACCESS_KEY")), env.GetEnv(tea.String("ALICLOUD_SECRET_KEY")))
	if _err != nil {
		return _err
	}

	describeInstancesRequest := &ecs20140526.DescribeInstancesRequest{
		RegionId: env.GetEnv(tea.String("ALICLOUD_REGION")),
	}

	resp, _err := client.DescribeInstances(describeInstancesRequest)
	if _err != nil {
		return _err
	}
	result, _ := json.MarshalIndent(resp, "", "  ")
	viper.SetConfigType("json")
	viper.ReadConfig(bytes.NewBuffer(result))
	var instanceInfo []interface{}
	viper.UnmarshalKey("body.Instances.Instance", &instanceInfo)

	info := &Info{}
	for _, instance := range instanceInfo {
		result, _ := json.Marshal(instance)
		viper.SetConfigType("json")
		viper.ReadConfig(bytes.NewBuffer(result))
		instanceId := viper.Get("InstanceId").(string)
		info.Name = viper.Get("InstanceName").(string) + instanceId
		var tagInfo []map[string]string
		viper.UnmarshalKey("Tags.tag", &tagInfo)
		for _, v := range tagInfo {
			if v["TagKey"] == "SshUser" {
				info.SshUser = v["TagValue"]
			} else {
				info.SshUser = "root"
			}
		}
		info.PrivateIpAddress = viper.Get("VpcAttributes.PrivateIpAddress.IpAddress").(interface{})
		info.PublicIpAddress = viper.Get("PublicIpAddress.IpAddress").(interface{})
		file, _ := json.MarshalIndent(info, "", " ")
		f, err := os.OpenFile(os.Args[1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()
		fmt.Fprintf(f, "%s", file)
	}
	return _err
}


func main() {
	err := RequestInstances()
	if err != nil {
		panic(err)
	}
}