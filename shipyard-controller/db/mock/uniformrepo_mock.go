// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package db_mock

import (
	"github.com/keptn/keptn/shipyard-controller/models"
	"sync"
)

// UniformRepoMock is a mock implementation of db.UniformRepo.
//
// 	func TestSomethingThatUsesUniformRepo(t *testing.T) {
//
// 		// make and configure a mocked db.UniformRepo
// 		mockedUniformRepo := &UniformRepoMock{
// 			CreateOrUpdateSubscriptionFunc: func(integrationID string, subscription models.Subscription) error {
// 				panic("mock out the CreateOrUpdateSubscription method")
// 			},
// 			CreateOrUpdateUniformIntegrationFunc: func(integration models.Integration) error {
// 				panic("mock out the CreateOrUpdateUniformIntegration method")
// 			},
// 			CreateUniformIntegrationFunc: func(integration models.Integration) error {
// 				panic("mock out the CreateUniformIntegration method")
// 			},
// 			DeleteServiceFromSubscriptionsFunc: func(subscriptionName string) error {
// 				panic("mock out the DeleteServiceFromSubscriptions method")
// 			},
// 			DeleteSubscriptionFunc: func(integrationID string, subscriptionID string) error {
// 				panic("mock out the DeleteSubscription method")
// 			},
// 			DeleteUniformIntegrationFunc: func(id string) error {
// 				panic("mock out the DeleteUniformIntegration method")
// 			},
// 			GetSubscriptionFunc: func(integrationID string, subscriptionID string) (*models.Subscription, error) {
// 				panic("mock out the GetSubscription method")
// 			},
// 			GetSubscriptionsFunc: func(integrationID string) ([]models.Subscription, error) {
// 				panic("mock out the GetSubscriptions method")
// 			},
// 			GetUniformIntegrationsFunc: func(filter models.GetUniformIntegrationsParams) ([]models.Integration, error) {
// 				panic("mock out the GetUniformIntegrations method")
// 			},
// 			UpdateLastSeenFunc: func(integrationID string) (*models.Integration, error) {
// 				panic("mock out the UpdateLastSeen method")
// 			},
// 		}
//
// 		// use mockedUniformRepo in code that requires db.UniformRepo
// 		// and then make assertions.
//
// 	}
type UniformRepoMock struct {
	// CreateOrUpdateSubscriptionFunc mocks the CreateOrUpdateSubscription method.
	CreateOrUpdateSubscriptionFunc func(integrationID string, subscription models.Subscription) error

	// CreateOrUpdateUniformIntegrationFunc mocks the CreateOrUpdateUniformIntegration method.
	CreateOrUpdateUniformIntegrationFunc func(integration models.Integration) error

	// CreateUniformIntegrationFunc mocks the CreateUniformIntegration method.
	CreateUniformIntegrationFunc func(integration models.Integration) error

	// DeleteServiceFromSubscriptionsFunc mocks the DeleteServiceFromSubscriptions method.
	DeleteServiceFromSubscriptionsFunc func(subscriptionName string) error

	// DeleteSubscriptionFunc mocks the DeleteSubscription method.
	DeleteSubscriptionFunc func(integrationID string, subscriptionID string) error

	// DeleteUniformIntegrationFunc mocks the DeleteUniformIntegration method.
	DeleteUniformIntegrationFunc func(id string) error

	// GetSubscriptionFunc mocks the GetSubscription method.
	GetSubscriptionFunc func(integrationID string, subscriptionID string) (*models.Subscription, error)

	// GetSubscriptionsFunc mocks the GetSubscriptions method.
	GetSubscriptionsFunc func(integrationID string) ([]models.Subscription, error)

	// GetUniformIntegrationsFunc mocks the GetUniformIntegrations method.
	GetUniformIntegrationsFunc func(filter models.GetUniformIntegrationsParams) ([]models.Integration, error)

	// UpdateLastSeenFunc mocks the UpdateLastSeen method.
	UpdateLastSeenFunc func(integrationID string) (*models.Integration, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateOrUpdateSubscription holds details about calls to the CreateOrUpdateSubscription method.
		CreateOrUpdateSubscription []struct {
			// IntegrationID is the integrationID argument value.
			IntegrationID string
			// Subscription is the subscription argument value.
			Subscription models.Subscription
		}
		// CreateOrUpdateUniformIntegration holds details about calls to the CreateOrUpdateUniformIntegration method.
		CreateOrUpdateUniformIntegration []struct {
			// Integration is the integration argument value.
			Integration models.Integration
		}
		// CreateUniformIntegration holds details about calls to the CreateUniformIntegration method.
		CreateUniformIntegration []struct {
			// Integration is the integration argument value.
			Integration models.Integration
		}
		// DeleteServiceFromSubscriptions holds details about calls to the DeleteServiceFromSubscriptions method.
		DeleteServiceFromSubscriptions []struct {
			// SubscriptionName is the subscriptionName argument value.
			SubscriptionName string
		}
		// DeleteSubscription holds details about calls to the DeleteSubscription method.
		DeleteSubscription []struct {
			// IntegrationID is the integrationID argument value.
			IntegrationID string
			// SubscriptionID is the subscriptionID argument value.
			SubscriptionID string
		}
		// DeleteUniformIntegration holds details about calls to the DeleteUniformIntegration method.
		DeleteUniformIntegration []struct {
			// ID is the id argument value.
			ID string
		}
		// GetSubscription holds details about calls to the GetSubscription method.
		GetSubscription []struct {
			// IntegrationID is the integrationID argument value.
			IntegrationID string
			// SubscriptionID is the subscriptionID argument value.
			SubscriptionID string
		}
		// GetSubscriptions holds details about calls to the GetSubscriptions method.
		GetSubscriptions []struct {
			// IntegrationID is the integrationID argument value.
			IntegrationID string
		}
		// GetUniformIntegrations holds details about calls to the GetUniformIntegrations method.
		GetUniformIntegrations []struct {
			// Filter is the filter argument value.
			Filter models.GetUniformIntegrationsParams
		}
		// UpdateLastSeen holds details about calls to the UpdateLastSeen method.
		UpdateLastSeen []struct {
			// IntegrationID is the integrationID argument value.
			IntegrationID string
		}
	}
	lockCreateOrUpdateSubscription       sync.RWMutex
	lockCreateOrUpdateUniformIntegration sync.RWMutex
	lockCreateUniformIntegration         sync.RWMutex
	lockDeleteServiceFromSubscriptions   sync.RWMutex
	lockDeleteSubscription               sync.RWMutex
	lockDeleteUniformIntegration         sync.RWMutex
	lockGetSubscription                  sync.RWMutex
	lockGetSubscriptions                 sync.RWMutex
	lockGetUniformIntegrations           sync.RWMutex
	lockUpdateLastSeen                   sync.RWMutex
}

