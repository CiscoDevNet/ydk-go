// The MIB module to describe the OSPF Version 2
// Protocol.  Note that some objects in this MIB
// module may pose a significant security risk.
// Refer to the Security Considerations section
// in RFC 4750 for more information.
// 
// Copyright (C) The IETF Trust (2006).
// This version of this MIB module is part of
// RFC 4750;  see the RFC itself for full legal
// notices.
package ospf_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ospf_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:OSPF-MIB OSPF-MIB}", reflect.TypeOf(OSPFMIB{}))
    ydk.RegisterEntity("OSPF-MIB:OSPF-MIB", reflect.TypeOf(OSPFMIB{}))
}

// Status represents and 'disabled' indicates that it is not.
type Status string

const (
    Status_enabled Status = "enabled"

    Status_disabled Status = "disabled"
)

// OspfAuthenticationType represents The authentication type.
type OspfAuthenticationType string

const (
    OspfAuthenticationType_none OspfAuthenticationType = "none"

    OspfAuthenticationType_simplePassword OspfAuthenticationType = "simplePassword"

    OspfAuthenticationType_md5 OspfAuthenticationType = "md5"
)

// OSPFMIB
type OSPFMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ospfgeneralgroup OSPFMIB_Ospfgeneralgroup

    // Information describing the configured parameters and cumulative statistics
    // of the router's attached areas. The interfaces and virtual links are
    // configured as part of these areas.  Area 0.0.0.0, by definition, is the
    // backbone area.
    Ospfareatable OSPFMIB_Ospfareatable

    // The set of metrics that will be advertised by a default Area Border Router
    // into a stub area.
    Ospfstubareatable OSPFMIB_Ospfstubareatable

    // The OSPF Process's link state database (LSDB). The LSDB contains the link
    // state advertisements from throughout the areas that the device is attached
    // to.
    Ospflsdbtable OSPFMIB_Ospflsdbtable

    // The Address Range Table acts as an adjunct to the Area Table.  It describes
    // those Address Range Summaries that are configured to be propagated from an
    // Area to reduce the amount of information about it that is known beyond its
    // borders.  It contains a set of IP address ranges specified by an IP
    // address/IP network mask pair. For example, class B address range of X.X.X.X
    // with a network mask of 255.255.0.0 includes all IP addresses from X.X.0.0
    // to X.X.255.255.  Note that this table is obsoleted and is replaced by the
    // Area Aggregate Table.
    Ospfarearangetable OSPFMIB_Ospfarearangetable

    // The Host/Metric Table indicates what hosts are directly  attached to the
    // router, what metrics and types of service should be advertised for them,
    // and what areas they are found within.
    Ospfhosttable OSPFMIB_Ospfhosttable

    // The OSPF Interface Table describes the interfaces from the viewpoint of
    // OSPF. It augments the ipAddrTable with OSPF specific information.
    Ospfiftable OSPFMIB_Ospfiftable

    // The Metric Table describes the metrics to be advertised for a specified
    // interface at the various types of service. As such, this table is an
    // adjunct of the OSPF Interface Table.  Types of service, as defined by RFC
    // 791, have the ability to request low delay, high bandwidth, or reliable
    // linkage.  For the purposes of this specification, the measure of bandwidth:
    // Metric = referenceBandwidth / ifSpeed  is the default value. The default
    // reference bandwidth is 10^8. For multiple link interfaces, note that
    // ifSpeed is the sum of the individual link speeds.  This yields a number
    // having the following typical values:  Network Type/bit rate   Metric  >=
    // 100 MBPS                 1 Ethernet/802.3             10 E1                
    // 48 T1 (ESF)                   65 64 KBPS                    1562 56 KBPS   
    // 1785 19.2 KBPS                  5208 9.6 KBPS                   10416 
    // Routes that are not specified use the default (TOS 0) metric.  Note that
    // the default reference bandwidth can be configured using the general group
    // object ospfReferenceBandwidth.
    Ospfifmetrictable OSPFMIB_Ospfifmetrictable

    // Information about this router's virtual interfaces that the OSPF Process is
    // configured to carry on.
    Ospfvirtiftable OSPFMIB_Ospfvirtiftable

    // A table describing all non-virtual neighbors in the locality of the OSPF
    // router.
    Ospfnbrtable OSPFMIB_Ospfnbrtable

    // This table describes all virtual neighbors. Since virtual links are
    // configured in the Virtual Interface Table, this table is read-only.
    Ospfvirtnbrtable OSPFMIB_Ospfvirtnbrtable

    // The OSPF Process's external LSA link state database.  This table is
    // identical to the OSPF LSDB Table in format, but contains only external link
    // state advertisements.  The purpose is to allow external  LSAs to be
    // displayed once for the router rather than once in each non-stub area.  Note
    // that external LSAs are also in the AS-scope link state database.
    Ospfextlsdbtable OSPFMIB_Ospfextlsdbtable

    // The Area Aggregate Table acts as an adjunct to the Area Table.  It
    // describes those address aggregates that are configured to be propagated
    // from an area. Its purpose is to reduce the amount of information that is
    // known beyond an Area's borders.  It contains a set of IP address ranges
    // specified by an IP address/IP network mask pair. For example, a class B
    // address range of X.X.X.X with a network mask of 255.255.0.0 includes all IP
    // addresses from X.X.0.0 to X.X.255.255.  Note that if ranges are configured
    // such that one range subsumes another range (e.g., 10.0.0.0 mask 255.0.0.0
    // and 10.1.0.0 mask 255.255.0.0), the most specific match is the preferred
    // one.
    Ospfareaaggregatetable OSPFMIB_Ospfareaaggregatetable

    // The OSPF Process's link-local link state database for non-virtual links.
    // This table is identical to the OSPF LSDB Table in format, but contains only
    // link-local Link State Advertisements for non-virtual links.  The purpose is
    // to allow link-local LSAs to be displayed for each non-virtual interface. 
    // This table is implemented to support type-9 LSAs that are defined in 'The
    // OSPF Opaque LSA Option'.
    Ospflocallsdbtable OSPFMIB_Ospflocallsdbtable

    // The OSPF Process's link-local link state database for virtual links.  This
    // table is identical to the OSPF LSDB Table in format, but contains only
    // link-local Link State Advertisements for virtual links.  The purpose is to
    // allow link-local LSAs to be displayed for each virtual interface.  This
    // table is implemented to support type-9 LSAs that are defined in 'The OSPF
    // Opaque LSA Option'.
    Ospfvirtlocallsdbtable OSPFMIB_Ospfvirtlocallsdbtable

    // The OSPF Process's AS-scope LSA link state database. The database contains
    // the AS-scope Link State Advertisements from throughout the areas that the
    // device is attached to.  This table is identical to the OSPF LSDB Table in
    // format, but contains only AS-scope Link State Advertisements.  The purpose
    // is to allow AS-scope LSAs to be displayed once for the router rather than
    // once in each non-stub area.
    Ospfaslsdbtable OSPFMIB_Ospfaslsdbtable

    // This table maintains per-area, per-LSA-type counters.
    Ospfarealsacounttable OSPFMIB_Ospfarealsacounttable
}

func (oSPFMIB *OSPFMIB) GetFilter() yfilter.YFilter { return oSPFMIB.YFilter }

func (oSPFMIB *OSPFMIB) SetFilter(yf yfilter.YFilter) { oSPFMIB.YFilter = yf }

func (oSPFMIB *OSPFMIB) GetGoName(yname string) string {
    if yname == "ospfGeneralGroup" { return "Ospfgeneralgroup" }
    if yname == "ospfAreaTable" { return "Ospfareatable" }
    if yname == "ospfStubAreaTable" { return "Ospfstubareatable" }
    if yname == "ospfLsdbTable" { return "Ospflsdbtable" }
    if yname == "ospfAreaRangeTable" { return "Ospfarearangetable" }
    if yname == "ospfHostTable" { return "Ospfhosttable" }
    if yname == "ospfIfTable" { return "Ospfiftable" }
    if yname == "ospfIfMetricTable" { return "Ospfifmetrictable" }
    if yname == "ospfVirtIfTable" { return "Ospfvirtiftable" }
    if yname == "ospfNbrTable" { return "Ospfnbrtable" }
    if yname == "ospfVirtNbrTable" { return "Ospfvirtnbrtable" }
    if yname == "ospfExtLsdbTable" { return "Ospfextlsdbtable" }
    if yname == "ospfAreaAggregateTable" { return "Ospfareaaggregatetable" }
    if yname == "ospfLocalLsdbTable" { return "Ospflocallsdbtable" }
    if yname == "ospfVirtLocalLsdbTable" { return "Ospfvirtlocallsdbtable" }
    if yname == "ospfAsLsdbTable" { return "Ospfaslsdbtable" }
    if yname == "ospfAreaLsaCountTable" { return "Ospfarealsacounttable" }
    return ""
}

func (oSPFMIB *OSPFMIB) GetSegmentPath() string {
    return "OSPF-MIB:OSPF-MIB"
}

func (oSPFMIB *OSPFMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfGeneralGroup" {
        return &oSPFMIB.Ospfgeneralgroup
    }
    if childYangName == "ospfAreaTable" {
        return &oSPFMIB.Ospfareatable
    }
    if childYangName == "ospfStubAreaTable" {
        return &oSPFMIB.Ospfstubareatable
    }
    if childYangName == "ospfLsdbTable" {
        return &oSPFMIB.Ospflsdbtable
    }
    if childYangName == "ospfAreaRangeTable" {
        return &oSPFMIB.Ospfarearangetable
    }
    if childYangName == "ospfHostTable" {
        return &oSPFMIB.Ospfhosttable
    }
    if childYangName == "ospfIfTable" {
        return &oSPFMIB.Ospfiftable
    }
    if childYangName == "ospfIfMetricTable" {
        return &oSPFMIB.Ospfifmetrictable
    }
    if childYangName == "ospfVirtIfTable" {
        return &oSPFMIB.Ospfvirtiftable
    }
    if childYangName == "ospfNbrTable" {
        return &oSPFMIB.Ospfnbrtable
    }
    if childYangName == "ospfVirtNbrTable" {
        return &oSPFMIB.Ospfvirtnbrtable
    }
    if childYangName == "ospfExtLsdbTable" {
        return &oSPFMIB.Ospfextlsdbtable
    }
    if childYangName == "ospfAreaAggregateTable" {
        return &oSPFMIB.Ospfareaaggregatetable
    }
    if childYangName == "ospfLocalLsdbTable" {
        return &oSPFMIB.Ospflocallsdbtable
    }
    if childYangName == "ospfVirtLocalLsdbTable" {
        return &oSPFMIB.Ospfvirtlocallsdbtable
    }
    if childYangName == "ospfAsLsdbTable" {
        return &oSPFMIB.Ospfaslsdbtable
    }
    if childYangName == "ospfAreaLsaCountTable" {
        return &oSPFMIB.Ospfarealsacounttable
    }
    return nil
}

func (oSPFMIB *OSPFMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospfGeneralGroup"] = &oSPFMIB.Ospfgeneralgroup
    children["ospfAreaTable"] = &oSPFMIB.Ospfareatable
    children["ospfStubAreaTable"] = &oSPFMIB.Ospfstubareatable
    children["ospfLsdbTable"] = &oSPFMIB.Ospflsdbtable
    children["ospfAreaRangeTable"] = &oSPFMIB.Ospfarearangetable
    children["ospfHostTable"] = &oSPFMIB.Ospfhosttable
    children["ospfIfTable"] = &oSPFMIB.Ospfiftable
    children["ospfIfMetricTable"] = &oSPFMIB.Ospfifmetrictable
    children["ospfVirtIfTable"] = &oSPFMIB.Ospfvirtiftable
    children["ospfNbrTable"] = &oSPFMIB.Ospfnbrtable
    children["ospfVirtNbrTable"] = &oSPFMIB.Ospfvirtnbrtable
    children["ospfExtLsdbTable"] = &oSPFMIB.Ospfextlsdbtable
    children["ospfAreaAggregateTable"] = &oSPFMIB.Ospfareaaggregatetable
    children["ospfLocalLsdbTable"] = &oSPFMIB.Ospflocallsdbtable
    children["ospfVirtLocalLsdbTable"] = &oSPFMIB.Ospfvirtlocallsdbtable
    children["ospfAsLsdbTable"] = &oSPFMIB.Ospfaslsdbtable
    children["ospfAreaLsaCountTable"] = &oSPFMIB.Ospfarealsacounttable
    return children
}

func (oSPFMIB *OSPFMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oSPFMIB *OSPFMIB) GetBundleName() string { return "cisco_ios_xe" }

func (oSPFMIB *OSPFMIB) GetYangName() string { return "OSPF-MIB" }

