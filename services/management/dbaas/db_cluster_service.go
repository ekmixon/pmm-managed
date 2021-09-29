package dbaas

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dbaascontrollerv1beta1 "github.com/percona-platform/dbaas-api/gen/controller"
	dbaasv1beta1 "github.com/percona/pmm/api/managementpb/dbaas"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"
)

type DBClusterService struct {
	db               *reform.DB
	l                *logrus.Entry
	controllerClient dbaasClient
	grafanaClient    grafanaClient
}

// NewDBClusterService creates DB Clusters Service.
func NewDBClusterService(db *reform.DB, client dbaasClient, grafanaClient grafanaClient) dbaasv1beta1.DBClusterServer {
	l := logrus.WithField("component", "dbaas_db_cluster")
	return &DBClusterService{db: db, l: l, controllerClient: client, grafanaClient: grafanaClient}
}

// ListDBClusters returns a list of all DB clusters.
func (s DBClusterService) ListDBClusters(ctx context.Context, req *dbaasv1beta1.ListDBClustersRequest) (*dbaasv1beta1.ListDBClustersResponse, error) {
	kubernetesCluster, err := models.FindKubernetesClusterByName(s.db.Querier, req.KubernetesClusterName)
	if err != nil {
		return nil, err
	}

	pxcClusters, err := s.listPXCClusters(ctx, kubernetesCluster.KubeConfig)
	if err != nil {
		return nil, err
	}

	psmdbClusters, err := s.listPSMDBClusters(ctx, kubernetesCluster.KubeConfig)
	if err != nil {
		return nil, err
	}

	return &dbaasv1beta1.ListDBClustersResponse{
		PxcClusters:   pxcClusters,
		PsmdbClusters: psmdbClusters,
	}, nil
}

func (s DBClusterService) listPSMDBClusters(ctx context.Context, kubeConfig string) ([]*dbaasv1beta1.PSMDBCluster, error) {
	in := dbaascontrollerv1beta1.ListPSMDBClustersRequest{
		KubeAuth: &dbaascontrollerv1beta1.KubeAuth{
			Kubeconfig: kubeConfig,
		},
	}

	out, err := s.controllerClient.ListPSMDBClusters(ctx, &in)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can't get list of PSMDB clusters: %s", err.Error())
	}

	clusters := make([]*dbaasv1beta1.PSMDBCluster, len(out.Clusters))
	for i, c := range out.Clusters {
		var computeResources *dbaasv1beta1.ComputeResources
		var diskSize int64
		if c.Params.Replicaset != nil {
			diskSize = c.Params.Replicaset.DiskSize
			if c.Params.Replicaset.ComputeResources != nil {
				computeResources = &dbaasv1beta1.ComputeResources{
					CpuM:        c.Params.Replicaset.ComputeResources.CpuM,
					MemoryBytes: c.Params.Replicaset.ComputeResources.MemoryBytes,
				}
			}
		}

		cluster := dbaasv1beta1.PSMDBCluster{
			Name: c.Name,
			Params: &dbaasv1beta1.PSMDBClusterParams{
				ClusterSize: c.Params.ClusterSize,
				Replicaset: &dbaasv1beta1.PSMDBClusterParams_ReplicaSet{
					ComputeResources: computeResources,
					DiskSize:         diskSize,
				},
			},
			State: psmdbStates()[c.State],
			Operation: &dbaasv1beta1.RunningOperation{
				TotalSteps:    c.Operation.TotalSteps,
				FinishedSteps: c.Operation.FinishedSteps,
				Message:       c.Operation.Message,
			},
			Exposed: c.Exposed,
		}

		clusters[i] = &cluster
	}

	return clusters, nil
}

