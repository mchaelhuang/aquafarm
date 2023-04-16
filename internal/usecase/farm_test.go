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

func TestNewFarmUC(t *testing.T) {
	f := NewFarmUC(nil, nil)
	assert.NotNil(t, f)
}

func TestFarmUC_GetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		farmRepo internal.FarmRepo
	}
	type args struct {
		filter entity.FarmFilter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Farm
		wantErr error
	}{
		{
			name: "error from farmRepo.Get",
			fields: fields{
				farmRepo: func() *MockFarmRepo {
					mock := NewMockFarmRepo(ctrl)
					mock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, errors.New("error test"))
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
				farmRepo: func() *MockFarmRepo {
					mock := NewMockFarmRepo(ctrl)
					filter := entity.FarmFilter{Farm: entity.Farm{Name: "test"}}
					farms := []entity.Farm{{ID: 111, Name: "test"}}
					mock.EXPECT().Get(gomock.Any(), filter).Return(farms, nil)
					return mock
				}(),
			},
			args: args{
				filter: entity.FarmFilter{
					Farm: entity.Farm{Name: "test"},
				},
			},
			want:    []entity.Farm{{ID: 111, Name: "test"}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FarmUC{
				farmRepo: tt.fields.farmRepo,
			}
			got, err := f.GetList(context.Background(), tt.args.filter)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFarmUC_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		farmRepo internal.FarmRepo
		pondRepo internal.PondRepo
	}
	type args struct {
		filter entity.FarmFilter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Farm
		wantErr error
	}{
		{
			name: "error from farmRepo.Get",
			fields: fields{
				farmRepo: func() *MockFarmRepo {
					mock := NewMockFarmRepo(ctrl)
					mock.EXPECT().Get(gomock.Any(), gomock.Any()).
						Return(nil, errors.New("error test"))
					return mock
				}(),
			},
			args:    args{},
			want:    entity.Farm{},
			wantErr: errors.New("error test"),
		},
		{
			name: "error from pondRepo.Get",
			fields: fields{
				farmRepo: func() *MockFarmRepo {
					mock := NewMockFarmRepo(ctrl)
					farms := []entity.Farm{{ID: 10}}
					mock.EXPECT().Get(gomock.Any(), gomock.Any()).
						Return(farms, nil)
					return mock
				}(),
				pondRepo: func() *MockPondRepo {
					mock := NewMockPondRepo(ctrl)
					mock.EXPECT().Get(gomock.Any(), gomock.Any()).
						Return(nil, errors.New("error test"))
					return mock
				}(),
			},
			args:    args{},
			want:    entity.Farm{},
			wantErr: errors.New("error test"),
		},
		{
			name: "success",
			fields: fields{
				farmRepo: func() *MockFarmRepo {
					mock := NewMockFarmRepo(ctrl)
					filter := entity.FarmFilter{
						Farm: entity.Farm{Name: "Test"},
					}
					farms := []entity.Farm{{ID: 10}}
					mock.EXPECT().Get(gomock.Any(), filter).
						Return(farms, nil)
					return mock
				}(),
				pondRepo: func() *MockPondRepo {
					mock := NewMockPondRepo(ctrl)
					ponds := []entity.Pond{
						{ID: 200}, {ID: 201},
					}
					filter := entity.PondFilter{
						Pond: entity.Pond{FarmID: 10},
					}
					mock.EXPECT().Get(gomock.Any(), filter).
						Return(ponds, nil)
					return mock
				}(),
			},
			args: args{
				filter: entity.FarmFilter{
					Farm: entity.Farm{Name: "Test"},
				},
			},
			want: entity.Farm{
				ID: 10,
				Ponds: []entity.Pond{
					{ID: 200}, {ID: 201},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FarmUC{
				farmRepo: tt.fields.farmRepo,
				pondRepo: tt.fields.pondRepo,
			}
			got, err := f.Get(context.Background(), tt.args.filter)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFarmUC_Store(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		farmRepo internal.FarmRepo
	}
	type args struct {
		farm entity.Farm
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr error
	}{
		{
			name: "error from farmRepo.Store",
			fields: fields{
				farmRepo: func() *MockFarmRepo {
					mock := NewMockFarmRepo(ctrl)
					mock.EXPECT().Store(gomock.Any(), gomock.Any()).Return(0, errors.New("error test"))
					return mock
				}(),
			},
			args:    args{},
			want:    0,
			wantErr: errors.New("error test"),
		},
		{
			name: "success",
			fields: fields{
				farmRepo: func() *MockFarmRepo {
					mock := NewMockFarmRepo(ctrl)
					farm := entity.Farm{Name: "test"}
					mock.EXPECT().Store(gomock.Any(), farm).Return(1, nil)
					return mock
				}(),
			},
			args: args{
				farm: entity.Farm{Name: "test"},
			},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FarmUC{
				farmRepo: tt.fields.farmRepo,
			}
			got, err := f.Store(context.Background(), tt.args.farm)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFarmUC_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		farmRepo internal.FarmRepo
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "error from farmRepo.Delete",
			fields: fields{
				farmRepo: func() *MockFarmRepo {
					mock := NewMockFarmRepo(ctrl)
					mock.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(errors.New("error test"))
					return mock
				}(),
			},
			args:    args{},
			wantErr: errors.New("error test"),
		},
		{
			name: "success",
			fields: fields{
				farmRepo: func() *MockFarmRepo {
					mock := NewMockFarmRepo(ctrl)
					mock.EXPECT().Delete(gomock.Any(), 10).Return(nil)
					return mock
				}(),
			},
			args: args{
				id: 10,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FarmUC{
				farmRepo: tt.fields.farmRepo,
			}
			err := f.Delete(context.Background(), tt.args.id)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
