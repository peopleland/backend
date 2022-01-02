package model

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func DropMintRecordCollection(fc *f.FaunaClient) {
	_, err := fc.Query(f.Delete(f.Collection(MintRecordCollectionName)))
	if err != nil {
		panic(err)
	}
}

func CreateMintRecordMeta(fc *f.FaunaClient) {
	_, err := fc.Query(
		f.CreateCollection(
			f.Obj{"name": MintRecordCollectionName},
		),
	)
	if err != nil {
		panic(err)
	}

	_, err = fc.Query(
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

func DropOpenerRecordCollection(fc *f.FaunaClient) {
	_, err := fc.Query(f.Delete(f.Collection(OpenerRecordCollectionName)))
	if err != nil {
		panic(err)
	}
}

func CreateOpenerRecordMeta(fc *f.FaunaClient) {
	_, err := fc.Query(
		f.CreateCollection(
			f.Obj{"name": OpenerRecordCollectionName},
		),
	)
	if err != nil {
		panic(err)
	}

	_, err = fc.Query(
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

	_, err = fc.Query(
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
