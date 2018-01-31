// This module augments ietf-netconf-monitoring with extra
// monitoring data.
package tailf_netconf_monitoring

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tailf_netconf_monitoring"))
}

type RestHttps struct {
}

func (id RestHttps) String() string {
	return "tailf-netconf-monitoring:rest-https"
}

type CliSsh struct {
}

func (id CliSsh) String() string {
	return "tailf-netconf-monitoring:cli-ssh"
}

type CliConsole struct {
}

func (id CliConsole) String() string {
	return "tailf-netconf-monitoring:cli-console"
}

type WebuiHttps struct {
}

func (id WebuiHttps) String() string {
	return "tailf-netconf-monitoring:webui-https"
}

type SnmpUdp struct {
}

func (id SnmpUdp) String() string {
	return "tailf-netconf-monitoring:snmp-udp"
}

type WebuiHttp struct {
}

func (id WebuiHttp) String() string {
	return "tailf-netconf-monitoring:webui-http"
}

type RestHttp struct {
}

func (id RestHttp) String() string {
	return "tailf-netconf-monitoring:rest-http"
}

type NetconfTcp struct {
}

func (id NetconfTcp) String() string {
	return "tailf-netconf-monitoring:netconf-tcp"
}

type CliTcp struct {
}

func (id CliTcp) String() string {
	return "tailf-netconf-monitoring:cli-tcp"
}

