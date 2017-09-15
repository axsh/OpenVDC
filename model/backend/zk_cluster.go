package backend

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/samuel/go-zookeeper/zk"
)

type ZkCluster struct {
	zkConnection
	// TODO: Add mutex
}

func NewZkClusterBackend() *ZkCluster {
	return &ZkCluster{}
}

func (z *ZkCluster) Register(key string, value []byte) error {
	if !z.isConnected() {
		return ErrConnectionNotReady
	}
	absKey, _ := z.canonKey(key)

	var errRetry = fmt.Errorf("")

	doRegist := func() error {
		exists, stat, wa, err := z.connection().ExistsW(absKey)
		if exists {
			if z.connection().SessionID() == stat.EphemeralOwner {
				// Same owner we can reuse the zNode.
				_, err = z.connection().Set(absKey, value, stat.Version)
				if err != nil {
					log.WithError(err).WithField("zNode", absKey).Error("Failed to reuse the key")
					return err
				}
				return nil
			} else {
				// There are two possible cases here:
				//  1. The zNode has lost owner. will be cleared after session timeout.
				//     => Wait for NodeDeleted event.
				//  2. Someone still owns the zNode.
				//     => Give up sometime later.
				var errTout = fmt.Errorf("Timed out for node registration")

			Done:
				for {
					select {
					case ev := <-wa:
						if ev.Type == zk.EventNodeDeleted {
							break Done
						}
					case <-time.After(10 * time.Second):
						return errTout
					}
				}
			}
		}
		_, err = z.connection().Create(absKey, value, zk.FlagEphemeral, defaultACL)
		if err != nil {
			if err == zk.ErrNodeExists {
				return errRetry
			}
			return err
		}
		return nil
	}

	for i := 0; i < 3; i++ {
		err := doRegist()
		if err == nil {
			return nil
		} else if err != errRetry {
			return err
		}
		log.Warnf("Retrying registration (%d)", i+1)
		time.Sleep(100 * time.Millisecond)
	}
	log.Errorf("Retry exceede.")
	return fmt.Errorf("Retry exceeded")
}

func (z *ZkCluster) Find(key string) (value []byte, err error) {
	if !z.isConnected() {
		return nil, ErrConnectionNotReady
	}
	absKey, _ := z.canonKey(key)
	value, _, err = z.connection().Get(absKey)
	return
}

func (z *ZkCluster) UnRegister(key string) error {
	if !z.isConnected() {
		return ErrConnectionNotReady
	}
	absKey, _ := z.canonKey(key)

	return z.connection().Delete(absKey, versionAny)
}
