package cmd

import (
	"context"
	"fmt"

	"k8s.io/client-go/rest"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	buildv1api "github.com/openshift/api/build/v1"
	imagev1api "github.com/openshift/api/image/v1"

	buildv1clientapi "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	imagev1clientapi "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
)

var envVar = new(EnvironmentVariables) // creating instance of struct EnvironmentVariables from cmd.openshiftclient.go

// Source :
var (
	Source string
)

func createTypeMeta(kind string, APIVersion string) metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       kind,
		APIVersion: APIVersion,
	}
}

func createObjectType(name string, namespace string) metav1.ObjectMeta {
	if namespace != "" {
		return metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		}
	} else {
		return metav1.ObjectMeta{
			Name: name,
		}
	}
}

func createBuildSpec(uri string) buildv1api.BuildConfigSpec {
	return buildv1api.BuildConfigSpec{
		CommonSpec: buildv1api.CommonSpec{
			Source: buildv1api.BuildSource{
				Type: buildv1api.BuildSourceType("Git"),
				Git: &buildv1api.GitBuildSource{
					URI: uri,
				},
			},
			Strategy: buildv1api.BuildStrategy{
				Type: buildv1api.BuildStrategyType("Docker"),
			},
			Output: buildv1api.BuildOutput{
				To: &corev1.ObjectReference{
					Kind: "ImageStreamTag",
					Name: "my-ruby-image:latest",
				},
			},
		},
	}
}

func createBuildConfig() error {
	buildclient := newBuildConfigClient()
	buildConfig := buildv1api.BuildConfig{
		TypeMeta:   createTypeMeta("BuildConfig", "build.openshift.io/v1"),
		ObjectMeta: createObjectType("my-app-docker-build", ""),
		Spec:       createBuildSpec(Source),
	}

	_, builderror := buildclient.BuildConfigs(envVar.Namespace).Create(context.TODO(), &buildConfig, metav1.CreateOptions{})

	if builderror != nil {
		fmt.Println("The build failed,", builderror)
		return builderror
	}
	return nil
}

func createImageStream() error {
	imagestreamClient := newImageStreamClient()
	imageStream := imagev1api.ImageStream{
		TypeMeta:   createTypeMeta("ImageStream", "image.openshift.io/v1"),
		ObjectMeta: createObjectType("my-ruby-image", "regina-build"),
	}

	_, imgerror := imagestreamClient.ImageStreams(envVar.Namespace).Create(context.TODO(), &imageStream, metav1.CreateOptions{})

	if imgerror != nil {
		fmt.Println("The image stream failed,", imgerror)
		return imgerror
	}

	return nil
}

func newImageStreamClient() *imagev1clientapi.ImageV1Client {
	config := rest.Config{
		Host:        envVar.Host,
		BearerToken: envVar.Bearertoken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	myClientSet, _ := imagev1clientapi.NewForConfig(&config)
	return myClientSet
}

func newBuildConfigClient() *buildv1clientapi.BuildV1Client {
	config := rest.Config{
		Host:        envVar.Host,
		BearerToken: envVar.Bearertoken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	myClientSet, _ := buildv1clientapi.NewForConfig(&config)
	return myClientSet
}

//Build image from dockerfile at github source
func Build() (error, error) {
	// buildclient := newBuildConfigClient()
	// buildconfig := createBuildConfig()

	// imagestreamclient := newImageStreamClient()
	// imagestream := createImageStream()

	imgerror := createImageStream()
	builderror := createBuildConfig()

	return imgerror, builderror
}
