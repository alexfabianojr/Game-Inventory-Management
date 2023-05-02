package walletBusiness

import "game-inventory-management/internal/domain/wallet"

func CalculateWalletBalance(walletEvents []wallet.WalletEventStore) int64 {
	balance := int64(0)
	for _, event := range walletEvents {
		switch event.Type {
		case wallet.Credit, wallet.TradeIn:
			balance += event.Value
		case wallet.Debit, wallet.TradeOut:
			balance -= event.Value
		}
	}
	return balance
}
