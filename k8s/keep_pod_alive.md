## Keeping the container alive:

In the deployment file, pod:

```yaml
containers:
- name: web-server
  image: mrsabath/web-ms:v2
  imagePullPolicy: Always
        command: [ "--" ]
        command: [ "/bin/bash", "-c", "--" ]
        args: [ "while true; do sleep 30; done;" ]        
```

In Dockerfile:

* In your Dockerfile use this command:
  * `CMD ["sh", "-c", "tail -f /dev/null"]`
* Build your docker image.
* Push it to your cluster or similar, just to make sure the image it's available.
* `kubectl run debug-container -it --image=<your-image>`


In the GoLang code

```
func main() {
	// start a controller
	controller := examplecontroller.ExampleController{
		ExampleClient: exampleClient,
		ExampleScheme: exampleScheme,
	}

	// create context to cancel controller when main exit
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	go controller.Run(ctx)

	// a go idiom to sleep forever without eating CPU
	select {}
}
```
