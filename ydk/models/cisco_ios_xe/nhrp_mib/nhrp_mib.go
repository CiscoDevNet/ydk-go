// This MIB contains managed object definitions for the Next
// Hop Resolution Procol, NHRP, as defined in RFC 2332 [17].
package nhrp_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package nhrp_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:NHRP-MIB NHRP-MIB}", reflect.TypeOf(NHRPMIB{}))
    ydk.RegisterEntity("NHRP-MIB:NHRP-MIB", reflect.TypeOf(NHRPMIB{}))
}

// NHRPMIB
type NHRPMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Nhrpgeneralobjects NHRPMIB_Nhrpgeneralobjects

    // This table contains mappings between internetwork layer addresses and NBMA
    // subnetwork layer addresses.
    Nhrpcachetable NHRPMIB_Nhrpcachetable

    // This table will track Purge Request Information.
    Nhrppurgereqtable NHRPMIB_Nhrppurgereqtable

    // Information about NHRP clients (NHCs) managed by this agent.
    Nhrpclienttable NHRPMIB_Nhrpclienttable

    // A table of Registration Request Information that needs to be maintained by
    // the NHCs (clients).
    Nhrpclientregistrationtable NHRPMIB_Nhrpclientregistrationtable

    // A table of NHSes that are available for use by this NHC (client). By
    // default, the agent will add an entry to this table that corresponds to the
    // client's default router.
    Nhrpclientnhstable NHRPMIB_Nhrpclientnhstable

    // This table contains statistics collected by NHRP clients.
    Nhrpclientstattable NHRPMIB_Nhrpclientstattable

    // This table contains information for a set of NHSes associated with this
    // agent.
    Nhrpservertable NHRPMIB_Nhrpservertable

    // This table extends the nhrpCacheTable for NHSes.  If the nhrpCacheTable has
    // a row added due to an NHS or based on information regarding an NHS then a
    // row is also added in this table.  The rows in this table will be created
    // when rows in the nhrpCacheTable are created.  However, there may be rows
    // created in the nhrpCacheTable which do not have corresponding rows in this
    // table.  For example, if the nhrpCacheTable has a row added due to a Next
    // Hop Client which is co-resident on the same device as the NHS, a row will
    // not be added to this table.
    Nhrpservercachetable NHRPMIB_Nhrpservercachetable

    // A table of NHCs that are available for use by this NHS (Server).
    Nhrpservernhctable NHRPMIB_Nhrpservernhctable

    // Statistics collected by Next Hop Servers.
    Nhrpserverstattable NHRPMIB_Nhrpserverstattable
}

func (nHRPMIB *NHRPMIB) GetEntityData() *types.CommonEntityData {
    nHRPMIB.EntityData.YFilter = nHRPMIB.YFilter
    nHRPMIB.EntityData.YangName = "NHRP-MIB"
    nHRPMIB.EntityData.BundleName = "cisco_ios_xe"
    nHRPMIB.EntityData.ParentYangName = "NHRP-MIB"
    nHRPMIB.EntityData.SegmentPath = "NHRP-MIB:NHRP-MIB"
    nHRPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nHRPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nHRPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nHRPMIB.EntityData.Children = make(map[string]types.YChild)
    nHRPMIB.EntityData.Children["nhrpGeneralObjects"] = types.YChild{"Nhrpgeneralobjects", &nHRPMIB.Nhrpgeneralobjects}
    nHRPMIB.EntityData.Children["nhrpCacheTable"] = types.YChild{"Nhrpcachetable", &nHRPMIB.Nhrpcachetable}
    nHRPMIB.EntityData.Children["nhrpPurgeReqTable"] = types.YChild{"Nhrppurgereqtable", &nHRPMIB.Nhrppurgereqtable}
    nHRPMIB.EntityData.Children["nhrpClientTable"] = types.YChild{"Nhrpclienttable", &nHRPMIB.Nhrpclienttable}
    nHRPMIB.EntityData.Children["nhrpClientRegistrationTable"] = types.YChild{"Nhrpclientregistrationtable", &nHRPMIB.Nhrpclientregistrationtable}
    nHRPMIB.EntityData.Children["nhrpClientNhsTable"] = types.YChild{"Nhrpclientnhstable", &nHRPMIB.Nhrpclientnhstable}
    nHRPMIB.EntityData.Children["nhrpClientStatTable"] = types.YChild{"Nhrpclientstattable", &nHRPMIB.Nhrpclientstattable}
    nHRPMIB.EntityData.Children["nhrpServerTable"] = types.YChild{"Nhrpservertable", &nHRPMIB.Nhrpservertable}
    nHRPMIB.EntityData.Children["nhrpServerCacheTable"] = types.YChild{"Nhrpservercachetable", &nHRPMIB.Nhrpservercachetable}
    nHRPMIB.EntityData.Children["nhrpServerNhcTable"] = types.YChild{"Nhrpservernhctable", &nHRPMIB.Nhrpservernhctable}
    nHRPMIB.EntityData.Children["nhrpServerStatTable"] = types.YChild{"Nhrpserverstattable", &nHRPMIB.Nhrpserverstattable}
    nHRPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nHRPMIB.EntityData)
}

// NHRPMIB_Nhrpgeneralobjects
type NHRPMIB_Nhrpgeneralobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This scalar is used for creating rows in the nhrpClientTable and the
    // nhrpServerTable. The value of this variable is a currently unused value for
    // nhrpClientIndex and nhrpServerIndex.     The value returned when reading
    // this variable must be unique for the NHC's and NHS's indices associated
    // with this row. Subsequent attempts to read this variable must return
    // different values.  NOTE:  this object exists in the General Group because
    // it is to be used in establishing rows in the nhrpClientTable and the
    // nhrpServerTable.  In other words, the value retrieved from this object
    // could become the value of nhrpClientIndex and nhprServerIndex.  In the
    // situation of an agent re-initialization the value of this object must be
    // saved in non-volatile storage.  This variable will return the special value
    // 0 if no new rows can be created. The type is interface{} with range:
    // 0..4294967295.
    Nhrpnextindex interface{}
}

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetEntityData() *types.CommonEntityData {
    nhrpgeneralobjects.EntityData.YFilter = nhrpgeneralobjects.YFilter
    nhrpgeneralobjects.EntityData.YangName = "nhrpGeneralObjects"
    nhrpgeneralobjects.EntityData.BundleName = "cisco_ios_xe"
    nhrpgeneralobjects.EntityData.ParentYangName = "NHRP-MIB"
    nhrpgeneralobjects.EntityData.SegmentPath = "nhrpGeneralObjects"
    nhrpgeneralobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpgeneralobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpgeneralobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpgeneralobjects.EntityData.Children = make(map[string]types.YChild)
    nhrpgeneralobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    nhrpgeneralobjects.EntityData.Leafs["nhrpNextIndex"] = types.YLeaf{"Nhrpnextindex", nhrpgeneralobjects.Nhrpnextindex}
    return &(nhrpgeneralobjects.EntityData)
}

// NHRPMIB_Nhrpcachetable
// This table contains mappings between internetwork layer
// addresses and NBMA subnetwork layer addresses.
type NHRPMIB_Nhrpcachetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cached mapping between an internetwork layer address and an NBMA address.
    // Entries can be created by the network administrator using the
    // nhrpCacheRowStatus column, or they may be added dynamically based on
    // protocol operation (including NHRP, SCSP, and others, such as ATMARP). 
    // When created based by NHRP protocol operations this entry is largely based
    // on contents contained in the Client Information Entry (CIE).  Zero or more
    // Client Information Entries (CIEs) may be included in the NHRP Packet. For a
    // complete description of the CIE, refer to Section 5.2.0.1 of RFC 2332 [17].
    // The type is slice of NHRPMIB_Nhrpcachetable_Nhrpcacheentry.
    Nhrpcacheentry []NHRPMIB_Nhrpcachetable_Nhrpcacheentry
}

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetEntityData() *types.CommonEntityData {
    nhrpcachetable.EntityData.YFilter = nhrpcachetable.YFilter
    nhrpcachetable.EntityData.YangName = "nhrpCacheTable"
    nhrpcachetable.EntityData.BundleName = "cisco_ios_xe"
    nhrpcachetable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpcachetable.EntityData.SegmentPath = "nhrpCacheTable"
    nhrpcachetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpcachetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpcachetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpcachetable.EntityData.Children = make(map[string]types.YChild)
    nhrpcachetable.EntityData.Children["nhrpCacheEntry"] = types.YChild{"Nhrpcacheentry", nil}
    for i := range nhrpcachetable.Nhrpcacheentry {
        nhrpcachetable.EntityData.Children[types.GetSegmentPath(&nhrpcachetable.Nhrpcacheentry[i])] = types.YChild{"Nhrpcacheentry", &nhrpcachetable.Nhrpcacheentry[i]}
    }
    nhrpcachetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nhrpcachetable.EntityData)
}

