// The Structure of Management Information for the
// Cisco enterprise.
package cisco_smi

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_smi"))
}

type CiscoProducts struct {
}

func (id CiscoProducts) String() string {
	return "CISCO-SMI:ciscoProducts"
}

type Local struct {
}

func (id Local) String() string {
	return "CISCO-SMI:local"
}

type Temporary struct {
}

func (id Temporary) String() string {
	return "CISCO-SMI:temporary"
}

type Pakmon struct {
}

func (id Pakmon) String() string {
	return "CISCO-SMI:pakmon"
}

type Workgroup struct {
}

func (id Workgroup) String() string {
	return "CISCO-SMI:workgroup"
}

type OtherEnterprises struct {
}

func (id OtherEnterprises) String() string {
	return "CISCO-SMI:otherEnterprises"
}

type CiscoSB struct {
}

func (id CiscoSB) String() string {
	return "CISCO-SMI:ciscoSB"
}

type CiscoSMB struct {
}

func (id CiscoSMB) String() string {
	return "CISCO-SMI:ciscoSMB"
}

type CiscoAgentCapability struct {
}

func (id CiscoAgentCapability) String() string {
	return "CISCO-SMI:ciscoAgentCapability"
}

type CiscoConfig struct {
}

func (id CiscoConfig) String() string {
	return "CISCO-SMI:ciscoConfig"
}

type CiscoMgmt struct {
}

func (id CiscoMgmt) String() string {
	return "CISCO-SMI:ciscoMgmt"
}

type CiscoExperiment struct {
}

func (id CiscoExperiment) String() string {
	return "CISCO-SMI:ciscoExperiment"
}

type CiscoAdmin struct {
}

func (id CiscoAdmin) String() string {
	return "CISCO-SMI:ciscoAdmin"
}

type CiscoProxy struct {
}

func (id CiscoProxy) String() string {
	return "CISCO-SMI:ciscoProxy"
}

type CiscoRptrGroupObjectID struct {
}

func (id CiscoRptrGroupObjectID) String() string {
	return "CISCO-SMI:ciscoRptrGroupObjectID"
}

type CiscoUnknownRptrGroup struct {
}

func (id CiscoUnknownRptrGroup) String() string {
	return "CISCO-SMI:ciscoUnknownRptrGroup"
}

type Cisco2505RptrGroup struct {
}

func (id Cisco2505RptrGroup) String() string {
	return "CISCO-SMI:cisco2505RptrGroup"
}

type Cisco2507RptrGroup struct {
}

func (id Cisco2507RptrGroup) String() string {
	return "CISCO-SMI:cisco2507RptrGroup"
}

type Cisco2516RptrGroup struct {
}

func (id Cisco2516RptrGroup) String() string {
	return "CISCO-SMI:cisco2516RptrGroup"
}

type CiscoWsx5020RptrGroup struct {
}

func (id CiscoWsx5020RptrGroup) String() string {
	return "CISCO-SMI:ciscoWsx5020RptrGroup"
}

type CiscoChipSets struct {
}

func (id CiscoChipSets) String() string {
	return "CISCO-SMI:ciscoChipSets"
}

type CiscoChipSetSaint1 struct {
}

func (id CiscoChipSetSaint1) String() string {
	return "CISCO-SMI:ciscoChipSetSaint1"
}

type CiscoChipSetSaint2 struct {
}

func (id CiscoChipSetSaint2) String() string {
	return "CISCO-SMI:ciscoChipSetSaint2"
}

type CiscoChipSetSaint3 struct {
}

func (id CiscoChipSetSaint3) String() string {
	return "CISCO-SMI:ciscoChipSetSaint3"
}

type CiscoChipSetSaint4 struct {
}

func (id CiscoChipSetSaint4) String() string {
	return "CISCO-SMI:ciscoChipSetSaint4"
}

type CiscoModules struct {
}