// CreateOrUpdateSubscription calls CreateOrUpdateSubscriptionFunc.
func (mock *UniformRepoMock) CreateOrUpdateSubscription(integrationID string, subscription models.Subscription) error {
	if mock.CreateOrUpdateSubscriptionFunc == nil {
		panic("UniformRepoMock.CreateOrUpdateSubscriptionFunc: method is nil but UniformRepo.CreateOrUpdateSubscription was just called")
	}
	callInfo := struct {
		IntegrationID string
		Subscription  models.Subscription
	}{
		IntegrationID: integrationID,
		Subscription:  subscription,
	}
	mock.lockCreateOrUpdateSubscription.Lock()
	mock.calls.CreateOrUpdateSubscription = append(mock.calls.CreateOrUpdateSubscription, callInfo)
	mock.lockCreateOrUpdateSubscription.Unlock()
	return mock.CreateOrUpdateSubscriptionFunc(integrationID, subscription)
}

// CreateOrUpdateSubscriptionCalls gets all the calls that were made to CreateOrUpdateSubscription.
// Check the length with:
//     len(mockedUniformRepo.CreateOrUpdateSubscriptionCalls())
func (mock *UniformRepoMock) CreateOrUpdateSubscriptionCalls() []struct {
	IntegrationID string
	Subscription  models.Subscription
} {
	var calls []struct {
		IntegrationID string
		Subscription  models.Subscription
	}
	mock.lockCreateOrUpdateSubscription.RLock()
	calls = mock.calls.CreateOrUpdateSubscription
	mock.lockCreateOrUpdateSubscription.RUnlock()
	return calls
}

