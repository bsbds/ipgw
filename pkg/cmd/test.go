package cmd

import (
	"strconv"

	"github.com/neucn/ipgw/pkg/console"
	"github.com/neucn/ipgw/pkg/handler"
	"github.com/urfave/cli/v2"
)

var (
	TestCommand = &cli.Command{
		Name:  "test",
		Usage: "test whether is connected to the campus network and whether has logged in ipgw",
		Action: func(ctx *cli.Context) error {
			mark := uint32(0)
			if c := ctx.String("fwmark"); c != "" {
				num, err := strconv.ParseUint(c, 10, 32)
				if err != nil {
					return err
				}
				mark = uint32(num)
			}
			h := handler.NewIpgwHandler(mark)
			connected, loggedIn := h.IsConnectedAndLoggedIn()
			console.Info("campus network:   ")
			if connected {
				console.InfoL("connected")
			} else {
				console.InfoL("disconnected")
			}
			console.Info("ipgw logged in:   ")
			if loggedIn {
				console.InfoL("yes")
			} else {
				console.InfoL("no")
			}
			return nil
		},
		OnUsageError: onUsageError,
	}
)
