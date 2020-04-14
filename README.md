# Webcam checker

Quick and easy way to check if your webcam is on. OSX only

```
ctx := comtext.Background()
isOn, err := webcamchecker.IsWebcamOn(ctx)
if err != nil {
  panic(err)
}
fmt.Println("webcam is", isOn)
```