// CreateOrUpdateUniformIntegration calls CreateOrUpdateUniformIntegrationFunc.
func (mock *UniformRepoMock) CreateOrUpdateUniformIntegration(integration models.Integration) error {
	if mock.CreateOrUpdateUniformIntegrationFunc == nil {
		panic("UniformRepoMock.CreateOrUpdateUniformIntegrationFunc: method is nil but UniformRepo.CreateOrUpdateUniformIntegration was just called")
	}
	callInfo := struct {
		Integration models.Integration
	}{
		Integration: integration,
	}
	mock.lockCreateOrUpdateUniformIntegration.Lock()
	mock.calls.CreateOrUpdateUniformIntegration = append(mock.calls.CreateOrUpdateUniformIntegration, callInfo)
	mock.lockCreateOrUpdateUniformIntegration.Unlock()
	return mock.CreateOrUpdateUniformIntegrationFunc(integration)
}

// CreateOrUpdateUniformIntegrationCalls gets all the calls that were made to CreateOrUpdateUniformIntegration.
// Check the length with:
//     len(mockedUniformRepo.CreateOrUpdateUniformIntegrationCalls())
func (mock *UniformRepoMock) CreateOrUpdateUniformIntegrationCalls() []struct {
	Integration models.Integration
} {
	var calls []struct {
		Integration models.Integration
	}
	mock.lockCreateOrUpdateUniformIntegration.RLock()
	calls = mock.calls.CreateOrUpdateUniformIntegration
	mock.lockCreateOrUpdateUniformIntegration.RUnlock()
	return calls
}

// CreateUniformIntegration calls CreateUniformIntegrationFunc.
func (mock *UniformRepoMock) CreateUniformIntegration(integration models.Integration) error {
	if mock.CreateUniformIntegrationFunc == nil {
		panic("UniformRepoMock.CreateUniformIntegrationFunc: method is nil but UniformRepo.CreateUniformIntegration was just called")
	}
	callInfo := struct {
		Integration models.Integration
	}{
		Integration: integration,
	}
	mock.lockCreateUniformIntegration.Lock()
	mock.calls.CreateUniformIntegration = append(mock.calls.CreateUniformIntegration, callInfo)
	mock.lockCreateUniformIntegration.Unlock()
	return mock.CreateUniformIntegrationFunc(integration)
}

// CreateUniformIntegrationCalls gets all the calls that were made to CreateUniformIntegration.
// Check the length with:
//     len(mockedUniformRepo.CreateUniformIntegrationCalls())
func (mock *UniformRepoMock) CreateUniformIntegrationCalls() []struct {
	Integration models.Integration
} {
	var calls []struct {
		Integration models.Integration
	}
	mock.lockCreateUniformIntegration.RLock()
	calls = mock.calls.CreateUniformIntegration
	mock.lockCreateUniformIntegration.RUnlock()
	return calls
}

// DeleteServiceFromSubscriptions calls DeleteServiceFromSubscriptionsFunc.
func (mock *UniformRepoMock) DeleteServiceFromSubscriptions(subscriptionName string) error {
	if mock.DeleteServiceFromSubscriptionsFunc == nil {
		panic("UniformRepoMock.DeleteServiceFromSubscriptionsFunc: method is nil but UniformRepo.DeleteServiceFromSubscriptions was just called")
	}
	callInfo := struct {
		SubscriptionName string
	}{
		SubscriptionName: subscriptionName,
	}
	mock.lockDeleteServiceFromSubscriptions.Lock()
	mock.calls.DeleteServiceFromSubscriptions = append(mock.calls.DeleteServiceFromSubscriptions, callInfo)
	mock.lockDeleteServiceFromSubscriptions.Unlock()
	return mock.DeleteServiceFromSubscriptionsFunc(subscriptionName)
}

