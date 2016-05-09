::Do not work for avi
::ffmpeg -i videos/avi/s02.avi -reset_timestamps 1 -codec: copy -map 0 -segment_time 3 -f segment avisplit/%04d.avi

::Do not work for avi
::ffmpeg -i videos/avi/s02.avi -codec: copy -map 0 -segment_time 3 -f segment avisplit/%04d.avi

::Work for bars_100.avi
::ffmpeg -i videos/avi/bars_100.avi -codec: copy -map 0 -segment_time 3 -f segment avisplit/%04d.avi
::But does not work for s02.avi, 6min video put in one segment
::ffmpeg -report -i videos/avi/s02.avi -codec: copy -map 0 -segment_time 3 -f segment avisplit/%04d.avi
::Not split correctly, 6min video put in one segment
::ffmpeg -i videos/avi/s02.avi -c copy -map 0 -segment_time 3 -f segment avisplit/%04d.avi

::SPLIT AND ENCODING AT A TIME (s02 has error in the audio stream mp3)
::Workes, 125 segments created but this also including transcoding,
::In this case,convert to video H264 and audio mp3, take long time to process
::ffmpeg -i videos/avi/s02.avi -c:v h264 -c:a mp3 -segment_time 3 -f segment avisplit/%04d.avi
::We can directly convert it to MP4 segments because it was transcoded
::ffmpeg -i videos/avi/s02.avi -c:v h264 -c:a mp3 -segment_time 3 -f segment avisplit/%04d.mp4
::We can also define the aac encoder for audio
::ffmpeg -i videos/avi/s02.avi -c:v h264 -c:a aac -segment_time 3 -f segment avisplit/%04d.mp4

::SPLIT - Create only one segment 0000.avi, not split
::ffmpeg -i videos/avi/s02.avi -c:v copy -c:a copy -segment_time 3 -f segment avisplit/%04d.avi

::SOLUTION IS TO MAINTAIN THE SAME CONTAINER WHEN SPLITTING TO AVOID OF ANY INVALID DATA

::Worked - cut 30s, the difficulty when using this command is ffmpeg will cut the video using key frame
::therefore, sometimes it will not create exactly 3 seconds segments and overlap segments
::ffmpeg -i videos/avi/s02.avi -acodec copy -vcodec copy -ss 00:00:00 -t 00:00:30 output.avi

::Time 19s, output size 100MB(huge increase), quality: same, audio & video not sync
::ffmpeg -ss 00:06:27 -t 3 -i videos/avi/s02.avi -c:v copy -c:a copy -y avisplit/s02-0130.avi

::Time 11s, output size 105MB(huge increase), quality: same, audio & video not sync, overlap
::ffmpeg -ss 00:06:27 -t 3 -i videos/avi/s02.avi -c:v copy -c:a copy -y avisplit/s02-0130.avi

::FORCE KEY FRAME each 3 seconds
::ffmpeg -i videos/avi/s02.avi -force_key_frames "expr:gte(t,n_forced*3)" videos/avi/s02b.avi

::---------------------------------------------------------------------------------------------------------------
::WORKED: can use this one to segment the video, however the quality is WORST 57MB --> 32MB (reduce in size), time: 28s
::ffmpeg -ss 00:00:00 -t 3 -i videos/avi/s02.avi -y avisplit/s02-0001.avi

::WORK: this one work as above and have a same quality, the only different is it cost 1'30 to split & transcode, 43.6MB ouput
::ffmpeg -ss 00:00:00 -t 3 -i videos/avi/s02.avi -c:v libx264 -y avisplit/s02-0001.avi

::WORK: this one work as above and have a same quality, the only different is it cost 1'33 to split & transcode, 43.6MB ouput
::This is the best quality I have, segment played without blink, segments can be put in MP4(bc transcoded)
::ffmpeg -ss 00:00:00 -t 3 -i videos/avi/s02.avi -c:v libx264 -c:a aac -y avisplit/s02-0001.avi
::---------------------------------------------------------------------------------------------------------------
