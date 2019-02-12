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

    
    NhrpGeneralObjects NHRPMIB_NhrpGeneralObjects

    // This table contains mappings between internetwork layer addresses and NBMA
    // subnetwork layer addresses.
    NhrpCacheTable NHRPMIB_NhrpCacheTable

    // This table will track Purge Request Information.
    NhrpPurgeReqTable NHRPMIB_NhrpPurgeReqTable

    // Information about NHRP clients (NHCs) managed by this agent.
    NhrpClientTable NHRPMIB_NhrpClientTable

    // A table of Registration Request Information that needs to be maintained by
    // the NHCs (clients).
    NhrpClientRegistrationTable NHRPMIB_NhrpClientRegistrationTable

    // A table of NHSes that are available for use by this NHC (client). By
    // default, the agent will add an entry to this table that corresponds to the
    // client's default router.
    NhrpClientNhsTable NHRPMIB_NhrpClientNhsTable

    // This table contains statistics collected by NHRP clients.
    NhrpClientStatTable NHRPMIB_NhrpClientStatTable

    // This table contains information for a set of NHSes associated with this
    // agent.
    NhrpServerTable NHRPMIB_NhrpServerTable

    // This table extends the nhrpCacheTable for NHSes.  If the nhrpCacheTable has
    // a row added due to an NHS or based on information regarding an NHS then a
    // row is also added in this table.  The rows in this table will be created
    // when rows in the nhrpCacheTable are created.  However, there may be rows
    // created in the nhrpCacheTable which do not have corresponding rows in this
    // table.  For example, if the nhrpCacheTable has a row added due to a Next
    // Hop Client which is co-resident on the same device as the NHS, a row will
    // not be added to this table.
    NhrpServerCacheTable NHRPMIB_NhrpServerCacheTable

    // A table of NHCs that are available for use by this NHS (Server).
    NhrpServerNhcTable NHRPMIB_NhrpServerNhcTable

    // Statistics collected by Next Hop Servers.
    NhrpServerStatTable NHRPMIB_NhrpServerStatTable
}

func (nHRPMIB *NHRPMIB) GetEntityData() *types.CommonEntityData {
    nHRPMIB.EntityData.YFilter = nHRPMIB.YFilter
    nHRPMIB.EntityData.YangName = "NHRP-MIB"
    nHRPMIB.EntityData.BundleName = "cisco_ios_xe"
    nHRPMIB.EntityData.ParentYangName = "NHRP-MIB"
    nHRPMIB.EntityData.SegmentPath = "NHRP-MIB:NHRP-MIB"
    nHRPMIB.EntityData.AbsolutePath = nHRPMIB.EntityData.SegmentPath
    nHRPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nHRPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nHRPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nHRPMIB.EntityData.Children = types.NewOrderedMap()
    nHRPMIB.EntityData.Children.Append("nhrpGeneralObjects", types.YChild{"NhrpGeneralObjects", &nHRPMIB.NhrpGeneralObjects})
    nHRPMIB.EntityData.Children.Append("nhrpCacheTable", types.YChild{"NhrpCacheTable", &nHRPMIB.NhrpCacheTable})
    nHRPMIB.EntityData.Children.Append("nhrpPurgeReqTable", types.YChild{"NhrpPurgeReqTable", &nHRPMIB.NhrpPurgeReqTable})
    nHRPMIB.EntityData.Children.Append("nhrpClientTable", types.YChild{"NhrpClientTable", &nHRPMIB.NhrpClientTable})
    nHRPMIB.EntityData.Children.Append("nhrpClientRegistrationTable", types.YChild{"NhrpClientRegistrationTable", &nHRPMIB.NhrpClientRegistrationTable})
    nHRPMIB.EntityData.Children.Append("nhrpClientNhsTable", types.YChild{"NhrpClientNhsTable", &nHRPMIB.NhrpClientNhsTable})
    nHRPMIB.EntityData.Children.Append("nhrpClientStatTable", types.YChild{"NhrpClientStatTable", &nHRPMIB.NhrpClientStatTable})
    nHRPMIB.EntityData.Children.Append("nhrpServerTable", types.YChild{"NhrpServerTable", &nHRPMIB.NhrpServerTable})
    nHRPMIB.EntityData.Children.Append("nhrpServerCacheTable", types.YChild{"NhrpServerCacheTable", &nHRPMIB.NhrpServerCacheTable})
    nHRPMIB.EntityData.Children.Append("nhrpServerNhcTable", types.YChild{"NhrpServerNhcTable", &nHRPMIB.NhrpServerNhcTable})
    nHRPMIB.EntityData.Children.Append("nhrpServerStatTable", types.YChild{"NhrpServerStatTable", &nHRPMIB.NhrpServerStatTable})
    nHRPMIB.EntityData.Leafs = types.NewOrderedMap()

    nHRPMIB.EntityData.YListKeys = []string {}

    return &(nHRPMIB.EntityData)
}

// NHRPMIB_NhrpGeneralObjects
type NHRPMIB_NhrpGeneralObjects struct {
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
    NhrpNextIndex interface{}
}

func (nhrpGeneralObjects *NHRPMIB_NhrpGeneralObjects) GetEntityData() *types.CommonEntityData {
    nhrpGeneralObjects.EntityData.YFilter = nhrpGeneralObjects.YFilter
    nhrpGeneralObjects.EntityData.YangName = "nhrpGeneralObjects"
    nhrpGeneralObjects.EntityData.BundleName = "cisco_ios_xe"
    nhrpGeneralObjects.EntityData.ParentYangName = "NHRP-MIB"
    nhrpGeneralObjects.EntityData.SegmentPath = "nhrpGeneralObjects"
    nhrpGeneralObjects.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/" + nhrpGeneralObjects.EntityData.SegmentPath
    nhrpGeneralObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpGeneralObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpGeneralObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpGeneralObjects.EntityData.Children = types.NewOrderedMap()
    nhrpGeneralObjects.EntityData.Leafs = types.NewOrderedMap()
    nhrpGeneralObjects.EntityData.Leafs.Append("nhrpNextIndex", types.YLeaf{"NhrpNextIndex", nhrpGeneralObjects.NhrpNextIndex})

    nhrpGeneralObjects.EntityData.YListKeys = []string {}

    return &(nhrpGeneralObjects.EntityData)
}

// NHRPMIB_NhrpCacheTable
// This table contains mappings between internetwork layer
// addresses and NBMA subnetwork layer addresses.
type NHRPMIB_NhrpCacheTable struct {
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
    // The type is slice of NHRPMIB_NhrpCacheTable_NhrpCacheEntry.
    NhrpCacheEntry []*NHRPMIB_NhrpCacheTable_NhrpCacheEntry
}

func (nhrpCacheTable *NHRPMIB_NhrpCacheTable) GetEntityData() *types.CommonEntityData {
    nhrpCacheTable.EntityData.YFilter = nhrpCacheTable.YFilter
    nhrpCacheTable.EntityData.YangName = "nhrpCacheTable"
    nhrpCacheTable.EntityData.BundleName = "cisco_ios_xe"
    nhrpCacheTable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpCacheTable.EntityData.SegmentPath = "nhrpCacheTable"
    nhrpCacheTable.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/" + nhrpCacheTable.EntityData.SegmentPath
    nhrpCacheTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpCacheTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpCacheTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpCacheTable.EntityData.Children = types.NewOrderedMap()
    nhrpCacheTable.EntityData.Children.Append("nhrpCacheEntry", types.YChild{"NhrpCacheEntry", nil})
    for i := range nhrpCacheTable.NhrpCacheEntry {
        nhrpCacheTable.EntityData.Children.Append(types.GetSegmentPath(nhrpCacheTable.NhrpCacheEntry[i]), types.YChild{"NhrpCacheEntry", nhrpCacheTable.NhrpCacheEntry[i]})
    }
    nhrpCacheTable.EntityData.Leafs = types.NewOrderedMap()

    nhrpCacheTable.EntityData.YListKeys = []string {}

    return &(nhrpCacheTable.EntityData)
}

// NHRPMIB_NhrpCacheTable_NhrpCacheEntry
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
type NHRPMIB_NhrpCacheTable_NhrpCacheEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The internetwork layer address type of this Next
    // Hop Resolution Cache entry. The value of this object indicates how to
    // interpret the values of nhrpCacheInternetworkAddr and
    // nhrpCacheNextHopInternetworkAddr. The type is AddressFamilyNumbers.
    NhrpCacheInternetworkAddrType interface{}

    // This attribute is a key. The value of the internetwork address of the
    // destination. The type is string with length: 0..64.
    NhrpCacheInternetworkAddr interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. An identifier for this entry that has local
    // significance within the scope of the General Group.  This identifier is
    // used here to uniquely identify this row, and also used in the
    // 'nhrpPurgeTable' for the value of the 'nhrpPurgeCacheIdentifier'. The type
    // is interface{} with range: 1..4294967295.
    NhrpCacheIndex interface{}

    // The number of bits that define the internetwork layer prefix associated
    // with the nhrpCacheInternetworkAddr. The type is interface{} with range:
    // 0..255.
    NhrpCachePrefixLength interface{}

    // The value of the internetwork address of the next hop. The type is string
    // with length: 0..64.
    NhrpCacheNextHopInternetworkAddr interface{}

    // The NBMA address type. The value of this object indicates how to interpret
    // the values of nhrpCacheNbmaAddr and nhrpCacheNbmaSubaddr. The type is
    // AddressFamilyNumbers.
    NhrpCacheNbmaAddrType interface{}

    // The value of the NBMA subnetwork address of the next hop. The type is
    // string with length: 0..64.
    NhrpCacheNbmaAddr interface{}

    // The value of the NBMA subaddress of the next hop. If there is no subaddress
    // concept for the NBMA address family, this value will be a zero-length OCTET
    // STRING. The type is string with length: 0..64.
    NhrpCacheNbmaSubaddr interface{}

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
    // NhrpCacheType.
    NhrpCacheType interface{}

    // An indication of the state of this entry. The values are:  'incomplete(1)'
    // The client has sent a NHRP Resolution                 Request but has not
    // yet received the                 NHRP Resolution Reply.   'ackReply(2)'  
    // For a client or server, this is a                 cached valid mapping. 
    // 'nakReply(3)'   For a client or server, this is a                 cached
    // NAK mapping. The type is NhrpCacheState.
    NhrpCacheState interface{}

    // True(1) is returned if the value of 'nhrpCacheType' is not
    // 'administrativelyAdded'.  Since the value of 'nhrpCacheType' was not
    // configured by a user, the value of 'nhrpCacheHoldingTime' is considered
    // valid.  In other words, the value of 'nhrpCacheHoldingTime' represents the
    // Holding Time for the cache Entry.  If 'nhrpCacheType has been configured by
    // a user, (i.e. the value of 'nhrpCacheType' is 'administrativelyAdded') then
    // false(2) will be returned. This indicates that the value of
    // 'nhrpCacheHoldingTime' is undefined because this row could possibly be
    // backed up in nonvolatile storage. The type is bool.
    NhrpCacheHoldingTimeValid interface{}

    // If the value of 'nhrpCacheHoldingTimeValid is true(1) then this object
    // represents the number of seconds that the cache entry will remain in this
    // table.  When this value reaches 0 (zero) the row should be deleted.  If the
    // value of 'nhrpCacheHoldingTimeValid is false(2) then this object is
    // undefined. The type is interface{} with range: 0..65535. Units are seconds.
    NhrpCacheHoldingTime interface{}

    // The maximum transmission unit (MTU) that was negotiated or registered for
    // this entity. In other words, this is the actual MTU being used. The type is
    // interface{} with range: 0..65535.
    NhrpCacheNegotiatedMtu interface{}

    // An object which reflects the Preference value of the Client Information
    // Entry (CIE).  Zero or more Client Information Entries (CIEs) may be
    // included in the NHRP Packet.  One of the fields in the CIE is the
    // Preference.  For a complete description of the CIE, refer to Section
    // 5.2.0.1 of  RFC 2332 [17]. The type is interface{} with range: 0..255.
    NhrpCachePreference interface{}

    // This value only has meaning when the 'nhrpCacheType' has the value of
    // 'administrativelyAdded'.  When the row is created due to being
    // 'administrativelyAdded', this object reflects whether this row is kept in
    // volatile storage and lost upon reboot or if this row is backed up by
    // non-volatile or permanent storage.  If the value of 'nhrpCacheType' has a
    // value which is not 'administrativelyAdded, then the value of this object is
    // 'other(1)'. The type is StorageType.
    NhrpCacheStorageType interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    NhrpCacheRowStatus interface{}
}

