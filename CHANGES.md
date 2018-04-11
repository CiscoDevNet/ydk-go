### 2018-04-09 version 0.7.1

#### Python, C++ & Go

##### Bundle improvements
  * Updated [`cisco-ios-xr`](https://github.com/CiscoDevNet/ydk-gen/blob/master/profiles/bundles/cisco-ios-xr_6_3_2.json) to support Cisco IOS XR 6.3.2 release
  * Updated [`cisco-ios-xe`](https://github.com/CiscoDevNet/ydk-gen/blob/master/profiles/bundles/cisco-ios-xe_16_8_1.json) to support Cisco IOS XE 16.8.1 release
  * Also updated [`openconfig`](https://github.com/CiscoDevNet/ydk-gen/blob/master/profiles/bundles/openconfig_0_1_5.json) and [`ietf`](https://github.com/CiscoDevNet/ydk-gen/blob/master/profiles/bundles/ietf_0_1_5.json) bundles
##### CRUD / Netconf / Codec / Path API improvements
  * Introduced support for multiple entities in Python and C++ ([#713](https://github.com/CiscoDevNet/ydk-gen/issues/713), [#719](https://github.com/CiscoDevNet/ydk-gen/issues/719), [#726](https://github.com/CiscoDevNet/ydk-gen/issues/726), [#736](https://github.com/CiscoDevNet/ydk-gen/issues/736))
  * Added support for yang models with more than 256 typedefs ([#678](https://github.com/CiscoDevNet/ydk-gen/issues/678), [#607](https://github.com/CiscoDevNet/ydk-gen/issues/607))
  * Fixed segfault with a `cisco-ios-xe` model ([#627](https://github.com/CiscoDevNet/ydk-gen/issues/627))
  * Changed default arguments to be more pythonic ([#682](https://github.com/CiscoDevNet/ydk-gen/issues/682))
  * Handled XML escape characters included in yang models ([#683](https://github.com/CiscoDevNet/ydk-gen/issues/683))
  * Improved handling XML declaration in XML payloads ([#662](https://github.com/CiscoDevNet/ydk-gen/issues/662))
  * Fixed support for yang models with lists as top-level nodes ([#728](https://github.com/CiscoDevNet/ydk-gen/issues/728))
  * Added support for yang 1.1 `action` statement in path API ([#717](https://github.com/CiscoDevNet/ydk-gen/issues/717))
##### Netconf provider improvements
  * Added support for connecting to devices with no `get-schema` support ([#554](https://github.com/CiscoDevNet/ydk-gen/issues/544))
##### ydk-gen improvements
  * Updated leafs in python model APIs to use native python types. ([604](https://github.com/CiscoDevNet/ydk-gen/issues/604))
  * Improved the size and performance of Golang model APIs ([604](https://github.com/CiscoDevNet/ydk-gen/issues/604))
  * Fixed issue with handling of some typedefs in Golang ([706](https://github.com/CiscoDevNet/ydk-gen/issues/706), [747](https://github.com/CiscoDevNet/ydk-gen/issues/747))
##### Documentation improvements
  * Improved enum documentation ([716](https://github.com/CiscoDevNet/ydk-gen/issues/716))
  * Enhanced table of contents for documentation ([715](https://github.com/CiscoDevNet/ydk-gen/issues/715))
##### Testing/error improvements
  * Improved ydk-gen error reporting and fixed `--one-class-per-module` option of generating python packages ([604](https://github.com/CiscoDevNet/ydk-gen/issues/604))
  * Added coverage for Golang and C++ ([740](https://github.com/CiscoDevNet/ydk-gen/issues/740), [705](https://github.com/CiscoDevNet/ydk-gen/issues/705))
##### Installation improvements
  * Introduced automated docker builds to produce docker images with `ydk-gen`, `ydk-py` and `ydk-go` pre-installed ([724](https://github.com/CiscoDevNet/ydk-gen/issues/724))
  * Removed `epel-release` as one of the requirements for libydk RPM ([#627](https://github.com/CiscoDevNet/ydk-gen/issues/627))
  * Added testing for `libydk` packages ([604](https://github.com/CiscoDevNet/ydk-gen/issues/604))

### 2018-01-31 version 0.7.0

#### Go
##### Introduced Go language YDK support (alpha version)
* Added support for all existing `ydk core` services, providers, types and errors in Go
* Added support for all existing `ydk bundles` including `ietf`, `openconfig`, `cisco-ios-xr` and `cisco-ios-xe` in Go
* [#673](https://github.com/CiscoDevNet/ydk-gen/pull/673), [#663](https://github.com/CiscoDevNet/ydk-gen/pull/), [#660](https://github.com/CiscoDevNet/ydk-gen/pull/660), [#658](https://github.com/CiscoDevNet/ydk-gen/pull/658), [#606](https://github.com/CiscoDevNet/ydk-gen/pull/606), [#605](https://github.com/CiscoDevNet/ydk-gen/pull/605)
