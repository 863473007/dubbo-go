// Copyright 2016-2019 hxmhlt
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster_impl

import (
	"context"
	"fmt"
	"testing"
)
import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/dubbo/go-for-apache-dubbo/cluster/directory"
	"github.com/dubbo/go-for-apache-dubbo/common"
	"github.com/dubbo/go-for-apache-dubbo/protocol"
	"github.com/dubbo/go-for-apache-dubbo/protocol/invocation"
)

func Test_RegAwareInvokeSuccess(t *testing.T) {

	regAwareCluster := NewRegistryAwareCluster()

	invokers := []protocol.Invoker{}
	for i := 0; i < 10; i++ {
		url, _ := common.NewURL(context.TODO(), fmt.Sprintf("dubbo://192.168.1.%v:20000/com.ikurento.user.UserProvider", i))
		invokers = append(invokers, NewMockInvoker(url, 1))
	}

	staticDir := directory.NewStaticDirectory(invokers)
	clusterInvoker := regAwareCluster.Join(staticDir)
	result := clusterInvoker.Invoke(&invocation.RPCInvocation{})
	assert.NoError(t, result.Error())
	count = 0
}

func TestDestroy(t *testing.T) {
	regAwareCluster := NewRegistryAwareCluster()

	invokers := []protocol.Invoker{}
	for i := 0; i < 10; i++ {
		url, _ := common.NewURL(context.TODO(), fmt.Sprintf("dubbo://192.168.1.%v:20000/com.ikurento.user.UserProvider", i))
		invokers = append(invokers, NewMockInvoker(url, 1))
	}

	staticDir := directory.NewStaticDirectory(invokers)
	clusterInvoker := regAwareCluster.Join(staticDir)
	assert.Equal(t, true, clusterInvoker.IsAvailable())
	result := clusterInvoker.Invoke(&invocation.RPCInvocation{})
	assert.NoError(t, result.Error())
	count = 0
	clusterInvoker.Destroy()
	assert.Equal(t, false, clusterInvoker.IsAvailable())

}
