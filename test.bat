::THIS WORK WELL WITH OSMO4 PLAYER, produce only one segment
::MP4Box -dash-ctx stream.txt -dash 3000 -segment-ext mp4 -out sample.mpd -add videos/trcd/001.mp4#video -add videos/trcd/001.mp4#audio -fps 30 sample.mp4
::MP4Box -dash-ctx stream.txt -dash 3000 -segment-ext mp4 -out sample.mpd -add videos/trcd/002.mp4#video -add videos/trcd/002.mp4#audio -fps 30 sample.mp4
::MP4Box -dash-ctx stream.txt -dash 3000 -segment-ext mp4 -out sample.mpd -add videos/trcd/003.mp4#video -add videos/trcd/003.mp4#audio -fps 30 sample.mp4

::ADD LIVE PROFILE -->PRODUCE 12 SEGMENTS, LOST SOUND AT THE END
::MP4Box -dash-ctx stream.txt -dash 3000 -profile live -segment-ext mp4 -out sample.mpd -add videos/trcd/001.mp4#video -add videos/trcd/001.mp4#audio -fps 30 sample.mp4
::MP4Box -dash-ctx stream.txt -dash 3000 -profile live -segment-ext mp4 -out sample.mpd -add videos/trcd/002.mp4#video -add videos/trcd/002.mp4#audio -fps 30 sample.mp4
::MP4Box -dash-ctx stream.txt -dash 3000 -profile live -segment-ext mp4 -out sample.mpd -add videos/trcd/003.mp4#video -add videos/trcd/003.mp4#audio -fps 30 sample.mp4

::ADD LIVE PROFILE -->PRODUCE 12 SEGMENTS, LOST SOUND in the Middle
::MP4Box -dash-ctx stream.txt -dash 3000 -profile live -segment-ext mp4 -out sample.mpd -add videos/trcd/001.mp4#video -add videos/trcd/001.mp4#audio -fps 30 sample.mp4
::MP4Box -dash-ctx stream.txt -dash 3000 -profile live -segment-ext mp4 -out sample.mpd -add videos/trcd/002.mp4#video -add videos/trcd/002.mp4#audio -fps 30 sample.mp4
::MP4Box -dash-ctx stream.txt -dash 3000 -profile live -segment-ext mp4 -out sample.mpd -add videos/trcd/003.mp4#video -add videos/trcd/003.mp4#audio -fps 30 sample.mp4


::MP4Box -dash-ctx stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out sample.mpd -add videos/trcd/001.mp4#video -add videos/trcd/001.mp4#audio -fps 30 sample.mp4
::MP4Box -dash-ctx stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out sample.mpd -add videos/trcd/002.mp4#video -add videos/trcd/002.mp4#audio -fps 30 sample.mp4
:;MP4Box -dash-ctx stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out sample.mpd -add videos/trcd/003.mp4#video -add videos/trcd/003.mp4#audio -fps 30 sample.mp4


::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0000.mp4#video bunnydash/0000.mp4#audio -fps 30 
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0001.mp4#video bunnydash/0001.mp4#audio -fps 30 
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0002.mp4#video bunnydash/0002.mp4#audio -fps 30 
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0003.mp4#video bunnydash/0003.mp4#audio -fps 30 
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0004.mp4#video bunnydash/0004.mp4#audio -fps 30 
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0005.mp4#video bunnydash/0005.mp4#audio -fps 30 
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0006.mp4#video bunnydash/0006.mp4#audio -fps 30 
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0007.mp4#video bunnydash/0007.mp4#audio -fps 30 
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0008.mp4#video bunnydash/0008.mp4#audio -fps 30 
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0009.mp4#video bunnydash/0009.mp4#audio -fps 30 
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0010.mp4#video bunnydash/0010.mp4#audio -fps 30 
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0011.mp4#video bunnydash/0011.mp4#audio -fps 30 
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0012.mp4#video bunnydash/0012.mp4#audio -fps 30 


