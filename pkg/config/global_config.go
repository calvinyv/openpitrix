// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/coreos/etcd/mvcc/mvccpb"

	"openpitrix.io/openpitrix/pkg/etcd"
	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/utils/yaml"
)

type GlobalConfig struct {
	Repo    RepoServiceConfig      `json:"repo"`
	Cluster ClusterServiceConfig   `json:"cluster"`
	Runtime map[string]ImageConfig `json:"runtime"`
}

type RepoServiceConfig struct {
	Cron string `json:"cron"`
}

type ClusterServiceConfig struct {
	Plugins       []string `json:"plugins"`
	FrontgateConf string   `json:"frontgate_conf"`
}

type ImageConfig struct {
	ApiServer string `json:"api_server"`
	Zone      string `json:"zone"`
	ImageId   string `json:"image_id"`
}

func (g *GlobalConfig) GetRuntimeImageId(apiServer, zone string) (string, error) {
	if strings.HasPrefix(apiServer, "https://") {
		apiServer = strings.Split(apiServer, "https://")[1]
	}
	for _, imageConfig := range g.Runtime {
		if imageConfig.ApiServer == apiServer && imageConfig.Zone == zone {
			return imageConfig.ImageId, nil
		}
	}
	logger.Errorf("No such runtime image with api server [%s] zone [%s]. ", apiServer, zone)
	return "", fmt.Errorf("no such runtime image with api server [%s] zone [%s]. ", apiServer, zone)
}

const InitialGlobalConfig = `
repo:
  # cron usage: https://godoc.org/github.com/robfig/cron#hdr-Usage
  #
  #   "@every 1h30m" means Every hour thirty
  #   "@hourly" means Every hour
  #   "0 30 * * * *" means Every hour on the half hour
  #
  #	  Field name   | Mandatory? | Allowed values  | Allowed special characters
  #	  ----------   | ---------- | --------------  | --------------------------
  #	  Seconds      | Yes        | 0-59            | * / , -
  #	  Minutes      | Yes        | 0-59            | * / , -
  #	  Hours        | Yes        | 0-23            | * / , -
  #	  Day of month | Yes        | 1-31            | * / , - ?
  #	  Month        | Yes        | 1-12 or JAN-DEC | * / , -
  #	  Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
  #
  cron: "0 30 4 * * *"
cluster:
  plugins:
    - qingcloud
    - kubernetes
runtime:
  qingcloud_pek3a:
    api_server: api.qingcloud.com
    zone: pek3a
    image_id: img-abcdefgh
  qingcloud_sh1a:
    api_server: api.qingcloud.com
    zone: sh1a
    image_id: img-abcdefgh
`

const (
	EtcdPrefix      = "openpitrix/"
	GlobalConfigKey = "global_config"
	DlockKey        = "dlock_" + GlobalConfigKey
)

type Watcher chan *GlobalConfig

func WatchGlobalConfig(etcd *etcd.Etcd, watcher Watcher) error {
	ctx := context.Background()
	var globalConfig GlobalConfig
	err := etcd.Dlock(ctx, DlockKey, func() error {
		// get value
		get, err := etcd.Get(ctx, GlobalConfigKey)
		if err != nil {
			return err
		}
		// parse value
		if get.Count == 0 {
			logger.Debugf("Cannot get global config, put the initial string. [%s]", InitialGlobalConfig)
			globalConfig = DecodeInitConfig()
			_, err = etcd.Put(ctx, GlobalConfigKey, InitialGlobalConfig)
			if err != nil {
				return err
			}
		} else {
			err = yaml.Decode(get.Kvs[0].Value, &globalConfig)
			if err != nil {
				return err
			}
		}
		logger.Debugf("Global config update to [%+v]", globalConfig)
		// send it back
		watcher <- &globalConfig
		return nil
	})

	// watch
	go func() {
		logger.Debugf("Start watch global config")
		watchRes := etcd.Watch(ctx, GlobalConfigKey)
		for res := range watchRes {
			for _, ev := range res.Events {
				if ev.Type == mvccpb.PUT {
					globalConfig = GlobalConfig{}
					//logger.Debugf("Got updated global config from etcd, try to decode with yaml")
					err = yaml.Decode(ev.Kv.Value, &globalConfig)
					if err != nil {
						logger.Errorf("Watch global config from etcd found error: %+v", err)
					} else {
						watcher <- &globalConfig
					}
				}
			}
		}
	}()
	return err
}

func DecodeInitConfig() GlobalConfig {
	var globalConfig GlobalConfig
	err := yaml.Decode([]byte(InitialGlobalConfig), &globalConfig)
	if err != nil {
		fmt.Print("InitialGlobalConfig is invalid, please fix it")
		panic(err)
	}
	return globalConfig
}

func EncodeGlobalConfig(conf GlobalConfig) string {
	out, err := yaml.Encode(conf)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func init() {
	DecodeInitConfig()
}
