package commands

import (
	"time"

	"encoding/json"
	pb "github.com/sonm-io/core/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func init() {
	hubRootCmd.AddCommand(hubPingCmd, hubStatusCmd, hubFindCmd)
}

// --- hub commands
var hubRootCmd = &cobra.Command{
	Use:     "hub",
	Short:   "Operations with hub",
	PreRunE: checkHubAddressIsSet,
}

var hubPingCmd = &cobra.Command{
	Use:     "ping",
	Short:   "Ping the hub",
	PreRunE: checkHubAddressIsSet,
	Run: func(cmd *cobra.Command, args []string) {
		itr, err := NewGrpcInteractor(hubAddress, timeout)
		if err != nil {
			showError(cmd, "Cannot connect to hub", err)
			return
		}
		hubPingCmdRunner(cmd, itr)
	},
}

var hubStatusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show hub status",
	PreRunE: checkHubAddressIsSet,
	Run: func(cmd *cobra.Command, args []string) {
		itr, err := NewGrpcInteractor(hubAddress, timeout)
		if err != nil {
			showError(cmd, "Cannot connect to hub", err)
			return
		}

		hubStatusCmdRunner(cmd, itr)
	},
}

var hubFindCmd = &cobra.Command{
	Use:   "find",
	Short: "Find hubs into SONM network",
	Run: func(cmd *cobra.Command, _ []string) {
		itr, err := NewGrpcInteractor(hubAddress, timeout)
		if err != nil {
			showError(cmd, "Cannot connect to hub", err)
			return
		}
		hubFindCmdRunner(cmd, itr)
	},
}

func hubPingCmdRunner(cmd *cobra.Command, interactor CliInteractor) {
	_, err := interactor.HubPing(context.Background())
	if err != nil {
		showError(cmd, "Ping failed", err)
		return
	}

	showOk(cmd)
}

func hubStatusCmdRunner(cmd *cobra.Command, interactor CliInteractor) {
	stat, err := interactor.HubStatus(context.Background())
	if err != nil {
		showError(cmd, "Cannot get status", err)
		return
	}

	printHubStatus(cmd, stat)
}

func hubFindCmdRunner(cmd *cobra.Command, interactor CliInteractor) {
	hubs, err := interactor.HubFind(context.Background(), 30*time.Second)
	if err != nil {
		showError(cmd, "Cannot get Hubs", err)
		return
	}

	for _, h := range hubs {
		if h.PublicKey != nil {
			cmd.Printf("Hub: %s\r\n", h.Address)
		} else {
			cmd.Printf("Hub: %s (ANON)\r\n", h.Address)
		}
	}
}

func printHubStatus(cmd *cobra.Command, stat *pb.HubStatusReply) {
	if isSimpleFormat() {
		cmd.Printf("Connected miners: %d\r\n", stat.MinerCount)
		cmd.Printf("Uptime:           %s\r\n", (time.Second * time.Duration(stat.Uptime)).String())
	} else {
		b, _ := json.Marshal(stat)
		cmd.Println(string(b))
	}
}
