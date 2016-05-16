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
