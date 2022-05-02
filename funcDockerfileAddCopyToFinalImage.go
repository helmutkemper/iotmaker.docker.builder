package iotmakerdockerbuilder

// DockerfileAddCopyToFinalImage
//
// English:
//
//  Add one instruction 'COPY --from=builder /app/`dst` `src`' to final image builder.
//
// Português:
//
//  Adiciona uma instrução 'COPY --from=builder /app/`dst` `src`' ao builder da imagem final.
//
//
func (e *ContainerBuilder) DockerfileAddCopyToFinalImage(src, dst string) {
	e.autoDockerfile.AddCopyToFinalImage(src, dst)
}
