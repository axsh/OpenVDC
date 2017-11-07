package cmd

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/axsh/openvdc/api"
	"github.com/axsh/openvdc/cmd/openvdc/internal/util"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var copyCmd = &cobra.Command{
	Use:   "copy [File src path] [Instance ID]:/[file dest path]",
	Short: "Copy files to and between instances",
	Long:  "Copy files to and between instances",
	Example: `
	% openvdc copy 1.txt i-xxxxxxx:/tmp/1.txt
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			log.Fatalf("Please provide a source path.")
		}
		if len(args) == 1 {
			log.Fatalf("Please provide a destination path.")
		}

		//src := args[0]
		dest := args[1]

		path := strings.Split(dest, ":")
        	fmt.Sprintf("value: %s", path[0])

        	instanceID, instanceDir := path[0], path[1]

        	if instanceID == "" {
               		log.Fatalf("Invalid Instance ID")
        	}

        	if instanceDir == "" {
                	log.Fatalf("Invalid destination path")
        	}

		req := &api.CopyRequest{
			InstanceId: instanceID,
		}

		return util.RemoteCall(func(conn *grpc.ClientConn) error {
			c := api.NewInstanceClient(conn)
			res, err := c.Copy(context.Background(), req)
			if err != nil {
				log.WithError(err).Fatal("Disconnected abnormaly")
				return err
			}

			fmt.Println(res.GetAddress())
			return err
		})
	},
}
