package copy

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"

	"github.com/axsh/openvdc/api"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	ClientConfig *ssh.ClientConfig
	Host         string
	Port         string
}

func NewClient(cr *api.CopyReply) (*Client, error) {
	host, port, err := net.SplitHostPort(cr.GetAddress())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Invalid ssh host address: %s", cr.GetAddress()))
	}

	config := &ssh.ClientConfig{
		User: cr.GetInstanceId(),
	}

	return &Client{
		ClientConfig: config,
		Host:         host,
		Port:         port,
	}, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func (c *Client) CopyFile(filePath string, instanceDir string) error {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("Source file not found")
		}
		return errors.Wrap(err, "Couldn't open file")
	}
	defer file.Close()

	srcFilename := filepath.Base(filePath)

	fileinfo, err := file.Stat()
	if err != nil {
		return errors.Wrap(err, "File.Stat failed")
	}

	cl, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", c.Host, c.Port), c.ClientConfig)

	if err != nil {
		return errors.Wrap(err, "ssh.Dial failed")
	}

	session, err := cl.NewSession()
	if err != nil {
		return errors.Wrap(err, "Failed to create session")
	}
	defer session.Close()

	go func() {
		w, _ := session.StdinPipe()
		defer w.Close()
		fmt.Fprintln(w, "C0655", int64(fileinfo.Size()), srcFilename)
		io.Copy(w, file) //Todo: handle error
		fmt.Fprintln(w, "\x00")
	}()

	session.Run("/usr/bin/scp -t " + instanceDir)

	return nil
}
