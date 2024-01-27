package show

import (
	"fmt"

	"github.com/warewulf/warewulf/internal/pkg/api/container"
	"github.com/warewulf/warewulf/internal/pkg/api/routes/wwapiv1"

	"github.com/spf13/cobra"
)

func CobraRunE(cmd *cobra.Command, args []string) (err error) {

	csp := &wwapiv1.ContainerShowParameter{
		ContainerName: args[0],
	}

	var r *wwapiv1.ContainerShowResponse
	r, err = container.ContainerShow(csp)
	if err != nil {
		return
	}

	if !ShowAll {
		fmt.Printf("%s\n", r.Rootfs)
	} else {
		kernelVersion := r.KernelVersion
		if kernelVersion == "" {
			kernelVersion = "not found"
		}
		fmt.Printf("Name: %s\n", r.Name)
		fmt.Printf("KernelVersion: %s\n", kernelVersion)
		fmt.Printf("Rootfs: %s\n", r.Rootfs)
		fmt.Printf("Nr nodes: %d\n", len(r.Nodes))
		fmt.Printf("Nodes: %s\n", r.Nodes)
	}
	return
}