// NHRPMIB_Nhrpcachetable_Nhrpcacheentry
// A cached mapping between an internetwork layer address
// and an NBMA address. Entries can be created by the
// network administrator using the nhrpCacheRowStatus
// column, or they may be added dynamically based on
// protocol operation (including NHRP, SCSP, and others,
// such as ATMARP).
// 
// When created based by NHRP protocol operations
// this entry is largely based on contents contained in
// the Client Information Entry (CIE).
// 
// Zero or more Client Information Entries (CIEs) may be
// included in the NHRP Packet. For a complete description
// of the CIE, refer to Section 5.2.0.1 of
// RFC 2332 [17].
type NHRPMIB_Nhrpcachetable_Nhrpcacheentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The internetwork layer address type of this Next
    // Hop Resolution Cache entry. The value of this object indicates how to
    // interpret the values of nhrpCacheInternetworkAddr and
    // nhrpCacheNextHopInternetworkAddr. The type is AddressFamilyNumbers.
    Nhrpcacheinternetworkaddrtype interface{}

    // This attribute is a key. The value of the internetwork address of the
    // destination. The type is string with length: 0..64.
    Nhrpcacheinternetworkaddr interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. An identifier for this entry that has local
    // significance within the scope of the General Group.  This identifier is
    // used here to uniquely identify this row, and also used in the
    // 'nhrpPurgeTable' for the value of the 'nhrpPurgeCacheIdentifier'. The type
    // is interface{} with range: 1..4294967295.
    Nhrpcacheindex interface{}

    // The number of bits that define the internetwork layer prefix associated
    // with the nhrpCacheInternetworkAddr. The type is interface{} with range:
    // 0..255.
    Nhrpcacheprefixlength interface{}

    // The value of the internetwork address of the next hop. The type is string
    // with length: 0..64.
    Nhrpcachenexthopinternetworkaddr interface{}

    // The NBMA address type. The value of this object indicates how to interpret
    // the values of nhrpCacheNbmaAddr and nhrpCacheNbmaSubaddr. The type is
    // AddressFamilyNumbers.
    Nhrpcachenbmaaddrtype interface{}

    // The value of the NBMA subnetwork address of the next hop. The type is
    // string with length: 0..64.
    Nhrpcachenbmaaddr interface{}

    // The value of the NBMA subaddress of the next hop. If there is no subaddress
    // concept for the NBMA address family, this value will be a zero-length OCTET
    // STRING. The type is string with length: 0..64.
    Nhrpcachenbmasubaddr interface{}

    // An indication of how this cache entry was created. The values are: 
    // 'other(1)'                   The entry was added by some                   
    // other means.  'register(2)'                In a server, added based on a   
    // client registration.  'resolveAuthoritative(3)'    In a client, added based
    // on                              receiving an Authoritative                 
    // NHRP Resolution Reply.     'resolveNonauthoritative(4)' In a client, added
    // based on                              receiving a Nonauthoritative         
    // NHRP Resolution Reply.  'transit(5)'                 In a transit server,
    // added by                              examining a forwarded NHRP           
    // packet.  'administrativelyAdded(6)'   In a client or server,               
    // manually added by the                              administrator. The      
    // StorageType of this entry is                              reflected in     
    // 'nhrpCacheStorageType'.  'atmarp(7)'                  The entry was added
    // due to an                              ATMARP.  'scsp(8)'                  
    // The entry was added due to                              SCSP.   When the
    // entry is under creation using the nhrpCacheRowStatus column, the only value
    // that can be specified by the administrator is 'administrativelyAdded'.
    // Attempting to set any other value will cause an 'inconsistentValue' error. 
    // The value cannot be modified once the entry is active. The type is
    // Nhrpcachetype.
    Nhrpcachetype interface{}

    // An indication of the state of this entry. The values are:  'incomplete(1)'
    // The client has sent a NHRP Resolution                 Request but has not
    // yet received the                 NHRP Resolution Reply.   'ackReply(2)'  
    // For a client or server, this is a                 cached valid mapping. 
    // 'nakReply(3)'   For a client or server, this is a                 cached
    // NAK mapping. The type is Nhrpcachestate.
    Nhrpcachestate interface{}

    // True(1) is returned if the value of 'nhrpCacheType' is not
    // 'administrativelyAdded'.  Since the value of 'nhrpCacheType' was not
    // configured by a user, the value of 'nhrpCacheHoldingTime' is considered
    // valid.  In other words, the value of 'nhrpCacheHoldingTime' represents the
    // Holding Time for the cache Entry.  If 'nhrpCacheType has been configured by
    // a user, (i.e. the value of 'nhrpCacheType' is 'administrativelyAdded') then
    // false(2) will be returned. This indicates that the value of
    // 'nhrpCacheHoldingTime' is undefined because this row could possibly be
    // backed up in nonvolatile storage. The type is bool.
    Nhrpcacheholdingtimevalid interface{}

    // If the value of 'nhrpCacheHoldingTimeValid is true(1) then this object
    // represents the number of seconds that the cache entry will remain in this
    // table.  When this value reaches 0 (zero) the row should be deleted.  If the
    // value of 'nhrpCacheHoldingTimeValid is false(2) then this object is
    // undefined. The type is interface{} with range: 0..65535. Units are seconds.
    Nhrpcacheholdingtime interface{}

    // The maximum transmission unit (MTU) that was negotiated or registered for
    // this entity. In other words, this is the actual MTU being used. The type is
    // interface{} with range: 0..65535.
    Nhrpcachenegotiatedmtu interface{}

    // An object which reflects the Preference value of the Client Information
    // Entry (CIE).  Zero or more Client Information Entries (CIEs) may be
    // included in the NHRP Packet.  One of the fields in the CIE is the
    // Preference.  For a complete description of the CIE, refer to Section
    // 5.2.0.1 of  RFC 2332 [17]. The type is interface{} with range: 0..255.
    Nhrpcachepreference interface{}

    // This value only has meaning when the 'nhrpCacheType' has the value of
    // 'administrativelyAdded'.  When the row is created due to being
    // 'administrativelyAdded', this object reflects whether this row is kept in
    // volatile storage and lost upon reboot or if this row is backed up by
    // non-volatile or permanent storage.  If the value of 'nhrpCacheType' has a
    // value which is not 'administrativelyAdded, then the value of this object is
    // 'other(1)'. The type is StorageType.
    Nhrpcachestoragetype interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    Nhrpcacherowstatus interface{}
}

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetEntityData() *types.CommonEntityData {
    nhrpcacheentry.EntityData.YFilter = nhrpcacheentry.YFilter
    nhrpcacheentry.EntityData.YangName = "nhrpCacheEntry"
    nhrpcacheentry.EntityData.BundleName = "cisco_ios_xe"
    nhrpcacheentry.EntityData.ParentYangName = "nhrpCacheTable"
    nhrpcacheentry.EntityData.SegmentPath = "nhrpCacheEntry" + "[nhrpCacheInternetworkAddrType='" + fmt.Sprintf("%v", nhrpcacheentry.Nhrpcacheinternetworkaddrtype) + "']" + "[nhrpCacheInternetworkAddr='" + fmt.Sprintf("%v", nhrpcacheentry.Nhrpcacheinternetworkaddr) + "']" + "[ifIndex='" + fmt.Sprintf("%v", nhrpcacheentry.Ifindex) + "']" + "[nhrpCacheIndex='" + fmt.Sprintf("%v", nhrpcacheentry.Nhrpcacheindex) + "']"
    nhrpcacheentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpcacheentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpcacheentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpcacheentry.EntityData.Children = make(map[string]types.YChild)
    nhrpcacheentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nhrpcacheentry.EntityData.Leafs["nhrpCacheInternetworkAddrType"] = types.YLeaf{"Nhrpcacheinternetworkaddrtype", nhrpcacheentry.Nhrpcacheinternetworkaddrtype}
    nhrpcacheentry.EntityData.Leafs["nhrpCacheInternetworkAddr"] = types.YLeaf{"Nhrpcacheinternetworkaddr", nhrpcacheentry.Nhrpcacheinternetworkaddr}
    nhrpcacheentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", nhrpcacheentry.Ifindex}
    nhrpcacheentry.EntityData.Leafs["nhrpCacheIndex"] = types.YLeaf{"Nhrpcacheindex", nhrpcacheentry.Nhrpcacheindex}
    nhrpcacheentry.EntityData.Leafs["nhrpCachePrefixLength"] = types.YLeaf{"Nhrpcacheprefixlength", nhrpcacheentry.Nhrpcacheprefixlength}
    nhrpcacheentry.EntityData.Leafs["nhrpCacheNextHopInternetworkAddr"] = types.YLeaf{"Nhrpcachenexthopinternetworkaddr", nhrpcacheentry.Nhrpcachenexthopinternetworkaddr}
    nhrpcacheentry.EntityData.Leafs["nhrpCacheNbmaAddrType"] = types.YLeaf{"Nhrpcachenbmaaddrtype", nhrpcacheentry.Nhrpcachenbmaaddrtype}
    nhrpcacheentry.EntityData.Leafs["nhrpCacheNbmaAddr"] = types.YLeaf{"Nhrpcachenbmaaddr", nhrpcacheentry.Nhrpcachenbmaaddr}
    nhrpcacheentry.EntityData.Leafs["nhrpCacheNbmaSubaddr"] = types.YLeaf{"Nhrpcachenbmasubaddr", nhrpcacheentry.Nhrpcachenbmasubaddr}
    nhrpcacheentry.EntityData.Leafs["nhrpCacheType"] = types.YLeaf{"Nhrpcachetype", nhrpcacheentry.Nhrpcachetype}
    nhrpcacheentry.EntityData.Leafs["nhrpCacheState"] = types.YLeaf{"Nhrpcachestate", nhrpcacheentry.Nhrpcachestate}
    nhrpcacheentry.EntityData.Leafs["nhrpCacheHoldingTimeValid"] = types.YLeaf{"Nhrpcacheholdingtimevalid", nhrpcacheentry.Nhrpcacheholdingtimevalid}
    nhrpcacheentry.EntityData.Leafs["nhrpCacheHoldingTime"] = types.YLeaf{"Nhrpcacheholdingtime", nhrpcacheentry.Nhrpcacheholdingtime}
    nhrpcacheentry.EntityData.Leafs["nhrpCacheNegotiatedMtu"] = types.YLeaf{"Nhrpcachenegotiatedmtu", nhrpcacheentry.Nhrpcachenegotiatedmtu}
    nhrpcacheentry.EntityData.Leafs["nhrpCachePreference"] = types.YLeaf{"Nhrpcachepreference", nhrpcacheentry.Nhrpcachepreference}
    nhrpcacheentry.EntityData.Leafs["nhrpCacheStorageType"] = types.YLeaf{"Nhrpcachestoragetype", nhrpcacheentry.Nhrpcachestoragetype}
    nhrpcacheentry.EntityData.Leafs["nhrpCacheRowStatus"] = types.YLeaf{"Nhrpcacherowstatus", nhrpcacheentry.Nhrpcacherowstatus}
    return &(nhrpcacheentry.EntityData)
}

// NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachestate represents                 cached NAK mapping.
type NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachestate string

const (
    NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachestate_incomplete NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachestate = "incomplete"

    NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachestate_ackReply NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachestate = "ackReply"

    NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachestate_nakReply NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachestate = "nakReply"
)

// NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype represents The value cannot be modified once the entry is active.
type NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype string

const (
    NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype_other NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype = "other"

    NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype_register NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype = "register"

    NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype_resolveAuthoritative NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype = "resolveAuthoritative"

    NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype_resoveNonauthoritative NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype = "resoveNonauthoritative"

    NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype_transit NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype = "transit"

    NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype_administrativelyAdded NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype = "administrativelyAdded"

    NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype_atmarp NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype = "atmarp"

    NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype_scsp NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcachetype = "scsp"
)

// NHRPMIB_Nhrppurgereqtable
// This table will track Purge Request Information.
type NHRPMIB_Nhrppurgereqtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information regarding a Purge Request. The type is slice of
    // NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry.
    Nhrppurgereqentry []NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry
}

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetEntityData() *types.CommonEntityData {
    nhrppurgereqtable.EntityData.YFilter = nhrppurgereqtable.YFilter
    nhrppurgereqtable.EntityData.YangName = "nhrpPurgeReqTable"
    nhrppurgereqtable.EntityData.BundleName = "cisco_ios_xe"
    nhrppurgereqtable.EntityData.ParentYangName = "NHRP-MIB"
    nhrppurgereqtable.EntityData.SegmentPath = "nhrpPurgeReqTable"
    nhrppurgereqtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrppurgereqtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrppurgereqtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrppurgereqtable.EntityData.Children = make(map[string]types.YChild)
    nhrppurgereqtable.EntityData.Children["nhrpPurgeReqEntry"] = types.YChild{"Nhrppurgereqentry", nil}
    for i := range nhrppurgereqtable.Nhrppurgereqentry {
        nhrppurgereqtable.EntityData.Children[types.GetSegmentPath(&nhrppurgereqtable.Nhrppurgereqentry[i])] = types.YChild{"Nhrppurgereqentry", &nhrppurgereqtable.Nhrppurgereqentry[i]}
    }
    nhrppurgereqtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nhrppurgereqtable.EntityData)
}

// NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry
// Information regarding a Purge Request.
type NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index for this entry that has local
    // significance within the scope of this table. The type is interface{} with
    // range: 1..4294967295.
    Nhrppurgeindex interface{}

    // This object identifies which row in 'nhrpCacheTable' is being purged.  This
    // object should have the same value as the 'nhrpCacheIndex' in the
    // 'nhrpCacheTable'. The type is interface{} with range: 1..4294967295.
    Nhrppurgecacheidentifier interface{}

    // In the case of NHRP Purge Requests, this specifies the equivalence class of
    // addresses which match the first 'Prefix Length' bit positions of the Client
    // Protocol Address specified in the Client Information Entry (CIE). The type
    // is interface{} with range: 0..255.
    Nhrppurgeprefixlength interface{}

    // The Request ID used in the purge request. The type is interface{} with
    // range: 0..4294967295.
    Nhrppurgerequestid interface{}

    // An indication of whether this Purge Request has the 'N' Bit cleared (off).
    // The type is bool.
    Nhrppurgereplyexpected interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    Nhrppurgerowstatus interface{}
}

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetEntityData() *types.CommonEntityData {
    nhrppurgereqentry.EntityData.YFilter = nhrppurgereqentry.YFilter
    nhrppurgereqentry.EntityData.YangName = "nhrpPurgeReqEntry"
    nhrppurgereqentry.EntityData.BundleName = "cisco_ios_xe"
    nhrppurgereqentry.EntityData.ParentYangName = "nhrpPurgeReqTable"
    nhrppurgereqentry.EntityData.SegmentPath = "nhrpPurgeReqEntry" + "[nhrpPurgeIndex='" + fmt.Sprintf("%v", nhrppurgereqentry.Nhrppurgeindex) + "']"
    nhrppurgereqentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrppurgereqentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrppurgereqentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrppurgereqentry.EntityData.Children = make(map[string]types.YChild)
    nhrppurgereqentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nhrppurgereqentry.EntityData.Leafs["nhrpPurgeIndex"] = types.YLeaf{"Nhrppurgeindex", nhrppurgereqentry.Nhrppurgeindex}
    nhrppurgereqentry.EntityData.Leafs["nhrpPurgeCacheIdentifier"] = types.YLeaf{"Nhrppurgecacheidentifier", nhrppurgereqentry.Nhrppurgecacheidentifier}
    nhrppurgereqentry.EntityData.Leafs["nhrpPurgePrefixLength"] = types.YLeaf{"Nhrppurgeprefixlength", nhrppurgereqentry.Nhrppurgeprefixlength}
    nhrppurgereqentry.EntityData.Leafs["nhrpPurgeRequestID"] = types.YLeaf{"Nhrppurgerequestid", nhrppurgereqentry.Nhrppurgerequestid}
    nhrppurgereqentry.EntityData.Leafs["nhrpPurgeReplyExpected"] = types.YLeaf{"Nhrppurgereplyexpected", nhrppurgereqentry.Nhrppurgereplyexpected}
    nhrppurgereqentry.EntityData.Leafs["nhrpPurgeRowStatus"] = types.YLeaf{"Nhrppurgerowstatus", nhrppurgereqentry.Nhrppurgerowstatus}
    return &(nhrppurgereqentry.EntityData)
}

// NHRPMIB_Nhrpclienttable
// Information about NHRP clients (NHCs) managed by this
// agent.
type NHRPMIB_Nhrpclienttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single NHC. The type is slice of
    // NHRPMIB_Nhrpclienttable_Nhrpcliententry.
    Nhrpcliententry []NHRPMIB_Nhrpclienttable_Nhrpcliententry
}

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetEntityData() *types.CommonEntityData {
    nhrpclienttable.EntityData.YFilter = nhrpclienttable.YFilter
    nhrpclienttable.EntityData.YangName = "nhrpClientTable"
    nhrpclienttable.EntityData.BundleName = "cisco_ios_xe"
    nhrpclienttable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpclienttable.EntityData.SegmentPath = "nhrpClientTable"
    nhrpclienttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpclienttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpclienttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpclienttable.EntityData.Children = make(map[string]types.YChild)
    nhrpclienttable.EntityData.Children["nhrpClientEntry"] = types.YChild{"Nhrpcliententry", nil}
    for i := range nhrpclienttable.Nhrpcliententry {
        nhrpclienttable.EntityData.Children[types.GetSegmentPath(&nhrpclienttable.Nhrpcliententry[i])] = types.YChild{"Nhrpcliententry", &nhrpclienttable.Nhrpcliententry[i]}
    }
    nhrpclienttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nhrpclienttable.EntityData)
}

// NHRPMIB_Nhrpclienttable_Nhrpcliententry
// Information about a single NHC.
type NHRPMIB_Nhrpclienttable_Nhrpcliententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An identifier for the NHRP client that is unique
    // within the scope of this agent.  The 'nhrpNextIndex' value should be
    // consulted (read), prior to creating a row in this table, and the value
    // returned from reading 'nhrpNextIndex' should be used as this object's
    // value. The type is interface{} with range: 1..4294967295.
    Nhrpclientindex interface{}

    // The type of the internetwork layer address of this client. This object
    // indicates how the value of nhrpClientInternetworkAddr is to be interpreted.
    // The type is AddressFamilyNumbers.
    Nhrpclientinternetworkaddrtype interface{}

    // The value of the internetwork layer address of this client. The type is
    // string with length: 0..64.
    Nhrpclientinternetworkaddr interface{}

    // The type of the NBMA subnetwork address of this client. This object
    // indicates how the values of nhrpClientNbmaAddr and nhrpClientNbmaSubaddr
    // are to be interpreted. The type is AddressFamilyNumbers.
    Nhrpclientnbmaaddrtype interface{}

    // The NBMA subnetwork address of this client. The type is string with length:
    // 0..64.
    Nhrpclientnbmaaddr interface{}

    // The NBMA subaddress of this client. For NBMA address families without a
    // subaddress concept, this will be a zero-length OCTET STRING. The type is
    // string with length: 0..64.
    Nhrpclientnbmasubaddr interface{}

    // The number of seconds that the client will wait before timing out an NHRP
    // initial request.  This object only has meaning for the initial timeout
    // period. The type is interface{} with range: 1..900. Units are seconds.
    Nhrpclientinitialrequesttimeout interface{}

    // The number of times the client will retry the registration request before
    // failure. A value of 0 means don't retry. A value of 65535 means retry
    // forever. The type is interface{} with range: 0..65535.
    Nhrpclientregistrationrequestretries interface{}

    // The number of times the client will retry the resolution request before
    // failure. A value of 0 means don't retry. A value of 65535 means retry
    // forever. The type is interface{} with range: 0..65535.
    Nhrpclientresolutionrequestretries interface{}

    // The number of times the client will retry a purge request before failure. A
    // value of 0 means don't retry. A value of 65535 means retry forever. The
    // type is interface{} with range: 0..65535.
    Nhrpclientpurgerequestretries interface{}

    // The default maximum transmission unit (MTU) of the LIS/LAG which this
    // client should use. This object will be initialized by the agent to the
    // default MTU of the LIS/LAG (which is 9180) unless a different MTU value is
    // specified during creation of this Client. The type is interface{} with
    // range: 0..65535.
    Nhrpclientdefaultmtu interface{}

    // The hold time the client will register. The type is interface{} with range:
    // 0..65535. Units are seconds.
    Nhrpclientholdtime interface{}

    // The Request ID used to register this client with its server. According to
    // Section 5.2.3 of the NHRP Specification, RFC 2332 [17], the Request ID must
    // be kept in non-volatile storage, so that if an NHC crashes and 
    // re-initializes, it will use a different  Request ID during the registration
    // process when reregistering with the same NHS. The type is interface{} with
    // range: 0..4294967295.
    Nhrpclientrequestid interface{}

    // This object defines whether this row is kept in volatile storage and lost
    // upon a Client crash or reboot situation, or if this row is backed up by
    // nonvolatile or permanent storage. The type is StorageType.
    Nhrpclientstoragetype interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    Nhrpclientrowstatus interface{}
}

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetEntityData() *types.CommonEntityData {
    nhrpcliententry.EntityData.YFilter = nhrpcliententry.YFilter
    nhrpcliententry.EntityData.YangName = "nhrpClientEntry"
    nhrpcliententry.EntityData.BundleName = "cisco_ios_xe"
    nhrpcliententry.EntityData.ParentYangName = "nhrpClientTable"
    nhrpcliententry.EntityData.SegmentPath = "nhrpClientEntry" + "[nhrpClientIndex='" + fmt.Sprintf("%v", nhrpcliententry.Nhrpclientindex) + "']"
    nhrpcliententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpcliententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpcliententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpcliententry.EntityData.Children = make(map[string]types.YChild)
    nhrpcliententry.EntityData.Leafs = make(map[string]types.YLeaf)
    nhrpcliententry.EntityData.Leafs["nhrpClientIndex"] = types.YLeaf{"Nhrpclientindex", nhrpcliententry.Nhrpclientindex}
    nhrpcliententry.EntityData.Leafs["nhrpClientInternetworkAddrType"] = types.YLeaf{"Nhrpclientinternetworkaddrtype", nhrpcliententry.Nhrpclientinternetworkaddrtype}
    nhrpcliententry.EntityData.Leafs["nhrpClientInternetworkAddr"] = types.YLeaf{"Nhrpclientinternetworkaddr", nhrpcliententry.Nhrpclientinternetworkaddr}
    nhrpcliententry.EntityData.Leafs["nhrpClientNbmaAddrType"] = types.YLeaf{"Nhrpclientnbmaaddrtype", nhrpcliententry.Nhrpclientnbmaaddrtype}
    nhrpcliententry.EntityData.Leafs["nhrpClientNbmaAddr"] = types.YLeaf{"Nhrpclientnbmaaddr", nhrpcliententry.Nhrpclientnbmaaddr}
    nhrpcliententry.EntityData.Leafs["nhrpClientNbmaSubaddr"] = types.YLeaf{"Nhrpclientnbmasubaddr", nhrpcliententry.Nhrpclientnbmasubaddr}
    nhrpcliententry.EntityData.Leafs["nhrpClientInitialRequestTimeout"] = types.YLeaf{"Nhrpclientinitialrequesttimeout", nhrpcliententry.Nhrpclientinitialrequesttimeout}
    nhrpcliententry.EntityData.Leafs["nhrpClientRegistrationRequestRetries"] = types.YLeaf{"Nhrpclientregistrationrequestretries", nhrpcliententry.Nhrpclientregistrationrequestretries}
    nhrpcliententry.EntityData.Leafs["nhrpClientResolutionRequestRetries"] = types.YLeaf{"Nhrpclientresolutionrequestretries", nhrpcliententry.Nhrpclientresolutionrequestretries}
    nhrpcliententry.EntityData.Leafs["nhrpClientPurgeRequestRetries"] = types.YLeaf{"Nhrpclientpurgerequestretries", nhrpcliententry.Nhrpclientpurgerequestretries}
    nhrpcliententry.EntityData.Leafs["nhrpClientDefaultMtu"] = types.YLeaf{"Nhrpclientdefaultmtu", nhrpcliententry.Nhrpclientdefaultmtu}
    nhrpcliententry.EntityData.Leafs["nhrpClientHoldTime"] = types.YLeaf{"Nhrpclientholdtime", nhrpcliententry.Nhrpclientholdtime}
    nhrpcliententry.EntityData.Leafs["nhrpClientRequestID"] = types.YLeaf{"Nhrpclientrequestid", nhrpcliententry.Nhrpclientrequestid}
    nhrpcliententry.EntityData.Leafs["nhrpClientStorageType"] = types.YLeaf{"Nhrpclientstoragetype", nhrpcliententry.Nhrpclientstoragetype}
    nhrpcliententry.EntityData.Leafs["nhrpClientRowStatus"] = types.YLeaf{"Nhrpclientrowstatus", nhrpcliententry.Nhrpclientrowstatus}
    return &(nhrpcliententry.EntityData)
}

