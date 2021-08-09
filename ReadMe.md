# Solution

## Application

I slightly modified the application so Fiber will always listen to a fixed port (8080 in the example).

## Dockerize

Being the first time when i worked with a Go application I relied on Docker documentation to prepare the Dockerfile. I used the multistaged build as the resulting image was much smaller in this case.
<https://docs.docker.com/language/golang/build-images/>

To obtain the image run teh following command after changing dir to `j2m-sre-test`:
`docker build -t j2m-test .`

## Docker compose

In order to test the application I created a docker-compose.yml file which spins two containers, one obtained from our application and the second one using the redis:alpine image.
The REDIS_URL is passed as an environment variable.

- Start:
`docker compose up -d`
- Open localhost:8080 in a browser, refresh a few times to see the incrementing of the counter.
- Stop:
`docker compose down`

## Registry

I've built and published the image to the GitLab registry using the following commands.
Normally a PAT (personal access token) would be preffered instead of user and password considering that one such PAT can be used for CI/CD pipelines as well.

```bash
$ docker login registry.gitlab.com
Username: betejb
Password: 
Login Succeeded

bogdan.betej@PC1819 MINGW64 /c/NTT/Extra/jobtome-gitlab/j2m-sre-test (feature/prepare-build)
$ docker build -t registry.gitlab.com/betejb/j2m-sre-test .
.....

$ docker push registry.gitlab.com/betejb/j2m-sre-test
Using default tag: latest
The push refers to repository [registry.gitlab.com/betejb/j2m-sre-test]
2deefee5a2c8: Pushed
208cc594e520: Pushed
07363fa84210: Pushed
latest: digest: sha256:02a3a7547ce78a3b8210c9dcf5c6094aa0f8c43efe6665e51b999e1cf64414c6 size: 949
```

## Kubernetes

The implementation is a very basic one, manifest files are in k8s folder
I created two deployments one for redis and one for our Go application (a simpler option would have been to create two containers in the same deployment, in such a case the redis would have been addressed by localhost:6379), two services, a namespace and an ingress object.

The ingress will route requests received at `http://j2m-sre-test.go.dev` (implicit port HTTP 80) to the `j2m-sre-test` service.
Our Go container will connect to redis via the dedicated service `redis-go`

We assume a secret (`regcred`) is created to allow the download of images from GitLab registry. It should contain the PAT mentioned above.