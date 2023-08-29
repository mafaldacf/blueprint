package memcached

import (
	"fmt"
	"reflect"

	"gitlab.mpi-sws.org/cld/blueprint/blueprint/pkg/blueprint"
	"gitlab.mpi-sws.org/cld/blueprint/blueprint/pkg/core/backend"
	"gitlab.mpi-sws.org/cld/blueprint/blueprint/pkg/core/pointer"
	"gitlab.mpi-sws.org/cld/blueprint/blueprint/pkg/core/process"
	"gitlab.mpi-sws.org/cld/blueprint/blueprint/pkg/core/service"
	"gitlab.mpi-sws.org/cld/blueprint/plugins/golang"
)

type MemcachedProcess struct {
	process.ProcessNode
	backend.Cache

	InstanceName string
	Addr         *pointer.Address
}

type MemcachedGoClient struct {
	golang.Service
	backend.Cache

	InstanceName string
	Addr         *pointer.Address
}

func newMemcachedProcess(name string, addr blueprint.IRNode) (*MemcachedProcess, error) {
	addrNode, is_addr := addr.(*pointer.Address)
	if !is_addr {
		return nil, fmt.Errorf("%s expected %s to be an address but found %s", name, addr.Name(), reflect.TypeOf(addr).String())
	}

	proc := &MemcachedProcess{}
	proc.InstanceName = name
	proc.Addr = addrNode
	return proc, nil
}

func newMemcachedGoClient(name string, addr blueprint.IRNode) (*MemcachedGoClient, error) {
	addrNode, is_addr := addr.(*pointer.Address)
	if !is_addr {
		return nil, fmt.Errorf("%s expected %s to be an address but found %s", name, addr.Name(), reflect.TypeOf(addr).String())
	}

	client := &MemcachedGoClient{}
	client.InstanceName = name
	client.Addr = addrNode
	return client, nil
}

func (n *MemcachedProcess) String() string {
	return n.InstanceName + " = MemcachedProcess(" + n.Addr.Name() + ")"
}

func (n *MemcachedProcess) Name() string {
	return n.InstanceName
}

func (n *MemcachedGoClient) String() string {
	return n.InstanceName + " = MemcachedClient(" + n.Addr.Name() + ")"
}

func (n *MemcachedGoClient) Name() string {
	return n.InstanceName
}

func (n *MemcachedGoClient) GetInterface() *service.ServiceInterface {
	return nil
}

func (node *MemcachedGoClient) ImplementsGolangNode()    {}
func (node *MemcachedGoClient) ImplementsGolangService() {}