func (nhrpCacheEntry *NHRPMIB_NhrpCacheTable_NhrpCacheEntry) GetEntityData() *types.CommonEntityData {
    nhrpCacheEntry.EntityData.YFilter = nhrpCacheEntry.YFilter
    nhrpCacheEntry.EntityData.YangName = "nhrpCacheEntry"
    nhrpCacheEntry.EntityData.BundleName = "cisco_ios_xe"
    nhrpCacheEntry.EntityData.ParentYangName = "nhrpCacheTable"
    nhrpCacheEntry.EntityData.SegmentPath = "nhrpCacheEntry" + types.AddKeyToken(nhrpCacheEntry.NhrpCacheInternetworkAddrType, "nhrpCacheInternetworkAddrType") + types.AddKeyToken(nhrpCacheEntry.NhrpCacheInternetworkAddr, "nhrpCacheInternetworkAddr") + types.AddKeyToken(nhrpCacheEntry.IfIndex, "ifIndex") + types.AddKeyToken(nhrpCacheEntry.NhrpCacheIndex, "nhrpCacheIndex")
    nhrpCacheEntry.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/nhrpCacheTable/" + nhrpCacheEntry.EntityData.SegmentPath
    nhrpCacheEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpCacheEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpCacheEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpCacheEntry.EntityData.Children = types.NewOrderedMap()
    nhrpCacheEntry.EntityData.Leafs = types.NewOrderedMap()
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheInternetworkAddrType", types.YLeaf{"NhrpCacheInternetworkAddrType", nhrpCacheEntry.NhrpCacheInternetworkAddrType})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheInternetworkAddr", types.YLeaf{"NhrpCacheInternetworkAddr", nhrpCacheEntry.NhrpCacheInternetworkAddr})
    nhrpCacheEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", nhrpCacheEntry.IfIndex})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheIndex", types.YLeaf{"NhrpCacheIndex", nhrpCacheEntry.NhrpCacheIndex})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCachePrefixLength", types.YLeaf{"NhrpCachePrefixLength", nhrpCacheEntry.NhrpCachePrefixLength})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheNextHopInternetworkAddr", types.YLeaf{"NhrpCacheNextHopInternetworkAddr", nhrpCacheEntry.NhrpCacheNextHopInternetworkAddr})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheNbmaAddrType", types.YLeaf{"NhrpCacheNbmaAddrType", nhrpCacheEntry.NhrpCacheNbmaAddrType})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheNbmaAddr", types.YLeaf{"NhrpCacheNbmaAddr", nhrpCacheEntry.NhrpCacheNbmaAddr})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheNbmaSubaddr", types.YLeaf{"NhrpCacheNbmaSubaddr", nhrpCacheEntry.NhrpCacheNbmaSubaddr})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheType", types.YLeaf{"NhrpCacheType", nhrpCacheEntry.NhrpCacheType})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheState", types.YLeaf{"NhrpCacheState", nhrpCacheEntry.NhrpCacheState})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheHoldingTimeValid", types.YLeaf{"NhrpCacheHoldingTimeValid", nhrpCacheEntry.NhrpCacheHoldingTimeValid})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheHoldingTime", types.YLeaf{"NhrpCacheHoldingTime", nhrpCacheEntry.NhrpCacheHoldingTime})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheNegotiatedMtu", types.YLeaf{"NhrpCacheNegotiatedMtu", nhrpCacheEntry.NhrpCacheNegotiatedMtu})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCachePreference", types.YLeaf{"NhrpCachePreference", nhrpCacheEntry.NhrpCachePreference})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheStorageType", types.YLeaf{"NhrpCacheStorageType", nhrpCacheEntry.NhrpCacheStorageType})
    nhrpCacheEntry.EntityData.Leafs.Append("nhrpCacheRowStatus", types.YLeaf{"NhrpCacheRowStatus", nhrpCacheEntry.NhrpCacheRowStatus})

    nhrpCacheEntry.EntityData.YListKeys = []string {"NhrpCacheInternetworkAddrType", "NhrpCacheInternetworkAddr", "IfIndex", "NhrpCacheIndex"}

    return &(nhrpCacheEntry.EntityData)
}

// NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheState represents                 cached NAK mapping.
type NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheState string

const (
    NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheState_incomplete NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheState = "incomplete"

    NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheState_ackReply NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheState = "ackReply"

    NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheState_nakReply NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheState = "nakReply"
)

// NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType represents The value cannot be modified once the entry is active.
type NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType string

const (
    NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType_other NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType = "other"

    NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType_register NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType = "register"

    NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType_resolveAuthoritative NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType = "resolveAuthoritative"

    NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType_resoveNonauthoritative NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType = "resoveNonauthoritative"

    NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType_transit NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType = "transit"

    NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType_administrativelyAdded NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType = "administrativelyAdded"

    NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType_atmarp NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType = "atmarp"

    NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType_scsp NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheType = "scsp"
)

// NHRPMIB_NhrpPurgeReqTable
// This table will track Purge Request Information.
type NHRPMIB_NhrpPurgeReqTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information regarding a Purge Request. The type is slice of
    // NHRPMIB_NhrpPurgeReqTable_NhrpPurgeReqEntry.
    NhrpPurgeReqEntry []*NHRPMIB_NhrpPurgeReqTable_NhrpPurgeReqEntry
}

func (nhrpPurgeReqTable *NHRPMIB_NhrpPurgeReqTable) GetEntityData() *types.CommonEntityData {
    nhrpPurgeReqTable.EntityData.YFilter = nhrpPurgeReqTable.YFilter
    nhrpPurgeReqTable.EntityData.YangName = "nhrpPurgeReqTable"
    nhrpPurgeReqTable.EntityData.BundleName = "cisco_ios_xe"
    nhrpPurgeReqTable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpPurgeReqTable.EntityData.SegmentPath = "nhrpPurgeReqTable"
    nhrpPurgeReqTable.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/" + nhrpPurgeReqTable.EntityData.SegmentPath
    nhrpPurgeReqTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpPurgeReqTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpPurgeReqTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpPurgeReqTable.EntityData.Children = types.NewOrderedMap()
    nhrpPurgeReqTable.EntityData.Children.Append("nhrpPurgeReqEntry", types.YChild{"NhrpPurgeReqEntry", nil})
    for i := range nhrpPurgeReqTable.NhrpPurgeReqEntry {
        nhrpPurgeReqTable.EntityData.Children.Append(types.GetSegmentPath(nhrpPurgeReqTable.NhrpPurgeReqEntry[i]), types.YChild{"NhrpPurgeReqEntry", nhrpPurgeReqTable.NhrpPurgeReqEntry[i]})
    }
    nhrpPurgeReqTable.EntityData.Leafs = types.NewOrderedMap()

    nhrpPurgeReqTable.EntityData.YListKeys = []string {}

    return &(nhrpPurgeReqTable.EntityData)
}

// NHRPMIB_NhrpPurgeReqTable_NhrpPurgeReqEntry
// Information regarding a Purge Request.
type NHRPMIB_NhrpPurgeReqTable_NhrpPurgeReqEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index for this entry that has local
    // significance within the scope of this table. The type is interface{} with
    // range: 1..4294967295.
    NhrpPurgeIndex interface{}

    // This object identifies which row in 'nhrpCacheTable' is being purged.  This
    // object should have the same value as the 'nhrpCacheIndex' in the
    // 'nhrpCacheTable'. The type is interface{} with range: 1..4294967295.
    NhrpPurgeCacheIdentifier interface{}

    // In the case of NHRP Purge Requests, this specifies the equivalence class of
    // addresses which match the first 'Prefix Length' bit positions of the Client
    // Protocol Address specified in the Client Information Entry (CIE). The type
    // is interface{} with range: 0..255.
    NhrpPurgePrefixLength interface{}

    // The Request ID used in the purge request. The type is interface{} with
    // range: 0..4294967295.
    NhrpPurgeRequestID interface{}

    // An indication of whether this Purge Request has the 'N' Bit cleared (off).
    // The type is bool.
    NhrpPurgeReplyExpected interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    NhrpPurgeRowStatus interface{}
}

func (nhrpPurgeReqEntry *NHRPMIB_NhrpPurgeReqTable_NhrpPurgeReqEntry) GetEntityData() *types.CommonEntityData {
    nhrpPurgeReqEntry.EntityData.YFilter = nhrpPurgeReqEntry.YFilter
    nhrpPurgeReqEntry.EntityData.YangName = "nhrpPurgeReqEntry"
    nhrpPurgeReqEntry.EntityData.BundleName = "cisco_ios_xe"
    nhrpPurgeReqEntry.EntityData.ParentYangName = "nhrpPurgeReqTable"
    nhrpPurgeReqEntry.EntityData.SegmentPath = "nhrpPurgeReqEntry" + types.AddKeyToken(nhrpPurgeReqEntry.NhrpPurgeIndex, "nhrpPurgeIndex")
    nhrpPurgeReqEntry.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/nhrpPurgeReqTable/" + nhrpPurgeReqEntry.EntityData.SegmentPath
    nhrpPurgeReqEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpPurgeReqEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpPurgeReqEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpPurgeReqEntry.EntityData.Children = types.NewOrderedMap()
    nhrpPurgeReqEntry.EntityData.Leafs = types.NewOrderedMap()
    nhrpPurgeReqEntry.EntityData.Leafs.Append("nhrpPurgeIndex", types.YLeaf{"NhrpPurgeIndex", nhrpPurgeReqEntry.NhrpPurgeIndex})
    nhrpPurgeReqEntry.EntityData.Leafs.Append("nhrpPurgeCacheIdentifier", types.YLeaf{"NhrpPurgeCacheIdentifier", nhrpPurgeReqEntry.NhrpPurgeCacheIdentifier})
    nhrpPurgeReqEntry.EntityData.Leafs.Append("nhrpPurgePrefixLength", types.YLeaf{"NhrpPurgePrefixLength", nhrpPurgeReqEntry.NhrpPurgePrefixLength})
    nhrpPurgeReqEntry.EntityData.Leafs.Append("nhrpPurgeRequestID", types.YLeaf{"NhrpPurgeRequestID", nhrpPurgeReqEntry.NhrpPurgeRequestID})
    nhrpPurgeReqEntry.EntityData.Leafs.Append("nhrpPurgeReplyExpected", types.YLeaf{"NhrpPurgeReplyExpected", nhrpPurgeReqEntry.NhrpPurgeReplyExpected})
    nhrpPurgeReqEntry.EntityData.Leafs.Append("nhrpPurgeRowStatus", types.YLeaf{"NhrpPurgeRowStatus", nhrpPurgeReqEntry.NhrpPurgeRowStatus})

    nhrpPurgeReqEntry.EntityData.YListKeys = []string {"NhrpPurgeIndex"}

    return &(nhrpPurgeReqEntry.EntityData)
}

// NHRPMIB_NhrpClientTable
// Information about NHRP clients (NHCs) managed by this
// agent.
type NHRPMIB_NhrpClientTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single NHC. The type is slice of
    // NHRPMIB_NhrpClientTable_NhrpClientEntry.
    NhrpClientEntry []*NHRPMIB_NhrpClientTable_NhrpClientEntry
}

