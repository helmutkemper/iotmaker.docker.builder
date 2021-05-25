package iotmakerdockerbuilder

// SetSshIdRsaFile (english):
//
// SetSshIdRsaFile (português):
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
