# Some Go related notes.

## Load dependencies


first in the script:

```bash
function load_dep {
  if [ ! $(basename "$PWD") = "proxy" ]; then
    cd proxy || exit
  fi
  local SETCD_SRC="../../fr8r-secure"

  if [ ! -d "$SETCD_SRC" ]; then
    git clone "git@github.ibm.com:alchemy-containers/fr8r-secure.git" "$SETCD_SRC"
  fi

  # 'setcd_src' directory is expected by Dockerfile, so copy fr8r-secure src there
  local SETCD_LIB=setcd_src
  echo "PWD: $PWD"
  rm -rf $SETCD_LIB
  TARGET="$SETCD_LIB/github.ibm.com/alchemy-containers/"
  mkdir -p "$TARGET"
  cp -r "$SETCD_SRC" "$TARGET"
}
```


then in a Dockerfile:

```bash
COPY src/ ./src/
COPY setcd_src/ ./src/
COPY config.toml /api-proxy/config.toml

RUN apt-get update \
 && apt-get -y install --no-install-recommends ca-certificates curl git \
 && curl -sSLO https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz \
 && tar -C /usr/local -xzf go1.7.3.linux-amd64.tar.gz \
 && rm -f go1.7.3.linux-amd64.tar.gz \
 && export PATH=$PATH:/usr/local/go/bin \
 && export GOPATH=/api-proxy \
 && go get "github.com/golang/glog" \
 && go get "github.com/spf13/viper" \
 && go get "github.com/coreos/etcd/client" \
 && go get "github.com/coreos/pkg/capnslog" \
 && go get "github.com/coreos/etcd/clientv3" \
 && go get "github.ibm.com/alchemy-containers/fr8r-secure/etcd/encconfig" \
 && go get "github.ibm.com/alchemy-containers/fr8r-secure/etcd/skeysapiv3" \
 && go get "github.ibm.com/alchemy-containers/fr8r-secure/etcd/encwrap" \
 && go build -o bin/api-proxy src/api-proxy/api-proxy.go \
 && go build -o bin/etcd_client src/misc/etcdmain/etcd_client.go \
 && rm -rf /usr/local/go \
 && apt-get -y remove git \
 && apt-get clean autoclean \
 && apt-get -y autoremove \
 && rm -rf /var/lib/apt/lists/* \
 && rm -rf ./src/
```

## Build fake response:
Replace some string arguments (e.g. image)

```golang
func newFakeRequestInitContainer(initImage, image string) *http.Request {
	// TODO: Delete what we don't need for unit tests
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(fmt.Sprintf(`
			{
				"kind": "AdmissionReview",
				"apiVersion": "admission.k8s.io/v1beta1",
				"request": {
					"uid": "ed782967-1c99-11e8-936d-08002789d446",
					"kind": {
						"group": "",
						"version": "v1",
						"kind": "Pod"
					},
					"resource": {
						"group": "",
						"version": "v1",
						"resource": "pods"
					},
					"namespace": "default",
					"operation": "CREATE",
					"userInfo": {
						"username": "minikube-user",
						"groups": [
							"system:masters",
							"system:authenticated"
						]
					},
					"object": {
						"metadata": {
							"name": "nginx",
							"namespace": "default",
							"creationTimestamp": null
						},
						"spec": {
							"volumes": [
								{
									"name": "default-token-xff4f",
									"secret": {
										"secretName": "default-token-xff4f"
									}
								}
							],
									"initContainers" : [
					{
									"name": "nginx",
									"image": "%s",
									"ports": [
										{
											"hostPort": 8080,
											"containerPort": 8080,
											"protocol": "TCP"
										}
									],
									"resources": {},
									"volumeMounts": [
										{
											"name": "default-token-xff4f",
											"readOnly": true,
											"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
										}
									],
									"terminationMessagePath": "/dev/termination-log",
									"terminationMessagePolicy": "File",
									"imagePullPolicy": "Always"
						}
									],
							"containers": [
									 {
									"name": "statsd",
									"image": "%s",
									"ports": [
										{
											"hostPort": 8080,
											"containerPort": 8080,
											"protocol": "TCP"
										}
									],
									"resources": {},
									"volumeMounts": [
										{
											"name": "default-token-xff4f",
											"readOnly": true,
											"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
										}
									],
									"terminationMessagePath": "/dev/termination-log",
									"terminationMessagePolicy": "File",
									"imagePullPolicy": "Always"
								}
							],
							"restartPolicy": "Always",
							"terminationGracePeriodSeconds": 30,
							"dnsPolicy": "ClusterFirst",
							"serviceAccountName": "default",
							"serviceAccount": "default",
							"hostNetwork": true,
							"securityContext": {},
							"imagePullSecrets": [
								{
									"name": "regsecret"
								}
							],
							"schedulerName": "default-scheduler",
							"tolerations": [
								{
									"key": "node.kubernetes.io/not-ready",
									"operator": "Exists",
									"effect": "NoExecute",
									"tolerationSeconds": 300
								},
								{
									"key": "node.kubernetes.io/unreachable",
									"operator": "Exists",
									"effect": "NoExecute",
									"tolerationSeconds": 300
								}
							]
						},
						"status": {}
					},
					"oldObject": null
				}
			}`, initImage, image)))
	req.Header.Set("Content-Type", "application/json")
	return req
}
```