func (nhrpClientTable *NHRPMIB_NhrpClientTable) GetEntityData() *types.CommonEntityData {
    nhrpClientTable.EntityData.YFilter = nhrpClientTable.YFilter
    nhrpClientTable.EntityData.YangName = "nhrpClientTable"
    nhrpClientTable.EntityData.BundleName = "cisco_ios_xe"
    nhrpClientTable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpClientTable.EntityData.SegmentPath = "nhrpClientTable"
    nhrpClientTable.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/" + nhrpClientTable.EntityData.SegmentPath
    nhrpClientTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpClientTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpClientTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpClientTable.EntityData.Children = types.NewOrderedMap()
    nhrpClientTable.EntityData.Children.Append("nhrpClientEntry", types.YChild{"NhrpClientEntry", nil})
    for i := range nhrpClientTable.NhrpClientEntry {
        nhrpClientTable.EntityData.Children.Append(types.GetSegmentPath(nhrpClientTable.NhrpClientEntry[i]), types.YChild{"NhrpClientEntry", nhrpClientTable.NhrpClientEntry[i]})
    }
    nhrpClientTable.EntityData.Leafs = types.NewOrderedMap()

    nhrpClientTable.EntityData.YListKeys = []string {}

    return &(nhrpClientTable.EntityData)
}

// NHRPMIB_NhrpClientTable_NhrpClientEntry
// Information about a single NHC.
type NHRPMIB_NhrpClientTable_NhrpClientEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An identifier for the NHRP client that is unique
    // within the scope of this agent.  The 'nhrpNextIndex' value should be
    // consulted (read), prior to creating a row in this table, and the value
    // returned from reading 'nhrpNextIndex' should be used as this object's
    // value. The type is interface{} with range: 1..4294967295.
    NhrpClientIndex interface{}

    // The type of the internetwork layer address of this client. This object
    // indicates how the value of nhrpClientInternetworkAddr is to be interpreted.
    // The type is AddressFamilyNumbers.
    NhrpClientInternetworkAddrType interface{}

    // The value of the internetwork layer address of this client. The type is
    // string with length: 0..64.
    NhrpClientInternetworkAddr interface{}

    // The type of the NBMA subnetwork address of this client. This object
    // indicates how the values of nhrpClientNbmaAddr and nhrpClientNbmaSubaddr
    // are to be interpreted. The type is AddressFamilyNumbers.
    NhrpClientNbmaAddrType interface{}

    // The NBMA subnetwork address of this client. The type is string with length:
    // 0..64.
    NhrpClientNbmaAddr interface{}

    // The NBMA subaddress of this client. For NBMA address families without a
    // subaddress concept, this will be a zero-length OCTET STRING. The type is
    // string with length: 0..64.
    NhrpClientNbmaSubaddr interface{}

    // The number of seconds that the client will wait before timing out an NHRP
    // initial request.  This object only has meaning for the initial timeout
    // period. The type is interface{} with range: 1..900. Units are seconds.
    NhrpClientInitialRequestTimeout interface{}

    // The number of times the client will retry the registration request before
    // failure. A value of 0 means don't retry. A value of 65535 means retry
    // forever. The type is interface{} with range: 0..65535.
    NhrpClientRegistrationRequestRetries interface{}

    // The number of times the client will retry the resolution request before
    // failure. A value of 0 means don't retry. A value of 65535 means retry
    // forever. The type is interface{} with range: 0..65535.
    NhrpClientResolutionRequestRetries interface{}

    // The number of times the client will retry a purge request before failure. A
    // value of 0 means don't retry. A value of 65535 means retry forever. The
    // type is interface{} with range: 0..65535.
    NhrpClientPurgeRequestRetries interface{}

    // The default maximum transmission unit (MTU) of the LIS/LAG which this
    // client should use. This object will be initialized by the agent to the
    // default MTU of the LIS/LAG (which is 9180) unless a different MTU value is
    // specified during creation of this Client. The type is interface{} with
    // range: 0..65535.
    NhrpClientDefaultMtu interface{}

    // The hold time the client will register. The type is interface{} with range:
    // 0..65535. Units are seconds.
    NhrpClientHoldTime interface{}

    // The Request ID used to register this client with its server. According to
    // Section 5.2.3 of the NHRP Specification, RFC 2332 [17], the Request ID must
    // be kept in non-volatile storage, so that if an NHC crashes and 
    // re-initializes, it will use a different  Request ID during the registration
    // process when reregistering with the same NHS. The type is interface{} with
    // range: 0..4294967295.
    NhrpClientRequestID interface{}

    // This object defines whether this row is kept in volatile storage and lost
    // upon a Client crash or reboot situation, or if this row is backed up by
    // nonvolatile or permanent storage. The type is StorageType.
    NhrpClientStorageType interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    NhrpClientRowStatus interface{}
}

func (nhrpClientEntry *NHRPMIB_NhrpClientTable_NhrpClientEntry) GetEntityData() *types.CommonEntityData {
    nhrpClientEntry.EntityData.YFilter = nhrpClientEntry.YFilter
    nhrpClientEntry.EntityData.YangName = "nhrpClientEntry"
    nhrpClientEntry.EntityData.BundleName = "cisco_ios_xe"
    nhrpClientEntry.EntityData.ParentYangName = "nhrpClientTable"
    nhrpClientEntry.EntityData.SegmentPath = "nhrpClientEntry" + types.AddKeyToken(nhrpClientEntry.NhrpClientIndex, "nhrpClientIndex")
    nhrpClientEntry.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/nhrpClientTable/" + nhrpClientEntry.EntityData.SegmentPath
    nhrpClientEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpClientEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpClientEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpClientEntry.EntityData.Children = types.NewOrderedMap()
    nhrpClientEntry.EntityData.Leafs = types.NewOrderedMap()
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientIndex", types.YLeaf{"NhrpClientIndex", nhrpClientEntry.NhrpClientIndex})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientInternetworkAddrType", types.YLeaf{"NhrpClientInternetworkAddrType", nhrpClientEntry.NhrpClientInternetworkAddrType})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientInternetworkAddr", types.YLeaf{"NhrpClientInternetworkAddr", nhrpClientEntry.NhrpClientInternetworkAddr})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientNbmaAddrType", types.YLeaf{"NhrpClientNbmaAddrType", nhrpClientEntry.NhrpClientNbmaAddrType})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientNbmaAddr", types.YLeaf{"NhrpClientNbmaAddr", nhrpClientEntry.NhrpClientNbmaAddr})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientNbmaSubaddr", types.YLeaf{"NhrpClientNbmaSubaddr", nhrpClientEntry.NhrpClientNbmaSubaddr})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientInitialRequestTimeout", types.YLeaf{"NhrpClientInitialRequestTimeout", nhrpClientEntry.NhrpClientInitialRequestTimeout})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientRegistrationRequestRetries", types.YLeaf{"NhrpClientRegistrationRequestRetries", nhrpClientEntry.NhrpClientRegistrationRequestRetries})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientResolutionRequestRetries", types.YLeaf{"NhrpClientResolutionRequestRetries", nhrpClientEntry.NhrpClientResolutionRequestRetries})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientPurgeRequestRetries", types.YLeaf{"NhrpClientPurgeRequestRetries", nhrpClientEntry.NhrpClientPurgeRequestRetries})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientDefaultMtu", types.YLeaf{"NhrpClientDefaultMtu", nhrpClientEntry.NhrpClientDefaultMtu})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientHoldTime", types.YLeaf{"NhrpClientHoldTime", nhrpClientEntry.NhrpClientHoldTime})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientRequestID", types.YLeaf{"NhrpClientRequestID", nhrpClientEntry.NhrpClientRequestID})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientStorageType", types.YLeaf{"NhrpClientStorageType", nhrpClientEntry.NhrpClientStorageType})
    nhrpClientEntry.EntityData.Leafs.Append("nhrpClientRowStatus", types.YLeaf{"NhrpClientRowStatus", nhrpClientEntry.NhrpClientRowStatus})

    nhrpClientEntry.EntityData.YListKeys = []string {"NhrpClientIndex"}

    return &(nhrpClientEntry.EntityData)
}

// NHRPMIB_NhrpClientRegistrationTable
// A table of Registration Request Information that
// needs to be maintained by the NHCs (clients).
type NHRPMIB_NhrpClientRegistrationTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An NHC needs to maintain registration request information between the NHC
    // and the NHS.  An entry in this table represents information for a single
    // registration request. The type is slice of
    // NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry.
    NhrpClientRegistrationEntry []*NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry
}

func (nhrpClientRegistrationTable *NHRPMIB_NhrpClientRegistrationTable) GetEntityData() *types.CommonEntityData {
    nhrpClientRegistrationTable.EntityData.YFilter = nhrpClientRegistrationTable.YFilter
    nhrpClientRegistrationTable.EntityData.YangName = "nhrpClientRegistrationTable"
    nhrpClientRegistrationTable.EntityData.BundleName = "cisco_ios_xe"
    nhrpClientRegistrationTable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpClientRegistrationTable.EntityData.SegmentPath = "nhrpClientRegistrationTable"
    nhrpClientRegistrationTable.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/" + nhrpClientRegistrationTable.EntityData.SegmentPath
    nhrpClientRegistrationTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpClientRegistrationTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpClientRegistrationTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpClientRegistrationTable.EntityData.Children = types.NewOrderedMap()
    nhrpClientRegistrationTable.EntityData.Children.Append("nhrpClientRegistrationEntry", types.YChild{"NhrpClientRegistrationEntry", nil})
    for i := range nhrpClientRegistrationTable.NhrpClientRegistrationEntry {
        nhrpClientRegistrationTable.EntityData.Children.Append(types.GetSegmentPath(nhrpClientRegistrationTable.NhrpClientRegistrationEntry[i]), types.YChild{"NhrpClientRegistrationEntry", nhrpClientRegistrationTable.NhrpClientRegistrationEntry[i]})
    }
    nhrpClientRegistrationTable.EntityData.Leafs = types.NewOrderedMap()

    nhrpClientRegistrationTable.EntityData.YListKeys = []string {}

    return &(nhrpClientRegistrationTable.EntityData)
}

// NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry
// An NHC needs to maintain registration request information
// between the NHC and the NHS.  An entry in this table
// represents information for a single registration request.
type NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to nhrp_mib.NHRPMIB_NhrpClientTable_NhrpClientEntry_NhrpClientIndex
    NhrpClientIndex interface{}

    // This attribute is a key. An identifier for this entry such that it
    // identifies a specific Registration Request from the NHC represented by the
    // nhrpClientIndex. The type is interface{} with range: 1..4294967295.
    NhrpClientRegIndex interface{}

    // The Uniqueness indicator for this Registration Request. If this object has
    // the value of requestUnique(1), then the Uniqueness bit is set in the the
    // NHRP Registration Request represented by this row.  The value cannot be
    // changed once the row is created. The type is NhrpClientRegUniqueness.
    NhrpClientRegUniqueness interface{}

    // The registration state of this client. The values are: 'other(1)'          
    // The state of the registration                        request is not one of 
    // 'registering',                        'ackRegisterReply' or                
    // 'nakRegisterReply'.  'registering(2)'        A registration request has    
    // been issued and a registration                         reply is expected. 
    // 'ackRegisterReply(3)'   A positive registration reply                      
    // has been received.  'nakRegisterReply(4)'   The client has received a      
    // negative registration                         reply (NAK). The type is
    // NhrpClientRegState.
    NhrpClientRegState interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    NhrpClientRegRowStatus interface{}
}

