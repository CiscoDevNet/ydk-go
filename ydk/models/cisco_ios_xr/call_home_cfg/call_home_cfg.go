// This module contains a collection of YANG definitions
// for Cisco IOS-XR call-home package configuration.
// 
// This module contains definitions
// for the following management objects:
//   call-home: Set CallHome parameters
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    // pattern: [a-zA-Z0-9._/-]+.
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
    callHome.EntityData.AbsolutePath = callHome.EntityData.SegmentPath
    callHome.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    callHome.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    callHome.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    callHome.EntityData.Children = types.NewOrderedMap()
    callHome.EntityData.Children.Append("mail-servers", types.YChild{"MailServers", &callHome.MailServers})
    callHome.EntityData.Children.Append("syslog-throttling", types.YChild{"SyslogThrottling", &callHome.SyslogThrottling})
    callHome.EntityData.Children.Append("http-proxy", types.YChild{"HttpProxy", &callHome.HttpProxy})
    callHome.EntityData.Children.Append("profiles", types.YChild{"Profiles", &callHome.Profiles})
    callHome.EntityData.Children.Append("alert-groups", types.YChild{"AlertGroups", &callHome.AlertGroups})
    callHome.EntityData.Children.Append("data-privacies", types.YChild{"DataPrivacies", &callHome.DataPrivacies})
    callHome.EntityData.Children.Append("alert-group-config", types.YChild{"AlertGroupConfig", &callHome.AlertGroupConfig})
    callHome.EntityData.Children.Append("authorization", types.YChild{"Authorization", &callHome.Authorization})
    callHome.EntityData.Leafs = types.NewOrderedMap()
    callHome.EntityData.Leafs.Append("customer-id", types.YLeaf{"CustomerId", callHome.CustomerId})
    callHome.EntityData.Leafs.Append("phone-number", types.YLeaf{"PhoneNumber", callHome.PhoneNumber})
    callHome.EntityData.Leafs.Append("contact-smart-licensing", types.YLeaf{"ContactSmartLicensing", callHome.ContactSmartLicensing})
    callHome.EntityData.Leafs.Append("contact-email-address", types.YLeaf{"ContactEmailAddress", callHome.ContactEmailAddress})
    callHome.EntityData.Leafs.Append("rate-limit", types.YLeaf{"RateLimit", callHome.RateLimit})
    callHome.EntityData.Leafs.Append("site-id", types.YLeaf{"SiteId", callHome.SiteId})
    callHome.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", callHome.Vrf})
    callHome.EntityData.Leafs.Append("street-address", types.YLeaf{"StreetAddress", callHome.StreetAddress})
    callHome.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", callHome.SourceInterface})
    callHome.EntityData.Leafs.Append("contract-id", types.YLeaf{"ContractId", callHome.ContractId})
    callHome.EntityData.Leafs.Append("reply-to", types.YLeaf{"ReplyTo", callHome.ReplyTo})
    callHome.EntityData.Leafs.Append("from", types.YLeaf{"From", callHome.From})
    callHome.EntityData.Leafs.Append("active", types.YLeaf{"Active", callHome.Active})

    callHome.EntityData.YListKeys = []string {}

    return &(callHome.EntityData)
}

// CallHome_MailServers
// List of call-home mail_server
type CallHome_MailServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Email server. The type is slice of CallHome_MailServers_MailServer.
    MailServer []*CallHome_MailServers_MailServer
}

func (mailServers *CallHome_MailServers) GetEntityData() *types.CommonEntityData {
    mailServers.EntityData.YFilter = mailServers.YFilter
    mailServers.EntityData.YangName = "mail-servers"
    mailServers.EntityData.BundleName = "cisco_ios_xr"
    mailServers.EntityData.ParentYangName = "call-home"
    mailServers.EntityData.SegmentPath = "mail-servers"
    mailServers.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/" + mailServers.EntityData.SegmentPath
    mailServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mailServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mailServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mailServers.EntityData.Children = types.NewOrderedMap()
    mailServers.EntityData.Children.Append("mail-server", types.YChild{"MailServer", nil})
    for i := range mailServers.MailServer {
        mailServers.EntityData.Children.Append(types.GetSegmentPath(mailServers.MailServer[i]), types.YChild{"MailServer", mailServers.MailServer[i]})
    }
    mailServers.EntityData.Leafs = types.NewOrderedMap()

    mailServers.EntityData.YListKeys = []string {}

    return &(mailServers.EntityData)
}

