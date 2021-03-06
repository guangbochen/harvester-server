package user

import (
	"context"

	"github.com/rancher/harvester-server/pkg/config"
)

const (
	userRbacControllerAgentName = "user-rbac-controller"
)

func Register(ctx context.Context, management *config.Management) error {
	users := management.HarvesterFactory.Harvester().V1alpha1().User()

	userRBACController := &userRBACHandler{
		users:                   users,
		clusterRoleBindings:     management.RbacFactory.Rbac().V1().ClusterRoleBinding(),
		clusterRoleBindingCache: management.RbacFactory.Rbac().V1().ClusterRoleBinding().Cache(),
	}

	users.OnChange(ctx, userRbacControllerAgentName, userRBACController.OnChanged)
	return nil
}