func (nhrpClientRegistrationEntry *NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry) GetEntityData() *types.CommonEntityData {
    nhrpClientRegistrationEntry.EntityData.YFilter = nhrpClientRegistrationEntry.YFilter
    nhrpClientRegistrationEntry.EntityData.YangName = "nhrpClientRegistrationEntry"
    nhrpClientRegistrationEntry.EntityData.BundleName = "cisco_ios_xe"
    nhrpClientRegistrationEntry.EntityData.ParentYangName = "nhrpClientRegistrationTable"
    nhrpClientRegistrationEntry.EntityData.SegmentPath = "nhrpClientRegistrationEntry" + types.AddKeyToken(nhrpClientRegistrationEntry.NhrpClientIndex, "nhrpClientIndex") + types.AddKeyToken(nhrpClientRegistrationEntry.NhrpClientRegIndex, "nhrpClientRegIndex")
    nhrpClientRegistrationEntry.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/nhrpClientRegistrationTable/" + nhrpClientRegistrationEntry.EntityData.SegmentPath
    nhrpClientRegistrationEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpClientRegistrationEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpClientRegistrationEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpClientRegistrationEntry.EntityData.Children = types.NewOrderedMap()
    nhrpClientRegistrationEntry.EntityData.Leafs = types.NewOrderedMap()
    nhrpClientRegistrationEntry.EntityData.Leafs.Append("nhrpClientIndex", types.YLeaf{"NhrpClientIndex", nhrpClientRegistrationEntry.NhrpClientIndex})
    nhrpClientRegistrationEntry.EntityData.Leafs.Append("nhrpClientRegIndex", types.YLeaf{"NhrpClientRegIndex", nhrpClientRegistrationEntry.NhrpClientRegIndex})
    nhrpClientRegistrationEntry.EntityData.Leafs.Append("nhrpClientRegUniqueness", types.YLeaf{"NhrpClientRegUniqueness", nhrpClientRegistrationEntry.NhrpClientRegUniqueness})
    nhrpClientRegistrationEntry.EntityData.Leafs.Append("nhrpClientRegState", types.YLeaf{"NhrpClientRegState", nhrpClientRegistrationEntry.NhrpClientRegState})
    nhrpClientRegistrationEntry.EntityData.Leafs.Append("nhrpClientRegRowStatus", types.YLeaf{"NhrpClientRegRowStatus", nhrpClientRegistrationEntry.NhrpClientRegRowStatus})

    nhrpClientRegistrationEntry.EntityData.YListKeys = []string {"NhrpClientIndex", "NhrpClientRegIndex"}

    return &(nhrpClientRegistrationEntry.EntityData)
}

// NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegState represents                         reply (NAK).
type NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegState string

const (
    NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegState_other NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegState = "other"

    NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegState_registering NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegState = "registering"

    NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegState_ackRegisterReply NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegState = "ackRegisterReply"

    NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegState_nakRegisterReply NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegState = "nakRegisterReply"
)

// NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegUniqueness represents be changed once the row is created.
type NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegUniqueness string

const (
    NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegUniqueness_requestUnique NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegUniqueness = "requestUnique"

    NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegUniqueness_requestNotUnique NHRPMIB_NhrpClientRegistrationTable_NhrpClientRegistrationEntry_NhrpClientRegUniqueness = "requestNotUnique"
)

// NHRPMIB_NhrpClientNhsTable
// A table of NHSes that are available for use by this NHC
// (client). By default, the agent will add an entry to this
// table that corresponds to the client's default router.
type NHRPMIB_NhrpClientNhsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An NHS that may be used by an NHC. The type is slice of
    // NHRPMIB_NhrpClientNhsTable_NhrpClientNhsEntry.
    NhrpClientNhsEntry []*NHRPMIB_NhrpClientNhsTable_NhrpClientNhsEntry
}

func (nhrpClientNhsTable *NHRPMIB_NhrpClientNhsTable) GetEntityData() *types.CommonEntityData {
    nhrpClientNhsTable.EntityData.YFilter = nhrpClientNhsTable.YFilter
    nhrpClientNhsTable.EntityData.YangName = "nhrpClientNhsTable"
    nhrpClientNhsTable.EntityData.BundleName = "cisco_ios_xe"
    nhrpClientNhsTable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpClientNhsTable.EntityData.SegmentPath = "nhrpClientNhsTable"
    nhrpClientNhsTable.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/" + nhrpClientNhsTable.EntityData.SegmentPath
    nhrpClientNhsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpClientNhsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpClientNhsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpClientNhsTable.EntityData.Children = types.NewOrderedMap()
    nhrpClientNhsTable.EntityData.Children.Append("nhrpClientNhsEntry", types.YChild{"NhrpClientNhsEntry", nil})
    for i := range nhrpClientNhsTable.NhrpClientNhsEntry {
        nhrpClientNhsTable.EntityData.Children.Append(types.GetSegmentPath(nhrpClientNhsTable.NhrpClientNhsEntry[i]), types.YChild{"NhrpClientNhsEntry", nhrpClientNhsTable.NhrpClientNhsEntry[i]})
    }
    nhrpClientNhsTable.EntityData.Leafs = types.NewOrderedMap()

    nhrpClientNhsTable.EntityData.YListKeys = []string {}

    return &(nhrpClientNhsTable.EntityData)
}

// NHRPMIB_NhrpClientNhsTable_NhrpClientNhsEntry
// An NHS that may be used by an NHC.
type NHRPMIB_NhrpClientNhsTable_NhrpClientNhsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to nhrp_mib.NHRPMIB_NhrpClientTable_NhrpClientEntry_NhrpClientIndex
    NhrpClientIndex interface{}

    // This attribute is a key. An identifier for an NHS available to an NHC. The
    // type is interface{} with range: 1..4294967295.
    NhrpClientNhsIndex interface{}

    // The type of the internetwork layer address of the NHRP server represented
    // in this entry. This object indicates how the value of
    // nhrpClientNhsInternetworkAddr is to be interpreted. The type is
    // AddressFamilyNumbers.
    NhrpClientNhsInternetworkAddrType interface{}

    // The value of the destination internetwork layer address of the NHRP server
    // represented by this    entry.  If this value is not known, this will be a
    // zero-length OCTET STRING. The type is string with length: 0..64.
    NhrpClientNhsInternetworkAddr interface{}

    // The type of the NBMA subnetwork address of the NHRP Server represented by
    // this entry. This object indicates how the values of nhrpClientNhsNbmaAddr
    // and nhrpClientNhsNbmaSubaddr are to be interpreted. The type is
    // AddressFamilyNumbers.
    NhrpClientNhsNbmaAddrType interface{}

    // The NBMA subnetwork address of the NHS. The type of the address is
    // indicated by the corresponding value of nhrpClientNhsNbmaAddrType. The type
    // is string with length: 0..64.
    NhrpClientNhsNbmaAddr interface{}

    // The NBMA subaddress of the NHS. For NMBA address families that do not have
    // the concept of subaddress,      this will be a zero-length OCTET STRING.
    // The type is string with length: 0..64.
    NhrpClientNhsNbmaSubaddr interface{}

    // An indication of whether this NHS is in use by the NHC. The type is bool.
    NhrpClientNhsInUse interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    NhrpClientNhsRowStatus interface{}
}

func (nhrpClientNhsEntry *NHRPMIB_NhrpClientNhsTable_NhrpClientNhsEntry) GetEntityData() *types.CommonEntityData {
    nhrpClientNhsEntry.EntityData.YFilter = nhrpClientNhsEntry.YFilter
    nhrpClientNhsEntry.EntityData.YangName = "nhrpClientNhsEntry"
    nhrpClientNhsEntry.EntityData.BundleName = "cisco_ios_xe"
    nhrpClientNhsEntry.EntityData.ParentYangName = "nhrpClientNhsTable"
    nhrpClientNhsEntry.EntityData.SegmentPath = "nhrpClientNhsEntry" + types.AddKeyToken(nhrpClientNhsEntry.NhrpClientIndex, "nhrpClientIndex") + types.AddKeyToken(nhrpClientNhsEntry.NhrpClientNhsIndex, "nhrpClientNhsIndex")
    nhrpClientNhsEntry.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/nhrpClientNhsTable/" + nhrpClientNhsEntry.EntityData.SegmentPath
    nhrpClientNhsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpClientNhsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpClientNhsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpClientNhsEntry.EntityData.Children = types.NewOrderedMap()
    nhrpClientNhsEntry.EntityData.Leafs = types.NewOrderedMap()
    nhrpClientNhsEntry.EntityData.Leafs.Append("nhrpClientIndex", types.YLeaf{"NhrpClientIndex", nhrpClientNhsEntry.NhrpClientIndex})
    nhrpClientNhsEntry.EntityData.Leafs.Append("nhrpClientNhsIndex", types.YLeaf{"NhrpClientNhsIndex", nhrpClientNhsEntry.NhrpClientNhsIndex})
    nhrpClientNhsEntry.EntityData.Leafs.Append("nhrpClientNhsInternetworkAddrType", types.YLeaf{"NhrpClientNhsInternetworkAddrType", nhrpClientNhsEntry.NhrpClientNhsInternetworkAddrType})
    nhrpClientNhsEntry.EntityData.Leafs.Append("nhrpClientNhsInternetworkAddr", types.YLeaf{"NhrpClientNhsInternetworkAddr", nhrpClientNhsEntry.NhrpClientNhsInternetworkAddr})
    nhrpClientNhsEntry.EntityData.Leafs.Append("nhrpClientNhsNbmaAddrType", types.YLeaf{"NhrpClientNhsNbmaAddrType", nhrpClientNhsEntry.NhrpClientNhsNbmaAddrType})
    nhrpClientNhsEntry.EntityData.Leafs.Append("nhrpClientNhsNbmaAddr", types.YLeaf{"NhrpClientNhsNbmaAddr", nhrpClientNhsEntry.NhrpClientNhsNbmaAddr})
    nhrpClientNhsEntry.EntityData.Leafs.Append("nhrpClientNhsNbmaSubaddr", types.YLeaf{"NhrpClientNhsNbmaSubaddr", nhrpClientNhsEntry.NhrpClientNhsNbmaSubaddr})
    nhrpClientNhsEntry.EntityData.Leafs.Append("nhrpClientNhsInUse", types.YLeaf{"NhrpClientNhsInUse", nhrpClientNhsEntry.NhrpClientNhsInUse})
    nhrpClientNhsEntry.EntityData.Leafs.Append("nhrpClientNhsRowStatus", types.YLeaf{"NhrpClientNhsRowStatus", nhrpClientNhsEntry.NhrpClientNhsRowStatus})

    nhrpClientNhsEntry.EntityData.YListKeys = []string {"NhrpClientIndex", "NhrpClientNhsIndex"}

    return &(nhrpClientNhsEntry.EntityData)
}

// NHRPMIB_NhrpClientStatTable
// This table contains statistics collected by NHRP
// clients.
type NHRPMIB_NhrpClientStatTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics collected by a NHRP client. The type is slice of
    // NHRPMIB_NhrpClientStatTable_NhrpClientStatEntry.
    NhrpClientStatEntry []*NHRPMIB_NhrpClientStatTable_NhrpClientStatEntry
}

func (nhrpClientStatTable *NHRPMIB_NhrpClientStatTable) GetEntityData() *types.CommonEntityData {
    nhrpClientStatTable.EntityData.YFilter = nhrpClientStatTable.YFilter
    nhrpClientStatTable.EntityData.YangName = "nhrpClientStatTable"
    nhrpClientStatTable.EntityData.BundleName = "cisco_ios_xe"
    nhrpClientStatTable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpClientStatTable.EntityData.SegmentPath = "nhrpClientStatTable"
    nhrpClientStatTable.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/" + nhrpClientStatTable.EntityData.SegmentPath
    nhrpClientStatTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpClientStatTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpClientStatTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpClientStatTable.EntityData.Children = types.NewOrderedMap()
    nhrpClientStatTable.EntityData.Children.Append("nhrpClientStatEntry", types.YChild{"NhrpClientStatEntry", nil})
    for i := range nhrpClientStatTable.NhrpClientStatEntry {
        nhrpClientStatTable.EntityData.Children.Append(types.GetSegmentPath(nhrpClientStatTable.NhrpClientStatEntry[i]), types.YChild{"NhrpClientStatEntry", nhrpClientStatTable.NhrpClientStatEntry[i]})
    }
    nhrpClientStatTable.EntityData.Leafs = types.NewOrderedMap()

    nhrpClientStatTable.EntityData.YListKeys = []string {}

    return &(nhrpClientStatTable.EntityData)
}

