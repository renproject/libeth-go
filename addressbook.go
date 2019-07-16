package libeth

import (
	"github.com/ethereum/go-ethereum/common"
)

// AddressBook maps contract names to their addresses.
type AddressBook map[string]common.Address

// MainnetAddressBook is the address book for the Mainnet Ren network on the
// main Ethereum network.
var MainnetAddressBook = AddressBook{
	// Ren contracts
	"DarknodeRegistry":      common.HexToAddress("0x34bd421C7948Bc16f826Fd99f9B785929b121633"),
	"DarknodeRegistryStore": common.HexToAddress("0x06df0657ba5e8f5339e742212669f6e7ee3c5057"),
	"DarknodePayment":       common.HexToAddress("0x5a7802E66b067cB1770ee5b1165AA201690A8B6a"),
	"DarknodePaymentStore":  common.HexToAddress("0x731Ea4Ba77fF184d89dBeB160A0078274Acbe9D2"),
	// Tokens
	"DGX":  common.HexToAddress("0x4f3AfEC4E5a3F2A6a1A411DEF7D7dFe50eE057bF"),
	"TUSD": common.HexToAddress("0x0000000000085d4780B73119b644AE5ecd22b376"),
	"REN":  common.HexToAddress("0x408e41876cccdc0f92210600ef50372656052a38"),
	"WBTC": common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"),
	"ZRX":  common.HexToAddress("0xE41d2489571d322189246DaFA5ebDe1F4699F498"),
	"OMG":  common.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"),
	"USDC": common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
	"GUSD": common.HexToAddress("0x056fd409e1d7a124bd7017459dfea2f387b6d5cd"),
	"DAI":  common.HexToAddress("0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359"),
	"PAX":  common.HexToAddress("0x8e870d67f660d95d5be530380d0ec0bd388289e1"),
}

// RopstenAddressBook is the empty address book for the Ropsten Ethereum
// network.
var RopstenAddressBook = AddressBook{}

// TestnetAddressBook is the address book for the Testnet Ren network on the
// Kovan Ethereum network.
var TestnetAddressBook = AddressBook{
	// Ren contracts
	"DarknodeRegistry":      common.HexToAddress("0x1C6309618338D0EDf9a7Ea8eA18E060fD323020D"),
	"DarknodeRegistryStore": common.HexToAddress("0x88e4477e4fdd677aee2dc9376471d45c198669fa"),
	"DarknodePayment":       common.HexToAddress("0x8E11B87547f4072CC8A094F2888201CAF4EA0B9e"),
	"DarknodePaymentStore":  common.HexToAddress("0xA9411C3AD1fBE168fd119A3B32fB481a0b9877A9"),
	// Shifter contracts
	"ShifterRegistry": common.HexToAddress("0x1a62cE8d00d061A05A70dB318f7dfE9Ecd9fFcB0"),
	"zBTC":            common.HexToAddress("0x1aFf7F90Bab456637a17d666D647Ea441A189F2d"),
	"BTCShifter":      common.HexToAddress("0x8a0E8dfC2389726DF1c0bAB874dd2C9A6031b28f"),
	"zZEC":            common.HexToAddress("0x51f41985134842Aa009CACdD28B2C69fA97C4738"),
	"ZECShifter":      common.HexToAddress("0xEeC88ec22E6582409631530F44E48057E7Ed9bBa"),
	// Tokens
	"WBTC": common.HexToAddress("0xA1D3EEcb76285B4435550E4D963B8042A8bffbF0"),
	"REN":  common.HexToAddress("0x2CD647668494c1B15743AB283A0f980d90a87394"),
	"ZRX":  common.HexToAddress("0x6EB628dCeFA95802899aD3A9EE0C7650Ac63d543"),
	"OMG":  common.HexToAddress("0x66497ba75dD127b46316d806c077B06395918064"),
	"USDC": common.HexToAddress("0x3f0a4aed397c66d7b7dde1d170321f87656b14cc"),
	"GUSD": common.HexToAddress("0xA9CF366E9fb4F7959452d7a17A6F88ee2A20e9DA"),
	"DAI":  common.HexToAddress("0xc4375b7de8af5a38a93548eb8453a498222c4ff2"),
	"TUSD": common.HexToAddress("0x525389752ffe6487d33EF53FBcD4E5D3AD7937a0"),
	"DGX":  common.HexToAddress("0x7d6D31326b12B6CBd7f054231D47CbcD16082b71"),
	"PAX":  common.HexToAddress("0x3584087444dabf2e0d29284766142ac5c3a9a2b7"),
}

