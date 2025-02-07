package topology

import (
	"reflect"

	k8scorev1 "k8s.io/api/core/v1"
	apimachinerytypes "k8s.io/apimachinery/pkg/types"
)

func serviceConforms( //nolint: gocyclo
	existingService,
	renderedService *k8scorev1.Service,
	expectedOwnerUID apimachinerytypes.UID,
) bool {
	if !reflect.DeepEqual(existingService.Spec.Selector, renderedService.Spec.Selector) {
		// bad selector somehow
		return false
	}

	if !reflect.DeepEqual(existingService.Spec.Type, renderedService.Spec.Type) {
		// somehow bad type
		return false
	}

	for _, expectedPort := range renderedService.Spec.Ports {
		var expectedPortExists bool

		for _, actualPort := range existingService.Spec.Ports {
			if expectedPort.Name != actualPort.Name {
				continue
			}

			if expectedPort.Port != actualPort.Port {
				break
			}

			if !reflect.DeepEqual(expectedPort.TargetPort, actualPort.TargetPort) {
				break
			}

			expectedPortExists = true
		}

		if !expectedPortExists {
			// port doesnt exist or is wrong
			return false
		}
	}

	if existingService.ObjectMeta.Labels == nil {
		// obviously our labels don't exist, so we need to enforce that
		return false
	}

	for k, v := range renderedService.ObjectMeta.Labels {
		var expectedLabelExists bool

		for nk, nv := range existingService.ObjectMeta.Labels {
			if k == nk && v == nv {
				expectedLabelExists = true

				break
			}
		}

		if !expectedLabelExists {
			// missing some expected label, and/or value is wrong
			return false
		}
	}

	if len(existingService.ObjectMeta.OwnerReferences) != 1 {
		// we should have only one owner reference, the extractor
		return false
	}

	if existingService.ObjectMeta.OwnerReferences[0].UID != expectedOwnerUID {
		// owner ref uid is not us
		return false
	}

	return true
}
