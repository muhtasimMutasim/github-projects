#!/usr/bin/env python

import youtube_dl
import os, sys
import argparse


def verify_dir(path_to_dir: str):
    """ Verifies wether directory exists or not. """
    isdir = os.path.isdir(path_to_dir) 
    #print(f"\n\nDoes directory exist: {isdir}\n\n")
    if isdir == True:
        return path_to_dir
    else:
        raise FileExistsError(
            'The desired directory for download does not exist. Check Directory again.')


def arg_parser():
    """ Returns all the arguments used. """
    parser = argparse.ArgumentParser()

    # Add an argument
    parser.add_argument('-u', '--url', type=str, required=True, help='The url of the video.')
    parser.add_argument('-f', '--file', type=str, required=True, help='The name of the downloaded audio file.')
    parser.add_argument('-d', '--dir', type=str, required=False, help='Directory to download the file.')
    parser.add_argument('--video', action='store_true', default=False,required=False, help='Directory to download the file.')
    
    args = parser.parse_args()
    #print(f"\n\nAll args: \n{args}\n\n")
    return args


def run(video_url: str = None, path_to_dir: str = None, 
        filename: str= None, download_val: bool = False):
    """ Runs logic to download the video as .mp3 file. """

    if video_url == None:
        raise ValueError('The video_url cannot be a "None" value, needs to be URL.')

    filename = f"{filename}.m4a"

    if path_to_dir != None:
        filename = f"{path_to_dir}/{filename}"

    video_info = youtube_dl.YoutubeDL().extract_info(
        url = video_url, download=download_val
    )
    options={
        'format':'bestaudio/best',
        'postprocessors': [{
        'key': 'FFmpegExtractAudio',
        'preferredcodec': 'mp3',
        'preferredquality': '192',
        }],
        'keepvideo':download_val,
        'outtmpl':filename,
    }

    with youtube_dl.YoutubeDL(options) as ydl:
        ydl.download([video_info['webpage_url']])

    print("Download complete... {}".format(filename))


def main():
    
    args = arg_parser()

    url = args.url
    file = args.file
    a_dir = verify_dir(args.dir)
    video_opt = args.video
    
    run(video_url=url, filename=file, path_to_dir=a_dir, download_val=video_opt)


if __name__=='__main__':
    if int(sys.version_info[0]) < 3:
        raise Exception('Version of Python is not 3.8! This script needs to have python 3.8 or higher.')
    if int(sys.version_info[1]) < 8:
        raise Exception('Version of Python is not 3.8! This script needs to have python 3.8 or higher.')
    
    main()
