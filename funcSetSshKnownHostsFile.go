package iotmakerdockerbuilder

// SetSshKnownHostsFile
//
// English: set a sseh knownhosts file
//
//     var err error
//     var usr *user.User
//     var path string
//     var file []byte
//     usr, err = user.Current()
//     if err != nil {
//       panic(err)
//     }
//
//     path = filepath.Join(usr.HomeDir, ".ssh", "known_hosts")
//     file, err = ioutil.ReadFile(path)
//     if err != nil {
//       panic(err)
//     }
//
//     var container = ContainerBuilder{}
//     container.SetSshKnownHostsFile(string(file))
//
// Português: define o arquivo knownhosts do ssh
//
//     var err error
//     var usr *user.User
//     var path string
//     var file []byte
//     usr, err = user.Current()
//     if err != nil {
//       panic(err)
//     }
//
//     path = filepath.Join(usr.HomeDir, ".ssh", "known_hosts")
//     file, err = ioutil.ReadFile(path)
//     if err != nil {
//       panic(err)
//     }
//
//     var container = ContainerBuilder{}
//     container.SetSshKnownHostsFile(string(file))
func (e *ContainerBuilder) SetSshKnownHostsFile(value string) {
	e.contentKnownHostsFile = value
}
