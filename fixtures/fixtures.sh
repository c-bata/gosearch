#!/bin/bash

GOSEARCH_DB=$1

if [ -z "$GOSEARCH_DB" ]; then
  GOSEARCH_DB="gosearch"
fi

echo Target database is ${GOSEARCH_DB}

echo "Drop database"
mongo 127.0.0.1/${GOSEARCH_DB} --eval "db.dropDatabase()"

echo "Provide initial data via fixtures"
mongo 127.0.0.1/${GOSEARCH_DB} fixtures/*.js
