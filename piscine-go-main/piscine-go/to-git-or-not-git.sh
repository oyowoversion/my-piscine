#!/bin/bash

curl -s https://acad.learn2earn.org.ng/assets/superhero/all.json | jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'
