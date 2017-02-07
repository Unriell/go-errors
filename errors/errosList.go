package errors

import "go.bq.com/pdtdev/go-errors/dto"

var values = map[string]dto.ErrorDto{
	"BQ0000000": dto.ErrorDto{
		ID:     "BQ0000000",
		Status: 500,
		Msg:    "Error Not found",
	},
	"BQ0000001": dto.ErrorDto{
		ID:     "BQ0000001",
		Status: 400,
		Msg:    "Invalid inbound entity",
	},
	"BQ0000002": dto.ErrorDto{
		ID:     "BQ0000002",
		Status: 404,
		Msg:    "Resource not found",
	},
	"BQ0000003": dto.ErrorDto{
		ID:     "BQ0000003",
		Status: 500,
		Msg:    "Undefined Error",
	},
	"BQ0000004": dto.ErrorDto{
		ID:     "BQ0000004",
		Status: 500,
		Msg:    "Connector is down",
	},
	"BQ0000005": dto.ErrorDto{
		ID:     "BQ0000005",
		Status: 401,
		Msg:    "Authentication error",
	},
	"BQ0000006": dto.ErrorDto{
		ID:     "BQ0000006",
		Status: 500,
		Msg:    "private keys errors",
	},
	"BQ0000007": dto.ErrorDto{
		ID:     "BQ0000007",
		Status: 500,
		Msg:    "public keys errors",
	},
	"BQ0000008": dto.ErrorDto{
		ID:     "BQ0000008",
		Status: 500,
		Msg:    "sign token error",
	},
	"BQ0000009": dto.ErrorDto{
		ID:     "BQ0000009",
		Status: 415,
		Msg:    "content-type application/json it's required",
	},
	"BQ0000010": dto.ErrorDto{
		ID:     "BQ0000010",
		Status: 500,
		Msg:    "mysql fatal error",
	},
	"BQ0000011": dto.ErrorDto{
		ID:     "BQ0000011",
		Status: 400,
		Msg:    "unssuported auth algorithm",
	},
}
