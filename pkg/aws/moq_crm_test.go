// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package aws

import (
	"github.com/integr8ly/cluster-service/pkg/clusterservice"
	"sync"
)

var (
	lockClusterResourceManagerMockDeleteResourcesForCluster sync.RWMutex
	lockClusterResourceManagerMockGetName                   sync.RWMutex
)

// Ensure, that ClusterResourceManagerMock does implement ClusterResourceManager.
// If this is not the case, regenerate this file with moq.
var _ ClusterResourceManager = &ClusterResourceManagerMock{}

// ClusterResourceManagerMock is a mock implementation of ClusterResourceManager.
//
//     func TestSomethingThatUsesClusterResourceManager(t *testing.T) {
//
//         // make and configure a mocked ClusterResourceManager
//         mockedClusterResourceManager := &ClusterResourceManagerMock{
//             DeleteResourcesForClusterFunc: func(clusterId string, tags map[string]string, dryRun bool) ([]*clusterservice.ReportItem, error) {
// 	               panic("mock out the DeleteResourcesForCluster method")
//             },
//             GetNameFunc: func() string {
// 	               panic("mock out the GetName method")
//             },
//         }
//
//         // use mockedClusterResourceManager in code that requires ClusterResourceManager
//         // and then make assertions.
//
//     }
type ClusterResourceManagerMock struct {
	// DeleteResourcesForClusterFunc mocks the DeleteResourcesForCluster method.
	DeleteResourcesForClusterFunc func(clusterId string, tags map[string]string, dryRun bool) ([]*clusterservice.ReportItem, error)

	// GetNameFunc mocks the GetName method.
	GetNameFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// DeleteResourcesForCluster holds details about calls to the DeleteResourcesForCluster method.
		DeleteResourcesForCluster []struct {
			// ClusterId is the clusterId argument value.
			ClusterId string
			// Tags is the tags argument value.
			Tags map[string]string
			// DryRun is the dryRun argument value.
			DryRun bool
		}
		// GetName holds details about calls to the GetName method.
		GetName []struct {
		}
	}
}

// DeleteResourcesForCluster calls DeleteResourcesForClusterFunc.
func (mock *ClusterResourceManagerMock) DeleteResourcesForCluster(clusterId string, tags map[string]string, dryRun bool) ([]*clusterservice.ReportItem, error) {
	if mock.DeleteResourcesForClusterFunc == nil {
		panic("ClusterResourceManagerMock.DeleteResourcesForClusterFunc: method is nil but ClusterResourceManager.DeleteResourcesForCluster was just called")
	}
	callInfo := struct {
		ClusterId string
		Tags      map[string]string
		DryRun    bool
	}{
		ClusterId: clusterId,
		Tags:      tags,
		DryRun:    dryRun,
	}
	lockClusterResourceManagerMockDeleteResourcesForCluster.Lock()
	mock.calls.DeleteResourcesForCluster = append(mock.calls.DeleteResourcesForCluster, callInfo)
	lockClusterResourceManagerMockDeleteResourcesForCluster.Unlock()
	return mock.DeleteResourcesForClusterFunc(clusterId, tags, dryRun)
}

// DeleteResourcesForClusterCalls gets all the calls that were made to DeleteResourcesForCluster.
// Check the length with:
//     len(mockedClusterResourceManager.DeleteResourcesForClusterCalls())
func (mock *ClusterResourceManagerMock) DeleteResourcesForClusterCalls() []struct {
	ClusterId string
	Tags      map[string]string
	DryRun    bool
} {
	var calls []struct {
		ClusterId string
		Tags      map[string]string
		DryRun    bool
	}
	lockClusterResourceManagerMockDeleteResourcesForCluster.RLock()
	calls = mock.calls.DeleteResourcesForCluster
	lockClusterResourceManagerMockDeleteResourcesForCluster.RUnlock()
	return calls
}

// GetName calls GetNameFunc.
func (mock *ClusterResourceManagerMock) GetName() string {
	if mock.GetNameFunc == nil {
		panic("ClusterResourceManagerMock.GetNameFunc: method is nil but ClusterResourceManager.GetName was just called")
	}
	callInfo := struct {
	}{}
	lockClusterResourceManagerMockGetName.Lock()
	mock.calls.GetName = append(mock.calls.GetName, callInfo)
	lockClusterResourceManagerMockGetName.Unlock()
	return mock.GetNameFunc()
}

// GetNameCalls gets all the calls that were made to GetName.
// Check the length with:
//     len(mockedClusterResourceManager.GetNameCalls())
func (mock *ClusterResourceManagerMock) GetNameCalls() []struct {
} {
	var calls []struct {
	}
	lockClusterResourceManagerMockGetName.RLock()
	calls = mock.calls.GetName
	lockClusterResourceManagerMockGetName.RUnlock()
	return calls
}