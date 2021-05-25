package iotmakerdockerbuilder

// FindCurrentIPV4Address
//
// English: Inspects the docker's network and returns the current IP of the container
//
// Português: Inspeciona a rede do docker e devolve o IP atual do container
func (e *ContainerBuilder) FindCurrentIPV4Address() (IP string, err error) {
	var id string
	if e.network == nil {
		id, err = e.dockerSys.NetworkFindIdByName("bridge")
		if err != nil {
			return
		}
		IP, err = e.findCurrentIPV4AddressSupport(id)
	} else {
		IP, err = e.findCurrentIPV4AddressSupport(e.network.GetNetworkID())
	}

	return
}