func (oSPFMIB *OSPFMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (oSPFMIB *OSPFMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (oSPFMIB *OSPFMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (oSPFMIB *OSPFMIB) SetParent(parent types.Entity) { oSPFMIB.parent = parent }

func (oSPFMIB *OSPFMIB) GetParent() types.Entity { return oSPFMIB.parent }

func (oSPFMIB *OSPFMIB) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfgeneralgroup
type OSPFMIB_Ospfgeneralgroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A 32-bit integer uniquely identifying the router in the Autonomous System.
    // By convention, to ensure uniqueness, this should default to the value of
    // one of the router's IP interface addresses.  This object is persistent and
    // when written the entity SHOULD save the change to non-volatile storage. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfrouterid interface{}

    // The administrative status of OSPF in the router.  The value 'enabled'
    // denotes that the OSPF Process is active on at least one interface;
    // 'disabled' disables it on all interfaces.  This object is persistent and
    // when written the entity SHOULD save the change to non-volatile storage. The
    // type is Status.
    Ospfadminstat interface{}

    // The current version number of the OSPF protocol is 2. The type is
    // Ospfversionnumber.
    Ospfversionnumber interface{}

    // A flag to note whether this router is an Area Border Router. The type is
    // bool.
    Ospfareabdrrtrstatus interface{}

    // A flag to note whether this router is configured as an Autonomous System
    // Border Router.  This object is persistent and when written the entity
    // SHOULD save the change to non-volatile storage. The type is bool.
    Ospfasbdrrtrstatus interface{}

    // The number of external (LS type-5) link state advertisements in the link
    // state database. The type is interface{} with range: 0..4294967295.
    Ospfexternlsacount interface{}

    // The 32-bit sum of the LS checksums of the external link state
    // advertisements contained in the link state database.  This sum can be used
    // to determine if there has been a change in a router's link state database
    // and to compare the link state database of two routers.  The value should be
    // treated as unsigned when comparing two sums of checksums. The type is
    // interface{} with range: -2147483648..2147483647.
    Ospfexternlsacksumsum interface{}

    // The router's support for type-of-service routing.  This object is
    // persistent and when written the entity SHOULD save the change to
    // non-volatile storage. The type is bool.
    Ospftossupport interface{}

    // The number of new link state advertisements that have been originated. 
    // This number is incremented each time the router originates a new LSA. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ospfDiscontinuityTime. The type is interface{} with range: 0..4294967295.
    Ospforiginatenewlsas interface{}

    // The number of link state advertisements received that are determined to be
    // new instantiations. This number does not include newer instantiations of
    // self-originated link state advertisements.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ospfDiscontinuityTime. The type
    // is interface{} with range: 0..4294967295.
    Ospfrxnewlsas interface{}

    // The maximum number of non-default AS-external LSAs entries that can be
    // stored in the link state database.  If the value is -1, then there is no
    // limit.  When the number of non-default AS-external LSAs in a router's link
    // state database reaches ospfExtLsdbLimit, the router enters overflow state. 
    // The router never holds more than ospfExtLsdbLimit non-default AS-external
    // LSAs in its database.  OspfExtLsdbLimit MUST be set identically in all
    // routers attached to the OSPF backbone and/or any regular OSPF area (i.e.,
    // OSPF stub areas and NSSAs are excluded).  This object is persistent and
    // when written the entity SHOULD save the change to non-volatile storage. The
    // type is interface{} with range: -1..2147483647.
    Ospfextlsdblimit interface{}

    // A bit mask indicating whether the router is forwarding IP multicast (Class
    // D) datagrams based on the algorithms defined in the multicast extensions to
    // OSPF.  Bit 0, if set, indicates that the router can  forward IP multicast
    // datagrams in the router's directly attached areas (called intra-area
    // multicast routing).  Bit 1, if set, indicates that the router can forward
    // IP multicast datagrams between OSPF areas (called inter-area multicast
    // routing).  Bit 2, if set, indicates that the router can forward IP
    // multicast datagrams between Autonomous Systems (called inter-AS multicast
    // routing).  Only certain combinations of bit settings are allowed, namely: 0
    // (no multicast forwarding is enabled), 1 (intra-area multicasting only), 3
    // (intra-area and inter-area multicasting), 5 (intra-area and inter-AS
    // multicasting), and 7 (multicasting everywhere).  By default, no multicast
    // forwarding is enabled.  This object is persistent and when written the
    // entity SHOULD save the change to non-volatile storage. The type is
    // interface{} with range: -2147483648..2147483647.
    Ospfmulticastextensions interface{}

    // The number of seconds that, after entering OverflowState, a router will
    // attempt to leave OverflowState.  This allows the router to again originate
    // non-default AS-external LSAs.  When set to 0, the router will not leave
    // overflow state until restarted.  This object is persistent and when written
    // the entity SHOULD save the change to non-volatile storage. The type is
    // interface{} with range: 0..2147483647.
    Ospfexitoverflowinterval interface{}

    // The router's support for demand routing. This object is persistent and when
    // written the entity SHOULD save the change to non-volatile storage. The type
    // is bool.
    Ospfdemandextensions interface{}

    // Indicates metrics used to choose among multiple AS-external LSAs.  When
    // RFC1583Compatibility is set to enabled, only cost will be used when
    // choosing among multiple AS-external LSAs advertising the same destination. 
    // When RFC1583Compatibility is set to disabled, preference will be driven
    // first by type of path using cost only to break ties.  This object is
    // persistent and when written the entity SHOULD save the change to
    // non-volatile storage. The type is bool.
    Ospfrfc1583Compatibility interface{}

    // The router's support for Opaque LSA types. The type is bool.
    Ospfopaquelsasupport interface{}

    // Reference bandwidth in kilobits/second for  calculating default interface
    // metrics.  The default value is 100,000 KBPS (100 MBPS).  This object is
    // persistent and when written the entity SHOULD save the change to
    // non-volatile storage. The type is interface{} with range: 0..4294967295.
    // Units are kilobits per second.
    Ospfreferencebandwidth interface{}

    // The router's support for OSPF graceful restart. Options include: no restart
    // support, only planned restarts, or both planned and unplanned restarts. 
    // This object is persistent and when written the entity SHOULD save the
    // change to non-volatile storage. The type is Ospfrestartsupport.
    Ospfrestartsupport interface{}

    // Configured OSPF graceful restart timeout interval.  This object is
    // persistent and when written the entity SHOULD save the change to
    // non-volatile storage. The type is interface{} with range: 1..1800. Units
    // are seconds.
    Ospfrestartinterval interface{}

    // Indicates if strict LSA checking is enabled for graceful restart.  This
    // object is persistent and when written the entity SHOULD save the change to
    // non-volatile  storage. The type is bool.
    Ospfrestartstrictlsachecking interface{}

    // Current status of OSPF graceful restart. The type is Ospfrestartstatus.
    Ospfrestartstatus interface{}

    // Remaining time in current OSPF graceful restart interval. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Ospfrestartage interface{}

    // Describes the outcome of the last attempt at a graceful restart.  If the
    // value is 'none', no restart has yet been attempted.  If the value is
    // 'inProgress', a restart attempt is currently underway. The type is
    // Ospfrestartexitreason.
    Ospfrestartexitreason interface{}

    // The number of AS-scope link state advertisements in the AS-scope link state
    // database. The type is interface{} with range: 0..4294967295.
    Ospfaslsacount interface{}

    // The 32-bit unsigned sum of the LS checksums of the AS link state
    // advertisements contained in the AS-scope link state database.  This sum can
    // be used to determine if there has been a change in a router's AS-scope link
    // state database, and to compare the AS-scope link state database of two
    // routers. The type is interface{} with range: 0..4294967295.
    Ospfaslsacksumsum interface{}

    // The router's support for stub router functionality. The type is bool.
    Ospfstubroutersupport interface{}

    // This object controls the advertisement of stub router LSAs by the router. 
    // The value doNotAdvertise will result in the advertisement of a standard
    // router LSA and is the default value.  This object is persistent and when
    // written the entity SHOULD save the change to non-volatile storage. The type
    // is Ospfstubrouteradvertisement.
    Ospfstubrouteradvertisement interface{}

    // The value of sysUpTime on the most recent occasion at which any one of this
    // MIB's counters suffered a discontinuity.  If no such discontinuities have
    // occurred since the last re-initialization of the local management
    // subsystem, then this object contains a zero value. The type is interface{}
    // with range: 0..4294967295.
    Ospfdiscontinuitytime interface{}
}

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetFilter() yfilter.YFilter { return ospfgeneralgroup.YFilter }

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) SetFilter(yf yfilter.YFilter) { ospfgeneralgroup.YFilter = yf }

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetGoName(yname string) string {
    if yname == "ospfRouterId" { return "Ospfrouterid" }
    if yname == "ospfAdminStat" { return "Ospfadminstat" }
    if yname == "ospfVersionNumber" { return "Ospfversionnumber" }
    if yname == "ospfAreaBdrRtrStatus" { return "Ospfareabdrrtrstatus" }
    if yname == "ospfASBdrRtrStatus" { return "Ospfasbdrrtrstatus" }
    if yname == "ospfExternLsaCount" { return "Ospfexternlsacount" }
    if yname == "ospfExternLsaCksumSum" { return "Ospfexternlsacksumsum" }
    if yname == "ospfTOSSupport" { return "Ospftossupport" }
    if yname == "ospfOriginateNewLsas" { return "Ospforiginatenewlsas" }
    if yname == "ospfRxNewLsas" { return "Ospfrxnewlsas" }
    if yname == "ospfExtLsdbLimit" { return "Ospfextlsdblimit" }
    if yname == "ospfMulticastExtensions" { return "Ospfmulticastextensions" }
    if yname == "ospfExitOverflowInterval" { return "Ospfexitoverflowinterval" }
    if yname == "ospfDemandExtensions" { return "Ospfdemandextensions" }
    if yname == "ospfRFC1583Compatibility" { return "Ospfrfc1583Compatibility" }
    if yname == "ospfOpaqueLsaSupport" { return "Ospfopaquelsasupport" }
    if yname == "ospfReferenceBandwidth" { return "Ospfreferencebandwidth" }
    if yname == "ospfRestartSupport" { return "Ospfrestartsupport" }
    if yname == "ospfRestartInterval" { return "Ospfrestartinterval" }
    if yname == "ospfRestartStrictLsaChecking" { return "Ospfrestartstrictlsachecking" }
    if yname == "ospfRestartStatus" { return "Ospfrestartstatus" }
    if yname == "ospfRestartAge" { return "Ospfrestartage" }
    if yname == "ospfRestartExitReason" { return "Ospfrestartexitreason" }
    if yname == "ospfAsLsaCount" { return "Ospfaslsacount" }
    if yname == "ospfAsLsaCksumSum" { return "Ospfaslsacksumsum" }
    if yname == "ospfStubRouterSupport" { return "Ospfstubroutersupport" }
    if yname == "ospfStubRouterAdvertisement" { return "Ospfstubrouteradvertisement" }
    if yname == "ospfDiscontinuityTime" { return "Ospfdiscontinuitytime" }
    return ""
}

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetSegmentPath() string {
    return "ospfGeneralGroup"
}

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfRouterId"] = ospfgeneralgroup.Ospfrouterid
    leafs["ospfAdminStat"] = ospfgeneralgroup.Ospfadminstat
    leafs["ospfVersionNumber"] = ospfgeneralgroup.Ospfversionnumber
    leafs["ospfAreaBdrRtrStatus"] = ospfgeneralgroup.Ospfareabdrrtrstatus
    leafs["ospfASBdrRtrStatus"] = ospfgeneralgroup.Ospfasbdrrtrstatus
    leafs["ospfExternLsaCount"] = ospfgeneralgroup.Ospfexternlsacount
    leafs["ospfExternLsaCksumSum"] = ospfgeneralgroup.Ospfexternlsacksumsum
    leafs["ospfTOSSupport"] = ospfgeneralgroup.Ospftossupport
    leafs["ospfOriginateNewLsas"] = ospfgeneralgroup.Ospforiginatenewlsas
    leafs["ospfRxNewLsas"] = ospfgeneralgroup.Ospfrxnewlsas
    leafs["ospfExtLsdbLimit"] = ospfgeneralgroup.Ospfextlsdblimit
    leafs["ospfMulticastExtensions"] = ospfgeneralgroup.Ospfmulticastextensions
    leafs["ospfExitOverflowInterval"] = ospfgeneralgroup.Ospfexitoverflowinterval
    leafs["ospfDemandExtensions"] = ospfgeneralgroup.Ospfdemandextensions
    leafs["ospfRFC1583Compatibility"] = ospfgeneralgroup.Ospfrfc1583Compatibility
    leafs["ospfOpaqueLsaSupport"] = ospfgeneralgroup.Ospfopaquelsasupport
    leafs["ospfReferenceBandwidth"] = ospfgeneralgroup.Ospfreferencebandwidth
    leafs["ospfRestartSupport"] = ospfgeneralgroup.Ospfrestartsupport
    leafs["ospfRestartInterval"] = ospfgeneralgroup.Ospfrestartinterval
    leafs["ospfRestartStrictLsaChecking"] = ospfgeneralgroup.Ospfrestartstrictlsachecking
    leafs["ospfRestartStatus"] = ospfgeneralgroup.Ospfrestartstatus
    leafs["ospfRestartAge"] = ospfgeneralgroup.Ospfrestartage
    leafs["ospfRestartExitReason"] = ospfgeneralgroup.Ospfrestartexitreason
    leafs["ospfAsLsaCount"] = ospfgeneralgroup.Ospfaslsacount
    leafs["ospfAsLsaCksumSum"] = ospfgeneralgroup.Ospfaslsacksumsum
    leafs["ospfStubRouterSupport"] = ospfgeneralgroup.Ospfstubroutersupport
    leafs["ospfStubRouterAdvertisement"] = ospfgeneralgroup.Ospfstubrouteradvertisement
    leafs["ospfDiscontinuityTime"] = ospfgeneralgroup.Ospfdiscontinuitytime
    return leafs
}

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetBundleName() string { return "cisco_ios_xe" }

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetYangName() string { return "ospfGeneralGroup" }

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) SetParent(parent types.Entity) { ospfgeneralgroup.parent = parent }

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetParent() types.Entity { return ospfgeneralgroup.parent }

func (ospfgeneralgroup *OSPFMIB_Ospfgeneralgroup) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason represents a restart attempt is currently underway.
type OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason string

const (
    OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason_none OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason = "none"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason_inProgress OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason = "inProgress"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason_completed OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason = "completed"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason_timedOut OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason = "timedOut"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason_topologyChanged OSPFMIB_Ospfgeneralgroup_Ospfrestartexitreason = "topologyChanged"
)

// OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus represents Current status of OSPF graceful restart.
type OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus string

const (
    OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus_notRestarting OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus = "notRestarting"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus_plannedRestart OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus = "plannedRestart"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus_unplannedRestart OSPFMIB_Ospfgeneralgroup_Ospfrestartstatus = "unplannedRestart"
)

// OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport represents storage.
type OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport string

const (
    OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport_none OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport = "none"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport_plannedOnly OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport = "plannedOnly"

    OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport_plannedAndUnplanned OSPFMIB_Ospfgeneralgroup_Ospfrestartsupport = "plannedAndUnplanned"
)

// OSPFMIB_Ospfgeneralgroup_Ospfstubrouteradvertisement represents storage.
type OSPFMIB_Ospfgeneralgroup_Ospfstubrouteradvertisement string

const (
    OSPFMIB_Ospfgeneralgroup_Ospfstubrouteradvertisement_doNotAdvertise OSPFMIB_Ospfgeneralgroup_Ospfstubrouteradvertisement = "doNotAdvertise"

    OSPFMIB_Ospfgeneralgroup_Ospfstubrouteradvertisement_advertise OSPFMIB_Ospfgeneralgroup_Ospfstubrouteradvertisement = "advertise"
)

// OSPFMIB_Ospfgeneralgroup_Ospfversionnumber represents The current version number of the OSPF protocol is 2.
type OSPFMIB_Ospfgeneralgroup_Ospfversionnumber string

const (
    OSPFMIB_Ospfgeneralgroup_Ospfversionnumber_version2 OSPFMIB_Ospfgeneralgroup_Ospfversionnumber = "version2"
)

// OSPFMIB_Ospfareatable
// Information describing the configured parameters and
// cumulative statistics of the router's attached areas.
// The interfaces and virtual links are configured
// as part of these areas.  Area 0.0.0.0, by definition,
// is the backbone area.
type OSPFMIB_Ospfareatable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information describing the configured parameters and cumulative statistics
    // of one of the router's attached areas. The interfaces and virtual links are
    // configured as part of these areas.  Area 0.0.0.0, by definition, is the
    // backbone area.  Information in this table is persistent and when this
    // object is written the entity SHOULD save the change to non-volatile
    // storage. The type is slice of OSPFMIB_Ospfareatable_Ospfareaentry.
    Ospfareaentry []OSPFMIB_Ospfareatable_Ospfareaentry
}

func (ospfareatable *OSPFMIB_Ospfareatable) GetFilter() yfilter.YFilter { return ospfareatable.YFilter }

func (ospfareatable *OSPFMIB_Ospfareatable) SetFilter(yf yfilter.YFilter) { ospfareatable.YFilter = yf }

func (ospfareatable *OSPFMIB_Ospfareatable) GetGoName(yname string) string {
    if yname == "ospfAreaEntry" { return "Ospfareaentry" }
    return ""
}

func (ospfareatable *OSPFMIB_Ospfareatable) GetSegmentPath() string {
    return "ospfAreaTable"
}

func (ospfareatable *OSPFMIB_Ospfareatable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfAreaEntry" {
        for _, c := range ospfareatable.Ospfareaentry {
            if ospfareatable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfareatable_Ospfareaentry{}
        ospfareatable.Ospfareaentry = append(ospfareatable.Ospfareaentry, child)
        return &ospfareatable.Ospfareaentry[len(ospfareatable.Ospfareaentry)-1]
    }
    return nil
}

func (ospfareatable *OSPFMIB_Ospfareatable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfareatable.Ospfareaentry {
        children[ospfareatable.Ospfareaentry[i].GetSegmentPath()] = &ospfareatable.Ospfareaentry[i]
    }
    return children
}

func (ospfareatable *OSPFMIB_Ospfareatable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfareatable *OSPFMIB_Ospfareatable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfareatable *OSPFMIB_Ospfareatable) GetYangName() string { return "ospfAreaTable" }

func (ospfareatable *OSPFMIB_Ospfareatable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfareatable *OSPFMIB_Ospfareatable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfareatable *OSPFMIB_Ospfareatable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfareatable *OSPFMIB_Ospfareatable) SetParent(parent types.Entity) { ospfareatable.parent = parent }

func (ospfareatable *OSPFMIB_Ospfareatable) GetParent() types.Entity { return ospfareatable.parent }

func (ospfareatable *OSPFMIB_Ospfareatable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfareatable_Ospfareaentry
// Information describing the configured parameters and
// cumulative statistics of one of the router's attached areas.
// The interfaces and virtual links are configured as part of
// these areas.  Area 0.0.0.0, by definition, is the backbone
// area.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfareatable_Ospfareaentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A 32-bit integer uniquely identifying an area.
    // Area ID 0.0.0.0 is used for the OSPF backbone. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfareaid interface{}

    // The authentication type specified for an area. The type is
    // OspfAuthenticationType.
    Ospfauthtype interface{}

    // Indicates if an area is a stub area, NSSA, or standard area.  Type-5
    // AS-external LSAs and type-11 Opaque LSAs are not imported into stub areas
    // or NSSAs.  NSSAs import AS-external data as type-7 LSAs. The type is
    // Ospfimportasextern.
    Ospfimportasextern interface{}

    // The number of times that the intra-area route table has been calculated
    // using this area's link state database.  This is typically done using
    // Dijkstra's algorithm.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, and at other times as
    // indicated by the value of ospfDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    Ospfspfruns interface{}

    // The total number of Area Border Routers reachable within this area.  This
    // is initially zero and is calculated in each Shortest Path First (SPF) pass.
    // The type is interface{} with range: 0..4294967295.
    Ospfareabdrrtrcount interface{}

    // The total number of Autonomous System Border Routers reachable within this
    // area.  This is initially zero and is calculated in each SPF pass. The type
    // is interface{} with range: 0..4294967295.
    Ospfasbdrrtrcount interface{}

    // The total number of link state advertisements in this area's link state
    // database, excluding AS-external LSAs. The type is interface{} with range:
    // 0..4294967295.
    Ospfarealsacount interface{}

    // The 32-bit sum of the link state advertisements' LS checksums contained in
    // this area's link state database.  This sum excludes external (LS type-5)
    // link state advertisements. The sum can be used to determine if there has
    // been a change in a router's link state database, and to compare the link
    // state database of two routers.  The value should be treated as unsigned
    // when comparing two sums of checksums. The type is interface{} with range:
    // -2147483648..2147483647.
    Ospfarealsacksumsum interface{}

    // The variable ospfAreaSummary controls the import of summary LSAs into stub
    // and NSSA areas. It has no effect on other areas.  If it is noAreaSummary,
    // the router will not originate summary LSAs into the stub or NSSA area. It
    // will rely entirely on its default route.  If it is sendAreaSummary, the
    // router will both summarize and propagate summary LSAs. The type is
    // Ospfareasummary.
    Ospfareasummary interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfareastatus interface{}

    // Indicates an NSSA border router's ability to perform NSSA translation of
    // type-7 LSAs into type-5 LSAs. The type is Ospfareanssatranslatorrole.
    Ospfareanssatranslatorrole interface{}

    // Indicates if and how an NSSA border router is performing NSSA translation
    // of type-7 LSAs into type-5  LSAs.  When this object is set to enabled, the
    // NSSA Border router's OspfAreaNssaExtTranslatorRole has been set to always. 
    // When this object is set to elected, a candidate NSSA Border router is
    // Translating type-7 LSAs into type-5. When this object is set to disabled, a
    // candidate NSSA border router is NOT translating type-7 LSAs into type-5.
    // The type is Ospfareanssatranslatorstate.
    Ospfareanssatranslatorstate interface{}

    // The number of seconds after an elected translator determines its services
    // are no longer required, that it should continue to perform its translation
    // duties. The type is interface{} with range: 0..2147483647. Units are
    // seconds.
    Ospfareanssatranslatorstabilityinterval interface{}

    // Indicates the number of translator state changes that have occurred since
    // the last boot-up.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ospfDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    Ospfareanssatranslatorevents interface{}

    // The total number of Opaque Area and AS link-state  advertisements in the
    // link state database of this area. The type is interface{} with range:
    // 0..4294967295.
    Cospfopaquearealsacount interface{}

    // The 32-bit unsigned sum of the Opaque Area and AS  link-state
    // advertisements' LS checksums contained  link state database of this area. 
    // The sum can be  used to determine if there has been a change in the  link
    // state database for Opaque Area and AS link-state advertisements. The type
    // is interface{} with range: 0..4294967295.
    Cospfopaquearealsacksumsum interface{}

    // Indicates an NSSA Border router's ability to perform NSSA translation of
    // type-7 LSAs into type-5 LSAs. The type is Cospfareanssatranslatorrole.
    Cospfareanssatranslatorrole interface{}

    // Indicates if and how an NSSA Border router is performing NSSA translation
    // of type-7 LSAs into type-5 LSAs. When this object set to enabled, the NSSA
    // Border router's cospfAreaNssaExtTranslatorRole has been set to always. When
    // this object is set to elected, a candidate NSSA Border router is
    // Translating type-7 LSAs into type-5. When this object is set to disabled, a
    // candidate NSSA Border router is NOT translating type-7 LSAs into type-5.
    // The type is Cospfareanssatranslatorstate.
    Cospfareanssatranslatorstate interface{}

    // Indicates the number of Translator State changes that have occurred since
    // the last boot-up. The type is interface{} with range: 0..4294967295.
    Cospfareanssatranslatorevents interface{}
}

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetFilter() yfilter.YFilter { return ospfareaentry.YFilter }

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) SetFilter(yf yfilter.YFilter) { ospfareaentry.YFilter = yf }

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetGoName(yname string) string {
    if yname == "ospfAreaId" { return "Ospfareaid" }
    if yname == "ospfAuthType" { return "Ospfauthtype" }
    if yname == "ospfImportAsExtern" { return "Ospfimportasextern" }
    if yname == "ospfSpfRuns" { return "Ospfspfruns" }
    if yname == "ospfAreaBdrRtrCount" { return "Ospfareabdrrtrcount" }
    if yname == "ospfAsBdrRtrCount" { return "Ospfasbdrrtrcount" }
    if yname == "ospfAreaLsaCount" { return "Ospfarealsacount" }
    if yname == "ospfAreaLsaCksumSum" { return "Ospfarealsacksumsum" }
    if yname == "ospfAreaSummary" { return "Ospfareasummary" }
    if yname == "ospfAreaStatus" { return "Ospfareastatus" }
    if yname == "ospfAreaNssaTranslatorRole" { return "Ospfareanssatranslatorrole" }
    if yname == "ospfAreaNssaTranslatorState" { return "Ospfareanssatranslatorstate" }
    if yname == "ospfAreaNssaTranslatorStabilityInterval" { return "Ospfareanssatranslatorstabilityinterval" }
    if yname == "ospfAreaNssaTranslatorEvents" { return "Ospfareanssatranslatorevents" }
    if yname == "cospfOpaqueAreaLsaCount" { return "Cospfopaquearealsacount" }
    if yname == "cospfOpaqueAreaLsaCksumSum" { return "Cospfopaquearealsacksumsum" }
    if yname == "cospfAreaNssaTranslatorRole" { return "Cospfareanssatranslatorrole" }
    if yname == "cospfAreaNssaTranslatorState" { return "Cospfareanssatranslatorstate" }
    if yname == "cospfAreaNssaTranslatorEvents" { return "Cospfareanssatranslatorevents" }
    return ""
}

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetSegmentPath() string {
    return "ospfAreaEntry" + "[ospfAreaId='" + fmt.Sprintf("%v", ospfareaentry.Ospfareaid) + "']"
}

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfAreaId"] = ospfareaentry.Ospfareaid
    leafs["ospfAuthType"] = ospfareaentry.Ospfauthtype
    leafs["ospfImportAsExtern"] = ospfareaentry.Ospfimportasextern
    leafs["ospfSpfRuns"] = ospfareaentry.Ospfspfruns
    leafs["ospfAreaBdrRtrCount"] = ospfareaentry.Ospfareabdrrtrcount
    leafs["ospfAsBdrRtrCount"] = ospfareaentry.Ospfasbdrrtrcount
    leafs["ospfAreaLsaCount"] = ospfareaentry.Ospfarealsacount
    leafs["ospfAreaLsaCksumSum"] = ospfareaentry.Ospfarealsacksumsum
    leafs["ospfAreaSummary"] = ospfareaentry.Ospfareasummary
    leafs["ospfAreaStatus"] = ospfareaentry.Ospfareastatus
    leafs["ospfAreaNssaTranslatorRole"] = ospfareaentry.Ospfareanssatranslatorrole
    leafs["ospfAreaNssaTranslatorState"] = ospfareaentry.Ospfareanssatranslatorstate
    leafs["ospfAreaNssaTranslatorStabilityInterval"] = ospfareaentry.Ospfareanssatranslatorstabilityinterval
    leafs["ospfAreaNssaTranslatorEvents"] = ospfareaentry.Ospfareanssatranslatorevents
    leafs["cospfOpaqueAreaLsaCount"] = ospfareaentry.Cospfopaquearealsacount
    leafs["cospfOpaqueAreaLsaCksumSum"] = ospfareaentry.Cospfopaquearealsacksumsum
    leafs["cospfAreaNssaTranslatorRole"] = ospfareaentry.Cospfareanssatranslatorrole
    leafs["cospfAreaNssaTranslatorState"] = ospfareaentry.Cospfareanssatranslatorstate
    leafs["cospfAreaNssaTranslatorEvents"] = ospfareaentry.Cospfareanssatranslatorevents
    return leafs
}

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetYangName() string { return "ospfAreaEntry" }

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) SetParent(parent types.Entity) { ospfareaentry.parent = parent }

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetParent() types.Entity { return ospfareaentry.parent }

