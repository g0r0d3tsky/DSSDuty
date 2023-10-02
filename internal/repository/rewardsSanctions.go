package repository

import (
	"context"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/domain"
	"github.com/google/uuid"
)

func (rw rw) CreateRewardsSanctions(ctx context.Context, rewSan *domain.RewardsSanctions) error {
	if rewSan == nil {
		return nil
	}
	if _, err := rw.store.Exec(ctx,
		`INSERT INTO "RewardsSanctions" (id, user_id, rewards, sanctions, info, timestamp)
 				VALUES ($1,$2, $3, $4, $5, $6)`,
		rewSan.Id, rewSan.UserId, rewSan.Rewards, rewSan.Sanctions, rewSan.Info, rewSan.DateTime); err != nil {
		return err
	}
	return nil
}
func (rw rw) GetRewardSanctionsByUserId(ctx context.Context, userId uuid.UUID) ([]*domain.RewardsSanctions, error) {
	var rewSans []*domain.RewardsSanctions

	rows, err := rw.store.Query(
		ctx,
		`SELECT * FROM "RewardsSanctions" WHERE user_id=$1`,
		userId)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rewSan := &domain.RewardsSanctions{}

		if err := rows.Scan(&rewSan.Id, &rewSan.UserId, &rewSan.Rewards, &rewSan.Sanctions, &rewSan.Info,
			&rewSan.DateTime); err != nil {
			return nil, err
		}

		rewSans = append(rewSans, rewSan)
	}

	return rewSans, nil
}
func (rw rw) UpdateRewardsSanctionsByUserId(ctx context.Context, rewSan *domain.RewardsSanctions) error {
	if _, err := rw.store.Exec(
		ctx,
		`UPDATE "RewardsSanctions" SET user_id=$2, rewards=$3, sanctions=$4, info=$5, timestamp=$6 WHERE id=$1`,
		rewSan.Id, rewSan.UserId, rewSan.Rewards, rewSan.Sanctions, rewSan.Info,
	); err != nil {
		return err
	}
	return nil
}
func (rw rw) DeleteRewardsSanctions(ctx context.Context, rewSanId uuid.UUID) error {
	if _, err := rw.store.Exec(ctx,
		`DELETE FROM "RewardsSanctions" WHERE id=$1`,
		rewSanId); err != nil {
		return err
	}
	return nil
}
