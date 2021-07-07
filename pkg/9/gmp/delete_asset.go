package gmp

import "encoding/xml"

type DeleteAssetCommand struct {
	XMLName  xml.Name `xml:"delete_asset"`
	AssetID  string   `xml:"asset_id,attr,omitempty"`
	ReportID string   `xml:"report_id,attr,omitempty"`
}

type DeleteAssetResponse struct {
	XMLName    xml.Name `xml:"delete_asset_response"`
	Status     string   `xml:"status,attr"`
	StatusText string   `xml:"status_text,attr"`
}
