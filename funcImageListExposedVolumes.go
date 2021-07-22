package iotmakerdockerbuilder

// ImageListExposedVolumes
//
// English: Lists all volumes defined in the image.
//
//   Note: Use the AddFileOrFolderToLinkBetweenConputerHostAndContainer() function to link folders and files
//   between the host computer and the container
//
// Português: Lista todos os volumes definidos na imagem.
//
//   Nota: Use a função AddFileOrFolderToLinkBetweenConputerHostAndContainer() para vincular pastas e arquivos
//   entre o computador hospedeiro e o container
func (e *ContainerBuilder) ImageListExposedVolumes() (list []string, err error) {

	list, err = e.dockerSys.ImageListExposedVolumes(e.imageID)
	return
}
