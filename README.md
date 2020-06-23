## Kodo - Command line interface for OpenShift
Kodo is a command line interface to build and deploy applications on Openshift.

## Setup
- `git clone https://github.com/cli-playground/kodo.git` 
- `cd kodo`
- To build run,  `go build -o bin/kodo main.go`

## Contributing
 1. Fork from https://github.com/cli-playground/kodo.git
 2. Create Your feature branch (`git checkout -b feature/fooBar`)
 3. Commit your changes (`git commit -am 'Add some fooBar'`)
 4. Push to the branch (`git push origin feature/fooBar`)
 5. Create a new pull request

## Dependencies
 - "github.com/spf13/cobra" 
 - "k8s.io/client-go/kubernetes"  
 - "k8s.io/client-go/rest" 
 - "k8s.io/apimachinery/pkg/apis/meta/v1" 
 - "github.com/openshift/api/route/v1" 
 - "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1" 
 - "k8s.io/api/apps/v1" 
 - "k8s.io/api/core/v1" 
 - "k8s.io/apimachinery/pkg/apis/meta/v1" 

## Commands List

Before running commands, Set Path to $HOME/kodo/bin/.

1. Count Resources: Command to count resources like pods running on cluster\
    `kodo count <resources>` \
    For example, `kodo count pods --server=<server url> --token=<token>` will output count of running pods of a given cluster.

2. Deploy : Command to deploy an image \
    `kodo deploy --image=<image> --replicas=<no of replicas> --port=<port number> --token=<token> --server=<cluster url>  --namespace=<namespace>` \

3. Build: Command to build an image from source \
    `kodo build --source=<dockerfile source> --namespace=<namespace> 
 --token=<Token>--server=<cluster url>` \

4. Help: Command to help user to list all available commands and flags\
    `kodo help`

5. Version : Command to check current version \
    `kodo version`
