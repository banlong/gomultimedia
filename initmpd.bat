::Worked - Produce dash video from a file, appendable: run 2 times will create playlist 2 time (I do not use -add here)
::MP4Box -dash 3000 -mpd-refresh 3 -profile live -rap -dash-ctx stream.txt  -out qh.mpd videos/quangha.mp4#video videos/quangha.mp4#audio

::Issue - No Video
::MP4Box -dash 3000 -mpd-refresh 3 -profile live -rap -segment-name %s_ -dash-ctx stream.txt -out bn.mpd videos/Bunny.mp4#video videos/Bunny.mp4#audio

::Issue - No Sound
::MP4Box -dash 3000 -mpd-refresh 3 -profile live -rap -segment-name %s_ -dash-ctx stream.txt -out bn.mpd videos/Bunny.mp4#video -add videos/Bunny.mp4#audio

::Worked - Name the segment
::MP4Box -dash 3000 -mpd-refresh 3 -profile live -rap -segment-name %s_ -dash-ctx stream.txt -out bn.mpd -add videos/Bunny.mp4#video -add videos/Bunny.mp4#audio

::Worked - Name the segment start with bunny
::MP4Box -dash-ctx stream.txt -dash 3000 -mpd-refresh 3 -profile live -segment-ext mp4 -out bn.mpd -add videos/Bunny.mp4#video -add videos/Bunny.mp4#audio -fps 30 bunny

::Worked
::MP4Box -dash-ctx videos/dashout/stream.txt -dash 3000 -mpd-refresh 3 -profile live -segment-ext mp4 -out videos/dashout/qh.mpd -add videos/quangha.mp4#video -add videos/quangha.mp4#audio -fps 30 sample