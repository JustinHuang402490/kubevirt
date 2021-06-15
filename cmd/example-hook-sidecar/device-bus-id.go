/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2018 Red Hat, Inc.
 *
 */

package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net"
	"os"
	"strings"

	//"strings"

	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	//"github.com/clbanning/mxj/v2"

	vmSchema "kubevirt.io/client-go/api/v1"
	"kubevirt.io/client-go/log"
	"kubevirt.io/kubevirt/pkg/hooks"
	hooksInfo "kubevirt.io/kubevirt/pkg/hooks/info"
	hooksV1alpha1 "kubevirt.io/kubevirt/pkg/hooks/v1alpha1"
	hooksV1alpha2 "kubevirt.io/kubevirt/pkg/hooks/v1alpha2"
	domainSchema "kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/api"
)

const (
	baseBoardManufacturerAnnotation = "smbios.vm.kubevirt.io/baseBoardManufacturer"
	deviceBusIdAnnotation           = "custom.kubevirt.io/domain.devices.hostdevices.source.address"
	onDefineDomainLoggingMessage    = "Hook's OnDefineDomain callback method has been called"
)

type infoServer struct {
	Version string
}

func (s infoServer) Info(ctx context.Context, params *hooksInfo.InfoParams) (*hooksInfo.InfoResult, error) {
	log.Log.Info("Hook's Info method has been called")

	return &hooksInfo.InfoResult{
		Name: "smbios",
		Versions: []string{
			s.Version,
		},
		HookPoints: []*hooksInfo.HookPoint{
			{
				Name:     hooksInfo.OnDefineDomainHookPointName,
				Priority: 0,
			},
		},
	}, nil
}

type v1alpha1Server struct{}
type v1alpha2Server struct{}

func (s v1alpha2Server) OnDefineDomain(ctx context.Context, params *hooksV1alpha2.OnDefineDomainParams) (*hooksV1alpha2.OnDefineDomainResult, error) {
	log.Log.Info(onDefineDomainLoggingMessage)
	newDomainXML, err := onDefineDomain(params.GetVmi(), params.GetDomainXML())
	if err != nil {
		return nil, err
	}
	return &hooksV1alpha2.OnDefineDomainResult{
		DomainXML: newDomainXML,
	}, nil
}
func (s v1alpha2Server) PreCloudInitIso(_ context.Context, params *hooksV1alpha2.PreCloudInitIsoParams) (*hooksV1alpha2.PreCloudInitIsoResult, error) {
	return &hooksV1alpha2.PreCloudInitIsoResult{
		CloudInitData: params.GetCloudInitData(),
	}, nil
}

func (s v1alpha1Server) OnDefineDomain(ctx context.Context, params *hooksV1alpha1.OnDefineDomainParams) (*hooksV1alpha1.OnDefineDomainResult, error) {
	log.Log.Info(onDefineDomainLoggingMessage)
	newDomainXML, err := onDefineDomain(params.GetVmi(), params.GetDomainXML())
	if err != nil {
		return nil, err
	}
	return &hooksV1alpha1.OnDefineDomainResult{
		DomainXML: newDomainXML,
	}, nil
}

