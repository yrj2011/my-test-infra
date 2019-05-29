package azure

// Copyright 2017 Microsoft Corporation
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// EnvironmentFilepathName captures the name of the environment variable containing the path to the file
// to be used while populating the Azure Environment.
const EnvironmentFilepathName = "AZURE_ENVIRONMENT_FILEPATH"

var environments = map[string]Environment{
	"AZURECHINACLOUD":        ChinaCloud,
	"AZUREGERMANCLOUD":       GermanCloud,
	"AZUREPUBLICCLOUD":       PublicCloud,
	"AZUREUSGOVERNMENTCLOUD": USGovernmentCloud,
}

// Environment represents a set of endpoints for each of Azure's Clouds.
type Environment struct {
	Name                         string `json:"name"`
	ManagementPortalURL          string `json:"managementPortalURL"`
	PublishSettingsURL           string `json:"publishSettingsURL"`
	ServiceManagementEndpoint    string `json:"serviceManagementEndpoint"`
	ResourceManagerEndpoint      string `json:"resourceManagerEndpoint"`
	ActiveDirectoryEndpoint      string `json:"activeDirectoryEndpoint"`
	GalleryEndpoint              string `json:"galleryEndpoint"`
	KeyVaultEndpoint             string `json:"keyVaultEndpoint"`
	GraphEndpoint                string `json:"graphEndpoint"`
	ServiceBusEndpoint           string `json:"serviceBusEndpoint"`
	BatchManagementEndpoint      string `json:"batchManagementEndpoint"`
	StorageEndpointSuffix        string `json:"storageEndpointSuffix"`
	SQLDatabaseDNSSuffix         string `json:"sqlDatabaseDNSSuffix"`
	TrafficManagerDNSSuffix      string `json:"trafficManagerDNSSuffix"`
	KeyVaultDNSSuffix            string `json:"keyVaultDNSSuffix"`
	ServiceBusEndpointSuffix     string `json:"serviceBusEndpointSuffix"`
	ServiceManagementVMDNSSuffix string `json:"serviceManagementVMDNSSuffix"`
	ResourceManagerVMDNSSuffix   string `json:"resourceManagerVMDNSSuffix"`
	ContainerRegistryDNSSuffix   string `json:"containerRegistryDNSSuffix"`
	TokenAudience                string `json:"tokenAudience"`
}

