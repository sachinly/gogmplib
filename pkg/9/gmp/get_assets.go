package gmp

import (
	"encoding/xml"
	"time"
)

type GetAssetsCommand struct {
	XMLName          xml.Name `xml:"get_assets"`
	AssetID          string   `xml:"id,attr,omitempty"`
	Filter           string   `xml:"filter,attr,omitempty"`
	FiltID           string   `xml:"filt_id,attr,omitempty"`
	IgnorePagination string   `xml:"ignore_pagination,attr,omitempty"`
	// type = host | os
	Type string `xml:"type,attr"`
}

type AssetOwner struct {
	Name string `xml:"name"`
}

type SshCredential struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name"`
}

type GetAssetsResponseAssetPermissions struct {
	Permission []GetAssetsResponseAssetPermissionsPermission `xml:"permission"`
}

type GetAssetsResponseAssetPermissionsPermission struct {
	Name string `xml:"name"`
}

type AssetUserTagsTag struct {
	ID      string `xml:"id,attr"`
	Name    string `xml:"name"`
	Value   string `xml:"value"`
	Comment string `xml:"comment"`
}

type AssetUserTags struct {
	Count int                `xml:"count"`
	Tag   []AssetUserTagsTag `xml:"tag"`
}

type IdentifierSource struct {
	ID      string `xml:"id,attr"`
	Type    string `xml:"type"`
	Data    string `xml:"data"`
	Deleted string `xml:"deleted"`
}

type IdentifierOS struct {
	ID    string `xml:"id,attr"`
	Title string `xml:"title"`
}

type AssetIdentifier struct {
	Name             string           `xml:"name"`
	Value            string           `xml:"value"`
	CreationTime     time.Time        `xml:"creation_time"`
	ModificationTime time.Time        `xml:"modification_time"`
	Source           IdentifierSource `xml:"source"`
	OS               IdentifierOS     `xml:"os"`
}

type Identifiers struct {
	Identifier []AssetIdentifier `xml:"identifier"`
}
type GetAssetsResponseAssetHostSeverity struct {
	Value string `xml:"value"`
}

type GetAssetsResponseAssetHostDetailSource struct {
	ID   string `xml:"id,attr"`
	Type string `xml:"type"`
}

type GetAssetsResponseAssetHostDetail struct {
	Name   string                                 `xml:"name"`
	Value  string                                 `xml:"value"`
	Source GetAssetsResponseAssetHostDetailSource `xml:"source"`
}
type GetAssetsResponseAssetHost struct {
	Severity GetAssetsResponseAssetHostSeverity `xml:"severity"`
	Detail   []GetAssetsResponseAssetHostDetail   `xml:"detail"`
}

type Asset struct {
	ID               string                            `xml:"id,attr"`
	Owner            AssetOwner                        `xml:"owner"`
	Name             string                            `xml:"name"`
	Comment          string                            `xml:"comment"`
	CreationTime     time.Time                         `xml:"creation_time"`
	ModificationTime time.Time                         `xml:"modification_time"`
	Writable         string                            `xml:"writable"`
	InUse            string                            `xml:"in_use"`
	Permissions      GetAssetsResponseAssetPermissions `xml:"permissions"`
	UserTags         AssetUserTags                     `xml:"user_tags"`
	Identifiers      Identifiers                       `xml:"identifiers"`
	Type             string                            `xml:"type"`
	Host             GetAssetsResponseAssetHost        `xml:"host"`
}

type GetAssetsResponseFiltersKeywordsKeyword struct {
	Column   string `xml:"column"`
	Relation string `xml:"relation"`
	Value    string `xml:"value"`
}

type GetAssetsResponseFilters struct {
	ID       string                                    `xml:"id,attr"`
	Term     string                                    `xml:"term"`
	Name     string                                    `xml:"name"`
	Keywords []GetAssetsResponseFiltersKeywordsKeyword `xml:"keywords"`
}

type GetAssetsResponseSortField struct {
	Order string `xml:"order"`
}

type GetAssetsResponseSort struct {
	Value string                     `xml:",chardata"`
	Field GetAssetsResponseSortField `xml:"field"`
}

type GetAssetsResponseResults struct {
	Start string `xml:"start,attr"`
	Max   string `xml:"max,attr"`
}

type GetAssetsResponseResultCount struct {
	Filtered int `xml:"filtered"`
	Page     int `xml:"page"`
}
type GetAssetsResponse struct {
	XMLName    xml.Name                     `xml:"get_assets_response"`
	Status     string                       `xml:"status,attr"`
	StatusText string                       `xml:"status_text,attr"`
	Asset      []Asset                      `xml:"asset"`
	Filters    GetAssetsResponseFilters     `xml:"filters"`
	Sort       GetAssetsResponseSort        `xml:"sort"`
	Assets     GetAssetsResponseResults     `xml:"assets"`
	AssetCount GetAssetsResponseResultCount `xml:"asset_count"`
}
