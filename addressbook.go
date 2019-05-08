package libeth

import (
	"github.com/ethereum/go-ethereum/common"
)

type AddressBook map[string]common.Address

var MainnetAddressBook = AddressBook{
	"DarknodePayment":      common.HexToAddress("0x5a7802E66b067cB1770ee5b1165AA201690A8B6a"),
	"DarknodePaymentStore": common.HexToAddress("0x731Ea4Ba77fF184d89dBeB160A0078274Acbe9D2"),
	"DarknodeRegistry":     common.HexToAddress("0x34bd421C7948Bc16f826Fd99f9B785929b121633"),
	"DGX":                  common.HexToAddress("0x4f3AfEC4E5a3F2A6a1A411DEF7D7dFe50eE057bF"),
	"TUSD":                 common.HexToAddress("0x0000000000085d4780B73119b644AE5ecd22b376"),
	"REN":                  common.HexToAddress("0x408e41876cccdc0f92210600ef50372656052a38"),
	"WBTC":                 common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"),
	"ZRX":                  common.HexToAddress("0xE41d2489571d322189246DaFA5ebDe1F4699F498"),
	"OMG":                  common.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"),
	"USDC":                 common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
	"GUSD":                 common.HexToAddress("0x056fd409e1d7a124bd7017459dfea2f387b6d5cd"),
	"DAI":                  common.HexToAddress("0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359"),
	"PAX":                  common.HexToAddress("0x8e870d67f660d95d5be530380d0ec0bd388289e1"),
	"ETHSwap":              common.HexToAddress("0x50F2b7aB72551b399aC81081484ef0b0F23aa38C"),
	"WBTCSwap":             common.HexToAddress("0x048babB592a39Dd5B4B0A52CD5638cab536b1bdD"),
	"RENSwap":              common.HexToAddress("0x4B77Ab393928f62307824Ea94D631ec887cA6E49"),
	"TUSDSwap":             common.HexToAddress("0x517F5e4E7FDA6792B338227d4F6EB34c83c5499D"),
	"OMGSwap":              common.HexToAddress("0x665f881Fd19E821F39c306dC3f0D83dC1Db9a18C"),
	"ZRXSwap":              common.HexToAddress("0x2f6A213FA2A67FB126e716D134Ac33e51EF7711B"),
	"DGXSwap":              common.HexToAddress("0x98b37BA3826A1F9295fFB819e209eb0c52BdA7E7"),
	"USDCSwap":             common.HexToAddress("0xCd1d2dD16a424C4857eE433542B3Ba8DdB08aC44"),
	"GUSDSwap":             common.HexToAddress("0x42a28d8F40e2cf0Ad62d39d63104F1dE6f782374"),
	"DAISwap":              common.HexToAddress("0x3F6C8124691D20d67C522a61507AAc2a28376caE"),
	"PAXSwap":              common.HexToAddress("0x2d92eB5FbC9e2B2026aD371B31dc267240cDFE40"),
}

var RopstenAddressBook = AddressBook{}

