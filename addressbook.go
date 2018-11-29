package beth

import (
	"github.com/ethereum/go-ethereum/common"
)

type AddressBook map[string]common.Address

type Address struct {
	Name    string
	Address string
}

var MainnetAddresses = []Address{
	Address{
		Name:    "DGX",
		Address: "0x4f3AfEC4E5a3F2A6a1A411DEF7D7dFe50eE057bF",
	},
	Address{
		Name:    "TUSD",
		Address: "0x8dd5fbCe2F6a956C3022bA3663759011Dd51e73E",
	},
	Address{
		Name:    "REN",
		Address: "0x21C482f153D0317fe85C60bE1F7fa079019fcEbD",
	},
	Address{
		Name:    "ZRX",
		Address: "0xE41d2489571d322189246DaFA5ebDe1F4699F498",
	},
	Address{
		Name:    "OMG",
		Address: "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07",
	},
}

var RopstenAddresses = []Address{
	Address{
		Name:    "RenExOrderbook",
		Address: "0xA1D3EEcb76285B4435550E4D963B8042A8b11111",
	},
	Address{
		Name:    "RenExSettlement",
		Address: "0xA1D3EEcb76285B4435550E4D963B8042A8b11111",
	},
	Address{
		Name:    "WBTC",
		Address: "0xA1D3EEcb76285B4435550E4D963B8042A8b11111",
	},
	Address{
		Name:    "SwapperdETH",
		Address: "0x2218fa20c33765e7e01671ee6aaca75fbaf3a974",
	},
	Address{
		Name:    "SwapperdWBTC",
		Address: "0x2218fa20c33765e7e01671ee6aaca75fbaf3a974",
	},
}

var KovanAddresses = []Address{
	Address{
		Name:    "RenExOrderbook",
		Address: "0xA1D3EEcb76285B4435550E4D963B8042A8b11111",
	},
	Address{
		Name:    "RenExSettlement",
		Address: "0xA1D3EEcb76285B4435550E4D963B8042A8b11111",
	},
	Address{
		Name:    "WBTC",
		Address: "0xA1D3EEcb76285B4435550E4D963B8042A8bffbF0",
	},
	Address{
		Name:    "SwapperdETH",
		Address: "0x2218fa20c33765e7e01671ee6aaca75fbaf3a974",
	},
	Address{
		Name:    "DGX",
		Address: "0x932F4580B261e9781A6c3c102133C8fDd4503DFc",
	},
	Address{
		Name:    "TUSD",
		Address: "0x525389752ffe6487d33EF53FBcD4E5D3AD7937a0",
	},
	Address{
		Name:    "REN",
		Address: "0x2CD647668494c1B15743AB283A0f980d90a87394",
	},
	Address{
		Name:    "ZRX",
		Address: "0x6EB628dCeFA95802899aD3A9EE0C7650Ac63d543",
	},
	Address{
		Name:    "OMG",
		Address: "0x66497ba75dD127b46316d806c077B06395918064",
	},
}

func DefaultAddressBook(network int64) AddressBook {
	addrs := []Address{}
	addrBook := AddressBook{}

	switch network {
	case 1:
		addrs = MainnetAddresses
	case 3:
		addrs = RopstenAddresses
	case 42:
		addrs = KovanAddresses
	default:
		return addrBook
	}

	for _, addr := range addrs {
		addrBook[addr.Name] = common.HexToAddress(addr.Address)
	}
	return addrBook
}