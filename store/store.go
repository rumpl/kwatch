package store

import "github.com/hashicorp/go-memdb"

// Deployment model
type Deployment struct {
	Name  string
	Image string
}

// Store persists deployments
type Store interface {
	Insert(deployment *Deployment) error
	List() ([]*Deployment, error)
}

type store struct {
	db *memdb.MemDB
}

// New creates a new store
func New() (Store, error) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"deployments": {
				Name: "deployments",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	return &store{
		db: db,
	}, nil
}

func (s *store) Insert(deployment *Deployment) error {
	txn := s.db.Txn(true)
	if err := txn.Insert("deployments", deployment); err != nil {
		return err
	}
	// Such an odd api, you the commit doesn't fail...
	txn.Commit()

	return nil
}

func (s *store) List() ([]*Deployment, error) {
	txn := s.db.Txn(false)

	it, err := txn.Get("deployments", "id")
	if err != nil {
		return []*Deployment{}, err
	}

	deps := []*Deployment{}
	for d := it.Next(); d != nil; d = it.Next() {
		de := d.(*Deployment)
		deps = append(deps, de)
	}

	return deps, nil
}
