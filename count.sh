#!/bin/bash
count=$(find "$(pwd)" -type f | wc -l)
printf "\t\vTotal files * 5: %d\v\n" $(($count * 5))