func onDefineDomain(vmiJSON []byte, domainXML []byte) ([]byte, error) {
	log.Log.Info(onDefineDomainLoggingMessage)

	vmiSpec := vmSchema.VirtualMachineInstance{}
	err := json.Unmarshal(vmiJSON, &vmiSpec)
	if err != nil {
		log.Log.Reason(err).Errorf("Failed to unmarshal given VMI spec: %s", vmiJSON)
		panic(err)
	}

	annotations := vmiSpec.GetAnnotations()

	if _, found := annotations[deviceBusIdAnnotation]; !found {
		log.Log.Info("Device bus id hook sidecar was requested, but no attributes provided. Returning original domain spec")
		return domainXML, nil
	}

	//domainSpecMap, err := mxj.NewMapXml(domainSchema.DomainSpec{})
	domainSpec := domainSchema.DomainSpec{}
	err = xml.Unmarshal(domainXML, &domainSpec)
	if err != nil {
		log.Log.Reason(err).Errorf("Failed to unmarshal given domain spec: %s", domainXML)
		panic(err)
	}

	////domainSpec.OS.SMBios = &domainSchema.SMBios{Mode: "sysinfo"}

	/*if domainSpec.SysInfo == nil {
		domainSpec.SysInfo = &domainSchema.SysInfo{}
	}*/
	////domainSpec.SysInfo.Type = "smbios"
	indexRecord := 0
	if deviceBusId, found := annotations[deviceBusIdAnnotation]; found {
		fmt.Println("\ndeviceBusId = ", deviceBusId)
		//runes := []rune(deviceBusId)

		/*for i, value := range runes {
			fmt.Println("index = ", i, ", value = ", string(value))
		}*/

		devices := domainSpec.Devices
		hostDevices := devices.HostDevices
		fmt.Println("\thostDevices = ", hostDevices)

		for index, hostDevice := range hostDevices {
			address := hostDevice.Source.Address
			//if hostDevice.Source != nil {
			fmt.Println("\n\tdomainSpec.Devices.hostDevice.Source = ", hostDevice.Source)
			fmt.Println("\tdomainSpec.Devices.hostDevice.Source.Address = ", address)

			strAddress := fmt.Sprintf("%s", address)
			//}

			//composeDeviceBusId := (
			//    ("0x"+ string(runes[0:1])) +
			//    ("0x"+ string(runes[3:4])) +
			//    ("0x"+ string(runes[6])))
			//fmt.Println("composeDeviceBusId = ", composeDeviceBusId)
			if strings.Contains(strAddress, deviceBusId) {
				fmt.Println("\t\tDevice bus id = ", deviceBusId)
				romFile := "/usr/share/NVIDIA_GTX1080Ti.dump"
				fmt.Println("\t\tRom file location = ", romFile)

				fmt.Println("\t\tindex = ", index)
				indexRecord = index
				//domainSpec.Devices.HostDevices[index]. = append(domainSpec.Devices.HostDevices, romFile)
				domainSpec.Devices.HostDevices[index].RomFile = romFile

				//domainSpec.Devices.HostDevices[index].Type = "changedPCI"
			} else {
				fmt.Println("\t\tstrAddress = ", strAddress)
			}
		}
		//if domainSpec.Devices.HostDevices
		//domainSpec.SysInfo.BaseBoard = append(domainSpec.SysInfo.BaseBoard, domainSchema.Entry{
		//	Name:  "manufacturer",
		//	Value: baseBoardManufacturer,
		//})
	}

	//newDomainXML, err := m.Xml()

	fmt.Println("domainSpec.Devices.HostDevices[indexRecord].RomFile = ", domainSpec.Devices.HostDevices[indexRecord].RomFile)
	newDomainXML, err := xml.Marshal(domainSpec)
	if err != nil {
		log.Log.Reason(err).Errorf("Failed to marshal updated domain spec: %+v", domainSpec)
		panic(err)
	}

	log.Log.Info("Successfully updated original domain spec with requested device bus id attributes")

	return newDomainXML, nil
}

/*func ensureFinalPathExists(path string, m mxj.Map) ([]byte, error) {
	pathSplitted := strings.Split(path, ".")
	var vp = pathSplitted[0]
	for _, p := range pathSplitted[1:] {
		if !m.Exists(vp) {
			err := m.SetValueForPath(map[string]interface{}{}, vp)
			if err != nil {
				return nil, err
			}
		}
		vp = vp + "." + p
	}
	return nil, nil
}*/

func main() {
	log.InitializeLogging("device-bus-id-hook-sidecar")

	var version string
	pflag.StringVar(&version, "version", "", "hook version to use")
	pflag.Parse()

	socketPath := hooks.HookSocketsSharedDirectory + "/smbios.sock"
	socket, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Log.Reason(err).Errorf("Failed to initialized socket on path: %s", socket)
		log.Log.Error("Check whether given directory exists and socket name is not already taken by other file")
		panic(err)
	}
	defer os.Remove(socketPath)

	server := grpc.NewServer([]grpc.ServerOption{}...)

	if version == "" {
		panic(fmt.Errorf("usage: \n        /example-hook-sidecar --version v1alpha1|v1alpha2"))
	}
	hooksInfo.RegisterInfoServer(server, infoServer{Version: version})
	hooksV1alpha1.RegisterCallbacksServer(server, v1alpha1Server{})
	hooksV1alpha2.RegisterCallbacksServer(server, v1alpha2Server{})
	log.Log.Infof("Starting hook server exposing 'info' and 'v1alpha1' services on socket %s", socketPath)
	server.Serve(socket)
}
