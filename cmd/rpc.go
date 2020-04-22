package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/BlockscapeLab/peerdx/rpc"
)

func createRPCCmd() *cobra.Command {
	rpcCmd := &cobra.Command{
		Use:   "rpc",
		Short: "command that make use of the rpc endpoint of nodes",
		Long:  `tools to get an overview of your node's peers via rpc`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	compareCmd := &cobra.Command{
		Use:   "compare [node1 address] [node2 address] <[...]>",
		Short: "compare peers of nodes via rpc",
		Long:  `calls /net_info on all provided nodes and compares the results`,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			loadConfig()
			rpcAddrs := make([]rpc.RPCAddr, 0, len(args))
			for _, arg := range args {
				rpcAddr, err := rpc.CreateRPCAddr(arg, "")
				if err != nil {
					log.Printf("'%s' is not a valid address: %s\n", arg, err.Error())
					os.Exit(1)
				}
				rpcAddrs = append(rpcAddrs, rpcAddr)
			}

			rpc.GetNetInfoAndCompare(rpcAddrs)
		},
	}

	listCmd := &cobra.Command{
		Use:   "list [node address]",
		Short: "list detailed information about peers of a node via rpc",
		Long:  `calls /net_info on provided node and lists detailed easy to read information`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			loadConfig()

			rpcAddr, err := rpc.CreateRPCAddr(args[0], "")
			if err != nil {
				log.Printf("'%s' is not a valid address: %s\n", args[0], err.Error())
				os.Exit(1)
			}

			rpc.ListDetailedPeerInfo(rpcAddr)
		},
	}

	rpcCmd.AddCommand(listCmd, compareCmd)

	return rpcCmd
}
