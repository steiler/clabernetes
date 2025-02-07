package topology_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	clabernetesapistopologyv1alpha1 "github.com/srl-labs/clabernetes/apis/topology/v1alpha1"
	clabernetescontrollerstopology "github.com/srl-labs/clabernetes/controllers/topology"
	clabernetestesthelper "github.com/srl-labs/clabernetes/testhelper"
)

const testAllocateTunnelIdsTestName = "tunnels/allocate-tunnel-ids"

// TestAllocateTunnelIds ensures that the tunnel clabernetes controllers VXLAN tunnel ID allocation
// process works as advertised. None of this is "hard" necessarily, but there are a lot of moving
// parts in play to ensure that we use the tunnel IDs consistently and also obviously don't stomp
// on any existing tunnel IDs.
func TestAllocateTunnelIds(t *testing.T) {
	testCases := []struct {
		name             string
		statusTunnels    map[string][]*clabernetesapistopologyv1alpha1.Tunnel
		processedTunnels map[string][]*clabernetesapistopologyv1alpha1.Tunnel
	}{
		{
			name:          "simple",
			statusTunnels: map[string][]*clabernetesapistopologyv1alpha1.Tunnel{},
			processedTunnels: map[string][]*clabernetesapistopologyv1alpha1.Tunnel{
				"srl1": {
					{
						ID:             0,
						LocalNodeName:  "srl1",
						RemoteName:     "topo-1-srl2.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl2",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
				"srl2": {
					{
						ID:             0,
						LocalNodeName:  "srl2",
						RemoteName:     "topo-1-srl1.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl1",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
			},
		},
		{
			name: "simple-existing-status",
			statusTunnels: map[string][]*clabernetesapistopologyv1alpha1.Tunnel{
				"srl1": {
					{
						ID:             0,
						LocalNodeName:  "srl1",
						RemoteName:     "topo-1-srl2.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl2",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
				"srl2": {
					{
						ID:             0,
						LocalNodeName:  "srl2",
						RemoteName:     "topo-1-srl1.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl1",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
			},
			processedTunnels: map[string][]*clabernetesapistopologyv1alpha1.Tunnel{
				"srl1": {
					{
						ID:             0,
						LocalNodeName:  "srl1",
						RemoteName:     "topo-1-srl2.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl2",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
				"srl2": {
					{
						ID:             0,
						LocalNodeName:  "srl2",
						RemoteName:     "topo-1-srl1.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl1",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
			},
		},
		{
			name: "simple-already-allocated-ids",
			statusTunnels: map[string][]*clabernetesapistopologyv1alpha1.Tunnel{
				"srl1": {
					{
						ID:             1,
						LocalNodeName:  "srl1",
						RemoteName:     "topo-1-srl2.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl2",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
				"srl2": {
					{
						ID:             1,
						LocalNodeName:  "srl2",
						RemoteName:     "topo-1-srl1.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl1",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
			},
			processedTunnels: map[string][]*clabernetesapistopologyv1alpha1.Tunnel{
				"srl1": {
					{
						ID:             0,
						LocalNodeName:  "srl1",
						RemoteName:     "topo-1-srl2.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl2",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
				"srl2": {
					{
						ID:             0,
						LocalNodeName:  "srl2",
						RemoteName:     "topo-1-srl1.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl1",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
			},
		},
		{
			name: "simple-weirdly-allocated-ids",
			statusTunnels: map[string][]*clabernetesapistopologyv1alpha1.Tunnel{
				"srl1": {
					{
						ID:             0,
						LocalNodeName:  "srl1",
						RemoteName:     "topo-1-srl2.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl2",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
				"srl2": {
					{
						ID:             1,
						LocalNodeName:  "srl2",
						RemoteName:     "topo-1-srl1.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl1",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
			},
			processedTunnels: map[string][]*clabernetesapistopologyv1alpha1.Tunnel{
				"srl1": {
					{
						ID:             0,
						LocalNodeName:  "srl1",
						RemoteName:     "topo-1-srl2.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl2",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
				"srl2": {
					{
						ID:             0,
						LocalNodeName:  "srl2",
						RemoteName:     "topo-1-srl1.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl1",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
				},
			},
		},
		{
			name:          "meshy-links",
			statusTunnels: map[string][]*clabernetesapistopologyv1alpha1.Tunnel{},
			processedTunnels: map[string][]*clabernetesapistopologyv1alpha1.Tunnel{
				"srl1": {
					{
						ID:             0,
						LocalNodeName:  "srl1",
						RemoteName:     "topo-1-srl2.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl2",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
					{
						ID:             0,
						LocalNodeName:  "srl1",
						RemoteName:     "topo-1-srl3.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl3",
						LocalLinkName:  "e1-2",
						RemoteLinkName: "e1-1",
					},
				},
				"srl2": {
					{
						ID:             0,
						LocalNodeName:  "srl2",
						RemoteName:     "topo-1-srl1.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl1",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-1",
					},
					{
						ID:             0,
						LocalNodeName:  "srl2",
						RemoteName:     "topo-1-srl3.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl3",
						LocalLinkName:  "e1-2",
						RemoteLinkName: "e1-2",
					},
					{
						ID:             0,
						LocalNodeName:  "srl2",
						RemoteName:     "topo-1-srl4.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl4",
						LocalLinkName:  "e1-3",
						RemoteLinkName: "e1-1",
					},
				},
				"srl3": {
					{
						ID:             0,
						LocalNodeName:  "srl3",
						RemoteName:     "topo-1-srl1.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl1",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-2",
					},
					{
						ID:             0,
						LocalNodeName:  "srl3",
						RemoteName:     "topo-1-srl2.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl2",
						LocalLinkName:  "e1-2",
						RemoteLinkName: "e1-2",
					},
					{
						ID:             0,
						LocalNodeName:  "srl3",
						RemoteName:     "topo-1-srl4.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl4",
						LocalLinkName:  "e1-3",
						RemoteLinkName: "e1-2",
					},
				},
				"srl4": {
					{
						ID:             0,
						LocalNodeName:  "srl4",
						RemoteName:     "topo-1-srl2.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl2",
						LocalLinkName:  "e1-1",
						RemoteLinkName: "e1-3",
					},
					{
						ID:             0,
						LocalNodeName:  "srl4",
						RemoteName:     "topo-1-srl3.clabernetes.svc.cluster.local",
						RemoteNodeName: "srl3",
						LocalLinkName:  "e1-2",
						RemoteLinkName: "e1-3",
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				clabernetescontrollerstopology.AllocateTunnelIDs(
					testCase.statusTunnels,
					testCase.processedTunnels,
				)

				got := testCase.processedTunnels

				if *clabernetestesthelper.Update {
					clabernetestesthelper.WriteTestFixtureJSON(
						t,
						fmt.Sprintf(
							"golden/%s/%s.json",
							testAllocateTunnelIdsTestName,
							testCase.name,
						),
						got,
					)
				}

				var want map[string][]*clabernetesapistopologyv1alpha1.Tunnel

				err := json.Unmarshal(
					clabernetestesthelper.ReadTestFixtureFile(
						t,
						fmt.Sprintf(
							"golden/%s/%s.json",
							testAllocateTunnelIdsTestName,
							testCase.name,
						),
					),
					&want,
				)
				if err != nil {
					t.Fatal(err)
				}

				if !reflect.DeepEqual(got, want) {
					clabernetestesthelper.FailOutput(t, got, want)
				}
			},
		)
	}
}