var (
	// PublicCloud is the default public Azure cloud environment
	PublicCloud = Environment{
		Name:                         "AzurePublicCloud",
		ManagementPortalURL:          "http://manage.windowsazure.com/",
		PublishSettingsURL:           "http://manage.windowsazure.com/publishsettings/index",
		ServiceManagementEndpoint:    "http://management.core.windows.net/",
		ResourceManagerEndpoint:      "http://management.azure.com/",
		ActiveDirectoryEndpoint:      "http://login.microsoftonline.com/",
		GalleryEndpoint:              "http://gallery.azure.com/",
		KeyVaultEndpoint:             "http://vault.azure.net/",
		GraphEndpoint:                "http://graph.windows.net/",
		ServiceBusEndpoint:           "http://servicebus.windows.net/",
		BatchManagementEndpoint:      "http://batch.core.windows.net/",
		StorageEndpointSuffix:        "core.windows.net",
		SQLDatabaseDNSSuffix:         "database.windows.net",
		TrafficManagerDNSSuffix:      "trafficmanager.net",
		KeyVaultDNSSuffix:            "vault.azure.net",
		ServiceBusEndpointSuffix:     "servicebus.windows.net",
		ServiceManagementVMDNSSuffix: "cloudapp.net",
		ResourceManagerVMDNSSuffix:   "cloudapp.azure.com",
		ContainerRegistryDNSSuffix:   "azurecr.io",
		TokenAudience:                "http://management.azure.com/",
	}

	// USGovernmentCloud is the cloud environment for the US Government
	USGovernmentCloud = Environment{
		Name:                         "AzureUSGovernmentCloud",
		ManagementPortalURL:          "http://manage.windowsazure.us/",
		PublishSettingsURL:           "http://manage.windowsazure.us/publishsettings/index",
		ServiceManagementEndpoint:    "http://management.core.usgovcloudapi.net/",
		ResourceManagerEndpoint:      "http://management.usgovcloudapi.net/",
		ActiveDirectoryEndpoint:      "http://login.microsoftonline.us/",
		GalleryEndpoint:              "http://gallery.usgovcloudapi.net/",
		KeyVaultEndpoint:             "http://vault.usgovcloudapi.net/",
		GraphEndpoint:                "http://graph.windows.net/",
		ServiceBusEndpoint:           "http://servicebus.usgovcloudapi.net/",
		BatchManagementEndpoint:      "http://batch.core.usgovcloudapi.net/",
		StorageEndpointSuffix:        "core.usgovcloudapi.net",
		SQLDatabaseDNSSuffix:         "database.usgovcloudapi.net",
		TrafficManagerDNSSuffix:      "usgovtrafficmanager.net",
		KeyVaultDNSSuffix:            "vault.usgovcloudapi.net",
		ServiceBusEndpointSuffix:     "servicebus.usgovcloudapi.net",
		ServiceManagementVMDNSSuffix: "usgovcloudapp.net",
		ResourceManagerVMDNSSuffix:   "cloudapp.windowsazure.us",
		ContainerRegistryDNSSuffix:   "azurecr.io",
		TokenAudience:                "http://management.usgovcloudapi.net/",
	}

	// ChinaCloud is the cloud environment operated in China
	ChinaCloud = Environment{
		Name:                         "AzureChinaCloud",
		ManagementPortalURL:          "http://manage.chinacloudapi.com/",
		PublishSettingsURL:           "http://manage.chinacloudapi.com/publishsettings/index",
		ServiceManagementEndpoint:    "http://management.core.chinacloudapi.cn/",
		ResourceManagerEndpoint:      "http://management.chinacloudapi.cn/",
		ActiveDirectoryEndpoint:      "http://login.chinacloudapi.cn/",
		GalleryEndpoint:              "http://gallery.chinacloudapi.cn/",
		KeyVaultEndpoint:             "http://vault.azure.cn/",
		GraphEndpoint:                "http://graph.chinacloudapi.cn/",
		ServiceBusEndpoint:           "http://servicebus.chinacloudapi.cn/",
		BatchManagementEndpoint:      "http://batch.chinacloudapi.cn/",
		StorageEndpointSuffix:        "core.chinacloudapi.cn",
		SQLDatabaseDNSSuffix:         "database.chinacloudapi.cn",
		TrafficManagerDNSSuffix:      "trafficmanager.cn",
		KeyVaultDNSSuffix:            "vault.azure.cn",
		ServiceBusEndpointSuffix:     "servicebus.chinacloudapi.cn",
		ServiceManagementVMDNSSuffix: "chinacloudapp.cn",
		ResourceManagerVMDNSSuffix:   "cloudapp.azure.cn",
		ContainerRegistryDNSSuffix:   "azurecr.io",
		TokenAudience:                "http://management.chinacloudapi.cn/",
	}

	// GermanCloud is the cloud environment operated in Germany
	GermanCloud = Environment{
		Name:                         "AzureGermanCloud",
		ManagementPortalURL:          "http://portal.microsoftazure.de/",
		PublishSettingsURL:           "http://manage.microsoftazure.de/publishsettings/index",
		ServiceManagementEndpoint:    "http://management.core.cloudapi.de/",
		ResourceManagerEndpoint:      "http://management.microsoftazure.de/",
		ActiveDirectoryEndpoint:      "http://login.microsoftonline.de/",
		GalleryEndpoint:              "http://gallery.cloudapi.de/",
		KeyVaultEndpoint:             "http://vault.microsoftazure.de/",
		GraphEndpoint:                "http://graph.cloudapi.de/",
		ServiceBusEndpoint:           "http://servicebus.cloudapi.de/",
		BatchManagementEndpoint:      "http://batch.cloudapi.de/",
		StorageEndpointSuffix:        "core.cloudapi.de",
		SQLDatabaseDNSSuffix:         "database.cloudapi.de",
		TrafficManagerDNSSuffix:      "azuretrafficmanager.de",
		KeyVaultDNSSuffix:            "vault.microsoftazure.de",
		ServiceBusEndpointSuffix:     "servicebus.cloudapi.de",
		ServiceManagementVMDNSSuffix: "azurecloudapp.de",
		ResourceManagerVMDNSSuffix:   "cloudapp.microsoftazure.de",
		ContainerRegistryDNSSuffix:   "azurecr.io",
		TokenAudience:                "http://management.microsoftazure.de/",
	}
)

// EnvironmentFromName returns an Environment based on the common name specified.
func EnvironmentFromName(name string) (Environment, error) {
	// IMPORTANT
	// As per @radhikagupta5:
	// This is technical debt, fundamentally here because Kubernetes is not currently accepting
	// contributions to the providers. Once that is an option, the provider should be updated to
	// directly call `EnvironmentFromFile`. Until then, we rely on dispatching Azure Stack environment creation
	// from this method based on the name that is provided to us.
	if strings.EqualFold(name, "AZURESTACKCLOUD") {
		return EnvironmentFromFile(os.Getenv(EnvironmentFilepathName))
	}

	name = strings.ToUpper(name)
	env, ok := environments[name]
	if !ok {
		return env, fmt.Errorf("autorest/azure: There is no cloud environment matching the name %q", name)
	}

	return env, nil
}

// EnvironmentFromFile loads an Environment from a configuration file available on disk.
// This function is particularly useful in the Hybrid Cloud model, where one must define their own
// endpoints.
func EnvironmentFromFile(location string) (unmarshaled Environment, err error) {
	fileContents, err := ioutil.ReadFile(location)
	if err != nil {
		return
	}

	err = json.Unmarshal(fileContents, &unmarshaled)

	return
}
