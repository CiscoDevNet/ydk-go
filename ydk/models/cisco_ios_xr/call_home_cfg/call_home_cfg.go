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

// CallHome
// Set CallHome parameters
type CallHome struct {
    EntityData types.CommonEntityData
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
    // pattern: b'[a-zA-Z0-9./-]+'.
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

func (callHome *CallHome) GetEntityData() *types.CommonEntityData {
    callHome.EntityData.YFilter = callHome.YFilter
    callHome.EntityData.YangName = "call-home"
    callHome.EntityData.BundleName = "cisco_ios_xr"
    callHome.EntityData.ParentYangName = "Cisco-IOS-XR-call-home-cfg"
    callHome.EntityData.SegmentPath = "Cisco-IOS-XR-call-home-cfg:call-home"
    callHome.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    callHome.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    callHome.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    callHome.EntityData.Children = make(map[string]types.YChild)
    callHome.EntityData.Children["mail-servers"] = types.YChild{"MailServers", &callHome.MailServers}
    callHome.EntityData.Children["syslog-throttling"] = types.YChild{"SyslogThrottling", &callHome.SyslogThrottling}
    callHome.EntityData.Children["smart-licensing"] = types.YChild{"SmartLicensing", &callHome.SmartLicensing}
    callHome.EntityData.Children["http-proxy"] = types.YChild{"HttpProxy", &callHome.HttpProxy}
    callHome.EntityData.Children["profiles"] = types.YChild{"Profiles", &callHome.Profiles}
    callHome.EntityData.Children["alert-groups"] = types.YChild{"AlertGroups", &callHome.AlertGroups}
    callHome.EntityData.Children["data-privacies"] = types.YChild{"DataPrivacies", &callHome.DataPrivacies}
    callHome.EntityData.Children["alert-group-config"] = types.YChild{"AlertGroupConfig", &callHome.AlertGroupConfig}
    callHome.EntityData.Children["authorization"] = types.YChild{"Authorization", &callHome.Authorization}
    callHome.EntityData.Leafs = make(map[string]types.YLeaf)
    callHome.EntityData.Leafs["customer-id"] = types.YLeaf{"CustomerId", callHome.CustomerId}
    callHome.EntityData.Leafs["phone-number"] = types.YLeaf{"PhoneNumber", callHome.PhoneNumber}
    callHome.EntityData.Leafs["contact-smart-licensing"] = types.YLeaf{"ContactSmartLicensing", callHome.ContactSmartLicensing}
    callHome.EntityData.Leafs["contact-email-address"] = types.YLeaf{"ContactEmailAddress", callHome.ContactEmailAddress}
    callHome.EntityData.Leafs["rate-limit"] = types.YLeaf{"RateLimit", callHome.RateLimit}
    callHome.EntityData.Leafs["site-id"] = types.YLeaf{"SiteId", callHome.SiteId}
    callHome.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", callHome.Vrf}
    callHome.EntityData.Leafs["street-address"] = types.YLeaf{"StreetAddress", callHome.StreetAddress}
    callHome.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", callHome.SourceInterface}
    callHome.EntityData.Leafs["contract-id"] = types.YLeaf{"ContractId", callHome.ContractId}
    callHome.EntityData.Leafs["reply-to"] = types.YLeaf{"ReplyTo", callHome.ReplyTo}
    callHome.EntityData.Leafs["from"] = types.YLeaf{"From", callHome.From}
    callHome.EntityData.Leafs["active"] = types.YLeaf{"Active", callHome.Active}
    return &(callHome.EntityData)
}

// CallHome_MailServers
// List of call-home mail_server
type CallHome_MailServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Email server. The type is slice of CallHome_MailServers_MailServer.
    MailServer []CallHome_MailServers_MailServer
}

func (mailServers *CallHome_MailServers) GetEntityData() *types.CommonEntityData {
    mailServers.EntityData.YFilter = mailServers.YFilter
    mailServers.EntityData.YangName = "mail-servers"
    mailServers.EntityData.BundleName = "cisco_ios_xr"
    mailServers.EntityData.ParentYangName = "call-home"
    mailServers.EntityData.SegmentPath = "mail-servers"
    mailServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mailServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mailServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mailServers.EntityData.Children = make(map[string]types.YChild)
    mailServers.EntityData.Children["mail-server"] = types.YChild{"MailServer", nil}
    for i := range mailServers.MailServer {
        mailServers.EntityData.Children[types.GetSegmentPath(&mailServers.MailServer[i])] = types.YChild{"MailServer", &mailServers.MailServer[i]}
    }
    mailServers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mailServers.EntityData)
}

// CallHome_MailServers_MailServer
// Email server
type CallHome_MailServers_MailServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Email server. The type is string.
    MailServAddress interface{}

    // Mail server with lower # will be used first. The type is interface{} with
    // range: 1..100.
    Priority interface{}
}

