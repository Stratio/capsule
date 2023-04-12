package utils

import (
	"context"
	"strings"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	capsulev1beta1 "github.com/clastix/capsule/api/v1beta1"
	"github.com/clastix/capsule/pkg/utils"
)

func IsCapsuleUser(ctx context.Context, req admission.Request, clt client.Client, userGroups []string, capsuleUserName string) bool {
	groupList := utils.NewUserGroupList(req.UserInfo.Groups)
	// if the user is a ServiceAccount belonging to the kube-system namespace, definitely, it's not a Capsule user
	// and we can skip the check in case of Capsule user group assigned to system:authenticated
	// (ref: https://github.com/clastix/capsule/issues/234)
	if groupList.Find("system:serviceaccounts:kube-system") {
		return false
	}
	// nolint:nestif
	if req.UserInfo.Username == capsuleUserName {
		return true
	}
	// nolint:nestif
	if sets.NewString(req.UserInfo.Groups...).Has("system:serviceaccounts") {
		parts := strings.Split(req.UserInfo.Username, ":")

		targetNamespace := parts[2]

		if len(targetNamespace) > 0 {
			tl := &capsulev1beta1.TenantList{}
			if err := clt.List(ctx, tl, client.MatchingFieldsSelector{Selector: fields.OneTermEqualSelector(".status.namespaces", targetNamespace)}); err != nil {
				return false
			}

			if len(tl.Items) == 1 {
				return true
			}
		}
	}

	for _, group := range userGroups {
		if groupList.Find(group) {
			return true
		}
	}

	return false
}
