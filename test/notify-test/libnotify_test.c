#include<stdio.h>
#include<libnotify/notify.h>

//Compilation Command: gcc libnotify_test.c -o libnotify_test `pkg-config --cflags --libs gtk+-2.0` -lnotify

int main() {
	NotifyNotification *notify;
	notify_init("some-name");
	char *title = "TEST";
	char *message= "Test notification using libnotify!";
	notify = notify_notification_new(title, message, NULL);
	notify_notification_show(notify, NULL);
	return 0;
}