func (mailServer *CallHome_MailServers_MailServer) GetEntityData() *types.CommonEntityData {
    mailServer.EntityData.YFilter = mailServer.YFilter
    mailServer.EntityData.YangName = "mail-server"
    mailServer.EntityData.BundleName = "cisco_ios_xr"
    mailServer.EntityData.ParentYangName = "mail-servers"
    mailServer.EntityData.SegmentPath = "mail-server" + "[mail-serv-address='" + fmt.Sprintf("%v", mailServer.MailServAddress) + "']"
    mailServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mailServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mailServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mailServer.EntityData.Children = make(map[string]types.YChild)
    mailServer.EntityData.Leafs = make(map[string]types.YLeaf)
    mailServer.EntityData.Leafs["mail-serv-address"] = types.YLeaf{"MailServAddress", mailServer.MailServAddress}
    mailServer.EntityData.Leafs["priority"] = types.YLeaf{"Priority", mailServer.Priority}
    return &(mailServer.EntityData)
}

// CallHome_SyslogThrottling
// Enable or disable call-home syslog message
// throttling
type CallHome_SyslogThrottling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Active syslog throttling. The type is bool.
    Active interface{}
}

func (syslogThrottling *CallHome_SyslogThrottling) GetEntityData() *types.CommonEntityData {
    syslogThrottling.EntityData.YFilter = syslogThrottling.YFilter
    syslogThrottling.EntityData.YangName = "syslog-throttling"
    syslogThrottling.EntityData.BundleName = "cisco_ios_xr"
    syslogThrottling.EntityData.ParentYangName = "call-home"
    syslogThrottling.EntityData.SegmentPath = "syslog-throttling"
    syslogThrottling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syslogThrottling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syslogThrottling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syslogThrottling.EntityData.Children = make(map[string]types.YChild)
    syslogThrottling.EntityData.Leafs = make(map[string]types.YLeaf)
    syslogThrottling.EntityData.Leafs["active"] = types.YLeaf{"Active", syslogThrottling.Active}
    return &(syslogThrottling.EntityData)
}

// CallHome_SmartLicensing
// Enable/disable licensing messages. By default is
// enabled.
type CallHome_SmartLicensing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // To specify existing profile name used for TG so that licensing message. The
    // type is string.
    ProfileName interface{}

    // Active the smart-licensing. The type is interface{}.
    Active interface{}
}

func (smartLicensing *CallHome_SmartLicensing) GetEntityData() *types.CommonEntityData {
    smartLicensing.EntityData.YFilter = smartLicensing.YFilter
    smartLicensing.EntityData.YangName = "smart-licensing"
    smartLicensing.EntityData.BundleName = "cisco_ios_xr"
    smartLicensing.EntityData.ParentYangName = "call-home"
    smartLicensing.EntityData.SegmentPath = "smart-licensing"
    smartLicensing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    smartLicensing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    smartLicensing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    smartLicensing.EntityData.Children = make(map[string]types.YChild)
    smartLicensing.EntityData.Leafs = make(map[string]types.YLeaf)
    smartLicensing.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", smartLicensing.ProfileName}
    smartLicensing.EntityData.Leafs["active"] = types.YLeaf{"Active", smartLicensing.Active}
    return &(smartLicensing.EntityData)
}

// CallHome_HttpProxy
// http proxy server address and port
type CallHome_HttpProxy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // http proxy server address. The type is string.
    ServerAddress interface{}

    // http proxy server's port. The type is interface{} with range: 1..65535.
    Port interface{}
}

func (httpProxy *CallHome_HttpProxy) GetEntityData() *types.CommonEntityData {
    httpProxy.EntityData.YFilter = httpProxy.YFilter
    httpProxy.EntityData.YangName = "http-proxy"
    httpProxy.EntityData.BundleName = "cisco_ios_xr"
    httpProxy.EntityData.ParentYangName = "call-home"
    httpProxy.EntityData.SegmentPath = "http-proxy"
    httpProxy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    httpProxy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    httpProxy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    httpProxy.EntityData.Children = make(map[string]types.YChild)
    httpProxy.EntityData.Leafs = make(map[string]types.YLeaf)
    httpProxy.EntityData.Leafs["server-address"] = types.YLeaf{"ServerAddress", httpProxy.ServerAddress}
    httpProxy.EntityData.Leafs["port"] = types.YLeaf{"Port", httpProxy.Port}
    return &(httpProxy.EntityData)
}

// CallHome_Profiles
// List of profiles
type CallHome_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A specific profile. The type is slice of CallHome_Profiles_Profile.
    Profile []CallHome_Profiles_Profile
}

func (profiles *CallHome_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "call-home"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profiles.EntityData.Children = make(map[string]types.YChild)
    profiles.EntityData.Children["profile"] = types.YChild{"Profile", nil}
    for i := range profiles.Profile {
        profiles.EntityData.Children[types.GetSegmentPath(&profiles.Profile[i])] = types.YChild{"Profile", &profiles.Profile[i]}
    }
    profiles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(profiles.EntityData)
}