// NHRPMIB_Nhrpclientregistrationtable
// A table of Registration Request Information that
// needs to be maintained by the NHCs (clients).
type NHRPMIB_Nhrpclientregistrationtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An NHC needs to maintain registration request information between the NHC
    // and the NHS.  An entry in this table represents information for a single
    // registration request. The type is slice of
    // NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry.
    Nhrpclientregistrationentry []NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry
}

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetEntityData() *types.CommonEntityData {
    nhrpclientregistrationtable.EntityData.YFilter = nhrpclientregistrationtable.YFilter
    nhrpclientregistrationtable.EntityData.YangName = "nhrpClientRegistrationTable"
    nhrpclientregistrationtable.EntityData.BundleName = "cisco_ios_xe"
    nhrpclientregistrationtable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpclientregistrationtable.EntityData.SegmentPath = "nhrpClientRegistrationTable"
    nhrpclientregistrationtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpclientregistrationtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpclientregistrationtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpclientregistrationtable.EntityData.Children = make(map[string]types.YChild)
    nhrpclientregistrationtable.EntityData.Children["nhrpClientRegistrationEntry"] = types.YChild{"Nhrpclientregistrationentry", nil}
    for i := range nhrpclientregistrationtable.Nhrpclientregistrationentry {
        nhrpclientregistrationtable.EntityData.Children[types.GetSegmentPath(&nhrpclientregistrationtable.Nhrpclientregistrationentry[i])] = types.YChild{"Nhrpclientregistrationentry", &nhrpclientregistrationtable.Nhrpclientregistrationentry[i]}
    }
    nhrpclientregistrationtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nhrpclientregistrationtable.EntityData)
}

// NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry
// An NHC needs to maintain registration request information
// between the NHC and the NHS.  An entry in this table
// represents information for a single registration request.
type NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to nhrp_mib.NHRPMIB_Nhrpclienttable_Nhrpcliententry_Nhrpclientindex
    Nhrpclientindex interface{}

    // This attribute is a key. An identifier for this entry such that it
    // identifies a specific Registration Request from the NHC represented by the
    // nhrpClientIndex. The type is interface{} with range: 1..4294967295.
    Nhrpclientregindex interface{}

    // The Uniqueness indicator for this Registration Request. If this object has
    // the value of requestUnique(1), then the Uniqueness bit is set in the the
    // NHRP Registration Request represented by this row.  The value cannot be
    // changed once the row is created. The type is Nhrpclientreguniqueness.
    Nhrpclientreguniqueness interface{}

    // The registration state of this client. The values are: 'other(1)'          
    // The state of the registration                        request is not one of 
    // 'registering',                        'ackRegisterReply' or                
    // 'nakRegisterReply'.  'registering(2)'        A registration request has    
    // been issued and a registration                         reply is expected. 
    // 'ackRegisterReply(3)'   A positive registration reply                      
    // has been received.  'nakRegisterReply(4)'   The client has received a      
    // negative registration                         reply (NAK). The type is
    // Nhrpclientregstate.
    Nhrpclientregstate interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    Nhrpclientregrowstatus interface{}
}

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetEntityData() *types.CommonEntityData {
    nhrpclientregistrationentry.EntityData.YFilter = nhrpclientregistrationentry.YFilter
    nhrpclientregistrationentry.EntityData.YangName = "nhrpClientRegistrationEntry"
    nhrpclientregistrationentry.EntityData.BundleName = "cisco_ios_xe"
    nhrpclientregistrationentry.EntityData.ParentYangName = "nhrpClientRegistrationTable"
    nhrpclientregistrationentry.EntityData.SegmentPath = "nhrpClientRegistrationEntry" + "[nhrpClientIndex='" + fmt.Sprintf("%v", nhrpclientregistrationentry.Nhrpclientindex) + "']" + "[nhrpClientRegIndex='" + fmt.Sprintf("%v", nhrpclientregistrationentry.Nhrpclientregindex) + "']"
    nhrpclientregistrationentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpclientregistrationentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpclientregistrationentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpclientregistrationentry.EntityData.Children = make(map[string]types.YChild)
    nhrpclientregistrationentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nhrpclientregistrationentry.EntityData.Leafs["nhrpClientIndex"] = types.YLeaf{"Nhrpclientindex", nhrpclientregistrationentry.Nhrpclientindex}
    nhrpclientregistrationentry.EntityData.Leafs["nhrpClientRegIndex"] = types.YLeaf{"Nhrpclientregindex", nhrpclientregistrationentry.Nhrpclientregindex}
    nhrpclientregistrationentry.EntityData.Leafs["nhrpClientRegUniqueness"] = types.YLeaf{"Nhrpclientreguniqueness", nhrpclientregistrationentry.Nhrpclientreguniqueness}
    nhrpclientregistrationentry.EntityData.Leafs["nhrpClientRegState"] = types.YLeaf{"Nhrpclientregstate", nhrpclientregistrationentry.Nhrpclientregstate}
    nhrpclientregistrationentry.EntityData.Leafs["nhrpClientRegRowStatus"] = types.YLeaf{"Nhrpclientregrowstatus", nhrpclientregistrationentry.Nhrpclientregrowstatus}
    return &(nhrpclientregistrationentry.EntityData)
}

// NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientregstate represents                         reply (NAK).
type NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientregstate string

const (
    NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientregstate_other NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientregstate = "other"

    NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientregstate_registering NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientregstate = "registering"

    NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientregstate_ackRegisterReply NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientregstate = "ackRegisterReply"

    NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientregstate_nakRegisterReply NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientregstate = "nakRegisterReply"
)

// NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientreguniqueness represents be changed once the row is created.
type NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientreguniqueness string

const (
    NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientreguniqueness_requestUnique NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientreguniqueness = "requestUnique"

    NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientreguniqueness_requestNotUnique NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry_Nhrpclientreguniqueness = "requestNotUnique"
)

// NHRPMIB_Nhrpclientnhstable
// A table of NHSes that are available for use by this NHC
// (client). By default, the agent will add an entry to this
// table that corresponds to the client's default router.
type NHRPMIB_Nhrpclientnhstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An NHS that may be used by an NHC. The type is slice of
    // NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry.
    Nhrpclientnhsentry []NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry
}

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetEntityData() *types.CommonEntityData {
    nhrpclientnhstable.EntityData.YFilter = nhrpclientnhstable.YFilter
    nhrpclientnhstable.EntityData.YangName = "nhrpClientNhsTable"
    nhrpclientnhstable.EntityData.BundleName = "cisco_ios_xe"
    nhrpclientnhstable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpclientnhstable.EntityData.SegmentPath = "nhrpClientNhsTable"
    nhrpclientnhstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpclientnhstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpclientnhstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpclientnhstable.EntityData.Children = make(map[string]types.YChild)
    nhrpclientnhstable.EntityData.Children["nhrpClientNhsEntry"] = types.YChild{"Nhrpclientnhsentry", nil}
    for i := range nhrpclientnhstable.Nhrpclientnhsentry {
        nhrpclientnhstable.EntityData.Children[types.GetSegmentPath(&nhrpclientnhstable.Nhrpclientnhsentry[i])] = types.YChild{"Nhrpclientnhsentry", &nhrpclientnhstable.Nhrpclientnhsentry[i]}
    }
    nhrpclientnhstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nhrpclientnhstable.EntityData)
}

// NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry
// An NHS that may be used by an NHC.
type NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to nhrp_mib.NHRPMIB_Nhrpclienttable_Nhrpcliententry_Nhrpclientindex
    Nhrpclientindex interface{}

    // This attribute is a key. An identifier for an NHS available to an NHC. The
    // type is interface{} with range: 1..4294967295.
    Nhrpclientnhsindex interface{}

    // The type of the internetwork layer address of the NHRP server represented
    // in this entry. This object indicates how the value of
    // nhrpClientNhsInternetworkAddr is to be interpreted. The type is
    // AddressFamilyNumbers.
    Nhrpclientnhsinternetworkaddrtype interface{}

    // The value of the destination internetwork layer address of the NHRP server
    // represented by this    entry.  If this value is not known, this will be a
    // zero-length OCTET STRING. The type is string with length: 0..64.
    Nhrpclientnhsinternetworkaddr interface{}

    // The type of the NBMA subnetwork address of the NHRP Server represented by
    // this entry. This object indicates how the values of nhrpClientNhsNbmaAddr
    // and nhrpClientNhsNbmaSubaddr are to be interpreted. The type is
    // AddressFamilyNumbers.
    Nhrpclientnhsnbmaaddrtype interface{}

    // The NBMA subnetwork address of the NHS. The type of the address is
    // indicated by the corresponding value of nhrpClientNhsNbmaAddrType. The type
    // is string with length: 0..64.
    Nhrpclientnhsnbmaaddr interface{}

    // The NBMA subaddress of the NHS. For NMBA address families that do not have
    // the concept of subaddress,      this will be a zero-length OCTET STRING.
    // The type is string with length: 0..64.
    Nhrpclientnhsnbmasubaddr interface{}

    // An indication of whether this NHS is in use by the NHC. The type is bool.
    Nhrpclientnhsinuse interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    Nhrpclientnhsrowstatus interface{}
}

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetEntityData() *types.CommonEntityData {
    nhrpclientnhsentry.EntityData.YFilter = nhrpclientnhsentry.YFilter
    nhrpclientnhsentry.EntityData.YangName = "nhrpClientNhsEntry"
    nhrpclientnhsentry.EntityData.BundleName = "cisco_ios_xe"
    nhrpclientnhsentry.EntityData.ParentYangName = "nhrpClientNhsTable"
    nhrpclientnhsentry.EntityData.SegmentPath = "nhrpClientNhsEntry" + "[nhrpClientIndex='" + fmt.Sprintf("%v", nhrpclientnhsentry.Nhrpclientindex) + "']" + "[nhrpClientNhsIndex='" + fmt.Sprintf("%v", nhrpclientnhsentry.Nhrpclientnhsindex) + "']"
    nhrpclientnhsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpclientnhsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpclientnhsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpclientnhsentry.EntityData.Children = make(map[string]types.YChild)
    nhrpclientnhsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nhrpclientnhsentry.EntityData.Leafs["nhrpClientIndex"] = types.YLeaf{"Nhrpclientindex", nhrpclientnhsentry.Nhrpclientindex}
    nhrpclientnhsentry.EntityData.Leafs["nhrpClientNhsIndex"] = types.YLeaf{"Nhrpclientnhsindex", nhrpclientnhsentry.Nhrpclientnhsindex}
    nhrpclientnhsentry.EntityData.Leafs["nhrpClientNhsInternetworkAddrType"] = types.YLeaf{"Nhrpclientnhsinternetworkaddrtype", nhrpclientnhsentry.Nhrpclientnhsinternetworkaddrtype}
    nhrpclientnhsentry.EntityData.Leafs["nhrpClientNhsInternetworkAddr"] = types.YLeaf{"Nhrpclientnhsinternetworkaddr", nhrpclientnhsentry.Nhrpclientnhsinternetworkaddr}
    nhrpclientnhsentry.EntityData.Leafs["nhrpClientNhsNbmaAddrType"] = types.YLeaf{"Nhrpclientnhsnbmaaddrtype", nhrpclientnhsentry.Nhrpclientnhsnbmaaddrtype}
    nhrpclientnhsentry.EntityData.Leafs["nhrpClientNhsNbmaAddr"] = types.YLeaf{"Nhrpclientnhsnbmaaddr", nhrpclientnhsentry.Nhrpclientnhsnbmaaddr}
    nhrpclientnhsentry.EntityData.Leafs["nhrpClientNhsNbmaSubaddr"] = types.YLeaf{"Nhrpclientnhsnbmasubaddr", nhrpclientnhsentry.Nhrpclientnhsnbmasubaddr}
    nhrpclientnhsentry.EntityData.Leafs["nhrpClientNhsInUse"] = types.YLeaf{"Nhrpclientnhsinuse", nhrpclientnhsentry.Nhrpclientnhsinuse}
    nhrpclientnhsentry.EntityData.Leafs["nhrpClientNhsRowStatus"] = types.YLeaf{"Nhrpclientnhsrowstatus", nhrpclientnhsentry.Nhrpclientnhsrowstatus}
    return &(nhrpclientnhsentry.EntityData)
}