// CallHome_MailServers_MailServer
// Email server
type CallHome_MailServers_MailServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    mailServer.EntityData.SegmentPath = "mail-server" + types.AddKeyToken(mailServer.MailServAddress, "mail-serv-address")
    mailServer.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/mail-servers/" + mailServer.EntityData.SegmentPath
    mailServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mailServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mailServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mailServer.EntityData.Children = types.NewOrderedMap()
    mailServer.EntityData.Leafs = types.NewOrderedMap()
    mailServer.EntityData.Leafs.Append("mail-serv-address", types.YLeaf{"MailServAddress", mailServer.MailServAddress})
    mailServer.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", mailServer.Priority})

    mailServer.EntityData.YListKeys = []string {"MailServAddress"}

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
    syslogThrottling.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/" + syslogThrottling.EntityData.SegmentPath
    syslogThrottling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syslogThrottling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syslogThrottling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syslogThrottling.EntityData.Children = types.NewOrderedMap()
    syslogThrottling.EntityData.Leafs = types.NewOrderedMap()
    syslogThrottling.EntityData.Leafs.Append("active", types.YLeaf{"Active", syslogThrottling.Active})

    syslogThrottling.EntityData.YListKeys = []string {}

    return &(syslogThrottling.EntityData)
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
    httpProxy.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/" + httpProxy.EntityData.SegmentPath
    httpProxy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    httpProxy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    httpProxy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    httpProxy.EntityData.Children = types.NewOrderedMap()
    httpProxy.EntityData.Leafs = types.NewOrderedMap()
    httpProxy.EntityData.Leafs.Append("server-address", types.YLeaf{"ServerAddress", httpProxy.ServerAddress})
    httpProxy.EntityData.Leafs.Append("port", types.YLeaf{"Port", httpProxy.Port})

    httpProxy.EntityData.YListKeys = []string {}

    return &(httpProxy.EntityData)
}

// CallHome_Profiles
// List of profiles
type CallHome_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A specific profile. The type is slice of CallHome_Profiles_Profile.
    Profile []*CallHome_Profiles_Profile
}

func (profiles *CallHome_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "call-home"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/" + profiles.EntityData.SegmentPath
    profiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profiles.EntityData.Children = types.NewOrderedMap()
    profiles.EntityData.Children.Append("profile", types.YChild{"Profile", nil})
    for i := range profiles.Profile {
        profiles.EntityData.Children.Append(types.GetSegmentPath(profiles.Profile[i]), types.YChild{"Profile", profiles.Profile[i]})
    }
    profiles.EntityData.Leafs = types.NewOrderedMap()

    profiles.EntityData.YListKeys = []string {}

    return &(profiles.EntityData)
}