func (ospfareaentry *OSPFMIB_Ospfareatable_Ospfareaentry) GetParentYangName() string { return "ospfAreaTable" }

// OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorrole represents type-5 LSAs.
type OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorrole string

const (
    OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorrole_always OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorrole = "always"

    OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorrole_candidate OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorrole = "candidate"
)

// OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate represents Border router is NOT translating type-7 LSAs into type-5.
type OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate string

const (
    OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate_enabled OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate = "enabled"

    OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate_elected OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate = "elected"

    OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate_disabled OSPFMIB_Ospfareatable_Ospfareaentry_Cospfareanssatranslatorstate = "disabled"
)

// OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorrole represents type-5 LSAs.
type OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorrole string

const (
    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorrole_always OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorrole = "always"

    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorrole_candidate OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorrole = "candidate"
)

// OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate represents border router is NOT translating type-7 LSAs into type-5.
type OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate string

const (
    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate_enabled OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate = "enabled"

    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate_elected OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate = "elected"

    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate_disabled OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareanssatranslatorstate = "disabled"
)

// OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareasummary represents summarize and propagate summary LSAs.
type OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareasummary string

const (
    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareasummary_noAreaSummary OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareasummary = "noAreaSummary"

    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareasummary_sendAreaSummary OSPFMIB_Ospfareatable_Ospfareaentry_Ospfareasummary = "sendAreaSummary"
)

// OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern represents AS-external data as type-7 LSAs
type OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern string

const (
    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern_importExternal OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern = "importExternal"

    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern_importNoExternal OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern = "importNoExternal"

    OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern_importNssa OSPFMIB_Ospfareatable_Ospfareaentry_Ospfimportasextern = "importNssa"
)

// OSPFMIB_Ospfstubareatable
// The set of metrics that will be advertised
// by a default Area Border Router into a stub area.
type OSPFMIB_Ospfstubareatable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The metric for a given Type of Service that will be advertised by a default
    // Area Border Router into a stub area.  Information in this table is
    // persistent and when this object is written the entity SHOULD save the
    // change to non-volatile storage. The type is slice of
    // OSPFMIB_Ospfstubareatable_Ospfstubareaentry.
    Ospfstubareaentry []OSPFMIB_Ospfstubareatable_Ospfstubareaentry
}

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetFilter() yfilter.YFilter { return ospfstubareatable.YFilter }

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) SetFilter(yf yfilter.YFilter) { ospfstubareatable.YFilter = yf }

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetGoName(yname string) string {
    if yname == "ospfStubAreaEntry" { return "Ospfstubareaentry" }
    return ""
}

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetSegmentPath() string {
    return "ospfStubAreaTable"
}

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfStubAreaEntry" {
        for _, c := range ospfstubareatable.Ospfstubareaentry {
            if ospfstubareatable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfstubareatable_Ospfstubareaentry{}
        ospfstubareatable.Ospfstubareaentry = append(ospfstubareatable.Ospfstubareaentry, child)
        return &ospfstubareatable.Ospfstubareaentry[len(ospfstubareatable.Ospfstubareaentry)-1]
    }
    return nil
}

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfstubareatable.Ospfstubareaentry {
        children[ospfstubareatable.Ospfstubareaentry[i].GetSegmentPath()] = &ospfstubareatable.Ospfstubareaentry[i]
    }
    return children
}

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetYangName() string { return "ospfStubAreaTable" }

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) SetParent(parent types.Entity) { ospfstubareatable.parent = parent }

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetParent() types.Entity { return ospfstubareatable.parent }

func (ospfstubareatable *OSPFMIB_Ospfstubareatable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfstubareatable_Ospfstubareaentry
// The metric for a given Type of Service that
// will be advertised by a default Area Border
// Router into a stub area.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfstubareatable_Ospfstubareaentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The 32-bit identifier for the stub area.  On
    // creation, this can be derived from the instance. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfstubareaid interface{}

    // This attribute is a key. The Type of Service associated with the metric. 
    // On creation, this can be derived from  the instance. The type is
    // interface{} with range: 0..30.
    Ospfstubtos interface{}

    // The metric value applied at the indicated Type of Service.  By default,
    // this equals the least metric at the Type of Service among the interfaces to
    // other areas. The type is interface{} with range: 0..16777215.
    Ospfstubmetric interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfstubstatus interface{}

    // This variable displays the type of metric advertised as a default route.
    // The type is Ospfstubmetrictype.
    Ospfstubmetrictype interface{}
}

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetFilter() yfilter.YFilter { return ospfstubareaentry.YFilter }

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) SetFilter(yf yfilter.YFilter) { ospfstubareaentry.YFilter = yf }

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetGoName(yname string) string {
    if yname == "ospfStubAreaId" { return "Ospfstubareaid" }
    if yname == "ospfStubTOS" { return "Ospfstubtos" }
    if yname == "ospfStubMetric" { return "Ospfstubmetric" }
    if yname == "ospfStubStatus" { return "Ospfstubstatus" }
    if yname == "ospfStubMetricType" { return "Ospfstubmetrictype" }
    return ""
}

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetSegmentPath() string {
    return "ospfStubAreaEntry" + "[ospfStubAreaId='" + fmt.Sprintf("%v", ospfstubareaentry.Ospfstubareaid) + "']" + "[ospfStubTOS='" + fmt.Sprintf("%v", ospfstubareaentry.Ospfstubtos) + "']"
}

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfStubAreaId"] = ospfstubareaentry.Ospfstubareaid
    leafs["ospfStubTOS"] = ospfstubareaentry.Ospfstubtos
    leafs["ospfStubMetric"] = ospfstubareaentry.Ospfstubmetric
    leafs["ospfStubStatus"] = ospfstubareaentry.Ospfstubstatus
    leafs["ospfStubMetricType"] = ospfstubareaentry.Ospfstubmetrictype
    return leafs
}

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetYangName() string { return "ospfStubAreaEntry" }

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) SetParent(parent types.Entity) { ospfstubareaentry.parent = parent }

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetParent() types.Entity { return ospfstubareaentry.parent }

func (ospfstubareaentry *OSPFMIB_Ospfstubareatable_Ospfstubareaentry) GetParentYangName() string { return "ospfStubAreaTable" }

// OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype represents advertised as a default route.
type OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype string

const (
    OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype_ospfMetric OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype = "ospfMetric"

    OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype_comparableCost OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype = "comparableCost"

    OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype_nonComparable OSPFMIB_Ospfstubareatable_Ospfstubareaentry_Ospfstubmetrictype = "nonComparable"
)

// OSPFMIB_Ospflsdbtable
// The OSPF Process's link state database (LSDB).
// The LSDB contains the link state advertisements
// from throughout the areas that the device is attached to.
type OSPFMIB_Ospflsdbtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_Ospflsdbtable_Ospflsdbentry.
    Ospflsdbentry []OSPFMIB_Ospflsdbtable_Ospflsdbentry
}

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetFilter() yfilter.YFilter { return ospflsdbtable.YFilter }

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) SetFilter(yf yfilter.YFilter) { ospflsdbtable.YFilter = yf }

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetGoName(yname string) string {
    if yname == "ospfLsdbEntry" { return "Ospflsdbentry" }
    return ""
}

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetSegmentPath() string {
    return "ospfLsdbTable"
}

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfLsdbEntry" {
        for _, c := range ospflsdbtable.Ospflsdbentry {
            if ospflsdbtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospflsdbtable_Ospflsdbentry{}
        ospflsdbtable.Ospflsdbentry = append(ospflsdbtable.Ospflsdbentry, child)
        return &ospflsdbtable.Ospflsdbentry[len(ospflsdbtable.Ospflsdbentry)-1]
    }
    return nil
}

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospflsdbtable.Ospflsdbentry {
        children[ospflsdbtable.Ospflsdbentry[i].GetSegmentPath()] = &ospflsdbtable.Ospflsdbentry[i]
    }
    return children
}

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetBundleName() string { return "cisco_ios_xe" }

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetYangName() string { return "ospfLsdbTable" }

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) SetParent(parent types.Entity) { ospflsdbtable.parent = parent }

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetParent() types.Entity { return ospflsdbtable.parent }

func (ospflsdbtable *OSPFMIB_Ospflsdbtable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospflsdbtable_Ospflsdbentry
// A single link state advertisement.
type OSPFMIB_Ospflsdbtable_Ospflsdbentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The 32-bit identifier of the area from which the
    // LSA was received. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospflsdbareaid interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format.  Note: External link
    // state advertisements are permitted for backward compatibility, but should
    // be displayed in the ospfAsLsdbTable rather than here. The type is
    // Ospflsdbtype.
    Ospflsdbtype interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing either a Router ID or an IP address; it identifies the piece of
    // the routing domain that is being described by the advertisement. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospflsdblsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospflsdbrouterid interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate Link State Advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    Ospflsdbsequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    Ospflsdbage interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless  datagrams;
    // it is commonly referred to as the Fletcher checksum. The type is
    // interface{} with range: -2147483648..2147483647.
    Ospflsdbchecksum interface{}

    // The entire link state advertisement, including its header.  Note that for
    // variable length LSAs, SNMP agents may not be able to return the largest
    // string size. The type is string with length: 1..65535.
    Ospflsdbadvertisement interface{}
}

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetFilter() yfilter.YFilter { return ospflsdbentry.YFilter }

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) SetFilter(yf yfilter.YFilter) { ospflsdbentry.YFilter = yf }

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetGoName(yname string) string {
    if yname == "ospfLsdbAreaId" { return "Ospflsdbareaid" }
    if yname == "ospfLsdbType" { return "Ospflsdbtype" }
    if yname == "ospfLsdbLsid" { return "Ospflsdblsid" }
    if yname == "ospfLsdbRouterId" { return "Ospflsdbrouterid" }
    if yname == "ospfLsdbSequence" { return "Ospflsdbsequence" }
    if yname == "ospfLsdbAge" { return "Ospflsdbage" }
    if yname == "ospfLsdbChecksum" { return "Ospflsdbchecksum" }
    if yname == "ospfLsdbAdvertisement" { return "Ospflsdbadvertisement" }
    return ""
}

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetSegmentPath() string {
    return "ospfLsdbEntry" + "[ospfLsdbAreaId='" + fmt.Sprintf("%v", ospflsdbentry.Ospflsdbareaid) + "']" + "[ospfLsdbType='" + fmt.Sprintf("%v", ospflsdbentry.Ospflsdbtype) + "']" + "[ospfLsdbLsid='" + fmt.Sprintf("%v", ospflsdbentry.Ospflsdblsid) + "']" + "[ospfLsdbRouterId='" + fmt.Sprintf("%v", ospflsdbentry.Ospflsdbrouterid) + "']"
}

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfLsdbAreaId"] = ospflsdbentry.Ospflsdbareaid
    leafs["ospfLsdbType"] = ospflsdbentry.Ospflsdbtype
    leafs["ospfLsdbLsid"] = ospflsdbentry.Ospflsdblsid
    leafs["ospfLsdbRouterId"] = ospflsdbentry.Ospflsdbrouterid
    leafs["ospfLsdbSequence"] = ospflsdbentry.Ospflsdbsequence
    leafs["ospfLsdbAge"] = ospflsdbentry.Ospflsdbage
    leafs["ospfLsdbChecksum"] = ospflsdbentry.Ospflsdbchecksum
    leafs["ospfLsdbAdvertisement"] = ospflsdbentry.Ospflsdbadvertisement
    return leafs
}

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetYangName() string { return "ospfLsdbEntry" }

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) SetParent(parent types.Entity) { ospflsdbentry.parent = parent }

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetParent() types.Entity { return ospflsdbentry.parent }

func (ospflsdbentry *OSPFMIB_Ospflsdbtable_Ospflsdbentry) GetParentYangName() string { return "ospfLsdbTable" }

// OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype represents in the ospfAsLsdbTable rather than here.
type OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype string

const (
    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_routerLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "routerLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_networkLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "networkLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_summaryLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "summaryLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_asSummaryLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "asSummaryLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_asExternalLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "asExternalLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_multicastLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "multicastLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_nssaExternalLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "nssaExternalLink"

    OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype_areaOpaqueLink OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbtype = "areaOpaqueLink"
)

// OSPFMIB_Ospfarearangetable
// The Address Range Table acts as an adjunct to the Area
// Table.  It describes those Address Range Summaries that
// are configured to be propagated from an Area to reduce
// the amount of information about it that is known beyond
// its borders.  It contains a set of IP address ranges
// specified by an IP address/IP network mask pair.
// For example, class B address range of X.X.X.X
// with a network mask of 255.255.0.0 includes all IP
// addresses from X.X.0.0 to X.X.255.255.
// 
// Note that this table is obsoleted and is replaced
// by the Area Aggregate Table.
type OSPFMIB_Ospfarearangetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single area address range.  Information in this table is persistent and
    // when this object is written the entity SHOULD save the change to
    // non-volatile storage. The type is slice of
    // OSPFMIB_Ospfarearangetable_Ospfarearangeentry.
    Ospfarearangeentry []OSPFMIB_Ospfarearangetable_Ospfarearangeentry
}

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetFilter() yfilter.YFilter { return ospfarearangetable.YFilter }

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) SetFilter(yf yfilter.YFilter) { ospfarearangetable.YFilter = yf }

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetGoName(yname string) string {
    if yname == "ospfAreaRangeEntry" { return "Ospfarearangeentry" }
    return ""
}

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetSegmentPath() string {
    return "ospfAreaRangeTable"
}

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfAreaRangeEntry" {
        for _, c := range ospfarearangetable.Ospfarearangeentry {
            if ospfarearangetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfarearangetable_Ospfarearangeentry{}
        ospfarearangetable.Ospfarearangeentry = append(ospfarearangetable.Ospfarearangeentry, child)
        return &ospfarearangetable.Ospfarearangeentry[len(ospfarearangetable.Ospfarearangeentry)-1]
    }
    return nil
}

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfarearangetable.Ospfarearangeentry {
        children[ospfarearangetable.Ospfarearangeentry[i].GetSegmentPath()] = &ospfarearangetable.Ospfarearangeentry[i]
    }
    return children
}

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetYangName() string { return "ospfAreaRangeTable" }

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) SetParent(parent types.Entity) { ospfarearangetable.parent = parent }

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetParent() types.Entity { return ospfarearangetable.parent }

