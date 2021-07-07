package gmp

import (
	"encoding/xml"
	"time"
)

type GetPortListsCommand struct {
	XMLName    xml.Name `xml:"get_port_lists"`
	Filter     string   `xml:"filter,attr,omitempty"`
	FiltID     string   `xml:"filt_id,attr,omitempty"`
	PortListID string   `xml:"port_list_id,attr,omitempty"`
	Details    bool     `xml:"details,attr,omitempty"`
	Targets    bool     `xml:"targets,attr,omitempty"`
	Trash      bool     `xml:"trash,attr,omitempty"`
}

type getPortListsResponseFiltersKeywordsKeyword struct {
	Column   string `xml:"column"`
	Relation string `xml:"relation"`
	Value    string `xml:"value"`
}

type getPortListsResponseFiltersKeywords struct {
	Keyword []getPortListsResponseFiltersKeywordsKeyword `xml:"keyword"`
}

type getPortListsResponseFilters struct {
	ID       string                              `xml:"id,attr"`
	Term     string                              `xml:"term"`
	Name     string                              `xml:"name"`
	Keywords getPortListsResponseFiltersKeywords `xml:"keywords"`
}

type getPortListsResponseSortField struct {
	Value string `xml:",chardata"`
	Order string `xml:"order"`
}

type getPortListsResponseSort struct {
	Value string                        `xml:",chardata"`
	Field getPortListsResponseSortField `xml:"field"`
}

type getPortListsResponsePortLists struct {
	Start int `xml:"start"`
	Max   int `xml:"max"`
}

type getPortListsResponsePortListCount struct {
	Filtered int `xml:"filtered"`
	Page     int `xml:"page"`
}

type GetPortListsResponse struct {
	XMLName       xml.Name                          `xml:"get_port_lists_response"`
	Status        string                            `xml:"status,attr"`
	StatusText    string                            `xml:"status_text,attr"`
	PortList        []getPortListsResponsePortList    `xml:"port_list"`
	Filters       getPortListsResponseFilters       `xml:"filters"`
	Sort          getPortListsResponseSort          `xml:"sort"`
	PortLists     getPortListsResponsePortLists     `xml:"port_lists"`
	PortListCount getPortListsResponsePortListCount `xml:"port_list_count"`
}

type getPortListsResponsePortListPortCount struct {
	All int `xml:"all"`
	Tcp int `xml:"tcp"`
	Udp int `xml:"udp"`
}

type getPortListsResponsePortListPortRangesPortRange struct {
	ID      string `xml:"id"`
	Start   int    `xml:"start"`
	End     int    `xml:"end"`
	Type    string `xml:"type"`
	Comment string `xml:"comment"`
}

type getPortListsResponsePortListPortRanges struct {
	PortRange []getPortListsResponsePortListPortRangesPortRange `xml:"port_range"`
}

type getPortListsResponsePortListOwner struct {
	Name string `xml:"name"`
}

type getPortListsResponsePortListTargetsTarget struct {
	ID          string `xml:"id"`
	Name        string `xml:"name"`
	Permissions string `xml:"permissions"`
}

type getPortListsResponsePortListTargets struct {
	Target []getPortListsResponsePortListTargetsTarget `xml:"target"`
}

type getPortListsResponsePortListPermissions struct {
	Permission []getPortListsResponsePortListPermissionsPermission `xml:"permission"`
}

type getPortListsResponsePortListPermissionsPermission struct {
	Name string `xml:"name"`
}

type getPortListsResponsePortListUserTags struct {
	Count int `xml:"count"`
}

type getPortListsResponsePortList struct {
	ID               string                                  `xml:"id,attr"`
	Owner            getPortListsResponsePortListOwner       `xml:"owner"`
	Name             string                                  `xml:"name"`
	Comment          string                                  `xml:"comment"`
	CreationTime     time.Time                               `xml:"creation_time"`
	ModificationTime time.Time                               `xml:"modification_time"`
	Writable         bool                                    `xml:"writable"`
	InUse            bool                                    `xml:"in_use"`
	PortCount        getPortListsResponsePortListPortCount   `xml:"port_count"`
	PortRanges       getPortListsResponsePortListPortRanges  `xml:"port_ranges"`
	Targets          getPortListsResponsePortListTargets     `xml:"targets"`
	Permissions      getPortListsResponsePortListPermissions `xml:"permissions"`
	UserTags         getPortListsResponsePortListUserTags    `xml:"user_tags"`
}