// CallHome_Profiles_Profile
// A specific profile
type CallHome_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (profile *CallHome_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + types.AddKeyToken(profile.ProfileName, "profile-name")
    profile.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/" + profile.EntityData.SegmentPath
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Children.Append("report-type", types.YChild{"ReportType", &profile.ReportType})
    profile.EntityData.Children.Append("methods", types.YChild{"Methods", &profile.Methods})
    profile.EntityData.Children.Append("addresses", types.YChild{"Addresses", &profile.Addresses})
    profile.EntityData.Children.Append("subscribe-alert-group", types.YChild{"SubscribeAlertGroup", &profile.SubscribeAlertGroup})
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", profile.ProfileName})
    profile.EntityData.Leafs.Append("create", types.YLeaf{"Create", profile.Create})
    profile.EntityData.Leafs.Append("message-format", types.YLeaf{"MessageFormat", profile.MessageFormat})
    profile.EntityData.Leafs.Append("anonymous", types.YLeaf{"Anonymous", profile.Anonymous})
    profile.EntityData.Leafs.Append("message-size-limit", types.YLeaf{"MessageSizeLimit", profile.MessageSizeLimit})
    profile.EntityData.Leafs.Append("active", types.YLeaf{"Active", profile.Active})

    profile.EntityData.YListKeys = []string {"ProfileName"}

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
    reportType.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/" + reportType.EntityData.SegmentPath
    reportType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reportType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reportType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reportType.EntityData.Children = types.NewOrderedMap()
    reportType.EntityData.Children.Append("reporting-callhome-data", types.YChild{"ReportingCallhomeData", &reportType.ReportingCallhomeData})
    reportType.EntityData.Children.Append("reporting-licensing-data", types.YChild{"ReportingLicensingData", &reportType.ReportingLicensingData})
    reportType.EntityData.Leafs = types.NewOrderedMap()

    reportType.EntityData.YListKeys = []string {}

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
    reportingCallhomeData.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/report-type/" + reportingCallhomeData.EntityData.SegmentPath
    reportingCallhomeData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reportingCallhomeData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reportingCallhomeData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reportingCallhomeData.EntityData.Children = types.NewOrderedMap()
    reportingCallhomeData.EntityData.Leafs = types.NewOrderedMap()
    reportingCallhomeData.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", reportingCallhomeData.Enable})

    reportingCallhomeData.EntityData.YListKeys = []string {}

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
    reportingLicensingData.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/report-type/" + reportingLicensingData.EntityData.SegmentPath
    reportingLicensingData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reportingLicensingData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reportingLicensingData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reportingLicensingData.EntityData.Children = types.NewOrderedMap()
    reportingLicensingData.EntityData.Leafs = types.NewOrderedMap()
    reportingLicensingData.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", reportingLicensingData.Enable})

    reportingLicensingData.EntityData.YListKeys = []string {}

    return &(reportingLicensingData.EntityData)
}

// CallHome_Profiles_Profile_Methods
// Transport method (http or email)
type CallHome_Profiles_Profile_Methods struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transport method. The type is slice of
    // CallHome_Profiles_Profile_Methods_Method.
    Method []*CallHome_Profiles_Profile_Methods_Method
}

func (methods *CallHome_Profiles_Profile_Methods) GetEntityData() *types.CommonEntityData {
    methods.EntityData.YFilter = methods.YFilter
    methods.EntityData.YangName = "methods"
    methods.EntityData.BundleName = "cisco_ios_xr"
    methods.EntityData.ParentYangName = "profile"
    methods.EntityData.SegmentPath = "methods"
    methods.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/" + methods.EntityData.SegmentPath
    methods.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    methods.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    methods.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    methods.EntityData.Children = types.NewOrderedMap()
    methods.EntityData.Children.Append("method", types.YChild{"Method", nil})
    for i := range methods.Method {
        methods.EntityData.Children.Append(types.GetSegmentPath(methods.Method[i]), types.YChild{"Method", methods.Method[i]})
    }
    methods.EntityData.Leafs = types.NewOrderedMap()

    methods.EntityData.YListKeys = []string {}

    return &(methods.EntityData)
}

// CallHome_Profiles_Profile_Methods_Method
// Transport method
type CallHome_Profiles_Profile_Methods_Method struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    method.EntityData.SegmentPath = "method" + types.AddKeyToken(method.Method, "method")
    method.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/methods/" + method.EntityData.SegmentPath
    method.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    method.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    method.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    method.EntityData.Children = types.NewOrderedMap()
    method.EntityData.Leafs = types.NewOrderedMap()
    method.EntityData.Leafs.Append("method", types.YLeaf{"Method", method.Method})
    method.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", method.Enable})

    method.EntityData.YListKeys = []string {"Method"}

    return &(method.EntityData)
}

// CallHome_Profiles_Profile_Addresses
// List of destination address
type CallHome_Profiles_Profile_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A specific address. The type is slice of
    // CallHome_Profiles_Profile_Addresses_Address.
    Address []*CallHome_Profiles_Profile_Addresses_Address
}

