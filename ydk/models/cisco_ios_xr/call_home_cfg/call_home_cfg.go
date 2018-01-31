// This module contains a collection of YANG definitions
// for Cisco IOS-XR call-home package configuration.
// 
// This module contains definitions
// for the following management objects:
//   call-home: Set CallHome parameters
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package call_home_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package call_home_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-call-home-cfg call-home}", reflect.TypeOf(CallHome{}))
    ydk.RegisterEntity("Cisco-IOS-XR-call-home-cfg:call-home", reflect.TypeOf(CallHome{}))
}

// CallHomeDayOfWeek represents Call home day of week
type CallHomeDayOfWeek string

const (
    // Sunday
    CallHomeDayOfWeek_sunday CallHomeDayOfWeek = "sunday"

    // Monday
    CallHomeDayOfWeek_monday CallHomeDayOfWeek = "monday"

    // Tuesday
    CallHomeDayOfWeek_tuesday CallHomeDayOfWeek = "tuesday"

    // Wednesday
    CallHomeDayOfWeek_wednesday CallHomeDayOfWeek = "wednesday"

    // Thursday
    CallHomeDayOfWeek_thursday CallHomeDayOfWeek = "thursday"

    // Friday
    CallHomeDayOfWeek_friday CallHomeDayOfWeek = "friday"

    // Saturday
    CallHomeDayOfWeek_saturday CallHomeDayOfWeek = "saturday"
)

// DataPrivacyLevel represents Data privacy level
type DataPrivacyLevel string

const (
    // Normal
    DataPrivacyLevel_normal DataPrivacyLevel = "normal"

    // High
    DataPrivacyLevel_high DataPrivacyLevel = "high"

    // HostName
    DataPrivacyLevel_host_name DataPrivacyLevel = "host-name"
)

// CallHomeMailSendInterval represents Call home mail send interval
type CallHomeMailSendInterval string

const (
    // Daily call-home message
    CallHomeMailSendInterval_daily CallHomeMailSendInterval = "daily"

    // Weekly call-home message
    CallHomeMailSendInterval_weekly CallHomeMailSendInterval = "weekly"

    // Monthly call-home message
    CallHomeMailSendInterval_monthly CallHomeMailSendInterval = "monthly"
)

// CallHomeEventSeverity represents Call home event severity
type CallHomeEventSeverity string

const (
    // Debugging event
    CallHomeEventSeverity_debugging CallHomeEventSeverity = "debugging"

    // Normal event
    CallHomeEventSeverity_normal CallHomeEventSeverity = "normal"

    // Notification event
    CallHomeEventSeverity_notification CallHomeEventSeverity = "notification"

    // Warning event
    CallHomeEventSeverity_warning CallHomeEventSeverity = "warning"

    // Minor event
    CallHomeEventSeverity_minor CallHomeEventSeverity = "minor"

    // Major event
    CallHomeEventSeverity_major CallHomeEventSeverity = "major"

    // Critical event
    CallHomeEventSeverity_critical CallHomeEventSeverity = "critical"

    // Fatal event
    CallHomeEventSeverity_fatal CallHomeEventSeverity = "fatal"

    // Disaster event
    CallHomeEventSeverity_disaster CallHomeEventSeverity = "disaster"

    // Catastrophic event
    CallHomeEventSeverity_catastrophic CallHomeEventSeverity = "catastrophic"
)

// SnapshotInterval represents Snapshot interval
type SnapshotInterval string

const (
    // Daily call-home message
    SnapshotInterval_daily SnapshotInterval = "daily"

    // Weekly call-home message
    SnapshotInterval_weekly SnapshotInterval = "weekly"

    // Monthly call-home message
    SnapshotInterval_monthly SnapshotInterval = "monthly"
)

// CallHomeTransMethod represents Call home trans method
type CallHomeTransMethod string

const (
    // To add email address to lthis profile
    CallHomeTransMethod_email CallHomeTransMethod = "email"

    // To add destination address(1-200) characters
    CallHomeTransMethod_http CallHomeTransMethod = "http"
)

// CallHome
// Set CallHome parameters
type CallHome struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Customer identification for Cisco Smart Call Home. The type is string with
    // length: 1..64.
    CustomerId interface{}

    // Phone number of the contact person. The type is string with length: 1..17.
    PhoneNumber interface{}

    // System Contact is Smart Licensing. The type is bool.
    ContactSmartLicensing interface{}

    // Contact person's email address. The type is string with length: 1..194.
    ContactEmailAddress interface{}

    // Call-home event trigger rate-limit threshold per minute. The type is
    // interface{} with range: 1..5.
    RateLimit interface{}

    // Site identification for Cisco Smart Call Home. The type is string with
    // length: 1..200.
    SiteId interface{}

    // Vrf routing/forwarding instance name. The type is string with length:
    // 1..32.
    Vrf interface{}

    // Street address, city, state, and zip code. The type is string with length:
    // 1..200.
    StreetAddress interface{}

    // Source interface name to send call-home messages. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Contract identification for Cisco Smart Call Home. The type is string with
    // length: 1..64.
    ContractId interface{}

    // Call home msg's reply-to email address. The type is string.
    ReplyTo interface{}

    // Call home msg's from email address. The type is string.
    From interface{}

    // Enable call-home service. The type is interface{}.
    Active interface{}

    // List of call-home mail_server.
    MailServers CallHome_MailServers

    // Enable or disable call-home syslog message throttling.
    SyslogThrottling CallHome_SyslogThrottling

    // Enable/disable licensing messages. By default is enabled.
    SmartLicensing CallHome_SmartLicensing

    // http proxy server address and port.
    HttpProxy CallHome_HttpProxy

    // List of profiles.
    Profiles CallHome_Profiles

    // List of alert-group.
    AlertGroups CallHome_AlertGroups

    // Set call-home data-privacy.
    DataPrivacies CallHome_DataPrivacies

    // alert-group config.
    AlertGroupConfig CallHome_AlertGroupConfig

    // Config aaa authorization, default username is callhome.
    Authorization CallHome_Authorization
}

func (callHome *CallHome) GetFilter() yfilter.YFilter { return callHome.YFilter }

func (callHome *CallHome) SetFilter(yf yfilter.YFilter) { callHome.YFilter = yf }

func (callHome *CallHome) GetGoName(yname string) string {
    if yname == "customer-id" { return "CustomerId" }
    if yname == "phone-number" { return "PhoneNumber" }
    if yname == "contact-smart-licensing" { return "ContactSmartLicensing" }
    if yname == "contact-email-address" { return "ContactEmailAddress" }
    if yname == "rate-limit" { return "RateLimit" }
    if yname == "site-id" { return "SiteId" }
    if yname == "vrf" { return "Vrf" }
    if yname == "street-address" { return "StreetAddress" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "contract-id" { return "ContractId" }
    if yname == "reply-to" { return "ReplyTo" }
    if yname == "from" { return "From" }
    if yname == "active" { return "Active" }
    if yname == "mail-servers" { return "MailServers" }
    if yname == "syslog-throttling" { return "SyslogThrottling" }
    if yname == "smart-licensing" { return "SmartLicensing" }
    if yname == "http-proxy" { return "HttpProxy" }
    if yname == "profiles" { return "Profiles" }
    if yname == "alert-groups" { return "AlertGroups" }
    if yname == "data-privacies" { return "DataPrivacies" }
    if yname == "alert-group-config" { return "AlertGroupConfig" }
    if yname == "authorization" { return "Authorization" }
    return ""
}

func (callHome *CallHome) GetSegmentPath() string {
    return "Cisco-IOS-XR-call-home-cfg:call-home"
}

func (callHome *CallHome) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mail-servers" {
        return &callHome.MailServers
    }
    if childYangName == "syslog-throttling" {
        return &callHome.SyslogThrottling
    }
    if childYangName == "smart-licensing" {
        return &callHome.SmartLicensing
    }
    if childYangName == "http-proxy" {
        return &callHome.HttpProxy
    }
    if childYangName == "profiles" {
        return &callHome.Profiles
    }
    if childYangName == "alert-groups" {
        return &callHome.AlertGroups
    }
    if childYangName == "data-privacies" {
        return &callHome.DataPrivacies
    }
    if childYangName == "alert-group-config" {
        return &callHome.AlertGroupConfig
    }
    if childYangName == "authorization" {
        return &callHome.Authorization
    }
    return nil
}