## Convert object to JSON, format and log

```golang
// Log the JSON format of the object
func logJSON(msg string, v interface{}) {
	s := getJSON(v)
	glog.Infof("*** JSON for %v:\n %v", msg, s)

}

func getJSON(v interface{}) string {
	// Dump the object so it can be used for testing
	b, er := json.MarshalIndent(v, "", "    ")
	if er != nil {
		panic(er)
	}
	b2 := append(b, '\n')
	return string(b2)
}
```

## Compare two JSON representations
Test code for creating Fake k8s objects, testing mutation
functions, then compare the output with results specified
in Expect file, using jsondiff.

```golang
import "github.com/nsf/jsondiff"

func TestMutateInitialization(t *testing.T) {

	ret := ctiv1.ClusterTI{
		Info: ctiv1.ClusterTISpec{
			ClusterName:   "testCluster",
			ClusterRegion: "testRegion",
		},
	}
	clInfo := NewCigKubeTest(ret)

	icc, err := loadInitContainerConfig("ConfigFile.yaml")
	if err != nil {
		t.Errorf("Error loading InitContainerConfig %v", err)
		return
	}

	whsvr := &WebhookServer{
		initcontainerConfig: icc,
		server:              &http.Server{},
		createVaultCert:     true,
		clusterInfo:         clInfo,
	}

	req := getFakeAdmissionRequest()
	pod := getFakePod()

	// get test result of running mutateInitialization method:
	result, err := whsvr.mutateInitialization(pod, &req)
	if err != nil {
		t.Errorf("Error executing mutateInitialization %v", err)
		return
	}

	// one of the Annotation fields is dynamically set (UUID), so it needs to be
	// changed to static, for comparison
	annot := result.Annotations
	annot["admission.trusted.identity/ti-secret-key"] = "ti-secret-XXX"
	result.Annotations = annot

	// convert the result JSON to []byte
	resultB, err := json.Marshal(result)
	if err != nil {
		t.Errorf("Error marshal Result to []byte: %v", err)
		return
	}

	// get the exepected result from a file
	dat, err := ioutil.ReadFile("ExpectMutateInit.json")
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	expect := InitContainerConfig{}
	json.Unmarshal(dat, &expect)

	// convert the expect JSON to []byte
	expectB, err := json.Marshal(expect)
	if err != nil {
		t.Errorf("Error marshal Expected to []byte: %v", err)
		return
	}

	// Execute JSON diff on both byte representations of JSON
	// when using DefaultHTMLOptions, `text` can be treated as
	// HTML with <pre> to show differences with colors
	opts := jsondiff.DefaultHTMLOptions()
	diff, text := jsondiff.Compare(expectB, resultB, &opts)

	if diff == jsondiff.FullMatch {
		t.Logf("Results match expections: %v", diff)
	} else {
		t.Errorf("Results do not match expections: %v %v", diff, text)
		return
	}
}

func getContentOfTheFile(file string) string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func getFakePod() corev1.Pod {
	s := getContentOfTheFile("FakePod.json")
	pod := corev1.Pod{}
	json.Unmarshal([]byte(s), &pod)
	return pod
}


```
