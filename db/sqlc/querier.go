// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	DeleteAccount(ctx context.Context, id int64) (Account, error)
	ReadAccount(ctx context.Context, id int64) (Account, error)
	ReadAccountForUpdate(ctx context.Context, id int64) (Account, error)
	ReadAllEntries(ctx context.Context, arg ReadAllEntriesParams) ([]Entry, error)
	ReadAllTransfers(ctx context.Context, arg ReadAllTransfersParams) ([]Transfer, error)
	ReadEntry(ctx context.Context, id int64) (Entry, error)
	ReadTransfer(ctx context.Context, id int64) (Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateBalance(ctx context.Context, arg UpdateBalanceParams) (Account, error)
}

var _ Querier = (*Queries)(nil)