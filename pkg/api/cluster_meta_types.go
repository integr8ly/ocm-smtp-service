package api

import "github.com/jinzhu/gorm"

type ClusterMeta struct {
	Meta
	ClusterID string
}

type ClusterList []*ClusterMeta
type ClusterIndex map[string]*ClusterMeta

func (l ClusterList) Index() ClusterIndex {
	index := ClusterIndex{}
	for _, o := range l {
		index[o.ID] = o
	}
	return index
}

func (org *ClusterMeta) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", NewID())
}
