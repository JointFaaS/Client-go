package strategy

type ResourceStrategy interface {
	SelectN([]CloudResource, int32) ([]CloudResource, error)
	Select([]CloudResource) (CloudResource, error)
}

type CloudResource interface {
	GetResourceMetrics() (*ResourceMetrics, error)
}

type ResourceMetrics struct {
	// Network Latency
	Distance int32

	// this field will be compared to Total
	Used int32

	// a custom number to represent the total size of resource
	Total int32

	// how many requests(128MB 1s) a unit of resource can handle
	StandardUnitToRequests int32
}
