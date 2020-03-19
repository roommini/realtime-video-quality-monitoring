# Real-time Video Quality Monitoring
## Deployment:
1. setup Grafana and prometheus by deploying the docker stack:
```bash
docker stack deploy rtvqm -c docker-stack.yml
```
2. build the exporter via:
```bash
go build main.go collector.go
```
3. stream webcam to ffprobe and pipe the results to the exporter via:
```bash
ffprobe -f lavfi movie=/dev/video0,signalstats="stat=tout+vrep+brng" -show_frames -print_format csv | ./main -port 2112

```
## snippets:
**capture stream from webcam:**
``` bash
ffmpeg -f v4l2 -framerate 25 -video_size 640x480 -i /dev/video0 -f mpegts udp://127.0.0.1:9999
```
**probe udp stream:**
``` bash
ffprobe -f lavfi movie="udp\\\://127.0.0.1\\\:9999",signalstats="stat=tout+vrep+brng",deflicker=bypass=1 -show_frames
```

**play captured stream:**
``` bash
ffplay udp://127.0.1:9999
```
**QCtool probbing command:**
```bash
ffprobe -f lavfi -i "movie=EXAMPLE.mov:s=v+a[in0][in1],[in0]signalstats=stat=tout+vrep+brng,cropdetect=reset=1:round=1,idet=half_life=1,deflicker=bypass=1,split[a][b];[a]field=top[a1];[b]field=bottom,split[b1][b2];[a1][b1]psnr[c1];[c1][b2]ssim[out0];[in1]ebur128=metadata=1,astats=metadata=1:reset=1:length=0.4[out1]" -show_frames -show_versions -of xml=x=1:q=1 -noprivate | gzip > EXAMPLE.mov.qctools.xml.gz
```

