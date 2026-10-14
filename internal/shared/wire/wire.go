//go:build wireinject
// +build wireinject

package wire

import (
	"core-ledger-service/internal/domain_payment"
	"core-ledger-service/internal/domain_wallet"
	"core-ledger-service/internal/domain_wallet/presentation/messaging"
	"core-ledger-service/internal/shared/router"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeCreateWalletConsumer(db *gorm.DB, cfg messaging.KafkaConsumerConfig) (*messaging.CreateWalletUserConsumer, error) {
	wire.Build(
		domain_wallet.WalletModuleSet,
	)
	return &messaging.CreateWalletUserConsumer{}, nil
}

func InitializeRouter(db *gorm.DB) (*router.RouterGroup, error) {
	wire.Build(
		domain_wallet.WalletModuleSet,
		domain_payment.PaymentModuleSet,
		router.RouterSet,
	)
	return &router.RouterGroup{}, nil
}