// DeleteServiceFromSubscriptionsCalls gets all the calls that were made to DeleteServiceFromSubscriptions.
// Check the length with:
//     len(mockedUniformRepo.DeleteServiceFromSubscriptionsCalls())
func (mock *UniformRepoMock) DeleteServiceFromSubscriptionsCalls() []struct {
	SubscriptionName string
} {
	var calls []struct {
		SubscriptionName string
	}
	mock.lockDeleteServiceFromSubscriptions.RLock()
	calls = mock.calls.DeleteServiceFromSubscriptions
	mock.lockDeleteServiceFromSubscriptions.RUnlock()
	return calls
}

// DeleteSubscription calls DeleteSubscriptionFunc.
func (mock *UniformRepoMock) DeleteSubscription(integrationID string, subscriptionID string) error {
	if mock.DeleteSubscriptionFunc == nil {
		panic("UniformRepoMock.DeleteSubscriptionFunc: method is nil but UniformRepo.DeleteSubscription was just called")
	}
	callInfo := struct {
		IntegrationID  string
		SubscriptionID string
	}{
		IntegrationID:  integrationID,
		SubscriptionID: subscriptionID,
	}
	mock.lockDeleteSubscription.Lock()
	mock.calls.DeleteSubscription = append(mock.calls.DeleteSubscription, callInfo)
	mock.lockDeleteSubscription.Unlock()
	return mock.DeleteSubscriptionFunc(integrationID, subscriptionID)
}

// DeleteSubscriptionCalls gets all the calls that were made to DeleteSubscription.
// Check the length with:
//     len(mockedUniformRepo.DeleteSubscriptionCalls())
func (mock *UniformRepoMock) DeleteSubscriptionCalls() []struct {
	IntegrationID  string
	SubscriptionID string
} {
	var calls []struct {
		IntegrationID  string
		SubscriptionID string
	}
	mock.lockDeleteSubscription.RLock()
	calls = mock.calls.DeleteSubscription
	mock.lockDeleteSubscription.RUnlock()
	return calls
}