// CallHome_Profiles_Profile
// A specific profile
type CallHome_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (profile *CallHome_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Children["report-type"] = types.YChild{"ReportType", &profile.ReportType}
    profile.EntityData.Children["methods"] = types.YChild{"Methods", &profile.Methods}
    profile.EntityData.Children["addresses"] = types.YChild{"Addresses", &profile.Addresses}
    profile.EntityData.Children["subscribe-alert-group"] = types.YChild{"SubscribeAlertGroup", &profile.SubscribeAlertGroup}
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    profile.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", profile.ProfileName}
    profile.EntityData.Leafs["create"] = types.YLeaf{"Create", profile.Create}
    profile.EntityData.Leafs["message-format"] = types.YLeaf{"MessageFormat", profile.MessageFormat}
    profile.EntityData.Leafs["anonymous"] = types.YLeaf{"Anonymous", profile.Anonymous}
    profile.EntityData.Leafs["message-size-limit"] = types.YLeaf{"MessageSizeLimit", profile.MessageSizeLimit}
    profile.EntityData.Leafs["active"] = types.YLeaf{"Active", profile.Active}
    return &(profile.EntityData)
}

// CallHome_Profiles_Profile_ReportType
// Choose what data to report
type CallHome_Profiles_Profile_ReportType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Report smart call-home data.
    ReportingCallhomeData CallHome_Profiles_Profile_ReportType_ReportingCallhomeData

    // Report smart licensing data.
    ReportingLicensingData CallHome_Profiles_Profile_ReportType_ReportingLicensingData
}

func (reportType *CallHome_Profiles_Profile_ReportType) GetEntityData() *types.CommonEntityData {
    reportType.EntityData.YFilter = reportType.YFilter
    reportType.EntityData.YangName = "report-type"
    reportType.EntityData.BundleName = "cisco_ios_xr"
    reportType.EntityData.ParentYangName = "profile"
    reportType.EntityData.SegmentPath = "report-type"
    reportType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reportType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reportType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reportType.EntityData.Children = make(map[string]types.YChild)
    reportType.EntityData.Children["reporting-callhome-data"] = types.YChild{"ReportingCallhomeData", &reportType.ReportingCallhomeData}
    reportType.EntityData.Children["reporting-licensing-data"] = types.YChild{"ReportingLicensingData", &reportType.ReportingLicensingData}
    reportType.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reportType.EntityData)
}

// CallHome_Profiles_Profile_ReportType_ReportingCallhomeData
// Report smart call-home data
type CallHome_Profiles_Profile_ReportType_ReportingCallhomeData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable report smart call-home data. The type is bool.
    Enable interface{}
}

func (reportingCallhomeData *CallHome_Profiles_Profile_ReportType_ReportingCallhomeData) GetEntityData() *types.CommonEntityData {
    reportingCallhomeData.EntityData.YFilter = reportingCallhomeData.YFilter
    reportingCallhomeData.EntityData.YangName = "reporting-callhome-data"
    reportingCallhomeData.EntityData.BundleName = "cisco_ios_xr"
    reportingCallhomeData.EntityData.ParentYangName = "report-type"
    reportingCallhomeData.EntityData.SegmentPath = "reporting-callhome-data"
    reportingCallhomeData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reportingCallhomeData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reportingCallhomeData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reportingCallhomeData.EntityData.Children = make(map[string]types.YChild)
    reportingCallhomeData.EntityData.Leafs = make(map[string]types.YLeaf)
    reportingCallhomeData.EntityData.Leafs["enable"] = types.YLeaf{"Enable", reportingCallhomeData.Enable}
    return &(reportingCallhomeData.EntityData)
}

// CallHome_Profiles_Profile_ReportType_ReportingLicensingData
// Report smart licensing data
type CallHome_Profiles_Profile_ReportType_ReportingLicensingData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable report smart licensing data. The type is bool.
    Enable interface{}
}

func (reportingLicensingData *CallHome_Profiles_Profile_ReportType_ReportingLicensingData) GetEntityData() *types.CommonEntityData {
    reportingLicensingData.EntityData.YFilter = reportingLicensingData.YFilter
    reportingLicensingData.EntityData.YangName = "reporting-licensing-data"
    reportingLicensingData.EntityData.BundleName = "cisco_ios_xr"
    reportingLicensingData.EntityData.ParentYangName = "report-type"
    reportingLicensingData.EntityData.SegmentPath = "reporting-licensing-data"
    reportingLicensingData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reportingLicensingData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reportingLicensingData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reportingLicensingData.EntityData.Children = make(map[string]types.YChild)
    reportingLicensingData.EntityData.Leafs = make(map[string]types.YLeaf)
    reportingLicensingData.EntityData.Leafs["enable"] = types.YLeaf{"Enable", reportingLicensingData.Enable}
    return &(reportingLicensingData.EntityData)
}

// CallHome_Profiles_Profile_Methods
// Transport method (http or email)
type CallHome_Profiles_Profile_Methods struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transport method. The type is slice of
    // CallHome_Profiles_Profile_Methods_Method.
    Method []CallHome_Profiles_Profile_Methods_Method
}

