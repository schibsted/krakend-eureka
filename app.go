//
// Copyright 2011 - 2018 Schibsted Products & Technology AS.
// Licensed under the terms of the Apache 2.0 license. See LICENSE in the project root.
//

package eureka

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/hudl/fargo"

	"os"
)

func NewLocalAppInstance(port int, appName string) (*fargo.Instance, error) {
	hostname, _ := os.Hostname()
	return &fargo.Instance{
		HostName:         hostname,
		Port:             port,
		SecurePort:       443,
		App:              appName,
		IPAddr:           "127.0.0.1",
		VipAddress:       appName,
		SecureVipAddress: "127.0.0.1",
		DataCenterInfo:   fargo.DataCenterInfo{Name: fargo.MyOwn},
		Status:           fargo.STARTING,
		Overriddenstatus: fargo.UNKNOWN,
		HealthCheckUrl:   fmt.Sprintf("http://%s:%d/healthcheck", hostname, port),
		StatusPageUrl:    fmt.Sprintf("http://%s:%d/healthcheck", hostname, port),
		HomePageUrl:      fmt.Sprintf("http://%s:%d/", hostname, port),
	}, nil

}

func NewAwsAppInstance(port int, appName string, instanceMeta *ec2metadata.EC2Metadata) (*fargo.Instance, error) {
	instanceMetadata := getAwsInstanceMetadata(instanceMeta)

	hostname := instanceMetadata.HostName
	secVipAddr, _ := instanceMeta.GetMetadata("public-ipv4")
	ipAddr, _ := instanceMeta.GetMetadata("local-ipv4")
	dataCenter := fargo.DataCenterInfo{Name: fargo.Amazon}
	dataCenter.Metadata = *instanceMetadata

	uniqueIdFunc := func(fargo.Instance) string {
		return instanceMetadata.InstanceID
	}

	return &fargo.Instance{
		HostName:         hostname,
		Port:             port,
		SecurePort:       443,
		App:              appName,
		IPAddr:           ipAddr,
		VipAddress:       appName,
		SecureVipAddress: secVipAddr,
		DataCenterInfo:   dataCenter,
		Status:           fargo.STARTING,
		Overriddenstatus: fargo.UNKNOWN,
		HealthCheckUrl:   fmt.Sprintf("http://%s:%d/healthcheck", hostname, port),
		StatusPageUrl:    fmt.Sprintf("http://%s:%d/healthcheck", hostname, port),
		HomePageUrl:      fmt.Sprintf("http://%s:%d/", hostname, port),
		UniqueID:         uniqueIdFunc,
	}, nil
}

func getAwsInstanceMetadata(instanceMeta *ec2metadata.EC2Metadata) *fargo.AmazonMetadataType {
	hostName, _ := instanceMeta.GetMetadata("public-hostname")
	instanceID, _ := instanceMeta.GetMetadata("instance-id")
	publicIP, _ := instanceMeta.GetMetadata("public-ipv4")
	az, _ := instanceMeta.GetMetadata("placement/availability-zone")
	localHostname, _ := instanceMeta.GetMetadata("local-hostname")
	publicHostname, _ := instanceMeta.GetMetadata("public-hostname")
	amiManifestPath, _ := instanceMeta.GetMetadata("ami-manifest-path")
	localIpv4, _ := instanceMeta.GetMetadata("local-ipv4")
	amiID, _ := instanceMeta.GetMetadata("ami-id")
	instanceType, _ := instanceMeta.GetMetadata("instance-type")

	if hostName == "" {
		hostName = localHostname
		publicHostname = localHostname
	}

	return &fargo.AmazonMetadataType{
		LocalHostname:    localHostname,
		AvailabilityZone: az,
		InstanceID:       instanceID,
		PublicIpv4:       publicIP,
		PublicHostname:   publicHostname,
		AmiManifestPath:  amiManifestPath,
		LocalIpv4:        localIpv4,
		HostName:         hostName,
		AmiID:            amiID,
		InstanceType:     instanceType,
	}
}
