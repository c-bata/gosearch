# GoSearch

The search engine in golang.

## Run Application

#### Setup

```
docker-machine create --driver virtualbox default
VBoxManage controlvm "default" natpf1 "mongo,tcp,127.0.0.1,27017,,27017"
VBoxManage controlvm "default" natpf1 "redis,tcp,127.0.0.1,6379,,6379"
```

#### Run

```
docker-compose -f docker-compose-db.yml run
```

**Option(Throw fixture data)**

```
./fixtures/fixtures.sh
```
