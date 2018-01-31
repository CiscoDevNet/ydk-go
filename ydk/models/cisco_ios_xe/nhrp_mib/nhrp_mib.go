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
    parent types.Entity
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

func (nHRPMIB *NHRPMIB) GetFilter() yfilter.YFilter { return nHRPMIB.YFilter }

func (nHRPMIB *NHRPMIB) SetFilter(yf yfilter.YFilter) { nHRPMIB.YFilter = yf }

func (nHRPMIB *NHRPMIB) GetGoName(yname string) string {
    if yname == "nhrpGeneralObjects" { return "Nhrpgeneralobjects" }
    if yname == "nhrpCacheTable" { return "Nhrpcachetable" }
    if yname == "nhrpPurgeReqTable" { return "Nhrppurgereqtable" }
    if yname == "nhrpClientTable" { return "Nhrpclienttable" }
    if yname == "nhrpClientRegistrationTable" { return "Nhrpclientregistrationtable" }
    if yname == "nhrpClientNhsTable" { return "Nhrpclientnhstable" }
    if yname == "nhrpClientStatTable" { return "Nhrpclientstattable" }
    if yname == "nhrpServerTable" { return "Nhrpservertable" }
    if yname == "nhrpServerCacheTable" { return "Nhrpservercachetable" }
    if yname == "nhrpServerNhcTable" { return "Nhrpservernhctable" }
    if yname == "nhrpServerStatTable" { return "Nhrpserverstattable" }
    return ""
}

func (nHRPMIB *NHRPMIB) GetSegmentPath() string {
    return "NHRP-MIB:NHRP-MIB"
}

func (nHRPMIB *NHRPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nhrpGeneralObjects" {
        return &nHRPMIB.Nhrpgeneralobjects
    }
    if childYangName == "nhrpCacheTable" {
        return &nHRPMIB.Nhrpcachetable
    }
    if childYangName == "nhrpPurgeReqTable" {
        return &nHRPMIB.Nhrppurgereqtable
    }
    if childYangName == "nhrpClientTable" {
        return &nHRPMIB.Nhrpclienttable
    }
    if childYangName == "nhrpClientRegistrationTable" {
        return &nHRPMIB.Nhrpclientregistrationtable
    }
    if childYangName == "nhrpClientNhsTable" {
        return &nHRPMIB.Nhrpclientnhstable
    }
    if childYangName == "nhrpClientStatTable" {
        return &nHRPMIB.Nhrpclientstattable
    }
    if childYangName == "nhrpServerTable" {
        return &nHRPMIB.Nhrpservertable
    }
    if childYangName == "nhrpServerCacheTable" {
        return &nHRPMIB.Nhrpservercachetable
    }
    if childYangName == "nhrpServerNhcTable" {
        return &nHRPMIB.Nhrpservernhctable
    }
    if childYangName == "nhrpServerStatTable" {
        return &nHRPMIB.Nhrpserverstattable
    }
    return nil
}

func (nHRPMIB *NHRPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nhrpGeneralObjects"] = &nHRPMIB.Nhrpgeneralobjects
    children["nhrpCacheTable"] = &nHRPMIB.Nhrpcachetable
    children["nhrpPurgeReqTable"] = &nHRPMIB.Nhrppurgereqtable
    children["nhrpClientTable"] = &nHRPMIB.Nhrpclienttable
    children["nhrpClientRegistrationTable"] = &nHRPMIB.Nhrpclientregistrationtable
    children["nhrpClientNhsTable"] = &nHRPMIB.Nhrpclientnhstable
    children["nhrpClientStatTable"] = &nHRPMIB.Nhrpclientstattable
    children["nhrpServerTable"] = &nHRPMIB.Nhrpservertable
    children["nhrpServerCacheTable"] = &nHRPMIB.Nhrpservercachetable
    children["nhrpServerNhcTable"] = &nHRPMIB.Nhrpservernhctable
    children["nhrpServerStatTable"] = &nHRPMIB.Nhrpserverstattable
    return children
}

func (nHRPMIB *NHRPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nHRPMIB *NHRPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (nHRPMIB *NHRPMIB) GetYangName() string { return "NHRP-MIB" }

func (nHRPMIB *NHRPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nHRPMIB *NHRPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nHRPMIB *NHRPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nHRPMIB *NHRPMIB) SetParent(parent types.Entity) { nHRPMIB.parent = parent }

func (nHRPMIB *NHRPMIB) GetParent() types.Entity { return nHRPMIB.parent }

func (nHRPMIB *NHRPMIB) GetParentYangName() string { return "NHRP-MIB" }

// NHRPMIB_Nhrpgeneralobjects
type NHRPMIB_Nhrpgeneralobjects struct {
    parent types.Entity
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

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetFilter() yfilter.YFilter { return nhrpgeneralobjects.YFilter }

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) SetFilter(yf yfilter.YFilter) { nhrpgeneralobjects.YFilter = yf }

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetGoName(yname string) string {
    if yname == "nhrpNextIndex" { return "Nhrpnextindex" }
    return ""
}

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetSegmentPath() string {
    return "nhrpGeneralObjects"
}

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nhrpNextIndex"] = nhrpgeneralobjects.Nhrpnextindex
    return leafs
}

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetYangName() string { return "nhrpGeneralObjects" }

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) SetParent(parent types.Entity) { nhrpgeneralobjects.parent = parent }

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetParent() types.Entity { return nhrpgeneralobjects.parent }

func (nhrpgeneralobjects *NHRPMIB_Nhrpgeneralobjects) GetParentYangName() string { return "NHRP-MIB" }

// NHRPMIB_Nhrpcachetable
// This table contains mappings between internetwork layer
// addresses and NBMA subnetwork layer addresses.
type NHRPMIB_Nhrpcachetable struct {
    parent types.Entity
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

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetFilter() yfilter.YFilter { return nhrpcachetable.YFilter }

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) SetFilter(yf yfilter.YFilter) { nhrpcachetable.YFilter = yf }

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetGoName(yname string) string {
    if yname == "nhrpCacheEntry" { return "Nhrpcacheentry" }
    return ""
}

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetSegmentPath() string {
    return "nhrpCacheTable"
}

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nhrpCacheEntry" {
        for _, c := range nhrpcachetable.Nhrpcacheentry {
            if nhrpcachetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NHRPMIB_Nhrpcachetable_Nhrpcacheentry{}
        nhrpcachetable.Nhrpcacheentry = append(nhrpcachetable.Nhrpcacheentry, child)
        return &nhrpcachetable.Nhrpcacheentry[len(nhrpcachetable.Nhrpcacheentry)-1]
    }
    return nil
}

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nhrpcachetable.Nhrpcacheentry {
        children[nhrpcachetable.Nhrpcacheentry[i].GetSegmentPath()] = &nhrpcachetable.Nhrpcacheentry[i]
    }
    return children
}

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetYangName() string { return "nhrpCacheTable" }

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) SetParent(parent types.Entity) { nhrpcachetable.parent = parent }

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetParent() types.Entity { return nhrpcachetable.parent }

