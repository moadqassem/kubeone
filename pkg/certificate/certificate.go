/*
Copyright 2019 The KubeOne Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package certificate

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"k8c.io/kubeone/pkg/configupload"
	"k8c.io/kubeone/pkg/templates/resources"

	certutil "k8s.io/client-go/util/cert"
	"k8s.io/client-go/util/keyutil"
)

// CAKeyPair parses generated PKI CA certificate and key
func CAKeyPair(config *configupload.Configuration) (*rsa.PrivateKey, *x509.Certificate, error) {
	caCert, found := config.KubernetesPKI[KubernetesCACertPath]
	if !found {
		return nil, nil, fmt.Errorf("%q not found", KubernetesCACertPath)
	}

	caKey, found := config.KubernetesPKI[KubernetesCAKeyPath]
	if !found {
		return nil, nil, fmt.Errorf("%q not found", KubernetesCAKeyPath)
	}

	certs, err := certutil.ParseCertsPEM(caCert)
	if err != nil {
		return nil, nil, err
	}

	if len(certs) == 0 {
		return nil, nil, errors.New("ca.crt does not contain at least one valid certificate")
	}

	possibleKey, err := keyutil.ParsePrivateKeyPEM(caKey)
	if err != nil {
		return nil, nil, err
	}

	rsaKey, ok := possibleKey.(*rsa.PrivateKey)
	if !ok {
		return nil, nil, errors.New("private key is not a RSA private key")
	}

	return rsaKey, certs[0], nil
}

func NewSignedTLSCert(name, namespace, domain string, caKey crypto.Signer, caCert *x509.Certificate) (map[string]string, error) {
	serviceCommonName := strings.Join([]string{name, namespace, "svc"}, ".")
	serviceFQDNCommonName := strings.Join([]string{serviceCommonName, domain, ""}, ".")

	altdnsNames := []string{
		serviceFQDNCommonName,
		serviceCommonName,
	}

	newKPKey, err := newPrivateKey()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate private key")
	}

	certCfg := certutil.Config{
		AltNames: certutil.AltNames{
			DNSNames: altdnsNames,
		},
		CommonName: serviceCommonName,
		Usages:     []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	newKPCert, err := newSignedCert(&certCfg, newKPKey, caCert, caKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate certificate")
	}

	return map[string]string{
		resources.TLSCertName:          string(encodeCertPEM(newKPCert)),
		resources.TLSKeyName:           string(encodePrivateKeyPEM(newKPKey)),
		resources.KubernetesCACertName: string(encodeCertPEM(caCert)),
	}, nil
}

// GetCertificateSANs combines host name and subject alternative names into a list of SANs after transformation
func GetCertificateSANs(host string, alternativeNames []string) []string {
	certSANS := []string{strings.ToLower(host)}
	for _, name := range alternativeNames {
		certSANS = append(certSANS, strings.ToLower(name))
	}
	return certSANS
}