::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0000.mp4 bunnydash/0000.mp4 -fps 30
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0001.mp4 bunnydash/0001.mp4 -fps 30
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0002.mp4 bunnydash/0002.mp4 -fps 30
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0003.mp4 bunnydash/0003.mp4 -fps 30
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0004.mp4 bunnydash/0004.mp4 -fps 30
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0005.mp4 bunnydash/0005.mp4 -fps 30
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0006.mp4 bunnydash/0006.mp4 -fps 30
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0007.mp4 bunnydash/0007.mp4 -fps 30
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0008.mp4 bunnydash/0008.mp4 -fps 30
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0009.mp4 bunnydash/0009.mp4 -fps 30
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0010.mp4 bunnydash/0010.mp4 -fps 30
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0011.mp4 bunnydash/0011.mp4 -fps 30
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile dashavc264:live -segment-ext mp4 -out bunnydash/bunny.mpd bunnydash/0012.mp4 bunnydash/0012.mp4 -fps 30


::THIS GENERATE SEGMENT BUT THE MPD CONTAIN THE SEGMENT NAME OF THE LAST SEGMENT, NEED TO SPECIFY THE SEGMENT NAME
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile live  -fps 30 -out bunnydash/bunny.mpd bunnydash/0000.mp4#video bunnydash/0000.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile live  -fps 30 -out bunnydash/bunny.mpd bunnydash/0001.mp4#video bunnydash/0001.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile live  -fps 30 -out bunnydash/bunny.mpd bunnydash/0002.mp4#video bunnydash/0002.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile live  -fps 30 -out bunnydash/bunny.mpd bunnydash/0003.mp4#video bunnydash/0003.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile live  -fps 30 -out bunnydash/bunny.mpd bunnydash/0004.mp4#video bunnydash/0004.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile live  -fps 30 -out bunnydash/bunny.mpd bunnydash/0005.mp4#video bunnydash/0005.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile live  -fps 30 -out bunnydash/bunny.mpd bunnydash/0006.mp4#video bunnydash/0006.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile live  -fps 30 -out bunnydash/bunny.mpd bunnydash/0007.mp4#video bunnydash/0007.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile live  -fps 30 -out bunnydash/bunny.mpd bunnydash/0008.mp4#video bunnydash/0008.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile live  -fps 30 -out bunnydash/bunny.mpd bunnydash/0009.mp4#video bunnydash/0009.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile live  -fps 30 -out bunnydash/bunny.mpd bunnydash/0010.mp4#video bunnydash/0010.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile live  -fps 30 -out bunnydash/bunny.mpd bunnydash/0011.mp4#video bunnydash/0011.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -profile live  -fps 30 -out bunnydash/bunny.mpd bunnydash/0012.mp4#video bunnydash/0012.mp4#audio

::SAME AS ABOVE
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000   -profile live -rap -single-file -fps 30 -out bunnydash/bunny.mpd bunnydash/0000.mp4#video bunnydash/0000.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000   -profile live -rap -single-file -fps 30 -out bunnydash/bunny.mpd bunnydash/0001.mp4#video bunnydash/0001.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000   -profile live -rap -single-file -fps 30 -out bunnydash/bunny.mpd bunnydash/0002.mp4#video bunnydash/0002.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000   -profile live -rap -single-file -fps 30 -out bunnydash/bunny.mpd bunnydash/0003.mp4#video bunnydash/0003.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000   -profile live -rap -single-file -fps 30 -out bunnydash/bunny.mpd bunnydash/0004.mp4#video bunnydash/0004.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000   -profile live -rap -single-file -fps 30 -out bunnydash/bunny.mpd bunnydash/0005.mp4#video bunnydash/0005.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000   -profile live -rap -single-file -fps 30 -out bunnydash/bunny.mpd bunnydash/0006.mp4#video bunnydash/0006.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000   -profile live -rap -single-file -fps 30 -out bunnydash/bunny.mpd bunnydash/0007.mp4#video bunnydash/0007.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000   -profile live -rap -single-file -fps 30 -out bunnydash/bunny.mpd bunnydash/0008.mp4#video bunnydash/0008.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000   -profile live -rap -single-file -fps 30 -out bunnydash/bunny.mpd bunnydash/0009.mp4#video bunnydash/0009.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000   -profile live -rap -single-file -fps 30 -out bunnydash/bunny.mpd bunnydash/0010.mp4#video bunnydash/0010.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000   -profile live -rap -single-file -fps 30 -out bunnydash/bunny.mpd bunnydash/0011.mp4#video bunnydash/0011.mp4#audio
::MP4Box -dash-ctx bunnydash/stream.txt -dash 3000   -profile live -rap -single-file -fps 30 -out bunnydash/bunny.mpd bunnydash/0012.mp4#video bunnydash/0012.mp4#audio


MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -mpd-refresh 3 -profile live -single-file -out bunnydash/bun.mpd -add bunnydash/0000.mp4#video -add bunnydash/0000.mp4#audio -fps 30 bun
MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -mpd-refresh 3 -profile live -single-file -out bunnydash/bun.mpd -add bunnydash/0001.mp4#video -add bunnydash/0001.mp4#audio -fps 30 bun
MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -mpd-refresh 3 -profile live -single-file -out bunnydash/bun.mpd -add bunnydash/0002.mp4#video -add bunnydash/0002.mp4#audio -fps 30 bun
MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -mpd-refresh 3 -profile live -single-file -out bunnydash/bun.mpd -add bunnydash/0003.mp4#video -add bunnydash/0003.mp4#audio -fps 30 bun
MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -mpd-refresh 3 -profile live -single-file -out bunnydash/bun.mpd -add bunnydash/0004.mp4#video -add bunnydash/0004.mp4#audio -fps 30 bun
MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -mpd-refresh 3 -profile live -single-file -out bunnydash/bun.mpd -add bunnydash/0005.mp4#video -add bunnydash/0005.mp4#audio -fps 30 bun
MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -mpd-refresh 3 -profile live -single-file -out bunnydash/bun.mpd -add bunnydash/0006.mp4#video -add bunnydash/0006.mp4#audio -fps 30 bun
MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -mpd-refresh 3 -profile live -single-file -out bunnydash/bun.mpd -add bunnydash/0007.mp4#video -add bunnydash/0007.mp4#audio -fps 30 bun
MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -mpd-refresh 3 -profile live -single-file -out bunnydash/bun.mpd -add bunnydash/0008.mp4#video -add bunnydash/0008.mp4#audio -fps 30 bun
MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -mpd-refresh 3 -profile live -single-file -out bunnydash/bun.mpd -add bunnydash/0009.mp4#video -add bunnydash/0009.mp4#audio -fps 30 bun
MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -mpd-refresh 3 -profile live -single-file -out bunnydash/bun.mpd -add bunnydash/0010.mp4#video -add bunnydash/0010.mp4#audio -fps 30 bun
MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -mpd-refresh 3 -profile live -single-file -out bunnydash/bun.mpd -add bunnydash/0011.mp4#video -add bunnydash/0011.mp4#audio -fps 30 bun
MP4Box -dash-ctx bunnydash/stream.txt -dash 3000 -mpd-refresh 3 -profile live -single-file -out bunnydash/bun.mpd -add bunnydash/0012.mp4#video -add bunnydash/0012.mp4#audio -fps 30 bun


MP4Box -dash-ctx factory/mpd/stream.txt -dash 3000 -segment-name test_ -profile h264:live -bs-switching no -rap -frag 3000 factory/input/seg.mp4#video factory/input/seg.mp4#audio -out factory/mpd/kd2.mpd

MP4Box -dash-ctx factory/mpd/stream.txt -dash 3000 -segment-name test_ -profile h264:live -bs-switching no -rap -frag 3000 factory/input/seg.mp4#video factory/input/seg.mp4#audio -out factory/mpd/kd2.mpd

MP4Box -dash-ctx factory/mpd/stream.txt -dash 3000 -segment-name test_ -profile live -bs-switching no -rap -frag 3000 factory/input/seg.mp4#video factory/input/seg.mp4#audio -out factory/mpd/kd2.mpd

MP4Box -dash-ctx factory/mpd/stream.txt -dash 3000 -profile live -single-segment -single-file -bs-switching no -rap -frag 3000 factory/input/seg.mp4#video factory/input/seg.mp4#audio -out factory/mpd/kd2.mpd

MP4Box -dash-ctx factory/mpd/stream.txt -dash 3000 -profile live -url-template -bs-switching no -rap -frag 3000 factory/input/seg.mp4#video factory/input/seg.mp4#audio -out factory/mpd/kd2.mpd


