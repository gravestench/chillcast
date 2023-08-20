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

## Usage & Configuration
The first time you run chillcast you should see some errors get printed. This 
is to point you to your config files. You need to set up your oath token for
your twitch account. The config files will be inside of `~/.config/chillcast`
and the twitch integration service should print the link you can go to
in order to get your oauth token.

After you get your oauth token, put it into your config file at:
```
~/.config/chillcast/twitch_integration.json
```

Explore the other config files for the profanity detecion and stuff.

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
