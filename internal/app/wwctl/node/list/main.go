package list

import (
	"fmt"
	"github.com/hpcng/warewulf/internal/pkg/node"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"github.com/spf13/cobra"
	"os"
)



func CobraRunE(cmd *cobra.Command, args []string) error {
	var err error
	var nodes []node.NodeInfo

	n, err := node.New()
	if err != nil {
		wwlog.Printf(wwlog.ERROR, "Could not open node configuration: %s\n", err)
		os.Exit(1)
	}

	if len(args) > 0 {
		nodes, err = n.SearchByNameList(args)
	} else {
		nodes, err = n.FindAllNodes()
	}
	if err != nil {
		wwlog.Printf(wwlog.ERROR, "Cloud not get nodeList: %s\n", err)
		os.Exit(1)
	}

	if ShowNet == true {
		fmt.Printf("%-22s %-10s %-20s %-16s %-16s %-16s %s\n", "NODE NAME", "DEVICE", "HWADDR", "IPADDR", "NETMASK", "GATEWAY", "TYPE")

		for _, node := range nodes {
			for name, dev := range node.NetDevs {
				fmt.Printf("%-22s %-10s %-20s %-16s %-16s %-16s %s\n", node.Fqdn, name, dev.Hwaddr, dev.Ipaddr, dev.Netmask, dev.Gateway, dev.Type)
			}
		}

	} else if ShowIpmi == true {
		fmt.Printf("%-22s %-16s %-20s %-20s\n", "NODE NAME", "IPMI IPADDR", "IPMI USERNAME", "IPMI PASSWORD")

		for _, node := range nodes {
			fmt.Printf("%-22s %-16s %-20s %-20s\n", node.Fqdn, node.IpmiIpaddr, node.IpmiUserName, node.IpmiPassword)
		}

	} else {
		fmt.Printf("%-22s %-30s %-30s %-16s %-16s\n", "NODE NAME", "KERNEL NAME", "VNFS IMAGE", "SYSTEM OVERLAY", "RUNTIME OVERLAY")

		for _, node := range nodes {
			fmt.Printf("%-22s %-30s %-30s %-16s %-16s\n", node.Fqdn, node.KernelVersion, node.Vnfs, node.SystemOverlay, node.RuntimeOverlay)
		}
	}


	return nil
}