func (addresses *CallHome_Profiles_Profile_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "profile"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/" + addresses.EntityData.SegmentPath
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("address", types.YChild{"Address", nil})
    for i := range addresses.Address {
        addresses.EntityData.Children.Append(types.GetSegmentPath(addresses.Address[i]), types.YChild{"Address", addresses.Address[i]})
    }
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// CallHome_Profiles_Profile_Addresses_Address
// A specific address
type CallHome_Profiles_Profile_Addresses_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    address.EntityData.SegmentPath = "address" + types.AddKeyToken(address.Method, "method") + types.AddKeyToken(address.DestinationAddr, "destination-addr")
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/addresses/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("method", types.YLeaf{"Method", address.Method})
    address.EntityData.Leafs.Append("destination-addr", types.YLeaf{"DestinationAddr", address.DestinationAddr})
    address.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", address.Enable})

    address.EntityData.YListKeys = []string {"Method", "DestinationAddr"}

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
    subscribeAlertGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/" + subscribeAlertGroup.EntityData.SegmentPath
    subscribeAlertGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscribeAlertGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscribeAlertGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscribeAlertGroup.EntityData.Children = types.NewOrderedMap()
    subscribeAlertGroup.EntityData.Children.Append("environment", types.YChild{"Environment", &subscribeAlertGroup.Environment})
    subscribeAlertGroup.EntityData.Children.Append("configuration", types.YChild{"Configuration", &subscribeAlertGroup.Configuration})
    subscribeAlertGroup.EntityData.Children.Append("snapshot", types.YChild{"Snapshot", &subscribeAlertGroup.Snapshot})
    subscribeAlertGroup.EntityData.Children.Append("inventory", types.YChild{"Inventory", &subscribeAlertGroup.Inventory})
    subscribeAlertGroup.EntityData.Children.Append("crash", types.YChild{"Crash", &subscribeAlertGroup.Crash})
    subscribeAlertGroup.EntityData.Children.Append("syslogs", types.YChild{"Syslogs", &subscribeAlertGroup.Syslogs})
    subscribeAlertGroup.EntityData.Leafs = types.NewOrderedMap()

    subscribeAlertGroup.EntityData.YListKeys = []string {}

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
    environment.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/subscribe-alert-group/" + environment.EntityData.SegmentPath
    environment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    environment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    environment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    environment.EntityData.Children = types.NewOrderedMap()
    environment.EntityData.Leafs = types.NewOrderedMap()
    environment.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", environment.Severity})

    environment.EntityData.YListKeys = []string {}

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
    configuration.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/subscribe-alert-group/" + configuration.EntityData.SegmentPath
    configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuration.EntityData.Children = types.NewOrderedMap()
    configuration.EntityData.Children.Append("periodic", types.YChild{"Periodic", &configuration.Periodic})
    configuration.EntityData.Leafs = types.NewOrderedMap()
    configuration.EntityData.Leafs.Append("subscribe", types.YLeaf{"Subscribe", configuration.Subscribe})

    configuration.EntityData.YListKeys = []string {}

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
    periodic.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/subscribe-alert-group/configuration/" + periodic.EntityData.SegmentPath
    periodic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    periodic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    periodic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    periodic.EntityData.Children = types.NewOrderedMap()
    periodic.EntityData.Leafs = types.NewOrderedMap()
    periodic.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", periodic.Interval})
    periodic.EntityData.Leafs.Append("day", types.YLeaf{"Day", periodic.Day})
    periodic.EntityData.Leafs.Append("weekday", types.YLeaf{"Weekday", periodic.Weekday})
    periodic.EntityData.Leafs.Append("hour", types.YLeaf{"Hour", periodic.Hour})
    periodic.EntityData.Leafs.Append("minute", types.YLeaf{"Minute", periodic.Minute})

    periodic.EntityData.YListKeys = []string {}

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
    snapshot.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/subscribe-alert-group/" + snapshot.EntityData.SegmentPath
    snapshot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snapshot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snapshot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snapshot.EntityData.Children = types.NewOrderedMap()
    snapshot.EntityData.Children.Append("periodic", types.YChild{"Periodic", &snapshot.Periodic})
    snapshot.EntityData.Leafs = types.NewOrderedMap()

    snapshot.EntityData.YListKeys = []string {}

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
    periodic.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/subscribe-alert-group/snapshot/" + periodic.EntityData.SegmentPath
    periodic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    periodic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    periodic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    periodic.EntityData.Children = types.NewOrderedMap()
    periodic.EntityData.Leafs = types.NewOrderedMap()
    periodic.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", periodic.Interval})
    periodic.EntityData.Leafs.Append("day", types.YLeaf{"Day", periodic.Day})
    periodic.EntityData.Leafs.Append("weekday", types.YLeaf{"Weekday", periodic.Weekday})
    periodic.EntityData.Leafs.Append("hour", types.YLeaf{"Hour", periodic.Hour})
    periodic.EntityData.Leafs.Append("minute", types.YLeaf{"Minute", periodic.Minute})

    periodic.EntityData.YListKeys = []string {}

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
    inventory.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/subscribe-alert-group/" + inventory.EntityData.SegmentPath
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = types.NewOrderedMap()
    inventory.EntityData.Children.Append("periodic", types.YChild{"Periodic", &inventory.Periodic})
    inventory.EntityData.Leafs = types.NewOrderedMap()
    inventory.EntityData.Leafs.Append("subscribe", types.YLeaf{"Subscribe", inventory.Subscribe})

    inventory.EntityData.YListKeys = []string {}

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
    periodic.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/subscribe-alert-group/inventory/" + periodic.EntityData.SegmentPath
    periodic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    periodic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    periodic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    periodic.EntityData.Children = types.NewOrderedMap()
    periodic.EntityData.Leafs = types.NewOrderedMap()
    periodic.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", periodic.Interval})
    periodic.EntityData.Leafs.Append("day", types.YLeaf{"Day", periodic.Day})
    periodic.EntityData.Leafs.Append("weekday", types.YLeaf{"Weekday", periodic.Weekday})
    periodic.EntityData.Leafs.Append("hour", types.YLeaf{"Hour", periodic.Hour})
    periodic.EntityData.Leafs.Append("minute", types.YLeaf{"Minute", periodic.Minute})

    periodic.EntityData.YListKeys = []string {}

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
    crash.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/subscribe-alert-group/" + crash.EntityData.SegmentPath
    crash.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crash.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crash.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crash.EntityData.Children = types.NewOrderedMap()
    crash.EntityData.Leafs = types.NewOrderedMap()
    crash.EntityData.Leafs.Append("subscribe", types.YLeaf{"Subscribe", crash.Subscribe})

    crash.EntityData.YListKeys = []string {}

    return &(crash.EntityData)
}

// CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs
// syslog info
type CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Syslog message pattern to be matched. The type is slice of
    // CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog.
    Syslog []*CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog
}

func (syslogs *CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs) GetEntityData() *types.CommonEntityData {
    syslogs.EntityData.YFilter = syslogs.YFilter
    syslogs.EntityData.YangName = "syslogs"
    syslogs.EntityData.BundleName = "cisco_ios_xr"
    syslogs.EntityData.ParentYangName = "subscribe-alert-group"
    syslogs.EntityData.SegmentPath = "syslogs"
    syslogs.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/subscribe-alert-group/" + syslogs.EntityData.SegmentPath
    syslogs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syslogs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syslogs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syslogs.EntityData.Children = types.NewOrderedMap()
    syslogs.EntityData.Children.Append("syslog", types.YChild{"Syslog", nil})
    for i := range syslogs.Syslog {
        syslogs.EntityData.Children.Append(types.GetSegmentPath(syslogs.Syslog[i]), types.YChild{"Syslog", syslogs.Syslog[i]})
    }
    syslogs.EntityData.Leafs = types.NewOrderedMap()

    syslogs.EntityData.YListKeys = []string {}

    return &(syslogs.EntityData)
}

// CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog
// Syslog message pattern to be matched
type CallHome_Profiles_Profile_SubscribeAlertGroup_Syslogs_Syslog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    syslog.EntityData.SegmentPath = "syslog" + types.AddKeyToken(syslog.SyslogPattern, "syslog-pattern")
    syslog.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/profiles/profile/subscribe-alert-group/syslogs/" + syslog.EntityData.SegmentPath
    syslog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syslog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syslog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syslog.EntityData.Children = types.NewOrderedMap()
    syslog.EntityData.Leafs = types.NewOrderedMap()
    syslog.EntityData.Leafs.Append("syslog-pattern", types.YLeaf{"SyslogPattern", syslog.SyslogPattern})
    syslog.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", syslog.Severity})

    syslog.EntityData.YListKeys = []string {"SyslogPattern"}

    return &(syslog.EntityData)
}

