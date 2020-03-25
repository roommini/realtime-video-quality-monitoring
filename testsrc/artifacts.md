# Video Artifacts

## raw
**command**: ```ffplay -f lavfi testsrc```

<img src="./images/raw.gif"></img>

## flicker
**command**: ```ffplay -f lavfi testsrc,vignette='PI/30+random(10)*PI/5':eval=frame```

<img src="./images/flicker.gif"></img>

## noise
**command**: ```ffplay -f lavfi testsrc,noise=c0s=90:allf=t```

<img src="./images/noise.gif"></img>

<img src="./images/noise.png"></img>

- Good saperation between random noise, flicker and raw
- flicker variance is high

## saturation
**command**: ```ffplay -f lavfi testsrc,hue='s=sin(2*PI*t^0.5)'```

<img src="./images/saturation.gif"></img>

<img src="./images/satavg.png"></img>

- The sine signal is visable!

## blur
**command**: ```ffplay -f lavfi testsrc,unsharp=7:7:-2:7:7:-2```

<img src="./images/blur.gif"></img>

## bright
**command**: ```ffplay -f lavfi testsrc,lutyuv='y=2*val'```

<img src="./images/bright.gif"></img>

<img src="./images/yavg.png"></img>

## dark
**command**: ```ffplay -f lavfi testsrc,lutyuv='y=val/2'```

<img src="./images/dark.gif"></img>

## snippet for playing with multiple artifacts
```bash
ffplay -f lavfi -i testsrc -vf "split=4[a][b][c][d];[b]vignette='PI/30+random(10)*PI/5':eval=frame[x];[c]noise=c0s=90:allf=t[y];[d]hue='s=sin(2*PI*t^0.5)'[z];[a][x][y][z]hstack=4"
```
