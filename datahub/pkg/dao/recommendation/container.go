package recommendation

import (
	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
)

// ContainerOperation defines container measurement operation of recommendation database
type ContainerOperation interface {
	AddPodRecommendations([]*datahub_v1alpha1.PodRecommendation) error
	ListPodRecommendations(podNamespacedName *datahub_v1alpha1.NamespacedName, timeRange *datahub_v1alpha1.TimeRange) ([]*datahub_v1alpha1.PodRecommendation, error)
}