func (callHome *CallHome) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mail-servers"] = &callHome.MailServers
    children["syslog-throttling"] = &callHome.SyslogThrottling
    children["smart-licensing"] = &callHome.SmartLicensing
    children["http-proxy"] = &callHome.HttpProxy
    children["profiles"] = &callHome.Profiles
    children["alert-groups"] = &callHome.AlertGroups
    children["data-privacies"] = &callHome.DataPrivacies
    children["alert-group-config"] = &callHome.AlertGroupConfig
    children["authorization"] = &callHome.Authorization
    return children
}

func (callHome *CallHome) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["customer-id"] = callHome.CustomerId
    leafs["phone-number"] = callHome.PhoneNumber
    leafs["contact-smart-licensing"] = callHome.ContactSmartLicensing
    leafs["contact-email-address"] = callHome.ContactEmailAddress
    leafs["rate-limit"] = callHome.RateLimit
    leafs["site-id"] = callHome.SiteId
    leafs["vrf"] = callHome.Vrf
    leafs["street-address"] = callHome.StreetAddress
    leafs["source-interface"] = callHome.SourceInterface
    leafs["contract-id"] = callHome.ContractId
    leafs["reply-to"] = callHome.ReplyTo
    leafs["from"] = callHome.From
    leafs["active"] = callHome.Active
    return leafs
}

func (callHome *CallHome) GetBundleName() string { return "cisco_ios_xr" }

func (callHome *CallHome) GetYangName() string { return "call-home" }

func (callHome *CallHome) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (callHome *CallHome) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (callHome *CallHome) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (callHome *CallHome) SetParent(parent types.Entity) { callHome.parent = parent }

func (callHome *CallHome) GetParent() types.Entity { return callHome.parent }

func (callHome *CallHome) GetParentYangName() string { return "Cisco-IOS-XR-call-home-cfg" }

// CallHome_MailServers
// List of call-home mail_server
type CallHome_MailServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Email server. The type is slice of CallHome_MailServers_MailServer.
    MailServer []CallHome_MailServers_MailServer
}

func (mailServers *CallHome_MailServers) GetFilter() yfilter.YFilter { return mailServers.YFilter }

func (mailServers *CallHome_MailServers) SetFilter(yf yfilter.YFilter) { mailServers.YFilter = yf }

func (mailServers *CallHome_MailServers) GetGoName(yname string) string {
    if yname == "mail-server" { return "MailServer" }
    return ""
}

func (mailServers *CallHome_MailServers) GetSegmentPath() string {
    return "mail-servers"
}

func (mailServers *CallHome_MailServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mail-server" {
        for _, c := range mailServers.MailServer {
            if mailServers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CallHome_MailServers_MailServer{}
        mailServers.MailServer = append(mailServers.MailServer, child)
        return &mailServers.MailServer[len(mailServers.MailServer)-1]
    }
    return nil
}

func (mailServers *CallHome_MailServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mailServers.MailServer {
        children[mailServers.MailServer[i].GetSegmentPath()] = &mailServers.MailServer[i]
    }
    return children
}

func (mailServers *CallHome_MailServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mailServers *CallHome_MailServers) GetBundleName() string { return "cisco_ios_xr" }

func (mailServers *CallHome_MailServers) GetYangName() string { return "mail-servers" }

func (mailServers *CallHome_MailServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mailServers *CallHome_MailServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mailServers *CallHome_MailServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mailServers *CallHome_MailServers) SetParent(parent types.Entity) { mailServers.parent = parent }

func (mailServers *CallHome_MailServers) GetParent() types.Entity { return mailServers.parent }

func (mailServers *CallHome_MailServers) GetParentYangName() string { return "call-home" }

// CallHome_MailServers_MailServer
// Email server
type CallHome_MailServers_MailServer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Email server. The type is string.
    MailServAddress interface{}

    // Mail server with lower # will be used first. The type is interface{} with
    // range: 1..100.
    Priority interface{}
}

func (mailServer *CallHome_MailServers_MailServer) GetFilter() yfilter.YFilter { return mailServer.YFilter }

func (mailServer *CallHome_MailServers_MailServer) SetFilter(yf yfilter.YFilter) { mailServer.YFilter = yf }

func (mailServer *CallHome_MailServers_MailServer) GetGoName(yname string) string {
    if yname == "mail-serv-address" { return "MailServAddress" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (mailServer *CallHome_MailServers_MailServer) GetSegmentPath() string {
    return "mail-server" + "[mail-serv-address='" + fmt.Sprintf("%v", mailServer.MailServAddress) + "']"
}

func (mailServer *CallHome_MailServers_MailServer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mailServer *CallHome_MailServers_MailServer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mailServer *CallHome_MailServers_MailServer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mail-serv-address"] = mailServer.MailServAddress
    leafs["priority"] = mailServer.Priority
    return leafs
}

func (mailServer *CallHome_MailServers_MailServer) GetBundleName() string { return "cisco_ios_xr" }

func (mailServer *CallHome_MailServers_MailServer) GetYangName() string { return "mail-server" }

func (mailServer *CallHome_MailServers_MailServer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mailServer *CallHome_MailServers_MailServer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mailServer *CallHome_MailServers_MailServer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mailServer *CallHome_MailServers_MailServer) SetParent(parent types.Entity) { mailServer.parent = parent }

func (mailServer *CallHome_MailServers_MailServer) GetParent() types.Entity { return mailServer.parent }

func (mailServer *CallHome_MailServers_MailServer) GetParentYangName() string { return "mail-servers" }

// CallHome_SyslogThrottling
// Enable or disable call-home syslog message
// throttling
type CallHome_SyslogThrottling struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Active syslog throttling. The type is bool.
    Active interface{}
}

func (syslogThrottling *CallHome_SyslogThrottling) GetFilter() yfilter.YFilter { return syslogThrottling.YFilter }

func (syslogThrottling *CallHome_SyslogThrottling) SetFilter(yf yfilter.YFilter) { syslogThrottling.YFilter = yf }

func (syslogThrottling *CallHome_SyslogThrottling) GetGoName(yname string) string {
    if yname == "active" { return "Active" }
    return ""
}

func (syslogThrottling *CallHome_SyslogThrottling) GetSegmentPath() string {
    return "syslog-throttling"
}

func (syslogThrottling *CallHome_SyslogThrottling) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (syslogThrottling *CallHome_SyslogThrottling) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (syslogThrottling *CallHome_SyslogThrottling) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active"] = syslogThrottling.Active
    return leafs
}