// NHRPMIB_Nhrpclientstattable
// This table contains statistics collected by NHRP
// clients.
type NHRPMIB_Nhrpclientstattable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics collected by a NHRP client. The type is slice of
    // NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry.
    Nhrpclientstatentry []NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry
}

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetEntityData() *types.CommonEntityData {
    nhrpclientstattable.EntityData.YFilter = nhrpclientstattable.YFilter
    nhrpclientstattable.EntityData.YangName = "nhrpClientStatTable"
    nhrpclientstattable.EntityData.BundleName = "cisco_ios_xe"
    nhrpclientstattable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpclientstattable.EntityData.SegmentPath = "nhrpClientStatTable"
    nhrpclientstattable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpclientstattable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpclientstattable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpclientstattable.EntityData.Children = make(map[string]types.YChild)
    nhrpclientstattable.EntityData.Children["nhrpClientStatEntry"] = types.YChild{"Nhrpclientstatentry", nil}
    for i := range nhrpclientstattable.Nhrpclientstatentry {
        nhrpclientstattable.EntityData.Children[types.GetSegmentPath(&nhrpclientstattable.Nhrpclientstatentry[i])] = types.YChild{"Nhrpclientstatentry", &nhrpclientstattable.Nhrpclientstatentry[i]}
    }
    nhrpclientstattable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nhrpclientstattable.EntityData)
}

// NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry
// Statistics collected by a NHRP client.
type NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to nhrp_mib.NHRPMIB_Nhrpclienttable_Nhrpcliententry_Nhrpclientindex
    Nhrpclientindex interface{}

    // The number of NHRP Resolution Requests transmitted by this client. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpclientstattxresolvereq interface{}

    // The number of positively acknowledged NHRP Resolution Replies received by
    // this client.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Client
    // re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpclientstatrxresolvereplyack interface{}

    // The number of NAKed NHRP Resolution Replies received by this client that
    // contained the code indicating 'Administratively Prohibited'.  
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpclientstatrxresolvereplynakprohibited interface{}

    // The number of NAKed NHRP Resolution Replies received by this client that
    // contained the code indicating 'Insufficient Resources'.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, at NHRP Client re-initialization and at other times as indicated by
    // the value of nhrpClientStatDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Nhrpclientstatrxresolvereplynakinsufresources interface{}

    // The number of NAKed NHRP Resolution Replies received by this client that
    // contained the code indicating 'No Internetworking Layer Address to NBMA
    // Address Binding Exists'.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, at NHRP Client
    // re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpclientstatrxresolvereplynaknobinding interface{}

    // The number of NAKed NHRP Resolution Replies received by this client that
    // contained the code indicating 'Binding Exists But Is Not Unique'. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpclientstatrxresolvereplynaknotunique interface{}

    // The number of NHRP Registration Requests transmitted by this client. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpclientstattxregisterreq interface{}

    // The number of positively acknowledged NHRP Registration Replies received by
    // this client.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Client
    // re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpclientstatrxregisterack interface{}

    // The number of NAKed NHRP Registration Replies received by this client that
    // contained the code indicating 'Administratively Prohibited'. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpclientstatrxregisternakprohibited interface{}

    // The number of NAKed NHRP Registration Replies received by this client that
    // contained the code indicating 'Insufficient Resources'.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, at NHRP Client re-initialization and at other times as indicated by
    // the value of nhrpClientStatDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Nhrpclientstatrxregisternakinsufresources interface{}

    // The number of NAKed NHRP Registration Replies received by this client that
    // contained the code indicating 'Unique Internetworking Layer Address Already
    // Registered'.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Client
    // re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpclientstatrxregisternakalreadyreg interface{}

    // The number of NHRP Purge Requests received by this client.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, at NHRP Client re-initialization and at other times as
    // indicated by the value of nhrpClientStatDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Nhrpclientstatrxpurgereq interface{}

    // The number of NHRP Purge Requests transmitted by this client. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpclientstattxpurgereq interface{}

    // The number of NHRP Purge Replies received by this client.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, at NHRP Client re-initialization and at other times as
    // indicated by the value of nhrpClientStatDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Nhrpclientstatrxpurgereply interface{}

    // The number of NHRP Purge Replies transmitted by this client. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpclientstattxpurgereply interface{}

    // The number of NHRP Error Indication packets transmitted by this client. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpclientstattxerrorindication interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'Unrecognized Extension'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Client re-initialization and at other times as indicated by the value
    // of nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpclientstatrxerrunrecognizedextension interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'NHRP Loop Detected'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Client re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpclientstatrxerrloopdetected interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'Protocol Address Unreachable'.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, at NHRP Client re-initialization and at other times as indicated by
    // the value of nhrpClientStatDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Nhrpclientstatrxerrprotoaddrunreachable interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'Protocol Error'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Client re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpclientstatrxerrprotoerror interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'NHRP SDU Size  Exceeded'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Client re-initialization and at other times as indicated by the value
    // of nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpclientstatrxerrsdusizeexceeded interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'Invalid Extension'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Client re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpclientstatrxerrinvalidextension interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'Authentication Failure'.      Discontinuities in the value
    // of this counter can occur at re-initialization of the management system, at
    // NHRP Client re-initialization and at other times as indicated by the value
    // of nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpclientstatrxerrauthenticationfailure interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'Hop Count Exceeded'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Client re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpclientstatrxerrhopcountexceeded interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this Client's counters suffered a discontinuity.  If no such
    // discontinuities have occurred since the last re-initialization of the local
    // management subsystem or the NHRP Client re-initialization associated with
    // this entry, then this object contains a zero value. The type is interface{}
    // with range: 0..4294967295.
    Nhrpclientstatdiscontinuitytime interface{}
}

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetEntityData() *types.CommonEntityData {
    nhrpclientstatentry.EntityData.YFilter = nhrpclientstatentry.YFilter
    nhrpclientstatentry.EntityData.YangName = "nhrpClientStatEntry"
    nhrpclientstatentry.EntityData.BundleName = "cisco_ios_xe"
    nhrpclientstatentry.EntityData.ParentYangName = "nhrpClientStatTable"
    nhrpclientstatentry.EntityData.SegmentPath = "nhrpClientStatEntry" + "[nhrpClientIndex='" + fmt.Sprintf("%v", nhrpclientstatentry.Nhrpclientindex) + "']"
    nhrpclientstatentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpclientstatentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpclientstatentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpclientstatentry.EntityData.Children = make(map[string]types.YChild)
    nhrpclientstatentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nhrpclientstatentry.EntityData.Leafs["nhrpClientIndex"] = types.YLeaf{"Nhrpclientindex", nhrpclientstatentry.Nhrpclientindex}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatTxResolveReq"] = types.YLeaf{"Nhrpclientstattxresolvereq", nhrpclientstatentry.Nhrpclientstattxresolvereq}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxResolveReplyAck"] = types.YLeaf{"Nhrpclientstatrxresolvereplyack", nhrpclientstatentry.Nhrpclientstatrxresolvereplyack}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxResolveReplyNakProhibited"] = types.YLeaf{"Nhrpclientstatrxresolvereplynakprohibited", nhrpclientstatentry.Nhrpclientstatrxresolvereplynakprohibited}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxResolveReplyNakInsufResources"] = types.YLeaf{"Nhrpclientstatrxresolvereplynakinsufresources", nhrpclientstatentry.Nhrpclientstatrxresolvereplynakinsufresources}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxResolveReplyNakNoBinding"] = types.YLeaf{"Nhrpclientstatrxresolvereplynaknobinding", nhrpclientstatentry.Nhrpclientstatrxresolvereplynaknobinding}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxResolveReplyNakNotUnique"] = types.YLeaf{"Nhrpclientstatrxresolvereplynaknotunique", nhrpclientstatentry.Nhrpclientstatrxresolvereplynaknotunique}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatTxRegisterReq"] = types.YLeaf{"Nhrpclientstattxregisterreq", nhrpclientstatentry.Nhrpclientstattxregisterreq}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxRegisterAck"] = types.YLeaf{"Nhrpclientstatrxregisterack", nhrpclientstatentry.Nhrpclientstatrxregisterack}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxRegisterNakProhibited"] = types.YLeaf{"Nhrpclientstatrxregisternakprohibited", nhrpclientstatentry.Nhrpclientstatrxregisternakprohibited}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxRegisterNakInsufResources"] = types.YLeaf{"Nhrpclientstatrxregisternakinsufresources", nhrpclientstatentry.Nhrpclientstatrxregisternakinsufresources}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxRegisterNakAlreadyReg"] = types.YLeaf{"Nhrpclientstatrxregisternakalreadyreg", nhrpclientstatentry.Nhrpclientstatrxregisternakalreadyreg}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxPurgeReq"] = types.YLeaf{"Nhrpclientstatrxpurgereq", nhrpclientstatentry.Nhrpclientstatrxpurgereq}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatTxPurgeReq"] = types.YLeaf{"Nhrpclientstattxpurgereq", nhrpclientstatentry.Nhrpclientstattxpurgereq}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxPurgeReply"] = types.YLeaf{"Nhrpclientstatrxpurgereply", nhrpclientstatentry.Nhrpclientstatrxpurgereply}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatTxPurgeReply"] = types.YLeaf{"Nhrpclientstattxpurgereply", nhrpclientstatentry.Nhrpclientstattxpurgereply}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatTxErrorIndication"] = types.YLeaf{"Nhrpclientstattxerrorindication", nhrpclientstatentry.Nhrpclientstattxerrorindication}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxErrUnrecognizedExtension"] = types.YLeaf{"Nhrpclientstatrxerrunrecognizedextension", nhrpclientstatentry.Nhrpclientstatrxerrunrecognizedextension}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxErrLoopDetected"] = types.YLeaf{"Nhrpclientstatrxerrloopdetected", nhrpclientstatentry.Nhrpclientstatrxerrloopdetected}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxErrProtoAddrUnreachable"] = types.YLeaf{"Nhrpclientstatrxerrprotoaddrunreachable", nhrpclientstatentry.Nhrpclientstatrxerrprotoaddrunreachable}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxErrProtoError"] = types.YLeaf{"Nhrpclientstatrxerrprotoerror", nhrpclientstatentry.Nhrpclientstatrxerrprotoerror}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxErrSduSizeExceeded"] = types.YLeaf{"Nhrpclientstatrxerrsdusizeexceeded", nhrpclientstatentry.Nhrpclientstatrxerrsdusizeexceeded}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxErrInvalidExtension"] = types.YLeaf{"Nhrpclientstatrxerrinvalidextension", nhrpclientstatentry.Nhrpclientstatrxerrinvalidextension}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxErrAuthenticationFailure"] = types.YLeaf{"Nhrpclientstatrxerrauthenticationfailure", nhrpclientstatentry.Nhrpclientstatrxerrauthenticationfailure}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatRxErrHopCountExceeded"] = types.YLeaf{"Nhrpclientstatrxerrhopcountexceeded", nhrpclientstatentry.Nhrpclientstatrxerrhopcountexceeded}
    nhrpclientstatentry.EntityData.Leafs["nhrpClientStatDiscontinuityTime"] = types.YLeaf{"Nhrpclientstatdiscontinuitytime", nhrpclientstatentry.Nhrpclientstatdiscontinuitytime}
    return &(nhrpclientstatentry.EntityData)
}

