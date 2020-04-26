#!/bin/bash

# FILEPATH=$(readlink -f "canvas-cli")

# chmod +x canvas-t

# ln -s $FILEPATH $HOME/.local/bin

echo ""
echo "Would you like to generate config.yaml at ~/.config/canvas-tui? This will over write any config already present there. (y/n)"
read generate
if [ "$generate" == "y" ]
then

  if [[ ! -d "$HOME/.config" ]]
  then
    mkdir "$HOME/.config"
  fi

  CONFPATH=$HOME/.config/canvas-cli/

  if [[ ! -d $CONFPATH ]]
  then
    mkdir $CONFPATH
  fi

  echo "What is your canvas token? Leave this blank if you don't know and will edit the config later."
  read TOKEN

  echo "What is your canvas domain? Example: wwu's domain is https://wwu.instructure.com/"
  read DOMAIN

  echo "canvasdomain: \"$DOMAIN\"" >> $CONFPATH/config.yaml
  echo "canvastoken: \"$TOKEN\"" >> $CONFPATH/config.yaml

else
  echo No config generated.
fi

  echo Make sure $HOME/.local/bin is in your PATH.
