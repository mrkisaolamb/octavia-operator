/*

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

package amphoracontrollers

const (
	// Common consts for Management network

	// LbMgmtNetName -
	LbMgmtNetName = "lb-mgmt-net"

	// LbMgmtNetDescription -
	LbMgmtNetDescription = "LBaaS Management Network"

	// LbMgmtSubnetName -
	LbMgmtSubnetName = "lb-mgmt-subnet"

	// LbMgmtSubnetDescription -
	LbMgmtSubnetDescription = "LBaaS Management Subnet"

	// IPv4 consts

	// LbMgmtSubnetCIDR -
	LbMgmtSubnetCIDR = "172.24.0.0/16"

	// LbMgmtSubnetAllocationPoolStart -
	LbMgmtSubnetAllocationPoolStart = "172.24.0.2"

	// LbMgmtSubnetAllocationPoolEnd -
	LbMgmtSubnetAllocationPoolEnd = "172.24.255.254"

	// LbMgmtSubnetGatewayIP -
	LbMgmtSubnetGatewayIP = ""

	// IPv6 consts
	// using Unique local address (fc00::/7)
	// with Global ID 6c:6261:6173 ("lbaas")

	// LbMgmtSubnetIPv6CIDR -
	LbMgmtSubnetIPv6CIDR = "fd6c:6261:6173:0001::/64"

	// LbMgmtSubnetIPv6AllocationPoolStart -
	LbMgmtSubnetIPv6AllocationPoolStart = "fd6c:6261:6173:0001::2"

	// LbMgmtSubnetIPv6AllocationPoolEnd -
	LbMgmtSubnetIPv6AllocationPoolEnd = "fd6c:6261:6173:0001:ffff:ffff:ffff:ffff"

	// LbMgmtSubnetIPv6AddressMode -
	LbMgmtSubnetIPv6AddressMode = "slaac"

	// LbMgmtSubnetIPv6RAMode -
	LbMgmtSubnetIPv6RAMode = "slaac"

	// LbMgmtSubnetIPv6GatewayIP -
	LbMgmtSubnetIPv6GatewayIP = ""
)
