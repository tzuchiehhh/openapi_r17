/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Represents a notification related to a single event that occurred.
type SmfEventExposureEventNotification struct {
	Event SmfEvent `json:"event" yaml:"event" bson:"event,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	TimeStamp *time.Time `json:"timeStamp" yaml:"timeStamp" bson:"timeStamp,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi string `json:"supi,omitempty" yaml:"supi" bson:"supi,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi     string  `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	UeIpAddr *IpAddr `json:"ueIpAddr,omitempty" yaml:"ueIpAddr" bson:"ueIpAddr,omitempty"`
	// Transaction Information.
	TransacInfos []TransactionInfo `json:"transacInfos,omitempty" yaml:"transacInfos" bson:"transacInfos,omitempty"`
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	SourceDnai string `json:"sourceDnai,omitempty" yaml:"sourceDnai" bson:"sourceDnai,omitempty"`
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	TargetDnai  string         `json:"targetDnai,omitempty" yaml:"targetDnai" bson:"targetDnai,omitempty"`
	DnaiChgType DnaiChangeType `json:"dnaiChgType,omitempty" yaml:"dnaiChgType" bson:"dnaiChgType,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	SourceUeIpv4Addr   string `json:"sourceUeIpv4Addr,omitempty" yaml:"sourceUeIpv4Addr" bson:"sourceUeIpv4Addr,omitempty"`
	SourceUeIpv6Prefix string `json:"sourceUeIpv6Prefix,omitempty" yaml:"sourceUeIpv6Prefix" bson:"sourceUeIpv6Prefix,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	TargetUeIpv4Addr   string           `json:"targetUeIpv4Addr,omitempty" yaml:"targetUeIpv4Addr" bson:"targetUeIpv4Addr,omitempty"`
	TargetUeIpv6Prefix string           `json:"targetUeIpv6Prefix,omitempty" yaml:"targetUeIpv6Prefix" bson:"targetUeIpv6Prefix,omitempty"`
	SourceTraRouting   *RouteToLocation `json:"sourceTraRouting,omitempty" yaml:"sourceTraRouting" bson:"sourceTraRouting,omitempty"`
	TargetTraRouting   *RouteToLocation `json:"targetTraRouting,omitempty" yaml:"targetTraRouting" bson:"targetTraRouting,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	UeMac string `json:"ueMac,omitempty" yaml:"ueMac" bson:"ueMac,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	AdIpv4Addr   string `json:"adIpv4Addr,omitempty" yaml:"adIpv4Addr" bson:"adIpv4Addr,omitempty"`
	AdIpv6Prefix string `json:"adIpv6Prefix,omitempty" yaml:"adIpv6Prefix" bson:"adIpv6Prefix,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	ReIpv4Addr   string     `json:"reIpv4Addr,omitempty" yaml:"reIpv4Addr" bson:"reIpv4Addr,omitempty"`
	ReIpv6Prefix string     `json:"reIpv6Prefix,omitempty" yaml:"reIpv6Prefix" bson:"reIpv6Prefix,omitempty"`
	PlmnId       *PlmnId    `json:"plmnId,omitempty" yaml:"plmnId" bson:"plmnId,omitempty"`
	AccType      AccessType `json:"accType,omitempty" yaml:"accType" bson:"accType,omitempty"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.
	PduSeId          int32                 `json:"pduSeId,omitempty" yaml:"pduSeId" bson:"pduSeId,omitempty"`
	RatType          RatType               `json:"ratType,omitempty" yaml:"ratType" bson:"ratType,omitempty"`
	DddStatus        DlDataDeliveryStatus  `json:"dddStatus,omitempty" yaml:"dddStatus" bson:"dddStatus,omitempty"`
	DddTraDescriptor *DddTrafficDescriptor `json:"dddTraDescriptor,omitempty" yaml:"dddTraDescriptor" bson:"dddTraDescriptor,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	MaxWaitTime *time.Time            `json:"maxWaitTime,omitempty" yaml:"maxWaitTime" bson:"maxWaitTime,omitempty"`
	CommFailure *CommunicationFailure `json:"commFailure,omitempty" yaml:"commFailure" bson:"commFailure,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4Addr     string         `json:"ipv4Addr,omitempty" yaml:"ipv4Addr" bson:"ipv4Addr,omitempty"`
	Ipv6Prefixes []string       `json:"ipv6Prefixes,omitempty" yaml:"ipv6Prefixes" bson:"ipv6Prefixes,omitempty"`
	Ipv6Addrs    []string       `json:"ipv6Addrs,omitempty" yaml:"ipv6Addrs" bson:"ipv6Addrs,omitempty"`
	PduSessType  PduSessionType `json:"pduSessType,omitempty" yaml:"pduSessType" bson:"pduSessType,omitempty"`
	// Unsigned integer identifying a QoS flow, within the range 0 to 63.
	Qfi int32 `json:"qfi,omitempty" yaml:"qfi" bson:"qfi,omitempty"`
	// String providing an application identifier.
	AppId string `json:"appId,omitempty" yaml:"appId" bson:"appId,omitempty"`
	// Descriptor(s) for non-IP traffic. It allows the encoding of multiple UL and/or DL flows. Each entry of the array describes a single Ethernet flow.
	EthFlowDescs []EthFlowDescription `json:"ethFlowDescs,omitempty" yaml:"ethFlowDescs" bson:"ethFlowDescs,omitempty"`
	// Contains the UL and/or DL Ethernet flows. Each entry of the array describes a single Ethernet flow.
	EthfDescs []EthFlowDescription `json:"ethfDescs,omitempty" yaml:"ethfDescs" bson:"ethfDescs,omitempty"`
	// Descriptor(s) for IP traffic. It allows the encoding of multiple UL and/or DL flows. Each entry of the array describes a single IP flow.
	FlowDescs []string `json:"flowDescs,omitempty" yaml:"flowDescs" bson:"flowDescs,omitempty"`
	// Contains the UL and/or DL IP flows. Each entry of the array describes a single IP flow.
	FDescs []string `json:"fDescs,omitempty" yaml:"fDescs" bson:"fDescs,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn      string  `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	Snssai   *Snssai `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
	UlDelays []int32 `json:"ulDelays,omitempty" yaml:"ulDelays" bson:"ulDelays,omitempty"`
	DlDelays []int32 `json:"dlDelays,omitempty" yaml:"dlDelays" bson:"dlDelays,omitempty"`
	RtDelays []int32 `json:"rtDelays,omitempty" yaml:"rtDelays" bson:"rtDelays,omitempty"`
	// Represents the packet delay measurement failure indicator.
	Pdmf         bool          `json:"pdmf,omitempty" yaml:"pdmf" bson:"pdmf,omitempty"`
	TimeWindow   *TimeWindow   `json:"timeWindow,omitempty" yaml:"timeWindow" bson:"timeWindow,omitempty"`
	SmNasFromUe  *SmNasFromUe  `json:"smNasFromUe,omitempty" yaml:"smNasFromUe" bson:"smNasFromUe,omitempty"`
	SmNasFromSmf *SmNasFromSmf `json:"smNasFromSmf,omitempty" yaml:"smNasFromSmf" bson:"smNasFromSmf,omitempty"`
	// Indicates whether the redundant transmission is setup or terminated. Set to \"true\" if  the redundant transmission is setup, otherwise set to \"false\" if the redundant  transmission is terminated. Default value is set to \"false\".
	UpRedTrans bool   `json:"upRedTrans,omitempty" yaml:"upRedTrans" bson:"upRedTrans,omitempty"`
	SsId       string `json:"ssId,omitempty" yaml:"ssId" bson:"ssId,omitempty"`
	BssId      string `json:"bssId,omitempty" yaml:"bssId" bson:"bssId,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	StartWlan *time.Time `json:"startWlan,omitempty" yaml:"startWlan" bson:"startWlan,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	EndWlan      *time.Time                              `json:"endWlan,omitempty" yaml:"endWlan" bson:"endWlan,omitempty"`
	PduSessInfos []SmfEventExposurePduSessionInformation `json:"pduSessInfos,omitempty" yaml:"pduSessInfos" bson:"pduSessInfos,omitempty"`
	UpfInfo      *UpfInformation                         `json:"upfInfo,omitempty" yaml:"upfInfo" bson:"upfInfo,omitempty"`
}