func (methods *CallHome_Profiles_Profile_Methods) GetEntityData() *types.CommonEntityData {
    methods.EntityData.YFilter = methods.YFilter
    methods.EntityData.YangName = "methods"
    methods.EntityData.BundleName = "cisco_ios_xr"
    methods.EntityData.ParentYangName = "profile"
    methods.EntityData.SegmentPath = "methods"
    methods.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    methods.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    methods.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    methods.EntityData.Children = make(map[string]types.YChild)
    methods.EntityData.Children["method"] = types.YChild{"Method", nil}
    for i := range methods.Method {
        methods.EntityData.Children[types.GetSegmentPath(&methods.Method[i])] = types.YChild{"Method", &methods.Method[i]}
    }
    methods.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(methods.EntityData)
}

// CallHome_Profiles_Profile_Methods_Method
// Transport method
type CallHome_Profiles_Profile_Methods_Method struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Transport Method. The type is CallHomeTransMethod.
    Method interface{}

    // Enable this transport method. The type is bool.
    Enable interface{}
}

func (method *CallHome_Profiles_Profile_Methods_Method) GetEntityData() *types.CommonEntityData {
    method.EntityData.YFilter = method.YFilter
    method.EntityData.YangName = "method"
    method.EntityData.BundleName = "cisco_ios_xr"
    method.EntityData.ParentYangName = "methods"
    method.EntityData.SegmentPath = "method" + "[method='" + fmt.Sprintf("%v", method.Method) + "']"
    method.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    method.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    method.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    method.EntityData.Children = make(map[string]types.YChild)
    method.EntityData.Leafs = make(map[string]types.YLeaf)
    method.EntityData.Leafs["method"] = types.YLeaf{"Method", method.Method}
    method.EntityData.Leafs["enable"] = types.YLeaf{"Enable", method.Enable}
    return &(method.EntityData)
}

// CallHome_Profiles_Profile_Addresses
// List of destination address
type CallHome_Profiles_Profile_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A specific address. The type is slice of
    // CallHome_Profiles_Profile_Addresses_Address.
    Address []CallHome_Profiles_Profile_Addresses_Address
}

func (addresses *CallHome_Profiles_Profile_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "profile"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = make(map[string]types.YChild)
    addresses.EntityData.Children["address"] = types.YChild{"Address", nil}
    for i := range addresses.Address {
        addresses.EntityData.Children[types.GetSegmentPath(&addresses.Address[i])] = types.YChild{"Address", &addresses.Address[i]}
    }
    addresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addresses.EntityData)
}

// CallHome_Profiles_Profile_Addresses_Address
// A specific address
type CallHome_Profiles_Profile_Addresses_Address struct {
    EntityData types.CommonEntityData
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

func (address *CallHome_Profiles_Profile_Addresses_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "addresses"
    address.EntityData.SegmentPath = "address" + "[method='" + fmt.Sprintf("%v", address.Method) + "']" + "[destination-addr='" + fmt.Sprintf("%v", address.DestinationAddr) + "']"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["method"] = types.YLeaf{"Method", address.Method}
    address.EntityData.Leafs["destination-addr"] = types.YLeaf{"DestinationAddr", address.DestinationAddr}
    address.EntityData.Leafs["enable"] = types.YLeaf{"Enable", address.Enable}
    return &(address.EntityData)
}

// CallHome_Profiles_Profile_SubscribeAlertGroup
// Subscribe to alert-group
type CallHome_Profiles_Profile_SubscribeAlertGroup struct {
    EntityData types.CommonEntityData
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

func (subscribeAlertGroup *CallHome_Profiles_Profile_SubscribeAlertGroup) GetEntityData() *types.CommonEntityData {
    subscribeAlertGroup.EntityData.YFilter = subscribeAlertGroup.YFilter
    subscribeAlertGroup.EntityData.YangName = "subscribe-alert-group"
    subscribeAlertGroup.EntityData.BundleName = "cisco_ios_xr"
    subscribeAlertGroup.EntityData.ParentYangName = "profile"
    subscribeAlertGroup.EntityData.SegmentPath = "subscribe-alert-group"
    subscribeAlertGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscribeAlertGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscribeAlertGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscribeAlertGroup.EntityData.Children = make(map[string]types.YChild)
    subscribeAlertGroup.EntityData.Children["environment"] = types.YChild{"Environment", &subscribeAlertGroup.Environment}
    subscribeAlertGroup.EntityData.Children["configuration"] = types.YChild{"Configuration", &subscribeAlertGroup.Configuration}
    subscribeAlertGroup.EntityData.Children["snapshot"] = types.YChild{"Snapshot", &subscribeAlertGroup.Snapshot}
    subscribeAlertGroup.EntityData.Children["inventory"] = types.YChild{"Inventory", &subscribeAlertGroup.Inventory}
    subscribeAlertGroup.EntityData.Children["crash"] = types.YChild{"Crash", &subscribeAlertGroup.Crash}
    subscribeAlertGroup.EntityData.Children["syslogs"] = types.YChild{"Syslogs", &subscribeAlertGroup.Syslogs}
    subscribeAlertGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(subscribeAlertGroup.EntityData)
}

// CallHome_Profiles_Profile_SubscribeAlertGroup_Environment
// environmental info
type CallHome_Profiles_Profile_SubscribeAlertGroup_Environment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Severity. The type is CallHomeEventSeverity.
    Severity interface{}
}

func (environment *CallHome_Profiles_Profile_SubscribeAlertGroup_Environment) GetEntityData() *types.CommonEntityData {
    environment.EntityData.YFilter = environment.YFilter
    environment.EntityData.YangName = "environment"
    environment.EntityData.BundleName = "cisco_ios_xr"
    environment.EntityData.ParentYangName = "subscribe-alert-group"
    environment.EntityData.SegmentPath = "environment"
    environment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    environment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    environment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    environment.EntityData.Children = make(map[string]types.YChild)
    environment.EntityData.Leafs = make(map[string]types.YLeaf)
    environment.EntityData.Leafs["severity"] = types.YLeaf{"Severity", environment.Severity}
    return &(environment.EntityData)
}

// CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration
// configuration info
type CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscribe the alert-group. The type is interface{}.
    Subscribe interface{}

    // Periodic call-home message.
    Periodic CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic
}

