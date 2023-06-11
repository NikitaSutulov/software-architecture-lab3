#!/bin/bash
curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "white"
curl -X POST http://localhost:17000 -d "update"