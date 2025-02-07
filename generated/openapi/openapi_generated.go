//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
  Copyright The Kubernetes Authors.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package openapi

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.Containerlab": schema_clabernetes_apis_topology_v1alpha1_Containerlab(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ContainerlabList": schema_clabernetes_apis_topology_v1alpha1_ContainerlabList(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ContainerlabSpec": schema_clabernetes_apis_topology_v1alpha1_ContainerlabSpec(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ContainerlabStatus": schema_clabernetes_apis_topology_v1alpha1_ContainerlabStatus(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ExposedPorts": schema_clabernetes_apis_topology_v1alpha1_ExposedPorts(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.FileFromConfigMap": schema_clabernetes_apis_topology_v1alpha1_FileFromConfigMap(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.Kne": schema_clabernetes_apis_topology_v1alpha1_Kne(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.KneList": schema_clabernetes_apis_topology_v1alpha1_KneList(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.KneSpec": schema_clabernetes_apis_topology_v1alpha1_KneSpec(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.KneStatus": schema_clabernetes_apis_topology_v1alpha1_KneStatus(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.LinkEndpoint": schema_clabernetes_apis_topology_v1alpha1_LinkEndpoint(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.TopologyCommonSpec": schema_clabernetes_apis_topology_v1alpha1_TopologyCommonSpec(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.TopologyStatus": schema_clabernetes_apis_topology_v1alpha1_TopologyStatus(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.Tunnel": schema_clabernetes_apis_topology_v1alpha1_Tunnel(
			ref,
		),
	}
}

func schema_clabernetes_apis_topology_v1alpha1_Containerlab(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Containerlab is an object that holds a \"normal\" containerlab topology file and any additional data necessary to \"clabernetes-ify\" it.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref: ref(
								"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ContainerlabSpec",
							),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref: ref(
								"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ContainerlabStatus",
							),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ContainerlabSpec", "github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ContainerlabStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_clabernetes_apis_topology_v1alpha1_ContainerlabList(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ContainerlabList is a list of Containerlab topology objects.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref: ref(
											"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.Containerlab",
										),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.Containerlab", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_clabernetes_apis_topology_v1alpha1_ContainerlabSpec(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ContainerlabSpec is the spec for a Containerlab topology resource.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"disableExpose": {
						SchemaProps: spec.SchemaProps{
							Description: "DisableExpose indicates if exposing nodes via LoadBalancer service should be disabled, by default any mapped ports in a containerlab topology will be exposed.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"disableAutoExpose": {
						SchemaProps: spec.SchemaProps{
							Description: "DisableAutoExpose disables the automagic exposing of ports for a given topology. When this setting is disabled clabernetes will not auto add ports so if you want to expose (via a load balancer service) you will need to have ports outlined in your containerlab config (or equivalent for kne). When this is `false` (default), clabernetes will add and expose the following list of ports to whatever ports you have already defined:\n\n21    - tcp - ftp 22    - tcp - ssh 23    - tcp - telnet 80    - tcp - http 161   - udp - snmp 443   - tcp - https 830   - tcp - netconf (over ssh) 5000  - tcp - telnet for vrnetlab qemu host 5900  - tcp - vnc 6030  - tcp - gnmi (arista default) 9339  - tcp - gnmi/gnoi 9340  - tcp - gribi 9559  - tcp - p4rt 57400 - tcp - gnmi (nokia srl/sros default)\n\nThis setting is *ignored completely* if `DisableExpose` is true!",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"insecureRegistries": {
						SchemaProps: spec.SchemaProps{
							Description: "InsecureRegistries is a slice of strings of insecure registries to configure in the launcher pods.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"filesFromConfigMap": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "FilesFromConfigMap is a slice of FileFromConfigMap that define the configmap/path and node and path on a launcher node that the file should be mounted to. If the path is not provided the configmap is mounted in its entirety (like normal k8s things), so you *probably* want to specify the sub path unless you are sure what you're doing!",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref: ref(
											"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.FileFromConfigMap",
										),
									},
								},
							},
						},
					},
					"containerlabDebug": {
						SchemaProps: spec.SchemaProps{
							Description: "ContainerlabDebug sets the `--debug` flag when invoking containerlab in the launcher pods. This is disabled by default.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Description: "Config is a \"normal\" containerlab configuration file.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"config"},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.FileFromConfigMap"},
	}
}

func schema_clabernetes_apis_topology_v1alpha1_ContainerlabStatus(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ContainerlabStatus is the status for a Containerlab topology resource.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"configs": {
						SchemaProps: spec.SchemaProps{
							Description: "Configs is a map of node name -> clab config -- in other words, this is the original containerlab configuration broken up and modified to use multi-node topology setup (via host links+vxlan). This is stored as a raw message so we don't have any weirdness w/ yaml tags instead of json tags in clab things, and so we kube builder doesnt poop itself.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configsHash": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigsHash is a hash of the last stored Configs data.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tunnels": {
						SchemaProps: spec.SchemaProps{
							Description: "Tunnels is a mapping of tunnels that need to be configured between nodes (nodes:[]tunnels).",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Ref: ref(
														"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.Tunnel",
													),
												},
											},
										},
									},
								},
							},
						},
					},
					"nodeExposedPorts": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeExposedPorts holds a map of (containerlab) nodes and their exposed ports (via load balancer).",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref(
											"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ExposedPorts",
										),
									},
								},
							},
						},
					},
					"nodeExposedPortsHash": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeExposedPortsHash is a hash of the last stored NodeExposedPorts data.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{
					"configs",
					"configsHash",
					"tunnels",
					"nodeExposedPorts",
					"nodeExposedPortsHash",
				},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ExposedPorts", "github.com/srl-labs/clabernetes/apis/topology/v1alpha1.Tunnel"},
	}
}