func (ospfarearangetable *OSPFMIB_Ospfarearangetable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfarearangetable_Ospfarearangeentry
// A single area address range.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfarearangetable_Ospfarearangeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The area that the address range is to be found
    // within. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfarearangeareaid interface{}

    // This attribute is a key. The IP address of the net or subnet indicated by
    // the range. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfarearangenet interface{}

    // The subnet mask that pertains to the net or subnet. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfarearangemask interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfarearangestatus interface{}

    // Subnets subsumed by ranges either trigger the advertisement of the
    // indicated summary (advertiseMatching) or result in the subnet's not being
    // advertised at all outside the area. The type is Ospfarearangeeffect.
    Ospfarearangeeffect interface{}
}

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetFilter() yfilter.YFilter { return ospfarearangeentry.YFilter }

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) SetFilter(yf yfilter.YFilter) { ospfarearangeentry.YFilter = yf }

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetGoName(yname string) string {
    if yname == "ospfAreaRangeAreaId" { return "Ospfarearangeareaid" }
    if yname == "ospfAreaRangeNet" { return "Ospfarearangenet" }
    if yname == "ospfAreaRangeMask" { return "Ospfarearangemask" }
    if yname == "ospfAreaRangeStatus" { return "Ospfarearangestatus" }
    if yname == "ospfAreaRangeEffect" { return "Ospfarearangeeffect" }
    return ""
}

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetSegmentPath() string {
    return "ospfAreaRangeEntry" + "[ospfAreaRangeAreaId='" + fmt.Sprintf("%v", ospfarearangeentry.Ospfarearangeareaid) + "']" + "[ospfAreaRangeNet='" + fmt.Sprintf("%v", ospfarearangeentry.Ospfarearangenet) + "']"
}

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfAreaRangeAreaId"] = ospfarearangeentry.Ospfarearangeareaid
    leafs["ospfAreaRangeNet"] = ospfarearangeentry.Ospfarearangenet
    leafs["ospfAreaRangeMask"] = ospfarearangeentry.Ospfarearangemask
    leafs["ospfAreaRangeStatus"] = ospfarearangeentry.Ospfarearangestatus
    leafs["ospfAreaRangeEffect"] = ospfarearangeentry.Ospfarearangeeffect
    return leafs
}

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetYangName() string { return "ospfAreaRangeEntry" }

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) SetParent(parent types.Entity) { ospfarearangeentry.parent = parent }

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetParent() types.Entity { return ospfarearangeentry.parent }

func (ospfarearangeentry *OSPFMIB_Ospfarearangetable_Ospfarearangeentry) GetParentYangName() string { return "ospfAreaRangeTable" }

// OSPFMIB_Ospfarearangetable_Ospfarearangeentry_Ospfarearangeeffect represents being advertised at all outside the area.
type OSPFMIB_Ospfarearangetable_Ospfarearangeentry_Ospfarearangeeffect string

const (
    OSPFMIB_Ospfarearangetable_Ospfarearangeentry_Ospfarearangeeffect_advertiseMatching OSPFMIB_Ospfarearangetable_Ospfarearangeentry_Ospfarearangeeffect = "advertiseMatching"

    OSPFMIB_Ospfarearangetable_Ospfarearangeentry_Ospfarearangeeffect_doNotAdvertiseMatching OSPFMIB_Ospfarearangetable_Ospfarearangeentry_Ospfarearangeeffect = "doNotAdvertiseMatching"
)

// OSPFMIB_Ospfhosttable
// The Host/Metric Table indicates what hosts are directly
// 
// attached to the router, what metrics and types
// of service should be advertised for them,
// and what areas they are found within.
type OSPFMIB_Ospfhosttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A metric to be advertised, for a given type of service, when a given host
    // is reachable.  Information in this table is persistent and when this object
    // is written the entity SHOULD save the change to non-volatile storage. The
    // type is slice of OSPFMIB_Ospfhosttable_Ospfhostentry.
    Ospfhostentry []OSPFMIB_Ospfhosttable_Ospfhostentry
}

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetFilter() yfilter.YFilter { return ospfhosttable.YFilter }

func (ospfhosttable *OSPFMIB_Ospfhosttable) SetFilter(yf yfilter.YFilter) { ospfhosttable.YFilter = yf }

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetGoName(yname string) string {
    if yname == "ospfHostEntry" { return "Ospfhostentry" }
    return ""
}

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetSegmentPath() string {
    return "ospfHostTable"
}

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfHostEntry" {
        for _, c := range ospfhosttable.Ospfhostentry {
            if ospfhosttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfhosttable_Ospfhostentry{}
        ospfhosttable.Ospfhostentry = append(ospfhosttable.Ospfhostentry, child)
        return &ospfhosttable.Ospfhostentry[len(ospfhosttable.Ospfhostentry)-1]
    }
    return nil
}

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfhosttable.Ospfhostentry {
        children[ospfhosttable.Ospfhostentry[i].GetSegmentPath()] = &ospfhosttable.Ospfhostentry[i]
    }
    return children
}

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetYangName() string { return "ospfHostTable" }

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfhosttable *OSPFMIB_Ospfhosttable) SetParent(parent types.Entity) { ospfhosttable.parent = parent }

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetParent() types.Entity { return ospfhosttable.parent }

func (ospfhosttable *OSPFMIB_Ospfhosttable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfhosttable_Ospfhostentry
// A metric to be advertised, for a given type of
// service, when a given host is reachable.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfhosttable_Ospfhostentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of the host. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfhostipaddress interface{}

    // This attribute is a key. The Type of Service of the route being configured.
    // The type is interface{} with range: 0..30.
    Ospfhosttos interface{}

    // The metric to be advertised. The type is interface{} with range: 0..65535.
    Ospfhostmetric interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfhoststatus interface{}

    // The OSPF area to which the host belongs. Deprecated by ospfHostCfgAreaID.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfhostareaid interface{}

    // To configure the OSPF area to which the host belongs. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfhostcfgareaid interface{}
}

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetFilter() yfilter.YFilter { return ospfhostentry.YFilter }

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) SetFilter(yf yfilter.YFilter) { ospfhostentry.YFilter = yf }

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetGoName(yname string) string {
    if yname == "ospfHostIpAddress" { return "Ospfhostipaddress" }
    if yname == "ospfHostTOS" { return "Ospfhosttos" }
    if yname == "ospfHostMetric" { return "Ospfhostmetric" }
    if yname == "ospfHostStatus" { return "Ospfhoststatus" }
    if yname == "ospfHostAreaID" { return "Ospfhostareaid" }
    if yname == "ospfHostCfgAreaID" { return "Ospfhostcfgareaid" }
    return ""
}

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetSegmentPath() string {
    return "ospfHostEntry" + "[ospfHostIpAddress='" + fmt.Sprintf("%v", ospfhostentry.Ospfhostipaddress) + "']" + "[ospfHostTOS='" + fmt.Sprintf("%v", ospfhostentry.Ospfhosttos) + "']"
}

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfHostIpAddress"] = ospfhostentry.Ospfhostipaddress
    leafs["ospfHostTOS"] = ospfhostentry.Ospfhosttos
    leafs["ospfHostMetric"] = ospfhostentry.Ospfhostmetric
    leafs["ospfHostStatus"] = ospfhostentry.Ospfhoststatus
    leafs["ospfHostAreaID"] = ospfhostentry.Ospfhostareaid
    leafs["ospfHostCfgAreaID"] = ospfhostentry.Ospfhostcfgareaid
    return leafs
}

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetYangName() string { return "ospfHostEntry" }

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) SetParent(parent types.Entity) { ospfhostentry.parent = parent }

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetParent() types.Entity { return ospfhostentry.parent }

func (ospfhostentry *OSPFMIB_Ospfhosttable_Ospfhostentry) GetParentYangName() string { return "ospfHostTable" }

// OSPFMIB_Ospfiftable
// The OSPF Interface Table describes the interfaces
// from the viewpoint of OSPF.
// It augments the ipAddrTable with OSPF specific information.
type OSPFMIB_Ospfiftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The OSPF interface entry describes one interface from the viewpoint of
    // OSPF.  Information in this table is persistent and when this object is
    // written the entity SHOULD save the change to non-volatile storage. The type
    // is slice of OSPFMIB_Ospfiftable_Ospfifentry.
    Ospfifentry []OSPFMIB_Ospfiftable_Ospfifentry
}

func (ospfiftable *OSPFMIB_Ospfiftable) GetFilter() yfilter.YFilter { return ospfiftable.YFilter }

func (ospfiftable *OSPFMIB_Ospfiftable) SetFilter(yf yfilter.YFilter) { ospfiftable.YFilter = yf }

func (ospfiftable *OSPFMIB_Ospfiftable) GetGoName(yname string) string {
    if yname == "ospfIfEntry" { return "Ospfifentry" }
    return ""
}

func (ospfiftable *OSPFMIB_Ospfiftable) GetSegmentPath() string {
    return "ospfIfTable"
}

func (ospfiftable *OSPFMIB_Ospfiftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfIfEntry" {
        for _, c := range ospfiftable.Ospfifentry {
            if ospfiftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfiftable_Ospfifentry{}
        ospfiftable.Ospfifentry = append(ospfiftable.Ospfifentry, child)
        return &ospfiftable.Ospfifentry[len(ospfiftable.Ospfifentry)-1]
    }
    return nil
}

func (ospfiftable *OSPFMIB_Ospfiftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfiftable.Ospfifentry {
        children[ospfiftable.Ospfifentry[i].GetSegmentPath()] = &ospfiftable.Ospfifentry[i]
    }
    return children
}

func (ospfiftable *OSPFMIB_Ospfiftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfiftable *OSPFMIB_Ospfiftable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfiftable *OSPFMIB_Ospfiftable) GetYangName() string { return "ospfIfTable" }

func (ospfiftable *OSPFMIB_Ospfiftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfiftable *OSPFMIB_Ospfiftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfiftable *OSPFMIB_Ospfiftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfiftable *OSPFMIB_Ospfiftable) SetParent(parent types.Entity) { ospfiftable.parent = parent }

func (ospfiftable *OSPFMIB_Ospfiftable) GetParent() types.Entity { return ospfiftable.parent }

func (ospfiftable *OSPFMIB_Ospfiftable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfiftable_Ospfifentry
// The OSPF interface entry describes one interface
// from the viewpoint of OSPF.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfiftable_Ospfifentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of this OSPF interface. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfifipaddress interface{}

    // This attribute is a key. For the purpose of easing the instancing of
    // addressed and addressless interfaces; this variable takes the value 0 on
    // interfaces with IP addresses and the corresponding value of ifIndex for
    // interfaces having no IP address. The type is interface{} with range:
    // 0..2147483647.
    Ospfaddresslessif interface{}

    // A 32-bit integer uniquely identifying the area to which the interface
    // connects.  Area ID 0.0.0.0 is used for the OSPF backbone. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfifareaid interface{}

    // The OSPF interface type. By way of a default, this field may be intuited
    // from the corresponding value of ifType. Broadcast LANs, such as Ethernet
    // and IEEE 802.5, take the value 'broadcast', X.25 and similar technologies
    // take the value 'nbma', and links that are definitively point to point take
    // the value 'pointToPoint'. The type is Ospfiftype.
    Ospfiftype interface{}

    // The OSPF interface's administrative status. The value formed on the
    // interface, and the interface will be advertised as an internal route to
    // some area. The value 'disabled' denotes that the interface is external to
    // OSPF. The type is Status.
    Ospfifadminstat interface{}

    // The priority of this interface.  Used in multi-access networks, this field
    // is used in the designated router election algorithm.  The value 0 signifies
    // that the router is not eligible to become the designated router on this
    // particular network.  In the event of a tie in this value, routers will use
    // their Router ID as a tie breaker. The type is interface{} with range:
    // 0..255.
    Ospfifrtrpriority interface{}

    // The estimated number of seconds it takes to transmit a link state update
    // packet over this interface.  Note that the minimal value SHOULD be 1
    // second. The type is interface{} with range: 0..3600. Units are seconds.
    Ospfiftransitdelay interface{}

    // The number of seconds between link state advertisement retransmissions, for
    // adjacencies belonging to this interface.  This value is also used when
    // retransmitting  database description and Link State request packets. Note
    // that minimal value SHOULD be 1 second. The type is interface{} with range:
    // 0..3600. Units are seconds.
    Ospfifretransinterval interface{}

    // The length of time, in seconds, between the Hello packets that the router
    // sends on the interface.  This value must be the same for all routers
    // attached to a common network. The type is interface{} with range: 1..65535.
    // Units are seconds.
    Ospfifhellointerval interface{}

    // The number of seconds that a router's Hello packets have not been seen
    // before its neighbors declare the router down. This should be some multiple
    // of the Hello interval.  This value must be the same for all routers
    // attached to a common network. The type is interface{} with range:
    // 0..2147483647. Units are seconds.
    Ospfifrtrdeadinterval interface{}

    // The larger time interval, in seconds, between the Hello packets sent to an
    // inactive non-broadcast multi-access neighbor. The type is interface{} with
    // range: 0..2147483647. Units are seconds.
    Ospfifpollinterval interface{}

    // The OSPF Interface State. The type is Ospfifstate.
    Ospfifstate interface{}

    // The IP address of the designated router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfifdesignatedrouter interface{}

    // The IP address of the backup designated router. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfifbackupdesignatedrouter interface{}

    // The number of times this OSPF interface has changed its state or an error
    // has occurred.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ospfDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ospfifevents interface{}

    // The cleartext password used as an OSPF authentication key when
    // simplePassword security is enabled.  This object does not access any OSPF
    // cryptogaphic (e.g., MD5) authentication key under any circumstance.  If the
    // key length is shorter than 8 octets, the agent will left adjust and zero
    // fill to 8 octets.  Unauthenticated interfaces need no authentication key,
    // and simple password authentication cannot use a key of more than 8 octets. 
    // Note that the use of simplePassword authentication is NOT recommended when
    // there is concern regarding attack upon the OSPF system.  SimplePassword
    // authentication is only sufficient to protect against accidental
    // misconfigurations because it re-uses cleartext passwords [RFC1704].  When
    // read, ospfIfAuthKey always returns an octet string of length zero. The type
    // is string with length: 0..256.
    Ospfifauthkey interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfifstatus interface{}

    // The way multicasts should be forwarded on this interface: not forwarded,
    // forwarded as data link multicasts, or forwarded as data link unicasts. 
    // Data link multicasting is not meaningful on point-to-point and NBMA
    // interfaces, and setting ospfMulticastForwarding to 0 effectively disables
    // all multicast forwarding. The type is Ospfifmulticastforwarding.
    Ospfifmulticastforwarding interface{}

    // Indicates whether Demand OSPF procedures (hello suppression to FULL
    // neighbors and setting the DoNotAge flag on propagated LSAs) should be
    // performed on this interface. The type is bool.
    Ospfifdemand interface{}

    // The authentication type specified for an interface.  Note that this object
    // can be used to engage in significant attacks against an OSPF router. The
    // type is OspfAuthenticationType.
    Ospfifauthtype interface{}

    // The total number of link-local link state advertisements in this
    // interface's link-local link state database. The type is interface{} with
    // range: 0..4294967295.
    Ospfiflsacount interface{}

    // The 32-bit unsigned sum of the Link State Advertisements' LS checksums
    // contained in this interface's link-local link state database. The sum can
    // be used to determine if there has been a change in the interface's link
    // state database and to compare the interface link state database of routers
    // attached to the same subnet. The type is interface{} with range:
    // 0..4294967295.
    Ospfiflsacksumsum interface{}

    // The Router ID of the designated router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfifdesignatedrouterid interface{}

    // The Router ID of the backup designated router. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfifbackupdesignatedrouterid interface{}

    // The total number of link-local link state advertisements in this
    // interface's link-local link state database. The type is interface{} with
    // range: 0..4294967295.
    Cospfiflsacount interface{}

    // The 32-bit unsigned sum of the link-state advertisements' LS checksums
    // contained in this interface's link-local link  state database. The sum can
    // be used to determine if there has been a change in the interface's link
    // state database, and to compare the interface link-state database of routers
    // attached to the same subnet. The type is interface{} with range:
    // 0..4294967295.
    Cospfiflsacksumsum interface{}
}

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetFilter() yfilter.YFilter { return ospfifentry.YFilter }

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) SetFilter(yf yfilter.YFilter) { ospfifentry.YFilter = yf }

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetGoName(yname string) string {
    if yname == "ospfIfIpAddress" { return "Ospfifipaddress" }
    if yname == "ospfAddressLessIf" { return "Ospfaddresslessif" }
    if yname == "ospfIfAreaId" { return "Ospfifareaid" }
    if yname == "ospfIfType" { return "Ospfiftype" }
    if yname == "ospfIfAdminStat" { return "Ospfifadminstat" }
    if yname == "ospfIfRtrPriority" { return "Ospfifrtrpriority" }
    if yname == "ospfIfTransitDelay" { return "Ospfiftransitdelay" }
    if yname == "ospfIfRetransInterval" { return "Ospfifretransinterval" }
    if yname == "ospfIfHelloInterval" { return "Ospfifhellointerval" }
    if yname == "ospfIfRtrDeadInterval" { return "Ospfifrtrdeadinterval" }
    if yname == "ospfIfPollInterval" { return "Ospfifpollinterval" }
    if yname == "ospfIfState" { return "Ospfifstate" }
    if yname == "ospfIfDesignatedRouter" { return "Ospfifdesignatedrouter" }
    if yname == "ospfIfBackupDesignatedRouter" { return "Ospfifbackupdesignatedrouter" }
    if yname == "ospfIfEvents" { return "Ospfifevents" }
    if yname == "ospfIfAuthKey" { return "Ospfifauthkey" }
    if yname == "ospfIfStatus" { return "Ospfifstatus" }
    if yname == "ospfIfMulticastForwarding" { return "Ospfifmulticastforwarding" }
    if yname == "ospfIfDemand" { return "Ospfifdemand" }
    if yname == "ospfIfAuthType" { return "Ospfifauthtype" }
    if yname == "ospfIfLsaCount" { return "Ospfiflsacount" }
    if yname == "ospfIfLsaCksumSum" { return "Ospfiflsacksumsum" }
    if yname == "ospfIfDesignatedRouterId" { return "Ospfifdesignatedrouterid" }
    if yname == "ospfIfBackupDesignatedRouterId" { return "Ospfifbackupdesignatedrouterid" }
    if yname == "cospfIfLsaCount" { return "Cospfiflsacount" }
    if yname == "cospfIfLsaCksumSum" { return "Cospfiflsacksumsum" }
    return ""
}

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetSegmentPath() string {
    return "ospfIfEntry" + "[ospfIfIpAddress='" + fmt.Sprintf("%v", ospfifentry.Ospfifipaddress) + "']" + "[ospfAddressLessIf='" + fmt.Sprintf("%v", ospfifentry.Ospfaddresslessif) + "']"
}

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfIfIpAddress"] = ospfifentry.Ospfifipaddress
    leafs["ospfAddressLessIf"] = ospfifentry.Ospfaddresslessif
    leafs["ospfIfAreaId"] = ospfifentry.Ospfifareaid
    leafs["ospfIfType"] = ospfifentry.Ospfiftype
    leafs["ospfIfAdminStat"] = ospfifentry.Ospfifadminstat
    leafs["ospfIfRtrPriority"] = ospfifentry.Ospfifrtrpriority
    leafs["ospfIfTransitDelay"] = ospfifentry.Ospfiftransitdelay
    leafs["ospfIfRetransInterval"] = ospfifentry.Ospfifretransinterval
    leafs["ospfIfHelloInterval"] = ospfifentry.Ospfifhellointerval
    leafs["ospfIfRtrDeadInterval"] = ospfifentry.Ospfifrtrdeadinterval
    leafs["ospfIfPollInterval"] = ospfifentry.Ospfifpollinterval
    leafs["ospfIfState"] = ospfifentry.Ospfifstate
    leafs["ospfIfDesignatedRouter"] = ospfifentry.Ospfifdesignatedrouter
    leafs["ospfIfBackupDesignatedRouter"] = ospfifentry.Ospfifbackupdesignatedrouter
    leafs["ospfIfEvents"] = ospfifentry.Ospfifevents
    leafs["ospfIfAuthKey"] = ospfifentry.Ospfifauthkey
    leafs["ospfIfStatus"] = ospfifentry.Ospfifstatus
    leafs["ospfIfMulticastForwarding"] = ospfifentry.Ospfifmulticastforwarding
    leafs["ospfIfDemand"] = ospfifentry.Ospfifdemand
    leafs["ospfIfAuthType"] = ospfifentry.Ospfifauthtype
    leafs["ospfIfLsaCount"] = ospfifentry.Ospfiflsacount
    leafs["ospfIfLsaCksumSum"] = ospfifentry.Ospfiflsacksumsum
    leafs["ospfIfDesignatedRouterId"] = ospfifentry.Ospfifdesignatedrouterid
    leafs["ospfIfBackupDesignatedRouterId"] = ospfifentry.Ospfifbackupdesignatedrouterid
    leafs["cospfIfLsaCount"] = ospfifentry.Cospfiflsacount
    leafs["cospfIfLsaCksumSum"] = ospfifentry.Cospfiflsacksumsum
    return leafs
}

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetYangName() string { return "ospfIfEntry" }

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) SetParent(parent types.Entity) { ospfifentry.parent = parent }

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetParent() types.Entity { return ospfifentry.parent }

