package authenticators

import (
	"fmt"

	htpasswd "github.com/brian-avery/go-htpasswd"
	"github.com/fsnotify/fsnotify"
	"istio.io/istio/pkg/log"
)

type BasicAuthAuthenticator struct {
	htpasswdHandler *htpasswd.HtpasswdFile
}

//NewBasicAuthAdapter creates a new instance of the basic auth autenticator
func NewBasicAuthAdapter(path string) (*BasicAuthAuthenticator, error) {
	fmt.Println("Create crypt")
	log.Infof("Create cyrpt")
	crypt, err := htpasswd.New(path,
		[]htpasswd.PasswdParser{htpasswd.AcceptMd5, htpasswd.AcceptSha, htpasswd.AcceptBcrypt, htpasswd.AcceptSsha, htpasswd.AcceptPlain},
		nil)
	if err != nil {
		log.Errorf("Could not create htpasswd handler: %s", err.Error())
		return nil, fmt.Errorf("unable to create htpasswd handler: %s", err.Error())
	}

	adapter := &BasicAuthAuthenticator{
		htpasswdHandler: crypt,
	}
	fmt.Printf("Setting up monitor for: %s", path)
	log.Infof("Setting up monitor for: %s", path)
	adapter.monitor(path)
	return adapter, nil
}

//Validate validates a user token
func (auth BasicAuthAuthenticator) Validate(user string, password string) bool {
	ok := auth.htpasswdHandler.Match(user, password)
	if !ok {
		log.Infof("User %s not present in system", user)
		return false
	}

	log.Infof("validated password successfully")
	return true
}

//monitor accepts a path to an htpasswd file and monitors for changes, reloading the htpasswd file when they are detected
func (auth BasicAuthAuthenticator) monitor(path string) {
	fmt.Printf("monitor\n")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Errorf("failed to create watcher for HTPasswd file: %s", err.Error())
	}
	defer watcher.Close()

	go func() {
		for {
			fmt.Printf("Monitor\n")
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Infof("modified file: %s", event.Name)
					if err := auth.htpasswdHandler.Reload(nil); err != nil {
						log.Errorf("could not read modified file contents: %s", err.Error())
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Errorf("error:%s", err.Error())
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Errorf("Error adding watcher for path:%s %s", path, err.Error())
	}
}
