package iotmakerdockerbuilder

// SetSshIdRsaFile
//
// English:
//
//  Set a id_rsa file from shh
//
// Example:
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
//     path = filepath.Join(usr.HomeDir, ".ssh", "id_rsa")
//     file, err = ioutil.ReadFile(path)
//     if err != nil {
//       panic(err)
//     }
//
//     var container = ContainerBuilder{}
//     container.SetSshIdRsaFile(string(file))
//
// Português:
//
//  Define o arquivo id_rsa do shh
//
// Exemplo:
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
//     path = filepath.Join(usr.HomeDir, ".ssh", "id_rsa")
//     file, err = ioutil.ReadFile(path)
//     if err != nil {
//       panic(err)
//     }
//
//     var container = ContainerBuilder{}
//     container.SetSshIdRsaFile(string(file))
func (e *ContainerBuilder) SetSshIdRsaFile(value string) {
	e.contentIdRsaFile = value
}
