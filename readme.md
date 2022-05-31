You can work directly with go files from your local machine or you can 
work through the docker image.
You will figure out how to work through the local environment and how 
to run through docker:
1) Open a terminal in `<project_dir>/app`
2) Run `docker-compose up`, you should see `tail: can't open 
'open-connection': No such file or directory`. The container is up and
container is running.
3) Open another terminal, go to `<project_dir>/app`, run `docker-compose
 exec go sh`. Next, the terminal in the container should open, the 
current pointer will be to `/go/src/app`.

Running tets:
1) Navigate to `<project_dir>/app/src`, inside docker container 
`/go/src/app/src`
2) Run `go test`

_Telegram: @agoodminute_