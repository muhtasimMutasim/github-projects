# Youtube Video Downloader

Python script to download youtube videos, and convert youtube videos as mp3 files.


## Requirements
- Python 3.8 or higher.
- A software called [poetry](https://python-poetry.org/docs/)
- (Optional) A software called [pyenv](https://github.com/pyenv/pyenv)

## Usage

### To download only the audio use this:
download.py -u "Youtube url" -f "Name of file" -d /name/of/directory/for/download


### To download both audio and video use this:
download.py -u "Youtube url" -f "Name of file" -d /name/of/directory/for/download --video

- By adding the "--video" option the script downloads both the audio and video.
