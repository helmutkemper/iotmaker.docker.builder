package iotmakerdockerbuilder

// SetBuildFolderPath (english):
//
// SetBuildFolderPath (português): Define o caminho da pasta a ser transformada em imagem
//   value: caminho da pasta a ser transformada em imagem
//
//     Nota: A pasta deve conter um arquivo dockerfile, mas, como diferentes usos podem ter diferentes dockerfiles,
//     será dada a seguinte ordem na busca pelo arquivo: "Dockerfile-iotmaker", "Dockerfile", "dockerfile" na pasta raiz.
//     Se não houver encontrado, será feita uma busca recusiva por "Dockerfile" e "dockerfile"
//
func (e *ContainerBuilder) SetBuildFolderPath(value string) {
	e.buildPath = value
}