MP4Box -dash 3000 -rap -frag-rap -profile dashavc264:onDemand -out asia.mpd videos/asia.mp4
mp4fragment videos/asia.mp4 videos/asia-frag.mp4
mp4split --media-segment %04llu.mp4 asia_dashinit.mp4
mp4split asia_dashinit.mp4




ffmpeg -i videos/asia-frag.mp4 -acodec copy -vcodec copy -force_key_frames "expr:gte(t,n_forced*2)" -segment_time 3 -f segment factory/segments/%04d.mp4

ffmpeg -i videos/asia-frag.mp4 -acodec copy -vcodec copy -force_key_frames  expr:gte(t,n_forced*GOP_LEN_IN_SECONDS) -segment_time 3 -f segment factory/segments/%04d.mp4

ffmpeg -i videos/asia-frag.mp4 -acodec copy -vcodec copy -x264opts keyint=25:min-keyint=25:scenecut=-1 -segment_time 3 -f segment factory/segments/%04d.mp4

ffmpeg -i videos/asia-frag.mp4 -acodec copy -vcodec copy -force_key_frames keyint=25:min-keyint=25:scenecut=-1 -segment_time 3 -f segment factory/segments/%04d.mp4

//fps = 30, 3s-->90f
ffmpeg -i videos/asia-frag.mp4 -acodec copy -vcodec copy -g 90 -segment_time 3 -f segment factory/segments/%04d.mp4


ffmpeg -i videos/asia-frag.mp4 -acodec copy -vcodec copy -force_key_frames  expr:gte(t,n_forced*3) -r 24 -segment_time 3 -f segment factory/segments/%04d.mp4

ffmpeg -i videos/asia-frag.mp4 -codec copy -map 0 -f segment -segment_list out.csv -segment_times 1,2,3,5,8,13,21 factory/segments/%04d.mp4

ffmpeg -i videos/asia-frag.mp4 -vcodec libx264 -preset veryslow -x264-params keyint=90:no-scenecut=1 -acodec copy videos/asia-copy.mp4

ffmpeg -i videos/asia.mp4 -vcodec copy -preset ultrafast -x264-params keyint=90:no-scenecut=1 -acodec copy videos/asia-copy.mp4

ffmpeg -i videos/asia.mp4 -vcodec copy -preset ultrafast -x264-params keyint=90:no-scenecut=1 -acodec copy videos/asia-copy.mp4



ffmpeg -i videos/asia.mp4 -acodec copy -vcodec copy -force_key_frames  expr:gte(t,n_forced*3) videos/asia-copy.mp4

ffmpeg -i videos/asia-copy.mp4 -vcodec copy -acodec copy -preset ultrafast -f segment -segment_times 3 factory/segments/%04d.mp4


ffmpeg -i videos/asia-copy.mp4 -acodec copy -vcodec copy -force_key_frames  expr:gte(t,n_forced*3) -segment_time 3 -f segment factory/segments/%04d.mp4

MP4Box -splits 3 videos/asia.mp4

ffmpeg -i videos/asia-frag.mp4 -codec copy -map 0 -segment_list list.txt -segment_times 3 -f segment factory/segments/%04d.mp4

//This split script produce the list.txt contains all the file name
ffmpeg -i videos/asia-copy.mp4 -acodec copy -vcodec copy -segment_list list.txt -force_key_frames  expr:gte(t,n_forced*1) -segment_time 1 -f segment factory/segments/%04d.mp4

ffmpeg -i videos/asia.mp4 -acodec copy -vcodec copy -segment_list list.txt -force_key_frames  expr:gte(t,n_forced*1) -segment_time 1 -f segment factory/segments/%04d.mp4

MP4Box -dash 3000 -rap -frag-rap  -profile dashavc264:onDemand -out videos/asia-copy.mp4 videos/asia.mp4

ffmpeg -i videos/asia-frag.mp4 -acodec copy -vcodec copy -segment_list list.txt -segment_time 3 -f segment factory/segments/%04d.mp4

same result
MP4Box -split 3 videos/duck.mp4 -out factory/segments/%s.mp4
MP4Box -split 3 videos/duck.mp4 -out factory/segments/%d.mp4
MP4Box -split 3 videos/duck.mp4 -out factory/segments/ -segment-name hello

MP4Box -dash 4000 -frag 4000 -rap -segment-name segment_ videos/asia.mp4