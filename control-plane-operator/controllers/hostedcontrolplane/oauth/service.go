package oauth

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	routev1 "github.com/openshift/api/route/v1"
	hyperv1 "github.com/openshift/hypershift/api/v1alpha1"
	"github.com/openshift/hypershift/control-plane-operator/controllers/hostedcontrolplane/config"
	"github.com/openshift/hypershift/control-plane-operator/controllers/hostedcontrolplane/manifests"
	"github.com/openshift/hypershift/control-plane-operator/controllers/hostedcontrolplane/util"
)

const (
	OAuthServerPort   = 6443
	RouteExternalPort = 443
)

var (
	oauthServerLabels = map[string]string{
		"app":                         "oauth-openshift",
		hyperv1.ControlPlaneComponent: "oauth-openshift",
	}
)

func ReconcileService(svc *corev1.Service, ownerRef config.OwnerRef, strategy *hyperv1.ServicePublishingStrategy) error {
	ownerRef.ApplyTo(svc)
	svc.Spec.Selector = oauthServerLabels
	var portSpec corev1.ServicePort
	if len(svc.Spec.Ports) > 0 {
		portSpec = svc.Spec.Ports[0]
	} else {
		svc.Spec.Ports = []corev1.ServicePort{portSpec}
	}
	portSpec.Port = int32(OAuthServerPort)
	portSpec.Protocol = corev1.ProtocolTCP
	portSpec.TargetPort = intstr.FromInt(OAuthServerPort)
	switch strategy.Type {
	case hyperv1.NodePort:
		svc.Spec.Type = corev1.ServiceTypeNodePort
		if portSpec.NodePort == 0 && strategy.NodePort != nil {
			portSpec.NodePort = strategy.NodePort.Port
		}
	case hyperv1.Route:
		svc.Spec.Type = corev1.ServiceTypeClusterIP
	default:
		return fmt.Errorf("invalid publishing strategy for OAuth service: %s", strategy.Type)
	}
	svc.Spec.Ports[0] = portSpec
	return nil
}

func ReconcileServiceStatus(svc *corev1.Service, route *routev1.Route, strategy *hyperv1.ServicePublishingStrategy) (host string, port int32, err error) {
	switch strategy.Type {
	case hyperv1.Route:
		if route.Spec.Host == "" {
			return
		}
		port = RouteExternalPort
		host = route.Spec.Host
	case hyperv1.NodePort:
		if strategy.NodePort == nil {
			err = fmt.Errorf("strategy details not specified for OAuth nodeport type service")
			return
		}
		if len(svc.Spec.Ports) == 0 {
			return
		}
		if svc.Spec.Ports[0].NodePort == 0 {
			return
		}
		port = svc.Spec.Ports[0].NodePort
		host = strategy.NodePort.Address
	}
	return
}

func ReconcileOAuthServerCertWorkerManifest(cm *corev1.ConfigMap, ownerRef config.OwnerRef, oauthServerCert *corev1.Secret) error {
	ownerRef.ApplyTo(cm)
	certConfigMap := manifests.OAuthServerCert()
	if err := reconcileOAuthServerCertConfigMap(certConfigMap, oauthServerCert); err != nil {
		return err
	}
	return util.ReconcileWorkerManifest(cm, certConfigMap)
}

func reconcileOAuthServerCertConfigMap(certConfigMap *corev1.ConfigMap, sourceSecret *corev1.Secret) error {
	if _, hasCertKey := sourceSecret.Data[corev1.TLSCertKey]; !hasCertKey {
		return fmt.Errorf("source secret %s/%s does not have a cert key", sourceSecret.Namespace, sourceSecret.Name)
	}

	if certConfigMap.Data == nil {
		certConfigMap.Data = map[string]string{}
	}
	certConfigMap.Data["ca-bundle.crt"] = string(sourceSecret.Data[corev1.TLSCertKey])
	return nil
}
