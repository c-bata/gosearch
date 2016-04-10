# GoSearch

[![Build Status](https://travis-ci.org/c-bata/gosearch.svg?branch=master)](https://travis-ci.org/c-bata/gosearch)

The search engine in golang.

## Flow

1. Crawling pages.
2. Build an inverted index. Use Kagome that is a morphological analysis engine.
3. Web frontend with gin returns the results.

## Run Application

#### Setup

```
docker-machine create --driver virtualbox default
VBoxManage controlvm "default" natpf1 "mongo,tcp,127.0.0.1,27017,,27017"
VBoxManage controlvm "default" natpf1 "redis,tcp,127.0.0.1,6379,,6379"
```

#### Run

```
docker-compose -f docker-compose-db.yml run -d
go get github.com/tools/godep
godep restore
go run run.go
```

Open http://localhost:8080/?keyword=KEYWORD in your browser.

**Option(Throw fixture data)**

```
./fixtures/fixtures.sh
```
