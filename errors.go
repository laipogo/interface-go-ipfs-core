package iface

import "errors"

var ErrIsDir = errors.New("object is a directory")
var ErrOffline = errors.New("can't resolve, ipfs node is offline")
