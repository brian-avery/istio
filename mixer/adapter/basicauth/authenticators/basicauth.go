package authenticators

import (
	"fmt"

	htpasswd "github.com/brian-avery/go-htpasswd"
	"github.com/fsnotify/fsnotify"
	"istio.io/istio/pkg/log"
)

const (
	HashAlgorithmBcrypt = "$2y$"
	HashAlgorithmMD5    = "$apr1$"
	HashAlgorithmSha1   = "{SHA}"
)

type BasicAuthAuthenticator struct {
	Users           map[string]string
	htpasswdHandler *htpasswd.HtpasswdFile
}

func NewBasicAuthAdapter(path string) (*BasicAuthAuthenticator, error) {
	crypt, err := htpasswd.New(path,
		[]htpasswd.PasswdParser{htpasswd.AcceptMd5, htpasswd.AcceptSha, htpasswd.AcceptBcrypt, htpasswd.AcceptSsha, htpasswd.AcceptPlain},
		htpasswdBadLineHandler)
	if err != nil {
		log.Errorf("Could not create htpasswd handler: %s", err.Error())
		return nil, fmt.Errorf("unable to create htpasswd handler: %s", err.Error())
	}

	adapter := &BasicAuthAuthenticator{
		htpasswdHandler: crypt,
	}
	//adapter.monitor(path)
	return adapter, nil
}

func (auth BasicAuthAuthenticator) Validate(user string, password string) bool {
	log.Infof("Validate user: %s password: %s", user, password)
	ok := auth.htpasswdHandler.Match(user, password)
	if !ok {
		log.Infof("User %s not present in system", user)
		return false
	}

	log.Infof("failed to validate password")
	return false
}

func (auth BasicAuthAuthenticator) monitor(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Warnf("faild to create watcher for HTPasswd file: %s", err.Error())
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Infof("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Infof("modified file:", event.Name)
					if err := auth.htpasswdHandler.Reload(htpasswdBadLineHandler); err != nil {
						log.Errorf("could not read modified file contents: %s", err.Error())
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Errorf("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Errorf(err.Error())
	}
	<-done
}

func htpasswdBadLineHandler(err error) {
	log.Errorf(err.Error())
}