// DeleteUniformIntegration calls DeleteUniformIntegrationFunc.
func (mock *UniformRepoMock) DeleteUniformIntegration(id string) error {
	if mock.DeleteUniformIntegrationFunc == nil {
		panic("UniformRepoMock.DeleteUniformIntegrationFunc: method is nil but UniformRepo.DeleteUniformIntegration was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockDeleteUniformIntegration.Lock()
	mock.calls.DeleteUniformIntegration = append(mock.calls.DeleteUniformIntegration, callInfo)
	mock.lockDeleteUniformIntegration.Unlock()
	return mock.DeleteUniformIntegrationFunc(id)
}

// DeleteUniformIntegrationCalls gets all the calls that were made to DeleteUniformIntegration.
// Check the length with:
//     len(mockedUniformRepo.DeleteUniformIntegrationCalls())
func (mock *UniformRepoMock) DeleteUniformIntegrationCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockDeleteUniformIntegration.RLock()
	calls = mock.calls.DeleteUniformIntegration
	mock.lockDeleteUniformIntegration.RUnlock()
	return calls
}

// GetSubscription calls GetSubscriptionFunc.
func (mock *UniformRepoMock) GetSubscription(integrationID string, subscriptionID string) (*models.Subscription, error) {
	if mock.GetSubscriptionFunc == nil {
		panic("UniformRepoMock.GetSubscriptionFunc: method is nil but UniformRepo.GetSubscription was just called")
	}
	callInfo := struct {
		IntegrationID  string
		SubscriptionID string
	}{
		IntegrationID:  integrationID,
		SubscriptionID: subscriptionID,
	}
	mock.lockGetSubscription.Lock()
	mock.calls.GetSubscription = append(mock.calls.GetSubscription, callInfo)
	mock.lockGetSubscription.Unlock()
	return mock.GetSubscriptionFunc(integrationID, subscriptionID)
}

// GetSubscriptionCalls gets all the calls that were made to GetSubscription.
// Check the length with:
//     len(mockedUniformRepo.GetSubscriptionCalls())
func (mock *UniformRepoMock) GetSubscriptionCalls() []struct {
	IntegrationID  string
	SubscriptionID string
} {
	var calls []struct {
		IntegrationID  string
		SubscriptionID string
	}
	mock.lockGetSubscription.RLock()
	calls = mock.calls.GetSubscription
	mock.lockGetSubscription.RUnlock()
	return calls
}

// GetSubscriptions calls GetSubscriptionsFunc.
func (mock *UniformRepoMock) GetSubscriptions(integrationID string) ([]models.Subscription, error) {
	if mock.GetSubscriptionsFunc == nil {
		panic("UniformRepoMock.GetSubscriptionsFunc: method is nil but UniformRepo.GetSubscriptions was just called")
	}
	callInfo := struct {
		IntegrationID string
	}{
		IntegrationID: integrationID,
	}
	mock.lockGetSubscriptions.Lock()
	mock.calls.GetSubscriptions = append(mock.calls.GetSubscriptions, callInfo)
	mock.lockGetSubscriptions.Unlock()
	return mock.GetSubscriptionsFunc(integrationID)
}

// GetSubscriptionsCalls gets all the calls that were made to GetSubscriptions.
// Check the length with:
//     len(mockedUniformRepo.GetSubscriptionsCalls())
func (mock *UniformRepoMock) GetSubscriptionsCalls() []struct {
	IntegrationID string
} {
	var calls []struct {
		IntegrationID string
	}
	mock.lockGetSubscriptions.RLock()
	calls = mock.calls.GetSubscriptions
	mock.lockGetSubscriptions.RUnlock()
	return calls
}

// GetUniformIntegrations calls GetUniformIntegrationsFunc.
func (mock *UniformRepoMock) GetUniformIntegrations(filter models.GetUniformIntegrationsParams) ([]models.Integration, error) {
	if mock.GetUniformIntegrationsFunc == nil {
		panic("UniformRepoMock.GetUniformIntegrationsFunc: method is nil but UniformRepo.GetUniformIntegrations was just called")
	}
	callInfo := struct {
		Filter models.GetUniformIntegrationsParams
	}{
		Filter: filter,
	}
	mock.lockGetUniformIntegrations.Lock()
	mock.calls.GetUniformIntegrations = append(mock.calls.GetUniformIntegrations, callInfo)
	mock.lockGetUniformIntegrations.Unlock()
	return mock.GetUniformIntegrationsFunc(filter)
}

// GetUniformIntegrationsCalls gets all the calls that were made to GetUniformIntegrations.
// Check the length with:
//     len(mockedUniformRepo.GetUniformIntegrationsCalls())
func (mock *UniformRepoMock) GetUniformIntegrationsCalls() []struct {
	Filter models.GetUniformIntegrationsParams
} {
	var calls []struct {
		Filter models.GetUniformIntegrationsParams
	}
	mock.lockGetUniformIntegrations.RLock()
	calls = mock.calls.GetUniformIntegrations
	mock.lockGetUniformIntegrations.RUnlock()
	return calls
}

// UpdateLastSeen calls UpdateLastSeenFunc.
func (mock *UniformRepoMock) UpdateLastSeen(integrationID string) (*models.Integration, error) {
	if mock.UpdateLastSeenFunc == nil {
		panic("UniformRepoMock.UpdateLastSeenFunc: method is nil but UniformRepo.UpdateLastSeen was just called")
	}
	callInfo := struct {
		IntegrationID string
	}{
		IntegrationID: integrationID,
	}
	mock.lockUpdateLastSeen.Lock()
	mock.calls.UpdateLastSeen = append(mock.calls.UpdateLastSeen, callInfo)
	mock.lockUpdateLastSeen.Unlock()
	return mock.UpdateLastSeenFunc(integrationID)
}

// UpdateLastSeenCalls gets all the calls that were made to UpdateLastSeen.
// Check the length with:
//     len(mockedUniformRepo.UpdateLastSeenCalls())
func (mock *UniformRepoMock) UpdateLastSeenCalls() []struct {
	IntegrationID string
} {
	var calls []struct {
		IntegrationID string
	}
	mock.lockUpdateLastSeen.RLock()
	calls = mock.calls.UpdateLastSeen
	mock.lockUpdateLastSeen.RUnlock()
	return calls
}
