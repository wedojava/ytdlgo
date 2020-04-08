# ytdlgo

## Build on POSIX

```
cd scripts
sh ./build.sh
```

## Usage

### Download Once

#### Download youtube video links directly

```
ytdlgo
```
 - It will download videos from youtube group by the `./list.txt` contains
 - every line contains subfolder name and video page url, split by `|`

#### Download youtube video by playlist or channel

```
ytdlgo 2
```
 - It will scan `configs/channelmap.txt` and download all videos
 - `channelmap.txt` is split by `|`, before is subfolder name, after is channel or playlist link url.
 - every channel or playlist at one line.

### Download by task

```
ytdlgo 1
```
 - It will scan `configs/channelmap.txt` and download all videos
 - `channelmap.txt` is split by `|`, before is subfolder name, after is channel or playlist link url.
 - every channel or playlist at one line.
 - It will start scan at `**:15` every hour
