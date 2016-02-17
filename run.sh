#!/usr/bin/env bash

beaconwatcher &
simulator 1 | logwatcher &
