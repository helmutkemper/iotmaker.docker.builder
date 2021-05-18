package iotmaker_docker_builder

// FindCurrentIPV4Address (english):
//
// FindCurrentIPV4Address (português): Inspeciona a rede e devolve o IP atual do container
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
