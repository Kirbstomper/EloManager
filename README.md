# EloManager
Elo Management in Go


This is to provide a backend to manging players and their Elo Rating

Using this to help players better understand where they rank in comparison to others in their area for various fighting games played
between the major cities in north carolina


Using Litestream [https://litestream.io/] to handle auto replication of the sqlite database to a s3 compatable bucket

Using s6-overlay[https://github.com/just-containers/s6-overlay] to run both my applicationa and litestream in a single docker container

Running the backend as a docker container
- Build
docker build -t elomapper .

Then you should make your environmental variables availible

export LITESTREAM_ACCESS_KEY_ID=xxxxxxxxxxxxxx
export LITESTREAM_SECRET_ACCESS_KEY=xxxxxxxxxx
export REPLICA_URL=xxxxxxxx
- Run
docker run \
  -p 8080:8080 \
  -v ${PWD}:/data \
  -e REPLICA_URL \
  -e LITESTREAM_ACCESS_KEY_ID \
  -e LITESTREAM_SECRET_ACCESS_KEY \
  elomapper