// NHRPMIB_NhrpClientStatTable_NhrpClientStatEntry
// Statistics collected by a NHRP client.
type NHRPMIB_NhrpClientStatTable_NhrpClientStatEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to nhrp_mib.NHRPMIB_NhrpClientTable_NhrpClientEntry_NhrpClientIndex
    NhrpClientIndex interface{}

    // The number of NHRP Resolution Requests transmitted by this client. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpClientStatTxResolveReq interface{}

    // The number of positively acknowledged NHRP Resolution Replies received by
    // this client.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Client
    // re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpClientStatRxResolveReplyAck interface{}

    // The number of NAKed NHRP Resolution Replies received by this client that
    // contained the code indicating 'Administratively Prohibited'.  
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpClientStatRxResolveReplyNakProhibited interface{}

    // The number of NAKed NHRP Resolution Replies received by this client that
    // contained the code indicating 'Insufficient Resources'.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, at NHRP Client re-initialization and at other times as indicated by
    // the value of nhrpClientStatDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    NhrpClientStatRxResolveReplyNakInsufResources interface{}

    // The number of NAKed NHRP Resolution Replies received by this client that
    // contained the code indicating 'No Internetworking Layer Address to NBMA
    // Address Binding Exists'.  Discontinuities in the value of this counter can
    // occur at re-initialization of the management system, at NHRP Client
    // re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpClientStatRxResolveReplyNakNoBinding interface{}

    // The number of NAKed NHRP Resolution Replies received by this client that
    // contained the code indicating 'Binding Exists But Is Not Unique'. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpClientStatRxResolveReplyNakNotUnique interface{}

    // The number of NHRP Registration Requests transmitted by this client. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpClientStatTxRegisterReq interface{}

    // The number of positively acknowledged NHRP Registration Replies received by
    // this client.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Client
    // re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpClientStatRxRegisterAck interface{}

    // The number of NAKed NHRP Registration Replies received by this client that
    // contained the code indicating 'Administratively Prohibited'. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpClientStatRxRegisterNakProhibited interface{}

    // The number of NAKed NHRP Registration Replies received by this client that
    // contained the code indicating 'Insufficient Resources'.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, at NHRP Client re-initialization and at other times as indicated by
    // the value of nhrpClientStatDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    NhrpClientStatRxRegisterNakInsufResources interface{}

    // The number of NAKed NHRP Registration Replies received by this client that
    // contained the code indicating 'Unique Internetworking Layer Address Already
    // Registered'.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Client
    // re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpClientStatRxRegisterNakAlreadyReg interface{}

    // The number of NHRP Purge Requests received by this client.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, at NHRP Client re-initialization and at other times as
    // indicated by the value of nhrpClientStatDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    NhrpClientStatRxPurgeReq interface{}

    // The number of NHRP Purge Requests transmitted by this client. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpClientStatTxPurgeReq interface{}

    // The number of NHRP Purge Replies received by this client.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, at NHRP Client re-initialization and at other times as
    // indicated by the value of nhrpClientStatDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    NhrpClientStatRxPurgeReply interface{}

    // The number of NHRP Purge Replies transmitted by this client. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpClientStatTxPurgeReply interface{}

    // The number of NHRP Error Indication packets transmitted by this client. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Client re-initialization and at other
    // times as indicated by the value of nhrpClientStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpClientStatTxErrorIndication interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'Unrecognized Extension'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Client re-initialization and at other times as indicated by the value
    // of nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpClientStatRxErrUnrecognizedExtension interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'NHRP Loop Detected'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Client re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpClientStatRxErrLoopDetected interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'Protocol Address Unreachable'.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, at NHRP Client re-initialization and at other times as indicated by
    // the value of nhrpClientStatDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    NhrpClientStatRxErrProtoAddrUnreachable interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'Protocol Error'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Client re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpClientStatRxErrProtoError interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'NHRP SDU Size  Exceeded'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Client re-initialization and at other times as indicated by the value
    // of nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpClientStatRxErrSduSizeExceeded interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'Invalid Extension'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Client re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpClientStatRxErrInvalidExtension interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'Authentication Failure'.      Discontinuities in the value
    // of this counter can occur at re-initialization of the management system, at
    // NHRP Client re-initialization and at other times as indicated by the value
    // of nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpClientStatRxErrAuthenticationFailure interface{}

    // The number of NHRP Error Indication packets received by this client with
    // the error code 'Hop Count Exceeded'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Client re-initialization and at other times as indicated by the value of
    // nhrpClientStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpClientStatRxErrHopCountExceeded interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this Client's counters suffered a discontinuity.  If no such
    // discontinuities have occurred since the last re-initialization of the local
    // management subsystem or the NHRP Client re-initialization associated with
    // this entry, then this object contains a zero value. The type is interface{}
    // with range: 0..4294967295.
    NhrpClientStatDiscontinuityTime interface{}
}

func (nhrpClientStatEntry *NHRPMIB_NhrpClientStatTable_NhrpClientStatEntry) GetEntityData() *types.CommonEntityData {
    nhrpClientStatEntry.EntityData.YFilter = nhrpClientStatEntry.YFilter
    nhrpClientStatEntry.EntityData.YangName = "nhrpClientStatEntry"
    nhrpClientStatEntry.EntityData.BundleName = "cisco_ios_xe"
    nhrpClientStatEntry.EntityData.ParentYangName = "nhrpClientStatTable"
    nhrpClientStatEntry.EntityData.SegmentPath = "nhrpClientStatEntry" + types.AddKeyToken(nhrpClientStatEntry.NhrpClientIndex, "nhrpClientIndex")
    nhrpClientStatEntry.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/nhrpClientStatTable/" + nhrpClientStatEntry.EntityData.SegmentPath
    nhrpClientStatEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpClientStatEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpClientStatEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpClientStatEntry.EntityData.Children = types.NewOrderedMap()
    nhrpClientStatEntry.EntityData.Leafs = types.NewOrderedMap()
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientIndex", types.YLeaf{"NhrpClientIndex", nhrpClientStatEntry.NhrpClientIndex})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatTxResolveReq", types.YLeaf{"NhrpClientStatTxResolveReq", nhrpClientStatEntry.NhrpClientStatTxResolveReq})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxResolveReplyAck", types.YLeaf{"NhrpClientStatRxResolveReplyAck", nhrpClientStatEntry.NhrpClientStatRxResolveReplyAck})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxResolveReplyNakProhibited", types.YLeaf{"NhrpClientStatRxResolveReplyNakProhibited", nhrpClientStatEntry.NhrpClientStatRxResolveReplyNakProhibited})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxResolveReplyNakInsufResources", types.YLeaf{"NhrpClientStatRxResolveReplyNakInsufResources", nhrpClientStatEntry.NhrpClientStatRxResolveReplyNakInsufResources})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxResolveReplyNakNoBinding", types.YLeaf{"NhrpClientStatRxResolveReplyNakNoBinding", nhrpClientStatEntry.NhrpClientStatRxResolveReplyNakNoBinding})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxResolveReplyNakNotUnique", types.YLeaf{"NhrpClientStatRxResolveReplyNakNotUnique", nhrpClientStatEntry.NhrpClientStatRxResolveReplyNakNotUnique})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatTxRegisterReq", types.YLeaf{"NhrpClientStatTxRegisterReq", nhrpClientStatEntry.NhrpClientStatTxRegisterReq})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxRegisterAck", types.YLeaf{"NhrpClientStatRxRegisterAck", nhrpClientStatEntry.NhrpClientStatRxRegisterAck})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxRegisterNakProhibited", types.YLeaf{"NhrpClientStatRxRegisterNakProhibited", nhrpClientStatEntry.NhrpClientStatRxRegisterNakProhibited})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxRegisterNakInsufResources", types.YLeaf{"NhrpClientStatRxRegisterNakInsufResources", nhrpClientStatEntry.NhrpClientStatRxRegisterNakInsufResources})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxRegisterNakAlreadyReg", types.YLeaf{"NhrpClientStatRxRegisterNakAlreadyReg", nhrpClientStatEntry.NhrpClientStatRxRegisterNakAlreadyReg})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxPurgeReq", types.YLeaf{"NhrpClientStatRxPurgeReq", nhrpClientStatEntry.NhrpClientStatRxPurgeReq})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatTxPurgeReq", types.YLeaf{"NhrpClientStatTxPurgeReq", nhrpClientStatEntry.NhrpClientStatTxPurgeReq})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxPurgeReply", types.YLeaf{"NhrpClientStatRxPurgeReply", nhrpClientStatEntry.NhrpClientStatRxPurgeReply})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatTxPurgeReply", types.YLeaf{"NhrpClientStatTxPurgeReply", nhrpClientStatEntry.NhrpClientStatTxPurgeReply})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatTxErrorIndication", types.YLeaf{"NhrpClientStatTxErrorIndication", nhrpClientStatEntry.NhrpClientStatTxErrorIndication})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxErrUnrecognizedExtension", types.YLeaf{"NhrpClientStatRxErrUnrecognizedExtension", nhrpClientStatEntry.NhrpClientStatRxErrUnrecognizedExtension})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxErrLoopDetected", types.YLeaf{"NhrpClientStatRxErrLoopDetected", nhrpClientStatEntry.NhrpClientStatRxErrLoopDetected})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxErrProtoAddrUnreachable", types.YLeaf{"NhrpClientStatRxErrProtoAddrUnreachable", nhrpClientStatEntry.NhrpClientStatRxErrProtoAddrUnreachable})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxErrProtoError", types.YLeaf{"NhrpClientStatRxErrProtoError", nhrpClientStatEntry.NhrpClientStatRxErrProtoError})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxErrSduSizeExceeded", types.YLeaf{"NhrpClientStatRxErrSduSizeExceeded", nhrpClientStatEntry.NhrpClientStatRxErrSduSizeExceeded})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxErrInvalidExtension", types.YLeaf{"NhrpClientStatRxErrInvalidExtension", nhrpClientStatEntry.NhrpClientStatRxErrInvalidExtension})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxErrAuthenticationFailure", types.YLeaf{"NhrpClientStatRxErrAuthenticationFailure", nhrpClientStatEntry.NhrpClientStatRxErrAuthenticationFailure})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatRxErrHopCountExceeded", types.YLeaf{"NhrpClientStatRxErrHopCountExceeded", nhrpClientStatEntry.NhrpClientStatRxErrHopCountExceeded})
    nhrpClientStatEntry.EntityData.Leafs.Append("nhrpClientStatDiscontinuityTime", types.YLeaf{"NhrpClientStatDiscontinuityTime", nhrpClientStatEntry.NhrpClientStatDiscontinuityTime})

    nhrpClientStatEntry.EntityData.YListKeys = []string {"NhrpClientIndex"}

    return &(nhrpClientStatEntry.EntityData)
}

// NHRPMIB_NhrpServerTable
// This table contains information for a set of NHSes
// associated with this agent.
type NHRPMIB_NhrpServerTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single NHS. The type is slice of
    // NHRPMIB_NhrpServerTable_NhrpServerEntry.
    NhrpServerEntry []*NHRPMIB_NhrpServerTable_NhrpServerEntry
}

func (nhrpServerTable *NHRPMIB_NhrpServerTable) GetEntityData() *types.CommonEntityData {
    nhrpServerTable.EntityData.YFilter = nhrpServerTable.YFilter
    nhrpServerTable.EntityData.YangName = "nhrpServerTable"
    nhrpServerTable.EntityData.BundleName = "cisco_ios_xe"
    nhrpServerTable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpServerTable.EntityData.SegmentPath = "nhrpServerTable"
    nhrpServerTable.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/" + nhrpServerTable.EntityData.SegmentPath
    nhrpServerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpServerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpServerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpServerTable.EntityData.Children = types.NewOrderedMap()
    nhrpServerTable.EntityData.Children.Append("nhrpServerEntry", types.YChild{"NhrpServerEntry", nil})
    for i := range nhrpServerTable.NhrpServerEntry {
        nhrpServerTable.EntityData.Children.Append(types.GetSegmentPath(nhrpServerTable.NhrpServerEntry[i]), types.YChild{"NhrpServerEntry", nhrpServerTable.NhrpServerEntry[i]})
    }
    nhrpServerTable.EntityData.Leafs = types.NewOrderedMap()

    nhrpServerTable.EntityData.YListKeys = []string {}

    return &(nhrpServerTable.EntityData)
}

