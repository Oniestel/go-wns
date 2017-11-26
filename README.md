# go-wns
Simple Go package for working with Windows Push Notification Services (WNS).

## Example
```
client := wns.Client{}
client.Init("{Package Security Identifier}", "{Client Secret}")
notification := wns.NewToast().SetTemplate("ToastText02").
  AddText("1", "Title").
  AddText("2", "Text")

success, err := client.Send("{Channel URL}", notification)
if err != nil {
  log.Println(err.Code, err.Message)
}
```

You can find more templates [here](https://msdn.microsoft.com/en-us/library/windows/apps/hh761494).
