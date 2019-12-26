// Copyright 2018 The UniChain Team Authors
// This file is part of the unichain project.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package asset

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/unichainplatform/unichain/common"
	"github.com/unichainplatform/unichain/rawdb"
	"github.com/unichainplatform/unichain/state"
)

var assetDB = getStateDB()

var ast = getAsset()

func getStateDB() *state.StateDB {
	db := rawdb.NewMemoryDatabase()
	trieDB := state.NewDatabase(db)
	stateDB, err := state.New(common.Hash{}, trieDB)
	if err != nil {
		//t.Fatal("test getStateDB failure ", err)
		return nil
	}
	return stateDB
}
func getAsset() *Asset {
	return NewAsset(assetDB)
}

func TestAsset_InitAssetCount(t *testing.T) {
	type fields struct {
		sdb *state.StateDB
	}
	db := rawdb.NewMemoryDatabase()
	trieDB := state.NewDatabase(db)
	stateDB, err := state.New(common.Hash{}, trieDB)
	if err != nil {
		//t.Fatal("test getStateDB failure ", err)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		//
		{"init", fields{stateDB}},
	}
	for _, tt := range tests {
		a := &Asset{
			sdb: tt.fields.sdb,
		}
		a.InitAssetCount()
	}
	ast1 := NewAsset(stateDB)
	num, _ := ast1.getAssetCount()
	if num != 0 {
		t.Errorf("InitAssetCount err")
	}
}

func TestNewAsset(t *testing.T) {
	type args struct {
		sdb *state.StateDB
	}

	tests := []struct {
		name string
		args args
		want *Asset
	}{
		//
		//{"newnil", args{nil}, nil},
		{"new", args{assetDB}, ast},
	}
	for _, tt := range tests {
		if got := NewAsset(tt.args.sdb); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewAsset() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
func TestAsset_GetAssetObjectByName(t *testing.T) {
	type fields struct {
		sdb *state.StateDB
	}
	type args struct {
		assetName string
	}

	ao, _ := NewAssetObject("uni", 0, "zz", big.NewInt(1000), 10, common.Name(""), common.Name("a123456789aeee"), big.NewInt(9999999999), common.Name(""), "")
	//ao.SetAssetID(0)
	ast.addNewAssetObject(ao)
	ao1, _ := NewAssetObject("uni2", 0, "zz2", big.NewInt(1000), 10, common.Name(""), common.Name("a123456789aeee"), big.NewInt(9999999999), common.Name(""), "")
	//ao1.SetAssetID(1)
	ast.addNewAssetObject(ao1)
	ao2, _ := NewAssetObject("uni0", 0, "zz0", big.NewInt(1000), 0, common.Name(""), common.Name("a123456789aeee"), big.NewInt(9999999999), common.Name(""), "")
	//ao1.SetAssetID(2)
	ast.addNewAssetObject(ao2)
	ao3, _ := NewAssetObject("unic", 0, "zzc", big.NewInt(1000), 0, common.Name(""), common.Name("a123456789aeee"), big.NewInt(9999999999), common.Name("a123456789aeee"), "")
	//ao3.SetAssetID(3)
	ast.addNewAssetObject(ao3)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AssetObject
		wantErr bool
	}{
		// TODO: Add test cases.
		{"getall", fields{assetDB}, args{"uni"}, ao, false},
		{"getall2", fields{assetDB}, args{"uni2"}, ao1, false},
		{"getall3", fields{assetDB}, args{"uni0"}, ao2, false},
		{"getall4", fields{assetDB}, args{"unic"}, ao3, false},
	}
	for _, tt := range tests {
		a := &Asset{
			sdb: tt.fields.sdb,
		}
		got, err := a.GetAssetObjectByName(tt.args.assetName)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Asset.GetAssetObjectByName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Asset.GetAssetObjectByName() = %v, want %v", tt.name, got, tt.want)
		}
		t.Logf("GetAssetObjectByName asset dec=%v", got.Decimals)
	}
}

