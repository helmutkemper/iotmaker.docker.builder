package iotmakerdockerbuilder

// ImageListExposedVolumes
//
// English: Lists all volumes defined in the image.
//
//   Note: Use the AddFiileOrFolderToLinkBetweenConputerHostAndContainer() function to link folders and files
//   between the host computer and the container
//
// Português: Lista todos os volumes definidos na imagem.
//
//   Nota: Use a função AddFiileOrFolderToLinkBetweenConputerHostAndContainer() para vincular pastas e arquivos
//   entre o computador hospedeiro e o container
func (e *ContainerBuilder) ImageListExposedVolumes() (list []string, err error) {

	list, err = e.dockerSys.ImageListExposedVolumes(e.imageID)
	return
}