func (id CiscoModules) String() string {
	return "CISCO-SMI:ciscoModules"
}

type Lightstream struct {
}

func (id Lightstream) String() string {
	return "CISCO-SMI:lightstream"
}

type Ciscoworks struct {
}

func (id Ciscoworks) String() string {
	return "CISCO-SMI:ciscoworks"
}

type Newport struct {
}

func (id Newport) String() string {
	return "CISCO-SMI:newport"
}

type CiscoPartnerProducts struct {
}

func (id CiscoPartnerProducts) String() string {
	return "CISCO-SMI:ciscoPartnerProducts"
}

type CiscoPolicy struct {
}

func (id CiscoPolicy) String() string {
	return "CISCO-SMI:ciscoPolicy"
}

type CiscoPIB struct {
}

func (id CiscoPIB) String() string {
	return "CISCO-SMI:ciscoPIB"
}

type CiscoPolicyAuto struct {
}

func (id CiscoPolicyAuto) String() string {
	return "CISCO-SMI:ciscoPolicyAuto"
}

type CiscoPibToMib struct {
}

func (id CiscoPibToMib) String() string {
	return "CISCO-SMI:ciscoPibToMib"
}

type CiscoDomains struct {
}

func (id CiscoDomains) String() string {
	return "CISCO-SMI:ciscoDomains"
}

type CiscoTDomainUdpIpv4 struct {
}

func (id CiscoTDomainUdpIpv4) String() string {
	return "CISCO-SMI:ciscoTDomainUdpIpv4"
}

type CiscoTDomainUdpIpv6 struct {
}

func (id CiscoTDomainUdpIpv6) String() string {
	return "CISCO-SMI:ciscoTDomainUdpIpv6"
}

type CiscoTDomainTcpIpv4 struct {
}

func (id CiscoTDomainTcpIpv4) String() string {
	return "CISCO-SMI:ciscoTDomainTcpIpv4"
}

type CiscoTDomainTcpIpv6 struct {
}

func (id CiscoTDomainTcpIpv6) String() string {
	return "CISCO-SMI:ciscoTDomainTcpIpv6"
}

type CiscoTDomainLocal struct {
}

func (id CiscoTDomainLocal) String() string {
	return "CISCO-SMI:ciscoTDomainLocal"
}

type CiscoTDomainClns struct {
}

func (id CiscoTDomainClns) String() string {
	return "CISCO-SMI:ciscoTDomainClns"
}

type CiscoTDomainCons struct {
}

func (id CiscoTDomainCons) String() string {
	return "CISCO-SMI:ciscoTDomainCons"
}

type CiscoTDomainDdp struct {
}

func (id CiscoTDomainDdp) String() string {
	return "CISCO-SMI:ciscoTDomainDdp"
}

type CiscoTDomainIpx struct {
}

func (id CiscoTDomainIpx) String() string {
	return "CISCO-SMI:ciscoTDomainIpx"
}

type CiscoTDomainSctpIpv4 struct {
}

func (id CiscoTDomainSctpIpv4) String() string {
	return "CISCO-SMI:ciscoTDomainSctpIpv4"
}

type CiscoTDomainSctpIpv6 struct {
}

func (id CiscoTDomainSctpIpv6) String() string {
	return "CISCO-SMI:ciscoTDomainSctpIpv6"
}

type CiscoCIB struct {
}

func (id CiscoCIB) String() string {
	return "CISCO-SMI:ciscoCIB"
}

type CiscoCibMmiGroup struct {
}

func (id CiscoCibMmiGroup) String() string {
	return "CISCO-SMI:ciscoCibMmiGroup"
}

type CiscoCibProvGroup struct {
}

func (id CiscoCibProvGroup) String() string {
	return "CISCO-SMI:ciscoCibProvGroup"
}

type CiscoPKI struct {
}

func (id CiscoPKI) String() string {
	return "CISCO-SMI:ciscoPKI"
}

