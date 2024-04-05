package constants

import (
	"net/http"
	"net/url"
)

const TESTTHIS = "testthis"

const (
	FILESA = iota
	FILESB
	FILESC
)

const (
	EVENT_CLICK = "Event_Click"
	EVENT_SEARCH = "Event_Search"
)

const (
	VW_INDEX = iota
	VW_APPHEADER
	VW_TABLE
	VW_SIDENAV
)

type EventData struct {
	Id        string
	EventType string
	Event     string
}

type HtmxInfo struct {
	Url string
}

type SubElement struct {
	Type string
	Lbl string
}

type ViewInteface interface {
	ProcessRequest(w http.ResponseWriter, d  url.Values/*ViewInfo*/)
}

type ViewInfo struct {
	Event   int
	Type    string
	Label   string
	ViewId  string
	ViewStr string
}

type RowData struct {
	Data []string
}