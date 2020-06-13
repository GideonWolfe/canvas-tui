#!/bin/sh

# FILEPATH=$(readlink -f "canvas-cli")

# chmod +x canvas-tui

# ln -s $FILEPATH $HOME/.local/bin

echo ""
echo "Would you like to generate config.yaml at ~/.config/canvas-tui? This will over write any config already present there. [y/N]"
read -r generate
if [ "$generate" = "y" ]
then

  if [ ! -d "${XDG_CONFIG_HOME:-$HOME/.config}" ]
  then
    mkdir "${XDG_CONFIG_HOME:-$HOME/.config}"
  fi

  CONFPATH=${XDG_CONFIG_HOME:-$HOME/.config}/canvas-tui/

  if [ ! -d "$CONFPATH" ]
  then
    mkdir "$CONFPATH"
  fi

  echo "What is your canvas token? Leave this blank if you don't know and will edit the config later."
  read -r TOKEN

  echo "What is your canvas domain? Example: wwu's domain is https://wwu.instructure.com/"
  read -r DOMAIN

  echo "canvasdomain: \"$DOMAIN\"" >> "$CONFPATH/config.yaml"
  echo "canvastoken: \"$TOKEN\"" >> "$CONFPATH/config.yaml"

else
  echo No config generated.
fi

echo "Make sure $HOME/.local/bin is in your PATH."
