/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291 V17.9.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 3.1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type ProseChargingInformation struct {
	AnnouncingPlmnID            *PlmnId              `json:"announcingPlmnID,omitempty" yaml:"announcingPlmnID" bson:"announcingPlmnID,omitempty"`
	AnnouncingUeHplmnIdentifier *PlmnId              `json:"announcingUeHplmnIdentifier,omitempty" yaml:"announcingUeHplmnIdentifier" bson:"announcingUeHplmnIdentifier,omitempty"`
	AnnouncingUeVplmnIdentifier *PlmnId              `json:"announcingUeVplmnIdentifier,omitempty" yaml:"announcingUeVplmnIdentifier" bson:"announcingUeVplmnIdentifier,omitempty"`
	MonitoringUeHplmnIdentifier *PlmnId              `json:"monitoringUeHplmnIdentifier,omitempty" yaml:"monitoringUeHplmnIdentifier" bson:"monitoringUeHplmnIdentifier,omitempty"`
	MonitoringUeVplmnIdentifier *PlmnId              `json:"monitoringUeVplmnIdentifier,omitempty" yaml:"monitoringUeVplmnIdentifier" bson:"monitoringUeVplmnIdentifier,omitempty"`
	DiscovererUeHplmnIdentifier *PlmnId              `json:"discovererUeHplmnIdentifier,omitempty" yaml:"discovererUeHplmnIdentifier" bson:"discovererUeHplmnIdentifier,omitempty"`
	DiscovererUeVplmnIdentifier *PlmnId              `json:"discovererUeVplmnIdentifier,omitempty" yaml:"discovererUeVplmnIdentifier" bson:"discovererUeVplmnIdentifier,omitempty"`
	DiscovereeUeHplmnIdentifier *PlmnId              `json:"discovereeUeHplmnIdentifier,omitempty" yaml:"discovereeUeHplmnIdentifier" bson:"discovereeUeHplmnIdentifier,omitempty"`
	DiscovereeUeVplmnIdentifier *PlmnId              `json:"discovereeUeVplmnIdentifier,omitempty" yaml:"discovereeUeVplmnIdentifier" bson:"discovereeUeVplmnIdentifier,omitempty"`
	MonitoredPlmnIdentifier     *PlmnId              `json:"monitoredPlmnIdentifier,omitempty" yaml:"monitoredPlmnIdentifier" bson:"monitoredPlmnIdentifier,omitempty"`
	ProseApplicationID          string               `json:"proseApplicationID,omitempty" yaml:"proseApplicationID" bson:"proseApplicationID,omitempty"`
	ApplicationId               string               `json:"ApplicationId,omitempty" yaml:"ApplicationId" bson:"ApplicationId,omitempty"`
	ApplicationSpecificDataList []string             `json:"applicationSpecificDataList,omitempty" yaml:"applicationSpecificDataList" bson:"applicationSpecificDataList,omitempty"`
	ProseFunctionality          ProseFunctionality   `json:"proseFunctionality,omitempty" yaml:"proseFunctionality" bson:"proseFunctionality,omitempty"`
	ProseEventType              ProseEventType       `json:"proseEventType,omitempty" yaml:"proseEventType" bson:"proseEventType,omitempty"`
	DirectDiscoveryModel        DirectDiscoveryModel `json:"directDiscoveryModel,omitempty" yaml:"directDiscoveryModel" bson:"directDiscoveryModel,omitempty"`
	ValidityPeriod              int32                `json:"validityPeriod,omitempty" yaml:"validityPeriod" bson:"validityPeriod,omitempty"`
	RoleOfUE                    RoleOfUe             `json:"roleOfUE,omitempty" yaml:"roleOfUE" bson:"roleOfUE,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ProseRequestTimestamp *time.Time `json:"proseRequestTimestamp,omitempty" yaml:"proseRequestTimestamp" bson:"proseRequestTimestamp,omitempty"`
	PC3ProtocolCause      int32      `json:"pC3ProtocolCause,omitempty" yaml:"pC3ProtocolCause" bson:"pC3ProtocolCause,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	MonitoringUEIdentifier   string     `json:"monitoringUEIdentifier,omitempty" yaml:"monitoringUEIdentifier" bson:"monitoringUEIdentifier,omitempty"`
	RequestedPLMNIdentifier  *PlmnId    `json:"requestedPLMNIdentifier,omitempty" yaml:"requestedPLMNIdentifier" bson:"requestedPLMNIdentifier,omitempty"`
	TimeWindow               int32      `json:"timeWindow,omitempty" yaml:"timeWindow" bson:"timeWindow,omitempty"`
	RangeClass               RangeClass `json:"rangeClass,omitempty" yaml:"rangeClass" bson:"rangeClass,omitempty"`
	ProximityAlertIndication bool       `json:"proximityAlertIndication,omitempty" yaml:"proximityAlertIndication" bson:"proximityAlertIndication,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ProximityAlertTimestamp *time.Time `json:"proximityAlertTimestamp,omitempty" yaml:"proximityAlertTimestamp" bson:"proximityAlertTimestamp,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ProximityCancellationTimestamp *time.Time                `json:"proximityCancellationTimestamp,omitempty" yaml:"proximityCancellationTimestamp" bson:"proximityCancellationTimestamp,omitempty"`
	RelayIPAddress                 *IpAddr                   `json:"relayIPAddress,omitempty" yaml:"relayIPAddress" bson:"relayIPAddress,omitempty"`
	ProseUEToNetworkRelayUEID      string                    `json:"proseUEToNetworkRelayUEID,omitempty" yaml:"proseUEToNetworkRelayUEID" bson:"proseUEToNetworkRelayUEID,omitempty"`
	ProseDestinationLayer2ID       string                    `json:"proseDestinationLayer2ID,omitempty" yaml:"proseDestinationLayer2ID" bson:"proseDestinationLayer2ID,omitempty"`
	PFIContainerInformation        []PfiContainerInformation `json:"pFIContainerInformation,omitempty" yaml:"pFIContainerInformation" bson:"pFIContainerInformation,omitempty"`
	TransmissionDataContainer      []Pc5DataContainer        `json:"transmissionDataContainer,omitempty" yaml:"transmissionDataContainer" bson:"transmissionDataContainer,omitempty"`
	ReceptionDataContainer         []Pc5DataContainer        `json:"receptionDataContainer,omitempty" yaml:"receptionDataContainer" bson:"receptionDataContainer,omitempty"`
}