func (configuration *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration) GetEntityData() *types.CommonEntityData {
    configuration.EntityData.YFilter = configuration.YFilter
    configuration.EntityData.YangName = "configuration"
    configuration.EntityData.BundleName = "cisco_ios_xr"
    configuration.EntityData.ParentYangName = "subscribe-alert-group"
    configuration.EntityData.SegmentPath = "configuration"
    configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuration.EntityData.Children = make(map[string]types.YChild)
    configuration.EntityData.Children["periodic"] = types.YChild{"Periodic", &configuration.Periodic}
    configuration.EntityData.Leafs = make(map[string]types.YLeaf)
    configuration.EntityData.Leafs["subscribe"] = types.YLeaf{"Subscribe", configuration.Subscribe}
    return &(configuration.EntityData)
}

// CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic
// Periodic call-home message
type CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic struct {
    EntityData types.CommonEntityData
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

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Configuration_Periodic) GetEntityData() *types.CommonEntityData {
    periodic.EntityData.YFilter = periodic.YFilter
    periodic.EntityData.YangName = "periodic"
    periodic.EntityData.BundleName = "cisco_ios_xr"
    periodic.EntityData.ParentYangName = "configuration"
    periodic.EntityData.SegmentPath = "periodic"
    periodic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    periodic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    periodic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    periodic.EntityData.Children = make(map[string]types.YChild)
    periodic.EntityData.Leafs = make(map[string]types.YLeaf)
    periodic.EntityData.Leafs["interval"] = types.YLeaf{"Interval", periodic.Interval}
    periodic.EntityData.Leafs["day"] = types.YLeaf{"Day", periodic.Day}
    periodic.EntityData.Leafs["weekday"] = types.YLeaf{"Weekday", periodic.Weekday}
    periodic.EntityData.Leafs["hour"] = types.YLeaf{"Hour", periodic.Hour}
    periodic.EntityData.Leafs["minute"] = types.YLeaf{"Minute", periodic.Minute}
    return &(periodic.EntityData)
}

// CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot
// snapshot info
type CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Periodic call-home message.
    Periodic CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic
}

func (snapshot *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot) GetEntityData() *types.CommonEntityData {
    snapshot.EntityData.YFilter = snapshot.YFilter
    snapshot.EntityData.YangName = "snapshot"
    snapshot.EntityData.BundleName = "cisco_ios_xr"
    snapshot.EntityData.ParentYangName = "subscribe-alert-group"
    snapshot.EntityData.SegmentPath = "snapshot"
    snapshot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snapshot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snapshot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snapshot.EntityData.Children = make(map[string]types.YChild)
    snapshot.EntityData.Children["periodic"] = types.YChild{"Periodic", &snapshot.Periodic}
    snapshot.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(snapshot.EntityData)
}

// CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic
// Periodic call-home message
type CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic struct {
    EntityData types.CommonEntityData
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

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Snapshot_Periodic) GetEntityData() *types.CommonEntityData {
    periodic.EntityData.YFilter = periodic.YFilter
    periodic.EntityData.YangName = "periodic"
    periodic.EntityData.BundleName = "cisco_ios_xr"
    periodic.EntityData.ParentYangName = "snapshot"
    periodic.EntityData.SegmentPath = "periodic"
    periodic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    periodic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    periodic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    periodic.EntityData.Children = make(map[string]types.YChild)
    periodic.EntityData.Leafs = make(map[string]types.YLeaf)
    periodic.EntityData.Leafs["interval"] = types.YLeaf{"Interval", periodic.Interval}
    periodic.EntityData.Leafs["day"] = types.YLeaf{"Day", periodic.Day}
    periodic.EntityData.Leafs["weekday"] = types.YLeaf{"Weekday", periodic.Weekday}
    periodic.EntityData.Leafs["hour"] = types.YLeaf{"Hour", periodic.Hour}
    periodic.EntityData.Leafs["minute"] = types.YLeaf{"Minute", periodic.Minute}
    return &(periodic.EntityData)
}

// CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory
// inventory info
type CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscribe the alert-group. The type is interface{}.
    Subscribe interface{}

    // Periodic call-home message.
    Periodic CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic
}

