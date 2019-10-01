# Some Go related notes.

## Basics:

```go
// print documentation:
go doc fmt.Println

// get data type of the variable/constant
fmt.Printf("%T\n", myobject)

// switch
switch a {
  case 10: fmt.Println("10")
  case 10: fmt.Println("20")
  default: fmt.Println("other")
}

// arrays
var favNums[5] float64
favNums[0] = 123
favNums[0] = 3.44
favNums[0] = 22
favNums[0] = 1.3
favNums[0] = 4
fmt.Println(favNums[3])

favNums2 := [5]float64 {1,2,3,4,5}
for i, value := range favNums2 {
  fmt.Println(value, i)
}
// no indexes printed:
for _, value := range favNums2 {
  fmt.Println(value)
}

// slices, unbounded arrays
numSlice := []int {5,4,3,2,1}
numSlice2 := numSlice[3:5] // start with 3 and  include all before 5
numSlice3 := make([]int, 5, 10) // put 5, 10 times
copy(numSlice3, numSlice)
numSlice3 = append(numSlice3,0,-1)

// maps
nameAge := make(map[string] int)
nameAge["Tony"] = 20
nameAge["Anna"] = 25
delete(nameAge, "Tony")
fmt.Println(len(nameAge))

// variadic function
fmt.Println(subtractThem(1,2,3,4,5))

// uknown size of the arguments
func subtrackThem(args ...int) int {
  finalValue := 0
  for _, value := range args {
    finalValue -= value
  }
  return finalValue
}

// defer with recover
func safeDiv(num1, num2 int) int {

  defer func() {
    recover()
  }()
  solution := num1 / num2
  return solution
}
```
## Pointers
```go
package main

import "fmt"

// POINTERS

func main() {

	// We pass the value of a variable to the function
	x := 0
	changeXVal(x)
	fmt.Println("x =",x)

	// If we pass a reference to the variable we can change the value
	// in a function
	changeXValNow(&x)
	fmt.Println("x =",x)

	// Get the address x points to with &
	fmt.Println("Memory Address for x =", &x)

	// We can also generate a pointer with new

	yPtr := new(int)
	changeYValNow(yPtr)
	fmt.Println("y =", *yPtr)

}

func changeXVal(x int) {

	// Has no effect on the value of x in main()
	x = 2

}

// * signals that we are being sent a reference to the value
func changeXValNow(x *int){

	// Change the value at the memory address referenced by the pointer
	// * gives us access to the value the pointer points at

	*x = 2 // Store 2 in the memory address x refers to

}

func changeYValNow(yPtr *int){

	*yPtr = 100

}
```

## Struct
```go
package main

import "fmt"

// STRUCTS

func main() {

// Define a rectangle
rect1 := Rectangle{leftX: 0, TopY: 50, height: 10, width: 10}

// Leave off attribute names if we know the order
// rect1 := Rectangle{0, 50, 10, 10}

// We access values with the dot operator
fmt.Println("Rectangle is", rect1.width, "wide")

// Call the method area for Rectangle
fmt.Println("Area of the rectangle =", rect1.area())

}

// We can define our own types using struct
type Rectangle struct{
	leftX float64
	TopY float64
	height float64
	width float64
}

// We can define methods for our Rectangle by adding the receiver
// rect *Rectangle between func and the function name so we can
// call it with the dot operator
func (rect *Rectangle) area() float64{

	return rect.width * rect.height

}

```

## Struct and Interfaces
```go
package main

import "fmt"
import "math"

// STRUCTS AND INTERFACES

func main() {

	rect := Rectangle{20, 50}
	circ := Circle{4}

	fmt.Println("Rectangle Area =", getArea(rect))
	fmt.Println("Circle Area =", getArea(circ))

}

// An interface defines a list of methods that a type must implement
// If that type implements those methods the proper method is executed
// even if the original is referred to with the interface name

type Shape interface {
	area() float64
}

type Rectangle struct{
	height float64
	width float64
}

type Circle struct{
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64{

	return shape.area()

}
```
## Web Server in Go
```go
package main

import (
    "fmt"
    "net/http"
)

// CREATE A HTTP SERVER

// http.ResponseWriter assembles the servers response and writes to
// the client
// http.Request is the clients request
func handler(w http.ResponseWriter, r *http.Request) {

	// Writes to the client
    fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Earth\n")
}

func main() {

	// Calls for function handlers output to match the directory /
    http.HandleFunc("/", handler)

    // Calls for function handler2 output to match directory /earth
    http.HandleFunc("/earth", handler2)

    // Listen to port 8080 and handle requests
    http.ListenAndServe(":8080", nil)
}
```

## Go Channels
```go
package main

import "fmt"
import "time"
import "strconv"

// CHANNELS
// Channels allow us to pass data between go routines

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string){

	pizzaNum++

	// Convert int into a string
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make Dough and Send for Sauce");

	// Send the pizzaName onto the channel for the next
	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)

}

func addSauce(stringChan chan string){

	// Receive the value passed on the channel
	pizza := <- stringChan

	fmt.Println("Add Sauce and Send", pizza, "for Toppings")

	// Send the pizzaName onto the channel for the next
	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)

}

func addToppings(stringChan chan string){

	// Receive the value passed on the channel
	pizza := <- stringChan

	fmt.Println("Add Toppings to", pizza, "and Ship")

	time.Sleep(time.Millisecond * 10)

}

func main() {

	// Make creates a channel that can hold a string
	// int channel intChan := make(chan int)
	stringChan := make(chan string)

	// Cycle through and make 3 pizzas
	for i := 0; i < 3; i++{

		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond * 5000)

	}

}
```

## Load dependencies


first in the script:

```go
function load_dep {
  if [ ! $(basename "$PWD") = "proxy" ]; then
    cd proxy || exit
  fi
  local SETCD_SRC="../../fr8r-secure"

  if [ ! -d "$SETCD_SRC" ]; then
    git clone "git@github.ibm.com:alchemy-containers/fr8r-secure.git" "$SETCD_SRC"
  fi

  // 'setcd_src' directory is expected by Dockerfile, so copy fr8r-secure src there
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