func (ospfifentry *OSPFMIB_Ospfiftable_Ospfifentry) GetParentYangName() string { return "ospfIfTable" }

// OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding represents disables all multicast forwarding.
type OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding string

const (
    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding_blocked OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding = "blocked"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding_multicast OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding = "multicast"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding_unicast OSPFMIB_Ospfiftable_Ospfifentry_Ospfifmulticastforwarding = "unicast"
)

// OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate represents The OSPF Interface State.
type OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate string

const (
    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_down OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "down"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_loopback OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "loopback"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_waiting OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "waiting"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_pointToPoint OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "pointToPoint"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_designatedRouter OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "designatedRouter"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_backupDesignatedRouter OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "backupDesignatedRouter"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate_otherDesignatedRouter OSPFMIB_Ospfiftable_Ospfifentry_Ospfifstate = "otherDesignatedRouter"
)

// OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype represents value 'pointToPoint'.
type OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype string

const (
    OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype_broadcast OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype = "broadcast"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype_nbma OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype = "nbma"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype_pointToPoint OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype = "pointToPoint"

    OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype_pointToMultipoint OSPFMIB_Ospfiftable_Ospfifentry_Ospfiftype = "pointToMultipoint"
)

// OSPFMIB_Ospfifmetrictable
// The Metric Table describes the metrics to be advertised
// for a specified interface at the various types of service.
// As such, this table is an adjunct of the OSPF Interface
// Table.
// 
// Types of service, as defined by RFC 791, have the ability
// to request low delay, high bandwidth, or reliable linkage.
// 
// For the purposes of this specification, the measure of
// bandwidth:
// 
// Metric = referenceBandwidth / ifSpeed
// 
// is the default value.
// The default reference bandwidth is 10^8.
// For multiple link interfaces, note that ifSpeed is the sum
// of the individual link speeds.  This yields a number having
// the following typical values:
// 
// Network Type/bit rate   Metric
// 
// >= 100 MBPS                 1
// Ethernet/802.3             10
// E1                         48
// T1 (ESF)                   65
// 64 KBPS                    1562
// 56 KBPS                    1785
// 19.2 KBPS                  5208
// 9.6 KBPS                   10416
// 
// Routes that are not specified use the default
// (TOS 0) metric.
// 
// Note that the default reference bandwidth can be configured
// using the general group object ospfReferenceBandwidth.
type OSPFMIB_Ospfifmetrictable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A particular TOS metric for a non-virtual interface identified by the
    // interface index.  Information in this table is persistent and when this
    // object is written the entity SHOULD save the change to non-volatile
    // storage. The type is slice of OSPFMIB_Ospfifmetrictable_Ospfifmetricentry.
    Ospfifmetricentry []OSPFMIB_Ospfifmetrictable_Ospfifmetricentry
}

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetFilter() yfilter.YFilter { return ospfifmetrictable.YFilter }

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) SetFilter(yf yfilter.YFilter) { ospfifmetrictable.YFilter = yf }

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetGoName(yname string) string {
    if yname == "ospfIfMetricEntry" { return "Ospfifmetricentry" }
    return ""
}

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetSegmentPath() string {
    return "ospfIfMetricTable"
}

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfIfMetricEntry" {
        for _, c := range ospfifmetrictable.Ospfifmetricentry {
            if ospfifmetrictable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfifmetrictable_Ospfifmetricentry{}
        ospfifmetrictable.Ospfifmetricentry = append(ospfifmetrictable.Ospfifmetricentry, child)
        return &ospfifmetrictable.Ospfifmetricentry[len(ospfifmetrictable.Ospfifmetricentry)-1]
    }
    return nil
}

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfifmetrictable.Ospfifmetricentry {
        children[ospfifmetrictable.Ospfifmetricentry[i].GetSegmentPath()] = &ospfifmetrictable.Ospfifmetricentry[i]
    }
    return children
}

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetYangName() string { return "ospfIfMetricTable" }

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) SetParent(parent types.Entity) { ospfifmetrictable.parent = parent }

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetParent() types.Entity { return ospfifmetrictable.parent }

func (ospfifmetrictable *OSPFMIB_Ospfifmetrictable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfifmetrictable_Ospfifmetricentry
// A particular TOS metric for a non-virtual interface
// identified by the interface index.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfifmetrictable_Ospfifmetricentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of this OSPF interface.  On row
    // creation, this can be derived from the instance. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfifmetricipaddress interface{}

    // This attribute is a key. For the purpose of easing the instancing of
    // addressed and addressless interfaces; this variable takes the value 0 on
    // interfaces with IP addresses and the value of ifIndex for interfaces having
    // no IP address.  On row creation, this can be derived from the instance. The
    // type is interface{} with range: 0..2147483647.
    Ospfifmetricaddresslessif interface{}

    // This attribute is a key. The Type of Service metric being referenced. On
    // row creation, this can be derived from the instance. The type is
    // interface{} with range: 0..30.
    Ospfifmetrictos interface{}

    // The metric of using this Type of Service on this interface.  The default
    // value of the TOS 0 metric is 10^8 / ifSpeed. The type is interface{} with
    // range: 0..65535.
    Ospfifmetricvalue interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfifmetricstatus interface{}
}

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetFilter() yfilter.YFilter { return ospfifmetricentry.YFilter }

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) SetFilter(yf yfilter.YFilter) { ospfifmetricentry.YFilter = yf }

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetGoName(yname string) string {
    if yname == "ospfIfMetricIpAddress" { return "Ospfifmetricipaddress" }
    if yname == "ospfIfMetricAddressLessIf" { return "Ospfifmetricaddresslessif" }
    if yname == "ospfIfMetricTOS" { return "Ospfifmetrictos" }
    if yname == "ospfIfMetricValue" { return "Ospfifmetricvalue" }
    if yname == "ospfIfMetricStatus" { return "Ospfifmetricstatus" }
    return ""
}

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetSegmentPath() string {
    return "ospfIfMetricEntry" + "[ospfIfMetricIpAddress='" + fmt.Sprintf("%v", ospfifmetricentry.Ospfifmetricipaddress) + "']" + "[ospfIfMetricAddressLessIf='" + fmt.Sprintf("%v", ospfifmetricentry.Ospfifmetricaddresslessif) + "']" + "[ospfIfMetricTOS='" + fmt.Sprintf("%v", ospfifmetricentry.Ospfifmetrictos) + "']"
}

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfIfMetricIpAddress"] = ospfifmetricentry.Ospfifmetricipaddress
    leafs["ospfIfMetricAddressLessIf"] = ospfifmetricentry.Ospfifmetricaddresslessif
    leafs["ospfIfMetricTOS"] = ospfifmetricentry.Ospfifmetrictos
    leafs["ospfIfMetricValue"] = ospfifmetricentry.Ospfifmetricvalue
    leafs["ospfIfMetricStatus"] = ospfifmetricentry.Ospfifmetricstatus
    return leafs
}

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetYangName() string { return "ospfIfMetricEntry" }

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) SetParent(parent types.Entity) { ospfifmetricentry.parent = parent }

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetParent() types.Entity { return ospfifmetricentry.parent }

func (ospfifmetricentry *OSPFMIB_Ospfifmetrictable_Ospfifmetricentry) GetParentYangName() string { return "ospfIfMetricTable" }

// OSPFMIB_Ospfvirtiftable
// Information about this router's virtual interfaces
// that the OSPF Process is configured to carry on.
type OSPFMIB_Ospfvirtiftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single virtual interface.  Information in this table is
    // persistent and when this object is written the entity SHOULD save the
    // change to non-volatile storage. The type is slice of
    // OSPFMIB_Ospfvirtiftable_Ospfvirtifentry.
    Ospfvirtifentry []OSPFMIB_Ospfvirtiftable_Ospfvirtifentry
}

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetFilter() yfilter.YFilter { return ospfvirtiftable.YFilter }

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) SetFilter(yf yfilter.YFilter) { ospfvirtiftable.YFilter = yf }

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetGoName(yname string) string {
    if yname == "ospfVirtIfEntry" { return "Ospfvirtifentry" }
    return ""
}

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetSegmentPath() string {
    return "ospfVirtIfTable"
}

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfVirtIfEntry" {
        for _, c := range ospfvirtiftable.Ospfvirtifentry {
            if ospfvirtiftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfvirtiftable_Ospfvirtifentry{}
        ospfvirtiftable.Ospfvirtifentry = append(ospfvirtiftable.Ospfvirtifentry, child)
        return &ospfvirtiftable.Ospfvirtifentry[len(ospfvirtiftable.Ospfvirtifentry)-1]
    }
    return nil
}

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfvirtiftable.Ospfvirtifentry {
        children[ospfvirtiftable.Ospfvirtifentry[i].GetSegmentPath()] = &ospfvirtiftable.Ospfvirtifentry[i]
    }
    return children
}

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetYangName() string { return "ospfVirtIfTable" }

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) SetParent(parent types.Entity) { ospfvirtiftable.parent = parent }

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetParent() types.Entity { return ospfvirtiftable.parent }

func (ospfvirtiftable *OSPFMIB_Ospfvirtiftable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfvirtiftable_Ospfvirtifentry
// Information about a single virtual interface.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfvirtiftable_Ospfvirtifentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The transit area that the virtual link traverses. 
    // By definition, this is not 0.0.0.0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfvirtifareaid interface{}

    // This attribute is a key. The Router ID of the virtual neighbor. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfvirtifneighbor interface{}

    // The estimated number of seconds it takes to transmit a Link State update
    // packet over this interface.  Note that the minimal value SHOULD be 1
    // second. The type is interface{} with range: 0..3600. Units are seconds.
    Ospfvirtiftransitdelay interface{}

    // The number of seconds between link state avertisement retransmissions, for
    // adjacencies belonging to this interface.  This value is also used when
    // retransmitting database description and Link State request packets.  This
    // value should be well over the expected round-trip time.  Note that the
    // minimal value SHOULD be 1 second. The type is interface{} with range:
    // 0..3600. Units are seconds.
    Ospfvirtifretransinterval interface{}

    // The length of time, in seconds, between the Hello packets that the router
    // sends on the interface.  This value must be the same for the virtual
    // neighbor. The type is interface{} with range: 1..65535. Units are seconds.
    Ospfvirtifhellointerval interface{}

    // The number of seconds that a router's Hello packets have not been seen
    // before its neighbors declare the router down.  This should be some multiple
    // of the Hello interval.  This value must be the same for the virtual
    // neighbor. The type is interface{} with range: 0..2147483647. Units are
    // seconds.
    Ospfvirtifrtrdeadinterval interface{}

    // OSPF virtual interface states. The type is Ospfvirtifstate.
    Ospfvirtifstate interface{}

    // The number of state changes or error events on this virtual link. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ospfDiscontinuityTime. The type is interface{} with range: 0..4294967295.
    Ospfvirtifevents interface{}

    // The cleartext password used as an OSPF authentication key when
    // simplePassword security is enabled.  This object does not access any OSPF
    // cryptogaphic (e.g., MD5) authentication key under any circumstance.  If the
    // key length is shorter than 8 octets, the agent will left adjust and zero
    // fill to 8 octets.  Unauthenticated interfaces need no authentication key,
    // and simple password authentication cannot use a key of more than 8 octets. 
    // Note that the use of simplePassword authentication is NOT recommended when
    // there is concern regarding attack upon the OSPF system.  SimplePassword
    // authentication is only sufficient to protect against accidental
    // misconfigurations because it re-uses cleartext passwords.  [RFC1704]  When
    // read, ospfIfAuthKey always returns an octet string of length zero. The type
    // is string with length: 0..256.
    Ospfvirtifauthkey interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfvirtifstatus interface{}

    // The authentication type specified for a virtual interface.  Note that this
    // object can be used to engage in significant attacks against an OSPF router.
    // The type is OspfAuthenticationType.
    Ospfvirtifauthtype interface{}

    // The total number of link-local link state advertisements in this virtual
    // interface's link-local link state database. The type is interface{} with
    // range: 0..4294967295.
    Ospfvirtiflsacount interface{}

    // The 32-bit unsigned sum of the link state advertisements' LS checksums
    // contained in this virtual interface's link-local link state database. The
    // sum can be used to determine if there has been a change in the virtual
    // interface's link state database, and to compare the virtual interface link
    // state database of the virtual neighbors. The type is interface{} with
    // range: 0..4294967295.
    Ospfvirtiflsacksumsum interface{}

    // The total number of link-local link state advertisements in this virtual
    // interface's link-local link state database. The type is interface{} with
    // range: 0..4294967295.
    Cospfvirtiflsacount interface{}

    // The 32-bit unsigned sum of the link-state advertisements' LS checksums
    // contained in this virtual interface's link-local link state database. The
    // sum can be used to determine if there has been a change in the virtual
    // interface's link state database, and to compare the virtual interface
    // link-state database of the virtual neighbors. The type is interface{} with
    // range: 0..4294967295.
    Cospfvirtiflsacksumsum interface{}
}

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetFilter() yfilter.YFilter { return ospfvirtifentry.YFilter }

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) SetFilter(yf yfilter.YFilter) { ospfvirtifentry.YFilter = yf }

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetGoName(yname string) string {
    if yname == "ospfVirtIfAreaId" { return "Ospfvirtifareaid" }
    if yname == "ospfVirtIfNeighbor" { return "Ospfvirtifneighbor" }
    if yname == "ospfVirtIfTransitDelay" { return "Ospfvirtiftransitdelay" }
    if yname == "ospfVirtIfRetransInterval" { return "Ospfvirtifretransinterval" }
    if yname == "ospfVirtIfHelloInterval" { return "Ospfvirtifhellointerval" }
    if yname == "ospfVirtIfRtrDeadInterval" { return "Ospfvirtifrtrdeadinterval" }
    if yname == "ospfVirtIfState" { return "Ospfvirtifstate" }
    if yname == "ospfVirtIfEvents" { return "Ospfvirtifevents" }
    if yname == "ospfVirtIfAuthKey" { return "Ospfvirtifauthkey" }
    if yname == "ospfVirtIfStatus" { return "Ospfvirtifstatus" }
    if yname == "ospfVirtIfAuthType" { return "Ospfvirtifauthtype" }
    if yname == "ospfVirtIfLsaCount" { return "Ospfvirtiflsacount" }
    if yname == "ospfVirtIfLsaCksumSum" { return "Ospfvirtiflsacksumsum" }
    if yname == "cospfVirtIfLsaCount" { return "Cospfvirtiflsacount" }
    if yname == "cospfVirtIfLsaCksumSum" { return "Cospfvirtiflsacksumsum" }
    return ""
}

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetSegmentPath() string {
    return "ospfVirtIfEntry" + "[ospfVirtIfAreaId='" + fmt.Sprintf("%v", ospfvirtifentry.Ospfvirtifareaid) + "']" + "[ospfVirtIfNeighbor='" + fmt.Sprintf("%v", ospfvirtifentry.Ospfvirtifneighbor) + "']"
}

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfVirtIfAreaId"] = ospfvirtifentry.Ospfvirtifareaid
    leafs["ospfVirtIfNeighbor"] = ospfvirtifentry.Ospfvirtifneighbor
    leafs["ospfVirtIfTransitDelay"] = ospfvirtifentry.Ospfvirtiftransitdelay
    leafs["ospfVirtIfRetransInterval"] = ospfvirtifentry.Ospfvirtifretransinterval
    leafs["ospfVirtIfHelloInterval"] = ospfvirtifentry.Ospfvirtifhellointerval
    leafs["ospfVirtIfRtrDeadInterval"] = ospfvirtifentry.Ospfvirtifrtrdeadinterval
    leafs["ospfVirtIfState"] = ospfvirtifentry.Ospfvirtifstate
    leafs["ospfVirtIfEvents"] = ospfvirtifentry.Ospfvirtifevents
    leafs["ospfVirtIfAuthKey"] = ospfvirtifentry.Ospfvirtifauthkey
    leafs["ospfVirtIfStatus"] = ospfvirtifentry.Ospfvirtifstatus
    leafs["ospfVirtIfAuthType"] = ospfvirtifentry.Ospfvirtifauthtype
    leafs["ospfVirtIfLsaCount"] = ospfvirtifentry.Ospfvirtiflsacount
    leafs["ospfVirtIfLsaCksumSum"] = ospfvirtifentry.Ospfvirtiflsacksumsum
    leafs["cospfVirtIfLsaCount"] = ospfvirtifentry.Cospfvirtiflsacount
    leafs["cospfVirtIfLsaCksumSum"] = ospfvirtifentry.Cospfvirtiflsacksumsum
    return leafs
}

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetYangName() string { return "ospfVirtIfEntry" }

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) SetParent(parent types.Entity) { ospfvirtifentry.parent = parent }

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetParent() types.Entity { return ospfvirtifentry.parent }

