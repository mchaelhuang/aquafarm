package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mchaelhuang/aquafarm/internal"
	"github.com/mchaelhuang/aquafarm/internal/entity"
)

func TestNewStatsUC(t *testing.T) {
	s := NewStatsUC(nil, nil)
	assert.NotNil(t, s)
}

func TestStatsUC_Add(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		statsRepo internal.StatsRepo
	}
	type args struct {
		info entity.StatsRequestInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "error IncrEndpoint",
			fields: fields{
				statsRepo: func() *MockStatsRepo {
					mock := NewMockStatsRepo(ctrl)
					mock.EXPECT().IncrEndpoint(gomock.Any(), gomock.Any()).
						Return(errors.New("error test"))
					return mock
				}(),
			},
			args:    args{},
			wantErr: errors.New("error test"),
		},
		{
			name: "error CollectUserAgent",
			fields: fields{
				statsRepo: func() *MockStatsRepo {
					mock := NewMockStatsRepo(ctrl)
					mock.EXPECT().IncrEndpoint(gomock.Any(), gomock.Any()).
						Return(nil)
					mock.EXPECT().CollectUserAgent(gomock.Any(), gomock.Any()).
						Return(errors.New("error test"))
					return mock
				}(),
			},
			args:    args{},
			wantErr: errors.New("error test"),
		},
		{
			name: "success",
			fields: fields{
				statsRepo: func() *MockStatsRepo {
					mock := NewMockStatsRepo(ctrl)
					info := entity.StatsRequestInfo{
						Method: "GET", Endpoint: "/v1/farm",
					}
					mock.EXPECT().IncrEndpoint(gomock.Any(), info).
						Return(nil)
					mock.EXPECT().CollectUserAgent(gomock.Any(), info).
						Return(nil)
					return mock
				}(),
			},
			args: args{
				info: entity.StatsRequestInfo{
					Method: "GET", Endpoint: "/v1/farm",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatsUC{
				statsRepo: tt.fields.statsRepo,
			}
			err := s.Add(context.Background(), tt.args.info)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestStatsUC_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		statsRepo internal.StatsRepo
	}
	type args struct {
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]entity.StatsResult
		wantErr error
	}{
		{
			name: "error from GetEndpointCount",
			fields: fields{
				statsRepo: func() *MockStatsRepo {
					mock := NewMockStatsRepo(ctrl)
					mock.EXPECT().GetEndpointCount(gomock.Any()).
						Return(nil, errors.New("error test"))
					return mock
				}(),
			},
			args:    args{},
			want:    nil,
			wantErr: errors.New("error test"),
		},
		{
			name: "error from GetUniqueAgentCount",
			fields: fields{
				statsRepo: func() *MockStatsRepo {
					mock := NewMockStatsRepo(ctrl)
					endpoints := map[string]int{
						"POST /v1/farm": 2,
					}
					mock.EXPECT().GetEndpointCount(gomock.Any()).
						Return(endpoints, nil)
					mock.EXPECT().GetUniqueAgentCount(gomock.Any(), gomock.Any()).
						Return(0, errors.New("error test"))
					return mock
				}(),
			},
			args:    args{},
			want:    nil,
			wantErr: errors.New("error test"),
		},
		{
			name: "success",
			fields: fields{
				statsRepo: func() *MockStatsRepo {
					mock := NewMockStatsRepo(ctrl)
					endpoints := map[string]int{
						"GET /v1/farm":  3,
						"POST /v1/farm": 2,
					}
					mock.EXPECT().GetEndpointCount(gomock.Any()).
						Return(endpoints, nil)
					mock.EXPECT().GetUniqueAgentCount(gomock.Any(), "GET /v1/farm").
						Return(4, nil)
					mock.EXPECT().GetUniqueAgentCount(gomock.Any(), "POST /v1/farm").
						Return(5, nil)
					return mock
				}(),
			},
			args: args{},
			want: map[string]entity.StatsResult{
				"GET /v1/farm":  {Count: 3, UniqueUserAgent: 4},
				"POST /v1/farm": {Count: 2, UniqueUserAgent: 5},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatsUC{
				statsRepo: tt.fields.statsRepo,
			}
			got, err := s.Get(context.Background())
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
