# ChillCast 😎

ChillCast is a Twitch integration application designed to bring your streaming 
experience to the next level. With ChillCast, interacting with your Twitch 
community has never been easier or more fun!

For now, we just provide text to speech and profanity checking.

## Installation

To get ChillCast up and running on your machine, follow these steps:

```bash
git clone https://github.com/gravestench/chillcast
cd chillcast

go build ./cmd/chillcast

./chillcast
```

## Services & Configuration
The chillcast config files will be inside of `~/.config/chillcast`, where `~` is
your home folder. On windows, you home folder would be something like
`C:\Users\JohnDoe`. This means that chillcast will put config files in the 
directory `C:\Users\JohnDoe\.config\chillcast`

### Profanity Detection Service
The profanity detection service can be used to configure words to censor.
```
~/.config/chillcast/profanity_detection.json
```
```json
{
  "Profanity Detector": {
    "false-negatives (comma-delimited)": "",
    "false-positives (comma-delimited)": "",
    "profanities (comma-delimited)": ""
  }
}
```

### Soundboard service
The soundboard service can be used to trigger a pool of audio files in a 
specified directory. 

Each soundboard entry should be a copy of `Example #1` below, but with 
a unique name instead of `Example #1`.

The "literals" below are for a comma-separated list of literal strings to match 
against. This works the same for message triggers and audio file names.
Additionally, you can specify a regex pattern to match files in the specified 
directory.

although the full list of twitch events are given, almost all but the 
OnPrivateMessage event handling logic is actually implemented.
```
~/.config/chillcast/soundboard.json
```
```json
{
  "Example #1": {
    "audio filename literals": "fart1.wav,fart2.mp3,fart3.mp3",
    "audio filename regex": "fart\\d.(wav|mp3)",
    "audio files directory": "C:\\path\\to\\audio\\files",
    "enabled": false,
    "spam count": 3,
    "spam interval": "30s",
    "trigger literals": "fart",
    "trigger regex": "(butt|fart)",
    "twitch event": "OnPrivateMessage"
  },
  "_help": {
    "twitch event list (these are the possible events you can use to trigger)": "OnPrivateMessage"
  }
}
```
Here's an example config that matches many audio files in my `Downloads` folder:
```json
{
  "laugh": {
    "audio filename literals": "",
    "audio filename regex": ".*(laugh|lol).*\\.(wav|mp3)",
    "audio files directory": "/home/gravestench/Downloads",
    "enabled": true,
    "spam count": 5,
    "spam interval": "30s",
    "trigger literals": "",
    "trigger regex": "(laugh|lol)",
    "twitch event": "OnPrivateMessage"
  },
  "hello": {
    "audio filename literals": "",
    "audio filename regex": ".*[hH][eE][lL][lL][oO].*\\.(wav|mp3)",
    "audio files directory": "/home/gravestench/Downloads",
    "enabled": true,
    "spam count": 5,
    "spam interval": "30s",
    "trigger literals": "",
    "trigger regex": "(hello|hi|hey)",
    "twitch event": "OnPrivateMessage"
  },
  "wow": {
    "audio filename literals": "",
    "audio filename regex": ".*[wW][oO][wW].*\\.(wav|mp3)",
    "audio files directory": "/home/gravestench/Downloads",
    "enabled": true,
    "spam count": 10,
    "spam interval": "30s",
    "trigger literals": "wow",
    "trigger regex": "wow",
    "twitch event": "OnPrivateMessage"
  }
}
```

### Text To Speech
The TTS service can be used to configure how the text is converted into audible 
speech.
```
~/.config/chillcast/text_to_speech.json
```
```json
{
  "Text to speech": {
    "directory": "~/.config/chillcast/audio_files",
    "mplayer-handler": false
  }
}
```

### Twitch-integrated TTS
The twitch-integrated TTS service links the TTS and the twitch integration
together.
```
~/.config/chillcast/twitch_integrated_text_to_speech.json
```
```json
{
  "Twitch-Integrated TTS": {
    "_help :: possible voices (language codes)": "en,en-UK,en-AU,ja,de,es,ru,ar,cs,da,nl,fi,el,hi,hu,id,km,la,it,no,pl,sk,sv,th,tr,uk,vi,af,bg,ca,cy,et,fr,gu,is,jv,kn,ko,lv,ml,mr,ms,ne,pt,ro,si,sr,su,ta,te,tl,ur,zh,sw,sq,my,mk,hy,hr,eo,bs",
    "announce user": true,
    "enable": true,
    "remove profanity": false,
    "system voice": "en-UK",
    "user voice": "en-AU"
  }
}
```

### Twitch Integration
The twitch integration service is used by the other services. You need to get
your oauth token from https://twitchapps.com/tmi/

You can enable/disable specific events, too.
```
~/.config/chillcast/twitch_integration.json
```
```json
{
  "credentials": {
    "Oauth Key": "",
    "Username": ""
  },
  "handlers": {
    "OnClearChatMessage": true,
    "OnClearMessage": true,
    "OnConnect": true,
    "OnGlobalUserStateMessage": true,
    "OnNamesMessage": true,
    "OnNoticeMessage": true,
    "OnPingMessage": true,
    "OnPingSent": true,
    "OnPongMessage": true,
    "OnPrivateMessage": true,
    "OnReconnectMessage": true,
    "OnRoomStateMessage": true,
    "OnUnsetMessage": true,
    "OnUserJoinMessage": true,
    "OnUserNoticeMessage": true,
    "OnUserPartMessage": true,
    "OnUserStateMessage": true,
    "OnWhisperMessage": true
  }
}
```

## Contributing
We welcome contributions! Please see our CONTRIBUTING.md file for details on how 
you can contribute to the ChillCast project.

## License
ChillCast is released under the MIT License.
Credits

    Dylan Knuth
    

## Support

For support with chillcast, please open an issue on this GitHub repository or 
contact the maintainers directly.

## Roadmap

* [x] text to speech
* [x] profanity check 
* [ ] soundboard
* [ ] automated soundboard downloads