// CallHome_AlertGroups
// List of alert-group
type CallHome_AlertGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A specific alert-group. The type is slice of
    // CallHome_AlertGroups_AlertGroup.
    AlertGroup []*CallHome_AlertGroups_AlertGroup
}

func (alertGroups *CallHome_AlertGroups) GetEntityData() *types.CommonEntityData {
    alertGroups.EntityData.YFilter = alertGroups.YFilter
    alertGroups.EntityData.YangName = "alert-groups"
    alertGroups.EntityData.BundleName = "cisco_ios_xr"
    alertGroups.EntityData.ParentYangName = "call-home"
    alertGroups.EntityData.SegmentPath = "alert-groups"
    alertGroups.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/" + alertGroups.EntityData.SegmentPath
    alertGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alertGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alertGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alertGroups.EntityData.Children = types.NewOrderedMap()
    alertGroups.EntityData.Children.Append("alert-group", types.YChild{"AlertGroup", nil})
    for i := range alertGroups.AlertGroup {
        alertGroups.EntityData.Children.Append(types.GetSegmentPath(alertGroups.AlertGroup[i]), types.YChild{"AlertGroup", alertGroups.AlertGroup[i]})
    }
    alertGroups.EntityData.Leafs = types.NewOrderedMap()

    alertGroups.EntityData.YListKeys = []string {}

    return &(alertGroups.EntityData)
}

// CallHome_AlertGroups_AlertGroup
// A specific alert-group
type CallHome_AlertGroups_AlertGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. none. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
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
    alertGroup.EntityData.SegmentPath = "alert-group" + types.AddKeyToken(alertGroup.AlertGroupName, "alert-group-name")
    alertGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/alert-groups/" + alertGroup.EntityData.SegmentPath
    alertGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alertGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alertGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alertGroup.EntityData.Children = types.NewOrderedMap()
    alertGroup.EntityData.Leafs = types.NewOrderedMap()
    alertGroup.EntityData.Leafs.Append("alert-group-name", types.YLeaf{"AlertGroupName", alertGroup.AlertGroupName})
    alertGroup.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", alertGroup.Enable})
    alertGroup.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", alertGroup.Disable})

    alertGroup.EntityData.YListKeys = []string {"AlertGroupName"}

    return &(alertGroup.EntityData)
}

// CallHome_DataPrivacies
// Set call-home data-privacy
type CallHome_DataPrivacies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // level hostname. The type is slice of CallHome_DataPrivacies_DataPrivacy.
    DataPrivacy []*CallHome_DataPrivacies_DataPrivacy
}

func (dataPrivacies *CallHome_DataPrivacies) GetEntityData() *types.CommonEntityData {
    dataPrivacies.EntityData.YFilter = dataPrivacies.YFilter
    dataPrivacies.EntityData.YangName = "data-privacies"
    dataPrivacies.EntityData.BundleName = "cisco_ios_xr"
    dataPrivacies.EntityData.ParentYangName = "call-home"
    dataPrivacies.EntityData.SegmentPath = "data-privacies"
    dataPrivacies.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/" + dataPrivacies.EntityData.SegmentPath
    dataPrivacies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataPrivacies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataPrivacies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataPrivacies.EntityData.Children = types.NewOrderedMap()
    dataPrivacies.EntityData.Children.Append("data-privacy", types.YChild{"DataPrivacy", nil})
    for i := range dataPrivacies.DataPrivacy {
        dataPrivacies.EntityData.Children.Append(types.GetSegmentPath(dataPrivacies.DataPrivacy[i]), types.YChild{"DataPrivacy", dataPrivacies.DataPrivacy[i]})
    }
    dataPrivacies.EntityData.Leafs = types.NewOrderedMap()

    dataPrivacies.EntityData.YListKeys = []string {}

    return &(dataPrivacies.EntityData)
}