func (syslogThrottling *CallHome_SyslogThrottling) GetBundleName() string { return "cisco_ios_xr" }

func (syslogThrottling *CallHome_SyslogThrottling) GetYangName() string { return "syslog-throttling" }

func (syslogThrottling *CallHome_SyslogThrottling) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (syslogThrottling *CallHome_SyslogThrottling) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (syslogThrottling *CallHome_SyslogThrottling) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (syslogThrottling *CallHome_SyslogThrottling) SetParent(parent types.Entity) { syslogThrottling.parent = parent }

func (syslogThrottling *CallHome_SyslogThrottling) GetParent() types.Entity { return syslogThrottling.parent }

func (syslogThrottling *CallHome_SyslogThrottling) GetParentYangName() string { return "call-home" }

// CallHome_SmartLicensing
// Enable/disable licensing messages. By default is
// enabled.
type CallHome_SmartLicensing struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // To specify existing profile name used for TG so that licensing message. The
    // type is string.
    ProfileName interface{}

    // Active the smart-licensing. The type is interface{}.
    Active interface{}
}

func (smartLicensing *CallHome_SmartLicensing) GetFilter() yfilter.YFilter { return smartLicensing.YFilter }

func (smartLicensing *CallHome_SmartLicensing) SetFilter(yf yfilter.YFilter) { smartLicensing.YFilter = yf }

func (smartLicensing *CallHome_SmartLicensing) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "active" { return "Active" }
    return ""
}

func (smartLicensing *CallHome_SmartLicensing) GetSegmentPath() string {
    return "smart-licensing"
}

func (smartLicensing *CallHome_SmartLicensing) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (smartLicensing *CallHome_SmartLicensing) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (smartLicensing *CallHome_SmartLicensing) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = smartLicensing.ProfileName
    leafs["active"] = smartLicensing.Active
    return leafs
}

func (smartLicensing *CallHome_SmartLicensing) GetBundleName() string { return "cisco_ios_xr" }

func (smartLicensing *CallHome_SmartLicensing) GetYangName() string { return "smart-licensing" }

func (smartLicensing *CallHome_SmartLicensing) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (smartLicensing *CallHome_SmartLicensing) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (smartLicensing *CallHome_SmartLicensing) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (smartLicensing *CallHome_SmartLicensing) SetParent(parent types.Entity) { smartLicensing.parent = parent }

func (smartLicensing *CallHome_SmartLicensing) GetParent() types.Entity { return smartLicensing.parent }

func (smartLicensing *CallHome_SmartLicensing) GetParentYangName() string { return "call-home" }

// CallHome_HttpProxy
// http proxy server address and port
type CallHome_HttpProxy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // http proxy server address. The type is string.
    ServerAddress interface{}

    // http proxy server's port. The type is interface{} with range: 1..65535.
    Port interface{}
}

func (httpProxy *CallHome_HttpProxy) GetFilter() yfilter.YFilter { return httpProxy.YFilter }

func (httpProxy *CallHome_HttpProxy) SetFilter(yf yfilter.YFilter) { httpProxy.YFilter = yf }

func (httpProxy *CallHome_HttpProxy) GetGoName(yname string) string {
    if yname == "server-address" { return "ServerAddress" }
    if yname == "port" { return "Port" }
    return ""
}

func (httpProxy *CallHome_HttpProxy) GetSegmentPath() string {
    return "http-proxy"
}

func (httpProxy *CallHome_HttpProxy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (httpProxy *CallHome_HttpProxy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (httpProxy *CallHome_HttpProxy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-address"] = httpProxy.ServerAddress
    leafs["port"] = httpProxy.Port
    return leafs
}

func (httpProxy *CallHome_HttpProxy) GetBundleName() string { return "cisco_ios_xr" }

func (httpProxy *CallHome_HttpProxy) GetYangName() string { return "http-proxy" }

func (httpProxy *CallHome_HttpProxy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (httpProxy *CallHome_HttpProxy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (httpProxy *CallHome_HttpProxy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (httpProxy *CallHome_HttpProxy) SetParent(parent types.Entity) { httpProxy.parent = parent }

func (httpProxy *CallHome_HttpProxy) GetParent() types.Entity { return httpProxy.parent }

func (httpProxy *CallHome_HttpProxy) GetParentYangName() string { return "call-home" }

// CallHome_Profiles
// List of profiles
type CallHome_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A specific profile. The type is slice of CallHome_Profiles_Profile.
    Profile []CallHome_Profiles_Profile
}

func (profiles *CallHome_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *CallHome_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *CallHome_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *CallHome_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *CallHome_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CallHome_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *CallHome_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *CallHome_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *CallHome_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *CallHome_Profiles) GetYangName() string { return "profiles" }

func (profiles *CallHome_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *CallHome_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *CallHome_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *CallHome_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *CallHome_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *CallHome_Profiles) GetParentYangName() string { return "call-home" }

// CallHome_Profiles_Profile
// A specific profile
type CallHome_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // Create a profile. The type is interface{}.
    Create interface{}

    // none. The type is string.
    MessageFormat interface{}

    // Enable call-home anonymous reporting only. The type is bool.
    Anonymous interface{}

    // To specify message size limit for this profile. The type is interface{}
    // with range: 50..3145728.
    MessageSizeLimit interface{}

    // Activate the current profile. The type is interface{}.
    Active interface{}

    // Choose what data to report.
    ReportType CallHome_Profiles_Profile_ReportType

    // Transport method (http or email).
    Methods CallHome_Profiles_Profile_Methods

    // List of destination address.
    Addresses CallHome_Profiles_Profile_Addresses

    // Subscribe to alert-group.
    SubscribeAlertGroup CallHome_Profiles_Profile_SubscribeAlertGroup
}

func (profile *CallHome_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *CallHome_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *CallHome_Profiles_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "create" { return "Create" }
    if yname == "message-format" { return "MessageFormat" }
    if yname == "anonymous" { return "Anonymous" }
    if yname == "message-size-limit" { return "MessageSizeLimit" }
    if yname == "active" { return "Active" }
    if yname == "report-type" { return "ReportType" }
    if yname == "methods" { return "Methods" }
    if yname == "addresses" { return "Addresses" }
    if yname == "subscribe-alert-group" { return "SubscribeAlertGroup" }
    return ""
}

func (profile *CallHome_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
}

func (profile *CallHome_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "report-type" {
        return &profile.ReportType
    }
    if childYangName == "methods" {
        return &profile.Methods
    }
    if childYangName == "addresses" {
        return &profile.Addresses
    }
    if childYangName == "subscribe-alert-group" {
        return &profile.SubscribeAlertGroup
    }
    return nil
}

func (profile *CallHome_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["report-type"] = &profile.ReportType
    children["methods"] = &profile.Methods
    children["addresses"] = &profile.Addresses
    children["subscribe-alert-group"] = &profile.SubscribeAlertGroup
    return children
}

func (profile *CallHome_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    leafs["create"] = profile.Create
    leafs["message-format"] = profile.MessageFormat
    leafs["anonymous"] = profile.Anonymous
    leafs["message-size-limit"] = profile.MessageSizeLimit
    leafs["active"] = profile.Active
    return leafs
}

func (profile *CallHome_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *CallHome_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *CallHome_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *CallHome_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *CallHome_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *CallHome_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *CallHome_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *CallHome_Profiles_Profile) GetParentYangName() string { return "profiles" }

// CallHome_Profiles_Profile_ReportType
// Choose what data to report
type CallHome_Profiles_Profile_ReportType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Report smart call-home data.
    ReportingCallhomeData CallHome_Profiles_Profile_ReportType_ReportingCallhomeData

    // Report smart licensing data.
    ReportingLicensingData CallHome_Profiles_Profile_ReportType_ReportingLicensingData
}

func (reportType *CallHome_Profiles_Profile_ReportType) GetFilter() yfilter.YFilter { return reportType.YFilter }

func (reportType *CallHome_Profiles_Profile_ReportType) SetFilter(yf yfilter.YFilter) { reportType.YFilter = yf }

func (reportType *CallHome_Profiles_Profile_ReportType) GetGoName(yname string) string {
    if yname == "reporting-callhome-data" { return "ReportingCallhomeData" }
    if yname == "reporting-licensing-data" { return "ReportingLicensingData" }
    return ""
}

func (reportType *CallHome_Profiles_Profile_ReportType) GetSegmentPath() string {
    return "report-type"
}

func (reportType *CallHome_Profiles_Profile_ReportType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reporting-callhome-data" {
        return &reportType.ReportingCallhomeData
    }
    if childYangName == "reporting-licensing-data" {
        return &reportType.ReportingLicensingData
    }
    return nil
}

func (reportType *CallHome_Profiles_Profile_ReportType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["reporting-callhome-data"] = &reportType.ReportingCallhomeData
    children["reporting-licensing-data"] = &reportType.ReportingLicensingData
    return children
}

func (reportType *CallHome_Profiles_Profile_ReportType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (reportType *CallHome_Profiles_Profile_ReportType) GetBundleName() string { return "cisco_ios_xr" }

func (reportType *CallHome_Profiles_Profile_ReportType) GetYangName() string { return "report-type" }

func (reportType *CallHome_Profiles_Profile_ReportType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reportType *CallHome_Profiles_Profile_ReportType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reportType *CallHome_Profiles_Profile_ReportType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reportType *CallHome_Profiles_Profile_ReportType) SetParent(parent types.Entity) { reportType.parent = parent }

func (reportType *CallHome_Profiles_Profile_ReportType) GetParent() types.Entity { return reportType.parent }

func (reportType *CallHome_Profiles_Profile_ReportType) GetParentYangName() string { return "profile" }

// CallHome_Profiles_Profile_ReportType_ReportingCallhomeData
// Report smart call-home data
type CallHome_Profiles_Profile_ReportType_ReportingCallhomeData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable report smart call-home data. The type is bool.
    Enable interface{}
}

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetFilter() yfilter.YFilter { return reportingCallhomeData.YFilter }

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) SetFilter(yf yfilter.YFilter) { reportingCallhomeData.YFilter = yf }

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetSegmentPath() string {
    return "reporting-callhome-data"
}

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = reportingCallhomeData.Enable
    return leafs
}

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetBundleName() string { return "cisco_ios_xr" }

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetYangName() string { return "reporting-callhome-data" }

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) SetParent(parent types.Entity) { reportingCallhomeData.parent = parent }

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetParent() types.Entity { return reportingCallhomeData.parent }

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetParentYangName() string { return "report-type" }

