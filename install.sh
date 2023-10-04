#!/bin/bash

echo "██████╗░██╗░█████╗░██╗░░██╗"
echo "██╔══██╗██║██╔══██╗╚██╗██╔╝"
echo "██║░░██║██║██║░░╚═╝░╚███╔╝░"
echo "██║░░██║██║██║░░██╗░██╔██╗░"
echo "██████╔╝██║╚█████╔╝██╔╝╚██╗"
echo "╚═════╝░╚═╝░╚════╝░╚═╝░░╚═╝"

echo -e "\n[*] Checking package requirements...\n"

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

echo -e "\n[*] Setting up Dicx...\n"

go build

sudo mv ./dicx /usr/local/bin

sudo cp ./icons/dicx.png /usr/share/icons/dicx.png
sudo cp ./icons/dicx-failed.png /usr/share/icons/dicx-failed.png 

echo -e "Dicx installed successfully!\n"

echo -e "\nYou can now setup your preferred shortcut in your system to call dicx command\n"
