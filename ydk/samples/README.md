## Running Go Samples

Please set `$GOPATH` and install `ydk` using the below command:

```
go get github.com/CiscoDevNet/ydk-go/ydk
```

Please refer to the [YDK API documentation](http://ydk.cisco.com/go/docs) for an explanation of the different services.

To run these samples, execute the below.

For example, to run samples on a device supporting netconf and the `cisco-ios-xr` models, see the below.

To run a CRUD service sample, execute:

```
go run crud/cisco_ios_xr/shellutil_oper/nc_read_20/nc_read_20.go -device ssh://<username>:<password>@<host>:<port> [-v]
```

To run a Netconf service sample, execute:
```
go run netconf/cisco_ios_xr/infra_infra_locale_cfg/nc_edit_26/nc_edit_26.go -device ssh://<username>:<password>@<host>:<port> [-v]
```

To run an Executor service sample, execute:
```
go run executor/cisco_ios_xr/syslog_act/nc_execute_20/nc_execute_20.go -device ssh://<username>:<password>@<host>:<port> [-v]
```

To run a Codec service sample, execute:

```
go run codec/cisco_ios_xr/clns_isis_cfg/cd_encode_20/cd_encode_20.go [-v]
```