func (nhrpcachetable *NHRPMIB_Nhrpcachetable) GetParentYangName() string { return "NHRP-MIB" }

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
    parent types.Entity
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

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetFilter() yfilter.YFilter { return nhrpcacheentry.YFilter }

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) SetFilter(yf yfilter.YFilter) { nhrpcacheentry.YFilter = yf }

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetGoName(yname string) string {
    if yname == "nhrpCacheInternetworkAddrType" { return "Nhrpcacheinternetworkaddrtype" }
    if yname == "nhrpCacheInternetworkAddr" { return "Nhrpcacheinternetworkaddr" }
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "nhrpCacheIndex" { return "Nhrpcacheindex" }
    if yname == "nhrpCachePrefixLength" { return "Nhrpcacheprefixlength" }
    if yname == "nhrpCacheNextHopInternetworkAddr" { return "Nhrpcachenexthopinternetworkaddr" }
    if yname == "nhrpCacheNbmaAddrType" { return "Nhrpcachenbmaaddrtype" }
    if yname == "nhrpCacheNbmaAddr" { return "Nhrpcachenbmaaddr" }
    if yname == "nhrpCacheNbmaSubaddr" { return "Nhrpcachenbmasubaddr" }
    if yname == "nhrpCacheType" { return "Nhrpcachetype" }
    if yname == "nhrpCacheState" { return "Nhrpcachestate" }
    if yname == "nhrpCacheHoldingTimeValid" { return "Nhrpcacheholdingtimevalid" }
    if yname == "nhrpCacheHoldingTime" { return "Nhrpcacheholdingtime" }
    if yname == "nhrpCacheNegotiatedMtu" { return "Nhrpcachenegotiatedmtu" }
    if yname == "nhrpCachePreference" { return "Nhrpcachepreference" }
    if yname == "nhrpCacheStorageType" { return "Nhrpcachestoragetype" }
    if yname == "nhrpCacheRowStatus" { return "Nhrpcacherowstatus" }
    return ""
}

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetSegmentPath() string {
    return "nhrpCacheEntry" + "[nhrpCacheInternetworkAddrType='" + fmt.Sprintf("%v", nhrpcacheentry.Nhrpcacheinternetworkaddrtype) + "']" + "[nhrpCacheInternetworkAddr='" + fmt.Sprintf("%v", nhrpcacheentry.Nhrpcacheinternetworkaddr) + "']" + "[ifIndex='" + fmt.Sprintf("%v", nhrpcacheentry.Ifindex) + "']" + "[nhrpCacheIndex='" + fmt.Sprintf("%v", nhrpcacheentry.Nhrpcacheindex) + "']"
}

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nhrpCacheInternetworkAddrType"] = nhrpcacheentry.Nhrpcacheinternetworkaddrtype
    leafs["nhrpCacheInternetworkAddr"] = nhrpcacheentry.Nhrpcacheinternetworkaddr
    leafs["ifIndex"] = nhrpcacheentry.Ifindex
    leafs["nhrpCacheIndex"] = nhrpcacheentry.Nhrpcacheindex
    leafs["nhrpCachePrefixLength"] = nhrpcacheentry.Nhrpcacheprefixlength
    leafs["nhrpCacheNextHopInternetworkAddr"] = nhrpcacheentry.Nhrpcachenexthopinternetworkaddr
    leafs["nhrpCacheNbmaAddrType"] = nhrpcacheentry.Nhrpcachenbmaaddrtype
    leafs["nhrpCacheNbmaAddr"] = nhrpcacheentry.Nhrpcachenbmaaddr
    leafs["nhrpCacheNbmaSubaddr"] = nhrpcacheentry.Nhrpcachenbmasubaddr
    leafs["nhrpCacheType"] = nhrpcacheentry.Nhrpcachetype
    leafs["nhrpCacheState"] = nhrpcacheentry.Nhrpcachestate
    leafs["nhrpCacheHoldingTimeValid"] = nhrpcacheentry.Nhrpcacheholdingtimevalid
    leafs["nhrpCacheHoldingTime"] = nhrpcacheentry.Nhrpcacheholdingtime
    leafs["nhrpCacheNegotiatedMtu"] = nhrpcacheentry.Nhrpcachenegotiatedmtu
    leafs["nhrpCachePreference"] = nhrpcacheentry.Nhrpcachepreference
    leafs["nhrpCacheStorageType"] = nhrpcacheentry.Nhrpcachestoragetype
    leafs["nhrpCacheRowStatus"] = nhrpcacheentry.Nhrpcacherowstatus
    return leafs
}

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetYangName() string { return "nhrpCacheEntry" }

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) SetParent(parent types.Entity) { nhrpcacheentry.parent = parent }

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetParent() types.Entity { return nhrpcacheentry.parent }

func (nhrpcacheentry *NHRPMIB_Nhrpcachetable_Nhrpcacheentry) GetParentYangName() string { return "nhrpCacheTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information regarding a Purge Request. The type is slice of
    // NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry.
    Nhrppurgereqentry []NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry
}

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetFilter() yfilter.YFilter { return nhrppurgereqtable.YFilter }

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) SetFilter(yf yfilter.YFilter) { nhrppurgereqtable.YFilter = yf }

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetGoName(yname string) string {
    if yname == "nhrpPurgeReqEntry" { return "Nhrppurgereqentry" }
    return ""
}

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetSegmentPath() string {
    return "nhrpPurgeReqTable"
}

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nhrpPurgeReqEntry" {
        for _, c := range nhrppurgereqtable.Nhrppurgereqentry {
            if nhrppurgereqtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry{}
        nhrppurgereqtable.Nhrppurgereqentry = append(nhrppurgereqtable.Nhrppurgereqentry, child)
        return &nhrppurgereqtable.Nhrppurgereqentry[len(nhrppurgereqtable.Nhrppurgereqentry)-1]
    }
    return nil
}

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nhrppurgereqtable.Nhrppurgereqentry {
        children[nhrppurgereqtable.Nhrppurgereqentry[i].GetSegmentPath()] = &nhrppurgereqtable.Nhrppurgereqentry[i]
    }
    return children
}

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetBundleName() string { return "cisco_ios_xe" }

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetYangName() string { return "nhrpPurgeReqTable" }

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) SetParent(parent types.Entity) { nhrppurgereqtable.parent = parent }

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetParent() types.Entity { return nhrppurgereqtable.parent }

func (nhrppurgereqtable *NHRPMIB_Nhrppurgereqtable) GetParentYangName() string { return "NHRP-MIB" }

// NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry
// Information regarding a Purge Request.
type NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry struct {
    parent types.Entity
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

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetFilter() yfilter.YFilter { return nhrppurgereqentry.YFilter }

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) SetFilter(yf yfilter.YFilter) { nhrppurgereqentry.YFilter = yf }

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetGoName(yname string) string {
    if yname == "nhrpPurgeIndex" { return "Nhrppurgeindex" }
    if yname == "nhrpPurgeCacheIdentifier" { return "Nhrppurgecacheidentifier" }
    if yname == "nhrpPurgePrefixLength" { return "Nhrppurgeprefixlength" }
    if yname == "nhrpPurgeRequestID" { return "Nhrppurgerequestid" }
    if yname == "nhrpPurgeReplyExpected" { return "Nhrppurgereplyexpected" }
    if yname == "nhrpPurgeRowStatus" { return "Nhrppurgerowstatus" }
    return ""
}

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetSegmentPath() string {
    return "nhrpPurgeReqEntry" + "[nhrpPurgeIndex='" + fmt.Sprintf("%v", nhrppurgereqentry.Nhrppurgeindex) + "']"
}

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nhrpPurgeIndex"] = nhrppurgereqentry.Nhrppurgeindex
    leafs["nhrpPurgeCacheIdentifier"] = nhrppurgereqentry.Nhrppurgecacheidentifier
    leafs["nhrpPurgePrefixLength"] = nhrppurgereqentry.Nhrppurgeprefixlength
    leafs["nhrpPurgeRequestID"] = nhrppurgereqentry.Nhrppurgerequestid
    leafs["nhrpPurgeReplyExpected"] = nhrppurgereqentry.Nhrppurgereplyexpected
    leafs["nhrpPurgeRowStatus"] = nhrppurgereqentry.Nhrppurgerowstatus
    return leafs
}

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetBundleName() string { return "cisco_ios_xe" }

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetYangName() string { return "nhrpPurgeReqEntry" }

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) SetParent(parent types.Entity) { nhrppurgereqentry.parent = parent }

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetParent() types.Entity { return nhrppurgereqentry.parent }

func (nhrppurgereqentry *NHRPMIB_Nhrppurgereqtable_Nhrppurgereqentry) GetParentYangName() string { return "nhrpPurgeReqTable" }

// NHRPMIB_Nhrpclienttable
// Information about NHRP clients (NHCs) managed by this
// agent.
type NHRPMIB_Nhrpclienttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single NHC. The type is slice of
    // NHRPMIB_Nhrpclienttable_Nhrpcliententry.
    Nhrpcliententry []NHRPMIB_Nhrpclienttable_Nhrpcliententry
}

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetFilter() yfilter.YFilter { return nhrpclienttable.YFilter }

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) SetFilter(yf yfilter.YFilter) { nhrpclienttable.YFilter = yf }

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetGoName(yname string) string {
    if yname == "nhrpClientEntry" { return "Nhrpcliententry" }
    return ""
}

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetSegmentPath() string {
    return "nhrpClientTable"
}

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nhrpClientEntry" {
        for _, c := range nhrpclienttable.Nhrpcliententry {
            if nhrpclienttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NHRPMIB_Nhrpclienttable_Nhrpcliententry{}
        nhrpclienttable.Nhrpcliententry = append(nhrpclienttable.Nhrpcliententry, child)
        return &nhrpclienttable.Nhrpcliententry[len(nhrpclienttable.Nhrpcliententry)-1]
    }
    return nil
}

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nhrpclienttable.Nhrpcliententry {
        children[nhrpclienttable.Nhrpcliententry[i].GetSegmentPath()] = &nhrpclienttable.Nhrpcliententry[i]
    }
    return children
}

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetYangName() string { return "nhrpClientTable" }

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) SetParent(parent types.Entity) { nhrpclienttable.parent = parent }

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetParent() types.Entity { return nhrpclienttable.parent }