var KovanAddressBook = AddressBook{
	"RenShift":             common.HexToAddress("0xa68acf54957c70d9e88565398b621dc094c82836"),
	"RenShiftEx":           common.HexToAddress("0x0df3510a4128c0ca11518465f670db970e9302b7"),
	"RenShiftExAdapter":    common.HexToAddress("0x8cfbf788757e767392e707aca1ec18ce26e570fc"),
	"ETH-BTC-Reserve":      common.HexToAddress("0x1fe044862a3d674cd13dc45a74dc179894022379"),
	"ETH-ZEC-Reserve":      common.HexToAddress("0x39cd35f90c2d5a7cfa4f0a76a28da15ab4464770"),
	"ZEC-BTC-Reserve":      common.HexToAddress("0x0fc92017944be95d02bdbc213cbbf93cc4521dcb"),
	"ShiftedBTC":           common.HexToAddress("0xFd44199b94EA4398aEa3dD5E1014e550D4cC5b9B"),
	"ShiftedZEC":           common.HexToAddress("0x18344Da571edfa6c3E37C2c4ec93B93F70692B14"),
	"RenExOrderbook":       common.HexToAddress("0x0000000000000000000000000000000000000000"),
	"RenExSettlement":      common.HexToAddress("0x0000000000000000000000000000000000000000"),
	"DarknodeRegistry":     common.HexToAddress("0x1C6309618338D0EDf9a7Ea8eA18E060fD323020D"),
	"DarknodePayment":      common.HexToAddress("0x8E11B87547f4072CC8A094F2888201CAF4EA0B9e"),
	"DarknodePaymentStore": common.HexToAddress("0xA9411C3AD1fBE168fd119A3B32fB481a0b9877A9"),
	"BitcoinRenShift":      common.HexToAddress("0x2a8368d2a983a0aeae8da0ebc5b7c03a0ea66b37"),
	"ZCashRenShift":        common.HexToAddress("0xd67256552f93b39ac30083b4b679718a061feae6"),
	"ETHSwap":              common.HexToAddress("0xc3183C98F4799d6004d38fF6578E677fa0B31b33"),
	"WBTCSwap":             common.HexToAddress("0xfe6a471D51d009cBbA7E99fB57fF31e7da80E47A"),
	"RENSwap":              common.HexToAddress("0x5B3A355C68BD9F6EAF87687d7388cFBFE9eDa8f3"),
	"TUSDSwap":             common.HexToAddress("0x291c7859e04E3cdE794eDb4739E7352BCb20D289"),
	"OMGSwap":              common.HexToAddress("0x4087DfbD64e53d1f00790B7920394DE329197721"),
	"ZRXSwap":              common.HexToAddress("0x371720f220D697d451350C363f51De8FF4fAE6f8"),
	"DGXSwap":              common.HexToAddress("0x047E9eb23652F085942227eaac32941D6483b632"),
	"USDCSwap":             common.HexToAddress("0x003803380687eD0d07463a1071dd96d20b698fbb"),
	"GUSDSwap":             common.HexToAddress("0x0EE5d4B23B62292fbEd63eeC85eA315eD6925959"),
	"DAISwap":              common.HexToAddress("0x050298877821b80b96fB28E0fa518367C0817723"),
	"PAXSwap":              common.HexToAddress("0x859454f3557f60a3d928184f823246Eb3109C72D"),
	"WBTC":                 common.HexToAddress("0xA1D3EEcb76285B4435550E4D963B8042A8bffbF0"),
	"REN":                  common.HexToAddress("0x2CD647668494c1B15743AB283A0f980d90a87394"),
	"ZRX":                  common.HexToAddress("0x6EB628dCeFA95802899aD3A9EE0C7650Ac63d543"),
	"OMG":                  common.HexToAddress("0x66497ba75dD127b46316d806c077B06395918064"),
	"USDC":                 common.HexToAddress("0x3f0a4aed397c66d7b7dde1d170321f87656b14cc"),
	"GUSD":                 common.HexToAddress("0xA9CF366E9fb4F7959452d7a17A6F88ee2A20e9DA"),
	"DAI":                  common.HexToAddress("0xc4375b7de8af5a38a93548eb8453a498222c4ff2"),
	"TUSD":                 common.HexToAddress("0x525389752ffe6487d33EF53FBcD4E5D3AD7937a0"),
	"DGX":                  common.HexToAddress("0x7d6D31326b12B6CBd7f054231D47CbcD16082b71"),
	"PAX":                  common.HexToAddress("0x3584087444dabf2e0d29284766142ac5c3a9a2b7"),
}

func DefaultAddressBook(network int64) AddressBook {
	switch network {
	case 1:
		return MainnetAddressBook
	case 3:
		return RopstenAddressBook
	case 42:
		return KovanAddressBook
	default:
		return AddressBook{}
	}
}