func (inventory *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory) GetEntityData() *types.CommonEntityData {
    inventory.EntityData.YFilter = inventory.YFilter
    inventory.EntityData.YangName = "inventory"
    inventory.EntityData.BundleName = "cisco_ios_xr"
    inventory.EntityData.ParentYangName = "subscribe-alert-group"
    inventory.EntityData.SegmentPath = "inventory"
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = make(map[string]types.YChild)
    inventory.EntityData.Children["periodic"] = types.YChild{"Periodic", &inventory.Periodic}
    inventory.EntityData.Leafs = make(map[string]types.YLeaf)
    inventory.EntityData.Leafs["subscribe"] = types.YLeaf{"Subscribe", inventory.Subscribe}
    return &(inventory.EntityData)
}

// CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic
// Periodic call-home message
type CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic struct {
    EntityData types.CommonEntityData
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

func (periodic *CallHome_Profiles_Profile_SubscribeAlertGroup_Inventory_Periodic) GetEntityData() *types.CommonEntityData {
    periodic.EntityData.YFilter = periodic.YFilter
    periodic.EntityData.YangName = "periodic"
    periodic.EntityData.BundleName = "cisco_ios_xr"
    periodic.EntityData.ParentYangName = "inventory"
    periodic.EntityData.SegmentPath = "periodic"
    periodic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    periodic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    periodic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    periodic.EntityData.Children = make(map[string]types.YChild)
    periodic.EntityData.Leafs = make(map[string]types.YLeaf)
    periodic.EntityData.Leafs["interval"] = types.YLeaf{"Interval", periodic.Interval}
    periodic.EntityData.Leafs["day"] = types.YLeaf{"Day", periodic.Day}
    periodic.EntityData.Leafs["weekday"] = types.YLeaf{"Weekday", periodic.Weekday}
    periodic.EntityData.Leafs["hour"] = types.YLeaf{"Hour", periodic.Hour}
    periodic.EntityData.Leafs["minute"] = types.YLeaf{"Minute", periodic.Minute}
    return &(periodic.EntityData)
}

// CallHome_Profiles_Profile_SubscribeAlertGroup_Crash
// Crash info
type CallHome_Profiles_Profile_SubscribeAlertGroup_Crash struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscribe crash group. The type is interface{}.
    Subscribe interface{}
}

func (crash *CallHome_Profiles_Profile_SubscribeAlertGroup_Crash) GetEntityData() *types.CommonEntityData {
    crash.EntityData.YFilter = crash.YFilter
    crash.EntityData.YangName = "crash"
    crash.EntityData.BundleName = "cisco_ios_xr"
    crash.EntityData.ParentYangName = "subscribe-alert-group"
    crash.EntityData.SegmentPath = "crash"
    crash.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crash.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crash.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crash.EntityData.Children = make(map[string]types.YChild)
    crash.EntityData.Leafs = make(map[string]types.YLeaf)
    crash.EntityData.Leafs["subscribe"] = types.YLeaf{"Subscribe", crash.Subscribe}
    return &(crash.EntityData)
}

// CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs
// syslog info
type CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Syslog message pattern to be matched. The type is slice of
    // CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog.
    Syslog []CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog
}

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetEntityData() *types.CommonEntityData {
    syslogs.EntityData.YFilter = syslogs.YFilter
    syslogs.EntityData.YangName = "syslogs"
    syslogs.EntityData.BundleName = "cisco_ios_xr"
    syslogs.EntityData.ParentYangName = "subscribe-alert-group"
    syslogs.EntityData.SegmentPath = "syslogs"
    syslogs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syslogs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syslogs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syslogs.EntityData.Children = make(map[string]types.YChild)
    syslogs.EntityData.Children["syslog"] = types.YChild{"Syslog", nil}
    for i := range syslogs.Syslog {
        syslogs.EntityData.Children[types.GetSegmentPath(&syslogs.Syslog[i])] = types.YChild{"Syslog", &syslogs.Syslog[i]}
    }
    syslogs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(syslogs.EntityData)
}

// CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog
// Syslog message pattern to be matched
type CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Syslog message pattern to be matched. The type is
    // string with length: 1..80.
    SyslogPattern interface{}

    // Severity. The type is CallHomeEventSeverity.
    Severity interface{}
}

func (syslog *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog) GetEntityData() *types.CommonEntityData {
    syslog.EntityData.YFilter = syslog.YFilter
    syslog.EntityData.YangName = "syslog"
    syslog.EntityData.BundleName = "cisco_ios_xr"
    syslog.EntityData.ParentYangName = "syslogs"
    syslog.EntityData.SegmentPath = "syslog" + "[syslog-pattern='" + fmt.Sprintf("%v", syslog.SyslogPattern) + "']"
    syslog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syslog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syslog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syslog.EntityData.Children = make(map[string]types.YChild)
    syslog.EntityData.Leafs = make(map[string]types.YLeaf)
    syslog.EntityData.Leafs["syslog-pattern"] = types.YLeaf{"SyslogPattern", syslog.SyslogPattern}
    syslog.EntityData.Leafs["severity"] = types.YLeaf{"Severity", syslog.Severity}
    return &(syslog.EntityData)
}

