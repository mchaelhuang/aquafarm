package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mchaelhuang/aquafarm/internal"
	"github.com/mchaelhuang/aquafarm/internal/constant"
	"github.com/mchaelhuang/aquafarm/internal/entity"
)

func TestNewPondUC(t *testing.T) {
	p := NewPondUC(nil, nil, nil)
	assert.NotNil(t, p)
}

func TestPondUC_GetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		pondRepo internal.PondRepo
	}
	type args struct {
		filter entity.PondFilter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Pond
		wantErr error
	}{
		{
			name: "error pondRepo.Get",
			fields: fields{
				pondRepo: func() *MockPondRepo {
					mock := NewMockPondRepo(ctrl)
					mock.EXPECT().Get(gomock.Any(), gomock.Any()).
						Return(nil, errors.New("error test"))
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
				pondRepo: func() *MockPondRepo {
					mock := NewMockPondRepo(ctrl)
					filter := entity.PondFilter{
						Pond: entity.Pond{Label: "Test"},
					}
					ponds := []entity.Pond{
						{ID: 100},
					}
					mock.EXPECT().Get(gomock.Any(), filter).
						Return(ponds, nil)
					return mock
				}(),
			},
			args: args{
				filter: entity.PondFilter{
					Pond: entity.Pond{Label: "Test"},
				},
			},
			want: []entity.Pond{
				{ID: 100},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &PondUC{
				pondRepo: tt.fields.pondRepo,
			}
			got, err := f.GetList(context.Background(), tt.args.filter)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPondUC_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		farmRepo internal.FarmRepo
		pondRepo internal.PondRepo
	}
	type args struct {
		filter entity.PondFilter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Pond
		wantErr error
	}{
		{
			name: "error pondRepo.Get",
			fields: fields{
				pondRepo: func() *MockPondRepo {
					mock := NewMockPondRepo(ctrl)
					mock.EXPECT().Get(gomock.Any(), gomock.Any()).
						Return(nil, errors.New("error test"))
					return mock
				}(),
			},
			args:    args{},
			want:    entity.Pond{},
			wantErr: errors.New("error test"),
		},
		{
			name: "success",
			fields: fields{
				pondRepo: func() *MockPondRepo {
					mock := NewMockPondRepo(ctrl)
					pond := []entity.Pond{{ID: 100}}
					mock.EXPECT().Get(gomock.Any(), gomock.Any()).
						Return(pond, nil)
					return mock
				}(),
			},
			args:    args{},
			want:    entity.Pond{ID: 100},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &PondUC{
				farmRepo: tt.fields.farmRepo,
				pondRepo: tt.fields.pondRepo,
			}
			got, err := f.Get(context.Background(), tt.args.filter)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPondUC_Store(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		farmRepo internal.FarmRepo
		pondRepo internal.PondRepo
	}
	type args struct {
		pond entity.Pond
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr error
	}{
		{
			name: "error not found from farmRepo.Get",
			fields: fields{
				farmRepo: func() *MockFarmRepo {
					mock := NewMockFarmRepo(ctrl)
					mock.EXPECT().Get(gomock.Any(), gomock.Any()).
						Return(nil, constant.ErrNotFound)
					return mock
				}(),
			},
			args:    args{},
			want:    0,
			wantErr: constant.ErrIncorrectFarmID,
		},
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
			want:    0,
			wantErr: errors.New("error test"),
		},
		{
			name: "success",
			fields: fields{
				farmRepo: func() *MockFarmRepo {
					mock := NewMockFarmRepo(ctrl)
					filter := entity.FarmFilter{
						Farm: entity.Farm{ID: 100},
					}
					mock.EXPECT().Get(gomock.Any(), filter).
						Return(nil, nil)
					return mock
				}(),
				pondRepo: func() *MockPondRepo {
					mock := NewMockPondRepo(ctrl)
					pond := entity.Pond{
						ID: 201, FarmID: 100, Label: "Test",
					}
					mock.EXPECT().Store(gomock.Any(), pond).
						Return(201, nil)
					return mock
				}(),
			},
			args: args{
				pond: entity.Pond{
					ID: 201, FarmID: 100, Label: "Test",
				},
			},
			want:    201,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &PondUC{
				farmRepo: tt.fields.farmRepo,
				pondRepo: tt.fields.pondRepo,
			}
			got, err := f.Store(context.Background(), tt.args.pond)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPondUC_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		pondRepo internal.PondRepo
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
			name: "error pondRepo.Delete",
			fields: fields{
				pondRepo: func() *MockPondRepo {
					mock := NewMockPondRepo(ctrl)
					mock.EXPECT().Delete(gomock.Any(), gomock.Any()).
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
				pondRepo: func() *MockPondRepo {
					mock := NewMockPondRepo(ctrl)
					mock.EXPECT().Delete(gomock.Any(), 100).
						Return(nil)
					return mock
				}(),
			},
			args: args{
				id: 100,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &PondUC{
				pondRepo: tt.fields.pondRepo,
			}
			err := f.Delete(context.Background(), tt.args.id)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
