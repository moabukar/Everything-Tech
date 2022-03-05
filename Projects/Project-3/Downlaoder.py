import click
import youtube_dl

@click.group()
def apis():
    """A CLI for getting transcriptions of YouTube videos"""

@click.argument('link')
@apis.command()
def download_video(link): #function to downlaod videos, run the file adding downlaod-video "youtube video link"
   ydl_opts = {
       'format': 'mp4',
       'outtmpl': "./%(id)s.%(ext)s",
   }
   _id = link.strip()
   meta = youtube_dl.YoutubeDL(ydl_opts).extract_info(_id)
   save_location = meta['id'] + ".mp4"
   print(save_location)
   return save_location

@click.argument('link')
@click.option('-k', '--keep-video', is_flag=True, help="Pass this to keep the video")
@apis.command()
def download_audio(link, keep_video): #function to download audio, run the file adding downlaod-audio (add --keep-video if you want to download the video too) "youtube video link"
   ydl_opts = {
       'format': 'bestaudio/best',
       'postprocessors': [{
           'key': 'FFmpegExtractAudio',
           'preferredcodec': 'mp3',
           'preferredquality': '192',
       }],
       'ffmpeg-location': './',
       'outtmpl': "./%(id)s.%(ext)s",
       'keepvideo': 'True' if keep_video else 'False'
   }
   _id = link.strip()
   meta = youtube_dl.YoutubeDL(ydl_opts).extract_info(_id)
   save_location = meta['id'] + ".mp3"
   print(save_location)
   return save_location

def main():
    apis(prog_name='apis')

if __name__ == '__main__':
    main()