func (ospfvirtifentry *OSPFMIB_Ospfvirtiftable_Ospfvirtifentry) GetParentYangName() string { return "ospfVirtIfTable" }

// OSPFMIB_Ospfvirtiftable_Ospfvirtifentry_Ospfvirtifstate represents OSPF virtual interface states.
type OSPFMIB_Ospfvirtiftable_Ospfvirtifentry_Ospfvirtifstate string

const (
    OSPFMIB_Ospfvirtiftable_Ospfvirtifentry_Ospfvirtifstate_down OSPFMIB_Ospfvirtiftable_Ospfvirtifentry_Ospfvirtifstate = "down"

    OSPFMIB_Ospfvirtiftable_Ospfvirtifentry_Ospfvirtifstate_pointToPoint OSPFMIB_Ospfvirtiftable_Ospfvirtifentry_Ospfvirtifstate = "pointToPoint"
)

// OSPFMIB_Ospfnbrtable
// A table describing all non-virtual neighbors
// in the locality of the OSPF router.
type OSPFMIB_Ospfnbrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The information regarding a single neighbor.  Information in this table is
    // persistent and when this object is written the entity SHOULD save the
    // change to non-volatile  storage. The type is slice of
    // OSPFMIB_Ospfnbrtable_Ospfnbrentry.
    Ospfnbrentry []OSPFMIB_Ospfnbrtable_Ospfnbrentry
}

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetFilter() yfilter.YFilter { return ospfnbrtable.YFilter }

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) SetFilter(yf yfilter.YFilter) { ospfnbrtable.YFilter = yf }

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetGoName(yname string) string {
    if yname == "ospfNbrEntry" { return "Ospfnbrentry" }
    return ""
}

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetSegmentPath() string {
    return "ospfNbrTable"
}

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfNbrEntry" {
        for _, c := range ospfnbrtable.Ospfnbrentry {
            if ospfnbrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfnbrtable_Ospfnbrentry{}
        ospfnbrtable.Ospfnbrentry = append(ospfnbrtable.Ospfnbrentry, child)
        return &ospfnbrtable.Ospfnbrentry[len(ospfnbrtable.Ospfnbrentry)-1]
    }
    return nil
}

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfnbrtable.Ospfnbrentry {
        children[ospfnbrtable.Ospfnbrentry[i].GetSegmentPath()] = &ospfnbrtable.Ospfnbrentry[i]
    }
    return children
}

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetYangName() string { return "ospfNbrTable" }

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) SetParent(parent types.Entity) { ospfnbrtable.parent = parent }

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetParent() types.Entity { return ospfnbrtable.parent }

func (ospfnbrtable *OSPFMIB_Ospfnbrtable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfnbrtable_Ospfnbrentry
// The information regarding a single neighbor.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// 
// storage.
type OSPFMIB_Ospfnbrtable_Ospfnbrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address this neighbor is using in its IP
    // source address.  Note that, on addressless links, this will not be 0.0.0.0
    // but the  address of another of the neighbor's interfaces. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfnbripaddr interface{}

    // This attribute is a key. On an interface having an IP address, zero. On
    // addressless interfaces, the corresponding value of ifIndex in the Internet
    // Standard MIB. On row creation, this can be derived from the instance. The
    // type is interface{} with range: 0..2147483647.
    Ospfnbraddresslessindex interface{}

    // A 32-bit integer (represented as a type IpAddress) uniquely identifying the
    // neighboring router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfnbrrtrid interface{}

    // A bit mask corresponding to the neighbor's options field.  Bit 0, if set,
    // indicates that the system will operate on Type of Service metrics other
    // than TOS 0.  If zero, the neighbor will ignore all metrics except the TOS 0
    // metric.  Bit 1, if set, indicates that the associated area accepts and
    // operates on external information; if zero, it is a stub area.  Bit 2, if
    // set, indicates that the system is capable of routing IP multicast
    // datagrams, that is that it implements the multicast extensions to OSPF. 
    // Bit 3, if set, indicates that the associated area is an NSSA.  These areas
    // are capable of carrying type-7 external advertisements, which are
    // translated into type-5 external advertisements at NSSA borders. The type is
    // interface{} with range: -2147483648..2147483647.
    Ospfnbroptions interface{}

    // The priority of this neighbor in the designated router election algorithm. 
    // The value 0 signifies that the neighbor is not eligible to become the
    // designated router on this particular network. The type is interface{} with
    // range: 0..255.
    Ospfnbrpriority interface{}

    // The state of the relationship with this neighbor. The type is Ospfnbrstate.
    Ospfnbrstate interface{}

    // The number of times this neighbor relationship has changed state or an
    // error has occurred.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ospfDiscontinuityTime. The type is interface{}
    // with range: 0..4294967295.
    Ospfnbrevents interface{}

    // The current length of the retransmission queue. The type is interface{}
    // with range: 0..4294967295.
    Ospfnbrlsretransqlen interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfnbmanbrstatus interface{}

    // This variable displays the status of the entry; 'dynamic' and 'permanent'
    // refer to how the neighbor became known. The type is Ospfnbmanbrpermanence.
    Ospfnbmanbrpermanence interface{}

    // Indicates whether Hellos are being suppressed to the neighbor. The type is
    // bool.
    Ospfnbrhellosuppressed interface{}

    // Indicates whether the router is acting as a graceful restart helper for the
    // neighbor. The type is Ospfnbrrestarthelperstatus.
    Ospfnbrrestarthelperstatus interface{}

    // Remaining time in current OSPF graceful restart interval, if the router is
    // acting as a restart helper for the neighbor. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    Ospfnbrrestarthelperage interface{}

    // Describes the outcome of the last attempt at acting as a graceful restart
    // helper for the neighbor. The type is Ospfnbrrestarthelperexitreason.
    Ospfnbrrestarthelperexitreason interface{}
}

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetFilter() yfilter.YFilter { return ospfnbrentry.YFilter }

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) SetFilter(yf yfilter.YFilter) { ospfnbrentry.YFilter = yf }

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetGoName(yname string) string {
    if yname == "ospfNbrIpAddr" { return "Ospfnbripaddr" }
    if yname == "ospfNbrAddressLessIndex" { return "Ospfnbraddresslessindex" }
    if yname == "ospfNbrRtrId" { return "Ospfnbrrtrid" }
    if yname == "ospfNbrOptions" { return "Ospfnbroptions" }
    if yname == "ospfNbrPriority" { return "Ospfnbrpriority" }
    if yname == "ospfNbrState" { return "Ospfnbrstate" }
    if yname == "ospfNbrEvents" { return "Ospfnbrevents" }
    if yname == "ospfNbrLsRetransQLen" { return "Ospfnbrlsretransqlen" }
    if yname == "ospfNbmaNbrStatus" { return "Ospfnbmanbrstatus" }
    if yname == "ospfNbmaNbrPermanence" { return "Ospfnbmanbrpermanence" }
    if yname == "ospfNbrHelloSuppressed" { return "Ospfnbrhellosuppressed" }
    if yname == "ospfNbrRestartHelperStatus" { return "Ospfnbrrestarthelperstatus" }
    if yname == "ospfNbrRestartHelperAge" { return "Ospfnbrrestarthelperage" }
    if yname == "ospfNbrRestartHelperExitReason" { return "Ospfnbrrestarthelperexitreason" }
    return ""
}

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetSegmentPath() string {
    return "ospfNbrEntry" + "[ospfNbrIpAddr='" + fmt.Sprintf("%v", ospfnbrentry.Ospfnbripaddr) + "']" + "[ospfNbrAddressLessIndex='" + fmt.Sprintf("%v", ospfnbrentry.Ospfnbraddresslessindex) + "']"
}

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfNbrIpAddr"] = ospfnbrentry.Ospfnbripaddr
    leafs["ospfNbrAddressLessIndex"] = ospfnbrentry.Ospfnbraddresslessindex
    leafs["ospfNbrRtrId"] = ospfnbrentry.Ospfnbrrtrid
    leafs["ospfNbrOptions"] = ospfnbrentry.Ospfnbroptions
    leafs["ospfNbrPriority"] = ospfnbrentry.Ospfnbrpriority
    leafs["ospfNbrState"] = ospfnbrentry.Ospfnbrstate
    leafs["ospfNbrEvents"] = ospfnbrentry.Ospfnbrevents
    leafs["ospfNbrLsRetransQLen"] = ospfnbrentry.Ospfnbrlsretransqlen
    leafs["ospfNbmaNbrStatus"] = ospfnbrentry.Ospfnbmanbrstatus
    leafs["ospfNbmaNbrPermanence"] = ospfnbrentry.Ospfnbmanbrpermanence
    leafs["ospfNbrHelloSuppressed"] = ospfnbrentry.Ospfnbrhellosuppressed
    leafs["ospfNbrRestartHelperStatus"] = ospfnbrentry.Ospfnbrrestarthelperstatus
    leafs["ospfNbrRestartHelperAge"] = ospfnbrentry.Ospfnbrrestarthelperage
    leafs["ospfNbrRestartHelperExitReason"] = ospfnbrentry.Ospfnbrrestarthelperexitreason
    return leafs
}

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetYangName() string { return "ospfNbrEntry" }

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) SetParent(parent types.Entity) { ospfnbrentry.parent = parent }

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetParent() types.Entity { return ospfnbrentry.parent }

func (ospfnbrentry *OSPFMIB_Ospfnbrtable_Ospfnbrentry) GetParentYangName() string { return "ospfNbrTable" }

// OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbmanbrpermanence represents became known.
type OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbmanbrpermanence string

const (
    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbmanbrpermanence_dynamic OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbmanbrpermanence = "dynamic"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbmanbrpermanence_permanent OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbmanbrpermanence = "permanent"
)

// OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason represents as a graceful restart helper for the neighbor.
type OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason string

const (
    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason_none OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason = "none"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason_inProgress OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason = "inProgress"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason_completed OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason = "completed"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason_timedOut OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason = "timedOut"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason_topologyChanged OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperexitreason = "topologyChanged"
)

// OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperstatus represents as a graceful restart helper for the neighbor.
type OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperstatus string

const (
    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperstatus_notHelping OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperstatus = "notHelping"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperstatus_helping OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrrestarthelperstatus = "helping"
)

// OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate represents The state of the relationship with this neighbor.
type OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate string

const (
    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_down OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "down"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_attempt OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "attempt"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_init OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "init"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_twoWay OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "twoWay"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_exchangeStart OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "exchangeStart"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_exchange OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "exchange"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_loading OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "loading"

    OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate_full OSPFMIB_Ospfnbrtable_Ospfnbrentry_Ospfnbrstate = "full"
)

// OSPFMIB_Ospfvirtnbrtable
// This table describes all virtual neighbors.
// Since virtual links are configured
// in the Virtual Interface Table, this table is read-only.
type OSPFMIB_Ospfvirtnbrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Virtual neighbor information. The type is slice of
    // OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry.
    Ospfvirtnbrentry []OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry
}

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetFilter() yfilter.YFilter { return ospfvirtnbrtable.YFilter }

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) SetFilter(yf yfilter.YFilter) { ospfvirtnbrtable.YFilter = yf }

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetGoName(yname string) string {
    if yname == "ospfVirtNbrEntry" { return "Ospfvirtnbrentry" }
    return ""
}

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetSegmentPath() string {
    return "ospfVirtNbrTable"
}

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfVirtNbrEntry" {
        for _, c := range ospfvirtnbrtable.Ospfvirtnbrentry {
            if ospfvirtnbrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry{}
        ospfvirtnbrtable.Ospfvirtnbrentry = append(ospfvirtnbrtable.Ospfvirtnbrentry, child)
        return &ospfvirtnbrtable.Ospfvirtnbrentry[len(ospfvirtnbrtable.Ospfvirtnbrentry)-1]
    }
    return nil
}

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfvirtnbrtable.Ospfvirtnbrentry {
        children[ospfvirtnbrtable.Ospfvirtnbrentry[i].GetSegmentPath()] = &ospfvirtnbrtable.Ospfvirtnbrentry[i]
    }
    return children
}

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetYangName() string { return "ospfVirtNbrTable" }

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) SetParent(parent types.Entity) { ospfvirtnbrtable.parent = parent }

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetParent() types.Entity { return ospfvirtnbrtable.parent }

func (ospfvirtnbrtable *OSPFMIB_Ospfvirtnbrtable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry
// Virtual neighbor information.
type OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Transit Area Identifier. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfvirtnbrarea interface{}

    // This attribute is a key. A 32-bit integer uniquely identifying the
    // neighboring router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfvirtnbrrtrid interface{}

    // The IP address this virtual neighbor is using. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfvirtnbripaddr interface{}

    // A bit mask corresponding to the neighbor's options field.  Bit 1, if set,
    // indicates that the system will operate on Type of Service metrics other
    // than TOS 0.  If zero, the neighbor will ignore all metrics except the TOS 0
    // metric.  Bit 2, if set, indicates that the system is network multicast
    // capable, i.e., that it implements OSPF multicast routing. The type is
    // interface{} with range: -2147483648..2147483647.
    Ospfvirtnbroptions interface{}

    // The state of the virtual neighbor relationship. The type is
    // Ospfvirtnbrstate.
    Ospfvirtnbrstate interface{}

    // The number of times this virtual link has changed its state or an error has
    // occurred.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ospfDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ospfvirtnbrevents interface{}

    // The current length of the retransmission queue. The type is interface{}
    // with range: 0..4294967295.
    Ospfvirtnbrlsretransqlen interface{}

    // Indicates whether Hellos are being suppressed to the neighbor. The type is
    // bool.
    Ospfvirtnbrhellosuppressed interface{}

    // Indicates whether the router is acting as a graceful restart helper for the
    // neighbor. The type is Ospfvirtnbrrestarthelperstatus.
    Ospfvirtnbrrestarthelperstatus interface{}

    // Remaining time in current OSPF graceful restart interval, if the router is
    // acting as a restart helper for the neighbor. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    Ospfvirtnbrrestarthelperage interface{}

    // Describes the outcome of the last attempt at acting as a graceful restart
    // helper for the neighbor. The type is Ospfvirtnbrrestarthelperexitreason.
    Ospfvirtnbrrestarthelperexitreason interface{}
}

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetFilter() yfilter.YFilter { return ospfvirtnbrentry.YFilter }

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) SetFilter(yf yfilter.YFilter) { ospfvirtnbrentry.YFilter = yf }

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetGoName(yname string) string {
    if yname == "ospfVirtNbrArea" { return "Ospfvirtnbrarea" }
    if yname == "ospfVirtNbrRtrId" { return "Ospfvirtnbrrtrid" }
    if yname == "ospfVirtNbrIpAddr" { return "Ospfvirtnbripaddr" }
    if yname == "ospfVirtNbrOptions" { return "Ospfvirtnbroptions" }
    if yname == "ospfVirtNbrState" { return "Ospfvirtnbrstate" }
    if yname == "ospfVirtNbrEvents" { return "Ospfvirtnbrevents" }
    if yname == "ospfVirtNbrLsRetransQLen" { return "Ospfvirtnbrlsretransqlen" }
    if yname == "ospfVirtNbrHelloSuppressed" { return "Ospfvirtnbrhellosuppressed" }
    if yname == "ospfVirtNbrRestartHelperStatus" { return "Ospfvirtnbrrestarthelperstatus" }
    if yname == "ospfVirtNbrRestartHelperAge" { return "Ospfvirtnbrrestarthelperage" }
    if yname == "ospfVirtNbrRestartHelperExitReason" { return "Ospfvirtnbrrestarthelperexitreason" }
    return ""
}

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetSegmentPath() string {
    return "ospfVirtNbrEntry" + "[ospfVirtNbrArea='" + fmt.Sprintf("%v", ospfvirtnbrentry.Ospfvirtnbrarea) + "']" + "[ospfVirtNbrRtrId='" + fmt.Sprintf("%v", ospfvirtnbrentry.Ospfvirtnbrrtrid) + "']"
}

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfVirtNbrArea"] = ospfvirtnbrentry.Ospfvirtnbrarea
    leafs["ospfVirtNbrRtrId"] = ospfvirtnbrentry.Ospfvirtnbrrtrid
    leafs["ospfVirtNbrIpAddr"] = ospfvirtnbrentry.Ospfvirtnbripaddr
    leafs["ospfVirtNbrOptions"] = ospfvirtnbrentry.Ospfvirtnbroptions
    leafs["ospfVirtNbrState"] = ospfvirtnbrentry.Ospfvirtnbrstate
    leafs["ospfVirtNbrEvents"] = ospfvirtnbrentry.Ospfvirtnbrevents
    leafs["ospfVirtNbrLsRetransQLen"] = ospfvirtnbrentry.Ospfvirtnbrlsretransqlen
    leafs["ospfVirtNbrHelloSuppressed"] = ospfvirtnbrentry.Ospfvirtnbrhellosuppressed
    leafs["ospfVirtNbrRestartHelperStatus"] = ospfvirtnbrentry.Ospfvirtnbrrestarthelperstatus
    leafs["ospfVirtNbrRestartHelperAge"] = ospfvirtnbrentry.Ospfvirtnbrrestarthelperage
    leafs["ospfVirtNbrRestartHelperExitReason"] = ospfvirtnbrentry.Ospfvirtnbrrestarthelperexitreason
    return leafs
}

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetYangName() string { return "ospfVirtNbrEntry" }

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) SetParent(parent types.Entity) { ospfvirtnbrentry.parent = parent }

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetParent() types.Entity { return ospfvirtnbrentry.parent }

func (ospfvirtnbrentry *OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry) GetParentYangName() string { return "ospfVirtNbrTable" }

// OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason represents as a graceful restart helper for the neighbor.
type OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason string

const (
    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason_none OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason = "none"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason_inProgress OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason = "inProgress"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason_completed OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason = "completed"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason_timedOut OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason = "timedOut"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason_topologyChanged OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperexitreason = "topologyChanged"
)

// OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperstatus represents as a graceful restart helper for the neighbor.
type OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperstatus string

