# iotmaker.docker.builder

### English

This project creates a simple way to manipulate containers with the moby docker project

> Status: Starting to document in English

### Português

> Status: Documentando em ingles 

Este projeto cria uma API Golang simples para criar e manipular o docker.

### Create a docker network / Cria uma rede docker

```golang
  var err error
  var netDocker = dockerNetwork.ContainerBuilderNetwork{}
  err = netDocker.Init()
  if err != nil { panic(err) }

  // create a network named cache_delete_after_test, subnet 10.0.0.0/16 e gatway 10.0.0.1
  err = netDocker.NetworkCreate("cache_delete_after_test", "10.0.0.0/16", "10.0.0.1")
  if err != nil { panic(err) }
```

English: use the `container.SetNetworkDocker(&netDocker)` command to link the container to the network

Português: use o comando `container.SetNetworkDocker(&netDocker)` para ligar um container com o docker

### Container nats

```golang
  var err error

  // create a container
  var container = ContainerBuilder{}
  // set image name for docker pull
  container.SetImageName("nats:latest")
  // link container and network [optional] (next ip address is 10.0.0.2)
  container.SetNetworkDocker(&netDocker)
  // set a container name
  container.SetContainerName("container_delete_nats_after_test")
  // set a waits for the text to appear in the standard container output to proceed [optional]
  container.SetWaitStringWithTimeout("Listening for route connections on 0.0.0.0:6222", 10*time.Second)

  // inialize the container object
  err = container.Init()
  if err != nil { panic(err) }

  // image nats:latest pull command [optional]
  err = container.ImagePull()
  if err != nil { panic(err) }

  // container build and start from image nats:latest
  // waits for the text "Listening for route connections on 0.0.0.0:6222" to appear  in the standard container output
  // to proceed
  err = container.ContainerBuildFromImage()
  if err != nil { panic(err) }
```

### Container from github project

```golang
  var err error
  var container = ContainerBuilder{}
  // new image name delete:latest
  container.SetImageName("delete:latest")
  // container name container_delete_server_after_test
  container.SetContainerName("container_delete_server_after_test")
  // git project to clone https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git
  container.SetGitCloneToBuild("https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git")
    
  // see SetGitCloneToBuildWithUserPassworh(), SetGitCloneToBuildWithPrivateSshKey() and
  // SetGitCloneToBuildWithPrivateToken()
    
  // set a waits for the text to appear in the standard container output to proceed [optional]
  container.SetWaitStringWithTimeout("Stating server on port 3000", 10*time.Second)
  // change and open port 3000 to 3030
  container.AddPortToChange("3000", "3030")
  // replace container folder /static to host folder ./test/static
  err = container.AddFiileOrFolderToLinkBetweenConputerHostAndContainer("./test/static", "/static")
  if err != nil { panic(err) }
    
  // inicialize container object
  err = container.Init()
  if err != nil { panic(err) }
    
  // builder new image from git project
  err = container.ImageBuildFromServer()
  if err != nil { panic(err) }

  err = container.ContainerBuildFromImage()
  if err != nil { panic(err) }
```

### MongoDB ephemeral

```golang
  var err error
  var mongoDocker = &ContainerBuilder{}
  mongoDocker.SetImageName("mongo:latest")
  mongoDocker.SetContainerName("container_delete_mongo_after_test")
  mongoDocker.AddPortToOpen("27017")
  mongoDocker.SetEnvironmentVar(
    []string{
      "--host 0.0.0.0",
    },
  )
  mongoDocker.SetWaitStringWithTimeout(`"msg":"Waiting for connections","attr":{"port":27017`, 20*time.Second)
  err = mongoDocker.Init()
  if err != nil { panic(err) }

  err = mongoDocker.ContainerBuildFromImage()
  if err != nil { panic(err) }
```

### MongoDB

```golang
  var err error
  var mongoDocker = &ContainerBuilder{}
  mongoDocker.SetImageName("mongo:latest")
  mongoDocker.SetContainerName("container_delete_mongo_after_test")
  mongoDocker.AddPortToOpen("27017")
  mongoDocker.SetEnvironmentVar(
    []string{
      "--host 0.0.0.0",
    },
  )
  err = mongoDocker.AddFiileOrFolderToLinkBetweenConputerHostAndContainer("./test/data", "/data")
  mongoDocker.SetWaitStringWithTimeout(`"msg":"Waiting for connections","attr":{"port":27017`, 20*time.Second)
  err = mongoDocker.Init()
  err = mongoDocker.ContainerBuildFromImage()
```

### Container from folder

```golang
  var err error

  GarbageCollector()
  var container = ContainerBuilder{}
  // new image name delete:latest
  container.SetImageName("delete:latest")
  // set a folder path to make a new image
  container.SetBuildFolderPath("./test/server")
  // container name container_delete_server_after_test
  container.SetContainerName("container_delete_server_after_test")
  // set a waits for the text to appear in the standard container output to proceed [optional]
  container.SetWaitStringWithTimeout("starting server at port 3000", 10*time.Second)
  // change and open port 3000 to 3030
  container.AddPortToOpen("3000")
  // replace container folder /static to host folder ./test/static
  err = container.AddFiileOrFolderToLinkBetweenConputerHostAndContainer("./test/static", "/static")
  if err != nil { panic(err) }

  // inicialize container object
  err = container.Init()
  if err != nil { panic(err) }

  // builder new image from folder
  err = container.ImageBuildFromFolder()
  if err != nil { panic(err) }

  // build a new container from image
  err = container.ContainerBuildFromImage()
  if err != nil { panic(err) }
```