// CallHome_AlertGroups
// List of alert-group
type CallHome_AlertGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A specific alert-group. The type is slice of
    // CallHome_AlertGroups_AlertGroup.
    AlertGroup []CallHome_AlertGroups_AlertGroup
}

func (alertGroups *CallHome_AlertGroups) GetEntityData() *types.CommonEntityData {
    alertGroups.EntityData.YFilter = alertGroups.YFilter
    alertGroups.EntityData.YangName = "alert-groups"
    alertGroups.EntityData.BundleName = "cisco_ios_xr"
    alertGroups.EntityData.ParentYangName = "call-home"
    alertGroups.EntityData.SegmentPath = "alert-groups"
    alertGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alertGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alertGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alertGroups.EntityData.Children = make(map[string]types.YChild)
    alertGroups.EntityData.Children["alert-group"] = types.YChild{"AlertGroup", nil}
    for i := range alertGroups.AlertGroup {
        alertGroups.EntityData.Children[types.GetSegmentPath(&alertGroups.AlertGroup[i])] = types.YChild{"AlertGroup", &alertGroups.AlertGroup[i]}
    }
    alertGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(alertGroups.EntityData)
}

// CallHome_AlertGroups_AlertGroup
// A specific alert-group
type CallHome_AlertGroups_AlertGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    AlertGroupName interface{}

    // Enable the alert-group. The type is bool.
    Enable interface{}

    // Disable the alert-group. The type is bool.
    Disable interface{}
}

func (alertGroup *CallHome_AlertGroups_AlertGroup) GetEntityData() *types.CommonEntityData {
    alertGroup.EntityData.YFilter = alertGroup.YFilter
    alertGroup.EntityData.YangName = "alert-group"
    alertGroup.EntityData.BundleName = "cisco_ios_xr"
    alertGroup.EntityData.ParentYangName = "alert-groups"
    alertGroup.EntityData.SegmentPath = "alert-group" + "[alert-group-name='" + fmt.Sprintf("%v", alertGroup.AlertGroupName) + "']"
    alertGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alertGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alertGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alertGroup.EntityData.Children = make(map[string]types.YChild)
    alertGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    alertGroup.EntityData.Leafs["alert-group-name"] = types.YLeaf{"AlertGroupName", alertGroup.AlertGroupName}
    alertGroup.EntityData.Leafs["enable"] = types.YLeaf{"Enable", alertGroup.Enable}
    alertGroup.EntityData.Leafs["disable"] = types.YLeaf{"Disable", alertGroup.Disable}
    return &(alertGroup.EntityData)
}

// CallHome_DataPrivacies
// Set call-home data-privacy
type CallHome_DataPrivacies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // level hostname. The type is slice of CallHome_DataPrivacies_DataPrivacy.
    DataPrivacy []CallHome_DataPrivacies_DataPrivacy
}

func (dataPrivacies *CallHome_DataPrivacies) GetEntityData() *types.CommonEntityData {
    dataPrivacies.EntityData.YFilter = dataPrivacies.YFilter
    dataPrivacies.EntityData.YangName = "data-privacies"
    dataPrivacies.EntityData.BundleName = "cisco_ios_xr"
    dataPrivacies.EntityData.ParentYangName = "call-home"
    dataPrivacies.EntityData.SegmentPath = "data-privacies"
    dataPrivacies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataPrivacies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataPrivacies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataPrivacies.EntityData.Children = make(map[string]types.YChild)
    dataPrivacies.EntityData.Children["data-privacy"] = types.YChild{"DataPrivacy", nil}
    for i := range dataPrivacies.DataPrivacy {
        dataPrivacies.EntityData.Children[types.GetSegmentPath(&dataPrivacies.DataPrivacy[i])] = types.YChild{"DataPrivacy", &dataPrivacies.DataPrivacy[i]}
    }
    dataPrivacies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dataPrivacies.EntityData)
}

// CallHome_DataPrivacies_DataPrivacy
// level hostname
type CallHome_DataPrivacies_DataPrivacy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Data privacy type (hostname or level). The type is
    // string.
    HostName interface{}

    // Set call-home data-privacy level. The type is DataPrivacyLevel.
    Level interface{}
}

func (dataPrivacy *CallHome_DataPrivacies_DataPrivacy) GetEntityData() *types.CommonEntityData {
    dataPrivacy.EntityData.YFilter = dataPrivacy.YFilter
    dataPrivacy.EntityData.YangName = "data-privacy"
    dataPrivacy.EntityData.BundleName = "cisco_ios_xr"
    dataPrivacy.EntityData.ParentYangName = "data-privacies"
    dataPrivacy.EntityData.SegmentPath = "data-privacy" + "[host-name='" + fmt.Sprintf("%v", dataPrivacy.HostName) + "']"
    dataPrivacy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataPrivacy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataPrivacy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataPrivacy.EntityData.Children = make(map[string]types.YChild)
    dataPrivacy.EntityData.Leafs = make(map[string]types.YLeaf)
    dataPrivacy.EntityData.Leafs["host-name"] = types.YLeaf{"HostName", dataPrivacy.HostName}
    dataPrivacy.EntityData.Leafs["level"] = types.YLeaf{"Level", dataPrivacy.Level}
    return &(dataPrivacy.EntityData)
}

