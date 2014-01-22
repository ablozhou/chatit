#!/bin/bash

for((i=0;i<500;i++)) {
    chatit client 10.10.10.4:9001 &
}
