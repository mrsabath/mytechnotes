#!/bin/bash

helpme()
{
  cat <<HELPMEHELPME

Syntax: ${0} <type> <n>
Where:
  type - deployer, planner or executor
  n    - 1 (first executor), 2 (second) etc [optional]

HELPMEHELPME
}

# validate the arguments
if [[ "$1" == "-?" || "$1" == "-h" || "$1" == "--help" || "$1" == "" ]] ; then
  helpme
  exit 1
fi

seq=$2

if [[ "$seq" == "" ]]; then
  seq=1
fi

kubectl logs -f $(kubectl  get po | grep update-$1 | awk '{print $1}' | sed -n ${seq}p)