// CallHome_Profiles_Profile_ReportType_ReportingLicensingData
// Report smart licensing data
type CallHome_Profiles_Profile_ReportType_ReportingLicensingData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable report smart licensing data. The type is bool.
    Enable interface{}
}

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetFilter() yfilter.YFilter { return reportingLicensingData.YFilter }

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) SetFilter(yf yfilter.YFilter) { reportingLicensingData.YFilter = yf }

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetSegmentPath() string {
    return "reporting-licensing-data"
}

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = reportingLicensingData.Enable
    return leafs
}

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetBundleName() string { return "cisco_ios_xr" }

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetYangName() string { return "reporting-licensing-data" }

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) SetParent(parent types.Entity) { reportingLicensingData.parent = parent }

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetParent() types.Entity { return reportingLicensingData.parent }

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetParentYangName() string { return "report-type" }

// CallHome_Profiles_Profile_Methods
// Transport method (http or email)
type CallHome_Profiles_Profile_Methods struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transport method. The type is slice of
    // CallHome_Profiles_Profile_Methods_Method.
    Method []CallHome_Profiles_Profile_Methods_Method
}

func (methods *CallHome_Profiles_Profile_Methods) GetFilter() yfilter.YFilter { return methods.YFilter }

func (methods *CallHome_Profiles_Profile_Methods) SetFilter(yf yfilter.YFilter) { methods.YFilter = yf }

func (methods *CallHome_Profiles_Profile_Methods) GetGoName(yname string) string {
    if yname == "method" { return "Method" }
    return ""
}

func (methods *CallHome_Profiles_Profile_Methods) GetSegmentPath() string {
    return "methods"
}

func (methods *CallHome_Profiles_Profile_Methods) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "method" {
        for _, c := range methods.Method {
            if methods.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CallHome_Profiles_Profile_Methods_Method{}
        methods.Method = append(methods.Method, child)
        return &methods.Method[len(methods.Method)-1]
    }
    return nil
}

func (methods *CallHome_Profiles_Profile_Methods) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range methods.Method {
        children[methods.Method[i].GetSegmentPath()] = &methods.Method[i]
    }
    return children
}

func (methods *CallHome_Profiles_Profile_Methods) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (methods *CallHome_Profiles_Profile_Methods) GetBundleName() string { return "cisco_ios_xr" }

func (methods *CallHome_Profiles_Profile_Methods) GetYangName() string { return "methods" }

func (methods *CallHome_Profiles_Profile_Methods) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (methods *CallHome_Profiles_Profile_Methods) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (methods *CallHome_Profiles_Profile_Methods) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (methods *CallHome_Profiles_Profile_Methods) SetParent(parent types.Entity) { methods.parent = parent }

func (methods *CallHome_Profiles_Profile_Methods) GetParent() types.Entity { return methods.parent }

func (methods *CallHome_Profiles_Profile_Methods) GetParentYangName() string { return "profile" }

// CallHome_Profiles_Profile_Methods_Method
// Transport method
type CallHome_Profiles_Profile_Methods_Method struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Transport Method. The type is CallHomeTransMethod.
    Method interface{}

    // Enable this transport method. The type is bool.
    Enable interface{}
}

func (method *CallHome_Profiles_Profile_Methods_Method) GetFilter() yfilter.YFilter { return method.YFilter }

func (method *CallHome_Profiles_Profile_Methods_Method) SetFilter(yf yfilter.YFilter) { method.YFilter = yf }

func (method *CallHome_Profiles_Profile_Methods_Method) GetGoName(yname string) string {
    if yname == "method" { return "Method" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (method *CallHome_Profiles_Profile_Methods_Method) GetSegmentPath() string {
    return "method" + "[method='" + fmt.Sprintf("%v", method.Method) + "']"
}

func (method *CallHome_Profiles_Profile_Methods_Method) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (method *CallHome_Profiles_Profile_Methods_Method) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (method *CallHome_Profiles_Profile_Methods_Method) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method"] = method.Method
    leafs["enable"] = method.Enable
    return leafs
}

func (method *CallHome_Profiles_Profile_Methods_Method) GetBundleName() string { return "cisco_ios_xr" }

func (method *CallHome_Profiles_Profile_Methods_Method) GetYangName() string { return "method" }

func (method *CallHome_Profiles_Profile_Methods_Method) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (method *CallHome_Profiles_Profile_Methods_Method) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (method *CallHome_Profiles_Profile_Methods_Method) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (method *CallHome_Profiles_Profile_Methods_Method) SetParent(parent types.Entity) { method.parent = parent }

func (method *CallHome_Profiles_Profile_Methods_Method) GetParent() types.Entity { return method.parent }

func (method *CallHome_Profiles_Profile_Methods_Method) GetParentYangName() string { return "methods" }

// CallHome_Profiles_Profile_Addresses
// List of destination address
type CallHome_Profiles_Profile_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A specific address. The type is slice of
    // CallHome_Profiles_Profile_Addresses_Address.
    Address []CallHome_Profiles_Profile_Addresses_Address
}

