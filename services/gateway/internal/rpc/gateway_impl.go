package rpc

type GatewayImpl struct {	
}

type GatewayInterface interface {
	BasicQuery()
	AdvancedFilter()
}

func (g *GatewayImpl) BasicQuery() {
    // Implementation for basic query functionality
}

func (g *GatewayImpl) AdvancedFilter() {
    // Implementation for advanced filter functionality
}

// RegisterGateway registers the gateway service with the given name and implementation
