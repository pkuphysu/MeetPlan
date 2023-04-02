package config

import (
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/registry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func NewRegistry() (registry.Registry, error) {
	return etcd.NewEtcdRegistry([]string{})
}
func NewResolver() (discovery.Resolver, error) {
	return etcd.NewEtcdResolver([]string{})
}
