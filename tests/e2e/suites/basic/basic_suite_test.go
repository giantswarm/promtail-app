package basic

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"

	"github.com/giantswarm/apptest-framework/pkg/config"
	"github.com/giantswarm/apptest-framework/pkg/state"
	"github.com/giantswarm/apptest-framework/pkg/suite"
	"github.com/giantswarm/clustertest/pkg/logger"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	isUpgrade       = false
	lokiIngressName = "loki-gateway"
	lokiNamespace   = "loki"
)

type LokiData struct {
	url      string
	username string
	password string
}

type PromtailSecret struct {
	Clients []Client `yaml:"clients" json:"clients"`
}

type Client struct {
	Url       string    `yaml:"url" json:"url"`
	BasicAuth BasicAuth `yaml:"basic_auth" json:"basic_auth"`
}

type BasicAuth struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

func TestBasic(t *testing.T) {
	suite.New(config.MustLoad("../../config.yaml")).
		// The namespace to install the app into within the workload cluster
		WithInstallNamespace("kube-system").
		// If this is an upgrade test or not.
		// If true, the suite will first install the latest released version of the app before upgrading to the test version
		WithIsUpgrade(isUpgrade).
		WithValuesFile("./values.yaml").
		BeforeInstall(func() {
			// Do any pre-install checks here (ensure the cluster has needed pre-reqs)
		}).
		BeforeUpgrade(func() {
			// Perform any checks between installing the latest released version
			// and upgrading it to the version to test
			// E.g. ensure that the initial install has completed and has settled before upgrading
		}).
		Tests(func() {
			// Check Loki is reachable
			It("should Loki reachable", func() {

				By("reaching Loki")
				Eventually(func() error {
					lokiData, err := getLokiDataFromSecret()
					Expect(err).NotTo(HaveOccurred())

					logger.Log("making request to Loki at %s", lokiData.url)
					httpClient := newHttpClientWithProxy()
					req, _ := http.NewRequest("POST", lokiData.url, bytes.NewBuffer([]byte(`{}`)))
					req.Header.Set("Content-Type", "application/json")
					req.SetBasicAuth(lokiData.username, lokiData.password)
					resp, err := httpClient.Do(req)
					if err != nil {
						return err
					}
					defer resp.Body.Close()

					if resp.StatusCode != http.StatusNoContent {
						return fmt.Errorf("unexpected status code %d", resp.StatusCode)
					}

					return nil
				}).Should(Succeed())
			})

			// Check number of promtail pods are equal to the number of nodes in the workload cluster
			It("should have number of promtail pods equal to the number of nodes in the workload cluster", func() {

				By("getting number of nodes in the workload cluster")
				wcClient, err := state.GetFramework().WC(state.GetCluster().Name)
				Expect(err).NotTo(HaveOccurred())

				nodes := corev1.NodeList{}
				err = wcClient.List(state.GetContext(), &nodes)
				Expect(err).NotTo(HaveOccurred())
				numNodes := len(nodes.Items)
				Expect(numNodes).NotTo(BeZero())

				By("getting number of promtail pods")
				promtailLabelSelector, err := labels.Parse(fmt.Sprintf("app.kubernetes.io/name=%s", state.GetApplication().AppName))
				Expect(err).NotTo(HaveOccurred())
				promtailPods := &corev1.PodList{}
				err = wcClient.List(state.GetContext(), promtailPods, &client.ListOptions{Namespace: state.GetApplication().InstallNamespace, LabelSelector: promtailLabelSelector})
				Expect(err).NotTo(HaveOccurred())
				numPromtailPods := len(promtailPods.Items)
				Expect(numPromtailPods).NotTo(BeZero())

				By("checking number of promtail pods are equal to the number of nodes in the workload cluster")
				Expect(numPromtailPods).To(Equal(numNodes))
			})
		}).
		Run(t, "Basic Test")
}

// getLokiDataFromSecret gets Loki data from the promtail secret
func getLokiDataFromSecret() (LokiData, error) {
	wcClient, err := state.GetFramework().WC(state.GetCluster().Name)
	if err != nil {
		logger.Log("Failed to get workload cluster client: %v", err)
		return LokiData{}, err
	}

	secret := &corev1.Secret{}
	err = wcClient.Get(state.GetContext(), types.NamespacedName{Name: state.GetApplication().AppName, Namespace: state.GetApplication().InstallNamespace}, secret)
	if err != nil {
		logger.Log("Failed to get secret: %v", err)
		return LokiData{}, err
	}

	var content PromtailSecret
	err = yaml.Unmarshal(secret.Data["promtail.yaml"], &content)
	if err != nil {
		logger.Log("Failed to unmarshal promtail.yaml key: %v", err)
		return LokiData{}, err
	}

	logger.Log("Promtail secret content: %v", content)
	return LokiData{
		url:      content.Clients[0].Url,
		username: content.Clients[0].BasicAuth.Username,
		password: content.Clients[0].BasicAuth.Password,
	}, nil
}

// newHttpClientWithProxy creates a new http client with proxy settings
func newHttpClientWithProxy() *http.Client {
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			dialer := &net.Dialer{
				Resolver: &net.Resolver{
					PreferGo: true,
					Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
						if os.Getenv("HTTP_PROXY") != "" {
							u, err := url.Parse(os.Getenv("HTTP_PROXY"))
							if err != nil {
								logger.Log("error parsing HTTP_PROXY as a URL %s", os.Getenv("HTTP_PROXY"))
							} else {
								if addr == u.Host {
									// always use coredns for proxy address resolution.
									var d net.Dialer
									return d.Dial(network, address)
								}
							}
						}
						d := net.Dialer{
							Timeout: time.Millisecond * time.Duration(10000),
						}
						return d.DialContext(ctx, "udp", "8.8.4.4:53")
					},
				},
			}
			return dialer.DialContext(ctx, network, addr)
		},
	}

	if os.Getenv("HTTP_PROXY") != "" {
		logger.Log("detected need to use PROXY as HTTP_PROXY env var was set to %s", os.Getenv("HTTP_PROXY"))
		transport.Proxy = http.ProxyFromEnvironment
	}

	httpClient := &http.Client{
		Transport: transport,
	}
	return httpClient
}
