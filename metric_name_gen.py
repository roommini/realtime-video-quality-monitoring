txt = "frame|media_type=video|stream_index=0|key_frame=1|pkt_pts=1152267000|pkt_pts_time=12802.966667|pkt_dts=1152267000|pkt_dts_time=12802.966667|best_effort_timestamp=1152267000|best_effort_timestamp_time=12802.966667|pkt_duration=3000|pkt_duration_time=0.033333|pkt_pos=134796|pkt_size=614400|width=640|height=480|pix_fmt=yuv422p|sample_aspect_ratio=1:1|pict_type=I|coded_picture_number=0|display_picture_number=0|interlaced_frame=0|top_field_first=0|repeat_pict=0|color_range=unknown|color_space=unknown|color_primaries=unknown|color_transfer=unknown|chroma_location=unspecified|tag:lavfi.signalstats.YMIN=2|tag:lavfi.signalstats.YLOW=36|tag:lavfi.signalstats.YAVG=125.403|tag:lavfi.signalstats.YHIGH=204|tag:lavfi.signalstats.YMAX=242|tag:lavfi.signalstats.UMIN=112|tag:lavfi.signalstats.ULOW=119|tag:lavfi.signalstats.UAVG=127.015|tag:lavfi.signalstats.UHIGH=136|tag:lavfi.signalstats.UMAX=141|tag:lavfi.signalstats.VMIN=119|tag:lavfi.signalstats.VLOW=125|tag:lavfi.signalstats.VAVG=128.439|tag:lavfi.signalstats.VHIGH=132|tag:lavfi.signalstats.VMAX=137|tag:lavfi.signalstats.SATMIN=0|tag:lavfi.signalstats.SATLOW=2|tag:lavfi.signalstats.SATAVG=6.11771|tag:lavfi.signalstats.SATHIGH=10|tag:lavfi.signalstats.SATMAX=16|tag:lavfi.signalstats.HUEMED=134|tag:lavfi.signalstats.HUEAVG=165.054|tag:lavfi.signalstats.YDIF=1.98616|tag:lavfi.signalstats.UDIF=0.225208|tag:lavfi.signalstats.VDIF=0.0793099|tag:lavfi.signalstats.YBITDEPTH=8|tag:lavfi.signalstats.UBITDEPTH=8|tag:lavfi.signalstats.VBITDEPTH=8|tag:lavfi.signalstats.TOUT=0.00031901|tag:lavfi.signalstats.VREP=0.11875|tag:lavfi.signalstats.BRNG=2.27865e-05|tag:lavfi.deflicker.luminance=125.402771|tag:lavfi.deflicker.new_luminance=120.826736|tag:lavfi.deflicker.relative_change=-0.036491|tag:lavfi.bitplanenoise.0.1=0.113685|tag:lavfi.bitplanenoise.1.1=0.043893|tag:lavfi.bitplanenoise.2.1=0.022474"
def get_metrics(txt=txt):
    metrics = []
    lst = txt.split('|')
    for metric in lst[29:]:
        name = metric.split('=')[0]
        if name.startswith('tag:lavfi.bitplanenoise'):
            name = name[10:].replace('.','')
        else:
            name = name.split('.')[-1]
        metrics.append(name.lower())
    return metrics

def create_struct(metrics):
    for metric in metrics:
        print("%sMetric *prometheus.Desc"%metric)

def create_collector(metrics):
    for metric in metrics:
        print('%sMetric: prometheus.NewDesc("%s","%s metric",nil, nil,),'%(metric,metric,metric))

def create_ch(metrics):
    for metric in metrics:
        print("ch <- collector.%sMetric"%metric)

def update_ch(metrics):
    for i in range(len(metrics)):
        print('ch <- prometheus.MustNewConstMetric(collector.%sMetric, prometheus.GaugeValue, metrics[%s])'%(metrics[i],i))
        
metrics = get_metrics()
#create_struct(metrics)
#create_collector(metrics)
#create_ch(metrics)
update_ch(metrics)