# Golang SDK for Rudder.io

<!-- Start SDK Installation [installation] -->
## Installation

```bash
go get -u github.com/infra-rdc/rudder-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## Example Usage

### Example

```go
package main

import (
	"context"
	ruddergo "github.com/infra-rdc/rudder-go"
	"log"
)

func main() {
	s := ruddergo.New(
		ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.ChangeRequests.ListChangeRequests(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [ChangeRequests](docs/sdks/changerequests/README.md)

* [ListChangeRequests](docs/sdks/changerequests/README.md#listchangerequests) - List all change requests
* [ChangeRequestDetails](docs/sdks/changerequests/README.md#changerequestdetails) - Get a change request details
* [DeclineChangeRequest](docs/sdks/changerequests/README.md#declinechangerequest) - Decline a request details
* [UpdateChangeRequest](docs/sdks/changerequests/README.md#updatechangerequest) - Update a request details
* [AcceptChangeRequest](docs/sdks/changerequests/README.md#acceptchangerequest) - Accept a request details
* [ListUsers](docs/sdks/changerequests/README.md#listusers) - List user
* [SaveWorkflowUser](docs/sdks/changerequests/README.md#saveworkflowuser) - Update validated user list
* [RemoveValidatedUser](docs/sdks/changerequests/README.md#removevalidateduser) - Remove an user from validated user list

### [Archives](docs/sdks/archives/README.md)

* [Export](docs/sdks/archives/README.md#export) - Get a ZIP archive of the requested items and their dependencies
* [Import](docs/sdks/archives/README.md#import) - Import a ZIP archive of policies into Rudder

### [Campaigns](docs/sdks/campaigns/README.md)

* [GetAllCampaignEvents](docs/sdks/campaigns/README.md#getallcampaignevents) - Get all campaign events
* [GetCampaignEvent](docs/sdks/campaigns/README.md#getcampaignevent) - Get a campaign event details
* [SaveCampaignEvent](docs/sdks/campaigns/README.md#savecampaignevent) - Update an existing event
* [DeleteCampaignEvent](docs/sdks/campaigns/README.md#deletecampaignevent) - Delete a campaign event details
* [GetEventsCampaign](docs/sdks/campaigns/README.md#geteventscampaign) - Get campaign events for a campaign
* [ScheduleCampaign](docs/sdks/campaigns/README.md#schedulecampaign) - Schedule a campaign event for a campaign

### [Compliance](docs/sdks/compliance/README.md)

* [GetGlobalCompliance](docs/sdks/compliance/README.md#getglobalcompliance) - Global compliance
* [GetDirectivesCompliance](docs/sdks/compliance/README.md#getdirectivescompliance) - Compliance details for all directives
* [GetDirectiveComplianceID](docs/sdks/compliance/README.md#getdirectivecomplianceid) - Compliance details by directive
* [GetNodesCompliance](docs/sdks/compliance/README.md#getnodescompliance) - Compliance details for all nodes
* [GetNodeCompliance](docs/sdks/compliance/README.md#getnodecompliance) - Compliance details by node
* [GetRulesCompliance](docs/sdks/compliance/README.md#getrulescompliance) - Compliance details for all rules
* [GetRuleCompliance](docs/sdks/compliance/README.md#getrulecompliance) - Compliance details by rule

### [Cve](docs/sdks/cve/README.md)

* [GetAllCve](docs/sdks/cve/README.md#getallcve) - Get all CVE details
* [CheckCVE](docs/sdks/cve/README.md#checkcve) - Trigger a CVE check
* [GetCVECheckConfiguration](docs/sdks/cve/README.md#getcvecheckconfiguration) - Get CVE check config
* [UpdateCVECheckConfiguration](docs/sdks/cve/README.md#updatecvecheckconfiguration) - Update cve check config
* [GetLastCVECheck](docs/sdks/cve/README.md#getlastcvecheck) - Get last CVE check result
* [GetCVEList](docs/sdks/cve/README.md#getcvelist) - Get a list of CVE details
* [UpdateCVE](docs/sdks/cve/README.md#updatecve) - Update CVE database from remote source
* [ReadCVEfromFS](docs/sdks/cve/README.md#readcvefromfs) - Update CVE database from file system
* [GetCve](docs/sdks/cve/README.md#getcve) - Get a CVE details

### [DataSources](docs/sdks/datasources/README.md)

* [GetAllDataSources](docs/sdks/datasources/README.md#getalldatasources) - List all data sources
* [ReloadAllDatasourcesAllNodes](docs/sdks/datasources/README.md#reloadalldatasourcesallnodes) - Update properties from data sources
* [ReloadOneDatasourceAllNodes](docs/sdks/datasources/README.md#reloadonedatasourceallnodes) - Update properties from data sources
* [GetDataSource](docs/sdks/datasources/README.md#getdatasource) - Get data source configuration
* [DeleteDataSource](docs/sdks/datasources/README.md#deletedatasource) - Delete a data source
* [ReloadAllDatasourcesOneNode](docs/sdks/datasources/README.md#reloadalldatasourcesonenode) - Update properties for one node from all data sources
* [ReloadOneDatasourceOneNode](docs/sdks/datasources/README.md#reloadonedatasourceonenode) - Update properties for one node from a data source

### [Directives](docs/sdks/directives/README.md)

* [ListDirectives](docs/sdks/directives/README.md#listdirectives) - List all directives
* [CreateDirective](docs/sdks/directives/README.md#createdirective) - Create a directive
* [DirectiveDetails](docs/sdks/directives/README.md#directivedetails) - Get directive details
* [DeleteDirective](docs/sdks/directives/README.md#deletedirective) - Delete a directive

### [Groups](docs/sdks/groups/README.md)

* [ListGroups](docs/sdks/groups/README.md#listgroups) - List all groups
* [CreateGroup](docs/sdks/groups/README.md#creategroup) - Create a group
* [GetGroupCategoryDetails](docs/sdks/groups/README.md#getgroupcategorydetails) - Get group category details
* [DeleteGroupCategory](docs/sdks/groups/README.md#deletegroupcategory) - Delete group category
* [UpdateGroupCategory](docs/sdks/groups/README.md#updategroupcategory) - Update group category details
* [GetGroupTree](docs/sdks/groups/README.md#getgrouptree) - Get groups tree
* [GroupDetails](docs/sdks/groups/README.md#groupdetails) - Get group details
* [UpdateGroup](docs/sdks/groups/README.md#updategroup) - Update group details
* [DeleteGroup](docs/sdks/groups/README.md#deletegroup) - Delete a group
* [ReloadGroup](docs/sdks/groups/README.md#reloadgroup) - Reload a group

### [APIInfo](docs/sdks/apiinfo/README.md)

* [APIGeneralInformations](docs/sdks/apiinfo/README.md#apigeneralinformations) - List all endpoints
* [APIInformations](docs/sdks/apiinfo/README.md#apiinformations) - Get information about one API endpoint
* [APISubInformations](docs/sdks/apiinfo/README.md#apisubinformations) - Get information on endpoint in a section

### [Inventories](docs/sdks/inventories/README.md)

* [QueueInformation](docs/sdks/inventories/README.md#queueinformation) - Get information about inventory processing queue
* [UploadInventory](docs/sdks/inventories/README.md#uploadinventory) - Upload an inventory for processing
* [FileWatcherRestart](docs/sdks/inventories/README.md#filewatcherrestart) - Restart inventory watcher
* [FileWatcherStart](docs/sdks/inventories/README.md#filewatcherstart) - Start inventory watcher
* [FileWatcherStop](docs/sdks/inventories/README.md#filewatcherstop) - Stop inventory watcher

### [Techniques](docs/sdks/techniques/README.md)

* [Methods](docs/sdks/techniques/README.md#methods) - List methods
* [ReloadMethods](docs/sdks/techniques/README.md#reloadmethods) - Reload methods
* [ListTechniques](docs/sdks/techniques/README.md#listtechniques) - List all techniques
* [Techniques](docs/sdks/techniques/README.md#techniques) - Reload techniques
* [ListTechniquesVersions](docs/sdks/techniques/README.md#listtechniquesversions) - List versions
* [GetTechniqueAllVersion](docs/sdks/techniques/README.md#gettechniqueallversion) - Technique metadata by ID
* [ListTechniquesDirectives](docs/sdks/techniques/README.md#listtechniquesdirectives) - List all directives based on a technique
* [DeleteTechnique](docs/sdks/techniques/README.md#deletetechnique) - Delete technique
* [GetTechniqueAllVersionID](docs/sdks/techniques/README.md#gettechniqueallversionid) - Technique metadata by version and ID
* [ListTechniqueVersionDirectives](docs/sdks/techniques/README.md#listtechniqueversiondirectives) - List all directives based on a version of a technique
* [GetTechniquesResources](docs/sdks/techniques/README.md#gettechniquesresources) - Technique's resources
* [TechniqueRevisions](docs/sdks/techniques/README.md#techniquerevisions) - Technique's revisions

### [Nodes](docs/sdks/nodes/README.md)

* [ListAcceptedNodes](docs/sdks/nodes/README.md#listacceptednodes) - List managed nodes
* [CreateNodes](docs/sdks/nodes/README.md#createnodes) - Create one or several new nodes
* [ApplyPolicyAllNodes](docs/sdks/nodes/README.md#applypolicyallnodes) - Trigger an agent run on all nodes
* [ListPendingNodes](docs/sdks/nodes/README.md#listpendingnodes) - List pending nodes
* [ChangePendingNodeStatus](docs/sdks/nodes/README.md#changependingnodestatus) - Update pending Node status
* [GetNodesStatus](docs/sdks/nodes/README.md#getnodesstatus) - Get nodes acceptation status
* [NodeDetails](docs/sdks/nodes/README.md#nodedetails) - Get information about a node
* [UpdateNode](docs/sdks/nodes/README.md#updatenode) - Update node settings and properties
* [DeleteNode](docs/sdks/nodes/README.md#deletenode) - Delete a node
* [ApplyPolicy](docs/sdks/nodes/README.md#applypolicy) - Trigger an agent run
* [NodeInheritedProperties](docs/sdks/nodes/README.md#nodeinheritedproperties) - Get inherited node properties for a node

### [OpenSCAP](docs/sdks/openscap/README.md)

* [OpenscapReport](docs/sdks/openscap/README.md#openscapreport) - Get an OpenSCAP report

### [Parameters](docs/sdks/parameters/README.md)

* [ListParameters](docs/sdks/parameters/README.md#listparameters) - List all global properties
* [ParameterDetails](docs/sdks/parameters/README.md#parameterdetails) - Get the value of a global property
* [UpdateParameter](docs/sdks/parameters/README.md#updateparameter) - Update a global property's value
* [DeleteParameter](docs/sdks/parameters/README.md#deleteparameter) - Delete a global parameter

### [Plugins](docs/sdks/plugins/README.md)

* [PluginSettings](docs/sdks/plugins/README.md#pluginsettings) - Get plugins repository settings
* [UpdateSettings](docs/sdks/plugins/README.md#updatesettings) - Update plugins settings

### [Rules](docs/sdks/rules/README.md)

* [GetRuleCategoryDetails](docs/sdks/rules/README.md#getrulecategorydetails) - Get rule category details
* [DeleteRuleCategory](docs/sdks/rules/README.md#deleterulecategory) - Delete group category
* [UpdateRuleCategory](docs/sdks/rules/README.md#updaterulecategory) - Update rule category details
* [GetRuleTree](docs/sdks/rules/README.md#getruletree) - Get rules tree

### [ScaleOutRelay](docs/sdks/scaleoutrelay/README.md)

* [DemoteToNode](docs/sdks/scaleoutrelay/README.md#demotetonode) - Demote a relay to simple node
* [PromoteToRelay](docs/sdks/scaleoutrelay/README.md#promotetorelay) - Promote a node to relay

### [SecretManagement](docs/sdks/secretmanagement/README.md)

* [GetAllSecrets](docs/sdks/secretmanagement/README.md#getallsecrets) - List all secrets
* [UpdateSecret](docs/sdks/secretmanagement/README.md#updatesecret) - Update a secret
* [AddSecret](docs/sdks/secretmanagement/README.md#addsecret) - Create a secret
* [GetSecret](docs/sdks/secretmanagement/README.md#getsecret) - Get one secret
* [DeleteSecret](docs/sdks/secretmanagement/README.md#deletesecret) - Delete a secret

### [Settings](docs/sdks/settings/README.md)

* [GetAllSettings](docs/sdks/settings/README.md#getallsettings) - List all settings
* [GetAllowedNetworks](docs/sdks/settings/README.md#getallowednetworks) - Get allowed networks for a policy server
* [SetAllowedNetworks](docs/sdks/settings/README.md#setallowednetworks) - Set allowed networks for a policy server
* [ModifyAllowedNetworks](docs/sdks/settings/README.md#modifyallowednetworks) - Modify allowed networks for a policy server
* [GetSetting](docs/sdks/settings/README.md#getsetting) - Get the value of a setting
* [ModifySetting](docs/sdks/settings/README.md#modifysetting) - Set the value of a setting

### [Status](docs/sdks/status/README.md)

* [None](docs/sdks/status/README.md#none) - Check if Rudder is alive

### [System](docs/sdks/system/README.md)

* [ListArchives](docs/sdks/system/README.md#listarchives) - List archives
* [CreateArchive](docs/sdks/system/README.md#createarchive) - Create an archive
* [RestoreArchive](docs/sdks/system/README.md#restorearchive) - Restore an archive
* [GetZipArchive](docs/sdks/system/README.md#getziparchive) - Get an archive as a ZIP
* [GetHealthcheckResult](docs/sdks/system/README.md#gethealthcheckresult) - Get healthcheck
* [GetSystemInfo](docs/sdks/system/README.md#getsysteminfo) - Get server information
* [PurgeSoftware](docs/sdks/system/README.md#purgesoftware) - Trigger batch for cleaning unreferenced software
* [RegeneratePolicies](docs/sdks/system/README.md#regeneratepolicies) - Trigger a new policy generation
* [ReloadAll](docs/sdks/system/README.md#reloadall) - Reload both techniques and dynamic groups
* [ReloadGroups](docs/sdks/system/README.md#reloadgroups) - Reload dynamic groups
* [ReloadTechniques](docs/sdks/system/README.md#reloadtechniques) - Reload techniques
* [GetStatus](docs/sdks/system/README.md#getstatus) - Get server status
* [UpdatePolicies](docs/sdks/system/README.md#updatepolicies) - Trigger update of policies

### [SystemUpdateCampaigns](docs/sdks/systemupdatecampaigns/README.md)

* [GetCampaignResults](docs/sdks/systemupdatecampaigns/README.md#getcampaignresults) - Get a campaign result history
* [GetCampaignEventResult](docs/sdks/systemupdatecampaigns/README.md#getcampaigneventresult) - Get a campaign event result
* [GetSystemUpdateResultForNode](docs/sdks/systemupdatecampaigns/README.md#getsystemupdateresultfornode) - Get detailed campaign event result for a Node

### [UserManagement](docs/sdks/usermanagement/README.md)

* [AddUser](docs/sdks/usermanagement/README.md#adduser) - Add user
* [GetRole](docs/sdks/usermanagement/README.md#getrole) - List all roles
* [UpdateUser](docs/sdks/usermanagement/README.md#updateuser) - Update user's infos
* [GetUserInfo](docs/sdks/usermanagement/README.md#getuserinfo) - List all users
* [ReloadUserConf](docs/sdks/usermanagement/README.md#reloaduserconf) - Reload user
* [DeleteUser](docs/sdks/usermanagement/README.md#deleteuser) - Delete an user
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                          | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.DeleteParameterResponseBody | 500                                   | application/json                      |
| sdkerrors.SDKError                    | 4xx-5xx                               | */*                                   |

### Example

```go
package main

import (
	"context"
	"errors"
	ruddergo "github.com/infra-rdc/rudder-go"
	"github.com/infra-rdc/rudder-go/models/sdkerrors"
	"log"
)

func main() {
	s := ruddergo.New(
		ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)
	var parameterID string = "rudder_file_edit_header"
	ctx := context.Background()
	res, err := s.Parameters.DeleteParameter(ctx, parameterID)
	if err != nil {

		var e *sdkerrors.DeleteParameterResponseBody
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://rudder.example.local/rudder/api/latest` | None |

#### Example

```go
package main

import (
	"context"
	ruddergo "github.com/infra-rdc/rudder-go"
	"log"
)

func main() {
	s := ruddergo.New(
		ruddergo.WithServerIndex(0),
		ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.ChangeRequests.ListChangeRequests(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	ruddergo "github.com/infra-rdc/rudder-go"
	"log"
)

func main() {
	s := ruddergo.New(
		ruddergo.WithServerURL("https://rudder.example.local/rudder/api/latest"),
		ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.ChangeRequests.ListChangeRequests(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name        | Type        | Scheme      |
| ----------- | ----------- | ----------- |
| `APITokens` | apiKey      | API key     |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	ruddergo "github.com/infra-rdc/rudder-go"
	"log"
)

func main() {
	s := ruddergo.New(
		ruddergo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.ChangeRequests.ListChangeRequests(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!