func (nhrpclienttable *NHRPMIB_Nhrpclienttable) GetParentYangName() string { return "NHRP-MIB" }

// NHRPMIB_Nhrpclienttable_Nhrpcliententry
// Information about a single NHC.
type NHRPMIB_Nhrpclienttable_Nhrpcliententry struct {
    parent types.Entity
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

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetFilter() yfilter.YFilter { return nhrpcliententry.YFilter }

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) SetFilter(yf yfilter.YFilter) { nhrpcliententry.YFilter = yf }

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetGoName(yname string) string {
    if yname == "nhrpClientIndex" { return "Nhrpclientindex" }
    if yname == "nhrpClientInternetworkAddrType" { return "Nhrpclientinternetworkaddrtype" }
    if yname == "nhrpClientInternetworkAddr" { return "Nhrpclientinternetworkaddr" }
    if yname == "nhrpClientNbmaAddrType" { return "Nhrpclientnbmaaddrtype" }
    if yname == "nhrpClientNbmaAddr" { return "Nhrpclientnbmaaddr" }
    if yname == "nhrpClientNbmaSubaddr" { return "Nhrpclientnbmasubaddr" }
    if yname == "nhrpClientInitialRequestTimeout" { return "Nhrpclientinitialrequesttimeout" }
    if yname == "nhrpClientRegistrationRequestRetries" { return "Nhrpclientregistrationrequestretries" }
    if yname == "nhrpClientResolutionRequestRetries" { return "Nhrpclientresolutionrequestretries" }
    if yname == "nhrpClientPurgeRequestRetries" { return "Nhrpclientpurgerequestretries" }
    if yname == "nhrpClientDefaultMtu" { return "Nhrpclientdefaultmtu" }
    if yname == "nhrpClientHoldTime" { return "Nhrpclientholdtime" }
    if yname == "nhrpClientRequestID" { return "Nhrpclientrequestid" }
    if yname == "nhrpClientStorageType" { return "Nhrpclientstoragetype" }
    if yname == "nhrpClientRowStatus" { return "Nhrpclientrowstatus" }
    return ""
}

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetSegmentPath() string {
    return "nhrpClientEntry" + "[nhrpClientIndex='" + fmt.Sprintf("%v", nhrpcliententry.Nhrpclientindex) + "']"
}

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nhrpClientIndex"] = nhrpcliententry.Nhrpclientindex
    leafs["nhrpClientInternetworkAddrType"] = nhrpcliententry.Nhrpclientinternetworkaddrtype
    leafs["nhrpClientInternetworkAddr"] = nhrpcliententry.Nhrpclientinternetworkaddr
    leafs["nhrpClientNbmaAddrType"] = nhrpcliententry.Nhrpclientnbmaaddrtype
    leafs["nhrpClientNbmaAddr"] = nhrpcliententry.Nhrpclientnbmaaddr
    leafs["nhrpClientNbmaSubaddr"] = nhrpcliententry.Nhrpclientnbmasubaddr
    leafs["nhrpClientInitialRequestTimeout"] = nhrpcliententry.Nhrpclientinitialrequesttimeout
    leafs["nhrpClientRegistrationRequestRetries"] = nhrpcliententry.Nhrpclientregistrationrequestretries
    leafs["nhrpClientResolutionRequestRetries"] = nhrpcliententry.Nhrpclientresolutionrequestretries
    leafs["nhrpClientPurgeRequestRetries"] = nhrpcliententry.Nhrpclientpurgerequestretries
    leafs["nhrpClientDefaultMtu"] = nhrpcliententry.Nhrpclientdefaultmtu
    leafs["nhrpClientHoldTime"] = nhrpcliententry.Nhrpclientholdtime
    leafs["nhrpClientRequestID"] = nhrpcliententry.Nhrpclientrequestid
    leafs["nhrpClientStorageType"] = nhrpcliententry.Nhrpclientstoragetype
    leafs["nhrpClientRowStatus"] = nhrpcliententry.Nhrpclientrowstatus
    return leafs
}

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetYangName() string { return "nhrpClientEntry" }

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) SetParent(parent types.Entity) { nhrpcliententry.parent = parent }

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetParent() types.Entity { return nhrpcliententry.parent }

func (nhrpcliententry *NHRPMIB_Nhrpclienttable_Nhrpcliententry) GetParentYangName() string { return "nhrpClientTable" }

// NHRPMIB_Nhrpclientregistrationtable
// A table of Registration Request Information that
// needs to be maintained by the NHCs (clients).
type NHRPMIB_Nhrpclientregistrationtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An NHC needs to maintain registration request information between the NHC
    // and the NHS.  An entry in this table represents information for a single
    // registration request. The type is slice of
    // NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry.
    Nhrpclientregistrationentry []NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry
}

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetFilter() yfilter.YFilter { return nhrpclientregistrationtable.YFilter }

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) SetFilter(yf yfilter.YFilter) { nhrpclientregistrationtable.YFilter = yf }

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetGoName(yname string) string {
    if yname == "nhrpClientRegistrationEntry" { return "Nhrpclientregistrationentry" }
    return ""
}

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetSegmentPath() string {
    return "nhrpClientRegistrationTable"
}

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nhrpClientRegistrationEntry" {
        for _, c := range nhrpclientregistrationtable.Nhrpclientregistrationentry {
            if nhrpclientregistrationtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry{}
        nhrpclientregistrationtable.Nhrpclientregistrationentry = append(nhrpclientregistrationtable.Nhrpclientregistrationentry, child)
        return &nhrpclientregistrationtable.Nhrpclientregistrationentry[len(nhrpclientregistrationtable.Nhrpclientregistrationentry)-1]
    }
    return nil
}

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nhrpclientregistrationtable.Nhrpclientregistrationentry {
        children[nhrpclientregistrationtable.Nhrpclientregistrationentry[i].GetSegmentPath()] = &nhrpclientregistrationtable.Nhrpclientregistrationentry[i]
    }
    return children
}

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetYangName() string { return "nhrpClientRegistrationTable" }

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) SetParent(parent types.Entity) { nhrpclientregistrationtable.parent = parent }

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetParent() types.Entity { return nhrpclientregistrationtable.parent }

func (nhrpclientregistrationtable *NHRPMIB_Nhrpclientregistrationtable) GetParentYangName() string { return "NHRP-MIB" }

// NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry
// An NHC needs to maintain registration request information
// between the NHC and the NHS.  An entry in this table
// represents information for a single registration request.
type NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry struct {
    parent types.Entity
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

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetFilter() yfilter.YFilter { return nhrpclientregistrationentry.YFilter }

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) SetFilter(yf yfilter.YFilter) { nhrpclientregistrationentry.YFilter = yf }

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetGoName(yname string) string {
    if yname == "nhrpClientIndex" { return "Nhrpclientindex" }
    if yname == "nhrpClientRegIndex" { return "Nhrpclientregindex" }
    if yname == "nhrpClientRegUniqueness" { return "Nhrpclientreguniqueness" }
    if yname == "nhrpClientRegState" { return "Nhrpclientregstate" }
    if yname == "nhrpClientRegRowStatus" { return "Nhrpclientregrowstatus" }
    return ""
}

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetSegmentPath() string {
    return "nhrpClientRegistrationEntry" + "[nhrpClientIndex='" + fmt.Sprintf("%v", nhrpclientregistrationentry.Nhrpclientindex) + "']" + "[nhrpClientRegIndex='" + fmt.Sprintf("%v", nhrpclientregistrationentry.Nhrpclientregindex) + "']"
}

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nhrpClientIndex"] = nhrpclientregistrationentry.Nhrpclientindex
    leafs["nhrpClientRegIndex"] = nhrpclientregistrationentry.Nhrpclientregindex
    leafs["nhrpClientRegUniqueness"] = nhrpclientregistrationentry.Nhrpclientreguniqueness
    leafs["nhrpClientRegState"] = nhrpclientregistrationentry.Nhrpclientregstate
    leafs["nhrpClientRegRowStatus"] = nhrpclientregistrationentry.Nhrpclientregrowstatus
    return leafs
}

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetYangName() string { return "nhrpClientRegistrationEntry" }

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) SetParent(parent types.Entity) { nhrpclientregistrationentry.parent = parent }

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetParent() types.Entity { return nhrpclientregistrationentry.parent }

