#
# English: first stage of the process
#
# Português: primeira etapa do processo
FROM golang:1.16-alpine as builder
#
# English: creates the .ssh directory within the root directory
#
# Português: cria o diretório .ssh dentro do diretório root
RUN mkdir -p /root/.ssh/ && \
    #
    # English: adjust file access security
    #
    # Português: ajusta a segurança de acesso do arquivo
    chmod -R 600 /root/.ssh/ && \
    #
    # English: prepares the OS for installation
    #
    # Português: prepara o OS para instalação
    apk update && \
    #
    # English: install git and openssh
    #
    # Português: instala o git e o opnssh
    apk add --no-cache build-base git openssh && \
    #
    # English: install git, fakeroot, scanelf, openssl, apk-tools, libc-utils, attr, tar, pkgconf, patch, lzip, curl,
    #      /bin/sh, so:libc.musl-x86_64.so.1, so:libcrypto.so.1.1 and so:libz.so.1
    #
    # Português: instala git, fakeroot, scanelf, openssl, apk-tools, libc-utils, attr, tar, pkgconf, patch, lzip, curl, 
    #      /bin/sh, so:libc.musl-x86_64.so.1, so:libcrypto.so.1.1 e so:libz.so.1
    apk add --no-cache alpine-sdk && \
    #
    # English: clear the cache
    #
    # Português: limpa a cache
    rm -rf /var/cache/apk/*
