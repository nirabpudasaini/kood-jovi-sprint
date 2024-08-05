#!/bin/bash
touch a \! \\ \"
mkdir \`
cp \! \`/\!
if  [ "$MOVE_A" == yes ]; then
    mv a \`/a
 elif [ "$MOVE_A" == no ]; then
    rm a
 else
    echo "MOVE_A variable not recognized"
fi