const (
    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperstatus_notHelping OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperstatus = "notHelping"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperstatus_helping OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrrestarthelperstatus = "helping"
)

// OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate represents The state of the virtual neighbor relationship.
type OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate string

const (
    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_down OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "down"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_attempt OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "attempt"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_init OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "init"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_twoWay OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "twoWay"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_exchangeStart OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "exchangeStart"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_exchange OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "exchange"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_loading OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "loading"

    OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate_full OSPFMIB_Ospfvirtnbrtable_Ospfvirtnbrentry_Ospfvirtnbrstate = "full"
)

// OSPFMIB_Ospfextlsdbtable
// The OSPF Process's external LSA link state database.
// 
// This table is identical to the OSPF LSDB Table
// in format, but contains only external link state
// advertisements.  The purpose is to allow external
// 
// LSAs to be displayed once for the router rather
// than once in each non-stub area.
// 
// Note that external LSAs are also in the AS-scope link state
// database.
type OSPFMIB_Ospfextlsdbtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry.
    Ospfextlsdbentry []OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry
}

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetFilter() yfilter.YFilter { return ospfextlsdbtable.YFilter }

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) SetFilter(yf yfilter.YFilter) { ospfextlsdbtable.YFilter = yf }

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetGoName(yname string) string {
    if yname == "ospfExtLsdbEntry" { return "Ospfextlsdbentry" }
    return ""
}

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetSegmentPath() string {
    return "ospfExtLsdbTable"
}

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfExtLsdbEntry" {
        for _, c := range ospfextlsdbtable.Ospfextlsdbentry {
            if ospfextlsdbtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry{}
        ospfextlsdbtable.Ospfextlsdbentry = append(ospfextlsdbtable.Ospfextlsdbentry, child)
        return &ospfextlsdbtable.Ospfextlsdbentry[len(ospfextlsdbtable.Ospfextlsdbentry)-1]
    }
    return nil
}

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfextlsdbtable.Ospfextlsdbentry {
        children[ospfextlsdbtable.Ospfextlsdbentry[i].GetSegmentPath()] = &ospfextlsdbtable.Ospfextlsdbentry[i]
    }
    return children
}

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetYangName() string { return "ospfExtLsdbTable" }

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) SetParent(parent types.Entity) { ospfextlsdbtable.parent = parent }

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetParent() types.Entity { return ospfextlsdbtable.parent }

func (ospfextlsdbtable *OSPFMIB_Ospfextlsdbtable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry
// A single link state advertisement.
type OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // Ospfextlsdbtype.
    Ospfextlsdbtype interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing either a Router ID or an IP address; it identifies the piece of
    // the routing domain that is being described by the advertisement. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfextlsdblsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfextlsdbrouterid interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate link state advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    Ospfextlsdbsequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    Ospfextlsdbage interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred to as the Fletcher checksum. The type is interface{}
    // with range: -2147483648..2147483647.
    Ospfextlsdbchecksum interface{}

    // The entire link state advertisement, including its header. The type is
    // string with length: 36.
    Ospfextlsdbadvertisement interface{}
}

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetFilter() yfilter.YFilter { return ospfextlsdbentry.YFilter }

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) SetFilter(yf yfilter.YFilter) { ospfextlsdbentry.YFilter = yf }

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetGoName(yname string) string {
    if yname == "ospfExtLsdbType" { return "Ospfextlsdbtype" }
    if yname == "ospfExtLsdbLsid" { return "Ospfextlsdblsid" }
    if yname == "ospfExtLsdbRouterId" { return "Ospfextlsdbrouterid" }
    if yname == "ospfExtLsdbSequence" { return "Ospfextlsdbsequence" }
    if yname == "ospfExtLsdbAge" { return "Ospfextlsdbage" }
    if yname == "ospfExtLsdbChecksum" { return "Ospfextlsdbchecksum" }
    if yname == "ospfExtLsdbAdvertisement" { return "Ospfextlsdbadvertisement" }
    return ""
}

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetSegmentPath() string {
    return "ospfExtLsdbEntry" + "[ospfExtLsdbType='" + fmt.Sprintf("%v", ospfextlsdbentry.Ospfextlsdbtype) + "']" + "[ospfExtLsdbLsid='" + fmt.Sprintf("%v", ospfextlsdbentry.Ospfextlsdblsid) + "']" + "[ospfExtLsdbRouterId='" + fmt.Sprintf("%v", ospfextlsdbentry.Ospfextlsdbrouterid) + "']"
}

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfExtLsdbType"] = ospfextlsdbentry.Ospfextlsdbtype
    leafs["ospfExtLsdbLsid"] = ospfextlsdbentry.Ospfextlsdblsid
    leafs["ospfExtLsdbRouterId"] = ospfextlsdbentry.Ospfextlsdbrouterid
    leafs["ospfExtLsdbSequence"] = ospfextlsdbentry.Ospfextlsdbsequence
    leafs["ospfExtLsdbAge"] = ospfextlsdbentry.Ospfextlsdbage
    leafs["ospfExtLsdbChecksum"] = ospfextlsdbentry.Ospfextlsdbchecksum
    leafs["ospfExtLsdbAdvertisement"] = ospfextlsdbentry.Ospfextlsdbadvertisement
    return leafs
}

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetYangName() string { return "ospfExtLsdbEntry" }

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) SetParent(parent types.Entity) { ospfextlsdbentry.parent = parent }

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetParent() types.Entity { return ospfextlsdbentry.parent }

func (ospfextlsdbentry *OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry) GetParentYangName() string { return "ospfExtLsdbTable" }

// OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry_Ospfextlsdbtype represents format.
type OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry_Ospfextlsdbtype string

const (
    OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry_Ospfextlsdbtype_asExternalLink OSPFMIB_Ospfextlsdbtable_Ospfextlsdbentry_Ospfextlsdbtype = "asExternalLink"
)

// OSPFMIB_Ospfareaaggregatetable
// The Area Aggregate Table acts as an adjunct
// to the Area Table.  It describes those address aggregates
// that are configured to be propagated from an area.
// Its purpose is to reduce the amount of information
// that is known beyond an Area's borders.
// 
// It contains a set of IP address ranges
// specified by an IP address/IP network mask pair.
// For example, a class B address range of X.X.X.X
// with a network mask of 255.255.0.0 includes all IP
// addresses from X.X.0.0 to X.X.255.255.
// 
// Note that if ranges are configured such that one range
// subsumes another range (e.g., 10.0.0.0 mask 255.0.0.0
// and 10.1.0.0 mask 255.255.0.0),
// the most specific match is the preferred one.
type OSPFMIB_Ospfareaaggregatetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single area aggregate entry.  Information in this table is persistent and
    // when this object is written the entity SHOULD save the change to
    // non-volatile storage. The type is slice of
    // OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry.
    Ospfareaaggregateentry []OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry
}

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetFilter() yfilter.YFilter { return ospfareaaggregatetable.YFilter }

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) SetFilter(yf yfilter.YFilter) { ospfareaaggregatetable.YFilter = yf }

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetGoName(yname string) string {
    if yname == "ospfAreaAggregateEntry" { return "Ospfareaaggregateentry" }
    return ""
}

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetSegmentPath() string {
    return "ospfAreaAggregateTable"
}

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfAreaAggregateEntry" {
        for _, c := range ospfareaaggregatetable.Ospfareaaggregateentry {
            if ospfareaaggregatetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry{}
        ospfareaaggregatetable.Ospfareaaggregateentry = append(ospfareaaggregatetable.Ospfareaaggregateentry, child)
        return &ospfareaaggregatetable.Ospfareaaggregateentry[len(ospfareaaggregatetable.Ospfareaaggregateentry)-1]
    }
    return nil
}

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfareaaggregatetable.Ospfareaaggregateentry {
        children[ospfareaaggregatetable.Ospfareaaggregateentry[i].GetSegmentPath()] = &ospfareaaggregatetable.Ospfareaaggregateentry[i]
    }
    return children
}

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetYangName() string { return "ospfAreaAggregateTable" }

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) SetParent(parent types.Entity) { ospfareaaggregatetable.parent = parent }

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetParent() types.Entity { return ospfareaaggregatetable.parent }

func (ospfareaaggregatetable *OSPFMIB_Ospfareaaggregatetable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry
// A single area aggregate entry.
// 
// Information in this table is persistent and when this object
// is written the entity SHOULD save the change to non-volatile
// storage.
type OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The area within which the address aggregate is to
    // be found. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfareaaggregateareaid interface{}

    // This attribute is a key. The type of the address aggregate.  This field
    // specifies the Lsdb type that this address aggregate applies to. The type is
    // Ospfareaaggregatelsdbtype.
    Ospfareaaggregatelsdbtype interface{}

    // This attribute is a key. The IP address of the net or subnet indicated by
    // the range. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfareaaggregatenet interface{}

    // This attribute is a key. The subnet mask that pertains to the net or
    // subnet. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfareaaggregatemask interface{}

    // This object permits management of the table by facilitating actions such as
    // row creation, construction, and destruction.  The value of this object has
    // no effect on whether other objects in this conceptual row can be modified.
    // The type is RowStatus.
    Ospfareaaggregatestatus interface{}

    // Subnets subsumed by ranges either trigger the advertisement of the
    // indicated aggregate (advertiseMatching) or result in the subnet's not being
    // advertised at all outside the area. The type is Ospfareaaggregateeffect.
    Ospfareaaggregateeffect interface{}

    // External route tag to be included in NSSA (type-7) LSAs. The type is
    // interface{} with range: 0..4294967295.
    Ospfareaaggregateextroutetag interface{}
}

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetFilter() yfilter.YFilter { return ospfareaaggregateentry.YFilter }

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) SetFilter(yf yfilter.YFilter) { ospfareaaggregateentry.YFilter = yf }

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetGoName(yname string) string {
    if yname == "ospfAreaAggregateAreaID" { return "Ospfareaaggregateareaid" }
    if yname == "ospfAreaAggregateLsdbType" { return "Ospfareaaggregatelsdbtype" }
    if yname == "ospfAreaAggregateNet" { return "Ospfareaaggregatenet" }
    if yname == "ospfAreaAggregateMask" { return "Ospfareaaggregatemask" }
    if yname == "ospfAreaAggregateStatus" { return "Ospfareaaggregatestatus" }
    if yname == "ospfAreaAggregateEffect" { return "Ospfareaaggregateeffect" }
    if yname == "ospfAreaAggregateExtRouteTag" { return "Ospfareaaggregateextroutetag" }
    return ""
}

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetSegmentPath() string {
    return "ospfAreaAggregateEntry" + "[ospfAreaAggregateAreaID='" + fmt.Sprintf("%v", ospfareaaggregateentry.Ospfareaaggregateareaid) + "']" + "[ospfAreaAggregateLsdbType='" + fmt.Sprintf("%v", ospfareaaggregateentry.Ospfareaaggregatelsdbtype) + "']" + "[ospfAreaAggregateNet='" + fmt.Sprintf("%v", ospfareaaggregateentry.Ospfareaaggregatenet) + "']" + "[ospfAreaAggregateMask='" + fmt.Sprintf("%v", ospfareaaggregateentry.Ospfareaaggregatemask) + "']"
}

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfAreaAggregateAreaID"] = ospfareaaggregateentry.Ospfareaaggregateareaid
    leafs["ospfAreaAggregateLsdbType"] = ospfareaaggregateentry.Ospfareaaggregatelsdbtype
    leafs["ospfAreaAggregateNet"] = ospfareaaggregateentry.Ospfareaaggregatenet
    leafs["ospfAreaAggregateMask"] = ospfareaaggregateentry.Ospfareaaggregatemask
    leafs["ospfAreaAggregateStatus"] = ospfareaaggregateentry.Ospfareaaggregatestatus
    leafs["ospfAreaAggregateEffect"] = ospfareaaggregateentry.Ospfareaaggregateeffect
    leafs["ospfAreaAggregateExtRouteTag"] = ospfareaaggregateentry.Ospfareaaggregateextroutetag
    return leafs
}

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetYangName() string { return "ospfAreaAggregateEntry" }

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) SetParent(parent types.Entity) { ospfareaaggregateentry.parent = parent }

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetParent() types.Entity { return ospfareaaggregateentry.parent }

func (ospfareaaggregateentry *OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry) GetParentYangName() string { return "ospfAreaAggregateTable" }

// OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregateeffect represents being advertised at all outside the area.
type OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregateeffect string

const (
    OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregateeffect_advertiseMatching OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregateeffect = "advertiseMatching"

    OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregateeffect_doNotAdvertiseMatching OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregateeffect = "doNotAdvertiseMatching"
)

// OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregatelsdbtype represents aggregate applies to.
type OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregatelsdbtype string

const (
    OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregatelsdbtype_summaryLink OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregatelsdbtype = "summaryLink"

    OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregatelsdbtype_nssaExternalLink OSPFMIB_Ospfareaaggregatetable_Ospfareaaggregateentry_Ospfareaaggregatelsdbtype = "nssaExternalLink"
)

// OSPFMIB_Ospflocallsdbtable
// The OSPF Process's link-local link state database
// for non-virtual links.
// This table is identical to the OSPF LSDB Table
// in format, but contains only link-local Link State
// Advertisements for non-virtual links.  The purpose is
// to allow link-local LSAs to be displayed for each
// non-virtual interface.  This table is implemented to
// support type-9 LSAs that are defined
// in 'The OSPF Opaque LSA Option'.
type OSPFMIB_Ospflocallsdbtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry.
    Ospflocallsdbentry []OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry
}

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetFilter() yfilter.YFilter { return ospflocallsdbtable.YFilter }

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) SetFilter(yf yfilter.YFilter) { ospflocallsdbtable.YFilter = yf }

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetGoName(yname string) string {
    if yname == "ospfLocalLsdbEntry" { return "Ospflocallsdbentry" }
    return ""
}

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetSegmentPath() string {
    return "ospfLocalLsdbTable"
}

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfLocalLsdbEntry" {
        for _, c := range ospflocallsdbtable.Ospflocallsdbentry {
            if ospflocallsdbtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry{}
        ospflocallsdbtable.Ospflocallsdbentry = append(ospflocallsdbtable.Ospflocallsdbentry, child)
        return &ospflocallsdbtable.Ospflocallsdbentry[len(ospflocallsdbtable.Ospflocallsdbentry)-1]
    }
    return nil
}

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospflocallsdbtable.Ospflocallsdbentry {
        children[ospflocallsdbtable.Ospflocallsdbentry[i].GetSegmentPath()] = &ospflocallsdbtable.Ospflocallsdbentry[i]
    }
    return children
}

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetBundleName() string { return "cisco_ios_xe" }

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetYangName() string { return "ospfLocalLsdbTable" }

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) SetParent(parent types.Entity) { ospflocallsdbtable.parent = parent }

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetParent() types.Entity { return ospflocallsdbtable.parent }

func (ospflocallsdbtable *OSPFMIB_Ospflocallsdbtable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry
// A single link state advertisement.
type OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of the interface from which the LSA
    // was received if the interface is numbered. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospflocallsdbipaddress interface{}

    // This attribute is a key. The interface index of the interface from which
    // the LSA was received if the interface is unnumbered. The type is
    // interface{} with range: 0..2147483647.
    Ospflocallsdbaddresslessif interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // Ospflocallsdbtype.
    Ospflocallsdbtype interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing a 32-bit identifier in IP address format; it identifies the
    // piece of the routing domain that is being described by the advertisement.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospflocallsdblsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospflocallsdbrouterid interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate link state advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    Ospflocallsdbsequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    Ospflocallsdbage interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred to as the Fletcher checksum. The type is interface{}
    // with range: -2147483648..2147483647.
    Ospflocallsdbchecksum interface{}

    // The entire link state advertisement, including its header.  Note that for
    // variable length LSAs, SNMP agents may not be able to return the largest
    // string size. The type is string with length: 1..65535.
    Ospflocallsdbadvertisement interface{}
}

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetFilter() yfilter.YFilter { return ospflocallsdbentry.YFilter }

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) SetFilter(yf yfilter.YFilter) { ospflocallsdbentry.YFilter = yf }

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetGoName(yname string) string {
    if yname == "ospfLocalLsdbIpAddress" { return "Ospflocallsdbipaddress" }
    if yname == "ospfLocalLsdbAddressLessIf" { return "Ospflocallsdbaddresslessif" }
    if yname == "ospfLocalLsdbType" { return "Ospflocallsdbtype" }
    if yname == "ospfLocalLsdbLsid" { return "Ospflocallsdblsid" }
    if yname == "ospfLocalLsdbRouterId" { return "Ospflocallsdbrouterid" }
    if yname == "ospfLocalLsdbSequence" { return "Ospflocallsdbsequence" }
    if yname == "ospfLocalLsdbAge" { return "Ospflocallsdbage" }
    if yname == "ospfLocalLsdbChecksum" { return "Ospflocallsdbchecksum" }
    if yname == "ospfLocalLsdbAdvertisement" { return "Ospflocallsdbadvertisement" }
    return ""
}

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetSegmentPath() string {
    return "ospfLocalLsdbEntry" + "[ospfLocalLsdbIpAddress='" + fmt.Sprintf("%v", ospflocallsdbentry.Ospflocallsdbipaddress) + "']" + "[ospfLocalLsdbAddressLessIf='" + fmt.Sprintf("%v", ospflocallsdbentry.Ospflocallsdbaddresslessif) + "']" + "[ospfLocalLsdbType='" + fmt.Sprintf("%v", ospflocallsdbentry.Ospflocallsdbtype) + "']" + "[ospfLocalLsdbLsid='" + fmt.Sprintf("%v", ospflocallsdbentry.Ospflocallsdblsid) + "']" + "[ospfLocalLsdbRouterId='" + fmt.Sprintf("%v", ospflocallsdbentry.Ospflocallsdbrouterid) + "']"
}

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfLocalLsdbIpAddress"] = ospflocallsdbentry.Ospflocallsdbipaddress
    leafs["ospfLocalLsdbAddressLessIf"] = ospflocallsdbentry.Ospflocallsdbaddresslessif
    leafs["ospfLocalLsdbType"] = ospflocallsdbentry.Ospflocallsdbtype
    leafs["ospfLocalLsdbLsid"] = ospflocallsdbentry.Ospflocallsdblsid
    leafs["ospfLocalLsdbRouterId"] = ospflocallsdbentry.Ospflocallsdbrouterid
    leafs["ospfLocalLsdbSequence"] = ospflocallsdbentry.Ospflocallsdbsequence
    leafs["ospfLocalLsdbAge"] = ospflocallsdbentry.Ospflocallsdbage
    leafs["ospfLocalLsdbChecksum"] = ospflocallsdbentry.Ospflocallsdbchecksum
    leafs["ospfLocalLsdbAdvertisement"] = ospflocallsdbentry.Ospflocallsdbadvertisement
    return leafs
}

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetYangName() string { return "ospfLocalLsdbEntry" }

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) SetParent(parent types.Entity) { ospflocallsdbentry.parent = parent }

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetParent() types.Entity { return ospflocallsdbentry.parent }

