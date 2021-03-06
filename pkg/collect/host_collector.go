package collect

import (
	"github.com/pkg/errors"
	troubleshootv1beta2 "github.com/replicatedhq/troubleshoot/pkg/apis/troubleshoot/v1beta2"
)

type HostCollector struct {
	Collect *troubleshootv1beta2.HostCollect
}

type HostCollectors []*HostCollector

func (c *HostCollector) RunCollectorSync() (result map[string][]byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("recovered rom panic: %v", r)
		}
	}()

	if c.Collect.CPU != nil {
		result, err = HostCPU(c)
	} else if c.Collect.Memory != nil {
		result, err = HostMemory(c)
	} else if c.Collect.TCPLoadBalancer != nil {
		result, err = HostTCPLoadBalancer(c)
	} else if c.Collect.HTTPLoadBalancer != nil {
		result, err = HostHTTPLoadBalancer(c)
	} else if c.Collect.DiskUsage != nil {
		result, err = HostDiskUsage(c)
	} else if c.Collect.TCPPortStatus != nil {
		result, err = HostTCPPortStatus(c)
	} else if c.Collect.HTTP != nil {
		result, err = HostHTTP(c)
	} else if c.Collect.Time != nil {
		result, err = HostTime(c)
	} else if c.Collect.BlockDevices != nil {
		result, err = HostBlockDevices(c)
	} else if c.Collect.TCPConnect != nil {
		result, err = HostTCPConnect(c)
	} else if c.Collect.IPV4Interfaces != nil {
		result, err = HostIPV4Interfaces(c)
	} else if c.Collect.FilesystemPerformance != nil {
		result, err = HostFilesystemPerformance(c)
	} else if c.Collect.Certificate != nil {
		result, err = HostCertificate(c)
	} else {
		err = errors.New("no spec found to run")
		return
	}
	if err != nil {
		return
	}

	return
}

func (c *HostCollector) GetDisplayName() string {
	return c.Collect.GetName()
}
