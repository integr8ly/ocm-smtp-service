package presenters

import (
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/api"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/api/openapi"
)

func ConvertClusterMeta(clusterMeta openapi.ClusterMeta) *api.ClusterMeta {
	return &api.ClusterMeta{
		Meta: api.Meta{
			ID: clusterMeta.Id,
		},
		ClusterID: clusterMeta.ClusterID,
	}
}

func PresentClusterMeta(clusterMeta *api.ClusterMeta) openapi.ClusterMeta {
	reference := PresentReference(clusterMeta.ID, clusterMeta)
	return openapi.ClusterMeta{
		Id:        reference.Id,
		ClusterID: reference.Id,
	}
}
