# Kubebuilder
Kubebuilder is an SDK for rapidly building and publishing Kubernetes APIs in Go. It builds on top of the canonical techniques used to build the core Kubernetes APIs to provide simple abstractions that reduce boilerplate and toil.

Install: https://book.kubebuilder.io/getting_started/installation_and_setup.html


Create a new project:
```
cd $GOPATH/src
mkdir kubebuilder-test
cd kubebuilder-test/
kubebuilder
```

Make sure 'dep' is installed:

```
kubebuilder init --domain ibm --owner = "Mariusz"
2018/08/21 10:02:57 Dep is not installed. Follow steps at: https://golang.github.io/dep/docs/installation.html
brew install dep
brew upgrade dep
```