// NHRPMIB_NhrpServerTable_NhrpServerEntry
// Information about a single NHS.
type NHRPMIB_NhrpServerTable_NhrpServerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An identifier for the server that is unique within
    // the scope of this agent. The type is interface{} with range: 1..4294967295.
    NhrpServerIndex interface{}

    // The type of the internetwork layer address of this server. This object is
    // used to interpret the value of nhrpServerInternetworkAddr. The type is
    // AddressFamilyNumbers.
    NhrpServerInternetworkAddrType interface{}

    // The value of the internetwork layer address of this server. The type is
    // string with length: 0..64.
    NhrpServerInternetworkAddr interface{}

    // The type of the NBMA subnetwork address of this server. This object is used
    // to interpret the value of nhrpServerNbmaAddr. The type is
    // AddressFamilyNumbers.
    NhrpServerNbmaAddrType interface{}

    // The value of the NBMA subnetwork address of this server. The type is string
    // with length: 0..64.
    NhrpServerNbmaAddr interface{}

    // The value of the NBMA subaddress of this server. For NBMA address families
    // without a subaddress concept, this will be a zero-length OCTET STRING. The
    // type is string with length: 0..64.
    NhrpServerNbmaSubaddr interface{}

    // This object defines whether this row is kept in volatile storage and lost
    // upon a Server crash or reboot situation, or if this row is backed up by
    // nonvolatile or permanent storage. The type is StorageType.
    NhrpServerStorageType interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    NhrpServerRowStatus interface{}
}

func (nhrpServerEntry *NHRPMIB_NhrpServerTable_NhrpServerEntry) GetEntityData() *types.CommonEntityData {
    nhrpServerEntry.EntityData.YFilter = nhrpServerEntry.YFilter
    nhrpServerEntry.EntityData.YangName = "nhrpServerEntry"
    nhrpServerEntry.EntityData.BundleName = "cisco_ios_xe"
    nhrpServerEntry.EntityData.ParentYangName = "nhrpServerTable"
    nhrpServerEntry.EntityData.SegmentPath = "nhrpServerEntry" + types.AddKeyToken(nhrpServerEntry.NhrpServerIndex, "nhrpServerIndex")
    nhrpServerEntry.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/nhrpServerTable/" + nhrpServerEntry.EntityData.SegmentPath
    nhrpServerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpServerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpServerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpServerEntry.EntityData.Children = types.NewOrderedMap()
    nhrpServerEntry.EntityData.Leafs = types.NewOrderedMap()
    nhrpServerEntry.EntityData.Leafs.Append("nhrpServerIndex", types.YLeaf{"NhrpServerIndex", nhrpServerEntry.NhrpServerIndex})
    nhrpServerEntry.EntityData.Leafs.Append("nhrpServerInternetworkAddrType", types.YLeaf{"NhrpServerInternetworkAddrType", nhrpServerEntry.NhrpServerInternetworkAddrType})
    nhrpServerEntry.EntityData.Leafs.Append("nhrpServerInternetworkAddr", types.YLeaf{"NhrpServerInternetworkAddr", nhrpServerEntry.NhrpServerInternetworkAddr})
    nhrpServerEntry.EntityData.Leafs.Append("nhrpServerNbmaAddrType", types.YLeaf{"NhrpServerNbmaAddrType", nhrpServerEntry.NhrpServerNbmaAddrType})
    nhrpServerEntry.EntityData.Leafs.Append("nhrpServerNbmaAddr", types.YLeaf{"NhrpServerNbmaAddr", nhrpServerEntry.NhrpServerNbmaAddr})
    nhrpServerEntry.EntityData.Leafs.Append("nhrpServerNbmaSubaddr", types.YLeaf{"NhrpServerNbmaSubaddr", nhrpServerEntry.NhrpServerNbmaSubaddr})
    nhrpServerEntry.EntityData.Leafs.Append("nhrpServerStorageType", types.YLeaf{"NhrpServerStorageType", nhrpServerEntry.NhrpServerStorageType})
    nhrpServerEntry.EntityData.Leafs.Append("nhrpServerRowStatus", types.YLeaf{"NhrpServerRowStatus", nhrpServerEntry.NhrpServerRowStatus})

    nhrpServerEntry.EntityData.YListKeys = []string {"NhrpServerIndex"}

    return &(nhrpServerEntry.EntityData)
}

// NHRPMIB_NhrpServerCacheTable
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
type NHRPMIB_NhrpServerCacheTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Additional information kept by a NHS for a relevant Next Hop Resolution
    // Cache entry. The type is slice of
    // NHRPMIB_NhrpServerCacheTable_NhrpServerCacheEntry.
    NhrpServerCacheEntry []*NHRPMIB_NhrpServerCacheTable_NhrpServerCacheEntry
}

func (nhrpServerCacheTable *NHRPMIB_NhrpServerCacheTable) GetEntityData() *types.CommonEntityData {
    nhrpServerCacheTable.EntityData.YFilter = nhrpServerCacheTable.YFilter
    nhrpServerCacheTable.EntityData.YangName = "nhrpServerCacheTable"
    nhrpServerCacheTable.EntityData.BundleName = "cisco_ios_xe"
    nhrpServerCacheTable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpServerCacheTable.EntityData.SegmentPath = "nhrpServerCacheTable"
    nhrpServerCacheTable.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/" + nhrpServerCacheTable.EntityData.SegmentPath
    nhrpServerCacheTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpServerCacheTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpServerCacheTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpServerCacheTable.EntityData.Children = types.NewOrderedMap()
    nhrpServerCacheTable.EntityData.Children.Append("nhrpServerCacheEntry", types.YChild{"NhrpServerCacheEntry", nil})
    for i := range nhrpServerCacheTable.NhrpServerCacheEntry {
        nhrpServerCacheTable.EntityData.Children.Append(types.GetSegmentPath(nhrpServerCacheTable.NhrpServerCacheEntry[i]), types.YChild{"NhrpServerCacheEntry", nhrpServerCacheTable.NhrpServerCacheEntry[i]})
    }
    nhrpServerCacheTable.EntityData.Leafs = types.NewOrderedMap()

    nhrpServerCacheTable.EntityData.YListKeys = []string {}

    return &(nhrpServerCacheTable.EntityData)
}

// NHRPMIB_NhrpServerCacheTable_NhrpServerCacheEntry
// Additional information kept by a NHS for a relevant
// Next Hop Resolution Cache entry.
type NHRPMIB_NhrpServerCacheTable_NhrpServerCacheEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is AddressFamilyNumbers. Refers to
    // nhrp_mib.NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheInternetworkAddrType
    NhrpCacheInternetworkAddrType interface{}

    // This attribute is a key. The type is string with length: 0..64. Refers to
    // nhrp_mib.NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheInternetworkAddr
    NhrpCacheInternetworkAddr interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to nhrp_mib.NHRPMIB_NhrpCacheTable_NhrpCacheEntry_NhrpCacheIndex
    NhrpCacheIndex interface{}

    // An indication of whether this cache entry is authoritative, which means the
    // entry was added because of a direct registration request with this server
    // or by Server Cache Synchronization Protocol (SCSP) from an authoritative
    // source. The type is bool.
    NhrpServerCacheAuthoritative interface{}

    // The Uniqueness indicator for this cache entry used in duplicate address
    // detection. This value cannot be changed after the entry is active. The type
    // is bool.
    NhrpServerCacheUniqueness interface{}
}

func (nhrpServerCacheEntry *NHRPMIB_NhrpServerCacheTable_NhrpServerCacheEntry) GetEntityData() *types.CommonEntityData {
    nhrpServerCacheEntry.EntityData.YFilter = nhrpServerCacheEntry.YFilter
    nhrpServerCacheEntry.EntityData.YangName = "nhrpServerCacheEntry"
    nhrpServerCacheEntry.EntityData.BundleName = "cisco_ios_xe"
    nhrpServerCacheEntry.EntityData.ParentYangName = "nhrpServerCacheTable"
    nhrpServerCacheEntry.EntityData.SegmentPath = "nhrpServerCacheEntry" + types.AddKeyToken(nhrpServerCacheEntry.NhrpCacheInternetworkAddrType, "nhrpCacheInternetworkAddrType") + types.AddKeyToken(nhrpServerCacheEntry.NhrpCacheInternetworkAddr, "nhrpCacheInternetworkAddr") + types.AddKeyToken(nhrpServerCacheEntry.IfIndex, "ifIndex") + types.AddKeyToken(nhrpServerCacheEntry.NhrpCacheIndex, "nhrpCacheIndex")
    nhrpServerCacheEntry.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/nhrpServerCacheTable/" + nhrpServerCacheEntry.EntityData.SegmentPath
    nhrpServerCacheEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpServerCacheEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpServerCacheEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpServerCacheEntry.EntityData.Children = types.NewOrderedMap()
    nhrpServerCacheEntry.EntityData.Leafs = types.NewOrderedMap()
    nhrpServerCacheEntry.EntityData.Leafs.Append("nhrpCacheInternetworkAddrType", types.YLeaf{"NhrpCacheInternetworkAddrType", nhrpServerCacheEntry.NhrpCacheInternetworkAddrType})
    nhrpServerCacheEntry.EntityData.Leafs.Append("nhrpCacheInternetworkAddr", types.YLeaf{"NhrpCacheInternetworkAddr", nhrpServerCacheEntry.NhrpCacheInternetworkAddr})
    nhrpServerCacheEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", nhrpServerCacheEntry.IfIndex})
    nhrpServerCacheEntry.EntityData.Leafs.Append("nhrpCacheIndex", types.YLeaf{"NhrpCacheIndex", nhrpServerCacheEntry.NhrpCacheIndex})
    nhrpServerCacheEntry.EntityData.Leafs.Append("nhrpServerCacheAuthoritative", types.YLeaf{"NhrpServerCacheAuthoritative", nhrpServerCacheEntry.NhrpServerCacheAuthoritative})
    nhrpServerCacheEntry.EntityData.Leafs.Append("nhrpServerCacheUniqueness", types.YLeaf{"NhrpServerCacheUniqueness", nhrpServerCacheEntry.NhrpServerCacheUniqueness})

    nhrpServerCacheEntry.EntityData.YListKeys = []string {"NhrpCacheInternetworkAddrType", "NhrpCacheInternetworkAddr", "IfIndex", "NhrpCacheIndex"}

    return &(nhrpServerCacheEntry.EntityData)
}

// NHRPMIB_NhrpServerNhcTable
// A table of NHCs that are available for use by this NHS
// (Server).
type NHRPMIB_NhrpServerNhcTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An NHC that may be used by an NHS. The type is slice of
    // NHRPMIB_NhrpServerNhcTable_NhrpServerNhcEntry.
    NhrpServerNhcEntry []*NHRPMIB_NhrpServerNhcTable_NhrpServerNhcEntry
}

func (nhrpServerNhcTable *NHRPMIB_NhrpServerNhcTable) GetEntityData() *types.CommonEntityData {
    nhrpServerNhcTable.EntityData.YFilter = nhrpServerNhcTable.YFilter
    nhrpServerNhcTable.EntityData.YangName = "nhrpServerNhcTable"
    nhrpServerNhcTable.EntityData.BundleName = "cisco_ios_xe"
    nhrpServerNhcTable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpServerNhcTable.EntityData.SegmentPath = "nhrpServerNhcTable"
    nhrpServerNhcTable.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/" + nhrpServerNhcTable.EntityData.SegmentPath
    nhrpServerNhcTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpServerNhcTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpServerNhcTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpServerNhcTable.EntityData.Children = types.NewOrderedMap()
    nhrpServerNhcTable.EntityData.Children.Append("nhrpServerNhcEntry", types.YChild{"NhrpServerNhcEntry", nil})
    for i := range nhrpServerNhcTable.NhrpServerNhcEntry {
        nhrpServerNhcTable.EntityData.Children.Append(types.GetSegmentPath(nhrpServerNhcTable.NhrpServerNhcEntry[i]), types.YChild{"NhrpServerNhcEntry", nhrpServerNhcTable.NhrpServerNhcEntry[i]})
    }
    nhrpServerNhcTable.EntityData.Leafs = types.NewOrderedMap()

    nhrpServerNhcTable.EntityData.YListKeys = []string {}

    return &(nhrpServerNhcTable.EntityData)
}

