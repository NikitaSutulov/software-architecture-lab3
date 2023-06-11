#!/bin/bash
curl -X POST http://localhost:17000 -d "white"
curl -X POST http://localhost:17000 -d "bgrect 0.25 0.25 0.75 0.75"
curl -X POST http://localhost:17000 -d "figure 0.5 0.5"
curl -X POST http://localhost:17000 -d "green"
curl -X POST http://localhost:17000 -d "figure 0.6 0.6"
curl -X POST http://localhost:17000 -d "update"