// NHRPMIB_Nhrpservertable
// This table contains information for a set of NHSes
// associated with this agent.
type NHRPMIB_Nhrpservertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single NHS. The type is slice of
    // NHRPMIB_Nhrpservertable_Nhrpserverentry.
    Nhrpserverentry []NHRPMIB_Nhrpservertable_Nhrpserverentry
}

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetEntityData() *types.CommonEntityData {
    nhrpservertable.EntityData.YFilter = nhrpservertable.YFilter
    nhrpservertable.EntityData.YangName = "nhrpServerTable"
    nhrpservertable.EntityData.BundleName = "cisco_ios_xe"
    nhrpservertable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpservertable.EntityData.SegmentPath = "nhrpServerTable"
    nhrpservertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpservertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpservertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpservertable.EntityData.Children = make(map[string]types.YChild)
    nhrpservertable.EntityData.Children["nhrpServerEntry"] = types.YChild{"Nhrpserverentry", nil}
    for i := range nhrpservertable.Nhrpserverentry {
        nhrpservertable.EntityData.Children[types.GetSegmentPath(&nhrpservertable.Nhrpserverentry[i])] = types.YChild{"Nhrpserverentry", &nhrpservertable.Nhrpserverentry[i]}
    }
    nhrpservertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nhrpservertable.EntityData)
}

// NHRPMIB_Nhrpservertable_Nhrpserverentry
// Information about a single NHS.
type NHRPMIB_Nhrpservertable_Nhrpserverentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An identifier for the server that is unique within
    // the scope of this agent. The type is interface{} with range: 1..4294967295.
    Nhrpserverindex interface{}

    // The type of the internetwork layer address of this server. This object is
    // used to interpret the value of nhrpServerInternetworkAddr. The type is
    // AddressFamilyNumbers.
    Nhrpserverinternetworkaddrtype interface{}

    // The value of the internetwork layer address of this server. The type is
    // string with length: 0..64.
    Nhrpserverinternetworkaddr interface{}

    // The type of the NBMA subnetwork address of this server. This object is used
    // to interpret the value of nhrpServerNbmaAddr. The type is
    // AddressFamilyNumbers.
    Nhrpservernbmaaddrtype interface{}

    // The value of the NBMA subnetwork address of this server. The type is string
    // with length: 0..64.
    Nhrpservernbmaaddr interface{}

    // The value of the NBMA subaddress of this server. For NBMA address families
    // without a subaddress concept, this will be a zero-length OCTET STRING. The
    // type is string with length: 0..64.
    Nhrpservernbmasubaddr interface{}

    // This object defines whether this row is kept in volatile storage and lost
    // upon a Server crash or reboot situation, or if this row is backed up by
    // nonvolatile or permanent storage. The type is StorageType.
    Nhrpserverstoragetype interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    Nhrpserverrowstatus interface{}
}

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetEntityData() *types.CommonEntityData {
    nhrpserverentry.EntityData.YFilter = nhrpserverentry.YFilter
    nhrpserverentry.EntityData.YangName = "nhrpServerEntry"
    nhrpserverentry.EntityData.BundleName = "cisco_ios_xe"
    nhrpserverentry.EntityData.ParentYangName = "nhrpServerTable"
    nhrpserverentry.EntityData.SegmentPath = "nhrpServerEntry" + "[nhrpServerIndex='" + fmt.Sprintf("%v", nhrpserverentry.Nhrpserverindex) + "']"
    nhrpserverentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpserverentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpserverentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpserverentry.EntityData.Children = make(map[string]types.YChild)
    nhrpserverentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nhrpserverentry.EntityData.Leafs["nhrpServerIndex"] = types.YLeaf{"Nhrpserverindex", nhrpserverentry.Nhrpserverindex}
    nhrpserverentry.EntityData.Leafs["nhrpServerInternetworkAddrType"] = types.YLeaf{"Nhrpserverinternetworkaddrtype", nhrpserverentry.Nhrpserverinternetworkaddrtype}
    nhrpserverentry.EntityData.Leafs["nhrpServerInternetworkAddr"] = types.YLeaf{"Nhrpserverinternetworkaddr", nhrpserverentry.Nhrpserverinternetworkaddr}
    nhrpserverentry.EntityData.Leafs["nhrpServerNbmaAddrType"] = types.YLeaf{"Nhrpservernbmaaddrtype", nhrpserverentry.Nhrpservernbmaaddrtype}
    nhrpserverentry.EntityData.Leafs["nhrpServerNbmaAddr"] = types.YLeaf{"Nhrpservernbmaaddr", nhrpserverentry.Nhrpservernbmaaddr}
    nhrpserverentry.EntityData.Leafs["nhrpServerNbmaSubaddr"] = types.YLeaf{"Nhrpservernbmasubaddr", nhrpserverentry.Nhrpservernbmasubaddr}
    nhrpserverentry.EntityData.Leafs["nhrpServerStorageType"] = types.YLeaf{"Nhrpserverstoragetype", nhrpserverentry.Nhrpserverstoragetype}
    nhrpserverentry.EntityData.Leafs["nhrpServerRowStatus"] = types.YLeaf{"Nhrpserverrowstatus", nhrpserverentry.Nhrpserverrowstatus}
    return &(nhrpserverentry.EntityData)
}

// NHRPMIB_Nhrpservercachetable
// This table extends the nhrpCacheTable for
// NHSes.  If the nhrpCacheTable has a row added due to
// an NHS or based on information regarding an NHS then
// a row is also added in this table.
// 
// The rows in this table will be created when rows in
// the nhrpCacheTable are created.  However, there may
// be rows created in the nhrpCacheTable which do not
// have corresponding rows in this table.  For example,
// if the nhrpCacheTable has a row added due to a Next
// Hop Client which is co-resident on the same device
// as the NHS, a row will not be added to this table.
type NHRPMIB_Nhrpservercachetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Additional information kept by a NHS for a relevant Next Hop Resolution
    // Cache entry. The type is slice of
    // NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry.
    Nhrpservercacheentry []NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry
}

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetEntityData() *types.CommonEntityData {
    nhrpservercachetable.EntityData.YFilter = nhrpservercachetable.YFilter
    nhrpservercachetable.EntityData.YangName = "nhrpServerCacheTable"
    nhrpservercachetable.EntityData.BundleName = "cisco_ios_xe"
    nhrpservercachetable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpservercachetable.EntityData.SegmentPath = "nhrpServerCacheTable"
    nhrpservercachetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpservercachetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpservercachetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpservercachetable.EntityData.Children = make(map[string]types.YChild)
    nhrpservercachetable.EntityData.Children["nhrpServerCacheEntry"] = types.YChild{"Nhrpservercacheentry", nil}
    for i := range nhrpservercachetable.Nhrpservercacheentry {
        nhrpservercachetable.EntityData.Children[types.GetSegmentPath(&nhrpservercachetable.Nhrpservercacheentry[i])] = types.YChild{"Nhrpservercacheentry", &nhrpservercachetable.Nhrpservercacheentry[i]}
    }
    nhrpservercachetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nhrpservercachetable.EntityData)
}

// NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry
// Additional information kept by a NHS for a relevant
// Next Hop Resolution Cache entry.
type NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is AddressFamilyNumbers. Refers to
    // nhrp_mib.NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcacheinternetworkaddrtype
    Nhrpcacheinternetworkaddrtype interface{}

    // This attribute is a key. The type is string with length: 0..64. Refers to
    // nhrp_mib.NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcacheinternetworkaddr
    Nhrpcacheinternetworkaddr interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to nhrp_mib.NHRPMIB_Nhrpcachetable_Nhrpcacheentry_Nhrpcacheindex
    Nhrpcacheindex interface{}

    // An indication of whether this cache entry is authoritative, which means the
    // entry was added because of a direct registration request with this server
    // or by Server Cache Synchronization Protocol (SCSP) from an authoritative
    // source. The type is bool.
    Nhrpservercacheauthoritative interface{}

    // The Uniqueness indicator for this cache entry used in duplicate address
    // detection. This value cannot be changed after the entry is active. The type
    // is bool.
    Nhrpservercacheuniqueness interface{}
}

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetEntityData() *types.CommonEntityData {
    nhrpservercacheentry.EntityData.YFilter = nhrpservercacheentry.YFilter
    nhrpservercacheentry.EntityData.YangName = "nhrpServerCacheEntry"
    nhrpservercacheentry.EntityData.BundleName = "cisco_ios_xe"
    nhrpservercacheentry.EntityData.ParentYangName = "nhrpServerCacheTable"
    nhrpservercacheentry.EntityData.SegmentPath = "nhrpServerCacheEntry" + "[nhrpCacheInternetworkAddrType='" + fmt.Sprintf("%v", nhrpservercacheentry.Nhrpcacheinternetworkaddrtype) + "']" + "[nhrpCacheInternetworkAddr='" + fmt.Sprintf("%v", nhrpservercacheentry.Nhrpcacheinternetworkaddr) + "']" + "[ifIndex='" + fmt.Sprintf("%v", nhrpservercacheentry.Ifindex) + "']" + "[nhrpCacheIndex='" + fmt.Sprintf("%v", nhrpservercacheentry.Nhrpcacheindex) + "']"
    nhrpservercacheentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpservercacheentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpservercacheentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpservercacheentry.EntityData.Children = make(map[string]types.YChild)
    nhrpservercacheentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nhrpservercacheentry.EntityData.Leafs["nhrpCacheInternetworkAddrType"] = types.YLeaf{"Nhrpcacheinternetworkaddrtype", nhrpservercacheentry.Nhrpcacheinternetworkaddrtype}
    nhrpservercacheentry.EntityData.Leafs["nhrpCacheInternetworkAddr"] = types.YLeaf{"Nhrpcacheinternetworkaddr", nhrpservercacheentry.Nhrpcacheinternetworkaddr}
    nhrpservercacheentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", nhrpservercacheentry.Ifindex}
    nhrpservercacheentry.EntityData.Leafs["nhrpCacheIndex"] = types.YLeaf{"Nhrpcacheindex", nhrpservercacheentry.Nhrpcacheindex}
    nhrpservercacheentry.EntityData.Leafs["nhrpServerCacheAuthoritative"] = types.YLeaf{"Nhrpservercacheauthoritative", nhrpservercacheentry.Nhrpservercacheauthoritative}
    nhrpservercacheentry.EntityData.Leafs["nhrpServerCacheUniqueness"] = types.YLeaf{"Nhrpservercacheuniqueness", nhrpservercacheentry.Nhrpservercacheuniqueness}
    return &(nhrpservercacheentry.EntityData)
}