// NHRPMIB_NhrpServerNhcTable_NhrpServerNhcEntry
// An NHC that may be used by an NHS.
type NHRPMIB_NhrpServerNhcTable_NhrpServerNhcEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to nhrp_mib.NHRPMIB_NhrpServerTable_NhrpServerEntry_NhrpServerIndex
    NhrpServerIndex interface{}

    // This attribute is a key. An identifier for an NHC available to an NHS. The
    // type is interface{} with range: 1..4294967295.
    NhrpServerNhcIndex interface{}

    // The number of bits that define the internetwork layer prefix associated
    // with the nhrpServerNhcInternetworkAddr. The type is interface{} with range:
    // 0..255.
    NhrpServerNhcPrefixLength interface{}

    // The type of the internetwork layer address of the NHRP Client represented
    // in this entry. This object indicates how the value of
    // nhrpServerNhcInternetworkAddr is to be interpreted. The type is
    // AddressFamilyNumbers.
    NhrpServerNhcInternetworkAddrType interface{}

    // The value of the internetwork layer address of the NHRP Client represented
    // by this entry.  If this value is not known, this will be a zero-length
    // OCTET STRING. The type is string with length: 0..64.
    NhrpServerNhcInternetworkAddr interface{}

    // The type of the NBMA subnetwork address of the NHRP Client represented by
    // this entry. This object indicates how the values of nhrpServerNhcNbmaAddr
    // and nhrpServerNhcNbmaSubaddr are to be interpreted. The type is
    // AddressFamilyNumbers.
    NhrpServerNhcNbmaAddrType interface{}

    // The NBMA subnetwork address of the NHC. The type of the address is
    // indicated by the corresponding value of nhrpServerNbmaAddrType. The type is
    // string with length: 0..64.
    NhrpServerNhcNbmaAddr interface{}

    // The NBMA subaddress of the NHC. For NMBA address familes that do not have
    // the concept of subaddress, this will be a zero-length OCTET STRING. The
    // type is string with length: 0..64.
    NhrpServerNhcNbmaSubaddr interface{}

    // An indication of whether this NHC is in use by the NHS. The type is bool.
    NhrpServerNhcInUse interface{}

    // An object that allows entries in this table to be created and deleted using
    // the RowStatus convention. The type is RowStatus.
    NhrpServerNhcRowStatus interface{}
}

func (nhrpServerNhcEntry *NHRPMIB_NhrpServerNhcTable_NhrpServerNhcEntry) GetEntityData() *types.CommonEntityData {
    nhrpServerNhcEntry.EntityData.YFilter = nhrpServerNhcEntry.YFilter
    nhrpServerNhcEntry.EntityData.YangName = "nhrpServerNhcEntry"
    nhrpServerNhcEntry.EntityData.BundleName = "cisco_ios_xe"
    nhrpServerNhcEntry.EntityData.ParentYangName = "nhrpServerNhcTable"
    nhrpServerNhcEntry.EntityData.SegmentPath = "nhrpServerNhcEntry" + types.AddKeyToken(nhrpServerNhcEntry.NhrpServerIndex, "nhrpServerIndex") + types.AddKeyToken(nhrpServerNhcEntry.NhrpServerNhcIndex, "nhrpServerNhcIndex")
    nhrpServerNhcEntry.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/nhrpServerNhcTable/" + nhrpServerNhcEntry.EntityData.SegmentPath
    nhrpServerNhcEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpServerNhcEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpServerNhcEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpServerNhcEntry.EntityData.Children = types.NewOrderedMap()
    nhrpServerNhcEntry.EntityData.Leafs = types.NewOrderedMap()
    nhrpServerNhcEntry.EntityData.Leafs.Append("nhrpServerIndex", types.YLeaf{"NhrpServerIndex", nhrpServerNhcEntry.NhrpServerIndex})
    nhrpServerNhcEntry.EntityData.Leafs.Append("nhrpServerNhcIndex", types.YLeaf{"NhrpServerNhcIndex", nhrpServerNhcEntry.NhrpServerNhcIndex})
    nhrpServerNhcEntry.EntityData.Leafs.Append("nhrpServerNhcPrefixLength", types.YLeaf{"NhrpServerNhcPrefixLength", nhrpServerNhcEntry.NhrpServerNhcPrefixLength})
    nhrpServerNhcEntry.EntityData.Leafs.Append("nhrpServerNhcInternetworkAddrType", types.YLeaf{"NhrpServerNhcInternetworkAddrType", nhrpServerNhcEntry.NhrpServerNhcInternetworkAddrType})
    nhrpServerNhcEntry.EntityData.Leafs.Append("nhrpServerNhcInternetworkAddr", types.YLeaf{"NhrpServerNhcInternetworkAddr", nhrpServerNhcEntry.NhrpServerNhcInternetworkAddr})
    nhrpServerNhcEntry.EntityData.Leafs.Append("nhrpServerNhcNbmaAddrType", types.YLeaf{"NhrpServerNhcNbmaAddrType", nhrpServerNhcEntry.NhrpServerNhcNbmaAddrType})
    nhrpServerNhcEntry.EntityData.Leafs.Append("nhrpServerNhcNbmaAddr", types.YLeaf{"NhrpServerNhcNbmaAddr", nhrpServerNhcEntry.NhrpServerNhcNbmaAddr})
    nhrpServerNhcEntry.EntityData.Leafs.Append("nhrpServerNhcNbmaSubaddr", types.YLeaf{"NhrpServerNhcNbmaSubaddr", nhrpServerNhcEntry.NhrpServerNhcNbmaSubaddr})
    nhrpServerNhcEntry.EntityData.Leafs.Append("nhrpServerNhcInUse", types.YLeaf{"NhrpServerNhcInUse", nhrpServerNhcEntry.NhrpServerNhcInUse})
    nhrpServerNhcEntry.EntityData.Leafs.Append("nhrpServerNhcRowStatus", types.YLeaf{"NhrpServerNhcRowStatus", nhrpServerNhcEntry.NhrpServerNhcRowStatus})

    nhrpServerNhcEntry.EntityData.YListKeys = []string {"NhrpServerIndex", "NhrpServerNhcIndex"}

    return &(nhrpServerNhcEntry.EntityData)
}

// NHRPMIB_NhrpServerStatTable
// Statistics collected by Next Hop Servers.
type NHRPMIB_NhrpServerStatTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics for a particular NHS. The statistics are broken into received
    // (Rx), transmitted (Tx) and forwarded (Fw).  Forwarded (Fw) would be done by
    // a transit NHS. The type is slice of
    // NHRPMIB_NhrpServerStatTable_NhrpServerStatEntry.
    NhrpServerStatEntry []*NHRPMIB_NhrpServerStatTable_NhrpServerStatEntry
}

func (nhrpServerStatTable *NHRPMIB_NhrpServerStatTable) GetEntityData() *types.CommonEntityData {
    nhrpServerStatTable.EntityData.YFilter = nhrpServerStatTable.YFilter
    nhrpServerStatTable.EntityData.YangName = "nhrpServerStatTable"
    nhrpServerStatTable.EntityData.BundleName = "cisco_ios_xe"
    nhrpServerStatTable.EntityData.ParentYangName = "NHRP-MIB"
    nhrpServerStatTable.EntityData.SegmentPath = "nhrpServerStatTable"
    nhrpServerStatTable.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/" + nhrpServerStatTable.EntityData.SegmentPath
    nhrpServerStatTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpServerStatTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpServerStatTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpServerStatTable.EntityData.Children = types.NewOrderedMap()
    nhrpServerStatTable.EntityData.Children.Append("nhrpServerStatEntry", types.YChild{"NhrpServerStatEntry", nil})
    for i := range nhrpServerStatTable.NhrpServerStatEntry {
        nhrpServerStatTable.EntityData.Children.Append(types.GetSegmentPath(nhrpServerStatTable.NhrpServerStatEntry[i]), types.YChild{"NhrpServerStatEntry", nhrpServerStatTable.NhrpServerStatEntry[i]})
    }
    nhrpServerStatTable.EntityData.Leafs = types.NewOrderedMap()

    nhrpServerStatTable.EntityData.YListKeys = []string {}

    return &(nhrpServerStatTable.EntityData)
}

