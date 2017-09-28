package errors

var values = map[string]ErrorDto{
	"BQ0000000": ErrorDto{
		ID:     "BQ0000000",
		Status: 500,
		Msg:    "Error Not found",
	},
	"BQ0000001": ErrorDto{
		ID:     "BQ0000001",
		Status: 400,
		Msg:    "Invalid inbound entity",
	},
	"BQ0000002": ErrorDto{
		ID:     "BQ0000002",
		Status: 404,
		Msg:    "Resource not found",
	},
	"BQ0000003": ErrorDto{
		ID:     "BQ0000003",
		Status: 500,
		Msg:    "Undefined Error",
	},
	"BQ0000004": ErrorDto{
		ID:     "BQ0000004",
		Status: 500,
		Msg:    "Connector is down",
	},
	"BQ0000005": ErrorDto{
		ID:     "BQ0000005",
		Status: 401,
		Msg:    "Authentication error",
	},
	"BQ0000006": ErrorDto{
		ID:     "BQ0000006",
		Status: 500,
		Msg:    "private keys errors",
	},
	"BQ0000007": ErrorDto{
		ID:     "BQ0000007",
		Status: 500,
		Msg:    "public keys errors",
	},
	"BQ0000008": ErrorDto{
		ID:     "BQ0000008",
		Status: 500,
		Msg:    "sign token error",
	},
	"BQ0000009": ErrorDto{
		ID:     "BQ0000009",
		Status: 415,
		Msg:    "content-type application/json it's required",
	},
	"BQ0000010": ErrorDto{
		ID:     "BQ0000010",
		Status: 500,
		Msg:    "mysql fatal error",
	},
	"BQ0000011": ErrorDto{
		ID:     "BQ0000011",
		Status: 400,
		Msg:    "unssuported auth algorithm",
	},
	"BQ0000012": ErrorDto{
		ID:     "BQ0000012",
		Status: 409,
		Msg:    "Resource conflict",
	},
	"BQ0000013": ErrorDto{
		ID:     "BQ0000013",
		Status: 500,
		Msg:    "Redis fatal error",
	},
	"BQ0000014": ErrorDto{
		ID:     "BQ0000014",
		Status: 500,
		Msg:    "Unmarshal redis error",
	},
	"BQ0000015": ErrorDto{
		ID:     "BQ0000015",
		Status: 404,
		Msg:    "Redis not found",
	},
	"BQ0000016": ErrorDto{
		ID:     "BQ0000016",
		Status: 401,
		Msg:    "authorization error",
	},
	"BQ0000017": ErrorDto{
		ID:     "BQ0000017",
		Status: 400,
		Msg:    "invalid authorization_token format",
	},
	"BQ0000018": ErrorDto{
		ID:     "BQ0000018",
		Status: 401,
		Msg:    "Token expires",
	},
	"BQ0000019": ErrorDto{
		ID:     "BQ0000019",
		Status: 409,
		Msg:    "Notification conflict",
	},
}
