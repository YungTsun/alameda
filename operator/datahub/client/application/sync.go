package application

import (
	"context"
	"time"

	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/pkg/apis/autoscaling/v1alpha1"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"fmt"

	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
)

func SyncWithDatahub(client client.Client, conn *grpc.ClientConn) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	applicationList := autoscalingv1alpha1.AlamedaScalerList{}
	if err := client.List(ctx, nil, &applicationList); err != nil {
		return errors.Errorf(
			"Sync applications with datahub failed due to list applications from cluster failed: %s", err.Error())
	}
	datahubApplicationRepo := NewApplicationRepository(conn)
	if len(applicationList.Items) > 0 {
		if err := datahubApplicationRepo.CreateApplications(applicationList.Items); err != nil {
			return fmt.Errorf(
				"Sync applications with datahub failed due to register application failed: %s", err.Error())
		}
	}

	// Clean up unexisting applications from Datahub
	existingApplicationMap := make(map[string]bool)
	for _, application := range applicationList.Items {
		existingApplicationMap[fmt.Sprintf("%s/%s",
			application.GetNamespace(), application.GetName())] = true
	}

	applicationsFromDatahub, err := datahubApplicationRepo.ListApplications()
	if err != nil {
		return fmt.Errorf(
			"Sync applications with datahub failed due to list applications from datahub failed: %s",
			err.Error())
	}
	applicationsNeedDeleting := make([]*datahub_resources.Application, 0)
	for _, n := range applicationsFromDatahub {
		if _, exist := existingApplicationMap[fmt.Sprintf("%s/%s",
			n.GetObjectMeta().GetNamespace(), n.GetObjectMeta().GetName())]; exist {
			continue
		}
		applicationsNeedDeleting = append(applicationsNeedDeleting, n)
	}
	if len(applicationsNeedDeleting) > 0 {
		err = datahubApplicationRepo.DeleteApplications(applicationsNeedDeleting)
		if err != nil {
			return errors.Wrap(err, "delete applications from Datahub failed")
		}
	}

	return nil
}