// NHRPMIB_Nhrpservernhctable
// A table of NHCs that are available for use by this NHS
// (Server).
type NHRPMIB_Nhrpservernhctable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An NHC that may be used by an NHS. The type is slice of
    // NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry.
    Nhrpservernhcentry []NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry
}

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetEntityData() *types.CommonEntityData {
    nhrpservernhctable.EntityData.YFilter = nhrpservernhctable.YFilter
    nhrpservernhctable.EntityData.YangName = "nhrpServerNhcTable"
    nhrpservernhctable.EntityData.BundleName = "cisco_ios_xe"
    nhrpservernhctable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpservernhctable.EntityData.SegmentPath = "nhrpServerNhcTable"
    nhrpservernhctable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpservernhctable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpservernhctable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpservernhctable.EntityData.Children = make(map[string]types.YChild)
    nhrpservernhctable.EntityData.Children["nhrpServerNhcEntry"] = types.YChild{"Nhrpservernhcentry", nil}
    for i := range nhrpservernhctable.Nhrpservernhcentry {
        nhrpservernhctable.EntityData.Children[types.GetSegmentPath(&nhrpservernhctable.Nhrpservernhcentry[i])] = types.YChild{"Nhrpservernhcentry", &nhrpservernhctable.Nhrpservernhcentry[i]}
    }
    nhrpservernhctable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nhrpservernhctable.EntityData)
}

// NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry
// An NHC that may be used by an NHS.
type NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to nhrp_mib.NHRPMIB_Nhrpservertable_Nhrpserverentry_Nhrpserverindex
    Nhrpserverindex interface{}

    // This attribute is a key. An identifier for an NHC available to an NHS. The
    // type is interface{} with range: 1..4294967295.
    Nhrpservernhcindex interface{}

    // The number of bits that define the internetwork layer prefix associated
    // with the nhrpServerNhcInternetworkAddr. The type is interface{} with range:
    // 0..255.
    Nhrpservernhcprefixlength interface{}

    // The type of the internetwork layer address of the NHRP Client represented
    // in this entry. This object indicates how the value of
    // nhrpServerNhcInternetworkAddr is to be interpreted. The type is
    // AddressFamilyNumbers.
    Nhrpservernhcinternetworkaddrtype interface{}

    // The value of the internetwork layer address of the NHRP Client represented
    // by this entry.  If this value is not known, this will be a zero-length
    // OCTET STRING. The type is string with length: 0..64.
    Nhrpservernhcinternetworkaddr interface{}

    // The type of the NBMA subnetwork address of the NHRP Client represented by
    // this entry. This object indicates how the values of nhrpServerNhcNbmaAddr
    // and nhrpServerNhcNbmaSubaddr are to be interpreted. The type is
    // AddressFamilyNumbers.
    Nhrpservernhcnbmaaddrtype interface{}

    // The NBMA subnetwork address of the NHC. The type of the address is
    // indicated by the corresponding value of nhrpServerNbmaAddrType. The type is
    // string with length: 0..64.
    Nhrpservernhcnbmaaddr interface{}

    // The NBMA subaddress of the NHC. For NMBA address familes that do not have
    // the concept of subaddress, this will be a zero-length OCTET STRING. The
    // type is string with length: 0..64.
    Nhrpservernhcnbmasubaddr interface{}

    // An indication of whether this NHC is in use by the NHS. The type is bool.
    Nhrpservernhcinuse interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    Nhrpservernhcrowstatus interface{}
}

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetEntityData() *types.CommonEntityData {
    nhrpservernhcentry.EntityData.YFilter = nhrpservernhcentry.YFilter
    nhrpservernhcentry.EntityData.YangName = "nhrpServerNhcEntry"
    nhrpservernhcentry.EntityData.BundleName = "cisco_ios_xe"
    nhrpservernhcentry.EntityData.ParentYangName = "nhrpServerNhcTable"
    nhrpservernhcentry.EntityData.SegmentPath = "nhrpServerNhcEntry" + "[nhrpServerIndex='" + fmt.Sprintf("%v", nhrpservernhcentry.Nhrpserverindex) + "']" + "[nhrpServerNhcIndex='" + fmt.Sprintf("%v", nhrpservernhcentry.Nhrpservernhcindex) + "']"
    nhrpservernhcentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpservernhcentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpservernhcentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpservernhcentry.EntityData.Children = make(map[string]types.YChild)
    nhrpservernhcentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nhrpservernhcentry.EntityData.Leafs["nhrpServerIndex"] = types.YLeaf{"Nhrpserverindex", nhrpservernhcentry.Nhrpserverindex}
    nhrpservernhcentry.EntityData.Leafs["nhrpServerNhcIndex"] = types.YLeaf{"Nhrpservernhcindex", nhrpservernhcentry.Nhrpservernhcindex}
    nhrpservernhcentry.EntityData.Leafs["nhrpServerNhcPrefixLength"] = types.YLeaf{"Nhrpservernhcprefixlength", nhrpservernhcentry.Nhrpservernhcprefixlength}
    nhrpservernhcentry.EntityData.Leafs["nhrpServerNhcInternetworkAddrType"] = types.YLeaf{"Nhrpservernhcinternetworkaddrtype", nhrpservernhcentry.Nhrpservernhcinternetworkaddrtype}
    nhrpservernhcentry.EntityData.Leafs["nhrpServerNhcInternetworkAddr"] = types.YLeaf{"Nhrpservernhcinternetworkaddr", nhrpservernhcentry.Nhrpservernhcinternetworkaddr}
    nhrpservernhcentry.EntityData.Leafs["nhrpServerNhcNbmaAddrType"] = types.YLeaf{"Nhrpservernhcnbmaaddrtype", nhrpservernhcentry.Nhrpservernhcnbmaaddrtype}
    nhrpservernhcentry.EntityData.Leafs["nhrpServerNhcNbmaAddr"] = types.YLeaf{"Nhrpservernhcnbmaaddr", nhrpservernhcentry.Nhrpservernhcnbmaaddr}
    nhrpservernhcentry.EntityData.Leafs["nhrpServerNhcNbmaSubaddr"] = types.YLeaf{"Nhrpservernhcnbmasubaddr", nhrpservernhcentry.Nhrpservernhcnbmasubaddr}
    nhrpservernhcentry.EntityData.Leafs["nhrpServerNhcInUse"] = types.YLeaf{"Nhrpservernhcinuse", nhrpservernhcentry.Nhrpservernhcinuse}
    nhrpservernhcentry.EntityData.Leafs["nhrpServerNhcRowStatus"] = types.YLeaf{"Nhrpservernhcrowstatus", nhrpservernhcentry.Nhrpservernhcrowstatus}
    return &(nhrpservernhcentry.EntityData)
}

// NHRPMIB_Nhrpserverstattable
// Statistics collected by Next Hop Servers.
type NHRPMIB_Nhrpserverstattable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics for a particular NHS. The statistics are broken into received
    // (Rx), transmitted (Tx) and forwarded (Fw).  Forwarded (Fw) would be done by
    // a transit NHS. The type is slice of
    // NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry.
    Nhrpserverstatentry []NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry
}

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetEntityData() *types.CommonEntityData {
    nhrpserverstattable.EntityData.YFilter = nhrpserverstattable.YFilter
    nhrpserverstattable.EntityData.YangName = "nhrpServerStatTable"
    nhrpserverstattable.EntityData.BundleName = "cisco_ios_xe"
    nhrpserverstattable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpserverstattable.EntityData.SegmentPath = "nhrpServerStatTable"
    nhrpserverstattable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpserverstattable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpserverstattable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpserverstattable.EntityData.Children = make(map[string]types.YChild)
    nhrpserverstattable.EntityData.Children["nhrpServerStatEntry"] = types.YChild{"Nhrpserverstatentry", nil}
    for i := range nhrpserverstattable.Nhrpserverstatentry {
        nhrpserverstattable.EntityData.Children[types.GetSegmentPath(&nhrpserverstattable.Nhrpserverstatentry[i])] = types.YChild{"Nhrpserverstatentry", &nhrpserverstattable.Nhrpserverstatentry[i]}
    }
    nhrpserverstattable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nhrpserverstattable.EntityData)
}

// NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry
// Statistics for a particular NHS. The statistics are
// broken into received (Rx), transmitted (Tx)
// and forwarded (Fw).  Forwarded (Fw) would be done
// by a transit NHS.
type NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to nhrp_mib.NHRPMIB_Nhrpservertable_Nhrpserverentry_Nhrpserverindex
    Nhrpserverindex interface{}

    // The number of NHRP Resolution Requests received by this server. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Server re-initialization and at other
    // times as indicated by the value of nhrpServerStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpserverstatrxresolvereq interface{}

    // The number of positively acknowledged NHRP Resolution Replies transmitted
    // by this server.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxresolvereplyack interface{}

    // The number of NAKed NHRP Resolution Replies transmitted by this server with
    // the code 'Administratively Prohibited'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxresolvereplynakprohibited interface{}

    // The number of NAKed NHRP Resolution Replies transmitted by this server with
    // the code 'Insufficient Resources'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxresolvereplynakinsufresources interface{}

    // The number of NAKed NHRP Resolution Replies transmitted by this server with
    // the code 'No Internetworking Layer Address to NBMA Address Binding Exists'.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Server re-initialization and at other
    // times as indicated by the value of nhrpServerStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpserverstattxresolvereplynaknobinding interface{}

    // The number of NAKed NHRP Resolution Replies transmitted by this server with
    // the code 'Binding Exists But Is Not Unique'.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxresolvereplynaknotunique interface{}

    // The number of NHRP Registration Requests received by this server. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Server re-initialization and at other
    // times as indicated by the value of nhrpServerStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpserverstatrxregisterreq interface{}

    // The number of positively acknowledged NHRP Registration Replies transmitted
    // by this server.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxregisterack interface{}

    // The number of NAKed NHRP Registration Replies transmitted by this server
    // with the code 'Administratively Prohibited'.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxregisternakprohibited interface{}

    // The number of NAKed NHRP Registration Replies transmitted by this server
    // with the code 'Insufficient Resources'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxregisternakinsufresources interface{}

    // The number of NAKed NHRP Registration Replies transmitted by this server
    // with the code 'Unique Internetworking Layer Address Already Registered'. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Server re-initialization and at other
    // times as indicated by the value of nhrpServerStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpserverstattxregisternakalreadyreg interface{}

    // The number of NHRP Purge Requests received by this server.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, at NHRP Server re-initialization and at other times as
    // indicated by the value of nhrpServerStatDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Nhrpserverstatrxpurgereq interface{}

    // The number of NHRP Purge Requests transmitted by this server. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Server re-initialization and at other
    // times as indicated by the value of nhrpServerStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpserverstattxpurgereq interface{}

    // The number of NHRP Purge Replies received by this server.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, at NHRP Server re-initialization and at other times as
    // indicated by the value of nhrpServerStatDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Nhrpserverstatrxpurgereply interface{}

    // The number of NHRP Purge Replies transmitted by this server. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Server re-initialization and at other
    // times as indicated by the value of nhrpServerStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Nhrpserverstattxpurgereply interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code  'Unrecognized Extension'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatrxerrunrecognizedextension interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'NHRP Loop Detected'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatrxerrloopdetected interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'Protocol Address Unreachable'.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, at NHRP Server re-initialization and at other times as indicated by
    // the value of nhrpServerStatDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Nhrpserverstatrxerrprotoaddrunreachable interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'Protocol Error'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatrxerrprotoerror interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'NHRP SDU Size Exceeded'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatrxerrsdusizeexceeded interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'Invalid Extension'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatrxerrinvalidextension interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'Invalid Resolution Reply Received'.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, at NHRP Server re-initialization and at other times as indicated by
    // the value of nhrpServerStatDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Nhrpserverstatrxerrinvalidresreplyreceived interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'Authentication Failure'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatrxerrauthenticationfailure interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'Hop Count Exceeded'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatrxerrhopcountexceeded interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'Unrecognized Extension'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxerrunrecognizedextension interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'NHRP Loop Detected'.     Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxerrloopdetected interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'Protocol Address Unreachable'.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, at NHRP Server re-initialization and at other times as indicated by
    // the value of nhrpServerStatDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Nhrpserverstattxerrprotoaddrunreachable interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'Protocol Error'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxerrprotoerror interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'NHRP SDU Size Exceeded'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxerrsdusizeexceeded interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code  'Invalid Extension'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxerrinvalidextension interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'Authentication Failure'.     Discontinuities in the value
    // of this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxerrauthenticationfailure interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'Hop Count Exceeded'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstattxerrhopcountexceeded interface{}

    // The number of NHRP Resolution Requests forwarded by this server acting as a
    // transit NHS.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatfwresolvereq interface{}

    // The number of NHRP Resolution Replies forwarded by this server acting as a
    // transit NHS.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatfwresolvereply interface{}

    // The number of NHRP Registration Requests forwarded by this server acting as
    // a transit NHS.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatfwregisterreq interface{}

    // The number of NHRP Registration Replies forwarded by this server acting as
    // a transit NHS. Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatfwregisterreply interface{}

    // The number of NHRP Purge Requests forwarded by this server acting as a
    // transit NHS.   Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatfwpurgereq interface{}

    // The number of NHRP Purge Replies forwarded by this server acting as a
    // transit NHS.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatfwpurgereply interface{}

    // The number of NHRP Error Indication packets forwarded by this server acting
    // as a transit NHS.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Nhrpserverstatfwerrorindication interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this Server's counters suffered a discontinuity.  If no such
    // discontinuities have occurred since the last re-initialization of the   
    // local management subsystem or the NHRP Server re-initialization associated
    // with this entry, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    Nhrpserverstatdiscontinuitytime interface{}
}

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetEntityData() *types.CommonEntityData {
    nhrpserverstatentry.EntityData.YFilter = nhrpserverstatentry.YFilter
    nhrpserverstatentry.EntityData.YangName = "nhrpServerStatEntry"
    nhrpserverstatentry.EntityData.BundleName = "cisco_ios_xe"
    nhrpserverstatentry.EntityData.ParentYangName = "nhrpServerStatTable"
    nhrpserverstatentry.EntityData.SegmentPath = "nhrpServerStatEntry" + "[nhrpServerIndex='" + fmt.Sprintf("%v", nhrpserverstatentry.Nhrpserverindex) + "']"
    nhrpserverstatentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpserverstatentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpserverstatentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpserverstatentry.EntityData.Children = make(map[string]types.YChild)
    nhrpserverstatentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nhrpserverstatentry.EntityData.Leafs["nhrpServerIndex"] = types.YLeaf{"Nhrpserverindex", nhrpserverstatentry.Nhrpserverindex}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatRxResolveReq"] = types.YLeaf{"Nhrpserverstatrxresolvereq", nhrpserverstatentry.Nhrpserverstatrxresolvereq}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxResolveReplyAck"] = types.YLeaf{"Nhrpserverstattxresolvereplyack", nhrpserverstatentry.Nhrpserverstattxresolvereplyack}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxResolveReplyNakProhibited"] = types.YLeaf{"Nhrpserverstattxresolvereplynakprohibited", nhrpserverstatentry.Nhrpserverstattxresolvereplynakprohibited}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxResolveReplyNakInsufResources"] = types.YLeaf{"Nhrpserverstattxresolvereplynakinsufresources", nhrpserverstatentry.Nhrpserverstattxresolvereplynakinsufresources}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxResolveReplyNakNoBinding"] = types.YLeaf{"Nhrpserverstattxresolvereplynaknobinding", nhrpserverstatentry.Nhrpserverstattxresolvereplynaknobinding}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxResolveReplyNakNotUnique"] = types.YLeaf{"Nhrpserverstattxresolvereplynaknotunique", nhrpserverstatentry.Nhrpserverstattxresolvereplynaknotunique}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatRxRegisterReq"] = types.YLeaf{"Nhrpserverstatrxregisterreq", nhrpserverstatentry.Nhrpserverstatrxregisterreq}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxRegisterAck"] = types.YLeaf{"Nhrpserverstattxregisterack", nhrpserverstatentry.Nhrpserverstattxregisterack}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxRegisterNakProhibited"] = types.YLeaf{"Nhrpserverstattxregisternakprohibited", nhrpserverstatentry.Nhrpserverstattxregisternakprohibited}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxRegisterNakInsufResources"] = types.YLeaf{"Nhrpserverstattxregisternakinsufresources", nhrpserverstatentry.Nhrpserverstattxregisternakinsufresources}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxRegisterNakAlreadyReg"] = types.YLeaf{"Nhrpserverstattxregisternakalreadyreg", nhrpserverstatentry.Nhrpserverstattxregisternakalreadyreg}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatRxPurgeReq"] = types.YLeaf{"Nhrpserverstatrxpurgereq", nhrpserverstatentry.Nhrpserverstatrxpurgereq}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxPurgeReq"] = types.YLeaf{"Nhrpserverstattxpurgereq", nhrpserverstatentry.Nhrpserverstattxpurgereq}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatRxPurgeReply"] = types.YLeaf{"Nhrpserverstatrxpurgereply", nhrpserverstatentry.Nhrpserverstatrxpurgereply}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxPurgeReply"] = types.YLeaf{"Nhrpserverstattxpurgereply", nhrpserverstatentry.Nhrpserverstattxpurgereply}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatRxErrUnrecognizedExtension"] = types.YLeaf{"Nhrpserverstatrxerrunrecognizedextension", nhrpserverstatentry.Nhrpserverstatrxerrunrecognizedextension}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatRxErrLoopDetected"] = types.YLeaf{"Nhrpserverstatrxerrloopdetected", nhrpserverstatentry.Nhrpserverstatrxerrloopdetected}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatRxErrProtoAddrUnreachable"] = types.YLeaf{"Nhrpserverstatrxerrprotoaddrunreachable", nhrpserverstatentry.Nhrpserverstatrxerrprotoaddrunreachable}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatRxErrProtoError"] = types.YLeaf{"Nhrpserverstatrxerrprotoerror", nhrpserverstatentry.Nhrpserverstatrxerrprotoerror}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatRxErrSduSizeExceeded"] = types.YLeaf{"Nhrpserverstatrxerrsdusizeexceeded", nhrpserverstatentry.Nhrpserverstatrxerrsdusizeexceeded}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatRxErrInvalidExtension"] = types.YLeaf{"Nhrpserverstatrxerrinvalidextension", nhrpserverstatentry.Nhrpserverstatrxerrinvalidextension}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatRxErrInvalidResReplyReceived"] = types.YLeaf{"Nhrpserverstatrxerrinvalidresreplyreceived", nhrpserverstatentry.Nhrpserverstatrxerrinvalidresreplyreceived}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatRxErrAuthenticationFailure"] = types.YLeaf{"Nhrpserverstatrxerrauthenticationfailure", nhrpserverstatentry.Nhrpserverstatrxerrauthenticationfailure}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatRxErrHopCountExceeded"] = types.YLeaf{"Nhrpserverstatrxerrhopcountexceeded", nhrpserverstatentry.Nhrpserverstatrxerrhopcountexceeded}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxErrUnrecognizedExtension"] = types.YLeaf{"Nhrpserverstattxerrunrecognizedextension", nhrpserverstatentry.Nhrpserverstattxerrunrecognizedextension}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxErrLoopDetected"] = types.YLeaf{"Nhrpserverstattxerrloopdetected", nhrpserverstatentry.Nhrpserverstattxerrloopdetected}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxErrProtoAddrUnreachable"] = types.YLeaf{"Nhrpserverstattxerrprotoaddrunreachable", nhrpserverstatentry.Nhrpserverstattxerrprotoaddrunreachable}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxErrProtoError"] = types.YLeaf{"Nhrpserverstattxerrprotoerror", nhrpserverstatentry.Nhrpserverstattxerrprotoerror}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxErrSduSizeExceeded"] = types.YLeaf{"Nhrpserverstattxerrsdusizeexceeded", nhrpserverstatentry.Nhrpserverstattxerrsdusizeexceeded}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxErrInvalidExtension"] = types.YLeaf{"Nhrpserverstattxerrinvalidextension", nhrpserverstatentry.Nhrpserverstattxerrinvalidextension}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxErrAuthenticationFailure"] = types.YLeaf{"Nhrpserverstattxerrauthenticationfailure", nhrpserverstatentry.Nhrpserverstattxerrauthenticationfailure}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatTxErrHopCountExceeded"] = types.YLeaf{"Nhrpserverstattxerrhopcountexceeded", nhrpserverstatentry.Nhrpserverstattxerrhopcountexceeded}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatFwResolveReq"] = types.YLeaf{"Nhrpserverstatfwresolvereq", nhrpserverstatentry.Nhrpserverstatfwresolvereq}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatFwResolveReply"] = types.YLeaf{"Nhrpserverstatfwresolvereply", nhrpserverstatentry.Nhrpserverstatfwresolvereply}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatFwRegisterReq"] = types.YLeaf{"Nhrpserverstatfwregisterreq", nhrpserverstatentry.Nhrpserverstatfwregisterreq}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatFwRegisterReply"] = types.YLeaf{"Nhrpserverstatfwregisterreply", nhrpserverstatentry.Nhrpserverstatfwregisterreply}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatFwPurgeReq"] = types.YLeaf{"Nhrpserverstatfwpurgereq", nhrpserverstatentry.Nhrpserverstatfwpurgereq}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatFwPurgeReply"] = types.YLeaf{"Nhrpserverstatfwpurgereply", nhrpserverstatentry.Nhrpserverstatfwpurgereply}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatFwErrorIndication"] = types.YLeaf{"Nhrpserverstatfwerrorindication", nhrpserverstatentry.Nhrpserverstatfwerrorindication}
    nhrpserverstatentry.EntityData.Leafs["nhrpServerStatDiscontinuityTime"] = types.YLeaf{"Nhrpserverstatdiscontinuitytime", nhrpserverstatentry.Nhrpserverstatdiscontinuitytime}
    return &(nhrpserverstatentry.EntityData)
}