// NHRPMIB_NhrpServerStatTable_NhrpServerStatEntry
// Statistics for a particular NHS. The statistics are
// broken into received (Rx), transmitted (Tx)
// and forwarded (Fw).  Forwarded (Fw) would be done
// by a transit NHS.
type NHRPMIB_NhrpServerStatTable_NhrpServerStatEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to nhrp_mib.NHRPMIB_NhrpServerTable_NhrpServerEntry_NhrpServerIndex
    NhrpServerIndex interface{}

    // The number of NHRP Resolution Requests received by this server. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Server re-initialization and at other
    // times as indicated by the value of nhrpServerStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpServerStatRxResolveReq interface{}

    // The number of positively acknowledged NHRP Resolution Replies transmitted
    // by this server.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxResolveReplyAck interface{}

    // The number of NAKed NHRP Resolution Replies transmitted by this server with
    // the code 'Administratively Prohibited'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxResolveReplyNakProhibited interface{}

    // The number of NAKed NHRP Resolution Replies transmitted by this server with
    // the code 'Insufficient Resources'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxResolveReplyNakInsufResources interface{}

    // The number of NAKed NHRP Resolution Replies transmitted by this server with
    // the code 'No Internetworking Layer Address to NBMA Address Binding Exists'.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Server re-initialization and at other
    // times as indicated by the value of nhrpServerStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpServerStatTxResolveReplyNakNoBinding interface{}

    // The number of NAKed NHRP Resolution Replies transmitted by this server with
    // the code 'Binding Exists But Is Not Unique'.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxResolveReplyNakNotUnique interface{}

    // The number of NHRP Registration Requests received by this server. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Server re-initialization and at other
    // times as indicated by the value of nhrpServerStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpServerStatRxRegisterReq interface{}

    // The number of positively acknowledged NHRP Registration Replies transmitted
    // by this server.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxRegisterAck interface{}

    // The number of NAKed NHRP Registration Replies transmitted by this server
    // with the code 'Administratively Prohibited'.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxRegisterNakProhibited interface{}

    // The number of NAKed NHRP Registration Replies transmitted by this server
    // with the code 'Insufficient Resources'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxRegisterNakInsufResources interface{}

    // The number of NAKed NHRP Registration Replies transmitted by this server
    // with the code 'Unique Internetworking Layer Address Already Registered'. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Server re-initialization and at other
    // times as indicated by the value of nhrpServerStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpServerStatTxRegisterNakAlreadyReg interface{}

    // The number of NHRP Purge Requests received by this server.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, at NHRP Server re-initialization and at other times as
    // indicated by the value of nhrpServerStatDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    NhrpServerStatRxPurgeReq interface{}

    // The number of NHRP Purge Requests transmitted by this server. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Server re-initialization and at other
    // times as indicated by the value of nhrpServerStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpServerStatTxPurgeReq interface{}

    // The number of NHRP Purge Replies received by this server.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, at NHRP Server re-initialization and at other times as
    // indicated by the value of nhrpServerStatDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    NhrpServerStatRxPurgeReply interface{}

    // The number of NHRP Purge Replies transmitted by this server. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, at NHRP Server re-initialization and at other
    // times as indicated by the value of nhrpServerStatDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    NhrpServerStatTxPurgeReply interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code  'Unrecognized Extension'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatRxErrUnrecognizedExtension interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'NHRP Loop Detected'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatRxErrLoopDetected interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'Protocol Address Unreachable'.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, at NHRP Server re-initialization and at other times as indicated by
    // the value of nhrpServerStatDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    NhrpServerStatRxErrProtoAddrUnreachable interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'Protocol Error'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatRxErrProtoError interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'NHRP SDU Size Exceeded'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatRxErrSduSizeExceeded interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'Invalid Extension'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatRxErrInvalidExtension interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'Invalid Resolution Reply Received'.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, at NHRP Server re-initialization and at other times as indicated by
    // the value of nhrpServerStatDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    NhrpServerStatRxErrInvalidResReplyReceived interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'Authentication Failure'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatRxErrAuthenticationFailure interface{}

    // The number of NHRP Error Indication packets received by this server with
    // the error code 'Hop Count Exceeded'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatRxErrHopCountExceeded interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'Unrecognized Extension'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxErrUnrecognizedExtension interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'NHRP Loop Detected'.     Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxErrLoopDetected interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'Protocol Address Unreachable'.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, at NHRP Server re-initialization and at other times as indicated by
    // the value of nhrpServerStatDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    NhrpServerStatTxErrProtoAddrUnreachable interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'Protocol Error'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxErrProtoError interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'NHRP SDU Size Exceeded'.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxErrSduSizeExceeded interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code  'Invalid Extension'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxErrInvalidExtension interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'Authentication Failure'.     Discontinuities in the value
    // of this counter can occur at re-initialization of the management system, at
    // NHRP Server re-initialization and at other times as indicated by the value
    // of nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxErrAuthenticationFailure interface{}

    // The number of NHRP Error Indication packets transmitted by this server with
    // the error code 'Hop Count Exceeded'.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, at NHRP
    // Server re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatTxErrHopCountExceeded interface{}

    // The number of NHRP Resolution Requests forwarded by this server acting as a
    // transit NHS.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatFwResolveReq interface{}

    // The number of NHRP Resolution Replies forwarded by this server acting as a
    // transit NHS.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatFwResolveReply interface{}

    // The number of NHRP Registration Requests forwarded by this server acting as
    // a transit NHS.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatFwRegisterReq interface{}

    // The number of NHRP Registration Replies forwarded by this server acting as
    // a transit NHS. Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatFwRegisterReply interface{}

    // The number of NHRP Purge Requests forwarded by this server acting as a
    // transit NHS.   Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatFwPurgeReq interface{}

    // The number of NHRP Purge Replies forwarded by this server acting as a
    // transit NHS.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatFwPurgeReply interface{}

    // The number of NHRP Error Indication packets forwarded by this server acting
    // as a transit NHS.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, at NHRP Server
    // re-initialization and at other times as indicated by the value of
    // nhrpServerStatDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    NhrpServerStatFwErrorIndication interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this Server's counters suffered a discontinuity.  If no such
    // discontinuities have occurred since the last re-initialization of the   
    // local management subsystem or the NHRP Server re-initialization associated
    // with this entry, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    NhrpServerStatDiscontinuityTime interface{}
}

func (nhrpServerStatEntry *NHRPMIB_NhrpServerStatTable_NhrpServerStatEntry) GetEntityData() *types.CommonEntityData {
    nhrpServerStatEntry.EntityData.YFilter = nhrpServerStatEntry.YFilter
    nhrpServerStatEntry.EntityData.YangName = "nhrpServerStatEntry"
    nhrpServerStatEntry.EntityData.BundleName = "cisco_ios_xe"
    nhrpServerStatEntry.EntityData.ParentYangName = "nhrpServerStatTable"
    nhrpServerStatEntry.EntityData.SegmentPath = "nhrpServerStatEntry" + types.AddKeyToken(nhrpServerStatEntry.NhrpServerIndex, "nhrpServerIndex")
    nhrpServerStatEntry.EntityData.AbsolutePath = "NHRP-MIB:NHRP-MIB/nhrpServerStatTable/" + nhrpServerStatEntry.EntityData.SegmentPath
    nhrpServerStatEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhrpServerStatEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhrpServerStatEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhrpServerStatEntry.EntityData.Children = types.NewOrderedMap()
    nhrpServerStatEntry.EntityData.Leafs = types.NewOrderedMap()
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerIndex", types.YLeaf{"NhrpServerIndex", nhrpServerStatEntry.NhrpServerIndex})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatRxResolveReq", types.YLeaf{"NhrpServerStatRxResolveReq", nhrpServerStatEntry.NhrpServerStatRxResolveReq})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxResolveReplyAck", types.YLeaf{"NhrpServerStatTxResolveReplyAck", nhrpServerStatEntry.NhrpServerStatTxResolveReplyAck})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxResolveReplyNakProhibited", types.YLeaf{"NhrpServerStatTxResolveReplyNakProhibited", nhrpServerStatEntry.NhrpServerStatTxResolveReplyNakProhibited})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxResolveReplyNakInsufResources", types.YLeaf{"NhrpServerStatTxResolveReplyNakInsufResources", nhrpServerStatEntry.NhrpServerStatTxResolveReplyNakInsufResources})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxResolveReplyNakNoBinding", types.YLeaf{"NhrpServerStatTxResolveReplyNakNoBinding", nhrpServerStatEntry.NhrpServerStatTxResolveReplyNakNoBinding})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxResolveReplyNakNotUnique", types.YLeaf{"NhrpServerStatTxResolveReplyNakNotUnique", nhrpServerStatEntry.NhrpServerStatTxResolveReplyNakNotUnique})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatRxRegisterReq", types.YLeaf{"NhrpServerStatRxRegisterReq", nhrpServerStatEntry.NhrpServerStatRxRegisterReq})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxRegisterAck", types.YLeaf{"NhrpServerStatTxRegisterAck", nhrpServerStatEntry.NhrpServerStatTxRegisterAck})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxRegisterNakProhibited", types.YLeaf{"NhrpServerStatTxRegisterNakProhibited", nhrpServerStatEntry.NhrpServerStatTxRegisterNakProhibited})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxRegisterNakInsufResources", types.YLeaf{"NhrpServerStatTxRegisterNakInsufResources", nhrpServerStatEntry.NhrpServerStatTxRegisterNakInsufResources})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxRegisterNakAlreadyReg", types.YLeaf{"NhrpServerStatTxRegisterNakAlreadyReg", nhrpServerStatEntry.NhrpServerStatTxRegisterNakAlreadyReg})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatRxPurgeReq", types.YLeaf{"NhrpServerStatRxPurgeReq", nhrpServerStatEntry.NhrpServerStatRxPurgeReq})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxPurgeReq", types.YLeaf{"NhrpServerStatTxPurgeReq", nhrpServerStatEntry.NhrpServerStatTxPurgeReq})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatRxPurgeReply", types.YLeaf{"NhrpServerStatRxPurgeReply", nhrpServerStatEntry.NhrpServerStatRxPurgeReply})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxPurgeReply", types.YLeaf{"NhrpServerStatTxPurgeReply", nhrpServerStatEntry.NhrpServerStatTxPurgeReply})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatRxErrUnrecognizedExtension", types.YLeaf{"NhrpServerStatRxErrUnrecognizedExtension", nhrpServerStatEntry.NhrpServerStatRxErrUnrecognizedExtension})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatRxErrLoopDetected", types.YLeaf{"NhrpServerStatRxErrLoopDetected", nhrpServerStatEntry.NhrpServerStatRxErrLoopDetected})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatRxErrProtoAddrUnreachable", types.YLeaf{"NhrpServerStatRxErrProtoAddrUnreachable", nhrpServerStatEntry.NhrpServerStatRxErrProtoAddrUnreachable})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatRxErrProtoError", types.YLeaf{"NhrpServerStatRxErrProtoError", nhrpServerStatEntry.NhrpServerStatRxErrProtoError})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatRxErrSduSizeExceeded", types.YLeaf{"NhrpServerStatRxErrSduSizeExceeded", nhrpServerStatEntry.NhrpServerStatRxErrSduSizeExceeded})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatRxErrInvalidExtension", types.YLeaf{"NhrpServerStatRxErrInvalidExtension", nhrpServerStatEntry.NhrpServerStatRxErrInvalidExtension})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatRxErrInvalidResReplyReceived", types.YLeaf{"NhrpServerStatRxErrInvalidResReplyReceived", nhrpServerStatEntry.NhrpServerStatRxErrInvalidResReplyReceived})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatRxErrAuthenticationFailure", types.YLeaf{"NhrpServerStatRxErrAuthenticationFailure", nhrpServerStatEntry.NhrpServerStatRxErrAuthenticationFailure})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatRxErrHopCountExceeded", types.YLeaf{"NhrpServerStatRxErrHopCountExceeded", nhrpServerStatEntry.NhrpServerStatRxErrHopCountExceeded})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxErrUnrecognizedExtension", types.YLeaf{"NhrpServerStatTxErrUnrecognizedExtension", nhrpServerStatEntry.NhrpServerStatTxErrUnrecognizedExtension})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxErrLoopDetected", types.YLeaf{"NhrpServerStatTxErrLoopDetected", nhrpServerStatEntry.NhrpServerStatTxErrLoopDetected})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxErrProtoAddrUnreachable", types.YLeaf{"NhrpServerStatTxErrProtoAddrUnreachable", nhrpServerStatEntry.NhrpServerStatTxErrProtoAddrUnreachable})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxErrProtoError", types.YLeaf{"NhrpServerStatTxErrProtoError", nhrpServerStatEntry.NhrpServerStatTxErrProtoError})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxErrSduSizeExceeded", types.YLeaf{"NhrpServerStatTxErrSduSizeExceeded", nhrpServerStatEntry.NhrpServerStatTxErrSduSizeExceeded})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxErrInvalidExtension", types.YLeaf{"NhrpServerStatTxErrInvalidExtension", nhrpServerStatEntry.NhrpServerStatTxErrInvalidExtension})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxErrAuthenticationFailure", types.YLeaf{"NhrpServerStatTxErrAuthenticationFailure", nhrpServerStatEntry.NhrpServerStatTxErrAuthenticationFailure})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatTxErrHopCountExceeded", types.YLeaf{"NhrpServerStatTxErrHopCountExceeded", nhrpServerStatEntry.NhrpServerStatTxErrHopCountExceeded})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatFwResolveReq", types.YLeaf{"NhrpServerStatFwResolveReq", nhrpServerStatEntry.NhrpServerStatFwResolveReq})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatFwResolveReply", types.YLeaf{"NhrpServerStatFwResolveReply", nhrpServerStatEntry.NhrpServerStatFwResolveReply})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatFwRegisterReq", types.YLeaf{"NhrpServerStatFwRegisterReq", nhrpServerStatEntry.NhrpServerStatFwRegisterReq})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatFwRegisterReply", types.YLeaf{"NhrpServerStatFwRegisterReply", nhrpServerStatEntry.NhrpServerStatFwRegisterReply})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatFwPurgeReq", types.YLeaf{"NhrpServerStatFwPurgeReq", nhrpServerStatEntry.NhrpServerStatFwPurgeReq})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatFwPurgeReply", types.YLeaf{"NhrpServerStatFwPurgeReply", nhrpServerStatEntry.NhrpServerStatFwPurgeReply})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatFwErrorIndication", types.YLeaf{"NhrpServerStatFwErrorIndication", nhrpServerStatEntry.NhrpServerStatFwErrorIndication})
    nhrpServerStatEntry.EntityData.Leafs.Append("nhrpServerStatDiscontinuityTime", types.YLeaf{"NhrpServerStatDiscontinuityTime", nhrpServerStatEntry.NhrpServerStatDiscontinuityTime})

    nhrpServerStatEntry.EntityData.YListKeys = []string {"NhrpServerIndex"}

    return &(nhrpServerStatEntry.EntityData)
}