func (addresses *CallHome_Profiles_Profile_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *CallHome_Profiles_Profile_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *CallHome_Profiles_Profile_Addresses) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (addresses *CallHome_Profiles_Profile_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *CallHome_Profiles_Profile_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range addresses.Address {
            if addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CallHome_Profiles_Profile_Addresses_Address{}
        addresses.Address = append(addresses.Address, child)
        return &addresses.Address[len(addresses.Address)-1]
    }
    return nil
}

func (addresses *CallHome_Profiles_Profile_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addresses.Address {
        children[addresses.Address[i].GetSegmentPath()] = &addresses.Address[i]
    }
    return children
}

func (addresses *CallHome_Profiles_Profile_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *CallHome_Profiles_Profile_Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (addresses *CallHome_Profiles_Profile_Addresses) GetYangName() string { return "addresses" }

func (addresses *CallHome_Profiles_Profile_Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addresses *CallHome_Profiles_Profile_Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addresses *CallHome_Profiles_Profile_Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addresses *CallHome_Profiles_Profile_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *CallHome_Profiles_Profile_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *CallHome_Profiles_Profile_Addresses) GetParentYangName() string { return "profile" }

// CallHome_Profiles_Profile_Addresses_Address
// A specific address
type CallHome_Profiles_Profile_Addresses_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Transpotation Method. The type is
    // CallHomeTransMethod.
    Method interface{}

    // This attribute is a key. Destination address (1-200) characters. The type
    // is string with length: 1..200.
    DestinationAddr interface{}

    // Set the address. The type is bool.
    Enable interface{}
}

func (address *CallHome_Profiles_Profile_Addresses_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *CallHome_Profiles_Profile_Addresses_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *CallHome_Profiles_Profile_Addresses_Address) GetGoName(yname string) string {
    if yname == "method" { return "Method" }
    if yname == "destination-addr" { return "DestinationAddr" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (address *CallHome_Profiles_Profile_Addresses_Address) GetSegmentPath() string {
    return "address" + "[method='" + fmt.Sprintf("%v", address.Method) + "']" + "[destination-addr='" + fmt.Sprintf("%v", address.DestinationAddr) + "']"
}

func (address *CallHome_Profiles_Profile_Addresses_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *CallHome_Profiles_Profile_Addresses_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *CallHome_Profiles_Profile_Addresses_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method"] = address.Method
    leafs["destination-addr"] = address.DestinationAddr
    leafs["enable"] = address.Enable
    return leafs
}

func (address *CallHome_Profiles_Profile_Addresses_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *CallHome_Profiles_Profile_Addresses_Address) GetYangName() string { return "address" }

func (address *CallHome_Profiles_Profile_Addresses_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *CallHome_Profiles_Profile_Addresses_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *CallHome_Profiles_Profile_Addresses_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *CallHome_Profiles_Profile_Addresses_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *CallHome_Profiles_Profile_Addresses_Address) GetParent() types.Entity { return address.parent }

func (address *CallHome_Profiles_Profile_Addresses_Address) GetParentYangName() string { return "addresses" }

// CallHome_Profiles_Profile_SubscribeAlertGroup
// Subscribe to alert-group
type CallHome_Profiles_Profile_SubscribeAlertGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // environmental info.
    Environment CallHome_Profiles_Profile_SubscribeAlertGroup_Environment

    // configuration info.
    Configuration CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration

    // snapshot info.
    Snapshot CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot

    // inventory info.
    Inventory CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory

    // Crash info.
    Crash CallHome_Profiles_Profile_SubscribeAlertGroup_Crash

    // syslog info.
    Syslogs CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs
}

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetFilter() yfilter.YFilter { return subscribeAlertGroup.YFilter }

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) SetFilter(yf yfilter.YFilter) { subscribeAlertGroup.YFilter = yf }

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetGoName(yname string) string {
    if yname == "environment" { return "Environment" }
    if yname == "configuration" { return "Configuration" }
    if yname == "snapshot" { return "Snapshot" }
    if yname == "inventory" { return "Inventory" }
    if yname == "crash" { return "Crash" }
    if yname == "syslogs" { return "Syslogs" }
    return ""
}

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetSegmentPath() string {
    return "subscribe-alert-group"
}

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "environment" {
        return &subscribeAlertGroup.Environment
    }
    if childYangName == "configuration" {
        return &subscribeAlertGroup.Configuration
    }
    if childYangName == "snapshot" {
        return &subscribeAlertGroup.Snapshot
    }
    if childYangName == "inventory" {
        return &subscribeAlertGroup.Inventory
    }
    if childYangName == "crash" {
        return &subscribeAlertGroup.Crash
    }
    if childYangName == "syslogs" {
        return &subscribeAlertGroup.Syslogs
    }
    return nil
}

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["environment"] = &subscribeAlertGroup.Environment
    children["configuration"] = &subscribeAlertGroup.Configuration
    children["snapshot"] = &subscribeAlertGroup.Snapshot
    children["inventory"] = &subscribeAlertGroup.Inventory
    children["crash"] = &subscribeAlertGroup.Crash
    children["syslogs"] = &subscribeAlertGroup.Syslogs
    return children
}

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetBundleName() string { return "cisco_ios_xr" }

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetYangName() string { return "subscribe-alert-group" }

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) SetParent(parent types.Entity) { subscribeAlertGroup.parent = parent }

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetParent() types.Entity { return subscribeAlertGroup.parent }

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetParentYangName() string { return "profile" }

// CallHome_Profiles_Profile_SubscribeAlertGroup_Environment
// environmental info
type CallHome_Profiles_Profile_SubscribeAlertGroup_Environment struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Severity. The type is CallHomeEventSeverity.
    Severity interface{}
}

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetFilter() yfilter.YFilter { return environment.YFilter }

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) SetFilter(yf yfilter.YFilter) { environment.YFilter = yf }

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetGoName(yname string) string {
    if yname == "severity" { return "Severity" }
    return ""
}

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetSegmentPath() string {
    return "environment"
}

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["severity"] = environment.Severity
    return leafs
}

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetBundleName() string { return "cisco_ios_xr" }

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetYangName() string { return "environment" }

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) SetParent(parent types.Entity) { environment.parent = parent }

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetParent() types.Entity { return environment.parent }

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetParentYangName() string { return "subscribe-alert-group" }

// CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration
// configuration info
type CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscribe the alert-group. The type is interface{}.
    Subscribe interface{}

    // Periodic call-home message.
    Periodic CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic
}

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetFilter() yfilter.YFilter { return configuration.YFilter }

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) SetFilter(yf yfilter.YFilter) { configuration.YFilter = yf }

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetGoName(yname string) string {
    if yname == "subscribe" { return "Subscribe" }
    if yname == "periodic" { return "Periodic" }
    return ""
}

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetSegmentPath() string {
    return "configuration"
}

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "periodic" {
        return &configuration.Periodic
    }
    return nil
}

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["periodic"] = &configuration.Periodic
    return children
}

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscribe"] = configuration.Subscribe
    return leafs
}

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetBundleName() string { return "cisco_ios_xr" }

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetYangName() string { return "configuration" }

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) SetParent(parent types.Entity) { configuration.parent = parent }

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetParent() types.Entity { return configuration.parent }

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetParentYangName() string { return "subscribe-alert-group" }

// CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic
// Periodic call-home message
type CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // none. The type is CallHomeMailSendInterval.
    Interval interface{}

    // Day. The type is interface{} with range: 0..31.
    Day interface{}

    // Day of week. The type is CallHomeDayOfWeek.
    Weekday interface{}

    // Hour. The type is interface{} with range: 0..23.
    Hour interface{}

    // Minute. The type is interface{} with range: 0..59.
    Minute interface{}
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetFilter() yfilter.YFilter { return periodic.YFilter }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) SetFilter(yf yfilter.YFilter) { periodic.YFilter = yf }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetGoName(yname string) string {
    if yname == "interval" { return "Interval" }
    if yname == "day" { return "Day" }
    if yname == "weekday" { return "Weekday" }
    if yname == "hour" { return "Hour" }
    if yname == "minute" { return "Minute" }
    return ""
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetSegmentPath() string {
    return "periodic"
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval"] = periodic.Interval
    leafs["day"] = periodic.Day
    leafs["weekday"] = periodic.Weekday
    leafs["hour"] = periodic.Hour
    leafs["minute"] = periodic.Minute
    return leafs
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetBundleName() string { return "cisco_ios_xr" }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetYangName() string { return "periodic" }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) SetParent(parent types.Entity) { periodic.parent = parent }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetParent() types.Entity { return periodic.parent }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetParentYangName() string { return "configuration" }

// CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot
// snapshot info
type CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Periodic call-home message.
    Periodic CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic
}

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetFilter() yfilter.YFilter { return snapshot.YFilter }

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) SetFilter(yf yfilter.YFilter) { snapshot.YFilter = yf }

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetGoName(yname string) string {
    if yname == "periodic" { return "Periodic" }
    return ""
}

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetSegmentPath() string {
    return "snapshot"
}

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "periodic" {
        return &snapshot.Periodic
    }
    return nil
}

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["periodic"] = &snapshot.Periodic
    return children
}

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetBundleName() string { return "cisco_ios_xr" }

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetYangName() string { return "snapshot" }

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) SetParent(parent types.Entity) { snapshot.parent = parent }

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetParent() types.Entity { return snapshot.parent }

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetParentYangName() string { return "subscribe-alert-group" }

// CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic
// Periodic call-home message
type CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // none. The type is SnapshotInterval.
    Interval interface{}

    // Day of month. The type is interface{} with range: 0..31.
    Day interface{}

    // Day of week. The type is CallHomeDayOfWeek.
    Weekday interface{}

    // Hour. The type is interface{} with range: 0..23.
    Hour interface{}

    // Minute. The type is interface{} with range: 0..59.
    Minute interface{}
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetFilter() yfilter.YFilter { return periodic.YFilter }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) SetFilter(yf yfilter.YFilter) { periodic.YFilter = yf }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetGoName(yname string) string {
    if yname == "interval" { return "Interval" }
    if yname == "day" { return "Day" }
    if yname == "weekday" { return "Weekday" }
    if yname == "hour" { return "Hour" }
    if yname == "minute" { return "Minute" }
    return ""
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetSegmentPath() string {
    return "periodic"
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval"] = periodic.Interval
    leafs["day"] = periodic.Day
    leafs["weekday"] = periodic.Weekday
    leafs["hour"] = periodic.Hour
    leafs["minute"] = periodic.Minute
    return leafs
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetBundleName() string { return "cisco_ios_xr" }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetYangName() string { return "periodic" }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) SetParent(parent types.Entity) { periodic.parent = parent }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetParent() types.Entity { return periodic.parent }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetParentYangName() string { return "snapshot" }

// CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory
// inventory info
type CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscribe the alert-group. The type is interface{}.
    Subscribe interface{}

    // Periodic call-home message.
    Periodic CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic
}

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetFilter() yfilter.YFilter { return inventory.YFilter }

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) SetFilter(yf yfilter.YFilter) { inventory.YFilter = yf }

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetGoName(yname string) string {
    if yname == "subscribe" { return "Subscribe" }
    if yname == "periodic" { return "Periodic" }
    return ""
}

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetSegmentPath() string {
    return "inventory"
}

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "periodic" {
        return &inventory.Periodic
    }
    return nil
}

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["periodic"] = &inventory.Periodic
    return children
}

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscribe"] = inventory.Subscribe
    return leafs
}

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetBundleName() string { return "cisco_ios_xr" }

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetYangName() string { return "inventory" }

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) SetParent(parent types.Entity) { inventory.parent = parent }

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetParent() types.Entity { return inventory.parent }

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetParentYangName() string { return "subscribe-alert-group" }

// CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic
// Periodic call-home message
type CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // none. The type is CallHomeMailSendInterval.
    Interval interface{}

    // Day of month. The type is interface{} with range: 0..31.
    Day interface{}

    // Day of week. The type is CallHomeDayOfWeek.
    Weekday interface{}

    // Hour. The type is interface{} with range: 0..23.
    Hour interface{}

    // Minute. The type is interface{} with range: 0..59.
    Minute interface{}
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetFilter() yfilter.YFilter { return periodic.YFilter }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) SetFilter(yf yfilter.YFilter) { periodic.YFilter = yf }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetGoName(yname string) string {
    if yname == "interval" { return "Interval" }
    if yname == "day" { return "Day" }
    if yname == "weekday" { return "Weekday" }
    if yname == "hour" { return "Hour" }
    if yname == "minute" { return "Minute" }
    return ""
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetSegmentPath() string {
    return "periodic"
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval"] = periodic.Interval
    leafs["day"] = periodic.Day
    leafs["weekday"] = periodic.Weekday
    leafs["hour"] = periodic.Hour
    leafs["minute"] = periodic.Minute
    return leafs
}

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetBundleName() string { return "cisco_ios_xr" }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetYangName() string { return "periodic" }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) SetParent(parent types.Entity) { periodic.parent = parent }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetParent() types.Entity { return periodic.parent }

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetParentYangName() string { return "inventory" }

// CallHome_Profiles_Profile_SubscribeAlertGroup_Crash
// Crash info
type CallHome_Profiles_Profile_SubscribeAlertGroup_Crash struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscribe crash group. The type is interface{}.
    Subscribe interface{}
}

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetFilter() yfilter.YFilter { return crash.YFilter }

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) SetFilter(yf yfilter.YFilter) { crash.YFilter = yf }

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetGoName(yname string) string {
    if yname == "subscribe" { return "Subscribe" }
    return ""
}

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetSegmentPath() string {
    return "crash"
}

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscribe"] = crash.Subscribe
    return leafs
}

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetBundleName() string { return "cisco_ios_xr" }

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetYangName() string { return "crash" }

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) SetParent(parent types.Entity) { crash.parent = parent }

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetParent() types.Entity { return crash.parent }

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetParentYangName() string { return "subscribe-alert-group" }

// CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs
// syslog info
type CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Syslog message pattern to be matched. The type is slice of
    // CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog.
    Syslog []CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog
}

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetFilter() yfilter.YFilter { return syslogs.YFilter }

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) SetFilter(yf yfilter.YFilter) { syslogs.YFilter = yf }

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetGoName(yname string) string {
    if yname == "syslog" { return "Syslog" }
    return ""
}

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetSegmentPath() string {
    return "syslogs"
}

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "syslog" {
        for _, c := range syslogs.Syslog {
            if syslogs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog{}
        syslogs.Syslog = append(syslogs.Syslog, child)
        return &syslogs.Syslog[len(syslogs.Syslog)-1]
    }
    return nil
}

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range syslogs.Syslog {
        children[syslogs.Syslog[i].GetSegmentPath()] = &syslogs.Syslog[i]
    }
    return children
}

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetBundleName() string { return "cisco_ios_xr" }

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetYangName() string { return "syslogs" }

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) SetParent(parent types.Entity) { syslogs.parent = parent }

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetParent() types.Entity { return syslogs.parent }

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetParentYangName() string { return "subscribe-alert-group" }

// CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog
// Syslog message pattern to be matched
type CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Syslog message pattern to be matched. The type is
    // string with length: 1..80.
    SyslogPattern interface{}

    // Severity. The type is CallHomeEventSeverity.
    Severity interface{}
}

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetFilter() yfilter.YFilter { return syslog.YFilter }

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) SetFilter(yf yfilter.YFilter) { syslog.YFilter = yf }

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetGoName(yname string) string {
    if yname == "syslog-pattern" { return "SyslogPattern" }
    if yname == "severity" { return "Severity" }
    return ""
}

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetSegmentPath() string {
    return "syslog" + "[syslog-pattern='" + fmt.Sprintf("%v", syslog.SyslogPattern) + "']"
}

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["syslog-pattern"] = syslog.SyslogPattern
    leafs["severity"] = syslog.Severity
    return leafs
}

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetBundleName() string { return "cisco_ios_xr" }

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetYangName() string { return "syslog" }

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) SetParent(parent types.Entity) { syslog.parent = parent }

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetParent() types.Entity { return syslog.parent }

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetParentYangName() string { return "syslogs" }

// CallHome_AlertGroups
// List of alert-group
type CallHome_AlertGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A specific alert-group. The type is slice of
    // CallHome_AlertGroups_AlertGroup.
    AlertGroup []CallHome_AlertGroups_AlertGroup
}

func (alertGroups *CallHome_AlertGroups) GetFilter() yfilter.YFilter { return alertGroups.YFilter }

func (alertGroups *CallHome_AlertGroups) SetFilter(yf yfilter.YFilter) { alertGroups.YFilter = yf }

func (alertGroups *CallHome_AlertGroups) GetGoName(yname string) string {
    if yname == "alert-group" { return "AlertGroup" }
    return ""
}

func (alertGroups *CallHome_AlertGroups) GetSegmentPath() string {
    return "alert-groups"
}

func (alertGroups *CallHome_AlertGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alert-group" {
        for _, c := range alertGroups.AlertGroup {
            if alertGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CallHome_AlertGroups_AlertGroup{}
        alertGroups.AlertGroup = append(alertGroups.AlertGroup, child)
        return &alertGroups.AlertGroup[len(alertGroups.AlertGroup)-1]
    }
    return nil
}

func (alertGroups *CallHome_AlertGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range alertGroups.AlertGroup {
        children[alertGroups.AlertGroup[i].GetSegmentPath()] = &alertGroups.AlertGroup[i]
    }
    return children
}

func (alertGroups *CallHome_AlertGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (alertGroups *CallHome_AlertGroups) GetBundleName() string { return "cisco_ios_xr" }

func (alertGroups *CallHome_AlertGroups) GetYangName() string { return "alert-groups" }

func (alertGroups *CallHome_AlertGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alertGroups *CallHome_AlertGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alertGroups *CallHome_AlertGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alertGroups *CallHome_AlertGroups) SetParent(parent types.Entity) { alertGroups.parent = parent }

func (alertGroups *CallHome_AlertGroups) GetParent() types.Entity { return alertGroups.parent }

func (alertGroups *CallHome_AlertGroups) GetParentYangName() string { return "call-home" }

// CallHome_AlertGroups_AlertGroup
// A specific alert-group
type CallHome_AlertGroups_AlertGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    AlertGroupName interface{}

    // Enable the alert-group. The type is bool.
    Enable interface{}

    // Disable the alert-group. The type is bool.
    Disable interface{}
}

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetFilter() yfilter.YFilter { return alertGroup.YFilter }

func (alertGroup *CallHome_AlertGroups_AlertGroup) SetFilter(yf yfilter.YFilter) { alertGroup.YFilter = yf }

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetGoName(yname string) string {
    if yname == "alert-group-name" { return "AlertGroupName" }
    if yname == "enable" { return "Enable" }
    if yname == "disable" { return "Disable" }
    return ""
}

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetSegmentPath() string {
    return "alert-group" + "[alert-group-name='" + fmt.Sprintf("%v", alertGroup.AlertGroupName) + "']"
}

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["alert-group-name"] = alertGroup.AlertGroupName
    leafs["enable"] = alertGroup.Enable
    leafs["disable"] = alertGroup.Disable
    return leafs
}

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetBundleName() string { return "cisco_ios_xr" }

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetYangName() string { return "alert-group" }

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alertGroup *CallHome_AlertGroups_AlertGroup) SetParent(parent types.Entity) { alertGroup.parent = parent }

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetParent() types.Entity { return alertGroup.parent }

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetParentYangName() string { return "alert-groups" }

// CallHome_DataPrivacies
// Set call-home data-privacy
type CallHome_DataPrivacies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // level hostname. The type is slice of CallHome_DataPrivacies_DataPrivacy.
    DataPrivacy []CallHome_DataPrivacies_DataPrivacy
}

func (dataPrivacies *CallHome_DataPrivacies) GetFilter() yfilter.YFilter { return dataPrivacies.YFilter }

func (dataPrivacies *CallHome_DataPrivacies) SetFilter(yf yfilter.YFilter) { dataPrivacies.YFilter = yf }

func (dataPrivacies *CallHome_DataPrivacies) GetGoName(yname string) string {
    if yname == "data-privacy" { return "DataPrivacy" }
    return ""
}

func (dataPrivacies *CallHome_DataPrivacies) GetSegmentPath() string {
    return "data-privacies"
}

func (dataPrivacies *CallHome_DataPrivacies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data-privacy" {
        for _, c := range dataPrivacies.DataPrivacy {
            if dataPrivacies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CallHome_DataPrivacies_DataPrivacy{}
        dataPrivacies.DataPrivacy = append(dataPrivacies.DataPrivacy, child)
        return &dataPrivacies.DataPrivacy[len(dataPrivacies.DataPrivacy)-1]
    }
    return nil
}

func (dataPrivacies *CallHome_DataPrivacies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dataPrivacies.DataPrivacy {
        children[dataPrivacies.DataPrivacy[i].GetSegmentPath()] = &dataPrivacies.DataPrivacy[i]
    }
    return children
}

func (dataPrivacies *CallHome_DataPrivacies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dataPrivacies *CallHome_DataPrivacies) GetBundleName() string { return "cisco_ios_xr" }

func (dataPrivacies *CallHome_DataPrivacies) GetYangName() string { return "data-privacies" }

func (dataPrivacies *CallHome_DataPrivacies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataPrivacies *CallHome_DataPrivacies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataPrivacies *CallHome_DataPrivacies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataPrivacies *CallHome_DataPrivacies) SetParent(parent types.Entity) { dataPrivacies.parent = parent }