func (ospflocallsdbentry *OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry) GetParentYangName() string { return "ospfLocalLsdbTable" }

// OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry_Ospflocallsdbtype represents advertisement format.
type OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry_Ospflocallsdbtype string

const (
    OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry_Ospflocallsdbtype_localOpaqueLink OSPFMIB_Ospflocallsdbtable_Ospflocallsdbentry_Ospflocallsdbtype = "localOpaqueLink"
)

// OSPFMIB_Ospfvirtlocallsdbtable
// The OSPF Process's link-local link state database
// for virtual links.
// 
// This table is identical to the OSPF LSDB Table
// in format, but contains only link-local Link State
// Advertisements for virtual links.  The purpose is to
// allow link-local LSAs to be displayed for each virtual
// interface.  This table is implemented to support type-9 LSAs
// that are defined in 'The OSPF Opaque LSA Option'.
type OSPFMIB_Ospfvirtlocallsdbtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry.
    Ospfvirtlocallsdbentry []OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry
}

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetFilter() yfilter.YFilter { return ospfvirtlocallsdbtable.YFilter }

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) SetFilter(yf yfilter.YFilter) { ospfvirtlocallsdbtable.YFilter = yf }

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetGoName(yname string) string {
    if yname == "ospfVirtLocalLsdbEntry" { return "Ospfvirtlocallsdbentry" }
    return ""
}

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetSegmentPath() string {
    return "ospfVirtLocalLsdbTable"
}

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfVirtLocalLsdbEntry" {
        for _, c := range ospfvirtlocallsdbtable.Ospfvirtlocallsdbentry {
            if ospfvirtlocallsdbtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry{}
        ospfvirtlocallsdbtable.Ospfvirtlocallsdbentry = append(ospfvirtlocallsdbtable.Ospfvirtlocallsdbentry, child)
        return &ospfvirtlocallsdbtable.Ospfvirtlocallsdbentry[len(ospfvirtlocallsdbtable.Ospfvirtlocallsdbentry)-1]
    }
    return nil
}

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfvirtlocallsdbtable.Ospfvirtlocallsdbentry {
        children[ospfvirtlocallsdbtable.Ospfvirtlocallsdbentry[i].GetSegmentPath()] = &ospfvirtlocallsdbtable.Ospfvirtlocallsdbentry[i]
    }
    return children
}

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetYangName() string { return "ospfVirtLocalLsdbTable" }

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) SetParent(parent types.Entity) { ospfvirtlocallsdbtable.parent = parent }

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetParent() types.Entity { return ospfvirtlocallsdbtable.parent }

func (ospfvirtlocallsdbtable *OSPFMIB_Ospfvirtlocallsdbtable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry
// A single link state advertisement.
type OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The transit area that the virtual link traverses. 
    // By definition, this is not 0.0.0.0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfvirtlocallsdbtransitarea interface{}

    // This attribute is a key. The Router ID of the virtual neighbor. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfvirtlocallsdbneighbor interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // Ospfvirtlocallsdbtype.
    Ospfvirtlocallsdbtype interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing a 32-bit identifier in IP address format; it identifies the
    // piece of the routing domain that is being described by the advertisement.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfvirtlocallsdblsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfvirtlocallsdbrouterid interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate link state advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    Ospfvirtlocallsdbsequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    Ospfvirtlocallsdbage interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that  an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred to as the Fletcher checksum. The type is interface{}
    // with range: -2147483648..2147483647.
    Ospfvirtlocallsdbchecksum interface{}

    // The entire link state advertisement, including its header. The type is
    // string with length: 1..65535.
    Ospfvirtlocallsdbadvertisement interface{}
}

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetFilter() yfilter.YFilter { return ospfvirtlocallsdbentry.YFilter }

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) SetFilter(yf yfilter.YFilter) { ospfvirtlocallsdbentry.YFilter = yf }

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetGoName(yname string) string {
    if yname == "ospfVirtLocalLsdbTransitArea" { return "Ospfvirtlocallsdbtransitarea" }
    if yname == "ospfVirtLocalLsdbNeighbor" { return "Ospfvirtlocallsdbneighbor" }
    if yname == "ospfVirtLocalLsdbType" { return "Ospfvirtlocallsdbtype" }
    if yname == "ospfVirtLocalLsdbLsid" { return "Ospfvirtlocallsdblsid" }
    if yname == "ospfVirtLocalLsdbRouterId" { return "Ospfvirtlocallsdbrouterid" }
    if yname == "ospfVirtLocalLsdbSequence" { return "Ospfvirtlocallsdbsequence" }
    if yname == "ospfVirtLocalLsdbAge" { return "Ospfvirtlocallsdbage" }
    if yname == "ospfVirtLocalLsdbChecksum" { return "Ospfvirtlocallsdbchecksum" }
    if yname == "ospfVirtLocalLsdbAdvertisement" { return "Ospfvirtlocallsdbadvertisement" }
    return ""
}

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetSegmentPath() string {
    return "ospfVirtLocalLsdbEntry" + "[ospfVirtLocalLsdbTransitArea='" + fmt.Sprintf("%v", ospfvirtlocallsdbentry.Ospfvirtlocallsdbtransitarea) + "']" + "[ospfVirtLocalLsdbNeighbor='" + fmt.Sprintf("%v", ospfvirtlocallsdbentry.Ospfvirtlocallsdbneighbor) + "']" + "[ospfVirtLocalLsdbType='" + fmt.Sprintf("%v", ospfvirtlocallsdbentry.Ospfvirtlocallsdbtype) + "']" + "[ospfVirtLocalLsdbLsid='" + fmt.Sprintf("%v", ospfvirtlocallsdbentry.Ospfvirtlocallsdblsid) + "']" + "[ospfVirtLocalLsdbRouterId='" + fmt.Sprintf("%v", ospfvirtlocallsdbentry.Ospfvirtlocallsdbrouterid) + "']"
}

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfVirtLocalLsdbTransitArea"] = ospfvirtlocallsdbentry.Ospfvirtlocallsdbtransitarea
    leafs["ospfVirtLocalLsdbNeighbor"] = ospfvirtlocallsdbentry.Ospfvirtlocallsdbneighbor
    leafs["ospfVirtLocalLsdbType"] = ospfvirtlocallsdbentry.Ospfvirtlocallsdbtype
    leafs["ospfVirtLocalLsdbLsid"] = ospfvirtlocallsdbentry.Ospfvirtlocallsdblsid
    leafs["ospfVirtLocalLsdbRouterId"] = ospfvirtlocallsdbentry.Ospfvirtlocallsdbrouterid
    leafs["ospfVirtLocalLsdbSequence"] = ospfvirtlocallsdbentry.Ospfvirtlocallsdbsequence
    leafs["ospfVirtLocalLsdbAge"] = ospfvirtlocallsdbentry.Ospfvirtlocallsdbage
    leafs["ospfVirtLocalLsdbChecksum"] = ospfvirtlocallsdbentry.Ospfvirtlocallsdbchecksum
    leafs["ospfVirtLocalLsdbAdvertisement"] = ospfvirtlocallsdbentry.Ospfvirtlocallsdbadvertisement
    return leafs
}

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetYangName() string { return "ospfVirtLocalLsdbEntry" }

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) SetParent(parent types.Entity) { ospfvirtlocallsdbentry.parent = parent }

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetParent() types.Entity { return ospfvirtlocallsdbentry.parent }

func (ospfvirtlocallsdbentry *OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry) GetParentYangName() string { return "ospfVirtLocalLsdbTable" }

// OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry_Ospfvirtlocallsdbtype represents advertisement format.
type OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry_Ospfvirtlocallsdbtype string

const (
    OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry_Ospfvirtlocallsdbtype_localOpaqueLink OSPFMIB_Ospfvirtlocallsdbtable_Ospfvirtlocallsdbentry_Ospfvirtlocallsdbtype = "localOpaqueLink"
)

// OSPFMIB_Ospfaslsdbtable
// The OSPF Process's AS-scope LSA link state database.
// The database contains the AS-scope Link State
// Advertisements from throughout the areas that
// the device is attached to.
// 
// This table is identical to the OSPF LSDB Table
// in format, but contains only AS-scope Link State
// Advertisements.  The purpose is to allow AS-scope
// LSAs to be displayed once for the router rather
// than once in each non-stub area.
type OSPFMIB_Ospfaslsdbtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single link state advertisement. The type is slice of
    // OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry.
    Ospfaslsdbentry []OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry
}

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetFilter() yfilter.YFilter { return ospfaslsdbtable.YFilter }

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) SetFilter(yf yfilter.YFilter) { ospfaslsdbtable.YFilter = yf }

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetGoName(yname string) string {
    if yname == "ospfAsLsdbEntry" { return "Ospfaslsdbentry" }
    return ""
}

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetSegmentPath() string {
    return "ospfAsLsdbTable"
}

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfAsLsdbEntry" {
        for _, c := range ospfaslsdbtable.Ospfaslsdbentry {
            if ospfaslsdbtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry{}
        ospfaslsdbtable.Ospfaslsdbentry = append(ospfaslsdbtable.Ospfaslsdbentry, child)
        return &ospfaslsdbtable.Ospfaslsdbentry[len(ospfaslsdbtable.Ospfaslsdbentry)-1]
    }
    return nil
}

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfaslsdbtable.Ospfaslsdbentry {
        children[ospfaslsdbtable.Ospfaslsdbentry[i].GetSegmentPath()] = &ospfaslsdbtable.Ospfaslsdbentry[i]
    }
    return children
}

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetYangName() string { return "ospfAsLsdbTable" }

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) SetParent(parent types.Entity) { ospfaslsdbtable.parent = parent }

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetParent() types.Entity { return ospfaslsdbtable.parent }

func (ospfaslsdbtable *OSPFMIB_Ospfaslsdbtable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry
// A single link state advertisement.
type OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // Ospfaslsdbtype.
    Ospfaslsdbtype interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing either a Router ID or an IP address;  it identifies the piece of
    // the routing domain that is being described by the advertisement. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfaslsdblsid interface{}

    // This attribute is a key. The 32-bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfaslsdbrouterid interface{}

    // The sequence number field is a signed 32-bit integer.  It starts with the
    // value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus,
    // a typical sequence number will be very negative. It is used to detect old
    // and duplicate link state advertisements.  The space of sequence numbers is
    // linearly ordered.  The larger the sequence number, the more recent the
    // advertisement. The type is interface{} with range: -2147483648..2147483647.
    Ospfaslsdbsequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: -2147483648..2147483647. Units are seconds.
    Ospfaslsdbage interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field.  The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum.  The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred to as the Fletcher checksum. The type is interface{}
    // with range: -2147483648..2147483647.
    Ospfaslsdbchecksum interface{}

    // The entire link state advertisement, including its header. The type is
    // string with length: 1..65535.
    Ospfaslsdbadvertisement interface{}
}

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetFilter() yfilter.YFilter { return ospfaslsdbentry.YFilter }

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) SetFilter(yf yfilter.YFilter) { ospfaslsdbentry.YFilter = yf }

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetGoName(yname string) string {
    if yname == "ospfAsLsdbType" { return "Ospfaslsdbtype" }
    if yname == "ospfAsLsdbLsid" { return "Ospfaslsdblsid" }
    if yname == "ospfAsLsdbRouterId" { return "Ospfaslsdbrouterid" }
    if yname == "ospfAsLsdbSequence" { return "Ospfaslsdbsequence" }
    if yname == "ospfAsLsdbAge" { return "Ospfaslsdbage" }
    if yname == "ospfAsLsdbChecksum" { return "Ospfaslsdbchecksum" }
    if yname == "ospfAsLsdbAdvertisement" { return "Ospfaslsdbadvertisement" }
    return ""
}

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetSegmentPath() string {
    return "ospfAsLsdbEntry" + "[ospfAsLsdbType='" + fmt.Sprintf("%v", ospfaslsdbentry.Ospfaslsdbtype) + "']" + "[ospfAsLsdbLsid='" + fmt.Sprintf("%v", ospfaslsdbentry.Ospfaslsdblsid) + "']" + "[ospfAsLsdbRouterId='" + fmt.Sprintf("%v", ospfaslsdbentry.Ospfaslsdbrouterid) + "']"
}

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfAsLsdbType"] = ospfaslsdbentry.Ospfaslsdbtype
    leafs["ospfAsLsdbLsid"] = ospfaslsdbentry.Ospfaslsdblsid
    leafs["ospfAsLsdbRouterId"] = ospfaslsdbentry.Ospfaslsdbrouterid
    leafs["ospfAsLsdbSequence"] = ospfaslsdbentry.Ospfaslsdbsequence
    leafs["ospfAsLsdbAge"] = ospfaslsdbentry.Ospfaslsdbage
    leafs["ospfAsLsdbChecksum"] = ospfaslsdbentry.Ospfaslsdbchecksum
    leafs["ospfAsLsdbAdvertisement"] = ospfaslsdbentry.Ospfaslsdbadvertisement
    return leafs
}

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetYangName() string { return "ospfAsLsdbEntry" }

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) SetParent(parent types.Entity) { ospfaslsdbentry.parent = parent }

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetParent() types.Entity { return ospfaslsdbentry.parent }

func (ospfaslsdbentry *OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry) GetParentYangName() string { return "ospfAsLsdbTable" }

// OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry_Ospfaslsdbtype represents advertisement format.
type OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry_Ospfaslsdbtype string

const (
    OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry_Ospfaslsdbtype_asExternalLink OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry_Ospfaslsdbtype = "asExternalLink"

    OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry_Ospfaslsdbtype_asOpaqueLink OSPFMIB_Ospfaslsdbtable_Ospfaslsdbentry_Ospfaslsdbtype = "asOpaqueLink"
)

// OSPFMIB_Ospfarealsacounttable
// This table maintains per-area, per-LSA-type counters
type OSPFMIB_Ospfarealsacounttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry with a number of link advertisements  of a given type for a given
    // area. The type is slice of
    // OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry.
    Ospfarealsacountentry []OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry
}

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetFilter() yfilter.YFilter { return ospfarealsacounttable.YFilter }

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) SetFilter(yf yfilter.YFilter) { ospfarealsacounttable.YFilter = yf }

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetGoName(yname string) string {
    if yname == "ospfAreaLsaCountEntry" { return "Ospfarealsacountentry" }
    return ""
}

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetSegmentPath() string {
    return "ospfAreaLsaCountTable"
}

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfAreaLsaCountEntry" {
        for _, c := range ospfarealsacounttable.Ospfarealsacountentry {
            if ospfarealsacounttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry{}
        ospfarealsacounttable.Ospfarealsacountentry = append(ospfarealsacounttable.Ospfarealsacountentry, child)
        return &ospfarealsacounttable.Ospfarealsacountentry[len(ospfarealsacounttable.Ospfarealsacountentry)-1]
    }
    return nil
}

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfarealsacounttable.Ospfarealsacountentry {
        children[ospfarealsacounttable.Ospfarealsacountentry[i].GetSegmentPath()] = &ospfarealsacounttable.Ospfarealsacountentry[i]
    }
    return children
}

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetBundleName() string { return "cisco_ios_xe" }

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetYangName() string { return "ospfAreaLsaCountTable" }

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) SetParent(parent types.Entity) { ospfarealsacounttable.parent = parent }

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetParent() types.Entity { return ospfarealsacounttable.parent }

func (ospfarealsacounttable *OSPFMIB_Ospfarealsacounttable) GetParentYangName() string { return "OSPF-MIB" }

// OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry
// An entry with a number of link advertisements
// 
// of a given type for a given area.
type OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This entry Area ID. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfarealsacountareaid interface{}

    // This attribute is a key. This entry LSA type. The type is
    // Ospfarealsacountlsatype.
    Ospfarealsacountlsatype interface{}

    // Number of LSAs of a given type for a given area. The type is interface{}
    // with range: 0..4294967295.
    Ospfarealsacountnumber interface{}
}

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetFilter() yfilter.YFilter { return ospfarealsacountentry.YFilter }

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) SetFilter(yf yfilter.YFilter) { ospfarealsacountentry.YFilter = yf }

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetGoName(yname string) string {
    if yname == "ospfAreaLsaCountAreaId" { return "Ospfarealsacountareaid" }
    if yname == "ospfAreaLsaCountLsaType" { return "Ospfarealsacountlsatype" }
    if yname == "ospfAreaLsaCountNumber" { return "Ospfarealsacountnumber" }
    return ""
}

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetSegmentPath() string {
    return "ospfAreaLsaCountEntry" + "[ospfAreaLsaCountAreaId='" + fmt.Sprintf("%v", ospfarealsacountentry.Ospfarealsacountareaid) + "']" + "[ospfAreaLsaCountLsaType='" + fmt.Sprintf("%v", ospfarealsacountentry.Ospfarealsacountlsatype) + "']"
}

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfAreaLsaCountAreaId"] = ospfarealsacountentry.Ospfarealsacountareaid
    leafs["ospfAreaLsaCountLsaType"] = ospfarealsacountentry.Ospfarealsacountlsatype
    leafs["ospfAreaLsaCountNumber"] = ospfarealsacountentry.Ospfarealsacountnumber
    return leafs
}

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetBundleName() string { return "cisco_ios_xe" }

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetYangName() string { return "ospfAreaLsaCountEntry" }

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) SetParent(parent types.Entity) { ospfarealsacountentry.parent = parent }

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetParent() types.Entity { return ospfarealsacountentry.parent }

func (ospfarealsacountentry *OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry) GetParentYangName() string { return "ospfAreaLsaCountTable" }

// OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype represents This entry LSA type.
type OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype string

const (
    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_routerLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "routerLink"

    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_networkLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "networkLink"

    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_summaryLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "summaryLink"

    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_asSummaryLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "asSummaryLink"

    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_multicastLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "multicastLink"

    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_nssaExternalLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "nssaExternalLink"

    OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype_areaOpaqueLink OSPFMIB_Ospfarealsacounttable_Ospfarealsacountentry_Ospfarealsacountlsatype = "areaOpaqueLink"
)