func TestAsset_addNewAssetObject(t *testing.T) {
	type fields struct {
		sdb *state.StateDB
	}
	type args struct {
		ao *AssetObject
	}

	ao3, _ := NewAssetObject("uni3", 0, "zz3", big.NewInt(1000), 10, common.Name(""), common.Name("a123456789aeee"), big.NewInt(9999999999), common.Name(""), "")
	//ao1.SetAssetID(3)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
		{"addnil", fields{assetDB}, args{nil}, 0, true},
		{"add", fields{assetDB}, args{ao3}, 4, false},
	}
	for _, tt := range tests {
		a := &Asset{
			sdb: tt.fields.sdb,
		}
		got, err := a.addNewAssetObject(tt.args.ao)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Asset.addNewAssetObject() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Asset.addNewAssetObject() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestAsset_GetAssetIDByName(t *testing.T) {
	type fields struct {
		sdb *state.StateDB
	}
	type args struct {
		assetName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		//
		{"normal", fields{assetDB}, args{""}, 0, true},
		{"normal", fields{assetDB}, args{"uni"}, 0, false},
		{"wrong", fields{assetDB}, args{"uni2"}, 1, false},
	}
	for _, tt := range tests {
		a := &Asset{
			sdb: tt.fields.sdb,
		}
		got, err := a.GetAssetIDByName(tt.args.assetName)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Asset.GetAssetIDByName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Asset.GetAssetIDByName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestAsset_GetAssetObjectByID(t *testing.T) {
	type fields struct {
		sdb *state.StateDB
	}
	type args struct {
		id uint64
	}

	ao, _ := NewAssetObject("uni", 0, "zz", big.NewInt(1000), 10, common.Name(""), common.Name("a123456789aeee"), big.NewInt(9999999999), common.Name(""), "")
	ao.SetAssetID(0)
	ast.IssueAssetObject(ao)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AssetObject
		wantErr bool
	}{
		//
		{"assetnotexist", fields{assetDB}, args{222}, nil, true},
		{"normal2", fields{assetDB}, args{0}, ao, false},
	}
	for _, tt := range tests {
		a := &Asset{
			sdb: tt.fields.sdb,
		}
		got, err := a.GetAssetObjectByID(tt.args.id)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Asset.GetAssetObjectByID() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Asset.GetAssetObjectByID() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestAsset_getAssetCount(t *testing.T) {
	type fields struct {
		sdb *state.StateDB
	}

	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
		{"get", fields{assetDB}, 5, false},
	}
	for _, tt := range tests {
		a := &Asset{
			sdb: tt.fields.sdb,
		}
		got, err := a.getAssetCount()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Asset.getAssetCount() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Asset.getAssetCount() = %v, want %v", tt.name, got, tt.want)
		}
	}
	ao, _ := NewAssetObject("uni2", 0, "zz2", big.NewInt(1000), 10, common.Name(""), common.Name("a123456789aeee"), big.NewInt(9999999999), common.Name(""), "")
	//ao.SetAssetID(1)
	ast.IssueAssetObject(ao)
	num, err := ast.getAssetCount()
	if err != nil {
		t.Errorf("get asset count err")
	}
	if num != 5 {
		t.Errorf("test asset count err")
	}
}

// func TestAsset_GetAllAssetObject(t *testing.T) {
// 	type fields struct {
// 		sdb *state.StateDB
// 	}
// 	aslice := make([]*AssetObject, 0)
// 	ao, _ := ast.GetAssetObjectByID(1)
// 	aslice = append(aslice, ao)
// 	ao, _ = ast.GetAssetObjectByID(2)
// 	aslice = append(aslice, ao)
// 	ao, _ = ast.GetAssetObjectByID(3)
// 	aslice = append(aslice, ao)

// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    []*AssetObject
// 		wantErr bool
// 	}{
// 		//
// 		//{"getall", fields{assetDB}, aslice, false},
// 	}
// 	for _, tt := range tests {
// 		a := &Asset{
// 			sdb: tt.fields.sdb,
// 		}
// 		got, err := a.GetAllAssetObject()
// 		if (err != nil) != tt.wantErr {
// 			t.Errorf("%q. Asset.GetAllAssetObject() error = %v, wantErr %v", tt.name, err, tt.wantErr)
// 			continue
// 		}
// 		if !reflect.DeepEqual(got, tt.want) {
// 			t.Errorf("%q. Asset.GetAllAssetObject() = %v, want %v", tt.name, got, tt.want)
// 		}
// 	}
// }

func TestAsset_SetAssetObject(t *testing.T) {
	type fields struct {
		sdb *state.StateDB
	}
	type args struct {
		ao *AssetObject
	}

	ao4, _ := NewAssetObject("uni4", 0, "zz4", big.NewInt(1000), 10, common.Name(""), common.Name("a123456789aeee"), big.NewInt(9999999999), common.Name(""), "")
	ao4.SetAssetID(54)
	ao5, _ := NewAssetObject("uni5", 0, "zz5", big.NewInt(1000), 10, common.Name(""), common.Name("a123456789aeee"), big.NewInt(9999999999), common.Name(""), "")
	ao5.SetAssetID(55)
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"setnil", fields{assetDB}, args{nil}, true},
		{"add", fields{assetDB}, args{ao4}, false},
		{"add2", fields{assetDB}, args{ao5}, false},
	}
	for _, tt := range tests {
		a := &Asset{
			sdb: tt.fields.sdb,
		}
		if err := a.SetAssetObject(tt.args.ao); (err != nil) != tt.wantErr {
			t.Errorf("%q. Asset.SetAssetObject() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestAsset_IssueAssetObject(t *testing.T) {
	type fields struct {
		sdb *state.StateDB
	}
	type args struct {
		ao *AssetObject
	}
	ao6, _ := NewAssetObject("uni6", 0, "zz6", big.NewInt(1000), 10, common.Name(""), common.Name("a123456789aeee"), big.NewInt(9999999999), common.Name(""), "")
	ao6.SetAssetID(11)
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"nil", fields{assetDB}, args{nil}, true},
		{"add", fields{assetDB}, args{ao6}, false},
	}
	for _, tt := range tests {
		a := &Asset{
			sdb: tt.fields.sdb,
		}
		if _, err := a.IssueAssetObject(tt.args.ao); (err != nil) != tt.wantErr {
			t.Errorf("%q. Asset.IssueAssetObject() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestAsset_IssueAsset(t *testing.T) {
	type fields struct {
		sdb *state.StateDB
	}
	type args struct {
		assetName string
		symbol    string
		amount    *big.Int
		dec       uint64
		founder   common.Name
		owner     common.Name
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"nilname", fields{assetDB}, args{"", "z", big.NewInt(1), 2, common.Name(""), common.Name("11")}, true},
		{"nilsym", fields{assetDB}, args{"22", "", big.NewInt(2), 2, common.Name(""), common.Name("11")}, true},
		{"exist", fields{assetDB}, args{"uni", "3", big.NewInt(2), 2, common.Name(""), common.Name("11")}, true},
		{"normal", fields{assetDB}, args{"uni22", "23", big.NewInt(2), 2, common.Name(""), common.Name("a112345698")}, true},
	}
	for _, tt := range tests {
		a := &Asset{
			sdb: tt.fields.sdb,
		}
		if _, err := a.IssueAsset(tt.args.assetName, 0, 0, tt.args.symbol, tt.args.amount, tt.args.dec, tt.args.founder, tt.args.owner, big.NewInt(9999999999), common.Name(""), ""); (err != nil) != tt.wantErr {
			t.Errorf("%q. Asset.IssueAsset() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestAsset_IncreaseAsset(t *testing.T) {
	type fields struct {
		sdb *state.StateDB
	}
	type args struct {
		accountName common.Name
		AssetID     uint64
		amount      *big.Int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"nilName", fields{assetDB}, args{common.Name(""), 1, big.NewInt(2)}, true},
		{"wrongID", fields{assetDB}, args{common.Name("11"), 0, big.NewInt(2)}, false},
		{"wrongAmount", fields{assetDB}, args{common.Name("11"), 0, big.NewInt(-2)}, true},
		{"normal", fields{assetDB}, args{common.Name("a123456789aeee"), 1, big.NewInt(50)}, false},
	}
	for _, tt := range tests {
		a := &Asset{
			sdb: tt.fields.sdb,
		}
		if err := a.IncreaseAsset(tt.args.accountName, tt.args.AssetID, tt.args.amount, 4); (err != nil) != tt.wantErr {
			t.Errorf("%q. Asset.IncreaseAsset() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestAsset_SetAssetNewOwner(t *testing.T) {
	type fields struct {
		sdb *state.StateDB
	}
	type args struct {
		accountName common.Name
		AssetID     uint64
		newOwner    common.Name
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases
		{"nilName", fields{assetDB}, args{common.Name(""), 1, common.Name("")}, true},
		{"wrongID", fields{assetDB}, args{common.Name("11"), 0, common.Name("")}, false},
		{"wrongAmount", fields{assetDB}, args{common.Name("11"), 123, common.Name("")}, true},
		{"normal", fields{assetDB}, args{common.Name("a123456789aeee"), 1, common.Name("a123456789afff")}, false},
	}
	for _, tt := range tests {
		a := &Asset{
			sdb: tt.fields.sdb,
		}
		if err := a.SetAssetNewOwner(tt.args.accountName, tt.args.AssetID, tt.args.newOwner); (err != nil) != tt.wantErr {
			t.Errorf("%q. Asset.SetAssetNewOwner() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestAsset_UpdateAsset(t *testing.T) {
	type fields struct {
		sdb *state.StateDB
	}
	type args struct {
		accountName common.Name
		AssetID     uint64
		Owner       common.Name
		founder     common.Name
		forkID      uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases
		{"nilname", fields{assetDB}, args{common.Name(""), 1, common.Name(""), common.Name(""), 0}, true},
		{"wrongAssetID", fields{assetDB}, args{common.Name("11"), 0, common.Name(""), common.Name(""), 0}, false},
		{"wrongamount", fields{assetDB}, args{common.Name("11"), 123, common.Name(""), common.Name(""), 0}, true},
		{"nilfounder", fields{assetDB}, args{common.Name("a123456789afff"), 1, common.Name("a123456789aeee"), common.Name(""), 0}, false},
		{"nilfounder", fields{assetDB}, args{common.Name("a123456789afff"), 1, common.Name("a123456789aeee"), common.Name(""), 4}, false},
		{"nilfounder", fields{assetDB}, args{common.Name("a123456789afff"), 1, common.Name("a123456789aeee"), common.Name("a123456789afff"), 4}, false},
		{"normal", fields{assetDB}, args{common.Name("a123456789afff"), 1, common.Name("a123456789afff"), common.Name("a123456789afff"), 0}, false},
	}
	for _, tt := range tests {
		a := &Asset{
			sdb: tt.fields.sdb,
		}
		if err := a.UpdateAsset(tt.args.accountName, tt.args.AssetID, tt.args.founder, tt.args.forkID); (err != nil) != tt.wantErr {
			t.Errorf("%q. Asset.updateAsset() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
func TestAsset_HasAccess(t *testing.T) {
	type fields struct {
		sdb *state.StateDB
	}
	type args struct {
		AssetID uint64
		name    common.Name
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases
		{"a0", fields{assetDB}, args{0, common.Name("")}, true},
		{"a1", fields{assetDB}, args{1, common.Name("")}, true},
		{"a2", fields{assetDB}, args{2, common.Name("")}, true},
		{"a3", fields{assetDB}, args{3, common.Name("a123456789aeee")}, true},
		{"a3_1", fields{assetDB}, args{3, common.Name("a123456789afff")}, false},
	}
	for _, tt := range tests {
		a := &Asset{
			sdb: tt.fields.sdb,
		}
		if has := a.HasAccess(tt.args.AssetID, tt.args.name); has != tt.wantErr {
			t.Errorf("%q. Asset.HasAccess() error = %v, wantErr %v", tt.name, has, tt.wantErr)
		}
	}
}