func (nhrpclientregistrationentry *NHRPMIB_Nhrpclientregistrationtable_Nhrpclientregistrationentry) GetParentYangName() string { return "nhrpClientRegistrationTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An NHS that may be used by an NHC. The type is slice of
    // NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry.
    Nhrpclientnhsentry []NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry
}

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetFilter() yfilter.YFilter { return nhrpclientnhstable.YFilter }

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) SetFilter(yf yfilter.YFilter) { nhrpclientnhstable.YFilter = yf }

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetGoName(yname string) string {
    if yname == "nhrpClientNhsEntry" { return "Nhrpclientnhsentry" }
    return ""
}

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetSegmentPath() string {
    return "nhrpClientNhsTable"
}

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nhrpClientNhsEntry" {
        for _, c := range nhrpclientnhstable.Nhrpclientnhsentry {
            if nhrpclientnhstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry{}
        nhrpclientnhstable.Nhrpclientnhsentry = append(nhrpclientnhstable.Nhrpclientnhsentry, child)
        return &nhrpclientnhstable.Nhrpclientnhsentry[len(nhrpclientnhstable.Nhrpclientnhsentry)-1]
    }
    return nil
}

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nhrpclientnhstable.Nhrpclientnhsentry {
        children[nhrpclientnhstable.Nhrpclientnhsentry[i].GetSegmentPath()] = &nhrpclientnhstable.Nhrpclientnhsentry[i]
    }
    return children
}

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetYangName() string { return "nhrpClientNhsTable" }

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) SetParent(parent types.Entity) { nhrpclientnhstable.parent = parent }

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetParent() types.Entity { return nhrpclientnhstable.parent }

func (nhrpclientnhstable *NHRPMIB_Nhrpclientnhstable) GetParentYangName() string { return "NHRP-MIB" }

// NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry
// An NHS that may be used by an NHC.
type NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry struct {
    parent types.Entity
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

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetFilter() yfilter.YFilter { return nhrpclientnhsentry.YFilter }

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) SetFilter(yf yfilter.YFilter) { nhrpclientnhsentry.YFilter = yf }

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetGoName(yname string) string {
    if yname == "nhrpClientIndex" { return "Nhrpclientindex" }
    if yname == "nhrpClientNhsIndex" { return "Nhrpclientnhsindex" }
    if yname == "nhrpClientNhsInternetworkAddrType" { return "Nhrpclientnhsinternetworkaddrtype" }
    if yname == "nhrpClientNhsInternetworkAddr" { return "Nhrpclientnhsinternetworkaddr" }
    if yname == "nhrpClientNhsNbmaAddrType" { return "Nhrpclientnhsnbmaaddrtype" }
    if yname == "nhrpClientNhsNbmaAddr" { return "Nhrpclientnhsnbmaaddr" }
    if yname == "nhrpClientNhsNbmaSubaddr" { return "Nhrpclientnhsnbmasubaddr" }
    if yname == "nhrpClientNhsInUse" { return "Nhrpclientnhsinuse" }
    if yname == "nhrpClientNhsRowStatus" { return "Nhrpclientnhsrowstatus" }
    return ""
}

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetSegmentPath() string {
    return "nhrpClientNhsEntry" + "[nhrpClientIndex='" + fmt.Sprintf("%v", nhrpclientnhsentry.Nhrpclientindex) + "']" + "[nhrpClientNhsIndex='" + fmt.Sprintf("%v", nhrpclientnhsentry.Nhrpclientnhsindex) + "']"
}

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nhrpClientIndex"] = nhrpclientnhsentry.Nhrpclientindex
    leafs["nhrpClientNhsIndex"] = nhrpclientnhsentry.Nhrpclientnhsindex
    leafs["nhrpClientNhsInternetworkAddrType"] = nhrpclientnhsentry.Nhrpclientnhsinternetworkaddrtype
    leafs["nhrpClientNhsInternetworkAddr"] = nhrpclientnhsentry.Nhrpclientnhsinternetworkaddr
    leafs["nhrpClientNhsNbmaAddrType"] = nhrpclientnhsentry.Nhrpclientnhsnbmaaddrtype
    leafs["nhrpClientNhsNbmaAddr"] = nhrpclientnhsentry.Nhrpclientnhsnbmaaddr
    leafs["nhrpClientNhsNbmaSubaddr"] = nhrpclientnhsentry.Nhrpclientnhsnbmasubaddr
    leafs["nhrpClientNhsInUse"] = nhrpclientnhsentry.Nhrpclientnhsinuse
    leafs["nhrpClientNhsRowStatus"] = nhrpclientnhsentry.Nhrpclientnhsrowstatus
    return leafs
}

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetYangName() string { return "nhrpClientNhsEntry" }

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) SetParent(parent types.Entity) { nhrpclientnhsentry.parent = parent }

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetParent() types.Entity { return nhrpclientnhsentry.parent }

func (nhrpclientnhsentry *NHRPMIB_Nhrpclientnhstable_Nhrpclientnhsentry) GetParentYangName() string { return "nhrpClientNhsTable" }

// NHRPMIB_Nhrpclientstattable
// This table contains statistics collected by NHRP
// clients.
type NHRPMIB_Nhrpclientstattable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics collected by a NHRP client. The type is slice of
    // NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry.
    Nhrpclientstatentry []NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry
}

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetFilter() yfilter.YFilter { return nhrpclientstattable.YFilter }

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) SetFilter(yf yfilter.YFilter) { nhrpclientstattable.YFilter = yf }

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetGoName(yname string) string {
    if yname == "nhrpClientStatEntry" { return "Nhrpclientstatentry" }
    return ""
}

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetSegmentPath() string {
    return "nhrpClientStatTable"
}

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nhrpClientStatEntry" {
        for _, c := range nhrpclientstattable.Nhrpclientstatentry {
            if nhrpclientstattable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry{}
        nhrpclientstattable.Nhrpclientstatentry = append(nhrpclientstattable.Nhrpclientstatentry, child)
        return &nhrpclientstattable.Nhrpclientstatentry[len(nhrpclientstattable.Nhrpclientstatentry)-1]
    }
    return nil
}

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nhrpclientstattable.Nhrpclientstatentry {
        children[nhrpclientstattable.Nhrpclientstatentry[i].GetSegmentPath()] = &nhrpclientstattable.Nhrpclientstatentry[i]
    }
    return children
}

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetYangName() string { return "nhrpClientStatTable" }

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) SetParent(parent types.Entity) { nhrpclientstattable.parent = parent }

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetParent() types.Entity { return nhrpclientstattable.parent }

func (nhrpclientstattable *NHRPMIB_Nhrpclientstattable) GetParentYangName() string { return "NHRP-MIB" }

// NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry
// Statistics collected by a NHRP client.
type NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry struct {
    parent types.Entity
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

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetFilter() yfilter.YFilter { return nhrpclientstatentry.YFilter }

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) SetFilter(yf yfilter.YFilter) { nhrpclientstatentry.YFilter = yf }

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetGoName(yname string) string {
    if yname == "nhrpClientIndex" { return "Nhrpclientindex" }
    if yname == "nhrpClientStatTxResolveReq" { return "Nhrpclientstattxresolvereq" }
    if yname == "nhrpClientStatRxResolveReplyAck" { return "Nhrpclientstatrxresolvereplyack" }
    if yname == "nhrpClientStatRxResolveReplyNakProhibited" { return "Nhrpclientstatrxresolvereplynakprohibited" }
    if yname == "nhrpClientStatRxResolveReplyNakInsufResources" { return "Nhrpclientstatrxresolvereplynakinsufresources" }
    if yname == "nhrpClientStatRxResolveReplyNakNoBinding" { return "Nhrpclientstatrxresolvereplynaknobinding" }
    if yname == "nhrpClientStatRxResolveReplyNakNotUnique" { return "Nhrpclientstatrxresolvereplynaknotunique" }
    if yname == "nhrpClientStatTxRegisterReq" { return "Nhrpclientstattxregisterreq" }
    if yname == "nhrpClientStatRxRegisterAck" { return "Nhrpclientstatrxregisterack" }
    if yname == "nhrpClientStatRxRegisterNakProhibited" { return "Nhrpclientstatrxregisternakprohibited" }
    if yname == "nhrpClientStatRxRegisterNakInsufResources" { return "Nhrpclientstatrxregisternakinsufresources" }
    if yname == "nhrpClientStatRxRegisterNakAlreadyReg" { return "Nhrpclientstatrxregisternakalreadyreg" }
    if yname == "nhrpClientStatRxPurgeReq" { return "Nhrpclientstatrxpurgereq" }
    if yname == "nhrpClientStatTxPurgeReq" { return "Nhrpclientstattxpurgereq" }
    if yname == "nhrpClientStatRxPurgeReply" { return "Nhrpclientstatrxpurgereply" }
    if yname == "nhrpClientStatTxPurgeReply" { return "Nhrpclientstattxpurgereply" }
    if yname == "nhrpClientStatTxErrorIndication" { return "Nhrpclientstattxerrorindication" }
    if yname == "nhrpClientStatRxErrUnrecognizedExtension" { return "Nhrpclientstatrxerrunrecognizedextension" }
    if yname == "nhrpClientStatRxErrLoopDetected" { return "Nhrpclientstatrxerrloopdetected" }
    if yname == "nhrpClientStatRxErrProtoAddrUnreachable" { return "Nhrpclientstatrxerrprotoaddrunreachable" }
    if yname == "nhrpClientStatRxErrProtoError" { return "Nhrpclientstatrxerrprotoerror" }
    if yname == "nhrpClientStatRxErrSduSizeExceeded" { return "Nhrpclientstatrxerrsdusizeexceeded" }
    if yname == "nhrpClientStatRxErrInvalidExtension" { return "Nhrpclientstatrxerrinvalidextension" }
    if yname == "nhrpClientStatRxErrAuthenticationFailure" { return "Nhrpclientstatrxerrauthenticationfailure" }
    if yname == "nhrpClientStatRxErrHopCountExceeded" { return "Nhrpclientstatrxerrhopcountexceeded" }
    if yname == "nhrpClientStatDiscontinuityTime" { return "Nhrpclientstatdiscontinuitytime" }
    return ""
}

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetSegmentPath() string {
    return "nhrpClientStatEntry" + "[nhrpClientIndex='" + fmt.Sprintf("%v", nhrpclientstatentry.Nhrpclientindex) + "']"
}

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nhrpClientIndex"] = nhrpclientstatentry.Nhrpclientindex
    leafs["nhrpClientStatTxResolveReq"] = nhrpclientstatentry.Nhrpclientstattxresolvereq
    leafs["nhrpClientStatRxResolveReplyAck"] = nhrpclientstatentry.Nhrpclientstatrxresolvereplyack
    leafs["nhrpClientStatRxResolveReplyNakProhibited"] = nhrpclientstatentry.Nhrpclientstatrxresolvereplynakprohibited
    leafs["nhrpClientStatRxResolveReplyNakInsufResources"] = nhrpclientstatentry.Nhrpclientstatrxresolvereplynakinsufresources
    leafs["nhrpClientStatRxResolveReplyNakNoBinding"] = nhrpclientstatentry.Nhrpclientstatrxresolvereplynaknobinding
    leafs["nhrpClientStatRxResolveReplyNakNotUnique"] = nhrpclientstatentry.Nhrpclientstatrxresolvereplynaknotunique
    leafs["nhrpClientStatTxRegisterReq"] = nhrpclientstatentry.Nhrpclientstattxregisterreq
    leafs["nhrpClientStatRxRegisterAck"] = nhrpclientstatentry.Nhrpclientstatrxregisterack
    leafs["nhrpClientStatRxRegisterNakProhibited"] = nhrpclientstatentry.Nhrpclientstatrxregisternakprohibited
    leafs["nhrpClientStatRxRegisterNakInsufResources"] = nhrpclientstatentry.Nhrpclientstatrxregisternakinsufresources
    leafs["nhrpClientStatRxRegisterNakAlreadyReg"] = nhrpclientstatentry.Nhrpclientstatrxregisternakalreadyreg
    leafs["nhrpClientStatRxPurgeReq"] = nhrpclientstatentry.Nhrpclientstatrxpurgereq
    leafs["nhrpClientStatTxPurgeReq"] = nhrpclientstatentry.Nhrpclientstattxpurgereq
    leafs["nhrpClientStatRxPurgeReply"] = nhrpclientstatentry.Nhrpclientstatrxpurgereply
    leafs["nhrpClientStatTxPurgeReply"] = nhrpclientstatentry.Nhrpclientstattxpurgereply
    leafs["nhrpClientStatTxErrorIndication"] = nhrpclientstatentry.Nhrpclientstattxerrorindication
    leafs["nhrpClientStatRxErrUnrecognizedExtension"] = nhrpclientstatentry.Nhrpclientstatrxerrunrecognizedextension
    leafs["nhrpClientStatRxErrLoopDetected"] = nhrpclientstatentry.Nhrpclientstatrxerrloopdetected
    leafs["nhrpClientStatRxErrProtoAddrUnreachable"] = nhrpclientstatentry.Nhrpclientstatrxerrprotoaddrunreachable
    leafs["nhrpClientStatRxErrProtoError"] = nhrpclientstatentry.Nhrpclientstatrxerrprotoerror
    leafs["nhrpClientStatRxErrSduSizeExceeded"] = nhrpclientstatentry.Nhrpclientstatrxerrsdusizeexceeded
    leafs["nhrpClientStatRxErrInvalidExtension"] = nhrpclientstatentry.Nhrpclientstatrxerrinvalidextension
    leafs["nhrpClientStatRxErrAuthenticationFailure"] = nhrpclientstatentry.Nhrpclientstatrxerrauthenticationfailure
    leafs["nhrpClientStatRxErrHopCountExceeded"] = nhrpclientstatentry.Nhrpclientstatrxerrhopcountexceeded
    leafs["nhrpClientStatDiscontinuityTime"] = nhrpclientstatentry.Nhrpclientstatdiscontinuitytime
    return leafs
}

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetYangName() string { return "nhrpClientStatEntry" }

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) SetParent(parent types.Entity) { nhrpclientstatentry.parent = parent }

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetParent() types.Entity { return nhrpclientstatentry.parent }

func (nhrpclientstatentry *NHRPMIB_Nhrpclientstattable_Nhrpclientstatentry) GetParentYangName() string { return "nhrpClientStatTable" }

// NHRPMIB_Nhrpservertable
// This table contains information for a set of NHSes
// associated with this agent.
type NHRPMIB_Nhrpservertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single NHS. The type is slice of
    // NHRPMIB_Nhrpservertable_Nhrpserverentry.
    Nhrpserverentry []NHRPMIB_Nhrpservertable_Nhrpserverentry
}

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetFilter() yfilter.YFilter { return nhrpservertable.YFilter }

func (nhrpservertable *NHRPMIB_Nhrpservertable) SetFilter(yf yfilter.YFilter) { nhrpservertable.YFilter = yf }

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetGoName(yname string) string {
    if yname == "nhrpServerEntry" { return "Nhrpserverentry" }
    return ""
}

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetSegmentPath() string {
    return "nhrpServerTable"
}

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nhrpServerEntry" {
        for _, c := range nhrpservertable.Nhrpserverentry {
            if nhrpservertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NHRPMIB_Nhrpservertable_Nhrpserverentry{}
        nhrpservertable.Nhrpserverentry = append(nhrpservertable.Nhrpserverentry, child)
        return &nhrpservertable.Nhrpserverentry[len(nhrpservertable.Nhrpserverentry)-1]
    }
    return nil
}

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nhrpservertable.Nhrpserverentry {
        children[nhrpservertable.Nhrpserverentry[i].GetSegmentPath()] = &nhrpservertable.Nhrpserverentry[i]
    }
    return children
}

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetYangName() string { return "nhrpServerTable" }

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpservertable *NHRPMIB_Nhrpservertable) SetParent(parent types.Entity) { nhrpservertable.parent = parent }

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetParent() types.Entity { return nhrpservertable.parent }

func (nhrpservertable *NHRPMIB_Nhrpservertable) GetParentYangName() string { return "NHRP-MIB" }

