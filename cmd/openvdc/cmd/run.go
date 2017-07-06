package cmd

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/axsh/openvdc/api"
	"github.com/axsh/openvdc/cmd/openvdc/internal/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func init() {
	runCmd.Flags().Bool("auto-recovery", true, "Set auto recovery flag")
}

var runCmd = &cobra.Command{
	Use:   "run [ResourceTemplate ID/URI]",
	Short: "Run an instance",
	Long:  "Run an instance",
	Example: `
	% openvdc run centos/7/lxc
	% openvdc run https://raw.githubusercontent.com/axsh/openvdc/master/templates/centos/7/lxc.json
	` + util.ExampleMergeTemplateOptions("openvdc run"),
	DisableFlagParsing: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := util.PreRunHelpFlagCheckAndQuit(cmd, args)
		if err != nil {
			return err
		}
		err = cmd.ParseFlags(args)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		left := cmd.Flags().Args()
		if len(left) < 1 {
			return pflag.ErrHelp
		}

		templateSlug := left[0]
		for i, a := range args {
			if a == templateSlug {
				left = args[i:]
				break
			}
		}
		req := prepareCreateAPICall(templateSlug, left, cmd.Flags())
		return util.RemoteCall(func(conn *grpc.ClientConn) error {
			c := api.NewInstanceClient(conn)
			res, err := c.Run(context.Background(), req)
			if err != nil {
				log.WithError(err).Fatal("Disconnected abnormaly")
				return err
			}
			fmt.Println(res.GetInstanceId())
			return err
		})
	}}