func schema_clabernetes_apis_topology_v1alpha1_ExposedPorts(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExposedPorts holds information about exposed ports.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"loadBalancerAddress": {
						SchemaProps: spec.SchemaProps{
							Description: "LoadBalancerAddress holds the address assigned to the load balancer exposing ports for a given node.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tcpPorts": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "TCPPorts is a list of TCP ports exposed on the LoadBalancer service.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: 0,
										Type:    []string{"integer"},
										Format:  "int32",
									},
								},
							},
						},
					},
					"udpPorts": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "UDPPorts is a list of UDP ports exposed on the LoadBalancer service.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: 0,
										Type:    []string{"integer"},
										Format:  "int32",
									},
								},
							},
						},
					},
				},
				Required: []string{"loadBalancerAddress", "tcpPorts", "udpPorts"},
			},
		},
	}
}

func schema_clabernetes_apis_topology_v1alpha1_FileFromConfigMap(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FileFromConfigMap represents a file that you would like to mount (from a configmap) in the launcher pod for a given node.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nodeName": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeName is the name of the node (as in node from the clab topology) that the file should be mounted for.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"filePath": {
						SchemaProps: spec.SchemaProps{
							Description: "FilePath is the path to mount the file.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configMapName": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigMapName is the name of the configmap to mount.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configMapPath": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigMapPath is the path/key in the configmap to mount, if not specified the configmap will be mounted without a sub-path.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"nodeName", "filePath", "configMapName"},
			},
		},
	}
}