func (s DBClusterService) listPXCClusters(ctx context.Context, kubeConfig string) ([]*dbaasv1beta1.PXCCluster, error) {
	in := dbaascontrollerv1beta1.ListPXCClustersRequest{
		KubeAuth: &dbaascontrollerv1beta1.KubeAuth{
			Kubeconfig: kubeConfig,
		},
	}

	out, err := s.controllerClient.ListPXCClusters(ctx, &in)
	if err != nil {
		return nil, err
	}

	pxcClusters := make([]*dbaasv1beta1.PXCCluster, len(out.Clusters))
	for i, c := range out.Clusters {
		cluster := dbaasv1beta1.PXCCluster{
			Name: c.Name,
			Params: &dbaasv1beta1.PXCClusterParams{
				ClusterSize: c.Params.ClusterSize,
			},
			State: pxcStates()[c.State],
			Operation: &dbaasv1beta1.RunningOperation{
				TotalSteps:    c.Operation.TotalSteps,
				FinishedSteps: c.Operation.FinishedSteps,
				Message:       c.Operation.Message,
			},
			Exposed: c.Exposed,
		}

		if c.Params.Pxc != nil {
			cluster.Params.Pxc = &dbaasv1beta1.PXCClusterParams_PXC{
				DiskSize: c.Params.Pxc.DiskSize,
			}
			if c.Params.Pxc.ComputeResources != nil {
				cluster.Params.Pxc.ComputeResources = &dbaasv1beta1.ComputeResources{
					CpuM:        c.Params.Pxc.ComputeResources.CpuM,
					MemoryBytes: c.Params.Pxc.ComputeResources.MemoryBytes,
				}
			}
		}

		if c.Params.Haproxy != nil {
			if c.Params.Haproxy.ComputeResources != nil {
				cluster.Params.Haproxy = &dbaasv1beta1.PXCClusterParams_HAProxy{
					ComputeResources: &dbaasv1beta1.ComputeResources{
						CpuM:        c.Params.Haproxy.ComputeResources.CpuM,
						MemoryBytes: c.Params.Haproxy.ComputeResources.MemoryBytes,
					},
				}
			}
		} else if c.Params.Proxysql != nil {
			if c.Params.Proxysql.ComputeResources != nil {
				cluster.Params.Proxysql = &dbaasv1beta1.PXCClusterParams_ProxySQL{
					DiskSize: c.Params.Proxysql.DiskSize,
					ComputeResources: &dbaasv1beta1.ComputeResources{
						CpuM:        c.Params.Proxysql.ComputeResources.CpuM,
						MemoryBytes: c.Params.Proxysql.ComputeResources.MemoryBytes,
					},
				}
			}
		}

		pxcClusters[i] = &cluster
	}
	return pxcClusters, nil
}

// DeleteDBCluster deletes DB cluster by given name and type.
func (s DBClusterService) DeleteDBCluster(ctx context.Context, req *dbaasv1beta1.DeleteDBClusterRequest) (*dbaasv1beta1.DeleteDBClusterResponse, error) {
	kubernetesCluster, err := models.FindKubernetesClusterByName(s.db.Querier, req.KubernetesClusterName)
	if err != nil {
		return nil, err
	}

	var clusterType string
	switch req.ClusterType {
	case dbaasv1beta1.DBClusterType_DB_CLUSTER_TYPE_PXC:
		in := dbaascontrollerv1beta1.DeletePXCClusterRequest{
			Name: req.Name,
			KubeAuth: &dbaascontrollerv1beta1.KubeAuth{
				Kubeconfig: kubernetesCluster.KubeConfig,
			},
		}

		_, err = s.controllerClient.DeletePXCCluster(ctx, &in)
		if err != nil {
			return nil, err
		}
		clusterType = "pxc"
	case dbaasv1beta1.DBClusterType_DB_CLUSTER_TYPE_PSMDB:
		in := dbaascontrollerv1beta1.DeletePSMDBClusterRequest{
			Name: req.Name,
			KubeAuth: &dbaascontrollerv1beta1.KubeAuth{
				Kubeconfig: kubernetesCluster.KubeConfig,
			},
		}

		_, err = s.controllerClient.DeletePSMDBCluster(ctx, &in)
		if err != nil {
			return nil, err
		}
		clusterType = "psmdb"
	default:
		return nil, status.Error(codes.InvalidArgument, "unexpected DB cluster type")
	}

	err = s.grafanaClient.DeleteAPIKeysWithPrefix(ctx, fmt.Sprintf("%s-%s-%s", clusterType, req.KubernetesClusterName, req.Name))
	if err != nil {
		// ignore if API Key is not deleted.
		s.l.Warnf("Couldn't delete API key: %s", err)
	}

	return &dbaasv1beta1.DeleteDBClusterResponse{}, nil
}
