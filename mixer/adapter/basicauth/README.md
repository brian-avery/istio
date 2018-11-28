This is an adapter for Istio's Mixer that is designed to support Basic Authentication.
This reads an htpasswd file as specified in the htpasswd_file argument in the yaml.
The following htpasswd encodings are supported:
* SSHA
* SHA
* Bcrypt
* MD5
* Apache APR1 Crypt


Installation:

Building
A makefile is provided in order to make the process easier. To build the adapter do a make build. To push it, do a make push with the hub variable set to a valid repo. Finally, deploy will deploy it to the Kubernetes cluster connected to kubectl.
As an example: hub=docker.io/myrepo make build push deploy.