// CallHome_AlertGroupConfig
// alert-group config
type CallHome_AlertGroupConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // snapshot for adding CLI command.
    SnapshotCommands CallHome_AlertGroupConfig_SnapshotCommands
}

func (alertGroupConfig *CallHome_AlertGroupConfig) GetEntityData() *types.CommonEntityData {
    alertGroupConfig.EntityData.YFilter = alertGroupConfig.YFilter
    alertGroupConfig.EntityData.YangName = "alert-group-config"
    alertGroupConfig.EntityData.BundleName = "cisco_ios_xr"
    alertGroupConfig.EntityData.ParentYangName = "call-home"
    alertGroupConfig.EntityData.SegmentPath = "alert-group-config"
    alertGroupConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alertGroupConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alertGroupConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alertGroupConfig.EntityData.Children = make(map[string]types.YChild)
    alertGroupConfig.EntityData.Children["snapshot-commands"] = types.YChild{"SnapshotCommands", &alertGroupConfig.SnapshotCommands}
    alertGroupConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(alertGroupConfig.EntityData)
}

// CallHome_AlertGroupConfig_SnapshotCommands
// snapshot for adding CLI command
type CallHome_AlertGroupConfig_SnapshotCommands struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A specific CLI cmd for snapshot. The type is slice of
    // CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand.
    SnapshotCommand []CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand
}

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetEntityData() *types.CommonEntityData {
    snapshotCommands.EntityData.YFilter = snapshotCommands.YFilter
    snapshotCommands.EntityData.YangName = "snapshot-commands"
    snapshotCommands.EntityData.BundleName = "cisco_ios_xr"
    snapshotCommands.EntityData.ParentYangName = "alert-group-config"
    snapshotCommands.EntityData.SegmentPath = "snapshot-commands"
    snapshotCommands.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snapshotCommands.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snapshotCommands.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snapshotCommands.EntityData.Children = make(map[string]types.YChild)
    snapshotCommands.EntityData.Children["snapshot-command"] = types.YChild{"SnapshotCommand", nil}
    for i := range snapshotCommands.SnapshotCommand {
        snapshotCommands.EntityData.Children[types.GetSegmentPath(&snapshotCommands.SnapshotCommand[i])] = types.YChild{"SnapshotCommand", &snapshotCommands.SnapshotCommand[i]}
    }
    snapshotCommands.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(snapshotCommands.EntityData)
}

// CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand
// A specific CLI cmd for snapshot
type CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. new added command. The type is string with length:
    // 1..127.
    Command interface{}

    // enable snapshot cmd. The type is interface{}.
    Active interface{}
}

func (snapshotCommand *CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand) GetEntityData() *types.CommonEntityData {
    snapshotCommand.EntityData.YFilter = snapshotCommand.YFilter
    snapshotCommand.EntityData.YangName = "snapshot-command"
    snapshotCommand.EntityData.BundleName = "cisco_ios_xr"
    snapshotCommand.EntityData.ParentYangName = "snapshot-commands"
    snapshotCommand.EntityData.SegmentPath = "snapshot-command" + "[command='" + fmt.Sprintf("%v", snapshotCommand.Command) + "']"
    snapshotCommand.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snapshotCommand.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snapshotCommand.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snapshotCommand.EntityData.Children = make(map[string]types.YChild)
    snapshotCommand.EntityData.Leafs = make(map[string]types.YLeaf)
    snapshotCommand.EntityData.Leafs["command"] = types.YLeaf{"Command", snapshotCommand.Command}
    snapshotCommand.EntityData.Leafs["active"] = types.YLeaf{"Active", snapshotCommand.Active}
    return &(snapshotCommand.EntityData)
}

// CallHome_Authorization
// Config aaa authorization, default username is
// callhome
type CallHome_Authorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Username for authorization. default is callhome. The type is string with
    // length: 1..64.
    Username interface{}

    // Enable call-home aaa-authorization. The type is interface{}.
    Active interface{}
}

func (authorization *CallHome_Authorization) GetEntityData() *types.CommonEntityData {
    authorization.EntityData.YFilter = authorization.YFilter
    authorization.EntityData.YangName = "authorization"
    authorization.EntityData.BundleName = "cisco_ios_xr"
    authorization.EntityData.ParentYangName = "call-home"
    authorization.EntityData.SegmentPath = "authorization"
    authorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorization.EntityData.Children = make(map[string]types.YChild)
    authorization.EntityData.Leafs = make(map[string]types.YLeaf)
    authorization.EntityData.Leafs["username"] = types.YLeaf{"Username", authorization.Username}
    authorization.EntityData.Leafs["active"] = types.YLeaf{"Active", authorization.Active}
    return &(authorization.EntityData)
}

