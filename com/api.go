package com

type FnApiResponse struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

// /api/v3
// version
type VersionA struct {
	Appname string `json:"appname"`
	Version string `json:"version"`
}

// mode
type ModeA struct {
	Mode string `json:"mode"`
}

// wallet/create?password=(string)
type CreateA struct {
	Address string `json:"address"`
}

// wallet/import?prikey=(string)&password=(string)
type ImportA struct {
	Address string `json:"address"`
}

// wallet/list
type ListA struct {
	AddressList []string `json:"address_list"`
}

// wallet/remove?address=(string)&password=(string)
type RemoveA struct {
	Success bool `json:"success"`
}

// wallet/prikey?address=x&password=x
type PrikeyA struct {
	Prikey string `json:"prikey"`
}

// wallet/setroot?address=x
type SetrootA struct {
	Success bool `json:"success"`
}

// wallet/getroot
type GetrootA struct {
	Address string `json:"address"`
}

