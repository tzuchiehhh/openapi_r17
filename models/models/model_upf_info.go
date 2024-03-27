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

// Information of an UPF NF Instance
type UpfInfo struct {
	SNssaiUpfInfoList     []SnssaiUpfInfoItem    `json:"sNssaiUpfInfoList" yaml:"sNssaiUpfInfoList" bson:"sNssaiUpfInfoList,omitempty"`
	SmfServingArea        []string               `json:"smfServingArea,omitempty" yaml:"smfServingArea" bson:"smfServingArea,omitempty"`
	InterfaceUpfInfoList  []InterfaceUpfInfoItem `json:"interfaceUpfInfoList,omitempty" yaml:"interfaceUpfInfoList" bson:"interfaceUpfInfoList,omitempty"`
	IwkEpsInd             bool                   `json:"iwkEpsInd,omitempty" yaml:"iwkEpsInd" bson:"iwkEpsInd,omitempty"`
	PduSessionTypes       []PduSessionType       `json:"pduSessionTypes,omitempty" yaml:"pduSessionTypes" bson:"pduSessionTypes,omitempty"`
	AtsssCapability       *AtsssCapability       `json:"atsssCapability,omitempty" yaml:"atsssCapability" bson:"atsssCapability,omitempty"`
	UeIpAddrInd           bool                   `json:"ueIpAddrInd,omitempty" yaml:"ueIpAddrInd" bson:"ueIpAddrInd,omitempty"`
	TaiList               []Tai                  `json:"taiList,omitempty" yaml:"taiList" bson:"taiList,omitempty"`
	TaiRangeList          []TaiRange             `json:"taiRangeList,omitempty" yaml:"taiRangeList" bson:"taiRangeList,omitempty"`
	WAgfInfo              *WAgfInfo              `json:"wAgfInfo,omitempty" yaml:"wAgfInfo" bson:"wAgfInfo,omitempty"`
	TngfInfo              *TngfInfo              `json:"tngfInfo,omitempty" yaml:"tngfInfo" bson:"tngfInfo,omitempty"`
	TwifInfo              *TwifInfo              `json:"twifInfo,omitempty" yaml:"twifInfo" bson:"twifInfo,omitempty"`
	Priority              int32                  `json:"priority,omitempty" yaml:"priority" bson:"priority,omitempty"`
	RedundantGtpu         bool                   `json:"redundantGtpu,omitempty" yaml:"redundantGtpu" bson:"redundantGtpu,omitempty"`
	Ipups                 bool                   `json:"ipups,omitempty" yaml:"ipups" bson:"ipups,omitempty"`
	DataForwarding        bool                   `json:"dataForwarding,omitempty" yaml:"dataForwarding" bson:"dataForwarding,omitempty"`
	SupportedPfcpFeatures string                 `json:"supportedPfcpFeatures,omitempty" yaml:"supportedPfcpFeatures" bson:"supportedPfcpFeatures,omitempty"`
}