func schema_clabernetes_apis_topology_v1alpha1_Kne(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Kne is an object that holds a \"normal\" kne topology file and any additional data necessary to \"clabernetes-ify\" it.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref: ref(
								"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.KneSpec",
							),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref: ref(
								"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.KneStatus",
							),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.KneSpec", "github.com/srl-labs/clabernetes/apis/topology/v1alpha1.KneStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_clabernetes_apis_topology_v1alpha1_KneList(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KneList is a list of Kne topology objects.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref: ref(
											"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.Kne",
										),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.Kne", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_clabernetes_apis_topology_v1alpha1_KneSpec(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KneSpec is the spec for a Kne topology resource.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"disableExpose": {
						SchemaProps: spec.SchemaProps{
							Description: "DisableExpose indicates if exposing nodes via LoadBalancer service should be disabled, by default any mapped ports in a containerlab topology will be exposed.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"disableAutoExpose": {
						SchemaProps: spec.SchemaProps{
							Description: "DisableAutoExpose disables the automagic exposing of ports for a given topology. When this setting is disabled clabernetes will not auto add ports so if you want to expose (via a load balancer service) you will need to have ports outlined in your containerlab config (or equivalent for kne). When this is `false` (default), clabernetes will add and expose the following list of ports to whatever ports you have already defined:\n\n21    - tcp - ftp 22    - tcp - ssh 23    - tcp - telnet 80    - tcp - http 161   - udp - snmp 443   - tcp - https 830   - tcp - netconf (over ssh) 5000  - tcp - telnet for vrnetlab qemu host 5900  - tcp - vnc 6030  - tcp - gnmi (arista default) 9339  - tcp - gnmi/gnoi 9340  - tcp - gribi 9559  - tcp - p4rt 57400 - tcp - gnmi (nokia srl/sros default)\n\nThis setting is *ignored completely* if `DisableExpose` is true!",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"insecureRegistries": {
						SchemaProps: spec.SchemaProps{
							Description: "InsecureRegistries is a slice of strings of insecure registries to configure in the launcher pods.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"filesFromConfigMap": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "FilesFromConfigMap is a slice of FileFromConfigMap that define the configmap/path and node and path on a launcher node that the file should be mounted to. If the path is not provided the configmap is mounted in its entirety (like normal k8s things), so you *probably* want to specify the sub path unless you are sure what you're doing!",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref: ref(
											"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.FileFromConfigMap",
										),
									},
								},
							},
						},
					},
					"containerlabDebug": {
						SchemaProps: spec.SchemaProps{
							Description: "ContainerlabDebug sets the `--debug` flag when invoking containerlab in the launcher pods. This is disabled by default.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"topology": {
						SchemaProps: spec.SchemaProps{
							Description: "Topology is a \"normal\" kne topology proto file.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"topology"},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.FileFromConfigMap"},
	}
}

func schema_clabernetes_apis_topology_v1alpha1_KneStatus(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KneStatus is the status for a Kne topology resource.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"configs": {
						SchemaProps: spec.SchemaProps{
							Description: "Configs is a map of node name -> clab config -- in other words, this is the original containerlab configuration broken up and modified to use multi-node topology setup (via host links+vxlan). This is stored as a raw message so we don't have any weirdness w/ yaml tags instead of json tags in clab things, and so we kube builder doesnt poop itself.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configsHash": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigsHash is a hash of the last stored Configs data.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tunnels": {
						SchemaProps: spec.SchemaProps{
							Description: "Tunnels is a mapping of tunnels that need to be configured between nodes (nodes:[]tunnels).",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Ref: ref(
														"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.Tunnel",
													),
												},
											},
										},
									},
								},
							},
						},
					},
					"nodeExposedPorts": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeExposedPorts holds a map of (containerlab) nodes and their exposed ports (via load balancer).",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref(
											"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ExposedPorts",
										),
									},
								},
							},
						},
					},
					"nodeExposedPortsHash": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeExposedPortsHash is a hash of the last stored NodeExposedPorts data.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{
					"configs",
					"configsHash",
					"tunnels",
					"nodeExposedPorts",
					"nodeExposedPortsHash",
				},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ExposedPorts", "github.com/srl-labs/clabernetes/apis/topology/v1alpha1.Tunnel"},
	}
}

func schema_clabernetes_apis_topology_v1alpha1_LinkEndpoint(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "LinkEndpoint is a simple struct to hold node/interface name info for a given link.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nodeName": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeName is the name of the node this link resides on.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"interfaceName": {
						SchemaProps: spec.SchemaProps{
							Description: "InterfaceName is the name of the interface on the node this link is on.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"nodeName", "interfaceName"},
			},
		},
	}
}