// CallHome_DataPrivacies_DataPrivacy
// level hostname
type CallHome_DataPrivacies_DataPrivacy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    dataPrivacy.EntityData.SegmentPath = "data-privacy" + types.AddKeyToken(dataPrivacy.HostName, "host-name")
    dataPrivacy.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/data-privacies/" + dataPrivacy.EntityData.SegmentPath
    dataPrivacy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataPrivacy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataPrivacy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataPrivacy.EntityData.Children = types.NewOrderedMap()
    dataPrivacy.EntityData.Leafs = types.NewOrderedMap()
    dataPrivacy.EntityData.Leafs.Append("host-name", types.YLeaf{"HostName", dataPrivacy.HostName})
    dataPrivacy.EntityData.Leafs.Append("level", types.YLeaf{"Level", dataPrivacy.Level})

    dataPrivacy.EntityData.YListKeys = []string {"HostName"}

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
    alertGroupConfig.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/" + alertGroupConfig.EntityData.SegmentPath
    alertGroupConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alertGroupConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alertGroupConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alertGroupConfig.EntityData.Children = types.NewOrderedMap()
    alertGroupConfig.EntityData.Children.Append("snapshot-commands", types.YChild{"SnapshotCommands", &alertGroupConfig.SnapshotCommands})
    alertGroupConfig.EntityData.Leafs = types.NewOrderedMap()

    alertGroupConfig.EntityData.YListKeys = []string {}

    return &(alertGroupConfig.EntityData)
}

// CallHome_AlertGroupConfig_SnapshotCommands
// snapshot for adding CLI command
type CallHome_AlertGroupConfig_SnapshotCommands struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A specific CLI cmd for snapshot. The type is slice of
    // CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand.
    SnapshotCommand []*CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand
}

func (snapshotCommands *CallHome_AlertGroupConfig_SnapshotCommands) GetEntityData() *types.CommonEntityData {
    snapshotCommands.EntityData.YFilter = snapshotCommands.YFilter
    snapshotCommands.EntityData.YangName = "snapshot-commands"
    snapshotCommands.EntityData.BundleName = "cisco_ios_xr"
    snapshotCommands.EntityData.ParentYangName = "alert-group-config"
    snapshotCommands.EntityData.SegmentPath = "snapshot-commands"
    snapshotCommands.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/alert-group-config/" + snapshotCommands.EntityData.SegmentPath
    snapshotCommands.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snapshotCommands.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snapshotCommands.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snapshotCommands.EntityData.Children = types.NewOrderedMap()
    snapshotCommands.EntityData.Children.Append("snapshot-command", types.YChild{"SnapshotCommand", nil})
    for i := range snapshotCommands.SnapshotCommand {
        snapshotCommands.EntityData.Children.Append(types.GetSegmentPath(snapshotCommands.SnapshotCommand[i]), types.YChild{"SnapshotCommand", snapshotCommands.SnapshotCommand[i]})
    }
    snapshotCommands.EntityData.Leafs = types.NewOrderedMap()

    snapshotCommands.EntityData.YListKeys = []string {}

    return &(snapshotCommands.EntityData)
}

// CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand
// A specific CLI cmd for snapshot
type CallHome_AlertGroupConfig_SnapshotCommands_SnapshotCommand struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    snapshotCommand.EntityData.SegmentPath = "snapshot-command" + types.AddKeyToken(snapshotCommand.Command, "command")
    snapshotCommand.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/alert-group-config/snapshot-commands/" + snapshotCommand.EntityData.SegmentPath
    snapshotCommand.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snapshotCommand.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snapshotCommand.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snapshotCommand.EntityData.Children = types.NewOrderedMap()
    snapshotCommand.EntityData.Leafs = types.NewOrderedMap()
    snapshotCommand.EntityData.Leafs.Append("command", types.YLeaf{"Command", snapshotCommand.Command})
    snapshotCommand.EntityData.Leafs.Append("active", types.YLeaf{"Active", snapshotCommand.Active})

    snapshotCommand.EntityData.YListKeys = []string {"Command"}

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
    authorization.EntityData.AbsolutePath = "Cisco-IOS-XR-call-home-cfg:call-home/" + authorization.EntityData.SegmentPath
    authorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorization.EntityData.Children = types.NewOrderedMap()
    authorization.EntityData.Leafs = types.NewOrderedMap()
    authorization.EntityData.Leafs.Append("username", types.YLeaf{"Username", authorization.Username})
    authorization.EntityData.Leafs.Append("active", types.YLeaf{"Active", authorization.Active})

    authorization.EntityData.YListKeys = []string {}

    return &(authorization.EntityData)
}

