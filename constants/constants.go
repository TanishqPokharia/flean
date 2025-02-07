package constants

import "github.com/charmbracelet/bubbles/table"

var (
	FilePermission    = 0755
	Dependencies      = []string{"get_it", "dartz", "flutter_bloc", "flutter_riverpod", "provider", "riverpod", "basic_utils", "dio", "intl", "uuid", "go_router", "auto_route", "freezed", "json_serializable"}
	DependencyColumns = []table.Column{
		{Width: 4, Title: "SNo."},
		{Width: 16, Title: "Name"},
		{Width: 16, Title: "Category"},
	}
	DependencyRows = []table.Row{
		{"1", "flutter_bloc", "state management"},
		{"2", "flutter_riverpod", "state management"},
		{"3", "provider", "state management"},
		{"4", "get_it", "service locator"},
		{"5", "dartz", "utility"},
		{"6", "dio", "http networking"},
		{"7", "http", "http networking"},
		{"8", "basic_utils", "utility"},
		{"9", "intl", "utility"},
		{"10", "uuid", "utility"},
		{"11", "go_router", "routing"},
		{"12", "auto_route", "routing"},
		{"13", "freezed", "json generation"},
		{"14", "json_serializable", "json generation"},
	}
)
