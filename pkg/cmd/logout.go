package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/neucn/ipgw/pkg/console"
	"github.com/neucn/ipgw/pkg/handler"
	"github.com/urfave/cli/v2"
)

var (
	LogoutCommand = &cli.Command{
		Name:  "logout",
		Usage: "logout ipgw",
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
			if !connected {
				return errors.New("not in campus network")
			}
			if !loggedIn {
				return errors.New("not logged in yet")
			}
			info := h.GetInfo()
			if err := h.Logout(); err != nil {
				return fmt.Errorf("fail to logout account '%s':\n\t%v", info.Username, err)
			}
			console.InfoF("logout account '%s' successfully\n", info.Username)
			return nil
		},
		OnUsageError: onUsageError,
	}
)
