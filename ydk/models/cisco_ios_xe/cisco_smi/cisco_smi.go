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

type Ciscoproducts struct {
}

func (id Ciscoproducts) String() string {
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

type Otherenterprises struct {
}

func (id Otherenterprises) String() string {
	return "CISCO-SMI:otherEnterprises"
}

type Ciscosb struct {
}

func (id Ciscosb) String() string {
	return "CISCO-SMI:ciscoSB"
}

type Ciscosmb struct {
}

func (id Ciscosmb) String() string {
	return "CISCO-SMI:ciscoSMB"
}

type Ciscoagentcapability struct {
}

func (id Ciscoagentcapability) String() string {
	return "CISCO-SMI:ciscoAgentCapability"
}

type Ciscoconfig struct {
}

func (id Ciscoconfig) String() string {
	return "CISCO-SMI:ciscoConfig"
}

type Ciscomgmt struct {
}

func (id Ciscomgmt) String() string {
	return "CISCO-SMI:ciscoMgmt"
}

type Ciscoexperiment struct {
}

func (id Ciscoexperiment) String() string {
	return "CISCO-SMI:ciscoExperiment"
}

type Ciscoadmin struct {
}

func (id Ciscoadmin) String() string {
	return "CISCO-SMI:ciscoAdmin"
}

type Ciscoproxy struct {
}

func (id Ciscoproxy) String() string {
	return "CISCO-SMI:ciscoProxy"
}

type Ciscorptrgroupobjectid struct {
}

func (id Ciscorptrgroupobjectid) String() string {
	return "CISCO-SMI:ciscoRptrGroupObjectID"
}

type Ciscounknownrptrgroup struct {
}

func (id Ciscounknownrptrgroup) String() string {
	return "CISCO-SMI:ciscoUnknownRptrGroup"
}

type Cisco2505Rptrgroup struct {
}

func (id Cisco2505Rptrgroup) String() string {
	return "CISCO-SMI:cisco2505RptrGroup"
}

type Cisco2507Rptrgroup struct {
}

func (id Cisco2507Rptrgroup) String() string {
	return "CISCO-SMI:cisco2507RptrGroup"
}

type Cisco2516Rptrgroup struct {
}

func (id Cisco2516Rptrgroup) String() string {
	return "CISCO-SMI:cisco2516RptrGroup"
}

type Ciscowsx5020Rptrgroup struct {
}

func (id Ciscowsx5020Rptrgroup) String() string {
	return "CISCO-SMI:ciscoWsx5020RptrGroup"
}

type Ciscochipsets struct {
}

func (id Ciscochipsets) String() string {
	return "CISCO-SMI:ciscoChipSets"
}

type Ciscochipsetsaint1 struct {
}

func (id Ciscochipsetsaint1) String() string {
	return "CISCO-SMI:ciscoChipSetSaint1"
}

type Ciscochipsetsaint2 struct {
}

func (id Ciscochipsetsaint2) String() string {
	return "CISCO-SMI:ciscoChipSetSaint2"
}

type Ciscochipsetsaint3 struct {
}

func (id Ciscochipsetsaint3) String() string {
	return "CISCO-SMI:ciscoChipSetSaint3"
}

type Ciscochipsetsaint4 struct {
}

func (id Ciscochipsetsaint4) String() string {
	return "CISCO-SMI:ciscoChipSetSaint4"
}

type Ciscomodules struct {
}

func (id Ciscomodules) String() string {
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

type Ciscopartnerproducts struct {
}

func (id Ciscopartnerproducts) String() string {
	return "CISCO-SMI:ciscoPartnerProducts"
}

type Ciscopolicy struct {
}

func (id Ciscopolicy) String() string {
	return "CISCO-SMI:ciscoPolicy"
}

type Ciscopib struct {
}

func (id Ciscopib) String() string {
	return "CISCO-SMI:ciscoPIB"
}

type Ciscopolicyauto struct {
}

func (id Ciscopolicyauto) String() string {
	return "CISCO-SMI:ciscoPolicyAuto"
}

type Ciscopibtomib struct {
}

func (id Ciscopibtomib) String() string {
	return "CISCO-SMI:ciscoPibToMib"
}

type Ciscodomains struct {
}

func (id Ciscodomains) String() string {
	return "CISCO-SMI:ciscoDomains"
}

type Ciscotdomainudpipv4 struct {
}

func (id Ciscotdomainudpipv4) String() string {
	return "CISCO-SMI:ciscoTDomainUdpIpv4"
}

type Ciscotdomainudpipv6 struct {
}

func (id Ciscotdomainudpipv6) String() string {
	return "CISCO-SMI:ciscoTDomainUdpIpv6"
}

type Ciscotdomaintcpipv4 struct {
}

func (id Ciscotdomaintcpipv4) String() string {
	return "CISCO-SMI:ciscoTDomainTcpIpv4"
}

type Ciscotdomaintcpipv6 struct {
}

func (id Ciscotdomaintcpipv6) String() string {
	return "CISCO-SMI:ciscoTDomainTcpIpv6"
}

type Ciscotdomainlocal struct {
}

func (id Ciscotdomainlocal) String() string {
	return "CISCO-SMI:ciscoTDomainLocal"
}

type Ciscotdomainclns struct {
}

func (id Ciscotdomainclns) String() string {
	return "CISCO-SMI:ciscoTDomainClns"
}

type Ciscotdomaincons struct {
}

func (id Ciscotdomaincons) String() string {
	return "CISCO-SMI:ciscoTDomainCons"
}

type Ciscotdomainddp struct {
}

func (id Ciscotdomainddp) String() string {
	return "CISCO-SMI:ciscoTDomainDdp"
}

type Ciscotdomainipx struct {
}

func (id Ciscotdomainipx) String() string {
	return "CISCO-SMI:ciscoTDomainIpx"
}

type Ciscotdomainsctpipv4 struct {
}

func (id Ciscotdomainsctpipv4) String() string {
	return "CISCO-SMI:ciscoTDomainSctpIpv4"
}

type Ciscotdomainsctpipv6 struct {
}

func (id Ciscotdomainsctpipv6) String() string {
	return "CISCO-SMI:ciscoTDomainSctpIpv6"
}

type Ciscocib struct {
}

func (id Ciscocib) String() string {
	return "CISCO-SMI:ciscoCIB"
}

type Ciscocibmmigroup struct {
}

func (id Ciscocibmmigroup) String() string {
	return "CISCO-SMI:ciscoCibMmiGroup"
}

type Ciscocibprovgroup struct {
}

func (id Ciscocibprovgroup) String() string {
	return "CISCO-SMI:ciscoCibProvGroup"
}

type Ciscopki struct {
}

func (id Ciscopki) String() string {
	return "CISCO-SMI:ciscoPKI"
}

