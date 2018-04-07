// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openstackbotanist

import "github.com/gardener/gardener/pkg/operation"

// OpenStackBotanist is a struct which has methods that perform OpenStack cloud-specific operations for a Shoot cluster.
type OpenStackBotanist struct {
	*operation.Operation
	CloudProviderName string
}

const (
	// DomainName is a constant for the key in a cloud provider secret and backup secret that holds the OpenStack domain name.
	DomainName = "domainName"
	// TenantName is a constant for the key in a cloud provider secret and backup secret that holds the OpenStack tenant name.
	TenantName = "tenantName"
	// UserName is a constant for the key in a cloud provider secret and backup secret that holds the OpenStack username.
	UserName = "username"
	// Password is a constant for the key in a cloud provider secret and backup secret that holds the OpenStack password.
	Password = "password"
	// AuthURL is a constant for the key in a backup secret that holds the OpenStack authentication URL.
	AuthURL = "authURL"
)
