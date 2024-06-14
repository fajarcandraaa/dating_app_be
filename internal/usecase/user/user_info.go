package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/fajarcandraaa/dating_app_be/internal/dto"
	userEntity "github.com/fajarcandraaa/dating_app_be/internal/entity/user"
	userPresentation "github.com/fajarcandraaa/dating_app_be/internal/presentation/user"
	"github.com/fajarcandraaa/dating_app_be/internal/repository"
	"github.com/go-redis/redis/v8"
)

type userInfoUsecase struct {
	repo *repository.Repositories
	rds  *redis.Client
}

func NewUserInfoUsecase(repo *repository.Repositories, rds *redis.Client) *userInfoUsecase {
	return &userInfoUsecase{repo, rds}
}

var _ UserInfoUsecaseContract = &userInfoUsecase{}

// GetRandomProfile implements UserInfoUsecaseContract.
func (u *userInfoUsecase) GetRandomProfile(ctx context.Context, userCode string) (*userPresentation.User, error) {
	// TODO :
	// 1. Have to check package_id from redis, ✔
	// 2. From that, we will have 2 choice about the package, it's premium or not ✔
	// 3. If non-Premi :
	// 	- Check existing key <user_id>_view_quota 
	// 	- If exist, calculate reamaining data base on that key (max total = 10) ✔
	// 	- Check <user_id>_view in redis, then the data from that key will be used to exeption params when doing query to get random profile ✔
	// 	- give the profile feedback, then store user_code from the profile that has been review on <user_id>_view in redis ✔
	// 	- update quota data <user_id>_view_quota ✔
	// 4. If Premium :
	// 	- Check the feature_code, it have 2 posibility : no_quota_view & verified_label ✔
	// 	- If feature_code = no_quota_view, check existing key <user_id>_view in redis, 
	// 	then the data from that key will be used to exeption params when doing query to get random profile ✔
	// 	- If feature_code = verified_label, It have same behaviour with non-premi package ✔
	var (
		temp []string
		userInfo        userPresentation.User
		codeExceptions  = []string{}
		expiration      = 1 * 24 * time.Hour
		rdsKey          = fmt.Sprintf("login_%s", userCode)
		rdsKeyView      = fmt.Sprintf("%s_view", userCode)
		rdsKeyViewQuota = fmt.Sprintf("%s_view_Quota", userCode)
		remainingQuota  = 0
		maxDailyQuota   = 10
	)

	res, err := u.rds.Get(ctx, rdsKey).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(res), &userInfo)
	if err != nil {
		return nil, err
	}

	switch userInfo.PackageId { // check subsctiption witch package_id
	// NOTE : 
	// 1. base on data in databse with data seeder, non-premi have id -> 1, and premium have id -> 2
	case 2:
		// NOTE :
		// Feature code for No Qota View -> "NQA"
		// Feature code for Verified Label -> "VL"
		if userInfo.FeatureCode == "NQA" { // check feature package here
			hasView, errRedisView := u.rds.Get(ctx, rdsKeyView).Result()
			if errRedisView != redis.Nil {
				err = json.Unmarshal([]byte(hasView), &temp)
				if err != nil {
					return nil, err
				}
				codeExceptions = append(codeExceptions, temp...)
				res, err := u.getAndStore(ctx, userInfo.Gender, rdsKeyView, userCode, codeExceptions)
				if err != nil {
					return nil, err
				}
				return res, nil
			} else {
				res, err := u.getAndStore(ctx, userInfo.Gender, rdsKeyView, userCode, codeExceptions)
				if err != nil {
					return nil, err
				}
				return res, nil
			}
		} else {
			quotaView, errRedisQuota := u.rds.Get(ctx, rdsKeyViewQuota).Result() // cek existing daily quota view

			if errRedisQuota != redis.Nil {
				existingQuota, _ := strconv.Atoi(quotaView)
				if existingQuota <= maxDailyQuota {
					hasView, err := u.rds.Get(ctx, rdsKeyView).Result()
					if err != nil {
						return nil, err
					}
					err = json.Unmarshal([]byte(hasView), &temp)
					if err != nil {
						return nil, err
					}

					if hasView != "" { // if history view is exist on redis
						codeExceptions = append(codeExceptions, temp...)
						res, err := u.getAndStore(ctx, userInfo.Gender, rdsKeyView, userCode, codeExceptions)
						if err != nil {
							return nil, err
						}
						go func() {
							existingQuota -= 1
							err = u.rds.Set(context.Background(), rdsKeyViewQuota, remainingQuota, expiration).Err()
							if err != nil {
								log.Printf("ERROR Redis Set : %s", err.Error())
							}
						}()
						return res, nil
					} else { // if there is no history view on redis
						res, err := u.getAndStore(ctx, userInfo.Gender, rdsKeyView, userCode, codeExceptions)
						if err != nil {
							return nil, err
						}
						go func() {
							existingQuota -= 1
							err = u.rds.Set(context.Background(), rdsKeyViewQuota, remainingQuota, expiration).Err()
							if err != nil {
								log.Printf("ERROR Redis Set : %s", err.Error())
							}
						}()
						return res, nil
					}
				} else {
					return nil, fmt.Errorf("out of daily limitation to view")
				}
			} else {
				remainingQuota = maxDailyQuota // set ramaining quota that will be used to view other profile
				hasView, errRedisView := u.rds.Get(ctx, rdsKeyView).Result()

				if errRedisView != redis.Nil { // if history view is exist on redis
					err = json.Unmarshal([]byte(hasView), &temp)
					if err != nil {
						return nil, err
					}
					codeExceptions = append(codeExceptions, temp...)
					res, err := u.getAndStore(ctx, userInfo.Gender, rdsKeyView, userCode, codeExceptions)
					if err != nil {
						return nil, err
					}
					go func() {
						remainingQuota -= 1
						err = u.rds.Set(context.Background(), rdsKeyViewQuota, remainingQuota, expiration).Err()
						if err != nil {
							log.Printf("ERROR Redis Set : %s", err.Error())
						}
					}()
					return res, nil
				} else { // if there is no history view on redis
					res, err := u.getAndStore(ctx, userInfo.Gender, rdsKeyView, userCode, codeExceptions)
					if err != nil {
						return nil, err
					}
					go func() {
						remainingQuota -= 1
						err = u.rds.Set(context.Background(), rdsKeyViewQuota, remainingQuota, expiration).Err()
						if err != nil {
							log.Printf("ERROR Redis Set : %s", err.Error())
						}
					}()
					return res, nil
				}
			}
		}
	default:
		quotaView, errRedisQuota := u.rds.Get(ctx, rdsKeyViewQuota).Result() // cek existing daily quota view

		if errRedisQuota == redis.Nil {
			remainingQuota = maxDailyQuota // set ramaining quota that will be used as daily quota to view other profile
			hasView, errRedisView := u.rds.Get(ctx, rdsKeyView).Result()

			// if history view is exist on redis
			if errRedisView != redis.Nil {
				err = json.Unmarshal([]byte(hasView), &temp)
				if err != nil {
					return nil, err
				}
				codeExceptions = append(codeExceptions, temp...)
				res, err := u.getAndStore(ctx, userInfo.Gender, rdsKeyView, userCode, codeExceptions)
				if err != nil {
					return nil, err
				}
				go func() {
					remainingQuota -= 1
					dataStore, err := json.Marshal(remainingQuota)
					if err != nil {
						log.Printf("ERROR Redis Set Remaining Quota : %s", err.Error())
					}
					err = u.rds.Set(context.Background(), rdsKeyViewQuota, dataStore, expiration).Err()
					if err != nil {
						log.Printf("ERROR Redis Set : %s", err.Error())
					}
				}()
				return res, nil
			} else { // if there is no history view on redis
				res, err := u.getAndStore(ctx, userInfo.Gender, rdsKeyView, userCode, codeExceptions)
				if err != nil {
					return nil, err
				}
				go func() {
					remainingQuota -= 1
					dataStore, err := json.Marshal(remainingQuota)
					if err != nil {
						log.Printf("ERROR Redis Set Remaining Quota : %s", err.Error())
					}
					err = u.rds.Set(context.Background(), rdsKeyViewQuota, dataStore, expiration).Err()
					if err != nil {
						log.Printf("ERROR Redis Set : %s", err.Error())
					}
				}()
				return res, nil
			}
		} else {
			existingQuota, _ := strconv.Atoi(quotaView)
			if existingQuota <= maxDailyQuota {
				hasView, err := u.rds.Get(ctx, rdsKeyView).Result()
				if err != nil {
					return nil, err
				}
				err = json.Unmarshal([]byte(hasView), &temp)
				if err != nil {
					return nil, err
				}

				if hasView != "" { 
					codeExceptions = append(codeExceptions, temp...)
					res, err := u.getAndStore(ctx, userInfo.Gender, rdsKeyView, userCode, codeExceptions)
					if err != nil {
						return nil, err
					}
					go func() {
						existingQuota -= 1
						dataStore, err := json.Marshal(existingQuota)
						if err != nil {
							log.Printf("ERROR Redis Set Remaining Quota : %s", err.Error())
						}
						err = u.rds.Set(context.Background(), rdsKeyViewQuota, dataStore, expiration).Err()
						if err != nil {
							log.Printf("ERROR Redis Set : %s", err.Error())
						}
					}()
					return res, nil
				} else {
					res, err := u.getAndStore(ctx, userInfo.Gender, rdsKeyView, userCode, codeExceptions)
					if err != nil {
						return nil, err
					}
					go func() {
						existingQuota -= 1
						dataStore, err := json.Marshal(remainingQuota)
						if err != nil {
							log.Printf("ERROR Redis Set Remaining Quota : %s", err.Error())
						}
						err = u.rds.Set(context.Background(), rdsKeyViewQuota, dataStore, expiration).Err()
						if err != nil {
							log.Printf("ERROR Redis Set Set Remaining Quota : %s", err.Error())
						}
					}()
					return res, nil
				}
			} else {
				return nil, fmt.Errorf("out of daily limitation to view")
			}
		}
	}

}

func (u *userInfoUsecase) getAndStore(ctx context.Context, gender userEntity.UserGender, key, userCode string, codeExceptions []string) (*userPresentation.User, error) {
	var (
		expiration = 1 * 24 * time.Hour
	)

	randData, err := u.repo.User.GetRandom(ctx, gender, userCode, codeExceptions)
	if err != nil {
		return nil, err
	}

	codeExceptions = append(codeExceptions, *randData.UserCode)
	dataStore, err := json.Marshal(codeExceptions)
	if err != nil {
		log.Printf("ERROR Redis Set reviewing profile : %s", err.Error())
	}
	err = u.rds.Set(context.Background(), key, dataStore, expiration).Err()
	if err != nil {
		return nil, err
	}

	dtoTransform := dto.TransformEntityToPresentation(randData)
	return &dtoTransform, nil
}
