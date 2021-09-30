#!/bin/sh

R CMD Rserve --RS-source rserve.conf --no-save --RS-enable-remote

while true; do
  sleep 1
done