// DevnetAddressBook is the address book for the Devnet Ren network on the
// Kovan Ethereum network.
var DevnetAddressBook = AddressBook{
	// Ren contracts
	"DarknodeRegistry":      common.HexToAddress("0x6E1a6b85f05bfec5c24C7a26E302cB28e639651c"),
	"DarknodeRegistryStore": common.HexToAddress("0xC126a308dd07Adfa4a445686dcF7CbC423185593"),
	"DarknodePayment":       common.HexToAddress("0x1f1b1d015Fc31d425C616cC35E39e31686DA69A8"),
	"DarknodePaymentStore":  common.HexToAddress("0x6341DF1012E862f766Fcd72e0fCAAc5a3839CFef"),
	// Shifter contracts
	"ShifterRegistry": common.HexToAddress("0x797029C0e76Bb9c8827802c68a5df6e6D93aC14f"),
	"zBTC":            common.HexToAddress("0x4eB1403f565c3e3145Afc3634F16e2F092545C2a"),
	"BTCShifter":      common.HexToAddress("0x7a40fE9FB464510215C41Eae1216973514eeEBB1"),
	"zZEC":            common.HexToAddress("0x92E9aC7c07bb40eD9DAF1DbD88E7Dd77f5754529"),
	"ZECShifter":      common.HexToAddress("0x5345425363A49904De0746E7c991B05380E42E27"),
	// Tokens
	"WBTC": common.HexToAddress("0xA1D3EEcb76285B4435550E4D963B8042A8bffbF0"),
	"REN":  common.HexToAddress("0x2CD647668494c1B15743AB283A0f980d90a87394"),
	"ZRX":  common.HexToAddress("0x6EB628dCeFA95802899aD3A9EE0C7650Ac63d543"),
	"OMG":  common.HexToAddress("0x66497ba75dD127b46316d806c077B06395918064"),
	"USDC": common.HexToAddress("0x3f0a4aed397c66d7b7dde1d170321f87656b14cc"),
	"GUSD": common.HexToAddress("0xA9CF366E9fb4F7959452d7a17A6F88ee2A20e9DA"),
	"DAI":  common.HexToAddress("0xc4375b7de8af5a38a93548eb8453a498222c4ff2"),
	"TUSD": common.HexToAddress("0x525389752ffe6487d33EF53FBcD4E5D3AD7937a0"),
	"DGX":  common.HexToAddress("0x7d6D31326b12B6CBd7f054231D47CbcD16082b71"),
	"PAX":  common.HexToAddress("0x3584087444dabf2e0d29284766142ac5c3a9a2b7"),
}