// NHRPMIB_Nhrpservertable_Nhrpserverentry
// Information about a single NHS.
type NHRPMIB_Nhrpservertable_Nhrpserverentry struct {
    parent types.Entity
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

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetFilter() yfilter.YFilter { return nhrpserverentry.YFilter }

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) SetFilter(yf yfilter.YFilter) { nhrpserverentry.YFilter = yf }

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetGoName(yname string) string {
    if yname == "nhrpServerIndex" { return "Nhrpserverindex" }
    if yname == "nhrpServerInternetworkAddrType" { return "Nhrpserverinternetworkaddrtype" }
    if yname == "nhrpServerInternetworkAddr" { return "Nhrpserverinternetworkaddr" }
    if yname == "nhrpServerNbmaAddrType" { return "Nhrpservernbmaaddrtype" }
    if yname == "nhrpServerNbmaAddr" { return "Nhrpservernbmaaddr" }
    if yname == "nhrpServerNbmaSubaddr" { return "Nhrpservernbmasubaddr" }
    if yname == "nhrpServerStorageType" { return "Nhrpserverstoragetype" }
    if yname == "nhrpServerRowStatus" { return "Nhrpserverrowstatus" }
    return ""
}

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetSegmentPath() string {
    return "nhrpServerEntry" + "[nhrpServerIndex='" + fmt.Sprintf("%v", nhrpserverentry.Nhrpserverindex) + "']"
}

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nhrpServerIndex"] = nhrpserverentry.Nhrpserverindex
    leafs["nhrpServerInternetworkAddrType"] = nhrpserverentry.Nhrpserverinternetworkaddrtype
    leafs["nhrpServerInternetworkAddr"] = nhrpserverentry.Nhrpserverinternetworkaddr
    leafs["nhrpServerNbmaAddrType"] = nhrpserverentry.Nhrpservernbmaaddrtype
    leafs["nhrpServerNbmaAddr"] = nhrpserverentry.Nhrpservernbmaaddr
    leafs["nhrpServerNbmaSubaddr"] = nhrpserverentry.Nhrpservernbmasubaddr
    leafs["nhrpServerStorageType"] = nhrpserverentry.Nhrpserverstoragetype
    leafs["nhrpServerRowStatus"] = nhrpserverentry.Nhrpserverrowstatus
    return leafs
}

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetYangName() string { return "nhrpServerEntry" }

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) SetParent(parent types.Entity) { nhrpserverentry.parent = parent }

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetParent() types.Entity { return nhrpserverentry.parent }

func (nhrpserverentry *NHRPMIB_Nhrpservertable_Nhrpserverentry) GetParentYangName() string { return "nhrpServerTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Additional information kept by a NHS for a relevant Next Hop Resolution
    // Cache entry. The type is slice of
    // NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry.
    Nhrpservercacheentry []NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry
}

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetFilter() yfilter.YFilter { return nhrpservercachetable.YFilter }

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) SetFilter(yf yfilter.YFilter) { nhrpservercachetable.YFilter = yf }

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetGoName(yname string) string {
    if yname == "nhrpServerCacheEntry" { return "Nhrpservercacheentry" }
    return ""
}

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetSegmentPath() string {
    return "nhrpServerCacheTable"
}

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nhrpServerCacheEntry" {
        for _, c := range nhrpservercachetable.Nhrpservercacheentry {
            if nhrpservercachetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry{}
        nhrpservercachetable.Nhrpservercacheentry = append(nhrpservercachetable.Nhrpservercacheentry, child)
        return &nhrpservercachetable.Nhrpservercacheentry[len(nhrpservercachetable.Nhrpservercacheentry)-1]
    }
    return nil
}

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nhrpservercachetable.Nhrpservercacheentry {
        children[nhrpservercachetable.Nhrpservercacheentry[i].GetSegmentPath()] = &nhrpservercachetable.Nhrpservercacheentry[i]
    }
    return children
}

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetYangName() string { return "nhrpServerCacheTable" }

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) SetParent(parent types.Entity) { nhrpservercachetable.parent = parent }

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetParent() types.Entity { return nhrpservercachetable.parent }

func (nhrpservercachetable *NHRPMIB_Nhrpservercachetable) GetParentYangName() string { return "NHRP-MIB" }

// NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry
// Additional information kept by a NHS for a relevant
// Next Hop Resolution Cache entry.
type NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry struct {
    parent types.Entity
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

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetFilter() yfilter.YFilter { return nhrpservercacheentry.YFilter }

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) SetFilter(yf yfilter.YFilter) { nhrpservercacheentry.YFilter = yf }

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetGoName(yname string) string {
    if yname == "nhrpCacheInternetworkAddrType" { return "Nhrpcacheinternetworkaddrtype" }
    if yname == "nhrpCacheInternetworkAddr" { return "Nhrpcacheinternetworkaddr" }
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "nhrpCacheIndex" { return "Nhrpcacheindex" }
    if yname == "nhrpServerCacheAuthoritative" { return "Nhrpservercacheauthoritative" }
    if yname == "nhrpServerCacheUniqueness" { return "Nhrpservercacheuniqueness" }
    return ""
}

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetSegmentPath() string {
    return "nhrpServerCacheEntry" + "[nhrpCacheInternetworkAddrType='" + fmt.Sprintf("%v", nhrpservercacheentry.Nhrpcacheinternetworkaddrtype) + "']" + "[nhrpCacheInternetworkAddr='" + fmt.Sprintf("%v", nhrpservercacheentry.Nhrpcacheinternetworkaddr) + "']" + "[ifIndex='" + fmt.Sprintf("%v", nhrpservercacheentry.Ifindex) + "']" + "[nhrpCacheIndex='" + fmt.Sprintf("%v", nhrpservercacheentry.Nhrpcacheindex) + "']"
}

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nhrpCacheInternetworkAddrType"] = nhrpservercacheentry.Nhrpcacheinternetworkaddrtype
    leafs["nhrpCacheInternetworkAddr"] = nhrpservercacheentry.Nhrpcacheinternetworkaddr
    leafs["ifIndex"] = nhrpservercacheentry.Ifindex
    leafs["nhrpCacheIndex"] = nhrpservercacheentry.Nhrpcacheindex
    leafs["nhrpServerCacheAuthoritative"] = nhrpservercacheentry.Nhrpservercacheauthoritative
    leafs["nhrpServerCacheUniqueness"] = nhrpservercacheentry.Nhrpservercacheuniqueness
    return leafs
}

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetYangName() string { return "nhrpServerCacheEntry" }

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) SetParent(parent types.Entity) { nhrpservercacheentry.parent = parent }

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetParent() types.Entity { return nhrpservercacheentry.parent }

func (nhrpservercacheentry *NHRPMIB_Nhrpservercachetable_Nhrpservercacheentry) GetParentYangName() string { return "nhrpServerCacheTable" }

// NHRPMIB_Nhrpservernhctable
// A table of NHCs that are available for use by this NHS
// (Server).
type NHRPMIB_Nhrpservernhctable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An NHC that may be used by an NHS. The type is slice of
    // NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry.
    Nhrpservernhcentry []NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry
}

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetFilter() yfilter.YFilter { return nhrpservernhctable.YFilter }

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) SetFilter(yf yfilter.YFilter) { nhrpservernhctable.YFilter = yf }

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetGoName(yname string) string {
    if yname == "nhrpServerNhcEntry" { return "Nhrpservernhcentry" }
    return ""
}

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetSegmentPath() string {
    return "nhrpServerNhcTable"
}

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nhrpServerNhcEntry" {
        for _, c := range nhrpservernhctable.Nhrpservernhcentry {
            if nhrpservernhctable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry{}
        nhrpservernhctable.Nhrpservernhcentry = append(nhrpservernhctable.Nhrpservernhcentry, child)
        return &nhrpservernhctable.Nhrpservernhcentry[len(nhrpservernhctable.Nhrpservernhcentry)-1]
    }
    return nil
}

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nhrpservernhctable.Nhrpservernhcentry {
        children[nhrpservernhctable.Nhrpservernhcentry[i].GetSegmentPath()] = &nhrpservernhctable.Nhrpservernhcentry[i]
    }
    return children
}

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetYangName() string { return "nhrpServerNhcTable" }

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) SetParent(parent types.Entity) { nhrpservernhctable.parent = parent }

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetParent() types.Entity { return nhrpservernhctable.parent }

func (nhrpservernhctable *NHRPMIB_Nhrpservernhctable) GetParentYangName() string { return "NHRP-MIB" }

// NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry
// An NHC that may be used by an NHS.
type NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry struct {
    parent types.Entity
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

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetFilter() yfilter.YFilter { return nhrpservernhcentry.YFilter }

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) SetFilter(yf yfilter.YFilter) { nhrpservernhcentry.YFilter = yf }

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetGoName(yname string) string {
    if yname == "nhrpServerIndex" { return "Nhrpserverindex" }
    if yname == "nhrpServerNhcIndex" { return "Nhrpservernhcindex" }
    if yname == "nhrpServerNhcPrefixLength" { return "Nhrpservernhcprefixlength" }
    if yname == "nhrpServerNhcInternetworkAddrType" { return "Nhrpservernhcinternetworkaddrtype" }
    if yname == "nhrpServerNhcInternetworkAddr" { return "Nhrpservernhcinternetworkaddr" }
    if yname == "nhrpServerNhcNbmaAddrType" { return "Nhrpservernhcnbmaaddrtype" }
    if yname == "nhrpServerNhcNbmaAddr" { return "Nhrpservernhcnbmaaddr" }
    if yname == "nhrpServerNhcNbmaSubaddr" { return "Nhrpservernhcnbmasubaddr" }
    if yname == "nhrpServerNhcInUse" { return "Nhrpservernhcinuse" }
    if yname == "nhrpServerNhcRowStatus" { return "Nhrpservernhcrowstatus" }
    return ""
}

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetSegmentPath() string {
    return "nhrpServerNhcEntry" + "[nhrpServerIndex='" + fmt.Sprintf("%v", nhrpservernhcentry.Nhrpserverindex) + "']" + "[nhrpServerNhcIndex='" + fmt.Sprintf("%v", nhrpservernhcentry.Nhrpservernhcindex) + "']"
}

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nhrpServerIndex"] = nhrpservernhcentry.Nhrpserverindex
    leafs["nhrpServerNhcIndex"] = nhrpservernhcentry.Nhrpservernhcindex
    leafs["nhrpServerNhcPrefixLength"] = nhrpservernhcentry.Nhrpservernhcprefixlength
    leafs["nhrpServerNhcInternetworkAddrType"] = nhrpservernhcentry.Nhrpservernhcinternetworkaddrtype
    leafs["nhrpServerNhcInternetworkAddr"] = nhrpservernhcentry.Nhrpservernhcinternetworkaddr
    leafs["nhrpServerNhcNbmaAddrType"] = nhrpservernhcentry.Nhrpservernhcnbmaaddrtype
    leafs["nhrpServerNhcNbmaAddr"] = nhrpservernhcentry.Nhrpservernhcnbmaaddr
    leafs["nhrpServerNhcNbmaSubaddr"] = nhrpservernhcentry.Nhrpservernhcnbmasubaddr
    leafs["nhrpServerNhcInUse"] = nhrpservernhcentry.Nhrpservernhcinuse
    leafs["nhrpServerNhcRowStatus"] = nhrpservernhcentry.Nhrpservernhcrowstatus
    return leafs
}

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetYangName() string { return "nhrpServerNhcEntry" }

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) SetParent(parent types.Entity) { nhrpservernhcentry.parent = parent }

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetParent() types.Entity { return nhrpservernhcentry.parent }

func (nhrpservernhcentry *NHRPMIB_Nhrpservernhctable_Nhrpservernhcentry) GetParentYangName() string { return "nhrpServerNhcTable" }

// NHRPMIB_Nhrpserverstattable
// Statistics collected by Next Hop Servers.
type NHRPMIB_Nhrpserverstattable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics for a particular NHS. The statistics are broken into received
    // (Rx), transmitted (Tx) and forwarded (Fw).  Forwarded (Fw) would be done by
    // a transit NHS. The type is slice of
    // NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry.
    Nhrpserverstatentry []NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry
}

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetFilter() yfilter.YFilter { return nhrpserverstattable.YFilter }

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) SetFilter(yf yfilter.YFilter) { nhrpserverstattable.YFilter = yf }

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetGoName(yname string) string {
    if yname == "nhrpServerStatEntry" { return "Nhrpserverstatentry" }
    return ""
}

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetSegmentPath() string {
    return "nhrpServerStatTable"
}

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nhrpServerStatEntry" {
        for _, c := range nhrpserverstattable.Nhrpserverstatentry {
            if nhrpserverstattable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry{}
        nhrpserverstattable.Nhrpserverstatentry = append(nhrpserverstattable.Nhrpserverstatentry, child)
        return &nhrpserverstattable.Nhrpserverstatentry[len(nhrpserverstattable.Nhrpserverstatentry)-1]
    }
    return nil
}

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nhrpserverstattable.Nhrpserverstatentry {
        children[nhrpserverstattable.Nhrpserverstatentry[i].GetSegmentPath()] = &nhrpserverstattable.Nhrpserverstatentry[i]
    }
    return children
}

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetYangName() string { return "nhrpServerStatTable" }

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) SetParent(parent types.Entity) { nhrpserverstattable.parent = parent }

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetParent() types.Entity { return nhrpserverstattable.parent }

func (nhrpserverstattable *NHRPMIB_Nhrpserverstattable) GetParentYangName() string { return "NHRP-MIB" }

// NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry
// Statistics for a particular NHS. The statistics are
// broken into received (Rx), transmitted (Tx)
// and forwarded (Fw).  Forwarded (Fw) would be done
// by a transit NHS.
type NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry struct {
    parent types.Entity
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

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetFilter() yfilter.YFilter { return nhrpserverstatentry.YFilter }

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) SetFilter(yf yfilter.YFilter) { nhrpserverstatentry.YFilter = yf }

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetGoName(yname string) string {
    if yname == "nhrpServerIndex" { return "Nhrpserverindex" }
    if yname == "nhrpServerStatRxResolveReq" { return "Nhrpserverstatrxresolvereq" }
    if yname == "nhrpServerStatTxResolveReplyAck" { return "Nhrpserverstattxresolvereplyack" }
    if yname == "nhrpServerStatTxResolveReplyNakProhibited" { return "Nhrpserverstattxresolvereplynakprohibited" }
    if yname == "nhrpServerStatTxResolveReplyNakInsufResources" { return "Nhrpserverstattxresolvereplynakinsufresources" }
    if yname == "nhrpServerStatTxResolveReplyNakNoBinding" { return "Nhrpserverstattxresolvereplynaknobinding" }
    if yname == "nhrpServerStatTxResolveReplyNakNotUnique" { return "Nhrpserverstattxresolvereplynaknotunique" }
    if yname == "nhrpServerStatRxRegisterReq" { return "Nhrpserverstatrxregisterreq" }
    if yname == "nhrpServerStatTxRegisterAck" { return "Nhrpserverstattxregisterack" }
    if yname == "nhrpServerStatTxRegisterNakProhibited" { return "Nhrpserverstattxregisternakprohibited" }
    if yname == "nhrpServerStatTxRegisterNakInsufResources" { return "Nhrpserverstattxregisternakinsufresources" }
    if yname == "nhrpServerStatTxRegisterNakAlreadyReg" { return "Nhrpserverstattxregisternakalreadyreg" }
    if yname == "nhrpServerStatRxPurgeReq" { return "Nhrpserverstatrxpurgereq" }
    if yname == "nhrpServerStatTxPurgeReq" { return "Nhrpserverstattxpurgereq" }
    if yname == "nhrpServerStatRxPurgeReply" { return "Nhrpserverstatrxpurgereply" }
    if yname == "nhrpServerStatTxPurgeReply" { return "Nhrpserverstattxpurgereply" }
    if yname == "nhrpServerStatRxErrUnrecognizedExtension" { return "Nhrpserverstatrxerrunrecognizedextension" }
    if yname == "nhrpServerStatRxErrLoopDetected" { return "Nhrpserverstatrxerrloopdetected" }
    if yname == "nhrpServerStatRxErrProtoAddrUnreachable" { return "Nhrpserverstatrxerrprotoaddrunreachable" }
    if yname == "nhrpServerStatRxErrProtoError" { return "Nhrpserverstatrxerrprotoerror" }
    if yname == "nhrpServerStatRxErrSduSizeExceeded" { return "Nhrpserverstatrxerrsdusizeexceeded" }
    if yname == "nhrpServerStatRxErrInvalidExtension" { return "Nhrpserverstatrxerrinvalidextension" }
    if yname == "nhrpServerStatRxErrInvalidResReplyReceived" { return "Nhrpserverstatrxerrinvalidresreplyreceived" }
    if yname == "nhrpServerStatRxErrAuthenticationFailure" { return "Nhrpserverstatrxerrauthenticationfailure" }
    if yname == "nhrpServerStatRxErrHopCountExceeded" { return "Nhrpserverstatrxerrhopcountexceeded" }
    if yname == "nhrpServerStatTxErrUnrecognizedExtension" { return "Nhrpserverstattxerrunrecognizedextension" }
    if yname == "nhrpServerStatTxErrLoopDetected" { return "Nhrpserverstattxerrloopdetected" }
    if yname == "nhrpServerStatTxErrProtoAddrUnreachable" { return "Nhrpserverstattxerrprotoaddrunreachable" }
    if yname == "nhrpServerStatTxErrProtoError" { return "Nhrpserverstattxerrprotoerror" }
    if yname == "nhrpServerStatTxErrSduSizeExceeded" { return "Nhrpserverstattxerrsdusizeexceeded" }
    if yname == "nhrpServerStatTxErrInvalidExtension" { return "Nhrpserverstattxerrinvalidextension" }
    if yname == "nhrpServerStatTxErrAuthenticationFailure" { return "Nhrpserverstattxerrauthenticationfailure" }
    if yname == "nhrpServerStatTxErrHopCountExceeded" { return "Nhrpserverstattxerrhopcountexceeded" }
    if yname == "nhrpServerStatFwResolveReq" { return "Nhrpserverstatfwresolvereq" }
    if yname == "nhrpServerStatFwResolveReply" { return "Nhrpserverstatfwresolvereply" }
    if yname == "nhrpServerStatFwRegisterReq" { return "Nhrpserverstatfwregisterreq" }
    if yname == "nhrpServerStatFwRegisterReply" { return "Nhrpserverstatfwregisterreply" }
    if yname == "nhrpServerStatFwPurgeReq" { return "Nhrpserverstatfwpurgereq" }
    if yname == "nhrpServerStatFwPurgeReply" { return "Nhrpserverstatfwpurgereply" }
    if yname == "nhrpServerStatFwErrorIndication" { return "Nhrpserverstatfwerrorindication" }
    if yname == "nhrpServerStatDiscontinuityTime" { return "Nhrpserverstatdiscontinuitytime" }
    return ""
}

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetSegmentPath() string {
    return "nhrpServerStatEntry" + "[nhrpServerIndex='" + fmt.Sprintf("%v", nhrpserverstatentry.Nhrpserverindex) + "']"
}

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nhrpServerIndex"] = nhrpserverstatentry.Nhrpserverindex
    leafs["nhrpServerStatRxResolveReq"] = nhrpserverstatentry.Nhrpserverstatrxresolvereq
    leafs["nhrpServerStatTxResolveReplyAck"] = nhrpserverstatentry.Nhrpserverstattxresolvereplyack
    leafs["nhrpServerStatTxResolveReplyNakProhibited"] = nhrpserverstatentry.Nhrpserverstattxresolvereplynakprohibited
    leafs["nhrpServerStatTxResolveReplyNakInsufResources"] = nhrpserverstatentry.Nhrpserverstattxresolvereplynakinsufresources
    leafs["nhrpServerStatTxResolveReplyNakNoBinding"] = nhrpserverstatentry.Nhrpserverstattxresolvereplynaknobinding
    leafs["nhrpServerStatTxResolveReplyNakNotUnique"] = nhrpserverstatentry.Nhrpserverstattxresolvereplynaknotunique
    leafs["nhrpServerStatRxRegisterReq"] = nhrpserverstatentry.Nhrpserverstatrxregisterreq
    leafs["nhrpServerStatTxRegisterAck"] = nhrpserverstatentry.Nhrpserverstattxregisterack
    leafs["nhrpServerStatTxRegisterNakProhibited"] = nhrpserverstatentry.Nhrpserverstattxregisternakprohibited
    leafs["nhrpServerStatTxRegisterNakInsufResources"] = nhrpserverstatentry.Nhrpserverstattxregisternakinsufresources
    leafs["nhrpServerStatTxRegisterNakAlreadyReg"] = nhrpserverstatentry.Nhrpserverstattxregisternakalreadyreg
    leafs["nhrpServerStatRxPurgeReq"] = nhrpserverstatentry.Nhrpserverstatrxpurgereq
    leafs["nhrpServerStatTxPurgeReq"] = nhrpserverstatentry.Nhrpserverstattxpurgereq
    leafs["nhrpServerStatRxPurgeReply"] = nhrpserverstatentry.Nhrpserverstatrxpurgereply
    leafs["nhrpServerStatTxPurgeReply"] = nhrpserverstatentry.Nhrpserverstattxpurgereply
    leafs["nhrpServerStatRxErrUnrecognizedExtension"] = nhrpserverstatentry.Nhrpserverstatrxerrunrecognizedextension
    leafs["nhrpServerStatRxErrLoopDetected"] = nhrpserverstatentry.Nhrpserverstatrxerrloopdetected
    leafs["nhrpServerStatRxErrProtoAddrUnreachable"] = nhrpserverstatentry.Nhrpserverstatrxerrprotoaddrunreachable
    leafs["nhrpServerStatRxErrProtoError"] = nhrpserverstatentry.Nhrpserverstatrxerrprotoerror
    leafs["nhrpServerStatRxErrSduSizeExceeded"] = nhrpserverstatentry.Nhrpserverstatrxerrsdusizeexceeded
    leafs["nhrpServerStatRxErrInvalidExtension"] = nhrpserverstatentry.Nhrpserverstatrxerrinvalidextension
    leafs["nhrpServerStatRxErrInvalidResReplyReceived"] = nhrpserverstatentry.Nhrpserverstatrxerrinvalidresreplyreceived
    leafs["nhrpServerStatRxErrAuthenticationFailure"] = nhrpserverstatentry.Nhrpserverstatrxerrauthenticationfailure
    leafs["nhrpServerStatRxErrHopCountExceeded"] = nhrpserverstatentry.Nhrpserverstatrxerrhopcountexceeded
    leafs["nhrpServerStatTxErrUnrecognizedExtension"] = nhrpserverstatentry.Nhrpserverstattxerrunrecognizedextension
    leafs["nhrpServerStatTxErrLoopDetected"] = nhrpserverstatentry.Nhrpserverstattxerrloopdetected
    leafs["nhrpServerStatTxErrProtoAddrUnreachable"] = nhrpserverstatentry.Nhrpserverstattxerrprotoaddrunreachable
    leafs["nhrpServerStatTxErrProtoError"] = nhrpserverstatentry.Nhrpserverstattxerrprotoerror
    leafs["nhrpServerStatTxErrSduSizeExceeded"] = nhrpserverstatentry.Nhrpserverstattxerrsdusizeexceeded
    leafs["nhrpServerStatTxErrInvalidExtension"] = nhrpserverstatentry.Nhrpserverstattxerrinvalidextension
    leafs["nhrpServerStatTxErrAuthenticationFailure"] = nhrpserverstatentry.Nhrpserverstattxerrauthenticationfailure
    leafs["nhrpServerStatTxErrHopCountExceeded"] = nhrpserverstatentry.Nhrpserverstattxerrhopcountexceeded
    leafs["nhrpServerStatFwResolveReq"] = nhrpserverstatentry.Nhrpserverstatfwresolvereq
    leafs["nhrpServerStatFwResolveReply"] = nhrpserverstatentry.Nhrpserverstatfwresolvereply
    leafs["nhrpServerStatFwRegisterReq"] = nhrpserverstatentry.Nhrpserverstatfwregisterreq
    leafs["nhrpServerStatFwRegisterReply"] = nhrpserverstatentry.Nhrpserverstatfwregisterreply
    leafs["nhrpServerStatFwPurgeReq"] = nhrpserverstatentry.Nhrpserverstatfwpurgereq
    leafs["nhrpServerStatFwPurgeReply"] = nhrpserverstatentry.Nhrpserverstatfwpurgereply
    leafs["nhrpServerStatFwErrorIndication"] = nhrpserverstatentry.Nhrpserverstatfwerrorindication
    leafs["nhrpServerStatDiscontinuityTime"] = nhrpserverstatentry.Nhrpserverstatdiscontinuitytime
    return leafs
}

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetBundleName() string { return "cisco_ios_xe" }

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetYangName() string { return "nhrpServerStatEntry" }

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) SetParent(parent types.Entity) { nhrpserverstatentry.parent = parent }

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetParent() types.Entity { return nhrpserverstatentry.parent }

func (nhrpserverstatentry *NHRPMIB_Nhrpserverstattable_Nhrpserverstatentry) GetParentYangName() string { return "nhrpServerStatTable" }

