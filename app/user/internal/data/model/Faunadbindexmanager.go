package model

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

type FaunaMetaManager struct {
	fc *f.FaunaClient
}

func (manager *FaunaMetaManager) createCollectionByName(name string) {
	_, err := manager.fc.Query(
		f.CreateCollection(
			f.Obj{"name": name},
		),
	)
	if err != nil {
		panic(err)
	}
}

func (manager *FaunaMetaManager) dropCollectionByName(name string) {
	_, err := manager.fc.Query(f.Delete(f.Collection(name)))
	if err != nil {
		panic(err)
	}
}

func (manager *FaunaMetaManager) DropMintRecordCollection() {
	manager.dropCollectionByName(MintRecordCollectionName)
}

func (manager *FaunaMetaManager) DropOpenerRecordCollection() {
	manager.dropCollectionByName(OpenerRecordCollectionName)
}

func (manager *FaunaMetaManager) DropOpenerGameRoundInfoCollection() {
	manager.dropCollectionByName(OpenerGameRoundInfoCollectionName)
}

// CreateMintRecordMeta mint_records
func (manager *FaunaMetaManager) CreateMintRecordMeta() {
	manager.createCollectionByName(MintRecordCollectionName)

	_, err := manager.fc.Query(
		f.CreateIndex(f.Obj{
			"name":   MintRecordsByMintAddressAndXAndYSortTsDescIndex,
			"source": f.Collection(MintRecordCollectionName),
			"terms": f.Arr{
				f.Obj{"field": f.Arr{"data", "mint_address"}},
				f.Obj{"field": f.Arr{"data", "x"}},
				f.Obj{"field": f.Arr{"data", "y"}},
			},
			"values": f.Arr{
				f.Obj{"field": "ts", "reverse": true},
				f.Obj{"field": "ref"},
			},
		}))
	if err != nil {
		panic(err)
	}
}

func (manager *FaunaMetaManager) CreateOpenerRecordMeta() {
	manager.createCollectionByName(OpenerRecordCollectionName)

	_, err := manager.fc.Query(
		f.CreateIndex(f.Obj{
			"name":   OpenerRecordByTokenId,
			"unique": true,
			"source": f.Collection(OpenerRecordCollectionName),
			"terms": f.Arr{
				f.Obj{"field": f.Arr{"data", "token_id"}},
			},
		}))
	if err != nil {
		panic(err)
	}

	_, err = manager.fc.Query(
		f.CreateIndex(f.Obj{
			"name":   OpenerRecordSortTokenIdDescIndex,
			"source": f.Collection(OpenerRecordCollectionName),
			"values": f.Arr{
				f.Obj{"field": f.Arr{"data", "token_id"}, "reverse": true},
				f.Obj{"field": "ref"},
			},
		}))
	if err != nil {
		panic(err)
	}
}

func (manager *FaunaMetaManager) CreateOpenerGameRoundInfoMeta() {
	manager.createCollectionByName(OpenerGameRoundInfoCollectionName)

	_, err := manager.fc.Query(
		f.CreateIndex(f.Obj{
			"name":   OpenerGameRoundInfoByRoundNumberIndex,
			"unique": true,
			"source": f.Collection(OpenerGameRoundInfoCollectionName),
			"terms": f.Arr{
				f.Obj{"field": f.Arr{"data", "round_number"}},
			},
		}))
	if err != nil {
		panic(err)
	}
}
