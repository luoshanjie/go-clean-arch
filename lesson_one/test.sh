#!/usr/bin/env sh
curl -X POST http://localhost/api/v1/user/add \
     -H "Content-Type: application/json" \
     -d "{'name': 'tom', 'age': 10}"