func schema_clabernetes_apis_topology_v1alpha1_TopologyCommonSpec(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TopologyCommonSpec holds fields that are common across different CR types for their spec.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"disableExpose": {
						SchemaProps: spec.SchemaProps{
							Description: "DisableExpose indicates if exposing nodes via LoadBalancer service should be disabled, by default any mapped ports in a containerlab topology will be exposed.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"disableAutoExpose": {
						SchemaProps: spec.SchemaProps{
							Description: "DisableAutoExpose disables the automagic exposing of ports for a given topology. When this setting is disabled clabernetes will not auto add ports so if you want to expose (via a load balancer service) you will need to have ports outlined in your containerlab config (or equivalent for kne). When this is `false` (default), clabernetes will add and expose the following list of ports to whatever ports you have already defined:\n\n21    - tcp - ftp 22    - tcp - ssh 23    - tcp - telnet 80    - tcp - http 161   - udp - snmp 443   - tcp - https 830   - tcp - netconf (over ssh) 5000  - tcp - telnet for vrnetlab qemu host 5900  - tcp - vnc 6030  - tcp - gnmi (arista default) 9339  - tcp - gnmi/gnoi 9340  - tcp - gribi 9559  - tcp - p4rt 57400 - tcp - gnmi (nokia srl/sros default)\n\nThis setting is *ignored completely* if `DisableExpose` is true!",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"insecureRegistries": {
						SchemaProps: spec.SchemaProps{
							Description: "InsecureRegistries is a slice of strings of insecure registries to configure in the launcher pods.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"filesFromConfigMap": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "FilesFromConfigMap is a slice of FileFromConfigMap that define the configmap/path and node and path on a launcher node that the file should be mounted to. If the path is not provided the configmap is mounted in its entirety (like normal k8s things), so you *probably* want to specify the sub path unless you are sure what you're doing!",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref: ref(
											"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.FileFromConfigMap",
										),
									},
								},
							},
						},
					},
					"containerlabDebug": {
						SchemaProps: spec.SchemaProps{
							Description: "ContainerlabDebug sets the `--debug` flag when invoking containerlab in the launcher pods. This is disabled by default.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.FileFromConfigMap"},
	}
}

func schema_clabernetes_apis_topology_v1alpha1_TopologyStatus(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TopologyStatus is a common struct used inline as CR status for topology resources.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"configs": {
						SchemaProps: spec.SchemaProps{
							Description: "Configs is a map of node name -> clab config -- in other words, this is the original containerlab configuration broken up and modified to use multi-node topology setup (via host links+vxlan). This is stored as a raw message so we don't have any weirdness w/ yaml tags instead of json tags in clab things, and so we kube builder doesnt poop itself.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configsHash": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigsHash is a hash of the last stored Configs data.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tunnels": {
						SchemaProps: spec.SchemaProps{
							Description: "Tunnels is a mapping of tunnels that need to be configured between nodes (nodes:[]tunnels).",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Ref: ref(
														"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.Tunnel",
													),
												},
											},
										},
									},
								},
							},
						},
					},
					"nodeExposedPorts": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeExposedPorts holds a map of (containerlab) nodes and their exposed ports (via load balancer).",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref(
											"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ExposedPorts",
										),
									},
								},
							},
						},
					},
					"nodeExposedPortsHash": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeExposedPortsHash is a hash of the last stored NodeExposedPorts data.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{
					"configs",
					"configsHash",
					"tunnels",
					"nodeExposedPorts",
					"nodeExposedPortsHash",
				},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/topology/v1alpha1.ExposedPorts", "github.com/srl-labs/clabernetes/apis/topology/v1alpha1.Tunnel"},
	}
}

func schema_clabernetes_apis_topology_v1alpha1_Tunnel(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Tunnel represents a VXLAN tunnel between clabernetes nodes (as configured by containerlab).",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"id": {
						SchemaProps: spec.SchemaProps{
							Description: "ID is the VXLAN ID (vnid) for the tunnel.",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"localNodeName": {
						SchemaProps: spec.SchemaProps{
							Description: "LocalNodeName is the name of the local node for this tunnel.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"remoteName": {
						SchemaProps: spec.SchemaProps{
							Description: "RemoteName is the name of the service to contact the remote end of the tunnel.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"remoteNodeName": {
						SchemaProps: spec.SchemaProps{
							Description: "RemoteNodeName is the name of the remote node.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"localLinkName": {
						SchemaProps: spec.SchemaProps{
							Description: "LocalLinkName is the local link name for the local end of the tunnel.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"remoteLinkName": {
						SchemaProps: spec.SchemaProps{
							Description: "RemoteLinkName is the remote link name for the remote end of the tunnel.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{
					"id",
					"localNodeName",
					"remoteName",
					"remoteNodeName",
					"localLinkName",
					"remoteLinkName",
				},
			},
		},
	}
}
