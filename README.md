# Latitude Go SDK

Everything you need to deploy and manage single-tenant, high-performance bare metal servers.

The Latitude.sh API gives you programmatic access to all resources available on the platform. With full control over the entire Latitude.sh platform, you can build integrations, custom dashboards, and manage your servers from your own clients.

<div align="left">
   <a href="https://github.com/speakeasy-sdks/latitude-go-sdk/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/latitude-go-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
   <a href="https://docs.latitude.sh/reference/summary"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/latitude-go-sdk
```
<!-- End SDK Installation -->
## Authentication

To create an API token, go to the Settings & Billing → API Keys page on the dashboard.

The key is shown to your only when it's created. If you lose it, you'll have to roll the key or create a new one.


## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/speakeasy-sdks/latitude-go-sdk"
    "github.com/speakeasy-sdks/latitude-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/latitude-go-sdk/pkg/models/operations"
)

func main() {
    s := latitude.New()
    
    req := operations.DeleteAPIKeyRequest{
        Security: operations.DeleteAPIKeySecurity{
            Bearer: shared.SchemeBearer{
                APIKey: "YOUR_API_KEY_HERE",
            },
        },
        PathParams: operations.DeleteAPIKeyPathParams{
            ID: "unde",
        },
    }
    
    res, err := s.APIKeys.DeleteAPIKey(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations


### APIKeys

* `DeleteAPIKey` - Delete an API Key
* `GetAPIKeys` - List all API Keys
* `PostAPIKey` - Create an API Key
* `UpdateAPIKey` - Regenerate an API Key

### APIVersion

* `GetCurrentVersion` - Get current API version
* `UpdateCurrentVersion` - Update current API version

### Account

* `GetUserProfile` - Retrieve the User Profile
* `GetUserTeams` - List all User Teams
* `PatchUserProfile` - Update the User Profile

### Bandwidth

* `GetTrafficConsumption` - Retrieve Traffic consumption
* `GetTrafficQuota` - Retrieve Traffic Quota

### BandwidthPackages

* `GetPlansBandwidth` - List all bandwidth packages available
* `UpdatePlansBandwidth` - Buy or remove bandwidth packages

### DeployConfig

* `GetServerDeployConfig` - Retrieve the Server Deploy Config
* `UpdateServerDeployConfig` - Update the Server Deploy Config

### IPAddresses

* `GetIP` - Retrieve an IP
* `GetIps` - List IPs

### IPMICredentials

* `CreateIpmiSession` - Generate IPMI credentials

### Members

* `DestroyTeamMember` - Remove a Team Member
* `GetTeamMembers` - List all Team Members
* `PostTeamMembers` - Add a Team Member

### OperatingSystems

* `GetPlansOperatingSystem` - List all operating systems available

### Plans

* `GetPlan` - Retrieve a Plan
* `GetPlans` - List all Plans

### PowerActions

* `CreateServerAction` - Run Server Action

### Projects

* `CreateProject` - Create a Project
* `DeleteProject` - Delete a Project
* `GetProject` - Retrieve a Project
* `GetProjects` - List all Projects
* `UpdateProject` - Update a Project

### Regions

* `GetRegion` - Retrieve a Region
* `GetRegions` - List all Regions

### RescueMode

* `ServerExitRescueMode` - Exits rescue mode for a Server
* `ServerStartRescueMode` - Puts a Server in rescue mode

### Roles

* `GetRoleID` - Retrieve a Role
* `GetRoles` - List all Roles

### SSHKeys

* `DeleteProjectSSHKey` - Delete a Project SSH Key
* `GetProjectSSHKey` - Retrieve a Project SSH Key
* `GetProjectSSHKeys` - List all Project SSH Keys
* `PostProjectSSHKey` - Create a Project SSH Key
* `PutProjectSSHKey` - Update a Project SSH Key

### ServerReinstall

* `CreateServerReinstall` - Run Server Reinstall

### Servers

* `CreateServer` - Deploy a new server
* `DestroyServer` - Remove a Server
* `GetServer` - Retrieve a Server
* `GetServers` - List all Servers
* `UpdateServer` - Update a Server

### Teams

* `GetTeam` - Retrieve the Current Team
* `PatchCurrentTeam` - Update a Team
* `PostTeam` - Create a Team

### UserData

* `DeleteProjectUserData` - Delete a Project User Data
* `GetProjectUserData` - Retrieve a Project User Data
* `GetProjectUsersData` - List all Project User Data
* `PostProjectUserData` - Create a Project User Data
* `PutProjectUserData` - Update a Project User Data

### VPNSessions

* `DeleteVpnSession` - Delete a VPN Session
* `GetVpnSessions` - List all Active VPN Sessions
* `PostVpnSession` - Create a VPN Session
* `PutVpnSession` - Refresh a VPN Session

### VirtualNetworkAssignments

* `AssignServerVirtualNetwork` - Assign a server to a virtual network
* `DeleteVirtualNetworksAssignments` - Delete a Virtual Network Assignment
* `GetVirtualNetworksAssignments` - List all servers assigned to virtual networks

### VirtualNetworks

* `CreateVirtualNetwork` - Create a Virtual Network
* `DestroyVirtualNetwork` - Delete a Virtual Network
* `GetVirtualNetwork` - Retrieve a Virtual Network
* `GetVirtualNetworks` - List all Virtual Networks
* `UpdateVirtualNetwork` - Update a Virtual Network
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
