package iotmakerdockerbuilder

// SetImageBuildOptionsExtraHosts (english): Add hostname mappings at build-time. Use the same values as the docker
// client --add-host parameter.
//
//   values:
//     somehost:162.242.195.82
//     otherhost:50.31.209.229
//
//   An entry with the ip address and hostname is created in /etc/hosts inside containers for this build, e.g:
//
//     162.242.195.82 somehost
//     50.31.209.229 otherhost
//
// SetImageBuildOptionsExtraHosts (português): Adiciona itens ao mapa de hostname durante o processo de construção da
// imagem. Use os mesmos valores que em docker client --add-host parameter.
//
//   values:
//     somehost:162.242.195.82
//     otherhost:50.31.209.229
//
//   An entry with the ip address and hostname is created in /etc/hosts inside containers for this build, e.g:
//   Uma nova entrada com o endereço ip e hostname será criada dentro de /etc/hosts do container. Exemplo:
//
//     162.242.195.82 somehost
//     50.31.209.229 otherhost
func (e *ContainerBuilder) SetImageBuildOptionsExtraHosts(values []string) {
	e.buildOptions.ExtraHosts = values
}
