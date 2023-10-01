#!/bin/env bash

echo "██████╗░██╗░█████╗░██╗░░██╗"
echo "██╔══██╗██║██╔══██╗╚██╗██╔╝"
echo "██║░░██║██║██║░░╚═╝░╚███╔╝░"
echo "██║░░██║██║██║░░██╗░██╔██╗░"
echo "██████╔╝██║╚█████╔╝██╔╝╚██╗"
echo "╚═════╝░╚═╝░╚════╝░╚═╝░░╚═╝"

echo "\n [*] Checking package requirements... \n"

if which xclip >/dev/null 2>&1; then
        echo " xclip is installed"
else
        echo " Installing xclip..."
        sudo apt install xclip -y
        echo " xclip install successfully!"
fi

if which notify-send >/dev/null 2>&1; then
        echo " notify-send is installed"
else
        echo " Installing notify-send..."
        sudo apt install libnotify-bin -y
        echo " notify-send installed successfully!"
fi

if which go >/dev/null 2>&1; then
        echo " Golang is installed"
else
        echo " Installing golang..."
        sudo snap install go --classic
        echo " golang installed successfully!"
fi

echo "\n [*] Setting up Dicx... \n"

go build

sudo mv ./dicx /usr/local/bin

echo "Dicx installed successfully! \n"

echo "\n You can now setup your preferred shortcut in your system to call dicx command \n"
