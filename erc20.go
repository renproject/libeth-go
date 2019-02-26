package libeth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type erc20 struct {
	account *account
	cerc20  *CompatibleERC20
}

type ERC20 interface {
	BalanceOf(ctx context.Context, who string) (*big.Int, error)
	Allowance(ctx context.Context, owner, spender string) (*big.Int, error)
	Transfer(ctx context.Context, to string, amount *big.Int, speed TxExecutionSpeed, sendAll bool) (*types.Transaction, error)
	Approve(ctx context.Context, spender string, amount *big.Int, speed TxExecutionSpeed) (*types.Transaction, error)
	TransferFrom(ctx context.Context, from, to string, amount *big.Int, speed TxExecutionSpeed) (*types.Transaction, error)
}

func (account *account) NewERC20(addressOrAlias string) (ERC20, error) {
	address, ok := account.addressBook[addressOrAlias]
	if !ok {
		address = common.HexToAddress(addressOrAlias)
	}
	compatibleERC20, err := NewCompatibleERC20(address, bind.ContractBackend(account.EthClient()))
	if err != nil {
		return nil, err
	}
	return &erc20{
		account: account,
		cerc20:  compatibleERC20,
	}, nil
}

func (erc20 *erc20) BalanceOf(ctx context.Context, who string) (*big.Int, error) {
	client := erc20.account.Client()
	whoAddress, err := client.Resolve(who)
	if err != nil {
		return nil, err
	}

	var balance *big.Int
	return balance, client.Get(ctx, func() error {
		bal, err := erc20.cerc20.BalanceOf(&bind.CallOpts{}, whoAddress)
		if err != nil {
			return err
		}
		balance = bal
		return nil
	})
}

func (erc20 *erc20) Allowance(ctx context.Context, owner, spender string) (*big.Int, error) {
	client := erc20.account.Client()
	ownerAddress, err := client.Resolve(owner)
	if err != nil {
		return nil, err
	}

	spenderAddress, err := client.Resolve(spender)
	if err != nil {
		return nil, err
	}

	var allowance *big.Int
	return allowance, client.Get(ctx, func() error {
		alw, err := erc20.cerc20.Allowance(&bind.CallOpts{}, ownerAddress, spenderAddress)
		if err != nil {
			return err
		}
		allowance = alw
		return nil
	})
}

func (erc20 *erc20) Transfer(ctx context.Context, to string, amount *big.Int, speed TxExecutionSpeed, sendAll bool) (*types.Transaction, error) {
	if sendAll {
		balance, err := erc20.BalanceOf(ctx, erc20.account.Address().String())
		if err != nil {
			return nil, err
		}
		amount = balance
	}

	client := erc20.account.Client()
	toAddress, err := client.Resolve(to)
	if err != nil {
		return nil, err
	}

	return erc20.account.Transact(
		ctx,
		speed,
		nil,
		func(tops *bind.TransactOpts) (*types.Transaction, error) {
			tx, err := erc20.cerc20.Transfer(tops, toAddress, amount)
			if err != nil {
				return tx, err
			}
			return tx, nil
		},
		nil,
		1,
	)
}

func (erc20 *erc20) Approve(ctx context.Context, spender string, amount *big.Int, speed TxExecutionSpeed) (*types.Transaction, error) {
	client := erc20.account.Client()
	spenderAddress, err := client.Resolve(spender)
	if err != nil {
		return nil, err
	}

	return erc20.account.Transact(
		ctx,
		speed,
		nil,
		func(tops *bind.TransactOpts) (*types.Transaction, error) {
			tx, err := erc20.cerc20.Approve(tops, spenderAddress, amount)
			if err != nil {
				return tx, err
			}
			return tx, nil
		},
		nil,
		1,
	)
}

func (erc20 *erc20) TransferFrom(ctx context.Context, from, to string, amount *big.Int, speed TxExecutionSpeed) (*types.Transaction, error) {
	client := erc20.account.Client()
	fromAddress, err := client.Resolve(from)
	if err != nil {
		return nil, err
	}
	toAddress, err := client.Resolve(to)
	if err != nil {
		return nil, err
	}
	return erc20.account.Transact(
		ctx,
		speed,
		nil,
		func(tops *bind.TransactOpts) (*types.Transaction, error) {
			tx, err := erc20.cerc20.TransferFrom(tops, fromAddress, toAddress, amount)
			if err != nil {
				return tx, err
			}
			return tx, nil
		},
		nil,
		1,
	)
}