// LocalnetAddressBook is the address book for the Localnet Ren network on the
// Kovan Ethereum network.
var LocalnetAddressBook = AddressBook{
	// Ren contracts
	"DarknodeRegistry":      common.HexToAddress("0xA7F5B11657AA2796B9355DceF075202C26507B9B"),
	"DarknodeRegistryStore": common.HexToAddress("0x46d016F50837a5DF8fe229127e54fb18B621bAeF"),
	"DarknodePayment":       common.HexToAddress("0x7c71E53853863ce0a3BE7D024EF99aba7d872bfe"),
	"DarknodePaymentStore":  common.HexToAddress("0x72Acdf4f0E3245262E46Bd8daCc207Df7CF3A534"),
	// Shifter contracts
	"ShifterRegistry": common.HexToAddress("0x33666067A7741B9e88520285C96B776E73281811"),
	"zBTC":            common.HexToAddress("0xDE027035d33CEB2757685E325de1A0b924aA73E6"),
	"BTCShifter":      common.HexToAddress("0x7012ECc13De5Ce416C14C013d9b02b7c37154b37"),
	"zZEC":            common.HexToAddress("0xB7b7be50B13E6817afBb30C93161D0eB388b8f08"),
	"ZECShifter":      common.HexToAddress("0x69AC72Cb35B1AA818e90842C048719a3246ba0BE"),
	// Tokens
	"WBTC": common.HexToAddress("0xA1D3EEcb76285B4435550E4D963B8042A8bffbF0"),
	"REN":  common.HexToAddress("0x2CD647668494c1B15743AB283A0f980d90a87394"),
	"ZRX":  common.HexToAddress("0x6EB628dCeFA95802899aD3A9EE0C7650Ac63d543"),
	"OMG":  common.HexToAddress("0x66497ba75dD127b46316d806c077B06395918064"),
	"USDC": common.HexToAddress("0x3f0a4aed397c66d7b7dde1d170321f87656b14cc"),
	"GUSD": common.HexToAddress("0xA9CF366E9fb4F7959452d7a17A6F88ee2A20e9DA"),
	"DAI":  common.HexToAddress("0xc4375b7de8af5a38a93548eb8453a498222c4ff2"),
	"TUSD": common.HexToAddress("0x525389752ffe6487d33EF53FBcD4E5D3AD7937a0"),
	"DGX":  common.HexToAddress("0x7d6D31326b12B6CBd7f054231D47CbcD16082b71"),
	"PAX":  common.HexToAddress("0x3584087444dabf2e0d29284766142ac5c3a9a2b7"),
}

func NetworkAddressBook(network RenNetwork) AddressBook {
	switch network {
	case Mainnet:
		return MainnetAddressBook
	case Ropsten:
		return RopstenAddressBook
	case Testnet:
		return TestnetAddressBook
	case Devnet:
		return DevnetAddressBook
	case Localnet:
		return LocalnetAddressBook
	default:
		return AddressBook{}
	}
}

type Contracts map[common.Address]string

var ContractABIs = map[string]string{
	"0xF228173631461642c737855d0a4Be25ce89Dd486": "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"bytes\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"commitment\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"LogMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
	"0x44831d9CDD36FB829eb2D9aA904b1078C496DB56": `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"sender","type":"address"},{"name":"recipient","type":"address"},{"name":"amount","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"addedValue","type":"uint256"}],"name":"increaseAllowance","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_amount","type":"uint256"}],"name":"mint","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"account","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"renounceOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"owner","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"isOwner","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_amount","type":"uint256"}],"name":"burn","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"subtractedValue","type":"uint256"}],"name":"decreaseAllowance","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"recipient","type":"address"},{"name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"owner","type":"address"},{"name":"spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[{"name":"_name","type":"string"},{"name":"_symbol","type":"string"},{"name":"_decimals","type":"uint8"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"previousOwner","type":"address"},{"indexed":true,"name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"}]`,
	"0xd56aed911840bd15316ae3038ac7381ebbebd909": "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"bytes\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"spent\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"commitment\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"}],\"name\":\"verifySig\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
	"0x981A342073215904C5069CD64204898501689E0f": "[{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"shiftedTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"newShiftedToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_commitment\",\"type\":\"bytes32\"}],\"name\":\"sigHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"spent\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_commitment\",\"type\":\"bytes32\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"shiftIn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"bytes\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"shiftOut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isShiftedToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_commitment\",\"type\":\"bytes32\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"verifySig\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_vault\",\"type\":\"address\"},{\"name\":\"_fee\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_token\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogShiftIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_token\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogShiftOut\",\"type\":\"event\"}]",
}
