package repositories
import (
	"database/sql"
	"context"
	)
func BeginTransaction(runnersRepository *RunnersRepository,
	resultsRepository *ResultsRepository) error {
	ctx := context.Background()
	transaction, err := resultsRepository.dbHandler.
	BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
	return err
	}
	runnersRepository.transaction = transaction
	resultsRepository.transaction = transaction
	return nil
	}
