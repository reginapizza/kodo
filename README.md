**<h1>Command line utilities using cobra and Go</h1>**
>A go program to develop command line utilities with commands, subcommands, aliase etc.

**Dependencies**
> "github.com/spf13/cobra" <br />
> "k8s.io/client-go/kubernetes" <br />
> "k8s.io/client-go/rest" <br />
> "k8s.io/apimachinery/pkg/apis/meta/v1" <br />
> "github.com/openshift/api/route/v1" <br />
> "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1" <br />
> "k8s.io/api/apps/v1" <br />
> "k8s.io/api/core/v1" <br />
> "k8s.io/apimachinery/pkg/apis/meta/v1" <br />

**How to Install**
> git clone https://github.com/cli-playground/kodo.git
> cd kodo

**Commands List**<br/>
Command to build project -> go build -o bin/kodo main.go<br/>
To try out different commands,go to folder bin/kodo and try running below commands<br />
>  count       command to count <resources> e.g you can try running count pods to get number of running pods on cluster url<br />
>  deploy      Command to deploy an image<br />
>  help        Help about available commands<br />
>  version     Version details<br />
  