package repository

import (
	"fmt"
	"testing"

	"github.com/alfisar/jastip-import/config"

	"github.com/alfisar/jastip-import/domain"

	"github.com/stretchr/testify/require"
)

var (
	poolData *domain.Config
	repo     = NewTravelSchRepository()
)

func TestInsert(t *testing.T) {
	config.Init()
	poolData = domain.DataPool.Get().(*domain.Config)
	data := domain.TravelSchRequest{
		UserID:      1,
		Location:    "Indonesia",
		PeriodStart: "2025-02-14",
		PeriodEnd:   "2025-02-20",
		Status:      1,
	}
	_, err := repo.Create(poolData.DBSql, data)
	require.Nil(t, err)

}

func TestGet(t *testing.T) {
	where := map[string]interface{}{
		"status": 1,
	}
	result, count, err := repo.GetList(poolData.DBSql, where, "", 0, 10)
	fmt.Println(result)
	fmt.Println(count)
	require.Nil(t, err)

}

func TestUpdate(t *testing.T) {
	where := map[string]interface{}{
		"id": 1,
	}

	updates := map[string]any{
		"locations": "Thailand",
	}
	err := repo.Update(poolData.DBSql, where, updates)
	require.Nil(t, err)

}
func TestGetFailed(t *testing.T) {
	where := map[string]interface{}{
		"status": 0,
	}
	result, _, _ := repo.GetList(poolData.DBSql, where, "", 0, 10)

	require.Empty(t, result)

}

func TestUpdateFailed(t *testing.T) {
	where := map[string]interface{}{
		"id": 100,
	}

	updates := map[string]any{
		"locations": "Thailand",
	}
	err := repo.Update(poolData.DBSql, where, updates)
	require.NotNil(t, err)

}

func TestDelete(t *testing.T) {
	where := map[string]interface{}{
		"id": 1,
	}
	err := repo.Delete(poolData.DBSql, where)
	require.Nil(t, err)
}

func TestInsertFailed(t *testing.T) {
	data := domain.TravelSchRequest{
		UserID:      10,
		Location:    "Indonesia",
		PeriodStart: "2025-14-02",
		PeriodEnd:   "2025-20-02",
		Status:      1,
	}
	poolData.DBSql = nil
	_, err := repo.Create(poolData.DBSql, data)
	require.NotNil(t, err)

}
func TestGetFailedConn(t *testing.T) {
	where := map[string]interface{}{
		"status": 1,
	}
	poolData.DBSql = nil
	_, _, err := repo.GetList(poolData.DBSql, where, "", 1, 10)
	require.NotNil(t, err)

}

func TestUpdateFailedConn(t *testing.T) {
	where := map[string]interface{}{
		"id": 1,
	}

	updates := map[string]any{
		"location": "Thailand",
	}
	poolData.DBSql = nil
	err := repo.Update(poolData.DBSql, where, updates)
	require.NotNil(t, err)

}

func TestDeleteFailedConn(t *testing.T) {
	poolData.DBSql = nil
	where := map[string]interface{}{
		"id": 1,
	}
	err := repo.Delete(poolData.DBSql, where)
	require.NotNil(t, err)
}