func (dataPrivacies *CallHome_DataPrivacies) GetParent() types.Entity { return dataPrivacies.parent }

func (dataPrivacies *CallHome_DataPrivacies) GetParentYangName() string { return "call-home" }

// CallHome_DataPrivacies_DataPrivacy
// level hostname
type CallHome_DataPrivacies_DataPrivacy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Data privacy type (hostname or level). The type is
    // string.
    HostName interface{}

    // Set call-home data-privacy level. The type is DataPrivacyLevel.
    Level interface{}
}

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetFilter() yfilter.YFilter { return dataPrivacy.YFilter }

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) SetFilter(yf yfilter.YFilter) { dataPrivacy.YFilter = yf }

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetGoName(yname string) string {
    if yname == "host-name" { return "HostName" }
    if yname == "level" { return "Level" }
    return ""
}

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetSegmentPath() string {
    return "data-privacy" + "[host-name='" + fmt.Sprintf("%v", dataPrivacy.HostName) + "']"
}

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-name"] = dataPrivacy.HostName
    leafs["level"] = dataPrivacy.Level
    return leafs
}

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetBundleName() string { return "cisco_ios_xr" }

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetYangName() string { return "data-privacy" }

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) SetParent(parent types.Entity) { dataPrivacy.parent = parent }

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetParent() types.Entity { return dataPrivacy.parent }

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetParentYangName() string { return "data-privacies" }

// CallHome_AlertGroupConfig
// alert-group config
type CallHome_AlertGroupConfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // snapshot for adding CLI command.
    SnapshotCommands CallHome_AlertGroupConfig_SnapshotCommands
}

func (alertGroupConfig *CallHome_AlertGroupConfig) GetFilter() yfilter.YFilter { return alertGroupConfig.YFilter }

func (alertGroupConfig *CallHome_AlertGroupConfig) SetFilter(yf yfilter.YFilter) { alertGroupConfig.YFilter = yf }

func (alertGroupConfig *CallHome_AlertGroupConfig) GetGoName(yname string) string {
    if yname == "snapshot-commands" { return "SnapshotCommands" }
    return ""
}

func (alertGroupConfig *CallHome_AlertGroupConfig) GetSegmentPath() string {
    return "alert-group-config"
}

func (alertGroupConfig *CallHome_AlertGroupConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snapshot-commands" {
        return &alertGroupConfig.SnapshotCommands
    }
    return nil
}

func (alertGroupConfig *CallHome_AlertGroupConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["snapshot-commands"] = &alertGroupConfig.SnapshotCommands
    return children
}

func (alertGroupConfig *CallHome_AlertGroupConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (alertGroupConfig *CallHome_AlertGroupConfig) GetBundleName() string { return "cisco_ios_xr" }

func (alertGroupConfig *CallHome_AlertGroupConfig) GetYangName() string { return "alert-group-config" }

func (alertGroupConfig *CallHome_AlertGroupConfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alertGroupConfig *CallHome_AlertGroupConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alertGroupConfig *CallHome_AlertGroupConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alertGroupConfig *CallHome_AlertGroupConfig) SetParent(parent types.Entity) { alertGroupConfig.parent = parent }

func (alertGroupConfig *CallHome_AlertGroupConfig) GetParent() types.Entity { return alertGroupConfig.parent }

func (alertGroupConfig *CallHome_AlertGroupConfig) GetParentYangName() string { return "call-home" }

// CallHome_AlertGroupConfig_SnapshotCommands
// snapshot for adding CLI command
type CallHome_AlertGroupConfig_SnapshotCommands struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A specific CLI cmd for snapshot. The type is slice of
    // CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand.
    SnapshotCommand []CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand
}

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetFilter() yfilter.YFilter { return snapshotCommands.YFilter }

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) SetFilter(yf yfilter.YFilter) { snapshotCommands.YFilter = yf }

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetGoName(yname string) string {
    if yname == "snapshot-command" { return "SnapshotCommand" }
    return ""
}

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetSegmentPath() string {
    return "snapshot-commands"
}

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snapshot-command" {
        for _, c := range snapshotCommands.SnapshotCommand {
            if snapshotCommands.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand{}
        snapshotCommands.SnapshotCommand = append(snapshotCommands.SnapshotCommand, child)
        return &snapshotCommands.SnapshotCommand[len(snapshotCommands.SnapshotCommand)-1]
    }
    return nil
}

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range snapshotCommands.SnapshotCommand {
        children[snapshotCommands.SnapshotCommand[i].GetSegmentPath()] = &snapshotCommands.SnapshotCommand[i]
    }
    return children
}

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetBundleName() string { return "cisco_ios_xr" }

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetYangName() string { return "snapshot-commands" }

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) SetParent(parent types.Entity) { snapshotCommands.parent = parent }

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetParent() types.Entity { return snapshotCommands.parent }

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetParentYangName() string { return "alert-group-config" }

// CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand
// A specific CLI cmd for snapshot
type CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. new added command. The type is string with length:
    // 1..127.
    Command interface{}

    // enable snapshot cmd. The type is interface{}.
    Active interface{}
}

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetFilter() yfilter.YFilter { return snapshotCommand.YFilter }

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) SetFilter(yf yfilter.YFilter) { snapshotCommand.YFilter = yf }

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetGoName(yname string) string {
    if yname == "command" { return "Command" }
    if yname == "active" { return "Active" }
    return ""
}

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetSegmentPath() string {
    return "snapshot-command" + "[command='" + fmt.Sprintf("%v", snapshotCommand.Command) + "']"
}

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["command"] = snapshotCommand.Command
    leafs["active"] = snapshotCommand.Active
    return leafs
}

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetBundleName() string { return "cisco_ios_xr" }

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetYangName() string { return "snapshot-command" }

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) SetParent(parent types.Entity) { snapshotCommand.parent = parent }

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetParent() types.Entity { return snapshotCommand.parent }

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetParentYangName() string { return "snapshot-commands" }

// CallHome_Authorization
// Config aaa authorization, default username is
// callhome
type CallHome_Authorization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Username for authorization. default is callhome. The type is string with
    // length: 1..64.
    Username interface{}

    // Enable call-home aaa-authorization. The type is interface{}.
    Active interface{}
}

func (authorization *CallHome_Authorization) GetFilter() yfilter.YFilter { return authorization.YFilter }

func (authorization *CallHome_Authorization) SetFilter(yf yfilter.YFilter) { authorization.YFilter = yf }

func (authorization *CallHome_Authorization) GetGoName(yname string) string {
    if yname == "username" { return "Username" }
    if yname == "active" { return "Active" }
    return ""
}

func (authorization *CallHome_Authorization) GetSegmentPath() string {
    return "authorization"
}

func (authorization *CallHome_Authorization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authorization *CallHome_Authorization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authorization *CallHome_Authorization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["username"] = authorization.Username
    leafs["active"] = authorization.Active
    return leafs
}

func (authorization *CallHome_Authorization) GetBundleName() string { return "cisco_ios_xr" }

func (authorization *CallHome_Authorization) GetYangName() string { return "authorization" }

func (authorization *CallHome_Authorization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authorization *CallHome_Authorization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authorization *CallHome_Authorization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authorization *CallHome_Authorization) SetParent(parent types.Entity) { authorization.parent = parent }

func (authorization *CallHome_Authorization) GetParent() types.Entity { return authorization.parent }

func (authorization *CallHome_Authorization) GetParentYangName() string { return "call-home" }

