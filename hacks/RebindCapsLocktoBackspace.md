# Rebind the Caps Lock key to Backspace

* Start Automator
* Select Application
* From Library select `Run Shell Scripts`
* Past the following:
```sh
hidutil property --set '{"UserKeyMapping":[{"HIDKeyboardModifierMappingSrc":0x700000039,"HIDKeyboardModifierMappingDst":0x70000002A}]}'
```
* Test it by running `Run`
* Save it, e.g. iCloud/Automator/Rebind-Backspace-to-CapsLock
* System Preferences --> Users and Groups --> + (add it from iCloud/Automator/Rebind-Backspace-to-CapsLock)
* Test by rebooting 
