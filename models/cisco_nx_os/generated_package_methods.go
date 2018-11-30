package cisco_nx_os


import (
    "runtime"
    "path"
)

func GetCapabilities() map[string]string {
    caps := make(map[string]string)
    caps["Cisco-NX-OS-device"] = "2018-07-17"
    return caps
}

func GetNamespaces() map[string]string {
    namespaces := make(map[string]string)
    namespaces["Cisco-NX-OS-device"] = "http://cisco.com/ns/yang/cisco-nx-os-device"
    return namespaces
}

func GetModelsPath() string {
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        panic("No caller information")
    }
    return path.Join(path.Dir(filename), "_yang")
}

