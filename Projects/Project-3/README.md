# Here we create a Python script to Download videos and audio from YouTube

### Dependancies: 
 <details>
 <summary>youtube-dl and click</summary>
 - open a terminal with admin privlages and run <code>pip install youtube-dl click</code>
 </details> 

 <details>
 <summary> <a href="https://ffbinaries.com/downloads">FFbinaries</a> </summary>
 - Downlaod <code>ffmeg</code>, <code>ffprobe</code> and <code>ffplay</code> from the above link<br>
 - Extract the zip files and sotre the <code>exe</code> files in the root directory for the project
</details>

## Running the code  
open a terminal in the root directory for the project  
***Replace video link with the link of the video you want to download***  
Run `.\Downlaoder.py downlaod-video "video link"` to downlaod a video  
Run `.\Downlaoder.py download-audio "video link"` to downlaod the audio for a video  
Run `.\Downlaoder.py downlaod-audo --keep-video "video link"` to keep the video and download the audio for a video  
