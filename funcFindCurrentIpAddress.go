package iotmaker_docker_builder

// FindCurrentIpAddress (english):
//
// FindCurrentIpAddress (português): Inspeciona a rede e devolve o IP atual do container
func (e *ContainerBuilder) FindCurrentIpAddress() (IP string, err error) {
	var id string
	if e.network == nil {
		id, err = e.dockerSys.NetworkFindIdByName("bridge")
		if err != nil {
			return
		}
		IP, err = e.findCurrentIpAddressSupport(id)
	} else {
		IP, err = e.findCurrentIpAddressSupport(e.network.GetNetworkID())
	}

	return
}
