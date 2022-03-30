// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2alpha1 "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/typed/cilium.io/v2alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCiliumV2alpha1 struct {
	*testing.Fake
}

func (c *FakeCiliumV2alpha1) CiliumClusterwideEnvoyConfigs() v2alpha1.CiliumClusterwideEnvoyConfigInterface {
	return &FakeCiliumClusterwideEnvoyConfigs{c}
}

func (c *FakeCiliumV2alpha1) CiliumEgressNATPolicies() v2alpha1.CiliumEgressNATPolicyInterface {
	return &FakeCiliumEgressNATPolicies{c}
}

func (c *FakeCiliumV2alpha1) CiliumEndpointSlices() v2alpha1.CiliumEndpointSliceInterface {
	return &FakeCiliumEndpointSlices{c}
}

func (c *FakeCiliumV2alpha1) CiliumEnvoyConfigs() v2alpha1.CiliumEnvoyConfigInterface {
	return &FakeCiliumEnvoyConfigs{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCiliumV2alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
